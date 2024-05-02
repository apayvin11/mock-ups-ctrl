package main

import (
	"log"
	"time"

	"github.com/alex11prog/mock-ups-ctrl/model"
	"github.com/tbrandon/mbserver"
)

const confPath = "conf/config.toml"

func main() {
	conf := model.NewConfig(confPath)
	serv := mbserver.NewServer()
	defer serv.Close()
	if err := serv.ListenTCP(conf.ModbusAddr); err != nil {
		log.Fatalf("%v\n", err)
	}

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
