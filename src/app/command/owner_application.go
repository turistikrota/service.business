package command

import (
	"context"
	"time"

	"github.com/9ssi7/vkn"
	"github.com/turistikrota/service.owner/src/domain/account"

	"github.com/mixarchitecture/chain"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/ssibrahimbas/KPSPublic"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/cipher"
)

type OwnerApplicationCommand struct {
	UserUUID          string
	UserName          string
	NickName          string
	RealName          string
	OwnerType         owner.Type
	Individual        OwnerApplicationIndividualCommand
	Corporation       OwnerApplicationCorporationCommand
	hashedIndividual  owner.Individual
	hashedCorporation owner.Corporation
}

type OwnerApplicationIndividualCommand struct {
	FirstName      string
	LastName       string
	IdentityNumber string
	SerialNumber   string
	Province       string
	District       string
	Address        string
	DateOfBirth    time.Time
}
type OwnerApplicationCorporationCommand struct {
	TaxNumber string
	Province  string
	District  string
	Address   string
	Type      string
	TaxOffice string
	Title     string
}

type OwnerApplicationResult struct {
	OwnerUUID string
}

type OwnerApplicationHandler decorator.CommandHandler[OwnerApplicationCommand, *OwnerApplicationResult]

type ownerApplicationHandler struct {
	repo            owner.Repository
	factory         owner.Factory
	events          owner.Events
	accountRepo     account.Repository
	identityService KPSPublic.Service
	vknService      vkn.Vkn
}

type OwnerApplicationHandlerConfig struct {
	Repo            owner.Repository
	Factory         owner.Factory
	Events          owner.Events
	AccountRepo     account.Repository
	IdentityService KPSPublic.Service
	CqrsBase        decorator.Base
}

func NewOwnerApplicationHandler(config OwnerApplicationHandlerConfig) OwnerApplicationHandler {
	return decorator.ApplyCommandDecorators[OwnerApplicationCommand, *OwnerApplicationResult](
		ownerApplicationHandler{
			repo:            config.Repo,
			factory:         config.Factory,
			events:          config.Events,
			accountRepo:     config.AccountRepo,
			identityService: config.IdentityService,
		},
		config.CqrsBase,
	)
}

type ownerApplicationChain struct {
	command   OwnerApplicationCommand
	ownerUUID string
	entity    *owner.Entity
}

func (h ownerApplicationHandler) Handle(ctx context.Context, command OwnerApplicationCommand) (*OwnerApplicationResult, *i18np.Error) {
	ch := chain.New[*ownerApplicationChain, OwnerApplicationResult]()
	ch.Use(h.checkAccount, h.create, h.hash, h.validate, h.verifyByType, h.checkExists, h.checkNickName, h.save, h.end)
	return ch.Run(ctx, &ownerApplicationChain{command: command})
}

func (h ownerApplicationHandler) checkAccount(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	_, err := h.accountRepo.GetByUserUUID(ctx, account.UserUnique{
		UserUUID: chain.command.UserUUID,
		Name:     chain.command.UserName,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h ownerApplicationHandler) checkExists(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	if chain.command.OwnerType == owner.Types.Individual {
		return h.checkIndividualExists(ctx, chain)
	}
	return h.checkCorporationExists(ctx, chain)
}

func (h ownerApplicationHandler) checkIndividualExists(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	_, notFound, err := h.repo.GetByIndividual(ctx, chain.command.hashedIndividual)
	if err != nil {
		return nil, err
	}
	if !notFound {
		return nil, h.factory.Errors.IndividualAlreadyExists()
	}
	return nil, nil
}

func (h ownerApplicationHandler) checkCorporationExists(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	_, notFound, err := h.repo.GetByCorporation(ctx, chain.command.hashedCorporation)
	if err != nil {
		return nil, err
	}
	if !notFound {
		return nil, h.factory.Errors.CorporationAlreadyExists()
	}
	return nil, nil
}

func (h ownerApplicationHandler) create(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	chain.entity = h.factory.NewOwner(owner.NewOwnerParams{
		UserUUID:    chain.command.UserUUID,
		UserName:    chain.command.UserName,
		NickName:    chain.command.NickName,
		RealName:    chain.command.RealName,
		OwnerType:   chain.command.OwnerType,
		Individual:  chain.command.hashedIndividual,
		Corporation: chain.command.hashedCorporation,
	})
	return nil, nil
}

func (h ownerApplicationHandler) validate(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	if err := h.factory.Validate(chain.entity); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h ownerApplicationHandler) verifyByType(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	if chain.entity.OwnerType == owner.Types.Individual {
		return h.validateIndividual(ctx, chain)
	}
	return h.validateCorporation(ctx, chain)
}

func (h ownerApplicationHandler) checkNickName(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	res, err := h.repo.CheckNickName(ctx, chain.command.NickName)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	if res {
		return nil, h.factory.Errors.NickNameAlreadyExists()
	}
	return nil, nil
}

func (h ownerApplicationHandler) hash(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	if chain.command.OwnerType == owner.Types.Individual {
		return h.hashIndividual(ctx, chain)
	}
	return h.hashCorporation(ctx, chain)
}

func (h ownerApplicationHandler) hashIndividual(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	id, error := cipher.Hash(chain.command.Individual.IdentityNumber)
	if error != nil {
		return nil, h.factory.Errors.Failed("failed to hash identity number")
	}
	serial, error := cipher.Hash(chain.command.Individual.SerialNumber)
	if error != nil {
		return nil, h.factory.Errors.Failed("failed to hash serial number")
	}
	chain.command.hashedIndividual = owner.Individual{
		FirstName:      chain.command.Individual.FirstName,
		LastName:       chain.command.Individual.LastName,
		IdentityNumber: id,
		SerialNumber:   serial,
		Province:       chain.command.Individual.Province,
		District:       chain.command.Individual.District,
		Address:        chain.command.Individual.Address,
		DateOfBirth:    chain.command.Individual.DateOfBirth,
	}
	chain.entity.Individual = chain.command.hashedIndividual
	return nil, nil
}

func (h ownerApplicationHandler) hashCorporation(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	tax, err := cipher.Hash(chain.command.Corporation.TaxNumber)
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to hash tax number")
	}
	chain.command.hashedCorporation = owner.Corporation{
		TaxNumber: tax,
		Province:  chain.command.Corporation.Province,
		District:  chain.command.Corporation.District,
		Address:   chain.command.Corporation.Address,
		Type:      owner.CorporationType(chain.command.Corporation.Type),
	}
	chain.entity.Corporation = chain.command.hashedCorporation
	return nil, nil
}

func (h ownerApplicationHandler) validateIndividual(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	res, err := h.identityService.Verify(KPSPublic.VerifyConfig{
		SerialNumber:     chain.command.Individual.SerialNumber,
		TCIdentityNumber: chain.command.Individual.IdentityNumber,
		LastName:         chain.entity.Individual.LastName,
		FirstName:        chain.entity.Individual.FirstName,
		BirthYear:        chain.entity.Individual.DateOfBirth.Year(),
		BirthMonth:       int(chain.entity.Individual.DateOfBirth.Month()),
		BirthDay:         chain.entity.Individual.DateOfBirth.Day(),
	})
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	if !res {
		return nil, h.factory.Errors.IdentityVerificationFailed()
	}
	return nil, nil
}

func (h ownerApplicationHandler) validateCorporation(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	res, err := h.vknService.GetRecipient(chain.command.Corporation.TaxNumber)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	if res == nil || res.Data.TaxOffice == "" || res.Data.Title == "" {
		return nil, h.factory.Errors.CorporationVerificationFailed()
	}
	chain.command.hashedCorporation.Title = res.Data.Title
	chain.command.hashedCorporation.TaxOffice = res.Data.TaxOffice
	chain.entity.Corporation.TaxOffice = res.Data.TaxOffice
	chain.entity.Corporation.Title = res.Data.Title
	return nil, nil
}

func (h ownerApplicationHandler) save(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	res, err := h.repo.Create(ctx, chain.entity)
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to save owner")
	}
	chain.ownerUUID = res.UUID
	return nil, nil
}

func (h ownerApplicationHandler) end(ctx context.Context, chain *ownerApplicationChain) (*OwnerApplicationResult, *i18np.Error) {
	h.events.Created(&owner.EventOwnerCreated{
		Owner: chain.entity,
	})
	return &OwnerApplicationResult{
		OwnerUUID: chain.ownerUUID,
	}, nil
}
