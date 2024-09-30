/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
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

// checks if the IppoolShadowPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IppoolShadowPool{}

// IppoolShadowPool Shadow Pool is a tracking object created on behalf of an IP pool, for each VRF.
type IppoolShadowPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                   `json:"ObjectType"`
	IpV4Blocks []IppoolIpV4Block        `json:"IpV4Blocks,omitempty"`
	IpV4Config NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
	IpV6Blocks []IppoolIpV6Block        `json:"IpV6Blocks,omitempty"`
	IpV6Config NullableIppoolIpV6Config `json:"IpV6Config,omitempty"`
	// Number of IPv4 addresses currently in use.
	V4Assigned *int64 `json:"V4Assigned,omitempty"`
	// Number of IPv4 addresses in this pool.
	V4Size *int64 `json:"V4Size,omitempty"`
	// Number of IPv6 addresses currently in use.
	V6Assigned *int64 `json:"V6Assigned,omitempty"`
	// Number of IPv6 addresses in this pool.
	V6Size *int64 `json:"V6Size,omitempty"`
	// An array of relationships to ippoolShadowBlock resources.
	IpBlockHeads []IppoolShadowBlockRelationship `json:"IpBlockHeads,omitempty"`
	Pool         NullableIppoolPoolRelationship  `json:"Pool,omitempty"`
	// An array of relationships to ippoolReservation resources.
	Reservations         []IppoolReservationRelationship `json:"Reservations,omitempty"`
	Vrf                  NullableVrfVrfRelationship      `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolShadowPool IppoolShadowPool

// NewIppoolShadowPool instantiates a new IppoolShadowPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolShadowPool(classId string, objectType string) *IppoolShadowPool {
	this := IppoolShadowPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewIppoolShadowPoolWithDefaults instantiates a new IppoolShadowPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolShadowPoolWithDefaults() *IppoolShadowPool {
	this := IppoolShadowPool{}
	var classId string = "ippool.ShadowPool"
	this.ClassId = classId
	var objectType string = "ippool.ShadowPool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolShadowPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolShadowPool) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ippool.ShadowPool" of the ClassId field.
func (o *IppoolShadowPool) GetDefaultClassId() interface{} {
	return "ippool.ShadowPool"
}

// GetObjectType returns the ObjectType field value
func (o *IppoolShadowPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolShadowPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ippool.ShadowPool" of the ObjectType field.
func (o *IppoolShadowPool) GetDefaultObjectType() interface{} {
	return "ippool.ShadowPool"
}

// GetIpV4Blocks returns the IpV4Blocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV4Blocks() []IppoolIpV4Block {
	if o == nil {
		var ret []IppoolIpV4Block
		return ret
	}
	return o.IpV4Blocks
}

// GetIpV4BlocksOk returns a tuple with the IpV4Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV4BlocksOk() ([]IppoolIpV4Block, bool) {
	if o == nil || IsNil(o.IpV4Blocks) {
		return nil, false
	}
	return o.IpV4Blocks, true
}

// HasIpV4Blocks returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV4Blocks() bool {
	if o != nil && !IsNil(o.IpV4Blocks) {
		return true
	}

	return false
}

// SetIpV4Blocks gets a reference to the given []IppoolIpV4Block and assigns it to the IpV4Blocks field.
func (o *IppoolShadowPool) SetIpV4Blocks(v []IppoolIpV4Block) {
	o.IpV4Blocks = v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV4Config() IppoolIpV4Config {
	if o == nil || IsNil(o.IpV4Config.Get()) {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config.Get()
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV4Config.Get(), o.IpV4Config.IsSet()
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV4Config() bool {
	if o != nil && o.IpV4Config.IsSet() {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given NullableIppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolShadowPool) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config.Set(&v)
}

// SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil
func (o *IppoolShadowPool) SetIpV4ConfigNil() {
	o.IpV4Config.Set(nil)
}

// UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
func (o *IppoolShadowPool) UnsetIpV4Config() {
	o.IpV4Config.Unset()
}

// GetIpV6Blocks returns the IpV6Blocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV6Blocks() []IppoolIpV6Block {
	if o == nil {
		var ret []IppoolIpV6Block
		return ret
	}
	return o.IpV6Blocks
}

// GetIpV6BlocksOk returns a tuple with the IpV6Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV6BlocksOk() ([]IppoolIpV6Block, bool) {
	if o == nil || IsNil(o.IpV6Blocks) {
		return nil, false
	}
	return o.IpV6Blocks, true
}

// HasIpV6Blocks returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV6Blocks() bool {
	if o != nil && !IsNil(o.IpV6Blocks) {
		return true
	}

	return false
}

// SetIpV6Blocks gets a reference to the given []IppoolIpV6Block and assigns it to the IpV6Blocks field.
func (o *IppoolShadowPool) SetIpV6Blocks(v []IppoolIpV6Block) {
	o.IpV6Blocks = v
}

// GetIpV6Config returns the IpV6Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV6Config() IppoolIpV6Config {
	if o == nil || IsNil(o.IpV6Config.Get()) {
		var ret IppoolIpV6Config
		return ret
	}
	return *o.IpV6Config.Get()
}

// GetIpV6ConfigOk returns a tuple with the IpV6Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV6ConfigOk() (*IppoolIpV6Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV6Config.Get(), o.IpV6Config.IsSet()
}

// HasIpV6Config returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV6Config() bool {
	if o != nil && o.IpV6Config.IsSet() {
		return true
	}

	return false
}

// SetIpV6Config gets a reference to the given NullableIppoolIpV6Config and assigns it to the IpV6Config field.
func (o *IppoolShadowPool) SetIpV6Config(v IppoolIpV6Config) {
	o.IpV6Config.Set(&v)
}

// SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil
func (o *IppoolShadowPool) SetIpV6ConfigNil() {
	o.IpV6Config.Set(nil)
}

// UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
func (o *IppoolShadowPool) UnsetIpV6Config() {
	o.IpV6Config.Unset()
}

// GetV4Assigned returns the V4Assigned field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV4Assigned() int64 {
	if o == nil || IsNil(o.V4Assigned) {
		var ret int64
		return ret
	}
	return *o.V4Assigned
}

// GetV4AssignedOk returns a tuple with the V4Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV4AssignedOk() (*int64, bool) {
	if o == nil || IsNil(o.V4Assigned) {
		return nil, false
	}
	return o.V4Assigned, true
}

// HasV4Assigned returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV4Assigned() bool {
	if o != nil && !IsNil(o.V4Assigned) {
		return true
	}

	return false
}

// SetV4Assigned gets a reference to the given int64 and assigns it to the V4Assigned field.
func (o *IppoolShadowPool) SetV4Assigned(v int64) {
	o.V4Assigned = &v
}

// GetV4Size returns the V4Size field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV4Size() int64 {
	if o == nil || IsNil(o.V4Size) {
		var ret int64
		return ret
	}
	return *o.V4Size
}

// GetV4SizeOk returns a tuple with the V4Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV4SizeOk() (*int64, bool) {
	if o == nil || IsNil(o.V4Size) {
		return nil, false
	}
	return o.V4Size, true
}

// HasV4Size returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV4Size() bool {
	if o != nil && !IsNil(o.V4Size) {
		return true
	}

	return false
}

// SetV4Size gets a reference to the given int64 and assigns it to the V4Size field.
func (o *IppoolShadowPool) SetV4Size(v int64) {
	o.V4Size = &v
}

// GetV6Assigned returns the V6Assigned field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV6Assigned() int64 {
	if o == nil || IsNil(o.V6Assigned) {
		var ret int64
		return ret
	}
	return *o.V6Assigned
}

// GetV6AssignedOk returns a tuple with the V6Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV6AssignedOk() (*int64, bool) {
	if o == nil || IsNil(o.V6Assigned) {
		return nil, false
	}
	return o.V6Assigned, true
}

// HasV6Assigned returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV6Assigned() bool {
	if o != nil && !IsNil(o.V6Assigned) {
		return true
	}

	return false
}

// SetV6Assigned gets a reference to the given int64 and assigns it to the V6Assigned field.
func (o *IppoolShadowPool) SetV6Assigned(v int64) {
	o.V6Assigned = &v
}

// GetV6Size returns the V6Size field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV6Size() int64 {
	if o == nil || IsNil(o.V6Size) {
		var ret int64
		return ret
	}
	return *o.V6Size
}

// GetV6SizeOk returns a tuple with the V6Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV6SizeOk() (*int64, bool) {
	if o == nil || IsNil(o.V6Size) {
		return nil, false
	}
	return o.V6Size, true
}

// HasV6Size returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV6Size() bool {
	if o != nil && !IsNil(o.V6Size) {
		return true
	}

	return false
}

// SetV6Size gets a reference to the given int64 and assigns it to the V6Size field.
func (o *IppoolShadowPool) SetV6Size(v int64) {
	o.V6Size = &v
}

// GetIpBlockHeads returns the IpBlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpBlockHeads() []IppoolShadowBlockRelationship {
	if o == nil {
		var ret []IppoolShadowBlockRelationship
		return ret
	}
	return o.IpBlockHeads
}

// GetIpBlockHeadsOk returns a tuple with the IpBlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpBlockHeadsOk() ([]IppoolShadowBlockRelationship, bool) {
	if o == nil || IsNil(o.IpBlockHeads) {
		return nil, false
	}
	return o.IpBlockHeads, true
}

// HasIpBlockHeads returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpBlockHeads() bool {
	if o != nil && !IsNil(o.IpBlockHeads) {
		return true
	}

	return false
}

// SetIpBlockHeads gets a reference to the given []IppoolShadowBlockRelationship and assigns it to the IpBlockHeads field.
func (o *IppoolShadowPool) SetIpBlockHeads(v []IppoolShadowBlockRelationship) {
	o.IpBlockHeads = v
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolShadowPool) SetPool(v IppoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *IppoolShadowPool) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *IppoolShadowPool) UnsetPool() {
	o.Pool.Unset()
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetReservations() []IppoolReservationRelationship {
	if o == nil {
		var ret []IppoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetReservationsOk() ([]IppoolReservationRelationship, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []IppoolReservationRelationship and assigns it to the Reservations field.
func (o *IppoolShadowPool) SetReservations(v []IppoolReservationRelationship) {
	o.Reservations = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolShadowPool) SetVrf(v VrfVrfRelationship) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *IppoolShadowPool) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *IppoolShadowPool) UnsetVrf() {
	o.Vrf.Unset()
}

func (o IppoolShadowPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IppoolShadowPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPool, errPoolAbstractPool := json.Marshal(o.PoolAbstractPool)
	if errPoolAbstractPool != nil {
		return map[string]interface{}{}, errPoolAbstractPool
	}
	errPoolAbstractPool = json.Unmarshal([]byte(serializedPoolAbstractPool), &toSerialize)
	if errPoolAbstractPool != nil {
		return map[string]interface{}{}, errPoolAbstractPool
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.IpV4Blocks != nil {
		toSerialize["IpV4Blocks"] = o.IpV4Blocks
	}
	if o.IpV4Config.IsSet() {
		toSerialize["IpV4Config"] = o.IpV4Config.Get()
	}
	if o.IpV6Blocks != nil {
		toSerialize["IpV6Blocks"] = o.IpV6Blocks
	}
	if o.IpV6Config.IsSet() {
		toSerialize["IpV6Config"] = o.IpV6Config.Get()
	}
	if !IsNil(o.V4Assigned) {
		toSerialize["V4Assigned"] = o.V4Assigned
	}
	if !IsNil(o.V4Size) {
		toSerialize["V4Size"] = o.V4Size
	}
	if !IsNil(o.V6Assigned) {
		toSerialize["V6Assigned"] = o.V6Assigned
	}
	if !IsNil(o.V6Size) {
		toSerialize["V6Size"] = o.V6Size
	}
	if o.IpBlockHeads != nil {
		toSerialize["IpBlockHeads"] = o.IpBlockHeads
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}
	if o.Vrf.IsSet() {
		toSerialize["Vrf"] = o.Vrf.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IppoolShadowPool) UnmarshalJSON(data []byte) (err error) {
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
	type IppoolShadowPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                   `json:"ObjectType"`
		IpV4Blocks []IppoolIpV4Block        `json:"IpV4Blocks,omitempty"`
		IpV4Config NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
		IpV6Blocks []IppoolIpV6Block        `json:"IpV6Blocks,omitempty"`
		IpV6Config NullableIppoolIpV6Config `json:"IpV6Config,omitempty"`
		// Number of IPv4 addresses currently in use.
		V4Assigned *int64 `json:"V4Assigned,omitempty"`
		// Number of IPv4 addresses in this pool.
		V4Size *int64 `json:"V4Size,omitempty"`
		// Number of IPv6 addresses currently in use.
		V6Assigned *int64 `json:"V6Assigned,omitempty"`
		// Number of IPv6 addresses in this pool.
		V6Size *int64 `json:"V6Size,omitempty"`
		// An array of relationships to ippoolShadowBlock resources.
		IpBlockHeads []IppoolShadowBlockRelationship `json:"IpBlockHeads,omitempty"`
		Pool         NullableIppoolPoolRelationship  `json:"Pool,omitempty"`
		// An array of relationships to ippoolReservation resources.
		Reservations []IppoolReservationRelationship `json:"Reservations,omitempty"`
		Vrf          NullableVrfVrfRelationship      `json:"Vrf,omitempty"`
	}

	varIppoolShadowPoolWithoutEmbeddedStruct := IppoolShadowPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIppoolShadowPoolWithoutEmbeddedStruct)
	if err == nil {
		varIppoolShadowPool := _IppoolShadowPool{}
		varIppoolShadowPool.ClassId = varIppoolShadowPoolWithoutEmbeddedStruct.ClassId
		varIppoolShadowPool.ObjectType = varIppoolShadowPoolWithoutEmbeddedStruct.ObjectType
		varIppoolShadowPool.IpV4Blocks = varIppoolShadowPoolWithoutEmbeddedStruct.IpV4Blocks
		varIppoolShadowPool.IpV4Config = varIppoolShadowPoolWithoutEmbeddedStruct.IpV4Config
		varIppoolShadowPool.IpV6Blocks = varIppoolShadowPoolWithoutEmbeddedStruct.IpV6Blocks
		varIppoolShadowPool.IpV6Config = varIppoolShadowPoolWithoutEmbeddedStruct.IpV6Config
		varIppoolShadowPool.V4Assigned = varIppoolShadowPoolWithoutEmbeddedStruct.V4Assigned
		varIppoolShadowPool.V4Size = varIppoolShadowPoolWithoutEmbeddedStruct.V4Size
		varIppoolShadowPool.V6Assigned = varIppoolShadowPoolWithoutEmbeddedStruct.V6Assigned
		varIppoolShadowPool.V6Size = varIppoolShadowPoolWithoutEmbeddedStruct.V6Size
		varIppoolShadowPool.IpBlockHeads = varIppoolShadowPoolWithoutEmbeddedStruct.IpBlockHeads
		varIppoolShadowPool.Pool = varIppoolShadowPoolWithoutEmbeddedStruct.Pool
		varIppoolShadowPool.Reservations = varIppoolShadowPoolWithoutEmbeddedStruct.Reservations
		varIppoolShadowPool.Vrf = varIppoolShadowPoolWithoutEmbeddedStruct.Vrf
		*o = IppoolShadowPool(varIppoolShadowPool)
	} else {
		return err
	}

	varIppoolShadowPool := _IppoolShadowPool{}

	err = json.Unmarshal(data, &varIppoolShadowPool)
	if err == nil {
		o.PoolAbstractPool = varIppoolShadowPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpV4Blocks")
		delete(additionalProperties, "IpV4Config")
		delete(additionalProperties, "IpV6Blocks")
		delete(additionalProperties, "IpV6Config")
		delete(additionalProperties, "V4Assigned")
		delete(additionalProperties, "V4Size")
		delete(additionalProperties, "V6Assigned")
		delete(additionalProperties, "V6Size")
		delete(additionalProperties, "IpBlockHeads")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Reservations")
		delete(additionalProperties, "Vrf")

		// remove fields from embedded structs
		reflectPoolAbstractPool := reflect.ValueOf(o.PoolAbstractPool)
		for i := 0; i < reflectPoolAbstractPool.Type().NumField(); i++ {
			t := reflectPoolAbstractPool.Type().Field(i)

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

type NullableIppoolShadowPool struct {
	value *IppoolShadowPool
	isSet bool
}

func (v NullableIppoolShadowPool) Get() *IppoolShadowPool {
	return v.value
}

func (v *NullableIppoolShadowPool) Set(val *IppoolShadowPool) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolShadowPool) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolShadowPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolShadowPool(val *IppoolShadowPool) *NullableIppoolShadowPool {
	return &NullableIppoolShadowPool{value: val, isSet: true}
}

func (v NullableIppoolShadowPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolShadowPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
