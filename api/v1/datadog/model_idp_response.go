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

// IdpResponse The IdP response object.
type IdpResponse struct {
	// Identity provider response.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewIdpResponse instantiates a new IdpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpResponse(message string) *IdpResponse {
	this := IdpResponse{}
	this.Message = message
	return &this
}

// NewIdpResponseWithDefaults instantiates a new IdpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpResponseWithDefaults() *IdpResponse {
	this := IdpResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *IdpResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *IdpResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *IdpResponse) SetMessage(v string) {
	o.Message = v
}

func (o IdpResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

func (o *IdpResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Message *string `json:"message"`
	}{}
	all := struct {
		Message string `json:"message"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Message == nil {
		return fmt.Errorf("Required field message missing")
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

	o.Message = all.Message
	return nil
}

type NullableIdpResponse struct {
	value *IdpResponse
	isSet bool
}

func (v NullableIdpResponse) Get() *IdpResponse {
	return v.value
}

func (v *NullableIdpResponse) Set(val *IdpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpResponse(val *IdpResponse) *NullableIdpResponse {
	return &NullableIdpResponse{value: val, isSet: true}
}

func (v NullableIdpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
