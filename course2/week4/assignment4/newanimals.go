package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

func NewAnimal(t string) Animal {
	if t == "cow" {
		return Cow{}
	} else if t == "bird" {
		return Bird{}
	} else {
		return Snake{}
	}
}

func GetInfo(a *Animal, i string) {
	if i == "eat" {
		(*a).Eat()
	} else if i == "move" {
		(*a).Move()
	} else {
		(*a).Speak()
	}
}

type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)

	fmt.Println("Program allows to create new animal or fetch information about existing animals.")
	fmt.Println("To create animal write: newanimal <name: string> <type: cow/bird/snake>")
	fmt.Println("To get the animal and its information write: query <name: string> <info: eat/move/speak>")
	fmt.Println("To exit the program press Ctrl+C.")

	for {
		var cmd, name, rest string

		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &cmd, &name, &rest)

		if cmd != "newanimal" && cmd != "query" {
			fmt.Printf("Invalid command %s, available commands: newanimal, query\n", cmd)
			continue
		}

		if cmd == "newanimal" {
			if rest != "cow" && rest != "bird" && rest != "snake" {
				fmt.Printf("Invalid animal type %s, available types: cow, bird, snake\n", rest)
				continue
			}

			_, ok := animals[name]
			if ok {
				fmt.Printf("Animal %s already exists, please enter unique name\n", name)
				continue
			}

			a := NewAnimal(rest)
			animals[name] = a

			fmt.Println("Created it!")
		}

		if cmd == "query" {
			if rest != "eat" && rest != "move" && rest != "speak" {
				fmt.Printf("Invalid animal info %s, available info: eat, move, speak\n", rest)
				continue
			}

			i, ok := animals[name]
			if !ok {
				fmt.Printf("Animal %s doesn't exist, try again\n", name)
				continue
			}

			GetInfo(&i, rest)
		}
	}
}
