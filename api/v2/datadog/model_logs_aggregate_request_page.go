/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// LogsAggregateRequestPage Paging settings
type LogsAggregateRequestPage struct {
	// The returned paging point to use to get the next results
	Cursor *string `json:"cursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewLogsAggregateRequestPage instantiates a new LogsAggregateRequestPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsAggregateRequestPage() *LogsAggregateRequestPage {
	this := LogsAggregateRequestPage{}
	return &this
}

// NewLogsAggregateRequestPageWithDefaults instantiates a new LogsAggregateRequestPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsAggregateRequestPageWithDefaults() *LogsAggregateRequestPage {
	this := LogsAggregateRequestPage{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *LogsAggregateRequestPage) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateRequestPage) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *LogsAggregateRequestPage) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *LogsAggregateRequestPage) SetCursor(v string) {
	o.Cursor = &v
}

func (o LogsAggregateRequestPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

func (o *LogsAggregateRequestPage) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Cursor *string `json:"cursor,omitempty"`
	}{}
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

	o.Cursor = all.Cursor
	return nil
}

type NullableLogsAggregateRequestPage struct {
	value *LogsAggregateRequestPage
	isSet bool
}

func (v NullableLogsAggregateRequestPage) Get() *LogsAggregateRequestPage {
	return v.value
}

func (v *NullableLogsAggregateRequestPage) Set(val *LogsAggregateRequestPage) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsAggregateRequestPage) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsAggregateRequestPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsAggregateRequestPage(val *LogsAggregateRequestPage) *NullableLogsAggregateRequestPage {
	return &NullableLogsAggregateRequestPage{value: val, isSet: true}
}

func (v NullableLogsAggregateRequestPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsAggregateRequestPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
