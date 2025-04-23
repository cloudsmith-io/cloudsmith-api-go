# \UserApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserSelf**](UserApi.md#UserSelf) | **Get** /user/self/ | Provide a brief for the current user (if any).
[**UserTokenCreate**](UserApi.md#UserTokenCreate) | **Post** /user/token/ | Create or retrieve API token for a user.
[**UserTokensCreate**](UserApi.md#UserTokensCreate) | **Post** /user/tokens/ | Create an API key for the user that is currently authenticated.
[**UserTokensList**](UserApi.md#UserTokensList) | **Get** /user/tokens/ | Retrieve the API key assigned to the user that is currently authenticated.
[**UserTokensRefresh**](UserApi.md#UserTokensRefresh) | **Put** /user/tokens/{slug_perm}/refresh/ | Refresh the specified API key for the user that is currently authenticated.



## UserSelf

> UserBrief UserSelf(ctx).Execute()

Provide a brief for the current user (if any).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.UserSelf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserSelf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSelf`: UserBrief
	fmt.Fprintf(os.Stdout, "Response from `UserApi.UserSelf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSelfRequest struct via the builder pattern


### Return type

[**UserBrief**](UserBrief.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokenCreate

> UserAuthToken UserTokenCreate(ctx).Data(data).Execute()

Create or retrieve API token for a user.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
	data := *openapiclient.NewUserAuthTokenRequest() // UserAuthTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.UserTokenCreate(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokenCreate`: UserAuthToken
	fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**UserAuthTokenRequest**](UserAuthTokenRequest.md) |  | 

### Return type

[**UserAuthToken**](UserAuthToken.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokensCreate

> UserAuthenticationToken UserTokensCreate(ctx).Execute()

Create an API key for the user that is currently authenticated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.UserTokensCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokensCreate`: UserAuthenticationToken
	fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTokensCreate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokensCreateRequest struct via the builder pattern


### Return type

[**UserAuthenticationToken**](UserAuthenticationToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokensList

> UserTokensList200Response UserTokensList(ctx).Page(page).PageSize(pageSize).Execute()

Retrieve the API key assigned to the user that is currently authenticated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
	page := int64(56) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(56) // int64 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.UserTokensList(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTokensList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokensList`: UserTokensList200Response
	fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**UserTokensList200Response**](UserTokensList200Response.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokensRefresh

> UserAuthenticationToken UserTokensRefresh(ctx, slugPerm).Execute()

Refresh the specified API key for the user that is currently authenticated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
	slugPerm := "slugPerm_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.UserTokensRefresh(context.Background(), slugPerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTokensRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokensRefresh`: UserAuthenticationToken
	fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTokensRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokensRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserAuthenticationToken**](UserAuthenticationToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

