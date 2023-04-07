/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.236.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// RepositoryAuditLog struct for RepositoryAuditLog
type RepositoryAuditLog struct {
	Actor          NullableString `json:"actor"`
	ActorIpAddress NullableString `json:"actor_ip_address"`
	ActorKind      *string        `json:"actor_kind,omitempty"`
	ActorLocation  GeoIpLocation  `json:"actor_location"`
	ActorSlugPerm  NullableString `json:"actor_slug_perm"`
	ActorUrl       NullableString `json:"actor_url,omitempty"`
	Context        string         `json:"context"`
	Event          string         `json:"event"`
	EventAt        time.Time      `json:"event_at"`
	Object         string         `json:"object"`
	ObjectKind     string         `json:"object_kind"`
	ObjectSlugPerm string         `json:"object_slug_perm"`
	Uuid           *string        `json:"uuid,omitempty"`
}

// NewRepositoryAuditLog instantiates a new RepositoryAuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryAuditLog(actor NullableString, actorIpAddress NullableString, actorLocation GeoIpLocation, actorSlugPerm NullableString, context string, event string, eventAt time.Time, object string, objectKind string, objectSlugPerm string) *RepositoryAuditLog {
	this := RepositoryAuditLog{}
	this.Actor = actor
	this.ActorIpAddress = actorIpAddress
	this.ActorLocation = actorLocation
	this.ActorSlugPerm = actorSlugPerm
	this.Context = context
	this.Event = event
	this.EventAt = eventAt
	this.Object = object
	this.ObjectKind = objectKind
	this.ObjectSlugPerm = objectSlugPerm
	return &this
}

// NewRepositoryAuditLogWithDefaults instantiates a new RepositoryAuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryAuditLogWithDefaults() *RepositoryAuditLog {
	this := RepositoryAuditLog{}
	return &this
}

// GetActor returns the Actor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryAuditLog) GetActor() string {
	if o == nil || o.Actor.Get() == nil {
		var ret string
		return ret
	}

	return *o.Actor.Get()
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryAuditLog) GetActorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actor.Get(), o.Actor.IsSet()
}

// SetActor sets field value
func (o *RepositoryAuditLog) SetActor(v string) {
	o.Actor.Set(&v)
}

// GetActorIpAddress returns the ActorIpAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryAuditLog) GetActorIpAddress() string {
	if o == nil || o.ActorIpAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActorIpAddress.Get()
}

// GetActorIpAddressOk returns a tuple with the ActorIpAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryAuditLog) GetActorIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorIpAddress.Get(), o.ActorIpAddress.IsSet()
}

// SetActorIpAddress sets field value
func (o *RepositoryAuditLog) SetActorIpAddress(v string) {
	o.ActorIpAddress.Set(&v)
}

// GetActorKind returns the ActorKind field value if set, zero value otherwise.
func (o *RepositoryAuditLog) GetActorKind() string {
	if o == nil || isNil(o.ActorKind) {
		var ret string
		return ret
	}
	return *o.ActorKind
}

// GetActorKindOk returns a tuple with the ActorKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetActorKindOk() (*string, bool) {
	if o == nil || isNil(o.ActorKind) {
		return nil, false
	}
	return o.ActorKind, true
}

// HasActorKind returns a boolean if a field has been set.
func (o *RepositoryAuditLog) HasActorKind() bool {
	if o != nil && !isNil(o.ActorKind) {
		return true
	}

	return false
}

// SetActorKind gets a reference to the given string and assigns it to the ActorKind field.
func (o *RepositoryAuditLog) SetActorKind(v string) {
	o.ActorKind = &v
}

// GetActorLocation returns the ActorLocation field value
func (o *RepositoryAuditLog) GetActorLocation() GeoIpLocation {
	if o == nil {
		var ret GeoIpLocation
		return ret
	}

	return o.ActorLocation
}

// GetActorLocationOk returns a tuple with the ActorLocation field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetActorLocationOk() (*GeoIpLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorLocation, true
}

// SetActorLocation sets field value
func (o *RepositoryAuditLog) SetActorLocation(v GeoIpLocation) {
	o.ActorLocation = v
}

// GetActorSlugPerm returns the ActorSlugPerm field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryAuditLog) GetActorSlugPerm() string {
	if o == nil || o.ActorSlugPerm.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActorSlugPerm.Get()
}

// GetActorSlugPermOk returns a tuple with the ActorSlugPerm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryAuditLog) GetActorSlugPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorSlugPerm.Get(), o.ActorSlugPerm.IsSet()
}

// SetActorSlugPerm sets field value
func (o *RepositoryAuditLog) SetActorSlugPerm(v string) {
	o.ActorSlugPerm.Set(&v)
}

// GetActorUrl returns the ActorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryAuditLog) GetActorUrl() string {
	if o == nil || isNil(o.ActorUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ActorUrl.Get()
}

// GetActorUrlOk returns a tuple with the ActorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryAuditLog) GetActorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorUrl.Get(), o.ActorUrl.IsSet()
}

// HasActorUrl returns a boolean if a field has been set.
func (o *RepositoryAuditLog) HasActorUrl() bool {
	if o != nil && o.ActorUrl.IsSet() {
		return true
	}

	return false
}

// SetActorUrl gets a reference to the given NullableString and assigns it to the ActorUrl field.
func (o *RepositoryAuditLog) SetActorUrl(v string) {
	o.ActorUrl.Set(&v)
}

// SetActorUrlNil sets the value for ActorUrl to be an explicit nil
func (o *RepositoryAuditLog) SetActorUrlNil() {
	o.ActorUrl.Set(nil)
}

// UnsetActorUrl ensures that no value is present for ActorUrl, not even an explicit nil
func (o *RepositoryAuditLog) UnsetActorUrl() {
	o.ActorUrl.Unset()
}

// GetContext returns the Context field value
func (o *RepositoryAuditLog) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *RepositoryAuditLog) SetContext(v string) {
	o.Context = v
}

// GetEvent returns the Event field value
func (o *RepositoryAuditLog) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *RepositoryAuditLog) SetEvent(v string) {
	o.Event = v
}

// GetEventAt returns the EventAt field value
func (o *RepositoryAuditLog) GetEventAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventAt
}

// GetEventAtOk returns a tuple with the EventAt field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetEventAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventAt, true
}

// SetEventAt sets field value
func (o *RepositoryAuditLog) SetEventAt(v time.Time) {
	o.EventAt = v
}

// GetObject returns the Object field value
func (o *RepositoryAuditLog) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RepositoryAuditLog) SetObject(v string) {
	o.Object = v
}

// GetObjectKind returns the ObjectKind field value
func (o *RepositoryAuditLog) GetObjectKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectKind
}

// GetObjectKindOk returns a tuple with the ObjectKind field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetObjectKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectKind, true
}

// SetObjectKind sets field value
func (o *RepositoryAuditLog) SetObjectKind(v string) {
	o.ObjectKind = v
}

// GetObjectSlugPerm returns the ObjectSlugPerm field value
func (o *RepositoryAuditLog) GetObjectSlugPerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectSlugPerm
}

// GetObjectSlugPermOk returns a tuple with the ObjectSlugPerm field value
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetObjectSlugPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectSlugPerm, true
}

// SetObjectSlugPerm sets field value
func (o *RepositoryAuditLog) SetObjectSlugPerm(v string) {
	o.ObjectSlugPerm = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RepositoryAuditLog) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAuditLog) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RepositoryAuditLog) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RepositoryAuditLog) SetUuid(v string) {
	o.Uuid = &v
}

func (o RepositoryAuditLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actor"] = o.Actor.Get()
	}
	if true {
		toSerialize["actor_ip_address"] = o.ActorIpAddress.Get()
	}
	if !isNil(o.ActorKind) {
		toSerialize["actor_kind"] = o.ActorKind
	}
	if true {
		toSerialize["actor_location"] = o.ActorLocation
	}
	if true {
		toSerialize["actor_slug_perm"] = o.ActorSlugPerm.Get()
	}
	if o.ActorUrl.IsSet() {
		toSerialize["actor_url"] = o.ActorUrl.Get()
	}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["event_at"] = o.EventAt
	}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["object_kind"] = o.ObjectKind
	}
	if true {
		toSerialize["object_slug_perm"] = o.ObjectSlugPerm
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryAuditLog struct {
	value *RepositoryAuditLog
	isSet bool
}

func (v NullableRepositoryAuditLog) Get() *RepositoryAuditLog {
	return v.value
}

func (v *NullableRepositoryAuditLog) Set(val *RepositoryAuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryAuditLog(val *RepositoryAuditLog) *NullableRepositoryAuditLog {
	return &NullableRepositoryAuditLog{value: val, isSet: true}
}

func (v NullableRepositoryAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
