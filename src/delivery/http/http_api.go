package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/delivery/http/dto"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h Server) BusinessApplication(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessApplication()
	h.parseBody(ctx, d)
	account := current_account.Parse(ctx)
	res, err := h.app.Commands.BusinessApplication.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.BusinessApplication, func() interface{} {
		return dto.Response.BusinessApplication(res)
	})
}

func (h Server) BusinessUserRemove(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetailUser()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.BusinessUserRemove.Handle(ctx.UserContext(), d.ToRemoveUserCommand(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.BusinessUserRemove)
}

func (h Server) AdminBusinessVerify(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.BusinessVerifyByAdmin.Handle(ctx.UserContext(), d.ToVerifyCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) AdminBusinessDelete(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.BusinessDeleteByAdmin.Handle(ctx.UserContext(), d.ToDeleteCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) AdminBusinessRecover(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.BusinessRecoverByAdmin.Handle(ctx.UserContext(), d.ToRecoverCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) AdminBusinessReject(ctx *fiber.Ctx) error {
	detail := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, detail)
	d := dto.Request.BusinessReject()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.BusinessRejectByAdmin.Handle(ctx.UserContext(), d.ToCommand(detail.NickName, current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.Ok)
}

func (h Server) BusinessUserPermAdd(ctx *fiber.Ctx) error {
	detail := dto.Request.BusinessShipDetailUser()
	d := dto.Request.BusinessPermissionAdd()
	h.parseParams(ctx, detail)
	h.parseBody(ctx, d)
	d.LoadDetail(detail)
	_, err := h.app.Commands.BusinessUserPermAdd.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.BusinessUserPermAdd)
}

func (h Server) BusinessUserPermRemove(ctx *fiber.Ctx) error {
	detail := dto.Request.BusinessShipDetailUser()
	d := dto.Request.BusinessPermissionRemove()
	h.parseParams(ctx, detail)
	h.parseBody(ctx, d)
	d.LoadDetail(detail)
	_, err := h.app.Commands.BusinessUserPermRemove.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID, current_account.Parse(ctx).Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.BusinessUserPermRemove)
}

func (h Server) BusinessAdminView(ctx *fiber.Ctx) error {
	business := h.parseBusiness(ctx)
	res, err := h.app.Queries.AdminViewBusiness.Handle(ctx.UserContext(), query.AdminViewBusinessQuery{
		NickName: business.Entity.NickName,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.BusinessAdminView, func() interface{} {
		return dto.Response.BusinessAdminView(res.Business)
	})
}

func (h Server) BusinessUserList(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.ListMyBusinessUsers.Handle(ctx.UserContext(), d.ToUserListQuery(account.Name))
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.BusinessUserList, func() interface{} {
		return dto.Response.ListMyBusinessUsers(res)
	})
}

func (h Server) ListMyBusinesses(ctx *fiber.Ctx) error {
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.ListMyBusinesses.Handle(ctx.UserContext(), query.ListMyBusinessesQuery{
		UserName: account.Name,
		UserUUID: current_user.Parse(ctx).UUID,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.ListMyBusinesses, func() interface{} {
		return dto.Response.ListMyBusinesses(res)
	})
}

func (h Server) AdminListAll(ctx *fiber.Ctx) error {
	d := dto.Request.Pagination()
	h.parseQuery(ctx, d)
	d.Default()
	res, err := h.app.Queries.AdminListAll.Handle(ctx.UserContext(), query.AdminListBusinessQuery{
		Offset: (*d.Page - 1) * *d.Limit,
		Limit:  *d.Limit,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.AdminListAll, func() interface{} {
		return dto.Response.AdminListAll(res)
	})
}

func (h Server) AdminView(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.AdminViewBusiness.Handle(ctx.UserContext(), d.ToAdminViewQuery())
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return dto.Response.AdminView(res)
	})
}

func (h Server) ViewBusiness(ctx *fiber.Ctx) error {
	d := dto.Request.Business()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.ViewBusiness.Handle(ctx.UserContext(), d.ToViewQuery())
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.ViewBusiness, func() interface{} {
		return dto.Response.ViewBusiness(res)
	})
}

func (h Server) BusinessEnable(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	_, err := h.app.Commands.BusinessEnable.Handle(ctx.UserContext(), d.ToEnableCommand(current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.BusinessEnable)
}

func (h Server) BusinessDisable(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessShipDetail()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	_, err := h.app.Commands.BusinessDisable.Handle(ctx.UserContext(), d.ToDisableCommand(current_user.Parse(ctx).UUID, account.Name))
	return result.IfSuccess(err, ctx, *h.i18n, Messages.Success.BusinessDisable)
}

func (h Server) BusinessSelect(ctx *fiber.Ctx) error {
	d := dto.Request.BusinessSelect()
	h.parseParams(ctx, d)
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.GetWithUserBusiness.Handle(ctx.UserContext(), d.ToGetQuery(current_user.Parse(ctx).UUID, account.Name))
	if err != nil {
		return err
	}
	if res.Business == nil {
		return result.ErrorDetail(Messages.Error.BusinessNotFound, dto.Response.BusinessSelectNotFound())
	}
	ctx.Cookie(h.CreateServerSideCookie(".s.o.n", res.Business.Entity.NickName))
	return result.Success(Messages.Success.BusinessSelect)
}

func (h Server) BusinessGetSelected(ctx *fiber.Ctx) error {
	nickName := ctx.Cookies(".s.o.n")
	if nickName == "" {
		return result.ErrorDetail(Messages.Error.BusinessNotSelected, dto.Response.BusinessSelectNotFound())
	}
	account := current_account.Parse(ctx)
	res, err := h.app.Queries.GetWithUserBusiness.Handle(ctx.UserContext(), query.GetWithUserBusinessQuery{
		NickName: nickName,
		UserName: account.Name,
		UserUUID: current_user.Parse(ctx).UUID,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.BusinessGetSelected, func() interface{} {
		return dto.Response.SelectBusiness(res)
	})
}

func (h Server) InviteCreate(ctx *fiber.Ctx) error {
	d := dto.Request.InviteCreate()
	h.parseBody(ctx, d)
	account := current_account.Parse(ctx)
	business := h.parseBusiness(ctx)
	res, err := h.app.Commands.InviteCreate.Handle(ctx.UserContext(), d.ToCommand(business.Entity.NickName, business.Entity.UUID, current_user.Parse(ctx).UUID, account.Name))
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

func (h Server) InviteGetByBusinessUUID(ctx *fiber.Ctx) error {
	business := h.parseBusiness(ctx)
	res, err := h.app.Queries.InviteGetByBusinessUUID.Handle(ctx.UserContext(), query.InviteGetByBusinessUUIDQuery{
		BusinessUUID: business.Entity.UUID,
	})
	return result.IfSuccessDetail(err, ctx, *h.i18n, Messages.Success.Ok, func() interface{} {
		return res.Invites
	})
}
