package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Name               string
	TargetPhone        string
	Messages           []string
	Country            string
	IsIntroMessage     bool
	IntroMessage       string
	MorningMessageHour string
}

func ParseConfigs(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Config file is missing")
		panic(err)
	}
	var configs Config
	err = json.Unmarshal(data, &configs)
	if err != nil {
		panic(err)
	}
	return configs
}
