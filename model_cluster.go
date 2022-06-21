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

// Cluster struct for Cluster
type Cluster struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	CockroachVersion string            `json:"cockroach_version"`
	Plan             Plan              `json:"plan"`
	CloudProvider    ApiCloudProvider  `json:"cloud_provider"`
	AccountId        *string           `json:"account_id,omitempty"`
	State            ClusterStateType  `json:"state"`
	CreatorId        string            `json:"creator_id"`
	OperationStatus  ClusterStatusType `json:"operation_status"`
	Config           ClusterConfig     `json:"config"`
	Regions          []Region          `json:"regions"`
	CreatedAt        *time.Time        `json:"created_at,omitempty"`
	UpdatedAt        *time.Time        `json:"updated_at,omitempty"`
	DeletedAt        *time.Time        `json:"deleted_at,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(id string, name string, cockroachVersion string, plan Plan, cloudProvider ApiCloudProvider, state ClusterStateType, creatorId string, operationStatus ClusterStatusType, config ClusterConfig, regions []Region) *Cluster {
	this := Cluster{}
	this.Id = id
	this.Name = name
	this.CockroachVersion = cockroachVersion
	this.Plan = plan
	this.CloudProvider = cloudProvider
	this.State = state
	this.CreatorId = creatorId
	this.OperationStatus = operationStatus
	this.Config = config
	this.Regions = regions
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	var plan Plan = PLAN_UNSPECIFIED
	this.Plan = plan
	var cloudProvider ApiCloudProvider = CLOUD_PROVIDER_UNSPECIFIED
	this.CloudProvider = cloudProvider
	var state ClusterStateType = CLUSTER_STATE_UNSPECIFIED
	this.State = state
	var operationStatus ClusterStatusType = CLUSTER_STATUS_UNSPECIFIED
	this.OperationStatus = operationStatus
	return &this
}

// GetId returns the Id field value
func (o *Cluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetCockroachVersion returns the CockroachVersion field value
func (o *Cluster) GetCockroachVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CockroachVersion
}

// GetCockroachVersionOk returns a tuple with the CockroachVersion field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCockroachVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CockroachVersion, true
}

// SetCockroachVersion sets field value
func (o *Cluster) SetCockroachVersion(v string) {
	o.CockroachVersion = v
}

// GetPlan returns the Plan field value
func (o *Cluster) GetPlan() Plan {
	if o == nil {
		var ret Plan
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetPlanOk() (*Plan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *Cluster) SetPlan(v Plan) {
	o.Plan = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *Cluster) GetCloudProvider() ApiCloudProvider {
	if o == nil {
		var ret ApiCloudProvider
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCloudProviderOk() (*ApiCloudProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *Cluster) SetCloudProvider(v ApiCloudProvider) {
	o.CloudProvider = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Cluster) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Cluster) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Cluster) SetAccountId(v string) {
	o.AccountId = &v
}

// GetState returns the State field value
func (o *Cluster) GetState() ClusterStateType {
	if o == nil {
		var ret ClusterStateType
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetStateOk() (*ClusterStateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Cluster) SetState(v ClusterStateType) {
	o.State = v
}

// GetCreatorId returns the CreatorId field value
func (o *Cluster) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *Cluster) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetOperationStatus returns the OperationStatus field value
func (o *Cluster) GetOperationStatus() ClusterStatusType {
	if o == nil {
		var ret ClusterStatusType
		return ret
	}

	return o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetOperationStatusOk() (*ClusterStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationStatus, true
}

// SetOperationStatus sets field value
func (o *Cluster) SetOperationStatus(v ClusterStatusType) {
	o.OperationStatus = v
}

// GetConfig returns the Config field value
func (o *Cluster) GetConfig() ClusterConfig {
	if o == nil {
		var ret ClusterConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetConfigOk() (*ClusterConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Cluster) SetConfig(v ClusterConfig) {
	o.Config = v
}

// GetRegions returns the Regions field value
func (o *Cluster) GetRegions() []Region {
	if o == nil {
		var ret []Region
		return ret
	}

	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetRegionsOk() ([]Region, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regions, true
}

// SetRegions sets field value
func (o *Cluster) SetRegions(v []Region) {
	o.Regions = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Cluster) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Cluster) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Cluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Cluster) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Cluster) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Cluster) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Cluster) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Cluster) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Cluster) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cockroach_version"] = o.CockroachVersion
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["creator_id"] = o.CreatorId
	}
	if true {
		toSerialize["operation_status"] = o.OperationStatus
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["regions"] = o.Regions
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
