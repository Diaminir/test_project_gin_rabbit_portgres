package rabbitconsum

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	tr *amqp091.Channel
}

func NewRabbitLoginAndCreate() (*RabbitMQ, error) {
	connFig := "amqp://guest:guest@rabbitmq:5672/"
	conn, err := amqp091.Dial(connFig)
	if err != nil {
		return &RabbitMQ{}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return &RabbitMQ{}, err
	}
	if err := ch.ExchangeDeclare("logs", "direct", true, false, false, false, nil); err != nil {
		return &RabbitMQ{}, err
	}
	q, err := ch.QueueDeclare("consumer", true, false, true, false, nil)
	if err != nil {
		return nil, err
	}
	if err := ch.QueueBind(q.Name, "rabbit", "logs", false, nil); err != nil {
		return nil, err
	}
	fmt.Println("rabbit успешно создался")
	return &RabbitMQ{
		tr: ch,
	}, nil
}

func (rb *RabbitMQ) ConnClose() {
	rb.tr.Close()
}

func (rb *RabbitMQ) ConsumeMessages() (<-chan amqp091.Delivery, error) {
	msgs, err := rb.tr.Consume("consumer", "rabbitConsume", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
