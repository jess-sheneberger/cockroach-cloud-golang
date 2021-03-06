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
	"fmt"
)

// PageDirection the model 'PageDirection'
type PageDirection string

// List of PageDirection
const (
	NEXT PageDirection = "PAGE_DIRECTION_NEXT"
	LAST PageDirection = "PAGE_DIRECTION_LAST"
)

// All allowed values of PageDirection enum
var AllowedPageDirectionEnumValues = []PageDirection{
	"PAGE_DIRECTION_NEXT",
	"PAGE_DIRECTION_LAST",
}

func (v *PageDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PageDirection(value)
	for _, existing := range AllowedPageDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PageDirection", value)
}

// NewPageDirectionFromValue returns a pointer to a valid PageDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPageDirectionFromValue(v string) (*PageDirection, error) {
	ev := PageDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PageDirection: valid values are %v", v, AllowedPageDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PageDirection) IsValid() bool {
	for _, existing := range AllowedPageDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PageDirection value
func (v PageDirection) Ptr() *PageDirection {
	return &v
}

type NullablePageDirection struct {
	value *PageDirection
	isSet bool
}

func (v NullablePageDirection) Get() *PageDirection {
	return v.value
}

func (v *NullablePageDirection) Set(val *PageDirection) {
	v.value = val
	v.isSet = true
}

func (v NullablePageDirection) IsSet() bool {
	return v.isSet
}

func (v *NullablePageDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageDirection(val *PageDirection) *NullablePageDirection {
	return &NullablePageDirection{value: val, isSet: true}
}

func (v NullablePageDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
