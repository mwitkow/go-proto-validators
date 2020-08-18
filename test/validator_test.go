// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package validatortest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	stableBytes = make([]byte, 12)
)

const (
	uuid4 = "fbe91ff5-fee7-40d3-89a8-f3db6cf210be"
	uuid1 = "66bb25e2-2e0d-11e9-b210-d663bd873d93"
)

func buildProto3(someString string, someInt uint32, identifier string, someValue int64, someDoubleStrict float64,
	someFloatStrict float32, someDouble float64, someFloat float32, nonEmptyString string, repeatedCount uint32,
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

		SomeInt:    someInt,
		SomeIntRep: []uint32{someInt, 12, 13, 14, 15, 16},

		SomeEmbedded:       nil,
		SomeEmbeddedExists: goodEmbeddedProto3,
		SomeEmbeddedRep:    []*ValidatorMessage3_EmbeddedMessage{goodEmbeddedProto3},

		StrictSomeDouble:           someDoubleStrict,
		StrictSomeDoubleRep:        []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeDoubleRepNonNull: []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeFloat:            someFloatStrict,
		StrictSomeFloatRep:         []float32{someFloatStrict, 0.5, 0.55, 0.6},
		StrictSomeFloatRepNonNull:  []float32{someFloatStrict, 0.5, 0.55, 0.6},

		SomeDouble:    someDouble,
		SomeDoubleRep: []float64{someDouble, 0.5, 0.55, 0.6},
		SomeFloat:     someFloat,
		SomeFloatRep:  []float32{someFloat, 0.5, 0.55, 0.6},

		SomeNonEmptyString: nonEmptyString,
		SomeStringEqReq:    someStringLength,
		SomeStringLtReq:    someStringLength,
		SomeStringGtReq:    someStringLength,

		SomeBytesLtReq: someBytes,
		SomeBytesGtReq: someBytes,
		SomeBytesEqReq: someBytes,

		RepeatedBaseType: []int32{},

		UUIDAny:       optionalUUIDAny,
		UUID4NotEmpty: uuid4,

		SomeEnum:         EnumProto3(someEnum),
		SomeEmbeddedEnum: ValidatorMessage3_EmbeddedEnum(someEmbeddedEnum),
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
		StringReq:       &someString,
		StringOpt:       nil,
		StringNoQuotes:  &someString,
		StringUnescaped: &someString,

		IntReq: &someInt,
		IntRep: []uint32{someInt, 12, 13, 14, 15, 16},

		EmbeddedReq: goodEmbeddedProto2,
		EmbeddedRep: []*ValidatorMessage_EmbeddedMessage{goodEmbeddedProto2},

		StrictSomeDoubleReq: &someDoubleStrict,
		StrictSomeDoubleRep: []float64{someDoubleStrict, 0.5, 0.55, 0.6},
		StrictSomeFloatReq:  &someFloatStrict,
		StrictSomeFloatRep:  []float32{someFloatStrict, 0.5, 0.55, 0.6},

		SomeDoubleReq: &someDouble,
		SomeDoubleRep: []float64{someDouble, 0.5, 0.55, 0.6},
		SomeFloatReq:  &someFloat,
		SomeFloatRep:  []float32{someFloat, 0.5, 0.55, 0.6},

		SomeNonEmptyString: &nonEmptyString,
		SomeStringEqReq:    &someStringLength,
		SomeStringLtReq:    &someStringLength,
		SomeStringGtReq:    &someStringLength,
		SomeBytesLtReq:     someBytes,
		SomeBytesGtReq:     someBytes,
		SomeBytesEqReq:     someBytes,
		RepeatedBaseType:   []int32{},

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
	assert.NoError(t, goodProto3.Validate())
}

func TestGoodProto2(t *testing.T) {
	goodProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 1, 1)
	assert.NoError(t, goodProto2.Validate())
}

func TestStringRegex(t *testing.T) {
	tooLongProto3 := buildProto3("toolong", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, tooLongProto3.Validate())

	tooLongProto2 := buildProto2("toolong", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, tooLongProto2.Validate())

	mismatchProto3 := buildProto3("-%ab", 11, "bad#", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, mismatchProto3.Validate())

	mismatchProto2 := buildProto2("-%ab", 11, "bad#", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, mismatchProto2.Validate())
}

func TestIntLowerBounds(t *testing.T) {
	lowerThan10Proto3 := buildProto3("-%ab", 9, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan10Proto3.Validate())

	lowerThan10Proto2 := buildProto2("-%ab", 9, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan10Proto2.Validate())

	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", -1, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan0Proto3.Validate())

	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", -1, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan0Proto2.Validate())
}

func TestIntUpperBounds(t *testing.T) {
	greaterThan100Proto3 := buildProto3("-%ab", 11, "abba", 101, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, greaterThan100Proto3.Validate())

	greaterThan100Proto2 := buildProto2("-%ab", 11, "abba", 101, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, greaterThan100Proto2.Validate())
}

func TestDoubleStrictLowerBounds(t *testing.T) {
	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.3, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan035EpsilonProto3.Validate())

	lowerThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.3, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan035EpsilonProto2.Validate())

	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.300000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, greaterThan035EpsilonProto3.Validate())

	greaterThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.300000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, greaterThan035EpsilonProto2.Validate())
}

func TestDoubleStrictUpperBounds(t *testing.T) {
	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.70000000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, greaterThan065EpsilonProto3.Validate())

	greaterThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.70000000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, greaterThan065EpsilonProto2.Validate())

	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.6999999999, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, lowerThan065EpsilonProto3.Validate())

	lowerThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.6999999999, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, lowerThan065EpsilonProto2.Validate())
}

func TestFloatStrictLowerBounds(t *testing.T) {
	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.2999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan035EpsilonProto3.Validate())

	lowerThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.2999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan035EpsilonProto2.Validate())

	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.3000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, greaterThan035EpsilonProto3.Validate())

	greaterThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.3000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, greaterThan035EpsilonProto2.Validate())
}

func TestFloatStrictUpperBounds(t *testing.T) {
	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.7000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, greaterThan065EpsilonProto3.Validate())

	greaterThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.7000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, greaterThan065EpsilonProto2.Validate())

	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.6999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, lowerThan065EpsilonProto3.Validate())

	lowerThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.6999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, lowerThan065EpsilonProto2.Validate())
}

func TestDoubleNonStrictLowerBounds(t *testing.T) {
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.2499999, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan0Proto3.Validate())

	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.2499999, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan0Proto2.Validate())

	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.25, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto3.Validate())

	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.25, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto2.Validate())
}

func TestDoubleNonStrictUpperBounds(t *testing.T) {
	higherThan1Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75111111, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, higherThan1Proto3.Validate())

	higherThan1Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75111111, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, higherThan1Proto2.Validate())

	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto3.Validate())

	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto2.Validate())
}

func TestFloatNonStrictLowerBounds(t *testing.T) {
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.2499999, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan0Proto3.Validate())

	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.2499999, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, lowerThan0Proto2.Validate())

	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.25, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto3.Validate())

	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.25, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto2.Validate())
}

func TestFloatNonStrictUpperBounds(t *testing.T) {
	higherThan1Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75111111, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, higherThan1Proto3.Validate())

	higherThan1Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75111111, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, higherThan1Proto2.Validate())

	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto3.Validate())

	equalTo0Proto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, equalTo0Proto2.Validate())
}

func TestStringNonEmpty(t *testing.T) {
	emptyStringProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, emptyStringProto3.Validate())

	emptyStringProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, emptyStringProto2.Validate())

	nonEmptyStringProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, nonEmptyStringProto3.Validate())

	nonEmptyStringProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, nonEmptyStringProto2.Validate())
}

func TestRepeatedEltsCount(t *testing.T) {
	notEnoughEltsProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 1, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, notEnoughEltsProto3.Validate())

	notEnoughEltsProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 1, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, notEnoughEltsProto2.Validate())

	tooManyEltsProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 14, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, tooManyEltsProto3.Validate())

	tooManyEltsProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 14, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, tooManyEltsProto2.Validate())

	validEltsCountProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, validEltsCountProto3.Validate())

	validEltsCountProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, validEltsCountProto2.Validate())
}

func TestMsgExist(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)

	someProto3.SomeEmbedded = nil
	assert.NoError(t, someProto3.Validate())

	someProto3.SomeEmbeddedExists = nil
	err := someProto3.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid field SomeEmbeddedExists:")
}

func TestStringLengthValidator(t *testing.T) {
	stringLengthErrorProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "abc456", stableBytes, uuid1, uuid4, 0, 0)
	assert.Error(t, stringLengthErrorProto3.Validate())

	stringLengthSuccess := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, stringLengthSuccess.Validate())
}

func TestBytesLengthValidator(t *testing.T) {
	stringLengthErrorProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "abc456", []byte("anc"), uuid1, uuid4, 0, 0)
	assert.Error(t, stringLengthErrorProto3.Validate())

	stringLengthSuccess := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)
	assert.NoError(t, stringLengthSuccess.Validate())
}

func TestValueIsInEnum(t *testing.T) {
	outOfTopLevelEnumProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 2, 0)
	assert.Error(t, outOfTopLevelEnumProto3.Validate())

	outOfTopLevelEnumProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 2, 0)
	assert.Error(t, outOfTopLevelEnumProto2.Validate())

	outOfEmbeddedEnumProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 2)
	assert.Error(t, outOfEmbeddedEnumProto3.Validate())

	outOfEmbeddedEnumProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 2)
	assert.Error(t, outOfEmbeddedEnumProto2.Validate())
}

func TestNestedError3(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)

	someProto3.SomeEmbeddedExists.SomeValue = 101 // should be less than 101
	err := someProto3.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid field SomeEmbeddedExists.SomeValue:")
}

func TestCustomError_Proto3(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, uuid4, 0, 0)

	someProto3.CustomErrorInt = 30
	err := someProto3.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid field CustomErrorInt: My Custom Error")
}

func TestMapAlwaysPassesUntilFixedProperly(t *testing.T) {
	example := &ValidatorMapMessage3{}
	assert.NoError(t, example.Validate())
}

func TestMapPrimitiveTypeValuesAreValidates(t *testing.T) {
	mapError := &ValidatorMapMessage3{
		SomeStringMap: map[string]string{
			"foo": "toolong",
		},
	}
	assert.Error(t, mapError.Validate())

	mapSucces := &ValidatorMapMessage3{
		SomeStringMap: map[string]string{
			"foo": "abba",
		},
	}
	assert.NoError(t, mapSucces.Validate())
}

func TestMapNestedStructsAreValidated(t *testing.T) {
	mapError1 := &ValidatorMapMessage3{
		SomeExtMap: map[string]*ValueType{
			"foo": {Something: "foo"},
			"bar": {Something: ""},
		},
	}
	assert.Error(t, mapError1.Validate())

	mapError2 := &ValidatorMapMessage3{
		SomeNestedMap: map[int32]*ValidatorMapMessage3_NestedType{
			0: nil,
		},
	}
	assert.Error(t, mapError2.Validate())

	mapSucces1 := &ValidatorMapMessage3{
		SomeExtMap: map[string]*ValueType{
			"foo": {Something: "foo"},
			"bar": {Something: "bar"},
		},
	}
	assert.NoError(t, mapSucces1.Validate())

	mapSucces := &ValidatorMapMessage3{
		SomeNestedMap: map[int32]*ValidatorMapMessage3_NestedType{
			0: {},
		},
	}
	assert.NoError(t, mapSucces.Validate())
}

func TestOneOf_Required(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
	}
	err := example.Validate()
	require.Error(t, err, "oneof.required should fail if none of the oneof fields are set")
	assert.Contains(t, err.Error(), "invalid field Something:")
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
	require.Error(t, err, "nested message in oneof should fail validation on ExternalMsg")
	assert.Contains(t, err.Error(), "OneMsg.Identifier", "error must err on the ExternalMsg.Identifier")
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
	require.Error(t, err, "nested message in oneof should fail validation on ThreeInt")
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
	assert.NoError(t, err, "This message should pass all validation")
}

func TestOneOf_Regex(t *testing.T) {
	example := &OneOfMessage3{
		SomeInt: 30,
		Something: &OneOfMessage3_FiveRegex{
			FiveRegex: "11", // fail
		},
	}
	err := example.Validate()
	require.Error(t, err, "regex applied to oneof field should fail validation on FiveRegex")
	assert.Contains(t, err.Error(), "FiveRegex", "error must err on the FiveRegex")

	example = &OneOfMessage3{
		SomeInt: 30,
		Something: &OneOfMessage3_FiveRegex{
			FiveRegex: "aaa", // pass
		},
	}
	err = example.Validate()
	assert.NoError(t, err, "This message should pass all validation")
}

func TestUUID4Validation(t *testing.T) {
	testcases := map[string]struct {
		uuid string
		fail bool
	}{
		"UUID1": {
			uuid: uuid1,
			fail: true,
		},
		"UUID4": {
			uuid: uuid4,
			fail: false,
		},
		"EmptyField": {
			uuid: "",
			fail: true,
		},
		"NonUUID1": {
			uuid: "66bb25e2-2e0d",
			fail: true,
		},
		"NonUUID2": {
			uuid: "1234abcd",
			fail: true,
		},
	}

	for name := range testcases {
		tc := testcases[name]
		t.Run(name, func(t *testing.T) {
			t.Run("Proto2", func(t *testing.T) {
				msg := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, tc.uuid, 0, 0)

				if tc.fail {
					assert.Error(t, msg.Validate())
				} else {
					assert.NoError(t, msg.Validate())
				}
			})

			t.Run("Proto3", func(t *testing.T) {
				msg := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes, uuid1, tc.uuid, 0, 0)

				if tc.fail {
					assert.Error(t, msg.Validate())
				} else {
					assert.NoError(t, msg.Validate())
				}
			})
		})
	}
}
