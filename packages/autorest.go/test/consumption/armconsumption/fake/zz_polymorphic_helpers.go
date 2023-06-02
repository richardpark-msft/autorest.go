//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armconsumption"
	"encoding/json"
)

func unmarshalChargeSummaryClassification(rawMsg json.RawMessage) (armconsumption.ChargeSummaryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armconsumption.ChargeSummaryClassification
	switch m["kind"] {
	case string(armconsumption.ChargeSummaryKindLegacy):
		b = &armconsumption.LegacyChargeSummary{}
	case string(armconsumption.ChargeSummaryKindModern):
		b = &armconsumption.ModernChargeSummary{}
	default:
		b = &armconsumption.ChargeSummary{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalChargeSummaryClassificationArray(rawMsg json.RawMessage) ([]armconsumption.ChargeSummaryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armconsumption.ChargeSummaryClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalChargeSummaryClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalReservationRecommendationClassification(rawMsg json.RawMessage) (armconsumption.ReservationRecommendationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armconsumption.ReservationRecommendationClassification
	switch m["kind"] {
	case string(armconsumption.ReservationRecommendationKindLegacy):
		b = &armconsumption.LegacyReservationRecommendation{}
	case string(armconsumption.ReservationRecommendationKindModern):
		b = &armconsumption.ModernReservationRecommendation{}
	default:
		b = &armconsumption.ReservationRecommendation{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalReservationRecommendationClassificationArray(rawMsg json.RawMessage) ([]armconsumption.ReservationRecommendationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armconsumption.ReservationRecommendationClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalReservationRecommendationClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalUsageDetailClassification(rawMsg json.RawMessage) (armconsumption.UsageDetailClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armconsumption.UsageDetailClassification
	switch m["kind"] {
	case string(armconsumption.UsageDetailsKindLegacy):
		b = &armconsumption.LegacyUsageDetail{}
	case string(armconsumption.UsageDetailsKindModern):
		b = &armconsumption.ModernUsageDetail{}
	default:
		b = &armconsumption.UsageDetail{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalUsageDetailClassificationArray(rawMsg json.RawMessage) ([]armconsumption.UsageDetailClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armconsumption.UsageDetailClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalUsageDetailClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}