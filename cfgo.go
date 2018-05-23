package cfgo

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	envVar          = "GOENV"
	defaultFileType = "yaml"
	defaultEnv      = "default"
	defaultPath     = "./config"
	localEnv        = "local"
)

type Cfg struct {
	Path     string
	Scope    string
	FileType string
}

type Config interface {
	getEnv() string
	load(string) error
	merge(string) error
	setCoreUp()
	setDefaults()
}

var _ Config = &Cfg{}

// Load loads the configuration of a file named as the environment the current
// services is running on
func Load(c Config) error {
	c.setDefaults()
	c.setCoreUp()

	if err := c.load(defaultEnv); err != nil {
		return fmt.Errorf("error reading default configuration file:", err)
	}

	env := c.getEnv()
	if err := c.merge(env); err != nil {
		return fmt.Errorf("error reading %s configuration file:", env, err)
	}

	return nil
}

func Get(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func (c *Cfg) load(env string) error {
	viper.SetConfigName(env)
	viper.AddConfigPath(c.Path)
	return viper.ReadInConfig()
}

func (c *Cfg) merge(env string) error {
	viper.SetConfigName(env)
	viper.AddConfigPath(c.Path)
	return viper.MergeInConfig()
}

func (c *Cfg) setCoreUp() {
	viper.SetConfigType(c.FileType)
	viper.SetEnvPrefix(c.Scope)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func (c *Cfg) getEnv() string {
	env := os.Getenv(envVar)
	if env == "" {
		env = localEnv
	}
	return env
}

func (c *Cfg) setDefaults() {
	if c.Path == "" {
		c.Path = defaultPath
	}

	if c.FileType == "" {
		c.FileType = defaultFileType
	}
}
