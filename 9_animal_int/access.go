package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	sliceOfAnimals := []animal{}

	fmt.Println("To create an animal for the database, you will need to use the following Syntax:")
	fmt.Println("newanimal (name of animal) (type of animal)")
	fmt.Println("Upon creation of an animal, you can then query the animal(s) by doing the following:")
	fmt.Println("query (name of animal) (action of animal)")

	for {

		fmt.Print(">")
		userQuery, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		userQuery = userQuery[:len(userQuery)-2]
		sliceOfQuery := strings.Split(userQuery, " ")
		if len(sliceOfQuery) > 3 || len(sliceOfQuery) < 3 {
			fmt.Println("Invalid input!")
			continue
		}

		determinant := sliceOfQuery[0]
		fmt.Print(",,,", sliceOfQuery[2])
		switch determinant {
		case "newanimal":

			if sliceOfQuery[2] == "cow" {
				sliceOfAnimals = append(sliceOfAnimals, cow{name: sliceOfQuery[1], Food_eaten: "grass", Locomotion_method: "walk", Spoken_sound: "moo"})
				fmt.Println("Created!")
			} else if sliceOfQuery[2] == "bird" {
				sliceOfAnimals = append(sliceOfAnimals, bird{name: sliceOfQuery[1], Food_eaten: "worms", Locomotion_method: "fly", Spoken_sound: "peep"})
				fmt.Println("Created!")
			} else if sliceOfQuery[2] == "snake" {
				sliceOfAnimals = append(sliceOfAnimals, snake{name: sliceOfQuery[1], Food_eaten: "mice", Locomotion_method: "slither", Spoken_sound: "hsss"})
				fmt.Println("Created!")
			} else {
				fmt.Println("Invalid input!...")
			}

		case "query":

			for _, animal := range sliceOfAnimals {
				if animal.getName() == sliceOfQuery[1] {
					if sliceOfQuery[2] == "move" {
						animal.move()
					} else if sliceOfQuery[2] == "eat" {
						animal.eat()
					} else if sliceOfQuery[2] == "speak" {
						animal.speak()
					}
				}
			}

		default:
			fmt.Println("Invalid input!")

		}
	}

}

type animal interface {
	eat()
	move()
	speak()
	getName() string
}

type cow struct {
	name              string
	Food_eaten        string
	Locomotion_method string
	Spoken_sound      string
}

type snake struct {
	name              string
	Food_eaten        string
	Locomotion_method string
	Spoken_sound      string
}

type bird struct {
	name              string
	Food_eaten        string
	Locomotion_method string
	Spoken_sound      string
}

func (c cow) getName() string {
	return c.name
}

func (s snake) getName() string {
	return s.name
}

func (b bird) getName() string {
	return b.name
}

func (c cow) eat() {
	fmt.Printf("%s eats %s\n", c.name, c.Food_eaten)
}

func (c cow) move() {
	fmt.Printf("%s %s\n", c.name, c.Locomotion_method)
}

func (c cow) speak() {
	fmt.Printf("%s %s\n", c.name, c.Spoken_sound)
}

func (s snake) eat() {
	fmt.Printf("%s eats %s\n", s.name, s.Food_eaten)
}

func (s snake) move() {
	fmt.Printf("%s %s\n", s.name, s.Locomotion_method)
}

func (s snake) speak() {
	fmt.Printf("%s %s\n", s.name, s.Food_eaten)
}

func (b bird) eat() {
	fmt.Printf("%s eats %s\n", b.name, b.Food_eaten)
}

func (b bird) move() {
	fmt.Printf("%s %s\n", b.name, b.Locomotion_method)
}

func (b bird) speak() {
	fmt.Printf("%s %s\n", b.name, b.Spoken_sound)
}
