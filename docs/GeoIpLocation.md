# GeoIpLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **NullableString** |  | 
**Continent** | **NullableString** |  | 
**Country** | **NullableString** |  | 
**CountryCode** | Pointer to **string** |  | [optional] [readonly] 
**Latitude** | Pointer to **NullableFloat64** |  | [optional] 
**Longitude** | Pointer to **NullableFloat64** |  | [optional] 
**PostalCode** | **NullableString** |  | 

## Methods

### NewGeoIpLocation

`func NewGeoIpLocation(city NullableString, continent NullableString, country NullableString, postalCode NullableString, ) *GeoIpLocation`

NewGeoIpLocation instantiates a new GeoIpLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoIpLocationWithDefaults

`func NewGeoIpLocationWithDefaults() *GeoIpLocation`

NewGeoIpLocationWithDefaults instantiates a new GeoIpLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *GeoIpLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeoIpLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeoIpLocation) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *GeoIpLocation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GeoIpLocation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetContinent

`func (o *GeoIpLocation) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *GeoIpLocation) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *GeoIpLocation) SetContinent(v string)`

SetContinent sets Continent field to given value.


### SetContinentNil

`func (o *GeoIpLocation) SetContinentNil(b bool)`

 SetContinentNil sets the value for Continent to be an explicit nil

### UnsetContinent
`func (o *GeoIpLocation) UnsetContinent()`

UnsetContinent ensures that no value is present for Continent, not even an explicit nil
### GetCountry

`func (o *GeoIpLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoIpLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoIpLocation) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *GeoIpLocation) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *GeoIpLocation) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryCode

`func (o *GeoIpLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GeoIpLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GeoIpLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GeoIpLocation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLatitude

`func (o *GeoIpLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GeoIpLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GeoIpLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GeoIpLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *GeoIpLocation) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *GeoIpLocation) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *GeoIpLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GeoIpLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GeoIpLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GeoIpLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *GeoIpLocation) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *GeoIpLocation) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetPostalCode

`func (o *GeoIpLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *GeoIpLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *GeoIpLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *GeoIpLocation) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *GeoIpLocation) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


