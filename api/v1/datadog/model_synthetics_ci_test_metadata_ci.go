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

// SyntheticsCITestMetadataCi Describe CI provider.
type SyntheticsCITestMetadataCi struct {
	// Name of the pipeline.
	Pipeline *string `json:"pipeline,omitempty"`
	// Name of the CI provider.
	Provider *string `json:"provider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewSyntheticsCITestMetadataCi instantiates a new SyntheticsCITestMetadataCi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsCITestMetadataCi() *SyntheticsCITestMetadataCi {
	this := SyntheticsCITestMetadataCi{}
	return &this
}

// NewSyntheticsCITestMetadataCiWithDefaults instantiates a new SyntheticsCITestMetadataCi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsCITestMetadataCiWithDefaults() *SyntheticsCITestMetadataCi {
	this := SyntheticsCITestMetadataCi{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *SyntheticsCITestMetadataCi) GetPipeline() string {
	if o == nil || o.Pipeline == nil {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITestMetadataCi) GetPipelineOk() (*string, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *SyntheticsCITestMetadataCi) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *SyntheticsCITestMetadataCi) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SyntheticsCITestMetadataCi) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITestMetadataCi) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SyntheticsCITestMetadataCi) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SyntheticsCITestMetadataCi) SetProvider(v string) {
	o.Provider = &v
}

func (o SyntheticsCITestMetadataCi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsCITestMetadataCi) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Pipeline *string `json:"pipeline,omitempty"`
		Provider *string `json:"provider,omitempty"`
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

	o.Pipeline = all.Pipeline
	o.Provider = all.Provider
	return nil
}

type NullableSyntheticsCITestMetadataCi struct {
	value *SyntheticsCITestMetadataCi
	isSet bool
}

func (v NullableSyntheticsCITestMetadataCi) Get() *SyntheticsCITestMetadataCi {
	return v.value
}

func (v *NullableSyntheticsCITestMetadataCi) Set(val *SyntheticsCITestMetadataCi) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsCITestMetadataCi) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsCITestMetadataCi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsCITestMetadataCi(val *SyntheticsCITestMetadataCi) *NullableSyntheticsCITestMetadataCi {
	return &NullableSyntheticsCITestMetadataCi{value: val, isSet: true}
}

func (v NullableSyntheticsCITestMetadataCi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsCITestMetadataCi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
