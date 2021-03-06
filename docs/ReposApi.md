# \ReposApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposAllList**](ReposApi.md#ReposAllList) | **Get** /repos/ | Get a list of all repositories associated with current user.
[**ReposCreate**](ReposApi.md#ReposCreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
[**ReposList**](ReposApi.md#ReposList) | **Get** /repos/{owner}/ | Get a list of all repositories within a namespace.
[**ReposPartialUpdate**](ReposApi.md#ReposPartialUpdate) | **Patch** /repos/{owner}/{identifier}/ | Update details about a repository in a given namespace.
[**ReposRead**](ReposApi.md#ReposRead) | **Get** /repos/{owner}/{identifier}/ | Get a specific repository.



## ReposAllList

> []Repository ReposAllList(ctx, optional)

Get a list of all repositories associated with current user.

Get a list of all repositories associated with current user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReposAllListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReposAllListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int64**| Number of results to return per page. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreate

> Repository ReposCreate(ctx, owner, optional)

Create a new repository in a given namespace.

Create a new repository in a given namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
 **optional** | ***ReposCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReposCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of ReposCreate**](ReposCreate.md)|  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposDelete

> ReposDelete(ctx, owner, identifier)

Delete a repository in a given namespace.

Delete a repository in a given namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
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


## ReposList

> []Repository ReposList(ctx, owner, optional)

Get a list of all repositories within a namespace.

Get a list of all repositories within a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
 **optional** | ***ReposListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReposListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int64**| Number of results to return per page. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPartialUpdate

> Repository ReposPartialUpdate(ctx, owner, identifier, optional)

Update details about a repository in a given namespace.

Update details about a repository in a given namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***ReposPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReposPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of ReposPartialUpdate**](ReposPartialUpdate.md)|  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRead

> Repository ReposRead(ctx, owner, identifier)

Get a specific repository.

Get a specific repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**identifier** | **string**|  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

