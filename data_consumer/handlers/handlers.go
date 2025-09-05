package handlers

import (
	"data_consumer/db"
	"data_consumer/dto"
	rabbitconsum "data_consumer/rabbitConsum"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	pg *db.Postgres
	rb *rabbitconsum.RabbitMQ
}

func NewApp(db *db.Postgres, cons *rabbitconsum.RabbitMQ) *App {
	return &App{
		pg: db,
		rb: cons,
	}
}

func (app *App) HandlerShowAllOrders(c *gin.Context) {
	dataSlice := []dto.DbSearchDTO{}
	dataDb, err := app.pg.DbSearch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorDTO(err))
		return
	}
	for {
		if dataDb.Next() {
			dataText, err := dataDb.Values()
			if err != nil {
				return
			}
			data := dto.NewDbSearchDTO(dataText[0].(int32), dataText[1].(string), dataText[2].(float64), dataText[3].(time.Time))
			dataSlice = append(dataSlice, data)
		} else {
			break
		}
	}
	c.JSON(http.StatusOK, dataSlice)
}

func (app *App) RabbitTiDB() error {
	msgs, err := app.rb.ConsumeMessages()
	if err != nil {
		return err
	}
	var consRabbit dto.ConsumeRabbitDTO
	go func() {
		for msg := range msgs {
			if err := json.Unmarshal(msg.Body, &consRabbit); err != nil {
				fmt.Println("ошибка преобразования JSON:", err)
				return
			}
			if err := app.pg.DbPut(consRabbit); err != nil {
				fmt.Println("ошибка записи в БД:", err)
				return
			}
		}
	}()
	return nil
}
