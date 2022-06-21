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

// ListSQLUsersResponse struct for ListSQLUsersResponse
type ListSQLUsersResponse struct {
	Users []SQLUser `json:"users"`
	Pagination *KeysetPaginationResponse `json:"pagination,omitempty"`
}

// NewListSQLUsersResponse instantiates a new ListSQLUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSQLUsersResponse(users []SQLUser) *ListSQLUsersResponse {
	this := ListSQLUsersResponse{}
	this.Users = users
	return &this
}

// NewListSQLUsersResponseWithDefaults instantiates a new ListSQLUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSQLUsersResponseWithDefaults() *ListSQLUsersResponse {
	this := ListSQLUsersResponse{}
	return &this
}

// GetUsers returns the Users field value
func (o *ListSQLUsersResponse) GetUsers() []SQLUser {
	if o == nil {
		var ret []SQLUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ListSQLUsersResponse) GetUsersOk() ([]SQLUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ListSQLUsersResponse) SetUsers(v []SQLUser) {
	o.Users = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListSQLUsersResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSQLUsersResponse) GetPaginationOk() (*KeysetPaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListSQLUsersResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ListSQLUsersResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}

func (o ListSQLUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["users"] = o.Users
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListSQLUsersResponse struct {
	value *ListSQLUsersResponse
	isSet bool
}

func (v NullableListSQLUsersResponse) Get() *ListSQLUsersResponse {
	return v.value
}

func (v *NullableListSQLUsersResponse) Set(val *ListSQLUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSQLUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSQLUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSQLUsersResponse(val *ListSQLUsersResponse) *NullableListSQLUsersResponse {
	return &NullableListSQLUsersResponse{value: val, isSet: true}
}

func (v NullableListSQLUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSQLUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


