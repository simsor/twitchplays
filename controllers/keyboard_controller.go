package controllers

import (
	"fmt"
	"strings"
)

// KeyboardController sends keyboard events
type KeyboardController struct {
}

var keys = map[string]byte{
	"a":      0x41,
	"b":      0x42,
	"start":  0x0D,
	"select": 0x53,
	"left":   0x25,
	"right":  0x27,
	"up":     0x26,
	"down":   0x28,
}

// NewKeyboardController creates a new keyboard controller
func NewKeyboardController() (k *KeyboardController, err error) {
	return &KeyboardController{}, nil
}

// Press simulates a button press corresponding to the given "key"
func (c KeyboardController) Press(button string) {
	value, ok := keys[strings.ToLower(button)]
	if !ok {
		fmt.Printf("Button %s doesn't seem to be valid\n", button)
	}

	pressAndHold(value, 200)
}
