# CMEKKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CMEKStatus**](CMEKStatus.md) |  | [optional] [default to UNKNOWN_STATUS]
**UserMessage** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**CMEKKeySpecification**](CMEKKeySpecification.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCMEKKeyInfo

`func NewCMEKKeyInfo() *CMEKKeyInfo`

NewCMEKKeyInfo instantiates a new CMEKKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMEKKeyInfoWithDefaults

`func NewCMEKKeyInfoWithDefaults() *CMEKKeyInfo`

NewCMEKKeyInfoWithDefaults instantiates a new CMEKKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CMEKKeyInfo) GetStatus() CMEKStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CMEKKeyInfo) GetStatusOk() (*CMEKStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CMEKKeyInfo) SetStatus(v CMEKStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CMEKKeyInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserMessage

`func (o *CMEKKeyInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *CMEKKeyInfo) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *CMEKKeyInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *CMEKKeyInfo) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.

### GetSpec

`func (o *CMEKKeyInfo) GetSpec() CMEKKeySpecification`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CMEKKeyInfo) GetSpecOk() (*CMEKKeySpecification, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CMEKKeyInfo) SetSpec(v CMEKKeySpecification)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CMEKKeyInfo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CMEKKeyInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CMEKKeyInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CMEKKeyInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CMEKKeyInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CMEKKeyInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CMEKKeyInfo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CMEKKeyInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CMEKKeyInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


