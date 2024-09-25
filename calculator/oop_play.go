package main

import "fmt"

type Animal struct {
	Pet
}
type Pet interface {
	Sound()
}
type Dog struct{}

func (d Dog) Sound() {
	fmt.Println("bark")
}

type Cat struct{}

func (c Cat) Sound() {
	fmt.Println("meawon")
}

type Duck struct{}

func (d Duck) Sound() {
	fmt.Println("quack")
}

type Lion struct{}

func (l Lion) Sound() {
	fmt.Println("rows")
}

type Baby struct{}

func (b Baby) Sound() {
	fmt.Println("i cry")
}
func main() {
	animal1 := Dog{}
	animal2 := Cat{}
	animal3 := Duck{}
	animal4 := Lion{}
	animal5 := Baby{}

	pet1 := Animal{Pet: animal1}
	pet2 := Animal{Pet: animal2}
	pet3 := Animal{Pet: animal3}
	pet4 := Animal{Pet: animal4}

	pet1.Sound()
	pet2.Sound()
	pet3.Sound()
	pet4.Sound()

}
