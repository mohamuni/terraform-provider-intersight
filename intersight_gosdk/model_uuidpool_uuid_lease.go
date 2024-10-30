/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the UuidpoolUuidLease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UuidpoolUuidLease{}

// UuidpoolUuidLease UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
type UuidpoolUuidLease struct {
	PoolAbstractLease
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                        `json:"ObjectType"`
	Reservation *UuidpoolReservationReference `json:"Reservation,omitempty"`
	// UUID Prefix+Suffix numbers.
	Uuid                 *string                                `json:"Uuid,omitempty" validate:"regexp=^$|^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"`
	AssignedToEntity     NullableMoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
	Pool                 NullableUuidpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           NullableUuidpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             NullableUuidpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolUuidLease UuidpoolUuidLease

// NewUuidpoolUuidLease instantiates a new UuidpoolUuidLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolUuidLease(classId string, objectType string) *UuidpoolUuidLease {
	this := UuidpoolUuidLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	var hasDuplicate bool = false
	this.HasDuplicate = &hasDuplicate
	return &this
}

// NewUuidpoolUuidLeaseWithDefaults instantiates a new UuidpoolUuidLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolUuidLeaseWithDefaults() *UuidpoolUuidLease {
	this := UuidpoolUuidLease{}
	var classId string = "uuidpool.UuidLease"
	this.ClassId = classId
	var objectType string = "uuidpool.UuidLease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolUuidLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolUuidLease) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "uuidpool.UuidLease" of the ClassId field.
func (o *UuidpoolUuidLease) GetDefaultClassId() interface{} {
	return "uuidpool.UuidLease"
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolUuidLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolUuidLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "uuidpool.UuidLease" of the ObjectType field.
func (o *UuidpoolUuidLease) GetDefaultObjectType() interface{} {
	return "uuidpool.UuidLease"
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetReservation() UuidpoolReservationReference {
	if o == nil || IsNil(o.Reservation) {
		var ret UuidpoolReservationReference
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetReservationOk() (*UuidpoolReservationReference, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given UuidpoolReservationReference and assigns it to the Reservation field.
func (o *UuidpoolUuidLease) SetReservation(v UuidpoolReservationReference) {
	o.Reservation = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UuidpoolUuidLease) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolUuidLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolUuidLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *UuidpoolUuidLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *UuidpoolUuidLease) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *UuidpoolUuidLease) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolUuidLease) GetPool() UuidpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolUuidLease) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableUuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolUuidLease) SetPool(v UuidpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *UuidpoolUuidLease) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *UuidpoolUuidLease) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolUuidLease) GetPoolMember() UuidpoolPoolMemberRelationship {
	if o == nil || IsNil(o.PoolMember.Get()) {
		var ret UuidpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember.Get()
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolUuidLease) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolMember.Get(), o.PoolMember.IsSet()
}

// HasPoolMember returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasPoolMember() bool {
	if o != nil && o.PoolMember.IsSet() {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given NullableUuidpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *UuidpoolUuidLease) SetPoolMember(v UuidpoolPoolMemberRelationship) {
	o.PoolMember.Set(&v)
}

// SetPoolMemberNil sets the value for PoolMember to be an explicit nil
func (o *UuidpoolUuidLease) SetPoolMemberNil() {
	o.PoolMember.Set(nil)
}

// UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
func (o *UuidpoolUuidLease) UnsetPoolMember() {
	o.PoolMember.Unset()
}

// GetUniverse returns the Universe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolUuidLease) GetUniverse() UuidpoolUniverseRelationship {
	if o == nil || IsNil(o.Universe.Get()) {
		var ret UuidpoolUniverseRelationship
		return ret
	}
	return *o.Universe.Get()
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolUuidLease) GetUniverseOk() (*UuidpoolUniverseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Universe.Get(), o.Universe.IsSet()
}

// HasUniverse returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasUniverse() bool {
	if o != nil && o.Universe.IsSet() {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given NullableUuidpoolUniverseRelationship and assigns it to the Universe field.
func (o *UuidpoolUuidLease) SetUniverse(v UuidpoolUniverseRelationship) {
	o.Universe.Set(&v)
}

// SetUniverseNil sets the value for Universe to be an explicit nil
func (o *UuidpoolUuidLease) SetUniverseNil() {
	o.Universe.Set(nil)
}

// UnsetUniverse ensures that no value is present for Universe, not even an explicit nil
func (o *UuidpoolUuidLease) UnsetUniverse() {
	o.Universe.Unset()
}

func (o UuidpoolUuidLease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UuidpoolUuidLease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractLease, errPoolAbstractLease := json.Marshal(o.PoolAbstractLease)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	errPoolAbstractLease = json.Unmarshal([]byte(serializedPoolAbstractLease), &toSerialize)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Reservation) {
		toSerialize["Reservation"] = o.Reservation
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.PoolMember.IsSet() {
		toSerialize["PoolMember"] = o.PoolMember.Get()
	}
	if o.Universe.IsSet() {
		toSerialize["Universe"] = o.Universe.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UuidpoolUuidLease) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type UuidpoolUuidLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                        `json:"ObjectType"`
		Reservation *UuidpoolReservationReference `json:"Reservation,omitempty"`
		// UUID Prefix+Suffix numbers.
		Uuid             *string                                `json:"Uuid,omitempty" validate:"regexp=^$|^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"`
		AssignedToEntity NullableMoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
		Pool             NullableUuidpoolPoolRelationship       `json:"Pool,omitempty"`
		PoolMember       NullableUuidpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
		Universe         NullableUuidpoolUniverseRelationship   `json:"Universe,omitempty"`
	}

	varUuidpoolUuidLeaseWithoutEmbeddedStruct := UuidpoolUuidLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUuidpoolUuidLeaseWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolUuidLease := _UuidpoolUuidLease{}
		varUuidpoolUuidLease.ClassId = varUuidpoolUuidLeaseWithoutEmbeddedStruct.ClassId
		varUuidpoolUuidLease.ObjectType = varUuidpoolUuidLeaseWithoutEmbeddedStruct.ObjectType
		varUuidpoolUuidLease.Reservation = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Reservation
		varUuidpoolUuidLease.Uuid = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Uuid
		varUuidpoolUuidLease.AssignedToEntity = varUuidpoolUuidLeaseWithoutEmbeddedStruct.AssignedToEntity
		varUuidpoolUuidLease.Pool = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Pool
		varUuidpoolUuidLease.PoolMember = varUuidpoolUuidLeaseWithoutEmbeddedStruct.PoolMember
		varUuidpoolUuidLease.Universe = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Universe
		*o = UuidpoolUuidLease(varUuidpoolUuidLease)
	} else {
		return err
	}

	varUuidpoolUuidLease := _UuidpoolUuidLease{}

	err = json.Unmarshal(data, &varUuidpoolUuidLease)
	if err == nil {
		o.PoolAbstractLease = varUuidpoolUuidLease.PoolAbstractLease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Reservation")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolAbstractLease := reflect.ValueOf(o.PoolAbstractLease)
		for i := 0; i < reflectPoolAbstractLease.Type().NumField(); i++ {
			t := reflectPoolAbstractLease.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUuidpoolUuidLease struct {
	value *UuidpoolUuidLease
	isSet bool
}

func (v NullableUuidpoolUuidLease) Get() *UuidpoolUuidLease {
	return v.value
}

func (v *NullableUuidpoolUuidLease) Set(val *UuidpoolUuidLease) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolUuidLease) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolUuidLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolUuidLease(val *UuidpoolUuidLease) *NullableUuidpoolUuidLease {
	return &NullableUuidpoolUuidLease{value: val, isSet: true}
}

func (v NullableUuidpoolUuidLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolUuidLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
