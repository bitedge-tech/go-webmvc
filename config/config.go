package config

import (
	"log"

	"github.com/spf13/viper"
)

type LogFileConfig struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type LogConfig struct {
	Level  string        `mapstructure:"level"`
	Format string        `mapstructure:"format"`
	Output string        `mapstructure:"output"`
	File   LogFileConfig `mapstructure:"file"`
}

type NatsConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type JwtConfig struct {
	Secret string `mapstructure:"secret"`
}

type Config struct {
	App struct {
		Port     string `mapstructure:"port"`
		Env      string `mapstructure:"env"`
		Timezone string `mapstructure:"timezone"`
	} `mapstructure:"app"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`

	Nats NatsConfig `mapstructure:"nats"`

	Log LogConfig `mapstructure:"log"`

	Jwt JwtConfig `mapstructure:"jwt"`

	Minio struct {
		Endpoint   string `mapstructure:"endpoint"`
		AccessKey  string `mapstructure:"access_key"`
		SecretKey  string `mapstructure:"secret_key"`
		BucketName string `mapstructure:"bucket_name"`
	}
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
