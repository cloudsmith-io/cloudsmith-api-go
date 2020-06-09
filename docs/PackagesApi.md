# \PackagesApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PackagesCopy**](PackagesApi.md#PackagesCopy) | **Post** /packages/{owner}/{repo}/{identifier}/copy/ | Copy a package to another repository.
[**PackagesDelete**](PackagesApi.md#PackagesDelete) | **Delete** /packages/{owner}/{repo}/{identifier}/ | Delete a specific package in a repository.
[**PackagesList**](PackagesApi.md#PackagesList) | **Get** /packages/{owner}/{repo}/ | Get a list of all packages associated with repository.
[**PackagesMove**](PackagesApi.md#PackagesMove) | **Post** /packages/{owner}/{repo}/{identifier}/move/ | Move a package to another repository.
[**PackagesRead**](PackagesApi.md#PackagesRead) | **Get** /packages/{owner}/{repo}/{identifier}/ | Get a specific package in a repository.
[**PackagesResync**](PackagesApi.md#PackagesResync) | **Post** /packages/{owner}/{repo}/{identifier}/resync/ | Schedule a package for resynchronisation.
[**PackagesStatus**](PackagesApi.md#PackagesStatus) | **Get** /packages/{owner}/{repo}/{identifier}/status/ | Get the synchronisation status for a package.
[**PackagesUploadAlpine**](PackagesApi.md#PackagesUploadAlpine) | **Post** /packages/{owner}/{repo}/upload/alpine/ | Create a new Alpine package
[**PackagesUploadCargo**](PackagesApi.md#PackagesUploadCargo) | **Post** /packages/{owner}/{repo}/upload/cargo/ | Create a new Cargo package
[**PackagesUploadCocoapods**](PackagesApi.md#PackagesUploadCocoapods) | **Post** /packages/{owner}/{repo}/upload/cocoapods/ | Create a new CocoaPods package
[**PackagesUploadComposer**](PackagesApi.md#PackagesUploadComposer) | **Post** /packages/{owner}/{repo}/upload/composer/ | Create a new Composer package
[**PackagesUploadConan**](PackagesApi.md#PackagesUploadConan) | **Post** /packages/{owner}/{repo}/upload/conan/ | Create a new Conan package
[**PackagesUploadCran**](PackagesApi.md#PackagesUploadCran) | **Post** /packages/{owner}/{repo}/upload/cran/ | Create a new CRAN package
[**PackagesUploadDart**](PackagesApi.md#PackagesUploadDart) | **Post** /packages/{owner}/{repo}/upload/dart/ | Create a new Dart package
[**PackagesUploadDeb**](PackagesApi.md#PackagesUploadDeb) | **Post** /packages/{owner}/{repo}/upload/deb/ | Create a new Debian package
[**PackagesUploadDocker**](PackagesApi.md#PackagesUploadDocker) | **Post** /packages/{owner}/{repo}/upload/docker/ | Create a new Docker package
[**PackagesUploadGo**](PackagesApi.md#PackagesUploadGo) | **Post** /packages/{owner}/{repo}/upload/go/ | Create a new Go package
[**PackagesUploadHelm**](PackagesApi.md#PackagesUploadHelm) | **Post** /packages/{owner}/{repo}/upload/helm/ | Create a new Helm package
[**PackagesUploadLuarocks**](PackagesApi.md#PackagesUploadLuarocks) | **Post** /packages/{owner}/{repo}/upload/luarocks/ | Create a new LuaRocks package
[**PackagesUploadMaven**](PackagesApi.md#PackagesUploadMaven) | **Post** /packages/{owner}/{repo}/upload/maven/ | Create a new Maven package
[**PackagesUploadNpm**](PackagesApi.md#PackagesUploadNpm) | **Post** /packages/{owner}/{repo}/upload/npm/ | Create a new npm package
[**PackagesUploadNuget**](PackagesApi.md#PackagesUploadNuget) | **Post** /packages/{owner}/{repo}/upload/nuget/ | Create a new NuGet package
[**PackagesUploadPython**](PackagesApi.md#PackagesUploadPython) | **Post** /packages/{owner}/{repo}/upload/python/ | Create a new Python package
[**PackagesUploadRaw**](PackagesApi.md#PackagesUploadRaw) | **Post** /packages/{owner}/{repo}/upload/raw/ | Create a new Raw package
[**PackagesUploadRpm**](PackagesApi.md#PackagesUploadRpm) | **Post** /packages/{owner}/{repo}/upload/rpm/ | Create a new RedHat package
[**PackagesUploadRuby**](PackagesApi.md#PackagesUploadRuby) | **Post** /packages/{owner}/{repo}/upload/ruby/ | Create a new Ruby package
[**PackagesUploadVagrant**](PackagesApi.md#PackagesUploadVagrant) | **Post** /packages/{owner}/{repo}/upload/vagrant/ | Create a new Vagrant package
[**PackagesValidateUploadAlpine**](PackagesApi.md#PackagesValidateUploadAlpine) | **Post** /packages/{owner}/{repo}/validate-upload/alpine/ | Validate parameters for create Alpine package
[**PackagesValidateUploadCargo**](PackagesApi.md#PackagesValidateUploadCargo) | **Post** /packages/{owner}/{repo}/validate-upload/cargo/ | Validate parameters for create Cargo package
[**PackagesValidateUploadCocoapods**](PackagesApi.md#PackagesValidateUploadCocoapods) | **Post** /packages/{owner}/{repo}/validate-upload/cocoapods/ | Validate parameters for create CocoaPods package
[**PackagesValidateUploadComposer**](PackagesApi.md#PackagesValidateUploadComposer) | **Post** /packages/{owner}/{repo}/validate-upload/composer/ | Validate parameters for create Composer package
[**PackagesValidateUploadConan**](PackagesApi.md#PackagesValidateUploadConan) | **Post** /packages/{owner}/{repo}/validate-upload/conan/ | Validate parameters for create Conan package
[**PackagesValidateUploadCran**](PackagesApi.md#PackagesValidateUploadCran) | **Post** /packages/{owner}/{repo}/validate-upload/cran/ | Validate parameters for create CRAN package
[**PackagesValidateUploadDart**](PackagesApi.md#PackagesValidateUploadDart) | **Post** /packages/{owner}/{repo}/validate-upload/dart/ | Validate parameters for create Dart package
[**PackagesValidateUploadDeb**](PackagesApi.md#PackagesValidateUploadDeb) | **Post** /packages/{owner}/{repo}/validate-upload/deb/ | Validate parameters for create Debian package
[**PackagesValidateUploadDocker**](PackagesApi.md#PackagesValidateUploadDocker) | **Post** /packages/{owner}/{repo}/validate-upload/docker/ | Validate parameters for create Docker package
[**PackagesValidateUploadGo**](PackagesApi.md#PackagesValidateUploadGo) | **Post** /packages/{owner}/{repo}/validate-upload/go/ | Validate parameters for create Go package
[**PackagesValidateUploadHelm**](PackagesApi.md#PackagesValidateUploadHelm) | **Post** /packages/{owner}/{repo}/validate-upload/helm/ | Validate parameters for create Helm package
[**PackagesValidateUploadLuarocks**](PackagesApi.md#PackagesValidateUploadLuarocks) | **Post** /packages/{owner}/{repo}/validate-upload/luarocks/ | Validate parameters for create LuaRocks package
[**PackagesValidateUploadMaven**](PackagesApi.md#PackagesValidateUploadMaven) | **Post** /packages/{owner}/{repo}/validate-upload/maven/ | Validate parameters for create Maven package
[**PackagesValidateUploadNpm**](PackagesApi.md#PackagesValidateUploadNpm) | **Post** /packages/{owner}/{repo}/validate-upload/npm/ | Validate parameters for create npm package
[**PackagesValidateUploadNuget**](PackagesApi.md#PackagesValidateUploadNuget) | **Post** /packages/{owner}/{repo}/validate-upload/nuget/ | Validate parameters for create NuGet package
[**PackagesValidateUploadPython**](PackagesApi.md#PackagesValidateUploadPython) | **Post** /packages/{owner}/{repo}/validate-upload/python/ | Validate parameters for create Python package
[**PackagesValidateUploadRaw**](PackagesApi.md#PackagesValidateUploadRaw) | **Post** /packages/{owner}/{repo}/validate-upload/raw/ | Validate parameters for create Raw package
[**PackagesValidateUploadRpm**](PackagesApi.md#PackagesValidateUploadRpm) | **Post** /packages/{owner}/{repo}/validate-upload/rpm/ | Validate parameters for create RedHat package
[**PackagesValidateUploadRuby**](PackagesApi.md#PackagesValidateUploadRuby) | **Post** /packages/{owner}/{repo}/validate-upload/ruby/ | Validate parameters for create Ruby package
[**PackagesValidateUploadVagrant**](PackagesApi.md#PackagesValidateUploadVagrant) | **Post** /packages/{owner}/{repo}/validate-upload/vagrant/ | Validate parameters for create Vagrant package



## PackagesCopy

> PackageCopy PackagesCopy(ctx, owner, repo, identifier, optional)

Copy a package to another repository.

Copy a package to another repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***PackagesCopyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesCopyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**optional.Interface of PackagesCopy**](PackagesCopy.md)|  | 

### Return type

[**PackageCopy**](PackageCopy.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesDelete

> PackagesDelete(ctx, owner, repo, identifier)

Delete a specific package in a repository.

Delete a specific package in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesList

> []Package PackagesList(ctx, owner, repo, optional)

Get a list of all packages associated with repository.

Get a list of all packages associated with repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **query** | **optional.String**| A search term for querying names, filenames, versions, distributions, architectures, formats or statuses of packages. | 

### Return type

[**[]Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesMove

> PackageMove PackagesMove(ctx, owner, repo, identifier, optional)

Move a package to another repository.

Move a package to another repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***PackagesMoveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesMoveOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**optional.Interface of PackagesMove**](PackagesMove.md)|  | 

### Return type

[**PackageMove**](PackageMove.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesRead

> Package PackagesRead(ctx, owner, repo, identifier)

Get a specific package in a repository.

Get a specific package in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesResync

> Package PackagesResync(ctx, owner, repo, identifier, optional)

Schedule a package for resynchronisation.

Schedule a package for resynchronisation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***PackagesResyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesResyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**optional.Interface of PackagesResync**](PackagesResync.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesStatus

> PackageStatus PackagesStatus(ctx, owner, repo, identifier)

Get the synchronisation status for a package.

Get the synchronisation status for a package.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

[**PackageStatus**](PackageStatus.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadAlpine

> Package PackagesUploadAlpine(ctx, owner, repo, optional)

Create a new Alpine package

Create a new Alpine package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadAlpineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadAlpineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadAlpine**](PackagesUploadAlpine.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadCargo

> Package PackagesUploadCargo(ctx, owner, repo, optional)

Create a new Cargo package

Create a new Cargo package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadCargoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadCargoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadCargo**](PackagesUploadCargo.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadCocoapods

> Package PackagesUploadCocoapods(ctx, owner, repo, optional)

Create a new CocoaPods package

Create a new CocoaPods package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadCocoapodsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadCocoapodsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadCocoapods**](PackagesUploadCocoapods.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadComposer

> Package PackagesUploadComposer(ctx, owner, repo, optional)

Create a new Composer package

Create a new Composer package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadComposerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadComposerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadComposer**](PackagesUploadComposer.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadConan

> ConanPackageUpload PackagesUploadConan(ctx, owner, repo, optional)

Create a new Conan package

Create a new Conan package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadConanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadConanOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadConan**](PackagesUploadConan.md)|  | 

### Return type

[**ConanPackageUpload**](ConanPackageUpload.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadCran

> Package PackagesUploadCran(ctx, owner, repo, optional)

Create a new CRAN package

Create a new CRAN package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadCranOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadCranOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadCran**](PackagesUploadCran.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadDart

> Package PackagesUploadDart(ctx, owner, repo, optional)

Create a new Dart package

Create a new Dart package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadDartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadDartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadDart**](PackagesUploadDart.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadDeb

> Package PackagesUploadDeb(ctx, owner, repo, optional)

Create a new Debian package

Create a new Debian package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadDebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadDebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadDeb**](PackagesUploadDeb.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadDocker

> Package PackagesUploadDocker(ctx, owner, repo, optional)

Create a new Docker package

Create a new Docker package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadDockerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadDockerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadDocker**](PackagesUploadDocker.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadGo

> Package PackagesUploadGo(ctx, owner, repo, optional)

Create a new Go package

Create a new Go package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadGoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadGoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadGo**](PackagesUploadGo.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadHelm

> Package PackagesUploadHelm(ctx, owner, repo, optional)

Create a new Helm package

Create a new Helm package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadHelmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadHelmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadHelm**](PackagesUploadHelm.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadLuarocks

> Package PackagesUploadLuarocks(ctx, owner, repo, optional)

Create a new LuaRocks package

Create a new LuaRocks package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadLuarocksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadLuarocksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadLuarocks**](PackagesUploadLuarocks.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadMaven

> MavenPackageUpload PackagesUploadMaven(ctx, owner, repo, optional)

Create a new Maven package

Create a new Maven package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadMavenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadMavenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadMaven**](PackagesUploadMaven.md)|  | 

### Return type

[**MavenPackageUpload**](MavenPackageUpload.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadNpm

> Package PackagesUploadNpm(ctx, owner, repo, optional)

Create a new npm package

Create a new npm package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadNpmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadNpmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadNpm**](PackagesUploadNpm.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadNuget

> Package PackagesUploadNuget(ctx, owner, repo, optional)

Create a new NuGet package

Create a new NuGet package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadNugetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadNugetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadNuget**](PackagesUploadNuget.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadPython

> Package PackagesUploadPython(ctx, owner, repo, optional)

Create a new Python package

Create a new Python package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadPythonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadPythonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadPython**](PackagesUploadPython.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadRaw

> RawPackageUpload PackagesUploadRaw(ctx, owner, repo, optional)

Create a new Raw package

Create a new Raw package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadRawOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadRawOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadRaw**](PackagesUploadRaw.md)|  | 

### Return type

[**RawPackageUpload**](RawPackageUpload.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadRpm

> Package PackagesUploadRpm(ctx, owner, repo, optional)

Create a new RedHat package

Create a new RedHat package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadRpmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadRpmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadRpm**](PackagesUploadRpm.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadRuby

> Package PackagesUploadRuby(ctx, owner, repo, optional)

Create a new Ruby package

Create a new Ruby package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadRubyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadRubyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadRuby**](PackagesUploadRuby.md)|  | 

### Return type

[**Package**](Package.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesUploadVagrant

> VagrantPackageUpload PackagesUploadVagrant(ctx, owner, repo, optional)

Create a new Vagrant package

Create a new Vagrant package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesUploadVagrantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesUploadVagrantOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesUploadVagrant**](PackagesUploadVagrant.md)|  | 

### Return type

[**VagrantPackageUpload**](VagrantPackageUpload.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadAlpine

> PackagesValidateUploadAlpine(ctx, owner, repo, optional)

Validate parameters for create Alpine package

Validate parameters for create Alpine package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadAlpineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadAlpineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadAlpine**](PackagesValidateUploadAlpine.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadCargo

> PackagesValidateUploadCargo(ctx, owner, repo, optional)

Validate parameters for create Cargo package

Validate parameters for create Cargo package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadCargoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadCargoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadCargo**](PackagesValidateUploadCargo.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadCocoapods

> PackagesValidateUploadCocoapods(ctx, owner, repo, optional)

Validate parameters for create CocoaPods package

Validate parameters for create CocoaPods package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadCocoapodsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadCocoapodsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadCocoapods**](PackagesValidateUploadCocoapods.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadComposer

> PackagesValidateUploadComposer(ctx, owner, repo, optional)

Validate parameters for create Composer package

Validate parameters for create Composer package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadComposerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadComposerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadComposer**](PackagesValidateUploadComposer.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadConan

> PackagesValidateUploadConan(ctx, owner, repo, optional)

Validate parameters for create Conan package

Validate parameters for create Conan package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadConanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadConanOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadConan**](PackagesValidateUploadConan.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadCran

> PackagesValidateUploadCran(ctx, owner, repo, optional)

Validate parameters for create CRAN package

Validate parameters for create CRAN package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadCranOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadCranOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadCran**](PackagesValidateUploadCran.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadDart

> PackagesValidateUploadDart(ctx, owner, repo, optional)

Validate parameters for create Dart package

Validate parameters for create Dart package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadDartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadDartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadDart**](PackagesValidateUploadDart.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadDeb

> PackagesValidateUploadDeb(ctx, owner, repo, optional)

Validate parameters for create Debian package

Validate parameters for create Debian package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadDebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadDebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadDeb**](PackagesValidateUploadDeb.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadDocker

> PackagesValidateUploadDocker(ctx, owner, repo, optional)

Validate parameters for create Docker package

Validate parameters for create Docker package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadDockerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadDockerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadDocker**](PackagesValidateUploadDocker.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadGo

> PackagesValidateUploadGo(ctx, owner, repo, optional)

Validate parameters for create Go package

Validate parameters for create Go package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadGoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadGoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadGo**](PackagesValidateUploadGo.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadHelm

> PackagesValidateUploadHelm(ctx, owner, repo, optional)

Validate parameters for create Helm package

Validate parameters for create Helm package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadHelmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadHelmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadHelm**](PackagesValidateUploadHelm.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadLuarocks

> PackagesValidateUploadLuarocks(ctx, owner, repo, optional)

Validate parameters for create LuaRocks package

Validate parameters for create LuaRocks package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadLuarocksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadLuarocksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadLuarocks**](PackagesValidateUploadLuarocks.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadMaven

> PackagesValidateUploadMaven(ctx, owner, repo, optional)

Validate parameters for create Maven package

Validate parameters for create Maven package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadMavenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadMavenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadMaven**](PackagesValidateUploadMaven.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadNpm

> PackagesValidateUploadNpm(ctx, owner, repo, optional)

Validate parameters for create npm package

Validate parameters for create npm package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadNpmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadNpmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadNpm**](PackagesValidateUploadNpm.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadNuget

> PackagesValidateUploadNuget(ctx, owner, repo, optional)

Validate parameters for create NuGet package

Validate parameters for create NuGet package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadNugetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadNugetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadNuget**](PackagesValidateUploadNuget.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadPython

> PackagesValidateUploadPython(ctx, owner, repo, optional)

Validate parameters for create Python package

Validate parameters for create Python package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadPythonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadPythonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadPython**](PackagesValidateUploadPython.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadRaw

> PackagesValidateUploadRaw(ctx, owner, repo, optional)

Validate parameters for create Raw package

Validate parameters for create Raw package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadRawOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadRawOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadRaw**](PackagesValidateUploadRaw.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadRpm

> PackagesValidateUploadRpm(ctx, owner, repo, optional)

Validate parameters for create RedHat package

Validate parameters for create RedHat package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadRpmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadRpmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadRpm**](PackagesValidateUploadRpm.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadRuby

> PackagesValidateUploadRuby(ctx, owner, repo, optional)

Validate parameters for create Ruby package

Validate parameters for create Ruby package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadRubyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadRubyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadRuby**](PackagesValidateUploadRuby.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesValidateUploadVagrant

> PackagesValidateUploadVagrant(ctx, owner, repo, optional)

Validate parameters for create Vagrant package

Validate parameters for create Vagrant package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***PackagesValidateUploadVagrantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackagesValidateUploadVagrantOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PackagesValidateUploadVagrant**](PackagesValidateUploadVagrant.md)|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

