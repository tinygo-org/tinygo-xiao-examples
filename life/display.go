package main

import (
	"errors"
	"image/color"
)

var (
	errBufferSize = errors.New("invalid size buffer")
)

type DisplayBuffer struct {
	buffer []byte
	width  int16
	height int16
}

func NewDisplayBuffer(width, height int16) *DisplayBuffer {
	return &DisplayBuffer{
		buffer: make([]byte, width*height/8),
		width:  width,
		height: height,
	}
}

func (d DisplayBuffer) Size() (x, y int16) {
	return d.width, d.height
}

func (d *DisplayBuffer) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || x >= d.width || y < 0 || y >= d.height {
		return
	}
	byteIndex := x + (y/8)*d.width
	if c.R != 0 || c.G != 0 || c.B != 0 {
		d.buffer[byteIndex] |= 1 << uint8(y%8)
	} else {
		d.buffer[byteIndex] &^= 1 << uint8(y%8)
	}
}

func (d DisplayBuffer) Display() error {
	return nil
}

func (d *DisplayBuffer) GetPixel(x int16, y int16) bool {
	if x < 0 || x >= d.width || y < 0 || y >= d.height {
		return false
	}
	byteIndex := x + (y/8)*d.width
	return (d.buffer[byteIndex] >> uint8(y%8) & 0x1) == 1
}

func (d *DisplayBuffer) SetBuffer(buffer []byte) error {
	if len(buffer) != len(d.buffer) {
		return errBufferSize
	}
	for i := 0; i < len(d.buffer); i++ {
		d.buffer[i] = buffer[i]
	}
	return nil
}

func (d *DisplayBuffer) GetBuffer() []byte {
	return d.buffer
}
