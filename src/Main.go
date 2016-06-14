package main

import "fmt"
//import "log"

func main() {
	em := Employee{"Emily", 26, female}
	fmt.Println("Name: " , em.name, em.age, em.getSex())

	bum := Business{"Emily's Business", nil}

	fmt.Println(bum.getName())
}

type Employee struct {
	name string
	age int
	sexNum Sex
}

type Sex int

const (
	male Sex = iota
	female
)

var sexes = [...]string {
	"male",
	"female",
}

func (day Employee) getSex() string {
	return sexes[day.sexNum]
}

type Business struct {
	name string
	employees []Employee
}

func (b Business) getName() (name string) {
	return b.name
}

