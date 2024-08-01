# RepositoryGeoIpStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoipEnabled** | Pointer to **bool** | If checked, any access to the website or downloads for this repository is allowed/denied according to the configured Geo/IP restriction rules. | [optional] [readonly] 

## Methods

### NewRepositoryGeoIpStatus

`func NewRepositoryGeoIpStatus() *RepositoryGeoIpStatus`

NewRepositoryGeoIpStatus instantiates a new RepositoryGeoIpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGeoIpStatusWithDefaults

`func NewRepositoryGeoIpStatusWithDefaults() *RepositoryGeoIpStatus`

NewRepositoryGeoIpStatusWithDefaults instantiates a new RepositoryGeoIpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoipEnabled

`func (o *RepositoryGeoIpStatus) GetGeoipEnabled() bool`

GetGeoipEnabled returns the GeoipEnabled field if non-nil, zero value otherwise.

### GetGeoipEnabledOk

`func (o *RepositoryGeoIpStatus) GetGeoipEnabledOk() (*bool, bool)`

GetGeoipEnabledOk returns a tuple with the GeoipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoipEnabled

`func (o *RepositoryGeoIpStatus) SetGeoipEnabled(v bool)`

SetGeoipEnabled sets GeoipEnabled field to given value.

### HasGeoipEnabled

`func (o *RepositoryGeoIpStatus) HasGeoipEnabled() bool`

HasGeoipEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


