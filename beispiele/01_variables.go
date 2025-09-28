package foundations

import (
	"fmt"
)

func basicTypes() {
	// Fixed Size
	//
	//   Integers: [u]int(8|16|32|64)
	//   Floats:   float(32|64)
	//   Complex:  complex(64|128)
	//
	//   byte = uint8
	//   rune = int32  -- represents a unicode character

	// Implementation-Specific Sizes:
	//   [u]int   --  usually 32 or 64 bits
	//   uintptr  --  an unsigned integer large enough to store the
	//                uninterpreted bits of a pointer value
}

func variables() {
	// Creates a new variable `a` of type `int` with default "zero" value.
	var a int
	fmt.Println("a:", a)

	// Creates a new variable `b` of type `float64` with value `2.0`.
	var b float64 = 2.0

	// If we set a value, we can make it shorter with `:=`:
	// Creates variables `c` (`int`) and `d` (`float64`).
	// Go has some type inference (but do not rely too much on it).
	c := 3
	d := 4.0

	// Important: `:=` declares a new variable. Assigning a new value to an
	// existing is done with a simple `=`.
	c := 5 // error: "no new variables on left side of :="
	c = 6  // all good

	// We can also group declarations:
	var (
		aa int
		bb float64 = 2.0
	)

	// We can also shadow variables.
	fmt.Println("a:", a)
	{
		a := 7
		fmt.Println("a:", a)
	}
	fmt.Println("a:", a)

	// Variables can contain functions.
	doNoting := func(_ any) {}

	// Compiler throws errors if variables are unused.
	doNoting(a)
	doNoting(b)
	doNoting(c)
	doNoting(d)
	doNoting(aa)
	doNoting(bb)
}

func constants() {
	// Work mostly like variables.
	const a int = 1
	const (
		c float32 = 2.0
		d string  = "3"
	)

	// However:

	// They need to have a value.
	const e int // error: "missing init expr for e"

	// They can be "untyped integer/float"s.
	const (
		untypedInt       = 42
		typedInt   int32 = 42
	)

	var someByte byte = 123
	someByte -= untypedInt // works fine
	someByte -= typedInt   // error: "mismatched types byte and int32"
}

func strings() {
	// Strings are sequences of UTF-8 characters.
	var hello string = "Hello - Привет - 你好 - مرحبًا"

	// Use \ to escape special characters.
	multiLine := "line 1\nline 2\nquotation mark: \""

	// Use \u to set unicode characters:
	specialChars := "\u00BF" // ¿

	// Use \x to encode single bytes:
	hexChars := "\xE2\x82\xAC" // €

	//
	// With backticks (`), we can write raw string without escaping.

	rawString := `Writing \n has no effect here.
But we can make line breaks in the string.
And "quotation marks".

Great to write regular expressions or SQL/Prometheus queries.

	Keep in mind, indentions will be in the result.`

	backtick := `Works great until you need a ` + "`" + ` in your string.`
}

func pointers() {
	// variable
	a := 42

	// pointer to variable
	aPtr := &a

	// dereference a pointer
	fmt.Println(*aPtr) // 42
	*aPtr = 123
	fmt.Println(a) // 123

	// unassigned pointer: *T
	var bPtr *int
	fmt.Println(bPtr) // nil
}
