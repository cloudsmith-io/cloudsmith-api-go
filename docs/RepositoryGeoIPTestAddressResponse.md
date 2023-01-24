# RepositoryGeoIPTestAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]RepositoryGeoIPTestAddressResponseDict**](RepositoryGeoIPTestAddressResponseDict.md) | The IP address test results ordered by allowed | 

## Methods

### NewRepositoryGeoIPTestAddressResponse

`func NewRepositoryGeoIPTestAddressResponse(addresses []RepositoryGeoIPTestAddressResponseDict, ) *RepositoryGeoIPTestAddressResponse`

NewRepositoryGeoIPTestAddressResponse instantiates a new RepositoryGeoIPTestAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIPTestAddressResponseWithDefaults

`func NewRepositoryGeoIPTestAddressResponseWithDefaults() *RepositoryGeoIPTestAddressResponse`

NewRepositoryGeoIPTestAddressResponseWithDefaults instantiates a new RepositoryGeoIPTestAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *RepositoryGeoIPTestAddressResponse) GetAddresses() []RepositoryGeoIPTestAddressResponseDict`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RepositoryGeoIPTestAddressResponse) GetAddressesOk() (*[]RepositoryGeoIPTestAddressResponseDict, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RepositoryGeoIPTestAddressResponse) SetAddresses(v []RepositoryGeoIPTestAddressResponseDict)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


