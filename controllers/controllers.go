package controllers

import "strings"

// KeyInput represents a valid input sent by someone
type KeyInput struct {
	Sender string
	Key    string
}

// Controller is an interface representing a way of controling the game
type Controller interface {
	Press(button string)
}

var buttons = []string{"a", "b", "start", "select", "up", "down", "left", "right"}

// IsValidButton checks if the given key name is a valid button
func IsValidButton(key string) bool {
	for _, b := range buttons {
		if b == strings.ToLower(key) {
			return true
		}
	}
	return false
}
