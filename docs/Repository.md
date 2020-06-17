# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdnUrl** | **string** |  | [optional] 
**CreatedAt** | **string** |  | [optional] 
**DeletedAt** | **string** |  | [optional] 
**Description** | **string** |  | [optional] 
**GpgKeys** | [**[]ReposGpgKeys**](_repos__gpg_keys.md) |  | [optional] 
**IndexFiles** | **bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**IsOpenSource** | **bool** |  | [optional] 
**IsPrivate** | **bool** |  | [optional] 
**IsPublic** | **bool** |  | [optional] 
**Name** | **string** | A descriptive name for the repository. | 
**Namespace** | **string** |  | [optional] 
**NamespaceUrl** | **string** |  | [optional] 
**NumDownloads** | **int64** | The number of downloads for packages in the repository. | [optional] 
**PackageCount** | **int64** | The number of packages in the repository. | [optional] 
**PackageGroupCount** | **int64** | The number of groups in the repository. | [optional] 
**RepositoryType** | **int64** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Open-Source repositories are always visible to everyone and are restricted by licensing, but are free to use and come with generous bandwidth/storage. You can only select Open-Source at repository creation time. | [optional] 
**RepositoryTypeStr** | **string** |  | [optional] 
**SelfHtmlUrl** | **string** |  | [optional] 
**SelfUrl** | **string** |  | [optional] 
**Size** | **int64** | The calculated size of the repository. | [optional] 
**SizeStr** | **string** |  | [optional] 
**Slug** | **string** |  | [optional] 
**SlugPerm** | **string** |  | [optional] 
**StorageRegion** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


