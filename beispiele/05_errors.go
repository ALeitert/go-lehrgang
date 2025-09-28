package foundations

import (
	"errors"
	"fmt"
	"io"
)

// There is no try/catch. Functions return errors.
// Standard: error is last parameter.
func floatDivide(a, b int) (float64, error) {
	if b == 0 {
		// Standard: other return values have zero-values.
		return 0, errors.New("divider cannot be 0")
	}

	// Return nil if no error.
	return float64(a) / float64(b), nil
}

func handleErrors() error {
	// Standard: handle error directly
	_, err := floatDivide(42, 0)
	if err != nil {
		fmt.Println("there was an error", err)

		// Check for specific errors with `errors.Is()`
		if errors.Is(err, io.ErrUnexpectedEOF) {
			fmt.Printf("specific error")
		}

		// We can wrap errors and hand them upwards.
		return fmt.Errorf("there was an error: %w", err)
	}

	return nil
}

func panics() {
	// A panic ends the program directly.
	// Should be avoided.
	panic("this is the end")
}

func recoverFromPanic() {
	defer func() {
		// Returns non-nil if there was a panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	// Use case: middleware in API servers.

	panics()
}
