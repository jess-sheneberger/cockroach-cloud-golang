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

// ClusterConfig struct for ClusterConfig
type ClusterConfig struct {
	Dedicated  *DedicatedHardwareConfig `json:"dedicated,omitempty"`
	Serverless *ServerlessClusterConfig `json:"serverless,omitempty"`
}

// NewClusterConfig instantiates a new ClusterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfig() *ClusterConfig {
	this := ClusterConfig{}
	return &this
}

// NewClusterConfigWithDefaults instantiates a new ClusterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigWithDefaults() *ClusterConfig {
	this := ClusterConfig{}
	return &this
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *ClusterConfig) GetDedicated() DedicatedHardwareConfig {
	if o == nil || o.Dedicated == nil {
		var ret DedicatedHardwareConfig
		return ret
	}
	return *o.Dedicated
}

// GetDedicatedOk returns a tuple with the Dedicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfig) GetDedicatedOk() (*DedicatedHardwareConfig, bool) {
	if o == nil || o.Dedicated == nil {
		return nil, false
	}
	return o.Dedicated, true
}

// HasDedicated returns a boolean if a field has been set.
func (o *ClusterConfig) HasDedicated() bool {
	if o != nil && o.Dedicated != nil {
		return true
	}

	return false
}

// SetDedicated gets a reference to the given DedicatedHardwareConfig and assigns it to the Dedicated field.
func (o *ClusterConfig) SetDedicated(v DedicatedHardwareConfig) {
	o.Dedicated = &v
}

// GetServerless returns the Serverless field value if set, zero value otherwise.
func (o *ClusterConfig) GetServerless() ServerlessClusterConfig {
	if o == nil || o.Serverless == nil {
		var ret ServerlessClusterConfig
		return ret
	}
	return *o.Serverless
}

// GetServerlessOk returns a tuple with the Serverless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfig) GetServerlessOk() (*ServerlessClusterConfig, bool) {
	if o == nil || o.Serverless == nil {
		return nil, false
	}
	return o.Serverless, true
}

// HasServerless returns a boolean if a field has been set.
func (o *ClusterConfig) HasServerless() bool {
	if o != nil && o.Serverless != nil {
		return true
	}

	return false
}

// SetServerless gets a reference to the given ServerlessClusterConfig and assigns it to the Serverless field.
func (o *ClusterConfig) SetServerless(v ServerlessClusterConfig) {
	o.Serverless = &v
}

func (o ClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dedicated != nil {
		toSerialize["dedicated"] = o.Dedicated
	}
	if o.Serverless != nil {
		toSerialize["serverless"] = o.Serverless
	}
	return json.Marshal(toSerialize)
}

type NullableClusterConfig struct {
	value *ClusterConfig
	isSet bool
}

func (v NullableClusterConfig) Get() *ClusterConfig {
	return v.value
}

func (v *NullableClusterConfig) Set(val *ClusterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfig(val *ClusterConfig) *NullableClusterConfig {
	return &NullableClusterConfig{value: val, isSet: true}
}

func (v NullableClusterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
