/*
CockroachDB Cloud API

This is an early access, experimental version of the Cloud API. The interface and output is subject to change, and there may be bugs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 2022-03-31
Contact: support@cockroachlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateClusterSpecification struct for UpdateClusterSpecification
type UpdateClusterSpecification struct {
	Dedicated *DedicatedClusterUpdateSpecification `json:"dedicated,omitempty"`
	Serverless *ServerlessClusterUpdateSpecification `json:"serverless,omitempty"`
}

// NewUpdateClusterSpecification instantiates a new UpdateClusterSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClusterSpecification() *UpdateClusterSpecification {
	this := UpdateClusterSpecification{}
	return &this
}

// NewUpdateClusterSpecificationWithDefaults instantiates a new UpdateClusterSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClusterSpecificationWithDefaults() *UpdateClusterSpecification {
	this := UpdateClusterSpecification{}
	return &this
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetDedicated() DedicatedClusterUpdateSpecification {
	if o == nil || o.Dedicated == nil {
		var ret DedicatedClusterUpdateSpecification
		return ret
	}
	return *o.Dedicated
}

// GetDedicatedOk returns a tuple with the Dedicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterSpecification) GetDedicatedOk() (*DedicatedClusterUpdateSpecification, bool) {
	if o == nil || o.Dedicated == nil {
		return nil, false
	}
	return o.Dedicated, true
}

// HasDedicated returns a boolean if a field has been set.
func (o *UpdateClusterSpecification) HasDedicated() bool {
	if o != nil && o.Dedicated != nil {
		return true
	}

	return false
}

// SetDedicated gets a reference to the given DedicatedClusterUpdateSpecification and assigns it to the Dedicated field.
func (o *UpdateClusterSpecification) SetDedicated(v DedicatedClusterUpdateSpecification) {
	o.Dedicated = &v
}

// GetServerless returns the Serverless field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetServerless() ServerlessClusterUpdateSpecification {
	if o == nil || o.Serverless == nil {
		var ret ServerlessClusterUpdateSpecification
		return ret
	}
	return *o.Serverless
}

// GetServerlessOk returns a tuple with the Serverless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterSpecification) GetServerlessOk() (*ServerlessClusterUpdateSpecification, bool) {
	if o == nil || o.Serverless == nil {
		return nil, false
	}
	return o.Serverless, true
}

// HasServerless returns a boolean if a field has been set.
func (o *UpdateClusterSpecification) HasServerless() bool {
	if o != nil && o.Serverless != nil {
		return true
	}

	return false
}

// SetServerless gets a reference to the given ServerlessClusterUpdateSpecification and assigns it to the Serverless field.
func (o *UpdateClusterSpecification) SetServerless(v ServerlessClusterUpdateSpecification) {
	o.Serverless = &v
}

func (o UpdateClusterSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dedicated != nil {
		toSerialize["dedicated"] = o.Dedicated
	}
	if o.Serverless != nil {
		toSerialize["serverless"] = o.Serverless
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateClusterSpecification struct {
	value *UpdateClusterSpecification
	isSet bool
}

func (v NullableUpdateClusterSpecification) Get() *UpdateClusterSpecification {
	return v.value
}

func (v *NullableUpdateClusterSpecification) Set(val *UpdateClusterSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClusterSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClusterSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClusterSpecification(val *UpdateClusterSpecification) *NullableUpdateClusterSpecification {
	return &NullableUpdateClusterSpecification{value: val, isSet: true}
}

func (v NullableUpdateClusterSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClusterSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

