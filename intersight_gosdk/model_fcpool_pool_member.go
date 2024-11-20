/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the FcpoolPoolMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FcpoolPoolMember{}

// FcpoolPoolMember PoolMember represents a single WWN ID that is part of a pool.
type FcpoolPoolMember struct {
	PoolAbstractIdPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// WWN ID of this pool member.
	WwnId                *string                               `json:"WwnId,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
	AssignedToEntity     NullableMoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	BlockHead            NullableFcpoolFcBlockRelationship     `json:"BlockHead,omitempty"`
	Peer                 NullableFcpoolLeaseRelationship       `json:"Peer,omitempty"`
	Pool                 NullableFcpoolPoolRelationship        `json:"Pool,omitempty"`
	Reservation          NullableFcpoolReservationRelationship `json:"Reservation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolPoolMember FcpoolPoolMember

// NewFcpoolPoolMember instantiates a new FcpoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolPoolMember(classId string, objectType string) *FcpoolPoolMember {
	this := FcpoolPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	var assignedByAnother bool = false
	this.AssignedByAnother = &assignedByAnother
	var reserved bool = false
	this.Reserved = &reserved
	return &this
}

// NewFcpoolPoolMemberWithDefaults instantiates a new FcpoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolPoolMemberWithDefaults() *FcpoolPoolMember {
	this := FcpoolPoolMember{}
	var classId string = "fcpool.PoolMember"
	this.ClassId = classId
	var objectType string = "fcpool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fcpool.PoolMember" of the ClassId field.
func (o *FcpoolPoolMember) GetDefaultClassId() interface{} {
	return "fcpool.PoolMember"
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fcpool.PoolMember" of the ObjectType field.
func (o *FcpoolPoolMember) GetDefaultObjectType() interface{} {
	return "fcpool.PoolMember"
}

// GetWwnId returns the WwnId field value if set, zero value otherwise.
func (o *FcpoolPoolMember) GetWwnId() string {
	if o == nil || IsNil(o.WwnId) {
		var ret string
		return ret
	}
	return *o.WwnId
}

// GetWwnIdOk returns a tuple with the WwnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetWwnIdOk() (*string, bool) {
	if o == nil || IsNil(o.WwnId) {
		return nil, false
	}
	return o.WwnId, true
}

// HasWwnId returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasWwnId() bool {
	if o != nil && !IsNil(o.WwnId) {
		return true
	}

	return false
}

// SetWwnId gets a reference to the given string and assigns it to the WwnId field.
func (o *FcpoolPoolMember) SetWwnId(v string) {
	o.WwnId = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *FcpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *FcpoolPoolMember) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *FcpoolPoolMember) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolMember) GetBlockHead() FcpoolFcBlockRelationship {
	if o == nil || IsNil(o.BlockHead.Get()) {
		var ret FcpoolFcBlockRelationship
		return ret
	}
	return *o.BlockHead.Get()
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolMember) GetBlockHeadOk() (*FcpoolFcBlockRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockHead.Get(), o.BlockHead.IsSet()
}

// HasBlockHead returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasBlockHead() bool {
	if o != nil && o.BlockHead.IsSet() {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given NullableFcpoolFcBlockRelationship and assigns it to the BlockHead field.
func (o *FcpoolPoolMember) SetBlockHead(v FcpoolFcBlockRelationship) {
	o.BlockHead.Set(&v)
}

// SetBlockHeadNil sets the value for BlockHead to be an explicit nil
func (o *FcpoolPoolMember) SetBlockHeadNil() {
	o.BlockHead.Set(nil)
}

// UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
func (o *FcpoolPoolMember) UnsetBlockHead() {
	o.BlockHead.Unset()
}

// GetPeer returns the Peer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolMember) GetPeer() FcpoolLeaseRelationship {
	if o == nil || IsNil(o.Peer.Get()) {
		var ret FcpoolLeaseRelationship
		return ret
	}
	return *o.Peer.Get()
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolMember) GetPeerOk() (*FcpoolLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peer.Get(), o.Peer.IsSet()
}

// HasPeer returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer.IsSet() {
		return true
	}

	return false
}

// SetPeer gets a reference to the given NullableFcpoolLeaseRelationship and assigns it to the Peer field.
func (o *FcpoolPoolMember) SetPeer(v FcpoolLeaseRelationship) {
	o.Peer.Set(&v)
}

// SetPeerNil sets the value for Peer to be an explicit nil
func (o *FcpoolPoolMember) SetPeerNil() {
	o.Peer.Set(nil)
}

// UnsetPeer ensures that no value is present for Peer, not even an explicit nil
func (o *FcpoolPoolMember) UnsetPeer() {
	o.Peer.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolMember) GetPool() FcpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolMember) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableFcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolPoolMember) SetPool(v FcpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *FcpoolPoolMember) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *FcpoolPoolMember) UnsetPool() {
	o.Pool.Unset()
}

// GetReservation returns the Reservation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolMember) GetReservation() FcpoolReservationRelationship {
	if o == nil || IsNil(o.Reservation.Get()) {
		var ret FcpoolReservationRelationship
		return ret
	}
	return *o.Reservation.Get()
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolMember) GetReservationOk() (*FcpoolReservationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reservation.Get(), o.Reservation.IsSet()
}

// HasReservation returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasReservation() bool {
	if o != nil && o.Reservation.IsSet() {
		return true
	}

	return false
}

// SetReservation gets a reference to the given NullableFcpoolReservationRelationship and assigns it to the Reservation field.
func (o *FcpoolPoolMember) SetReservation(v FcpoolReservationRelationship) {
	o.Reservation.Set(&v)
}

// SetReservationNil sets the value for Reservation to be an explicit nil
func (o *FcpoolPoolMember) SetReservationNil() {
	o.Reservation.Set(nil)
}

// UnsetReservation ensures that no value is present for Reservation, not even an explicit nil
func (o *FcpoolPoolMember) UnsetReservation() {
	o.Reservation.Unset()
}

func (o FcpoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FcpoolPoolMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractIdPoolMember, errPoolAbstractIdPoolMember := json.Marshal(o.PoolAbstractIdPoolMember)
	if errPoolAbstractIdPoolMember != nil {
		return map[string]interface{}{}, errPoolAbstractIdPoolMember
	}
	errPoolAbstractIdPoolMember = json.Unmarshal([]byte(serializedPoolAbstractIdPoolMember), &toSerialize)
	if errPoolAbstractIdPoolMember != nil {
		return map[string]interface{}{}, errPoolAbstractIdPoolMember
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.WwnId) {
		toSerialize["WwnId"] = o.WwnId
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.BlockHead.IsSet() {
		toSerialize["BlockHead"] = o.BlockHead.Get()
	}
	if o.Peer.IsSet() {
		toSerialize["Peer"] = o.Peer.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.Reservation.IsSet() {
		toSerialize["Reservation"] = o.Reservation.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FcpoolPoolMember) UnmarshalJSON(data []byte) (err error) {
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
	type FcpoolPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// WWN ID of this pool member.
		WwnId            *string                               `json:"WwnId,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
		AssignedToEntity NullableMoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
		BlockHead        NullableFcpoolFcBlockRelationship     `json:"BlockHead,omitempty"`
		Peer             NullableFcpoolLeaseRelationship       `json:"Peer,omitempty"`
		Pool             NullableFcpoolPoolRelationship        `json:"Pool,omitempty"`
		Reservation      NullableFcpoolReservationRelationship `json:"Reservation,omitempty"`
	}

	varFcpoolPoolMemberWithoutEmbeddedStruct := FcpoolPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFcpoolPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varFcpoolPoolMember := _FcpoolPoolMember{}
		varFcpoolPoolMember.ClassId = varFcpoolPoolMemberWithoutEmbeddedStruct.ClassId
		varFcpoolPoolMember.ObjectType = varFcpoolPoolMemberWithoutEmbeddedStruct.ObjectType
		varFcpoolPoolMember.WwnId = varFcpoolPoolMemberWithoutEmbeddedStruct.WwnId
		varFcpoolPoolMember.AssignedToEntity = varFcpoolPoolMemberWithoutEmbeddedStruct.AssignedToEntity
		varFcpoolPoolMember.BlockHead = varFcpoolPoolMemberWithoutEmbeddedStruct.BlockHead
		varFcpoolPoolMember.Peer = varFcpoolPoolMemberWithoutEmbeddedStruct.Peer
		varFcpoolPoolMember.Pool = varFcpoolPoolMemberWithoutEmbeddedStruct.Pool
		varFcpoolPoolMember.Reservation = varFcpoolPoolMemberWithoutEmbeddedStruct.Reservation
		*o = FcpoolPoolMember(varFcpoolPoolMember)
	} else {
		return err
	}

	varFcpoolPoolMember := _FcpoolPoolMember{}

	err = json.Unmarshal(data, &varFcpoolPoolMember)
	if err == nil {
		o.PoolAbstractIdPoolMember = varFcpoolPoolMember.PoolAbstractIdPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "WwnId")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Reservation")

		// remove fields from embedded structs
		reflectPoolAbstractIdPoolMember := reflect.ValueOf(o.PoolAbstractIdPoolMember)
		for i := 0; i < reflectPoolAbstractIdPoolMember.Type().NumField(); i++ {
			t := reflectPoolAbstractIdPoolMember.Type().Field(i)

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

type NullableFcpoolPoolMember struct {
	value *FcpoolPoolMember
	isSet bool
}

func (v NullableFcpoolPoolMember) Get() *FcpoolPoolMember {
	return v.value
}

func (v *NullableFcpoolPoolMember) Set(val *FcpoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolPoolMember(val *FcpoolPoolMember) *NullableFcpoolPoolMember {
	return &NullableFcpoolPoolMember{value: val, isSet: true}
}

func (v NullableFcpoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
