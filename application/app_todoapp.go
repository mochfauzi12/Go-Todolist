package application

import (
	"Go-Todolist/domain_todolist/controller/restapi"
	"Go-Todolist/domain_todolist/gateway/withgorm"
	"Go-Todolist/domain_todolist/usecase/getalltodo"
	"Go-Todolist/domain_todolist/usecase/runtodocheck"
	"Go-Todolist/domain_todolist/usecase/runtodocreate"
	"Go-Todolist/shared/config"
	"Go-Todolist/shared/gogen"
	"Go-Todolist/shared/infrastructure/logger"
	"Go-Todolist/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)

	primaryDriver.AddUsecase(

		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
