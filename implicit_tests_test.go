package attest

import (
	"log"
	"testing"
)

func TestAttest(t *testing.T) {
	test := Test{t}
	test.Attest(true, "attest.Test.Attest has failed an implicit test.")
}
func TestAttestEquals(t *testing.T) {
	test := Test{t}
	test.AttestEquals(
		"attest.Test.AttestEquals has failed an implicit test.",
		"attest.Test.AttestEquals has failed an implicit test.")
}
func TestAttestOrDo(t *testing.T) {
	test := Test{t}
	test.AttestOrDo(true, func(t *Test) {
		log.Printf("attest.Test.AttestOrDo has failed an implicit test")
	})
}
func TestAttestNil(t *testing.T) {
	test := Test{t}
	test.AttestNil(nil)
}
func TestAttestNotNil(t *testing.T) {
	test := Test{t}
	test.AttestNotNil("attest.Test.AttestNotNil failed an implicit test.")
}
func TestAttestGreaterThan(t *testing.T) {
	test := Test{t}
	test.AttestGreaterThan(2, 1)
	test.AttestGreaterThan(2.5, 1.3)
	test.AttestGreaterThan(int8(2), int8(1))
	test.AttestGreaterThan(int16(2), int16(1))
	test.AttestGreaterThan(int32(2), int32(1))
	test.AttestGreaterThan(int64(2), int64(1))
	test.AttestGreaterThan(float32(2.1), float32(1.3))
}
func TestAttestPositive(t *testing.T) {
	test := Test{t}
	test.AttestPositive(2)
	test.AttestPositive(2.5)
	test.AttestPositive(int8(2))
	test.AttestPositive(int16(2))
	test.AttestPositive(int32(2))
	test.AttestPositive(int64(2))
	test.AttestPositive(float32(2.1))
}
