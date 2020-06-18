# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdnUrl** | **string** | Base URL from which packages and other artifacts are downloaded. | [optional] 
**CreatedAt** | **string** |  | [optional] 
**DeletedAt** | **string** |  | [optional] 
**Description** | **string** | A description of the repository&#39;s purpose/contents. | [optional] 
**GpgKeys** | [**[]ReposGpgKeys**](_repos__gpg_keys.md) |  | [optional] 
**IndexFiles** | **\*bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**IsOpenSource** | **bool** |  | [optional] 
**IsPrivate** | **bool** |  | [optional] 
**IsPublic** | **bool** |  | [optional] 
**Name** | **string** | A descriptive name for the repository. | 
**Namespace** | **string** | Namespace to which this repository belongs. | [optional] 
**NamespaceUrl** | **string** | API endpoint where data about this namespace can be retrieved. | [optional] 
**NumDownloads** | **int64** | The number of downloads for packages in the repository. | [optional] 
**PackageCount** | **int64** | The number of packages in the repository. | [optional] 
**PackageGroupCount** | **int64** | The number of groups in the repository. | [optional] 
**RepositoryType** | **int64** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Open-Source repositories are always visible to everyone and are restricted by licensing, but are free to use and come with generous bandwidth/storage. You can only select Open-Source at repository creation time. | [optional] 
**RepositoryTypeStr** | **string** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users. | [optional] 
**SelfHtmlUrl** | **string** | Website URL for this repository. | [optional] 
**SelfUrl** | **string** | API endpoint where data about this repository can be retrieved. | [optional] 
**Size** | **int64** | The calculated size of the repository. | [optional] 
**SizeStr** | **string** | The calculated size of the repository (human readable). | [optional] 
**Slug** | **string** | The slug identifies the repository in URIs. | [optional] 
**SlugPerm** | **string** | The slug_perm immutably identifies the repository. It will never change once a repository has been created. | [optional] 
**StorageRegion** | **string** | The Cloudsmith region in which package files are stored. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


