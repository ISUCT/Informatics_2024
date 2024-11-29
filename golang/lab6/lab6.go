package lab6

import (
	"fmt"
)

type TV interface {
	GetChannel() int
	SetChannel(channel int)
}

type Television struct {
	channel int
}

func (t *Television) GetChannel() int {
	return t.channel
}

func (t *Television) SetChannel(channel int) {
	if channel > 0 {
		t.channel = channel
		fmt.Printf("Channel set to %d\n", channel)
	} else {
		fmt.Println("Invalid channel number")
	}
}

func (t *Television) Display() {
	fmt.Println("╔════════════════╗")
	fmt.Println("║                ║")
	fmt.Printf("║   Channel %2d   ║\n", t.GetChannel())
	fmt.Println("║                ║")
	fmt.Println("╚════════════════╝")
}

func RunLab6() {
	tv := &Television{}
	tv.SetChannel(5)
	tv.Display()

	tv.SetChannel(12)
	tv.Display()
}
