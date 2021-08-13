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

// MetricTagConfigurationUpdateAttributes Object containing the definition of a metric tag configuration to be updated.
type MetricTagConfigurationUpdateAttributes struct {
	// Toggle to include/exclude percentiles for a distribution metric. Defaults to false. Can only be applied to metrics that have a `metric_type` of `distribution`.
	IncludePercentiles *bool `json:"include_percentiles,omitempty"`
	// A list of tag keys that will be queryable for your metric.
	Tags *[]string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewMetricTagConfigurationUpdateAttributes instantiates a new MetricTagConfigurationUpdateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationUpdateAttributes() *MetricTagConfigurationUpdateAttributes {
	this := MetricTagConfigurationUpdateAttributes{}
	var includePercentiles bool = false
	this.IncludePercentiles = &includePercentiles
	return &this
}

// NewMetricTagConfigurationUpdateAttributesWithDefaults instantiates a new MetricTagConfigurationUpdateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationUpdateAttributesWithDefaults() *MetricTagConfigurationUpdateAttributes {
	this := MetricTagConfigurationUpdateAttributes{}
	var includePercentiles bool = false
	this.IncludePercentiles = &includePercentiles
	return &this
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *MetricTagConfigurationUpdateAttributes) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationUpdateAttributes) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *MetricTagConfigurationUpdateAttributes) HasIncludePercentiles() bool {
	if o != nil && o.IncludePercentiles != nil {
		return true
	}

	return false
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *MetricTagConfigurationUpdateAttributes) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricTagConfigurationUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricTagConfigurationUpdateAttributes) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricTagConfigurationUpdateAttributes) SetTags(v []string) {
	o.Tags = &v
}

func (o MetricTagConfigurationUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

func (o *MetricTagConfigurationUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		IncludePercentiles *bool     `json:"include_percentiles,omitempty"`
		Tags               *[]string `json:"tags,omitempty"`
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

	o.IncludePercentiles = all.IncludePercentiles
	o.Tags = all.Tags
	return nil
}

type NullableMetricTagConfigurationUpdateAttributes struct {
	value *MetricTagConfigurationUpdateAttributes
	isSet bool
}

func (v NullableMetricTagConfigurationUpdateAttributes) Get() *MetricTagConfigurationUpdateAttributes {
	return v.value
}

func (v *NullableMetricTagConfigurationUpdateAttributes) Set(val *MetricTagConfigurationUpdateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationUpdateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationUpdateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationUpdateAttributes(val *MetricTagConfigurationUpdateAttributes) *NullableMetricTagConfigurationUpdateAttributes {
	return &NullableMetricTagConfigurationUpdateAttributes{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationUpdateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationUpdateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
