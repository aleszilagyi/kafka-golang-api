package env

import (
	"github.com/aleszilagyi/kafka-golang-api/adapters/config/logger"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

var log *logger.StandardLogger

func init() {
	log = logger.NewLogger()
}

type AppEnv struct {
	Kafka    KafkaEnv `mapstructure:"kafka"`
	MysqlEnv MysqlEnv `mapstructure:"mysql"`
}

type KafkaEnv struct {
	KafkaTopics []string `mapstructure:"topics"`
}

type MysqlEnv struct {
	Table    string `mapstructure:"table"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func GetConf() *AppEnv {
	readEnv()
	envs := &AppEnv{}
	err := viper.UnmarshalKey("application", envs)
	if err != nil {
		log.Errorf("unable to decode into config struct, %v", err)
		panic(err)
	}
	return envs
}

func GetKafkaEnvs(configMap *ckafka.ConfigMap) {
	readEnv()
	viper.UnmarshalKey("application.kafka.configs", configMap)
}

func readEnv() {
	viper.AddConfigPath("./env")
	viper.SetConfigName("application")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
}
