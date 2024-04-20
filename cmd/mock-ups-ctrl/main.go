package main

import (
	"log"
	"time"

	"github.com/alex11prog/mock-ups-ctrl/internal/app/mockctrl"
	"github.com/tbrandon/mbserver"
)

const confPath = "conf/config.toml"

func main() {
	conf := mockctrl.NewConfig(confPath)
	serv := mbserver.NewServer()
	mockctrl.InitMockModbusParams(serv)
	err := serv.ListenTCP(conf.ModbusServerAddrPort)
	if err != nil {
		log.Printf("%v\n", err)
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
