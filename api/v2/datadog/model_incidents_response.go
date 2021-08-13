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

// IncidentsResponse Response with a list of incidents.
type IncidentsResponse struct {
	// An array of incidents.
	Data []IncidentResponseData `json:"data"`
	// Included related resources that the user requested.
	Included *[]IncidentResponseIncludedItem `json:"included,omitempty"`
	Meta     *IncidentServicesResponseMeta   `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewIncidentsResponse instantiates a new IncidentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentsResponse(data []IncidentResponseData) *IncidentsResponse {
	this := IncidentsResponse{}
	this.Data = data
	return &this
}

// NewIncidentsResponseWithDefaults instantiates a new IncidentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentsResponseWithDefaults() *IncidentsResponse {
	this := IncidentsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentsResponse) GetData() []IncidentResponseData {
	if o == nil {
		var ret []IncidentResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentsResponse) GetDataOk() (*[]IncidentResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentsResponse) SetData(v []IncidentResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentsResponse) GetIncluded() []IncidentResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentResponseIncludedItem
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsResponse) GetIncludedOk() (*[]IncidentResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncidentResponseIncludedItem and assigns it to the Included field.
func (o *IncidentsResponse) SetIncluded(v []IncidentResponseIncludedItem) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentsResponse) GetMeta() IncidentServicesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentServicesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsResponse) GetMetaOk() (*IncidentServicesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given IncidentServicesResponseMeta and assigns it to the Meta field.
func (o *IncidentsResponse) SetMeta(v IncidentServicesResponseMeta) {
	o.Meta = &v
}

func (o IncidentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *[]IncidentResponseData `json:"data"`
	}{}
	all := struct {
		Data     []IncidentResponseData          `json:"data"`
		Included *[]IncidentResponseIncludedItem `json:"included,omitempty"`
		Meta     *IncidentServicesResponseMeta   `json:"meta,omitempty"`
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

	if !o.ContainsUnparsedObject {
		if v := all.Data; v != nil {
			o.ContainsUnparsedObject = containsUnparsedObject(reflect.ValueOf(v))
		}
	}

	if !o.ContainsUnparsedObject {
		if v := all.Included; v != nil {
			o.ContainsUnparsedObject = containsUnparsedObject(reflect.ValueOf(*v))
		}
	}

	if v := all.Meta; !o.ContainsUnparsedObject && v != nil && v.ContainsUnparsedObject {
		o.ContainsUnparsedObject = true
	}

	o.Data = all.Data
	o.Included = all.Included
	o.Meta = all.Meta
	return nil
}

type NullableIncidentsResponse struct {
	value *IncidentsResponse
	isSet bool
}

func (v NullableIncidentsResponse) Get() *IncidentsResponse {
	return v.value
}

func (v *NullableIncidentsResponse) Set(val *IncidentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentsResponse(val *IncidentsResponse) *NullableIncidentsResponse {
	return &NullableIncidentsResponse{value: val, isSet: true}
}

func (v NullableIncidentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
