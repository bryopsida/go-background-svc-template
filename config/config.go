package config

import (
	"path"

	"github.com/bryopsida/go-background-svc-template/interfaces"
	"github.com/spf13/viper"
)

type viperConfig struct {
	viper *viper.Viper
}

func NewViperConfig() interfaces.IConfig {
	config := viperConfig{viper: viper.New()}
	config.initialize()
	config.setDefaults()
	return &config
}

func (c *viperConfig) setDefaults() {
	c.viper.SetDefault("database.path", path.Join("data", "db"))
}

func (c *viperConfig) initialize() {
	c.viper.SetConfigName("config")
	c.viper.SetConfigType("yaml")
	c.viper.AddConfigPath(".")
}
func (c *viperConfig) GetDatabasePath() string {
	return c.viper.GetString("database.path")
}
