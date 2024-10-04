package main

import "fmt"

type Vechile struct {
	Engine string
}
type Car struct {
	*Vechile
}

func NewVechile() *Vechile {
	return &Vechile{Engine: "am a car engine"}

}

type keke struct {
	*Vechile
}

func NewKeke() *keke {
	return &keke{Vechile: &Vechile{Engine: "i am a keke engine"}}
}
func main() {
	car := NewVechile()
	fmt.Println(car.Engine)
	keke := NewKeke()
	fmt.Println(keke.Engine)
}
