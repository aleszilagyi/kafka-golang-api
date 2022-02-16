package kafka

import (
	logWrapper "github.com/aleszilagyi/kafka-golang-api/adapters/config/logger"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var log = logWrapper.NewLogger()

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(messageChannel chan *ckafka.Message) {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		log.Error(err.Error())
		panic("App is shutting down due to reason: Error creating consumer")
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		log.Error(err.Error())
		panic("App is shutting down due to reason: Error subscribing to topics")
	}

	log.Info("Consumer successfully connected to broker")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			messageChannel <- msg
		} else {
			log.Error(err)
		}
	}
}
