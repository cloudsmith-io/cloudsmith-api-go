# RepositoryGeoIpRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to [**RepositoryGeoIpCountryCodeRules**](RepositoryGeoIpCountryCodeRules.md) |  | [optional] 
**Cidr** | Pointer to [**RepositoryGeoIpCidrRules**](RepositoryGeoIpCidrRules.md) |  | [optional] 

## Methods

### NewRepositoryGeoIpRules

`func NewRepositoryGeoIpRules() *RepositoryGeoIpRules`

NewRepositoryGeoIpRules instantiates a new RepositoryGeoIpRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIpRulesWithDefaults

`func NewRepositoryGeoIpRulesWithDefaults() *RepositoryGeoIpRules`

NewRepositoryGeoIpRulesWithDefaults instantiates a new RepositoryGeoIpRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *RepositoryGeoIpRules) GetCountryCode() RepositoryGeoIpCountryCodeRules`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RepositoryGeoIpRules) GetCountryCodeOk() (*RepositoryGeoIpCountryCodeRules, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RepositoryGeoIpRules) SetCountryCode(v RepositoryGeoIpCountryCodeRules)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *RepositoryGeoIpRules) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCidr

`func (o *RepositoryGeoIpRules) GetCidr() RepositoryGeoIpCidrRules`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RepositoryGeoIpRules) GetCidrOk() (*RepositoryGeoIpCidrRules, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RepositoryGeoIpRules) SetCidr(v RepositoryGeoIpCidrRules)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *RepositoryGeoIpRules) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


