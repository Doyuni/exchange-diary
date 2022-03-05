package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	typeEXT      = "yaml"
	defaultPhase = "dev"
)

// Config ...
type Config struct {
	DBConfig DBConfig `mapstructure:"db-config"`
	Client   Client   `mapstructure:"client"`
}

// DBConfig ...
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

// Client ...
type Client struct {
	Kakao Kakao `mapstructure:"kakao"`
}

// OAuthConfig ...
type OAuthConfig struct {
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
	RedirectURL  string `mapstructure:"redirect-url"`
}

// Kakao ...
type Kakao struct {
	Oauth   OAuthConfig `mapstructure:"oauth"`
	BaseURL string      `mapstructure:"base-url"`
}

// Load ...
func Load(path string, name string) (Config, error) {
	config := Config{}
	fmt.Println("Load config file - profile:", name)

	viper.AddConfigPath(path)
	if name == "" {
		name = defaultPhase
	}
	fmt.Println("[PHASE]: ", name)
	viper.SetConfigName(name)
	viper.SetConfigType(typeEXT)

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

// DatabaseConfig ...
func DatabaseConfig() *DBConfig {
	return &DBConfig{
		Host:     viper.GetString("db-config.host"),
		Port:     viper.GetInt("db-config.port"),
		User:     viper.GetString("db-config.user"),
		Name:     viper.GetString("db-config.name"),
		Password: viper.GetString("db-config.password"),
	}
}