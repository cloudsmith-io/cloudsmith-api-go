# RepositoryGeoIPList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | List all CIDR geographic rules within the repository | [optional] [readonly] 
**CountryCode** | Pointer to **string** | List all GeoIP geographic rules within the repository | [optional] [readonly] 

## Methods

### NewRepositoryGeoIPList

`func NewRepositoryGeoIPList() *RepositoryGeoIPList`

NewRepositoryGeoIPList instantiates a new RepositoryGeoIPList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIPListWithDefaults

`func NewRepositoryGeoIPListWithDefaults() *RepositoryGeoIPList`

NewRepositoryGeoIPListWithDefaults instantiates a new RepositoryGeoIPList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *RepositoryGeoIPList) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RepositoryGeoIPList) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RepositoryGeoIPList) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *RepositoryGeoIPList) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCountryCode

`func (o *RepositoryGeoIPList) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RepositoryGeoIPList) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RepositoryGeoIPList) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *RepositoryGeoIPList) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


