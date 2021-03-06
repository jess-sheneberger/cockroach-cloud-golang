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

// CMEKRegionSpecification CMEKRegionSpecification declares the customer-provided key specification that should be used in a given region.
type CMEKRegionSpecification struct {
	Region  *string               `json:"region,omitempty"`
	KeySpec *CMEKKeySpecification `json:"key_spec,omitempty"`
}

// NewCMEKRegionSpecification instantiates a new CMEKRegionSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKRegionSpecification() *CMEKRegionSpecification {
	this := CMEKRegionSpecification{}
	return &this
}

// NewCMEKRegionSpecificationWithDefaults instantiates a new CMEKRegionSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCMEKRegionSpecificationWithDefaults() *CMEKRegionSpecification {
	this := CMEKRegionSpecification{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CMEKRegionSpecification) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CMEKRegionSpecification) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CMEKRegionSpecification) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CMEKRegionSpecification) SetRegion(v string) {
	o.Region = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise.
func (o *CMEKRegionSpecification) GetKeySpec() CMEKKeySpecification {
	if o == nil || o.KeySpec == nil {
		var ret CMEKKeySpecification
		return ret
	}
	return *o.KeySpec
}

// GetKeySpecOk returns a tuple with the KeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CMEKRegionSpecification) GetKeySpecOk() (*CMEKKeySpecification, bool) {
	if o == nil || o.KeySpec == nil {
		return nil, false
	}
	return o.KeySpec, true
}

// HasKeySpec returns a boolean if a field has been set.
func (o *CMEKRegionSpecification) HasKeySpec() bool {
	if o != nil && o.KeySpec != nil {
		return true
	}

	return false
}

// SetKeySpec gets a reference to the given CMEKKeySpecification and assigns it to the KeySpec field.
func (o *CMEKRegionSpecification) SetKeySpec(v CMEKKeySpecification) {
	o.KeySpec = &v
}

func (o CMEKRegionSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.KeySpec != nil {
		toSerialize["key_spec"] = o.KeySpec
	}
	return json.Marshal(toSerialize)
}

type NullableCMEKRegionSpecification struct {
	value *CMEKRegionSpecification
	isSet bool
}

func (v NullableCMEKRegionSpecification) Get() *CMEKRegionSpecification {
	return v.value
}

func (v *NullableCMEKRegionSpecification) Set(val *CMEKRegionSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableCMEKRegionSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableCMEKRegionSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCMEKRegionSpecification(val *CMEKRegionSpecification) *NullableCMEKRegionSpecification {
	return &NullableCMEKRegionSpecification{value: val, isSet: true}
}

func (v NullableCMEKRegionSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCMEKRegionSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
