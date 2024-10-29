package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	RPCUrl     string `mapstructure:"RPC_URL"`
	RPCPort    string `mapstructure:"RPC_PORT"`
	PrivateKey string `mapstructure:"PRIVATE_KEY"`
	ChainID    uint64 `mapstructure:"CHAIN_ID"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	return &env
}
