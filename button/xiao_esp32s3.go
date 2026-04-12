//go:build xiao_esp32s3

package main

import (
	"machine"
)

const (
	led    = machine.LED
	button = machine.D1
)

func buttonPushed() bool {
	return button.Get()
}
