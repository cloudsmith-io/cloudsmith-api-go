# PackagesUploadRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here. | [optional] 
**Description** | Pointer to **string** | A textual description of this package. | [optional] 
**Name** | Pointer to **string** | The name of this package. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Summary** | Pointer to **string** | A one-liner synopsis of this package. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | Pointer to **string** | The raw version for this package. | [optional] 

## Methods

### NewPackagesUploadRaw

`func NewPackagesUploadRaw(packageFile string, ) *PackagesUploadRaw`

NewPackagesUploadRaw instantiates a new PackagesUploadRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadRawWithDefaults

`func NewPackagesUploadRawWithDefaults() *PackagesUploadRaw`

NewPackagesUploadRawWithDefaults instantiates a new PackagesUploadRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *PackagesUploadRaw) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PackagesUploadRaw) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PackagesUploadRaw) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PackagesUploadRaw) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDescription

`func (o *PackagesUploadRaw) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackagesUploadRaw) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackagesUploadRaw) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackagesUploadRaw) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PackagesUploadRaw) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackagesUploadRaw) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackagesUploadRaw) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackagesUploadRaw) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesUploadRaw) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadRaw) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadRaw) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesUploadRaw) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadRaw) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadRaw) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadRaw) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSummary

`func (o *PackagesUploadRaw) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PackagesUploadRaw) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PackagesUploadRaw) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PackagesUploadRaw) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadRaw) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadRaw) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadRaw) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadRaw) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesUploadRaw) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesUploadRaw) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesUploadRaw) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackagesUploadRaw) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


