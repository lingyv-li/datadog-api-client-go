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

// UsageIngestedSpansResponse Response containing the ingested spans usage for each hour for a given organization.
type UsageIngestedSpansResponse struct {
	// Get hourly usage for ingested spans.
	Usage *[]UsageIngestedSpansHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewUsageIngestedSpansResponse instantiates a new UsageIngestedSpansResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageIngestedSpansResponse() *UsageIngestedSpansResponse {
	this := UsageIngestedSpansResponse{}
	return &this
}

// NewUsageIngestedSpansResponseWithDefaults instantiates a new UsageIngestedSpansResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageIngestedSpansResponseWithDefaults() *UsageIngestedSpansResponse {
	this := UsageIngestedSpansResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageIngestedSpansResponse) GetUsage() []UsageIngestedSpansHour {
	if o == nil || o.Usage == nil {
		var ret []UsageIngestedSpansHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIngestedSpansResponse) GetUsageOk() (*[]UsageIngestedSpansHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageIngestedSpansResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageIngestedSpansHour and assigns it to the Usage field.
func (o *UsageIngestedSpansResponse) SetUsage(v []UsageIngestedSpansHour) {
	o.Usage = &v
}

func (o UsageIngestedSpansResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

func (o *UsageIngestedSpansResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Usage *[]UsageIngestedSpansHour `json:"usage,omitempty"`
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

type NullableUsageIngestedSpansResponse struct {
	value *UsageIngestedSpansResponse
	isSet bool
}

func (v NullableUsageIngestedSpansResponse) Get() *UsageIngestedSpansResponse {
	return v.value
}

func (v *NullableUsageIngestedSpansResponse) Set(val *UsageIngestedSpansResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageIngestedSpansResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageIngestedSpansResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageIngestedSpansResponse(val *UsageIngestedSpansResponse) *NullableUsageIngestedSpansResponse {
	return &NullableUsageIngestedSpansResponse{value: val, isSet: true}
}

func (v NullableUsageIngestedSpansResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageIngestedSpansResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
