package main
import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

const (
	RabbitMQAddr = "amqp://guest:guest@localhost:5672/"
)

func logErr(err error)  {
	if err!=nil{
		log.Fatal(err)
	}
}

func getBody()string{
	if len(os.Args)<2||os.Args[1]==""{
		return "hello"
	}
	return strings.Join(os.Args[1:]," ")
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

	body := getBody()
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}
