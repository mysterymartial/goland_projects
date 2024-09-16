package main

import "fmt"

type crew struct {
	Driver
}
type Driver interface {
	Drive()
}
type Truck struct{}

func (t Truck) Drive() {
	fmt.Println("i am a Truck driver")

}

type Tesla struct{}

func (t Tesla) Drive() {
	fmt.Println("i am a Tesla driver")

}

type Bike struct{}

func (b Bike) Drive() {
	fmt.Println("i am a Bike rider")

}
func main() {
	driver1 := Truck{}
	driver2 := Tesla{}
	driver3 := Bike{}

	team1 := crew{Driver: driver1}
	team2 := crew{Driver: driver2}
	team3 := crew{Driver: driver3}

	team1.Drive()
	team2.Drive()
	team3.Drive()

}
