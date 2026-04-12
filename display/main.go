package main

import (
	"machine"

	"runtime"

	"image/color"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func main() {

	display := newSSD1306Display()
	display.ClearDisplay()

	w, h := display.Size()
	x := int16(0)
	y := int16(0)
	deltaX := int16(1)
	deltaY := int16(1)

	traceTime := time.Now().UnixMilli() + 1000
	frames := 0
	ms := runtime.MemStats{}

	for {
		pixel := display.GetPixel(x, y)
		c := color.RGBA{255, 255, 255, 255}
		if pixel {
			c = color.RGBA{0, 0, 0, 255}
		}
		display.SetPixel(x, y, c)
		display.Display()

		x += deltaX
		y += deltaY

		if x == 0 || x == w-1 {
			deltaX = -deltaX
		}

		if y == 0 || y == h-1 {
			deltaY = -deltaY
		}

		frames++
		now := time.Now().UnixMilli()
		if now >= traceTime {
			runtime.ReadMemStats(&ms)
			println("TS", now, "| FPS", frames, "| HeapInuse", ms.HeapInuse)
			traceTime = now + 1000
			frames = 0
		}
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
