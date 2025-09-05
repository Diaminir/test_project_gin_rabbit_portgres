package servergin

import (
	"gin_db/handlers"

	"github.com/gin-gonic/gin"
)

type ServerGin struct {
	app *handlers.App
}

func NewServer(app *handlers.App) *ServerGin {
	return &ServerGin{
		app: app,
	}
}

func (s *ServerGin) StartServer() error {
	router := gin.Default()
	api := router.Group("/")
	{
		api.POST("/cars", s.app.HandlerCars)
	}
	err := router.Run(":8082")
	if err != nil {
		return err
	}
	return nil
}
