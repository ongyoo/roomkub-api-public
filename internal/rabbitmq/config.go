package rabbitmq

import (
	"github.com/caarlos0/env"
)

type Config struct {
	RabbitMQURI string `env:"RABBIT_MQ_URI,required"`
}

func ReadRabbitMQConfig() Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config
}
