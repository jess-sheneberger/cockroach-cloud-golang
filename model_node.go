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

// Node struct for Node
type Node struct {
	Name string `json:"name"`
	RegionName string `json:"region_name"`
	Status NodeStatus `json:"status"`
}

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode(name string, regionName string, status NodeStatus) *Node {
	this := Node{}
	this.Name = name
	this.RegionName = regionName
	this.Status = status
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	var status NodeStatus = NODE_STATUS_UNSPECIFIED
	this.Status = status
	return &this
}

// GetName returns the Name field value
func (o *Node) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Node) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Node) SetName(v string) {
	o.Name = v
}

// GetRegionName returns the RegionName field value
func (o *Node) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *Node) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *Node) SetRegionName(v string) {
	o.RegionName = v
}

// GetStatus returns the Status field value
func (o *Node) GetStatus() NodeStatus {
	if o == nil {
		var ret NodeStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Node) GetStatusOk() (*NodeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Node) SetStatus(v NodeStatus) {
	o.Status = v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["region_name"] = o.RegionName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


