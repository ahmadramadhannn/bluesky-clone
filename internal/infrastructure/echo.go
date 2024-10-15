package infrastructure

import (
	"log"

	"github.com/ahmadramadhannn/bluesky-clone/config"
	"github.com/ahmadramadhannn/bluesky-clone/internal/feature/auth/user"
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	api := e.Group("/api/v1")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	postgres, err := ConnectToPostgres(config.DATABASEURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewUserRepository(postgres)

	userService := user.NewUserService(userRepo)

	user.NewUserHandler(api, userService)

	e.Logger.Fatal(e.Start(":2000"))
}
