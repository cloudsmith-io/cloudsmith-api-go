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

// checks if the RepositoryPrivilegeDict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryPrivilegeDict{}

// RepositoryPrivilegeDict struct for RepositoryPrivilegeDict
type RepositoryPrivilegeDict struct {
	// The level of privilege that the user or team should be granted to the specified repository.
	Privilege string `json:"privilege"`
	// The service identifier (slug).
	Service *string `json:"service,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// The team identifier (slug).
	Team *string `json:"team,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// The user identifier (slug).
	User                 *string `json:"user,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryPrivilegeDict RepositoryPrivilegeDict

// NewRepositoryPrivilegeDict instantiates a new RepositoryPrivilegeDict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryPrivilegeDict(privilege string) *RepositoryPrivilegeDict {
	this := RepositoryPrivilegeDict{}
	this.Privilege = privilege
	return &this
}

// NewRepositoryPrivilegeDictWithDefaults instantiates a new RepositoryPrivilegeDict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryPrivilegeDictWithDefaults() *RepositoryPrivilegeDict {
	this := RepositoryPrivilegeDict{}
	return &this
}

// GetPrivilege returns the Privilege field value
func (o *RepositoryPrivilegeDict) GetPrivilege() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeDict) GetPrivilegeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privilege, true
}

// SetPrivilege sets field value
func (o *RepositoryPrivilegeDict) SetPrivilege(v string) {
	o.Privilege = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RepositoryPrivilegeDict) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeDict) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RepositoryPrivilegeDict) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *RepositoryPrivilegeDict) SetService(v string) {
	o.Service = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *RepositoryPrivilegeDict) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeDict) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *RepositoryPrivilegeDict) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *RepositoryPrivilegeDict) SetTeam(v string) {
	o.Team = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RepositoryPrivilegeDict) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeDict) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RepositoryPrivilegeDict) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *RepositoryPrivilegeDict) SetUser(v string) {
	o.User = &v
}

func (o RepositoryPrivilegeDict) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryPrivilegeDict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privilege"] = o.Privilege
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryPrivilegeDict) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privilege",
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

	varRepositoryPrivilegeDict := _RepositoryPrivilegeDict{}

	err = json.Unmarshal(data, &varRepositoryPrivilegeDict)

	if err != nil {
		return err
	}

	*o = RepositoryPrivilegeDict(varRepositoryPrivilegeDict)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "privilege")
		delete(additionalProperties, "service")
		delete(additionalProperties, "team")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryPrivilegeDict struct {
	value *RepositoryPrivilegeDict
	isSet bool
}

func (v NullableRepositoryPrivilegeDict) Get() *RepositoryPrivilegeDict {
	return v.value
}

func (v *NullableRepositoryPrivilegeDict) Set(val *RepositoryPrivilegeDict) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryPrivilegeDict) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryPrivilegeDict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryPrivilegeDict(val *RepositoryPrivilegeDict) *NullableRepositoryPrivilegeDict {
	return &NullableRepositoryPrivilegeDict{value: val, isSet: true}
}

func (v NullableRepositoryPrivilegeDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryPrivilegeDict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
