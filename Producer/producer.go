package main

import (
	"email-service-rabbitmq/shared"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel
var queueName = "email_queuee"

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var e shared.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	rawjson, err := json.Marshal(e)

	err = ch.Publish("", queueName, //Routing Key
		false, //Mandatory
		false, amqp.Publishing{
			ContentType: "application/json",
			Body:        rawjson,
		})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating message", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Email queued successfully ✅"))
}

func main() {
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

	http.HandleFunc("/send-email", sendEmailHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
