# OidcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidcToken** | **string** | Serialized OIDC token. | 
**ServiceSlug** | **string** | Slug for the service account. | 

## Methods

### NewOidcRequest

`func NewOidcRequest(oidcToken string, serviceSlug string, ) *OidcRequest`

NewOidcRequest instantiates a new OidcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcRequestWithDefaults

`func NewOidcRequestWithDefaults() *OidcRequest`

NewOidcRequestWithDefaults instantiates a new OidcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidcToken

`func (o *OidcRequest) GetOidcToken() string`

GetOidcToken returns the OidcToken field if non-nil, zero value otherwise.

### GetOidcTokenOk

`func (o *OidcRequest) GetOidcTokenOk() (*string, bool)`

GetOidcTokenOk returns a tuple with the OidcToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcToken

`func (o *OidcRequest) SetOidcToken(v string)`

SetOidcToken sets OidcToken field to given value.


### GetServiceSlug

`func (o *OidcRequest) GetServiceSlug() string`

GetServiceSlug returns the ServiceSlug field if non-nil, zero value otherwise.

### GetServiceSlugOk

`func (o *OidcRequest) GetServiceSlugOk() (*string, bool)`

GetServiceSlugOk returns a tuple with the ServiceSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSlug

`func (o *OidcRequest) SetServiceSlug(v string)`

SetServiceSlug sets ServiceSlug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


