package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/lfkodama/son-fullcycle-simulator/application/kafka"
	"github.com/lfkodama/son-fullcycle-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {

		fmt.Println(string(msg.Value))

		go kafka2.Produce(msg)
	}

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "readtest", producer)

	// for {
	// 	_ = 1
	// }

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}
