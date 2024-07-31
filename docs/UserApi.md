# \UserApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserEmailsPartialUpdate**](UserApi.md#UserEmailsPartialUpdate) | **Patch** /user/emails/{identifier}/ | 
[**UserPassword**](UserApi.md#UserPassword) | **Put** /user/password/ | Views for the current user.
[**UserSelf**](UserApi.md#UserSelf) | **Get** /user/self/ | Provide a brief for the current user (if any).
[**UserTokenCreate**](UserApi.md#UserTokenCreate) | **Post** /user/token/ | Retrieve/Create API key/token for the authenticated user.



## UserEmailsPartialUpdate

> EmailAddress UserEmailsPartialUpdate(ctx, identifier).Data(data).Execute()



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewEmailAddressRequestPatch() // EmailAddressRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserEmailsPartialUpdate(context.Background(), identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserEmailsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserEmailsPartialUpdate`: EmailAddress
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserEmailsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEmailsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**EmailAddressRequestPatch**](EmailAddressRequestPatch.md) |  | 

### Return type

[**EmailAddress**](EmailAddress.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPassword

> UserPassword(ctx).Data(data).Execute()

Views for the current user.



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
    data := *openapiclient.NewUserPasswordRequest("Password_example") // UserPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserApi.UserPassword(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**UserPasswordRequest**](UserPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

Retrieve/Create API key/token for the authenticated user.



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

