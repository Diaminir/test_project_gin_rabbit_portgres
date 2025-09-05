package rabbitreceive

import "github.com/rabbitmq/amqp091-go"

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
	return &RabbitMQ{
		tr: ch,
	}, nil

}

func (rb *RabbitMQ) ConnClose() {
	rb.tr.Close()
}

func (rb *RabbitMQ) PublishMessage(jsTrRabbit []byte) error {
	if err := rb.tr.Publish(
		"logs",
		"rabbit",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        jsTrRabbit,
		},
	); err != nil {
		return err
	}
	return nil
}
