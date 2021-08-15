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

// TableWidgetDefinitionType Type of the table widget.
type TableWidgetDefinitionType string

// List of TableWidgetDefinitionType
const (
	TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE TableWidgetDefinitionType = "query_table"
)

var allowedTableWidgetDefinitionTypeEnumValues = []TableWidgetDefinitionType{
	"query_table",
}

func (w *TableWidgetDefinitionType) GetAllowedValues() []TableWidgetDefinitionType {
	return allowedTableWidgetDefinitionTypeEnumValues
}

func (v *TableWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	ev, err := NewTableWidgetDefinitionTypeFromValue(value)
	if err != nil {
		return err
	}
	*v = *ev
	return nil
}

// NewTableWidgetDefinitionTypeFromValue returns a pointer to a valid TableWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTableWidgetDefinitionTypeFromValue(v string) (*TableWidgetDefinitionType, error) {
	ev := TableWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TableWidgetDefinitionType: valid values are %v", v, allowedTableWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TableWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedTableWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetDefinitionType value
func (v TableWidgetDefinitionType) Ptr() *TableWidgetDefinitionType {
	return &v
}

type NullableTableWidgetDefinitionType struct {
	value *TableWidgetDefinitionType
	isSet bool
}

func (v NullableTableWidgetDefinitionType) Get() *TableWidgetDefinitionType {
	return v.value
}

func (v *NullableTableWidgetDefinitionType) Set(val *TableWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableTableWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableTableWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableWidgetDefinitionType(val *TableWidgetDefinitionType) *NullableTableWidgetDefinitionType {
	return &NullableTableWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableTableWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
