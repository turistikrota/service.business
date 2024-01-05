package command

import (
	"context"
	"time"

	"github.com/9ssi7/vkn"

	"github.com/mixarchitecture/chain"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/ssibrahimbas/KPSPublic"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.shared/cipher"
)

type BusinessApplicationCommand struct {
	UserUUID          string
	UserName          string
	NickName          string
	RealName          string
	BusinessType      business.Type
	Individual        BusinessApplicationIndividualCommand
	Corporation       BusinessApplicationCorporationCommand
	hashedIndividual  business.Individual
	hashedCorporation business.Corporation
}

type BusinessApplicationIndividualCommand struct {
	FirstName      string
	LastName       string
	IdentityNumber string
	SerialNumber   string
	Province       string
	District       string
	Address        string
	DateOfBirth    time.Time
}
type BusinessApplicationCorporationCommand struct {
	TaxNumber string
	Province  string
	District  string
	Address   string
	Type      string
	TaxOffice string
	Title     string
}

type BusinessApplicationResult struct {
	BusinessUUID string
}

type BusinessApplicationHandler decorator.CommandHandler[BusinessApplicationCommand, *BusinessApplicationResult]

type businessApplicationHandler struct {
	repo            business.Repository
	factory         business.Factory
	events          business.Events
	identityService KPSPublic.Service
	vknService      vkn.Vkn
	cipher          cipher.Service
}

type BusinessApplicationHandlerConfig struct {
	Repo            business.Repository
	Factory         business.Factory
	Events          business.Events
	IdentityService KPSPublic.Service
	CqrsBase        decorator.Base
	Cipher          cipher.Service
	VknService      vkn.Vkn
}

func NewBusinessApplicationHandler(config BusinessApplicationHandlerConfig) BusinessApplicationHandler {
	return decorator.ApplyCommandDecorators[BusinessApplicationCommand, *BusinessApplicationResult](
		businessApplicationHandler{
			repo:            config.Repo,
			factory:         config.Factory,
			events:          config.Events,
			identityService: config.IdentityService,
			vknService:      config.VknService,
			cipher:          config.Cipher,
		},
		config.CqrsBase,
	)
}

type businessApplicationChain struct {
	command      BusinessApplicationCommand
	businessUUID string
	entity       *business.Entity
}

func (h businessApplicationHandler) Handle(ctx context.Context, command BusinessApplicationCommand) (*BusinessApplicationResult, *i18np.Error) {
	ch := chain.New[*businessApplicationChain, BusinessApplicationResult]()
	ch.Use(h.create, h.hash, h.validate, h.verifyByType, h.checkExists, h.checkNickName, h.save, h.end)
	return ch.Run(ctx, &businessApplicationChain{command: command})
}

func (h businessApplicationHandler) checkExists(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	if chain.command.BusinessType == business.Types.Individual {
		return h.checkIndividualExists(ctx, chain)
	}
	return h.checkCorporationExists(ctx, chain)
}

func (h businessApplicationHandler) checkIndividualExists(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	_, notFound, err := h.repo.GetByIndividual(ctx, chain.command.hashedIndividual)
	if err != nil {
		return nil, err
	}
	if !notFound {
		return nil, h.factory.Errors.IndividualAlreadyExists()
	}
	return nil, nil
}

func (h businessApplicationHandler) checkCorporationExists(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	_, notFound, err := h.repo.GetByCorporation(ctx, chain.command.hashedCorporation)
	if err != nil {
		return nil, err
	}
	if !notFound {
		return nil, h.factory.Errors.CorporationAlreadyExists()
	}
	return nil, nil
}

func (h businessApplicationHandler) create(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	chain.entity = h.factory.NewBusiness(business.NewBusinessParams{
		UserUUID:     chain.command.UserUUID,
		UserName:     chain.command.UserName,
		NickName:     chain.command.NickName,
		RealName:     chain.command.RealName,
		BusinessType: chain.command.BusinessType,
		Individual:   chain.command.hashedIndividual,
		Corporation:  chain.command.hashedCorporation,
	})
	return nil, nil
}

func (h businessApplicationHandler) validate(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	if err := h.factory.Validate(chain.entity); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h businessApplicationHandler) verifyByType(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	if chain.entity.BusinessType == business.Types.Individual {
		return h.validateIndividual(ctx, chain)
	}
	return h.validateCorporation(ctx, chain)
}

func (h businessApplicationHandler) checkNickName(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	res, err := h.repo.CheckNickName(ctx, chain.command.NickName)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	if res {
		return nil, h.factory.Errors.NickNameAlreadyExists()
	}
	return nil, nil
}

func (h businessApplicationHandler) hash(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	if chain.command.BusinessType == business.Types.Individual {
		return h.hashIndividual(ctx, chain)
	}
	return h.hashCorporation(ctx, chain)
}

func (h businessApplicationHandler) hashIndividual(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	id, error := h.cipher.Encrypt(chain.command.Individual.IdentityNumber)
	if error != nil {
		return nil, h.factory.Errors.Failed("failed to hash identity number")
	}
	serial, error := h.cipher.Encrypt(chain.command.Individual.SerialNumber)
	if error != nil {
		return nil, h.factory.Errors.Failed("failed to hash serial number")
	}
	chain.command.hashedIndividual = business.Individual{
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

func (h businessApplicationHandler) hashCorporation(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	tax, err := h.cipher.Encrypt(chain.command.Corporation.TaxNumber)
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to hash tax number")
	}
	chain.command.hashedCorporation = business.Corporation{
		TaxNumber: tax,
		Province:  chain.command.Corporation.Province,
		District:  chain.command.Corporation.District,
		Address:   chain.command.Corporation.Address,
		Type:      business.CorporationType(chain.command.Corporation.Type),
	}
	chain.entity.Corporation = chain.command.hashedCorporation
	return nil, nil
}

func (h businessApplicationHandler) validateIndividual(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
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

func (h businessApplicationHandler) validateCorporation(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
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

func (h businessApplicationHandler) save(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	res, err := h.repo.Create(ctx, chain.entity)
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to save business")
	}
	chain.businessUUID = res.UUID
	return nil, nil
}

func (h businessApplicationHandler) end(ctx context.Context, chain *businessApplicationChain) (*BusinessApplicationResult, *i18np.Error) {
	h.events.Created(&business.EventBusinessCreated{
		Business: chain.entity,
		UserUUID: chain.command.UserUUID,
		UserName: chain.command.UserName,
	})
	return &BusinessApplicationResult{
		BusinessUUID: chain.businessUUID,
	}, nil
}
