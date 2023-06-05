//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dictionarygroup

import "time"

// DictionaryClientGetArrayEmptyResponse contains the response from method DictionaryClient.GetArrayEmpty.
type DictionaryClientGetArrayEmptyResponse struct {
	// An empty dictionary {}
	Value map[string][]*string
}

// DictionaryClientGetArrayItemEmptyResponse contains the response from method DictionaryClient.GetArrayItemEmpty.
type DictionaryClientGetArrayItemEmptyResponse struct {
	// An array of array of strings {"0": ["1", "2", "3"], "1": [], "2": ["7", "8", "9"]}
	Value map[string][]*string
}

// DictionaryClientGetArrayItemNullResponse contains the response from method DictionaryClient.GetArrayItemNull.
type DictionaryClientGetArrayItemNullResponse struct {
	// An array of array of strings {"0": ["1", "2", "3"], "1": null, "2": ["7", "8", "9"]}
	Value map[string][]*string
}

// DictionaryClientGetArrayNullResponse contains the response from method DictionaryClient.GetArrayNull.
type DictionaryClientGetArrayNullResponse struct {
	// a null array
	Value map[string][]*string
}

// DictionaryClientGetArrayValidResponse contains the response from method DictionaryClient.GetArrayValid.
type DictionaryClientGetArrayValidResponse struct {
	// An array of array of strings {"0": ["1", "2", "3"], "1": ["4", "5", "6"], "2": ["7", "8", "9"]}
	Value map[string][]*string
}

// DictionaryClientGetBase64URLResponse contains the response from method DictionaryClient.GetBase64URL.
type DictionaryClientGetBase64URLResponse struct {
	// The base64url dictionary value {"0": "a string that gets encoded with base64url", "1": "test string", "2": "Lorem ipsum"}
	Value map[string][]byte
}

// DictionaryClientGetBooleanInvalidNullResponse contains the response from method DictionaryClient.GetBooleanInvalidNull.
type DictionaryClientGetBooleanInvalidNullResponse struct {
	// The dictionary value {"0": true, "1": null, "2": false }
	Value map[string]*bool
}

// DictionaryClientGetBooleanInvalidStringResponse contains the response from method DictionaryClient.GetBooleanInvalidString.
type DictionaryClientGetBooleanInvalidStringResponse struct {
	// The dictionary value [true, 'boolean', false]
	Value map[string]*bool
}

// DictionaryClientGetBooleanTfftResponse contains the response from method DictionaryClient.GetBooleanTfft.
type DictionaryClientGetBooleanTfftResponse struct {
	// The dictionary value {"0": true, "1": false, "2": false, "3": true }
	Value map[string]*bool
}

// DictionaryClientGetByteInvalidNullResponse contains the response from method DictionaryClient.GetByteInvalidNull.
type DictionaryClientGetByteInvalidNullResponse struct {
	// The byte dictionary value {"0": hex(FF FF FF FA), "1": null} with the first item base64 encoded
	Value map[string][]byte
}

// DictionaryClientGetByteValidResponse contains the response from method DictionaryClient.GetByteValid.
type DictionaryClientGetByteValidResponse struct {
	// The dictionary value {"0": hex(FF FF FF FA), "1": hex(01 02 03), "2": hex (25, 29, 43)} with each elementencoded in base
	// 64
	Value map[string][]byte
}

// DictionaryClientGetComplexEmptyResponse contains the response from method DictionaryClient.GetComplexEmpty.
type DictionaryClientGetComplexEmptyResponse struct {
	// Empty dictionary of complex type {}
	Value map[string]*Widget
}

// DictionaryClientGetComplexItemEmptyResponse contains the response from method DictionaryClient.GetComplexItemEmpty.
type DictionaryClientGetComplexItemEmptyResponse struct {
	// Dictionary of complex type with empty item [{'integer': 1 'string': '2'}, {}, {'integer': 5, 'string': '6'}]
	Value map[string]*Widget
}

// DictionaryClientGetComplexItemNullResponse contains the response from method DictionaryClient.GetComplexItemNull.
type DictionaryClientGetComplexItemNullResponse struct {
	// Dictionary of complex type with null item [{'integer': 1 'string': '2'}, null, {'integer': 5, 'string': '6'}]
	Value map[string]*Widget
}

// DictionaryClientGetComplexNullResponse contains the response from method DictionaryClient.GetComplexNull.
type DictionaryClientGetComplexNullResponse struct {
	// Dictionary of complex type with null value
	Value map[string]*Widget
}

// DictionaryClientGetComplexValidResponse contains the response from method DictionaryClient.GetComplexValid.
type DictionaryClientGetComplexValidResponse struct {
	// Dictionary of complex type with [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string':
	// '6'}]
	Value map[string]*Widget
}

// DictionaryClientGetDateInvalidCharsResponse contains the response from method DictionaryClient.GetDateInvalidChars.
type DictionaryClientGetDateInvalidCharsResponse struct {
	// The dictionary value {"0": "2011-03-22", "1": "date"}
	Value map[string]*time.Time
}

// DictionaryClientGetDateInvalidNullResponse contains the response from method DictionaryClient.GetDateInvalidNull.
type DictionaryClientGetDateInvalidNullResponse struct {
	// The dictionary value {"0": "2012-01-01", "1": null, "2": "1776-07-04"}
	Value map[string]*time.Time
}

// DictionaryClientGetDateTimeInvalidCharsResponse contains the response from method DictionaryClient.GetDateTimeInvalidChars.
type DictionaryClientGetDateTimeInvalidCharsResponse struct {
	// The dictionary value {"0": "2000-12-01t00:00:01z", "1": "date-time"}
	Value map[string]*time.Time
}

// DictionaryClientGetDateTimeInvalidNullResponse contains the response from method DictionaryClient.GetDateTimeInvalidNull.
type DictionaryClientGetDateTimeInvalidNullResponse struct {
	// The dictionary value {"0": "2000-12-01t00:00:01z", "1": null}
	Value map[string]*time.Time
}

// DictionaryClientGetDateTimeRFC1123ValidResponse contains the response from method DictionaryClient.GetDateTimeRFC1123Valid.
type DictionaryClientGetDateTimeRFC1123ValidResponse struct {
	// The dictionary value {"0": "Fri, 01 Dec 2000 00:00:01 GMT", "1": "Wed, 02 Jan 1980 00:11:35 GMT", "2": "Wed, 12 Oct 1492
	// 10:15:01 GMT"}
	Value map[string]*time.Time
}

// DictionaryClientGetDateTimeValidResponse contains the response from method DictionaryClient.GetDateTimeValid.
type DictionaryClientGetDateTimeValidResponse struct {
	// The dictionary value {"0": "2000-12-01t00:00:01z", "1": "1980-01-02T00:11:35+01:00", "2": "1492-10-12T10:15:01-08:00"}
	Value map[string]*time.Time
}

// DictionaryClientGetDateValidResponse contains the response from method DictionaryClient.GetDateValid.
type DictionaryClientGetDateValidResponse struct {
	// The dictionary value {"0": "2000-12-01", "1": "1980-01-02", "2": "1492-10-12"}
	Value map[string]*time.Time
}

// DictionaryClientGetDictionaryEmptyResponse contains the response from method DictionaryClient.GetDictionaryEmpty.
type DictionaryClientGetDictionaryEmptyResponse struct {
	// An dictionaries of dictionaries of type <string, string> with value {}
	Value map[string]map[string]*string
}

// DictionaryClientGetDictionaryItemEmptyResponse contains the response from method DictionaryClient.GetDictionaryItemEmpty.
type DictionaryClientGetDictionaryItemEmptyResponse struct {
	// An dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1":
	// {}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
	Value map[string]map[string]*string
}

// DictionaryClientGetDictionaryItemNullResponse contains the response from method DictionaryClient.GetDictionaryItemNull.
type DictionaryClientGetDictionaryItemNullResponse struct {
	// An dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1":
	// null, "2": {"7": "seven", "8": "eight", "9": "nine"}}
	Value map[string]map[string]*string
}

// DictionaryClientGetDictionaryNullResponse contains the response from method DictionaryClient.GetDictionaryNull.
type DictionaryClientGetDictionaryNullResponse struct {
	// An dictionaries of dictionaries with value null
	Value map[string]map[string]*string
}

// DictionaryClientGetDictionaryValidResponse contains the response from method DictionaryClient.GetDictionaryValid.
type DictionaryClientGetDictionaryValidResponse struct {
	// An dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1":
	// {"4": "four", "5": "five", "6": "six"}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
	Value map[string]map[string]*string
}

// DictionaryClientGetDoubleInvalidNullResponse contains the response from method DictionaryClient.GetDoubleInvalidNull.
type DictionaryClientGetDoubleInvalidNullResponse struct {
	// The dictionary value {"0": 0.0, "1": null, "2": 1.2e20}
	Value map[string]*float64
}

// DictionaryClientGetDoubleInvalidStringResponse contains the response from method DictionaryClient.GetDoubleInvalidString.
type DictionaryClientGetDoubleInvalidStringResponse struct {
	// The dictionary value {"0": 1.0, "1": "number", "2": 0.0}
	Value map[string]*float64
}

// DictionaryClientGetDoubleValidResponse contains the response from method DictionaryClient.GetDoubleValid.
type DictionaryClientGetDoubleValidResponse struct {
	// The dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
	Value map[string]*float64
}

// DictionaryClientGetDurationValidResponse contains the response from method DictionaryClient.GetDurationValid.
type DictionaryClientGetDurationValidResponse struct {
	// The dictionary value {"0": "P123DT22H14M12.011S", "1": "P5DT1H0M0S"}
	Value map[string]*string
}

// DictionaryClientGetEmptyResponse contains the response from method DictionaryClient.GetEmpty.
type DictionaryClientGetEmptyResponse struct {
	// The empty dictionary value {}
	Value map[string]*int32
}

// DictionaryClientGetEmptyStringKeyResponse contains the response from method DictionaryClient.GetEmptyStringKey.
type DictionaryClientGetEmptyStringKeyResponse struct {
	// Dictionary of <string>
	Value map[string]*string
}

// DictionaryClientGetFloatInvalidNullResponse contains the response from method DictionaryClient.GetFloatInvalidNull.
type DictionaryClientGetFloatInvalidNullResponse struct {
	// The dictionary value {"0": 0.0, "1": null, "2": 1.2e20}
	Value map[string]*float32
}

// DictionaryClientGetFloatInvalidStringResponse contains the response from method DictionaryClient.GetFloatInvalidString.
type DictionaryClientGetFloatInvalidStringResponse struct {
	// The dictionary value {"0": 1.0, "1": "number", "2": 0.0}
	Value map[string]*float32
}

// DictionaryClientGetFloatValidResponse contains the response from method DictionaryClient.GetFloatValid.
type DictionaryClientGetFloatValidResponse struct {
	// The dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
	Value map[string]*float32
}

// DictionaryClientGetIntInvalidNullResponse contains the response from method DictionaryClient.GetIntInvalidNull.
type DictionaryClientGetIntInvalidNullResponse struct {
	// The dictionary value {"0": 1, "1": null, "2": 0}
	Value map[string]*int32
}

// DictionaryClientGetIntInvalidStringResponse contains the response from method DictionaryClient.GetIntInvalidString.
type DictionaryClientGetIntInvalidStringResponse struct {
	// The dictionary value {"0": 1, "1": "integer", "2": 0}
	Value map[string]*int32
}

// DictionaryClientGetIntegerValidResponse contains the response from method DictionaryClient.GetIntegerValid.
type DictionaryClientGetIntegerValidResponse struct {
	// The dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
	Value map[string]*int32
}

// DictionaryClientGetInvalidResponse contains the response from method DictionaryClient.GetInvalid.
type DictionaryClientGetInvalidResponse struct {
	// Dictionary of <string>
	Value map[string]*string
}

// DictionaryClientGetLongInvalidNullResponse contains the response from method DictionaryClient.GetLongInvalidNull.
type DictionaryClientGetLongInvalidNullResponse struct {
	// The dictionary value {"0": 1, "1": null, "2": 0}
	Value map[string]*int64
}

// DictionaryClientGetLongInvalidStringResponse contains the response from method DictionaryClient.GetLongInvalidString.
type DictionaryClientGetLongInvalidStringResponse struct {
	// The dictionary value {"0": 1, "1": "integer", "2": 0}
	Value map[string]*int64
}

// DictionaryClientGetLongValidResponse contains the response from method DictionaryClient.GetLongValid.
type DictionaryClientGetLongValidResponse struct {
	// The dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
	Value map[string]*int64
}

// DictionaryClientGetNullKeyResponse contains the response from method DictionaryClient.GetNullKey.
type DictionaryClientGetNullKeyResponse struct {
	// Dictionary of <string>
	Value map[string]*string
}

// DictionaryClientGetNullResponse contains the response from method DictionaryClient.GetNull.
type DictionaryClientGetNullResponse struct {
	// The null dictionary value
	Value map[string]*int32
}

// DictionaryClientGetNullValueResponse contains the response from method DictionaryClient.GetNullValue.
type DictionaryClientGetNullValueResponse struct {
	// Dictionary of <string>
	Value map[string]*string
}

// DictionaryClientGetStringValidResponse contains the response from method DictionaryClient.GetStringValid.
type DictionaryClientGetStringValidResponse struct {
	// The dictionary value {"0": "foo1", "1": "foo2", "2": "foo3"}
	Value map[string]*string
}

// DictionaryClientGetStringWithInvalidResponse contains the response from method DictionaryClient.GetStringWithInvalid.
type DictionaryClientGetStringWithInvalidResponse struct {
	// The dictionary value {"0": "foo", "1": 123, "2": "foo2"}
	Value map[string]*string
}

// DictionaryClientGetStringWithNullResponse contains the response from method DictionaryClient.GetStringWithNull.
type DictionaryClientGetStringWithNullResponse struct {
	// The dictionary value {"0": "foo", "1": null, "2": "foo2"}
	Value map[string]*string
}

// DictionaryClientPutArrayValidResponse contains the response from method DictionaryClient.PutArrayValid.
type DictionaryClientPutArrayValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutBooleanTfftResponse contains the response from method DictionaryClient.PutBooleanTfft.
type DictionaryClientPutBooleanTfftResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutByteValidResponse contains the response from method DictionaryClient.PutByteValid.
type DictionaryClientPutByteValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutComplexValidResponse contains the response from method DictionaryClient.PutComplexValid.
type DictionaryClientPutComplexValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutDateTimeRFC1123ValidResponse contains the response from method DictionaryClient.PutDateTimeRFC1123Valid.
type DictionaryClientPutDateTimeRFC1123ValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutDateTimeValidResponse contains the response from method DictionaryClient.PutDateTimeValid.
type DictionaryClientPutDateTimeValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutDateValidResponse contains the response from method DictionaryClient.PutDateValid.
type DictionaryClientPutDateValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutDictionaryValidResponse contains the response from method DictionaryClient.PutDictionaryValid.
type DictionaryClientPutDictionaryValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutDoubleValidResponse contains the response from method DictionaryClient.PutDoubleValid.
type DictionaryClientPutDoubleValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutDurationValidResponse contains the response from method DictionaryClient.PutDurationValid.
type DictionaryClientPutDurationValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutEmptyResponse contains the response from method DictionaryClient.PutEmpty.
type DictionaryClientPutEmptyResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutFloatValidResponse contains the response from method DictionaryClient.PutFloatValid.
type DictionaryClientPutFloatValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutIntegerValidResponse contains the response from method DictionaryClient.PutIntegerValid.
type DictionaryClientPutIntegerValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutLongValidResponse contains the response from method DictionaryClient.PutLongValid.
type DictionaryClientPutLongValidResponse struct {
	// placeholder for future response values
}

// DictionaryClientPutStringValidResponse contains the response from method DictionaryClient.PutStringValid.
type DictionaryClientPutStringValidResponse struct {
	// placeholder for future response values
}
