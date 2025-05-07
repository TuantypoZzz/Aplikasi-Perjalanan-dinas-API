package configuration

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ() *amqp.Connection {
	// Ambil URL RabbitMQ dari environment variable
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	// Buat koneksi ke RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	log.Println("Successfully connected to RabbitMQ")
	return conn
}
