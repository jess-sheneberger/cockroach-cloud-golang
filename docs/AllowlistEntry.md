# AllowlistEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrIp** | **string** |  | 
**CidrMask** | **int32** |  | 
**Ui** | **bool** |  | 
**Sql** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewAllowlistEntry

`func NewAllowlistEntry(cidrIp string, cidrMask int32, ui bool, sql bool, ) *AllowlistEntry`

NewAllowlistEntry instantiates a new AllowlistEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistEntryWithDefaults

`func NewAllowlistEntryWithDefaults() *AllowlistEntry`

NewAllowlistEntryWithDefaults instantiates a new AllowlistEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrIp

`func (o *AllowlistEntry) GetCidrIp() string`

GetCidrIp returns the CidrIp field if non-nil, zero value otherwise.

### GetCidrIpOk

`func (o *AllowlistEntry) GetCidrIpOk() (*string, bool)`

GetCidrIpOk returns a tuple with the CidrIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIp

`func (o *AllowlistEntry) SetCidrIp(v string)`

SetCidrIp sets CidrIp field to given value.


### GetCidrMask

`func (o *AllowlistEntry) GetCidrMask() int32`

GetCidrMask returns the CidrMask field if non-nil, zero value otherwise.

### GetCidrMaskOk

`func (o *AllowlistEntry) GetCidrMaskOk() (*int32, bool)`

GetCidrMaskOk returns a tuple with the CidrMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrMask

`func (o *AllowlistEntry) SetCidrMask(v int32)`

SetCidrMask sets CidrMask field to given value.


### GetUi

`func (o *AllowlistEntry) GetUi() bool`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *AllowlistEntry) GetUiOk() (*bool, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *AllowlistEntry) SetUi(v bool)`

SetUi sets Ui field to given value.


### GetSql

`func (o *AllowlistEntry) GetSql() bool`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *AllowlistEntry) GetSqlOk() (*bool, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *AllowlistEntry) SetSql(v bool)`

SetSql sets Sql field to given value.


### GetName

`func (o *AllowlistEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllowlistEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllowlistEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllowlistEntry) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


