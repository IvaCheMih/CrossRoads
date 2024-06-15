package traffic

import (
	"fmt"
	"github.com/IvaCheMih/CrossRoads/traffic_light/traffic/dto"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"sync"
)

type TrafficHandlers struct {
	Id    int
	Queue int
	Light bool
	Mu    sync.Mutex
}

func CreateTrafficHandlers(id int) TrafficHandlers {
	return TrafficHandlers{
		Id:    id,
		Queue: 0,
		Light: false,
		Mu:    sync.Mutex{},
	}
}

func (t *TrafficHandlers) Collect(c *fiber.Ctx) error {

	t.Mu.Lock()

	var response = dto.CollectResponse{
		t.Queue,
	}

	fmt.Println("Collect: ", t.Queue)

	t.Mu.Unlock()

	return c.JSON(response)
}

func (t *TrafficHandlers) AddInQueue(c *fiber.Ctx) error {

	request, err := dto.GetAddInQueueRequest(c)
	if err != nil {
		log.Println(err)
	}

	t.Mu.Lock()

	t.Queue += request.Number

	fmt.Println("AddInQueue: ", request.Number)

	t.Mu.Unlock()

	return c.JSON(request.Number)
}

func (t *TrafficHandlers) Execute(c *fiber.Ctx) error {
	request, err := dto.GetCommandRequest(c)
	if err != nil {
		log.Println(err)
	}

	t.Mu.Lock()

	t.Light = request.CMD
	t.Queue -= request.Num

	fmt.Println("Execute: ", request.Num)

	t.Mu.Unlock()

	var response = dto.CommandResponse{
		Message: "[*] TrafficLight id: " + strconv.Itoa(t.Id) + ": command received",
	}

	return c.JSON(response)
}
