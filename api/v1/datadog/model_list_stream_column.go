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
)

// ListStreamColumn Widget column.
type ListStreamColumn struct {
	// Widget column field.
	Field string                `json:"field"`
	Width ListStreamColumnWidth `json:"width"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewListStreamColumn instantiates a new ListStreamColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStreamColumn(field string, width ListStreamColumnWidth) *ListStreamColumn {
	this := ListStreamColumn{}
	this.Field = field
	this.Width = width
	return &this
}

// NewListStreamColumnWithDefaults instantiates a new ListStreamColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStreamColumnWithDefaults() *ListStreamColumn {
	this := ListStreamColumn{}
	return &this
}

// GetField returns the Field field value
func (o *ListStreamColumn) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ListStreamColumn) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *ListStreamColumn) SetField(v string) {
	o.Field = v
}

// GetWidth returns the Width field value
func (o *ListStreamColumn) GetWidth() ListStreamColumnWidth {
	if o == nil {
		var ret ListStreamColumnWidth
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ListStreamColumn) GetWidthOk() (*ListStreamColumnWidth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ListStreamColumn) SetWidth(v ListStreamColumnWidth) {
	o.Width = v
}

func (o ListStreamColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}

func (o *ListStreamColumn) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Field *string                `json:"field"`
		Width *ListStreamColumnWidth `json:"width"`
	}{}
	all := struct {
		Field string                `json:"field"`
		Width ListStreamColumnWidth `json:"width"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Field == nil {
		return fmt.Errorf("Required field field missing")
	}
	if required.Width == nil {
		return fmt.Errorf("Required field width missing")
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

	if v := all.Width; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.ContainsUnparsedObject = true
		o.UnparsedObject = raw
		return nil
	}

	o.Field = all.Field
	o.Width = all.Width
	return nil
}

type NullableListStreamColumn struct {
	value *ListStreamColumn
	isSet bool
}

func (v NullableListStreamColumn) Get() *ListStreamColumn {
	return v.value
}

func (v *NullableListStreamColumn) Set(val *ListStreamColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableListStreamColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableListStreamColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStreamColumn(val *ListStreamColumn) *NullableListStreamColumn {
	return &NullableListStreamColumn{value: val, isSet: true}
}

func (v NullableListStreamColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStreamColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
