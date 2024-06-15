package main

import (
	"fmt"
	"github.com/IvaCheMih/CrossRoads/generator/services"
	"github.com/IvaCheMih/CrossRoads/generator/traffic"
	"log"
	"math/rand"
	"time"
)

func main() {
	err := services.GetFromEnv()
	if err != nil {
		log.Panicln(err)
	}

	time.Sleep(1 * time.Second)

	for {
		timer := time.NewTimer(30 * time.Millisecond)

		trafficLightId := rand.Intn(4) + 1

		number := rand.Intn(2)

		fmt.Printf("Traffic light id: %d, number: %d\n", trafficLightId, number)

		err = traffic.SendOne(trafficLightId, number)
		if err != nil {
			log.Panicln(err)
		}

		<-timer.C
	}
}
