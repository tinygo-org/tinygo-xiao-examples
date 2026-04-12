package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

var (
	displayBuffer *DisplayBuffer
	lifegame      *LifeGame

	textWhite = color.RGBA{255, 255, 255, 255}
	textBlack = color.RGBA{0, 0, 0, 255}
)

func main() {
	display := newSSD1306Display()
	display.ClearDisplay()
	displayBuffer = NewDisplayBuffer(display.Size())

	var err error
	lifegame, err = NewLifeGame(64, 32)
	if err != nil {
		return
	}
	lifegame.InitRandom()

	for {
		playLife()

		display.SetBuffer(displayBuffer.GetBuffer())
		display.Display()
		time.Sleep(5 * time.Millisecond)
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

func playLife() {
	lifegame.Update()
	cells := lifegame.GetCells()
	for y := range cells {
		for x := range cells[y] {
			color := textBlack
			if cells[y][x] {
				color = textWhite
			}

			displayBuffer.SetPixel(int16(x)*2+0, int16(y)*2+0, color)
			displayBuffer.SetPixel(int16(x)*2+0, int16(y)*2+1, color)
			displayBuffer.SetPixel(int16(x)*2+1, int16(y)*2+0, color)
			displayBuffer.SetPixel(int16(x)*2+1, int16(y)*2+1, color)
		}
	}
}
