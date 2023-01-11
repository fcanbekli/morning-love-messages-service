package main

import (
	"github.com/go-co-op/gocron"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"os"
	"time"
)

func main() {
	//Parse Configs
	cfg := ParseConfigs(os.Args[1])
	//Connect to Wp
	waConnect()
	//Send intro message (if required)
	if cfg.IsIntroMessage {
		SendMessage(cfg.TargetPhone, cfg.IntroMessage)
	}

	//Start cron job
	loc, err := time.LoadLocation(cfg.Country)
	if err != nil {
		panic(err)
	}

	s := gocron.NewScheduler(loc)

	s.Every(1).Day().At(cfg.MorningMessageHour).Do(func() {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		index := r.Intn(len(cfg.Messages))
		SendMessage(cfg.TargetPhone, cfg.Messages[index])
	})
	s.StartBlocking()
}