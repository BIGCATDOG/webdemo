package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const (
	RabbitMQAddr = "amqp://guest:guest@localhost:5672/"
)

func logErr(err error)  {
	if err!=nil{
		log.Fatal(err)
	}
}
func main()  {
	conn,err:=amqp.Dial(RabbitMQAddr)
	logErr(err)
	ch, err := conn.Channel()
	logErr(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	//err = ch.Qos(
	//	1,     // prefetch count
	//	0,     // prefetch size
	//	false, // global
	//)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	logErr(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Reveived mess type %s ",d.ContentType)
			log.Printf("Received a message: %s", d.Body)
			var dotCount = bytes.Count(d.Body,[]byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t*time.Second)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}