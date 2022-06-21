# ListAllowlistEntriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowlist** | [**[]AllowlistEntry**](AllowlistEntry.md) |  | 
**Propagating** | **bool** |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListAllowlistEntriesResponse

`func NewListAllowlistEntriesResponse(allowlist []AllowlistEntry, propagating bool, ) *ListAllowlistEntriesResponse`

NewListAllowlistEntriesResponse instantiates a new ListAllowlistEntriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllowlistEntriesResponseWithDefaults

`func NewListAllowlistEntriesResponseWithDefaults() *ListAllowlistEntriesResponse`

NewListAllowlistEntriesResponseWithDefaults instantiates a new ListAllowlistEntriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowlist

`func (o *ListAllowlistEntriesResponse) GetAllowlist() []AllowlistEntry`

GetAllowlist returns the Allowlist field if non-nil, zero value otherwise.

### GetAllowlistOk

`func (o *ListAllowlistEntriesResponse) GetAllowlistOk() (*[]AllowlistEntry, bool)`

GetAllowlistOk returns a tuple with the Allowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlist

`func (o *ListAllowlistEntriesResponse) SetAllowlist(v []AllowlistEntry)`

SetAllowlist sets Allowlist field to given value.


### GetPropagating

`func (o *ListAllowlistEntriesResponse) GetPropagating() bool`

GetPropagating returns the Propagating field if non-nil, zero value otherwise.

### GetPropagatingOk

`func (o *ListAllowlistEntriesResponse) GetPropagatingOk() (*bool, bool)`

GetPropagatingOk returns a tuple with the Propagating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagating

`func (o *ListAllowlistEntriesResponse) SetPropagating(v bool)`

SetPropagating sets Propagating field to given value.


### GetPagination

`func (o *ListAllowlistEntriesResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAllowlistEntriesResponse) GetPaginationOk() (*KeysetPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAllowlistEntriesResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAllowlistEntriesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


