package config

// import "github.com/spf13/viper"

// type Gemini struct {
// 	ApiKey string `mapstructure:"api_key"`
// 	Model  string `mapstructure:"model"`
// }
// type Database struct {
// 	Username string `mapstructure:"username"`
// 	Password string `mapstructure:"password"`
// 	Uri      string `mapstructure:"uri"`
// 	Name     string `mapstructure:"name"`
// }
// type Email struct {
// 	EmailKey string `mapstructure:"key"`
// }
// type Config struct {
// 	Database Database `mapstructure:"database"`
// 	Email    Email    `mapstructure:"email"`
// 	Port     string   `mapstructure:"port"`
// 	Jwt      Jwt      `mapstructure:"jwt"`
// 	Gemini   Gemini   `mapstructure:"gemini"`
// 	Domain  string   `mapstructure:"domain"`
// }
// type Jwt struct {
// 	JwtKey string `mapstructure:"jwtKey"`
// }

// func LoadConfig() (*Config, error) {
// 	viper.AddConfigPath("../")
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	if err := viper.ReadInConfig(); err != nil {
// 		return &Config{}, err
// 	}
// 	config := Config{}
// 	if err := viper.Unmarshal(&config); err != nil {
// 		return &config, err
// 	}
// 	return &config, nil
// }

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Username string
		Password string
		Uri      string
		Name     string
	}
	Email struct {
		Key string
	}
	Port  string
	JWT   string
	Gemini struct {
		ApiKey string
		Model  string
	}
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	cfg := &Config{
		Database: struct {
			Username string
			Password string
			Uri      string
			Name     string
		}{
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Uri:      viper.GetString("DATABASE_URI"),
		},
		Email: struct{ Key string }{
			Key: viper.GetString("EMAIL_KEY"),
		},
		Port: viper.GetString("PORT"),
		JWT:  viper.GetString("JWT"),
		Gemini: struct {
			ApiKey string
			Model  string
		}{
			ApiKey: viper.GetString("GEMINI_API_KEY"),
			Model:  viper.GetString("GEMINI_MODEL"),
		},
	}

	return cfg, nil
}
