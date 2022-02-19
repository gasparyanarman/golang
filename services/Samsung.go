package services

import "fmt"

type Samsung struct {
	Cameras int
	Cpu     float64
	Screen  float64
}

func (s Samsung) Call() {
	fmt.Println("Samsung is calling")
}

func (s Samsung) Message() {
	fmt.Println("Samsung is texting")
}
