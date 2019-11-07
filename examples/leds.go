package main

import (
	"fmt"
	"log"

	leds "github.com/hodgesds/goleds"
)

func main() {
	leds, err := leds.LEDs()
	if err != nil {
		log.Fatal(err)
	}

	for _, led := range leds {
		triggers, err := led.Triggers()
		if err != nil {
			log.Fatal(err)
		}
		for _, trigger := range triggers {
			fmt.Printf("%+v\n", trigger)
		}
	}
}
