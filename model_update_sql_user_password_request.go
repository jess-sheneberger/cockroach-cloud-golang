/*
CockroachDB Cloud API

This is an early access, experimental version of the Cloud API. The interface and output is subject to change, and there may be bugs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 2022-03-31
Contact: support@cockroachlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccloud

import (
	"encoding/json"
)

// UpdateSQLUserPasswordRequest struct for UpdateSQLUserPasswordRequest
type UpdateSQLUserPasswordRequest struct {
	Password string `json:"password"`
}

// NewUpdateSQLUserPasswordRequest instantiates a new UpdateSQLUserPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSQLUserPasswordRequest(password string) *UpdateSQLUserPasswordRequest {
	this := UpdateSQLUserPasswordRequest{}
	this.Password = password
	return &this
}

// NewUpdateSQLUserPasswordRequestWithDefaults instantiates a new UpdateSQLUserPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSQLUserPasswordRequestWithDefaults() *UpdateSQLUserPasswordRequest {
	this := UpdateSQLUserPasswordRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *UpdateSQLUserPasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateSQLUserPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateSQLUserPasswordRequest) SetPassword(v string) {
	o.Password = v
}

func (o UpdateSQLUserPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSQLUserPasswordRequest struct {
	value *UpdateSQLUserPasswordRequest
	isSet bool
}

func (v NullableUpdateSQLUserPasswordRequest) Get() *UpdateSQLUserPasswordRequest {
	return v.value
}

func (v *NullableUpdateSQLUserPasswordRequest) Set(val *UpdateSQLUserPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSQLUserPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSQLUserPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSQLUserPasswordRequest(val *UpdateSQLUserPasswordRequest) *NullableUpdateSQLUserPasswordRequest {
	return &NullableUpdateSQLUserPasswordRequest{value: val, isSet: true}
}

func (v NullableUpdateSQLUserPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSQLUserPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
