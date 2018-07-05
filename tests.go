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

func typeOf(val interface{}) string {
	return fmt.Sprintf("%T", val)
}

// AttestEquals that var1 is deeply equal to var2, and print a helpful message
// if not
func (t *Test) AttestEquals(var1, var2 interface{}) {
	t.Attest(
		typeOf(var1) == typeOf(var2),
		fmt.Sprintf(
			"%#v of type %T didn't match the type of %#v, %T; so they can't be"+
				" compared. ",
			var1,
			var1,
			var2,
			var2))
	t.Attest(
		var1 == var2,
		fmt.Sprintf(
			"Expected %#v (%v) was actually %#v (%v)",
			var1,
			var1,
			var2,
			var2))
}

func (t *Test) AttestNotEqual(var1, var2 interface{},
	msg string,
	fmt ...interface{},
) {
	if typeOf(var1) != typeOf(var2) {
		// types don't match, not equal by default.
		return
	}
	t.Attest(var1 != var2, msg, fmt...)
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
	t.Attest(
		variable == nil,
		"%#v was expected to be nil, but was not!",
		variable)
}

// AttestNotNil --  Log a message and fail if the variable is nil.
func (t *Test) AttestNotNil(variable interface{}) {
	t.Attest(
		variable != nil,
		"%#v was expected to have a value but instead was nil",
		variable)
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
		t.NativeTest.Fail()
	case int:
		t.Attest(
			variable.(int) > expected.(int),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	case int8:
		t.Attest(
			variable.(int8) > expected.(int8),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	case int16:
		t.Attest(
			variable.(int16) > expected.(int16),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	case int32:
		t.Attest(
			variable.(int32) > expected.(int32),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	case int64:
		t.Attest(
			variable.(int64) > expected.(int64),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	case float32:
		t.Attest(
			variable.(float32) > expected.(float32),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	case float64:
		t.Attest(
			variable.(float64) > expected.(float64),
			"Value (%#v) was less than expected (%#v).",
			variable,
			expected)
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// AttestLessThan -- log a message and fail if variable is negative.
func (t *Test) AttestLessThan(expected, variable interface{}) {
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check value of %#v: check isn't implemented for type %T",
			variable,
			variable)
		t.NativeTest.Fail()
	case int:
		t.Attest(
			variable.(int) < expected.(int),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
	case int8:
		t.Attest(
			variable.(int8) < expected.(int8),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
	case int16:
		t.Attest(
			variable.(int16) < expected.(int16),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
	case int32:
		t.Attest(
			variable.(int32) < expected.(int32),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
	case int64:
		t.Attest(
			variable.(int64) < expected.(int64),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
	case float32:
		t.Attest(
			variable.(float32) < expected.(float32),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
	case float64:
		t.Attest(
			variable.(float64) < expected.(float64),
			"Value (%#v) was not less than expected (%#v).",
			variable,
			expected)
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
		t.NativeTest.Fail()
	case int:
		t.Attest(
			variable.(int) > 0,
			"Value (%#v) was not positive.",
			variable)
	case int8:
		t.Attest(
			variable.(int8) > 0,
			"Value (%#v) was not positive.",
			variable)
	case int16:
		t.Attest(
			variable.(int16) > 0,
			"Value (%#v) was not positive.",
			variable)
	case int32:
		t.Attest(
			variable.(int32) > 0,
			"Value (%#v) was not positive.",
			variable)
	case int64:
		t.Attest(
			variable.(int64) > 0,
			"Value (%#v) was not positive.",
			variable)
	case float32:
		t.Attest(
			variable.(float32) > 0,
			"Value (%#v) was not positive.",
			variable)
	case float64:
		t.Attest(
			variable.(float64) > 0,
			"Value (%#v) was not positive.",
			variable)
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// AttestNegative -- log a message and fail if variable is negative.
func (t *Test) AttestNegative(variable interface{}) {
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check that %#v is negative: check isn't implemented for "+
				"type %T",
			variable,
			variable)
		t.NativeTest.Fail()
	case int:
		t.Attest(
			variable.(int) < 0,
			"Value (%#v) was not negative.",
			variable)
	case int8:
		t.Attest(
			variable.(int8) < 0,
			"Value (%#v) was not negative.",
			variable)
	case int16:
		t.Attest(
			variable.(int16) < 0,
			"Value (%#v) was not negative.",
			variable)
	case int32:
		t.Attest(
			variable.(int32) < 0,
			"Value (%#v) was not negative.",
			variable)
	case int64:
		t.Attest(
			variable.(int64) < 0,
			"Value (%#v) was not negative.",
			variable)
	case float32:
		t.Attest(
			variable.(float32) < 0,
			"Value (%#v) was not negative.",
			variable)
	case float64:
		t.Attest(
			variable.(float64) < 0,
			"Value (%#v) was not negative.",
			variable)
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
