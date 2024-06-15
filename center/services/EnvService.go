package services

import (
	"errors"
	"os"
)

func GetFromEnv() (error, string) {
	var exist bool

	TRIFFICLIGHT_URL, exist := os.LookupEnv("TRIFFICLIGHT_URL")
	if !exist {
		return errors.New("TRIFFICLIGHT_URL is not found"), ""
	}

	return nil, TRIFFICLIGHT_URL
}
