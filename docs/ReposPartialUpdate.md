# ReposPartialUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the repository&#39;s purpose/contents. | [optional] 
**IndexFiles** | Pointer to **bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**Name** | Pointer to **string** | A descriptive name for the repository. | [optional] 
**RepositoryTypeStr** | Pointer to **string** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users. | [optional] 
**Slug** | Pointer to **string** | The slug identifies the repository in URIs. | [optional] 

## Methods

### NewReposPartialUpdate

`func NewReposPartialUpdate() *ReposPartialUpdate`

NewReposPartialUpdate instantiates a new ReposPartialUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposPartialUpdateWithDefaults

`func NewReposPartialUpdateWithDefaults() *ReposPartialUpdate`

NewReposPartialUpdateWithDefaults instantiates a new ReposPartialUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReposPartialUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposPartialUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposPartialUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposPartialUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIndexFiles

`func (o *ReposPartialUpdate) GetIndexFiles() bool`

GetIndexFiles returns the IndexFiles field if non-nil, zero value otherwise.

### GetIndexFilesOk

`func (o *ReposPartialUpdate) GetIndexFilesOk() (*bool, bool)`

GetIndexFilesOk returns a tuple with the IndexFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFiles

`func (o *ReposPartialUpdate) SetIndexFiles(v bool)`

SetIndexFiles sets IndexFiles field to given value.

### HasIndexFiles

`func (o *ReposPartialUpdate) HasIndexFiles() bool`

HasIndexFiles returns a boolean if a field has been set.

### GetName

`func (o *ReposPartialUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposPartialUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposPartialUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposPartialUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepositoryTypeStr

`func (o *ReposPartialUpdate) GetRepositoryTypeStr() string`

GetRepositoryTypeStr returns the RepositoryTypeStr field if non-nil, zero value otherwise.

### GetRepositoryTypeStrOk

`func (o *ReposPartialUpdate) GetRepositoryTypeStrOk() (*string, bool)`

GetRepositoryTypeStrOk returns a tuple with the RepositoryTypeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryTypeStr

`func (o *ReposPartialUpdate) SetRepositoryTypeStr(v string)`

SetRepositoryTypeStr sets RepositoryTypeStr field to given value.

### HasRepositoryTypeStr

`func (o *ReposPartialUpdate) HasRepositoryTypeStr() bool`

HasRepositoryTypeStr returns a boolean if a field has been set.

### GetSlug

`func (o *ReposPartialUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ReposPartialUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ReposPartialUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ReposPartialUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


