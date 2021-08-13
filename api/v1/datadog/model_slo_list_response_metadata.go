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

// SLOListResponseMetadata The metadata object containing additional information about the list of SLOs.
type SLOListResponseMetadata struct {
	Page *SLOListResponseMetadataPage `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewSLOListResponseMetadata instantiates a new SLOListResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOListResponseMetadata() *SLOListResponseMetadata {
	this := SLOListResponseMetadata{}
	return &this
}

// NewSLOListResponseMetadataWithDefaults instantiates a new SLOListResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOListResponseMetadataWithDefaults() *SLOListResponseMetadata {
	this := SLOListResponseMetadata{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SLOListResponseMetadata) GetPage() SLOListResponseMetadataPage {
	if o == nil || o.Page == nil {
		var ret SLOListResponseMetadataPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOListResponseMetadata) GetPageOk() (*SLOListResponseMetadataPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SLOListResponseMetadata) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given SLOListResponseMetadataPage and assigns it to the Page field.
func (o *SLOListResponseMetadata) SetPage(v SLOListResponseMetadataPage) {
	o.Page = &v
}

func (o SLOListResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

func (o *SLOListResponseMetadata) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Page *SLOListResponseMetadataPage `json:"page,omitempty"`
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
	if v := all.Page; !o.ContainsUnparsedObject && v != nil && v.ContainsUnparsedObject {
		o.ContainsUnparsedObject = true
	}

	o.Page = all.Page
	return nil
}

type NullableSLOListResponseMetadata struct {
	value *SLOListResponseMetadata
	isSet bool
}

func (v NullableSLOListResponseMetadata) Get() *SLOListResponseMetadata {
	return v.value
}

func (v *NullableSLOListResponseMetadata) Set(val *SLOListResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOListResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOListResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOListResponseMetadata(val *SLOListResponseMetadata) *NullableSLOListResponseMetadata {
	return &NullableSLOListResponseMetadata{value: val, isSet: true}
}

func (v NullableSLOListResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOListResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
