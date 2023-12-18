package service

import (
	"fiber-wire-template/pkg/log"
	"fiber-wire-template/pkg/util/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
}

func NewService(logger *log.Logger, sid *sid.Sid) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
	}
}
