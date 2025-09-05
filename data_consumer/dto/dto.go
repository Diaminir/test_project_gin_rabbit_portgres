package dto

import "time"

type ConsumeRabbitDTO struct {
	CarInfo string  `json:"carInfo"`
	Price   float64 `json:"price"`
}

type DbSearchDTO struct {
	Order_id  int32   `json:"order_id"`
	CarInfo   string  `json:"carInfo"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"createdAt"`
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewErrorDTO(message error) ErrorDTO {
	return ErrorDTO{
		Message: message.Error(),
		Time:    time.Now(),
	}
}

func NewDbSearchDTO(order_id int32, carInfo string, price float64, createdAt time.Time) DbSearchDTO {
	return DbSearchDTO{
		Order_id:  order_id,
		CarInfo:   carInfo,
		Price:     price,
		CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
	}
}
