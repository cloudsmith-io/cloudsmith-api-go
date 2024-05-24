# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architectures** | Pointer to [**[]Architecture**](Architecture.md) |  | [optional] [readonly] 
**CdnUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumMd5** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha1** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha256** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha512** | Pointer to **string** |  | [optional] [readonly] 
**DependenciesChecksumMd5** | Pointer to **NullableString** | A checksum of all of the package&#39;s dependencies. | [optional] [readonly] 
**DependenciesUrl** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | A textual description of this package. | [optional] [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**Distro** | Pointer to [**NullableDistribution**](Distribution.md) |  | [optional] 
**DistroVersion** | Pointer to [**DistributionVersion**](DistributionVersion.md) |  | [optional] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**Epoch** | Pointer to **NullableInt64** | The epoch of the package version (if any). | [optional] [readonly] 
**Extension** | Pointer to **string** |  | [optional] [readonly] 
**Filename** | Pointer to **string** |  | [optional] [readonly] 
**Files** | Pointer to [**[]PackageFile**](PackageFile.md) |  | [optional] [readonly] 
**Format** | Pointer to **string** |  | [optional] [readonly] 
**FormatUrl** | Pointer to **string** |  | [optional] [readonly] 
**IdentifierPerm** | Pointer to **string** | Unique and permanent identifier for the package. | [optional] [readonly] 
**Indexed** | Pointer to **bool** |  | [optional] [readonly] 
**IsCancellable** | Pointer to **bool** |  | [optional] [readonly] 
**IsCopyable** | Pointer to **bool** |  | [optional] [readonly] 
**IsDeleteable** | Pointer to **bool** |  | [optional] [readonly] 
**IsDownloadable** | Pointer to **bool** |  | [optional] [readonly] 
**IsMoveable** | Pointer to **bool** |  | [optional] [readonly] 
**IsQuarantinable** | Pointer to **bool** |  | [optional] [readonly] 
**IsQuarantined** | Pointer to **bool** |  | [optional] [readonly] 
**IsResyncable** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecurityScannable** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncAwaiting** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncCompleted** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncFailed** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncInFlight** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncInProgress** | Pointer to **bool** |  | [optional] [readonly] 
**License** | Pointer to **NullableString** | The license of this package. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The name of this package. | [optional] [readonly] 
**Namespace** | Pointer to **string** |  | [optional] [readonly] 
**NamespaceUrl** | Pointer to **string** |  | [optional] [readonly] 
**NumFiles** | Pointer to **int64** |  | [optional] [readonly] 
**OriginRepository** | Pointer to **string** |  | [optional] [readonly] 
**OriginRepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**PackageType** | Pointer to **int64** | The type of package contents. | [optional] [readonly] 
**Release** | Pointer to **NullableString** | The release of the package version (if any). | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**RepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**SecurityScanCompletedAt** | Pointer to **NullableTime** | The datetime the security scanning was completed. | [optional] [readonly] 
**SecurityScanStartedAt** | Pointer to **NullableTime** | The datetime the security scanning was started. | [optional] [readonly] 
**SecurityScanStatus** | Pointer to **NullableString** |  | [optional] [readonly] [default to "Awaiting Security Scan"]
**SecurityScanStatusUpdatedAt** | Pointer to **NullableTime** | The datetime the security scanning status was updated. | [optional] [readonly] 
**SelfHtmlUrl** | Pointer to **string** |  | [optional] [readonly] 
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**SignatureUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Size** | Pointer to **int64** | The calculated size of the package. | [optional] [readonly] 
**Slug** | Pointer to **string** | The public unique identifier for the package. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Stage** | Pointer to **int64** | The synchronisation (in progress) stage of the package. | [optional] [readonly] 
**StageStr** | Pointer to **string** |  | [optional] [readonly] 
**StageUpdatedAt** | Pointer to **time.Time** | The datetime the package stage was updated at. | [optional] [readonly] 
**Status** | Pointer to **int64** | The synchronisation status of the package. | [optional] [readonly] 
**StatusReason** | Pointer to **NullableString** | A textual description for the synchronous status reason (if any | [optional] [readonly] 
**StatusStr** | Pointer to **string** |  | [optional] [readonly] 
**StatusUpdatedAt** | Pointer to **time.Time** | The datetime the package status was updated at. | [optional] [readonly] 
**StatusUrl** | Pointer to **string** |  | [optional] [readonly] 
**Subtype** | Pointer to **string** |  | [optional] [readonly] 
**Summary** | Pointer to **NullableString** | A one-liner synopsis of this package. | [optional] [readonly] 
**SyncFinishedAt** | Pointer to **NullableTime** | The datetime the package sync was finished at. | [optional] [readonly] 
**SyncProgress** | Pointer to **int64** | Synchronisation progress (from 0-100) | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TagsImmutable** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TypeDisplay** | Pointer to **string** |  | [optional] [readonly] 
**UploadedAt** | Pointer to **time.Time** | The date this package was uploaded. | [optional] [readonly] 
**Uploader** | Pointer to **string** |  | [optional] [readonly] 
**UploaderUrl** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] [readonly] 
**VersionOrig** | Pointer to **string** |  | [optional] [readonly] 
**VulnerabilityScanResultsUrl** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPackage

`func NewPackage() *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *Package) GetArchitectures() []Architecture`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *Package) GetArchitecturesOk() (*[]Architecture, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *Package) SetArchitectures(v []Architecture)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *Package) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetCdnUrl

`func (o *Package) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *Package) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *Package) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *Package) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### SetCdnUrlNil

`func (o *Package) SetCdnUrlNil(b bool)`

 SetCdnUrlNil sets the value for CdnUrl to be an explicit nil

### UnsetCdnUrl
`func (o *Package) UnsetCdnUrl()`

UnsetCdnUrl ensures that no value is present for CdnUrl, not even an explicit nil
### GetChecksumMd5

`func (o *Package) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *Package) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *Package) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *Package) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *Package) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *Package) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *Package) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *Package) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *Package) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *Package) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *Package) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *Package) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *Package) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *Package) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *Package) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *Package) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDependenciesChecksumMd5

`func (o *Package) GetDependenciesChecksumMd5() string`

GetDependenciesChecksumMd5 returns the DependenciesChecksumMd5 field if non-nil, zero value otherwise.

### GetDependenciesChecksumMd5Ok

`func (o *Package) GetDependenciesChecksumMd5Ok() (*string, bool)`

GetDependenciesChecksumMd5Ok returns a tuple with the DependenciesChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesChecksumMd5

`func (o *Package) SetDependenciesChecksumMd5(v string)`

SetDependenciesChecksumMd5 sets DependenciesChecksumMd5 field to given value.

### HasDependenciesChecksumMd5

`func (o *Package) HasDependenciesChecksumMd5() bool`

HasDependenciesChecksumMd5 returns a boolean if a field has been set.

### SetDependenciesChecksumMd5Nil

`func (o *Package) SetDependenciesChecksumMd5Nil(b bool)`

 SetDependenciesChecksumMd5Nil sets the value for DependenciesChecksumMd5 to be an explicit nil

### UnsetDependenciesChecksumMd5
`func (o *Package) UnsetDependenciesChecksumMd5()`

UnsetDependenciesChecksumMd5 ensures that no value is present for DependenciesChecksumMd5, not even an explicit nil
### GetDependenciesUrl

`func (o *Package) GetDependenciesUrl() string`

GetDependenciesUrl returns the DependenciesUrl field if non-nil, zero value otherwise.

### GetDependenciesUrlOk

`func (o *Package) GetDependenciesUrlOk() (*string, bool)`

GetDependenciesUrlOk returns a tuple with the DependenciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesUrl

`func (o *Package) SetDependenciesUrl(v string)`

SetDependenciesUrl sets DependenciesUrl field to given value.

### HasDependenciesUrl

`func (o *Package) HasDependenciesUrl() bool`

HasDependenciesUrl returns a boolean if a field has been set.

### GetDescription

`func (o *Package) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Package) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Package) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Package) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Package) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Package) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Package) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Package) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Package) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Package) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDistro

`func (o *Package) GetDistro() Distribution`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *Package) GetDistroOk() (*Distribution, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *Package) SetDistro(v Distribution)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *Package) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *Package) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *Package) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *Package) GetDistroVersion() DistributionVersion`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *Package) GetDistroVersionOk() (*DistributionVersion, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *Package) SetDistroVersion(v DistributionVersion)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *Package) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDownloads

`func (o *Package) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *Package) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *Package) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *Package) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEpoch

`func (o *Package) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *Package) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *Package) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *Package) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### SetEpochNil

`func (o *Package) SetEpochNil(b bool)`

 SetEpochNil sets the value for Epoch to be an explicit nil

### UnsetEpoch
`func (o *Package) UnsetEpoch()`

UnsetEpoch ensures that no value is present for Epoch, not even an explicit nil
### GetExtension

`func (o *Package) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Package) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Package) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Package) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFilename

`func (o *Package) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Package) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Package) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Package) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFiles

`func (o *Package) GetFiles() []PackageFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Package) GetFilesOk() (*[]PackageFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Package) SetFiles(v []PackageFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Package) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFormat

`func (o *Package) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Package) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Package) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Package) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *Package) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *Package) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *Package) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *Package) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetIdentifierPerm

`func (o *Package) GetIdentifierPerm() string`

GetIdentifierPerm returns the IdentifierPerm field if non-nil, zero value otherwise.

### GetIdentifierPermOk

`func (o *Package) GetIdentifierPermOk() (*string, bool)`

GetIdentifierPermOk returns a tuple with the IdentifierPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierPerm

`func (o *Package) SetIdentifierPerm(v string)`

SetIdentifierPerm sets IdentifierPerm field to given value.

### HasIdentifierPerm

`func (o *Package) HasIdentifierPerm() bool`

HasIdentifierPerm returns a boolean if a field has been set.

### GetIndexed

`func (o *Package) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *Package) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *Package) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *Package) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetIsCancellable

`func (o *Package) GetIsCancellable() bool`

GetIsCancellable returns the IsCancellable field if non-nil, zero value otherwise.

### GetIsCancellableOk

`func (o *Package) GetIsCancellableOk() (*bool, bool)`

GetIsCancellableOk returns a tuple with the IsCancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancellable

`func (o *Package) SetIsCancellable(v bool)`

SetIsCancellable sets IsCancellable field to given value.

### HasIsCancellable

`func (o *Package) HasIsCancellable() bool`

HasIsCancellable returns a boolean if a field has been set.

### GetIsCopyable

`func (o *Package) GetIsCopyable() bool`

GetIsCopyable returns the IsCopyable field if non-nil, zero value otherwise.

### GetIsCopyableOk

`func (o *Package) GetIsCopyableOk() (*bool, bool)`

GetIsCopyableOk returns a tuple with the IsCopyable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyable

`func (o *Package) SetIsCopyable(v bool)`

SetIsCopyable sets IsCopyable field to given value.

### HasIsCopyable

`func (o *Package) HasIsCopyable() bool`

HasIsCopyable returns a boolean if a field has been set.

### GetIsDeleteable

`func (o *Package) GetIsDeleteable() bool`

GetIsDeleteable returns the IsDeleteable field if non-nil, zero value otherwise.

### GetIsDeleteableOk

`func (o *Package) GetIsDeleteableOk() (*bool, bool)`

GetIsDeleteableOk returns a tuple with the IsDeleteable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteable

`func (o *Package) SetIsDeleteable(v bool)`

SetIsDeleteable sets IsDeleteable field to given value.

### HasIsDeleteable

`func (o *Package) HasIsDeleteable() bool`

HasIsDeleteable returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *Package) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *Package) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *Package) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *Package) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsMoveable

`func (o *Package) GetIsMoveable() bool`

GetIsMoveable returns the IsMoveable field if non-nil, zero value otherwise.

### GetIsMoveableOk

`func (o *Package) GetIsMoveableOk() (*bool, bool)`

GetIsMoveableOk returns a tuple with the IsMoveable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMoveable

`func (o *Package) SetIsMoveable(v bool)`

SetIsMoveable sets IsMoveable field to given value.

### HasIsMoveable

`func (o *Package) HasIsMoveable() bool`

HasIsMoveable returns a boolean if a field has been set.

### GetIsQuarantinable

`func (o *Package) GetIsQuarantinable() bool`

GetIsQuarantinable returns the IsQuarantinable field if non-nil, zero value otherwise.

### GetIsQuarantinableOk

`func (o *Package) GetIsQuarantinableOk() (*bool, bool)`

GetIsQuarantinableOk returns a tuple with the IsQuarantinable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantinable

`func (o *Package) SetIsQuarantinable(v bool)`

SetIsQuarantinable sets IsQuarantinable field to given value.

### HasIsQuarantinable

`func (o *Package) HasIsQuarantinable() bool`

HasIsQuarantinable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *Package) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *Package) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *Package) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *Package) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsResyncable

`func (o *Package) GetIsResyncable() bool`

GetIsResyncable returns the IsResyncable field if non-nil, zero value otherwise.

### GetIsResyncableOk

`func (o *Package) GetIsResyncableOk() (*bool, bool)`

GetIsResyncableOk returns a tuple with the IsResyncable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResyncable

`func (o *Package) SetIsResyncable(v bool)`

SetIsResyncable sets IsResyncable field to given value.

### HasIsResyncable

`func (o *Package) HasIsResyncable() bool`

HasIsResyncable returns a boolean if a field has been set.

### GetIsSecurityScannable

`func (o *Package) GetIsSecurityScannable() bool`

GetIsSecurityScannable returns the IsSecurityScannable field if non-nil, zero value otherwise.

### GetIsSecurityScannableOk

`func (o *Package) GetIsSecurityScannableOk() (*bool, bool)`

GetIsSecurityScannableOk returns a tuple with the IsSecurityScannable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityScannable

`func (o *Package) SetIsSecurityScannable(v bool)`

SetIsSecurityScannable sets IsSecurityScannable field to given value.

### HasIsSecurityScannable

`func (o *Package) HasIsSecurityScannable() bool`

HasIsSecurityScannable returns a boolean if a field has been set.

### GetIsSyncAwaiting

`func (o *Package) GetIsSyncAwaiting() bool`

GetIsSyncAwaiting returns the IsSyncAwaiting field if non-nil, zero value otherwise.

### GetIsSyncAwaitingOk

`func (o *Package) GetIsSyncAwaitingOk() (*bool, bool)`

GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncAwaiting

`func (o *Package) SetIsSyncAwaiting(v bool)`

SetIsSyncAwaiting sets IsSyncAwaiting field to given value.

### HasIsSyncAwaiting

`func (o *Package) HasIsSyncAwaiting() bool`

HasIsSyncAwaiting returns a boolean if a field has been set.

### GetIsSyncCompleted

`func (o *Package) GetIsSyncCompleted() bool`

GetIsSyncCompleted returns the IsSyncCompleted field if non-nil, zero value otherwise.

### GetIsSyncCompletedOk

`func (o *Package) GetIsSyncCompletedOk() (*bool, bool)`

GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncCompleted

`func (o *Package) SetIsSyncCompleted(v bool)`

SetIsSyncCompleted sets IsSyncCompleted field to given value.

### HasIsSyncCompleted

`func (o *Package) HasIsSyncCompleted() bool`

HasIsSyncCompleted returns a boolean if a field has been set.

### GetIsSyncFailed

`func (o *Package) GetIsSyncFailed() bool`

GetIsSyncFailed returns the IsSyncFailed field if non-nil, zero value otherwise.

### GetIsSyncFailedOk

`func (o *Package) GetIsSyncFailedOk() (*bool, bool)`

GetIsSyncFailedOk returns a tuple with the IsSyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncFailed

`func (o *Package) SetIsSyncFailed(v bool)`

SetIsSyncFailed sets IsSyncFailed field to given value.

### HasIsSyncFailed

`func (o *Package) HasIsSyncFailed() bool`

HasIsSyncFailed returns a boolean if a field has been set.

### GetIsSyncInFlight

`func (o *Package) GetIsSyncInFlight() bool`

GetIsSyncInFlight returns the IsSyncInFlight field if non-nil, zero value otherwise.

### GetIsSyncInFlightOk

`func (o *Package) GetIsSyncInFlightOk() (*bool, bool)`

GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInFlight

`func (o *Package) SetIsSyncInFlight(v bool)`

SetIsSyncInFlight sets IsSyncInFlight field to given value.

### HasIsSyncInFlight

`func (o *Package) HasIsSyncInFlight() bool`

HasIsSyncInFlight returns a boolean if a field has been set.

### GetIsSyncInProgress

`func (o *Package) GetIsSyncInProgress() bool`

GetIsSyncInProgress returns the IsSyncInProgress field if non-nil, zero value otherwise.

### GetIsSyncInProgressOk

`func (o *Package) GetIsSyncInProgressOk() (*bool, bool)`

GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInProgress

`func (o *Package) SetIsSyncInProgress(v bool)`

SetIsSyncInProgress sets IsSyncInProgress field to given value.

### HasIsSyncInProgress

`func (o *Package) HasIsSyncInProgress() bool`

HasIsSyncInProgress returns a boolean if a field has been set.

### GetLicense

`func (o *Package) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Package) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Package) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Package) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *Package) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *Package) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetName

`func (o *Package) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Package) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Package) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Package) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Package) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Package) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *Package) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Package) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Package) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Package) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *Package) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *Package) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *Package) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *Package) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumFiles

`func (o *Package) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *Package) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *Package) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *Package) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### GetOriginRepository

`func (o *Package) GetOriginRepository() string`

GetOriginRepository returns the OriginRepository field if non-nil, zero value otherwise.

### GetOriginRepositoryOk

`func (o *Package) GetOriginRepositoryOk() (*string, bool)`

GetOriginRepositoryOk returns a tuple with the OriginRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepository

`func (o *Package) SetOriginRepository(v string)`

SetOriginRepository sets OriginRepository field to given value.

### HasOriginRepository

`func (o *Package) HasOriginRepository() bool`

HasOriginRepository returns a boolean if a field has been set.

### GetOriginRepositoryUrl

`func (o *Package) GetOriginRepositoryUrl() string`

GetOriginRepositoryUrl returns the OriginRepositoryUrl field if non-nil, zero value otherwise.

### GetOriginRepositoryUrlOk

`func (o *Package) GetOriginRepositoryUrlOk() (*string, bool)`

GetOriginRepositoryUrlOk returns a tuple with the OriginRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepositoryUrl

`func (o *Package) SetOriginRepositoryUrl(v string)`

SetOriginRepositoryUrl sets OriginRepositoryUrl field to given value.

### HasOriginRepositoryUrl

`func (o *Package) HasOriginRepositoryUrl() bool`

HasOriginRepositoryUrl returns a boolean if a field has been set.

### GetPackageType

`func (o *Package) GetPackageType() int64`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *Package) GetPackageTypeOk() (*int64, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *Package) SetPackageType(v int64)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *Package) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetRelease

`func (o *Package) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Package) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Package) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Package) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *Package) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *Package) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetRepository

`func (o *Package) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Package) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Package) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Package) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *Package) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *Package) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *Package) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *Package) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *Package) GetSecurityScanCompletedAt() time.Time`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *Package) GetSecurityScanCompletedAtOk() (*time.Time, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *Package) SetSecurityScanCompletedAt(v time.Time)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *Package) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### SetSecurityScanCompletedAtNil

`func (o *Package) SetSecurityScanCompletedAtNil(b bool)`

 SetSecurityScanCompletedAtNil sets the value for SecurityScanCompletedAt to be an explicit nil

### UnsetSecurityScanCompletedAt
`func (o *Package) UnsetSecurityScanCompletedAt()`

UnsetSecurityScanCompletedAt ensures that no value is present for SecurityScanCompletedAt, not even an explicit nil
### GetSecurityScanStartedAt

`func (o *Package) GetSecurityScanStartedAt() time.Time`

GetSecurityScanStartedAt returns the SecurityScanStartedAt field if non-nil, zero value otherwise.

### GetSecurityScanStartedAtOk

`func (o *Package) GetSecurityScanStartedAtOk() (*time.Time, bool)`

GetSecurityScanStartedAtOk returns a tuple with the SecurityScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStartedAt

`func (o *Package) SetSecurityScanStartedAt(v time.Time)`

SetSecurityScanStartedAt sets SecurityScanStartedAt field to given value.

### HasSecurityScanStartedAt

`func (o *Package) HasSecurityScanStartedAt() bool`

HasSecurityScanStartedAt returns a boolean if a field has been set.

### SetSecurityScanStartedAtNil

`func (o *Package) SetSecurityScanStartedAtNil(b bool)`

 SetSecurityScanStartedAtNil sets the value for SecurityScanStartedAt to be an explicit nil

### UnsetSecurityScanStartedAt
`func (o *Package) UnsetSecurityScanStartedAt()`

UnsetSecurityScanStartedAt ensures that no value is present for SecurityScanStartedAt, not even an explicit nil
### GetSecurityScanStatus

`func (o *Package) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *Package) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *Package) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *Package) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### SetSecurityScanStatusNil

`func (o *Package) SetSecurityScanStatusNil(b bool)`

 SetSecurityScanStatusNil sets the value for SecurityScanStatus to be an explicit nil

### UnsetSecurityScanStatus
`func (o *Package) UnsetSecurityScanStatus()`

UnsetSecurityScanStatus ensures that no value is present for SecurityScanStatus, not even an explicit nil
### GetSecurityScanStatusUpdatedAt

`func (o *Package) GetSecurityScanStatusUpdatedAt() time.Time`

GetSecurityScanStatusUpdatedAt returns the SecurityScanStatusUpdatedAt field if non-nil, zero value otherwise.

### GetSecurityScanStatusUpdatedAtOk

`func (o *Package) GetSecurityScanStatusUpdatedAtOk() (*time.Time, bool)`

GetSecurityScanStatusUpdatedAtOk returns a tuple with the SecurityScanStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatusUpdatedAt

`func (o *Package) SetSecurityScanStatusUpdatedAt(v time.Time)`

SetSecurityScanStatusUpdatedAt sets SecurityScanStatusUpdatedAt field to given value.

### HasSecurityScanStatusUpdatedAt

`func (o *Package) HasSecurityScanStatusUpdatedAt() bool`

HasSecurityScanStatusUpdatedAt returns a boolean if a field has been set.

### SetSecurityScanStatusUpdatedAtNil

`func (o *Package) SetSecurityScanStatusUpdatedAtNil(b bool)`

 SetSecurityScanStatusUpdatedAtNil sets the value for SecurityScanStatusUpdatedAt to be an explicit nil

### UnsetSecurityScanStatusUpdatedAt
`func (o *Package) UnsetSecurityScanStatusUpdatedAt()`

UnsetSecurityScanStatusUpdatedAt ensures that no value is present for SecurityScanStatusUpdatedAt, not even an explicit nil
### GetSelfHtmlUrl

`func (o *Package) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *Package) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *Package) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *Package) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *Package) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *Package) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *Package) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *Package) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *Package) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *Package) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *Package) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *Package) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### SetSignatureUrlNil

`func (o *Package) SetSignatureUrlNil(b bool)`

 SetSignatureUrlNil sets the value for SignatureUrl to be an explicit nil

### UnsetSignatureUrl
`func (o *Package) UnsetSignatureUrl()`

UnsetSignatureUrl ensures that no value is present for SignatureUrl, not even an explicit nil
### GetSize

`func (o *Package) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Package) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Package) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Package) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlug

`func (o *Package) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Package) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Package) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Package) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *Package) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *Package) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *Package) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *Package) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStage

`func (o *Package) GetStage() int64`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Package) GetStageOk() (*int64, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Package) SetStage(v int64)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Package) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStageStr

`func (o *Package) GetStageStr() string`

GetStageStr returns the StageStr field if non-nil, zero value otherwise.

### GetStageStrOk

`func (o *Package) GetStageStrOk() (*string, bool)`

GetStageStrOk returns a tuple with the StageStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStr

`func (o *Package) SetStageStr(v string)`

SetStageStr sets StageStr field to given value.

### HasStageStr

`func (o *Package) HasStageStr() bool`

HasStageStr returns a boolean if a field has been set.

### GetStageUpdatedAt

`func (o *Package) GetStageUpdatedAt() time.Time`

GetStageUpdatedAt returns the StageUpdatedAt field if non-nil, zero value otherwise.

### GetStageUpdatedAtOk

`func (o *Package) GetStageUpdatedAtOk() (*time.Time, bool)`

GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUpdatedAt

`func (o *Package) SetStageUpdatedAt(v time.Time)`

SetStageUpdatedAt sets StageUpdatedAt field to given value.

### HasStageUpdatedAt

`func (o *Package) HasStageUpdatedAt() bool`

HasStageUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Package) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Package) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Package) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Package) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *Package) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *Package) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *Package) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *Package) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *Package) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *Package) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetStatusStr

`func (o *Package) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *Package) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *Package) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *Package) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *Package) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *Package) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *Package) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *Package) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStatusUrl

`func (o *Package) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *Package) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *Package) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *Package) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSubtype

`func (o *Package) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Package) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Package) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Package) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSummary

`func (o *Package) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Package) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Package) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Package) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *Package) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *Package) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetSyncFinishedAt

`func (o *Package) GetSyncFinishedAt() time.Time`

GetSyncFinishedAt returns the SyncFinishedAt field if non-nil, zero value otherwise.

### GetSyncFinishedAtOk

`func (o *Package) GetSyncFinishedAtOk() (*time.Time, bool)`

GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFinishedAt

`func (o *Package) SetSyncFinishedAt(v time.Time)`

SetSyncFinishedAt sets SyncFinishedAt field to given value.

### HasSyncFinishedAt

`func (o *Package) HasSyncFinishedAt() bool`

HasSyncFinishedAt returns a boolean if a field has been set.

### SetSyncFinishedAtNil

`func (o *Package) SetSyncFinishedAtNil(b bool)`

 SetSyncFinishedAtNil sets the value for SyncFinishedAt to be an explicit nil

### UnsetSyncFinishedAt
`func (o *Package) UnsetSyncFinishedAt()`

UnsetSyncFinishedAt ensures that no value is present for SyncFinishedAt, not even an explicit nil
### GetSyncProgress

`func (o *Package) GetSyncProgress() int64`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *Package) GetSyncProgressOk() (*int64, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *Package) SetSyncProgress(v int64)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *Package) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.

### GetTags

`func (o *Package) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Package) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Package) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Package) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagsImmutable

`func (o *Package) GetTagsImmutable() map[string]interface{}`

GetTagsImmutable returns the TagsImmutable field if non-nil, zero value otherwise.

### GetTagsImmutableOk

`func (o *Package) GetTagsImmutableOk() (*map[string]interface{}, bool)`

GetTagsImmutableOk returns a tuple with the TagsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsImmutable

`func (o *Package) SetTagsImmutable(v map[string]interface{})`

SetTagsImmutable sets TagsImmutable field to given value.

### HasTagsImmutable

`func (o *Package) HasTagsImmutable() bool`

HasTagsImmutable returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *Package) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *Package) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *Package) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *Package) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *Package) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *Package) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *Package) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *Package) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *Package) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *Package) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *Package) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *Package) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderUrl

`func (o *Package) GetUploaderUrl() string`

GetUploaderUrl returns the UploaderUrl field if non-nil, zero value otherwise.

### GetUploaderUrlOk

`func (o *Package) GetUploaderUrlOk() (*string, bool)`

GetUploaderUrlOk returns a tuple with the UploaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderUrl

`func (o *Package) SetUploaderUrl(v string)`

SetUploaderUrl sets UploaderUrl field to given value.

### HasUploaderUrl

`func (o *Package) HasUploaderUrl() bool`

HasUploaderUrl returns a boolean if a field has been set.

### GetVersion

`func (o *Package) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Package) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Package) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Package) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Package) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Package) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVersionOrig

`func (o *Package) GetVersionOrig() string`

GetVersionOrig returns the VersionOrig field if non-nil, zero value otherwise.

### GetVersionOrigOk

`func (o *Package) GetVersionOrigOk() (*string, bool)`

GetVersionOrigOk returns a tuple with the VersionOrig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrig

`func (o *Package) SetVersionOrig(v string)`

SetVersionOrig sets VersionOrig field to given value.

### HasVersionOrig

`func (o *Package) HasVersionOrig() bool`

HasVersionOrig returns a boolean if a field has been set.

### GetVulnerabilityScanResultsUrl

`func (o *Package) GetVulnerabilityScanResultsUrl() string`

GetVulnerabilityScanResultsUrl returns the VulnerabilityScanResultsUrl field if non-nil, zero value otherwise.

### GetVulnerabilityScanResultsUrlOk

`func (o *Package) GetVulnerabilityScanResultsUrlOk() (*string, bool)`

GetVulnerabilityScanResultsUrlOk returns a tuple with the VulnerabilityScanResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanResultsUrl

`func (o *Package) SetVulnerabilityScanResultsUrl(v string)`

SetVulnerabilityScanResultsUrl sets VulnerabilityScanResultsUrl field to given value.

### HasVulnerabilityScanResultsUrl

`func (o *Package) HasVulnerabilityScanResultsUrl() bool`

HasVulnerabilityScanResultsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


