// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

type CMYKColors string

const (
	CMYKColorsBlacK   CMYKColors = "blacK"
	CMYKColorsCyan    CMYKColors = "cyan"
	CMYKColorsMagenta CMYKColors = "Magenta"
	CMYKColorsYELLOW  CMYKColors = "YELLOW"
)

// PossibleCMYKColorsValues returns the possible values for the CMYKColors const type.
func PossibleCMYKColorsValues() []CMYKColors {
	return []CMYKColors{
		CMYKColorsBlacK,
		CMYKColorsCyan,
		CMYKColorsMagenta,
		CMYKColorsYELLOW,
	}
}

// ToPtr() returns a *CMYKColors pointing to the current value.
func (c CMYKColors) ToPtr() *CMYKColors {
	return &c
}

// GoblinSharkColor - Colors possible
type GoblinSharkColor string

const (
	GoblinSharkColorBrown GoblinSharkColor = "brown"
	GoblinSharkColorGray  GoblinSharkColor = "gray"
	// GoblinSharkColorLowerRed - Lowercase RED
	GoblinSharkColorLowerRed GoblinSharkColor = "red"
	GoblinSharkColorPink     GoblinSharkColor = "pink"
	// GoblinSharkColorUpperRed - Uppercase RED
	GoblinSharkColorUpperRed GoblinSharkColor = "RED"
)

// PossibleGoblinSharkColorValues returns the possible values for the GoblinSharkColor const type.
func PossibleGoblinSharkColorValues() []GoblinSharkColor {
	return []GoblinSharkColor{
		GoblinSharkColorBrown,
		GoblinSharkColorGray,
		GoblinSharkColorLowerRed,
		GoblinSharkColorPink,
		GoblinSharkColorUpperRed,
	}
}

// ToPtr() returns a *GoblinSharkColor pointing to the current value.
func (c GoblinSharkColor) ToPtr() *GoblinSharkColor {
	return &c
}

type MyKind string

const (
	MyKindKind1 MyKind = "Kind1"
)

// PossibleMyKindValues returns the possible values for the MyKind const type.
func PossibleMyKindValues() []MyKind {
	return []MyKind{
		MyKindKind1,
	}
}

// ToPtr() returns a *MyKind pointing to the current value.
func (c MyKind) ToPtr() *MyKind {
	return &c
}
