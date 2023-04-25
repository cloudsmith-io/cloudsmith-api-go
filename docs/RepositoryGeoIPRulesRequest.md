# RepositoryGeoIPRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | [**RepositoryGeoIPCidr**](RepositoryGeoIPCidr.md) |  | 
**CountryCode** | [**RepositoryGeoIPCountryCode**](RepositoryGeoIPCountryCode.md) |  | 

## Methods

### NewRepositoryGeoIPRulesRequest

`func NewRepositoryGeoIPRulesRequest(cidr RepositoryGeoIPCidr, countryCode RepositoryGeoIPCountryCode, ) *RepositoryGeoIPRulesRequest`

NewRepositoryGeoIPRulesRequest instantiates a new RepositoryGeoIPRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIPRulesRequestWithDefaults

`func NewRepositoryGeoIPRulesRequestWithDefaults() *RepositoryGeoIPRulesRequest`

NewRepositoryGeoIPRulesRequestWithDefaults instantiates a new RepositoryGeoIPRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *RepositoryGeoIPRulesRequest) GetCidr() RepositoryGeoIPCidr`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RepositoryGeoIPRulesRequest) GetCidrOk() (*RepositoryGeoIPCidr, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RepositoryGeoIPRulesRequest) SetCidr(v RepositoryGeoIPCidr)`

SetCidr sets Cidr field to given value.


### GetCountryCode

`func (o *RepositoryGeoIPRulesRequest) GetCountryCode() RepositoryGeoIPCountryCode`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RepositoryGeoIPRulesRequest) GetCountryCodeOk() (*RepositoryGeoIPCountryCode, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RepositoryGeoIPRulesRequest) SetCountryCode(v RepositoryGeoIPCountryCode)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


