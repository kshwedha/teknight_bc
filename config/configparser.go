package config

import (
	"fmt"

	"github.com/bigkevmcd/go-configparser"
)

func ParseConfig(fn string) *configparser.ConfigParser {
	p, err := configparser.NewConfigParserFromFile(fn)
	if err != nil {
		fmt.Print("!! Error reading a file.", err)
	}
	return p
}

func getConfigSections() []string {
	Config := ParseConfig(fileName())
	return Config.Sections()
}

func GetConfigVal(section string, key string) string {
	Config := ParseConfig(fileName())
	v, err := Config.Get(section, key)
	if err != nil {
		fmt.Print("!! Error reading a file/section/key.", err)
	}
	return v

}

func SetConfig() {
	Config := ParseConfig(fileName())
	err := Config.Set("section", "newoption", "value")
	if err != nil {
		fmt.Print("!! Error reading a file.", err)
	}
}

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func fileName() string {
	return "config/config.cfg"
}
