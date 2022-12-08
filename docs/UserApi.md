# \UserApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserSelf**](UserApi.md#UserSelf) | **Get** /user/self/ | Provide a brief for the current user (if any).
[**UserTokenCreate**](UserApi.md#UserTokenCreate) | **Post** /user/token/ | Retrieve the API key/token for the authenticated user.



## UserSelf

> UserBriefResponse UserSelf(ctx).Execute()

Provide a brief for the current user (if any).



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserSelf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSelf`: UserBriefResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserSelf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSelfRequest struct via the builder pattern


### Return type

[**UserBriefResponse**](UserBriefResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokenCreate

> UserAuthTokenResponse UserTokenCreate(ctx).Data(data).Execute()

Retrieve the API key/token for the authenticated user.



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
    data := *openapiclient.NewUserAuthTokenRequest() // UserAuthTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserTokenCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTokenCreate`: UserAuthTokenResponse
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

[**UserAuthTokenResponse**](UserAuthTokenResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

