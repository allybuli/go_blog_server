package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
	UserService
}

var ServiceGroupApp = new(ServiceGroup)
