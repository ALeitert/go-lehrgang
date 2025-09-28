package foundations

import "fmt"

// Functions start with `func` have name, parameters, and return types.
func randomNumber() int {
	// Chosen by fair dice roll
	// Guaranteed to be random.
	return 4
}

func isEven(num int) bool { return num%2 == 0 }
func isOdd(num int) bool  { return num%2 != 0 }

// If multiple parameters have the same type, listing it once is sufficient.
func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

// We can return multiple values and (optionally) name them.
// Good habit: naming them if not super clear which value is what.
func intDivide(a, b int) (div int, mod int) {
	return a / b, a % b
}

// If we have named return values, we can use them as variables.
// They initially have their type's zero-value.
func lcm(a, b int) (result int) {
	result = a * b
	result /= gcd(a, b)

	return // No value needed if named.
}

func deferStatement() (val int) {
	// defer states a function which is called after the function is done.
	defer fmt.Println("done")

	fmt.Println("do something")

	// They are called in reverse order.
	defer fmt.Println("almost done")

	// They allow us to have access to return values.
	defer fmt.Println(val)                     // current value
	defer func() { fmt.Println("val", val) }() // value during defer

	val = 42

	// They are also called if we panic.
	panic("not good")
}
