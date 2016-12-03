// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package validatortest

import (
	"strings"
	"testing"
)

func buildProto3(someString string, someInt uint32, identifier string, someValue int64, someDouble float64, someFloat float32) *ValidatorMessage3 {
	goodEmbeddedProto3 := &ValidatorMessage3_Embedded{
		Identifier: identifier,
		SomeValue:  someValue,
	}

	goodProto3 := &ValidatorMessage3{
		SomeString:                    someString,
		SomeStringRep:                 []string{someString, "xyz34"},
		SomeStringNoQuotes:            someString,
		SomeInt:                       someInt,
		SomeIntRep:                    []uint32{someInt, 12, 13, 14, 15, 16},
		SomeIntRepNonNull:             []uint32{someInt, 102},
		SomeEmbedded:                  nil,
		SomeEmbeddedNonNullable:       goodEmbeddedProto3,
		SomeEmbeddedExists:            goodEmbeddedProto3,
		SomeEmbeddedExistsNonNullable: goodEmbeddedProto3,
		SomeEmbeddedRep:               []*ValidatorMessage3_Embedded{goodEmbeddedProto3},
		SomeEmbeddedRepNonNullable:    []*ValidatorMessage3_Embedded{goodEmbeddedProto3},
		SomeDouble:                    someDouble,
		SomeDoubleRep:                 []float64{someDouble, 0.5, 0.55, 0.6},
		SomeDoubleRepNonNull:          []float64{someDouble, 0.5, 0.55, 0.6},
		SomeFloat:                     someFloat,
		SomeFloatRep:                  []float32{someFloat, 0.5, 0.55, 0.6},
		SomeFloatRepNonNull:           []float32{someFloat, 0.5, 0.55, 0.6},
	}
	return goodProto3
}

func buildProto2(someString string, someInt uint32, identifier string, someValue int64, someDouble float64, someFloat float32) *ValidatorMessage {
	goodEmbeddedProto2 := &ValidatorMessage_Embedded{
		Identifier: &identifier,
		SomeValue:  &someValue,
	}

	goodProto2 := &ValidatorMessage{
		StringReq:        &someString,
		StringReqNonNull: &someString,

		StringOpt:        &someString,
		StringOptNonNull: &someString,

		IntReq:        &someInt,
		IntReqNonNull: &someInt,
		IntRep:        []uint32{someInt, 12, 13, 14, 15, 16},
		IntRepNonNull: []uint32{someInt, 12, 13, 14, 15, 16},

		EmbeddedReq:            goodEmbeddedProto2,
		EmbeddedNonNull:        goodEmbeddedProto2,
		EmbeddedRep:            []*ValidatorMessage_Embedded{goodEmbeddedProto2},
		EmbeddedRepNonNullable: []*ValidatorMessage_Embedded{goodEmbeddedProto2},

		DoubleReq:        &someDouble,
		DoubleReqNonNull: &someDouble,
		DoubleRep:        []float64{someDouble, 0.5, 0.55, 0.6},
		DoubleRepNonNull: []float64{someDouble, 0.5, 0.55, 0.6},

		FloatReq:        &someFloat,
		FloatReqNonNull: &someFloat,
		FloatRep:        []float32{someFloat, 0.5, 0.55, 0.6},
		FloatRepNonNull: []float32{someFloat, 0.5, 0.55, 0.6},
	}
	return goodProto2
}

func TestGoodProto3(t *testing.T) {
	var err error
	goodProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5)
	err = goodProto3.Validate()
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestGoodProto2(t *testing.T) {
	var err error
	goodProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.5)
	err = goodProto2.Validate()
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestStringRegex(t *testing.T) {
	tooLong1Proto3 := buildProto3("toolong", 11, "abba", 99, 0.5, 0.5)
	if tooLong1Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong2Proto3 := buildProto3("-%ab", 11, "bad#", 99, 0.5, 0.5)
	if tooLong2Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong1Proto2 := buildProto2("toolong", 11, "abba", 99, 0.5, 0.5)
	if tooLong1Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong2Proto2 := buildProto2("-%ab", 11, "bad#", 99, 0.5, 0.5)
	if tooLong2Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestIntLowerBounds(t *testing.T) {
	lowerThan10Proto3 := buildProto3("-%ab", 9, "abba", 99, 0.5, 0.5)
	if lowerThan10Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan10Proto2 := buildProto2("-%ab", 9, "abba", 99, 0.5, 0.5)
	if lowerThan10Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", -1, 0.5, 0.5)
	if lowerThan0Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", -1, 0.5, 0.5)
	if lowerThan0Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestIntUpperBounds(t *testing.T) {
	greaterThan100Proto3 := buildProto3("-%ab", 11, "abba", 101, 0.5, 0.5)
	if greaterThan100Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan100Proto2 := buildProto2("-%ab", 11, "abba", 101, 0.5, 0.5)
	if greaterThan100Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestDoubleLowerBounds(t *testing.T) {
	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.3, 0.5)
	if lowerThan035EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.3, 0.5)
	if lowerThan035EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.300000001, 0.5)
	if greaterThan035EpsilonProto3.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
	greaterThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.300000001, 0.5)
	if greaterThan035EpsilonProto2.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
}

func TestDoubleUpperBounds(t *testing.T) {
	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.70000000001, 0.5)
	if greaterThan065EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.70000000001, 0.5)
	if greaterThan065EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.6999999999, 0.5)
	if lowerThan065EpsilonProto3.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
	lowerThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.6999999999, 0.5)
	if lowerThan065EpsilonProto2.Validate() != nil {
		t.Fatalf("unexpected fail in validator")
	}
}

func TestFloatLowerBounds(t *testing.T) {
	lowerThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.2999999)
	if lowerThan035EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.2999999)
	if lowerThan035EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan035EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.3000001)
	if err := greaterThan035EpsilonProto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	greaterThan035EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.3000001)
	if err := greaterThan035EpsilonProto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestFloatUpperBounds(t *testing.T) {
	greaterThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.7000001)
	if greaterThan065EpsilonProto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	greaterThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.7000001)
	if greaterThan065EpsilonProto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan065EpsilonProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.6999999)
	if err := lowerThan065EpsilonProto3.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
	lowerThan065EpsilonProto2 := buildProto2("-%ab", 11, "abba", 99, 0.5, 0.6999999)
	if err := lowerThan065EpsilonProto2.Validate(); err != nil {
		t.Fatalf("unexpected fail in validator %v", err)
	}
}

func TestMsgExist(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5)
	someProto3.SomeEmbedded = nil
	if err := someProto3.Validate(); err != nil {
		t.Fatalf("validate shouldn't fail on missing SomeEmbedded, not annotated")
	}
	someProto3.SomeEmbeddedExists = nil
	if err := someProto3.Validate(); err == nil {
		t.Fatalf("expected fail due to lacking SomeEmbeddedExists")
	}
}

func TestNestedError3(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5)
	someProto3.SomeEmbeddedExists.SomeValue = 101 // should be less than 101
	if err := someProto3.Validate(); err == nil {
		t.Fatalf("expected fail due to nested SomeEmbeddedExists.SomeValue being wrong")
	} else if !strings.HasPrefix(err.Error(), "invalid field SomeEmbeddedNonNullable.SomeValue:") {
		t.Fatalf("expected fieldError, got '%v'", err)
	}
}

func TestCustomError_Proto3(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99, 0.5, 0.5)
	someProto3.CustomErrorInt = 30
	expectedErr := "invalid field CustomErrorInt: My Custom Error"
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
