package gomydb

import (
	"errors"
	"github.com/msbranco/goconfig"
)

var (
	ErrNullString   = errors.New("null string")
	ErrConfigParser = errors.New("config file content syntax invalid")
)

type Config struct {
	config *goconfig.ConfigFile
}

func NewConfig(fileconfig string) *Config {
	if fileconfig == "" {
		panic(ErrNullString)
	}
	c, err := goconfig.ReadConfigFile(fileconfig)
	if err != nil {
		panic(ErrConfigParser)
	}
	return &Config{c}
}

func (c *Config) Get(section string, option string) (value string) {
	value, err := c.config.GetString(section, option)
	if err != nil {
		return ""
	}
	return value
}

// etc
//c, err := configfile.ReadConfigFile("config.cfg");
//// result is string :http://www.example.com/some/path"
//c.GetString("service-1", "url");
//c.GetInt64("service-1", "maxclients"); // result is int 200
//c.GetBool("service-1", "delegation"); // result is bool true
//
//// result is string "This is a multi-line\nentry"
//c.GetString("service-1", "comments");
