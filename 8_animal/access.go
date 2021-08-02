package main

import "fmt"

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Speak() {
	fmt.Println(animal.spoken)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

type Animal struct {
	name       string
	food       string
	locomotion string
	spoken     string
}

func main() {
	var animal string
	var action string
	animals := make([]Animal, 3, 3)
	animals[0] = Animal{name: "cow", food: "grass", locomotion: "walk", spoken: "moo"}
	animals[1] = Animal{name: "bird", food: "worms", locomotion: "fly", spoken: "peep"}
	animals[2] = Animal{name: "snake", food: "mice", locomotion: "slither", spoken: "hsss"}

	for {
		fmt.Printf(">")
		_, err := fmt.Scanf("%s %s", &animal, &action)
		if err == nil {
			for _, executor := range animals {
				if executor.name == animal {
					switch action {
					case "eat":
						executor.Eat()
						break
					case "move":
						executor.Move()
						break
					case "speak":
						executor.Speak()
						break
					default:
						fmt.Println("The action is not recognized")
					}
				}
			}

		} else {
			fmt.Println("There was an error. Expected: '<animal> <action>'")
		}
	}
}
