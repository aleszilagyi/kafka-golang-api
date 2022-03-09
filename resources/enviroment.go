package resources

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
	Kafka KafkaConfig `mapstructure:"kafka"`
	DbConfig    DbConfig    `mapstructure:"mysql"`
}

type DbConfig struct {
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Table    string `mapstructure:"table"`
	Database string `mapstructure:"database"`
	Hostname string `mapstructure:"hostname"`
}

type KafkaConfig struct {
	KafkaConfigs     *ckafka.ConfigMap `mapstructure:"configs"`
	Topics           []string          `mapstructure:"topics"`
}

func GetConf() *AppEnv {
	viper.AddConfigPath("./resources")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}

	conf := &AppEnv{}
	err = viper.UnmarshalKey("application", conf)
	if err != nil {
		log.Errorf("unable to decode into config struct, %v", err)
		panic(err)
	}
	return conf
}
