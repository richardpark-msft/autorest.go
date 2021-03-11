// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ArrayGetEmptyOptions contains the optional parameters for the Array.GetEmpty method.
type ArrayGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// ArrayGetNotProvidedOptions contains the optional parameters for the Array.GetNotProvided method.
type ArrayGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// ArrayGetValidOptions contains the optional parameters for the Array.GetValid method.
type ArrayGetValidOptions struct {
	// placeholder for future optional parameters
}

// ArrayPutEmptyOptions contains the optional parameters for the Array.PutEmpty method.
type ArrayPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// ArrayPutValidOptions contains the optional parameters for the Array.PutValid method.
type ArrayPutValidOptions struct {
	// placeholder for future optional parameters
}

type ArrayWrapper struct {
	Array *[]string `json:"array,omitempty"`
}

// ArrayWrapperResponse is the response envelope for operations that return a ArrayWrapper type.
type ArrayWrapperResponse struct {
	ArrayWrapper *ArrayWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Basic struct {
	Color *CMYKColors `json:"color,omitempty"`

	// Basic Id
	ID *int32 `json:"id,omitempty"`

	// Name property with a very long description that does not fit on a single line and a line break.
	Name *string `json:"name,omitempty"`
}

// BasicGetEmptyOptions contains the optional parameters for the Basic.GetEmpty method.
type BasicGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// BasicGetInvalidOptions contains the optional parameters for the Basic.GetInvalid method.
type BasicGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// BasicGetNotProvidedOptions contains the optional parameters for the Basic.GetNotProvided method.
type BasicGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// BasicGetNullOptions contains the optional parameters for the Basic.GetNull method.
type BasicGetNullOptions struct {
	// placeholder for future optional parameters
}

// BasicGetValidOptions contains the optional parameters for the Basic.GetValid method.
type BasicGetValidOptions struct {
	// placeholder for future optional parameters
}

// BasicPutValidOptions contains the optional parameters for the Basic.PutValid method.
type BasicPutValidOptions struct {
	// placeholder for future optional parameters
}

// BasicResponse is the response envelope for operations that return a Basic type.
type BasicResponse struct {
	Basic *Basic

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type BooleanWrapper struct {
	FieldFalse *bool `json:"field_false,omitempty"`
	FieldTrue  *bool `json:"field_true,omitempty"`
}

// BooleanWrapperResponse is the response envelope for operations that return a BooleanWrapper type.
type BooleanWrapperResponse struct {
	BooleanWrapper *BooleanWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type ByteWrapper struct {
	Field *[]byte `json:"field,omitempty"`
}

// ByteWrapperResponse is the response envelope for operations that return a ByteWrapper type.
type ByteWrapperResponse struct {
	ByteWrapper *ByteWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Cat struct {
	Pet
	Color *string `json:"color,omitempty"`
	Hates *[]Dog  `json:"hates,omitempty"`
}

type Cookiecuttershark struct {
	Shark
}

// MarshalJSON implements the json.Marshaller interface for type Cookiecuttershark.
func (c Cookiecuttershark) MarshalJSON() ([]byte, error) {
	objectMap := c.Shark.marshalInternal("cookiecuttershark")
	return json.Marshal(objectMap)
}

type DateWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Leap  *time.Time `json:"leap,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DateWrapper.
func (d DateWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.Field != nil {
		objectMap["field"] = (*dateType)(d.Field)
	}
	if d.Leap != nil {
		objectMap["leap"] = (*dateType)(d.Leap)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateWrapper.
func (d *DateWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			if val != nil {
				var aux dateType
				err = json.Unmarshal(*val, &aux)
				d.Field = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "leap":
			if val != nil {
				var aux dateType
				err = json.Unmarshal(*val, &aux)
				d.Leap = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DateWrapperResponse is the response envelope for operations that return a DateWrapper type.
type DateWrapperResponse struct {
	DateWrapper *DateWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type DatetimeWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DatetimeWrapper.
func (d DatetimeWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.Field != nil {
		objectMap["field"] = (*timeRFC3339)(d.Field)
	}
	if d.Now != nil {
		objectMap["now"] = (*timeRFC3339)(d.Now)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatetimeWrapper.
func (d *DatetimeWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				d.Field = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "now":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				d.Now = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DatetimeWrapperResponse is the response envelope for operations that return a DatetimeWrapper type.
type DatetimeWrapperResponse struct {
	DatetimeWrapper *DatetimeWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Datetimerfc1123Wrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Datetimerfc1123Wrapper.
func (d Datetimerfc1123Wrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.Field != nil {
		objectMap["field"] = (*timeRFC1123)(d.Field)
	}
	if d.Now != nil {
		objectMap["now"] = (*timeRFC1123)(d.Now)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Datetimerfc1123Wrapper.
func (d *Datetimerfc1123Wrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			if val != nil {
				var aux timeRFC1123
				err = json.Unmarshal(*val, &aux)
				d.Field = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "now":
			if val != nil {
				var aux timeRFC1123
				err = json.Unmarshal(*val, &aux)
				d.Now = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Datetimerfc1123WrapperResponse is the response envelope for operations that return a Datetimerfc1123Wrapper type.
type Datetimerfc1123WrapperResponse struct {
	Datetimerfc1123Wrapper *Datetimerfc1123Wrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DictionaryGetEmptyOptions contains the optional parameters for the Dictionary.GetEmpty method.
type DictionaryGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// DictionaryGetNotProvidedOptions contains the optional parameters for the Dictionary.GetNotProvided method.
type DictionaryGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// DictionaryGetNullOptions contains the optional parameters for the Dictionary.GetNull method.
type DictionaryGetNullOptions struct {
	// placeholder for future optional parameters
}

// DictionaryGetValidOptions contains the optional parameters for the Dictionary.GetValid method.
type DictionaryGetValidOptions struct {
	// placeholder for future optional parameters
}

// DictionaryPutEmptyOptions contains the optional parameters for the Dictionary.PutEmpty method.
type DictionaryPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// DictionaryPutValidOptions contains the optional parameters for the Dictionary.PutValid method.
type DictionaryPutValidOptions struct {
	// placeholder for future optional parameters
}

type DictionaryWrapper struct {
	// Dictionary of
	DefaultProgram *map[string]string `json:"defaultProgram,omitempty"`
}

// DictionaryWrapperResponse is the response envelope for operations that return a DictionaryWrapper type.
type DictionaryWrapperResponse struct {
	DictionaryWrapper *DictionaryWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Dog struct {
	Pet
	Food *string `json:"food,omitempty"`
}

// DotFishClassification provides polymorphic access to related types.
// Call the interface's GetDotFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DotFish, *DotSalmon
type DotFishClassification interface {
	// GetDotFish() returns the DotFish content of the underlying type.
	GetDotFish() *DotFish
}

type DotFish struct {
	FishType *string `json:"fish.type,omitempty"`
	Species  *string `json:"species,omitempty"`
}

// GetDotFish implements the DotFishClassification interface for type DotFish.
func (d *DotFish) GetDotFish() *DotFish { return d }

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFish.
func (d *DotFish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return d.unmarshalInternal(rawMsg)
}

func (d DotFish) marshalInternal(discValue string) map[string]interface{} {
	objectMap := make(map[string]interface{})
	d.FishType = &discValue
	objectMap["fish.type"] = d.FishType
	if d.Species != nil {
		objectMap["species"] = d.Species
	}
	return objectMap
}

func (d *DotFish) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fish.type":
			if val != nil {
				err = json.Unmarshal(*val, &d.FishType)
			}
			delete(rawMsg, key)
		case "species":
			if val != nil {
				err = json.Unmarshal(*val, &d.Species)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DotFishMarket struct {
	Fishes       *[]DotFishClassification `json:"fishes,omitempty"`
	Salmons      *[]DotSalmon             `json:"salmons,omitempty"`
	SampleFish   DotFishClassification    `json:"sampleFish,omitempty"`
	SampleSalmon *DotSalmon               `json:"sampleSalmon,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishMarket.
func (d *DotFishMarket) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishes":
			if val != nil {
				d.Fishes, err = unmarshalDotFishClassificationArray(*val)
			}
			delete(rawMsg, key)
		case "salmons":
			if val != nil {
				err = json.Unmarshal(*val, &d.Salmons)
			}
			delete(rawMsg, key)
		case "sampleFish":
			if val != nil {
				d.SampleFish, err = unmarshalDotFishClassification(*val)
			}
			delete(rawMsg, key)
		case "sampleSalmon":
			if val != nil {
				err = json.Unmarshal(*val, &d.SampleSalmon)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DotFishMarketResponse is the response envelope for operations that return a DotFishMarket type.
type DotFishMarketResponse struct {
	DotFishMarket *DotFishMarket

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DotFishResponse is the response envelope for operations that return a DotFish type.
type DotFishResponse struct {
	DotFish DotFishClassification

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishResponse.
func (d *DotFishResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalDotFishClassification(data)
	if err != nil {
		return err
	}
	d.DotFish = t
	return nil
}

type DotSalmon struct {
	DotFish
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DotSalmon.
func (d DotSalmon) MarshalJSON() ([]byte, error) {
	objectMap := d.DotFish.marshalInternal("DotSalmon")
	if d.Iswild != nil {
		objectMap["iswild"] = d.Iswild
	}
	if d.Location != nil {
		objectMap["location"] = d.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotSalmon.
func (d *DotSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "iswild":
			if val != nil {
				err = json.Unmarshal(*val, &d.Iswild)
			}
			delete(rawMsg, key)
		case "location":
			if val != nil {
				err = json.Unmarshal(*val, &d.Location)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return d.DotFish.unmarshalInternal(rawMsg)
}

type DoubleWrapper struct {
	Field1                                                                          *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

// DoubleWrapperResponse is the response envelope for operations that return a DoubleWrapper type.
type DoubleWrapperResponse struct {
	DoubleWrapper *DoubleWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type DurationWrapper struct {
	Field *string `json:"field,omitempty"`
}

// DurationWrapperResponse is the response envelope for operations that return a DurationWrapper type.
type DurationWrapperResponse struct {
	DurationWrapper *DurationWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// FishClassification provides polymorphic access to related types.
// Call the interface's GetFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Fish, *Cookiecuttershark, *Goblinshark, *Salmon, *Sawshark, *Shark, *SmartSalmon
type FishClassification interface {
	// GetFish() returns the Fish content of the underlying type.
	GetFish() *Fish
}

type Fish struct {
	Fishtype *string               `json:"fishtype,omitempty"`
	Length   *float32              `json:"length,omitempty"`
	Siblings *[]FishClassification `json:"siblings,omitempty"`
	Species  *string               `json:"species,omitempty"`
}

// GetFish implements the FishClassification interface for type Fish.
func (f *Fish) GetFish() *Fish { return f }

// UnmarshalJSON implements the json.Unmarshaller interface for type Fish.
func (f *Fish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return f.unmarshalInternal(rawMsg)
}

func (f Fish) marshalInternal(discValue string) map[string]interface{} {
	objectMap := make(map[string]interface{})
	f.Fishtype = &discValue
	objectMap["fishtype"] = f.Fishtype
	if f.Length != nil {
		objectMap["length"] = f.Length
	}
	if f.Siblings != nil {
		objectMap["siblings"] = f.Siblings
	}
	if f.Species != nil {
		objectMap["species"] = f.Species
	}
	return objectMap
}

func (f *Fish) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			if val != nil {
				err = json.Unmarshal(*val, &f.Fishtype)
			}
			delete(rawMsg, key)
		case "length":
			if val != nil {
				err = json.Unmarshal(*val, &f.Length)
			}
			delete(rawMsg, key)
		case "siblings":
			if val != nil {
				f.Siblings, err = unmarshalFishClassificationArray(*val)
			}
			delete(rawMsg, key)
		case "species":
			if val != nil {
				err = json.Unmarshal(*val, &f.Species)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// FishResponse is the response envelope for operations that return a Fish type.
type FishResponse struct {
	Fish FishClassification

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FishResponse.
func (f *FishResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalFishClassification(data)
	if err != nil {
		return err
	}
	f.Fish = t
	return nil
}

// FlattencomplexGetValidOptions contains the optional parameters for the Flattencomplex.GetValid method.
type FlattencomplexGetValidOptions struct {
	// placeholder for future optional parameters
}

type FloatWrapper struct {
	Field1 *float32 `json:"field1,omitempty"`
	Field2 *float32 `json:"field2,omitempty"`
}

// FloatWrapperResponse is the response envelope for operations that return a FloatWrapper type.
type FloatWrapperResponse struct {
	FloatWrapper *FloatWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Goblinshark struct {
	Shark
	// Colors possible
	Color   *GoblinSharkColor `json:"color,omitempty"`
	Jawsize *int32            `json:"jawsize,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Goblinshark.
func (g Goblinshark) MarshalJSON() ([]byte, error) {
	objectMap := g.Shark.marshalInternal("goblin")
	if g.Color != nil {
		objectMap["color"] = g.Color
	}
	if g.Jawsize != nil {
		objectMap["jawsize"] = g.Jawsize
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Goblinshark.
func (g *Goblinshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "color":
			if val != nil {
				err = json.Unmarshal(*val, &g.Color)
			}
			delete(rawMsg, key)
		case "jawsize":
			if val != nil {
				err = json.Unmarshal(*val, &g.Jawsize)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return g.Shark.unmarshalInternal(rawMsg)
}

// InheritanceGetValidOptions contains the optional parameters for the Inheritance.GetValid method.
type InheritanceGetValidOptions struct {
	// placeholder for future optional parameters
}

// InheritancePutValidOptions contains the optional parameters for the Inheritance.PutValid method.
type InheritancePutValidOptions struct {
	// placeholder for future optional parameters
}

type IntWrapper struct {
	Field1 *int32 `json:"field1,omitempty"`
	Field2 *int32 `json:"field2,omitempty"`
}

// IntWrapperResponse is the response envelope for operations that return a IntWrapper type.
type IntWrapperResponse struct {
	IntWrapper *IntWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type LongWrapper struct {
	Field1 *int64 `json:"field1,omitempty"`
	Field2 *int64 `json:"field2,omitempty"`
}

// LongWrapperResponse is the response envelope for operations that return a LongWrapper type.
type LongWrapperResponse struct {
	LongWrapper *LongWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type MyBaseHelperType struct {
	PropBH1 *string `json:"propBH1,omitempty"`
}

// MyBaseTypeClassification provides polymorphic access to related types.
// Call the interface's GetMyBaseType() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MyBaseType, *MyDerivedType
type MyBaseTypeClassification interface {
	// GetMyBaseType() returns the MyBaseType content of the underlying type.
	GetMyBaseType() *MyBaseType
}

type MyBaseType struct {
	Helper *MyBaseHelperType `json:"helper,omitempty"`
	Kind   *MyKind           `json:"kind,omitempty"`
	PropB1 *string           `json:"propB1,omitempty"`
}

// GetMyBaseType implements the MyBaseTypeClassification interface for type MyBaseType.
func (m *MyBaseType) GetMyBaseType() *MyBaseType { return m }

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseType.
func (m *MyBaseType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return m.unmarshalInternal(rawMsg)
}

func (m MyBaseType) marshalInternal(discValue MyKind) map[string]interface{} {
	objectMap := make(map[string]interface{})
	if m.Helper != nil {
		objectMap["helper"] = m.Helper
	}
	m.Kind = &discValue
	objectMap["kind"] = m.Kind
	if m.PropB1 != nil {
		objectMap["propB1"] = m.PropB1
	}
	return objectMap
}

func (m *MyBaseType) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helper":
			if val != nil {
				err = json.Unmarshal(*val, &m.Helper)
			}
			delete(rawMsg, key)
		case "kind":
			if val != nil {
				err = json.Unmarshal(*val, &m.Kind)
			}
			delete(rawMsg, key)
		case "propB1":
			if val != nil {
				err = json.Unmarshal(*val, &m.PropB1)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MyBaseTypeResponse is the response envelope for operations that return a MyBaseType type.
type MyBaseTypeResponse struct {
	MyBaseType MyBaseTypeClassification

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseTypeResponse.
func (m *MyBaseTypeResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalMyBaseTypeClassification(data)
	if err != nil {
		return err
	}
	m.MyBaseType = t
	return nil
}

type MyDerivedType struct {
	MyBaseType
	PropD1 *string `json:"propD1,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MyDerivedType.
func (m MyDerivedType) MarshalJSON() ([]byte, error) {
	objectMap := m.MyBaseType.marshalInternal(MyKindKind1)
	if m.PropD1 != nil {
		objectMap["propD1"] = m.PropD1
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyDerivedType.
func (m *MyDerivedType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "propD1":
			if val != nil {
				err = json.Unmarshal(*val, &m.PropD1)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return m.MyBaseType.unmarshalInternal(rawMsg)
}

type Pet struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PolymorphicrecursiveGetValidOptions contains the optional parameters for the Polymorphicrecursive.GetValid method.
type PolymorphicrecursiveGetValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphicrecursivePutValidOptions contains the optional parameters for the Polymorphicrecursive.PutValid method.
type PolymorphicrecursivePutValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetComplicatedOptions contains the optional parameters for the Polymorphism.GetComplicated method.
type PolymorphismGetComplicatedOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetComposedWithDiscriminatorOptions contains the optional parameters for the Polymorphism.GetComposedWithDiscriminator method.
type PolymorphismGetComposedWithDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetComposedWithoutDiscriminatorOptions contains the optional parameters for the Polymorphism.GetComposedWithoutDiscriminator method.
type PolymorphismGetComposedWithoutDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetDotSyntaxOptions contains the optional parameters for the Polymorphism.GetDotSyntax method.
type PolymorphismGetDotSyntaxOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetValidOptions contains the optional parameters for the Polymorphism.GetValid method.
type PolymorphismGetValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutComplicatedOptions contains the optional parameters for the Polymorphism.PutComplicated method.
type PolymorphismPutComplicatedOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutMissingDiscriminatorOptions contains the optional parameters for the Polymorphism.PutMissingDiscriminator method.
type PolymorphismPutMissingDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutValidMissingRequiredOptions contains the optional parameters for the Polymorphism.PutValidMissingRequired method.
type PolymorphismPutValidMissingRequiredOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutValidOptions contains the optional parameters for the Polymorphism.PutValid method.
type PolymorphismPutValidOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetBoolOptions contains the optional parameters for the Primitive.GetBool method.
type PrimitiveGetBoolOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetByteOptions contains the optional parameters for the Primitive.GetByte method.
type PrimitiveGetByteOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDateOptions contains the optional parameters for the Primitive.GetDate method.
type PrimitiveGetDateOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDateTimeOptions contains the optional parameters for the Primitive.GetDateTime method.
type PrimitiveGetDateTimeOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDateTimeRFC1123Options contains the optional parameters for the Primitive.GetDateTimeRFC1123 method.
type PrimitiveGetDateTimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDoubleOptions contains the optional parameters for the Primitive.GetDouble method.
type PrimitiveGetDoubleOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDurationOptions contains the optional parameters for the Primitive.GetDuration method.
type PrimitiveGetDurationOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetFloatOptions contains the optional parameters for the Primitive.GetFloat method.
type PrimitiveGetFloatOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetIntOptions contains the optional parameters for the Primitive.GetInt method.
type PrimitiveGetIntOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetLongOptions contains the optional parameters for the Primitive.GetLong method.
type PrimitiveGetLongOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetStringOptions contains the optional parameters for the Primitive.GetString method.
type PrimitiveGetStringOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutBoolOptions contains the optional parameters for the Primitive.PutBool method.
type PrimitivePutBoolOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutByteOptions contains the optional parameters for the Primitive.PutByte method.
type PrimitivePutByteOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDateOptions contains the optional parameters for the Primitive.PutDate method.
type PrimitivePutDateOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDateTimeOptions contains the optional parameters for the Primitive.PutDateTime method.
type PrimitivePutDateTimeOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDateTimeRFC1123Options contains the optional parameters for the Primitive.PutDateTimeRFC1123 method.
type PrimitivePutDateTimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// PrimitivePutDoubleOptions contains the optional parameters for the Primitive.PutDouble method.
type PrimitivePutDoubleOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDurationOptions contains the optional parameters for the Primitive.PutDuration method.
type PrimitivePutDurationOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutFloatOptions contains the optional parameters for the Primitive.PutFloat method.
type PrimitivePutFloatOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutIntOptions contains the optional parameters for the Primitive.PutInt method.
type PrimitivePutIntOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutLongOptions contains the optional parameters for the Primitive.PutLong method.
type PrimitivePutLongOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutStringOptions contains the optional parameters for the Primitive.PutString method.
type PrimitivePutStringOptions struct {
	// placeholder for future optional parameters
}

type ReadonlyObj struct {
	// ID - READ-ONLY
	ID   *string `json:"id,omitempty" azure:"ro"`
	Size *int32  `json:"size,omitempty"`
}

// ReadonlyObjResponse is the response envelope for operations that return a ReadonlyObj type.
type ReadonlyObjResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	ReadonlyObj *ReadonlyObj
}

// ReadonlypropertyGetValidOptions contains the optional parameters for the Readonlyproperty.GetValid method.
type ReadonlypropertyGetValidOptions struct {
	// placeholder for future optional parameters
}

// ReadonlypropertyPutValidOptions contains the optional parameters for the Readonlyproperty.PutValid method.
type ReadonlypropertyPutValidOptions struct {
	// placeholder for future optional parameters
}

// SalmonClassification provides polymorphic access to related types.
// Call the interface's GetSalmon() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Salmon, *SmartSalmon
type SalmonClassification interface {
	FishClassification
	// GetSalmon() returns the Salmon content of the underlying type.
	GetSalmon() *Salmon
}

type Salmon struct {
	Fish
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

// GetSalmon implements the SalmonClassification interface for type Salmon.
func (s *Salmon) GetSalmon() *Salmon { return s }

// MarshalJSON implements the json.Marshaller interface for type Salmon.
func (s Salmon) MarshalJSON() ([]byte, error) {
	objectMap := s.marshalInternal("salmon")
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Salmon.
func (s *Salmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return s.unmarshalInternal(rawMsg)
}

func (s Salmon) marshalInternal(discValue string) map[string]interface{} {
	objectMap := s.Fish.marshalInternal(discValue)
	if s.Iswild != nil {
		objectMap["iswild"] = s.Iswild
	}
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	return objectMap
}

func (s *Salmon) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "iswild":
			if val != nil {
				err = json.Unmarshal(*val, &s.Iswild)
			}
			delete(rawMsg, key)
		case "location":
			if val != nil {
				err = json.Unmarshal(*val, &s.Location)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return s.Fish.unmarshalInternal(rawMsg)
}

// SalmonResponse is the response envelope for operations that return a Salmon type.
type SalmonResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Salmon      SalmonClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SalmonResponse.
func (s *SalmonResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalSalmonClassification(data)
	if err != nil {
		return err
	}
	s.Salmon = t
	return nil
}

type Sawshark struct {
	Shark
	Picture *[]byte `json:"picture,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Sawshark.
func (s Sawshark) MarshalJSON() ([]byte, error) {
	objectMap := s.Shark.marshalInternal("sawshark")
	if s.Picture != nil {
		objectMap["picture"] = s.Picture
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Sawshark.
func (s *Sawshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "picture":
			if val != nil {
				err = json.Unmarshal(*val, &s.Picture)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return s.Shark.unmarshalInternal(rawMsg)
}

// SharkClassification provides polymorphic access to related types.
// Call the interface's GetShark() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Shark, *Cookiecuttershark, *Goblinshark, *Sawshark
type SharkClassification interface {
	FishClassification
	// GetShark() returns the Shark content of the underlying type.
	GetShark() *Shark
}

type Shark struct {
	Fish
	Age      *int32     `json:"age,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}

// GetShark implements the SharkClassification interface for type Shark.
func (s *Shark) GetShark() *Shark { return s }

// MarshalJSON implements the json.Marshaller interface for type Shark.
func (s Shark) MarshalJSON() ([]byte, error) {
	objectMap := s.marshalInternal("shark")
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Shark.
func (s *Shark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return s.unmarshalInternal(rawMsg)
}

func (s Shark) marshalInternal(discValue string) map[string]interface{} {
	objectMap := s.Fish.marshalInternal(discValue)
	if s.Age != nil {
		objectMap["age"] = s.Age
	}
	if s.Birthday != nil {
		objectMap["birthday"] = (*timeRFC3339)(s.Birthday)
	}
	return objectMap
}

func (s *Shark) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			if val != nil {
				err = json.Unmarshal(*val, &s.Age)
			}
			delete(rawMsg, key)
		case "birthday":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.Birthday = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return s.Fish.unmarshalInternal(rawMsg)
}

type Siamese struct {
	Cat
	Breed *string `json:"breed,omitempty"`
}

// SiameseResponse is the response envelope for operations that return a Siamese type.
type SiameseResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Siamese     *Siamese
}

type SmartSalmon struct {
	Salmon
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]interface{}
	CollegeDegree        *string `json:"college_degree,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SmartSalmon.
func (s SmartSalmon) MarshalJSON() ([]byte, error) {
	objectMap := s.Salmon.marshalInternal("smart_salmon")
	if s.CollegeDegree != nil {
		objectMap["college_degree"] = s.CollegeDegree
	}
	if s.AdditionalProperties != nil {
		for key, val := range *s.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmartSalmon.
func (s *SmartSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "college_degree":
			if val != nil {
				err = json.Unmarshal(*val, &s.CollegeDegree)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := s.Salmon.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		if s.AdditionalProperties == nil {
			s.AdditionalProperties = &map[string]interface{}{}
		}
		if val != nil {
			var aux interface{}
			err = json.Unmarshal(*val, &aux)
			(*s.AdditionalProperties)[key] = aux
		}
		delete(rawMsg, key)
		if err != nil {
			return err
		}
	}
	return nil
}

type StringWrapper struct {
	Empty *string `json:"empty,omitempty"`
	Field *string `json:"field,omitempty"`
	Null  *string `json:"null,omitempty"`
}

// StringWrapperResponse is the response envelope for operations that return a StringWrapper type.
type StringWrapperResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse   *http.Response
	StringWrapper *StringWrapper
}
