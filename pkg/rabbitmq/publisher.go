package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Emitter for publishing AMQP events
type Publisher struct {
	conn *amqp.Connection
}

type PublisherInterface interface {
	Push(key string, replyTo string, body []byte, correlationId string) error
	Reply(key string, body []byte, correlationId string) error
}

func (p *Publisher) setup() error {
	channel, err := p.conn.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()
	return declareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (p *Publisher) Push(key string, replyTo string, body []byte, correlationId string) error {
	channel, err := p.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	channel.Publish(
		getExchangeName(),
		key,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          body,
			CorrelationId: correlationId,
			ReplyTo:       "message.reply",
		},
	)

	log.Printf("Sending message:  %s -> %s", body, getExchangeName())
	return nil
}

func (p *Publisher) Reply(key string, body []byte, correlationId string) error {
	channel, err := p.conn.Channel()
	if err != nil {
		return err
	}

	q, _ := channel.QueueDeclare(
		"reply_queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	channel.QueueBind(
		q.Name,
		key,
		getExchangeName(),
		false,
		nil,
	)

	defer channel.Close()

	channel.Publish(
		getExchangeName(),
		key,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          body,
			CorrelationId: correlationId,
			ReplyTo:       key,
		},
	)

	log.Printf("Sending message:  %s -> %s", body, getExchangeName())
	return nil
}

// NewPublisher returns a new event.Emitter object
// ensuring that the object is initialised, without error
func NewPublisher(conn *amqp.Connection) (PublisherInterface, error) {
	publisher := &Publisher{
		conn: conn,
	}

	err := publisher.setup()
	if err != nil {
		return &Publisher{}, err
	}

	return publisher, nil
}
