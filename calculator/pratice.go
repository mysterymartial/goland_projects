package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	creditCard creditcard
}
type creditcard struct {
	Bank   string
	Number string
	pin    string
	cvc    string
}

func Newcreditcard() creditcard {
	return creditcard{
		Bank:   "Gtb",
		Number: "1000-4000-6789",
		pin:    "0000",
		cvc:    "123456",
	}
}
func NewPerson(name string, age int) Person {
	return Person{
		Name:       name,
		Age:        age,
		creditCard: Newcreditcard(),
	}

}

type Stringer interface {
	String() string
}
type Community struct {
	Name     string
	Interest string
}

func (c Community) string() string {
	return fmt.Sprintf("Name %s, interest %s",
		c.Name,
		c.Interest,
	)
}

type Hellower interface {
	Hello()
}
type Bayer interface {
	Bay()
}
type HelloBay struct{}

func (HelloBay) Hello() {
	fmt.Println("Hello Bay")

}
func (HelloBay) Bay() {
	fmt.Println("Hello Babo")

}

var _ Hellower = (*HelloBay)(nil)
var _ Bayer = (*HelloBay)(nil)
var try *string

func main() {
	community := Community{
		Name:     "Go lang boys",
		Interest: "writing go codes",
	}

	fmt.Println(community)

}
