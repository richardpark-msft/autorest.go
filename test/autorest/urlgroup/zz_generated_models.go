//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import "time"

// PathItemsClientGetAllWithValuesOptions contains the optional parameters for the PathItemsClient.GetAllWithValues method.
type PathItemsClientGetAllWithValuesOptions struct {
	// should contain value 'localStringQuery'
	LocalStringQuery *string
	// A string value 'pathItemStringQuery' that appears as a query parameter
	PathItemStringQuery *string
}

// PathItemsClientGetGlobalAndLocalQueryNullOptions contains the optional parameters for the PathItemsClient.GetGlobalAndLocalQueryNull
// method.
type PathItemsClientGetGlobalAndLocalQueryNullOptions struct {
	// should contain null value
	LocalStringQuery *string
	// A string value 'pathItemStringQuery' that appears as a query parameter
	PathItemStringQuery *string
}

// PathItemsClientGetGlobalQueryNullOptions contains the optional parameters for the PathItemsClient.GetGlobalQueryNull method.
type PathItemsClientGetGlobalQueryNullOptions struct {
	// should contain value 'localStringQuery'
	LocalStringQuery *string
	// A string value 'pathItemStringQuery' that appears as a query parameter
	PathItemStringQuery *string
}

// PathItemsClientGetLocalPathItemQueryNullOptions contains the optional parameters for the PathItemsClient.GetLocalPathItemQueryNull
// method.
type PathItemsClientGetLocalPathItemQueryNullOptions struct {
	// should contain value null
	LocalStringQuery *string
	// should contain value null
	PathItemStringQuery *string
}

// PathsClientArrayCSVInPathOptions contains the optional parameters for the PathsClient.ArrayCSVInPath method.
type PathsClientArrayCSVInPathOptions struct {
	// placeholder for future optional parameters
}

// PathsClientBase64URLOptions contains the optional parameters for the PathsClient.Base64URL method.
type PathsClientBase64URLOptions struct {
	// placeholder for future optional parameters
}

// PathsClientByteEmptyOptions contains the optional parameters for the PathsClient.ByteEmpty method.
type PathsClientByteEmptyOptions struct {
	// placeholder for future optional parameters
}

// PathsClientByteMultiByteOptions contains the optional parameters for the PathsClient.ByteMultiByte method.
type PathsClientByteMultiByteOptions struct {
	// placeholder for future optional parameters
}

// PathsClientByteNullOptions contains the optional parameters for the PathsClient.ByteNull method.
type PathsClientByteNullOptions struct {
	// placeholder for future optional parameters
}

// PathsClientDateNullOptions contains the optional parameters for the PathsClient.DateNull method.
type PathsClientDateNullOptions struct {
	// placeholder for future optional parameters
}

// PathsClientDateTimeNullOptions contains the optional parameters for the PathsClient.DateTimeNull method.
type PathsClientDateTimeNullOptions struct {
	// placeholder for future optional parameters
}

// PathsClientDateTimeValidOptions contains the optional parameters for the PathsClient.DateTimeValid method.
type PathsClientDateTimeValidOptions struct {
	// placeholder for future optional parameters
}

// PathsClientDateValidOptions contains the optional parameters for the PathsClient.DateValid method.
type PathsClientDateValidOptions struct {
	// placeholder for future optional parameters
}

// PathsClientDoubleDecimalNegativeOptions contains the optional parameters for the PathsClient.DoubleDecimalNegative method.
type PathsClientDoubleDecimalNegativeOptions struct {
	// placeholder for future optional parameters
}

// PathsClientDoubleDecimalPositiveOptions contains the optional parameters for the PathsClient.DoubleDecimalPositive method.
type PathsClientDoubleDecimalPositiveOptions struct {
	// placeholder for future optional parameters
}

// PathsClientEnumNullOptions contains the optional parameters for the PathsClient.EnumNull method.
type PathsClientEnumNullOptions struct {
	// placeholder for future optional parameters
}

// PathsClientEnumValidOptions contains the optional parameters for the PathsClient.EnumValid method.
type PathsClientEnumValidOptions struct {
	// placeholder for future optional parameters
}

// PathsClientFloatScientificNegativeOptions contains the optional parameters for the PathsClient.FloatScientificNegative
// method.
type PathsClientFloatScientificNegativeOptions struct {
	// placeholder for future optional parameters
}

// PathsClientFloatScientificPositiveOptions contains the optional parameters for the PathsClient.FloatScientificPositive
// method.
type PathsClientFloatScientificPositiveOptions struct {
	// placeholder for future optional parameters
}

// PathsClientGetBooleanFalseOptions contains the optional parameters for the PathsClient.GetBooleanFalse method.
type PathsClientGetBooleanFalseOptions struct {
	// placeholder for future optional parameters
}

// PathsClientGetBooleanTrueOptions contains the optional parameters for the PathsClient.GetBooleanTrue method.
type PathsClientGetBooleanTrueOptions struct {
	// placeholder for future optional parameters
}

// PathsClientGetIntNegativeOneMillionOptions contains the optional parameters for the PathsClient.GetIntNegativeOneMillion
// method.
type PathsClientGetIntNegativeOneMillionOptions struct {
	// placeholder for future optional parameters
}

// PathsClientGetIntOneMillionOptions contains the optional parameters for the PathsClient.GetIntOneMillion method.
type PathsClientGetIntOneMillionOptions struct {
	// placeholder for future optional parameters
}

// PathsClientGetNegativeTenBillionOptions contains the optional parameters for the PathsClient.GetNegativeTenBillion method.
type PathsClientGetNegativeTenBillionOptions struct {
	// placeholder for future optional parameters
}

// PathsClientGetTenBillionOptions contains the optional parameters for the PathsClient.GetTenBillion method.
type PathsClientGetTenBillionOptions struct {
	// placeholder for future optional parameters
}

// PathsClientStringEmptyOptions contains the optional parameters for the PathsClient.StringEmpty method.
type PathsClientStringEmptyOptions struct {
	// placeholder for future optional parameters
}

// PathsClientStringNullOptions contains the optional parameters for the PathsClient.StringNull method.
type PathsClientStringNullOptions struct {
	// placeholder for future optional parameters
}

// PathsClientStringURLEncodedOptions contains the optional parameters for the PathsClient.StringURLEncoded method.
type PathsClientStringURLEncodedOptions struct {
	// placeholder for future optional parameters
}

// PathsClientStringURLNonEncodedOptions contains the optional parameters for the PathsClient.StringURLNonEncoded method.
type PathsClientStringURLNonEncodedOptions struct {
	// placeholder for future optional parameters
}

// PathsClientStringUnicodeOptions contains the optional parameters for the PathsClient.StringUnicode method.
type PathsClientStringUnicodeOptions struct {
	// placeholder for future optional parameters
}

// PathsClientUnixTimeURLOptions contains the optional parameters for the PathsClient.UnixTimeURL method.
type PathsClientUnixTimeURLOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientArrayStringCSVEmptyOptions contains the optional parameters for the QueriesClient.ArrayStringCSVEmpty method.
type QueriesClientArrayStringCSVEmptyOptions struct {
	// an empty array [] of string using the csv-array format
	ArrayQuery []string
}

// QueriesClientArrayStringCSVNullOptions contains the optional parameters for the QueriesClient.ArrayStringCSVNull method.
type QueriesClientArrayStringCSVNullOptions struct {
	// a null array of string using the csv-array format
	ArrayQuery []string
}

// QueriesClientArrayStringCSVValidOptions contains the optional parameters for the QueriesClient.ArrayStringCSVValid method.
type QueriesClientArrayStringCSVValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format
	ArrayQuery []string
}

// QueriesClientArrayStringNoCollectionFormatEmptyOptions contains the optional parameters for the QueriesClient.ArrayStringNoCollectionFormatEmpty
// method.
type QueriesClientArrayStringNoCollectionFormatEmptyOptions struct {
	// Array-typed query parameter. Pass in ['hello', 'nihao', 'bonjour'].
	ArrayQuery []string
}

// QueriesClientArrayStringPipesValidOptions contains the optional parameters for the QueriesClient.ArrayStringPipesValid
// method.
type QueriesClientArrayStringPipesValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the pipes-array format
	ArrayQuery []string
}

// QueriesClientArrayStringSsvValidOptions contains the optional parameters for the QueriesClient.ArrayStringSsvValid method.
type QueriesClientArrayStringSsvValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the ssv-array format
	ArrayQuery []string
}

// QueriesClientArrayStringTsvValidOptions contains the optional parameters for the QueriesClient.ArrayStringTsvValid method.
type QueriesClientArrayStringTsvValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the tsv-array format
	ArrayQuery []string
}

// QueriesClientByteEmptyOptions contains the optional parameters for the QueriesClient.ByteEmpty method.
type QueriesClientByteEmptyOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientByteMultiByteOptions contains the optional parameters for the QueriesClient.ByteMultiByte method.
type QueriesClientByteMultiByteOptions struct {
	// '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array
	ByteQuery []byte
}

// QueriesClientByteNullOptions contains the optional parameters for the QueriesClient.ByteNull method.
type QueriesClientByteNullOptions struct {
	// null as byte array (no query parameters in uri)
	ByteQuery []byte
}

// QueriesClientDateNullOptions contains the optional parameters for the QueriesClient.DateNull method.
type QueriesClientDateNullOptions struct {
	// null as date (no query parameters in uri)
	DateQuery *time.Time
}

// QueriesClientDateTimeNullOptions contains the optional parameters for the QueriesClient.DateTimeNull method.
type QueriesClientDateTimeNullOptions struct {
	// null as date-time (no query parameters)
	DateTimeQuery *time.Time
}

// QueriesClientDateTimeValidOptions contains the optional parameters for the QueriesClient.DateTimeValid method.
type QueriesClientDateTimeValidOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientDateValidOptions contains the optional parameters for the QueriesClient.DateValid method.
type QueriesClientDateValidOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientDoubleDecimalNegativeOptions contains the optional parameters for the QueriesClient.DoubleDecimalNegative
// method.
type QueriesClientDoubleDecimalNegativeOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientDoubleDecimalPositiveOptions contains the optional parameters for the QueriesClient.DoubleDecimalPositive
// method.
type QueriesClientDoubleDecimalPositiveOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientDoubleNullOptions contains the optional parameters for the QueriesClient.DoubleNull method.
type QueriesClientDoubleNullOptions struct {
	// null numeric value
	DoubleQuery *float64
}

// QueriesClientEnumNullOptions contains the optional parameters for the QueriesClient.EnumNull method.
type QueriesClientEnumNullOptions struct {
	// null string value
	EnumQuery *URIColor
}

// QueriesClientEnumValidOptions contains the optional parameters for the QueriesClient.EnumValid method.
type QueriesClientEnumValidOptions struct {
	// 'green color' enum value
	EnumQuery *URIColor
}

// QueriesClientFloatNullOptions contains the optional parameters for the QueriesClient.FloatNull method.
type QueriesClientFloatNullOptions struct {
	// null numeric value
	FloatQuery *float32
}

// QueriesClientFloatScientificNegativeOptions contains the optional parameters for the QueriesClient.FloatScientificNegative
// method.
type QueriesClientFloatScientificNegativeOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientFloatScientificPositiveOptions contains the optional parameters for the QueriesClient.FloatScientificPositive
// method.
type QueriesClientFloatScientificPositiveOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientGetBooleanFalseOptions contains the optional parameters for the QueriesClient.GetBooleanFalse method.
type QueriesClientGetBooleanFalseOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientGetBooleanNullOptions contains the optional parameters for the QueriesClient.GetBooleanNull method.
type QueriesClientGetBooleanNullOptions struct {
	// null boolean value
	BoolQuery *bool
}

// QueriesClientGetBooleanTrueOptions contains the optional parameters for the QueriesClient.GetBooleanTrue method.
type QueriesClientGetBooleanTrueOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientGetIntNegativeOneMillionOptions contains the optional parameters for the QueriesClient.GetIntNegativeOneMillion
// method.
type QueriesClientGetIntNegativeOneMillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientGetIntNullOptions contains the optional parameters for the QueriesClient.GetIntNull method.
type QueriesClientGetIntNullOptions struct {
	// null integer value
	IntQuery *int32
}

// QueriesClientGetIntOneMillionOptions contains the optional parameters for the QueriesClient.GetIntOneMillion method.
type QueriesClientGetIntOneMillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientGetLongNullOptions contains the optional parameters for the QueriesClient.GetLongNull method.
type QueriesClientGetLongNullOptions struct {
	// null 64 bit integer value
	LongQuery *int64
}

// QueriesClientGetNegativeTenBillionOptions contains the optional parameters for the QueriesClient.GetNegativeTenBillion
// method.
type QueriesClientGetNegativeTenBillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientGetTenBillionOptions contains the optional parameters for the QueriesClient.GetTenBillion method.
type QueriesClientGetTenBillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientStringEmptyOptions contains the optional parameters for the QueriesClient.StringEmpty method.
type QueriesClientStringEmptyOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientStringNullOptions contains the optional parameters for the QueriesClient.StringNull method.
type QueriesClientStringNullOptions struct {
	// null string value
	StringQuery *string
}

// QueriesClientStringURLEncodedOptions contains the optional parameters for the QueriesClient.StringURLEncoded method.
type QueriesClientStringURLEncodedOptions struct {
	// placeholder for future optional parameters
}

// QueriesClientStringUnicodeOptions contains the optional parameters for the QueriesClient.StringUnicode method.
type QueriesClientStringUnicodeOptions struct {
	// placeholder for future optional parameters
}
