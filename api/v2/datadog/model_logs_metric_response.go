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

// LogsMetricResponse The log-based metric object.
type LogsMetricResponse struct {
	Data *LogsMetricResponseData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewLogsMetricResponse instantiates a new LogsMetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsMetricResponse() *LogsMetricResponse {
	this := LogsMetricResponse{}
	return &this
}

// NewLogsMetricResponseWithDefaults instantiates a new LogsMetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsMetricResponseWithDefaults() *LogsMetricResponse {
	this := LogsMetricResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LogsMetricResponse) GetData() LogsMetricResponseData {
	if o == nil || o.Data == nil {
		var ret LogsMetricResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricResponse) GetDataOk() (*LogsMetricResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LogsMetricResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LogsMetricResponseData and assigns it to the Data field.
func (o *LogsMetricResponse) SetData(v LogsMetricResponseData) {
	o.Data = &v
}

func (o LogsMetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *LogsMetricResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *LogsMetricResponseData `json:"data,omitempty"`
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
	if v := all.Data; !o.ContainsUnparsedObject && v != nil && v.ContainsUnparsedObject {
		o.ContainsUnparsedObject = true
	}

	o.Data = all.Data
	return nil
}

type NullableLogsMetricResponse struct {
	value *LogsMetricResponse
	isSet bool
}

func (v NullableLogsMetricResponse) Get() *LogsMetricResponse {
	return v.value
}

func (v *NullableLogsMetricResponse) Set(val *LogsMetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMetricResponse(val *LogsMetricResponse) *NullableLogsMetricResponse {
	return &NullableLogsMetricResponse{value: val, isSet: true}
}

func (v NullableLogsMetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
