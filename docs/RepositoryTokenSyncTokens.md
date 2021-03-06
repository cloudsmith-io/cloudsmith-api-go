# RepositoryTokenSyncTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The datetime the token was updated at. | [optional] 
**CreatedBy** | **string** |  | [optional] 
**CreatedByUrl** | **string** |  | [optional] 
**Default** | **bool** | If selected this is the default token for this repository. | [optional] 
**HasLimits** | **bool** |  | [optional] 
**Identifier** | **int64** |  | [optional] 
**IsActive** | **\*bool** | If enabled, the token will allow downloads based on configured restrictions (if any). | [optional] 
**IsLimited** | **bool** |  | [optional] 
**LimitDateRangeFrom** | **string** | The starting date/time the token is allowed to be used from. | [optional] 
**LimitDateRangeTo** | **string** | The ending date/time the token is allowed to be used until. | [optional] 
**LimitNumClients** | **int64** | The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitNumDownloads** | **int64** | The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitPackageQuery** | **string** | The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. For package formats that support dynamic metadata indexes, the contents of the metadata will also be filtered. | [optional] 
**LimitPathQuery** | **string** | The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash. | [optional] 
**Metadata** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**RefreshUrl** | **string** |  | [optional] 
**SelfUrl** | **string** |  | [optional] 
**SlugPerm** | **string** |  | [optional] 
**Token** | **string** |  | [optional] 
**UpdatedAt** | **string** | The datetime the token was updated at. | [optional] 
**UpdatedBy** | **string** |  | [optional] 
**UpdatedByUrl** | **string** |  | [optional] 
**User** | **string** |  | [optional] 
**UserUrl** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


