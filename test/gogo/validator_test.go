// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package validatortest

import (
	fmt "fmt"
	"strconv"
	"testing"

	"github.com/TheThingsIndustries/go-proto-validators/errors"
)

var (
	stableBytes = make([]byte, 12)
)

func buildProto3(someString string, someInt uint32, identifier string, someValue int64, someDoubleStrict float64,
	someFloatStrict float32, someDouble float64, someFloat float32, nonEmptyString string, repeatedCount uint32, someStringLength string, someBytes []byte) ValidatorMessage3 {
	goodEmbeddedProto3 := &ValidatorMessage3_Embedded{
		Identifier: identifier,
		SomeValue:  someValue,
	}

	goodProto3 := &ValidatorMessage3{
		SomeString:    someString,
		SomeStringRep: []string{someString, "xyz34"},
		// SomeStringNoQuotes:  someString,
		// SomeStringUnescaped: someString,

		SomeInt:           someInt,
		SomeIntRep:        []uint32{someInt, 12, 13, 14, 15, 16},
		SomeIntRepNonNull: []uint32{someInt, 102},

		SomeEmbedded:               nil,
		SomeEmbeddedNonNullable:    *goodEmbeddedProto3,
		SomeEmbeddedExists:         goodEmbeddedProto3,
		SomeEmbeddedRep:            []*ValidatorMessage3_Embedded{goodEmbeddedProto3},
		SomeEmbeddedRepNonNullable: []ValidatorMessage3_Embedded{*goodEmbeddedProto3},

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
	}

	goodProto3.Repeated = make([]int32, repeatedCount, repeatedCount)

	return *goodProto3
}

func TestGoodProto3(t *testing.T) {
	var err error
	goodProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
	err = goodProto3.Validate([]string{"SomeString", "SomeInt", "StrictSomeDouble", "StrictSomeFloat", "SomeDouble", "SomeFloat", "SomeNonEmptyString", "SomeStringLtReq", "SomeStringGtReq", "SomeStringEqReq", "SomeBytesLtReq", "SomeBytesGtReq", "SomeBytesEqReq"})
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestMessages(t *testing.T) {
	test1Proto3 := buildProto3("toolong", 0, "1234", -1, 0.5, 0.5, 0.5, 0.5, "", 1, "1234567890", stableBytes)
	for i, tc := range []struct {
		FieldMaskPaths     []string
		ExpectedErrorField string
		ExpectedErrorType  errors.Types
	}{
		{
			[]string{"SomeString"},
			"SomeString",
			errors.Types_STRING_REGEX,
		},
		{
			[]string{"someEmbeddedExists.Identifier"},
			"someEmbeddedExists.Identifier",
			errors.Types_STRING_REGEX,
		},
		{
			[]string{"someEmbeddedNonNullable.Identifier"},
			"someEmbeddedNonNullable.Identifier",
			errors.Types_STRING_REGEX,
		},
		{
			[]string{"SomeInt"},
			"SomeInt",
			errors.Types_INT_GT,
		},
		{
			[]string{"someEmbeddedNonNullable.SomeValue"},
			"someEmbeddedNonNullable.SomeValue",
			errors.Types_INT_GT,
		},
		{
			[]string{"SomeNonEmptyString"},
			"SomeNonEmptyString",
			errors.Types_STRING_NOT_EMPTY,
		},
		{
			[]string{"SomeNonEmptyString"},
			"SomeNonEmptyString",
			errors.Types_STRING_NOT_EMPTY,
		},
		{
			[]string{"Repeated"},
			"Repeated",
			errors.Types_REPEATED_COUNT_MIN,
		},
	} {
		t.Run(fmt.Sprintf("Case%s", strconv.Itoa(i)), func(t *testing.T) {
			err := test1Proto3.Validate(tc.FieldMaskPaths)
			if err == nil {
				t.Fatal(fmt.Sprintf("Error %s Expected on field %s", tc.ExpectedErrorType, tc.ExpectedErrorField))
			} else if tc.ExpectedErrorField != errors.GetFieldName(err.Error()) || tc.ExpectedErrorType.String() != errors.GetType(err.Error()) {
				t.Fatal(err)
			}
		})
	}

	test2Proto3 := buildProto3("-", 9, "bad#", 101, 0.5, 0.5, 0.5, 0.5, "x", 6, "1234567890", stableBytes)
	for i, tc := range []struct {
		FieldMaskPaths     []string
		ExpectedErrorField string
		ExpectedErrorType  errors.Types
	}{
		{
			[]string{"someEmbeddedNonNullable.SomeValue"},
			"someEmbeddedNonNullable.SomeValue",
			errors.Types_INT_LT,
		},
		{
			[]string{"Repeated"},
			"Repeated",
			errors.Types_REPEATED_COUNT_MAX,
		},
	} {
		t.Run(fmt.Sprintf("AdditionalCases/%s", strconv.Itoa(i)), func(t *testing.T) {
			err := test2Proto3.Validate(tc.FieldMaskPaths)
			if err == nil {
				t.Fatal(fmt.Sprintf("Error %s Expected on field %s", tc.ExpectedErrorType, tc.ExpectedErrorField))
			} else if tc.ExpectedErrorField != errors.GetFieldName(err.Error()) || tc.ExpectedErrorType.String() != errors.GetType(err.Error()) {
				t.Fatal(err)
			}
		})
	}
	test3Proto3 := buildProto3("-", 9, "bad#", 101, 0.3, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes)

	// Nil FieldMask
	if err := test3Proto3.Validate([]string{}); err != nil {
		t.Fatal(err)
	}

	//Misc
	if err := test3Proto3.Validate([]string{"SomeEmbedded"}); err != nil {
		t.Fatal(err)
	}
	test3Proto3.SomeEmbeddedExists = nil
	if err := test3Proto3.Validate([]string{}); err == nil {
		t.Fatal("No Error when MSG_EXISTS expected")
	} else {
		t.Logf("Successfully errored: %s", err)
	}

	test3Proto3.SomeEmbedded = &ValidatorMessage3_Embedded{Identifier: "&*&#&$(#$*(#"}
	if err := test3Proto3.Validate([]string{"SomeEmbedded.Identifier"}); err == nil {
		t.Fatal("No Error when STRING_REGEX expected")
	} else {
		t.Logf("Successfully errored: %s", err)
	}

	test3Proto3.SomeEmbeddedRep = []*ValidatorMessage3_Embedded{&ValidatorMessage3_Embedded{Identifier: "&*&#&$(#$*(#"}}
	if err := test3Proto3.Validate([]string{"SomeEmbeddedRep.Identifier"}); err == nil {
		t.Fatal("No Error when STRING_REGEX expected")
	} else {
		t.Logf("Successfully errored: %s", err)
	}
}

func TestFloat(t *testing.T) {
	for i, tc := range []struct {
		Message            ValidatorMessage3
		FieldMaskPaths     []string
		ExpectedErrorField string
		ExpectedErrorType  errors.Types
	}{
		{
			buildProto3("toolong", 0, "1234", -1, 0.3, 0.5, 0.5, 0.5, "", 2, "1234567890", stableBytes),
			[]string{"StrictSomeDouble"},
			"StrictSomeDouble",
			errors.Types_FLOAT_ELIPSON,
		},
		{
			buildProto3("toolong", 0, "1234", -1, 0.70000000001, 0.5, 0.5, 0.5, "", 2, "1234567890", stableBytes),
			[]string{"StrictSomeDouble"},
			"StrictSomeDouble",
			errors.Types_FLOAT_ELIPSON,
		},
		{
			buildProto3("toolong", 0, "1234", -1, 0.2999999, 0.5, 0.5, 0.5, "", 2, "1234567890", stableBytes),
			[]string{"StrictSomeDouble"},
			"StrictSomeDouble",
			errors.Types_FLOAT_ELIPSON,
		},
		{
			buildProto3("toolong", 0, "1234", -1, 0.25, 0.5, 0.2499999, 0.5, "", 2, "1234567890", stableBytes),
			[]string{"SomeDouble"},
			"SomeDouble",
			errors.Types_FLOAT_GTE,
		},
		{
			buildProto3("toolong", 0, "1234", -1, 0.25, 0.5, 0.75111111, 0.5, "", 2, "1234567890", stableBytes),
			[]string{"SomeDouble"},
			"SomeDouble",
			errors.Types_FLOAT_LTE,
		},
		{
			buildProto3("toolong", 0, "1234", -1, 0.25, 0.5, 0.75111111, 0.2499999, "", 2, "1234567890", stableBytes),
			[]string{"SomeFloat"},
			"SomeFloat",
			errors.Types_FLOAT_GTE,
		},
		{
			buildProto3("toolong", 0, "1234", -1, 0.25, 0.5, 0.75111111, 0.75111111, "", 2, "1234567890", stableBytes),
			[]string{"SomeFloat"},
			"SomeFloat",
			errors.Types_FLOAT_LTE,
		},
	} {
		t.Run(fmt.Sprintf("AdditionalCases/%s", strconv.Itoa(i)), func(t *testing.T) {
			err := tc.Message.Validate(tc.FieldMaskPaths)
			if err == nil {
				t.Fatal(fmt.Sprintf("Error %s Expected on field %s", tc.ExpectedErrorType, tc.ExpectedErrorField))
			} else if tc.ExpectedErrorField != errors.GetFieldName(err.Error()) || tc.ExpectedErrorType.String() != errors.GetType(err.Error()) {
				t.Fatal(err)
			}
		})
	}
}

// func TestDoubleStrictUpperBounds(t *testing.T) {
// 	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.70000000001, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
// 	if greaterThan065EpsilonProto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.6999999999, 0.5, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
// 	if lowerThan065EpsilonProto3.Validate() != nil {
// 		t.Fatalf("unexpected fail in validator")
// 	}
// }

// func TestFloatStrictLowerBounds(t *testing.T) {
// 	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.2999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
// 	if lowerThan035EpsilonProto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.3000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
// 	if err := greaterThan035EpsilonProto3.Validate(); err != nil {
// 		t.Fatalf("unexpected fail in validator %v", err)
// 	}
// }

// func TestFloatStrictUpperBounds(t *testing.T) {
// 	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.7000001, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
// 	if greaterThan065EpsilonProto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.6999999, 0.5, 0.5, "x", 4, "1234567890", stableBytes)
// 	if err := lowerThan065EpsilonProto3.Validate(); err != nil {
// 		t.Fatalf("unexpected fail in validator %v", err)
// 	}
// }

// func TestDoubleNonStrictLowerBounds(t *testing.T) {
// 	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.2499999, 0.5, "x", 4, "1234567890", stableBytes)
// 	if lowerThan0Proto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.25, 0.5, "x", 4, "1234567890", stableBytes)
// 	if err := equalTo0Proto3.Validate(); err != nil {
// 		t.Fatalf("unexpected fail in validator %v", err)
// 	}
// }

// func TestDoubleNonStrictUpperBounds(t *testing.T) {
// 	higherThan1Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75111111, 0.5, "x", 4, "1234567890", stableBytes)
// 	if higherThan1Proto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.75, 0.5, "x", 4, "1234567890", stableBytes)
// 	if err := equalTo0Proto3.Validate(); err != nil {
// 		t.Fatalf("unexpected fail in validator %v", err)
// 	}
// }

// func TestFloatNonStrictLowerBounds(t *testing.T) {
// 	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.2499999, "x", 4, "1234567890", stableBytes)
// 	if lowerThan0Proto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.25, "x", 4, "1234567890", stableBytes)
// 	if err := equalTo0Proto3.Validate(); err != nil {
// 		t.Fatalf("unexpected fail in validator %v", err)
// 	}
// }

// func TestFloatNonStrictUpperBounds(t *testing.T) {
// 	higherThan1Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75111111, "x", 4, "1234567890", stableBytes)
// 	if higherThan1Proto3.Validate() == nil {
// 		t.Fatalf("expected fail in validator, but it didn't happen")
// 	}
// 	equalTo0Proto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5, 0.5, 0.75, "x", 4, "1234567890", stableBytes)
// 	if err := equalTo0Proto3.Validate(); err != nil {
// 		t.Fatalf("unexpected fail in validator %v", err)
// 	}
// }

// func TestMapAlwaysPassesUntilFixedProperly(t *testing.T) {
// 	example := &ValidatorMapMessage3{}
// 	if err := example.Validate(); err != nil {
// 		t.Fatalf("map validators should always pass")
// 	}
// }

// func TestOneOf_NestedMessage(t *testing.T) {
// 	example := &OneOfMessage3{
// 		SomeInt: 30,
// 		Type: &OneOfMessage3_OneMsg{
// 			OneMsg: &ExternalMsg{
// 				Identifier: "999", // bad
// 				SomeValue:  99,    // good
// 			},
// 		},
// 		Something: &OneOfMessage3_ThreeInt{
// 			ThreeInt: 100, // > 20
// 		},
// 	}
// 	err := example.Validate()
// 	assert.Error(t, err, "nested message in oneof should fail validation on ExternalMsg")
// 	assert.Contains(t, err.Error(), "OneMsg.Identifier", "error must err on the ExternalMsg.Identifier")
// }

// func TestOneOf_NestedInt(t *testing.T) {
// 	example := &OneOfMessage3{
// 		SomeInt: 30,
// 		Type: &OneOfMessage3_OneMsg{
// 			OneMsg: &ExternalMsg{
// 				Identifier: "abba", // good
// 				SomeValue:  99,     // good
// 			},
// 		},
// 		Something: &OneOfMessage3_ThreeInt{
// 			ThreeInt: 19, // > 20
// 		},
// 	}
// 	err := example.Validate()
// 	assert.Error(t, err, "nested message in oneof should fail validation on ThreeInt")
// 	assert.Contains(t, err.Error(), "ThreeInt", "error must err on the ThreeInt.ThreeInt")
// }

// func TestOneOf_Passes(t *testing.T) {
// 	example := &OneOfMessage3{
// 		SomeInt: 30,
// 		Type: &OneOfMessage3_OneMsg{
// 			OneMsg: &ExternalMsg{
// 				Identifier: "abba", // good
// 				SomeValue:  99,     // good
// 			},
// 		},
// 		Something: &OneOfMessage3_FourInt{
// 			FourInt: 101, // > 101
// 		},
// 	}
// 	err := example.Validate()
// 	assert.NoError(t, err, "This message should pass all validation")
// }
