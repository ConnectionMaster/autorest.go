// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

// FloatEnum - List of float enums
type FloatEnum float32

const (
	FloatEnumFourHundredFive3       FloatEnum = 405.3
	FloatEnumFourHundredSix2        FloatEnum = 406.2
	FloatEnumFourHundredThree4      FloatEnum = 403.4
	FloatEnumFourHundredTwentyNine1 FloatEnum = 429.1
	FloatEnumTwoHundred4            FloatEnum = 200.4
)

// PossibleFloatEnumValues returns the possible values for the FloatEnum const type.
func PossibleFloatEnumValues() []FloatEnum {
	return []FloatEnum{
		FloatEnumFourHundredFive3,
		FloatEnumFourHundredSix2,
		FloatEnumFourHundredThree4,
		FloatEnumFourHundredTwentyNine1,
		FloatEnumTwoHundred4,
	}
}

// ToPtr returns a *FloatEnum pointing to the current value.
func (c FloatEnum) ToPtr() *FloatEnum {
	return &c
}

// IntEnum - List of integer enums
type IntEnum int32

const (
	IntEnumFourHundredFive       IntEnum = 405
	IntEnumFourHundredSix        IntEnum = 406
	IntEnumFourHundredThree      IntEnum = 403
	IntEnumFourHundredTwentyNine IntEnum = 429
	IntEnumTwoHundred            IntEnum = 200
)

// PossibleIntEnumValues returns the possible values for the IntEnum const type.
func PossibleIntEnumValues() []IntEnum {
	return []IntEnum{
		IntEnumFourHundredFive,
		IntEnumFourHundredSix,
		IntEnumFourHundredThree,
		IntEnumFourHundredTwentyNine,
		IntEnumTwoHundred,
	}
}

// ToPtr returns a *IntEnum pointing to the current value.
func (c IntEnum) ToPtr() *IntEnum {
	return &c
}
