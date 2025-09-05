package main

import (
	"data_consumer/db"
	"data_consumer/handlers"
	rabbitconsum "data_consumer/rabbitConsum"
	servergin "data_consumer/serverGin"
	"fmt"
)

func main() {
	postgresDB, err := db.NewDbLogin()
	ErrorErr(err)
	defer postgresDB.DbClose()
	rabbitMQ, err := rabbitconsum.NewRabbitLoginAndCreate()
	ErrorErr(err)
	defer rabbitMQ.ConnClose()
	handler := handlers.NewApp(postgresDB, rabbitMQ)
	err = handler.RabbitTiDB()
	ErrorErr(err)
	serverRun := servergin.NewServer(handler)
	if err := serverRun.StartServer(); err != nil {
		fmt.Println("Остановка сервера")
	}
}

func ErrorErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
