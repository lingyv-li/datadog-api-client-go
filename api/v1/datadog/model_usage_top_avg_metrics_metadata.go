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

// UsageTopAvgMetricsMetadata The object containing document metadata.
type UsageTopAvgMetricsMetadata struct {
	// The day value from the user request that contains the returned usage data. (If day was used the request)
	Day *time.Time `json:"day,omitempty"`
	// The month value from the user request that contains the returned usage data. (If month was used the request)
	Month      *time.Time                  `json:"month,omitempty"`
	Pagination *UsageAttributionPagination `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewUsageTopAvgMetricsMetadata instantiates a new UsageTopAvgMetricsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageTopAvgMetricsMetadata() *UsageTopAvgMetricsMetadata {
	this := UsageTopAvgMetricsMetadata{}
	return &this
}

// NewUsageTopAvgMetricsMetadataWithDefaults instantiates a new UsageTopAvgMetricsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageTopAvgMetricsMetadataWithDefaults() *UsageTopAvgMetricsMetadata {
	this := UsageTopAvgMetricsMetadata{}
	return &this
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *UsageTopAvgMetricsMetadata) GetDay() time.Time {
	if o == nil || o.Day == nil {
		var ret time.Time
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTopAvgMetricsMetadata) GetDayOk() (*time.Time, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *UsageTopAvgMetricsMetadata) HasDay() bool {
	if o != nil && o.Day != nil {
		return true
	}

	return false
}

// SetDay gets a reference to the given time.Time and assigns it to the Day field.
func (o *UsageTopAvgMetricsMetadata) SetDay(v time.Time) {
	o.Day = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *UsageTopAvgMetricsMetadata) GetMonth() time.Time {
	if o == nil || o.Month == nil {
		var ret time.Time
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTopAvgMetricsMetadata) GetMonthOk() (*time.Time, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *UsageTopAvgMetricsMetadata) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given time.Time and assigns it to the Month field.
func (o *UsageTopAvgMetricsMetadata) SetMonth(v time.Time) {
	o.Month = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *UsageTopAvgMetricsMetadata) GetPagination() UsageAttributionPagination {
	if o == nil || o.Pagination == nil {
		var ret UsageAttributionPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTopAvgMetricsMetadata) GetPaginationOk() (*UsageAttributionPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *UsageTopAvgMetricsMetadata) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given UsageAttributionPagination and assigns it to the Pagination field.
func (o *UsageTopAvgMetricsMetadata) SetPagination(v UsageAttributionPagination) {
	o.Pagination = &v
}

func (o UsageTopAvgMetricsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Day != nil {
		toSerialize["day"] = o.Day
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

func (o *UsageTopAvgMetricsMetadata) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Day        *time.Time                  `json:"day,omitempty"`
		Month      *time.Time                  `json:"month,omitempty"`
		Pagination *UsageAttributionPagination `json:"pagination,omitempty"`
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

	if v := all.Pagination; !o.ContainsUnparsedObject && v != nil && v.ContainsUnparsedObject {
		o.ContainsUnparsedObject = true
	}

	o.Day = all.Day
	o.Month = all.Month
	o.Pagination = all.Pagination
	return nil
}

type NullableUsageTopAvgMetricsMetadata struct {
	value *UsageTopAvgMetricsMetadata
	isSet bool
}

func (v NullableUsageTopAvgMetricsMetadata) Get() *UsageTopAvgMetricsMetadata {
	return v.value
}

func (v *NullableUsageTopAvgMetricsMetadata) Set(val *UsageTopAvgMetricsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageTopAvgMetricsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageTopAvgMetricsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageTopAvgMetricsMetadata(val *UsageTopAvgMetricsMetadata) *NullableUsageTopAvgMetricsMetadata {
	return &NullableUsageTopAvgMetricsMetadata{value: val, isSet: true}
}

func (v NullableUsageTopAvgMetricsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageTopAvgMetricsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
