/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.498.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	CreatedBy    *string    `json:"created_by,omitempty"`
	CreatedByUrl *string    `json:"created_by_url,omitempty"`
	// The description of the service
	Description *string `json:"description,omitempty"`
	// The API key of the service
	Key *string `json:"key,omitempty"`
	// The name of the service
	Name string `json:"name"`
	// The role of the service.
	Role *string `json:"role,omitempty"`
	// The slug of the service
	Slug  *string        `json:"slug,omitempty"`
	Teams []ServiceTeams `json:"teams,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(name string) *Service {
	this := Service{}
	this.Name = name
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Service) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Service) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Service) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Service) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Service) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Service) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUrl returns the CreatedByUrl field value if set, zero value otherwise.
func (o *Service) GetCreatedByUrl() string {
	if o == nil || IsNil(o.CreatedByUrl) {
		var ret string
		return ret
	}
	return *o.CreatedByUrl
}

// GetCreatedByUrlOk returns a tuple with the CreatedByUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedByUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUrl) {
		return nil, false
	}
	return o.CreatedByUrl, true
}

// HasCreatedByUrl returns a boolean if a field has been set.
func (o *Service) HasCreatedByUrl() bool {
	if o != nil && !IsNil(o.CreatedByUrl) {
		return true
	}

	return false
}

// SetCreatedByUrl gets a reference to the given string and assigns it to the CreatedByUrl field.
func (o *Service) SetCreatedByUrl(v string) {
	o.CreatedByUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Service) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Service) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Service) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Service) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Service) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Service) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Service) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Service) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Service) SetRole(v string) {
	o.Role = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Service) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Service) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Service) SetSlug(v string) {
	o.Slug = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *Service) GetTeams() []ServiceTeams {
	if o == nil || IsNil(o.Teams) {
		var ret []ServiceTeams
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTeamsOk() ([]ServiceTeams, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *Service) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []ServiceTeams and assigns it to the Teams field.
func (o *Service) SetTeams(v []ServiceTeams) {
	o.Teams = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CreatedByUrl) {
		toSerialize["created_by_url"] = o.CreatedByUrl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return toSerialize, nil
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
