package traffic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/IvaCheMih/CrossRoads/generator/services"
	"github.com/IvaCheMih/CrossRoads/generator/traffic/dto"
	"io"
	"net/http"
)

func SendOne(TL_id int, number int) error {
	var sendOneBody = dto.SendOneRequest{
		Number: number,
	}

	body, err := json.Marshal(sendOneBody)
	if err != nil {
		return err
	}

	url := services.Urls[TL_id] + "Quantity/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
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
		fmt.Println("client: could not read response body: %s\n", err)
	}
	fmt.Println(string(resBody))

	return err
}
