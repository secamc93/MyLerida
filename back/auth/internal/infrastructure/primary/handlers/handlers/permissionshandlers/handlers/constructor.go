package handlers

import (
	"auth/internal/application/usecase/usecasepermissions"
	"auth/pkg/logger"
)

type Permissionshandlers struct {
	usecasepermissions usecasepermissions.IPermissionsUseCase
	log                logger.ILogger
}

func New(usecasepermissions usecasepermissions.IPermissionsUseCase) *Permissionshandlers {
	return &Permissionshandlers{
		usecasepermissions: usecasepermissions,
		log:                logger.NewLogger(),
	}
}
