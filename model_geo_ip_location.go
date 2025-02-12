/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the GeoIpLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoIpLocation{}

// GeoIpLocation struct for GeoIpLocation
type GeoIpLocation struct {
	City                 NullableString  `json:"city"`
	Continent            NullableString  `json:"continent"`
	Country              NullableString  `json:"country"`
	CountryCode          *string         `json:"country_code,omitempty"`
	Latitude             NullableFloat64 `json:"latitude,omitempty"`
	Longitude            NullableFloat64 `json:"longitude,omitempty"`
	PostalCode           NullableString  `json:"postal_code"`
	AdditionalProperties map[string]interface{}
}

type _GeoIpLocation GeoIpLocation

// NewGeoIpLocation instantiates a new GeoIpLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoIpLocation(city NullableString, continent NullableString, country NullableString, postalCode NullableString) *GeoIpLocation {
	this := GeoIpLocation{}
	this.City = city
	this.Continent = continent
	this.Country = country
	this.PostalCode = postalCode
	return &this
}

// NewGeoIpLocationWithDefaults instantiates a new GeoIpLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoIpLocationWithDefaults() *GeoIpLocation {
	this := GeoIpLocation{}
	return &this
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GeoIpLocation) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoIpLocation) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *GeoIpLocation) SetCity(v string) {
	o.City.Set(&v)
}

// GetContinent returns the Continent field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GeoIpLocation) GetContinent() string {
	if o == nil || o.Continent.Get() == nil {
		var ret string
		return ret
	}

	return *o.Continent.Get()
}

// GetContinentOk returns a tuple with the Continent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoIpLocation) GetContinentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Continent.Get(), o.Continent.IsSet()
}

// SetContinent sets field value
func (o *GeoIpLocation) SetContinent(v string) {
	o.Continent.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GeoIpLocation) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoIpLocation) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *GeoIpLocation) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GeoIpLocation) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIpLocation) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GeoIpLocation) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GeoIpLocation) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeoIpLocation) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoIpLocation) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *GeoIpLocation) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *GeoIpLocation) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}

// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *GeoIpLocation) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *GeoIpLocation) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeoIpLocation) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoIpLocation) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *GeoIpLocation) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *GeoIpLocation) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}

// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *GeoIpLocation) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *GeoIpLocation) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GeoIpLocation) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoIpLocation) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *GeoIpLocation) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

func (o GeoIpLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoIpLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City.Get()
	toSerialize["continent"] = o.Continent.Get()
	toSerialize["country"] = o.Country.Get()
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	toSerialize["postal_code"] = o.PostalCode.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GeoIpLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"continent",
		"country",
		"postal_code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGeoIpLocation := _GeoIpLocation{}

	err = json.Unmarshal(data, &varGeoIpLocation)

	if err != nil {
		return err
	}

	*o = GeoIpLocation(varGeoIpLocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "continent")
		delete(additionalProperties, "country")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "postal_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGeoIpLocation struct {
	value *GeoIpLocation
	isSet bool
}

func (v NullableGeoIpLocation) Get() *GeoIpLocation {
	return v.value
}

func (v *NullableGeoIpLocation) Set(val *GeoIpLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoIpLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoIpLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoIpLocation(val *GeoIpLocation) *NullableGeoIpLocation {
	return &NullableGeoIpLocation{value: val, isSet: true}
}

func (v NullableGeoIpLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoIpLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
