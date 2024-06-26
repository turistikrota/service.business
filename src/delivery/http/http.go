package http

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/server/http"
	"github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/parser"
	"github.com/turistikrota/service.business/src/config"
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
	"github.com/turistikrota/service.business/src/app"
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
	router.Use(i18n.New(*h.i18n), h.cors(), h.deviceUUID())

	business := router.Group("/~:nickName", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.CurrentBusiness())
	business.Get("/", h.BusinessPermissions(config.Roles.Business.AdminView), h.wrapWithTimeout(h.BusinessAdminView))
	business.Get("/user", h.BusinessPermissions(config.Roles.Business.UserList), h.wrapWithTimeout(h.BusinessUserList))
	business.Patch("/user/@:userName/add-role", h.BusinessPermissions(config.Roles.Business.UserPermAdd), h.wrapWithTimeout(h.BusinessUserPermAdd))
	business.Patch("/user/@:userName/rm-role", h.BusinessPermissions(config.Roles.Business.UserPermRemove), h.wrapWithTimeout(h.BusinessUserPermRemove))
	business.Patch("/user/@:userName", h.BusinessPermissions(config.Roles.Business.UserRemove), h.wrapWithTimeout(h.BusinessUserRemove))
	business.Patch("/enable", h.BusinessPermissions(config.Roles.Business.Enable), h.wrapWithTimeout(h.BusinessEnable))
	business.Patch("/disable", h.BusinessPermissions(config.Roles.Business.Disable), h.wrapWithTimeout(h.BusinessDisable))
	business.Patch("/locale", h.BusinessPermissions(config.Roles.Business.LocaleSet), h.wrapWithTimeout(h.BusinessSetLocale))
	business.Put("/select", h.wrapWithTimeout(h.BusinessSelect))

	// invite business routes
	business.Post("/invite", h.BusinessPermissions(config.Roles.Business.InviteCreate), h.wrapWithTimeout(h.InviteCreate))
	business.Delete("/invite/:uuid", h.BusinessPermissions(config.Roles.Business.InviteDelete), h.wrapWithTimeout(h.InviteDelete))
	business.Get("/invite", h.BusinessPermissions(config.Roles.Business.InviteView), h.wrapWithTimeout(h.InviteGetByBusinessUUID))

	admin := router.Group("/admin", h.currentUserAccess(), h.requiredAccess())
	admin.Get("/", h.adminRoute(config.Roles.Business.AdminList), h.wrapWithTimeout(h.AdminListAll))
	admin.Patch("/:nickName/verify", h.adminRoute(config.Roles.Business.AdminVerify), h.wrapWithTimeout(h.AdminBusinessVerify))
	admin.Patch("/:nickName/reject", h.adminRoute(config.Roles.Business.AdminReject), h.wrapWithTimeout(h.AdminBusinessReject))
	admin.Delete("/:nickName", h.adminRoute(config.Roles.Business.AdminDelete), h.wrapWithTimeout(h.AdminBusinessDelete))
	admin.Patch("/:nickName/recover", h.adminRoute(config.Roles.Business.AdminRecover), h.wrapWithTimeout(h.AdminBusinessRecover))
	admin.Get("/:nickName", h.adminRoute(config.Roles.Business.AdminView), h.wrapWithTimeout(h.AdminView))
	admin.Get("/invites/:uuid", h.adminRoute(config.Roles.Business.InviteView), h.wrapWithTimeout(h.InviteGetByUUID))

	router.Get("/selected", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.BusinessGetSelected))
	router.Post("/", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.BusinessApplication))
	router.Get("/", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.ListMyBusinesses))

	// invite public routes
	router.Post("/join/:uuid", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteUse))
	router.Get("/join/:uuid", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteGetByUUID))
	router.Get("/invites", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteGetByEmail))

	router.Get("/:nickName", h.wrapWithTimeout(h.ViewBusiness))
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
	return current_account.New(current_account.Config{})
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
	})
}

func (h Server) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
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
		MsgKey: Messages.Error.RequiredAuth,
		I18n:   *h.i18n,
	})
}

func (h Server) cors() fiber.Handler {
	return cors.New(cors.Config{
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
