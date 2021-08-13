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

// MetricTagConfigurationResponse Response object which includes a single metric's tag configuration.
type MetricTagConfigurationResponse struct {
	Data *MetricTagConfiguration `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewMetricTagConfigurationResponse instantiates a new MetricTagConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationResponse() *MetricTagConfigurationResponse {
	this := MetricTagConfigurationResponse{}
	return &this
}

// NewMetricTagConfigurationResponseWithDefaults instantiates a new MetricTagConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationResponseWithDefaults() *MetricTagConfigurationResponse {
	this := MetricTagConfigurationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MetricTagConfigurationResponse) GetData() MetricTagConfiguration {
	if o == nil || o.Data == nil {
		var ret MetricTagConfiguration
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationResponse) GetDataOk() (*MetricTagConfiguration, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MetricTagConfigurationResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given MetricTagConfiguration and assigns it to the Data field.
func (o *MetricTagConfigurationResponse) SetData(v MetricTagConfiguration) {
	o.Data = &v
}

func (o MetricTagConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *MetricTagConfigurationResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *MetricTagConfiguration `json:"data,omitempty"`
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

type NullableMetricTagConfigurationResponse struct {
	value *MetricTagConfigurationResponse
	isSet bool
}

func (v NullableMetricTagConfigurationResponse) Get() *MetricTagConfigurationResponse {
	return v.value
}

func (v *NullableMetricTagConfigurationResponse) Set(val *MetricTagConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationResponse(val *MetricTagConfigurationResponse) *NullableMetricTagConfigurationResponse {
	return &NullableMetricTagConfigurationResponse{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
