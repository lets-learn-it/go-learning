package main

import (
	"log"

	"github.com/lets-learn-it/gin-auth-scope/api"
	"github.com/lets-learn-it/gin-auth-scope/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("something wrong while creating new server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
