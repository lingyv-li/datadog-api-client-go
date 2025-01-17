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

// MonitorGroupSearchResponseCountsStatus A facet item.
type MonitorGroupSearchResponseCountsStatus struct {
	// The number of found monitors with the listed value.
	Count *int64 `json:"count,omitempty"`
	// The facet value.
	Name interface{} `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewMonitorGroupSearchResponseCountsStatus instantiates a new MonitorGroupSearchResponseCountsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorGroupSearchResponseCountsStatus() *MonitorGroupSearchResponseCountsStatus {
	this := MonitorGroupSearchResponseCountsStatus{}
	return &this
}

// NewMonitorGroupSearchResponseCountsStatusWithDefaults instantiates a new MonitorGroupSearchResponseCountsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorGroupSearchResponseCountsStatusWithDefaults() *MonitorGroupSearchResponseCountsStatus {
	this := MonitorGroupSearchResponseCountsStatus{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MonitorGroupSearchResponseCountsStatus) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroupSearchResponseCountsStatus) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MonitorGroupSearchResponseCountsStatus) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *MonitorGroupSearchResponseCountsStatus) SetCount(v int64) {
	o.Count = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorGroupSearchResponseCountsStatus) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitorGroupSearchResponseCountsStatus) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorGroupSearchResponseCountsStatus) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *MonitorGroupSearchResponseCountsStatus) SetName(v interface{}) {
	o.Name = v
}

func (o MonitorGroupSearchResponseCountsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (o *MonitorGroupSearchResponseCountsStatus) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Count *int64      `json:"count,omitempty"`
		Name  interface{} `json:"name,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Count = all.Count
	o.Name = all.Name
	return nil
}

type NullableMonitorGroupSearchResponseCountsStatus struct {
	value *MonitorGroupSearchResponseCountsStatus
	isSet bool
}

func (v NullableMonitorGroupSearchResponseCountsStatus) Get() *MonitorGroupSearchResponseCountsStatus {
	return v.value
}

func (v *NullableMonitorGroupSearchResponseCountsStatus) Set(val *MonitorGroupSearchResponseCountsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorGroupSearchResponseCountsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorGroupSearchResponseCountsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorGroupSearchResponseCountsStatus(val *MonitorGroupSearchResponseCountsStatus) *NullableMonitorGroupSearchResponseCountsStatus {
	return &NullableMonitorGroupSearchResponseCountsStatus{value: val, isSet: true}
}

func (v NullableMonitorGroupSearchResponseCountsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorGroupSearchResponseCountsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
