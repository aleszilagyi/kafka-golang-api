package main

import (
	"database/sql"
	"encoding/json"

	courseUsecase "github.com/aleszilagyi/kafka-golang-api/domain/usecase"
	courseRepository "github.com/aleszilagyi/kafka-golang-api/adapters/repository"
	kafkaCourse "github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
	logWrapper "github.com/aleszilagyi/kafka-golang-api/adapters/config/logger"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log := logWrapper.NewLogger()

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/courses")
	if err != nil {
		log.Error(err)
	}
	repository := courseRepository.CourseMySQLRepository{Db: db}
	usecase := courseUsecase.CreateCourse{Repository: repository}

	var messageChannel = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "kafka-course",
		"group.id":          "kafka-course",
	}
	topics := []string{"courses"}
	consumer := kafkaCourse.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(messageChannel)

	for msg := range messageChannel {
		var input = kafkaCourse.NewCourseInputDto()
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)
		if err != nil {
			log.Error(err.Error())
			} else {
			log.Infof("Message consumed with output: %s", output)
		}
	}
}
