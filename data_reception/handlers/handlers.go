package handlers

import (
	"encoding/json"
	"fmt"
	"gin_db/db"
	"gin_db/dto"
	rabbitreceive "gin_db/rabbitReceive"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	pg *db.Postgres
	rb *rabbitreceive.RabbitMQ
}

func NewApp(db *db.Postgres, tr *rabbitreceive.RabbitMQ) *App {
	return &App{
		pg: db,
		rb: tr,
	}
}

func (app *App) HandlerCars(c *gin.Context) {
	var desCar dto.DesiredCarDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&desCar); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorDTO(err))
		return
	}
	infoDB, err := app.pg.DbSearch(desCar)
	// fmt.Println(infoDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorDTO(err))
		return
	}
	str := fmt.Sprintf("Марка: %s. Модель: %s. Двигатель: %s. Поколение: %s.\n", infoDB.Name, infoDB.Model, infoDB.Engine, infoDB.Genetation)
	trRabbit := dto.NewTranferRabbitDTO(str, infoDB.Price)
	// fmt.Println(trRabbit)
	jsTrRabbit, err := json.Marshal(trRabbit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorDTO(err))
		return
	}
	// fmt.Println(jsTrRabbit)
	if err := app.rb.PublishMessage(jsTrRabbit); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorDTO(err))
		return
	}
}
