package main

import (
  "log"
  "github.com/streadway/amqp"
  "syscall"
  "os"
  "fmt"
  "github.com/joho/godotenv"
)

func main() {
    er := godotenv.Load(".env")

    failOnError(er, "Error load .env file")

    user := os.Getenv("USERNAME")
    pass := os.Getenv("PASSWORD")
    host := os.Getenv("HOST")
    port := os.Getenv("PORT")

    amqp_str := fmt.Sprintf("amqp://%s:%s@%s:%s/",user,pass,host,port)

	connection, e := amqp.Dial(amqp_str)

	failOnError(e, "Failed to connect to RabbitMQ")

	defer connection.Close()

	channel, e := connection.Channel()
	failOnError(e, "Failed to open channel")
	defer channel.Close()

	queue, e := channel.QueueDeclare(
		"test", // name
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(e, "Failed to declare a queue")

	var stat syscall.Statfs_t

    wd, e := os.Getwd()

    syscall.Statfs(wd, &stat)

    // Available blocks * size per block = available space in bytes
    body := fmt.Sprintf("Available disk space: %v GB", (stat.Bavail * uint64(stat.Bsize)) / 1e9)

	e = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v", body)),
		})
	log.Printf("message sent: %s", body)
	failOnError(e, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}