package service

import (
	"fiber-wire-template/pkg/jwt"
	"fiber-wire-template/pkg/log"
	"fiber-wire-template/pkg/util/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
