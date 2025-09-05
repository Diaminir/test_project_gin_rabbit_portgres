package dto

import "time"

type DesiredCarDTO struct {
	Title      string `json:"title"`
	Model      string `json:"model"`
	Engine     string `json:"engine"`
	Genetation string `json:"generation"`
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type TransferRabbitDTO struct {
	CarInfo string  `json:"carInfo"`
	Price   float64 `json:"price"`
}

type CarInfoBD struct {
	Name       string  `json:"title"`
	Model      string  `json:"model"`
	Engine     string  `json:"engine"`
	Genetation string  `json:"generation"`
	Price      float64 `josn:"price"`
}

func NewErrorDTO(message error) ErrorDTO {
	return ErrorDTO{
		Message: message.Error(),
		Time:    time.Now(),
	}
}

func NewTranferRabbitDTO(carInfo string, price float64) TransferRabbitDTO {
	return TransferRabbitDTO{
		CarInfo: carInfo,
		Price:   price,
	}
}
