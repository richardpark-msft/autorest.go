//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AnimalNotFound.
func (a AnimalNotFound) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", a.Name)
	populate(objectMap, "reason", a.Reason)
	populate(objectMap, "someBaseProp", a.SomeBaseProp)
	objectMap["whatNotFound"] = "AnimalNotFound"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AnimalNotFound.
func (a *AnimalNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, &a.Name)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &a.Reason)
			delete(rawMsg, key)
		case "someBaseProp":
			err = unpopulate(val, &a.SomeBaseProp)
			delete(rawMsg, key)
		case "whatNotFound":
			err = unpopulate(val, &a.WhatNotFound)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LinkNotFound.
func (l LinkNotFound) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "reason", l.Reason)
	populate(objectMap, "someBaseProp", l.SomeBaseProp)
	objectMap["whatNotFound"] = "InvalidResourceLink"
	populate(objectMap, "whatSubAddress", l.WhatSubAddress)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LinkNotFound.
func (l *LinkNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "reason":
			err = unpopulate(val, &l.Reason)
			delete(rawMsg, key)
		case "someBaseProp":
			err = unpopulate(val, &l.SomeBaseProp)
			delete(rawMsg, key)
		case "whatNotFound":
			err = unpopulate(val, &l.WhatNotFound)
			delete(rawMsg, key)
		case "whatSubAddress":
			err = unpopulate(val, &l.WhatSubAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetHungryOrThirstyError.
func (p PetHungryOrThirstyError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionResponse", p.ActionResponse)
	populate(objectMap, "errorMessage", p.ErrorMessage)
	objectMap["errorType"] = "PetHungryOrThirstyError"
	populate(objectMap, "hungryOrThirsty", p.HungryOrThirsty)
	populate(objectMap, "reason", p.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetHungryOrThirstyError.
func (p *PetHungryOrThirstyError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionResponse":
			err = unpopulate(val, &p.ActionResponse)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &p.ErrorMessage)
			delete(rawMsg, key)
		case "errorType":
			err = unpopulate(val, &p.ErrorType)
			delete(rawMsg, key)
		case "hungryOrThirsty":
			err = unpopulate(val, &p.HungryOrThirsty)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &p.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetSadError.
func (p PetSadError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionResponse", p.ActionResponse)
	populate(objectMap, "errorMessage", p.ErrorMessage)
	objectMap["errorType"] = "PetSadError"
	populate(objectMap, "reason", p.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetSadError.
func (p *PetSadError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionResponse":
			err = unpopulate(val, &p.ActionResponse)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &p.ErrorMessage)
			delete(rawMsg, key)
		case "errorType":
			err = unpopulate(val, &p.ErrorType)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &p.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}