package attest

import (
	"fmt"
	"testing"
)

func TestAttestPanics(t *testing.T) {
	test := Test{t}
	test.AttestPanics(func(a ...interface{}) { panic(a[0].(string)) }, "test panic")
	defer func() {
		r := recover()
		test.Attest(r == nil, "Printing a word caused a panic.")
	}()
	fmt.Println("Test passed.")
}

func TestAttestNoPanic(t *testing.T) {
	test := Test{t}
	test.AttestNoPanic(
		func(a ...interface{}) { fmt.Printf(a[0].(string)) },
		"Test function shouldn't panic.\n")
	// the following is an explicit test on the implementation, not an implicit
	// test like the others.
	defer func() {
		r := recover()
		test.Attest(r != nil, "Test error function didn't panic!")
	}()
	panic("test panic")
}

// this function is called by EatError, it returns a value and a string
func returnsError() (string, error) {
	return "success", nil
}

func TestEatError(t *testing.T) {
	test := Test{t}
	test.Equals("success", test.EatError(returnsError()).(string))
}
