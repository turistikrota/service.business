package http

import (
	"strings"
	"time"

	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/turistikrota/service.business/app"
	"github.com/turistikrota/service.business/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	httpServer "github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_business"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
)

type srv struct {
	config      config.App
	app         app.Application
	i18n        *i18np.I18n
	validator   validation.Validator
	tknSrv      token.Service
	sessionSrv  session.Service
	httpHeaders config.HttpHeaders
}

type Config struct {
	Env         config.App
	App         app.Application
	I18n        *i18np.I18n
	Validator   validation.Validator
	HttpHeaders config.HttpHeaders
	TokenSrv    token.Service
	SessionSrv  session.Service
}

func New(config Config) server.Server {
	return srv{
		config:      config.Env,
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		tknSrv:      config.TokenSrv,
		sessionSrv:  config.SessionSrv,
		httpHeaders: config.HttpHeaders,
	}
}

func (h srv) Listen() error {
	return http.RunServer(http.Config{
		Host:        h.config.Http.Host,
		Port:        h.config.Http.Port,
		I18n:        h.i18n,
		AcceptLangs: h.config.I18n.Locales,
		Debug:       true,
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Use(h.cors(), h.deviceUUID())
			business := router.Group("/business", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess())
			business.Get("/", h.currentBusinessAccess(config.Roles.Business.AdminView), h.wrapWithTimeout(h.ViewMyBusiness))
			business.Get("/user", h.currentBusinessAccess(config.Roles.Business.UserList), h.wrapWithTimeout(h.ListMyBusinessUsers))
			business.Patch("/user/@:userName/add-role", h.currentBusinessAccess(config.Roles.Business.UserPermAdd), h.wrapWithTimeout(h.BusinessUserPermAdd))
			business.Patch("/user/@:userName/rm-role", h.currentBusinessAccess(config.Roles.Business.UserPermRemove), h.wrapWithTimeout(h.BusinessUserPermRemove))
			business.Patch("/user/@:userName", h.currentBusinessAccess(config.Roles.Business.UserRemove), h.wrapWithTimeout(h.BusinessUserRemove))
			business.Patch("/enable", h.currentBusinessAccess(config.Roles.Business.Enable), h.wrapWithTimeout(h.BusinessEnable))
			business.Patch("/disable", h.currentBusinessAccess(config.Roles.Business.Disable), h.wrapWithTimeout(h.BusinessDisable))
			business.Patch("/locale", h.currentBusinessAccess(config.Roles.Business.LocaleSet), h.wrapWithTimeout(h.BusinessSetLocale))

			// invite business routes
			business.Post("/invite", h.currentBusinessAccess(config.Roles.Business.InviteCreate), h.wrapWithTimeout(h.InviteCreate))
			business.Delete("/invite/:uuid", h.currentBusinessAccess(config.Roles.Business.InviteDelete), h.wrapWithTimeout(h.InviteDelete))
			business.Get("/invite", h.currentBusinessAccess(config.Roles.Business.InviteView), h.wrapWithTimeout(h.InviteListByBusinessUUID))

			admin := router.Group("/admin", h.currentUserAccess(), h.requiredAccess())
			admin.Get("/", h.adminRoute(config.Roles.Business.AdminList), h.wrapWithTimeout(h.AdminListBusinesses))
			admin.Patch("/:nickName/verify", h.adminRoute(config.Roles.Business.AdminVerify), h.wrapWithTimeout(h.AdminBusinessVerify))
			admin.Patch("/:nickName/reject", h.adminRoute(config.Roles.Business.AdminReject), h.wrapWithTimeout(h.AdminBusinessReject))
			admin.Delete("/:nickName", h.adminRoute(config.Roles.Business.AdminDelete), h.wrapWithTimeout(h.AdminBusinessDelete))
			admin.Patch("/:nickName/recover", h.adminRoute(config.Roles.Business.AdminRecover), h.wrapWithTimeout(h.AdminBusinessRecover))
			admin.Get("/:nickName", h.adminRoute(config.Roles.Business.AdminView), h.wrapWithTimeout(h.AdminViewBusiness))
			admin.Get("/invites/:uuid", h.adminRoute(config.Roles.Business.InviteView), h.wrapWithTimeout(h.InviteGetByUUID))

			router.Get("/selected", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.BusinessGetSelected))
			router.Post("/", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.BusinessApplication))
			router.Get("/", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.ListMyBusinesses))

			// invite 3rd user routes
			router.Post("/join/:uuid", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteUse))
			router.Get("/join/:uuid", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteGetByUUID))
			router.Get("/invites", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.InviteListByEmail))

			router.Get("/:nickName", h.wrapWithTimeout(h.ViewBusiness))
			router.Put("/~:nickName/select", h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess(), h.wrapWithTimeout(h.BusinessSelect))
			return router
		},
	})
}

func (h srv) currentBusinessAccess(roles ...string) fiber.Handler {
	roles = append(roles, config.Roles.Business.Super)
	return current_business.New(current_business.Config{
		Roles: roles,
	})
}

func (h srv) currentAccountAccess() fiber.Handler {
	return current_account.New(current_account.Config{})
}

func (h srv) parseBody(c *fiber.Ctx, d interface{}) {
	http.ParseBody(c, h.validator, *h.i18n, d)
}

func (h srv) parseParams(c *fiber.Ctx, d interface{}) {
	http.ParseParams(c, h.validator, *h.i18n, d)
}

func (h srv) parseQuery(c *fiber.Ctx, d interface{}) {
	http.ParseQuery(c, h.validator, *h.i18n, d)
}

func (h srv) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  httpServer.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h srv) rateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
	})
}

func (h srv) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h srv) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   *h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h srv) adminRoute(extra ...string) fiber.Handler {
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

func (h srv) cors() fiber.Handler {
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

func (h srv) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}

func (h srv) CreateServerSideCookie(key string, value string) *fiber.Cookie {
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
