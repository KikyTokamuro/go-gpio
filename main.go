package main

import (
	"log"

	"github.com/KikyTokamuro/go-gpio/gpio"
)

func main() {
	pin := 26

	status := gpio.Setup()
	if err := gpio.SetupError(status); err != nil {
		log.Fatal(err)
	}

	gpio.SetupGPIO(pin, gpio.OUTPUT, gpio.PUD_OFF)
	gpio.SetupGPIO(pin, gpio.INPUT, gpio.PUD_OFF)
}
