package event_stream

import (
	"context"

	"github.com/mixarchitecture/microp/events"
	"github.com/sirupsen/logrus"
	"github.com/turistikrota/service.business/src/app"
	"github.com/turistikrota/service.business/src/config"
	"github.com/turistikrota/service.business/src/delivery/event_stream/dto"
)

type Server struct {
	app    app.Application
	Topics config.Topics
	engine events.Engine
	ctx    context.Context
	dto    dto.Dto
}

type Config struct {
	App    app.Application
	Topics config.Topics
	Engine events.Engine
	Ctx    context.Context
}

func New(config Config) Server {
	return Server{
		app:    config.App,
		engine: config.Engine,
		Topics: config.Topics,
		ctx:    config.Ctx,
		dto:    dto.New(),
	}
}

func (s Server) Load() {
	logrus.Info("Loading event stream server")
}
