package main

import (
	"github.com/cilloparch/cillop/env"
	"github.com/cilloparch/cillop/events/nats"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.business/config"
	"github.com/turistikrota/service.business/server/http"
	"github.com/turistikrota/service.business/server/rpc"
	"github.com/turistikrota/service.business/service"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
)

func main() {
	config := config.App{}
	env.Load(&config)
	i18n := i18np.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     config.Nats.Url,
		Streams: config.Nats.Streams,
	})
	valid := validation.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadBusinessMongo(config)
	app := service.NewApplication(service.Config{
		App:         config,
		EventEngine: eventEngine,
		Mongo:       mongo,
		Validator:   valid,
		I18n:        i18n,
	})
	redis := redis.New(&redis.Config{
		Host:     config.Redis.Host,
		Port:     config.Redis.Port,
		Password: config.Redis.Pw,
		DB:       config.Redis.Db,
	})
	tknSrv := token.New(token.Config{
		Expiration:     config.TokenSrv.Expiration,
		PublicKeyFile:  config.Rsa.PublicKeyFile,
		PrivateKeyFile: config.Rsa.PrivateKeyFile,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       redis,
		EventEngine: eventEngine,
		Topic:       config.Session.Topic,
		TokenSrv:    tknSrv,
		Project:     config.TokenSrv.Project,
	})
	http := http.New(http.Config{
		Env:         config,
		App:         app,
		I18n:        i18n,
		Validator:   *valid,
		HttpHeaders: config.HttpHeaders,
		TokenSrv:    tknSrv,
		SessionSrv:  session.Service,
	})
	rpc := rpc.New(rpc.Config{
		Port: config.Rpc.Port,
		App:  app,
		I18n: *i18n,
	})
	go rpc.Listen()
	http.Listen()
}

func loadBusinessMongo(cnf config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  cnf.DB.MongoBusiness.Host,
		Port:  cnf.DB.MongoBusiness.Port,
		User:  cnf.DB.MongoBusiness.Username,
		Pass:  cnf.DB.MongoBusiness.Password,
		Db:    cnf.DB.MongoBusiness.Database,
		Query: cnf.DB.MongoBusiness.Query,
	})
	d, err := mongo.New(uri, cnf.DB.MongoBusiness.Database)
	if err != nil {
		panic(err)
	}
	return d
}
