package controllers

import (
	"syscall"
	"time"

	"github.com/AllenDang/w32"
)

var (
	moduser32               = syscall.NewLazyDLL("user32.dll")
	procSetForegroundWindow = moduser32.NewProc("SetForegroundWindow")
	procFindWindowW         = moduser32.NewProc("FindWindowW")
)

const (
	KEYEVENTF_KEYDOWN     = 0 //key UP
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002 //key UP
	KEYEVENTF_UNICODE     = 0x0004
	KEYEVENTF_SCANCODE    = 0x0008 // scancode
)

type HWND uintptr

func sendkey(vk uint16) {
	var inputs []w32.INPUT
	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki: w32.KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KEYEVENTF_KEYDOWN,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki: w32.KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KEYEVENTF_KEYUP,
			Time:        0,
			DwExtraInfo: 0,
		},
	})
	w32.SendInput(inputs)
}

func sendkeys(str string) {
	var inputs []w32.INPUT
	for _, s := range str {
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         0,
				WScan:       uint16(s),
				DwFlags:     KEYEVENTF_KEYDOWN | KEYEVENTF_UNICODE,
				Time:        0,
				DwExtraInfo: 0,
			},
		})

		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         0,
				WScan:       uint16(s),
				DwFlags:     KEYEVENTF_KEYUP | KEYEVENTF_UNICODE,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	w32.SendInput(inputs)
}

func pressAndHold(key byte, holdforMs time.Duration) {
	var inputs []w32.INPUT
	inputs = []w32.INPUT{
		{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:     uint16(key),
				DwFlags: KEYEVENTF_KEYDOWN | KEYEVENTF_UNICODE,
			},
		},
	}

	w32.SendInput(inputs)

	time.Sleep(holdforMs * time.Millisecond)

	inputs = []w32.INPUT{
		{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:     uint16(key),
				DwFlags: KEYEVENTF_KEYUP | KEYEVENTF_UNICODE,
			},
		},
	}

	w32.SendInput(inputs)

}
