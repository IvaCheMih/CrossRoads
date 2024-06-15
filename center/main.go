package main

import (
	"github.com/IvaCheMih/CrossRoads/center/crossroads"
	"github.com/IvaCheMih/CrossRoads/center/services"
	"log"
	"time"
)

var crossRoads crossroads.CrossRoads

func Init() {
	err, url := services.GetFromEnv()
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(1 * time.Second)

	crossRoads = crossroads.CreateCrossroads(url)

	err = crossroads.StartCrossRoads(crossRoads)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	Init()
}
