package main

import (
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
)

var (
	terminal *tinyterm.Terminal
	font     = &proggy.TinySZ8pt7b
)

func initDisplay() {
	display := newSSD1306Display()
	terminal = tinyterm.NewTerminal(display)

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})
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
	terminal.Write([]byte("\n" + msg))
	terminal.Display()
}
