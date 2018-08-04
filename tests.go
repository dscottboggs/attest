package attest

import (
	"fmt"
	"log"
	"testing"
)

// Test -- A structure for containing methods and data for asserting and
// testing assertion validity
type Test struct {
	*testing.T
}

func typeOf(val interface{}) string {
	return fmt.Sprintf("%T", val)
}

// Equals that var1 is deeply equal to var2. Optionally, you can pass an
// additional string and additional string formatters to be passed to
// Test.Attest. If no message is specified, a message will be logged simply
// stating that the two values weren't equal.
func (t *Test) Equals(
	var1, var2 interface{}, msgAndFormatters ...interface{},
) {
	if len(msgAndFormatters) > 0 {
		t.Attest(
			typeOf(var1) == typeOf(var2),
			msgAndFormatters[0].(string),
			msgAndFormatters[1:]...)
		t.Attest(
			var1 == var2,
			msgAndFormatters[0].(string),
			msgAndFormatters[1:]...)
	} else {
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
}

// NotEqual fails the test if var1 doesn't equal var2, with the given message
// and formatting.
func (t *Test) NotEqual(var1, var2 interface{},
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
	if !that {
		if len(formatters) == 0 {
			fmt.Println(message)
		} else {
			fmt.Printf(message+"\n", formatters...)
		}
		t.Fail()
	}
}

// AttestOrDo -- call `callback` with the Test as a parameter and fail the test
// should `that` be false.
func (t *Test) AttestOrDo(that bool,
	callback func(*Test, ...interface{}),
	cbArgs ...interface{},
) {
	if !that {
		callback(t, cbArgs...)
		t.Fail()
	}
}

// Nil -- Log a message and fail if the variable is not nil
func (t *Test) Nil(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Attest(
			variable == nil,
			"%#v was expected to be nil, but was not!",
			variable)
	} else if len(msgAndFmt) == 1 {
		t.Attest(
			variable == nil,
			msgAndFmt[0].(string))
	} else {
		t.Attest(
			variable == nil,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
}

// NotNil --  Log a message and fail if the variable is nil.
func (t *Test) NotNil(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Attest(
			variable != nil,
			"%#v was expected to have a value but instead was nil",
			variable)
	} else if len(msgAndFmt) == 1 {
		t.Attest(
			variable != nil,
			msgAndFmt[0].(string))
	} else {
		t.Attest(
			variable != nil,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
}

// GreaterThan -- log a message and fail if the variable is less than the
// expected value
func (t *Test) GreaterThan(expected,
	variable interface{},
	msgAndFmt ...interface{},
) {
	defMsg := fmt.Sprintf(
		"Value (%#v) was less than expected (%#v).",
		variable,
		expected)
	msg := func() string {
		if len(msgAndFmt) == 0 {
			return defMsg
		}
		if len(msgAndFmt) == 1 {
			return msgAndFmt[0].(string)
		}
		return fmt.Sprintf(msgAndFmt[0].(string), msgAndFmt[1:]...)
	}
	switch variable.(type) {
	default:
		log.Printf(
			"When trying check that %v was greater than %v, found non-numeric "+
				"types %T and %T",
			expected,
			variable,
			expected,
			variable)
		t.Fail()
	case int:
		t.Attest(variable.(int) > expected.(int), msg())
	case int8:
		t.Attest(variable.(int8) > expected.(int8), msg())
	case int16:
		t.Attest(variable.(int16) > expected.(int16), msg())
	case int32:
		t.Attest(variable.(int32) > expected.(int32), msg())
	case int64:
		t.Attest(variable.(int64) > expected.(int64), msg())
	case float32:
		t.Attest(variable.(float32) > expected.(float32), msg())
	case float64:
		t.Attest(variable.(float64) > expected.(float64), msg())
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// LessThan -- log a message and fail if variable is negative.
func (t *Test) LessThan(expected,
	variable interface{},
	msgAndFmt ...interface{},
) {
	defMsg := fmt.Sprintf(
		"Value (%#v) was greater than expected (%#v).",
		variable,
		expected)
	msg := func() string {
		if len(msgAndFmt) == 0 {
			return defMsg
		}
		if len(msgAndFmt) == 1 {
			return msgAndFmt[0].(string)
		}
		return fmt.Sprintf(msgAndFmt[0].(string), msgAndFmt[1:]...)
	}
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check value of %#v: check isn't implemented for type %T",
			variable,
			variable)
		t.Fail()
	case int:
		t.Attest(variable.(int) < expected.(int), msg())
	case int8:
		t.Attest(variable.(int8) < expected.(int8), msg())
	case int16:
		t.Attest(variable.(int16) < expected.(int16), msg())
	case int32:
		t.Attest(variable.(int32) < expected.(int32), msg())
	case int64:
		t.Attest(variable.(int64) < expected.(int64), msg())
	case float32:
		t.Attest(variable.(float32) < expected.(float32), msg())
	case float64:
		t.Attest(variable.(float64) < expected.(float64), msg())
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// Positive -- log a message and fail if variable is negative or zero.
func (t *Test) Positive(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		msgAndFmt = []interface{}{"%#v was not positive", variable}
	}
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check that %#v is positive: check isn't implemented for "+
				"type %T",
			variable,
			variable)
		t.Fail()
	case int:
		t.Attest(
			variable.(int) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int8:
		t.Attest(
			variable.(int8) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int16:
		t.Attest(
			variable.(int16) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int32:
		t.Attest(
			variable.(int32) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int64:
		t.Attest(
			variable.(int64) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float32:
		t.Attest(
			variable.(float32) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float64:
		t.Attest(
			variable.(float64) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// Negative -- log a message and fail if variable is positive or zero.
func (t *Test) Negative(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		msgAndFmt = []interface{}{"%#v was not positive", variable}
	}
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check that %#v is negative: check isn't implemented for "+
				"type %T",
			variable,
			variable)
		t.Fail()
	case int:
		t.Attest(
			variable.(int) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int8:
		t.Attest(
			variable.(int8) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int16:
		t.Attest(
			variable.(int16) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int32:
		t.Attest(
			variable.(int32) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int64:
		t.Attest(
			variable.(int64) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float32:
		t.Attest(
			variable.(float32) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float64:
		t.Attest(
			variable.(float64) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// In -- Check if the given value is a member the given iterable.
func (t *Test) In(iterable []interface{},
	value interface{},
	msgAndFmt ...interface{},
) {
	found := false
	for _, entry := range iterable {
		if entry == value {
			found = true
		}
	}
	if !found {
		if len(msgAndFmt) == 0 {
			msgAndFmt = []interface{}{"%v was not found in %v",
				value,
				iterable}
		}
		log.Printf(msgAndFmt[0].(string)+"\n", msgAndFmt[1:]...)
		t.Fail()
	}
}

// NotIn -- Fail the test if value is a member of iterable.
func (t *Test) NotIn(iterable []interface{},
	value interface{},
	msgAndFmt ...interface{},
) {
	found := false
	for _, entry := range iterable {
		if entry == value {
			found = true
		}
	}
	if found {
		if len(msgAndFmt) == 0 {
			msgAndFmt = []interface{}{"%v was found in %v", value, iterable}
		}
		log.Printf(msgAndFmt[0].(string)+"\n", msgAndFmt[1:]...)
		t.Fail()
	}
}

// TypeIs fails the test if the type of the value does not match the typestring,
// as determined by fmt.Sprintf("%T")
func (t *Test) TypeIs(typestring string, value interface{}) {
	if fmt.Sprintf("%T", value) == typestring {
		return
	}
	t.Errorf("Type of %#v wasn't %s.", value, typestring)
}

// TypeIsNot is the inverse of TypeIs; it fails the test if the type of value
// matches the typestring.
func (t *Test) TypeIsNot(typestring string, value interface{}) {
	if fmt.Sprintf("%T", value) == typestring {
		t.Errorf("Type of %#v was %s.", value, typestring)
	}
}

/* 						ERROR HANDLING TESTS

These tests are passed (possibly nil) errors. The test fails if the error is
not nil, and logs the error and, in some cases, an optional custom message.
*/

// AttestPanics -- Attest that when fun is called with args, it causes a panic.
// e.g. t.AttestPanics(func(){log.Printf("Panics, passes test."); panic()})
//			t.AttestPanics(func(){log.Printf("Doesn't panic, fails test.")})
func (t *Test) AttestPanics(fun func(...interface{}), args ...interface{}) {
	defer func() {
		r := recover()
		t.Attest(r != nil, "Function %v didn't cause a panic!", fun)
	}()
	fun(args...)
}

// AttestNoPanic -- the inverse of AttestPanics
func (t *Test) AttestNoPanic(fun func(...interface{}), args ...interface{}) {
	defer func() {
		r := recover()
		t.Attest(r == nil, "Function %v caused a panic!", fun)
	}()
	fun(args...)
}

// Handle -- log and fail for an arbitrary number of errors.
func (t *Test) Handle(e ...error) {
	for _, err := range e {
		if err != nil {
			log.Println(err)
			t.Fail()
		}
	}
}

// MessageHandle -- handle an error with a custom message.
func (t *Test) MessageHandle(err error, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Handle(err)
	} else {
		if err != nil {
			log.Printf(msgAndFmt[0].(string), msgAndFmt[1:]...)
			t.Fail()
		}
	}
}

// StopIf -- Fail the test and stop running it if an error is present, with
// optional message.
func (t *Test) StopIf(err error, msgAndFmt ...interface{}) {
	if err != nil {
		if len(msgAndFmt) == 0 {
			msgAndFmt = []interface{}{"Fatal error: %#v", err}
		}
		log.Printf(msgAndFmt[0].(string), msgAndFmt[1:]...)
		t.FailNow()
	}

}

// EatError accepts two values, the latter of which is a nillable error. If the
// error is not nil, the test is failed. Regardless, the first value is
// returned through the function.
func (t *Test) EatError(value interface{}, err error) interface{} {
	if err != nil {
		t.Errorf("When aquiring value %#v, got error %#v", value, err)
	}
	return value
}
