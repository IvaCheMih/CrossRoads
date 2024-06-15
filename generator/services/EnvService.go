package services

import (
	"errors"
	"os"
)

var Urls = map[int]string{}

func GetFromEnv() error {
	var exist bool

	Urls[1], exist = os.LookupEnv("TRIFFICLIGHT_1")
	if !exist {
		return errors.New("TRIFFICLIGHT_1 is not found")
	}

	Urls[2], exist = os.LookupEnv("TRIFFICLIGHT_2")
	if !exist {
		return errors.New("TRIFFICLIGHT_2 is not found")
	}

	Urls[3], exist = os.LookupEnv("TRIFFICLIGHT_3")
	if !exist {
		return errors.New("TRIFFICLIGHT_3 is not found")
	}

	Urls[4], exist = os.LookupEnv("TRIFFICLIGHT_4")
	if !exist {
		return errors.New("TRIFFICLIGHT_4 is not found")
	}

	return nil
}
