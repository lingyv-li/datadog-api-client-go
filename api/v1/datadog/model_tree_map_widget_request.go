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

// TreeMapWidgetRequest An updated treemap widget.
type TreeMapWidgetRequest struct {
	// The widget metrics query.
	Q *string `json:"q,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewTreeMapWidgetRequest instantiates a new TreeMapWidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTreeMapWidgetRequest() *TreeMapWidgetRequest {
	this := TreeMapWidgetRequest{}
	return &this
}

// NewTreeMapWidgetRequestWithDefaults instantiates a new TreeMapWidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTreeMapWidgetRequestWithDefaults() *TreeMapWidgetRequest {
	this := TreeMapWidgetRequest{}
	return &this
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TreeMapWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *TreeMapWidgetRequest) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TreeMapWidgetRequest) SetQ(v string) {
	o.Q = &v
}

func (o TreeMapWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	return json.Marshal(toSerialize)
}

func (o *TreeMapWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Q *string `json:"q,omitempty"`
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

	o.Q = all.Q
	return nil
}

type NullableTreeMapWidgetRequest struct {
	value *TreeMapWidgetRequest
	isSet bool
}

func (v NullableTreeMapWidgetRequest) Get() *TreeMapWidgetRequest {
	return v.value
}

func (v *NullableTreeMapWidgetRequest) Set(val *TreeMapWidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTreeMapWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTreeMapWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTreeMapWidgetRequest(val *TreeMapWidgetRequest) *NullableTreeMapWidgetRequest {
	return &NullableTreeMapWidgetRequest{value: val, isSet: true}
}

func (v NullableTreeMapWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTreeMapWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
