//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ArrayWrapper.
func (a ArrayWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "array", a.Array)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ByteWrapper.
func (b ByteWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateByteArray(objectMap, "field", b.Field, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ByteWrapper.
func (b *ByteWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = runtime.DecodeByteArray(string(val), &b.Field, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Cat.
func (c Cat) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "color", c.Color)
	populate(objectMap, "hates", c.Hates)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Cookiecuttershark.
func (c Cookiecuttershark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "age", c.Age)
	populateTimeRFC3339(objectMap, "birthday", c.Birthday)
	objectMap["fishtype"] = "cookiecuttershark"
	populate(objectMap, "length", c.Length)
	populate(objectMap, "siblings", c.Siblings)
	populate(objectMap, "species", c.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Cookiecuttershark.
func (c *Cookiecuttershark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &c.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &c.Birthday)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &c.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &c.Length)
			delete(rawMsg, key)
		case "siblings":
			c.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &c.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DateWrapper.
func (d DateWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateDateType(objectMap, "field", d.Field)
	populateDateType(objectMap, "leap", d.Leap)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateWrapper.
func (d *DateWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateDateType(val, "Field", &d.Field)
			delete(rawMsg, key)
		case "leap":
			err = unpopulateDateType(val, "Leap", &d.Leap)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DatetimeWrapper.
func (d DatetimeWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "field", d.Field)
	populateTimeRFC3339(objectMap, "now", d.Now)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatetimeWrapper.
func (d *DatetimeWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateTimeRFC3339(val, "Field", &d.Field)
			delete(rawMsg, key)
		case "now":
			err = unpopulateTimeRFC3339(val, "Now", &d.Now)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Datetimerfc1123Wrapper.
func (d Datetimerfc1123Wrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC1123(objectMap, "field", d.Field)
	populateTimeRFC1123(objectMap, "now", d.Now)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Datetimerfc1123Wrapper.
func (d *Datetimerfc1123Wrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateTimeRFC1123(val, "Field", &d.Field)
			delete(rawMsg, key)
		case "now":
			err = unpopulateTimeRFC1123(val, "Now", &d.Now)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DictionaryWrapper.
func (d DictionaryWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultProgram", d.DefaultProgram)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DotFishMarket.
func (d DotFishMarket) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fishes", d.Fishes)
	populate(objectMap, "salmons", d.Salmons)
	populate(objectMap, "sampleFish", d.SampleFish)
	populate(objectMap, "sampleSalmon", d.SampleSalmon)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishMarket.
func (d *DotFishMarket) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishes":
			d.Fishes, err = unmarshalDotFishClassificationArray(val)
			delete(rawMsg, key)
		case "salmons":
			err = unpopulate(val, "Salmons", &d.Salmons)
			delete(rawMsg, key)
		case "sampleFish":
			d.SampleFish, err = unmarshalDotFishClassification(val)
			delete(rawMsg, key)
		case "sampleSalmon":
			err = unpopulate(val, "SampleSalmon", &d.SampleSalmon)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DotSalmon.
func (d DotSalmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["fish.type"] = "DotSalmon"
	populate(objectMap, "iswild", d.Iswild)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "species", d.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotSalmon.
func (d *DotSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fish.type":
			err = unpopulate(val, "FishType", &d.FishType)
			delete(rawMsg, key)
		case "iswild":
			err = unpopulate(val, "Iswild", &d.Iswild)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &d.Location)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &d.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Fish.
func (f Fish) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["fishtype"] = f.Fishtype
	populate(objectMap, "length", f.Length)
	populate(objectMap, "siblings", f.Siblings)
	populate(objectMap, "species", f.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Fish.
func (f *Fish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			err = unpopulate(val, "Fishtype", &f.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &f.Length)
			delete(rawMsg, key)
		case "siblings":
			f.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &f.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Goblinshark.
func (g Goblinshark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "age", g.Age)
	populateTimeRFC3339(objectMap, "birthday", g.Birthday)
	populate(objectMap, "color", g.Color)
	objectMap["fishtype"] = "goblin"
	populate(objectMap, "jawsize", g.Jawsize)
	populate(objectMap, "length", g.Length)
	populate(objectMap, "siblings", g.Siblings)
	populate(objectMap, "species", g.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Goblinshark.
func (g *Goblinshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &g.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &g.Birthday)
			delete(rawMsg, key)
		case "color":
			err = unpopulate(val, "Color", &g.Color)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &g.Fishtype)
			delete(rawMsg, key)
		case "jawsize":
			err = unpopulate(val, "Jawsize", &g.Jawsize)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &g.Length)
			delete(rawMsg, key)
		case "siblings":
			g.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &g.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MyDerivedType.
func (m MyDerivedType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "helper", m.Helper)
	objectMap["kind"] = MyKindKind1
	populate(objectMap, "propB1", m.PropB1)
	populate(objectMap, "propD1", m.PropD1)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyDerivedType.
func (m *MyDerivedType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helper":
			err = unpopulate(val, "Helper", &m.Helper)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &m.Kind)
			delete(rawMsg, key)
		case "propB1":
			err = unpopulate(val, "PropB1", &m.PropB1)
			delete(rawMsg, key)
		case "propD1":
			err = unpopulate(val, "PropD1", &m.PropD1)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Salmon.
func (s Salmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["fishtype"] = "salmon"
	populate(objectMap, "iswild", s.Iswild)
	populate(objectMap, "length", s.Length)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Salmon.
func (s *Salmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "iswild":
			err = unpopulate(val, "Iswild", &s.Iswild)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Sawshark.
func (s Sawshark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "age", s.Age)
	populateTimeRFC3339(objectMap, "birthday", s.Birthday)
	objectMap["fishtype"] = "sawshark"
	populate(objectMap, "length", s.Length)
	populateByteArray(objectMap, "picture", s.Picture, runtime.Base64StdFormat)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Sawshark.
func (s *Sawshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &s.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &s.Birthday)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "picture":
			err = runtime.DecodeByteArray(string(val), &s.Picture, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Shark.
func (s Shark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "age", s.Age)
	populateTimeRFC3339(objectMap, "birthday", s.Birthday)
	objectMap["fishtype"] = "shark"
	populate(objectMap, "length", s.Length)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Shark.
func (s *Shark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &s.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &s.Birthday)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Siamese.
func (s Siamese) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "breed", s.Breed)
	populate(objectMap, "color", s.Color)
	populate(objectMap, "hates", s.Hates)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SmartSalmon.
func (s SmartSalmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "college_degree", s.CollegeDegree)
	objectMap["fishtype"] = "smart_salmon"
	populate(objectMap, "iswild", s.Iswild)
	populate(objectMap, "length", s.Length)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	if s.AdditionalProperties != nil {
		for key, val := range s.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmartSalmon.
func (s *SmartSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "college_degree":
			err = unpopulate(val, "CollegeDegree", &s.CollegeDegree)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "iswild":
			err = unpopulate(val, "Iswild", &s.Iswild)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		default:
			if s.AdditionalProperties == nil {
				s.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				s.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
