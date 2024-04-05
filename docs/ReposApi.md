# \ReposApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposCreate**](ReposApi.md#ReposCreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
[**ReposGeoipDisable**](ReposApi.md#ReposGeoipDisable) | **Post** /repos/{owner}/{identifier}/geoip/disable/ | Disable GeoIP for this repository.
[**ReposGeoipEnable**](ReposApi.md#ReposGeoipEnable) | **Post** /repos/{owner}/{identifier}/geoip/enable/ | Enable GeoIP for this repository.
[**ReposGeoipPartialUpdate**](ReposApi.md#ReposGeoipPartialUpdate) | **Patch** /repos/{owner}/{identifier}/geoip | Partially update repository geoip rules.
[**ReposGeoipRead**](ReposApi.md#ReposGeoipRead) | **Get** /repos/{owner}/{identifier}/geoip | List all repository geoip rules.
[**ReposGeoipTest**](ReposApi.md#ReposGeoipTest) | **Post** /repos/{owner}/{identifier}/geoip/test/ | Test a list of IP addresses against the repository&#39;s current GeoIP rules.
[**ReposGeoipUpdate**](ReposApi.md#ReposGeoipUpdate) | **Put** /repos/{owner}/{identifier}/geoip | Replace repository geoip rules.
[**ReposGpgCreate**](ReposApi.md#ReposGpgCreate) | **Post** /repos/{owner}/{identifier}/gpg/ | Set the active GPG key for the Repository.
[**ReposGpgList**](ReposApi.md#ReposGpgList) | **Get** /repos/{owner}/{identifier}/gpg/ | Retrieve the active GPG key for the Repository.
[**ReposGpgRegenerate**](ReposApi.md#ReposGpgRegenerate) | **Post** /repos/{owner}/{identifier}/gpg/regenerate/ | Regenerate GPG Key for the Repository.
[**ReposNamespaceList**](ReposApi.md#ReposNamespaceList) | **Get** /repos/{owner}/ | Get a list of all repositories within a namespace.
[**ReposPartialUpdate**](ReposApi.md#ReposPartialUpdate) | **Patch** /repos/{owner}/{identifier}/ | Update details about a repository in a given namespace.
[**ReposPrivilegesList**](ReposApi.md#ReposPrivilegesList) | **Get** /repos/{owner}/{identifier}/privileges | List all explicity created privileges for the repository.
[**ReposPrivilegesPartialUpdate**](ReposApi.md#ReposPrivilegesPartialUpdate) | **Patch** /repos/{owner}/{identifier}/privileges | Modify privileges for the repository.
[**ReposPrivilegesUpdate**](ReposApi.md#ReposPrivilegesUpdate) | **Put** /repos/{owner}/{identifier}/privileges | Replace all existing repository privileges with those specified.
[**ReposRead**](ReposApi.md#ReposRead) | **Get** /repos/{owner}/{identifier}/ | Get a specific repository.
[**ReposRsaCreate**](ReposApi.md#ReposRsaCreate) | **Post** /repos/{owner}/{identifier}/rsa/ | Set the active RSA key for the Repository.
[**ReposRsaList**](ReposApi.md#ReposRsaList) | **Get** /repos/{owner}/{identifier}/rsa/ | Retrieve the active RSA key for the Repository.
[**ReposRsaRegenerate**](ReposApi.md#ReposRsaRegenerate) | **Post** /repos/{owner}/{identifier}/rsa/regenerate/ | Regenerate RSA Key for the Repository.
[**ReposUpstreamCranCreate**](ReposApi.md#ReposUpstreamCranCreate) | **Post** /repos/{owner}/{identifier}/upstream/cran/ | Create a CRAN upstream config for this repository.
[**ReposUpstreamCranDelete**](ReposApi.md#ReposUpstreamCranDelete) | **Delete** /repos/{owner}/{identifier}/upstream/cran/{slug_perm}/ | Delete a CRAN upstream config for this repository.
[**ReposUpstreamCranList**](ReposApi.md#ReposUpstreamCranList) | **Get** /repos/{owner}/{identifier}/upstream/cran/ | List CRAN upstream configs for this repository.
[**ReposUpstreamCranPartialUpdate**](ReposApi.md#ReposUpstreamCranPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/cran/{slug_perm}/ | Partially update a CRAN upstream config for this repository.
[**ReposUpstreamCranRead**](ReposApi.md#ReposUpstreamCranRead) | **Get** /repos/{owner}/{identifier}/upstream/cran/{slug_perm}/ | Retrieve a CRAN upstream config for this repository.
[**ReposUpstreamCranUpdate**](ReposApi.md#ReposUpstreamCranUpdate) | **Put** /repos/{owner}/{identifier}/upstream/cran/{slug_perm}/ | Update a CRAN upstream config for this repository.
[**ReposUpstreamDartCreate**](ReposApi.md#ReposUpstreamDartCreate) | **Post** /repos/{owner}/{identifier}/upstream/dart/ | Create a Dart upstream config for this repository.
[**ReposUpstreamDartDelete**](ReposApi.md#ReposUpstreamDartDelete) | **Delete** /repos/{owner}/{identifier}/upstream/dart/{slug_perm}/ | Delete a Dart upstream config for this repository.
[**ReposUpstreamDartList**](ReposApi.md#ReposUpstreamDartList) | **Get** /repos/{owner}/{identifier}/upstream/dart/ | List Dart upstream configs for this repository.
[**ReposUpstreamDartPartialUpdate**](ReposApi.md#ReposUpstreamDartPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/dart/{slug_perm}/ | Partially update a Dart upstream config for this repository.
[**ReposUpstreamDartRead**](ReposApi.md#ReposUpstreamDartRead) | **Get** /repos/{owner}/{identifier}/upstream/dart/{slug_perm}/ | Retrieve a Dart upstream config for this repository.
[**ReposUpstreamDartUpdate**](ReposApi.md#ReposUpstreamDartUpdate) | **Put** /repos/{owner}/{identifier}/upstream/dart/{slug_perm}/ | Update a Dart upstream config for this repository.
[**ReposUpstreamDebCreate**](ReposApi.md#ReposUpstreamDebCreate) | **Post** /repos/{owner}/{identifier}/upstream/deb/ | Create a Debian upstream config for this repository.
[**ReposUpstreamDebDelete**](ReposApi.md#ReposUpstreamDebDelete) | **Delete** /repos/{owner}/{identifier}/upstream/deb/{slug_perm}/ | Delete a Debian upstream config for this repository.
[**ReposUpstreamDebList**](ReposApi.md#ReposUpstreamDebList) | **Get** /repos/{owner}/{identifier}/upstream/deb/ | List Debian upstream configs for this repository.
[**ReposUpstreamDebPartialUpdate**](ReposApi.md#ReposUpstreamDebPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/deb/{slug_perm}/ | Partially update a Debian upstream config for this repository.
[**ReposUpstreamDebRead**](ReposApi.md#ReposUpstreamDebRead) | **Get** /repos/{owner}/{identifier}/upstream/deb/{slug_perm}/ | Retrieve a Debian upstream config for this repository.
[**ReposUpstreamDebUpdate**](ReposApi.md#ReposUpstreamDebUpdate) | **Put** /repos/{owner}/{identifier}/upstream/deb/{slug_perm}/ | Update a Debian upstream config for this repository.
[**ReposUpstreamDockerCreate**](ReposApi.md#ReposUpstreamDockerCreate) | **Post** /repos/{owner}/{identifier}/upstream/docker/ | Create a Docker upstream config for this repository.
[**ReposUpstreamDockerDelete**](ReposApi.md#ReposUpstreamDockerDelete) | **Delete** /repos/{owner}/{identifier}/upstream/docker/{slug_perm}/ | Delete a Docker upstream config for this repository.
[**ReposUpstreamDockerList**](ReposApi.md#ReposUpstreamDockerList) | **Get** /repos/{owner}/{identifier}/upstream/docker/ | List Docker upstream configs for this repository.
[**ReposUpstreamDockerPartialUpdate**](ReposApi.md#ReposUpstreamDockerPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/docker/{slug_perm}/ | Partially update a Docker upstream config for this repository.
[**ReposUpstreamDockerRead**](ReposApi.md#ReposUpstreamDockerRead) | **Get** /repos/{owner}/{identifier}/upstream/docker/{slug_perm}/ | Retrieve a Docker upstream config for this repository.
[**ReposUpstreamDockerUpdate**](ReposApi.md#ReposUpstreamDockerUpdate) | **Put** /repos/{owner}/{identifier}/upstream/docker/{slug_perm}/ | Update a Docker upstream config for this repository.
[**ReposUpstreamHelmCreate**](ReposApi.md#ReposUpstreamHelmCreate) | **Post** /repos/{owner}/{identifier}/upstream/helm/ | Create a Helm upstream config for this repository.
[**ReposUpstreamHelmDelete**](ReposApi.md#ReposUpstreamHelmDelete) | **Delete** /repos/{owner}/{identifier}/upstream/helm/{slug_perm}/ | Delete a Helm upstream config for this repository.
[**ReposUpstreamHelmList**](ReposApi.md#ReposUpstreamHelmList) | **Get** /repos/{owner}/{identifier}/upstream/helm/ | List Helm upstream configs for this repository.
[**ReposUpstreamHelmPartialUpdate**](ReposApi.md#ReposUpstreamHelmPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/helm/{slug_perm}/ | Partially update a Helm upstream config for this repository.
[**ReposUpstreamHelmRead**](ReposApi.md#ReposUpstreamHelmRead) | **Get** /repos/{owner}/{identifier}/upstream/helm/{slug_perm}/ | Retrieve a Helm upstream config for this repository.
[**ReposUpstreamHelmUpdate**](ReposApi.md#ReposUpstreamHelmUpdate) | **Put** /repos/{owner}/{identifier}/upstream/helm/{slug_perm}/ | Update a Helm upstream config for this repository.
[**ReposUpstreamMavenCreate**](ReposApi.md#ReposUpstreamMavenCreate) | **Post** /repos/{owner}/{identifier}/upstream/maven/ | Create a Maven upstream config for this repository.
[**ReposUpstreamMavenDelete**](ReposApi.md#ReposUpstreamMavenDelete) | **Delete** /repos/{owner}/{identifier}/upstream/maven/{slug_perm}/ | Delete a Maven upstream config for this repository.
[**ReposUpstreamMavenList**](ReposApi.md#ReposUpstreamMavenList) | **Get** /repos/{owner}/{identifier}/upstream/maven/ | List Maven upstream configs for this repository.
[**ReposUpstreamMavenPartialUpdate**](ReposApi.md#ReposUpstreamMavenPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/maven/{slug_perm}/ | Partially update a Maven upstream config for this repository.
[**ReposUpstreamMavenRead**](ReposApi.md#ReposUpstreamMavenRead) | **Get** /repos/{owner}/{identifier}/upstream/maven/{slug_perm}/ | Retrieve a Maven upstream config for this repository.
[**ReposUpstreamMavenUpdate**](ReposApi.md#ReposUpstreamMavenUpdate) | **Put** /repos/{owner}/{identifier}/upstream/maven/{slug_perm}/ | Update a Maven upstream config for this repository.
[**ReposUpstreamNpmCreate**](ReposApi.md#ReposUpstreamNpmCreate) | **Post** /repos/{owner}/{identifier}/upstream/npm/ | Create a npm upstream config for this repository.
[**ReposUpstreamNpmDelete**](ReposApi.md#ReposUpstreamNpmDelete) | **Delete** /repos/{owner}/{identifier}/upstream/npm/{slug_perm}/ | Delete a npm upstream config for this repository.
[**ReposUpstreamNpmList**](ReposApi.md#ReposUpstreamNpmList) | **Get** /repos/{owner}/{identifier}/upstream/npm/ | List npm upstream configs for this repository.
[**ReposUpstreamNpmPartialUpdate**](ReposApi.md#ReposUpstreamNpmPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/npm/{slug_perm}/ | Partially update a npm upstream config for this repository.
[**ReposUpstreamNpmRead**](ReposApi.md#ReposUpstreamNpmRead) | **Get** /repos/{owner}/{identifier}/upstream/npm/{slug_perm}/ | Retrieve a npm upstream config for this repository.
[**ReposUpstreamNpmUpdate**](ReposApi.md#ReposUpstreamNpmUpdate) | **Put** /repos/{owner}/{identifier}/upstream/npm/{slug_perm}/ | Update a npm upstream config for this repository.
[**ReposUpstreamNugetCreate**](ReposApi.md#ReposUpstreamNugetCreate) | **Post** /repos/{owner}/{identifier}/upstream/nuget/ | Create a NuGet upstream config for this repository.
[**ReposUpstreamNugetDelete**](ReposApi.md#ReposUpstreamNugetDelete) | **Delete** /repos/{owner}/{identifier}/upstream/nuget/{slug_perm}/ | Delete a NuGet upstream config for this repository.
[**ReposUpstreamNugetList**](ReposApi.md#ReposUpstreamNugetList) | **Get** /repos/{owner}/{identifier}/upstream/nuget/ | List NuGet upstream configs for this repository.
[**ReposUpstreamNugetPartialUpdate**](ReposApi.md#ReposUpstreamNugetPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/nuget/{slug_perm}/ | Partially update a NuGet upstream config for this repository.
[**ReposUpstreamNugetRead**](ReposApi.md#ReposUpstreamNugetRead) | **Get** /repos/{owner}/{identifier}/upstream/nuget/{slug_perm}/ | Retrieve a NuGet upstream config for this repository.
[**ReposUpstreamNugetUpdate**](ReposApi.md#ReposUpstreamNugetUpdate) | **Put** /repos/{owner}/{identifier}/upstream/nuget/{slug_perm}/ | Update a NuGet upstream config for this repository.
[**ReposUpstreamPythonCreate**](ReposApi.md#ReposUpstreamPythonCreate) | **Post** /repos/{owner}/{identifier}/upstream/python/ | Create a Python upstream config for this repository.
[**ReposUpstreamPythonDelete**](ReposApi.md#ReposUpstreamPythonDelete) | **Delete** /repos/{owner}/{identifier}/upstream/python/{slug_perm}/ | Delete a Python upstream config for this repository.
[**ReposUpstreamPythonList**](ReposApi.md#ReposUpstreamPythonList) | **Get** /repos/{owner}/{identifier}/upstream/python/ | List Python upstream configs for this repository.
[**ReposUpstreamPythonPartialUpdate**](ReposApi.md#ReposUpstreamPythonPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/python/{slug_perm}/ | Partially update a Python upstream config for this repository.
[**ReposUpstreamPythonRead**](ReposApi.md#ReposUpstreamPythonRead) | **Get** /repos/{owner}/{identifier}/upstream/python/{slug_perm}/ | Retrieve a Python upstream config for this repository.
[**ReposUpstreamPythonUpdate**](ReposApi.md#ReposUpstreamPythonUpdate) | **Put** /repos/{owner}/{identifier}/upstream/python/{slug_perm}/ | Update a Python upstream config for this repository.
[**ReposUpstreamRpmCreate**](ReposApi.md#ReposUpstreamRpmCreate) | **Post** /repos/{owner}/{identifier}/upstream/rpm/ | Create a RedHat upstream config for this repository.
[**ReposUpstreamRpmDelete**](ReposApi.md#ReposUpstreamRpmDelete) | **Delete** /repos/{owner}/{identifier}/upstream/rpm/{slug_perm}/ | Delete a RedHat upstream config for this repository.
[**ReposUpstreamRpmList**](ReposApi.md#ReposUpstreamRpmList) | **Get** /repos/{owner}/{identifier}/upstream/rpm/ | List RedHat upstream configs for this repository.
[**ReposUpstreamRpmPartialUpdate**](ReposApi.md#ReposUpstreamRpmPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/rpm/{slug_perm}/ | Partially update a RedHat upstream config for this repository.
[**ReposUpstreamRpmRead**](ReposApi.md#ReposUpstreamRpmRead) | **Get** /repos/{owner}/{identifier}/upstream/rpm/{slug_perm}/ | Retrieve a RedHat upstream config for this repository.
[**ReposUpstreamRpmUpdate**](ReposApi.md#ReposUpstreamRpmUpdate) | **Put** /repos/{owner}/{identifier}/upstream/rpm/{slug_perm}/ | Update a RedHat upstream config for this repository.
[**ReposUpstreamRubyCreate**](ReposApi.md#ReposUpstreamRubyCreate) | **Post** /repos/{owner}/{identifier}/upstream/ruby/ | Create a Ruby upstream config for this repository.
[**ReposUpstreamRubyDelete**](ReposApi.md#ReposUpstreamRubyDelete) | **Delete** /repos/{owner}/{identifier}/upstream/ruby/{slug_perm}/ | Delete a Ruby upstream config for this repository.
[**ReposUpstreamRubyList**](ReposApi.md#ReposUpstreamRubyList) | **Get** /repos/{owner}/{identifier}/upstream/ruby/ | List Ruby upstream configs for this repository.
[**ReposUpstreamRubyPartialUpdate**](ReposApi.md#ReposUpstreamRubyPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/ruby/{slug_perm}/ | Partially update a Ruby upstream config for this repository.
[**ReposUpstreamRubyRead**](ReposApi.md#ReposUpstreamRubyRead) | **Get** /repos/{owner}/{identifier}/upstream/ruby/{slug_perm}/ | Retrieve a Ruby upstream config for this repository.
[**ReposUpstreamRubyUpdate**](ReposApi.md#ReposUpstreamRubyUpdate) | **Put** /repos/{owner}/{identifier}/upstream/ruby/{slug_perm}/ | Update a Ruby upstream config for this repository.
[**ReposUpstreamSwiftCreate**](ReposApi.md#ReposUpstreamSwiftCreate) | **Post** /repos/{owner}/{identifier}/upstream/swift/ | Create a Swift upstream config for this repository.
[**ReposUpstreamSwiftDelete**](ReposApi.md#ReposUpstreamSwiftDelete) | **Delete** /repos/{owner}/{identifier}/upstream/swift/{slug_perm}/ | Delete a Swift upstream config for this repository.
[**ReposUpstreamSwiftList**](ReposApi.md#ReposUpstreamSwiftList) | **Get** /repos/{owner}/{identifier}/upstream/swift/ | List Swift upstream configs for this repository.
[**ReposUpstreamSwiftPartialUpdate**](ReposApi.md#ReposUpstreamSwiftPartialUpdate) | **Patch** /repos/{owner}/{identifier}/upstream/swift/{slug_perm}/ | Partially update a Swift upstream config for this repository.
[**ReposUpstreamSwiftRead**](ReposApi.md#ReposUpstreamSwiftRead) | **Get** /repos/{owner}/{identifier}/upstream/swift/{slug_perm}/ | Retrieve a Swift upstream config for this repository.
[**ReposUpstreamSwiftUpdate**](ReposApi.md#ReposUpstreamSwiftUpdate) | **Put** /repos/{owner}/{identifier}/upstream/swift/{slug_perm}/ | Update a Swift upstream config for this repository.
[**ReposUserList**](ReposApi.md#ReposUserList) | **Get** /repos/ | Get a list of all repositories associated with current user.



## ReposCreate

> RepositoryCreate ReposCreate(ctx, owner).Data(data).Execute()

Create a new repository in a given namespace.



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
    data := *openapiclient.NewRepositoryCreateRequest("Name_example") // RepositoryCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreate(context.Background(), owner).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreate`: RepositoryCreate
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**RepositoryCreateRequest**](RepositoryCreateRequest.md) |  | 

### Return type

[**RepositoryCreate**](RepositoryCreate.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposDelete

> ReposDelete(ctx, owner, identifier).Execute()

Delete a repository in a given namespace.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposDelete(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteRequest struct via the builder pattern


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


## ReposGeoipDisable

> ReposGeoipDisable(ctx, owner, identifier).Data(data).Execute()

Disable GeoIP for this repository.



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
    identifier := "identifier_example" // string | 
    data := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposGeoipDisable(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **map[string]interface{}** |  | 

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


## ReposGeoipEnable

> ReposGeoipEnable(ctx, owner, identifier).Data(data).Execute()

Enable GeoIP for this repository.



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
    identifier := "identifier_example" // string | 
    data := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposGeoipEnable(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **map[string]interface{}** |  | 

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


## ReposGeoipPartialUpdate

> RepositoryGeoIpRules ReposGeoipPartialUpdate(ctx, owner, identifier).Data(data).Execute()

Partially update repository geoip rules.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryGeoIpRulesRequestPatch() // RepositoryGeoIpRulesRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipPartialUpdate`: RepositoryGeoIpRules
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGeoIpRulesRequestPatch**](RepositoryGeoIpRulesRequestPatch.md) |  | 

### Return type

[**RepositoryGeoIpRules**](RepositoryGeoIpRules.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGeoipRead

> RepositoryGeoIpRules ReposGeoipRead(ctx, owner, identifier).Execute()

List all repository geoip rules.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipRead(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipRead`: RepositoryGeoIpRules
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryGeoIpRules**](RepositoryGeoIpRules.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGeoipTest

> RepositoryGeoIpTestAddressResponse ReposGeoipTest(ctx, owner, identifier).Data(data).Execute()

Test a list of IP addresses against the repository's current GeoIP rules.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryGeoIpTestAddress([]string{"Addresses_example"}) // RepositoryGeoIpTestAddress |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipTest(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipTest`: RepositoryGeoIpTestAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGeoIpTestAddress**](RepositoryGeoIpTestAddress.md) |  | 

### Return type

[**RepositoryGeoIpTestAddressResponse**](RepositoryGeoIpTestAddressResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGeoipUpdate

> RepositoryGeoIpRules ReposGeoipUpdate(ctx, owner, identifier).Data(data).Execute()

Replace repository geoip rules.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryGeoIpRulesRequest(*openapiclient.NewRepositoryGeoIpCidr([]string{"Allow_example"}, []string{"Deny_example"}), *openapiclient.NewRepositoryGeoIpCountryCode([]string{"Allow_example"}, []string{"Deny_example"})) // RepositoryGeoIpRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipUpdate`: RepositoryGeoIpRules
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGeoIpRulesRequest**](RepositoryGeoIpRulesRequest.md) |  | 

### Return type

[**RepositoryGeoIpRules**](RepositoryGeoIpRules.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgCreate

> RepositoryGpgKey ReposGpgCreate(ctx, owner, identifier).Data(data).Execute()

Set the active GPG key for the Repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryGpgKeyCreate("GpgPrivateKey_example") // RepositoryGpgKeyCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGpgCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgCreate`: RepositoryGpgKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGpgCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGpgCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGpgKeyCreate**](RepositoryGpgKeyCreate.md) |  | 

### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgList

> RepositoryGpgKey ReposGpgList(ctx, owner, identifier).Execute()

Retrieve the active GPG key for the Repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGpgList(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgList`: RepositoryGpgKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGpgList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGpgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgRegenerate

> RepositoryGpgKey ReposGpgRegenerate(ctx, owner, identifier).Execute()

Regenerate GPG Key for the Repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGpgRegenerate(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgRegenerate`: RepositoryGpgKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGpgRegenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGpgRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposNamespaceList

> []Repository ReposNamespaceList(ctx, owner).Page(page).PageSize(pageSize).Execute()

Get a list of all repositories within a namespace.



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
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposNamespaceList(context.Background(), owner).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposNamespaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposNamespaceList`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPartialUpdate

> Repository ReposPartialUpdate(ctx, owner, identifier).Data(data).Execute()

Update details about a repository in a given namespace.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryRequestPatch() // RepositoryRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPartialUpdate`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryRequestPatch**](RepositoryRequestPatch.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesList

> RepositoryPrivilegeInput ReposPrivilegesList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List all explicity created privileges for the repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPrivilegesList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPrivilegesList`: RepositoryPrivilegeInput
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposPrivilegesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPrivilegesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**RepositoryPrivilegeInput**](RepositoryPrivilegeInput.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesPartialUpdate

> ReposPrivilegesPartialUpdate(ctx, owner, identifier).Data(data).Execute()

Modify privileges for the repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryPrivilegeInputRequestPatch() // RepositoryPrivilegeInputRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposPrivilegesPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPrivilegesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryPrivilegeInputRequestPatch**](RepositoryPrivilegeInputRequestPatch.md) |  | 

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


## ReposPrivilegesUpdate

> ReposPrivilegesUpdate(ctx, owner, identifier).Data(data).Execute()

Replace all existing repository privileges with those specified.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryPrivilegeInputRequest([]openapiclient.RepositoryPrivilegeDict{*openapiclient.NewRepositoryPrivilegeDict("Privilege_example")}) // RepositoryPrivilegeInputRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposPrivilegesUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPrivilegesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryPrivilegeInputRequest**](RepositoryPrivilegeInputRequest.md) |  | 

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


## ReposRead

> Repository ReposRead(ctx, owner, identifier).Execute()

Get a specific repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRead(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRead`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaCreate

> RepositoryRsaKey ReposRsaCreate(ctx, owner, identifier).Data(data).Execute()

Set the active RSA key for the Repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRepositoryRsaKeyCreate("RsaPrivateKey_example") // RepositoryRsaKeyCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRsaCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaCreate`: RepositoryRsaKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRsaCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRsaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryRsaKeyCreate**](RepositoryRsaKeyCreate.md) |  | 

### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaList

> RepositoryRsaKey ReposRsaList(ctx, owner, identifier).Execute()

Retrieve the active RSA key for the Repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRsaList(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaList`: RepositoryRsaKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRsaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRsaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaRegenerate

> RepositoryRsaKey ReposRsaRegenerate(ctx, owner, identifier).Execute()

Regenerate RSA Key for the Repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRsaRegenerate(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaRegenerate`: RepositoryRsaKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRsaRegenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRsaRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamCranCreate

> CranUpstream ReposUpstreamCranCreate(ctx, owner, identifier).Data(data).Execute()

Create a CRAN upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewCranUpstreamRequest("Name_example", "UpstreamUrl_example") // CranUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamCranCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamCranCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamCranCreate`: CranUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamCranCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamCranCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**CranUpstreamRequest**](CranUpstreamRequest.md) |  | 

### Return type

[**CranUpstream**](CranUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamCranDelete

> ReposUpstreamCranDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a CRAN upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamCranDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamCranDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamCranDeleteRequest struct via the builder pattern


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


## ReposUpstreamCranList

> []CranUpstream ReposUpstreamCranList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List CRAN upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamCranList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamCranList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamCranList`: []CranUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamCranList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamCranListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]CranUpstream**](CranUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamCranPartialUpdate

> CranUpstream ReposUpstreamCranPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a CRAN upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewCranUpstreamRequestPatch() // CranUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamCranPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamCranPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamCranPartialUpdate`: CranUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamCranPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamCranPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**CranUpstreamRequestPatch**](CranUpstreamRequestPatch.md) |  | 

### Return type

[**CranUpstream**](CranUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamCranRead

> CranUpstream ReposUpstreamCranRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a CRAN upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamCranRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamCranRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamCranRead`: CranUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamCranRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamCranReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CranUpstream**](CranUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamCranUpdate

> CranUpstream ReposUpstreamCranUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a CRAN upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewCranUpstreamRequest("Name_example", "UpstreamUrl_example") // CranUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamCranUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamCranUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamCranUpdate`: CranUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamCranUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamCranUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**CranUpstreamRequest**](CranUpstreamRequest.md) |  | 

### Return type

[**CranUpstream**](CranUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDartCreate

> DartUpstream ReposUpstreamDartCreate(ctx, owner, identifier).Data(data).Execute()

Create a Dart upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewDartUpstreamRequest("Name_example", "UpstreamUrl_example") // DartUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDartCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDartCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDartCreate`: DartUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDartCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDartCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DartUpstreamRequest**](DartUpstreamRequest.md) |  | 

### Return type

[**DartUpstream**](DartUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDartDelete

> ReposUpstreamDartDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Dart upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamDartDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDartDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDartDeleteRequest struct via the builder pattern


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


## ReposUpstreamDartList

> []DartUpstream ReposUpstreamDartList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Dart upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDartList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDartList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDartList`: []DartUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDartList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDartListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]DartUpstream**](DartUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDartPartialUpdate

> DartUpstream ReposUpstreamDartPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Dart upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewDartUpstreamRequestPatch() // DartUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDartPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDartPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDartPartialUpdate`: DartUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDartPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDartPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**DartUpstreamRequestPatch**](DartUpstreamRequestPatch.md) |  | 

### Return type

[**DartUpstream**](DartUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDartRead

> DartUpstream ReposUpstreamDartRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Dart upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDartRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDartRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDartRead`: DartUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDartRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDartReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DartUpstream**](DartUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDartUpdate

> DartUpstream ReposUpstreamDartUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Dart upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewDartUpstreamRequest("Name_example", "UpstreamUrl_example") // DartUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDartUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDartUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDartUpdate`: DartUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDartUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDartUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**DartUpstreamRequest**](DartUpstreamRequest.md) |  | 

### Return type

[**DartUpstream**](DartUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDebCreate

> DebUpstream ReposUpstreamDebCreate(ctx, owner, identifier).Data(data).Execute()

Create a Debian upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewDebUpstreamRequest([]string{"DistroVersions_example"}, "Name_example", "UpstreamUrl_example") // DebUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDebCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDebCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDebCreate`: DebUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDebCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDebCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DebUpstreamRequest**](DebUpstreamRequest.md) |  | 

### Return type

[**DebUpstream**](DebUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDebDelete

> ReposUpstreamDebDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Debian upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamDebDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDebDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDebDeleteRequest struct via the builder pattern


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


## ReposUpstreamDebList

> []DebUpstream ReposUpstreamDebList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Debian upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDebList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDebList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDebList`: []DebUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDebList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDebListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]DebUpstream**](DebUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDebPartialUpdate

> DebUpstream ReposUpstreamDebPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Debian upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewDebUpstreamRequestPatch() // DebUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDebPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDebPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDebPartialUpdate`: DebUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDebPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDebPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**DebUpstreamRequestPatch**](DebUpstreamRequestPatch.md) |  | 

### Return type

[**DebUpstream**](DebUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDebRead

> DebUpstream ReposUpstreamDebRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Debian upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDebRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDebRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDebRead`: DebUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDebRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDebReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DebUpstream**](DebUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDebUpdate

> DebUpstream ReposUpstreamDebUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Debian upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewDebUpstreamRequest([]string{"DistroVersions_example"}, "Name_example", "UpstreamUrl_example") // DebUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDebUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDebUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDebUpdate`: DebUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDebUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDebUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**DebUpstreamRequest**](DebUpstreamRequest.md) |  | 

### Return type

[**DebUpstream**](DebUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDockerCreate

> DockerUpstream ReposUpstreamDockerCreate(ctx, owner, identifier).Data(data).Execute()

Create a Docker upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewDockerUpstreamRequest("Name_example", "UpstreamUrl_example") // DockerUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDockerCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDockerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDockerCreate`: DockerUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDockerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDockerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**DockerUpstreamRequest**](DockerUpstreamRequest.md) |  | 

### Return type

[**DockerUpstream**](DockerUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDockerDelete

> ReposUpstreamDockerDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Docker upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamDockerDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDockerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDockerDeleteRequest struct via the builder pattern


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


## ReposUpstreamDockerList

> []DockerUpstream ReposUpstreamDockerList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Docker upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDockerList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDockerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDockerList`: []DockerUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDockerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDockerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]DockerUpstream**](DockerUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDockerPartialUpdate

> DockerUpstream ReposUpstreamDockerPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Docker upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewDockerUpstreamRequestPatch() // DockerUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDockerPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDockerPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDockerPartialUpdate`: DockerUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDockerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDockerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**DockerUpstreamRequestPatch**](DockerUpstreamRequestPatch.md) |  | 

### Return type

[**DockerUpstream**](DockerUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDockerRead

> DockerUpstream ReposUpstreamDockerRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Docker upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDockerRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDockerRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDockerRead`: DockerUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDockerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDockerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DockerUpstream**](DockerUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamDockerUpdate

> DockerUpstream ReposUpstreamDockerUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Docker upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewDockerUpstreamRequest("Name_example", "UpstreamUrl_example") // DockerUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamDockerUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamDockerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamDockerUpdate`: DockerUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamDockerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamDockerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**DockerUpstreamRequest**](DockerUpstreamRequest.md) |  | 

### Return type

[**DockerUpstream**](DockerUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamHelmCreate

> HelmUpstream ReposUpstreamHelmCreate(ctx, owner, identifier).Data(data).Execute()

Create a Helm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewHelmUpstreamRequest("Name_example", "UpstreamUrl_example") // HelmUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamHelmCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamHelmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamHelmCreate`: HelmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamHelmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamHelmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**HelmUpstreamRequest**](HelmUpstreamRequest.md) |  | 

### Return type

[**HelmUpstream**](HelmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamHelmDelete

> ReposUpstreamHelmDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Helm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamHelmDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamHelmDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamHelmDeleteRequest struct via the builder pattern


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


## ReposUpstreamHelmList

> []HelmUpstream ReposUpstreamHelmList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Helm upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamHelmList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamHelmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamHelmList`: []HelmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamHelmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamHelmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]HelmUpstream**](HelmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamHelmPartialUpdate

> HelmUpstream ReposUpstreamHelmPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Helm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewHelmUpstreamRequestPatch() // HelmUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamHelmPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamHelmPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamHelmPartialUpdate`: HelmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamHelmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamHelmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**HelmUpstreamRequestPatch**](HelmUpstreamRequestPatch.md) |  | 

### Return type

[**HelmUpstream**](HelmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamHelmRead

> HelmUpstream ReposUpstreamHelmRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Helm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamHelmRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamHelmRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamHelmRead`: HelmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamHelmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamHelmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**HelmUpstream**](HelmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamHelmUpdate

> HelmUpstream ReposUpstreamHelmUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Helm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewHelmUpstreamRequest("Name_example", "UpstreamUrl_example") // HelmUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamHelmUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamHelmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamHelmUpdate`: HelmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamHelmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamHelmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**HelmUpstreamRequest**](HelmUpstreamRequest.md) |  | 

### Return type

[**HelmUpstream**](HelmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamMavenCreate

> MavenUpstream ReposUpstreamMavenCreate(ctx, owner, identifier).Data(data).Execute()

Create a Maven upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewMavenUpstreamRequest("Name_example", "UpstreamUrl_example") // MavenUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamMavenCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamMavenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamMavenCreate`: MavenUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamMavenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamMavenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**MavenUpstreamRequest**](MavenUpstreamRequest.md) |  | 

### Return type

[**MavenUpstream**](MavenUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamMavenDelete

> ReposUpstreamMavenDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Maven upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamMavenDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamMavenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamMavenDeleteRequest struct via the builder pattern


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


## ReposUpstreamMavenList

> []MavenUpstream ReposUpstreamMavenList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Maven upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamMavenList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamMavenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamMavenList`: []MavenUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamMavenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamMavenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]MavenUpstream**](MavenUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamMavenPartialUpdate

> MavenUpstream ReposUpstreamMavenPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Maven upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewMavenUpstreamRequestPatch() // MavenUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamMavenPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamMavenPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamMavenPartialUpdate`: MavenUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamMavenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamMavenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**MavenUpstreamRequestPatch**](MavenUpstreamRequestPatch.md) |  | 

### Return type

[**MavenUpstream**](MavenUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamMavenRead

> MavenUpstream ReposUpstreamMavenRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Maven upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamMavenRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamMavenRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamMavenRead`: MavenUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamMavenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamMavenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MavenUpstream**](MavenUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamMavenUpdate

> MavenUpstream ReposUpstreamMavenUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Maven upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewMavenUpstreamRequest("Name_example", "UpstreamUrl_example") // MavenUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamMavenUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamMavenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamMavenUpdate`: MavenUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamMavenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamMavenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**MavenUpstreamRequest**](MavenUpstreamRequest.md) |  | 

### Return type

[**MavenUpstream**](MavenUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNpmCreate

> NpmUpstream ReposUpstreamNpmCreate(ctx, owner, identifier).Data(data).Execute()

Create a npm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewNpmUpstreamRequest("Name_example", "UpstreamUrl_example") // NpmUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNpmCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNpmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNpmCreate`: NpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNpmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**NpmUpstreamRequest**](NpmUpstreamRequest.md) |  | 

### Return type

[**NpmUpstream**](NpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNpmDelete

> ReposUpstreamNpmDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a npm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamNpmDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNpmDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNpmDeleteRequest struct via the builder pattern


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


## ReposUpstreamNpmList

> []NpmUpstream ReposUpstreamNpmList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List npm upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNpmList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNpmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNpmList`: []NpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNpmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNpmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]NpmUpstream**](NpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNpmPartialUpdate

> NpmUpstream ReposUpstreamNpmPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a npm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewNpmUpstreamRequestPatch() // NpmUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNpmPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNpmPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNpmPartialUpdate`: NpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**NpmUpstreamRequestPatch**](NpmUpstreamRequestPatch.md) |  | 

### Return type

[**NpmUpstream**](NpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNpmRead

> NpmUpstream ReposUpstreamNpmRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a npm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNpmRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNpmRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNpmRead`: NpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NpmUpstream**](NpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNpmUpdate

> NpmUpstream ReposUpstreamNpmUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a npm upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewNpmUpstreamRequest("Name_example", "UpstreamUrl_example") // NpmUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNpmUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNpmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNpmUpdate`: NpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**NpmUpstreamRequest**](NpmUpstreamRequest.md) |  | 

### Return type

[**NpmUpstream**](NpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNugetCreate

> NugetUpstream ReposUpstreamNugetCreate(ctx, owner, identifier).Data(data).Execute()

Create a NuGet upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewNugetUpstreamRequest("Name_example", "UpstreamUrl_example") // NugetUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNugetCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNugetCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNugetCreate`: NugetUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNugetCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNugetCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**NugetUpstreamRequest**](NugetUpstreamRequest.md) |  | 

### Return type

[**NugetUpstream**](NugetUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNugetDelete

> ReposUpstreamNugetDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a NuGet upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamNugetDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNugetDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNugetDeleteRequest struct via the builder pattern


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


## ReposUpstreamNugetList

> []NugetUpstream ReposUpstreamNugetList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List NuGet upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNugetList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNugetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNugetList`: []NugetUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNugetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNugetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]NugetUpstream**](NugetUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNugetPartialUpdate

> NugetUpstream ReposUpstreamNugetPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a NuGet upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewNugetUpstreamRequestPatch() // NugetUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNugetPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNugetPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNugetPartialUpdate`: NugetUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNugetPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNugetPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**NugetUpstreamRequestPatch**](NugetUpstreamRequestPatch.md) |  | 

### Return type

[**NugetUpstream**](NugetUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNugetRead

> NugetUpstream ReposUpstreamNugetRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a NuGet upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNugetRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNugetRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNugetRead`: NugetUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNugetRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNugetReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NugetUpstream**](NugetUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamNugetUpdate

> NugetUpstream ReposUpstreamNugetUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a NuGet upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewNugetUpstreamRequest("Name_example", "UpstreamUrl_example") // NugetUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamNugetUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamNugetUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamNugetUpdate`: NugetUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamNugetUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamNugetUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**NugetUpstreamRequest**](NugetUpstreamRequest.md) |  | 

### Return type

[**NugetUpstream**](NugetUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamPythonCreate

> PythonUpstream ReposUpstreamPythonCreate(ctx, owner, identifier).Data(data).Execute()

Create a Python upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewPythonUpstreamRequest("Name_example", "UpstreamUrl_example") // PythonUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamPythonCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamPythonCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamPythonCreate`: PythonUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamPythonCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamPythonCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PythonUpstreamRequest**](PythonUpstreamRequest.md) |  | 

### Return type

[**PythonUpstream**](PythonUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamPythonDelete

> ReposUpstreamPythonDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Python upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamPythonDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamPythonDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamPythonDeleteRequest struct via the builder pattern


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


## ReposUpstreamPythonList

> []PythonUpstream ReposUpstreamPythonList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Python upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamPythonList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamPythonList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamPythonList`: []PythonUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamPythonList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamPythonListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]PythonUpstream**](PythonUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamPythonPartialUpdate

> PythonUpstream ReposUpstreamPythonPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Python upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewPythonUpstreamRequestPatch() // PythonUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamPythonPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamPythonPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamPythonPartialUpdate`: PythonUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamPythonPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamPythonPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PythonUpstreamRequestPatch**](PythonUpstreamRequestPatch.md) |  | 

### Return type

[**PythonUpstream**](PythonUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamPythonRead

> PythonUpstream ReposUpstreamPythonRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Python upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamPythonRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamPythonRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamPythonRead`: PythonUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamPythonRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamPythonReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PythonUpstream**](PythonUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamPythonUpdate

> PythonUpstream ReposUpstreamPythonUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Python upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewPythonUpstreamRequest("Name_example", "UpstreamUrl_example") // PythonUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamPythonUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamPythonUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamPythonUpdate`: PythonUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamPythonUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamPythonUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**PythonUpstreamRequest**](PythonUpstreamRequest.md) |  | 

### Return type

[**PythonUpstream**](PythonUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRpmCreate

> RpmUpstream ReposUpstreamRpmCreate(ctx, owner, identifier).Data(data).Execute()

Create a RedHat upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRpmUpstreamRequest("DistroVersion_example", "Name_example", "UpstreamUrl_example") // RpmUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRpmCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRpmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRpmCreate`: RpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRpmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RpmUpstreamRequest**](RpmUpstreamRequest.md) |  | 

### Return type

[**RpmUpstream**](RpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRpmDelete

> ReposUpstreamRpmDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a RedHat upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamRpmDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRpmDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRpmDeleteRequest struct via the builder pattern


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


## ReposUpstreamRpmList

> []RpmUpstream ReposUpstreamRpmList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List RedHat upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRpmList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRpmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRpmList`: []RpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRpmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRpmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]RpmUpstream**](RpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRpmPartialUpdate

> RpmUpstream ReposUpstreamRpmPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a RedHat upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewRpmUpstreamRequestPatch() // RpmUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRpmPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRpmPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRpmPartialUpdate`: RpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**RpmUpstreamRequestPatch**](RpmUpstreamRequestPatch.md) |  | 

### Return type

[**RpmUpstream**](RpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRpmRead

> RpmUpstream ReposUpstreamRpmRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a RedHat upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRpmRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRpmRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRpmRead`: RpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RpmUpstream**](RpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRpmUpdate

> RpmUpstream ReposUpstreamRpmUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a RedHat upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewRpmUpstreamRequest("DistroVersion_example", "Name_example", "UpstreamUrl_example") // RpmUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRpmUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRpmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRpmUpdate`: RpmUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**RpmUpstreamRequest**](RpmUpstreamRequest.md) |  | 

### Return type

[**RpmUpstream**](RpmUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRubyCreate

> RubyUpstream ReposUpstreamRubyCreate(ctx, owner, identifier).Data(data).Execute()

Create a Ruby upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewRubyUpstreamRequest("Name_example", "UpstreamUrl_example") // RubyUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRubyCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRubyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRubyCreate`: RubyUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRubyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRubyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RubyUpstreamRequest**](RubyUpstreamRequest.md) |  | 

### Return type

[**RubyUpstream**](RubyUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRubyDelete

> ReposUpstreamRubyDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Ruby upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamRubyDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRubyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRubyDeleteRequest struct via the builder pattern


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


## ReposUpstreamRubyList

> []RubyUpstream ReposUpstreamRubyList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Ruby upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRubyList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRubyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRubyList`: []RubyUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRubyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRubyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]RubyUpstream**](RubyUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRubyPartialUpdate

> RubyUpstream ReposUpstreamRubyPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Ruby upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewRubyUpstreamRequestPatch() // RubyUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRubyPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRubyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRubyPartialUpdate`: RubyUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRubyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRubyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**RubyUpstreamRequestPatch**](RubyUpstreamRequestPatch.md) |  | 

### Return type

[**RubyUpstream**](RubyUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRubyRead

> RubyUpstream ReposUpstreamRubyRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Ruby upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRubyRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRubyRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRubyRead`: RubyUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRubyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRubyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RubyUpstream**](RubyUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamRubyUpdate

> RubyUpstream ReposUpstreamRubyUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Ruby upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewRubyUpstreamRequest("Name_example", "UpstreamUrl_example") // RubyUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamRubyUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamRubyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamRubyUpdate`: RubyUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamRubyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamRubyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**RubyUpstreamRequest**](RubyUpstreamRequest.md) |  | 

### Return type

[**RubyUpstream**](RubyUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamSwiftCreate

> SwiftUpstream ReposUpstreamSwiftCreate(ctx, owner, identifier).Data(data).Execute()

Create a Swift upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewSwiftUpstreamRequest("Name_example", "UpstreamUrl_example") // SwiftUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamSwiftCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamSwiftCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamSwiftCreate`: SwiftUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamSwiftCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamSwiftCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**SwiftUpstreamRequest**](SwiftUpstreamRequest.md) |  | 

### Return type

[**SwiftUpstream**](SwiftUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamSwiftDelete

> ReposUpstreamSwiftDelete(ctx, owner, identifier, slugPerm).Execute()

Delete a Swift upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReposApi.ReposUpstreamSwiftDelete(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamSwiftDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamSwiftDeleteRequest struct via the builder pattern


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


## ReposUpstreamSwiftList

> []SwiftUpstream ReposUpstreamSwiftList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

List Swift upstream configs for this repository.



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
    identifier := "identifier_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamSwiftList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamSwiftList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamSwiftList`: []SwiftUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamSwiftList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamSwiftListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]SwiftUpstream**](SwiftUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamSwiftPartialUpdate

> SwiftUpstream ReposUpstreamSwiftPartialUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Partially update a Swift upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewSwiftUpstreamRequestPatch() // SwiftUpstreamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamSwiftPartialUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamSwiftPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamSwiftPartialUpdate`: SwiftUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamSwiftPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamSwiftPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**SwiftUpstreamRequestPatch**](SwiftUpstreamRequestPatch.md) |  | 

### Return type

[**SwiftUpstream**](SwiftUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamSwiftRead

> SwiftUpstream ReposUpstreamSwiftRead(ctx, owner, identifier, slugPerm).Execute()

Retrieve a Swift upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamSwiftRead(context.Background(), owner, identifier, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamSwiftRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamSwiftRead`: SwiftUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamSwiftRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamSwiftReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SwiftUpstream**](SwiftUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpstreamSwiftUpdate

> SwiftUpstream ReposUpstreamSwiftUpdate(ctx, owner, identifier, slugPerm).Data(data).Execute()

Update a Swift upstream config for this repository.



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
    identifier := "identifier_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewSwiftUpstreamRequest("Name_example", "UpstreamUrl_example") // SwiftUpstreamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpstreamSwiftUpdate(context.Background(), owner, identifier, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpstreamSwiftUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpstreamSwiftUpdate`: SwiftUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpstreamSwiftUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpstreamSwiftUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**SwiftUpstreamRequest**](SwiftUpstreamRequest.md) |  | 

### Return type

[**SwiftUpstream**](SwiftUpstream.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUserList

> []Repository ReposUserList(ctx).Page(page).PageSize(pageSize).Execute()

Get a list of all repositories associated with current user.



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
    resp, r, err := apiClient.ReposApi.ReposUserList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUserList`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

