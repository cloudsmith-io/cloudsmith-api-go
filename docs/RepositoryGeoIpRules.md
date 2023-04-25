# RepositoryGeoIPRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | [**RepositoryGeoIPCidr**](RepositoryGeoIPCidr.md) |  | 
**CountryCode** | [**RepositoryGeoIPCountryCode**](RepositoryGeoIPCountryCode.md) |  | 

## Methods

### NewRepositoryGeoIPRules

`func NewRepositoryGeoIPRules(cidr RepositoryGeoIPCidr, countryCode RepositoryGeoIPCountryCode, ) *RepositoryGeoIPRules`

NewRepositoryGeoIPRules instantiates a new RepositoryGeoIPRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIPRulesWithDefaults

`func NewRepositoryGeoIPRulesWithDefaults() *RepositoryGeoIPRules`

NewRepositoryGeoIPRulesWithDefaults instantiates a new RepositoryGeoIPRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *RepositoryGeoIPRules) GetCidr() RepositoryGeoIPCidr`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RepositoryGeoIPRules) GetCidrOk() (*RepositoryGeoIPCidr, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RepositoryGeoIPRules) SetCidr(v RepositoryGeoIPCidr)`

SetCidr sets Cidr field to given value.


### GetCountryCode

`func (o *RepositoryGeoIPRules) GetCountryCode() RepositoryGeoIPCountryCode`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RepositoryGeoIPRules) GetCountryCodeOk() (*RepositoryGeoIPCountryCode, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RepositoryGeoIPRules) SetCountryCode(v RepositoryGeoIPCountryCode)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


