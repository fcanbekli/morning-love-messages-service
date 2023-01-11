package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	//_ = SendMessage("Canim babam")
	//	cfg := ParseConfigs("lover.json")

	// 3
	loc, err := time.LoadLocation("Turkey")
	if err != nil {
		// handle error
	}

	s := gocron.NewScheduler(loc)

	s.Every(1).Day().At("12:45").Do(func() {
		fmt.Println("Hello")
	})

	s.StartBlocking()
	//Parse Configs

	//Connect to Wp

	//Send intro message (if required)

	//Start cron job
}

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}
