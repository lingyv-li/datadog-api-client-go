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

// IncidentServiceUpdateAttributes The incident service's attributes for an update request.
type IncidentServiceUpdateAttributes struct {
	// Name of the incident service.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewIncidentServiceUpdateAttributes instantiates a new IncidentServiceUpdateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentServiceUpdateAttributes(name string) *IncidentServiceUpdateAttributes {
	this := IncidentServiceUpdateAttributes{}
	this.Name = name
	return &this
}

// NewIncidentServiceUpdateAttributesWithDefaults instantiates a new IncidentServiceUpdateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentServiceUpdateAttributesWithDefaults() *IncidentServiceUpdateAttributes {
	this := IncidentServiceUpdateAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *IncidentServiceUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IncidentServiceUpdateAttributes) SetName(v string) {
	o.Name = v
}

func (o IncidentServiceUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentServiceUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Name *string `json:"name"`
	}{}
	all := struct {
		Name string `json:"name"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Name == nil {
		return fmt.Errorf("Required field name missing")
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

	o.Name = all.Name
	return nil
}

type NullableIncidentServiceUpdateAttributes struct {
	value *IncidentServiceUpdateAttributes
	isSet bool
}

func (v NullableIncidentServiceUpdateAttributes) Get() *IncidentServiceUpdateAttributes {
	return v.value
}

func (v *NullableIncidentServiceUpdateAttributes) Set(val *IncidentServiceUpdateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentServiceUpdateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentServiceUpdateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentServiceUpdateAttributes(val *IncidentServiceUpdateAttributes) *NullableIncidentServiceUpdateAttributes {
	return &NullableIncidentServiceUpdateAttributes{value: val, isSet: true}
}

func (v NullableIncidentServiceUpdateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentServiceUpdateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
