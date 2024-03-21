/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.372.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the ServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceRequest{}

// ServiceRequest struct for ServiceRequest
type ServiceRequest struct {
	// The description of the service
	Description *string `json:"description,omitempty"`
	// The name of the service
	Name string `json:"name"`
	// The role of the service.
	Role  *string        `json:"role,omitempty"`
	Teams []ServiceTeams `json:"teams,omitempty"`
}

// NewServiceRequest instantiates a new ServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRequest(name string) *ServiceRequest {
	this := ServiceRequest{}
	this.Name = name
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewServiceRequestWithDefaults instantiates a new ServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRequestWithDefaults() *ServiceRequest {
	this := ServiceRequest{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *ServiceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceRequest) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ServiceRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ServiceRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ServiceRequest) SetRole(v string) {
	o.Role = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ServiceRequest) GetTeams() []ServiceTeams {
	if o == nil || IsNil(o.Teams) {
		var ret []ServiceTeams
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetTeamsOk() ([]ServiceTeams, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ServiceRequest) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []ServiceTeams and assigns it to the Teams field.
func (o *ServiceRequest) SetTeams(v []ServiceTeams) {
	o.Teams = v
}

func (o ServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return toSerialize, nil
}

type NullableServiceRequest struct {
	value *ServiceRequest
	isSet bool
}

func (v NullableServiceRequest) Get() *ServiceRequest {
	return v.value
}

func (v *NullableServiceRequest) Set(val *ServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRequest(val *ServiceRequest) *NullableServiceRequest {
	return &NullableServiceRequest{value: val, isSet: true}
}

func (v NullableServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
