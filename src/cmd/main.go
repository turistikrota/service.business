package main

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/env"
	"github.com/mixarchitecture/microp/events/nats"
	"github.com/mixarchitecture/microp/logs"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.owner/src/delivery"
	"github.com/turistikrota/service.owner/src/service"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
)

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18np.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     config.Nats.Url,
		Streams: config.Nats.Streams,
	})
	valid := validator.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadOwnerMongo(config)
	app := service.NewApplication(service.Config{
		App:         config,
		EventEngine: eventEngine,
		Mongo:       mongo,
		Validator:   valid,
	})
	redis := redis.New(&redis.Config{
		Host:     config.Redis.Host,
		Port:     config.Redis.Port,
		Password: config.Redis.Pw,
		DB:       config.Redis.Db,
	})
	tknSrv := token.New(token.Config{
		Expiration: config.TokenSrv.Expiration,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       redis,
		EventEngine: eventEngine,
		Topic:       config.Session.Topic,
		TokenSrv:    tknSrv,
		Project:     config.TokenSrv.Project,
	})
	delivery := delivery.New(delivery.Config{
		App:         app,
		Config:      config,
		I18n:        i18n,
		Validator:   valid,
		Ctx:         ctx,
		EventEngine: eventEngine,
		SessionSrv:  session.Service,
		TokenSrv:    tknSrv,
	})
	delivery.Load()
}

func loadOwnerMongo(cnf config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  cnf.DB.MongoOwner.Host,
		Port:  cnf.DB.MongoOwner.Port,
		User:  cnf.DB.MongoOwner.Username,
		Pass:  cnf.DB.MongoOwner.Password,
		Db:    cnf.DB.MongoOwner.Database,
		Query: cnf.DB.MongoOwner.Query,
	})
	d, err := mongo.New(uri, cnf.DB.MongoOwner.Database)
	if err != nil {
		panic(err)
	}
	return d
}
