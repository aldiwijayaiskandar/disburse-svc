package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Consumer for receiving AMPQ events
type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type ConsumerInterface interface {
	Listen(topics []string, listener func(delivery *amqp.Delivery)) error
	WaitReply(correlationId string) (*amqp.Delivery, error)
}

func (c Consumer) setup() error {
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

// NewConsumer returns a new Consumer
func NewConsumer(conn *amqp.Connection) (ConsumerInterface, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

// Listen will listen for all new Queue publications
// and print them to the console.
func (c Consumer) Listen(topics []string, listener func(delivery *amqp.Delivery)) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		c.queueName, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	for _, s := range topics {
		err = ch.QueueBind(
			q.Name,
			s,
			getExchangeName(),
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			listener(&d)
		}
	}()

	log.Printf("[*] Waiting for message [Exchange, Queue][%s, %s]. To exit press CTRL+C", getExchangeName(), q.Name)
	<-forever
	return nil
}

func (c Consumer) WaitReply(correlationId string) (*amqp.Delivery, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"reply_queue", // name
		false,         // durable
		false,         // delete when unused
		true,          // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	replyChan := make(chan *amqp.Delivery)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			if d.CorrelationId == correlationId {
				replyChan <- &d
				return
			}
		}
	}()

	log.Printf("[*] Waiting for message [Exchange, Queue][%s, %s]. To exit press CTRL+C", getExchangeName(), q.Name)

	return <-replyChan, nil
}
