package lab6

import (
	"fmt"
)

type TV struct {
	Channel int
}

func NewTV(channel int) *TV {
	tv := new(TV)
	tv.Channel = channel
	return tv
}

func (t *TV) GetChannel() int {
	return t.Channel
}

func (t *TV) SetChannel(channel int) {
	if channel > 0 {
		t.Channel = channel
		fmt.Printf("Channel set to %d\n", channel)
	} else {
		fmt.Println("Invalid channel number")
	}
}

func (t *TV) Display() {
	fmt.Println("╔════════════════╗")
	fmt.Println("║                ║")
	fmt.Printf("║   Channel %2d   ║\n", t.GetChannel())
	fmt.Println("║                ║")
	fmt.Println("╚════════════════╝")
	fmt.Println("        ||        ")
	fmt.Println("       ====        ")
}

func Lab6() {
	samsung := NewTV(2)
	samsung.SetChannel(5)
	samsung.Display()
	samsung.SetChannel(12)
	samsung.Display()
}
