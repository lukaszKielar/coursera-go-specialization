package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func (a *Animal) Info(i string) {
	if i == "eat" {
		a.Eat()
	} else if i == "move" {
		a.Move()
	} else if i == "speak" {
		a.Speak()
	} else {
		fmt.Printf("Invalid information %s, please use one of: eat, move, speak\n", i)
	}
}

func main() {
	animals := make(map[string]Animal)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}

	fmt.Println("Enter animal's name (cow, bird, snake) and information (eat, move, speak) after > prompt. To exit press ctrl+C.")
	for {
		var animal, info string

		fmt.Print("> ")
		fmt.Scanf("%s %s", &animal, &info)

		a, ok := animals[animal]
		if !ok {
			fmt.Printf("Invalid animal %s, please use one of: cow, bird, snake\n", animal)
			continue
		}

		a.Info(info)
	}
}
