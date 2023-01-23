# RepositoryGeoIPTestAddressResponseDict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | The result of the IP test | 
**CountryCode** | **NullableString** | The country code of the tested IP address | 
**IpAddress** | **string** | The IP address that was tested | 
**Reason** | **string** | The reason for the result | 

## Methods

### NewRepositoryGeoIPTestAddressResponseDict

`func NewRepositoryGeoIPTestAddressResponseDict(allowed bool, countryCode NullableString, ipAddress string, reason string, ) *RepositoryGeoIPTestAddressResponseDict`

NewRepositoryGeoIPTestAddressResponseDict instantiates a new RepositoryGeoIPTestAddressResponseDict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIPTestAddressResponseDictWithDefaults

`func NewRepositoryGeoIPTestAddressResponseDictWithDefaults() *RepositoryGeoIPTestAddressResponseDict`

NewRepositoryGeoIPTestAddressResponseDictWithDefaults instantiates a new RepositoryGeoIPTestAddressResponseDict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *RepositoryGeoIPTestAddressResponseDict) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *RepositoryGeoIPTestAddressResponseDict) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *RepositoryGeoIPTestAddressResponseDict) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetCountryCode

`func (o *RepositoryGeoIPTestAddressResponseDict) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RepositoryGeoIPTestAddressResponseDict) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RepositoryGeoIPTestAddressResponseDict) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *RepositoryGeoIPTestAddressResponseDict) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *RepositoryGeoIPTestAddressResponseDict) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetIpAddress

`func (o *RepositoryGeoIPTestAddressResponseDict) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RepositoryGeoIPTestAddressResponseDict) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RepositoryGeoIPTestAddressResponseDict) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetReason

`func (o *RepositoryGeoIPTestAddressResponseDict) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RepositoryGeoIPTestAddressResponseDict) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RepositoryGeoIPTestAddressResponseDict) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


