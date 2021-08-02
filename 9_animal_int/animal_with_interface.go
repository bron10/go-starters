package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalStruc struct {
	food, locomotion, sound string
}

type Cow struct {
	food, locomotion, sound string
}
type Bird struct {
	food, locomotion, sound string
}
type Snake struct {
	food, locomotion, sound string
}

func (v AnimalStruc) Eat() {
	fmt.Println(v.food)
}

func (v AnimalStruc) Move() {
	fmt.Println(v.locomotion)
}

func (v AnimalStruc) Speak() {
	fmt.Println(v.sound)
}

func main() {
	var animalInt Animal
	dataMap := map[string]AnimalStruc{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	for {
		var command, animalRequest, animalTypeAction string
		fmt.Print(">")
		fmt.Scan(&command, &animalRequest, &animalTypeAction)

		if command == "query" {
			animalInt = dataMap[animalRequest]
			fmt.Print("\n Response : ")
			switch animalTypeAction {
			case "eat":
				animalInt.Eat()
			case "move":
				animalInt.Move()
			case "speak":
				animalInt.Speak()
			}
		}
		if command == "newanimal" {
			dataMap[animalRequest] = dataMap[animalTypeAction]
			fmt.Println("Created it!")
		}
	}
}
