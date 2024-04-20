package mockctrl

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ModbusServerAddrPort string `toml:"modbus_server_addr_port"`
}

func NewConfig(configPath string) *Config {
	conf := &Config{}
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatal("toml decode file config error! ", err)
	}
	return conf
}
