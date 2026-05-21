package main

import (
	"log"
	"time"

	"github.com/SHAIK14/pulsemon/checker"
	"github.com/SHAIK14/pulsemon/config"
	"github.com/SHAIK14/pulsemon/logger"
	"github.com/SHAIK14/pulsemon/notifier"
)

func main() {
	config.LoadEnv(".env")

	data, err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}
	resultC := make(chan checker.Result, len(data.Services))
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		for _, service := range data.Services {
			go func(svc config.Service) {
				resCheck := checker.Check(svc)
				resultC <- resCheck
			}(service)

		}
		for i := 0; i < len(data.Services); i++ {
			result := <-resultC
			if !result.IsUp {
				notifier.Notify(result)
			}
			logger.Log(result)

		}
	}

}
