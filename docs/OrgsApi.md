# \OrgsApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsInvitesCreate**](OrgsApi.md#OrgsInvitesCreate) | **Post** /orgs/{org}/invites/ | Create an organization invite for a specific user
[**OrgsInvitesDelete**](OrgsApi.md#OrgsInvitesDelete) | **Delete** /orgs/{org}/invites/{slug_perm}/ | Delete a specific organization invite
[**OrgsInvitesExtend**](OrgsApi.md#OrgsInvitesExtend) | **Post** /orgs/{org}/invites/{slug_perm}/extend/ | Extend an organization invite.
[**OrgsInvitesList**](OrgsApi.md#OrgsInvitesList) | **Get** /orgs/{org}/invites/ | Get a list of all invites for an organization.
[**OrgsInvitesPartialUpdate**](OrgsApi.md#OrgsInvitesPartialUpdate) | **Patch** /orgs/{org}/invites/{slug_perm}/ | Update a specific organization invite.
[**OrgsInvitesResend**](OrgsApi.md#OrgsInvitesResend) | **Post** /orgs/{org}/invites/{slug_perm}/resend/ | Resend an organization invite.
[**OrgsList**](OrgsApi.md#OrgsList) | **Get** /orgs/ | Get a list of all the organizations you are associated with.
[**OrgsMembersDelete**](OrgsApi.md#OrgsMembersDelete) | **Delete** /orgs/{org}/members/{member}/ | Removes a member from the organization.
[**OrgsMembersList**](OrgsApi.md#OrgsMembersList) | **Get** /orgs/{org}/members/ | Get the details for all organization members.
[**OrgsMembersRead**](OrgsApi.md#OrgsMembersRead) | **Get** /orgs/{org}/members/{member}/ | Get the details for a specific organization member.
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsInvitesDelete(context.Background(), org, slugPerm).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    member := "member_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsMembersDelete(context.Background(), org, member).Execute()
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

> []OrganizationMembership OrgsMembersList(ctx, org).Page(page).PageSize(pageSize).Execute()

Get the details for all organization members.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsMembersList(context.Background(), org).Page(page).PageSize(pageSize).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    member := "member_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsMembersRemove(context.Background(), org, member).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    slugPerm := "slugPerm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsSamlGroupSyncDelete(context.Background(), org, slugPerm).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsServicesDelete(context.Background(), org, service).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | 
    team := "team_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsTeamsDelete(context.Background(), org, team).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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

