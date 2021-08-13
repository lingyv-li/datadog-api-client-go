/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"reflect"
)

// UsageLogsByRetentionResponse Response containing the indexed logs usage broken down by retention period for an organization during a given hour.
type UsageLogsByRetentionResponse struct {
	// Get hourly usage for indexed logs by retention period.
	Usage *[]UsageLogsByRetentionHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewUsageLogsByRetentionResponse instantiates a new UsageLogsByRetentionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLogsByRetentionResponse() *UsageLogsByRetentionResponse {
	this := UsageLogsByRetentionResponse{}
	return &this
}

// NewUsageLogsByRetentionResponseWithDefaults instantiates a new UsageLogsByRetentionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLogsByRetentionResponseWithDefaults() *UsageLogsByRetentionResponse {
	this := UsageLogsByRetentionResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageLogsByRetentionResponse) GetUsage() []UsageLogsByRetentionHour {
	if o == nil || o.Usage == nil {
		var ret []UsageLogsByRetentionHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsByRetentionResponse) GetUsageOk() (*[]UsageLogsByRetentionHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageLogsByRetentionResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageLogsByRetentionHour and assigns it to the Usage field.
func (o *UsageLogsByRetentionResponse) SetUsage(v []UsageLogsByRetentionHour) {
	o.Usage = &v
}

func (o UsageLogsByRetentionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

func (o *UsageLogsByRetentionResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Usage *[]UsageLogsByRetentionHour `json:"usage,omitempty"`
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

	if !o.ContainsUnparsedObject {
		if v := all.Usage; v != nil {
			o.ContainsUnparsedObject = containsUnparsedObject(reflect.ValueOf(*v))
		}
	}

	o.Usage = all.Usage
	return nil
}

type NullableUsageLogsByRetentionResponse struct {
	value *UsageLogsByRetentionResponse
	isSet bool
}

func (v NullableUsageLogsByRetentionResponse) Get() *UsageLogsByRetentionResponse {
	return v.value
}

func (v *NullableUsageLogsByRetentionResponse) Set(val *UsageLogsByRetentionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLogsByRetentionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLogsByRetentionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLogsByRetentionResponse(val *UsageLogsByRetentionResponse) *NullableUsageLogsByRetentionResponse {
	return &NullableUsageLogsByRetentionResponse{value: val, isSet: true}
}

func (v NullableUsageLogsByRetentionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLogsByRetentionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
