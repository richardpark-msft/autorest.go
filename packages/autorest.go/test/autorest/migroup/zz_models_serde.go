//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package migroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Cat.
func (c Cat) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "hisses", c.Hisses)
	populate(objectMap, "likesMilk", c.LikesMilk)
	populate(objectMap, "meows", c.Meows)
	populate(objectMap, "name", c.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Cat.
func (c *Cat) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "hisses":
			err = unpopulate(val, "Hisses", &c.Hisses)
			delete(rawMsg, key)
		case "likesMilk":
			err = unpopulate(val, "LikesMilk", &c.LikesMilk)
			delete(rawMsg, key)
		case "meows":
			err = unpopulate(val, "Meows", &c.Meows)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Feline.
func (f Feline) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "hisses", f.Hisses)
	populate(objectMap, "meows", f.Meows)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Feline.
func (f *Feline) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "hisses":
			err = unpopulate(val, "Hisses", &f.Hisses)
			delete(rawMsg, key)
		case "meows":
			err = unpopulate(val, "Meows", &f.Meows)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Horse.
func (h Horse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "isAShowHorse", h.IsAShowHorse)
	populate(objectMap, "name", h.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Horse.
func (h *Horse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "isAShowHorse":
			err = unpopulate(val, "IsAShowHorse", &h.IsAShowHorse)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &h.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Kitten.
func (k Kitten) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "eatsMiceYet", k.EatsMiceYet)
	populate(objectMap, "hisses", k.Hisses)
	populate(objectMap, "likesMilk", k.LikesMilk)
	populate(objectMap, "meows", k.Meows)
	populate(objectMap, "name", k.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Kitten.
func (k *Kitten) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "eatsMiceYet":
			err = unpopulate(val, "EatsMiceYet", &k.EatsMiceYet)
			delete(rawMsg, key)
		case "hisses":
			err = unpopulate(val, "Hisses", &k.Hisses)
			delete(rawMsg, key)
		case "likesMilk":
			err = unpopulate(val, "LikesMilk", &k.LikesMilk)
			delete(rawMsg, key)
		case "meows":
			err = unpopulate(val, "Meows", &k.Meows)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &k.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Pet.
func (p Pet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", p.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Pet.
func (p *Pet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
