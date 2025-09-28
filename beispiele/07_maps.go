package foundations

import (
	"fmt"
)

func goMaps() {
	// map[key type]value type
	var table map[string]int

	// Map-variables are pointers.
	// Subsequently, their zero-value is `nil`.
	fmt.Println(table == nil) // true

	// Initialize with `{}` or `make`.
	table = map[string]int{}
	table = make(map[string]int)

	// We can use `make` to give it an initial capacity.
	table = make(map[string]int, 10)

	// We can initialize it with pairs using `{}`.
	table = map[string]int{
		"Berlin":  3_685_265,
		"Hamburg": 1_862_565,
		"München": 1_505_005,
	}
}

func setGetDelete() {
	table := map[string]int{
		"Berlin":  3_685_265,
		"Hamburg": 1_862_565,
		"München": 1_505_005,
	}

	// Use `[key]` to access entries.
	fmt.Println(table["Berlin"]) // 3685265

	// Map return zero-value if there is no matching key.
	fmt.Println(table["Rostock"]) // 0

	// Use ok-pattern to check existence of key.
	pop, ok := table["Rostock"]
	fmt.Println(pop, ok) // 0, false

	table["Rostock"] = 205_307
	pop, ok = table["Rostock"]
	fmt.Println(pop, ok) // 205307, true

	// Delete with `delete`.
	delete(table, "Rostock")
	pop, ok = table["Rostock"]
	fmt.Println(pop, ok) // 0, false

	// Delete everything with `clear`.
	clear(table)
	fmt.Println(len(table)) // 0
}

func sets() {
	// No sets in Go.
	// Use `map[ ... ]struct{}` instead.
	set := map[string]struct{}{}

	_, ok := set["Bielefeld"]
	if !ok {
		fmt.Println("Bielefeld gibt es nicht!")
	}
}
