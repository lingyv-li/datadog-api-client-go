/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ListStreamWidgetRequest Updated list stream widget.
type ListStreamWidgetRequest struct {
	// Widget columns.
	Columns        []ListStreamColumn       `json:"columns"`
	Query          ListStreamQuery          `json:"query"`
	ResponseFormat ListStreamResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewListStreamWidgetRequest instantiates a new ListStreamWidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStreamWidgetRequest(columns []ListStreamColumn, query ListStreamQuery, responseFormat ListStreamResponseFormat) *ListStreamWidgetRequest {
	this := ListStreamWidgetRequest{}
	this.Columns = columns
	this.Query = query
	this.ResponseFormat = responseFormat
	return &this
}

// NewListStreamWidgetRequestWithDefaults instantiates a new ListStreamWidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStreamWidgetRequestWithDefaults() *ListStreamWidgetRequest {
	this := ListStreamWidgetRequest{}
	return &this
}

// GetColumns returns the Columns field value
func (o *ListStreamWidgetRequest) GetColumns() []ListStreamColumn {
	if o == nil {
		var ret []ListStreamColumn
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *ListStreamWidgetRequest) GetColumnsOk() (*[]ListStreamColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *ListStreamWidgetRequest) SetColumns(v []ListStreamColumn) {
	o.Columns = v
}

// GetQuery returns the Query field value
func (o *ListStreamWidgetRequest) GetQuery() ListStreamQuery {
	if o == nil {
		var ret ListStreamQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ListStreamWidgetRequest) GetQueryOk() (*ListStreamQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ListStreamWidgetRequest) SetQuery(v ListStreamQuery) {
	o.Query = v
}

// GetResponseFormat returns the ResponseFormat field value
func (o *ListStreamWidgetRequest) GetResponseFormat() ListStreamResponseFormat {
	if o == nil {
		var ret ListStreamResponseFormat
		return ret
	}

	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ListStreamWidgetRequest) GetResponseFormatOk() (*ListStreamResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value
func (o *ListStreamWidgetRequest) SetResponseFormat(v ListStreamResponseFormat) {
	o.ResponseFormat = v
}

func (o ListStreamWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["columns"] = o.Columns
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["response_format"] = o.ResponseFormat
	}
	return json.Marshal(toSerialize)
}

func (o *ListStreamWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Columns        *[]ListStreamColumn       `json:"columns"`
		Query          *ListStreamQuery          `json:"query"`
		ResponseFormat *ListStreamResponseFormat `json:"response_format"`
	}{}
	all := struct {
		Columns        []ListStreamColumn       `json:"columns"`
		Query          ListStreamQuery          `json:"query"`
		ResponseFormat ListStreamResponseFormat `json:"response_format"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Columns == nil {
		return fmt.Errorf("Required field columns missing")
	}
	if required.Query == nil {
		return fmt.Errorf("Required field query missing")
	}
	if required.ResponseFormat == nil {
		return fmt.Errorf("Required field response_format missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.ContainsUnparsedObject = true
		o.UnparsedObject = raw
		return nil
	}

	if !o.ContainsUnparsedObject {
		if v := all.Columns; v != nil {
			o.ContainsUnparsedObject = containsUnparsedObject(reflect.ValueOf(v))
		}
	}

	if v := all.Query; !o.ContainsUnparsedObject && v.ContainsUnparsedObject {
		o.ContainsUnparsedObject = true
	}

	if v := all.ResponseFormat; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.ContainsUnparsedObject = true
		o.UnparsedObject = raw
		return nil
	}

	o.Columns = all.Columns
	o.Query = all.Query
	o.ResponseFormat = all.ResponseFormat
	return nil
}

type NullableListStreamWidgetRequest struct {
	value *ListStreamWidgetRequest
	isSet bool
}

func (v NullableListStreamWidgetRequest) Get() *ListStreamWidgetRequest {
	return v.value
}

func (v *NullableListStreamWidgetRequest) Set(val *ListStreamWidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListStreamWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListStreamWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStreamWidgetRequest(val *ListStreamWidgetRequest) *NullableListStreamWidgetRequest {
	return &NullableListStreamWidgetRequest{value: val, isSet: true}
}

func (v NullableListStreamWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStreamWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
