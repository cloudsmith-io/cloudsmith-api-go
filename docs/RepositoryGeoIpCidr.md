# RepositoryGeoIpCidr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | **[]string** | The allowed CIDRs for this repository | 
**Deny** | **[]string** | The denied CIDRs for this repository | 

## Methods

### NewRepositoryGeoIpCidr

`func NewRepositoryGeoIpCidr(allow []string, deny []string, ) *RepositoryGeoIpCidr`

NewRepositoryGeoIpCidr instantiates a new RepositoryGeoIpCidr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIpCidrWithDefaults

`func NewRepositoryGeoIpCidrWithDefaults() *RepositoryGeoIpCidr`

NewRepositoryGeoIpCidrWithDefaults instantiates a new RepositoryGeoIpCidr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *RepositoryGeoIpCidr) GetAllow() []string`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *RepositoryGeoIpCidr) GetAllowOk() (*[]string, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *RepositoryGeoIpCidr) SetAllow(v []string)`

SetAllow sets Allow field to given value.


### GetDeny

`func (o *RepositoryGeoIpCidr) GetDeny() []string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *RepositoryGeoIpCidr) GetDenyOk() (*[]string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *RepositoryGeoIpCidr) SetDeny(v []string)`

SetDeny sets Deny field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


