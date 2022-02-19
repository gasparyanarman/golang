package services

import "fmt"

type Iphone struct {
	Cameras int
	Cpu     float64
	Screen  float64
}

func (i Iphone) Call() {
	fmt.Println("Iphone is calling")
}

func (i Iphone) Message() {
	fmt.Println("Iphone is texting")
}

func (i Iphone) Capture() {
	fmt.Println("I am taking a photo of you")
}
