// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newPolymorphismClient() *PolymorphismClient {
	return NewPolymorphismClient(nil)
}

// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func TestPolymorphismGetComplicated(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetComplicated(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	salmon, ok := result.SalmonClassification.(*SmartSalmon)
	if !ok {
		t.Fatal("fish wasn't a smart salmon")
	}
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	expectedFish := Fish{
		Fishtype: to.StringPtr("smart_salmon"),
		Length:   to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.StringPtr("goblin"),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species: to.StringPtr("king"),
	}
	expectedSalmon := Salmon{
		Fishtype: to.StringPtr("smart_salmon"),
		Length:   to.Float32Ptr(1),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.StringPtr("goblin"),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species: to.StringPtr("king"),
	}
	if r := cmp.Diff(salmon, &SmartSalmon{
		Fishtype: to.StringPtr("smart_salmon"),
		Length:   to.Float32Ptr(1),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.StringPtr("goblin"),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species: to.StringPtr("king"),
		AdditionalProperties: map[string]interface{}{
			"additionalProperty1": float64(1),
			"additionalProperty2": false,
			"additionalProperty3": "hello",
			"additionalProperty4": map[string]interface{}{
				"a": float64(1),
				"b": float64(2),
			},
			"additionalProperty5": []interface{}{
				float64(1), float64(3),
			},
		},
	}); r != "" {
		t.Fatal(r)
	}
	if r := cmp.Diff(result.SalmonClassification.GetSalmon(), &expectedSalmon); r != "" {
		t.Fatal(r)
	}
	if r := cmp.Diff(result.SalmonClassification.GetFish(), &expectedFish); r != "" {
		t.Fatal(r)
	}
}

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type specified on the wire.
func TestPolymorphismGetComposedWithDiscriminator(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetComposedWithDiscriminator(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.DotFishMarket, DotFishMarket{
		Fishes: []DotFishClassification{
			&DotSalmon{
				FishType: to.StringPtr("DotSalmon"),
				Species:  to.StringPtr("king"),
				Location: to.StringPtr("australia"),
				Iswild:   to.BoolPtr(false),
			},
			&DotSalmon{
				FishType: to.StringPtr("DotSalmon"),
				Species:  to.StringPtr("king"),
				Location: to.StringPtr("canada"),
				Iswild:   to.BoolPtr(true),
			},
		},
		Salmons: []*DotSalmon{
			{
				FishType: to.StringPtr("DotSalmon"),
				Species:  to.StringPtr("king"),
				Location: to.StringPtr("sweden"),
				Iswild:   to.BoolPtr(false),
			},
			{
				FishType: to.StringPtr("DotSalmon"),
				Species:  to.StringPtr("king"),
				Location: to.StringPtr("atlantic"),
				Iswild:   to.BoolPtr(true),
			},
		},
		SampleFish: &DotSalmon{
			FishType: to.StringPtr("DotSalmon"),
			Species:  to.StringPtr("king"),
			Location: to.StringPtr("australia"),
			Iswild:   to.BoolPtr(false),
		},
		SampleSalmon: &DotSalmon{
			FishType: to.StringPtr("DotSalmon"),
			Species:  to.StringPtr("king"),
			Location: to.StringPtr("sweden"),
			Iswild:   to.BoolPtr(false),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, without discriminator specified on wire. Deserialization must NOT fail and use the explicit type of the property.
func TestPolymorphismGetComposedWithoutDiscriminator(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetComposedWithoutDiscriminator(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.DotFishMarket, DotFishMarket{
		Fishes: []DotFishClassification{
			&DotFish{
				Species: to.StringPtr("king"),
			},
			&DotFish{
				Species: to.StringPtr("king"),
			},
		},
		Salmons: []*DotSalmon{
			{
				Species:  to.StringPtr("king"),
				Location: to.StringPtr("sweden"),
				Iswild:   to.BoolPtr(false),
			},
			{
				Species:  to.StringPtr("king"),
				Location: to.StringPtr("atlantic"),
				Iswild:   to.BoolPtr(true),
			},
		},
		SampleFish: &DotFish{
			Species: to.StringPtr("king"),
		},
		SampleSalmon: &DotSalmon{
			Species:  to.StringPtr("king"),
			Location: to.StringPtr("sweden"),
			Iswild:   to.BoolPtr(false),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
func TestPolymorphismGetDotSyntax(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetDotSyntax(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.DotFishClassification, &DotSalmon{
		FishType: to.StringPtr("DotSalmon"),
		Species:  to.StringPtr("king"),
		Location: to.StringPtr("sweden"),
		Iswild:   to.BoolPtr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetValid - Get complex types that are polymorphic
func TestPolymorphismGetValid(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	salmon, ok := result.FishClassification.(*Salmon)
	if !ok {
		t.Fatal("fish wasn't a salmon")
	}
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	if r := cmp.Diff(salmon, &Salmon{
		Fishtype: to.StringPtr("salmon"),
		Length:   to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.StringPtr("goblin"),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}); r != "" {
		t.Fatal(r)
	}
}

// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func TestPolymorphismPutComplicated(t *testing.T) {
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	client := newPolymorphismClient()
	result, err := client.PutComplicated(context.Background(), &SmartSalmon{
		Fishtype: to.StringPtr("smart_salmon"),
		Length:   to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.StringPtr("goblin"),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
		AdditionalProperties: map[string]interface{}{
			"additionalProperty1": float64(1),
			"additionalProperty2": false,
			"additionalProperty3": "hello",
			"additionalProperty4": map[string]interface{}{
				"a": float64(1),
				"b": float64(2),
			},
			"additionalProperty5": []interface{}{
				float64(1), float64(3),
			},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
func TestPolymorphismPutMissingDiscriminator(t *testing.T) {
	client := newPolymorphismClient()
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutMissingDiscriminator(context.Background(), &Salmon{
		Length: to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	expectedSalmon := &Salmon{
		Length: to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.StringPtr("goblin"),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}
	if r := cmp.Diff(expectedSalmon, result.SalmonClassification); r != "" {
		t.Fatal(r)
	}
}

// PutValid - Put complex types that are polymorphic
func TestPolymorphismPutValid(t *testing.T) {
	client := newPolymorphismClient()
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutValid(context.Background(), &Salmon{
		Length: to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Length:   to.Float32Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request should not be allowed from the client
func TestPolymorphismPutValidMissingRequired(t *testing.T) {
	t.Skip("client side validation not applicable to track 2")
}
