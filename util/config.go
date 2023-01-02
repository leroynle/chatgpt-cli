package util

import "github.com/spf13/viper"

type EnvConfig struct {
	ChatGPTToken string `mapstructure:"CHATGPTTOKEN"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config EnvConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
