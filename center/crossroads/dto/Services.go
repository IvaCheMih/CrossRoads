package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func GetRequestControl(c *fiber.Ctx) (ControlRequest, error) {
	body := c.Body()

	var request ControlRequest

	err := json.Unmarshal(body, &request)

	return request, err
}
