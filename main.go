package main

import (
	"fmt"
	"log"

	"github.com/SHAIK14/pulsemon/config"
)

func main() {
	data, err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data.Services {
		fmt.Println(d.Name, d.URL)
	}

}
