/**
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package attest

import (
	"log"
	"regexp"
	"testing"
)

func TestAttest(t *testing.T) {
	test := New(t)
	test.Attest(true, "attest.Test.Attest has failed an implicit test.")
}
func TestAttestNot(t *testing.T) {
	test := New(t)
	test.AttestNot(false, "attest.Test.AttestNot has failed an implicit test.")
	test.Not(false, "attest.Test.Not has failed an implicit test.")
}
func TestEquals(t *testing.T) {
	test := New(t)
	test.Equals(
		"attest.Test.Equals has failed an implicit test.",
		"attest.Test.Equals has failed an implicit test.")
}
func TestCompares(t *testing.T) {
	test := NewTest(t)
	test.Compares("987", 987)
	test.SimilarTo([]string{"5", "6", "7"}, []int{5, 6, 7})
}
func TestDoesNotCompare(t *testing.T) {
	test := NewTest(t)
	test.DoesNotCompare("two values that", "are not the same")
	test.NotSimilarTo(5, "var2")
}
func TestAttestOrDo(t *testing.T) {
	test := New(t)
	test.AttestOrDo(true, func(t *Test, args ...interface{}) {
		log.Printf("attest.Test.AttestOrDo has failed an implicit test")
	})
}
func TestNil(t *testing.T) {
	test := New(t)
	test.Nil(nil, "attest.Test.Nil as failed an implicit test")
}
func TestNotNil(t *testing.T) {
	test := New(t)
	test.NotNil(
		"non-nil value",
		"required message: %s",
		"attest.Test.NotNil failed an implicit test.", //formatters not required
	)
}
func TestGreaterThan(t *testing.T) {
	test := New(t)
	test.GreaterThan(1, 2)
	test.GreaterThan(1.3, 2.5)
	test.GreaterThan(int8(1), int8(2))
	test.GreaterThan(int16(1), int16(2))
	test.GreaterThan(int32(1), int32(2))
	test.GreaterThan(int64(1), int64(2))
	test.GreaterThan(float32(1.3), float32(2.1))
	test.GreaterThan(float64(1.3), float64(2.1))
}
func TestPositive(t *testing.T) {
	test := New(t)
	test.Positive(2)
	test.Positive(2.5)
	test.Positive(int8(2))
	test.Positive(int16(2))
	test.Positive(int32(2))
	test.Positive(int64(2))
	test.Positive(float32(2.1))
}
func TestNegative(t *testing.T) {
	test := New(t)
	test.Negative(-2)
	test.Negative(-2.5)
	test.Negative(int8(-2))
	test.Negative(int16(-2))
	test.Negative(int32(-2))
	test.Negative(int64(-2))
	test.Negative(float32(-2.1))
}
func TestLessThan(t *testing.T) {
	test := New(t)
	test.LessThan(2, 1)
	test.LessThan(2.5, 1.3)
	test.LessThan(int8(2), int8(1))
	test.LessThan(int16(2), int16(1))
	test.LessThan(int32(2), int32(1))
	test.LessThan(int64(2), int64(1))
	test.LessThan(float32(2.1), float32(1.3))
	test.LessThan(float64(2.1), float64(1.3))
}

func TestNotEqual(t *testing.T) {
	test := New(t)
	var1 := "test var 1"
	var2 := "test var 2"
	test.NotEqual(
		var1, var2, "The strings %s and %s were somehow equal", var1, var2)
	test.NotEqual(
		var1,
		2,
		"The differently-typed values %s and %d were somehow equal.",
		var1,
		2)
}

func TestMatches(t *testing.T) {
	test := New(t)
	var pattern = test.FailOnError(regexp.Compile("foo.*")).(*regexp.Regexp)
	var value = "seafood"
	test.Matches(pattern, value)
}

func TestDoesNotMatch(t *testing.T) {
	test := New(t)
	var pattern = test.FailOnError(regexp.Compile("doesn't match")).(*regexp.Regexp)
	var value = "zxcvbn"
	test.DoesNotMatch(pattern, value)
}
