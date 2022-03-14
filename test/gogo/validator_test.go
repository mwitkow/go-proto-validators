// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package validatortest

import (
	fmt "fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stableBytes = make([]byte, 12)
)

const (
	uuid4 = "fbe91ff5-fee7-40d3-89a8-f3db6cf210be"
	uuid1 = "66bb25e2-2e0d-11e9-b210-d663bd873d93"
)

func buildProto3(someString string, someInt uint32, identifier string,
	someValue int64, someDoubleStrict float64, someFloatStrict float32, someDouble float64,
	someFloat float32, nonEmptyString string, repeatedCount uint32,
	someStringLength string, someBytes []byte,
	optionalUUIDAny, uuid4 string,
	someEnum int32, someEmbeddedEnum int32) *ValidatorMessage3 {
	goodEmbeddedProto3 := &ValidatorMessage3_EmbeddedMessage{
		Identifier: identifier,
		SomeValue:  someValue,
	}

	goodProto3 := &ValidatorMessage3{
		SomeString:          someString,
		SomeStringRep:       []string{someString, "xyz34"},
		SomeStringNoQuotes:  someString,
		SomeStringUnescaped: someString,

		SomeInt:           someInt,
		SomeIntRep:        []uint32{someInt, 12, 13, 14, 15, 16},
		SomeIntRepNonNull: []uint32{someInt, 102},

		SomeEmbedded:               nil,
		SomeEmbeddedNonNullable:    *goodEmbeddedProto3,
		SomeEmbeddedExists:         goodEmbeddedProto3,
		SomeEmbeddedRep:            []*ValidatorMessage3_EmbeddedMessage{goodEmbeddedProto3},
		SomeEmbeddedRepNonNullable: []ValidatorMessage3_EmbeddedMessage{*goodEmbeddedProto3},

		StrictSomeDouble:           someDoubleStrict,
		StrictSomeDoubleRep:        []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeDoubleRepNonNull: []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeFloat:            someFloatStrict,
		StrictSomeFloatRep:         []float32{someFloatStrict, 0.5, 0.55, 0.6},
		StrictSomeFloatRepNonNull:  []float32{someFloatStrict, 0.5, 0.55, 0.6},

		SomeDouble:           someDouble,
		SomeDoubleRep:        []float64{someDouble, 0.5, 0.55, 0.6},
		SomeDoubleRepNonNull: []float64{someDouble, 0.5, 0.55, 0.6},
		SomeFloat:            someFloat,
		SomeFloatRep:         []float32{someFloat, 0.5, 0.55, 0.6},
		SomeFloatRepNonNull:  []float32{someFloat, 0.5, 0.55, 0.6},

		SomeNonEmptyString: nonEmptyString,
		SomeStringEqReq:    someStringLength,
		SomeStringLtReq:    someStringLength,
		SomeStringGtReq:    someStringLength,

		SomeBytesLtReq:   someBytes,
		SomeBytesGtReq:   someBytes,
		SomeBytesEqReq:   someBytes,
		RepeatedBaseType: []int32{},

		UUIDAny:       optionalUUIDAny,
		UUID4NotEmpty: uuid4,

		SomeEnum:         EnumProto3(someEnum),
		SomeEmbeddedEnum: ValidatorMessage3_EmbeddedEnum(someEmbeddedEnum),

		ValidatorMessage3_EmbeddedMessage: *goodEmbeddedProto3,
	}

	goodProto3.Repeated = make([]int32, repeatedCount)

	return goodProto3
}

func buildProto2(someString string, someInt uint32, identifier string,
	someValue int64, someDoubleStrict float64, someFloatStrict float32, someDouble float64,
	someFloat float32, nonEmptyString string, repeatedCount uint32,
	someStringLength string, someBytes []byte,
	optionalUUIDAny, uuid5 string,
	someEnum int32, someEmbeddedEnum int32) *ValidatorMessage {
	goodEmbeddedProto2 := &ValidatorMessage_EmbeddedMessage{
		Identifier: &identifier,
		SomeValue:  &someValue,
	}

	goodProto2 := &ValidatorMessage{
		StringReq:        &someString,
		StringReqNonNull: someString,

		StringOpt:        nil,
		StringOptNonNull: someString,

		StringUnescaped: &someString,

		IntReq:        &someInt,
		IntReqNonNull: someInt,
		IntRep:        []uint32{someInt, 12, 13, 14, 15, 16},
		IntRepNonNull: []uint32{someInt, 12, 13, 14, 15, 16},

		EmbeddedReq:            goodEmbeddedProto2,
		EmbeddedNonNull:        *goodEmbeddedProto2,
		EmbeddedRep:            []*ValidatorMessage_EmbeddedMessage{goodEmbeddedProto2},
		EmbeddedRepNonNullable: []ValidatorMessage_EmbeddedMessage{*goodEmbeddedProto2},

		StrictSomeDoubleReq:        &someDoubleStrict,
		StrictSomeDoubleReqNonNull: someDoubleStrict,
		StrictSomeDoubleRep:        []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeDoubleRepNonNull: []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeFloatReq:         &someFloatStrict,
		StrictSomeFloatReqNonNull:  someFloatStrict,
		StrictSomeFloatRep:         []float32{someFloatStrict, 0.5, 0.55, 0.6},
		StrictSomeFloatRepNonNull:  []float32{someFloatStrict, 0.5, 0.55, 0.6},

		SomeDoubleReq:        &someDouble,
		SomeDoubleReqNonNull: someDouble,
		SomeDoubleRep:        []float64{someDouble, 0.5, 0.55, 0.6},
		SomeDoubleRepNonNull: []float64{someDouble, 0.5, 0.55, 0.6},
		SomeFloatReq:         &someFloat,
		SomeFloatReqNonNull:  someFloat,
		SomeFloatRep:         []float32{someFloat, 0.5, 0.55, 0.6},
		SomeFloatRepNonNull:  []float32{someFloat, 0.5, 0.55, 0.6},

		SomeNonEmptyString: &nonEmptyString,
		SomeStringEqReq:    &someStringLength,
		SomeStringLtReq:    &someStringLength,
		SomeStringGtReq:    &someStringLength,

		SomeBytesLtReq:   someBytes,
		SomeBytesGtReq:   someBytes,
		SomeBytesEqReq:   someBytes,
		RepeatedBaseType: []int32{},

		UUIDAny:       &optionalUUIDAny,
		UUID4NotEmpty: &uuid5,

		SomeEnum:         (*EnumProto2)(&someEnum),
		SomeEmbeddedEnum: (*ValidatorMessage_EmbeddedEnum)(&someEmbeddedEnum),
	}

	goodProto2.Repeated = make([]int32, repeatedCount)

	return goodProto2
}

func TestGoodProto3(t *testing.T) {
	goodProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 1, 1)
	err := goodProto3.Validate()
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestGoodProto2(t *testing.T) {
	goodProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 1, 1)
	err := goodProto2.Validate()
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestStringRegex(t *testing.T) {
	tooLong1Proto3 := buildProto3("toolong", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if tooLong1Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong1Proto2 := buildProto2("toolong", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if tooLong1Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong2Proto3 := buildProto3("-%ab", 11, "bad#", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if tooLong2Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong2Proto2 := buildProto2("-%ab", 11, "bad#", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if tooLong2Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestIntLowerBounds(t *testing.T) {
	lowerThan10Proto3 := buildProto3("-%ab", 9, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan10Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan10Proto2 := buildProto2("-%ab", 9, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan10Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", -1, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan0Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", -1, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan0Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestIntUpperBounds(t *testing.T) {
	greaterThan100Proto3 := buildProto3("-%ab", 11, "abba", 101, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan100Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan100Proto2 := buildProto2("-%ab", 11, "abba", 101, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan100Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestDoubleStrictLowerBounds(t *testing.T) {
	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.3, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan035EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.3, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan035EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.300000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan035EpsilonProto3.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
	greaterThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.300000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan035EpsilonProto2.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
}

func TestDoubleStrictUpperBounds(t *testing.T) {
	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.70000000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan065EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.70000000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan065EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.6999999999, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan065EpsilonProto3.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
	lowerThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.6999999999, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan065EpsilonProto2.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
}

func TestFloatStrictLowerBounds(t *testing.T) {
	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.2999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan035EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.2999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan035EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.3000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := greaterThan035EpsilonProto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	greaterThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.3000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := greaterThan035EpsilonProto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestFloatStrictUpperBounds(t *testing.T) {
	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.7000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan065EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.7000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if greaterThan065EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.6999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := lowerThan065EpsilonProto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	lowerThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.6999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := lowerThan065EpsilonProto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestDoubleNonStrictLowerBounds(t *testing.T) {
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.2499999, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan0Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.2499999, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan0Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.25, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.25, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestDoubleNonStrictUpperBounds(t *testing.T) {
	higherThan1Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75111111, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if higherThan1Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	higherThan1Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75111111, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if higherThan1Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestFloatNonStrictLowerBounds(t *testing.T) {
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.2499999, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan0Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.2499999, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if lowerThan0Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.25, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.25, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestFloatNonStrictUpperBounds(t *testing.T) {
	higherThan1Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75111111, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if higherThan1Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	higherThan1Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75111111, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if higherThan1Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := equalTo0Proto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestStringNonEmpty(t *testing.T) {
	emptyStringProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if emptyStringProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	emptyStringProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if emptyStringProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	nonEmptyStringProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := nonEmptyStringProto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	nonEmptyStringProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := nonEmptyStringProto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestRepeatedEltsCount(t *testing.T) {
	notEnoughEltsProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 1, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if notEnoughEltsProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	notEnoughEltsProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 1, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if notEnoughEltsProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooManyEltsProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 14, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if tooManyEltsProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooManyEltsProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 14, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if tooManyEltsProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	validEltsCountProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := validEltsCountProto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	validEltsCountProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	if err := validEltsCountProto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestMsgExist(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	someProto3.SomeEmbedded = nil
	if err := someProto3.Validate(); err != nil {
		t.Fatalf("validate shouldn't fail on missing SomeEmbedded, not annotated")
	}
	someProto3.SomeEmbeddedExists = nil
	if err := someProto3.Validate(); err == nil {
		t.Fatalf("expected fail due to lacking SomeEmbeddedExists")
	} else if !strings.HasPrefix(err.Error(), "SomeEmbeddedExists: invalid field SomeEmbeddedExists:") {
		t.Fatalf("expected fieldError, got '%v'", err)
	}
}

func TestValueIsInEnum(t *testing.T) {
	outOfTopLevelEnumProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 2, 0)
	if err := outOfTopLevelEnumProto3.Validate(); err == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	outOfTopLevelEnumProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 2, 0)
	if err := outOfTopLevelEnumProto2.Validate(); err == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	outOfEmbeddedEnumProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 2)
	if err := outOfEmbeddedEnumProto3.Validate(); err == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	outOfEmbeddedEnumProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 2)
	if err := outOfEmbeddedEnumProto2.Validate(); err == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestNestedError3(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	someProto3.SomeEmbeddedExists.SomeValue = 101 // should be less than 101
	if err := someProto3.Validate(); err == nil {
		t.Fatalf("expected fail due to nested SomeEmbeddedExists.SomeValue being wrong")
	} else if !strings.HasPrefix(err.Error(), "SomeEmbeddedExists: invalid field SomeEmbeddedExists") {
		t.Fatalf("expected fieldError, got '%v'", err)
	}
}

func TestCustomError_Proto3(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	someProto3.CustomErrorInt = 30
	expectedErr := "CustomErrorInt: invalid field CustomErrorInt: My Custom Error"
	if err := someProto3.Validate(); err == nil {
		t.Fatalf("validate should fail on missing CustomErrorInt")
	} else if err.Error() != expectedErr {
		t.Fatalf("validation error should be '%s' but was '%s'", expectedErr, err.Error())
	}
}

func TestMapAlwaysPassesUntilFixedProperly(t *testing.T) {
	example := &ValidatorMapMessage3{}
	if err := example.Validate(); err != nil {
		t.Fatalf("map validators should always pass")
	}
}

func TestOneOf_Required(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
	}
	err := example.Validate()
	assert.Error(t, err, "oneof.required should fail if none of the oneof fields are set")
	assert.Contains(t, err.Error(), "Something", "error must err on the Something field")
}

func TestOneOf_NestedMessage(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
		Type: &OneOfMessage3_OneMsg{
			OneMsg: &ExternalMsg{
				Identifier: "999", // bad
				SomeValue:  99,    // good
			},
		},
		Something: &OneOfMessage3_ThreeInt{
			ThreeInt: 100, // > 20
		},
	}
	err := example.Validate()
	assert.Error(t, err, "nested message in oneof should fail validation on ExternalMsg")
	assert.Contains(t, err.Error(), "OneMsg: ", "error must err on the ExternalMsg.Identifier")
}

func TestOneOf_NestedInt(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
		Type: &OneOfMessage3_OneMsg{
			OneMsg: &ExternalMsg{
				Identifier: "abba", // good
				SomeValue:  99,     // good
			},
		},
		Something: &OneOfMessage3_ThreeInt{
			ThreeInt: 19, // > 20
		},
	}
	err := example.Validate()
	assert.Error(t, err, "nested message in oneof should fail validation on ThreeInt")
	assert.Contains(t, err.Error(), "ThreeInt", "error must err on the ThreeInt.ThreeInt")
}

func TestOneOf_Passes(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
		Type: &OneOfMessage3_OneMsg{
			OneMsg: &ExternalMsg{
				Identifier: "abba", // good
				SomeValue:  99,     // good
			},
		},
		Something: &OneOfMessage3_FourInt{
			FourInt: 101, // > 101
		},
	}
	err := example.Validate()
	if err != nil {
		assert.NoError(t, err, "This message should pass all validation")
	}
}

func TestOneOf_Regex(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
		Something: &OneOfMessage3_FiveRegex{
			FiveRegex: "11", // fail
		},
	}
	err := example.Validate()
	assert.Error(t, err, "regex applied to oneof field should fail validation on FiveRegex")
	assert.Contains(t, err.Error(), "FiveRegex", "error must err on the FiveRegex")

	example = &OneOfMessage3{
		SomeInt: 30,
		Something: &OneOfMessage3_FiveRegex{
			FiveRegex: "aaa", // pass
		},
	}
	err = example.Validate()
	if err != nil {
		assert.NoError(t, err, "This message should pass all validation")
	}
}

func TestUUID4Validation(t *testing.T) {
	testcases := []struct {
		uuid string
		fail bool
	}{
		{
			uuid: uuid1,
			fail: true,
		},
		{
			uuid: uuid4,
			fail: false,
		},
		{
			uuid: "",
			fail: true,
		},
		{
			uuid: "66bb25e2-2e0d",
			fail: true,
		},
		{
			uuid: "1234abcd",
			fail: true,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("proto2 uuid '%s'", tc.uuid), func(t *testing.T) {
			msg := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, tc.uuid, 0, 0)

			err := msg.Validate()
			failed := err != nil
			if tc.fail != failed {
				t.Errorf("Expected validation failure: %t, but got %t, err: %v", tc.fail, failed, err)
			}
		})

		t.Run(fmt.Sprintf("proto3 uuid '%s'", tc.uuid), func(t *testing.T) {
			msg := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, tc.uuid, 0, 0)

			err := msg.Validate()
			failed := err != nil
			if tc.fail != failed {
				t.Errorf("Expected validation failure: %t, but got %t, err: %v", tc.fail, failed, err)
			}
		})
	}
}
