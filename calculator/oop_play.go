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
func main() {
	animal1 := Dog{}
	animal2 := Cat{}
	animal3 := Duck{}

	pet1 := Animal{Pet: animal1}
	pet2 := Animal{Pet: animal2}
	pet3 := Animal{Pet: animal3}

	pet1.Sound()
	pet2.Sound()
	pet3.Sound()

}
