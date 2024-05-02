package model

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ModbusAddr string `toml:"modbus_addr"`
}

func NewConfig(configPath string) *Config {
	conf := &Config{}
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatal("toml decode file config error! ", err)
	}
	return conf
}
