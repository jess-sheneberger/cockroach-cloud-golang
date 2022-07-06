# KeysetPaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** |  | [optional] 
**Last** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] [default to ASC]

## Methods

### NewKeysetPaginationResponse

`func NewKeysetPaginationResponse() *KeysetPaginationResponse`

NewKeysetPaginationResponse instantiates a new KeysetPaginationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysetPaginationResponseWithDefaults

`func NewKeysetPaginationResponseWithDefaults() *KeysetPaginationResponse`

NewKeysetPaginationResponseWithDefaults instantiates a new KeysetPaginationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *KeysetPaginationResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *KeysetPaginationResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *KeysetPaginationResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *KeysetPaginationResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetLast

`func (o *KeysetPaginationResponse) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *KeysetPaginationResponse) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *KeysetPaginationResponse) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *KeysetPaginationResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLimit

`func (o *KeysetPaginationResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KeysetPaginationResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KeysetPaginationResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KeysetPaginationResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTime

`func (o *KeysetPaginationResponse) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *KeysetPaginationResponse) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *KeysetPaginationResponse) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *KeysetPaginationResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOrder

`func (o *KeysetPaginationResponse) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *KeysetPaginationResponse) GetOrderOk() (*SortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *KeysetPaginationResponse) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *KeysetPaginationResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


