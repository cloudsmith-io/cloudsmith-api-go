# SwiftPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorName** | Pointer to **string** | The name of the author of the package. | [optional] 
**AuthorOrg** | Pointer to **string** | The organization of the author. | [optional] 
**LicenseUrl** | Pointer to **NullableString** | The license URL of this package. | [optional] 
**Name** | **string** | The name of this package. | 
**PackageFile** | **string** | The primary file for the package. | 
**ReadmeUrl** | Pointer to **string** | The URL of the readme for the package. | [optional] 
**RepositoryUrl** | Pointer to **string** | The URL of the SCM repository for the package. | [optional] 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Scope** | **string** | A scope provides a namespace for related packages within the package registry. | 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | **string** | The raw version for this package. | 

## Methods

### NewSwiftPackageUploadRequest

`func NewSwiftPackageUploadRequest(name string, packageFile string, scope string, version string, ) *SwiftPackageUploadRequest`

NewSwiftPackageUploadRequest instantiates a new SwiftPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwiftPackageUploadRequestWithDefaults

`func NewSwiftPackageUploadRequestWithDefaults() *SwiftPackageUploadRequest`

NewSwiftPackageUploadRequestWithDefaults instantiates a new SwiftPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorName

`func (o *SwiftPackageUploadRequest) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *SwiftPackageUploadRequest) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *SwiftPackageUploadRequest) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *SwiftPackageUploadRequest) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetAuthorOrg

`func (o *SwiftPackageUploadRequest) GetAuthorOrg() string`

GetAuthorOrg returns the AuthorOrg field if non-nil, zero value otherwise.

### GetAuthorOrgOk

`func (o *SwiftPackageUploadRequest) GetAuthorOrgOk() (*string, bool)`

GetAuthorOrgOk returns a tuple with the AuthorOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorOrg

`func (o *SwiftPackageUploadRequest) SetAuthorOrg(v string)`

SetAuthorOrg sets AuthorOrg field to given value.

### HasAuthorOrg

`func (o *SwiftPackageUploadRequest) HasAuthorOrg() bool`

HasAuthorOrg returns a boolean if a field has been set.

### GetLicenseUrl

`func (o *SwiftPackageUploadRequest) GetLicenseUrl() string`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *SwiftPackageUploadRequest) GetLicenseUrlOk() (*string, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *SwiftPackageUploadRequest) SetLicenseUrl(v string)`

SetLicenseUrl sets LicenseUrl field to given value.

### HasLicenseUrl

`func (o *SwiftPackageUploadRequest) HasLicenseUrl() bool`

HasLicenseUrl returns a boolean if a field has been set.

### SetLicenseUrlNil

`func (o *SwiftPackageUploadRequest) SetLicenseUrlNil(b bool)`

 SetLicenseUrlNil sets the value for LicenseUrl to be an explicit nil

### UnsetLicenseUrl
`func (o *SwiftPackageUploadRequest) UnsetLicenseUrl()`

UnsetLicenseUrl ensures that no value is present for LicenseUrl, not even an explicit nil
### GetName

`func (o *SwiftPackageUploadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwiftPackageUploadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwiftPackageUploadRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPackageFile

`func (o *SwiftPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *SwiftPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *SwiftPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetReadmeUrl

`func (o *SwiftPackageUploadRequest) GetReadmeUrl() string`

GetReadmeUrl returns the ReadmeUrl field if non-nil, zero value otherwise.

### GetReadmeUrlOk

`func (o *SwiftPackageUploadRequest) GetReadmeUrlOk() (*string, bool)`

GetReadmeUrlOk returns a tuple with the ReadmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadmeUrl

`func (o *SwiftPackageUploadRequest) SetReadmeUrl(v string)`

SetReadmeUrl sets ReadmeUrl field to given value.

### HasReadmeUrl

`func (o *SwiftPackageUploadRequest) HasReadmeUrl() bool`

HasReadmeUrl returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *SwiftPackageUploadRequest) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *SwiftPackageUploadRequest) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *SwiftPackageUploadRequest) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *SwiftPackageUploadRequest) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetRepublish

`func (o *SwiftPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *SwiftPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *SwiftPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *SwiftPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetScope

`func (o *SwiftPackageUploadRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SwiftPackageUploadRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SwiftPackageUploadRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTags

`func (o *SwiftPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SwiftPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SwiftPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SwiftPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SwiftPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SwiftPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVersion

`func (o *SwiftPackageUploadRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SwiftPackageUploadRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SwiftPackageUploadRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


