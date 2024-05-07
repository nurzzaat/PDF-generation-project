package pkg

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                string `mapstructure:"APP_ENV"`
	ServerAddress         string `mapstructure:"SERVER_ADDRESS"`
	PORT                  string `mapstructure:"PORT"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPass                string `mapstructure:"DB_PASSWORD"`
	DBName                string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret     string `mapstructure:"ACCESS_TOKEN_SECRET"`
	UnidocLisenseKey      string `mapstructure:"UNIDOC_LICENSE_DIR"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The app is running in development env")
	}

	return &env

}
