# \WebhooksApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksCreate**](WebhooksApi.md#WebhooksCreate) | **Post** /webhooks/{owner}/{repo}/ | Create a specific webhook in a repository.
[**WebhooksDelete**](WebhooksApi.md#WebhooksDelete) | **Delete** /webhooks/{owner}/{repo}/{identifier}/ | Delete a specific webhook in a repository.
[**WebhooksList**](WebhooksApi.md#WebhooksList) | **Get** /webhooks/{owner}/{repo}/ | Get a list of all webhooks in a repository.
[**WebhooksPartialUpdate**](WebhooksApi.md#WebhooksPartialUpdate) | **Patch** /webhooks/{owner}/{repo}/{identifier}/ | Update a specific webhook in a repository.
[**WebhooksRead**](WebhooksApi.md#WebhooksRead) | **Get** /webhooks/{owner}/{repo}/{identifier}/ | Views for working with repository webhooks.



## WebhooksCreate

> RepositoryWebhook WebhooksCreate(ctx, owner, repo).Data(data).Execute()

Create a specific webhook in a repository.



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
	owner := "owner_example" // string | 
	repo := "repo_example" // string | 
	data := *openapiclient.NewRepositoryWebhookRequest([]string{"Events_example"}, "TargetUrl_example", []openapiclient.WebhookTemplate{nil}) // RepositoryWebhookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksApi.WebhooksCreate(context.Background(), owner, repo).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksCreate`: RepositoryWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryWebhookRequest**](RepositoryWebhookRequest.md) |  | 

### Return type

[**RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksDelete

> WebhooksDelete(ctx, owner, repo, identifier).Execute()

Delete a specific webhook in a repository.



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
	owner := "owner_example" // string | 
	repo := "repo_example" // string | 
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksApi.WebhooksDelete(context.Background(), owner, repo, identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksList

> []RepositoryWebhook WebhooksList(ctx, owner, repo).Page(page).PageSize(pageSize).Execute()

Get a list of all webhooks in a repository.



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
	owner := "owner_example" // string | 
	repo := "repo_example" // string | 
	page := int64(56) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(56) // int64 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksApi.WebhooksList(context.Background(), owner, repo).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksList`: []RepositoryWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPartialUpdate

> RepositoryWebhook WebhooksPartialUpdate(ctx, owner, repo, identifier).Data(data).Execute()

Update a specific webhook in a repository.



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
	owner := "owner_example" // string | 
	repo := "repo_example" // string | 
	identifier := "identifier_example" // string | 
	data := *openapiclient.NewRepositoryWebhookRequestPatch() // RepositoryWebhookRequestPatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksApi.WebhooksPartialUpdate(context.Background(), owner, repo, identifier).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksPartialUpdate`: RepositoryWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**RepositoryWebhookRequestPatch**](RepositoryWebhookRequestPatch.md) |  | 

### Return type

[**RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksRead

> RepositoryWebhook WebhooksRead(ctx, owner, repo, identifier).Execute()

Views for working with repository webhooks.



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
	owner := "owner_example" // string | 
	repo := "repo_example" // string | 
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksApi.WebhooksRead(context.Background(), owner, repo, identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksRead`: RepositoryWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

