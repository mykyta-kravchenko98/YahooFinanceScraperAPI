package configs

import (
	"github.com/spf13/viper"
)

type YahooFinanceConfig struct {
	URL string `mapstructure:"url"`
}

type PostgresDBConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"sslmode"`
}

type ServerConfig struct {
	Port      string `mapstructure:"port"`
	SyncTime1 string `mapstructure:"syncTime1"`
	SyncTime2 string `mapstructure:"syncTime2"`
}
type Config struct {
	PostgresDB   PostgresDBConfig   `mapstructure:"postgresDB"`
	YahooFinance YahooFinanceConfig `mapstructure:"yahooFinance"`
	Server       ServerConfig       `mapstructure:"server"`
}

var (
	vp     *viper.Viper
	config *Config
)

func LoadConfigs(env string) (*Config, error) {
	vp = viper.New()

	vp.SetConfigType("json")
	vp.SetConfigName(env)
	vp.AddConfigPath("../configs/")
	vp.AddConfigPath("../../configs/")
	vp.AddConfigPath("configs/")

	err := vp.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return &Config{}, err
	}

	return config, err
}

func GetConfig() *Config {
	return config
}
