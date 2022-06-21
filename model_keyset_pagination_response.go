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
	"time"
)

// KeysetPaginationResponse struct for KeysetPaginationResponse
type KeysetPaginationResponse struct {
	Next  *string    `json:"next,omitempty"`
	Last  *string    `json:"last,omitempty"`
	Limit *int32     `json:"limit,omitempty"`
	Time  *time.Time `json:"time,omitempty"`
	Order *SortOrder `json:"order,omitempty"`
}

// NewKeysetPaginationResponse instantiates a new KeysetPaginationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysetPaginationResponse() *KeysetPaginationResponse {
	this := KeysetPaginationResponse{}
	var order SortOrder = ASC
	this.Order = &order
	return &this
}

// NewKeysetPaginationResponseWithDefaults instantiates a new KeysetPaginationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeysetPaginationResponseWithDefaults() *KeysetPaginationResponse {
	this := KeysetPaginationResponse{}
	var order SortOrder = ASC
	this.Order = &order
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysetPaginationResponse) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *KeysetPaginationResponse) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *KeysetPaginationResponse) SetNext(v string) {
	o.Next = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysetPaginationResponse) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *KeysetPaginationResponse) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *KeysetPaginationResponse) SetLast(v string) {
	o.Last = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysetPaginationResponse) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *KeysetPaginationResponse) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *KeysetPaginationResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysetPaginationResponse) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *KeysetPaginationResponse) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *KeysetPaginationResponse) SetTime(v time.Time) {
	o.Time = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetOrder() SortOrder {
	if o == nil || o.Order == nil {
		var ret SortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeysetPaginationResponse) GetOrderOk() (*SortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *KeysetPaginationResponse) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given SortOrder and assigns it to the Order field.
func (o *KeysetPaginationResponse) SetOrder(v SortOrder) {
	o.Order = &v
}

func (o KeysetPaginationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableKeysetPaginationResponse struct {
	value *KeysetPaginationResponse
	isSet bool
}

func (v NullableKeysetPaginationResponse) Get() *KeysetPaginationResponse {
	return v.value
}

func (v *NullableKeysetPaginationResponse) Set(val *KeysetPaginationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeysetPaginationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeysetPaginationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeysetPaginationResponse(val *KeysetPaginationResponse) *NullableKeysetPaginationResponse {
	return &NullableKeysetPaginationResponse{value: val, isSet: true}
}

func (v NullableKeysetPaginationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeysetPaginationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
