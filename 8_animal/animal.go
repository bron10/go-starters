package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, sound string
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func promptUserInput(text string) string {
	var num string
	fmt.Printf("\nEnter %s : ", text)
	ptr, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Not a valid : ", err, ptr)
	}
	return num
}

func main() {
	dataMap := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	for {
		inputAnimal := promptUserInput("Select animal (cow, bird, snake)")
		action := promptUserInput("Select action (eat, move, speak)")

		fmt.Print("\n Response : ")
		switch action {
		case "eat":
			dataMap[inputAnimal].Eat()
		case "move":
			dataMap[inputAnimal].Move()
		case "speak":
			dataMap[inputAnimal].Speak()
		}
	}
}
