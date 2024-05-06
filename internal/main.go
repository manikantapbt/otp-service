package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/twilio/twilio-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"os"
	"os/signal"
	"otp-service/internal/gateway"
	v1 "otp-service/internal/gen/v1"
	"otp-service/internal/gen/v1/v1connect"
	"otp-service/internal/server"
	"otp-service/internal/service"
	"syscall"
	"time"
)

func main() {
	otpInterval := 10 * time.Minute
	generator := service.NewOtpGenerator("your_secret_key", otpInterval)
	sender := getOTPSender()
	iService := service.NewService(sender, generator)
	otpService := server.NewOtpService(iService)
	mux := http.NewServeMux()
	path, handler := v1connect.NewOTPServiceHandler(otpService)
	mux.Handle(path, handler)
	go func() {
		log.Println("Starting server on localhost:8082")
		if err := http.ListenAndServe("localhost:8082", h2c.NewHandler(mux, &http2.Server{})); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()
	conn, err := amqp.Dial("amqps://<username>:<password>@<host>/<virtual_host>")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	queueName := "teja"
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,         // queue
		"otp-listener", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Handle incoming messages
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
			}
		}()
		for msg := range msgs {
			request := &v1.GenerateOTPRequest{}
			err2 := proto.Unmarshal(msg.Body, request)
			if err2 != nil {
				log.Printf("Failed to unmarshal request: %v", err2)
			} else {
				iService.SendMessage(request)
			}
		}
	}()

	fmt.Println("RabbitMQ listener server started")

	// Wait for termination signal to gracefully shutdown the server
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Shutting down RabbitMQ listener server")
}

func getOTPSender() gateway.OTPSender {
	accountSid := "twilio_sid"
	authToken := "twilio_auth_token"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	from := "twilio_from_phone_number"
	return gateway.NewTwilioOTPSender(client, from, accountSid, authToken)
}
