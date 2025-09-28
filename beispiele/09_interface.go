package foundations

import "fmt"

// Interfaces are types that list functions.
type Duck interface {
	Looks()
	Swims()
	Quacks()
}

type Mallard struct{}

func (Mallard) Looks() {
	fmt.Println("Small with two wings, two feed, and a beak.")
}

func (Mallard) Swims() {
	fmt.Println("Floats without effort, feet no longer visible.")
}

func (Mallard) Quacks() {
	fmt.Println("Quack! Quack!")
}

type Dog struct{}

func (Dog) Looks() {
	fmt.Println("Sometimes large, sometimes small, usually furry.")
}

func (Dog) Swims() {
	fmt.Println("Can swim with medium effort.")
}

func (Dog) Barks() {
	fmt.Println("Wuff! Wuff")
}

func duckCheck() {
	toCheck := map[string]any{
		"Mallard": &Mallard{},
		"Dog":     &Dog{},
	}

	for name, animal := range toCheck {
		// We can check if an object implements an interface with a cast and the
		// ok-pattern.
		duck, ok := animal.(Duck)
		if ok {
			fmt.Printf("%s is a duck! ", name)
			makeNoise(duck)
		} else {
			fmt.Printf("%s is not a duck!\n", name)
		}
	}

	//

	fmt.Println()
	toCheck["number"] = 123

	for name, animal := range toCheck {
		// We can also use a switch for casting.
		switch a := animal.(type) {
		case Duck:
			fmt.Printf("%s is a duck! ", name)
			a.Quacks()

		case Mallard:
			fmt.Printf("%s is a mallard! ", name)
			a.Quacks()

		case Dog:
			fmt.Printf("%s is a dog! ", name)
			a.Barks()

		default:
			fmt.Printf("%s is of type %T\n", name, a)
		}

		// Question: Why does the dog not match with `case Dog`?
	}
}

func makeNoise(d Duck) {
	d.Quacks()
}
