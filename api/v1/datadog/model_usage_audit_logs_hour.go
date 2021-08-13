/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageAuditLogsHour Audit logs usage for a given organization for a given hour.
type UsageAuditLogsHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The total number of audit logs lines indexed during a given hour.
	LinesIndexed *int64 `json:"lines_indexed,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewUsageAuditLogsHour instantiates a new UsageAuditLogsHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAuditLogsHour() *UsageAuditLogsHour {
	this := UsageAuditLogsHour{}
	return &this
}

// NewUsageAuditLogsHourWithDefaults instantiates a new UsageAuditLogsHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAuditLogsHourWithDefaults() *UsageAuditLogsHour {
	this := UsageAuditLogsHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageAuditLogsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageAuditLogsHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageAuditLogsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetLinesIndexed returns the LinesIndexed field value if set, zero value otherwise.
func (o *UsageAuditLogsHour) GetLinesIndexed() int64 {
	if o == nil || o.LinesIndexed == nil {
		var ret int64
		return ret
	}
	return *o.LinesIndexed
}

// GetLinesIndexedOk returns a tuple with the LinesIndexed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsHour) GetLinesIndexedOk() (*int64, bool) {
	if o == nil || o.LinesIndexed == nil {
		return nil, false
	}
	return o.LinesIndexed, true
}

// HasLinesIndexed returns a boolean if a field has been set.
func (o *UsageAuditLogsHour) HasLinesIndexed() bool {
	if o != nil && o.LinesIndexed != nil {
		return true
	}

	return false
}

// SetLinesIndexed gets a reference to the given int64 and assigns it to the LinesIndexed field.
func (o *UsageAuditLogsHour) SetLinesIndexed(v int64) {
	o.LinesIndexed = &v
}

func (o UsageAuditLogsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.LinesIndexed != nil {
		toSerialize["lines_indexed"] = o.LinesIndexed
	}
	return json.Marshal(toSerialize)
}

func (o *UsageAuditLogsHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Hour         *time.Time `json:"hour,omitempty"`
		LinesIndexed *int64     `json:"lines_indexed,omitempty"`
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

	o.Hour = all.Hour
	o.LinesIndexed = all.LinesIndexed
	return nil
}

type NullableUsageAuditLogsHour struct {
	value *UsageAuditLogsHour
	isSet bool
}

func (v NullableUsageAuditLogsHour) Get() *UsageAuditLogsHour {
	return v.value
}

func (v *NullableUsageAuditLogsHour) Set(val *UsageAuditLogsHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAuditLogsHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAuditLogsHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAuditLogsHour(val *UsageAuditLogsHour) *NullableUsageAuditLogsHour {
	return &NullableUsageAuditLogsHour{value: val, isSet: true}
}

func (v NullableUsageAuditLogsHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAuditLogsHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
