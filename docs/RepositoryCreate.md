# RepositoryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdnUrl** | Pointer to **string** | Base URL from which packages and other artifacts are downloaded. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | A description of the repository&#39;s purpose/contents. | [optional] 
**GpgKeys** | Pointer to [**[]ReposGpgKeys**](ReposGpgKeys.md) |  | [optional] 
**IndexFiles** | Pointer to **bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**IsOpenSource** | Pointer to **bool** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | **string** | A descriptive name for the repository. | 
**Namespace** | Pointer to **string** | Namespace to which this repository belongs. | [optional] 
**NamespaceUrl** | Pointer to **string** | API endpoint where data about this namespace can be retrieved. | [optional] 
**NumDownloads** | Pointer to **int64** | The number of downloads for packages in the repository. | [optional] 
**PackageCount** | Pointer to **int64** | The number of packages in the repository. | [optional] 
**PackageGroupCount** | Pointer to **int64** | The number of groups in the repository. | [optional] 
**RepositoryType** | Pointer to **int64** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Open-Source repositories are always visible to everyone and are restricted by licensing, but are free to use and come with generous bandwidth/storage. You can only select Open-Source at repository creation time. | [optional] 
**RepositoryTypeStr** | Pointer to **string** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users. | [optional] 
**SelfHtmlUrl** | Pointer to **string** | Website URL for this repository. | [optional] 
**SelfUrl** | Pointer to **string** | API endpoint where data about this repository can be retrieved. | [optional] 
**Size** | Pointer to **int64** | The calculated size of the repository. | [optional] 
**SizeStr** | Pointer to **string** | The calculated size of the repository (human readable). | [optional] 
**Slug** | Pointer to **string** | The slug identifies the repository in URIs. | [optional] 
**SlugPerm** | Pointer to **string** | The slug_perm immutably identifies the repository. It will never change once a repository has been created. | [optional] 
**StorageRegion** | Pointer to **string** | The Cloudsmith region in which package files are stored. | [optional] 

## Methods

### NewRepositoryCreate

`func NewRepositoryCreate(name string, ) *RepositoryCreate`

NewRepositoryCreate instantiates a new RepositoryCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryCreateWithDefaults

`func NewRepositoryCreateWithDefaults() *RepositoryCreate`

NewRepositoryCreateWithDefaults instantiates a new RepositoryCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdnUrl

`func (o *RepositoryCreate) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *RepositoryCreate) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *RepositoryCreate) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *RepositoryCreate) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RepositoryCreate) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RepositoryCreate) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RepositoryCreate) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RepositoryCreate) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *RepositoryCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepositoryCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGpgKeys

`func (o *RepositoryCreate) GetGpgKeys() []ReposGpgKeys`

GetGpgKeys returns the GpgKeys field if non-nil, zero value otherwise.

### GetGpgKeysOk

`func (o *RepositoryCreate) GetGpgKeysOk() (*[]ReposGpgKeys, bool)`

GetGpgKeysOk returns a tuple with the GpgKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeys

`func (o *RepositoryCreate) SetGpgKeys(v []ReposGpgKeys)`

SetGpgKeys sets GpgKeys field to given value.

### HasGpgKeys

`func (o *RepositoryCreate) HasGpgKeys() bool`

HasGpgKeys returns a boolean if a field has been set.

### GetIndexFiles

`func (o *RepositoryCreate) GetIndexFiles() bool`

GetIndexFiles returns the IndexFiles field if non-nil, zero value otherwise.

### GetIndexFilesOk

`func (o *RepositoryCreate) GetIndexFilesOk() (*bool, bool)`

GetIndexFilesOk returns a tuple with the IndexFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFiles

`func (o *RepositoryCreate) SetIndexFiles(v bool)`

SetIndexFiles sets IndexFiles field to given value.

### HasIndexFiles

`func (o *RepositoryCreate) HasIndexFiles() bool`

HasIndexFiles returns a boolean if a field has been set.

### GetIsOpenSource

`func (o *RepositoryCreate) GetIsOpenSource() bool`

GetIsOpenSource returns the IsOpenSource field if non-nil, zero value otherwise.

### GetIsOpenSourceOk

`func (o *RepositoryCreate) GetIsOpenSourceOk() (*bool, bool)`

GetIsOpenSourceOk returns a tuple with the IsOpenSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenSource

`func (o *RepositoryCreate) SetIsOpenSource(v bool)`

SetIsOpenSource sets IsOpenSource field to given value.

### HasIsOpenSource

`func (o *RepositoryCreate) HasIsOpenSource() bool`

HasIsOpenSource returns a boolean if a field has been set.

### GetIsPrivate

`func (o *RepositoryCreate) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *RepositoryCreate) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *RepositoryCreate) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *RepositoryCreate) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsPublic

`func (o *RepositoryCreate) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *RepositoryCreate) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *RepositoryCreate) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *RepositoryCreate) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *RepositoryCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryCreate) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *RepositoryCreate) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RepositoryCreate) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RepositoryCreate) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RepositoryCreate) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *RepositoryCreate) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *RepositoryCreate) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *RepositoryCreate) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *RepositoryCreate) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumDownloads

`func (o *RepositoryCreate) GetNumDownloads() int64`

GetNumDownloads returns the NumDownloads field if non-nil, zero value otherwise.

### GetNumDownloadsOk

`func (o *RepositoryCreate) GetNumDownloadsOk() (*int64, bool)`

GetNumDownloadsOk returns a tuple with the NumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDownloads

`func (o *RepositoryCreate) SetNumDownloads(v int64)`

SetNumDownloads sets NumDownloads field to given value.

### HasNumDownloads

`func (o *RepositoryCreate) HasNumDownloads() bool`

HasNumDownloads returns a boolean if a field has been set.

### GetPackageCount

`func (o *RepositoryCreate) GetPackageCount() int64`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *RepositoryCreate) GetPackageCountOk() (*int64, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *RepositoryCreate) SetPackageCount(v int64)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *RepositoryCreate) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetPackageGroupCount

`func (o *RepositoryCreate) GetPackageGroupCount() int64`

GetPackageGroupCount returns the PackageGroupCount field if non-nil, zero value otherwise.

### GetPackageGroupCountOk

`func (o *RepositoryCreate) GetPackageGroupCountOk() (*int64, bool)`

GetPackageGroupCountOk returns a tuple with the PackageGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageGroupCount

`func (o *RepositoryCreate) SetPackageGroupCount(v int64)`

SetPackageGroupCount sets PackageGroupCount field to given value.

### HasPackageGroupCount

`func (o *RepositoryCreate) HasPackageGroupCount() bool`

HasPackageGroupCount returns a boolean if a field has been set.

### GetRepositoryType

`func (o *RepositoryCreate) GetRepositoryType() int64`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *RepositoryCreate) GetRepositoryTypeOk() (*int64, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *RepositoryCreate) SetRepositoryType(v int64)`

SetRepositoryType sets RepositoryType field to given value.

### HasRepositoryType

`func (o *RepositoryCreate) HasRepositoryType() bool`

HasRepositoryType returns a boolean if a field has been set.

### GetRepositoryTypeStr

`func (o *RepositoryCreate) GetRepositoryTypeStr() string`

GetRepositoryTypeStr returns the RepositoryTypeStr field if non-nil, zero value otherwise.

### GetRepositoryTypeStrOk

`func (o *RepositoryCreate) GetRepositoryTypeStrOk() (*string, bool)`

GetRepositoryTypeStrOk returns a tuple with the RepositoryTypeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryTypeStr

`func (o *RepositoryCreate) SetRepositoryTypeStr(v string)`

SetRepositoryTypeStr sets RepositoryTypeStr field to given value.

### HasRepositoryTypeStr

`func (o *RepositoryCreate) HasRepositoryTypeStr() bool`

HasRepositoryTypeStr returns a boolean if a field has been set.

### GetSelfHtmlUrl

`func (o *RepositoryCreate) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *RepositoryCreate) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *RepositoryCreate) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *RepositoryCreate) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *RepositoryCreate) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RepositoryCreate) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RepositoryCreate) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RepositoryCreate) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSize

`func (o *RepositoryCreate) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RepositoryCreate) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RepositoryCreate) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RepositoryCreate) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeStr

`func (o *RepositoryCreate) GetSizeStr() string`

GetSizeStr returns the SizeStr field if non-nil, zero value otherwise.

### GetSizeStrOk

`func (o *RepositoryCreate) GetSizeStrOk() (*string, bool)`

GetSizeStrOk returns a tuple with the SizeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeStr

`func (o *RepositoryCreate) SetSizeStr(v string)`

SetSizeStr sets SizeStr field to given value.

### HasSizeStr

`func (o *RepositoryCreate) HasSizeStr() bool`

HasSizeStr returns a boolean if a field has been set.

### GetSlug

`func (o *RepositoryCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RepositoryCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RepositoryCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *RepositoryCreate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RepositoryCreate) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RepositoryCreate) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RepositoryCreate) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RepositoryCreate) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStorageRegion

`func (o *RepositoryCreate) GetStorageRegion() string`

GetStorageRegion returns the StorageRegion field if non-nil, zero value otherwise.

### GetStorageRegionOk

`func (o *RepositoryCreate) GetStorageRegionOk() (*string, bool)`

GetStorageRegionOk returns a tuple with the StorageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRegion

`func (o *RepositoryCreate) SetStorageRegion(v string)`

SetStorageRegion sets StorageRegion field to given value.

### HasStorageRegion

`func (o *RepositoryCreate) HasStorageRegion() bool`

HasStorageRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


