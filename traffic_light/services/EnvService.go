package services

import (
	"errors"
	"os"
	"strconv"
)

func GetFromEnv() (error, string, int) {
	var exist bool

	TRIFFICLIGHT_PORT, exist := os.LookupEnv("TRIFFICLIGHT_PORT")
	if !exist {
		return errors.New("TRIFFICLIGHT_PORT is not found"), "", 0
	}

	ID, exist := os.LookupEnv("ID")
	if !exist {
		return errors.New("ID is not found"), "", 0
	}

	id, err := strconv.Atoi(ID)
	if err != nil {
		return errors.New("ID is not int"), "", 0
	}

	return nil, TRIFFICLIGHT_PORT, id
}
