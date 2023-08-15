package event_stream

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func (s Server) ListenAccountCreated(data []byte) {
	logrus.Info("Account created event received")
	d := s.dto.AccountCreated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountCreate.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountUpdated(data []byte) {
	logrus.Info("Account updated event received")
	d := s.dto.AccountUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountUpdate.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountDeleted(data []byte) {
	logrus.Info("Account deleted event received")
	d := s.dto.AccountDeleted()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountDelete.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountEnabled(data []byte) {
	logrus.Info("Account enabled event received")
	d := s.dto.AccountEnabled()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountEnable.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountDisabled(data []byte) {
	logrus.Info("Account disabled event received")
	d := s.dto.AccountDisabled()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountDisable.Handle(s.ctx, d.ToCommand())
}
