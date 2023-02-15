package config

import "github.com/spf13/viper"

type Config struct {
	APIPort      string `mapstructure: "API_PORT"`
	Postgres_URL string `mapstructure: "POSTGRES_DB_URL"`
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

}
