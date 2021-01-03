package queue

import (
	"encoding/json"
	"fmt"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/streadway/amqp"
)

type RabbitQueue struct {
	config         *common.Config
	enqueueConn    *amqp.Connection
	enqueueChannel *amqp.Channel
	dequeueConn    *amqp.Connection
	dequeueChannel *amqp.Channel
}

func newRabbitQueue(config *common.Config) Queue {
	return &RabbitQueue{config: config}
}

func (q *RabbitQueue) Enqueue(archive *domain.Archive) error {
	logger := common.GetLogger()
	config := q.config
	rabbitMQ := config.RabbitMQ
	rabbitMQUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitMQ.Username, rabbitMQ.Password, rabbitMQ.Host, rabbitMQ.Port)
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		logger.Printf("amqp - %v", rabbitMQUrl)
		logger.Printf("connect to rabbitMQ(amqp) fail: %v", err)
		return err
	}
	q.enqueueConn = conn
	// defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Printf("connect to rabbitMQ(amqp) fail: %v", err)
		return err
	}
	q.enqueueChannel = ch
	// defer ch.Close()

	_, err = ch.QueueDeclare(rabbitMQ.Queue, false, false, false, false, nil)
	if err != nil {
		logger.Printf("connect to rabbitMQ(amqp) fail: %v", err)
		return err
	}

	data, err := json.Marshal(archive)
	if err != nil {
		logger.Printf("Publish: marshal message fail: %v", err)
		return err
	}

	err = ch.Publish("", rabbitMQ.Queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})

	if err != nil {
		logger.Printf("Publish message fail: %v", err)
		return err
	}
	return nil
}

func (q *RabbitQueue) Channel() (chan *domain.Archive, error) {
	logger := common.GetLogger()
	config := q.config

	rabbitMQUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.RabbitMQ.Username, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port)
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		logger.Printf("amqp - %v", rabbitMQUrl)
		logger.Fatalf("connect to rabbitMQ(amqp) fail: %v", err)
	}
	q.dequeueConn = conn
	// defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("connect to rabbitMQ(amqp) fail: %v", err)
	}
	q.dequeueChannel = ch
	// defer ch.Close()

	_, err = ch.QueueDeclare(config.RabbitMQ.Queue, false, false, false, false, nil)
	if err != nil {
		logger.Fatalf("connect to rabbitMQ(amqp) fail: %v", err)
	}

	messageChannel, err := ch.Consume(config.RabbitMQ.Queue, "", true, false, false, false, nil)
	if err != nil {
		logger.Fatalf("consume message fail: %v", err)
	}

	c := make(chan *domain.Archive)
	go func() {
		for msg := range messageChannel {
			archive := new(domain.Archive)
			err := json.Unmarshal(msg.Body, archive)
			if err != nil {
				logger.Printf("unmarshal fail: %v", err)
				return
			}
			c <- archive
		}
	}()
	logger.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return c, nil
}

func (q *RabbitQueue) Close() {
	if q.enqueueChannel != nil {
		q.enqueueChannel.Close()
	}
	if q.enqueueConn != nil {
		q.enqueueChannel.Close()
	}

	if q.dequeueChannel != nil {
		q.dequeueChannel.Close()
	}
	if q.dequeueConn != nil {
		q.dequeueChannel.Close()
	}
}
