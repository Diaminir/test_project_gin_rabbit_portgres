package servergin

import (
	"data_consumer/handlers"

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
		api.GET("/orders", s.app.HandlerShowAllOrders)
	}
	err := router.Run(":8083")
	if err != nil {
		return err
	}
	return nil
}
