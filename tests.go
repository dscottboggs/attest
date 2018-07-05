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

// AttestEquals that var1 is deeply equal to var2. Optionally, you can pass an
// additional string and additional string formatters to be passed to
// Test.Attest. If no message is specified, a message will be logged simply
// stating that the two values weren't equal.
func (t *Test) AttestEquals(
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
	if !that {
		if len(formatters) == 0 {
			fmt.Println(message)
		} else {
			fmt.Printf(message+"\n", formatters...)
		}
		t.NativeTest.Fail()
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
		t.NativeTest.Fail()
	}
}

// AttestNil -- Log a message and fail if the variable is not nil
func (t *Test) AttestNil(variable interface{}, msgAndFmt ...interface{}) {
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

// AttestNotNil --  Log a message and fail if the variable is nil.
func (t *Test) AttestNotNil(variable interface{}, msgAndFmt ...interface{}) {
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

// AttestGreaterThan -- log a message and fail if the variable is less than the
// expected value
func (t *Test) AttestGreaterThan(expected,
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
		t.NativeTest.Fail()
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

// AttestLessThan -- log a message and fail if variable is negative.
func (t *Test) AttestLessThan(expected,
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
		t.NativeTest.Fail()
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

// AttestPositive -- log a message and fail if variable is negative.
func (t *Test) AttestPositive(variable interface{}, msgAndFmt ...interface{}) {
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
		t.NativeTest.Fail()
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

// AttestNegative -- log a message and fail if variable is negative.
func (t *Test) AttestNegative(variable interface{}, msgAndFmt ...interface{}) {
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
		t.NativeTest.Fail()
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

// Handle -- log and fail for an arbitrary number of errors.
func (t *Test) Handle(e ...error) {
	for _, err := range e {
		if err != nil {
			log.Println(err)
			t.NativeTest.Fail()
		}
	}
}

func (t *Test) MessageHandle(err error, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Handle(err)
	} else {
		if err != nil {
			log.Printf(msgAndFmt[0].(string), msgAndFmt[:1]...)
			t.NativeTest.Fail()
		}
	}
}
