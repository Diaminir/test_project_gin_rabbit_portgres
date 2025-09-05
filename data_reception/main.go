package main

import (
	"fmt"
	"gin_db/db"
	"gin_db/handlers"
	rabbitreceive "gin_db/rabbitReceive"
	servergin "gin_db/serverGin"
)

func main() {
	postgresDB, err := db.NewDbLogin()
	ErrorErr(err)
	defer postgresDB.DbClose()
	rabbitMQ, err := rabbitreceive.NewRabbitLoginAndCreate()
	ErrorErr(err)
	defer rabbitMQ.ConnClose()
	handler := handlers.NewApp(postgresDB, rabbitMQ)
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
