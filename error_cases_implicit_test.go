package attest

import (
	"testing"
)

func TestAttestPanics(t *testing.T) {
	test := Test{t}
	test.AttestPanics(
		func(a ...interface{}) { panic(a[0].(string)) },
		"test panic",
	)
}

func TestAttestNoPanic(t *testing.T) {
	test := Test{t}
	test.AttestNoPanic(
		func(a ...interface{}) { a[0] = a[1] },
		"args for",
		"callback func",
	)
}

// this function is called by EatError, it returns a value and a string
func returnsError() (string, error) {
	return "success", nil
}

func TestEatError(t *testing.T) {
	test := Test{t}
	test.Equals("success", test.EatError(returnsError()).(string))
}

// the following are explicit tests on the implementation, not implicit tests
// like the others.
func TestPanicCheckImplementationWithPanic(t *testing.T) {
	test := NewTest(t)
	defer func() {
		r := recover()
		test.Attest(r != nil, "Test error function didn't panic!")
	}()
	panic("test panic")
}
func TestPanicCheckImplementationWithNoPanic(t *testing.T) {
	test := NewTest(t)
	defer func() {
		r := recover()
		test.Attest(r == nil, "Basic math caused a panic.")
	}()
	_ = 2 + 2
}
