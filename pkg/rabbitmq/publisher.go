package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Emitter for publishing AMQP events
type Publisher struct {
	conn *amqp.Connection
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
func (p *Publisher) Push(event string, severity string) error {
	channel, err := p.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	channel.Publish(
		getExchangeName(),
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	log.Printf("Sending message: %s -> %s", event, getExchangeName())
	return nil
}

// NewEventEmitter returns a new event.Emitter object
// ensuring that the object is initialised, without error
func NewEventEmitter(conn *amqp.Connection) (Publisher, error) {
	emitter := Publisher{
		conn: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Publisher{}, err
	}

	return emitter, nil
}