package attest

import (
	"fmt"
	"log"
	"testing"
)

// Test -- A structure for containing methods and data for asserting and
// testing assertion validity
type Test struct {
	NativeTest *testing.T
}

// AttestEquals that var1 is deeply equal to var2, and print a helpful message
// if not
func (t *Test) AttestEquals(var1, var2 interface{}) {
	t.Attest(
		var1 == var2,
		fmt.Sprintf(
			"Expected %#v (%v) was actually %#v (%v)",
			var1,
			var1,
			var2,
			var2))
}

// Attest that `that` is true, or log `message` and fail the test.
func (t *Test) Attest(that bool, message string, formatters ...interface{}) {
	if len(formatters) > 0 {
		message = fmt.Sprintf(message, formatters...)
	}
	t.AttestOrDo(that, func(t *Test) {
		log.Println(message)
	})
}

// AttestOrDo -- call `callback` with the Test as a parameter and fail the test
// should `that` be false.
func (t *Test) AttestOrDo(that bool, callback func(*Test)) {
	if !that {
		callback(t)
		t.NativeTest.Fail()
	}
}

// AttestNil -- Log a message and fail if the variable is not nil
func (t *Test) AttestNil(variable interface{}) {
	if variable != nil {
		log.Printf("%#v was expected to be nil, but was not!", variable)
		t.NativeTest.Fail()
	}
}

// AttestNotNil --  Log a message and fail if the variable is nil.
func (t *Test) AttestNotNil(variable interface{}) {
	if variable == nil {
		log.Printf(
			"%#v was expected to have a value but instead was nil", variable)
		t.NativeTest.Fail()
	}
}

// AttestGreaterThan -- log a message and fail if the variable is less than the
// expected value
func (t *Test) AttestGreaterThan(expected, variable interface{}) {
	switch variable.(type) {
	default:
		log.Printf(
			"When trying check that %v was greater than %v, found non-numeric "+
				"types %T and %T",
			expected,
			variable,
			expected,
			variable)
	case int8:
		if variable.(int8) < expected.(int8) {
			log.Printf(
				"Value (%#v) was less than expected (%#v).", variable, expected)
			t.NativeTest.Fail()
		}
	case int16:
		if variable.(int16) < expected.(int16) {
			log.Printf(
				"Value (%#v) was less than expected (%#v).", variable, expected)
			t.NativeTest.Fail()
		}
	case int32:
		if variable.(int32) < expected.(int32) {
			log.Printf(
				"Value (%#v) was less than expected (%#v).", variable, expected)
			t.NativeTest.Fail()
		}
	case int64:
		if variable.(int64) < expected.(int64) {
			log.Printf(
				"Value (%#v) was less than expected (%#v).", variable, expected)
			t.NativeTest.Fail()
		}
	case float32:
		if variable.(float32) < expected.(float32) {
			log.Printf(
				"Value (%#v) was less than expected (%#v).", variable, expected)
			t.NativeTest.Fail()
		}
	case float64:
		if variable.(float64) < expected.(float64) {
			log.Printf(
				"Value (%#v) was less than expected (%#v).", variable, expected)
			t.NativeTest.Fail()
		}
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// AttestPositive -- log a message and fail if variable is negative.
func (t *Test) AttestPositive(variable interface{}) {
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check that %#v is positive: check isn't implemented for "+
				"type %T",
			variable,
			variable)
	case int8:
		if variable.(int8) < 0 {
			log.Printf(
				"Value (%#v) was not positive.", variable)
			t.NativeTest.Fail()
		}
	case int16:
		if variable.(int16) < 0 {
			log.Printf(
				"Value (%#v) was not positive.", variable)
			t.NativeTest.Fail()
		}
	case int32:
		if variable.(int32) < 0 {
			log.Printf(
				"Value (%#v) was not positive.", variable)
			t.NativeTest.Fail()
		}
	case int64:
		if variable.(int64) < 0 {
			log.Printf(
				"Value (%#v) was not positive.", variable)
			t.NativeTest.Fail()
		}
	case float32:
		if variable.(float32) < 0 {
			log.Printf(
				"Value (%#v) was not positive.", variable)
			t.NativeTest.Fail()
		}
	case float64:
		if variable.(float64) < 0 {
			log.Printf(
				"Value (%#v) was not positive.", variable)
			t.NativeTest.Fail()
		}
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// Handle -- log and fail for an arbitrary number of errors.
func (t *Test) Handle(e ...error) {
	for _, err := range e {
		if err != nil {
			log.Println(err)
			t.NativeTest.Fail()
		}
	}
}
