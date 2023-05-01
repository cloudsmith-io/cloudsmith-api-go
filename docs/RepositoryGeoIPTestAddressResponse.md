# RepositoryGeoIpTestAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]RepositoryGeoIpTestAddressResponseDict**](RepositoryGeoIpTestAddressResponseDict.md) | The IP address test results ordered by allowed | 

## Methods

### NewRepositoryGeoIpTestAddressResponse

`func NewRepositoryGeoIpTestAddressResponse(addresses []RepositoryGeoIpTestAddressResponseDict, ) *RepositoryGeoIpTestAddressResponse`

NewRepositoryGeoIpTestAddressResponse instantiates a new RepositoryGeoIpTestAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIpTestAddressResponseWithDefaults

`func NewRepositoryGeoIpTestAddressResponseWithDefaults() *RepositoryGeoIpTestAddressResponse`

NewRepositoryGeoIpTestAddressResponseWithDefaults instantiates a new RepositoryGeoIpTestAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *RepositoryGeoIpTestAddressResponse) GetAddresses() []RepositoryGeoIpTestAddressResponseDict`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RepositoryGeoIpTestAddressResponse) GetAddressesOk() (*[]RepositoryGeoIpTestAddressResponseDict, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RepositoryGeoIpTestAddressResponse) SetAddresses(v []RepositoryGeoIpTestAddressResponseDict)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


