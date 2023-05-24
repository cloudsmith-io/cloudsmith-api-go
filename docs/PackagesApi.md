# \PackagesApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PackagesCopy**](PackagesApi.md#PackagesCopy) | **Post** /packages/{owner}/{repo}/{identifier}/copy/ | Copy a package to another repository.
[**PackagesDelete**](PackagesApi.md#PackagesDelete) | **Delete** /packages/{owner}/{repo}/{identifier}/ | Delete a specific package in a repository.
[**PackagesDependencies**](PackagesApi.md#PackagesDependencies) | **Get** /packages/{owner}/{repo}/{identifier}/dependencies/ | Get the list of dependencies for a package. Transitive dependencies are included where supported.
[**PackagesList**](PackagesApi.md#PackagesList) | **Get** /packages/{owner}/{repo}/ | Get a list of all packages associated with repository.
[**PackagesMove**](PackagesApi.md#PackagesMove) | **Post** /packages/{owner}/{repo}/{identifier}/move/ | Move a package to another repository.
[**PackagesQuarantine**](PackagesApi.md#PackagesQuarantine) | **Post** /packages/{owner}/{repo}/{identifier}/quarantine/ | Quarantine or restore a package.
[**PackagesRead**](PackagesApi.md#PackagesRead) | **Get** /packages/{owner}/{repo}/{identifier}/ | Get a specific package in a repository.
[**PackagesResync**](PackagesApi.md#PackagesResync) | **Post** /packages/{owner}/{repo}/{identifier}/resync/ | Schedule a package for resynchronisation.
[**PackagesScan**](PackagesApi.md#PackagesScan) | **Post** /packages/{owner}/{repo}/{identifier}/scan/ | Schedule a package for scanning.
[**PackagesStatus**](PackagesApi.md#PackagesStatus) | **Get** /packages/{owner}/{repo}/{identifier}/status/ | Get the synchronisation status for a package.
[**PackagesTag**](PackagesApi.md#PackagesTag) | **Post** /packages/{owner}/{repo}/{identifier}/tag/ | Add/Replace/Remove tags for a package.
[**PackagesUploadAlpine**](PackagesApi.md#PackagesUploadAlpine) | **Post** /packages/{owner}/{repo}/upload/alpine/ | Create a new Alpine package
[**PackagesUploadCargo**](PackagesApi.md#PackagesUploadCargo) | **Post** /packages/{owner}/{repo}/upload/cargo/ | Create a new Cargo package
[**PackagesUploadCocoapods**](PackagesApi.md#PackagesUploadCocoapods) | **Post** /packages/{owner}/{repo}/upload/cocoapods/ | Create a new CocoaPods package
[**PackagesUploadComposer**](PackagesApi.md#PackagesUploadComposer) | **Post** /packages/{owner}/{repo}/upload/composer/ | Create a new Composer package
[**PackagesUploadConan**](PackagesApi.md#PackagesUploadConan) | **Post** /packages/{owner}/{repo}/upload/conan/ | Create a new Conan package
[**PackagesUploadConda**](PackagesApi.md#PackagesUploadConda) | **Post** /packages/{owner}/{repo}/upload/conda/ | Create a new Conda package
[**PackagesUploadCran**](PackagesApi.md#PackagesUploadCran) | **Post** /packages/{owner}/{repo}/upload/cran/ | Create a new CRAN package
[**PackagesUploadDart**](PackagesApi.md#PackagesUploadDart) | **Post** /packages/{owner}/{repo}/upload/dart/ | Create a new Dart package
[**PackagesUploadDeb**](PackagesApi.md#PackagesUploadDeb) | **Post** /packages/{owner}/{repo}/upload/deb/ | Create a new Debian package
[**PackagesUploadDocker**](PackagesApi.md#PackagesUploadDocker) | **Post** /packages/{owner}/{repo}/upload/docker/ | Create a new Docker package
[**PackagesUploadGo**](PackagesApi.md#PackagesUploadGo) | **Post** /packages/{owner}/{repo}/upload/go/ | Create a new Go package
[**PackagesUploadHelm**](PackagesApi.md#PackagesUploadHelm) | **Post** /packages/{owner}/{repo}/upload/helm/ | Create a new Helm package
[**PackagesUploadHex**](PackagesApi.md#PackagesUploadHex) | **Post** /packages/{owner}/{repo}/upload/hex/ | Create a new Hex package
[**PackagesUploadLuarocks**](PackagesApi.md#PackagesUploadLuarocks) | **Post** /packages/{owner}/{repo}/upload/luarocks/ | Create a new LuaRocks package
[**PackagesUploadMaven**](PackagesApi.md#PackagesUploadMaven) | **Post** /packages/{owner}/{repo}/upload/maven/ | Create a new Maven package
[**PackagesUploadNpm**](PackagesApi.md#PackagesUploadNpm) | **Post** /packages/{owner}/{repo}/upload/npm/ | Create a new npm package
[**PackagesUploadNuget**](PackagesApi.md#PackagesUploadNuget) | **Post** /packages/{owner}/{repo}/upload/nuget/ | Create a new NuGet package
[**PackagesUploadP2**](PackagesApi.md#PackagesUploadP2) | **Post** /packages/{owner}/{repo}/upload/p2/ | Create a new P2 package
[**PackagesUploadPython**](PackagesApi.md#PackagesUploadPython) | **Post** /packages/{owner}/{repo}/upload/python/ | Create a new Python package
[**PackagesUploadRaw**](PackagesApi.md#PackagesUploadRaw) | **Post** /packages/{owner}/{repo}/upload/raw/ | Create a new Raw package
[**PackagesUploadRpm**](PackagesApi.md#PackagesUploadRpm) | **Post** /packages/{owner}/{repo}/upload/rpm/ | Create a new RedHat package
[**PackagesUploadRuby**](PackagesApi.md#PackagesUploadRuby) | **Post** /packages/{owner}/{repo}/upload/ruby/ | Create a new Ruby package
[**PackagesUploadSwift**](PackagesApi.md#PackagesUploadSwift) | **Post** /packages/{owner}/{repo}/upload/swift/ | Create a new Swift package
[**PackagesUploadTerraform**](PackagesApi.md#PackagesUploadTerraform) | **Post** /packages/{owner}/{repo}/upload/terraform/ | Create a new Terraform package
[**PackagesUploadVagrant**](PackagesApi.md#PackagesUploadVagrant) | **Post** /packages/{owner}/{repo}/upload/vagrant/ | Create a new Vagrant package
[**PackagesValidateUploadAlpine**](PackagesApi.md#PackagesValidateUploadAlpine) | **Post** /packages/{owner}/{repo}/validate-upload/alpine/ | Validate parameters for create Alpine package
[**PackagesValidateUploadCargo**](PackagesApi.md#PackagesValidateUploadCargo) | **Post** /packages/{owner}/{repo}/validate-upload/cargo/ | Validate parameters for create Cargo package
[**PackagesValidateUploadCocoapods**](PackagesApi.md#PackagesValidateUploadCocoapods) | **Post** /packages/{owner}/{repo}/validate-upload/cocoapods/ | Validate parameters for create CocoaPods package
[**PackagesValidateUploadComposer**](PackagesApi.md#PackagesValidateUploadComposer) | **Post** /packages/{owner}/{repo}/validate-upload/composer/ | Validate parameters for create Composer package
[**PackagesValidateUploadConan**](PackagesApi.md#PackagesValidateUploadConan) | **Post** /packages/{owner}/{repo}/validate-upload/conan/ | Validate parameters for create Conan package
[**PackagesValidateUploadConda**](PackagesApi.md#PackagesValidateUploadConda) | **Post** /packages/{owner}/{repo}/validate-upload/conda/ | Validate parameters for create Conda package
[**PackagesValidateUploadCran**](PackagesApi.md#PackagesValidateUploadCran) | **Post** /packages/{owner}/{repo}/validate-upload/cran/ | Validate parameters for create CRAN package
[**PackagesValidateUploadDart**](PackagesApi.md#PackagesValidateUploadDart) | **Post** /packages/{owner}/{repo}/validate-upload/dart/ | Validate parameters for create Dart package
[**PackagesValidateUploadDeb**](PackagesApi.md#PackagesValidateUploadDeb) | **Post** /packages/{owner}/{repo}/validate-upload/deb/ | Validate parameters for create Debian package
[**PackagesValidateUploadDocker**](PackagesApi.md#PackagesValidateUploadDocker) | **Post** /packages/{owner}/{repo}/validate-upload/docker/ | Validate parameters for create Docker package
[**PackagesValidateUploadGo**](PackagesApi.md#PackagesValidateUploadGo) | **Post** /packages/{owner}/{repo}/validate-upload/go/ | Validate parameters for create Go package
[**PackagesValidateUploadHelm**](PackagesApi.md#PackagesValidateUploadHelm) | **Post** /packages/{owner}/{repo}/validate-upload/helm/ | Validate parameters for create Helm package
[**PackagesValidateUploadHex**](PackagesApi.md#PackagesValidateUploadHex) | **Post** /packages/{owner}/{repo}/validate-upload/hex/ | Validate parameters for create Hex package
[**PackagesValidateUploadLuarocks**](PackagesApi.md#PackagesValidateUploadLuarocks) | **Post** /packages/{owner}/{repo}/validate-upload/luarocks/ | Validate parameters for create LuaRocks package
[**PackagesValidateUploadMaven**](PackagesApi.md#PackagesValidateUploadMaven) | **Post** /packages/{owner}/{repo}/validate-upload/maven/ | Validate parameters for create Maven package
[**PackagesValidateUploadNpm**](PackagesApi.md#PackagesValidateUploadNpm) | **Post** /packages/{owner}/{repo}/validate-upload/npm/ | Validate parameters for create npm package
[**PackagesValidateUploadNuget**](PackagesApi.md#PackagesValidateUploadNuget) | **Post** /packages/{owner}/{repo}/validate-upload/nuget/ | Validate parameters for create NuGet package
[**PackagesValidateUploadP2**](PackagesApi.md#PackagesValidateUploadP2) | **Post** /packages/{owner}/{repo}/validate-upload/p2/ | Validate parameters for create P2 package
[**PackagesValidateUploadPython**](PackagesApi.md#PackagesValidateUploadPython) | **Post** /packages/{owner}/{repo}/validate-upload/python/ | Validate parameters for create Python package
[**PackagesValidateUploadRaw**](PackagesApi.md#PackagesValidateUploadRaw) | **Post** /packages/{owner}/{repo}/validate-upload/raw/ | Validate parameters for create Raw package
[**PackagesValidateUploadRpm**](PackagesApi.md#PackagesValidateUploadRpm) | **Post** /packages/{owner}/{repo}/validate-upload/rpm/ | Validate parameters for create RedHat package
[**PackagesValidateUploadRuby**](PackagesApi.md#PackagesValidateUploadRuby) | **Post** /packages/{owner}/{repo}/validate-upload/ruby/ | Validate parameters for create Ruby package
[**PackagesValidateUploadSwift**](PackagesApi.md#PackagesValidateUploadSwift) | **Post** /packages/{owner}/{repo}/validate-upload/swift/ | Validate parameters for create Swift package
[**PackagesValidateUploadTerraform**](PackagesApi.md#PackagesValidateUploadTerraform) | **Post** /packages/{owner}/{repo}/validate-upload/terraform/ | Validate parameters for create Terraform package
[**PackagesValidateUploadVagrant**](PackagesApi.md#PackagesValidateUploadVagrant) | **Post** /packages/{owner}/{repo}/validate-upload/vagrant/ | Validate parameters for create Vagrant package



## PackagesCopy

> PackageCopy PackagesCopy(ctx, owner, repo, identifier).Data(data).Execute()

Copy a package to another repository.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewPackageCopyRequest("Destination_example") // PackageCopyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesCopy(context.Background(), owner, repo, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesCopy`: PackageCopy
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesCopy`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PackageCopyRequest**](PackageCopyRequest.md) |  | 

### Return type

[**PackageCopy**](PackageCopy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesDelete

> PackagesDelete(ctx, owner, repo, identifier).Execute()

Delete a specific package in a repository.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDelete(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesDependencies

> PackageDependencies PackagesDependencies(ctx, owner, repo, identifier).Execute()

Get the list of dependencies for a package. Transitive dependencies are included where supported.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDependencies(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDependencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesDependencies`: PackageDependencies
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesDependencies`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageDependencies**](PackageDependencies.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesList

> []Package PackagesList(ctx, owner, repo).Page(page).PageSize(pageSize).Query(query).Sort(sort).Execute()

Get a list of all packages associated with repository.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)
    query := "query_example" // string | A search term for querying names, filenames, versions, distributions, architectures, formats or statuses of packages. (optional)
    sort := "sort_example" // string | A field for sorting objects in ascending or descending order. (optional) (default to "-date")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesList(context.Background(), owner, repo).Page(page).PageSize(pageSize).Query(query).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesList`: []Package
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **query** | **string** | A search term for querying names, filenames, versions, distributions, architectures, formats or statuses of packages. | 
 **sort** | **string** | A field for sorting objects in ascending or descending order. | [default to &quot;-date&quot;]

### Return type

[**[]Package**](Package.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesMove

> PackageMove PackagesMove(ctx, owner, repo, identifier).Data(data).Execute()

Move a package to another repository.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewPackageMoveRequest("Destination_example") // PackageMoveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesMove(context.Background(), owner, repo, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesMove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesMove`: PackageMove
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesMove`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PackageMoveRequest**](PackageMoveRequest.md) |  | 

### Return type

[**PackageMove**](PackageMove.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesQuarantine

> PackageQuarantine PackagesQuarantine(ctx, owner, repo, identifier).Data(data).Execute()

Quarantine or restore a package.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewPackageQuarantineRequest() // PackageQuarantineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesQuarantine(context.Background(), owner, repo, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesQuarantine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesQuarantine`: PackageQuarantine
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesQuarantine`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesQuarantineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PackageQuarantineRequest**](PackageQuarantineRequest.md) |  | 

### Return type

[**PackageQuarantine**](PackageQuarantine.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesRead

> Package PackagesRead(ctx, owner, repo, identifier).Execute()

Get a specific package in a repository.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRead(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesRead`: Package
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesResync

> PackageResync PackagesResync(ctx, owner, repo, identifier).Execute()

Schedule a package for resynchronisation.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesResync(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesResync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesResync`: PackageResync
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesResync`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesResyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageResync**](PackageResync.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesScan

> Package PackagesScan(ctx, owner, repo, identifier).Execute()

Schedule a package for scanning.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesScan(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesScan`: Package
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesScan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesStatus

> PackageStatus PackagesStatus(ctx, owner, repo, identifier).Execute()

Get the synchronisation status for a package.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesStatus(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesStatus`: PackageStatus
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageStatus**](PackageStatus.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesTag

> Package PackagesTag(ctx, owner, repo, identifier).Data(data).Execute()

Add/Replace/Remove tags for a package.



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewPackageTagRequest() // PackageTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesTag(context.Background(), owner, repo, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesTag`: Package
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesTag`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPackagesTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PackageTagRequest**](PackageTagRequest.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadAlpine

> AlpinePackageUpload PackagesUploadAlpine(ctx, owner, repo).Data(data).Execute()

Create a new Alpine package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewAlpinePackageUploadRequest("Distribution_example", "PackageFile_example") // AlpinePackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadAlpine(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadAlpine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadAlpine`: AlpinePackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadAlpine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadAlpineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**AlpinePackageUploadRequest**](AlpinePackageUploadRequest.md) |  | 

### Return type

[**AlpinePackageUpload**](AlpinePackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadCargo

> CargoPackageUpload PackagesUploadCargo(ctx, owner, repo).Data(data).Execute()

Create a new Cargo package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCargoPackageUploadRequest("PackageFile_example") // CargoPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadCargo(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadCargo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadCargo`: CargoPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadCargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CargoPackageUploadRequest**](CargoPackageUploadRequest.md) |  | 

### Return type

[**CargoPackageUpload**](CargoPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadCocoapods

> CocoapodsPackageUpload PackagesUploadCocoapods(ctx, owner, repo).Data(data).Execute()

Create a new CocoaPods package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCocoapodsPackageUploadRequest("PackageFile_example") // CocoapodsPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadCocoapods(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadCocoapods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadCocoapods`: CocoapodsPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadCocoapods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadCocoapodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CocoapodsPackageUploadRequest**](CocoapodsPackageUploadRequest.md) |  | 

### Return type

[**CocoapodsPackageUpload**](CocoapodsPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadComposer

> ComposerPackageUpload PackagesUploadComposer(ctx, owner, repo).Data(data).Execute()

Create a new Composer package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewComposerPackageUploadRequest("PackageFile_example") // ComposerPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadComposer(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadComposer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadComposer`: ComposerPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadComposer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadComposerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ComposerPackageUploadRequest**](ComposerPackageUploadRequest.md) |  | 

### Return type

[**ComposerPackageUpload**](ComposerPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadConan

> ConanPackageUpload PackagesUploadConan(ctx, owner, repo).Data(data).Execute()

Create a new Conan package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewConanPackageUploadRequest("InfoFile_example", "ManifestFile_example", "MetadataFile_example", "PackageFile_example") // ConanPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadConan(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadConan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadConan`: ConanPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadConan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadConanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ConanPackageUploadRequest**](ConanPackageUploadRequest.md) |  | 

### Return type

[**ConanPackageUpload**](ConanPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadConda

> CondaPackageUpload PackagesUploadConda(ctx, owner, repo).Data(data).Execute()

Create a new Conda package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCondaPackageUploadRequest("PackageFile_example") // CondaPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadConda(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadConda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadConda`: CondaPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadConda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadCondaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CondaPackageUploadRequest**](CondaPackageUploadRequest.md) |  | 

### Return type

[**CondaPackageUpload**](CondaPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadCran

> CranPackageUpload PackagesUploadCran(ctx, owner, repo).Data(data).Execute()

Create a new CRAN package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCranPackageUploadRequest("PackageFile_example") // CranPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadCran(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadCran``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadCran`: CranPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadCran`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadCranRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CranPackageUploadRequest**](CranPackageUploadRequest.md) |  | 

### Return type

[**CranPackageUpload**](CranPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadDart

> DartPackageUpload PackagesUploadDart(ctx, owner, repo).Data(data).Execute()

Create a new Dart package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewDartPackageUploadRequest("PackageFile_example") // DartPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadDart(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadDart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadDart`: DartPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadDart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadDartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DartPackageUploadRequest**](DartPackageUploadRequest.md) |  | 

### Return type

[**DartPackageUpload**](DartPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadDeb

> DebPackageUpload PackagesUploadDeb(ctx, owner, repo).Data(data).Execute()

Create a new Debian package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewDebPackageUploadRequest("Distribution_example", "PackageFile_example") // DebPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadDeb(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadDeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadDeb`: DebPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadDeb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadDebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DebPackageUploadRequest**](DebPackageUploadRequest.md) |  | 

### Return type

[**DebPackageUpload**](DebPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadDocker

> DockerPackageUpload PackagesUploadDocker(ctx, owner, repo).Data(data).Execute()

Create a new Docker package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewDockerPackageUploadRequest("PackageFile_example") // DockerPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadDocker(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadDocker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadDocker`: DockerPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadDocker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadDockerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DockerPackageUploadRequest**](DockerPackageUploadRequest.md) |  | 

### Return type

[**DockerPackageUpload**](DockerPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadGo

> GoPackageUpload PackagesUploadGo(ctx, owner, repo).Data(data).Execute()

Create a new Go package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewGoPackageUploadRequest("PackageFile_example") // GoPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadGo(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadGo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadGo`: GoPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadGo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadGoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GoPackageUploadRequest**](GoPackageUploadRequest.md) |  | 

### Return type

[**GoPackageUpload**](GoPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadHelm

> HelmPackageUpload PackagesUploadHelm(ctx, owner, repo).Data(data).Execute()

Create a new Helm package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewHelmPackageUploadRequest("PackageFile_example") // HelmPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadHelm(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadHelm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadHelm`: HelmPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**HelmPackageUploadRequest**](HelmPackageUploadRequest.md) |  | 

### Return type

[**HelmPackageUpload**](HelmPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadHex

> HexPackageUpload PackagesUploadHex(ctx, owner, repo).Data(data).Execute()

Create a new Hex package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewHexPackageUploadRequest("PackageFile_example") // HexPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadHex(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadHex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadHex`: HexPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadHex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadHexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**HexPackageUploadRequest**](HexPackageUploadRequest.md) |  | 

### Return type

[**HexPackageUpload**](HexPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadLuarocks

> LuarocksPackageUpload PackagesUploadLuarocks(ctx, owner, repo).Data(data).Execute()

Create a new LuaRocks package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewLuarocksPackageUploadRequest("PackageFile_example") // LuarocksPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadLuarocks(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadLuarocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadLuarocks`: LuarocksPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadLuarocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadLuarocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**LuarocksPackageUploadRequest**](LuarocksPackageUploadRequest.md) |  | 

### Return type

[**LuarocksPackageUpload**](LuarocksPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadMaven

> MavenPackageUpload PackagesUploadMaven(ctx, owner, repo).Data(data).Execute()

Create a new Maven package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewMavenPackageUploadRequest("PackageFile_example") // MavenPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadMaven(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadMaven``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadMaven`: MavenPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadMaven`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadMavenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**MavenPackageUploadRequest**](MavenPackageUploadRequest.md) |  | 

### Return type

[**MavenPackageUpload**](MavenPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadNpm

> NpmPackageUpload PackagesUploadNpm(ctx, owner, repo).Data(data).Execute()

Create a new npm package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewNpmPackageUploadRequest("PackageFile_example") // NpmPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadNpm(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadNpm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadNpm`: NpmPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadNpm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadNpmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**NpmPackageUploadRequest**](NpmPackageUploadRequest.md) |  | 

### Return type

[**NpmPackageUpload**](NpmPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadNuget

> NugetPackageUpload PackagesUploadNuget(ctx, owner, repo).Data(data).Execute()

Create a new NuGet package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewNugetPackageUploadRequest("PackageFile_example") // NugetPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadNuget(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadNuget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadNuget`: NugetPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadNuget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadNugetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**NugetPackageUploadRequest**](NugetPackageUploadRequest.md) |  | 

### Return type

[**NugetPackageUpload**](NugetPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadP2

> P2PackageUpload PackagesUploadP2(ctx, owner, repo).Data(data).Execute()

Create a new P2 package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewP2PackageUploadRequest("PackageFile_example") // P2PackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadP2(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadP2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadP2`: P2PackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadP2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadP2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**P2PackageUploadRequest**](P2PackageUploadRequest.md) |  | 

### Return type

[**P2PackageUpload**](P2PackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadPython

> PythonPackageUpload PackagesUploadPython(ctx, owner, repo).Data(data).Execute()

Create a new Python package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewPythonPackageUploadRequest("PackageFile_example") // PythonPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadPython(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadPython``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadPython`: PythonPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadPython`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadPythonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PythonPackageUploadRequest**](PythonPackageUploadRequest.md) |  | 

### Return type

[**PythonPackageUpload**](PythonPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadRaw

> RawPackageUpload PackagesUploadRaw(ctx, owner, repo).Data(data).Execute()

Create a new Raw package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewRawPackageUploadRequest("PackageFile_example") // RawPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadRaw(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadRaw`: RawPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadRaw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RawPackageUploadRequest**](RawPackageUploadRequest.md) |  | 

### Return type

[**RawPackageUpload**](RawPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadRpm

> RpmPackageUpload PackagesUploadRpm(ctx, owner, repo).Data(data).Execute()

Create a new RedHat package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewRpmPackageUploadRequest("Distribution_example", "PackageFile_example") // RpmPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadRpm(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadRpm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadRpm`: RpmPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadRpm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadRpmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RpmPackageUploadRequest**](RpmPackageUploadRequest.md) |  | 

### Return type

[**RpmPackageUpload**](RpmPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadRuby

> RubyPackageUpload PackagesUploadRuby(ctx, owner, repo).Data(data).Execute()

Create a new Ruby package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewRubyPackageUploadRequest("PackageFile_example") // RubyPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadRuby(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadRuby``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadRuby`: RubyPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadRuby`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadRubyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RubyPackageUploadRequest**](RubyPackageUploadRequest.md) |  | 

### Return type

[**RubyPackageUpload**](RubyPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadSwift

> SwiftPackageUpload PackagesUploadSwift(ctx, owner, repo).Data(data).Execute()

Create a new Swift package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewSwiftPackageUploadRequest("PackageFile_example", "Version_example") // SwiftPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadSwift(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadSwift``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadSwift`: SwiftPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadSwift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadSwiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**SwiftPackageUploadRequest**](SwiftPackageUploadRequest.md) |  | 

### Return type

[**SwiftPackageUpload**](SwiftPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadTerraform

> TerraformPackageUpload PackagesUploadTerraform(ctx, owner, repo).Data(data).Execute()

Create a new Terraform package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewTerraformPackageUploadRequest("PackageFile_example") // TerraformPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadTerraform(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadTerraform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadTerraform`: TerraformPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**TerraformPackageUploadRequest**](TerraformPackageUploadRequest.md) |  | 

### Return type

[**TerraformPackageUpload**](TerraformPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadVagrant

> VagrantPackageUpload PackagesUploadVagrant(ctx, owner, repo).Data(data).Execute()

Create a new Vagrant package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewVagrantPackageUploadRequest("Name_example", "PackageFile_example", "Provider_example", "Version_example") // VagrantPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesUploadVagrant(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesUploadVagrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesUploadVagrant`: VagrantPackageUpload
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesUploadVagrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesUploadVagrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**VagrantPackageUploadRequest**](VagrantPackageUploadRequest.md) |  | 

### Return type

[**VagrantPackageUpload**](VagrantPackageUpload.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadAlpine

> PackagesValidateUploadAlpine(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Alpine package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewAlpinePackageUploadRequest("Distribution_example", "PackageFile_example") // AlpinePackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadAlpine(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadAlpine``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadAlpineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**AlpinePackageUploadRequest**](AlpinePackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadCargo

> PackagesValidateUploadCargo(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Cargo package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCargoPackageUploadRequest("PackageFile_example") // CargoPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadCargo(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadCargo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadCargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CargoPackageUploadRequest**](CargoPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadCocoapods

> PackagesValidateUploadCocoapods(ctx, owner, repo).Data(data).Execute()

Validate parameters for create CocoaPods package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCocoapodsPackageUploadRequest("PackageFile_example") // CocoapodsPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadCocoapods(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadCocoapods``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadCocoapodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CocoapodsPackageUploadRequest**](CocoapodsPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadComposer

> PackagesValidateUploadComposer(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Composer package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewComposerPackageUploadRequest("PackageFile_example") // ComposerPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadComposer(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadComposer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadComposerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ComposerPackageUploadRequest**](ComposerPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadConan

> PackagesValidateUploadConan(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Conan package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewConanPackageUploadRequest("InfoFile_example", "ManifestFile_example", "MetadataFile_example", "PackageFile_example") // ConanPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadConan(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadConan``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadConanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ConanPackageUploadRequest**](ConanPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadConda

> PackagesValidateUploadConda(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Conda package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCondaPackageUploadRequest("PackageFile_example") // CondaPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadConda(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadConda``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadCondaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CondaPackageUploadRequest**](CondaPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadCran

> PackagesValidateUploadCran(ctx, owner, repo).Data(data).Execute()

Validate parameters for create CRAN package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewCranPackageUploadRequest("PackageFile_example") // CranPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadCran(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadCran``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadCranRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CranPackageUploadRequest**](CranPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadDart

> PackagesValidateUploadDart(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Dart package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewDartPackageUploadRequest("PackageFile_example") // DartPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadDart(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadDart``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadDartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DartPackageUploadRequest**](DartPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadDeb

> PackagesValidateUploadDeb(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Debian package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewDebPackageUploadRequest("Distribution_example", "PackageFile_example") // DebPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadDeb(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadDeb``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadDebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DebPackageUploadRequest**](DebPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadDocker

> PackagesValidateUploadDocker(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Docker package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewDockerPackageUploadRequest("PackageFile_example") // DockerPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadDocker(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadDocker``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadDockerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DockerPackageUploadRequest**](DockerPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadGo

> PackagesValidateUploadGo(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Go package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewGoPackageUploadRequest("PackageFile_example") // GoPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadGo(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadGo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadGoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GoPackageUploadRequest**](GoPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadHelm

> PackagesValidateUploadHelm(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Helm package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewHelmPackageUploadRequest("PackageFile_example") // HelmPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadHelm(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadHelm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**HelmPackageUploadRequest**](HelmPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadHex

> PackagesValidateUploadHex(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Hex package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewHexPackageUploadRequest("PackageFile_example") // HexPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadHex(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadHex``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadHexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**HexPackageUploadRequest**](HexPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadLuarocks

> PackagesValidateUploadLuarocks(ctx, owner, repo).Data(data).Execute()

Validate parameters for create LuaRocks package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewLuarocksPackageUploadRequest("PackageFile_example") // LuarocksPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadLuarocks(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadLuarocks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadLuarocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**LuarocksPackageUploadRequest**](LuarocksPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadMaven

> PackagesValidateUploadMaven(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Maven package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewMavenPackageUploadRequest("PackageFile_example") // MavenPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadMaven(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadMaven``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadMavenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**MavenPackageUploadRequest**](MavenPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadNpm

> PackagesValidateUploadNpm(ctx, owner, repo).Data(data).Execute()

Validate parameters for create npm package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewNpmPackageUploadRequest("PackageFile_example") // NpmPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadNpm(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadNpm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadNpmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**NpmPackageUploadRequest**](NpmPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadNuget

> PackagesValidateUploadNuget(ctx, owner, repo).Data(data).Execute()

Validate parameters for create NuGet package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewNugetPackageUploadRequest("PackageFile_example") // NugetPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadNuget(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadNuget``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadNugetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**NugetPackageUploadRequest**](NugetPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadP2

> PackagesValidateUploadP2(ctx, owner, repo).Data(data).Execute()

Validate parameters for create P2 package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewP2PackageUploadRequest("PackageFile_example") // P2PackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadP2(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadP2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadP2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**P2PackageUploadRequest**](P2PackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadPython

> PackagesValidateUploadPython(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Python package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewPythonPackageUploadRequest("PackageFile_example") // PythonPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadPython(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadPython``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadPythonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PythonPackageUploadRequest**](PythonPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadRaw

> PackagesValidateUploadRaw(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Raw package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewRawPackageUploadRequest("PackageFile_example") // RawPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadRaw(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadRaw``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RawPackageUploadRequest**](RawPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadRpm

> PackagesValidateUploadRpm(ctx, owner, repo).Data(data).Execute()

Validate parameters for create RedHat package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewRpmPackageUploadRequest("Distribution_example", "PackageFile_example") // RpmPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadRpm(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadRpm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadRpmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RpmPackageUploadRequest**](RpmPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadRuby

> PackagesValidateUploadRuby(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Ruby package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewRubyPackageUploadRequest("PackageFile_example") // RubyPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadRuby(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadRuby``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadRubyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RubyPackageUploadRequest**](RubyPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadSwift

> PackagesValidateUploadSwift(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Swift package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewSwiftPackageUploadRequest("PackageFile_example", "Version_example") // SwiftPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadSwift(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadSwift``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadSwiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**SwiftPackageUploadRequest**](SwiftPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadTerraform

> PackagesValidateUploadTerraform(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Terraform package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewTerraformPackageUploadRequest("PackageFile_example") // TerraformPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadTerraform(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadTerraform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**TerraformPackageUploadRequest**](TerraformPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadVagrant

> PackagesValidateUploadVagrant(ctx, owner, repo).Data(data).Execute()

Validate parameters for create Vagrant package



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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    data := *openapiclient.NewVagrantPackageUploadRequest("Name_example", "PackageFile_example", "Provider_example", "Version_example") // VagrantPackageUploadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesValidateUploadVagrant(context.Background(), owner, repo).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesValidateUploadVagrant``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPackagesValidateUploadVagrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**VagrantPackageUploadRequest**](VagrantPackageUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

