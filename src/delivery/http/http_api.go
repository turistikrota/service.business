package http

import (
	"github.com/gofiber/fiber/v2"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.owner/src/delivery/http/dto"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h Server) OwnerApplication(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerApplication()
	h.parseBody(ctx, d)
	account := current_account.Parse(ctx)
	res, err := h.app.Commands.OwnerApplication.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.OwnerApplication, func() interface{} {
		return dto.Response.OwnerApplication(res)
	})
}

func (h Server) OwnershipUserRemove(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetailUser()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipUserRemove.Handle(ctx.UserContext(), d.ToRemoveUserCommand(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.OwnershipUserRemove)
}

func (h Server) AdminOwnershipVerify(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipVerifyByAdmin.Handle(ctx.UserContext(), d.ToVerifyCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) AdminOwnershipDelete(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipDeleteByAdmin.Handle(ctx.UserContext(), d.ToDeleteCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) AdminOwnershipRecover(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.OwnershipRecoverByAdmin.Handle(ctx.UserContext(), d.ToRecoverCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) AdminOwnershipReject(ctx *fiber.Ctx) error {
	detail := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, detail)
	d := dto.Request.OwnershipReject()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.OwnershipRejectByAdmin.Handle(ctx.UserContext(), d.ToCommand(detail.NickName, current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) OwnershipUserPermAdd(ctx *fiber.Ctx) error {
	detail := dto.Request.OwnerShipDetailUser()
	d := dto.Request.OwnerPermissionAdd()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.OwnershipUserPermAdd.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.OwnershipUserPermAdd)
}

func (h Server) OwnershipUserPermRemove(ctx *fiber.Ctx) error {
	detail := dto.Request.OwnerShipDetailUser()
	d := dto.Request.OwnerPermissionRemove()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.OwnershipUserPermRemove.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.OwnershipUserPermRemove)
}

func (h Server) OwnershipAdminView(ctx *fiber.Ctx) error {
	ownership := h.parseOwner(ctx)
	l, a := httpI18n.GetLanguagesInContext(*h.i18n, ctx)
	return result.SuccessDetail(h.i18n.Translate(Messages.Success.OwnershipAdminView, l, a), dto.Response.OwnershipAdminView(&ownership.Entity))
}

func (h Server) OwnershipUserList(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.ListMyOwnershipUsers.Handle(ctx.UserContext(), d.ToUserListQuery(account.Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.OwnershipUserList, func() interface{} {
		return dto.Response.ListMyOwnershipUsers(res)
	})
}

func (h Server) ListMyOwnerships(ctx *fiber.Ctx) error {
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.ListMyOwnerships.Handle(ctx.UserContext(), query.ListMyOwnershipsQuery{
		UserName: account.Name,
		UserUUID: current_user.Parse(ctx).UUID,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.ListMyOwnerships, func() interface{} {
		return dto.Response.ListMyOwnerships(res)
	})
}

func (h Server) AdminListAll(ctx *fiber.Ctx) error {
	d := dto.Request.Pagination()
	h.parseQuery(ctx, d)
	d.Default()
	res, err := h.app.Queries.AdminListAll.Handle(ctx.UserContext(), query.AdminListOwnershipQuery{
		Offset: (*d.Page - 1) * *d.Limit,
		Limit:  *d.Limit,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.AdminListAll, func() interface{} {
		return dto.Response.AdminListAll(res)
	})
}

func (h Server) AdminView(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.AdminViewOwnership.Handle(ctx.UserContext(), d.ToAdminViewQuery())
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return dto.Response.AdminView(res)
	})
}

func (h Server) ViewOwnership(ctx *fiber.Ctx) error {
	d := dto.Request.Ownership()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.ViewOwnership.Handle(ctx.UserContext(), d.ToViewQuery())
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.ViewOwnership, func() interface{} {
		return dto.Response.ViewOwnership(res)
	})
}

func (h Server) OwnershipEnable(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	_, err := h.app.Commands.OwnershipEnable.Handle(ctx.UserContext(), d.ToEnableCommand(current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.OwnershipEnable)
}

func (h Server) OwnershipDisable(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerShipDetail()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	_, err := h.app.Commands.OwnershipDisable.Handle(ctx.UserContext(), d.ToDisableCommand(current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.OwnershipDisable)
}

func (h Server) OwnershipSelect(ctx *fiber.Ctx) error {
	d := dto.Request.OwnerSelect()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.GetWithUserOwnership.Handle(ctx.UserContext(), d.ToGetQuery(current_user.Parse(ctx).UUID, account.Name))
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
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.GetWithUserOwnership.Handle(ctx.UserContext(), query.GetWithUserOwnershipQuery{
		NickName: nickName,
		UserName: account.Name,
		UserUUID: current_user.Parse(ctx).UUID,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.OwnershipGetSelected, func() interface{} {
		return dto.Response.SelectOwnership(res)
	})
}

func (h Server) InviteCreate(ctx *fiber.Ctx) error {
	d := dto.Request.InviteCreate()
	h.parseBody(ctx, d)
	account := current_account.Parse(ctx)
	ownership := h.parseOwner(ctx)
	res, err := h.app.Commands.InviteCreate.Handle(ctx.UserContext(), d.ToCommand(ownership.Entity.NickName, ownership.Entity.UUID, current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res
	})
}

func (h Server) InviteDelete(ctx *fiber.Ctx) error {
	d := dto.Request.InviteDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Commands.InviteDelete.Handle(ctx.UserContext(), d.ToDelete(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res
	})
}

func (h Server) InviteUse(ctx *fiber.Ctx) error {
	d := dto.Request.InviteDetail()
	h.parseParams(ctx, d)
	u := current_user.Parse(ctx)
	account := current_account.Parse(ctx)
	res, err := h.app.Commands.InviteUse.Handle(ctx.UserContext(), d.ToUse(u.UUID, u.Email, account.Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res
	})
}

func (h Server) InviteGetByUUID(ctx *fiber.Ctx) error {
	d := dto.Request.InviteDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.InviteGetByUUID.Handle(ctx.UserContext(), d.ToGet())
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res.Invite
	})
}

func (h Server) InviteGetByEmail(ctx *fiber.Ctx) error {
	u := current_user.Parse(ctx)
	res, err := h.app.Queries.InviteGetByEmail.Handle(ctx.UserContext(), query.InviteGetByEmailQuery{
		UserEmail: u.Email,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res.Invites
	})
}

func (h Server) InviteGetByOwnerUUID(ctx *fiber.Ctx) error {
	ownership := h.parseOwner(ctx)
	res, err := h.app.Queries.InviteGetByOwnerUUID.Handle(ctx.UserContext(), query.InviteGetByOwnerUUIDQuery{
		OwnerUUID: ownership.Entity.UUID,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res.Invites
	})
}
