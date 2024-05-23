package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Coffee", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Read messages from a queue into a channel
	msgs, err := ch.Consume(
		q.Name,
		"",
		true, // auto-ack
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Print("Recieved a coffee request", string(d.Body))
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
