package main

import (
	"log"
	config "test-module/internal/config"
	orderservice "test-module/internal/service/order"
	psqlrepo "test-module/internal/storage/psql"
	nutshendler "test-module/internal/transport/nutsstreaming/hendler"
	"test-module/internal/transport/rest"
	"test-module/pkg/localcache"

	"github.com/nats-io/stan.go"
)

func main() {
	config := config.Config{}

	err := config.LoadConfig("configs/config.yml")
	if err != nil {
		log.Printf("Config loading error   #%v ", err)
		return
	}

	rep := psqlrepo.Repository{}
	err = rep.Init(config)
	if err != nil {
		log.Printf("Repository loading error   #%v ", err)
	}
	lc := localcache.NewCache()
	err = orderservice.Create(&lc, &rep)
	if err != nil {
		log.Printf("Order service init error   #%v ", err)
	}

	connection, err := stan.Connect(config.NutsCluster, config.NutsClient, stan.NatsURL("nats://"+config.NutsHost+":4222"))
	if err == nil {
		_, err = connection.Subscribe(config.NutsChenel, nutshendler.HandleOrder, stan.StartWithLastReceived())
		if err != nil {
			log.Printf("Order service init error   #%v ", err)
		}
	} else {
		log.Printf("Order service init error   #%v ", err)
	}

	rest.RegisterHendlers()
	//TODO: Make it in config
	rest.StartServer(8080)
}
