package foundations

import "fmt"

func loops() {
	// Go only has a for loop.
	// However, different headers give some convenience.

	// classic
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// only condition
	q := []int{0}
	for len(q) > 0 {
		// ...
	}

	// infinite loop
	for {
		// we can exit with break
		break
	}

	// `range` for supported types
	for i := range 10 {
		fmt.Println(i)
		continue // next iteration
		fmt.Println("not printed")
	}

	numbers := []int{1, 2, 3, 4, 5, 6}
	for i, num := range numbers {
		fmt.Printf("numbers[%d] = %d", i, num)
	}
}

func ifStatement() {
	// Basically like each other language, but parentheses are not needed.
	if 3 < 4 {
		fmt.Println("3 is smaller than 4.")
	} else if !true || false {
		fmt.Println("true implies false")
	} else {
		fmt.Println("every statement was false")
	}

	// Error: We always need to use braces.
	if 2 < 3 return

	if true {
	}
	// Error: else must be on same line as closing }.
	else {
	}
}

func switchStatement() {
	var a int

	// switch compares the given value to all cases in order.
	switch a {
	case 1:
		fmt.Println("a is 1")
		// No break needed

	case 2, 3: // We can list multiple cases at once.
		fmt.Println("a is 2 or 3")
		fallthrough // Also executes next case.

	case 4:
		fmt.Println("a is 4")

	default: // If no other case matches, optional `default` is choses.
		fmt.Println("a:", a)
	}
}
