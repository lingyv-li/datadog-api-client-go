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

// OrganizationSubscription Subscription definition.
type OrganizationSubscription struct {
	// The subscription type. Types available are `trial`, `free`, and `pro`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewOrganizationSubscription instantiates a new OrganizationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSubscription() *OrganizationSubscription {
	this := OrganizationSubscription{}
	return &this
}

// NewOrganizationSubscriptionWithDefaults instantiates a new OrganizationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSubscriptionWithDefaults() *OrganizationSubscription {
	this := OrganizationSubscription{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationSubscription) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSubscription) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationSubscription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationSubscription) SetType(v string) {
	o.Type = &v
}

func (o OrganizationSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *OrganizationSubscription) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Type *string `json:"type,omitempty"`
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

	o.Type = all.Type
	return nil
}

type NullableOrganizationSubscription struct {
	value *OrganizationSubscription
	isSet bool
}

func (v NullableOrganizationSubscription) Get() *OrganizationSubscription {
	return v.value
}

func (v *NullableOrganizationSubscription) Set(val *OrganizationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSubscription(val *OrganizationSubscription) *NullableOrganizationSubscription {
	return &NullableOrganizationSubscription{value: val, isSet: true}
}

func (v NullableOrganizationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
