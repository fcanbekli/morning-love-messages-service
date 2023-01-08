package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	//_ = SendMessage("Canim babam")
	cfg := ParseConfigs("lover.json")
	fmt.Println(cfg.Name)
	fmt.Println(cfg.TargetPhone)
	fmt.Println(cfg.Messages)

}

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	// 3
	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(1).Seconds().Do(func() {
		hello("John Doe")
	})

	// 5
	s.StartBlocking()
}
