package main

import (
	"github.com/IvaCheMih/CrossRoads/traffic_light/services"
	"github.com/IvaCheMih/CrossRoads/traffic_light/traffic"
	"github.com/gofiber/fiber/v2"
	"log"
)

var trafficHandlers traffic.TrafficHandlers
var port string
var id int

func Init() {

	var err error

	err, port, id = services.GetFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	trafficHandlers = traffic.CreateTrafficHandlers(id)

}

func main() {

	Init()

	server := fiber.New()

	server.Get("/Collect", trafficHandlers.Collect)

	server.Post("/Command", trafficHandlers.Execute)

	server.Post("/Quantity", trafficHandlers.AddInQueue)

	if err := server.Listen(port); err != nil {
		log.Fatal(err)
	}
}
