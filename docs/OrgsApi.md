# \OrgsApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsDenyPolicyCreate**](OrgsApi.md#OrgsDenyPolicyCreate) | **Post** /orgs/{org}/deny-policy/ | Create a package deny policy.
[**OrgsDenyPolicyDelete**](OrgsApi.md#OrgsDenyPolicyDelete) | **Delete** /orgs/{org}/deny-policy/{slug_perm}/ | Delete a package deny policy.
[**OrgsDenyPolicyList**](OrgsApi.md#OrgsDenyPolicyList) | **Get** /orgs/{org}/deny-policy/ | Get a list of all package deny policies.
[**OrgsDenyPolicyPartialUpdate**](OrgsApi.md#OrgsDenyPolicyPartialUpdate) | **Patch** /orgs/{org}/deny-policy/{slug_perm}/ | Partially update a package deny policy.
[**OrgsDenyPolicyRead**](OrgsApi.md#OrgsDenyPolicyRead) | **Get** /orgs/{org}/deny-policy/{slug_perm}/ | Get a package deny policy.
[**OrgsDenyPolicyUpdate**](OrgsApi.md#OrgsDenyPolicyUpdate) | **Put** /orgs/{org}/deny-policy/{slug_perm}/ | Update a package deny policy.
[**OrgsInvitesCreate**](OrgsApi.md#OrgsInvitesCreate) | **Post** /orgs/{org}/invites/ | Create an organization invite for a specific user
[**OrgsInvitesDelete**](OrgsApi.md#OrgsInvitesDelete) | **Delete** /orgs/{org}/invites/{slug_perm}/ | Delete a specific organization invite
[**OrgsInvitesExtend**](OrgsApi.md#OrgsInvitesExtend) | **Post** /orgs/{org}/invites/{slug_perm}/extend/ | Extend an organization invite.
[**OrgsInvitesList**](OrgsApi.md#OrgsInvitesList) | **Get** /orgs/{org}/invites/ | Get a list of all invites for an organization.
[**OrgsInvitesPartialUpdate**](OrgsApi.md#OrgsInvitesPartialUpdate) | **Patch** /orgs/{org}/invites/{slug_perm}/ | Update a specific organization invite.
[**OrgsInvitesResend**](OrgsApi.md#OrgsInvitesResend) | **Post** /orgs/{org}/invites/{slug_perm}/resend/ | Resend an organization invite.
[**OrgsLicensePolicyCreate**](OrgsApi.md#OrgsLicensePolicyCreate) | **Post** /orgs/{org}/license-policy/ | Create a package license policy.
[**OrgsLicensePolicyDelete**](OrgsApi.md#OrgsLicensePolicyDelete) | **Delete** /orgs/{org}/license-policy/{slug_perm}/ | Delete a package license policy.
[**OrgsLicensePolicyEvaluationCreate**](OrgsApi.md#OrgsLicensePolicyEvaluationCreate) | **Post** /orgs/{org}/license-policy/{policy_slug_perm}/evaluation/ | Create an evaluation request for this policy.
[**OrgsLicensePolicyEvaluationList**](OrgsApi.md#OrgsLicensePolicyEvaluationList) | **Get** /orgs/{org}/license-policy/{policy_slug_perm}/evaluation/ | List evaluation requests for this policy.
[**OrgsLicensePolicyEvaluationRead**](OrgsApi.md#OrgsLicensePolicyEvaluationRead) | **Get** /orgs/{org}/license-policy/{policy_slug_perm}/evaluation/{slug_perm}/ | Retrieve an evaluation request for this policy.
[**OrgsLicensePolicyList**](OrgsApi.md#OrgsLicensePolicyList) | **Get** /orgs/{org}/license-policy/ | Get a list of all package license policies.
[**OrgsLicensePolicyPartialUpdate**](OrgsApi.md#OrgsLicensePolicyPartialUpdate) | **Patch** /orgs/{org}/license-policy/{slug_perm}/ | Partially update a package license policy.
[**OrgsLicensePolicyRead**](OrgsApi.md#OrgsLicensePolicyRead) | **Get** /orgs/{org}/license-policy/{slug_perm}/ | Get a package license policy.
[**OrgsLicensePolicyUpdate**](OrgsApi.md#OrgsLicensePolicyUpdate) | **Put** /orgs/{org}/license-policy/{slug_perm}/ | Update a package license policy.
[**OrgsLicensePolicyViolationList**](OrgsApi.md#OrgsLicensePolicyViolationList) | **Get** /orgs/{org}/license-policy-violation/ | List all current license policy violations for this Organization.
[**OrgsList**](OrgsApi.md#OrgsList) | **Get** /orgs/ | Get a list of all the organizations you are associated with.
[**OrgsMembersDelete**](OrgsApi.md#OrgsMembersDelete) | **Delete** /orgs/{org}/members/{member}/ | Removes a member from the organization.
[**OrgsMembersList**](OrgsApi.md#OrgsMembersList) | **Get** /orgs/{org}/members/ | Get the details for all organization members.
[**OrgsMembersRead**](OrgsApi.md#OrgsMembersRead) | **Get** /orgs/{org}/members/{member}/ | Get the details for a specific organization member.
[**OrgsMembersRefresh**](OrgsApi.md#OrgsMembersRefresh) | **Post** /orgs/{org}/members/{member}/refresh/ | Refresh a member of the organization&#39;s API key.
[**OrgsMembersRemove**](OrgsApi.md#OrgsMembersRemove) | **Get** /orgs/{org}/members/{member}/remove/ | Removes a member from the organization (deprecated, use DELETE instead).
[**OrgsRead**](OrgsApi.md#OrgsRead) | **Get** /orgs/{org}/ | Get the details for the specific organization.
[**OrgsSamlGroupSyncCreate**](OrgsApi.md#OrgsSamlGroupSyncCreate) | **Post** /orgs/{org}/saml-group-sync/ | Create a new SAML Group Sync mapping within an organization.
[**OrgsSamlGroupSyncDelete**](OrgsApi.md#OrgsSamlGroupSyncDelete) | **Delete** /orgs/{org}/saml-group-sync/{slug_perm}/ | Delete a SAML Group Sync mapping from an organization.
[**OrgsSamlGroupSyncList**](OrgsApi.md#OrgsSamlGroupSyncList) | **Get** /orgs/{org}/saml-group-sync/ | Get the details of all SAML Group Sync mapping within an organization.
[**OrgsServicesCreate**](OrgsApi.md#OrgsServicesCreate) | **Post** /orgs/{org}/services/ | Create a service within an organization.
[**OrgsServicesDelete**](OrgsApi.md#OrgsServicesDelete) | **Delete** /orgs/{org}/services/{service}/ | Delete a specific service
[**OrgsServicesList**](OrgsApi.md#OrgsServicesList) | **Get** /orgs/{org}/services/ | Get a list of all services within an organization.
[**OrgsServicesPartialUpdate**](OrgsApi.md#OrgsServicesPartialUpdate) | **Patch** /orgs/{org}/services/{service}/ | Update a service within an organization.
[**OrgsServicesRead**](OrgsApi.md#OrgsServicesRead) | **Get** /orgs/{org}/services/{service}/ | Retrieve details of a single service within an organization.
[**OrgsServicesRefresh**](OrgsApi.md#OrgsServicesRefresh) | **Post** /orgs/{org}/services/{service}/refresh/ | Refresh service API token.
[**OrgsTeamsCreate**](OrgsApi.md#OrgsTeamsCreate) | **Post** /orgs/{org}/teams/ | Create a team for this organization.
[**OrgsTeamsDelete**](OrgsApi.md#OrgsTeamsDelete) | **Delete** /orgs/{org}/teams/{team}/ | Delete a specific team in a organization.
[**OrgsTeamsList**](OrgsApi.md#OrgsTeamsList) | **Get** /orgs/{org}/teams/ | Get the details of all teams within an organization.
[**OrgsTeamsMembersCreate**](OrgsApi.md#OrgsTeamsMembersCreate) | **Post** /orgs/{org}/teams/{team}/members | Add users to a team.
[**OrgsTeamsMembersList**](OrgsApi.md#OrgsTeamsMembersList) | **Get** /orgs/{org}/teams/{team}/members | List all members for the team.
[**OrgsTeamsMembersUpdate**](OrgsApi.md#OrgsTeamsMembersUpdate) | **Put** /orgs/{org}/teams/{team}/members | Replace all team members.
[**OrgsTeamsPartialUpdate**](OrgsApi.md#OrgsTeamsPartialUpdate) | **Patch** /orgs/{org}/teams/{team}/ | Update a specific team in a organization.
[**OrgsTeamsRead**](OrgsApi.md#OrgsTeamsRead) | **Get** /orgs/{org}/teams/{team}/ | Get the details of a specific team within an organization.
[**OrgsVulnerabilityPolicyCreate**](OrgsApi.md#OrgsVulnerabilityPolicyCreate) | **Post** /orgs/{org}/vulnerability-policy/ | Create a package vulnerability policy.
[**OrgsVulnerabilityPolicyDelete**](OrgsApi.md#OrgsVulnerabilityPolicyDelete) | **Delete** /orgs/{org}/vulnerability-policy/{slug_perm}/ | Delete a package vulnerability policy.
[**OrgsVulnerabilityPolicyEvaluationCreate**](OrgsApi.md#OrgsVulnerabilityPolicyEvaluationCreate) | **Post** /orgs/{org}/vulnerability-policy/{policy_slug_perm}/evaluation/ | Create an evaluation request for this policy.
[**OrgsVulnerabilityPolicyEvaluationList**](OrgsApi.md#OrgsVulnerabilityPolicyEvaluationList) | **Get** /orgs/{org}/vulnerability-policy/{policy_slug_perm}/evaluation/ | List evaluation requests for this policy.
[**OrgsVulnerabilityPolicyEvaluationRead**](OrgsApi.md#OrgsVulnerabilityPolicyEvaluationRead) | **Get** /orgs/{org}/vulnerability-policy/{policy_slug_perm}/evaluation/{slug_perm}/ | Retrieve an evaluation request for this policy.
[**OrgsVulnerabilityPolicyList**](OrgsApi.md#OrgsVulnerabilityPolicyList) | **Get** /orgs/{org}/vulnerability-policy/ | Get a list of all package vulnerability policies.
[**OrgsVulnerabilityPolicyPartialUpdate**](OrgsApi.md#OrgsVulnerabilityPolicyPartialUpdate) | **Patch** /orgs/{org}/vulnerability-policy/{slug_perm}/ | Partially update a package vulnerability policy.
[**OrgsVulnerabilityPolicyRead**](OrgsApi.md#OrgsVulnerabilityPolicyRead) | **Get** /orgs/{org}/vulnerability-policy/{slug_perm}/ | Get a package vulnerability policy.
[**OrgsVulnerabilityPolicyUpdate**](OrgsApi.md#OrgsVulnerabilityPolicyUpdate) | **Put** /orgs/{org}/vulnerability-policy/{slug_perm}/ | Update a package vulnerability policy.
[**OrgsVulnerabilityPolicyViolationList**](OrgsApi.md#OrgsVulnerabilityPolicyViolationList) | **Get** /orgs/{org}/vulnerability-policy-violation/ | List all current vulnerability policy violations for this Organization.



## OrgsDenyPolicyCreate

> PackageDenyPolicy OrgsDenyPolicyCreate(ctx, org).Data(data).Execute()

Create a package deny policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewPackageDenyPolicyRequest("PackageQueryString_example") // PackageDenyPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsDenyPolicyCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDenyPolicyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsDenyPolicyCreate`: PackageDenyPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsDenyPolicyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDenyPolicyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**PackageDenyPolicyRequest**](PackageDenyPolicyRequest.md) |  | 

### Return type

[**PackageDenyPolicy**](PackageDenyPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsDenyPolicyDelete

> OrgsDenyPolicyDelete(ctx, org, slugPerm).Execute()

Delete a package deny policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsDenyPolicyDelete(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDenyPolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDenyPolicyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsDenyPolicyList

> []PackageDenyPolicy OrgsDenyPolicyList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get a list of all package deny policies.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsDenyPolicyList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDenyPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsDenyPolicyList`: []PackageDenyPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsDenyPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDenyPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]PackageDenyPolicy**](PackageDenyPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsDenyPolicyPartialUpdate

> PackageDenyPolicy OrgsDenyPolicyPartialUpdate(ctx, org, slugPerm).Data(data).Execute()

Partially update a package deny policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewPackageDenyPolicyRequestPatch() // PackageDenyPolicyRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsDenyPolicyPartialUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDenyPolicyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsDenyPolicyPartialUpdate`: PackageDenyPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsDenyPolicyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDenyPolicyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PackageDenyPolicyRequestPatch**](PackageDenyPolicyRequestPatch.md) |  | 

### Return type

[**PackageDenyPolicy**](PackageDenyPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsDenyPolicyRead

> PackageDenyPolicy OrgsDenyPolicyRead(ctx, org, slugPerm).Execute()

Get a package deny policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsDenyPolicyRead(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDenyPolicyRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsDenyPolicyRead`: PackageDenyPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsDenyPolicyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDenyPolicyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PackageDenyPolicy**](PackageDenyPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsDenyPolicyUpdate

> PackageDenyPolicy OrgsDenyPolicyUpdate(ctx, org, slugPerm).Data(data).Execute()

Update a package deny policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewPackageDenyPolicyRequest("PackageQueryString_example") // PackageDenyPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsDenyPolicyUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDenyPolicyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsDenyPolicyUpdate`: PackageDenyPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsDenyPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDenyPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**PackageDenyPolicyRequest**](PackageDenyPolicyRequest.md) |  | 

### Return type

[**PackageDenyPolicy**](PackageDenyPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsInvitesCreate

> OrganizationInvite OrgsInvitesCreate(ctx, org).Data(data).Execute()

Create an organization invite for a specific user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewOrganizationInviteRequest() // OrganizationInviteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsInvitesCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsInvitesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsInvitesCreate`: OrganizationInvite
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsInvitesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsInvitesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**OrganizationInviteRequest**](OrganizationInviteRequest.md) |  | 

### Return type

[**OrganizationInvite**](OrganizationInvite.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsInvitesDelete

> OrgsInvitesDelete(ctx, org, slugPerm).Execute()

Delete a specific organization invite



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsInvitesDelete(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsInvitesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsInvitesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsInvitesExtend

> OrganizationInviteExtend OrgsInvitesExtend(ctx, org, slugPerm).Execute()

Extend an organization invite.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsInvitesExtend(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsInvitesExtend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsInvitesExtend`: OrganizationInviteExtend
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsInvitesExtend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsInvitesExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationInviteExtend**](OrganizationInviteExtend.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsInvitesList

> []OrganizationInvite OrgsInvitesList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get a list of all invites for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsInvitesList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsInvitesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsInvitesList`: []OrganizationInvite
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsInvitesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsInvitesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]OrganizationInvite**](OrganizationInvite.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsInvitesPartialUpdate

> OrganizationInvite OrgsInvitesPartialUpdate(ctx, org, slugPerm).Data(data).Execute()

Update a specific organization invite.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewOrganizationInviteUpdateRequestPatch() // OrganizationInviteUpdateRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsInvitesPartialUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsInvitesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsInvitesPartialUpdate`: OrganizationInvite
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsInvitesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsInvitesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationInviteUpdateRequestPatch**](OrganizationInviteUpdateRequestPatch.md) |  | 

### Return type

[**OrganizationInvite**](OrganizationInvite.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsInvitesResend

> OrganizationInviteExtend OrgsInvitesResend(ctx, org, slugPerm).Execute()

Resend an organization invite.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsInvitesResend(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsInvitesResend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsInvitesResend`: OrganizationInviteExtend
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsInvitesResend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsInvitesResendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationInviteExtend**](OrganizationInviteExtend.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyCreate

> OrganizationPackageLicensePolicy OrgsLicensePolicyCreate(ctx, org).Data(data).Execute()

Create a package license policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewOrganizationPackageLicensePolicyRequest("Name_example", []string{"SpdxIdentifiers_example"}) // OrganizationPackageLicensePolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyCreate`: OrganizationPackageLicensePolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**OrganizationPackageLicensePolicyRequest**](OrganizationPackageLicensePolicyRequest.md) |  | 

### Return type

[**OrganizationPackageLicensePolicy**](OrganizationPackageLicensePolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyDelete

> OrgsLicensePolicyDelete(ctx, org, slugPerm).Execute()

Delete a package license policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsLicensePolicyDelete(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyEvaluationCreate

> PackageLicensePolicyEvaluationRequest OrgsLicensePolicyEvaluationCreate(ctx, org, policySlugPerm).Data(data).Execute()

Create an evaluation request for this policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    policySlugPerm := "policySlugPerm_example" // string | 
    data := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyEvaluationCreate(context.Background(), org, policySlugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyEvaluationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyEvaluationCreate`: PackageLicensePolicyEvaluationRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyEvaluationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**policySlugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyEvaluationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **map[string]interface{}** |  | 

### Return type

[**PackageLicensePolicyEvaluationRequest**](PackageLicensePolicyEvaluationRequest.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyEvaluationList

> []PackageLicensePolicyEvaluationRequest OrgsLicensePolicyEvaluationList(ctx, org, policySlugPerm).Page(page).PageSize(pageSize).Execute()

List evaluation requests for this policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    policySlugPerm := "policySlugPerm_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyEvaluationList(context.Background(), org, policySlugPerm).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyEvaluationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyEvaluationList`: []PackageLicensePolicyEvaluationRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyEvaluationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**policySlugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyEvaluationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]PackageLicensePolicyEvaluationRequest**](PackageLicensePolicyEvaluationRequest.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyEvaluationRead

> PackageLicensePolicyEvaluationRequest OrgsLicensePolicyEvaluationRead(ctx, org, policySlugPerm, slugPerm).Execute()

Retrieve an evaluation request for this policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    policySlugPerm := "policySlugPerm_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyEvaluationRead(context.Background(), org, policySlugPerm, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyEvaluationRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyEvaluationRead`: PackageLicensePolicyEvaluationRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyEvaluationRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**policySlugPerm** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyEvaluationReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageLicensePolicyEvaluationRequest**](PackageLicensePolicyEvaluationRequest.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyList

> []OrganizationPackageLicensePolicy OrgsLicensePolicyList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get a list of all package license policies.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyList`: []OrganizationPackageLicensePolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]OrganizationPackageLicensePolicy**](OrganizationPackageLicensePolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyPartialUpdate

> OrganizationPackageLicensePolicy OrgsLicensePolicyPartialUpdate(ctx, org, slugPerm).Data(data).Execute()

Partially update a package license policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewOrganizationPackageLicensePolicyRequestPatch() // OrganizationPackageLicensePolicyRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyPartialUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyPartialUpdate`: OrganizationPackageLicensePolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationPackageLicensePolicyRequestPatch**](OrganizationPackageLicensePolicyRequestPatch.md) |  | 

### Return type

[**OrganizationPackageLicensePolicy**](OrganizationPackageLicensePolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyRead

> OrganizationPackageLicensePolicy OrgsLicensePolicyRead(ctx, org, slugPerm).Execute()

Get a package license policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyRead(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyRead`: OrganizationPackageLicensePolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationPackageLicensePolicy**](OrganizationPackageLicensePolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyUpdate

> OrganizationPackageLicensePolicy OrgsLicensePolicyUpdate(ctx, org, slugPerm).Data(data).Execute()

Update a package license policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewOrganizationPackageLicensePolicyRequest("Name_example", []string{"SpdxIdentifiers_example"}) // OrganizationPackageLicensePolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyUpdate`: OrganizationPackageLicensePolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationPackageLicensePolicyRequest**](OrganizationPackageLicensePolicyRequest.md) |  | 

### Return type

[**OrganizationPackageLicensePolicy**](OrganizationPackageLicensePolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsLicensePolicyViolationList

> PackageLicensePolicyViolationLogCursorPage OrgsLicensePolicyViolationList(ctx, org).Cursor(cursor).PageSize(pageSize).Execute()

List all current license policy violations for this Organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    cursor := "cursor_example" // string | The pagination cursor value. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsLicensePolicyViolationList(context.Background(), org).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsLicensePolicyViolationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsLicensePolicyViolationList`: PackageLicensePolicyViolationLogCursorPage
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsLicensePolicyViolationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsLicensePolicyViolationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**PackageLicensePolicyViolationLogCursorPage**](PackageLicensePolicyViolationLogCursorPage.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsList

> []Organization OrgsList(ctx).Page(page).PageSize(pageSize).Execute()

Get a list of all the organizations you are associated with.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsList`: []Organization
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]Organization**](Organization.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsMembersDelete

> OrgsMembersDelete(ctx, org, member).Execute()

Removes a member from the organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    member := "member_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsMembersDelete(context.Background(), org, member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsMembersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**member** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsMembersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsMembersList

> []OrganizationMembership OrgsMembersList(ctx, org).Page(page).PageSize(pageSize).IsActive(isActive).Execute()

Get the details for all organization members.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)
    isActive := true // bool | Filter for active/inactive users. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsMembersList(context.Background(), org).Page(page).PageSize(pageSize).IsActive(isActive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsMembersList`: []OrganizationMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsMembersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsMembersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **isActive** | **bool** | Filter for active/inactive users. | [default to false]

### Return type

[**[]OrganizationMembership**](OrganizationMembership.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsMembersRead

> OrganizationMembership OrgsMembersRead(ctx, org, member).Execute()

Get the details for a specific organization member.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    member := "member_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsMembersRead(context.Background(), org, member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsMembersRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsMembersRead`: OrganizationMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsMembersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**member** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsMembersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMembership**](OrganizationMembership.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsMembersRefresh

> OrgsMembersRefresh(ctx, org, member).Execute()

Refresh a member of the organization's API key.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    member := "member_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsMembersRefresh(context.Background(), org, member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsMembersRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**member** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsMembersRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsMembersRemove

> OrgsMembersRemove(ctx, org, member).Execute()

Removes a member from the organization (deprecated, use DELETE instead).



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    member := "member_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsMembersRemove(context.Background(), org, member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsMembersRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**member** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsMembersRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsRead

> Organization OrgsRead(ctx, org).Execute()

Get the details for the specific organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRead(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsRead`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsSamlGroupSyncCreate

> OrganizationGroupSync OrgsSamlGroupSyncCreate(ctx, org).Data(data).Execute()

Create a new SAML Group Sync mapping within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewOrganizationGroupSyncRequest("IdpKey_example", "IdpValue_example", "Organization_example", "Team_example") // OrganizationGroupSyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsSamlGroupSyncCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsSamlGroupSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsSamlGroupSyncCreate`: OrganizationGroupSync
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsSamlGroupSyncCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsSamlGroupSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**OrganizationGroupSyncRequest**](OrganizationGroupSyncRequest.md) |  | 

### Return type

[**OrganizationGroupSync**](OrganizationGroupSync.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsSamlGroupSyncDelete

> OrgsSamlGroupSyncDelete(ctx, org, slugPerm).Execute()

Delete a SAML Group Sync mapping from an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsSamlGroupSyncDelete(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsSamlGroupSyncDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsSamlGroupSyncDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsSamlGroupSyncList

> []OrganizationGroupSync OrgsSamlGroupSyncList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get the details of all SAML Group Sync mapping within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsSamlGroupSyncList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsSamlGroupSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsSamlGroupSyncList`: []OrganizationGroupSync
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsSamlGroupSyncList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsSamlGroupSyncListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]OrganizationGroupSync**](OrganizationGroupSync.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsServicesCreate

> Service OrgsServicesCreate(ctx, org).Data(data).Execute()

Create a service within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewServiceRequest("Name_example") // ServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsServicesCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsServicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsServicesCreate`: Service
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsServicesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsServicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**ServiceRequest**](ServiceRequest.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsServicesDelete

> OrgsServicesDelete(ctx, org, service).Execute()

Delete a specific service



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsServicesDelete(context.Background(), org, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsServicesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsServicesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsServicesList

> []Service OrgsServicesList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get a list of all services within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsServicesList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsServicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsServicesList`: []Service
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsServicesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsServicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]Service**](Service.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsServicesPartialUpdate

> Service OrgsServicesPartialUpdate(ctx, org, service).Data(data).Execute()

Update a service within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    service := "service_example" // string | 
    data := *openapiclient.NewServiceRequestPatch() // ServiceRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsServicesPartialUpdate(context.Background(), org, service).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsServicesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsServicesPartialUpdate`: Service
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsServicesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsServicesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ServiceRequestPatch**](ServiceRequestPatch.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsServicesRead

> Service OrgsServicesRead(ctx, org, service).Execute()

Retrieve details of a single service within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsServicesRead(context.Background(), org, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsServicesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsServicesRead`: Service
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsServicesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsServicesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Service**](Service.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsServicesRefresh

> Service OrgsServicesRefresh(ctx, org, service).Execute()

Refresh service API token.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsServicesRefresh(context.Background(), org, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsServicesRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsServicesRefresh`: Service
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsServicesRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsServicesRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Service**](Service.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsCreate

> OrganizationTeam OrgsTeamsCreate(ctx, org).Data(data).Execute()

Create a team for this organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewOrganizationTeamRequest("Name_example") // OrganizationTeamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsCreate`: OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**OrganizationTeamRequest**](OrganizationTeamRequest.md) |  | 

### Return type

[**OrganizationTeam**](OrganizationTeam.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsDelete

> OrgsTeamsDelete(ctx, org, team).Execute()

Delete a specific team in a organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsTeamsDelete(context.Background(), org, team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsList

> []OrganizationTeam OrgsTeamsList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get the details of all teams within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsList`: []OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]OrganizationTeam**](OrganizationTeam.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsMembersCreate

> OrganizationTeamMembers OrgsTeamsMembersCreate(ctx, org, team).Data(data).Execute()

Add users to a team.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 
    data := *openapiclient.NewOrganizationTeamMembers([]openapiclient.OrganizationTeamMembership{*openapiclient.NewOrganizationTeamMembership("Role_example", "User_example")}) // OrganizationTeamMembers |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsMembersCreate(context.Background(), org, team).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsMembersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsMembersCreate`: OrganizationTeamMembers
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsMembersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsMembersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationTeamMembers**](OrganizationTeamMembers.md) |  | 

### Return type

[**OrganizationTeamMembers**](OrganizationTeamMembers.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsMembersList

> OrganizationTeamMembers OrgsTeamsMembersList(ctx, org, team).Execute()

List all members for the team.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsMembersList(context.Background(), org, team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsMembersList`: OrganizationTeamMembers
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsMembersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsMembersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationTeamMembers**](OrganizationTeamMembers.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsMembersUpdate

> OrganizationTeamMembers OrgsTeamsMembersUpdate(ctx, org, team).Execute()

Replace all team members.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsMembersUpdate(context.Background(), org, team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsMembersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsMembersUpdate`: OrganizationTeamMembers
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsMembersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsMembersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationTeamMembers**](OrganizationTeamMembers.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsPartialUpdate

> OrganizationTeam OrgsTeamsPartialUpdate(ctx, org, team).Data(data).Execute()

Update a specific team in a organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 
    data := *openapiclient.NewOrganizationTeamRequestPatch() // OrganizationTeamRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsPartialUpdate(context.Background(), org, team).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsPartialUpdate`: OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationTeamRequestPatch**](OrganizationTeamRequestPatch.md) |  | 

### Return type

[**OrganizationTeam**](OrganizationTeam.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsTeamsRead

> OrganizationTeam OrgsTeamsRead(ctx, org, team).Execute()

Get the details of a specific team within an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsRead(context.Background(), org, team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsTeamsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsTeamsRead`: OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsTeamsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsTeamsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationTeam**](OrganizationTeam.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyCreate

> OrganizationPackageVulnerabilityPolicy OrgsVulnerabilityPolicyCreate(ctx, org).Data(data).Execute()

Create a package vulnerability policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    data := *openapiclient.NewOrganizationPackageVulnerabilityPolicyRequest("Name_example") // OrganizationPackageVulnerabilityPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyCreate(context.Background(), org).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyCreate`: OrganizationPackageVulnerabilityPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**OrganizationPackageVulnerabilityPolicyRequest**](OrganizationPackageVulnerabilityPolicyRequest.md) |  | 

### Return type

[**OrganizationPackageVulnerabilityPolicy**](OrganizationPackageVulnerabilityPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyDelete

> OrgsVulnerabilityPolicyDelete(ctx, org, slugPerm).Execute()

Delete a package vulnerability policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyDelete(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyEvaluationCreate

> PackageVulnerabilityPolicyEvaluationRequest OrgsVulnerabilityPolicyEvaluationCreate(ctx, org, policySlugPerm).Data(data).Execute()

Create an evaluation request for this policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    policySlugPerm := "policySlugPerm_example" // string | 
    data := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyEvaluationCreate(context.Background(), org, policySlugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyEvaluationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyEvaluationCreate`: PackageVulnerabilityPolicyEvaluationRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyEvaluationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**policySlugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyEvaluationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **map[string]interface{}** |  | 

### Return type

[**PackageVulnerabilityPolicyEvaluationRequest**](PackageVulnerabilityPolicyEvaluationRequest.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyEvaluationList

> []PackageVulnerabilityPolicyEvaluationRequest OrgsVulnerabilityPolicyEvaluationList(ctx, org, policySlugPerm).Page(page).PageSize(pageSize).Execute()

List evaluation requests for this policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    policySlugPerm := "policySlugPerm_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyEvaluationList(context.Background(), org, policySlugPerm).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyEvaluationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyEvaluationList`: []PackageVulnerabilityPolicyEvaluationRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyEvaluationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**policySlugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyEvaluationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]PackageVulnerabilityPolicyEvaluationRequest**](PackageVulnerabilityPolicyEvaluationRequest.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyEvaluationRead

> PackageVulnerabilityPolicyEvaluationRequest OrgsVulnerabilityPolicyEvaluationRead(ctx, org, policySlugPerm, slugPerm).Execute()

Retrieve an evaluation request for this policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    policySlugPerm := "policySlugPerm_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyEvaluationRead(context.Background(), org, policySlugPerm, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyEvaluationRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyEvaluationRead`: PackageVulnerabilityPolicyEvaluationRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyEvaluationRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**policySlugPerm** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyEvaluationReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageVulnerabilityPolicyEvaluationRequest**](PackageVulnerabilityPolicyEvaluationRequest.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyList

> []OrganizationPackageVulnerabilityPolicy OrgsVulnerabilityPolicyList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get a list of all package vulnerability policies.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyList`: []OrganizationPackageVulnerabilityPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]OrganizationPackageVulnerabilityPolicy**](OrganizationPackageVulnerabilityPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyPartialUpdate

> OrganizationPackageVulnerabilityPolicy OrgsVulnerabilityPolicyPartialUpdate(ctx, org, slugPerm).Data(data).Execute()

Partially update a package vulnerability policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewOrganizationPackageVulnerabilityPolicyRequestPatch() // OrganizationPackageVulnerabilityPolicyRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyPartialUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyPartialUpdate`: OrganizationPackageVulnerabilityPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationPackageVulnerabilityPolicyRequestPatch**](OrganizationPackageVulnerabilityPolicyRequestPatch.md) |  | 

### Return type

[**OrganizationPackageVulnerabilityPolicy**](OrganizationPackageVulnerabilityPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyRead

> OrganizationPackageVulnerabilityPolicy OrgsVulnerabilityPolicyRead(ctx, org, slugPerm).Execute()

Get a package vulnerability policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyRead(context.Background(), org, slugPerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyRead`: OrganizationPackageVulnerabilityPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationPackageVulnerabilityPolicy**](OrganizationPackageVulnerabilityPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyUpdate

> OrganizationPackageVulnerabilityPolicy OrgsVulnerabilityPolicyUpdate(ctx, org, slugPerm).Data(data).Execute()

Update a package vulnerability policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 
    data := *openapiclient.NewOrganizationPackageVulnerabilityPolicyRequest("Name_example") // OrganizationPackageVulnerabilityPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyUpdate(context.Background(), org, slugPerm).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyUpdate`: OrganizationPackageVulnerabilityPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**slugPerm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**OrganizationPackageVulnerabilityPolicyRequest**](OrganizationPackageVulnerabilityPolicyRequest.md) |  | 

### Return type

[**OrganizationPackageVulnerabilityPolicy**](OrganizationPackageVulnerabilityPolicy.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsVulnerabilityPolicyViolationList

> PackageVulnerabilityPolicyViolationLogCursorPage OrgsVulnerabilityPolicyViolationList(ctx, org).Cursor(cursor).PageSize(pageSize).Execute()

List all current vulnerability policy violations for this Organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
    org := "org_example" // string | 
    cursor := "cursor_example" // string | The pagination cursor value. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsVulnerabilityPolicyViolationList(context.Background(), org).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsVulnerabilityPolicyViolationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsVulnerabilityPolicyViolationList`: PackageVulnerabilityPolicyViolationLogCursorPage
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsVulnerabilityPolicyViolationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsVulnerabilityPolicyViolationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**PackageVulnerabilityPolicyViolationLogCursorPage**](PackageVulnerabilityPolicyViolationLogCursorPage.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

