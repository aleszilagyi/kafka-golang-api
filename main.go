package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/aleszilagyi/kafka-golang-api/adapters/config/logger"
	kafkaCourse "github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
	courseRepository "github.com/aleszilagyi/kafka-golang-api/adapters/repository"
	courseUsecase "github.com/aleszilagyi/kafka-golang-api/domain/usecase"
	"github.com/aleszilagyi/kafka-golang-api/env"
	"github.com/aleszilagyi/kafka-golang-api/ports"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

var (
	repository ports.RepositoryPort
	usecase    ports.CreateCoursePort
	envs       *env.AppEnv
	log        *logger.StandardLogger
)

func init() {
	envs = env.GetConf()
	log = logger.NewLogger()
}

func main() {
	log.Info("app starting")
	dbString := fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s", envs.MysqlEnv.User, envs.MysqlEnv.Password, envs.MysqlEnv.Port, envs.MysqlEnv.Table)

	configMapConsumer := &ckafka.ConfigMap{}
	env.GetKafkaEnvs(configMapConsumer)

	db, err := sql.Open("mysql", dbString)
	if err != nil {
		log.Error(err)
	}
	repository = courseRepository.CourseMySQLRepository{Db: db}
	usecase = courseUsecase.CreateCourse{Repository: repository}

	var messageChannel = make(chan *ckafka.Message)

	consumer := kafkaCourse.NewConsumer(configMapConsumer, envs.Kafka.KafkaTopics)
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
