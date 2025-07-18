# \FilesApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesAbort**](FilesApi.md#FilesAbort) | **Post** /files/{owner}/{repo}/{identifier}/abort/ | Abort a multipart file upload.
[**FilesComplete**](FilesApi.md#FilesComplete) | **Post** /files/{owner}/{repo}/{identifier}/complete/ | Complete a multipart file upload.
[**FilesCreate**](FilesApi.md#FilesCreate) | **Post** /files/{owner}/{repo}/ | Request URL(s) to upload new package file upload(s) to.
[**FilesInfo**](FilesApi.md#FilesInfo) | **Get** /files/{owner}/{repo}/{identifier}/info/ | Get upload information to perform a multipart file upload.
[**FilesValidate**](FilesApi.md#FilesValidate) | **Post** /files/{owner}/{repo}/validate/ | Validate parameters used for create.



## FilesAbort

> FilesAbort(ctx, owner, repo, identifier).Data(data).Execute()

Abort a multipart file upload.



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
	data := *openapiclient.NewPackageFileUploadRequest("Filename_example") // PackageFileUploadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesApi.FilesAbort(context.Background(), owner, repo, identifier).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesAbort``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFilesAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PackageFileUploadRequest**](PackageFileUploadRequest.md) |  | 

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


## FilesComplete

> PackageFileUpload FilesComplete(ctx, owner, repo, identifier).Data(data).Execute()

Complete a multipart file upload.



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
	data := *openapiclient.NewPackageFileUploadRequest("Filename_example") // PackageFileUploadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesApi.FilesComplete(context.Background(), owner, repo, identifier).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesComplete`: PackageFileUpload
	fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesComplete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiFilesCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PackageFileUploadRequest**](PackageFileUploadRequest.md) |  | 

### Return type

[**PackageFileUpload**](PackageFileUpload.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesCreate

> PackageFileUpload FilesCreate(ctx, owner, repo).Data(data).Execute()

Request URL(s) to upload new package file upload(s) to.



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
	data := *openapiclient.NewPackageFileUploadRequest("Filename_example") // PackageFileUploadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesApi.FilesCreate(context.Background(), owner, repo).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesCreate`: PackageFileUpload
	fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PackageFileUploadRequest**](PackageFileUploadRequest.md) |  | 

### Return type

[**PackageFileUpload**](PackageFileUpload.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesInfo

> PackageFilePartsUpload FilesInfo(ctx, owner, repo, identifier).Filename(filename).PartNumber(partNumber).Execute()

Get upload information to perform a multipart file upload.



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
	filename := "filename_example" // string | The filename of the file being uploaded
	partNumber := int64(56) // int64 | The part number to be uploaded next (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesApi.FilesInfo(context.Background(), owner, repo, identifier).Filename(filename).PartNumber(partNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesInfo`: PackageFilePartsUpload
	fmt.Fprintf(os.Stdout, "Response from `FilesApi.FilesInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiFilesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filename** | **string** | The filename of the file being uploaded | 
 **partNumber** | **int64** | The part number to be uploaded next | 

### Return type

[**PackageFilePartsUpload**](PackageFilePartsUpload.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesValidate

> FilesValidate(ctx, owner, repo).Data(data).Execute()

Validate parameters used for create.



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
	data := *openapiclient.NewPackageFileUploadRequest("Filename_example") // PackageFileUploadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesApi.FilesValidate(context.Background(), owner, repo).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FilesValidate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiFilesValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PackageFileUploadRequest**](PackageFileUploadRequest.md) |  | 

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

