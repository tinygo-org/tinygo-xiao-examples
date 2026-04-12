package main

import (
	"machine"
	"time"
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if buttonPushed() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
