/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"reflect"
)

// DashboardListUpdateItemsResponse Response containing a list of updated dashboards.
type DashboardListUpdateItemsResponse struct {
	// List of dashboards in the dashboard list.
	Dashboards *[]DashboardListItemResponse `json:"dashboards,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject         map[string]interface{} `json:-`
	ContainsUnparsedObject bool                   `json:-`
}

// NewDashboardListUpdateItemsResponse instantiates a new DashboardListUpdateItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardListUpdateItemsResponse() *DashboardListUpdateItemsResponse {
	this := DashboardListUpdateItemsResponse{}
	return &this
}

// NewDashboardListUpdateItemsResponseWithDefaults instantiates a new DashboardListUpdateItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardListUpdateItemsResponseWithDefaults() *DashboardListUpdateItemsResponse {
	this := DashboardListUpdateItemsResponse{}
	return &this
}

// GetDashboards returns the Dashboards field value if set, zero value otherwise.
func (o *DashboardListUpdateItemsResponse) GetDashboards() []DashboardListItemResponse {
	if o == nil || o.Dashboards == nil {
		var ret []DashboardListItemResponse
		return ret
	}
	return *o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListUpdateItemsResponse) GetDashboardsOk() (*[]DashboardListItemResponse, bool) {
	if o == nil || o.Dashboards == nil {
		return nil, false
	}
	return o.Dashboards, true
}

// HasDashboards returns a boolean if a field has been set.
func (o *DashboardListUpdateItemsResponse) HasDashboards() bool {
	if o != nil && o.Dashboards != nil {
		return true
	}

	return false
}

// SetDashboards gets a reference to the given []DashboardListItemResponse and assigns it to the Dashboards field.
func (o *DashboardListUpdateItemsResponse) SetDashboards(v []DashboardListItemResponse) {
	o.Dashboards = &v
}

func (o DashboardListUpdateItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Dashboards != nil {
		toSerialize["dashboards"] = o.Dashboards
	}
	return json.Marshal(toSerialize)
}

func (o *DashboardListUpdateItemsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Dashboards *[]DashboardListItemResponse `json:"dashboards,omitempty"`
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

	if !o.ContainsUnparsedObject {
		if v := all.Dashboards; v != nil {
			o.ContainsUnparsedObject = containsUnparsedObject(reflect.ValueOf(*v))
		}
	}

	o.Dashboards = all.Dashboards
	return nil
}

type NullableDashboardListUpdateItemsResponse struct {
	value *DashboardListUpdateItemsResponse
	isSet bool
}

func (v NullableDashboardListUpdateItemsResponse) Get() *DashboardListUpdateItemsResponse {
	return v.value
}

func (v *NullableDashboardListUpdateItemsResponse) Set(val *DashboardListUpdateItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardListUpdateItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardListUpdateItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardListUpdateItemsResponse(val *DashboardListUpdateItemsResponse) *NullableDashboardListUpdateItemsResponse {
	return &NullableDashboardListUpdateItemsResponse{value: val, isSet: true}
}

func (v NullableDashboardListUpdateItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardListUpdateItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
