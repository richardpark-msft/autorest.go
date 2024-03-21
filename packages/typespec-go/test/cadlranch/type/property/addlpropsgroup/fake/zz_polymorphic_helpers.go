// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"addlpropsgroup"
	"encoding/json"
)

func unmarshalExtendsUnknownAdditionalPropertiesDiscriminatedClassification(rawMsg json.RawMessage) (addlpropsgroup.ExtendsUnknownAdditionalPropertiesDiscriminatedClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b addlpropsgroup.ExtendsUnknownAdditionalPropertiesDiscriminatedClassification
	switch m["kind"] {
	case "derived":
		b = &addlpropsgroup.ExtendsUnknownAdditionalPropertiesDiscriminatedDerived{}
	default:
		b = &addlpropsgroup.ExtendsUnknownAdditionalPropertiesDiscriminated{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalIsUnknownAdditionalPropertiesDiscriminatedClassification(rawMsg json.RawMessage) (addlpropsgroup.IsUnknownAdditionalPropertiesDiscriminatedClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b addlpropsgroup.IsUnknownAdditionalPropertiesDiscriminatedClassification
	switch m["kind"] {
	case "derived":
		b = &addlpropsgroup.IsUnknownAdditionalPropertiesDiscriminatedDerived{}
	default:
		b = &addlpropsgroup.IsUnknownAdditionalPropertiesDiscriminated{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
