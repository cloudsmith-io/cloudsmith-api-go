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
	"time"
)

// checks if the NamespaceAuditLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceAuditLog{}

// NamespaceAuditLog struct for NamespaceAuditLog
type NamespaceAuditLog struct {
	Actor                NullableString `json:"actor"`
	ActorIpAddress       NullableString `json:"actor_ip_address"`
	ActorKind            *string        `json:"actor_kind,omitempty"`
	ActorLocation        GeoIpLocation  `json:"actor_location"`
	ActorSlugPerm        NullableString `json:"actor_slug_perm"`
	ActorUrl             NullableString `json:"actor_url,omitempty"`
	Context              string         `json:"context"`
	Event                string         `json:"event"`
	EventAt              time.Time      `json:"event_at"`
	Object               string         `json:"object"`
	ObjectKind           string         `json:"object_kind"`
	ObjectSlugPerm       string         `json:"object_slug_perm"`
	Target               string         `json:"target"`
	TargetKind           string         `json:"target_kind"`
	TargetSlugPerm       NullableString `json:"target_slug_perm,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Uuid                 *string        `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NamespaceAuditLog NamespaceAuditLog

// NewNamespaceAuditLog instantiates a new NamespaceAuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceAuditLog(actor NullableString, actorIpAddress NullableString, actorLocation GeoIpLocation, actorSlugPerm NullableString, context string, event string, eventAt time.Time, object string, objectKind string, objectSlugPerm string, target string, targetKind string) *NamespaceAuditLog {
	this := NamespaceAuditLog{}
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
	this.Target = target
	this.TargetKind = targetKind
	return &this
}

// NewNamespaceAuditLogWithDefaults instantiates a new NamespaceAuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceAuditLogWithDefaults() *NamespaceAuditLog {
	this := NamespaceAuditLog{}
	return &this
}

// GetActor returns the Actor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NamespaceAuditLog) GetActor() string {
	if o == nil || o.Actor.Get() == nil {
		var ret string
		return ret
	}

	return *o.Actor.Get()
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamespaceAuditLog) GetActorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actor.Get(), o.Actor.IsSet()
}

// SetActor sets field value
func (o *NamespaceAuditLog) SetActor(v string) {
	o.Actor.Set(&v)
}

// GetActorIpAddress returns the ActorIpAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NamespaceAuditLog) GetActorIpAddress() string {
	if o == nil || o.ActorIpAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActorIpAddress.Get()
}

// GetActorIpAddressOk returns a tuple with the ActorIpAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamespaceAuditLog) GetActorIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorIpAddress.Get(), o.ActorIpAddress.IsSet()
}

// SetActorIpAddress sets field value
func (o *NamespaceAuditLog) SetActorIpAddress(v string) {
	o.ActorIpAddress.Set(&v)
}

// GetActorKind returns the ActorKind field value if set, zero value otherwise.
func (o *NamespaceAuditLog) GetActorKind() string {
	if o == nil || IsNil(o.ActorKind) {
		var ret string
		return ret
	}
	return *o.ActorKind
}

// GetActorKindOk returns a tuple with the ActorKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetActorKindOk() (*string, bool) {
	if o == nil || IsNil(o.ActorKind) {
		return nil, false
	}
	return o.ActorKind, true
}

// HasActorKind returns a boolean if a field has been set.
func (o *NamespaceAuditLog) HasActorKind() bool {
	if o != nil && !IsNil(o.ActorKind) {
		return true
	}

	return false
}

// SetActorKind gets a reference to the given string and assigns it to the ActorKind field.
func (o *NamespaceAuditLog) SetActorKind(v string) {
	o.ActorKind = &v
}

// GetActorLocation returns the ActorLocation field value
func (o *NamespaceAuditLog) GetActorLocation() GeoIpLocation {
	if o == nil {
		var ret GeoIpLocation
		return ret
	}

	return o.ActorLocation
}

// GetActorLocationOk returns a tuple with the ActorLocation field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetActorLocationOk() (*GeoIpLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorLocation, true
}

// SetActorLocation sets field value
func (o *NamespaceAuditLog) SetActorLocation(v GeoIpLocation) {
	o.ActorLocation = v
}

// GetActorSlugPerm returns the ActorSlugPerm field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NamespaceAuditLog) GetActorSlugPerm() string {
	if o == nil || o.ActorSlugPerm.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActorSlugPerm.Get()
}

// GetActorSlugPermOk returns a tuple with the ActorSlugPerm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamespaceAuditLog) GetActorSlugPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorSlugPerm.Get(), o.ActorSlugPerm.IsSet()
}

// SetActorSlugPerm sets field value
func (o *NamespaceAuditLog) SetActorSlugPerm(v string) {
	o.ActorSlugPerm.Set(&v)
}

// GetActorUrl returns the ActorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamespaceAuditLog) GetActorUrl() string {
	if o == nil || IsNil(o.ActorUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ActorUrl.Get()
}

// GetActorUrlOk returns a tuple with the ActorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamespaceAuditLog) GetActorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorUrl.Get(), o.ActorUrl.IsSet()
}

// HasActorUrl returns a boolean if a field has been set.
func (o *NamespaceAuditLog) HasActorUrl() bool {
	if o != nil && o.ActorUrl.IsSet() {
		return true
	}

	return false
}

// SetActorUrl gets a reference to the given NullableString and assigns it to the ActorUrl field.
func (o *NamespaceAuditLog) SetActorUrl(v string) {
	o.ActorUrl.Set(&v)
}

// SetActorUrlNil sets the value for ActorUrl to be an explicit nil
func (o *NamespaceAuditLog) SetActorUrlNil() {
	o.ActorUrl.Set(nil)
}

// UnsetActorUrl ensures that no value is present for ActorUrl, not even an explicit nil
func (o *NamespaceAuditLog) UnsetActorUrl() {
	o.ActorUrl.Unset()
}

// GetContext returns the Context field value
func (o *NamespaceAuditLog) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *NamespaceAuditLog) SetContext(v string) {
	o.Context = v
}

// GetEvent returns the Event field value
func (o *NamespaceAuditLog) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NamespaceAuditLog) SetEvent(v string) {
	o.Event = v
}

// GetEventAt returns the EventAt field value
func (o *NamespaceAuditLog) GetEventAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventAt
}

// GetEventAtOk returns a tuple with the EventAt field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetEventAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventAt, true
}

// SetEventAt sets field value
func (o *NamespaceAuditLog) SetEventAt(v time.Time) {
	o.EventAt = v
}

// GetObject returns the Object field value
func (o *NamespaceAuditLog) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *NamespaceAuditLog) SetObject(v string) {
	o.Object = v
}

// GetObjectKind returns the ObjectKind field value
func (o *NamespaceAuditLog) GetObjectKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectKind
}

// GetObjectKindOk returns a tuple with the ObjectKind field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetObjectKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectKind, true
}

// SetObjectKind sets field value
func (o *NamespaceAuditLog) SetObjectKind(v string) {
	o.ObjectKind = v
}

// GetObjectSlugPerm returns the ObjectSlugPerm field value
func (o *NamespaceAuditLog) GetObjectSlugPerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectSlugPerm
}

// GetObjectSlugPermOk returns a tuple with the ObjectSlugPerm field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetObjectSlugPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectSlugPerm, true
}

// SetObjectSlugPerm sets field value
func (o *NamespaceAuditLog) SetObjectSlugPerm(v string) {
	o.ObjectSlugPerm = v
}

// GetTarget returns the Target field value
func (o *NamespaceAuditLog) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *NamespaceAuditLog) SetTarget(v string) {
	o.Target = v
}

// GetTargetKind returns the TargetKind field value
func (o *NamespaceAuditLog) GetTargetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetKind
}

// GetTargetKindOk returns a tuple with the TargetKind field value
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetTargetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetKind, true
}

// SetTargetKind sets field value
func (o *NamespaceAuditLog) SetTargetKind(v string) {
	o.TargetKind = v
}

// GetTargetSlugPerm returns the TargetSlugPerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamespaceAuditLog) GetTargetSlugPerm() string {
	if o == nil || IsNil(o.TargetSlugPerm.Get()) {
		var ret string
		return ret
	}
	return *o.TargetSlugPerm.Get()
}

// GetTargetSlugPermOk returns a tuple with the TargetSlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamespaceAuditLog) GetTargetSlugPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSlugPerm.Get(), o.TargetSlugPerm.IsSet()
}

// HasTargetSlugPerm returns a boolean if a field has been set.
func (o *NamespaceAuditLog) HasTargetSlugPerm() bool {
	if o != nil && o.TargetSlugPerm.IsSet() {
		return true
	}

	return false
}

// SetTargetSlugPerm gets a reference to the given NullableString and assigns it to the TargetSlugPerm field.
func (o *NamespaceAuditLog) SetTargetSlugPerm(v string) {
	o.TargetSlugPerm.Set(&v)
}

// SetTargetSlugPermNil sets the value for TargetSlugPerm to be an explicit nil
func (o *NamespaceAuditLog) SetTargetSlugPermNil() {
	o.TargetSlugPerm.Set(nil)
}

// UnsetTargetSlugPerm ensures that no value is present for TargetSlugPerm, not even an explicit nil
func (o *NamespaceAuditLog) UnsetTargetSlugPerm() {
	o.TargetSlugPerm.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NamespaceAuditLog) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceAuditLog) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NamespaceAuditLog) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NamespaceAuditLog) SetUuid(v string) {
	o.Uuid = &v
}

func (o NamespaceAuditLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceAuditLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actor"] = o.Actor.Get()
	toSerialize["actor_ip_address"] = o.ActorIpAddress.Get()
	if !IsNil(o.ActorKind) {
		toSerialize["actor_kind"] = o.ActorKind
	}
	toSerialize["actor_location"] = o.ActorLocation
	toSerialize["actor_slug_perm"] = o.ActorSlugPerm.Get()
	if o.ActorUrl.IsSet() {
		toSerialize["actor_url"] = o.ActorUrl.Get()
	}
	toSerialize["context"] = o.Context
	toSerialize["event"] = o.Event
	toSerialize["event_at"] = o.EventAt
	toSerialize["object"] = o.Object
	toSerialize["object_kind"] = o.ObjectKind
	toSerialize["object_slug_perm"] = o.ObjectSlugPerm
	toSerialize["target"] = o.Target
	toSerialize["target_kind"] = o.TargetKind
	if o.TargetSlugPerm.IsSet() {
		toSerialize["target_slug_perm"] = o.TargetSlugPerm.Get()
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NamespaceAuditLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actor",
		"actor_ip_address",
		"actor_location",
		"actor_slug_perm",
		"context",
		"event",
		"event_at",
		"object",
		"object_kind",
		"object_slug_perm",
		"target",
		"target_kind",
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

	varNamespaceAuditLog := _NamespaceAuditLog{}

	err = json.Unmarshal(data, &varNamespaceAuditLog)

	if err != nil {
		return err
	}

	*o = NamespaceAuditLog(varNamespaceAuditLog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actor")
		delete(additionalProperties, "actor_ip_address")
		delete(additionalProperties, "actor_kind")
		delete(additionalProperties, "actor_location")
		delete(additionalProperties, "actor_slug_perm")
		delete(additionalProperties, "actor_url")
		delete(additionalProperties, "context")
		delete(additionalProperties, "event")
		delete(additionalProperties, "event_at")
		delete(additionalProperties, "object")
		delete(additionalProperties, "object_kind")
		delete(additionalProperties, "object_slug_perm")
		delete(additionalProperties, "target")
		delete(additionalProperties, "target_kind")
		delete(additionalProperties, "target_slug_perm")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNamespaceAuditLog struct {
	value *NamespaceAuditLog
	isSet bool
}

func (v NullableNamespaceAuditLog) Get() *NamespaceAuditLog {
	return v.value
}

func (v *NullableNamespaceAuditLog) Set(val *NamespaceAuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceAuditLog(val *NamespaceAuditLog) *NullableNamespaceAuditLog {
	return &NullableNamespaceAuditLog{value: val, isSet: true}
}

func (v NullableNamespaceAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
