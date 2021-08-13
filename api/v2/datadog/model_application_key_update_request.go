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
)

// ApplicationKeyUpdateRequest Request used to update an application key.
type ApplicationKeyUpdateRequest struct {
	Data ApplicationKeyUpdateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewApplicationKeyUpdateRequest instantiates a new ApplicationKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyUpdateRequest(data ApplicationKeyUpdateData) *ApplicationKeyUpdateRequest {
	this := ApplicationKeyUpdateRequest{}
	this.Data = data
	return &this
}

// NewApplicationKeyUpdateRequestWithDefaults instantiates a new ApplicationKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyUpdateRequestWithDefaults() *ApplicationKeyUpdateRequest {
	this := ApplicationKeyUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ApplicationKeyUpdateRequest) GetData() ApplicationKeyUpdateData {
	if o == nil {
		var ret ApplicationKeyUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyUpdateRequest) GetDataOk() (*ApplicationKeyUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ApplicationKeyUpdateRequest) SetData(v ApplicationKeyUpdateData) {
	o.Data = v
}

func (o ApplicationKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *ApplicationKeyUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *ApplicationKeyUpdateData `json:"data"`
	}{}
	all := struct {
		Data ApplicationKeyUpdateData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
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
	if v := all.Data; !o.ContainsUnparsedObject && v.ContainsUnparsedObject {
		o.ContainsUnparsedObject = true
	}

	o.Data = all.Data
	return nil
}

type NullableApplicationKeyUpdateRequest struct {
	value *ApplicationKeyUpdateRequest
	isSet bool
}

func (v NullableApplicationKeyUpdateRequest) Get() *ApplicationKeyUpdateRequest {
	return v.value
}

func (v *NullableApplicationKeyUpdateRequest) Set(val *ApplicationKeyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyUpdateRequest(val *ApplicationKeyUpdateRequest) *NullableApplicationKeyUpdateRequest {
	return &NullableApplicationKeyUpdateRequest{value: val, isSet: true}
}

func (v NullableApplicationKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
