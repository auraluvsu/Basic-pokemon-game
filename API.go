package main

import (
	"log"
	"math/rand"
	"time"
)

type Animal interface {
	Says() string
	numOfLegs() int
	Habitat() string
}

type Cat struct {
	Name   string
	Breed  string
	Colour string
}

type Dog struct {
	Name  string
	Breed string
	Alive bool
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func CalculateValue(intChan chan int) {
	randNum := RandomNumber(100)
	intChan <- randNum
}
func main() {
	cat := Cat{
		Name:   "Sally",
		Breed:  "Persian",
		Colour: "Blonde",
	}
	dog := Dog{
		Name:  "Zeus",
		Breed: "German Shepherd",
		Alive: true,
	}
	PrintInfo(&cat)
	PrintInfo(&dog)

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and lives in the", a.Habitat())
}

func (c *Cat) Habitat() string {
	return "Jungle"
}

func (c *Cat) Says() string {
	return "Meow"
}
func (c *Cat) numOfLegs() int {
	return 4
}
func (d *Dog) Habitat() string {
	return "Home"
}

func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) numOfLegs() int {
	return 4
}
