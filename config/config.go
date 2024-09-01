package config

import (
	"os"

	"github.com/caarlos0/env/v6"
)

type ApiConfig struct {
	Port    string
	Secrets Secrets
	Swagger SwaggerConfig
	Root    Root
}

type Secrets struct {
	JWTSecret string `env:"JWT_KEY"`
	WDEKS     string `env:"WDEKS"`
}

type SwaggerConfig struct {
	Host string `env:"SWAGGER_HOST"`
}

type Root struct {
	RootRoleSlug string `env:"ROOT_ROLE_SLUG"`
}

func ReadApi() ApiConfig {
	var (
		secretsConfig Secrets
		swaggerConfig SwaggerConfig
		rootConfig    Root
	)

	privateKey := os.Getenv("PRIVATE_KEY")
	println(privateKey)

	if err := env.Parse(&secretsConfig); err != nil {
		panic(err)
	}

	if err := env.Parse(&swaggerConfig); err != nil {
		panic(err)
	}

	if err := env.Parse(&rootConfig); err != nil {
		panic(err)
	}

	return ApiConfig{
		Port:    os.Getenv("SERVICE_PORT"),
		Secrets: secretsConfig,
		Swagger: swaggerConfig,
		Root:    rootConfig,
	}
}
