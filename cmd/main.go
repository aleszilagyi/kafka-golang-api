package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/aleszilagyi/kafka-golang-api/adapters/config/logger"
	kafkaCourse "github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
	courseRepository "github.com/aleszilagyi/kafka-golang-api/adapters/repository"
	courseUsecase "github.com/aleszilagyi/kafka-golang-api/domain/usecase"
	"github.com/aleszilagyi/kafka-golang-api/resources"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

var (
	log *logger.StandardLogger
	env *resources.AppEnv
)

func init() {
	log = logger.NewLogger()
	env = resources.GetConf()
}

func main() {
	log.Info("application starting")

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DbConfig.User, env.DbConfig.Hostname, env.DbConfig.Password, env.DbConfig.Port, env.DbConfig.Database)

	configMapConsumer := env.Kafka.KafkaConfigs

	topics := env.Kafka.Topics

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
				log.Infof("message consumed with output: %s", output)
			}
		} else {
			log.Errorf("unmarshal failed for: %s", msg.Value)
		}

	}
}
