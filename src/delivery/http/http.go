package http

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/server/http"
	"github.com/mixarchitecture/microp/server/http/parser"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.owner/src/app"
)

type Server struct {
	app         app.Application
	i18n        *i18np.I18n
	validator   validator.Validator
	ctx         context.Context
	sessionSrv  session.Service
	tknSrv      token.Service
	httpHeaders config.HttpHeaders
}

type Config struct {
	App         app.Application
	I18n        *i18np.I18n
	Validator   validator.Validator
	Context     context.Context
	HttpHeaders config.HttpHeaders
	SessionSrv  session.Service
	TokenSrv    token.Service
}

func New(config Config) Server {
	return Server{
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		ctx:         config.Context,
		sessionSrv:  config.SessionSrv,
		tknSrv:      config.TokenSrv,
		httpHeaders: config.HttpHeaders,
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Use(h.cors(), h.deviceUUID())

	owner := router.Group("/~:nickName", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.CurrentOwner())
	owner.Get("/", h.OwnerPermissions(config.Roles.Owner.AdminView), h.wrapWithTimeout(h.OwnershipAdminView))
	owner.Get("/user", h.OwnerPermissions(config.Roles.Owner.UserList), h.wrapWithTimeout(h.OwnershipUserList))
	owner.Post("/user/@:userName/role", h.OwnerPermissions(config.Roles.Owner.UserPermAdd), h.wrapWithTimeout(h.OwnershipUserPermAdd))
	owner.Delete("/user/@:userName/role", h.OwnerPermissions(config.Roles.Owner.UserPermRemove), h.wrapWithTimeout(h.OwnershipUserPermRemove))
	owner.Delete("/user/@:userName", h.OwnerPermissions(config.Roles.Owner.UserRemove), h.wrapWithTimeout(h.OwnershipUserRemove))
	owner.Put("/enable", h.OwnerPermissions(config.Roles.Owner.Enable), h.wrapWithTimeout(h.OwnershipEnable))
	owner.Put("/disable", h.OwnerPermissions(config.Roles.Owner.Disable), h.wrapWithTimeout(h.OwnershipDisable))
	owner.Put("/select", h.wrapWithTimeout(h.OwnershipSelect))

	// invite owner routes
	owner.Post("/invite", h.OwnerPermissions(config.Roles.Owner.InviteCreate), h.wrapWithTimeout(h.InviteCreate))
	owner.Delete("/invite/:uuid", h.OwnerPermissions(config.Roles.Owner.InviteDelete), h.wrapWithTimeout(h.InviteDelete))
	owner.Get("/invite", h.OwnerPermissions(config.Roles.Owner.InviteView), h.wrapWithTimeout(h.InviteGetByOwnerUUID))

	admin := router.Group("/admin", h.currentUserAccess(), h.requiredAccess())
	admin.Get("/", h.adminRoute(config.Roles.Owner.AdminList), h.wrapWithTimeout(h.AdminListAll))
	admin.Patch("/:nickName/verify", h.adminRoute(config.Roles.Owner.AdminVerify), h.wrapWithTimeout(h.AdminOwnershipVerify))
	admin.Patch("/:nickName/reject", h.adminRoute(config.Roles.Owner.AdminReject), h.wrapWithTimeout(h.AdminOwnershipReject))
	admin.Delete("/:nickName", h.adminRoute(config.Roles.Owner.AdminDelete), h.wrapWithTimeout(h.AdminOwnershipDelete))
	admin.Patch("/:nickName/recover", h.adminRoute(config.Roles.Owner.AdminRecover), h.wrapWithTimeout(h.AdminOwnershipRecover))
	admin.Get("/:nickName", h.adminRoute(config.Roles.Owner.AdminView), h.wrapWithTimeout(h.AdminView))
	admin.Get("/invites/:uuid", h.adminRoute(config.Roles.Owner.InviteView), h.wrapWithTimeout(h.InviteGetByUUID))

	router.Get("/selected", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.OwnershipGetSelected))
	router.Post("/", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.OwnerApplication))
	router.Get("/", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.ListMyOwnerships))

	// invite public routes
	router.Post("/join/:uuid", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteUse))
	router.Get("/join/:uuid", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteGetByUUID))
	router.Get("/invites", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteGetByEmail))

	router.Get("/:nickName", h.wrapWithTimeout(h.ViewOwnership))
	return router
}

func (h Server) parseBody(c *fiber.Ctx, d interface{}) {
	parser.ParseBody(c, h.validator, *h.i18n, d)
}

func (h Server) parseParams(c *fiber.Ctx, d interface{}) {
	parser.ParseParams(c, h.validator, *h.i18n, d)
}

func (h Server) parseQuery(c *fiber.Ctx, d interface{}) {
	parser.ParseQuery(c, h.validator, *h.i18n, d)
}

func (h Server) currentAccountAccess() fiber.Handler {
	return current_account.New(current_account.Config{
		I18n:         h.i18n,
		RequiredKey:  Messages.Error.RequiredAccountSelect,
		ForbiddenKey: Messages.Error.AccountNotFound,
	})
}

func (h Server) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}

func (h Server) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   *h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h Server) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  http.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h Server) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h Server) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   *h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h Server) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     h.httpHeaders.AllowedOrigins,
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(h.httpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}

func (h Server) CreateServerSideCookie(key string, value string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Domain:   h.httpHeaders.Domain,
	}
}
