# RawPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **NullableString** | A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here. | [optional] 
**Description** | Pointer to **NullableString** | A textual description of this package. | [optional] 
**Name** | Pointer to **NullableString** | The name of this package. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Summary** | Pointer to **NullableString** | A one-liner synopsis of this package. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] 

## Methods

### NewRawPackageUploadRequest

`func NewRawPackageUploadRequest(packageFile string, ) *RawPackageUploadRequest`

NewRawPackageUploadRequest instantiates a new RawPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawPackageUploadRequestWithDefaults

`func NewRawPackageUploadRequestWithDefaults() *RawPackageUploadRequest`

NewRawPackageUploadRequestWithDefaults instantiates a new RawPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *RawPackageUploadRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *RawPackageUploadRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *RawPackageUploadRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *RawPackageUploadRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *RawPackageUploadRequest) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *RawPackageUploadRequest) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDescription

`func (o *RawPackageUploadRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RawPackageUploadRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RawPackageUploadRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RawPackageUploadRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RawPackageUploadRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RawPackageUploadRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *RawPackageUploadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawPackageUploadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawPackageUploadRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RawPackageUploadRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RawPackageUploadRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RawPackageUploadRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackageFile

`func (o *RawPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *RawPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *RawPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *RawPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *RawPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *RawPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *RawPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSummary

`func (o *RawPackageUploadRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RawPackageUploadRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RawPackageUploadRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RawPackageUploadRequest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *RawPackageUploadRequest) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *RawPackageUploadRequest) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetTags

`func (o *RawPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RawPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RawPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RawPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RawPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RawPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVersion

`func (o *RawPackageUploadRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RawPackageUploadRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RawPackageUploadRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RawPackageUploadRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RawPackageUploadRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RawPackageUploadRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


