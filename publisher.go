package main

import (
	"fmt"
	"log"
	"os"
	"time"

	message "github.com/dohernandez/proto-publisher/pkg/resources/proto"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Proto publisher: publishing message RegionWasCreated")

	// conn
	conn, err := amqp.Dial("amqp://admin:admin@localhost:9999/")
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:9999/")
	if err != nil {
		log.Printf("ERROR: fail init consumer: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: done init producer conn")

	// create channel
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR: fail create channel: %s", err.Error())
		os.Exit(1)
	}

	msg := message.RegionWasCreated{
		RegionId: "dad2d9ff-c1cd-44ba-8601-cd0d58f27e51",
		Country:  "US",
		Handle:   "SHUTDOWN-2020-W13",
		Name:     "SHUTDOWN-2020-W13",
		ListId:   "4bef27f1-eb6a-4b21-9a23-2a876d840194",
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		log.Printf("ERROR: fail marshal: %s", err.Error())
		os.Exit(1)
	}
	log.Printf("DEBUG: publishing data: %+v", data)

	// publish message
	err = amqpChannel.Publish(
		"delivery_regions", // exchange
		"region.created",   // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			Timestamp: time.Now(),
			Body:      data,
			Headers: amqp.Table{
				"Message-Type":      ".HelloFresh.Message.DeliveryRegions.RegionWasCreated",
				"Content-Encoding":  "binary",
				"Content-Type":      "application/octet-stream",
				"X-B3-Sampled":      "0",
				"X-B3-SpanId":       "",
				"X-B3-TraceId":      "",
				"ot-tracer-sampled": "false",
				"ot-tracer-spanid":  "",
				"ot-tracer-traceid": "",
			},
		},
	)
	if err != nil {
		log.Printf("ERROR: fail publish msg: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: published msg: %v", msg)
}
