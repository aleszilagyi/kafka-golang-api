package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	logWrapper "github.com/aleszilagyi/kafka-golang-api/adapters/config/logger"
	kafkaCourse "github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
	courseRepository "github.com/aleszilagyi/kafka-golang-api/adapters/repository"
	courseUsecase "github.com/aleszilagyi/kafka-golang-api/domain/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	viper "github.com/spf13/viper"
)

func main() {
	log := logWrapper.NewLogger()

	viper.SetConfigName("application")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbConfigs := viper.GetStringMapString("application.mysql")
	dbString := fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s", dbConfigs["user"], dbConfigs["password"], dbConfigs["port"], dbConfigs["table"])

	configMapConsumer := &ckafka.ConfigMap{}

	topics := viper.GetStringSlice("application.kafka.topics")
	viper.UnmarshalKey("application.kafka.configs", configMapConsumer)

	db, err := sql.Open("mysql", dbString)
	if err != nil {
		log.Error(err)
	}
	repository := courseRepository.CourseMySQLRepository{Db: db}
	usecase := courseUsecase.CreateCourse{Repository: repository}

	var messageChannel = make(chan *ckafka.Message)

	consumer := kafkaCourse.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(messageChannel)

	for msg := range messageChannel {
		var input = kafkaCourse.NewCourseInputDto()

		err := json.Unmarshal(msg.Value, &input)
		if err == nil {
			output, err := usecase.Execute(input)
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Infof("Message consumed with output: %s", output)
			}
		} else {
			log.Errorf("Unmarshal failed for: %s", msg.Value)
		}

	}
}
