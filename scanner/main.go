package main

import (
	"machine"

	"fmt"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/espradio"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
)

var (
	terminal *tinyterm.Terminal
	font     = &proggy.TinySZ8pt7b
)

func main() {
	display := newSSD1306Display()
	terminal = tinyterm.NewTerminal(display)

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})

	printMessage("initializing radio...")
	err := espradio.Enable(espradio.Config{})
	if err != nil {
		printMessage("could not enable radio: " + err.Error())
		return
	}

	printMessage("starting radio...")
	err = espradio.Start()
	if err != nil {
		printMessage("could not start radio: " + err.Error())
		return
	}

	for {
		printMessage("scanning WiFi...")
		aps, err := espradio.Scan()
		if err != nil {
			printMessage("could not scan wifi: " + err.Error())
			return
		}

		for _, ap := range aps {
			printMessage(ap.SSID + " " + fmt.Sprint(ap.RSSI))
		}

		terminal.Display()
		time.Sleep(5 * time.Second)
	}
}

func newSSD1306Display() *ssd1306.Device {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
		SDA:       machine.SDA_PIN,
		SCL:       machine.SCL_PIN,
	})
	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32, // or ssd1306.Address
		Width:   128,
		Height:  64,
	})
	return display
}

func printMessage(msg string) {
	println(msg)
	fmt.Fprintf(terminal, "\n%s", msg)
}
