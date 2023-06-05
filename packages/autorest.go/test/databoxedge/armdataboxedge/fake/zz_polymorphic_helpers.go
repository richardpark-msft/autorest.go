//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armdataboxedge"
	"encoding/json"
)

func unmarshalAddonClassification(rawMsg json.RawMessage) (armdataboxedge.AddonClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armdataboxedge.AddonClassification
	switch m["kind"] {
	case string(armdataboxedge.AddonTypeArcForKubernetes):
		b = &armdataboxedge.ArcAddon{}
	case string(armdataboxedge.AddonTypeIotEdge):
		b = &armdataboxedge.IoTAddon{}
	default:
		b = &armdataboxedge.Addon{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAddonClassificationArray(rawMsg json.RawMessage) ([]armdataboxedge.AddonClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armdataboxedge.AddonClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAddonClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalRoleClassification(rawMsg json.RawMessage) (armdataboxedge.RoleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armdataboxedge.RoleClassification
	switch m["kind"] {
	case string(armdataboxedge.RoleTypesCloudEdgeManagement):
		b = &armdataboxedge.CloudEdgeManagementRole{}
	case string(armdataboxedge.RoleTypesIOT):
		b = &armdataboxedge.IoTRole{}
	case string(armdataboxedge.RoleTypesKubernetes):
		b = &armdataboxedge.KubernetesRole{}
	case string(armdataboxedge.RoleTypesMEC):
		b = &armdataboxedge.MECRole{}
	default:
		b = &armdataboxedge.Role{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRoleClassificationArray(rawMsg json.RawMessage) ([]armdataboxedge.RoleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armdataboxedge.RoleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRoleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTriggerClassification(rawMsg json.RawMessage) (armdataboxedge.TriggerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armdataboxedge.TriggerClassification
	switch m["kind"] {
	case string(armdataboxedge.TriggerEventTypeFileEvent):
		b = &armdataboxedge.FileEventTrigger{}
	case string(armdataboxedge.TriggerEventTypePeriodicTimerEvent):
		b = &armdataboxedge.PeriodicTimerEventTrigger{}
	default:
		b = &armdataboxedge.Trigger{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTriggerClassificationArray(rawMsg json.RawMessage) ([]armdataboxedge.TriggerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armdataboxedge.TriggerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTriggerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
