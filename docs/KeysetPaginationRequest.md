# KeysetPaginationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartKey** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**PageDirection**](PageDirection.md) |  | [optional] [default to NEXT]
**Limit** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] [default to ASC]

## Methods

### NewKeysetPaginationRequest

`func NewKeysetPaginationRequest() *KeysetPaginationRequest`

NewKeysetPaginationRequest instantiates a new KeysetPaginationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysetPaginationRequestWithDefaults

`func NewKeysetPaginationRequestWithDefaults() *KeysetPaginationRequest`

NewKeysetPaginationRequestWithDefaults instantiates a new KeysetPaginationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartKey

`func (o *KeysetPaginationRequest) GetStartKey() string`

GetStartKey returns the StartKey field if non-nil, zero value otherwise.

### GetStartKeyOk

`func (o *KeysetPaginationRequest) GetStartKeyOk() (*string, bool)`

GetStartKeyOk returns a tuple with the StartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartKey

`func (o *KeysetPaginationRequest) SetStartKey(v string)`

SetStartKey sets StartKey field to given value.

### HasStartKey

`func (o *KeysetPaginationRequest) HasStartKey() bool`

HasStartKey returns a boolean if a field has been set.

### GetDirection

`func (o *KeysetPaginationRequest) GetDirection() PageDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KeysetPaginationRequest) GetDirectionOk() (*PageDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KeysetPaginationRequest) SetDirection(v PageDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *KeysetPaginationRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetLimit

`func (o *KeysetPaginationRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KeysetPaginationRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KeysetPaginationRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KeysetPaginationRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTime

`func (o *KeysetPaginationRequest) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *KeysetPaginationRequest) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *KeysetPaginationRequest) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *KeysetPaginationRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOrder

`func (o *KeysetPaginationRequest) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *KeysetPaginationRequest) GetOrderOk() (*SortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *KeysetPaginationRequest) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *KeysetPaginationRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


