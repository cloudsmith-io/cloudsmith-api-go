# RepositoryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdnUrl** | Pointer to **NullableString** | Base URL from which packages and other artifacts are downloaded. | [optional] [readonly] 
**ContentKind** | Pointer to **string** | The repository content kind determines whether this repository contains packages, or provides a distribution of packages from other repositories. You can only select the content kind at repository creation time. | [optional] [default to "Standard"]
**ContextualAuthRealm** | Pointer to **bool** | If checked, missing credentials for this repository where basic authentication is required shall present an enriched value in the &#39;WWW-Authenticate&#39; header containing the namespace and repository. This can be useful for tooling such as SBT where the authentication realm is used to distinguish and disambiguate credentials. | [optional] 
**CopyOwn** | Pointer to **bool** | If checked, users can copy any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**CopyPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to copy packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific copy setting. | [optional] [default to "Read"]
**CosignSigningEnabled** | Pointer to **bool** | When enabled, all pushed (or pulled from upstream) OCI packages and artifacts will be signed using cosign with the repository&#39;s ECDSA key. This generates a distinct cosign signature artifact per artifact. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**DefaultPrivilege** | Pointer to **string** | This defines the default level of privilege that all of your organization members have for this repository. This does not include collaborators, but applies to any member of the org regardless of their own membership role (i.e. it applies to owners, managers and members). Be careful if setting this to admin, because any member will be able to change settings. | [optional] [default to "None"]
**DeleteOwn** | Pointer to **bool** | If checked, users can delete any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**DeletePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to delete packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific delete setting. | [optional] [default to "Admin"]
**DeletedAt** | Pointer to **NullableTime** | The datetime the repository was manually deleted at. | [optional] [readonly] 
**Description** | Pointer to **string** | A description of the repository&#39;s purpose/contents. | [optional] 
**Distributes** | Pointer to **[]string** | The repositories distributed through this repo. Adding repos here is only valid if the content_kind is DISTRIBUTION. | [optional] 
**DockerRefreshTokensEnabled** | Pointer to **bool** | If checked, refresh tokens will be issued in addition to access tokens for Docker authentication. This allows unlimited extension of the lifetime of access tokens. | [optional] 
**EcdsaKeys** | Pointer to [**[]RepositoryEcdsaKey**](RepositoryEcdsaKey.md) |  | [optional] [readonly] 
**EnforceEula** | Pointer to **bool** | If checked, downloads will explicitly require acceptance of an EULA. | [optional] 
**GpgKeys** | Pointer to [**[]RepositoryGpgKey**](RepositoryGpgKey.md) |  | [optional] [readonly] 
**IndexFiles** | Pointer to **bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**IsOpenSource** | Pointer to **bool** |  | [optional] [readonly] 
**IsPrivate** | Pointer to **bool** |  | [optional] [readonly] 
**IsPublic** | Pointer to **bool** |  | [optional] [readonly] 
**ManageEntitlementsPrivilege** | Pointer to **string** | This defines the minimum level of privilege required for a user to manage entitlement tokens with private repositories. Management is the ability to create, alter, enable, disable or delete all tokens without a repository. | [optional] [default to "Admin"]
**MoveOwn** | Pointer to **bool** | If checked, users can move any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**MovePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to move packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific move setting. | [optional] [default to "Admin"]
**Name** | **string** | A descriptive name for the repository. | 
**Namespace** | Pointer to **string** | Namespace to which this repository belongs. | [optional] [readonly] 
**NamespaceUrl** | Pointer to **string** | API endpoint where data about this namespace can be retrieved. | [optional] [readonly] 
**NugetNativeSigningEnabled** | Pointer to **bool** | When enabled, all pushed (or pulled from upstream) nuget packages and artifacts will be signed using the repository&#39;s X.509 RSA certificate. Additionally, the nuget RepositorySignature index will list all of the repository&#39;s signing certificates including the ones from configured upstreams. | [optional] 
**NumDownloads** | Pointer to **int64** | The number of downloads for packages in the repository. | [optional] [readonly] 
**NumPolicyViolatedPackages** | Pointer to **int64** | Number of packages with policy violations in a repository. | [optional] [readonly] 
**NumQuarantinedPackages** | Pointer to **int64** | Number of quarantined packages in a repository. | [optional] [readonly] 
**OpenSourceLicense** | Pointer to **NullableString** | The SPDX identifier of the open source license. | [optional] 
**OpenSourceProjectUrl** | Pointer to **NullableString** | The URL to the Open-Source project, used for validating that the project meets the requirements for Open-Source. | [optional] 
**PackageCount** | Pointer to **int64** | The number of packages in the repository. | [optional] [readonly] 
**PackageGroupCount** | Pointer to **int64** | The number of groups in the repository. | [optional] [readonly] 
**ProxyNpmjs** | Pointer to **bool** | If checked, Npm packages that are not in the repository when requested by clients will automatically be proxied from the public npmjs.org registry. If there is at least one version for a package, others will not be proxied. | [optional] 
**ProxyPypi** | Pointer to **bool** | If checked, Python packages that are not in the repository when requested by clients will automatically be proxied from the public pypi.python.org registry. If there is at least one version for a package, others will not be proxied. | [optional] 
**RawPackageIndexEnabled** | Pointer to **bool** | If checked, HTML and JSON indexes will be generated that list all available raw packages in the repository. | [optional] 
**RawPackageIndexSignaturesEnabled** | Pointer to **bool** | If checked, the HTML and JSON indexes will display raw package GPG signatures alongside the index packages. | [optional] 
**ReplacePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to republish packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific republish setting. Please note that the user still requires the privilege to delete packages that will be replaced by the new package; otherwise the republish will fail. | [optional] [default to "Write"]
**ReplacePackagesByDefault** | Pointer to **bool** | If checked, uploaded packages will overwrite/replace any others with the same attributes (e.g. same version) by default. This only applies if the user has the required privilege for the republishing AND has the required privilege to delete existing packages that they don&#39;t own. | [optional] 
**RepositoryType** | Pointer to **int64** | The repository type changes how it is accessed and billed. Private repositories are visible only to you or authorized delegates. Open-Source repositories are always visible to everyone and are restricted by licensing, but are free to use and come with generous bandwidth/storage. You can only select Open-Source at repository creation time. | [optional] [readonly] 
**RepositoryTypeStr** | Pointer to **string** | The repository type changes how it is accessed and billed. Private repositories are visible only to you or authorized delegates. Public repositories are visible to all Cloudsmith users. | [optional] [default to "Public"]
**ResyncOwn** | Pointer to **bool** | If checked, users can resync any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**ResyncPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to resync packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific resync setting. | [optional] [default to "Admin"]
**ScanOwn** | Pointer to **bool** | If checked, users can scan any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**ScanPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to scan packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific scan setting. | [optional] [default to "Read"]
**SelfHtmlUrl** | Pointer to **string** | Website URL for this repository. | [optional] [readonly] 
**SelfUrl** | Pointer to **string** | API endpoint where data about this repository can be retrieved. | [optional] [readonly] 
**ShowSetupAll** | Pointer to **bool** | If checked, the Set Me Up help for all formats will always be shown, even if you don&#39;t have packages of that type uploaded. Otherwise, help will only be shown for packages that are in the repository. For example, if you have uploaded only NuGet packages, then the Set Me Up help for NuGet packages will be shown only. | [optional] 
**Size** | Pointer to **int64** | The calculated size of the repository. | [optional] [readonly] 
**SizeStr** | Pointer to **string** | The calculated size of the repository (human readable). | [optional] [readonly] 
**Slug** | Pointer to **string** | The slug identifies the repository in URIs. | [optional] 
**SlugPerm** | Pointer to **string** | The slug_perm immutably identifies the repository. It will never change once a repository has been created. | [optional] [readonly] 
**StorageRegion** | Pointer to **string** | The Cloudsmith region in which package files are stored. | [optional] [default to "default"]
**StrictNpmValidation** | Pointer to **bool** | If checked, npm packages will be validated strictly to ensure the package matches specifcation. You can turn this on if you want to guarantee that the packages will work with npm-cli and other tools correctly. | [optional] 
**TagPreReleasesAsLatest** | Pointer to **bool** | If checked, packages pushed with a pre-release component on that version will be marked with the &#39;latest&#39; tag. Note that if unchecked, a repository containing ONLY pre-release versions, will have no version marked latest which may cause incompatibility with native tools  | [optional] 
**UseDebianLabels** | Pointer to **bool** | If checked, a &#39;Label&#39; field will be present in Debian-based repositories. It will contain a string that identifies the entitlement token used to authenticate the repository, in the form of &#39;source&#x3D;t-&lt;identifier&gt;&#39;; or &#39;source&#x3D;none&#39; if no token was used. You can use this to help with pinning. | [optional] 
**UseDefaultCargoUpstream** | Pointer to **bool** | If checked, dependencies of uploaded Cargo crates which do not set an explicit value for \&quot;registry\&quot; will be assumed to be available from crates.io. If unchecked, dependencies with unspecified \&quot;registry\&quot; values will be assumed to be available in the registry being uploaded to. Uncheck this if you want to ensure that dependencies are only ever installed from Cloudsmith unless explicitly specified as belong to another registry. | [optional] 
**UseEntitlementsPrivilege** | Pointer to **string** | This defines the minimum level of privilege required for a user to see/use entitlement tokens with private repositories. If a user does not have the permission, they will only be able to download packages using other credentials, such as email/password via basic authentication. Use this if you want to force users to only use their user-based token, which is tied to their access (if removed, they can&#39;t use it). | [optional] [default to "Read"]
**UseNoarchPackages** | Pointer to **bool** | If checked, noarch packages (if supported) are enabled in installations/configurations. A noarch package is one that is not tied to specific system architecture (like i686). | [optional] 
**UseSourcePackages** | Pointer to **bool** | If checked, source packages (if supported) are enabled in installations/configurations. A source package is one that contains source code rather than built binaries. | [optional] 
**UseVulnerabilityScanning** | Pointer to **bool** | If checked, vulnerability scanning will be enabled for all supported packages within this repository. | [optional] 
**UserEntitlementsEnabled** | Pointer to **bool** | If checked, users can use and manage their own user-specific entitlement token for the repository (if private). Otherwise, user-specific entitlements are disabled for all users. | [optional] 
**ViewStatistics** | Pointer to **string** | This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applicable. If a user does not have the permission, they won&#39;t be able to view any statistics, either via the UI, API or CLI. | [optional] [default to "Read"]

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

### SetCdnUrlNil

`func (o *RepositoryCreate) SetCdnUrlNil(b bool)`

 SetCdnUrlNil sets the value for CdnUrl to be an explicit nil

### UnsetCdnUrl
`func (o *RepositoryCreate) UnsetCdnUrl()`

UnsetCdnUrl ensures that no value is present for CdnUrl, not even an explicit nil
### GetContentKind

`func (o *RepositoryCreate) GetContentKind() string`

GetContentKind returns the ContentKind field if non-nil, zero value otherwise.

### GetContentKindOk

`func (o *RepositoryCreate) GetContentKindOk() (*string, bool)`

GetContentKindOk returns a tuple with the ContentKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentKind

`func (o *RepositoryCreate) SetContentKind(v string)`

SetContentKind sets ContentKind field to given value.

### HasContentKind

`func (o *RepositoryCreate) HasContentKind() bool`

HasContentKind returns a boolean if a field has been set.

### GetContextualAuthRealm

`func (o *RepositoryCreate) GetContextualAuthRealm() bool`

GetContextualAuthRealm returns the ContextualAuthRealm field if non-nil, zero value otherwise.

### GetContextualAuthRealmOk

`func (o *RepositoryCreate) GetContextualAuthRealmOk() (*bool, bool)`

GetContextualAuthRealmOk returns a tuple with the ContextualAuthRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualAuthRealm

`func (o *RepositoryCreate) SetContextualAuthRealm(v bool)`

SetContextualAuthRealm sets ContextualAuthRealm field to given value.

### HasContextualAuthRealm

`func (o *RepositoryCreate) HasContextualAuthRealm() bool`

HasContextualAuthRealm returns a boolean if a field has been set.

### GetCopyOwn

`func (o *RepositoryCreate) GetCopyOwn() bool`

GetCopyOwn returns the CopyOwn field if non-nil, zero value otherwise.

### GetCopyOwnOk

`func (o *RepositoryCreate) GetCopyOwnOk() (*bool, bool)`

GetCopyOwnOk returns a tuple with the CopyOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOwn

`func (o *RepositoryCreate) SetCopyOwn(v bool)`

SetCopyOwn sets CopyOwn field to given value.

### HasCopyOwn

`func (o *RepositoryCreate) HasCopyOwn() bool`

HasCopyOwn returns a boolean if a field has been set.

### GetCopyPackages

`func (o *RepositoryCreate) GetCopyPackages() string`

GetCopyPackages returns the CopyPackages field if non-nil, zero value otherwise.

### GetCopyPackagesOk

`func (o *RepositoryCreate) GetCopyPackagesOk() (*string, bool)`

GetCopyPackagesOk returns a tuple with the CopyPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPackages

`func (o *RepositoryCreate) SetCopyPackages(v string)`

SetCopyPackages sets CopyPackages field to given value.

### HasCopyPackages

`func (o *RepositoryCreate) HasCopyPackages() bool`

HasCopyPackages returns a boolean if a field has been set.

### GetCosignSigningEnabled

`func (o *RepositoryCreate) GetCosignSigningEnabled() bool`

GetCosignSigningEnabled returns the CosignSigningEnabled field if non-nil, zero value otherwise.

### GetCosignSigningEnabledOk

`func (o *RepositoryCreate) GetCosignSigningEnabledOk() (*bool, bool)`

GetCosignSigningEnabledOk returns a tuple with the CosignSigningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosignSigningEnabled

`func (o *RepositoryCreate) SetCosignSigningEnabled(v bool)`

SetCosignSigningEnabled sets CosignSigningEnabled field to given value.

### HasCosignSigningEnabled

`func (o *RepositoryCreate) HasCosignSigningEnabled() bool`

HasCosignSigningEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryCreate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryCreate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryCreate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultPrivilege

`func (o *RepositoryCreate) GetDefaultPrivilege() string`

GetDefaultPrivilege returns the DefaultPrivilege field if non-nil, zero value otherwise.

### GetDefaultPrivilegeOk

`func (o *RepositoryCreate) GetDefaultPrivilegeOk() (*string, bool)`

GetDefaultPrivilegeOk returns a tuple with the DefaultPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivilege

`func (o *RepositoryCreate) SetDefaultPrivilege(v string)`

SetDefaultPrivilege sets DefaultPrivilege field to given value.

### HasDefaultPrivilege

`func (o *RepositoryCreate) HasDefaultPrivilege() bool`

HasDefaultPrivilege returns a boolean if a field has been set.

### GetDeleteOwn

`func (o *RepositoryCreate) GetDeleteOwn() bool`

GetDeleteOwn returns the DeleteOwn field if non-nil, zero value otherwise.

### GetDeleteOwnOk

`func (o *RepositoryCreate) GetDeleteOwnOk() (*bool, bool)`

GetDeleteOwnOk returns a tuple with the DeleteOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOwn

`func (o *RepositoryCreate) SetDeleteOwn(v bool)`

SetDeleteOwn sets DeleteOwn field to given value.

### HasDeleteOwn

`func (o *RepositoryCreate) HasDeleteOwn() bool`

HasDeleteOwn returns a boolean if a field has been set.

### GetDeletePackages

`func (o *RepositoryCreate) GetDeletePackages() string`

GetDeletePackages returns the DeletePackages field if non-nil, zero value otherwise.

### GetDeletePackagesOk

`func (o *RepositoryCreate) GetDeletePackagesOk() (*string, bool)`

GetDeletePackagesOk returns a tuple with the DeletePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePackages

`func (o *RepositoryCreate) SetDeletePackages(v string)`

SetDeletePackages sets DeletePackages field to given value.

### HasDeletePackages

`func (o *RepositoryCreate) HasDeletePackages() bool`

HasDeletePackages returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RepositoryCreate) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RepositoryCreate) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RepositoryCreate) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RepositoryCreate) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *RepositoryCreate) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *RepositoryCreate) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
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

### GetDistributes

`func (o *RepositoryCreate) GetDistributes() []string`

GetDistributes returns the Distributes field if non-nil, zero value otherwise.

### GetDistributesOk

`func (o *RepositoryCreate) GetDistributesOk() (*[]string, bool)`

GetDistributesOk returns a tuple with the Distributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributes

`func (o *RepositoryCreate) SetDistributes(v []string)`

SetDistributes sets Distributes field to given value.

### HasDistributes

`func (o *RepositoryCreate) HasDistributes() bool`

HasDistributes returns a boolean if a field has been set.

### GetDockerRefreshTokensEnabled

`func (o *RepositoryCreate) GetDockerRefreshTokensEnabled() bool`

GetDockerRefreshTokensEnabled returns the DockerRefreshTokensEnabled field if non-nil, zero value otherwise.

### GetDockerRefreshTokensEnabledOk

`func (o *RepositoryCreate) GetDockerRefreshTokensEnabledOk() (*bool, bool)`

GetDockerRefreshTokensEnabledOk returns a tuple with the DockerRefreshTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRefreshTokensEnabled

`func (o *RepositoryCreate) SetDockerRefreshTokensEnabled(v bool)`

SetDockerRefreshTokensEnabled sets DockerRefreshTokensEnabled field to given value.

### HasDockerRefreshTokensEnabled

`func (o *RepositoryCreate) HasDockerRefreshTokensEnabled() bool`

HasDockerRefreshTokensEnabled returns a boolean if a field has been set.

### GetEcdsaKeys

`func (o *RepositoryCreate) GetEcdsaKeys() []RepositoryEcdsaKey`

GetEcdsaKeys returns the EcdsaKeys field if non-nil, zero value otherwise.

### GetEcdsaKeysOk

`func (o *RepositoryCreate) GetEcdsaKeysOk() (*[]RepositoryEcdsaKey, bool)`

GetEcdsaKeysOk returns a tuple with the EcdsaKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaKeys

`func (o *RepositoryCreate) SetEcdsaKeys(v []RepositoryEcdsaKey)`

SetEcdsaKeys sets EcdsaKeys field to given value.

### HasEcdsaKeys

`func (o *RepositoryCreate) HasEcdsaKeys() bool`

HasEcdsaKeys returns a boolean if a field has been set.

### GetEnforceEula

`func (o *RepositoryCreate) GetEnforceEula() bool`

GetEnforceEula returns the EnforceEula field if non-nil, zero value otherwise.

### GetEnforceEulaOk

`func (o *RepositoryCreate) GetEnforceEulaOk() (*bool, bool)`

GetEnforceEulaOk returns a tuple with the EnforceEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceEula

`func (o *RepositoryCreate) SetEnforceEula(v bool)`

SetEnforceEula sets EnforceEula field to given value.

### HasEnforceEula

`func (o *RepositoryCreate) HasEnforceEula() bool`

HasEnforceEula returns a boolean if a field has been set.

### GetGpgKeys

`func (o *RepositoryCreate) GetGpgKeys() []RepositoryGpgKey`

GetGpgKeys returns the GpgKeys field if non-nil, zero value otherwise.

### GetGpgKeysOk

`func (o *RepositoryCreate) GetGpgKeysOk() (*[]RepositoryGpgKey, bool)`

GetGpgKeysOk returns a tuple with the GpgKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeys

`func (o *RepositoryCreate) SetGpgKeys(v []RepositoryGpgKey)`

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

### GetManageEntitlementsPrivilege

`func (o *RepositoryCreate) GetManageEntitlementsPrivilege() string`

GetManageEntitlementsPrivilege returns the ManageEntitlementsPrivilege field if non-nil, zero value otherwise.

### GetManageEntitlementsPrivilegeOk

`func (o *RepositoryCreate) GetManageEntitlementsPrivilegeOk() (*string, bool)`

GetManageEntitlementsPrivilegeOk returns a tuple with the ManageEntitlementsPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageEntitlementsPrivilege

`func (o *RepositoryCreate) SetManageEntitlementsPrivilege(v string)`

SetManageEntitlementsPrivilege sets ManageEntitlementsPrivilege field to given value.

### HasManageEntitlementsPrivilege

`func (o *RepositoryCreate) HasManageEntitlementsPrivilege() bool`

HasManageEntitlementsPrivilege returns a boolean if a field has been set.

### GetMoveOwn

`func (o *RepositoryCreate) GetMoveOwn() bool`

GetMoveOwn returns the MoveOwn field if non-nil, zero value otherwise.

### GetMoveOwnOk

`func (o *RepositoryCreate) GetMoveOwnOk() (*bool, bool)`

GetMoveOwnOk returns a tuple with the MoveOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveOwn

`func (o *RepositoryCreate) SetMoveOwn(v bool)`

SetMoveOwn sets MoveOwn field to given value.

### HasMoveOwn

`func (o *RepositoryCreate) HasMoveOwn() bool`

HasMoveOwn returns a boolean if a field has been set.

### GetMovePackages

`func (o *RepositoryCreate) GetMovePackages() string`

GetMovePackages returns the MovePackages field if non-nil, zero value otherwise.

### GetMovePackagesOk

`func (o *RepositoryCreate) GetMovePackagesOk() (*string, bool)`

GetMovePackagesOk returns a tuple with the MovePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovePackages

`func (o *RepositoryCreate) SetMovePackages(v string)`

SetMovePackages sets MovePackages field to given value.

### HasMovePackages

`func (o *RepositoryCreate) HasMovePackages() bool`

HasMovePackages returns a boolean if a field has been set.

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

### GetNugetNativeSigningEnabled

`func (o *RepositoryCreate) GetNugetNativeSigningEnabled() bool`

GetNugetNativeSigningEnabled returns the NugetNativeSigningEnabled field if non-nil, zero value otherwise.

### GetNugetNativeSigningEnabledOk

`func (o *RepositoryCreate) GetNugetNativeSigningEnabledOk() (*bool, bool)`

GetNugetNativeSigningEnabledOk returns a tuple with the NugetNativeSigningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetNativeSigningEnabled

`func (o *RepositoryCreate) SetNugetNativeSigningEnabled(v bool)`

SetNugetNativeSigningEnabled sets NugetNativeSigningEnabled field to given value.

### HasNugetNativeSigningEnabled

`func (o *RepositoryCreate) HasNugetNativeSigningEnabled() bool`

HasNugetNativeSigningEnabled returns a boolean if a field has been set.

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

### GetNumPolicyViolatedPackages

`func (o *RepositoryCreate) GetNumPolicyViolatedPackages() int64`

GetNumPolicyViolatedPackages returns the NumPolicyViolatedPackages field if non-nil, zero value otherwise.

### GetNumPolicyViolatedPackagesOk

`func (o *RepositoryCreate) GetNumPolicyViolatedPackagesOk() (*int64, bool)`

GetNumPolicyViolatedPackagesOk returns a tuple with the NumPolicyViolatedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPolicyViolatedPackages

`func (o *RepositoryCreate) SetNumPolicyViolatedPackages(v int64)`

SetNumPolicyViolatedPackages sets NumPolicyViolatedPackages field to given value.

### HasNumPolicyViolatedPackages

`func (o *RepositoryCreate) HasNumPolicyViolatedPackages() bool`

HasNumPolicyViolatedPackages returns a boolean if a field has been set.

### GetNumQuarantinedPackages

`func (o *RepositoryCreate) GetNumQuarantinedPackages() int64`

GetNumQuarantinedPackages returns the NumQuarantinedPackages field if non-nil, zero value otherwise.

### GetNumQuarantinedPackagesOk

`func (o *RepositoryCreate) GetNumQuarantinedPackagesOk() (*int64, bool)`

GetNumQuarantinedPackagesOk returns a tuple with the NumQuarantinedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQuarantinedPackages

`func (o *RepositoryCreate) SetNumQuarantinedPackages(v int64)`

SetNumQuarantinedPackages sets NumQuarantinedPackages field to given value.

### HasNumQuarantinedPackages

`func (o *RepositoryCreate) HasNumQuarantinedPackages() bool`

HasNumQuarantinedPackages returns a boolean if a field has been set.

### GetOpenSourceLicense

`func (o *RepositoryCreate) GetOpenSourceLicense() string`

GetOpenSourceLicense returns the OpenSourceLicense field if non-nil, zero value otherwise.

### GetOpenSourceLicenseOk

`func (o *RepositoryCreate) GetOpenSourceLicenseOk() (*string, bool)`

GetOpenSourceLicenseOk returns a tuple with the OpenSourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSourceLicense

`func (o *RepositoryCreate) SetOpenSourceLicense(v string)`

SetOpenSourceLicense sets OpenSourceLicense field to given value.

### HasOpenSourceLicense

`func (o *RepositoryCreate) HasOpenSourceLicense() bool`

HasOpenSourceLicense returns a boolean if a field has been set.

### SetOpenSourceLicenseNil

`func (o *RepositoryCreate) SetOpenSourceLicenseNil(b bool)`

 SetOpenSourceLicenseNil sets the value for OpenSourceLicense to be an explicit nil

### UnsetOpenSourceLicense
`func (o *RepositoryCreate) UnsetOpenSourceLicense()`

UnsetOpenSourceLicense ensures that no value is present for OpenSourceLicense, not even an explicit nil
### GetOpenSourceProjectUrl

`func (o *RepositoryCreate) GetOpenSourceProjectUrl() string`

GetOpenSourceProjectUrl returns the OpenSourceProjectUrl field if non-nil, zero value otherwise.

### GetOpenSourceProjectUrlOk

`func (o *RepositoryCreate) GetOpenSourceProjectUrlOk() (*string, bool)`

GetOpenSourceProjectUrlOk returns a tuple with the OpenSourceProjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSourceProjectUrl

`func (o *RepositoryCreate) SetOpenSourceProjectUrl(v string)`

SetOpenSourceProjectUrl sets OpenSourceProjectUrl field to given value.

### HasOpenSourceProjectUrl

`func (o *RepositoryCreate) HasOpenSourceProjectUrl() bool`

HasOpenSourceProjectUrl returns a boolean if a field has been set.

### SetOpenSourceProjectUrlNil

`func (o *RepositoryCreate) SetOpenSourceProjectUrlNil(b bool)`

 SetOpenSourceProjectUrlNil sets the value for OpenSourceProjectUrl to be an explicit nil

### UnsetOpenSourceProjectUrl
`func (o *RepositoryCreate) UnsetOpenSourceProjectUrl()`

UnsetOpenSourceProjectUrl ensures that no value is present for OpenSourceProjectUrl, not even an explicit nil
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

### GetProxyNpmjs

`func (o *RepositoryCreate) GetProxyNpmjs() bool`

GetProxyNpmjs returns the ProxyNpmjs field if non-nil, zero value otherwise.

### GetProxyNpmjsOk

`func (o *RepositoryCreate) GetProxyNpmjsOk() (*bool, bool)`

GetProxyNpmjsOk returns a tuple with the ProxyNpmjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNpmjs

`func (o *RepositoryCreate) SetProxyNpmjs(v bool)`

SetProxyNpmjs sets ProxyNpmjs field to given value.

### HasProxyNpmjs

`func (o *RepositoryCreate) HasProxyNpmjs() bool`

HasProxyNpmjs returns a boolean if a field has been set.

### GetProxyPypi

`func (o *RepositoryCreate) GetProxyPypi() bool`

GetProxyPypi returns the ProxyPypi field if non-nil, zero value otherwise.

### GetProxyPypiOk

`func (o *RepositoryCreate) GetProxyPypiOk() (*bool, bool)`

GetProxyPypiOk returns a tuple with the ProxyPypi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPypi

`func (o *RepositoryCreate) SetProxyPypi(v bool)`

SetProxyPypi sets ProxyPypi field to given value.

### HasProxyPypi

`func (o *RepositoryCreate) HasProxyPypi() bool`

HasProxyPypi returns a boolean if a field has been set.

### GetRawPackageIndexEnabled

`func (o *RepositoryCreate) GetRawPackageIndexEnabled() bool`

GetRawPackageIndexEnabled returns the RawPackageIndexEnabled field if non-nil, zero value otherwise.

### GetRawPackageIndexEnabledOk

`func (o *RepositoryCreate) GetRawPackageIndexEnabledOk() (*bool, bool)`

GetRawPackageIndexEnabledOk returns a tuple with the RawPackageIndexEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPackageIndexEnabled

`func (o *RepositoryCreate) SetRawPackageIndexEnabled(v bool)`

SetRawPackageIndexEnabled sets RawPackageIndexEnabled field to given value.

### HasRawPackageIndexEnabled

`func (o *RepositoryCreate) HasRawPackageIndexEnabled() bool`

HasRawPackageIndexEnabled returns a boolean if a field has been set.

### GetRawPackageIndexSignaturesEnabled

`func (o *RepositoryCreate) GetRawPackageIndexSignaturesEnabled() bool`

GetRawPackageIndexSignaturesEnabled returns the RawPackageIndexSignaturesEnabled field if non-nil, zero value otherwise.

### GetRawPackageIndexSignaturesEnabledOk

`func (o *RepositoryCreate) GetRawPackageIndexSignaturesEnabledOk() (*bool, bool)`

GetRawPackageIndexSignaturesEnabledOk returns a tuple with the RawPackageIndexSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPackageIndexSignaturesEnabled

`func (o *RepositoryCreate) SetRawPackageIndexSignaturesEnabled(v bool)`

SetRawPackageIndexSignaturesEnabled sets RawPackageIndexSignaturesEnabled field to given value.

### HasRawPackageIndexSignaturesEnabled

`func (o *RepositoryCreate) HasRawPackageIndexSignaturesEnabled() bool`

HasRawPackageIndexSignaturesEnabled returns a boolean if a field has been set.

### GetReplacePackages

`func (o *RepositoryCreate) GetReplacePackages() string`

GetReplacePackages returns the ReplacePackages field if non-nil, zero value otherwise.

### GetReplacePackagesOk

`func (o *RepositoryCreate) GetReplacePackagesOk() (*string, bool)`

GetReplacePackagesOk returns a tuple with the ReplacePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePackages

`func (o *RepositoryCreate) SetReplacePackages(v string)`

SetReplacePackages sets ReplacePackages field to given value.

### HasReplacePackages

`func (o *RepositoryCreate) HasReplacePackages() bool`

HasReplacePackages returns a boolean if a field has been set.

### GetReplacePackagesByDefault

`func (o *RepositoryCreate) GetReplacePackagesByDefault() bool`

GetReplacePackagesByDefault returns the ReplacePackagesByDefault field if non-nil, zero value otherwise.

### GetReplacePackagesByDefaultOk

`func (o *RepositoryCreate) GetReplacePackagesByDefaultOk() (*bool, bool)`

GetReplacePackagesByDefaultOk returns a tuple with the ReplacePackagesByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePackagesByDefault

`func (o *RepositoryCreate) SetReplacePackagesByDefault(v bool)`

SetReplacePackagesByDefault sets ReplacePackagesByDefault field to given value.

### HasReplacePackagesByDefault

`func (o *RepositoryCreate) HasReplacePackagesByDefault() bool`

HasReplacePackagesByDefault returns a boolean if a field has been set.

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

### GetResyncOwn

`func (o *RepositoryCreate) GetResyncOwn() bool`

GetResyncOwn returns the ResyncOwn field if non-nil, zero value otherwise.

### GetResyncOwnOk

`func (o *RepositoryCreate) GetResyncOwnOk() (*bool, bool)`

GetResyncOwnOk returns a tuple with the ResyncOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncOwn

`func (o *RepositoryCreate) SetResyncOwn(v bool)`

SetResyncOwn sets ResyncOwn field to given value.

### HasResyncOwn

`func (o *RepositoryCreate) HasResyncOwn() bool`

HasResyncOwn returns a boolean if a field has been set.

### GetResyncPackages

`func (o *RepositoryCreate) GetResyncPackages() string`

GetResyncPackages returns the ResyncPackages field if non-nil, zero value otherwise.

### GetResyncPackagesOk

`func (o *RepositoryCreate) GetResyncPackagesOk() (*string, bool)`

GetResyncPackagesOk returns a tuple with the ResyncPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncPackages

`func (o *RepositoryCreate) SetResyncPackages(v string)`

SetResyncPackages sets ResyncPackages field to given value.

### HasResyncPackages

`func (o *RepositoryCreate) HasResyncPackages() bool`

HasResyncPackages returns a boolean if a field has been set.

### GetScanOwn

`func (o *RepositoryCreate) GetScanOwn() bool`

GetScanOwn returns the ScanOwn field if non-nil, zero value otherwise.

### GetScanOwnOk

`func (o *RepositoryCreate) GetScanOwnOk() (*bool, bool)`

GetScanOwnOk returns a tuple with the ScanOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOwn

`func (o *RepositoryCreate) SetScanOwn(v bool)`

SetScanOwn sets ScanOwn field to given value.

### HasScanOwn

`func (o *RepositoryCreate) HasScanOwn() bool`

HasScanOwn returns a boolean if a field has been set.

### GetScanPackages

`func (o *RepositoryCreate) GetScanPackages() string`

GetScanPackages returns the ScanPackages field if non-nil, zero value otherwise.

### GetScanPackagesOk

`func (o *RepositoryCreate) GetScanPackagesOk() (*string, bool)`

GetScanPackagesOk returns a tuple with the ScanPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPackages

`func (o *RepositoryCreate) SetScanPackages(v string)`

SetScanPackages sets ScanPackages field to given value.

### HasScanPackages

`func (o *RepositoryCreate) HasScanPackages() bool`

HasScanPackages returns a boolean if a field has been set.

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

### GetShowSetupAll

`func (o *RepositoryCreate) GetShowSetupAll() bool`

GetShowSetupAll returns the ShowSetupAll field if non-nil, zero value otherwise.

### GetShowSetupAllOk

`func (o *RepositoryCreate) GetShowSetupAllOk() (*bool, bool)`

GetShowSetupAllOk returns a tuple with the ShowSetupAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSetupAll

`func (o *RepositoryCreate) SetShowSetupAll(v bool)`

SetShowSetupAll sets ShowSetupAll field to given value.

### HasShowSetupAll

`func (o *RepositoryCreate) HasShowSetupAll() bool`

HasShowSetupAll returns a boolean if a field has been set.

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

### GetStrictNpmValidation

`func (o *RepositoryCreate) GetStrictNpmValidation() bool`

GetStrictNpmValidation returns the StrictNpmValidation field if non-nil, zero value otherwise.

### GetStrictNpmValidationOk

`func (o *RepositoryCreate) GetStrictNpmValidationOk() (*bool, bool)`

GetStrictNpmValidationOk returns a tuple with the StrictNpmValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictNpmValidation

`func (o *RepositoryCreate) SetStrictNpmValidation(v bool)`

SetStrictNpmValidation sets StrictNpmValidation field to given value.

### HasStrictNpmValidation

`func (o *RepositoryCreate) HasStrictNpmValidation() bool`

HasStrictNpmValidation returns a boolean if a field has been set.

### GetTagPreReleasesAsLatest

`func (o *RepositoryCreate) GetTagPreReleasesAsLatest() bool`

GetTagPreReleasesAsLatest returns the TagPreReleasesAsLatest field if non-nil, zero value otherwise.

### GetTagPreReleasesAsLatestOk

`func (o *RepositoryCreate) GetTagPreReleasesAsLatestOk() (*bool, bool)`

GetTagPreReleasesAsLatestOk returns a tuple with the TagPreReleasesAsLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPreReleasesAsLatest

`func (o *RepositoryCreate) SetTagPreReleasesAsLatest(v bool)`

SetTagPreReleasesAsLatest sets TagPreReleasesAsLatest field to given value.

### HasTagPreReleasesAsLatest

`func (o *RepositoryCreate) HasTagPreReleasesAsLatest() bool`

HasTagPreReleasesAsLatest returns a boolean if a field has been set.

### GetUseDebianLabels

`func (o *RepositoryCreate) GetUseDebianLabels() bool`

GetUseDebianLabels returns the UseDebianLabels field if non-nil, zero value otherwise.

### GetUseDebianLabelsOk

`func (o *RepositoryCreate) GetUseDebianLabelsOk() (*bool, bool)`

GetUseDebianLabelsOk returns a tuple with the UseDebianLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDebianLabels

`func (o *RepositoryCreate) SetUseDebianLabels(v bool)`

SetUseDebianLabels sets UseDebianLabels field to given value.

### HasUseDebianLabels

`func (o *RepositoryCreate) HasUseDebianLabels() bool`

HasUseDebianLabels returns a boolean if a field has been set.

### GetUseDefaultCargoUpstream

`func (o *RepositoryCreate) GetUseDefaultCargoUpstream() bool`

GetUseDefaultCargoUpstream returns the UseDefaultCargoUpstream field if non-nil, zero value otherwise.

### GetUseDefaultCargoUpstreamOk

`func (o *RepositoryCreate) GetUseDefaultCargoUpstreamOk() (*bool, bool)`

GetUseDefaultCargoUpstreamOk returns a tuple with the UseDefaultCargoUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultCargoUpstream

`func (o *RepositoryCreate) SetUseDefaultCargoUpstream(v bool)`

SetUseDefaultCargoUpstream sets UseDefaultCargoUpstream field to given value.

### HasUseDefaultCargoUpstream

`func (o *RepositoryCreate) HasUseDefaultCargoUpstream() bool`

HasUseDefaultCargoUpstream returns a boolean if a field has been set.

### GetUseEntitlementsPrivilege

`func (o *RepositoryCreate) GetUseEntitlementsPrivilege() string`

GetUseEntitlementsPrivilege returns the UseEntitlementsPrivilege field if non-nil, zero value otherwise.

### GetUseEntitlementsPrivilegeOk

`func (o *RepositoryCreate) GetUseEntitlementsPrivilegeOk() (*string, bool)`

GetUseEntitlementsPrivilegeOk returns a tuple with the UseEntitlementsPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEntitlementsPrivilege

`func (o *RepositoryCreate) SetUseEntitlementsPrivilege(v string)`

SetUseEntitlementsPrivilege sets UseEntitlementsPrivilege field to given value.

### HasUseEntitlementsPrivilege

`func (o *RepositoryCreate) HasUseEntitlementsPrivilege() bool`

HasUseEntitlementsPrivilege returns a boolean if a field has been set.

### GetUseNoarchPackages

`func (o *RepositoryCreate) GetUseNoarchPackages() bool`

GetUseNoarchPackages returns the UseNoarchPackages field if non-nil, zero value otherwise.

### GetUseNoarchPackagesOk

`func (o *RepositoryCreate) GetUseNoarchPackagesOk() (*bool, bool)`

GetUseNoarchPackagesOk returns a tuple with the UseNoarchPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNoarchPackages

`func (o *RepositoryCreate) SetUseNoarchPackages(v bool)`

SetUseNoarchPackages sets UseNoarchPackages field to given value.

### HasUseNoarchPackages

`func (o *RepositoryCreate) HasUseNoarchPackages() bool`

HasUseNoarchPackages returns a boolean if a field has been set.

### GetUseSourcePackages

`func (o *RepositoryCreate) GetUseSourcePackages() bool`

GetUseSourcePackages returns the UseSourcePackages field if non-nil, zero value otherwise.

### GetUseSourcePackagesOk

`func (o *RepositoryCreate) GetUseSourcePackagesOk() (*bool, bool)`

GetUseSourcePackagesOk returns a tuple with the UseSourcePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSourcePackages

`func (o *RepositoryCreate) SetUseSourcePackages(v bool)`

SetUseSourcePackages sets UseSourcePackages field to given value.

### HasUseSourcePackages

`func (o *RepositoryCreate) HasUseSourcePackages() bool`

HasUseSourcePackages returns a boolean if a field has been set.

### GetUseVulnerabilityScanning

`func (o *RepositoryCreate) GetUseVulnerabilityScanning() bool`

GetUseVulnerabilityScanning returns the UseVulnerabilityScanning field if non-nil, zero value otherwise.

### GetUseVulnerabilityScanningOk

`func (o *RepositoryCreate) GetUseVulnerabilityScanningOk() (*bool, bool)`

GetUseVulnerabilityScanningOk returns a tuple with the UseVulnerabilityScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVulnerabilityScanning

`func (o *RepositoryCreate) SetUseVulnerabilityScanning(v bool)`

SetUseVulnerabilityScanning sets UseVulnerabilityScanning field to given value.

### HasUseVulnerabilityScanning

`func (o *RepositoryCreate) HasUseVulnerabilityScanning() bool`

HasUseVulnerabilityScanning returns a boolean if a field has been set.

### GetUserEntitlementsEnabled

`func (o *RepositoryCreate) GetUserEntitlementsEnabled() bool`

GetUserEntitlementsEnabled returns the UserEntitlementsEnabled field if non-nil, zero value otherwise.

### GetUserEntitlementsEnabledOk

`func (o *RepositoryCreate) GetUserEntitlementsEnabledOk() (*bool, bool)`

GetUserEntitlementsEnabledOk returns a tuple with the UserEntitlementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEntitlementsEnabled

`func (o *RepositoryCreate) SetUserEntitlementsEnabled(v bool)`

SetUserEntitlementsEnabled sets UserEntitlementsEnabled field to given value.

### HasUserEntitlementsEnabled

`func (o *RepositoryCreate) HasUserEntitlementsEnabled() bool`

HasUserEntitlementsEnabled returns a boolean if a field has been set.

### GetViewStatistics

`func (o *RepositoryCreate) GetViewStatistics() string`

GetViewStatistics returns the ViewStatistics field if non-nil, zero value otherwise.

### GetViewStatisticsOk

`func (o *RepositoryCreate) GetViewStatisticsOk() (*string, bool)`

GetViewStatisticsOk returns a tuple with the ViewStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStatistics

`func (o *RepositoryCreate) SetViewStatistics(v string)`

SetViewStatistics sets ViewStatistics field to given value.

### HasViewStatistics

`func (o *RepositoryCreate) HasViewStatistics() bool`

HasViewStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


