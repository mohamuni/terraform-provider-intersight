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

// checks if the IppoolPoolMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IppoolPoolMember{}

// IppoolPoolMember PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
type IppoolPoolMember struct {
	PoolAbstractIdPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType *string `json:"IpType,omitempty"`
	// IPv4 Address of this pool member.
	IpV4Address *string `json:"IpV4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// IPv6 Address of this pool member.
	IpV6Address          *string                               `json:"IpV6Address,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
	AssignedToEntity     NullableMoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	IpBlock              NullableIppoolShadowBlockRelationship `json:"IpBlock,omitempty"`
	Peer                 NullableIppoolIpLeaseRelationship     `json:"Peer,omitempty"`
	Pool                 NullableIppoolShadowPoolRelationship  `json:"Pool,omitempty"`
	Reservation          NullableIppoolReservationRelationship `json:"Reservation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolPoolMember IppoolPoolMember

// NewIppoolPoolMember instantiates a new IppoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolPoolMember(classId string, objectType string) *IppoolPoolMember {
	this := IppoolPoolMember{}
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

// NewIppoolPoolMemberWithDefaults instantiates a new IppoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolPoolMemberWithDefaults() *IppoolPoolMember {
	this := IppoolPoolMember{}
	var classId string = "ippool.PoolMember"
	this.ClassId = classId
	var objectType string = "ippool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ippool.PoolMember" of the ClassId field.
func (o *IppoolPoolMember) GetDefaultClassId() interface{} {
	return "ippool.PoolMember"
}

// GetObjectType returns the ObjectType field value
func (o *IppoolPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ippool.PoolMember" of the ObjectType field.
func (o *IppoolPoolMember) GetDefaultObjectType() interface{} {
	return "ippool.PoolMember"
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpType() string {
	if o == nil || IsNil(o.IpType) {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IpType) {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpType() bool {
	if o != nil && !IsNil(o.IpType) {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolPoolMember) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpV4Address() string {
	if o == nil || IsNil(o.IpV4Address) {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpV4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpV4Address) {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpV4Address() bool {
	if o != nil && !IsNil(o.IpV4Address) {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolPoolMember) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV6Address returns the IpV6Address field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpV6Address() string {
	if o == nil || IsNil(o.IpV6Address) {
		var ret string
		return ret
	}
	return *o.IpV6Address
}

// GetIpV6AddressOk returns a tuple with the IpV6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpV6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpV6Address) {
		return nil, false
	}
	return o.IpV6Address, true
}

// HasIpV6Address returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpV6Address() bool {
	if o != nil && !IsNil(o.IpV6Address) {
		return true
	}

	return false
}

// SetIpV6Address gets a reference to the given string and assigns it to the IpV6Address field.
func (o *IppoolPoolMember) SetIpV6Address(v string) {
	o.IpV6Address = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IppoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *IppoolPoolMember) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *IppoolPoolMember) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetIpBlock returns the IpBlock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolPoolMember) GetIpBlock() IppoolShadowBlockRelationship {
	if o == nil || IsNil(o.IpBlock.Get()) {
		var ret IppoolShadowBlockRelationship
		return ret
	}
	return *o.IpBlock.Get()
}

// GetIpBlockOk returns a tuple with the IpBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolPoolMember) GetIpBlockOk() (*IppoolShadowBlockRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpBlock.Get(), o.IpBlock.IsSet()
}

// HasIpBlock returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpBlock() bool {
	if o != nil && o.IpBlock.IsSet() {
		return true
	}

	return false
}

// SetIpBlock gets a reference to the given NullableIppoolShadowBlockRelationship and assigns it to the IpBlock field.
func (o *IppoolPoolMember) SetIpBlock(v IppoolShadowBlockRelationship) {
	o.IpBlock.Set(&v)
}

// SetIpBlockNil sets the value for IpBlock to be an explicit nil
func (o *IppoolPoolMember) SetIpBlockNil() {
	o.IpBlock.Set(nil)
}

// UnsetIpBlock ensures that no value is present for IpBlock, not even an explicit nil
func (o *IppoolPoolMember) UnsetIpBlock() {
	o.IpBlock.Unset()
}

// GetPeer returns the Peer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolPoolMember) GetPeer() IppoolIpLeaseRelationship {
	if o == nil || IsNil(o.Peer.Get()) {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.Peer.Get()
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolPoolMember) GetPeerOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peer.Get(), o.Peer.IsSet()
}

// HasPeer returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer.IsSet() {
		return true
	}

	return false
}

// SetPeer gets a reference to the given NullableIppoolIpLeaseRelationship and assigns it to the Peer field.
func (o *IppoolPoolMember) SetPeer(v IppoolIpLeaseRelationship) {
	o.Peer.Set(&v)
}

// SetPeerNil sets the value for Peer to be an explicit nil
func (o *IppoolPoolMember) SetPeerNil() {
	o.Peer.Set(nil)
}

// UnsetPeer ensures that no value is present for Peer, not even an explicit nil
func (o *IppoolPoolMember) UnsetPeer() {
	o.Peer.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolPoolMember) GetPool() IppoolShadowPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolPoolMember) GetPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableIppoolShadowPoolRelationship and assigns it to the Pool field.
func (o *IppoolPoolMember) SetPool(v IppoolShadowPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *IppoolPoolMember) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *IppoolPoolMember) UnsetPool() {
	o.Pool.Unset()
}

// GetReservation returns the Reservation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolPoolMember) GetReservation() IppoolReservationRelationship {
	if o == nil || IsNil(o.Reservation.Get()) {
		var ret IppoolReservationRelationship
		return ret
	}
	return *o.Reservation.Get()
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolPoolMember) GetReservationOk() (*IppoolReservationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reservation.Get(), o.Reservation.IsSet()
}

// HasReservation returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasReservation() bool {
	if o != nil && o.Reservation.IsSet() {
		return true
	}

	return false
}

// SetReservation gets a reference to the given NullableIppoolReservationRelationship and assigns it to the Reservation field.
func (o *IppoolPoolMember) SetReservation(v IppoolReservationRelationship) {
	o.Reservation.Set(&v)
}

// SetReservationNil sets the value for Reservation to be an explicit nil
func (o *IppoolPoolMember) SetReservationNil() {
	o.Reservation.Set(nil)
}

// UnsetReservation ensures that no value is present for Reservation, not even an explicit nil
func (o *IppoolPoolMember) UnsetReservation() {
	o.Reservation.Unset()
}

func (o IppoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IppoolPoolMember) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IpType) {
		toSerialize["IpType"] = o.IpType
	}
	if !IsNil(o.IpV4Address) {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if !IsNil(o.IpV6Address) {
		toSerialize["IpV6Address"] = o.IpV6Address
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.IpBlock.IsSet() {
		toSerialize["IpBlock"] = o.IpBlock.Get()
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

func (o *IppoolPoolMember) UnmarshalJSON(data []byte) (err error) {
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
	type IppoolPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
		IpType *string `json:"IpType,omitempty"`
		// IPv4 Address of this pool member.
		IpV4Address *string `json:"IpV4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// IPv6 Address of this pool member.
		IpV6Address      *string                               `json:"IpV6Address,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
		AssignedToEntity NullableMoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
		IpBlock          NullableIppoolShadowBlockRelationship `json:"IpBlock,omitempty"`
		Peer             NullableIppoolIpLeaseRelationship     `json:"Peer,omitempty"`
		Pool             NullableIppoolShadowPoolRelationship  `json:"Pool,omitempty"`
		Reservation      NullableIppoolReservationRelationship `json:"Reservation,omitempty"`
	}

	varIppoolPoolMemberWithoutEmbeddedStruct := IppoolPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIppoolPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varIppoolPoolMember := _IppoolPoolMember{}
		varIppoolPoolMember.ClassId = varIppoolPoolMemberWithoutEmbeddedStruct.ClassId
		varIppoolPoolMember.ObjectType = varIppoolPoolMemberWithoutEmbeddedStruct.ObjectType
		varIppoolPoolMember.IpType = varIppoolPoolMemberWithoutEmbeddedStruct.IpType
		varIppoolPoolMember.IpV4Address = varIppoolPoolMemberWithoutEmbeddedStruct.IpV4Address
		varIppoolPoolMember.IpV6Address = varIppoolPoolMemberWithoutEmbeddedStruct.IpV6Address
		varIppoolPoolMember.AssignedToEntity = varIppoolPoolMemberWithoutEmbeddedStruct.AssignedToEntity
		varIppoolPoolMember.IpBlock = varIppoolPoolMemberWithoutEmbeddedStruct.IpBlock
		varIppoolPoolMember.Peer = varIppoolPoolMemberWithoutEmbeddedStruct.Peer
		varIppoolPoolMember.Pool = varIppoolPoolMemberWithoutEmbeddedStruct.Pool
		varIppoolPoolMember.Reservation = varIppoolPoolMemberWithoutEmbeddedStruct.Reservation
		*o = IppoolPoolMember(varIppoolPoolMember)
	} else {
		return err
	}

	varIppoolPoolMember := _IppoolPoolMember{}

	err = json.Unmarshal(data, &varIppoolPoolMember)
	if err == nil {
		o.PoolAbstractIdPoolMember = varIppoolPoolMember.PoolAbstractIdPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Address")
		delete(additionalProperties, "IpV6Address")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "IpBlock")
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

type NullableIppoolPoolMember struct {
	value *IppoolPoolMember
	isSet bool
}

func (v NullableIppoolPoolMember) Get() *IppoolPoolMember {
	return v.value
}

func (v *NullableIppoolPoolMember) Set(val *IppoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolPoolMember(val *IppoolPoolMember) *NullableIppoolPoolMember {
	return &NullableIppoolPoolMember{value: val, isSet: true}
}

func (v NullableIppoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
