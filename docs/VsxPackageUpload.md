# VsxPackageUpload

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
**FreeableStorage** | Pointer to **int64** | Amount of storage that will be freed if this package is deleted | [optional] [readonly] 
**FullyQualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**IdentifierPerm** | Pointer to **string** | Unique and permanent identifier for the package. | [optional] [readonly] 
**Identifiers** | Pointer to **map[string]string** | Return a map of identifier field names and their values. | [optional] [readonly] 
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
**PolicyViolated** | Pointer to **bool** | Whether or not the package has violated any policy. | [optional] [readonly] 
**RawLicense** | Pointer to **NullableString** | The raw license string. | [optional] [readonly] 
**Release** | Pointer to **NullableString** | The release of the package version (if any). | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**RepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**SecurityScanCompletedAt** | Pointer to **NullableTime** | The datetime the security scanning was completed. | [optional] [readonly] 
**SecurityScanStartedAt** | Pointer to **NullableTime** | The datetime the security scanning was started. | [optional] [readonly] 
**SecurityScanStatus** | Pointer to **NullableString** |  | [optional] [readonly] [default to "Awaiting Security Scan"]
**SecurityScanStatusUpdatedAt** | Pointer to **NullableTime** | The datetime the security scanning status was updated. | [optional] [readonly] 
**SelfHtmlUrl** | Pointer to **string** |  | [optional] [readonly] 
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**SelfWebappUrl** | Pointer to **string** |  | [optional] [readonly] 
**SignatureUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Size** | Pointer to **int64** | The calculated size of the package. | [optional] [readonly] 
**Slug** | Pointer to **string** | The public unique identifier for the package. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**SpdxLicense** | Pointer to **NullableString** | The SPDX license identifier for this package. | [optional] [readonly] 
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
**TagsAutomatic** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TagsImmutable** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TypeDisplay** | Pointer to **string** |  | [optional] [readonly] 
**UploadedAt** | Pointer to **time.Time** | The date this package was uploaded. | [optional] [readonly] 
**Uploader** | Pointer to **string** |  | [optional] [readonly] 
**UploaderUrl** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] [readonly] 
**VersionOrig** | Pointer to **string** |  | [optional] [readonly] 
**VulnerabilityScanResultsUrl** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewVsxPackageUpload

`func NewVsxPackageUpload() *VsxPackageUpload`

NewVsxPackageUpload instantiates a new VsxPackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVsxPackageUploadWithDefaults

`func NewVsxPackageUploadWithDefaults() *VsxPackageUpload`

NewVsxPackageUploadWithDefaults instantiates a new VsxPackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *VsxPackageUpload) GetArchitectures() []Architecture`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *VsxPackageUpload) GetArchitecturesOk() (*[]Architecture, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *VsxPackageUpload) SetArchitectures(v []Architecture)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *VsxPackageUpload) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetCdnUrl

`func (o *VsxPackageUpload) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *VsxPackageUpload) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *VsxPackageUpload) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *VsxPackageUpload) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### SetCdnUrlNil

`func (o *VsxPackageUpload) SetCdnUrlNil(b bool)`

 SetCdnUrlNil sets the value for CdnUrl to be an explicit nil

### UnsetCdnUrl
`func (o *VsxPackageUpload) UnsetCdnUrl()`

UnsetCdnUrl ensures that no value is present for CdnUrl, not even an explicit nil
### GetChecksumMd5

`func (o *VsxPackageUpload) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *VsxPackageUpload) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *VsxPackageUpload) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *VsxPackageUpload) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *VsxPackageUpload) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *VsxPackageUpload) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *VsxPackageUpload) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *VsxPackageUpload) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *VsxPackageUpload) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *VsxPackageUpload) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *VsxPackageUpload) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *VsxPackageUpload) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *VsxPackageUpload) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *VsxPackageUpload) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *VsxPackageUpload) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *VsxPackageUpload) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDependenciesChecksumMd5

`func (o *VsxPackageUpload) GetDependenciesChecksumMd5() string`

GetDependenciesChecksumMd5 returns the DependenciesChecksumMd5 field if non-nil, zero value otherwise.

### GetDependenciesChecksumMd5Ok

`func (o *VsxPackageUpload) GetDependenciesChecksumMd5Ok() (*string, bool)`

GetDependenciesChecksumMd5Ok returns a tuple with the DependenciesChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesChecksumMd5

`func (o *VsxPackageUpload) SetDependenciesChecksumMd5(v string)`

SetDependenciesChecksumMd5 sets DependenciesChecksumMd5 field to given value.

### HasDependenciesChecksumMd5

`func (o *VsxPackageUpload) HasDependenciesChecksumMd5() bool`

HasDependenciesChecksumMd5 returns a boolean if a field has been set.

### SetDependenciesChecksumMd5Nil

`func (o *VsxPackageUpload) SetDependenciesChecksumMd5Nil(b bool)`

 SetDependenciesChecksumMd5Nil sets the value for DependenciesChecksumMd5 to be an explicit nil

### UnsetDependenciesChecksumMd5
`func (o *VsxPackageUpload) UnsetDependenciesChecksumMd5()`

UnsetDependenciesChecksumMd5 ensures that no value is present for DependenciesChecksumMd5, not even an explicit nil
### GetDependenciesUrl

`func (o *VsxPackageUpload) GetDependenciesUrl() string`

GetDependenciesUrl returns the DependenciesUrl field if non-nil, zero value otherwise.

### GetDependenciesUrlOk

`func (o *VsxPackageUpload) GetDependenciesUrlOk() (*string, bool)`

GetDependenciesUrlOk returns a tuple with the DependenciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesUrl

`func (o *VsxPackageUpload) SetDependenciesUrl(v string)`

SetDependenciesUrl sets DependenciesUrl field to given value.

### HasDependenciesUrl

`func (o *VsxPackageUpload) HasDependenciesUrl() bool`

HasDependenciesUrl returns a boolean if a field has been set.

### GetDescription

`func (o *VsxPackageUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VsxPackageUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VsxPackageUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VsxPackageUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VsxPackageUpload) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VsxPackageUpload) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *VsxPackageUpload) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VsxPackageUpload) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VsxPackageUpload) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VsxPackageUpload) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDistro

`func (o *VsxPackageUpload) GetDistro() Distribution`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *VsxPackageUpload) GetDistroOk() (*Distribution, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *VsxPackageUpload) SetDistro(v Distribution)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *VsxPackageUpload) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *VsxPackageUpload) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *VsxPackageUpload) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *VsxPackageUpload) GetDistroVersion() DistributionVersion`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *VsxPackageUpload) GetDistroVersionOk() (*DistributionVersion, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *VsxPackageUpload) SetDistroVersion(v DistributionVersion)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *VsxPackageUpload) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDownloads

`func (o *VsxPackageUpload) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *VsxPackageUpload) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *VsxPackageUpload) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *VsxPackageUpload) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEpoch

`func (o *VsxPackageUpload) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *VsxPackageUpload) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *VsxPackageUpload) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *VsxPackageUpload) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### SetEpochNil

`func (o *VsxPackageUpload) SetEpochNil(b bool)`

 SetEpochNil sets the value for Epoch to be an explicit nil

### UnsetEpoch
`func (o *VsxPackageUpload) UnsetEpoch()`

UnsetEpoch ensures that no value is present for Epoch, not even an explicit nil
### GetExtension

`func (o *VsxPackageUpload) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *VsxPackageUpload) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *VsxPackageUpload) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *VsxPackageUpload) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFilename

`func (o *VsxPackageUpload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VsxPackageUpload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VsxPackageUpload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *VsxPackageUpload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFiles

`func (o *VsxPackageUpload) GetFiles() []PackageFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VsxPackageUpload) GetFilesOk() (*[]PackageFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VsxPackageUpload) SetFiles(v []PackageFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *VsxPackageUpload) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFormat

`func (o *VsxPackageUpload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VsxPackageUpload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VsxPackageUpload) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VsxPackageUpload) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *VsxPackageUpload) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *VsxPackageUpload) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *VsxPackageUpload) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *VsxPackageUpload) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetFreeableStorage

`func (o *VsxPackageUpload) GetFreeableStorage() int64`

GetFreeableStorage returns the FreeableStorage field if non-nil, zero value otherwise.

### GetFreeableStorageOk

`func (o *VsxPackageUpload) GetFreeableStorageOk() (*int64, bool)`

GetFreeableStorageOk returns a tuple with the FreeableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeableStorage

`func (o *VsxPackageUpload) SetFreeableStorage(v int64)`

SetFreeableStorage sets FreeableStorage field to given value.

### HasFreeableStorage

`func (o *VsxPackageUpload) HasFreeableStorage() bool`

HasFreeableStorage returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *VsxPackageUpload) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *VsxPackageUpload) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *VsxPackageUpload) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *VsxPackageUpload) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### SetFullyQualifiedNameNil

`func (o *VsxPackageUpload) SetFullyQualifiedNameNil(b bool)`

 SetFullyQualifiedNameNil sets the value for FullyQualifiedName to be an explicit nil

### UnsetFullyQualifiedName
`func (o *VsxPackageUpload) UnsetFullyQualifiedName()`

UnsetFullyQualifiedName ensures that no value is present for FullyQualifiedName, not even an explicit nil
### GetIdentifierPerm

`func (o *VsxPackageUpload) GetIdentifierPerm() string`

GetIdentifierPerm returns the IdentifierPerm field if non-nil, zero value otherwise.

### GetIdentifierPermOk

`func (o *VsxPackageUpload) GetIdentifierPermOk() (*string, bool)`

GetIdentifierPermOk returns a tuple with the IdentifierPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierPerm

`func (o *VsxPackageUpload) SetIdentifierPerm(v string)`

SetIdentifierPerm sets IdentifierPerm field to given value.

### HasIdentifierPerm

`func (o *VsxPackageUpload) HasIdentifierPerm() bool`

HasIdentifierPerm returns a boolean if a field has been set.

### GetIdentifiers

`func (o *VsxPackageUpload) GetIdentifiers() map[string]string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *VsxPackageUpload) GetIdentifiersOk() (*map[string]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *VsxPackageUpload) SetIdentifiers(v map[string]string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *VsxPackageUpload) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetIndexed

`func (o *VsxPackageUpload) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *VsxPackageUpload) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *VsxPackageUpload) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *VsxPackageUpload) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetIsCancellable

`func (o *VsxPackageUpload) GetIsCancellable() bool`

GetIsCancellable returns the IsCancellable field if non-nil, zero value otherwise.

### GetIsCancellableOk

`func (o *VsxPackageUpload) GetIsCancellableOk() (*bool, bool)`

GetIsCancellableOk returns a tuple with the IsCancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancellable

`func (o *VsxPackageUpload) SetIsCancellable(v bool)`

SetIsCancellable sets IsCancellable field to given value.

### HasIsCancellable

`func (o *VsxPackageUpload) HasIsCancellable() bool`

HasIsCancellable returns a boolean if a field has been set.

### GetIsCopyable

`func (o *VsxPackageUpload) GetIsCopyable() bool`

GetIsCopyable returns the IsCopyable field if non-nil, zero value otherwise.

### GetIsCopyableOk

`func (o *VsxPackageUpload) GetIsCopyableOk() (*bool, bool)`

GetIsCopyableOk returns a tuple with the IsCopyable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyable

`func (o *VsxPackageUpload) SetIsCopyable(v bool)`

SetIsCopyable sets IsCopyable field to given value.

### HasIsCopyable

`func (o *VsxPackageUpload) HasIsCopyable() bool`

HasIsCopyable returns a boolean if a field has been set.

### GetIsDeleteable

`func (o *VsxPackageUpload) GetIsDeleteable() bool`

GetIsDeleteable returns the IsDeleteable field if non-nil, zero value otherwise.

### GetIsDeleteableOk

`func (o *VsxPackageUpload) GetIsDeleteableOk() (*bool, bool)`

GetIsDeleteableOk returns a tuple with the IsDeleteable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteable

`func (o *VsxPackageUpload) SetIsDeleteable(v bool)`

SetIsDeleteable sets IsDeleteable field to given value.

### HasIsDeleteable

`func (o *VsxPackageUpload) HasIsDeleteable() bool`

HasIsDeleteable returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *VsxPackageUpload) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *VsxPackageUpload) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *VsxPackageUpload) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *VsxPackageUpload) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsMoveable

`func (o *VsxPackageUpload) GetIsMoveable() bool`

GetIsMoveable returns the IsMoveable field if non-nil, zero value otherwise.

### GetIsMoveableOk

`func (o *VsxPackageUpload) GetIsMoveableOk() (*bool, bool)`

GetIsMoveableOk returns a tuple with the IsMoveable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMoveable

`func (o *VsxPackageUpload) SetIsMoveable(v bool)`

SetIsMoveable sets IsMoveable field to given value.

### HasIsMoveable

`func (o *VsxPackageUpload) HasIsMoveable() bool`

HasIsMoveable returns a boolean if a field has been set.

### GetIsQuarantinable

`func (o *VsxPackageUpload) GetIsQuarantinable() bool`

GetIsQuarantinable returns the IsQuarantinable field if non-nil, zero value otherwise.

### GetIsQuarantinableOk

`func (o *VsxPackageUpload) GetIsQuarantinableOk() (*bool, bool)`

GetIsQuarantinableOk returns a tuple with the IsQuarantinable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantinable

`func (o *VsxPackageUpload) SetIsQuarantinable(v bool)`

SetIsQuarantinable sets IsQuarantinable field to given value.

### HasIsQuarantinable

`func (o *VsxPackageUpload) HasIsQuarantinable() bool`

HasIsQuarantinable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *VsxPackageUpload) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *VsxPackageUpload) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *VsxPackageUpload) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *VsxPackageUpload) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsResyncable

`func (o *VsxPackageUpload) GetIsResyncable() bool`

GetIsResyncable returns the IsResyncable field if non-nil, zero value otherwise.

### GetIsResyncableOk

`func (o *VsxPackageUpload) GetIsResyncableOk() (*bool, bool)`

GetIsResyncableOk returns a tuple with the IsResyncable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResyncable

`func (o *VsxPackageUpload) SetIsResyncable(v bool)`

SetIsResyncable sets IsResyncable field to given value.

### HasIsResyncable

`func (o *VsxPackageUpload) HasIsResyncable() bool`

HasIsResyncable returns a boolean if a field has been set.

### GetIsSecurityScannable

`func (o *VsxPackageUpload) GetIsSecurityScannable() bool`

GetIsSecurityScannable returns the IsSecurityScannable field if non-nil, zero value otherwise.

### GetIsSecurityScannableOk

`func (o *VsxPackageUpload) GetIsSecurityScannableOk() (*bool, bool)`

GetIsSecurityScannableOk returns a tuple with the IsSecurityScannable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityScannable

`func (o *VsxPackageUpload) SetIsSecurityScannable(v bool)`

SetIsSecurityScannable sets IsSecurityScannable field to given value.

### HasIsSecurityScannable

`func (o *VsxPackageUpload) HasIsSecurityScannable() bool`

HasIsSecurityScannable returns a boolean if a field has been set.

### GetIsSyncAwaiting

`func (o *VsxPackageUpload) GetIsSyncAwaiting() bool`

GetIsSyncAwaiting returns the IsSyncAwaiting field if non-nil, zero value otherwise.

### GetIsSyncAwaitingOk

`func (o *VsxPackageUpload) GetIsSyncAwaitingOk() (*bool, bool)`

GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncAwaiting

`func (o *VsxPackageUpload) SetIsSyncAwaiting(v bool)`

SetIsSyncAwaiting sets IsSyncAwaiting field to given value.

### HasIsSyncAwaiting

`func (o *VsxPackageUpload) HasIsSyncAwaiting() bool`

HasIsSyncAwaiting returns a boolean if a field has been set.

### GetIsSyncCompleted

`func (o *VsxPackageUpload) GetIsSyncCompleted() bool`

GetIsSyncCompleted returns the IsSyncCompleted field if non-nil, zero value otherwise.

### GetIsSyncCompletedOk

`func (o *VsxPackageUpload) GetIsSyncCompletedOk() (*bool, bool)`

GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncCompleted

`func (o *VsxPackageUpload) SetIsSyncCompleted(v bool)`

SetIsSyncCompleted sets IsSyncCompleted field to given value.

### HasIsSyncCompleted

`func (o *VsxPackageUpload) HasIsSyncCompleted() bool`

HasIsSyncCompleted returns a boolean if a field has been set.

### GetIsSyncFailed

`func (o *VsxPackageUpload) GetIsSyncFailed() bool`

GetIsSyncFailed returns the IsSyncFailed field if non-nil, zero value otherwise.

### GetIsSyncFailedOk

`func (o *VsxPackageUpload) GetIsSyncFailedOk() (*bool, bool)`

GetIsSyncFailedOk returns a tuple with the IsSyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncFailed

`func (o *VsxPackageUpload) SetIsSyncFailed(v bool)`

SetIsSyncFailed sets IsSyncFailed field to given value.

### HasIsSyncFailed

`func (o *VsxPackageUpload) HasIsSyncFailed() bool`

HasIsSyncFailed returns a boolean if a field has been set.

### GetIsSyncInFlight

`func (o *VsxPackageUpload) GetIsSyncInFlight() bool`

GetIsSyncInFlight returns the IsSyncInFlight field if non-nil, zero value otherwise.

### GetIsSyncInFlightOk

`func (o *VsxPackageUpload) GetIsSyncInFlightOk() (*bool, bool)`

GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInFlight

`func (o *VsxPackageUpload) SetIsSyncInFlight(v bool)`

SetIsSyncInFlight sets IsSyncInFlight field to given value.

### HasIsSyncInFlight

`func (o *VsxPackageUpload) HasIsSyncInFlight() bool`

HasIsSyncInFlight returns a boolean if a field has been set.

### GetIsSyncInProgress

`func (o *VsxPackageUpload) GetIsSyncInProgress() bool`

GetIsSyncInProgress returns the IsSyncInProgress field if non-nil, zero value otherwise.

### GetIsSyncInProgressOk

`func (o *VsxPackageUpload) GetIsSyncInProgressOk() (*bool, bool)`

GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInProgress

`func (o *VsxPackageUpload) SetIsSyncInProgress(v bool)`

SetIsSyncInProgress sets IsSyncInProgress field to given value.

### HasIsSyncInProgress

`func (o *VsxPackageUpload) HasIsSyncInProgress() bool`

HasIsSyncInProgress returns a boolean if a field has been set.

### GetLicense

`func (o *VsxPackageUpload) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *VsxPackageUpload) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *VsxPackageUpload) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *VsxPackageUpload) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *VsxPackageUpload) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *VsxPackageUpload) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetName

`func (o *VsxPackageUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VsxPackageUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VsxPackageUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VsxPackageUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VsxPackageUpload) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VsxPackageUpload) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *VsxPackageUpload) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VsxPackageUpload) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VsxPackageUpload) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *VsxPackageUpload) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *VsxPackageUpload) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *VsxPackageUpload) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *VsxPackageUpload) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *VsxPackageUpload) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumFiles

`func (o *VsxPackageUpload) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *VsxPackageUpload) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *VsxPackageUpload) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *VsxPackageUpload) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### GetOriginRepository

`func (o *VsxPackageUpload) GetOriginRepository() string`

GetOriginRepository returns the OriginRepository field if non-nil, zero value otherwise.

### GetOriginRepositoryOk

`func (o *VsxPackageUpload) GetOriginRepositoryOk() (*string, bool)`

GetOriginRepositoryOk returns a tuple with the OriginRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepository

`func (o *VsxPackageUpload) SetOriginRepository(v string)`

SetOriginRepository sets OriginRepository field to given value.

### HasOriginRepository

`func (o *VsxPackageUpload) HasOriginRepository() bool`

HasOriginRepository returns a boolean if a field has been set.

### GetOriginRepositoryUrl

`func (o *VsxPackageUpload) GetOriginRepositoryUrl() string`

GetOriginRepositoryUrl returns the OriginRepositoryUrl field if non-nil, zero value otherwise.

### GetOriginRepositoryUrlOk

`func (o *VsxPackageUpload) GetOriginRepositoryUrlOk() (*string, bool)`

GetOriginRepositoryUrlOk returns a tuple with the OriginRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepositoryUrl

`func (o *VsxPackageUpload) SetOriginRepositoryUrl(v string)`

SetOriginRepositoryUrl sets OriginRepositoryUrl field to given value.

### HasOriginRepositoryUrl

`func (o *VsxPackageUpload) HasOriginRepositoryUrl() bool`

HasOriginRepositoryUrl returns a boolean if a field has been set.

### GetPackageType

`func (o *VsxPackageUpload) GetPackageType() int64`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *VsxPackageUpload) GetPackageTypeOk() (*int64, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *VsxPackageUpload) SetPackageType(v int64)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *VsxPackageUpload) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPolicyViolated

`func (o *VsxPackageUpload) GetPolicyViolated() bool`

GetPolicyViolated returns the PolicyViolated field if non-nil, zero value otherwise.

### GetPolicyViolatedOk

`func (o *VsxPackageUpload) GetPolicyViolatedOk() (*bool, bool)`

GetPolicyViolatedOk returns a tuple with the PolicyViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolated

`func (o *VsxPackageUpload) SetPolicyViolated(v bool)`

SetPolicyViolated sets PolicyViolated field to given value.

### HasPolicyViolated

`func (o *VsxPackageUpload) HasPolicyViolated() bool`

HasPolicyViolated returns a boolean if a field has been set.

### GetRawLicense

`func (o *VsxPackageUpload) GetRawLicense() string`

GetRawLicense returns the RawLicense field if non-nil, zero value otherwise.

### GetRawLicenseOk

`func (o *VsxPackageUpload) GetRawLicenseOk() (*string, bool)`

GetRawLicenseOk returns a tuple with the RawLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLicense

`func (o *VsxPackageUpload) SetRawLicense(v string)`

SetRawLicense sets RawLicense field to given value.

### HasRawLicense

`func (o *VsxPackageUpload) HasRawLicense() bool`

HasRawLicense returns a boolean if a field has been set.

### SetRawLicenseNil

`func (o *VsxPackageUpload) SetRawLicenseNil(b bool)`

 SetRawLicenseNil sets the value for RawLicense to be an explicit nil

### UnsetRawLicense
`func (o *VsxPackageUpload) UnsetRawLicense()`

UnsetRawLicense ensures that no value is present for RawLicense, not even an explicit nil
### GetRelease

`func (o *VsxPackageUpload) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *VsxPackageUpload) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *VsxPackageUpload) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *VsxPackageUpload) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *VsxPackageUpload) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *VsxPackageUpload) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetRepository

`func (o *VsxPackageUpload) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *VsxPackageUpload) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *VsxPackageUpload) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *VsxPackageUpload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *VsxPackageUpload) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *VsxPackageUpload) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *VsxPackageUpload) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *VsxPackageUpload) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *VsxPackageUpload) GetSecurityScanCompletedAt() time.Time`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *VsxPackageUpload) GetSecurityScanCompletedAtOk() (*time.Time, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *VsxPackageUpload) SetSecurityScanCompletedAt(v time.Time)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *VsxPackageUpload) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### SetSecurityScanCompletedAtNil

`func (o *VsxPackageUpload) SetSecurityScanCompletedAtNil(b bool)`

 SetSecurityScanCompletedAtNil sets the value for SecurityScanCompletedAt to be an explicit nil

### UnsetSecurityScanCompletedAt
`func (o *VsxPackageUpload) UnsetSecurityScanCompletedAt()`

UnsetSecurityScanCompletedAt ensures that no value is present for SecurityScanCompletedAt, not even an explicit nil
### GetSecurityScanStartedAt

`func (o *VsxPackageUpload) GetSecurityScanStartedAt() time.Time`

GetSecurityScanStartedAt returns the SecurityScanStartedAt field if non-nil, zero value otherwise.

### GetSecurityScanStartedAtOk

`func (o *VsxPackageUpload) GetSecurityScanStartedAtOk() (*time.Time, bool)`

GetSecurityScanStartedAtOk returns a tuple with the SecurityScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStartedAt

`func (o *VsxPackageUpload) SetSecurityScanStartedAt(v time.Time)`

SetSecurityScanStartedAt sets SecurityScanStartedAt field to given value.

### HasSecurityScanStartedAt

`func (o *VsxPackageUpload) HasSecurityScanStartedAt() bool`

HasSecurityScanStartedAt returns a boolean if a field has been set.

### SetSecurityScanStartedAtNil

`func (o *VsxPackageUpload) SetSecurityScanStartedAtNil(b bool)`

 SetSecurityScanStartedAtNil sets the value for SecurityScanStartedAt to be an explicit nil

### UnsetSecurityScanStartedAt
`func (o *VsxPackageUpload) UnsetSecurityScanStartedAt()`

UnsetSecurityScanStartedAt ensures that no value is present for SecurityScanStartedAt, not even an explicit nil
### GetSecurityScanStatus

`func (o *VsxPackageUpload) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *VsxPackageUpload) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *VsxPackageUpload) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *VsxPackageUpload) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### SetSecurityScanStatusNil

`func (o *VsxPackageUpload) SetSecurityScanStatusNil(b bool)`

 SetSecurityScanStatusNil sets the value for SecurityScanStatus to be an explicit nil

### UnsetSecurityScanStatus
`func (o *VsxPackageUpload) UnsetSecurityScanStatus()`

UnsetSecurityScanStatus ensures that no value is present for SecurityScanStatus, not even an explicit nil
### GetSecurityScanStatusUpdatedAt

`func (o *VsxPackageUpload) GetSecurityScanStatusUpdatedAt() time.Time`

GetSecurityScanStatusUpdatedAt returns the SecurityScanStatusUpdatedAt field if non-nil, zero value otherwise.

### GetSecurityScanStatusUpdatedAtOk

`func (o *VsxPackageUpload) GetSecurityScanStatusUpdatedAtOk() (*time.Time, bool)`

GetSecurityScanStatusUpdatedAtOk returns a tuple with the SecurityScanStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatusUpdatedAt

`func (o *VsxPackageUpload) SetSecurityScanStatusUpdatedAt(v time.Time)`

SetSecurityScanStatusUpdatedAt sets SecurityScanStatusUpdatedAt field to given value.

### HasSecurityScanStatusUpdatedAt

`func (o *VsxPackageUpload) HasSecurityScanStatusUpdatedAt() bool`

HasSecurityScanStatusUpdatedAt returns a boolean if a field has been set.

### SetSecurityScanStatusUpdatedAtNil

`func (o *VsxPackageUpload) SetSecurityScanStatusUpdatedAtNil(b bool)`

 SetSecurityScanStatusUpdatedAtNil sets the value for SecurityScanStatusUpdatedAt to be an explicit nil

### UnsetSecurityScanStatusUpdatedAt
`func (o *VsxPackageUpload) UnsetSecurityScanStatusUpdatedAt()`

UnsetSecurityScanStatusUpdatedAt ensures that no value is present for SecurityScanStatusUpdatedAt, not even an explicit nil
### GetSelfHtmlUrl

`func (o *VsxPackageUpload) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *VsxPackageUpload) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *VsxPackageUpload) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *VsxPackageUpload) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *VsxPackageUpload) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *VsxPackageUpload) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *VsxPackageUpload) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *VsxPackageUpload) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSelfWebappUrl

`func (o *VsxPackageUpload) GetSelfWebappUrl() string`

GetSelfWebappUrl returns the SelfWebappUrl field if non-nil, zero value otherwise.

### GetSelfWebappUrlOk

`func (o *VsxPackageUpload) GetSelfWebappUrlOk() (*string, bool)`

GetSelfWebappUrlOk returns a tuple with the SelfWebappUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfWebappUrl

`func (o *VsxPackageUpload) SetSelfWebappUrl(v string)`

SetSelfWebappUrl sets SelfWebappUrl field to given value.

### HasSelfWebappUrl

`func (o *VsxPackageUpload) HasSelfWebappUrl() bool`

HasSelfWebappUrl returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *VsxPackageUpload) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *VsxPackageUpload) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *VsxPackageUpload) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *VsxPackageUpload) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### SetSignatureUrlNil

`func (o *VsxPackageUpload) SetSignatureUrlNil(b bool)`

 SetSignatureUrlNil sets the value for SignatureUrl to be an explicit nil

### UnsetSignatureUrl
`func (o *VsxPackageUpload) UnsetSignatureUrl()`

UnsetSignatureUrl ensures that no value is present for SignatureUrl, not even an explicit nil
### GetSize

`func (o *VsxPackageUpload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VsxPackageUpload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VsxPackageUpload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VsxPackageUpload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlug

`func (o *VsxPackageUpload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VsxPackageUpload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VsxPackageUpload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *VsxPackageUpload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *VsxPackageUpload) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *VsxPackageUpload) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *VsxPackageUpload) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *VsxPackageUpload) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetSpdxLicense

`func (o *VsxPackageUpload) GetSpdxLicense() string`

GetSpdxLicense returns the SpdxLicense field if non-nil, zero value otherwise.

### GetSpdxLicenseOk

`func (o *VsxPackageUpload) GetSpdxLicenseOk() (*string, bool)`

GetSpdxLicenseOk returns a tuple with the SpdxLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxLicense

`func (o *VsxPackageUpload) SetSpdxLicense(v string)`

SetSpdxLicense sets SpdxLicense field to given value.

### HasSpdxLicense

`func (o *VsxPackageUpload) HasSpdxLicense() bool`

HasSpdxLicense returns a boolean if a field has been set.

### SetSpdxLicenseNil

`func (o *VsxPackageUpload) SetSpdxLicenseNil(b bool)`

 SetSpdxLicenseNil sets the value for SpdxLicense to be an explicit nil

### UnsetSpdxLicense
`func (o *VsxPackageUpload) UnsetSpdxLicense()`

UnsetSpdxLicense ensures that no value is present for SpdxLicense, not even an explicit nil
### GetStage

`func (o *VsxPackageUpload) GetStage() int64`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *VsxPackageUpload) GetStageOk() (*int64, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *VsxPackageUpload) SetStage(v int64)`

SetStage sets Stage field to given value.

### HasStage

`func (o *VsxPackageUpload) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStageStr

`func (o *VsxPackageUpload) GetStageStr() string`

GetStageStr returns the StageStr field if non-nil, zero value otherwise.

### GetStageStrOk

`func (o *VsxPackageUpload) GetStageStrOk() (*string, bool)`

GetStageStrOk returns a tuple with the StageStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStr

`func (o *VsxPackageUpload) SetStageStr(v string)`

SetStageStr sets StageStr field to given value.

### HasStageStr

`func (o *VsxPackageUpload) HasStageStr() bool`

HasStageStr returns a boolean if a field has been set.

### GetStageUpdatedAt

`func (o *VsxPackageUpload) GetStageUpdatedAt() time.Time`

GetStageUpdatedAt returns the StageUpdatedAt field if non-nil, zero value otherwise.

### GetStageUpdatedAtOk

`func (o *VsxPackageUpload) GetStageUpdatedAtOk() (*time.Time, bool)`

GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUpdatedAt

`func (o *VsxPackageUpload) SetStageUpdatedAt(v time.Time)`

SetStageUpdatedAt sets StageUpdatedAt field to given value.

### HasStageUpdatedAt

`func (o *VsxPackageUpload) HasStageUpdatedAt() bool`

HasStageUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *VsxPackageUpload) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VsxPackageUpload) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VsxPackageUpload) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VsxPackageUpload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *VsxPackageUpload) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *VsxPackageUpload) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *VsxPackageUpload) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *VsxPackageUpload) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *VsxPackageUpload) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *VsxPackageUpload) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetStatusStr

`func (o *VsxPackageUpload) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *VsxPackageUpload) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *VsxPackageUpload) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *VsxPackageUpload) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *VsxPackageUpload) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *VsxPackageUpload) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *VsxPackageUpload) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *VsxPackageUpload) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStatusUrl

`func (o *VsxPackageUpload) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *VsxPackageUpload) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *VsxPackageUpload) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *VsxPackageUpload) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSubtype

`func (o *VsxPackageUpload) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *VsxPackageUpload) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *VsxPackageUpload) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *VsxPackageUpload) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSummary

`func (o *VsxPackageUpload) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *VsxPackageUpload) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *VsxPackageUpload) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *VsxPackageUpload) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *VsxPackageUpload) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *VsxPackageUpload) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetSyncFinishedAt

`func (o *VsxPackageUpload) GetSyncFinishedAt() time.Time`

GetSyncFinishedAt returns the SyncFinishedAt field if non-nil, zero value otherwise.

### GetSyncFinishedAtOk

`func (o *VsxPackageUpload) GetSyncFinishedAtOk() (*time.Time, bool)`

GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFinishedAt

`func (o *VsxPackageUpload) SetSyncFinishedAt(v time.Time)`

SetSyncFinishedAt sets SyncFinishedAt field to given value.

### HasSyncFinishedAt

`func (o *VsxPackageUpload) HasSyncFinishedAt() bool`

HasSyncFinishedAt returns a boolean if a field has been set.

### SetSyncFinishedAtNil

`func (o *VsxPackageUpload) SetSyncFinishedAtNil(b bool)`

 SetSyncFinishedAtNil sets the value for SyncFinishedAt to be an explicit nil

### UnsetSyncFinishedAt
`func (o *VsxPackageUpload) UnsetSyncFinishedAt()`

UnsetSyncFinishedAt ensures that no value is present for SyncFinishedAt, not even an explicit nil
### GetSyncProgress

`func (o *VsxPackageUpload) GetSyncProgress() int64`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *VsxPackageUpload) GetSyncProgressOk() (*int64, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *VsxPackageUpload) SetSyncProgress(v int64)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *VsxPackageUpload) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.

### GetTagsAutomatic

`func (o *VsxPackageUpload) GetTagsAutomatic() map[string]interface{}`

GetTagsAutomatic returns the TagsAutomatic field if non-nil, zero value otherwise.

### GetTagsAutomaticOk

`func (o *VsxPackageUpload) GetTagsAutomaticOk() (*map[string]interface{}, bool)`

GetTagsAutomaticOk returns a tuple with the TagsAutomatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsAutomatic

`func (o *VsxPackageUpload) SetTagsAutomatic(v map[string]interface{})`

SetTagsAutomatic sets TagsAutomatic field to given value.

### HasTagsAutomatic

`func (o *VsxPackageUpload) HasTagsAutomatic() bool`

HasTagsAutomatic returns a boolean if a field has been set.

### GetTagsImmutable

`func (o *VsxPackageUpload) GetTagsImmutable() map[string]interface{}`

GetTagsImmutable returns the TagsImmutable field if non-nil, zero value otherwise.

### GetTagsImmutableOk

`func (o *VsxPackageUpload) GetTagsImmutableOk() (*map[string]interface{}, bool)`

GetTagsImmutableOk returns a tuple with the TagsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsImmutable

`func (o *VsxPackageUpload) SetTagsImmutable(v map[string]interface{})`

SetTagsImmutable sets TagsImmutable field to given value.

### HasTagsImmutable

`func (o *VsxPackageUpload) HasTagsImmutable() bool`

HasTagsImmutable returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *VsxPackageUpload) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *VsxPackageUpload) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *VsxPackageUpload) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *VsxPackageUpload) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *VsxPackageUpload) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *VsxPackageUpload) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *VsxPackageUpload) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *VsxPackageUpload) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *VsxPackageUpload) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *VsxPackageUpload) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *VsxPackageUpload) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *VsxPackageUpload) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderUrl

`func (o *VsxPackageUpload) GetUploaderUrl() string`

GetUploaderUrl returns the UploaderUrl field if non-nil, zero value otherwise.

### GetUploaderUrlOk

`func (o *VsxPackageUpload) GetUploaderUrlOk() (*string, bool)`

GetUploaderUrlOk returns a tuple with the UploaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderUrl

`func (o *VsxPackageUpload) SetUploaderUrl(v string)`

SetUploaderUrl sets UploaderUrl field to given value.

### HasUploaderUrl

`func (o *VsxPackageUpload) HasUploaderUrl() bool`

HasUploaderUrl returns a boolean if a field has been set.

### GetVersion

`func (o *VsxPackageUpload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VsxPackageUpload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VsxPackageUpload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VsxPackageUpload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *VsxPackageUpload) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *VsxPackageUpload) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVersionOrig

`func (o *VsxPackageUpload) GetVersionOrig() string`

GetVersionOrig returns the VersionOrig field if non-nil, zero value otherwise.

### GetVersionOrigOk

`func (o *VsxPackageUpload) GetVersionOrigOk() (*string, bool)`

GetVersionOrigOk returns a tuple with the VersionOrig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrig

`func (o *VsxPackageUpload) SetVersionOrig(v string)`

SetVersionOrig sets VersionOrig field to given value.

### HasVersionOrig

`func (o *VsxPackageUpload) HasVersionOrig() bool`

HasVersionOrig returns a boolean if a field has been set.

### GetVulnerabilityScanResultsUrl

`func (o *VsxPackageUpload) GetVulnerabilityScanResultsUrl() string`

GetVulnerabilityScanResultsUrl returns the VulnerabilityScanResultsUrl field if non-nil, zero value otherwise.

### GetVulnerabilityScanResultsUrlOk

`func (o *VsxPackageUpload) GetVulnerabilityScanResultsUrlOk() (*string, bool)`

GetVulnerabilityScanResultsUrlOk returns a tuple with the VulnerabilityScanResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanResultsUrl

`func (o *VsxPackageUpload) SetVulnerabilityScanResultsUrl(v string)`

SetVulnerabilityScanResultsUrl sets VulnerabilityScanResultsUrl field to given value.

### HasVulnerabilityScanResultsUrl

`func (o *VsxPackageUpload) HasVulnerabilityScanResultsUrl() bool`

HasVulnerabilityScanResultsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


