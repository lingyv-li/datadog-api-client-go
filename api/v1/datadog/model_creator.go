/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// Creator struct for Creator
type Creator struct {
	Email  *string `json:"email,omitempty"`
	Handle *string `json:"handle,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Creator) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Creator) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Creator) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Creator) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *Creator) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Creator) GetHandleOk() (string, bool) {
	if o == nil || o.Handle == nil {
		var ret string
		return ret, false
	}
	return *o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *Creator) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *Creator) SetHandle(v string) {
	o.Handle = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Creator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Creator) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Creator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Creator) SetName(v string) {
	o.Name = &v
}

type NullableCreator struct {
	Value        Creator
	ExplicitNull bool
}

func (v NullableCreator) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreator) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}