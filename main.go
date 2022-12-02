package main

import (
	_middlewares "backend/app/middlewares"
	_routes "backend/app/routes"
	_utils "backend/utils"

	"fmt"

	_driverFactory "backend/drivers"
	_dbDriver "backend/drivers/mysql"

	_userUseCase "backend/businesses/users"
	_userController "backend/controllers/users"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "3000"

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _utils.GetConfig("DB_PASSWORD"),
		DB_HOST: _utils.GetConfig("DB_HOST"),
		DB_PORT: _utils.GetConfig("DB_PORT"),
		DB_NAME: _utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT: _utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	app := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUseCase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware: configJWT.Init(),
		AuthController: *userCtrl,
	}

	routesInit.RouteRegister(app)

	var port string

	if port == "" {
		port = DEFAULT_PORT
	}

	var appPort string = fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))
}
