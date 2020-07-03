package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	kafkaBrokerAddr = []string{"172.17.40.166:9092"}
	groupId         = "yangchun"
)

func main() {
	fmt.Println("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_0_0_0
	// consumer
	consumer, err := sarama.NewConsumer(kafkaBrokerAddr, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()
	for i := 0; i < 24; i++ {
		go func(i int) {
			fmt.Println("start ", i)
			partition_consumer, err := consumer.ConsumePartition("test-topic-yangchun1", int32(i), sarama.OffsetOldest)
			if err != nil {
				fmt.Printf("try create partition_consumer error %s\n", err.Error())
				return
			}

			for {
				select {
				case msg := <-partition_consumer.Messages():
					fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
						msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				case err := <-partition_consumer.Errors():
					fmt.Printf("err :%s\n", err.Error())
					partition_consumer.Close()
				}
			}

		}(i)
	}
	time.Sleep(100 * time.Second)
}

//type consumerHandler struct{}
//
//// Setup is run at the beginning of a new session, before ConsumeClaim.
//func (c *consumerHandler) Setup(sarama.ConsumerGroupSession) error {
//	return nil
//}
//
//// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
//// but before the offsets are committed for the very last time.
//func (c *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
//	return nil
//}
//
//// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
//// Once the Messages() channel is closed, the Handler must finish its processing
//// loop and exit.
//func (c *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//
//	// NOTE:
//	// Do not move the code below to a goroutine.
//	// The `ConsumeClaim` itself is called within a goroutine, see:
//	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
//	msgChan := claim.Messages()
//	for {
//		message := <-msgChan
//		if message == nil {
//			fmt.Println(msgChan)
//			continue
//		}
//		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
//		session.MarkMessage(message, "")
//	}
//
//	return nil
//}
