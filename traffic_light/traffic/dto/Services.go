package dto

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAddInQueueRequest(c *fiber.Ctx) (AddInQueueRequest, error) {
	body := c.Body()

	var request = AddInQueueRequest{}

	err := json.Unmarshal(body, &request)
	if err != nil {
		return AddInQueueRequest{}, err
	}

	fmt.Println(request)

	return request, err
}

func GetCommandRequest(c *fiber.Ctx) (CommandRequest, error) {
	body := c.Body()

	var request = CommandRequest{}

	err := json.Unmarshal(body, &request)
	if err != nil {
		return CommandRequest{}, err
	}

	return request, err
}
