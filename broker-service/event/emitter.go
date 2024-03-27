package event

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setUp() error {
	channel, err := e.connection.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			log.Println(err)
		}
	}(channel)

	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	chanel, err := e.connection.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	defer func(chanel *amqp.Channel) {
		err := chanel.Close()
		if err != nil {
			log.Println(err)
		}
	}(chanel)

	log.Println("Pushing to channel")

	err = chanel.PublishWithContext(
		context.TODO(),
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		})
	if err != nil {
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}
	err := emitter.setUp()
	if err != nil {
		log.Println(err)
		return Emitter{}, err
	}

	return emitter, nil
}
