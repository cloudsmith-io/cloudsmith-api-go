# RepositoryGeoIpRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | [**RepositoryGeoIpCidr**](RepositoryGeoIpCidr.md) |  | 
**CountryCode** | [**RepositoryGeoIpCountryCode**](RepositoryGeoIpCountryCode.md) |  | 

## Methods

### NewRepositoryGeoIpRulesRequest

`func NewRepositoryGeoIpRulesRequest(cidr RepositoryGeoIpCidr, countryCode RepositoryGeoIpCountryCode, ) *RepositoryGeoIpRulesRequest`

NewRepositoryGeoIpRulesRequest instantiates a new RepositoryGeoIpRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIpRulesRequestWithDefaults

`func NewRepositoryGeoIpRulesRequestWithDefaults() *RepositoryGeoIpRulesRequest`

NewRepositoryGeoIpRulesRequestWithDefaults instantiates a new RepositoryGeoIpRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *RepositoryGeoIpRulesRequest) GetCidr() RepositoryGeoIpCidr`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RepositoryGeoIpRulesRequest) GetCidrOk() (*RepositoryGeoIpCidr, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RepositoryGeoIpRulesRequest) SetCidr(v RepositoryGeoIpCidr)`

SetCidr sets Cidr field to given value.


### GetCountryCode

`func (o *RepositoryGeoIpRulesRequest) GetCountryCode() RepositoryGeoIpCountryCode`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RepositoryGeoIpRulesRequest) GetCountryCodeOk() (*RepositoryGeoIpCountryCode, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RepositoryGeoIpRulesRequest) SetCountryCode(v RepositoryGeoIpCountryCode)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


