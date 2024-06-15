package crossroads

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/IvaCheMih/CrossRoads/center/crossroads/dto"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type CrossRoads struct {
	TrafficLight_1 TrafficLight
	TrafficLight_2 TrafficLight
	TrafficLight_3 TrafficLight
	TrafficLight_4 TrafficLight
	TimeRerOne     time.Duration `json:"time_rerOne"`
	URL            string        `json:"url"`
}

type TrafficLight struct {
	Id              int `json:"id"`
	QuantityInQueue int `json:"quantity_in_queue"`

	URL string `json:"url"`
}

func CreateCrossroads(url string) CrossRoads {
	return CrossRoads{
		TrafficLight_1: TrafficLight{
			Id:              1,
			QuantityInQueue: 0,
		},
		TrafficLight_2: TrafficLight{
			Id:              2,
			QuantityInQueue: 0,
		},
		TrafficLight_3: TrafficLight{
			Id:              3,
			QuantityInQueue: 0,
		},
		TrafficLight_4: TrafficLight{
			Id:              4,
			QuantityInQueue: 0,
		},

		TimeRerOne: 100 * time.Millisecond,
		URL:        url,
	}
}

func (c *CrossRoads) CollectInformationFromTLs() error {
	var er error
	var errChan = make(chan error, 4)

	go func() {
		er = <-errChan
	}()

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		err, num := c.collectInfoFromOne(c.TrafficLight_1.URL, c.TrafficLight_1.Id)
		if err != nil {
			errChan <- err
		}
		c.TrafficLight_1.QuantityInQueue += num

		wg.Done()
	}()

	go func() {
		err, num := c.collectInfoFromOne(c.TrafficLight_2.URL, c.TrafficLight_2.Id)
		if err != nil {
			errChan <- err
		}

		c.TrafficLight_2.QuantityInQueue += num
		wg.Done()
	}()

	go func() {
		err, num := c.collectInfoFromOne(c.TrafficLight_3.URL, c.TrafficLight_3.Id)
		if err != nil {
			errChan <- err
		}

		c.TrafficLight_3.QuantityInQueue += num

		wg.Done()
	}()

	go func() {
		err, num := c.collectInfoFromOne(c.TrafficLight_4.URL, c.TrafficLight_4.Id)
		if err != nil {
			errChan <- err
		}

		c.TrafficLight_4.QuantityInQueue += num

		wg.Done()
	}()

	wg.Wait()

	return er

}

func (c *CrossRoads) collectInfoFromOne(url string, id int) (error, int) {

	response := fiber.Get(c.URL + strconv.Itoa(id) + ":808" + strconv.Itoa(id) + "/Collect/")

	_, body, errs := response.Bytes()
	if errs != nil {
		return errs[0], 0
	}

	var quantity = dto.CollectOneResponse{}

	err := json.Unmarshal(body, &quantity)
	if err != nil {
		return err, 0
	}

	fmt.Println("collect from ", id, " quantity: ", quantity.Quantity)

	return err, quantity.Quantity
}

func (c *CrossRoads) FindOptimal() (int, []bool) {

	time1 := min(c.TrafficLight_1.QuantityInQueue, c.TrafficLight_3.QuantityInQueue)

	time2 := min(c.TrafficLight_2.QuantityInQueue, c.TrafficLight_4.QuantityInQueue)

	times := max(time1, time2)

	if times == time1 {
		return times, []bool{true, false, true, false}
	}

	return times, []bool{false, true, false, true}
}

func (c *CrossRoads) Execute(times int, commands []bool) error {

	var er error
	var errChan = make(chan error, 4)

	go func() {
		er = <-errChan
	}()

	wg := sync.WaitGroup{}
	wg.Add(4)

	for i, cmd := range commands {
		go func(i int, cmd bool) {
			num := 0

			if cmd {
				num = times
			}

			err := c.sendOneCommand(i+1, cmd, num)
			if err != nil {
				errChan <- err
			}

			wg.Done()
		}(i, cmd)
	}

	wg.Wait()

	c.dellFromQueue(times, commands)

	return er
}

func (c *CrossRoads) sendOneCommand(i int, cmd bool, times int) error {

	var sendOneBody = dto.SendOneBody{
		CMD: cmd,
		Num: times,
	}

	body, err := json.Marshal(sendOneBody)
	if err != nil {
		return err
	}

	fmt.Println("send ", sendOneBody.CMD, sendOneBody.Num)

	request, err := http.NewRequest(http.MethodPost, c.URL+strconv.Itoa(i)+":808"+strconv.Itoa(i)+"/Command/", bytes.NewReader(body))
	if err != nil {
		return err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println()
	log.Println(string(resBody))
	fmt.Println()

	return nil
}

func (c *CrossRoads) dellFromQueue(times int, commands []bool) {
	if commands[0] {
		c.TrafficLight_1.QuantityInQueue -= times
		c.TrafficLight_3.QuantityInQueue -= times
	} else {
		c.TrafficLight_2.QuantityInQueue -= times
		c.TrafficLight_4.QuantityInQueue -= times
	}

}
