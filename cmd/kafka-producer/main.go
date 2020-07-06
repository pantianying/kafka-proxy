package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	kafkaBrokerAddr = []string{"172.17.40.166:32400"}
	topic           = "test-topic-yangchun"
)

func main() {
	var sCfg = sarama.NewConfig()
	sCfg.Producer.Return.Successes = true
	sCfg.Producer.Timeout = 5 * time.Second
	producer, err := sarama.NewSyncProducer(kafkaBrokerAddr, sCfg)
	if err != nil {
		panic(err)
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder("hello world"),
		Key:   sarama.StringEncoder("test-key"),
	}
	for {
		p, offset, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		fmt.Println("send message ok:", p, offset)
		time.Sleep(1 * time.Second)
	}

}
