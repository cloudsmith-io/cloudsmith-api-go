# RepositoryCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentKind** | Pointer to **string** | The repository content kind determines whether this repository contains packages, or provides a distribution of packages from other repositories. You can only select the content kind at repository creation time. | [optional] [default to "Standard"]
**ContextualAuthRealm** | Pointer to **bool** | If checked, missing credentials for this repository where basic authentication is required shall present an enriched value in the &#39;WWW-Authenticate&#39; header containing the namespace and repository. This can be useful for tooling such as SBT where the authentication realm is used to distinguish and disambiguate credentials. | [optional] 
**CopyOwn** | Pointer to **bool** | If checked, users can copy any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**CopyPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to copy packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific copy setting. | [optional] [default to "Read"]
**DefaultPrivilege** | Pointer to **string** | This defines the default level of privilege that all of your organization members have for this repository. This does not include collaborators, but applies to any member of the org regardless of their own membership role (i.e. it applies to owners, managers and members). Be careful if setting this to admin, because any member will be able to change settings. | [optional] [default to "None"]
**DeleteOwn** | Pointer to **bool** | If checked, users can delete any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**DeletePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to delete packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific delete setting. | [optional] [default to "Admin"]
**Description** | Pointer to **string** | A description of the repository&#39;s purpose/contents. | [optional] 
**Distributes** | Pointer to **[]string** | The repositories distributed through this repo. Adding repos here is only valid if the content_kind is DISTRIBUTION. | [optional] 
**DockerRefreshTokensEnabled** | Pointer to **bool** | If checked, refresh tokens will be issued in addition to access tokens for Docker authentication. This allows unlimited extension of the lifetime of access tokens. | [optional] 
**EnforceEula** | Pointer to **bool** | If checked, downloads will explicitly require acceptance of an EULA. | [optional] 
**IndexFiles** | Pointer to **bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**ManageEntitlementsPrivilege** | Pointer to **string** | This defines the minimum level of privilege required for a user to manage entitlement tokens with private repositories. Management is the ability to create, alter, enable, disable or delete all tokens without a repository. | [optional] [default to "Admin"]
**MoveOwn** | Pointer to **bool** | If checked, users can move any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**MovePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to move packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific move setting. | [optional] [default to "Admin"]
**Name** | **string** | A descriptive name for the repository. | 
**ProxyNpmjs** | Pointer to **bool** | If checked, Npm packages that are not in the repository when requested by clients will automatically be proxied from the public npmjs.org registry. If there is at least one version for a package, others will not be proxied. | [optional] 
**ProxyPypi** | Pointer to **bool** | If checked, Python packages that are not in the repository when requested by clients will automatically be proxied from the public pypi.python.org registry. If there is at least one version for a package, others will not be proxied. | [optional] 
**RawPackageIndexEnabled** | Pointer to **bool** | If checked, HTML and JSON indexes will be generated that list all available raw packages in the repository. | [optional] 
**RawPackageIndexSignaturesEnabled** | Pointer to **bool** | If checked, the HTML and JSON indexes will display raw package GPG signatures alongside the index packages. | [optional] 
**ReplacePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to republish packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific republish setting. Please note that the user still requires the privilege to delete packages that will be replaced by the new package; otherwise the republish will fail. | [optional] [default to "Write"]
**ReplacePackagesByDefault** | Pointer to **bool** | If checked, uploaded packages will overwrite/replace any others with the same attributes (e.g. same version) by default. This only applies if the user has the required privilege for the republishing AND has the required privilege to delete existing packages that they don&#39;t own. | [optional] 
**RepositoryTypeStr** | Pointer to **string** | The repository type changes how it is accessed and billed. Private repositories are visible only to you or authorized delegates. Public repositories are visible to all Cloudsmith users. | [optional] [default to "Public"]
**ResyncOwn** | Pointer to **bool** | If checked, users can resync any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**ResyncPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to resync packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific resync setting. | [optional] [default to "Admin"]
**ScanOwn** | Pointer to **bool** | If checked, users can scan any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**ScanPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to scan packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific scan setting. | [optional] [default to "Read"]
**ShowSetupAll** | Pointer to **bool** | If checked, the Set Me Up help for all formats will always be shown, even if you don&#39;t have packages of that type uploaded. Otherwise, help will only be shown for packages that are in the repository. For example, if you have uploaded only NuGet packages, then the Set Me Up help for NuGet packages will be shown only. | [optional] 
**Slug** | Pointer to **string** | The slug identifies the repository in URIs. | [optional] 
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

### NewRepositoryCreateRequest

`func NewRepositoryCreateRequest(name string, ) *RepositoryCreateRequest`

NewRepositoryCreateRequest instantiates a new RepositoryCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryCreateRequestWithDefaults

`func NewRepositoryCreateRequestWithDefaults() *RepositoryCreateRequest`

NewRepositoryCreateRequestWithDefaults instantiates a new RepositoryCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentKind

`func (o *RepositoryCreateRequest) GetContentKind() string`

GetContentKind returns the ContentKind field if non-nil, zero value otherwise.

### GetContentKindOk

`func (o *RepositoryCreateRequest) GetContentKindOk() (*string, bool)`

GetContentKindOk returns a tuple with the ContentKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentKind

`func (o *RepositoryCreateRequest) SetContentKind(v string)`

SetContentKind sets ContentKind field to given value.

### HasContentKind

`func (o *RepositoryCreateRequest) HasContentKind() bool`

HasContentKind returns a boolean if a field has been set.

### GetContextualAuthRealm

`func (o *RepositoryCreateRequest) GetContextualAuthRealm() bool`

GetContextualAuthRealm returns the ContextualAuthRealm field if non-nil, zero value otherwise.

### GetContextualAuthRealmOk

`func (o *RepositoryCreateRequest) GetContextualAuthRealmOk() (*bool, bool)`

GetContextualAuthRealmOk returns a tuple with the ContextualAuthRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualAuthRealm

`func (o *RepositoryCreateRequest) SetContextualAuthRealm(v bool)`

SetContextualAuthRealm sets ContextualAuthRealm field to given value.

### HasContextualAuthRealm

`func (o *RepositoryCreateRequest) HasContextualAuthRealm() bool`

HasContextualAuthRealm returns a boolean if a field has been set.

### GetCopyOwn

`func (o *RepositoryCreateRequest) GetCopyOwn() bool`

GetCopyOwn returns the CopyOwn field if non-nil, zero value otherwise.

### GetCopyOwnOk

`func (o *RepositoryCreateRequest) GetCopyOwnOk() (*bool, bool)`

GetCopyOwnOk returns a tuple with the CopyOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOwn

`func (o *RepositoryCreateRequest) SetCopyOwn(v bool)`

SetCopyOwn sets CopyOwn field to given value.

### HasCopyOwn

`func (o *RepositoryCreateRequest) HasCopyOwn() bool`

HasCopyOwn returns a boolean if a field has been set.

### GetCopyPackages

`func (o *RepositoryCreateRequest) GetCopyPackages() string`

GetCopyPackages returns the CopyPackages field if non-nil, zero value otherwise.

### GetCopyPackagesOk

`func (o *RepositoryCreateRequest) GetCopyPackagesOk() (*string, bool)`

GetCopyPackagesOk returns a tuple with the CopyPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPackages

`func (o *RepositoryCreateRequest) SetCopyPackages(v string)`

SetCopyPackages sets CopyPackages field to given value.

### HasCopyPackages

`func (o *RepositoryCreateRequest) HasCopyPackages() bool`

HasCopyPackages returns a boolean if a field has been set.

### GetDefaultPrivilege

`func (o *RepositoryCreateRequest) GetDefaultPrivilege() string`

GetDefaultPrivilege returns the DefaultPrivilege field if non-nil, zero value otherwise.

### GetDefaultPrivilegeOk

`func (o *RepositoryCreateRequest) GetDefaultPrivilegeOk() (*string, bool)`

GetDefaultPrivilegeOk returns a tuple with the DefaultPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivilege

`func (o *RepositoryCreateRequest) SetDefaultPrivilege(v string)`

SetDefaultPrivilege sets DefaultPrivilege field to given value.

### HasDefaultPrivilege

`func (o *RepositoryCreateRequest) HasDefaultPrivilege() bool`

HasDefaultPrivilege returns a boolean if a field has been set.

### GetDeleteOwn

`func (o *RepositoryCreateRequest) GetDeleteOwn() bool`

GetDeleteOwn returns the DeleteOwn field if non-nil, zero value otherwise.

### GetDeleteOwnOk

`func (o *RepositoryCreateRequest) GetDeleteOwnOk() (*bool, bool)`

GetDeleteOwnOk returns a tuple with the DeleteOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOwn

`func (o *RepositoryCreateRequest) SetDeleteOwn(v bool)`

SetDeleteOwn sets DeleteOwn field to given value.

### HasDeleteOwn

`func (o *RepositoryCreateRequest) HasDeleteOwn() bool`

HasDeleteOwn returns a boolean if a field has been set.

### GetDeletePackages

`func (o *RepositoryCreateRequest) GetDeletePackages() string`

GetDeletePackages returns the DeletePackages field if non-nil, zero value otherwise.

### GetDeletePackagesOk

`func (o *RepositoryCreateRequest) GetDeletePackagesOk() (*string, bool)`

GetDeletePackagesOk returns a tuple with the DeletePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePackages

`func (o *RepositoryCreateRequest) SetDeletePackages(v string)`

SetDeletePackages sets DeletePackages field to given value.

### HasDeletePackages

`func (o *RepositoryCreateRequest) HasDeletePackages() bool`

HasDeletePackages returns a boolean if a field has been set.

### GetDescription

`func (o *RepositoryCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepositoryCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistributes

`func (o *RepositoryCreateRequest) GetDistributes() []string`

GetDistributes returns the Distributes field if non-nil, zero value otherwise.

### GetDistributesOk

`func (o *RepositoryCreateRequest) GetDistributesOk() (*[]string, bool)`

GetDistributesOk returns a tuple with the Distributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributes

`func (o *RepositoryCreateRequest) SetDistributes(v []string)`

SetDistributes sets Distributes field to given value.

### HasDistributes

`func (o *RepositoryCreateRequest) HasDistributes() bool`

HasDistributes returns a boolean if a field has been set.

### GetDockerRefreshTokensEnabled

`func (o *RepositoryCreateRequest) GetDockerRefreshTokensEnabled() bool`

GetDockerRefreshTokensEnabled returns the DockerRefreshTokensEnabled field if non-nil, zero value otherwise.

### GetDockerRefreshTokensEnabledOk

`func (o *RepositoryCreateRequest) GetDockerRefreshTokensEnabledOk() (*bool, bool)`

GetDockerRefreshTokensEnabledOk returns a tuple with the DockerRefreshTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRefreshTokensEnabled

`func (o *RepositoryCreateRequest) SetDockerRefreshTokensEnabled(v bool)`

SetDockerRefreshTokensEnabled sets DockerRefreshTokensEnabled field to given value.

### HasDockerRefreshTokensEnabled

`func (o *RepositoryCreateRequest) HasDockerRefreshTokensEnabled() bool`

HasDockerRefreshTokensEnabled returns a boolean if a field has been set.

### GetEnforceEula

`func (o *RepositoryCreateRequest) GetEnforceEula() bool`

GetEnforceEula returns the EnforceEula field if non-nil, zero value otherwise.

### GetEnforceEulaOk

`func (o *RepositoryCreateRequest) GetEnforceEulaOk() (*bool, bool)`

GetEnforceEulaOk returns a tuple with the EnforceEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceEula

`func (o *RepositoryCreateRequest) SetEnforceEula(v bool)`

SetEnforceEula sets EnforceEula field to given value.

### HasEnforceEula

`func (o *RepositoryCreateRequest) HasEnforceEula() bool`

HasEnforceEula returns a boolean if a field has been set.

### GetIndexFiles

`func (o *RepositoryCreateRequest) GetIndexFiles() bool`

GetIndexFiles returns the IndexFiles field if non-nil, zero value otherwise.

### GetIndexFilesOk

`func (o *RepositoryCreateRequest) GetIndexFilesOk() (*bool, bool)`

GetIndexFilesOk returns a tuple with the IndexFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFiles

`func (o *RepositoryCreateRequest) SetIndexFiles(v bool)`

SetIndexFiles sets IndexFiles field to given value.

### HasIndexFiles

`func (o *RepositoryCreateRequest) HasIndexFiles() bool`

HasIndexFiles returns a boolean if a field has been set.

### GetManageEntitlementsPrivilege

`func (o *RepositoryCreateRequest) GetManageEntitlementsPrivilege() string`

GetManageEntitlementsPrivilege returns the ManageEntitlementsPrivilege field if non-nil, zero value otherwise.

### GetManageEntitlementsPrivilegeOk

`func (o *RepositoryCreateRequest) GetManageEntitlementsPrivilegeOk() (*string, bool)`

GetManageEntitlementsPrivilegeOk returns a tuple with the ManageEntitlementsPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageEntitlementsPrivilege

`func (o *RepositoryCreateRequest) SetManageEntitlementsPrivilege(v string)`

SetManageEntitlementsPrivilege sets ManageEntitlementsPrivilege field to given value.

### HasManageEntitlementsPrivilege

`func (o *RepositoryCreateRequest) HasManageEntitlementsPrivilege() bool`

HasManageEntitlementsPrivilege returns a boolean if a field has been set.

### GetMoveOwn

`func (o *RepositoryCreateRequest) GetMoveOwn() bool`

GetMoveOwn returns the MoveOwn field if non-nil, zero value otherwise.

### GetMoveOwnOk

`func (o *RepositoryCreateRequest) GetMoveOwnOk() (*bool, bool)`

GetMoveOwnOk returns a tuple with the MoveOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveOwn

`func (o *RepositoryCreateRequest) SetMoveOwn(v bool)`

SetMoveOwn sets MoveOwn field to given value.

### HasMoveOwn

`func (o *RepositoryCreateRequest) HasMoveOwn() bool`

HasMoveOwn returns a boolean if a field has been set.

### GetMovePackages

`func (o *RepositoryCreateRequest) GetMovePackages() string`

GetMovePackages returns the MovePackages field if non-nil, zero value otherwise.

### GetMovePackagesOk

`func (o *RepositoryCreateRequest) GetMovePackagesOk() (*string, bool)`

GetMovePackagesOk returns a tuple with the MovePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovePackages

`func (o *RepositoryCreateRequest) SetMovePackages(v string)`

SetMovePackages sets MovePackages field to given value.

### HasMovePackages

`func (o *RepositoryCreateRequest) HasMovePackages() bool`

HasMovePackages returns a boolean if a field has been set.

### GetName

`func (o *RepositoryCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProxyNpmjs

`func (o *RepositoryCreateRequest) GetProxyNpmjs() bool`

GetProxyNpmjs returns the ProxyNpmjs field if non-nil, zero value otherwise.

### GetProxyNpmjsOk

`func (o *RepositoryCreateRequest) GetProxyNpmjsOk() (*bool, bool)`

GetProxyNpmjsOk returns a tuple with the ProxyNpmjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNpmjs

`func (o *RepositoryCreateRequest) SetProxyNpmjs(v bool)`

SetProxyNpmjs sets ProxyNpmjs field to given value.

### HasProxyNpmjs

`func (o *RepositoryCreateRequest) HasProxyNpmjs() bool`

HasProxyNpmjs returns a boolean if a field has been set.

### GetProxyPypi

`func (o *RepositoryCreateRequest) GetProxyPypi() bool`

GetProxyPypi returns the ProxyPypi field if non-nil, zero value otherwise.

### GetProxyPypiOk

`func (o *RepositoryCreateRequest) GetProxyPypiOk() (*bool, bool)`

GetProxyPypiOk returns a tuple with the ProxyPypi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPypi

`func (o *RepositoryCreateRequest) SetProxyPypi(v bool)`

SetProxyPypi sets ProxyPypi field to given value.

### HasProxyPypi

`func (o *RepositoryCreateRequest) HasProxyPypi() bool`

HasProxyPypi returns a boolean if a field has been set.

### GetRawPackageIndexEnabled

`func (o *RepositoryCreateRequest) GetRawPackageIndexEnabled() bool`

GetRawPackageIndexEnabled returns the RawPackageIndexEnabled field if non-nil, zero value otherwise.

### GetRawPackageIndexEnabledOk

`func (o *RepositoryCreateRequest) GetRawPackageIndexEnabledOk() (*bool, bool)`

GetRawPackageIndexEnabledOk returns a tuple with the RawPackageIndexEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPackageIndexEnabled

`func (o *RepositoryCreateRequest) SetRawPackageIndexEnabled(v bool)`

SetRawPackageIndexEnabled sets RawPackageIndexEnabled field to given value.

### HasRawPackageIndexEnabled

`func (o *RepositoryCreateRequest) HasRawPackageIndexEnabled() bool`

HasRawPackageIndexEnabled returns a boolean if a field has been set.

### GetRawPackageIndexSignaturesEnabled

`func (o *RepositoryCreateRequest) GetRawPackageIndexSignaturesEnabled() bool`

GetRawPackageIndexSignaturesEnabled returns the RawPackageIndexSignaturesEnabled field if non-nil, zero value otherwise.

### GetRawPackageIndexSignaturesEnabledOk

`func (o *RepositoryCreateRequest) GetRawPackageIndexSignaturesEnabledOk() (*bool, bool)`

GetRawPackageIndexSignaturesEnabledOk returns a tuple with the RawPackageIndexSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPackageIndexSignaturesEnabled

`func (o *RepositoryCreateRequest) SetRawPackageIndexSignaturesEnabled(v bool)`

SetRawPackageIndexSignaturesEnabled sets RawPackageIndexSignaturesEnabled field to given value.

### HasRawPackageIndexSignaturesEnabled

`func (o *RepositoryCreateRequest) HasRawPackageIndexSignaturesEnabled() bool`

HasRawPackageIndexSignaturesEnabled returns a boolean if a field has been set.

### GetReplacePackages

`func (o *RepositoryCreateRequest) GetReplacePackages() string`

GetReplacePackages returns the ReplacePackages field if non-nil, zero value otherwise.

### GetReplacePackagesOk

`func (o *RepositoryCreateRequest) GetReplacePackagesOk() (*string, bool)`

GetReplacePackagesOk returns a tuple with the ReplacePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePackages

`func (o *RepositoryCreateRequest) SetReplacePackages(v string)`

SetReplacePackages sets ReplacePackages field to given value.

### HasReplacePackages

`func (o *RepositoryCreateRequest) HasReplacePackages() bool`

HasReplacePackages returns a boolean if a field has been set.

### GetReplacePackagesByDefault

`func (o *RepositoryCreateRequest) GetReplacePackagesByDefault() bool`

GetReplacePackagesByDefault returns the ReplacePackagesByDefault field if non-nil, zero value otherwise.

### GetReplacePackagesByDefaultOk

`func (o *RepositoryCreateRequest) GetReplacePackagesByDefaultOk() (*bool, bool)`

GetReplacePackagesByDefaultOk returns a tuple with the ReplacePackagesByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePackagesByDefault

`func (o *RepositoryCreateRequest) SetReplacePackagesByDefault(v bool)`

SetReplacePackagesByDefault sets ReplacePackagesByDefault field to given value.

### HasReplacePackagesByDefault

`func (o *RepositoryCreateRequest) HasReplacePackagesByDefault() bool`

HasReplacePackagesByDefault returns a boolean if a field has been set.

### GetRepositoryTypeStr

`func (o *RepositoryCreateRequest) GetRepositoryTypeStr() string`

GetRepositoryTypeStr returns the RepositoryTypeStr field if non-nil, zero value otherwise.

### GetRepositoryTypeStrOk

`func (o *RepositoryCreateRequest) GetRepositoryTypeStrOk() (*string, bool)`

GetRepositoryTypeStrOk returns a tuple with the RepositoryTypeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryTypeStr

`func (o *RepositoryCreateRequest) SetRepositoryTypeStr(v string)`

SetRepositoryTypeStr sets RepositoryTypeStr field to given value.

### HasRepositoryTypeStr

`func (o *RepositoryCreateRequest) HasRepositoryTypeStr() bool`

HasRepositoryTypeStr returns a boolean if a field has been set.

### GetResyncOwn

`func (o *RepositoryCreateRequest) GetResyncOwn() bool`

GetResyncOwn returns the ResyncOwn field if non-nil, zero value otherwise.

### GetResyncOwnOk

`func (o *RepositoryCreateRequest) GetResyncOwnOk() (*bool, bool)`

GetResyncOwnOk returns a tuple with the ResyncOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncOwn

`func (o *RepositoryCreateRequest) SetResyncOwn(v bool)`

SetResyncOwn sets ResyncOwn field to given value.

### HasResyncOwn

`func (o *RepositoryCreateRequest) HasResyncOwn() bool`

HasResyncOwn returns a boolean if a field has been set.

### GetResyncPackages

`func (o *RepositoryCreateRequest) GetResyncPackages() string`

GetResyncPackages returns the ResyncPackages field if non-nil, zero value otherwise.

### GetResyncPackagesOk

`func (o *RepositoryCreateRequest) GetResyncPackagesOk() (*string, bool)`

GetResyncPackagesOk returns a tuple with the ResyncPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncPackages

`func (o *RepositoryCreateRequest) SetResyncPackages(v string)`

SetResyncPackages sets ResyncPackages field to given value.

### HasResyncPackages

`func (o *RepositoryCreateRequest) HasResyncPackages() bool`

HasResyncPackages returns a boolean if a field has been set.

### GetScanOwn

`func (o *RepositoryCreateRequest) GetScanOwn() bool`

GetScanOwn returns the ScanOwn field if non-nil, zero value otherwise.

### GetScanOwnOk

`func (o *RepositoryCreateRequest) GetScanOwnOk() (*bool, bool)`

GetScanOwnOk returns a tuple with the ScanOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOwn

`func (o *RepositoryCreateRequest) SetScanOwn(v bool)`

SetScanOwn sets ScanOwn field to given value.

### HasScanOwn

`func (o *RepositoryCreateRequest) HasScanOwn() bool`

HasScanOwn returns a boolean if a field has been set.

### GetScanPackages

`func (o *RepositoryCreateRequest) GetScanPackages() string`

GetScanPackages returns the ScanPackages field if non-nil, zero value otherwise.

### GetScanPackagesOk

`func (o *RepositoryCreateRequest) GetScanPackagesOk() (*string, bool)`

GetScanPackagesOk returns a tuple with the ScanPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPackages

`func (o *RepositoryCreateRequest) SetScanPackages(v string)`

SetScanPackages sets ScanPackages field to given value.

### HasScanPackages

`func (o *RepositoryCreateRequest) HasScanPackages() bool`

HasScanPackages returns a boolean if a field has been set.

### GetShowSetupAll

`func (o *RepositoryCreateRequest) GetShowSetupAll() bool`

GetShowSetupAll returns the ShowSetupAll field if non-nil, zero value otherwise.

### GetShowSetupAllOk

`func (o *RepositoryCreateRequest) GetShowSetupAllOk() (*bool, bool)`

GetShowSetupAllOk returns a tuple with the ShowSetupAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSetupAll

`func (o *RepositoryCreateRequest) SetShowSetupAll(v bool)`

SetShowSetupAll sets ShowSetupAll field to given value.

### HasShowSetupAll

`func (o *RepositoryCreateRequest) HasShowSetupAll() bool`

HasShowSetupAll returns a boolean if a field has been set.

### GetSlug

`func (o *RepositoryCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RepositoryCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RepositoryCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *RepositoryCreateRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStorageRegion

`func (o *RepositoryCreateRequest) GetStorageRegion() string`

GetStorageRegion returns the StorageRegion field if non-nil, zero value otherwise.

### GetStorageRegionOk

`func (o *RepositoryCreateRequest) GetStorageRegionOk() (*string, bool)`

GetStorageRegionOk returns a tuple with the StorageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRegion

`func (o *RepositoryCreateRequest) SetStorageRegion(v string)`

SetStorageRegion sets StorageRegion field to given value.

### HasStorageRegion

`func (o *RepositoryCreateRequest) HasStorageRegion() bool`

HasStorageRegion returns a boolean if a field has been set.

### GetStrictNpmValidation

`func (o *RepositoryCreateRequest) GetStrictNpmValidation() bool`

GetStrictNpmValidation returns the StrictNpmValidation field if non-nil, zero value otherwise.

### GetStrictNpmValidationOk

`func (o *RepositoryCreateRequest) GetStrictNpmValidationOk() (*bool, bool)`

GetStrictNpmValidationOk returns a tuple with the StrictNpmValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictNpmValidation

`func (o *RepositoryCreateRequest) SetStrictNpmValidation(v bool)`

SetStrictNpmValidation sets StrictNpmValidation field to given value.

### HasStrictNpmValidation

`func (o *RepositoryCreateRequest) HasStrictNpmValidation() bool`

HasStrictNpmValidation returns a boolean if a field has been set.

### GetTagPreReleasesAsLatest

`func (o *RepositoryCreateRequest) GetTagPreReleasesAsLatest() bool`

GetTagPreReleasesAsLatest returns the TagPreReleasesAsLatest field if non-nil, zero value otherwise.

### GetTagPreReleasesAsLatestOk

`func (o *RepositoryCreateRequest) GetTagPreReleasesAsLatestOk() (*bool, bool)`

GetTagPreReleasesAsLatestOk returns a tuple with the TagPreReleasesAsLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPreReleasesAsLatest

`func (o *RepositoryCreateRequest) SetTagPreReleasesAsLatest(v bool)`

SetTagPreReleasesAsLatest sets TagPreReleasesAsLatest field to given value.

### HasTagPreReleasesAsLatest

`func (o *RepositoryCreateRequest) HasTagPreReleasesAsLatest() bool`

HasTagPreReleasesAsLatest returns a boolean if a field has been set.

### GetUseDebianLabels

`func (o *RepositoryCreateRequest) GetUseDebianLabels() bool`

GetUseDebianLabels returns the UseDebianLabels field if non-nil, zero value otherwise.

### GetUseDebianLabelsOk

`func (o *RepositoryCreateRequest) GetUseDebianLabelsOk() (*bool, bool)`

GetUseDebianLabelsOk returns a tuple with the UseDebianLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDebianLabels

`func (o *RepositoryCreateRequest) SetUseDebianLabels(v bool)`

SetUseDebianLabels sets UseDebianLabels field to given value.

### HasUseDebianLabels

`func (o *RepositoryCreateRequest) HasUseDebianLabels() bool`

HasUseDebianLabels returns a boolean if a field has been set.

### GetUseDefaultCargoUpstream

`func (o *RepositoryCreateRequest) GetUseDefaultCargoUpstream() bool`

GetUseDefaultCargoUpstream returns the UseDefaultCargoUpstream field if non-nil, zero value otherwise.

### GetUseDefaultCargoUpstreamOk

`func (o *RepositoryCreateRequest) GetUseDefaultCargoUpstreamOk() (*bool, bool)`

GetUseDefaultCargoUpstreamOk returns a tuple with the UseDefaultCargoUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultCargoUpstream

`func (o *RepositoryCreateRequest) SetUseDefaultCargoUpstream(v bool)`

SetUseDefaultCargoUpstream sets UseDefaultCargoUpstream field to given value.

### HasUseDefaultCargoUpstream

`func (o *RepositoryCreateRequest) HasUseDefaultCargoUpstream() bool`

HasUseDefaultCargoUpstream returns a boolean if a field has been set.

### GetUseEntitlementsPrivilege

`func (o *RepositoryCreateRequest) GetUseEntitlementsPrivilege() string`

GetUseEntitlementsPrivilege returns the UseEntitlementsPrivilege field if non-nil, zero value otherwise.

### GetUseEntitlementsPrivilegeOk

`func (o *RepositoryCreateRequest) GetUseEntitlementsPrivilegeOk() (*string, bool)`

GetUseEntitlementsPrivilegeOk returns a tuple with the UseEntitlementsPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEntitlementsPrivilege

`func (o *RepositoryCreateRequest) SetUseEntitlementsPrivilege(v string)`

SetUseEntitlementsPrivilege sets UseEntitlementsPrivilege field to given value.

### HasUseEntitlementsPrivilege

`func (o *RepositoryCreateRequest) HasUseEntitlementsPrivilege() bool`

HasUseEntitlementsPrivilege returns a boolean if a field has been set.

### GetUseNoarchPackages

`func (o *RepositoryCreateRequest) GetUseNoarchPackages() bool`

GetUseNoarchPackages returns the UseNoarchPackages field if non-nil, zero value otherwise.

### GetUseNoarchPackagesOk

`func (o *RepositoryCreateRequest) GetUseNoarchPackagesOk() (*bool, bool)`

GetUseNoarchPackagesOk returns a tuple with the UseNoarchPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNoarchPackages

`func (o *RepositoryCreateRequest) SetUseNoarchPackages(v bool)`

SetUseNoarchPackages sets UseNoarchPackages field to given value.

### HasUseNoarchPackages

`func (o *RepositoryCreateRequest) HasUseNoarchPackages() bool`

HasUseNoarchPackages returns a boolean if a field has been set.

### GetUseSourcePackages

`func (o *RepositoryCreateRequest) GetUseSourcePackages() bool`

GetUseSourcePackages returns the UseSourcePackages field if non-nil, zero value otherwise.

### GetUseSourcePackagesOk

`func (o *RepositoryCreateRequest) GetUseSourcePackagesOk() (*bool, bool)`

GetUseSourcePackagesOk returns a tuple with the UseSourcePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSourcePackages

`func (o *RepositoryCreateRequest) SetUseSourcePackages(v bool)`

SetUseSourcePackages sets UseSourcePackages field to given value.

### HasUseSourcePackages

`func (o *RepositoryCreateRequest) HasUseSourcePackages() bool`

HasUseSourcePackages returns a boolean if a field has been set.

### GetUseVulnerabilityScanning

`func (o *RepositoryCreateRequest) GetUseVulnerabilityScanning() bool`

GetUseVulnerabilityScanning returns the UseVulnerabilityScanning field if non-nil, zero value otherwise.

### GetUseVulnerabilityScanningOk

`func (o *RepositoryCreateRequest) GetUseVulnerabilityScanningOk() (*bool, bool)`

GetUseVulnerabilityScanningOk returns a tuple with the UseVulnerabilityScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVulnerabilityScanning

`func (o *RepositoryCreateRequest) SetUseVulnerabilityScanning(v bool)`

SetUseVulnerabilityScanning sets UseVulnerabilityScanning field to given value.

### HasUseVulnerabilityScanning

`func (o *RepositoryCreateRequest) HasUseVulnerabilityScanning() bool`

HasUseVulnerabilityScanning returns a boolean if a field has been set.

### GetUserEntitlementsEnabled

`func (o *RepositoryCreateRequest) GetUserEntitlementsEnabled() bool`

GetUserEntitlementsEnabled returns the UserEntitlementsEnabled field if non-nil, zero value otherwise.

### GetUserEntitlementsEnabledOk

`func (o *RepositoryCreateRequest) GetUserEntitlementsEnabledOk() (*bool, bool)`

GetUserEntitlementsEnabledOk returns a tuple with the UserEntitlementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEntitlementsEnabled

`func (o *RepositoryCreateRequest) SetUserEntitlementsEnabled(v bool)`

SetUserEntitlementsEnabled sets UserEntitlementsEnabled field to given value.

### HasUserEntitlementsEnabled

`func (o *RepositoryCreateRequest) HasUserEntitlementsEnabled() bool`

HasUserEntitlementsEnabled returns a boolean if a field has been set.

### GetViewStatistics

`func (o *RepositoryCreateRequest) GetViewStatistics() string`

GetViewStatistics returns the ViewStatistics field if non-nil, zero value otherwise.

### GetViewStatisticsOk

`func (o *RepositoryCreateRequest) GetViewStatisticsOk() (*string, bool)`

GetViewStatisticsOk returns a tuple with the ViewStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStatistics

`func (o *RepositoryCreateRequest) SetViewStatistics(v string)`

SetViewStatistics sets ViewStatistics field to given value.

### HasViewStatistics

`func (o *RepositoryCreateRequest) HasViewStatistics() bool`

HasViewStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


