package generator

type validatorType uint8

const (
	unknownValidator validatorType = iota
	floatEpsilonValidator
	floatGtValidator
	floatGteValidator
	floatLtValidator
	floatLteValidator
	humanErrorValidator
	intGtValidator
	intLtValidator
	isInEnumValidator
	lengthEqValidator
	lengthGtValidator
	lengthLtValidator
	msgExistsValidator
	regexValidator
	repeatedCountMaxValidator
	repeatedCountMinValidator
	stringNotEmptyValidator
	uuidVerValidator
)

var validatorTypeToString = map[validatorType]string{
	unknownValidator:          "",
	floatEpsilonValidator:     "FloatEpsilon",
	floatGtValidator:          "FloatGt",
	floatGteValidator:         "FloatGte",
	floatLtValidator:          "FloatLt",
	floatLteValidator:         "FloatLte",
	humanErrorValidator:       "HumanError",
	intGtValidator:            "IntGt",
	intLtValidator:            "IntLt",
	isInEnumValidator:         "IsInEnum",
	lengthEqValidator:         "LengthEq",
	lengthGtValidator:         "LengthGt",
	lengthLtValidator:         "LengthLt",
	msgExistsValidator:        "MsgExists",
	regexValidator:            "Regex",
	repeatedCountMaxValidator: "RepeatedCountMax",
	repeatedCountMinValidator: "RepeatedCountMin",
	stringNotEmptyValidator:   "StringNotEmpty",
	uuidVerValidator:          "UuidVer",
}
