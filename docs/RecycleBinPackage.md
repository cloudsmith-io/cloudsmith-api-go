# RecycleBinPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionBy** | Pointer to **NullableString** | The name of the user who deleted the package. | [optional] [readonly] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**Filename** | Pointer to **string** |  | [optional] [readonly] 
**Format** | Pointer to **string** |  | [optional] [readonly] 
**FullyQualifiedName** | Pointer to **string** | The fully qualified name of the package. | [optional] [readonly] 
**Identifiers** | Pointer to **map[string]string** |  | [optional] [readonly] 
**InvokedRetentionRule** | Pointer to **map[string]string** | Information about the retention rule that triggered deletion (if any). | [optional] [readonly] 
**IsDeleteable** | Pointer to **bool** |  | [optional] [readonly] 
**IsQuarantined** | Pointer to **bool** |  | [optional] [readonly] 
**IsRestorable** | Pointer to **bool** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The name of this package. | [optional] [readonly] 
**PolicyViolated** | Pointer to **bool** | Whether or not the package has violated any policy. | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**SecurityScanCompletedAt** | Pointer to **NullableTime** | The datetime the security scanning was completed. | [optional] [readonly] 
**SecurityScanStatus** | Pointer to **NullableString** |  | [optional] [readonly] [default to "Awaiting Security Scan"]
**Size** | Pointer to **int64** | The calculated size of the package. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **int64** | The synchronisation status of the package. | [optional] [readonly] 
**StatusUpdatedAt** | Pointer to **time.Time** | The datetime the package status was updated at. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TypeDisplay** | Pointer to **string** |  | [optional] [readonly] 
**UploadedAt** | Pointer to **time.Time** | The date this package was uploaded. | [optional] [readonly] 
**Uploader** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] [readonly] 

## Methods

### NewRecycleBinPackage

`func NewRecycleBinPackage() *RecycleBinPackage`

NewRecycleBinPackage instantiates a new RecycleBinPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecycleBinPackageWithDefaults

`func NewRecycleBinPackageWithDefaults() *RecycleBinPackage`

NewRecycleBinPackageWithDefaults instantiates a new RecycleBinPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionBy

`func (o *RecycleBinPackage) GetActionBy() string`

GetActionBy returns the ActionBy field if non-nil, zero value otherwise.

### GetActionByOk

`func (o *RecycleBinPackage) GetActionByOk() (*string, bool)`

GetActionByOk returns a tuple with the ActionBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionBy

`func (o *RecycleBinPackage) SetActionBy(v string)`

SetActionBy sets ActionBy field to given value.

### HasActionBy

`func (o *RecycleBinPackage) HasActionBy() bool`

HasActionBy returns a boolean if a field has been set.

### SetActionByNil

`func (o *RecycleBinPackage) SetActionByNil(b bool)`

 SetActionByNil sets the value for ActionBy to be an explicit nil

### UnsetActionBy
`func (o *RecycleBinPackage) UnsetActionBy()`

UnsetActionBy ensures that no value is present for ActionBy, not even an explicit nil
### GetDownloads

`func (o *RecycleBinPackage) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *RecycleBinPackage) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *RecycleBinPackage) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *RecycleBinPackage) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetFilename

`func (o *RecycleBinPackage) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RecycleBinPackage) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RecycleBinPackage) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *RecycleBinPackage) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFormat

`func (o *RecycleBinPackage) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RecycleBinPackage) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RecycleBinPackage) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RecycleBinPackage) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *RecycleBinPackage) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *RecycleBinPackage) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *RecycleBinPackage) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *RecycleBinPackage) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetIdentifiers

`func (o *RecycleBinPackage) GetIdentifiers() map[string]string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *RecycleBinPackage) GetIdentifiersOk() (*map[string]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *RecycleBinPackage) SetIdentifiers(v map[string]string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *RecycleBinPackage) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetInvokedRetentionRule

`func (o *RecycleBinPackage) GetInvokedRetentionRule() map[string]string`

GetInvokedRetentionRule returns the InvokedRetentionRule field if non-nil, zero value otherwise.

### GetInvokedRetentionRuleOk

`func (o *RecycleBinPackage) GetInvokedRetentionRuleOk() (*map[string]string, bool)`

GetInvokedRetentionRuleOk returns a tuple with the InvokedRetentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokedRetentionRule

`func (o *RecycleBinPackage) SetInvokedRetentionRule(v map[string]string)`

SetInvokedRetentionRule sets InvokedRetentionRule field to given value.

### HasInvokedRetentionRule

`func (o *RecycleBinPackage) HasInvokedRetentionRule() bool`

HasInvokedRetentionRule returns a boolean if a field has been set.

### SetInvokedRetentionRuleNil

`func (o *RecycleBinPackage) SetInvokedRetentionRuleNil(b bool)`

 SetInvokedRetentionRuleNil sets the value for InvokedRetentionRule to be an explicit nil

### UnsetInvokedRetentionRule
`func (o *RecycleBinPackage) UnsetInvokedRetentionRule()`

UnsetInvokedRetentionRule ensures that no value is present for InvokedRetentionRule, not even an explicit nil
### GetIsDeleteable

`func (o *RecycleBinPackage) GetIsDeleteable() bool`

GetIsDeleteable returns the IsDeleteable field if non-nil, zero value otherwise.

### GetIsDeleteableOk

`func (o *RecycleBinPackage) GetIsDeleteableOk() (*bool, bool)`

GetIsDeleteableOk returns a tuple with the IsDeleteable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteable

`func (o *RecycleBinPackage) SetIsDeleteable(v bool)`

SetIsDeleteable sets IsDeleteable field to given value.

### HasIsDeleteable

`func (o *RecycleBinPackage) HasIsDeleteable() bool`

HasIsDeleteable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *RecycleBinPackage) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *RecycleBinPackage) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *RecycleBinPackage) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *RecycleBinPackage) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsRestorable

`func (o *RecycleBinPackage) GetIsRestorable() bool`

GetIsRestorable returns the IsRestorable field if non-nil, zero value otherwise.

### GetIsRestorableOk

`func (o *RecycleBinPackage) GetIsRestorableOk() (*bool, bool)`

GetIsRestorableOk returns a tuple with the IsRestorable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestorable

`func (o *RecycleBinPackage) SetIsRestorable(v bool)`

SetIsRestorable sets IsRestorable field to given value.

### HasIsRestorable

`func (o *RecycleBinPackage) HasIsRestorable() bool`

HasIsRestorable returns a boolean if a field has been set.

### GetName

`func (o *RecycleBinPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecycleBinPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecycleBinPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecycleBinPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RecycleBinPackage) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RecycleBinPackage) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPolicyViolated

`func (o *RecycleBinPackage) GetPolicyViolated() bool`

GetPolicyViolated returns the PolicyViolated field if non-nil, zero value otherwise.

### GetPolicyViolatedOk

`func (o *RecycleBinPackage) GetPolicyViolatedOk() (*bool, bool)`

GetPolicyViolatedOk returns a tuple with the PolicyViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolated

`func (o *RecycleBinPackage) SetPolicyViolated(v bool)`

SetPolicyViolated sets PolicyViolated field to given value.

### HasPolicyViolated

`func (o *RecycleBinPackage) HasPolicyViolated() bool`

HasPolicyViolated returns a boolean if a field has been set.

### GetRepository

`func (o *RecycleBinPackage) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RecycleBinPackage) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RecycleBinPackage) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RecycleBinPackage) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *RecycleBinPackage) GetSecurityScanCompletedAt() time.Time`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *RecycleBinPackage) GetSecurityScanCompletedAtOk() (*time.Time, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *RecycleBinPackage) SetSecurityScanCompletedAt(v time.Time)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *RecycleBinPackage) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### SetSecurityScanCompletedAtNil

`func (o *RecycleBinPackage) SetSecurityScanCompletedAtNil(b bool)`

 SetSecurityScanCompletedAtNil sets the value for SecurityScanCompletedAt to be an explicit nil

### UnsetSecurityScanCompletedAt
`func (o *RecycleBinPackage) UnsetSecurityScanCompletedAt()`

UnsetSecurityScanCompletedAt ensures that no value is present for SecurityScanCompletedAt, not even an explicit nil
### GetSecurityScanStatus

`func (o *RecycleBinPackage) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *RecycleBinPackage) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *RecycleBinPackage) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *RecycleBinPackage) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### SetSecurityScanStatusNil

`func (o *RecycleBinPackage) SetSecurityScanStatusNil(b bool)`

 SetSecurityScanStatusNil sets the value for SecurityScanStatus to be an explicit nil

### UnsetSecurityScanStatus
`func (o *RecycleBinPackage) UnsetSecurityScanStatus()`

UnsetSecurityScanStatus ensures that no value is present for SecurityScanStatus, not even an explicit nil
### GetSize

`func (o *RecycleBinPackage) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RecycleBinPackage) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RecycleBinPackage) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RecycleBinPackage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RecycleBinPackage) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RecycleBinPackage) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RecycleBinPackage) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RecycleBinPackage) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStatus

`func (o *RecycleBinPackage) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecycleBinPackage) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecycleBinPackage) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecycleBinPackage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *RecycleBinPackage) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *RecycleBinPackage) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *RecycleBinPackage) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *RecycleBinPackage) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *RecycleBinPackage) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecycleBinPackage) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecycleBinPackage) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecycleBinPackage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *RecycleBinPackage) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *RecycleBinPackage) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *RecycleBinPackage) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *RecycleBinPackage) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *RecycleBinPackage) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *RecycleBinPackage) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *RecycleBinPackage) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *RecycleBinPackage) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *RecycleBinPackage) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *RecycleBinPackage) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *RecycleBinPackage) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *RecycleBinPackage) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetVersion

`func (o *RecycleBinPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RecycleBinPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RecycleBinPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RecycleBinPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RecycleBinPackage) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RecycleBinPackage) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


