package http

import (
	"github.com/gofiber/fiber/v2"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.owner/src/delivery/http/dto"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h Server) OwnerApplication(ctx *fiber.Ctx) error {
	user := dto.Request.UserAccount()
	d := dto.Request.OwnerApplication()
	h.parseParams(ctx, user)
	d.LoadUser(user)
	h.parseBody(ctx, d)
	res, err := h.app.Commands.OwnerApplication.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.OwnerApplication, func() interface{} {
		return dto.Response.OwnerApplication(res)
	})
}

func (h Server) OwnershipUserAdd(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetailUser()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipUserAdd.Handle(ctx.UserContext(), d.ToAddUserCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.OwnershipUserAdd)
}

func (h Server) OwnershipUserRemove(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetailUser()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipUserRemove.Handle(ctx.UserContext(), d.ToRemoveUserCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.OwnershipUserRemove)
}

func (h Server) OwnershipUserPermAdd(ctx *fiber.Ctx) error {
	detail := dto.Request.OwnerShipDetailUser()
	d := dto.Request.OwnerPermissionAdd()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.OwnershipUserPermAdd.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.OwnershipUserPermAdd)
}

func (h Server) OwnershipUserPermRemove(ctx *fiber.Ctx) error {
	detail := dto.Request.OwnerShipDetailUser()
	d := dto.Request.OwnerPermissionRemove()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.OwnershipUserPermRemove.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.OwnershipUserPermRemove)
}

func (h Server) OwnershipAdminView(ctx *fiber.Ctx) error {
	ownership := h.parseOwner(ctx)
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	return result.SuccessDetail(h.i18n.Translate(Messages.Success.OwnershipAdminView, l, a), dto.Response.OwnershipAdminView(&ownership.Entity))
}

func (h Server) OwnershipUserList(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.ListMyOwnershipUsers.Handle(ctx.UserContext(), d.ToUserListQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.OwnershipUserList, func() interface{} {
		return dto.Response.ListMyOwnershipUsers(res)
	})
}

func (h Server) ListMyOwnerships(ctx *fiber.Ctx) error {
	d := dto.Request.UserAccount()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.ListMyOwnerships.Handle(ctx.UserContext(), d.ToListMyOwnershipsQuery(current_user.Parse(ctx).UUID))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.ListMyOwnerships, func() interface{} {
		return dto.Response.ListMyOwnerships(res)
	})
}

func (h Server) ViewOwnership(ctx *fiber.Ctx) error {
	d := dto.Request.Ownership()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.ViewOwnership.Handle(ctx.UserContext(), d.ToViewQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.ViewOwnership, func() interface{} {
		return dto.Response.ViewOwnership(res)
	})
}

func (h Server) OwnershipEnable(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipEnable.Handle(ctx.UserContext(), d.ToEnableCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.OwnershipEnable)
}

func (h Server) OwnershipDisable(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipDisable.Handle(ctx.UserContext(), d.ToDisableCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.OwnershipDisable)
}

func (h Server) OwnershipSelect(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerSelect()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.GetWithUserOwnership.Handle(ctx.UserContext(), d.ToGetQuery(current_user.Parse(ctx).UUID))
	if err != nil {
		return err
	}
	ctx.Cookie(h.CreateServerSideCookie(".s.o.n", res.Ownership.Entity.NickName))
	return result.Success(Messages.Success.OwnershipSelect)
}

func (h Server) OwnershipGetSelected(ctx *fiber.Ctx) error {
	nickName := ctx.Cookies(".s.o.n")
	if nickName == "" {
		return result.ErrorDetail(Messages.Error.OwnerNotSelected, dto.Response.OwnershipSelectNotFound())
	}
	d := dto.Request.UserAccount()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.GetWithUserOwnership.Handle(ctx.UserContext(), d.ToGetOwnershipQuery(nickName, current_user.Parse(ctx).UUID))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.OwnershipGetSelected, func() interface{} {
		return dto.Response.SelectOwnership(res)
	})
}
