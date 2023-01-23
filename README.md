# Go API client for cloudsmith

The API to the Cloudsmith Service

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.202.1
- Package version: 0.0.20
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://help.cloudsmith.io](https://help.cloudsmith.io)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import cloudsmith "github.com/cloudsmith-io/cloudsmith-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), cloudsmith.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), cloudsmith.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), cloudsmith.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), cloudsmith.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.cloudsmith.io/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditLogApi* | [**AuditLogNamespaceList**](docs/AuditLogApi.md#auditlognamespacelist) | **Get** /audit-log/{owner}/ | Lists audit log entries for a specific namespace.
*AuditLogApi* | [**AuditLogRepoList**](docs/AuditLogApi.md#auditlogrepolist) | **Get** /audit-log/{owner}/{repo}/ | Lists audit log entries for a specific repository.
*BadgesApi* | [**BadgesVersionList**](docs/BadgesApi.md#badgesversionlist) | **Get** /badges/version/{owner}/{repo}/{package_format}/{package_name}/{package_version}/{package_identifiers}/ | Get latest package version for a package or package group.
*DistrosApi* | [**DistrosList**](docs/DistrosApi.md#distroslist) | **Get** /distros/ | Get a list of all supported distributions.
*DistrosApi* | [**DistrosRead**](docs/DistrosApi.md#distrosread) | **Get** /distros/{slug}/ | View for viewing/listing distributions.
*EntitlementsApi* | [**EntitlementsCreate**](docs/EntitlementsApi.md#entitlementscreate) | **Post** /entitlements/{owner}/{repo}/ | Create a specific entitlement in a repository.
*EntitlementsApi* | [**EntitlementsDelete**](docs/EntitlementsApi.md#entitlementsdelete) | **Delete** /entitlements/{owner}/{repo}/{identifier}/ | Delete a specific entitlement in a repository.
*EntitlementsApi* | [**EntitlementsDisable**](docs/EntitlementsApi.md#entitlementsdisable) | **Post** /entitlements/{owner}/{repo}/{identifier}/disable/ | Disable an entitlement token in a repository.
*EntitlementsApi* | [**EntitlementsEnable**](docs/EntitlementsApi.md#entitlementsenable) | **Post** /entitlements/{owner}/{repo}/{identifier}/enable/ | Enable an entitlement token in a repository.
*EntitlementsApi* | [**EntitlementsList**](docs/EntitlementsApi.md#entitlementslist) | **Get** /entitlements/{owner}/{repo}/ | Get a list of all entitlements in a repository.
*EntitlementsApi* | [**EntitlementsPartialUpdate**](docs/EntitlementsApi.md#entitlementspartialupdate) | **Patch** /entitlements/{owner}/{repo}/{identifier}/ | Update a specific entitlement in a repository.
*EntitlementsApi* | [**EntitlementsRead**](docs/EntitlementsApi.md#entitlementsread) | **Get** /entitlements/{owner}/{repo}/{identifier}/ | Get a specific entitlement in a repository.
*EntitlementsApi* | [**EntitlementsRefresh**](docs/EntitlementsApi.md#entitlementsrefresh) | **Post** /entitlements/{owner}/{repo}/{identifier}/refresh/ | Refresh an entitlement token in a repository.
*EntitlementsApi* | [**EntitlementsReset**](docs/EntitlementsApi.md#entitlementsreset) | **Post** /entitlements/{owner}/{repo}/{identifier}/reset/ | Reset the statistics for an entitlement token in a repository.
*EntitlementsApi* | [**EntitlementsSync**](docs/EntitlementsApi.md#entitlementssync) | **Post** /entitlements/{owner}/{repo}/sync/ | Synchronise tokens from a source repository.
*FilesApi* | [**FilesAbort**](docs/FilesApi.md#filesabort) | **Post** /files/{owner}/{repo}/{identifier}/abort/ | Abort a multipart file upload.
*FilesApi* | [**FilesComplete**](docs/FilesApi.md#filescomplete) | **Post** /files/{owner}/{repo}/{identifier}/complete/ | Complete a multipart file upload.
*FilesApi* | [**FilesCreate**](docs/FilesApi.md#filescreate) | **Post** /files/{owner}/{repo}/ | Request URL(s) to upload new package file upload(s) to.
*FilesApi* | [**FilesInfo**](docs/FilesApi.md#filesinfo) | **Get** /files/{owner}/{repo}/{identifier}/info/ | Get upload information to perform a multipart file upload.
*FilesApi* | [**FilesValidate**](docs/FilesApi.md#filesvalidate) | **Post** /files/{owner}/{repo}/validate/ | Validate parameters used for create.
*FormatsApi* | [**FormatsList**](docs/FormatsApi.md#formatslist) | **Get** /formats/ | Get a list of all supported package formats.
*FormatsApi* | [**FormatsRead**](docs/FormatsApi.md#formatsread) | **Get** /formats/{slug}/ | Get a specific supported package format.
*MetricsApi* | [**MetricsEntitlementsAccountList**](docs/MetricsApi.md#metricsentitlementsaccountlist) | **Get** /metrics/entitlements/{owner}/ | View for listing entitlement token metrics, across an account.
*MetricsApi* | [**MetricsEntitlementsRepoList**](docs/MetricsApi.md#metricsentitlementsrepolist) | **Get** /metrics/entitlements/{owner}/{repo}/ | View for listing entitlement token metrics, for a repository.
*MetricsApi* | [**MetricsPackagesList**](docs/MetricsApi.md#metricspackageslist) | **Get** /metrics/packages/{owner}/{repo}/ | View for listing package usage metrics, for a repository.
*NamespacesApi* | [**NamespacesList**](docs/NamespacesApi.md#namespaceslist) | **Get** /namespaces/ | Get a list of all namespaces the user belongs to.
*NamespacesApi* | [**NamespacesRead**](docs/NamespacesApi.md#namespacesread) | **Get** /namespaces/{slug}/ | Views for working with namespaces.
*OrgsApi* | [**OrgsInvitesCreate**](docs/OrgsApi.md#orgsinvitescreate) | **Post** /orgs/{org}/invites/ | Create an organization invite for a specific user
*OrgsApi* | [**OrgsInvitesDelete**](docs/OrgsApi.md#orgsinvitesdelete) | **Delete** /orgs/{org}/invites/{slug_perm}/ | Delete a specific organization invite
*OrgsApi* | [**OrgsInvitesExtend**](docs/OrgsApi.md#orgsinvitesextend) | **Post** /orgs/{org}/invites/{slug_perm}/extend/ | Extend an organization invite.
*OrgsApi* | [**OrgsInvitesList**](docs/OrgsApi.md#orgsinviteslist) | **Get** /orgs/{org}/invites/ | Get a list of all invites for an organization.
*OrgsApi* | [**OrgsInvitesPartialUpdate**](docs/OrgsApi.md#orgsinvitespartialupdate) | **Patch** /orgs/{org}/invites/{slug_perm}/ | Update a specific organization invite.
*OrgsApi* | [**OrgsInvitesResend**](docs/OrgsApi.md#orgsinvitesresend) | **Post** /orgs/{org}/invites/{slug_perm}/resend/ | Resend an organization invite.
*OrgsApi* | [**OrgsList**](docs/OrgsApi.md#orgslist) | **Get** /orgs/ | Get a list of all the organizations you are associated with.
*OrgsApi* | [**OrgsMembersDelete**](docs/OrgsApi.md#orgsmembersdelete) | **Delete** /orgs/{org}/members/{member}/ | Removes a member from the organization.
*OrgsApi* | [**OrgsMembersList**](docs/OrgsApi.md#orgsmemberslist) | **Get** /orgs/{org}/members/ | Get the details for all organization members.
*OrgsApi* | [**OrgsMembersRead**](docs/OrgsApi.md#orgsmembersread) | **Get** /orgs/{org}/members/{member}/ | Get the details for a specific organization member.
*OrgsApi* | [**OrgsMembersRemove**](docs/OrgsApi.md#orgsmembersremove) | **Get** /orgs/{org}/members/{member}/remove/ | Removes a member from the organization (deprecated, use DELETE instead).
*OrgsApi* | [**OrgsRead**](docs/OrgsApi.md#orgsread) | **Get** /orgs/{org}/ | Get the details for the specific organization.
*OrgsApi* | [**OrgsSamlGroupSyncCreate**](docs/OrgsApi.md#orgssamlgroupsynccreate) | **Post** /orgs/{org}/saml-group-sync/ | Create a new SAML Group Sync mapping within an organization.
*OrgsApi* | [**OrgsSamlGroupSyncDelete**](docs/OrgsApi.md#orgssamlgroupsyncdelete) | **Delete** /orgs/{org}/saml-group-sync/{slug_perm}/ | Delete a SAML Group Sync mapping from an organization.
*OrgsApi* | [**OrgsSamlGroupSyncList**](docs/OrgsApi.md#orgssamlgroupsynclist) | **Get** /orgs/{org}/saml-group-sync/ | Get the details of all SAML Group Sync mapping within an organization.
*OrgsApi* | [**OrgsServicesCreate**](docs/OrgsApi.md#orgsservicescreate) | **Post** /orgs/{org}/services/ | Create a service within an organization.
*OrgsApi* | [**OrgsServicesDelete**](docs/OrgsApi.md#orgsservicesdelete) | **Delete** /orgs/{org}/services/{service}/ | Delete a specific service
*OrgsApi* | [**OrgsServicesList**](docs/OrgsApi.md#orgsserviceslist) | **Get** /orgs/{org}/services/ | Get a list of all services within an organization.
*OrgsApi* | [**OrgsServicesPartialUpdate**](docs/OrgsApi.md#orgsservicespartialupdate) | **Patch** /orgs/{org}/services/{service}/ | Update a service within an organization.
*OrgsApi* | [**OrgsServicesRead**](docs/OrgsApi.md#orgsservicesread) | **Get** /orgs/{org}/services/{service}/ | Retrieve details of a single service within an organization.
*OrgsApi* | [**OrgsServicesRefresh**](docs/OrgsApi.md#orgsservicesrefresh) | **Post** /orgs/{org}/services/{service}/refresh/ | Refresh service API token.
*OrgsApi* | [**OrgsTeamsCreate**](docs/OrgsApi.md#orgsteamscreate) | **Post** /orgs/{org}/teams/ | Create a team for this organization.
*OrgsApi* | [**OrgsTeamsDelete**](docs/OrgsApi.md#orgsteamsdelete) | **Delete** /orgs/{org}/teams/{team}/ | Delete a specific team in a organization.
*OrgsApi* | [**OrgsTeamsList**](docs/OrgsApi.md#orgsteamslist) | **Get** /orgs/{org}/teams/ | Get the details of all teams within an organization.
*OrgsApi* | [**OrgsTeamsMembersCreate**](docs/OrgsApi.md#orgsteamsmemberscreate) | **Post** /orgs/{org}/teams/{team}/members | Add users to a team.
*OrgsApi* | [**OrgsTeamsMembersList**](docs/OrgsApi.md#orgsteamsmemberslist) | **Get** /orgs/{org}/teams/{team}/members | List all members for the team.
*OrgsApi* | [**OrgsTeamsMembersUpdate**](docs/OrgsApi.md#orgsteamsmembersupdate) | **Put** /orgs/{org}/teams/{team}/members | Replace all team members.
*OrgsApi* | [**OrgsTeamsPartialUpdate**](docs/OrgsApi.md#orgsteamspartialupdate) | **Patch** /orgs/{org}/teams/{team}/ | Update a specific team in a organization.
*OrgsApi* | [**OrgsTeamsRead**](docs/OrgsApi.md#orgsteamsread) | **Get** /orgs/{org}/teams/{team}/ | Get the details of a specific team within an organization.
*PackagesApi* | [**PackagesCopy**](docs/PackagesApi.md#packagescopy) | **Post** /packages/{owner}/{repo}/{identifier}/copy/ | Copy a package to another repository.
*PackagesApi* | [**PackagesDelete**](docs/PackagesApi.md#packagesdelete) | **Delete** /packages/{owner}/{repo}/{identifier}/ | Delete a specific package in a repository.
*PackagesApi* | [**PackagesDependencies**](docs/PackagesApi.md#packagesdependencies) | **Get** /packages/{owner}/{repo}/{identifier}/dependencies/ | Get the direct (non-transitive) dependencies list for a package.
*PackagesApi* | [**PackagesList**](docs/PackagesApi.md#packageslist) | **Get** /packages/{owner}/{repo}/ | Get a list of all packages associated with repository.
*PackagesApi* | [**PackagesMove**](docs/PackagesApi.md#packagesmove) | **Post** /packages/{owner}/{repo}/{identifier}/move/ | Move a package to another repository.
*PackagesApi* | [**PackagesQuarantine**](docs/PackagesApi.md#packagesquarantine) | **Post** /packages/{owner}/{repo}/{identifier}/quarantine/ | Quarantine or restore a package.
*PackagesApi* | [**PackagesRead**](docs/PackagesApi.md#packagesread) | **Get** /packages/{owner}/{repo}/{identifier}/ | Get a specific package in a repository.
*PackagesApi* | [**PackagesResync**](docs/PackagesApi.md#packagesresync) | **Post** /packages/{owner}/{repo}/{identifier}/resync/ | Schedule a package for resynchronisation.
*PackagesApi* | [**PackagesScan**](docs/PackagesApi.md#packagesscan) | **Post** /packages/{owner}/{repo}/{identifier}/scan/ | Schedule a package for scanning.
*PackagesApi* | [**PackagesStatus**](docs/PackagesApi.md#packagesstatus) | **Get** /packages/{owner}/{repo}/{identifier}/status/ | Get the synchronisation status for a package.
*PackagesApi* | [**PackagesTag**](docs/PackagesApi.md#packagestag) | **Post** /packages/{owner}/{repo}/{identifier}/tag/ | Add/Replace/Remove tags for a package.
*PackagesApi* | [**PackagesUploadAlpine**](docs/PackagesApi.md#packagesuploadalpine) | **Post** /packages/{owner}/{repo}/upload/alpine/ | Create a new Alpine package
*PackagesApi* | [**PackagesUploadCargo**](docs/PackagesApi.md#packagesuploadcargo) | **Post** /packages/{owner}/{repo}/upload/cargo/ | Create a new Cargo package
*PackagesApi* | [**PackagesUploadCocoapods**](docs/PackagesApi.md#packagesuploadcocoapods) | **Post** /packages/{owner}/{repo}/upload/cocoapods/ | Create a new CocoaPods package
*PackagesApi* | [**PackagesUploadComposer**](docs/PackagesApi.md#packagesuploadcomposer) | **Post** /packages/{owner}/{repo}/upload/composer/ | Create a new Composer package
*PackagesApi* | [**PackagesUploadConan**](docs/PackagesApi.md#packagesuploadconan) | **Post** /packages/{owner}/{repo}/upload/conan/ | Create a new Conan package
*PackagesApi* | [**PackagesUploadConda**](docs/PackagesApi.md#packagesuploadconda) | **Post** /packages/{owner}/{repo}/upload/conda/ | Create a new Conda package
*PackagesApi* | [**PackagesUploadCran**](docs/PackagesApi.md#packagesuploadcran) | **Post** /packages/{owner}/{repo}/upload/cran/ | Create a new CRAN package
*PackagesApi* | [**PackagesUploadDart**](docs/PackagesApi.md#packagesuploaddart) | **Post** /packages/{owner}/{repo}/upload/dart/ | Create a new Dart package
*PackagesApi* | [**PackagesUploadDeb**](docs/PackagesApi.md#packagesuploaddeb) | **Post** /packages/{owner}/{repo}/upload/deb/ | Create a new Debian package
*PackagesApi* | [**PackagesUploadDocker**](docs/PackagesApi.md#packagesuploaddocker) | **Post** /packages/{owner}/{repo}/upload/docker/ | Create a new Docker package
*PackagesApi* | [**PackagesUploadGo**](docs/PackagesApi.md#packagesuploadgo) | **Post** /packages/{owner}/{repo}/upload/go/ | Create a new Go package
*PackagesApi* | [**PackagesUploadHelm**](docs/PackagesApi.md#packagesuploadhelm) | **Post** /packages/{owner}/{repo}/upload/helm/ | Create a new Helm package
*PackagesApi* | [**PackagesUploadLuarocks**](docs/PackagesApi.md#packagesuploadluarocks) | **Post** /packages/{owner}/{repo}/upload/luarocks/ | Create a new LuaRocks package
*PackagesApi* | [**PackagesUploadMaven**](docs/PackagesApi.md#packagesuploadmaven) | **Post** /packages/{owner}/{repo}/upload/maven/ | Create a new Maven package
*PackagesApi* | [**PackagesUploadNpm**](docs/PackagesApi.md#packagesuploadnpm) | **Post** /packages/{owner}/{repo}/upload/npm/ | Create a new npm package
*PackagesApi* | [**PackagesUploadNuget**](docs/PackagesApi.md#packagesuploadnuget) | **Post** /packages/{owner}/{repo}/upload/nuget/ | Create a new NuGet package
*PackagesApi* | [**PackagesUploadP2**](docs/PackagesApi.md#packagesuploadp2) | **Post** /packages/{owner}/{repo}/upload/p2/ | Create a new P2 package
*PackagesApi* | [**PackagesUploadPython**](docs/PackagesApi.md#packagesuploadpython) | **Post** /packages/{owner}/{repo}/upload/python/ | Create a new Python package
*PackagesApi* | [**PackagesUploadRaw**](docs/PackagesApi.md#packagesuploadraw) | **Post** /packages/{owner}/{repo}/upload/raw/ | Create a new Raw package
*PackagesApi* | [**PackagesUploadRpm**](docs/PackagesApi.md#packagesuploadrpm) | **Post** /packages/{owner}/{repo}/upload/rpm/ | Create a new RedHat package
*PackagesApi* | [**PackagesUploadRuby**](docs/PackagesApi.md#packagesuploadruby) | **Post** /packages/{owner}/{repo}/upload/ruby/ | Create a new Ruby package
*PackagesApi* | [**PackagesUploadTerraform**](docs/PackagesApi.md#packagesuploadterraform) | **Post** /packages/{owner}/{repo}/upload/terraform/ | Create a new Terraform package
*PackagesApi* | [**PackagesUploadVagrant**](docs/PackagesApi.md#packagesuploadvagrant) | **Post** /packages/{owner}/{repo}/upload/vagrant/ | Create a new Vagrant package
*PackagesApi* | [**PackagesValidateUploadAlpine**](docs/PackagesApi.md#packagesvalidateuploadalpine) | **Post** /packages/{owner}/{repo}/validate-upload/alpine/ | Validate parameters for create Alpine package
*PackagesApi* | [**PackagesValidateUploadCargo**](docs/PackagesApi.md#packagesvalidateuploadcargo) | **Post** /packages/{owner}/{repo}/validate-upload/cargo/ | Validate parameters for create Cargo package
*PackagesApi* | [**PackagesValidateUploadCocoapods**](docs/PackagesApi.md#packagesvalidateuploadcocoapods) | **Post** /packages/{owner}/{repo}/validate-upload/cocoapods/ | Validate parameters for create CocoaPods package
*PackagesApi* | [**PackagesValidateUploadComposer**](docs/PackagesApi.md#packagesvalidateuploadcomposer) | **Post** /packages/{owner}/{repo}/validate-upload/composer/ | Validate parameters for create Composer package
*PackagesApi* | [**PackagesValidateUploadConan**](docs/PackagesApi.md#packagesvalidateuploadconan) | **Post** /packages/{owner}/{repo}/validate-upload/conan/ | Validate parameters for create Conan package
*PackagesApi* | [**PackagesValidateUploadConda**](docs/PackagesApi.md#packagesvalidateuploadconda) | **Post** /packages/{owner}/{repo}/validate-upload/conda/ | Validate parameters for create Conda package
*PackagesApi* | [**PackagesValidateUploadCran**](docs/PackagesApi.md#packagesvalidateuploadcran) | **Post** /packages/{owner}/{repo}/validate-upload/cran/ | Validate parameters for create CRAN package
*PackagesApi* | [**PackagesValidateUploadDart**](docs/PackagesApi.md#packagesvalidateuploaddart) | **Post** /packages/{owner}/{repo}/validate-upload/dart/ | Validate parameters for create Dart package
*PackagesApi* | [**PackagesValidateUploadDeb**](docs/PackagesApi.md#packagesvalidateuploaddeb) | **Post** /packages/{owner}/{repo}/validate-upload/deb/ | Validate parameters for create Debian package
*PackagesApi* | [**PackagesValidateUploadDocker**](docs/PackagesApi.md#packagesvalidateuploaddocker) | **Post** /packages/{owner}/{repo}/validate-upload/docker/ | Validate parameters for create Docker package
*PackagesApi* | [**PackagesValidateUploadGo**](docs/PackagesApi.md#packagesvalidateuploadgo) | **Post** /packages/{owner}/{repo}/validate-upload/go/ | Validate parameters for create Go package
*PackagesApi* | [**PackagesValidateUploadHelm**](docs/PackagesApi.md#packagesvalidateuploadhelm) | **Post** /packages/{owner}/{repo}/validate-upload/helm/ | Validate parameters for create Helm package
*PackagesApi* | [**PackagesValidateUploadLuarocks**](docs/PackagesApi.md#packagesvalidateuploadluarocks) | **Post** /packages/{owner}/{repo}/validate-upload/luarocks/ | Validate parameters for create LuaRocks package
*PackagesApi* | [**PackagesValidateUploadMaven**](docs/PackagesApi.md#packagesvalidateuploadmaven) | **Post** /packages/{owner}/{repo}/validate-upload/maven/ | Validate parameters for create Maven package
*PackagesApi* | [**PackagesValidateUploadNpm**](docs/PackagesApi.md#packagesvalidateuploadnpm) | **Post** /packages/{owner}/{repo}/validate-upload/npm/ | Validate parameters for create npm package
*PackagesApi* | [**PackagesValidateUploadNuget**](docs/PackagesApi.md#packagesvalidateuploadnuget) | **Post** /packages/{owner}/{repo}/validate-upload/nuget/ | Validate parameters for create NuGet package
*PackagesApi* | [**PackagesValidateUploadP2**](docs/PackagesApi.md#packagesvalidateuploadp2) | **Post** /packages/{owner}/{repo}/validate-upload/p2/ | Validate parameters for create P2 package
*PackagesApi* | [**PackagesValidateUploadPython**](docs/PackagesApi.md#packagesvalidateuploadpython) | **Post** /packages/{owner}/{repo}/validate-upload/python/ | Validate parameters for create Python package
*PackagesApi* | [**PackagesValidateUploadRaw**](docs/PackagesApi.md#packagesvalidateuploadraw) | **Post** /packages/{owner}/{repo}/validate-upload/raw/ | Validate parameters for create Raw package
*PackagesApi* | [**PackagesValidateUploadRpm**](docs/PackagesApi.md#packagesvalidateuploadrpm) | **Post** /packages/{owner}/{repo}/validate-upload/rpm/ | Validate parameters for create RedHat package
*PackagesApi* | [**PackagesValidateUploadRuby**](docs/PackagesApi.md#packagesvalidateuploadruby) | **Post** /packages/{owner}/{repo}/validate-upload/ruby/ | Validate parameters for create Ruby package
*PackagesApi* | [**PackagesValidateUploadTerraform**](docs/PackagesApi.md#packagesvalidateuploadterraform) | **Post** /packages/{owner}/{repo}/validate-upload/terraform/ | Validate parameters for create Terraform package
*PackagesApi* | [**PackagesValidateUploadVagrant**](docs/PackagesApi.md#packagesvalidateuploadvagrant) | **Post** /packages/{owner}/{repo}/validate-upload/vagrant/ | Validate parameters for create Vagrant package
*QuotaApi* | [**QuotaHistoryRead**](docs/QuotaApi.md#quotahistoryread) | **Get** /quota/history/{owner}/ | Quota history for a given namespace.
*QuotaApi* | [**QuotaOssHistoryRead**](docs/QuotaApi.md#quotaosshistoryread) | **Get** /quota/oss/history/{owner}/ | Open-source Quota history for a given namespace.
*QuotaApi* | [**QuotaOssRead**](docs/QuotaApi.md#quotaossread) | **Get** /quota/oss/{owner}/ | Open-source Quota usage for a given namespace.
*QuotaApi* | [**QuotaRead**](docs/QuotaApi.md#quotaread) | **Get** /quota/{owner}/ | Quota usage for a given namespace.
*RatesApi* | [**RatesLimitsList**](docs/RatesApi.md#rateslimitslist) | **Get** /rates/limits/ | Endpoint to check rate limits for current user.
*ReposApi* | [**ReposCreate**](docs/ReposApi.md#reposcreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
*ReposApi* | [**ReposDelete**](docs/ReposApi.md#reposdelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
*ReposApi* | [**ReposGeoipDisable**](docs/ReposApi.md#reposgeoipdisable) | **Post** /repos/{owner}/{identifier}/geoip/disable/ | Disable GeoIP for this repository.
*ReposApi* | [**ReposGeoipEnable**](docs/ReposApi.md#reposgeoipenable) | **Post** /repos/{owner}/{identifier}/geoip/enable/ | Enable GeoIP for this repository.
*ReposApi* | [**ReposGeoipPartialUpdate**](docs/ReposApi.md#reposgeoippartialupdate) | **Patch** /repos/{owner}/{identifier}/geoip | Partially update existing repository geoip rules with those specified
*ReposApi* | [**ReposGeoipRead**](docs/ReposApi.md#reposgeoipread) | **Get** /repos/{owner}/{identifier}/geoip | List all created GeoIP rules for the repository.
*ReposApi* | [**ReposGeoipTest**](docs/ReposApi.md#reposgeoiptest) | **Post** /repos/{owner}/{identifier}/geoip/test/ | Test a list of IP addresses against the repository&#39;s current GeoIP rules.
*ReposApi* | [**ReposGeoipUpdate**](docs/ReposApi.md#reposgeoipupdate) | **Put** /repos/{owner}/{identifier}/geoip | Replace all existing repository geoip rules with those specified
*ReposApi* | [**ReposGpgCreate**](docs/ReposApi.md#reposgpgcreate) | **Post** /repos/{owner}/{identifier}/gpg/ | Set the active GPG key for the Repository.
*ReposApi* | [**ReposGpgList**](docs/ReposApi.md#reposgpglist) | **Get** /repos/{owner}/{identifier}/gpg/ | Retrieve the active GPG key for the Repository.
*ReposApi* | [**ReposGpgRegenerate**](docs/ReposApi.md#reposgpgregenerate) | **Post** /repos/{owner}/{identifier}/gpg/regenerate/ | Regenerate GPG Key for the Repository.
*ReposApi* | [**ReposNamespaceList**](docs/ReposApi.md#reposnamespacelist) | **Get** /repos/{owner}/ | Get a list of all repositories within a namespace.
*ReposApi* | [**ReposPartialUpdate**](docs/ReposApi.md#repospartialupdate) | **Patch** /repos/{owner}/{identifier}/ | Update details about a repository in a given namespace.
*ReposApi* | [**ReposPrivilegesList**](docs/ReposApi.md#reposprivilegeslist) | **Get** /repos/{owner}/{identifier}/privileges | List all explicity created privileges for the repository.
*ReposApi* | [**ReposPrivilegesPartialUpdate**](docs/ReposApi.md#reposprivilegespartialupdate) | **Patch** /repos/{owner}/{identifier}/privileges | Modify privileges for the repository.
*ReposApi* | [**ReposPrivilegesUpdate**](docs/ReposApi.md#reposprivilegesupdate) | **Put** /repos/{owner}/{identifier}/privileges | Replace all existing repository privileges with those specified.
*ReposApi* | [**ReposRead**](docs/ReposApi.md#reposread) | **Get** /repos/{owner}/{identifier}/ | Get a specific repository.
*ReposApi* | [**ReposRsaCreate**](docs/ReposApi.md#reposrsacreate) | **Post** /repos/{owner}/{identifier}/rsa/ | Set the active RSA key for the Repository.
*ReposApi* | [**ReposRsaList**](docs/ReposApi.md#reposrsalist) | **Get** /repos/{owner}/{identifier}/rsa/ | Retrieve the active RSA key for the Repository.
*ReposApi* | [**ReposRsaRegenerate**](docs/ReposApi.md#reposrsaregenerate) | **Post** /repos/{owner}/{identifier}/rsa/regenerate/ | Regenerate RSA Key for the Repository.
*ReposApi* | [**ReposUserList**](docs/ReposApi.md#reposuserlist) | **Get** /repos/ | Get a list of all repositories associated with current user.
*StatusApi* | [**StatusCheckBasic**](docs/StatusApi.md#statuscheckbasic) | **Get** /status/check/basic/ | Endpoint to check basic API connectivity.
*StorageRegionsApi* | [**StorageRegionsList**](docs/StorageRegionsApi.md#storageregionslist) | **Get** /storage-regions/ | Get a list of all available storage regions.
*StorageRegionsApi* | [**StorageRegionsRead**](docs/StorageRegionsApi.md#storageregionsread) | **Get** /storage-regions/{slug}/ | Get a specific storage region.
*UserApi* | [**UserSelf**](docs/UserApi.md#userself) | **Get** /user/self/ | Provide a brief for the current user (if any).
*UserApi* | [**UserTokenCreate**](docs/UserApi.md#usertokencreate) | **Post** /user/token/ | Retrieve the API key/token for the authenticated user.
*UsersApi* | [**UsersProfileRead**](docs/UsersApi.md#usersprofileread) | **Get** /users/profile/{slug}/ | Provide a brief for the specified user (if any).
*VulnerabilitiesApi* | [**VulnerabilitiesNamespaceList**](docs/VulnerabilitiesApi.md#vulnerabilitiesnamespacelist) | **Get** /vulnerabilities/{owner}/ | Lists scan results for a specific namespace.
*VulnerabilitiesApi* | [**VulnerabilitiesPackageList**](docs/VulnerabilitiesApi.md#vulnerabilitiespackagelist) | **Get** /vulnerabilities/{owner}/{repo}/{package}/ | Lists scan results for a specific package.
*VulnerabilitiesApi* | [**VulnerabilitiesRead**](docs/VulnerabilitiesApi.md#vulnerabilitiesread) | **Get** /vulnerabilities/{owner}/{repo}/{package}/{scan_id}/ | Returns a Scan Result.
*VulnerabilitiesApi* | [**VulnerabilitiesRepoList**](docs/VulnerabilitiesApi.md#vulnerabilitiesrepolist) | **Get** /vulnerabilities/{owner}/{repo}/ | Lists scan results for a specific repository.
*WebhooksApi* | [**WebhooksCreate**](docs/WebhooksApi.md#webhookscreate) | **Post** /webhooks/{owner}/{repo}/ | Create a specific webhook in a repository.
*WebhooksApi* | [**WebhooksDelete**](docs/WebhooksApi.md#webhooksdelete) | **Delete** /webhooks/{owner}/{repo}/{identifier}/ | Delete a specific webhook in a repository.
*WebhooksApi* | [**WebhooksList**](docs/WebhooksApi.md#webhookslist) | **Get** /webhooks/{owner}/{repo}/ | Get a list of all webhooks in a repository.
*WebhooksApi* | [**WebhooksPartialUpdate**](docs/WebhooksApi.md#webhookspartialupdate) | **Patch** /webhooks/{owner}/{repo}/{identifier}/ | Update a specific webhook in a repository.
*WebhooksApi* | [**WebhooksRead**](docs/WebhooksApi.md#webhooksread) | **Get** /webhooks/{owner}/{repo}/{identifier}/ | Views for working with repository webhooks.


## Documentation For Models

 - [AllocatedLimit](docs/AllocatedLimit.md)
 - [AlpinePackageUpload](docs/AlpinePackageUpload.md)
 - [AlpinePackageUploadRequest](docs/AlpinePackageUploadRequest.md)
 - [Architecture](docs/Architecture.md)
 - [CargoPackageUpload](docs/CargoPackageUpload.md)
 - [CargoPackageUploadRequest](docs/CargoPackageUploadRequest.md)
 - [CocoapodsPackageUpload](docs/CocoapodsPackageUpload.md)
 - [CocoapodsPackageUploadRequest](docs/CocoapodsPackageUploadRequest.md)
 - [CommonBandwidthMetrics](docs/CommonBandwidthMetrics.md)
 - [CommonBandwidthMetricsValue](docs/CommonBandwidthMetricsValue.md)
 - [CommonDownloadsMetrics](docs/CommonDownloadsMetrics.md)
 - [CommonDownloadsMetricsValue](docs/CommonDownloadsMetricsValue.md)
 - [CommonMetrics](docs/CommonMetrics.md)
 - [ComposerPackageUpload](docs/ComposerPackageUpload.md)
 - [ComposerPackageUploadRequest](docs/ComposerPackageUploadRequest.md)
 - [ConanPackageUpload](docs/ConanPackageUpload.md)
 - [ConanPackageUploadRequest](docs/ConanPackageUploadRequest.md)
 - [CondaPackageUpload](docs/CondaPackageUpload.md)
 - [CondaPackageUploadRequest](docs/CondaPackageUploadRequest.md)
 - [CranPackageUpload](docs/CranPackageUpload.md)
 - [CranPackageUploadRequest](docs/CranPackageUploadRequest.md)
 - [DartPackageUpload](docs/DartPackageUpload.md)
 - [DartPackageUploadRequest](docs/DartPackageUploadRequest.md)
 - [DebPackageUpload](docs/DebPackageUpload.md)
 - [DebPackageUploadRequest](docs/DebPackageUploadRequest.md)
 - [Distribution](docs/Distribution.md)
 - [DistributionFull](docs/DistributionFull.md)
 - [DistributionVersion](docs/DistributionVersion.md)
 - [DockerPackageUpload](docs/DockerPackageUpload.md)
 - [DockerPackageUploadRequest](docs/DockerPackageUploadRequest.md)
 - [EntitlementUsageMetrics](docs/EntitlementUsageMetrics.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [Eula](docs/Eula.md)
 - [Format](docs/Format.md)
 - [FormatSupport](docs/FormatSupport.md)
 - [GeoIpLocation](docs/GeoIpLocation.md)
 - [GoPackageUpload](docs/GoPackageUpload.md)
 - [GoPackageUploadRequest](docs/GoPackageUploadRequest.md)
 - [HelmPackageUpload](docs/HelmPackageUpload.md)
 - [HelmPackageUploadRequest](docs/HelmPackageUploadRequest.md)
 - [History](docs/History.md)
 - [HistoryFieldset](docs/HistoryFieldset.md)
 - [LuarocksPackageUpload](docs/LuarocksPackageUpload.md)
 - [LuarocksPackageUploadRequest](docs/LuarocksPackageUploadRequest.md)
 - [MavenPackageUpload](docs/MavenPackageUpload.md)
 - [MavenPackageUploadRequest](docs/MavenPackageUploadRequest.md)
 - [Namespace](docs/Namespace.md)
 - [NamespaceAuditLog](docs/NamespaceAuditLog.md)
 - [NpmPackageUpload](docs/NpmPackageUpload.md)
 - [NpmPackageUploadRequest](docs/NpmPackageUploadRequest.md)
 - [NugetPackageUpload](docs/NugetPackageUpload.md)
 - [NugetPackageUploadRequest](docs/NugetPackageUploadRequest.md)
 - [Organization](docs/Organization.md)
 - [OrganizationGroupSync](docs/OrganizationGroupSync.md)
 - [OrganizationGroupSyncRequest](docs/OrganizationGroupSyncRequest.md)
 - [OrganizationInvite](docs/OrganizationInvite.md)
 - [OrganizationInviteExtend](docs/OrganizationInviteExtend.md)
 - [OrganizationInviteRequest](docs/OrganizationInviteRequest.md)
 - [OrganizationInviteUpdate](docs/OrganizationInviteUpdate.md)
 - [OrganizationInviteUpdateRequestPatch](docs/OrganizationInviteUpdateRequestPatch.md)
 - [OrganizationMembership](docs/OrganizationMembership.md)
 - [OrganizationTeam](docs/OrganizationTeam.md)
 - [OrganizationTeamMembers](docs/OrganizationTeamMembers.md)
 - [OrganizationTeamMembership](docs/OrganizationTeamMembership.md)
 - [OrganizationTeamRequest](docs/OrganizationTeamRequest.md)
 - [OrganizationTeamRequestPatch](docs/OrganizationTeamRequestPatch.md)
 - [P2PackageUpload](docs/P2PackageUpload.md)
 - [P2PackageUploadRequest](docs/P2PackageUploadRequest.md)
 - [Package](docs/Package.md)
 - [PackageCopy](docs/PackageCopy.md)
 - [PackageCopyRequest](docs/PackageCopyRequest.md)
 - [PackageDependencies](docs/PackageDependencies.md)
 - [PackageDependency](docs/PackageDependency.md)
 - [PackageFile](docs/PackageFile.md)
 - [PackageFilePartsUpload](docs/PackageFilePartsUpload.md)
 - [PackageFileUpload](docs/PackageFileUpload.md)
 - [PackageFileUploadRequest](docs/PackageFileUploadRequest.md)
 - [PackageMove](docs/PackageMove.md)
 - [PackageMoveRequest](docs/PackageMoveRequest.md)
 - [PackageQuarantine](docs/PackageQuarantine.md)
 - [PackageQuarantineRequest](docs/PackageQuarantineRequest.md)
 - [PackageResync](docs/PackageResync.md)
 - [PackageStatus](docs/PackageStatus.md)
 - [PackageTag](docs/PackageTag.md)
 - [PackageTagRequest](docs/PackageTagRequest.md)
 - [PackageUsageMetrics](docs/PackageUsageMetrics.md)
 - [PackageVulnerability](docs/PackageVulnerability.md)
 - [PythonPackageUpload](docs/PythonPackageUpload.md)
 - [PythonPackageUploadRequest](docs/PythonPackageUploadRequest.md)
 - [Quota](docs/Quota.md)
 - [QuotaHistory](docs/QuotaHistory.md)
 - [RateCheck](docs/RateCheck.md)
 - [RawPackageUpload](docs/RawPackageUpload.md)
 - [RawPackageUploadRequest](docs/RawPackageUploadRequest.md)
 - [ReposGeoipRead200Response](docs/ReposGeoipRead200Response.md)
 - [ReposGeoipRead200ResponseCountryCode](docs/ReposGeoipRead200ResponseCountryCode.md)
 - [Repository](docs/Repository.md)
 - [RepositoryAuditLog](docs/RepositoryAuditLog.md)
 - [RepositoryCreate](docs/RepositoryCreate.md)
 - [RepositoryCreateRequest](docs/RepositoryCreateRequest.md)
 - [RepositoryGeoIPList](docs/RepositoryGeoIPList.md)
 - [RepositoryGeoIPTestAddress](docs/RepositoryGeoIPTestAddress.md)
 - [RepositoryGeoIPTestAddressResponse](docs/RepositoryGeoIPTestAddressResponse.md)
 - [RepositoryGeoIPTestAddressResponseDict](docs/RepositoryGeoIPTestAddressResponseDict.md)
 - [RepositoryGpgKey](docs/RepositoryGpgKey.md)
 - [RepositoryGpgKeyCreate](docs/RepositoryGpgKeyCreate.md)
 - [RepositoryPrivilegeDict](docs/RepositoryPrivilegeDict.md)
 - [RepositoryPrivilegeInput](docs/RepositoryPrivilegeInput.md)
 - [RepositoryPrivilegeInputRequest](docs/RepositoryPrivilegeInputRequest.md)
 - [RepositoryPrivilegeInputRequestPatch](docs/RepositoryPrivilegeInputRequestPatch.md)
 - [RepositoryRequestPatch](docs/RepositoryRequestPatch.md)
 - [RepositoryRsaKey](docs/RepositoryRsaKey.md)
 - [RepositoryRsaKeyCreate](docs/RepositoryRsaKeyCreate.md)
 - [RepositoryToken](docs/RepositoryToken.md)
 - [RepositoryTokenRefresh](docs/RepositoryTokenRefresh.md)
 - [RepositoryTokenRefreshRequest](docs/RepositoryTokenRefreshRequest.md)
 - [RepositoryTokenRequest](docs/RepositoryTokenRequest.md)
 - [RepositoryTokenRequestPatch](docs/RepositoryTokenRequestPatch.md)
 - [RepositoryTokenSync](docs/RepositoryTokenSync.md)
 - [RepositoryTokenSyncRequest](docs/RepositoryTokenSyncRequest.md)
 - [RepositoryWebhook](docs/RepositoryWebhook.md)
 - [RepositoryWebhookRequest](docs/RepositoryWebhookRequest.md)
 - [RepositoryWebhookRequestPatch](docs/RepositoryWebhookRequestPatch.md)
 - [ResourcesRateCheck](docs/ResourcesRateCheck.md)
 - [RpmPackageUpload](docs/RpmPackageUpload.md)
 - [RpmPackageUploadRequest](docs/RpmPackageUploadRequest.md)
 - [RubyPackageUpload](docs/RubyPackageUpload.md)
 - [RubyPackageUploadRequest](docs/RubyPackageUploadRequest.md)
 - [Service](docs/Service.md)
 - [ServiceRequest](docs/ServiceRequest.md)
 - [ServiceRequestPatch](docs/ServiceRequestPatch.md)
 - [ServiceTeams](docs/ServiceTeams.md)
 - [StatusBasic](docs/StatusBasic.md)
 - [StorageRegion](docs/StorageRegion.md)
 - [TerraformPackageUpload](docs/TerraformPackageUpload.md)
 - [TerraformPackageUploadRequest](docs/TerraformPackageUploadRequest.md)
 - [Usage](docs/Usage.md)
 - [UsageFieldset](docs/UsageFieldset.md)
 - [UsageLimits](docs/UsageLimits.md)
 - [UserAuthToken](docs/UserAuthToken.md)
 - [UserAuthTokenRequest](docs/UserAuthTokenRequest.md)
 - [UserBrief](docs/UserBrief.md)
 - [UserProfile](docs/UserProfile.md)
 - [VagrantPackageUpload](docs/VagrantPackageUpload.md)
 - [VagrantPackageUploadRequest](docs/VagrantPackageUploadRequest.md)
 - [Vulnerability](docs/Vulnerability.md)
 - [VulnerabilityScan](docs/VulnerabilityScan.md)
 - [VulnerabilityScanResults](docs/VulnerabilityScanResults.md)
 - [VulnerabilityScanResultsList](docs/VulnerabilityScanResultsList.md)
 - [VulnerabilityScanVersion](docs/VulnerabilityScanVersion.md)
 - [WebhookTemplate](docs/WebhookTemplate.md)


## Documentation For Authorization



### apikey

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.


### basic

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cloudsmith.io

