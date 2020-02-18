// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/urlgroup/internal/urlgroup"
	"time"
)

// QueriesOperations contains the methods for the Queries group.
type QueriesOperations interface {
	// ArrayStringCsvEmpty - Get an empty array [] of string using the csv-array format 
	ArrayStringCsvEmpty(ctx context.Context, arrayQuery []string) (*QueriesArrayStringCsvEmptyResponse, error)
	// ArrayStringCsvNull - Get a null array of string using the csv-array format 
	ArrayStringCsvNull(ctx context.Context, arrayQuery []string) (*QueriesArrayStringCsvNullResponse, error)
	// ArrayStringCsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format 
	ArrayStringCsvValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringCsvValidResponse, error)
	// ArrayStringPipesValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the pipes-array format 
	ArrayStringPipesValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringPipesValidResponse, error)
	// ArrayStringSsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the ssv-array format 
	ArrayStringSsvValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringSsvValidResponse, error)
	// ArrayStringTsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the tsv-array format 
	ArrayStringTsvValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringTsvValidResponse, error)
	// ByteEmpty - Get '' as byte array 
	ByteEmpty(ctx context.Context) (*QueriesByteEmptyResponse, error)
	// ByteMultiByte - Get '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array 
	ByteMultiByte(ctx context.Context, byteQuery []byte) (*QueriesByteMultiByteResponse, error)
	// ByteNull - Get null as byte array (no query parameters in uri) 
	ByteNull(ctx context.Context, byteQuery []byte) (*QueriesByteNullResponse, error)
	// DateNull - Get null as date - this should result in no query parameters in uri 
	DateNull(ctx context.Context, dateQuery time.Time) (*QueriesDateNullResponse, error)
	// DateTimeNull - Get null as date-time, should result in no query parameters in uri 
	DateTimeNull(ctx context.Context, dateTimeQuery time.Time) (*QueriesDateTimeNullResponse, error)
	// DateTimeValid - Get '2012-01-01T01:01:01Z' as date-time 
	DateTimeValid(ctx context.Context) (*QueriesDateTimeValidResponse, error)
	// DateValid - Get '2012-01-01' as date 
	DateValid(ctx context.Context) (*QueriesDateValidResponse, error)
	// DoubleDecimalNegative - Get '-9999999.999' numeric value 
	DoubleDecimalNegative(ctx context.Context) (*QueriesDoubleDecimalNegativeResponse, error)
	// DoubleDecimalPositive - Get '9999999.999' numeric value 
	DoubleDecimalPositive(ctx context.Context) (*QueriesDoubleDecimalPositiveResponse, error)
	// DoubleNull - Get null numeric value (no query parameter) 
	DoubleNull(ctx context.Context, doubleQuery float64) (*QueriesDoubleNullResponse, error)
	// EnumNull - Get null (no query parameter in url) 
	EnumNull(ctx context.Context, enumQuery UriColor) (*QueriesEnumNullResponse, error)
	// EnumValid - Get using uri with query parameter 'green color' 
	EnumValid(ctx context.Context, enumQuery UriColor) (*QueriesEnumValidResponse, error)
	// FloatNull - Get null numeric value (no query parameter) 
	FloatNull(ctx context.Context, floatQuery float32) (*QueriesFloatNullResponse, error)
	// FloatScientificNegative - Get '-1.034E-20' numeric value 
	FloatScientificNegative(ctx context.Context) (*QueriesFloatScientificNegativeResponse, error)
	// FloatScientificPositive - Get '1.034E+20' numeric value 
	FloatScientificPositive(ctx context.Context) (*QueriesFloatScientificPositiveResponse, error)
	// GetBooleanFalse - Get false Boolean value on path 
	GetBooleanFalse(ctx context.Context) (*QueriesGetBooleanFalseResponse, error)
	// GetBooleanNull - Get null Boolean value on query (query string should be absent) 
	GetBooleanNull(ctx context.Context, boolQuery bool) (*QueriesGetBooleanNullResponse, error)
	// GetBooleanTrue - Get true Boolean value on path 
	GetBooleanTrue(ctx context.Context) (*QueriesGetBooleanTrueResponse, error)
	// GetIntNegativeOneMillion - Get '-1000000' integer value 
	GetIntNegativeOneMillion(ctx context.Context) (*QueriesGetIntNegativeOneMillionResponse, error)
	// GetIntNull - Get null integer value (no query parameter) 
	GetIntNull(ctx context.Context, intQuery int32) (*QueriesGetIntNullResponse, error)
	// GetIntOneMillion - Get '1000000' integer value 
	GetIntOneMillion(ctx context.Context) (*QueriesGetIntOneMillionResponse, error)
	// GetLongNull - Get 'null 64 bit integer value (no query param in uri) 
	GetLongNull(ctx context.Context, longQuery int64) (*QueriesGetLongNullResponse, error)
	// GetNegativeTenBillion - Get '-10000000000' 64 bit integer value 
	GetNegativeTenBillion(ctx context.Context) (*QueriesGetNegativeTenBillionResponse, error)
	// GetTenBillion - Get '10000000000' 64 bit integer value 
	GetTenBillion(ctx context.Context) (*QueriesGetTenBillionResponse, error)
	// StringEmpty - Get '' 
	StringEmpty(ctx context.Context) (*QueriesStringEmptyResponse, error)
	// StringNull - Get null (no query parameter in url) 
	StringNull(ctx context.Context, stringQuery string) (*QueriesStringNullResponse, error)
	// StringURLEncoded - Get 'begin!*'();:@ &=+$,/?#[]end 
	StringURLEncoded(ctx context.Context) (*QueriesStringURLEncodedResponse, error)
	// StringUnicode - Get '啊齄丂狛狜隣郎隣兀﨩' multi-byte string value 
	StringUnicode(ctx context.Context) (*QueriesStringUnicodeResponse, error)
}

type queriesOperations struct {
	*Client
	azinternal.QueriesOperations
}

// ArrayStringCsvEmpty - Get an empty array [] of string using the csv-array format 
func (client *queriesOperations) ArrayStringCsvEmpty(ctx context.Context, arrayQuery []string) (*QueriesArrayStringCsvEmptyResponse, error) {
	req, err := client.ArrayStringCsvEmptyCreateRequest(*client.u, arrayQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ArrayStringCsvEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ArrayStringCsvNull - Get a null array of string using the csv-array format 
func (client *queriesOperations) ArrayStringCsvNull(ctx context.Context, arrayQuery []string) (*QueriesArrayStringCsvNullResponse, error) {
	req, err := client.ArrayStringCsvNullCreateRequest(*client.u, arrayQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ArrayStringCsvNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ArrayStringCsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format 
func (client *queriesOperations) ArrayStringCsvValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringCsvValidResponse, error) {
	req, err := client.ArrayStringCsvValidCreateRequest(*client.u, arrayQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ArrayStringCsvValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ArrayStringPipesValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the pipes-array format 
func (client *queriesOperations) ArrayStringPipesValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringPipesValidResponse, error) {
	req, err := client.ArrayStringPipesValidCreateRequest(*client.u, arrayQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ArrayStringPipesValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ArrayStringSsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the ssv-array format 
func (client *queriesOperations) ArrayStringSsvValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringSsvValidResponse, error) {
	req, err := client.ArrayStringSsvValidCreateRequest(*client.u, arrayQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ArrayStringSsvValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ArrayStringTsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the tsv-array format 
func (client *queriesOperations) ArrayStringTsvValid(ctx context.Context, arrayQuery []string) (*QueriesArrayStringTsvValidResponse, error) {
	req, err := client.ArrayStringTsvValidCreateRequest(*client.u, arrayQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ArrayStringTsvValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ByteEmpty - Get '' as byte array 
func (client *queriesOperations) ByteEmpty(ctx context.Context) (*QueriesByteEmptyResponse, error) {
	req, err := client.ByteEmptyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ByteEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ByteMultiByte - Get '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array 
func (client *queriesOperations) ByteMultiByte(ctx context.Context, byteQuery []byte) (*QueriesByteMultiByteResponse, error) {
	req, err := client.ByteMultiByteCreateRequest(*client.u, byteQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ByteMultiByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ByteNull - Get null as byte array (no query parameters in uri) 
func (client *queriesOperations) ByteNull(ctx context.Context, byteQuery []byte) (*QueriesByteNullResponse, error) {
	req, err := client.ByteNullCreateRequest(*client.u, byteQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ByteNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DateNull - Get null as date - this should result in no query parameters in uri 
func (client *queriesOperations) DateNull(ctx context.Context, dateQuery time.Time) (*QueriesDateNullResponse, error) {
	req, err := client.DateNullCreateRequest(*client.u, dateQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DateNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DateTimeNull - Get null as date-time, should result in no query parameters in uri 
func (client *queriesOperations) DateTimeNull(ctx context.Context, dateTimeQuery time.Time) (*QueriesDateTimeNullResponse, error) {
	req, err := client.DateTimeNullCreateRequest(*client.u, dateTimeQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DateTimeNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DateTimeValid - Get '2012-01-01T01:01:01Z' as date-time 
func (client *queriesOperations) DateTimeValid(ctx context.Context) (*QueriesDateTimeValidResponse, error) {
	req, err := client.DateTimeValidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DateTimeValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DateValid - Get '2012-01-01' as date 
func (client *queriesOperations) DateValid(ctx context.Context) (*QueriesDateValidResponse, error) {
	req, err := client.DateValidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DateValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DoubleDecimalNegative - Get '-9999999.999' numeric value 
func (client *queriesOperations) DoubleDecimalNegative(ctx context.Context) (*QueriesDoubleDecimalNegativeResponse, error) {
	req, err := client.DoubleDecimalNegativeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DoubleDecimalNegativeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DoubleDecimalPositive - Get '9999999.999' numeric value 
func (client *queriesOperations) DoubleDecimalPositive(ctx context.Context) (*QueriesDoubleDecimalPositiveResponse, error) {
	req, err := client.DoubleDecimalPositiveCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DoubleDecimalPositiveHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DoubleNull - Get null numeric value (no query parameter) 
func (client *queriesOperations) DoubleNull(ctx context.Context, doubleQuery float64) (*QueriesDoubleNullResponse, error) {
	req, err := client.DoubleNullCreateRequest(*client.u, doubleQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DoubleNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EnumNull - Get null (no query parameter in url) 
func (client *queriesOperations) EnumNull(ctx context.Context, enumQuery UriColor) (*QueriesEnumNullResponse, error) {
	req, err := client.EnumNullCreateRequest(*client.u, enumQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.EnumNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EnumValid - Get using uri with query parameter 'green color' 
func (client *queriesOperations) EnumValid(ctx context.Context, enumQuery UriColor) (*QueriesEnumValidResponse, error) {
	req, err := client.EnumValidCreateRequest(*client.u, enumQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.EnumValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FloatNull - Get null numeric value (no query parameter) 
func (client *queriesOperations) FloatNull(ctx context.Context, floatQuery float32) (*QueriesFloatNullResponse, error) {
	req, err := client.FloatNullCreateRequest(*client.u, floatQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.FloatNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FloatScientificNegative - Get '-1.034E-20' numeric value 
func (client *queriesOperations) FloatScientificNegative(ctx context.Context) (*QueriesFloatScientificNegativeResponse, error) {
	req, err := client.FloatScientificNegativeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.FloatScientificNegativeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FloatScientificPositive - Get '1.034E+20' numeric value 
func (client *queriesOperations) FloatScientificPositive(ctx context.Context) (*QueriesFloatScientificPositiveResponse, error) {
	req, err := client.FloatScientificPositiveCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.FloatScientificPositiveHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBooleanFalse - Get false Boolean value on path 
func (client *queriesOperations) GetBooleanFalse(ctx context.Context) (*QueriesGetBooleanFalseResponse, error) {
	req, err := client.GetBooleanFalseCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetBooleanFalseHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBooleanNull - Get null Boolean value on query (query string should be absent) 
func (client *queriesOperations) GetBooleanNull(ctx context.Context, boolQuery bool) (*QueriesGetBooleanNullResponse, error) {
	req, err := client.GetBooleanNullCreateRequest(*client.u, boolQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetBooleanNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBooleanTrue - Get true Boolean value on path 
func (client *queriesOperations) GetBooleanTrue(ctx context.Context) (*QueriesGetBooleanTrueResponse, error) {
	req, err := client.GetBooleanTrueCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetBooleanTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIntNegativeOneMillion - Get '-1000000' integer value 
func (client *queriesOperations) GetIntNegativeOneMillion(ctx context.Context) (*QueriesGetIntNegativeOneMillionResponse, error) {
	req, err := client.GetIntNegativeOneMillionCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetIntNegativeOneMillionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIntNull - Get null integer value (no query parameter) 
func (client *queriesOperations) GetIntNull(ctx context.Context, intQuery int32) (*QueriesGetIntNullResponse, error) {
	req, err := client.GetIntNullCreateRequest(*client.u, intQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetIntNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIntOneMillion - Get '1000000' integer value 
func (client *queriesOperations) GetIntOneMillion(ctx context.Context) (*QueriesGetIntOneMillionResponse, error) {
	req, err := client.GetIntOneMillionCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetIntOneMillionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLongNull - Get 'null 64 bit integer value (no query param in uri) 
func (client *queriesOperations) GetLongNull(ctx context.Context, longQuery int64) (*QueriesGetLongNullResponse, error) {
	req, err := client.GetLongNullCreateRequest(*client.u, longQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetLongNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNegativeTenBillion - Get '-10000000000' 64 bit integer value 
func (client *queriesOperations) GetNegativeTenBillion(ctx context.Context) (*QueriesGetNegativeTenBillionResponse, error) {
	req, err := client.GetNegativeTenBillionCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetNegativeTenBillionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTenBillion - Get '10000000000' 64 bit integer value 
func (client *queriesOperations) GetTenBillion(ctx context.Context) (*QueriesGetTenBillionResponse, error) {
	req, err := client.GetTenBillionCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetTenBillionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StringEmpty - Get '' 
func (client *queriesOperations) StringEmpty(ctx context.Context) (*QueriesStringEmptyResponse, error) {
	req, err := client.StringEmptyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.StringEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StringNull - Get null (no query parameter in url) 
func (client *queriesOperations) StringNull(ctx context.Context, stringQuery string) (*QueriesStringNullResponse, error) {
	req, err := client.StringNullCreateRequest(*client.u, stringQuery)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.StringNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StringURLEncoded - Get 'begin!*'();:@ &=+$,/?#[]end 
func (client *queriesOperations) StringURLEncoded(ctx context.Context) (*QueriesStringURLEncodedResponse, error) {
	req, err := client.StringURLEncodedCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.StringURLEncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StringUnicode - Get '啊齄丂狛狜隣郎隣兀﨩' multi-byte string value 
func (client *queriesOperations) StringUnicode(ctx context.Context) (*QueriesStringUnicodeResponse, error) {
	req, err := client.StringUnicodeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.StringUnicodeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ QueriesOperations = (*queriesOperations)(nil)