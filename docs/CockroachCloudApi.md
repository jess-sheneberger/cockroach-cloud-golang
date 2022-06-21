# \CockroachCloudApi

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CockroachCloudAddAllowlistEntry**](CockroachCloudApi.md#CockroachCloudAddAllowlistEntry) | **Post** /api/v1/clusters/{cluster_id}/networking/allowlist | Add a new CIDR address to the IP allowlist.
[**CockroachCloudAddAllowlistEntry2**](CockroachCloudApi.md#CockroachCloudAddAllowlistEntry2) | **Put** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Add a new CIDR address to the IP allowlist.
[**CockroachCloudCreateCluster**](CockroachCloudApi.md#CockroachCloudCreateCluster) | **Post** /api/v1/clusters | Create and initialize a new cluster.
[**CockroachCloudCreateSQLUser**](CockroachCloudApi.md#CockroachCloudCreateSQLUser) | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user.
[**CockroachCloudDeleteAllowlistEntry**](CockroachCloudApi.md#CockroachCloudDeleteAllowlistEntry) | **Delete** /api/v1/clusters/{cluster_id}/networking/allowlist/{cidr_ip}/{cidr_mask} | Delete an IP allowlist entry.
[**CockroachCloudDeleteCluster**](CockroachCloudApi.md#CockroachCloudDeleteCluster) | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data.
[**CockroachCloudDeleteSQLUser**](CockroachCloudApi.md#CockroachCloudDeleteSQLUser) | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user.
[**CockroachCloudEnableCMEK**](CockroachCloudApi.md#CockroachCloudEnableCMEK) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster.
[**CockroachCloudGetCMEKClusterInfo**](CockroachCloudApi.md#CockroachCloudGetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster.
[**CockroachCloudGetCluster**](CockroachCloudApi.md#CockroachCloudGetCluster) | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster.
[**CockroachCloudListAllowlistEntries**](CockroachCloudApi.md#CockroachCloudListAllowlistEntries) | **Get** /api/v1/clusters/{cluster_id}/networking/allowlist | Get the IP allowlist and propagation status for a cluster.
[**CockroachCloudListAvailableRegions**](CockroachCloudApi.md#CockroachCloudListAvailableRegions) | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes.
[**CockroachCloudListClusterNodes**](CockroachCloudApi.md#CockroachCloudListClusterNodes) | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster.
[**CockroachCloudListClusters**](CockroachCloudApi.md#CockroachCloudListClusters) | **Get** /api/v1/clusters | List clusters owned by an organization.
[**CockroachCloudListSQLUsers**](CockroachCloudApi.md#CockroachCloudListSQLUsers) | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster.
[**CockroachCloudUpdateAllowlistEntry**](CockroachCloudApi.md#CockroachCloudUpdateAllowlistEntry) | **Patch** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Update an IP allowlist entry.
[**CockroachCloudUpdateCMEKStatus**](CockroachCloudApi.md#CockroachCloudUpdateCMEKStatus) | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster.
[**CockroachCloudUpdateCluster**](CockroachCloudApi.md#CockroachCloudUpdateCluster) | **Patch** /api/v1/clusters/{cluster_id} | Scale or edit a cluster.
[**CockroachCloudUpdateSQLUserPassword**](CockroachCloudApi.md#CockroachCloudUpdateSQLUserPassword) | **Put** /api/v1/clusters/{cluster_id}/sql-users/{name}/password | Update a SQL user&#39;s password.



## CockroachCloudAddAllowlistEntry

> AllowlistEntry CockroachCloudAddAllowlistEntry(ctx, clusterId).AllowlistEntry(allowlistEntry).Execute()

Add a new CIDR address to the IP allowlist.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudAddAllowlistEntry(context.Background(), clusterId).AllowlistEntry(allowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudAddAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudAddAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudAddAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudAddAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudAddAllowlistEntry2

> AllowlistEntry CockroachCloudAddAllowlistEntry2(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).Execute()

Add a new CIDR address to the IP allowlist.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudAddAllowlistEntry2(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudAddAllowlistEntry2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudAddAllowlistEntry2`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudAddAllowlistEntry2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**entryCidrIp** | **string** |  | 
**entryCidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudAddAllowlistEntry2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudCreateCluster

> Cluster CockroachCloudCreateCluster(ctx).CreateClusterRequest(createClusterRequest).Execute()

Create and initialize a new cluster.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createClusterRequest := *openapiclient.NewCreateClusterRequest("Name_example", openapiclient.api.CloudProvider("CLOUD_PROVIDER_UNSPECIFIED"), *openapiclient.NewCreateClusterSpecification()) // CreateClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudCreateCluster(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudCreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudCreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudCreateSQLUser

> SQLUser CockroachCloudCreateSQLUser(ctx, clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()

Create a new SQL user.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    createSQLUserRequest := *openapiclient.NewCreateSQLUserRequest("Name_example", "Password_example") // CreateSQLUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudCreateSQLUser(context.Background(), clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudCreateSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudCreateSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudCreateSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudCreateSQLUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSQLUserRequest** | [**CreateSQLUserRequest**](CreateSQLUserRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudDeleteAllowlistEntry

> AllowlistEntry CockroachCloudDeleteAllowlistEntry(ctx, clusterId, cidrIp, cidrMask).Execute()

Delete an IP allowlist entry.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    cidrIp := "cidrIp_example" // string | 
    cidrMask := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudDeleteAllowlistEntry(context.Background(), clusterId, cidrIp, cidrMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudDeleteAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudDeleteAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudDeleteAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**cidrIp** | **string** |  | 
**cidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudDeleteAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudDeleteCluster

> Cluster CockroachCloudDeleteCluster(ctx, clusterId).Execute()

Delete a cluster and all of its data.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudDeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudDeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudDeleteCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudDeleteSQLUser

> SQLUser CockroachCloudDeleteSQLUser(ctx, clusterId, name).Execute()

Delete a SQL user.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudDeleteSQLUser(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudDeleteSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudDeleteSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudDeleteSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudDeleteSQLUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudEnableCMEK

> CMEKClusterInfo CockroachCloudEnableCMEK(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

Enable CMEK for a cluster.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    cMEKClusterSpecification := *openapiclient.NewCMEKClusterSpecification([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // CMEKClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudEnableCMEK(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudEnableCMEK``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudEnableCMEK`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudEnableCMEK`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudEnableCMEKRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cMEKClusterSpecification** | [**CMEKClusterSpecification**](CMEKClusterSpecification.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudGetCMEKClusterInfo

> CMEKClusterInfo CockroachCloudGetCMEKClusterInfo(ctx, clusterId).Execute()

Get CMEK-related information for a cluster.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudGetCMEKClusterInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudGetCMEKClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudGetCMEKClusterInfo`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudGetCMEKClusterInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudGetCMEKClusterInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudGetCluster

> Cluster CockroachCloudGetCluster(ctx, clusterId).Execute()

Get extended information about a cluster.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudGetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudGetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudGetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudGetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListAllowlistEntries

> ListAllowlistEntriesResponse CockroachCloudListAllowlistEntries(ctx, clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

Get the IP allowlist and propagation status for a cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListAllowlistEntries(context.Background(), clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListAllowlistEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListAllowlistEntries`: ListAllowlistEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListAllowlistEntriesResponse**](ListAllowlistEntriesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListAvailableRegions

> ListAvailableRegionsResponse CockroachCloudListAvailableRegions(ctx).Provider(provider).Serverless(serverless).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List the regions available for new clusters and nodes.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    provider := "provider_example" // string | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider. (optional) (default to "CLOUD_PROVIDER_UNSPECIFIED")
    serverless := true // bool | Optional filter to only show regions available for serverless clusters. (optional) (default to false)
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListAvailableRegions(context.Background()).Provider(provider).Serverless(serverless).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListAvailableRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListAvailableRegions`: ListAvailableRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListAvailableRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListAvailableRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider. | [default to &quot;CLOUD_PROVIDER_UNSPECIFIED&quot;]
 **serverless** | **bool** | Optional filter to only show regions available for serverless clusters. | [default to false]
 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListAvailableRegionsResponse**](ListAvailableRegionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListClusterNodes

> ListClusterNodesResponse CockroachCloudListClusterNodes(ctx, clusterId).RegionName(regionName).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List nodes for a cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    regionName := "regionName_example" // string | Optional filter to limit response to a single region. (optional)
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListClusterNodes(context.Background(), clusterId).RegionName(regionName).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListClusterNodes`: ListClusterNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionName** | **string** | Optional filter to limit response to a single region. | 
 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListClusterNodesResponse**](ListClusterNodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListClusters

> ListClustersResponse CockroachCloudListClusters(ctx).ShowInactive(showInactive).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List clusters owned by an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    showInactive := true // bool | If `true`, show clusters that have been deleted or failed to initialize. (optional) (default to false)
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListClusters(context.Background()).ShowInactive(showInactive).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListClusters`: ListClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInactive** | **bool** | If &#x60;true&#x60;, show clusters that have been deleted or failed to initialize. | [default to false]
 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListClustersResponse**](ListClustersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListSQLUsers

> ListSQLUsersResponse CockroachCloudListSQLUsers(ctx, clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List SQL users for a cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListSQLUsers(context.Background(), clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListSQLUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListSQLUsers`: ListSQLUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListSQLUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListSQLUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListSQLUsersResponse**](ListSQLUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateAllowlistEntry

> AllowlistEntry CockroachCloudUpdateAllowlistEntry(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).FieldMask(fieldMask).Execute()

Update an IP allowlist entry.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 
    fieldMask := "fieldMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateAllowlistEntry(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).FieldMask(fieldMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**entryCidrIp** | **string** |  | 
**entryCidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 
 **fieldMask** | **string** |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateCMEKStatus

> CMEKClusterInfo CockroachCloudUpdateCMEKStatus(ctx, clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()

Update the CMEK-related status for a cluster.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    updateCMEKStatusRequest := *openapiclient.NewUpdateCMEKStatusRequest(openapiclient.CMEKCustomerAction("UNKNOWN_ACTION")) // UpdateCMEKStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateCMEKStatus(context.Background(), clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateCMEKStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateCMEKStatus`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateCMEKStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateCMEKStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCMEKStatusRequest** | [**UpdateCMEKStatusRequest**](UpdateCMEKStatusRequest.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateCluster

> Cluster CockroachCloudUpdateCluster(ctx, clusterId).UpdateClusterSpecification(updateClusterSpecification).FieldMask(fieldMask).Execute()

Scale or edit a cluster.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    updateClusterSpecification := *openapiclient.NewUpdateClusterSpecification() // UpdateClusterSpecification | 
    fieldMask := "fieldMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateCluster(context.Background(), clusterId).UpdateClusterSpecification(updateClusterSpecification).FieldMask(fieldMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterSpecification** | [**UpdateClusterSpecification**](UpdateClusterSpecification.md) |  | 
 **fieldMask** | **string** |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateSQLUserPassword

> SQLUser CockroachCloudUpdateSQLUserPassword(ctx, clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()

Update a SQL user's password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    name := "name_example" // string | 
    updateSQLUserPasswordRequest := *openapiclient.NewUpdateSQLUserPasswordRequest("Password_example") // UpdateSQLUserPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateSQLUserPassword(context.Background(), clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateSQLUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateSQLUserPassword`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateSQLUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateSQLUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSQLUserPasswordRequest** | [**UpdateSQLUserPasswordRequest**](UpdateSQLUserPasswordRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

