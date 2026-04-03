package main

import (
	"email-service-rabbitmq/shared"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)

var ch *amqp.Channel
var queueName = "email_queuee"

func sendEmail(email shared.Email, FromEmail string, appPassword string) {
	m := gomail.NewMessage()
	m.SetHeader("From", FromEmail)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)

	m.SetBody("text/html", email.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, FromEmail, appPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email")
		return
	}
	log.Printf("Email sent successfully to %s\n", email.To)
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	FromEmail := os.Getenv("APP_EMAIL")
	appPassword := os.Getenv("APP_PASSWORD")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect rabittmQ %v", err)
	}
	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open the channel or connection")
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queueName,
		true,  // Durable ✅
		false, // AutoDelete ❌
		false, // Exclusive ❌
		false, // NoWait ❌
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare the queue %v", err)
	}
	fmt.Println(queue)

	msgs, err := ch.Consume(queueName, "", true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalf("Failed to register the consumer")
	}
	chanwait := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Println(msg.Body)
			var email shared.Email
			err := json.Unmarshal(msg.Body, &email)
			if err != nil {
				continue
			}
			sendEmail(email, FromEmail, appPassword)

		}

	}()
	<-chanwait

}
