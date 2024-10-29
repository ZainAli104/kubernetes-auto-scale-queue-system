package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TaskProcessor(task string) {
	log.Printf("Processing task: %s", task)
	time.Sleep(2 * time.Second)
	log.Printf("Completed task: %s", task)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Step 1: Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Step 2: Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Step 3: Declare a queue
	queueName := "task_queue"
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// Step 4: Start consuming messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	log.Printf("Waiting for messages in %s. To exit press CTRL+C", queueName)

	// Step 5: Process messages
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// Process the message
			TaskProcessor(string(d.Body))
		}
	}()

	<-forever
}
