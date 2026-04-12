//go:build xiao_esp32c3

package main

import (
	"machine"
)

const (
	led    = machine.D0
	button = machine.D1
)

func buttonPushed() bool {
	return !button.Get()
}
