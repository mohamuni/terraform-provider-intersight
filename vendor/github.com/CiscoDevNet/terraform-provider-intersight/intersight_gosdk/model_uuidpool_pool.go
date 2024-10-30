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

// checks if the UuidpoolPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UuidpoolPool{}

// UuidpoolPool Pool represents a collection of UUID items that can be allocated to server profiles.
type UuidpoolPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx.
	Prefix           *string             `json:"Prefix,omitempty" validate:"regexp=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}$"`
	UuidSuffixBlocks []UuidpoolUuidBlock `json:"UuidSuffixBlocks,omitempty"`
	// An array of relationships to uuidpoolBlock resources.
	BlockHeads   []UuidpoolBlockRelationship                  `json:"BlockHeads,omitempty"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to uuidpoolReservation resources.
	Reservations         []UuidpoolReservationRelationship `json:"Reservations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolPool UuidpoolPool

// NewUuidpoolPool instantiates a new UuidpoolPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolPool(classId string, objectType string) *UuidpoolPool {
	this := UuidpoolPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewUuidpoolPoolWithDefaults instantiates a new UuidpoolPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolPoolWithDefaults() *UuidpoolPool {
	this := UuidpoolPool{}
	var classId string = "uuidpool.Pool"
	this.ClassId = classId
	var objectType string = "uuidpool.Pool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolPool) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "uuidpool.Pool" of the ClassId field.
func (o *UuidpoolPool) GetDefaultClassId() interface{} {
	return "uuidpool.Pool"
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "uuidpool.Pool" of the ObjectType field.
func (o *UuidpoolPool) GetDefaultObjectType() interface{} {
	return "uuidpool.Pool"
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UuidpoolPool) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UuidpoolPool) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UuidpoolPool) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUuidSuffixBlocks returns the UuidSuffixBlocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolPool) GetUuidSuffixBlocks() []UuidpoolUuidBlock {
	if o == nil {
		var ret []UuidpoolUuidBlock
		return ret
	}
	return o.UuidSuffixBlocks
}

// GetUuidSuffixBlocksOk returns a tuple with the UuidSuffixBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolPool) GetUuidSuffixBlocksOk() ([]UuidpoolUuidBlock, bool) {
	if o == nil || IsNil(o.UuidSuffixBlocks) {
		return nil, false
	}
	return o.UuidSuffixBlocks, true
}

// HasUuidSuffixBlocks returns a boolean if a field has been set.
func (o *UuidpoolPool) HasUuidSuffixBlocks() bool {
	if o != nil && !IsNil(o.UuidSuffixBlocks) {
		return true
	}

	return false
}

// SetUuidSuffixBlocks gets a reference to the given []UuidpoolUuidBlock and assigns it to the UuidSuffixBlocks field.
func (o *UuidpoolPool) SetUuidSuffixBlocks(v []UuidpoolUuidBlock) {
	o.UuidSuffixBlocks = v
}

// GetBlockHeads returns the BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolPool) GetBlockHeads() []UuidpoolBlockRelationship {
	if o == nil {
		var ret []UuidpoolBlockRelationship
		return ret
	}
	return o.BlockHeads
}

// GetBlockHeadsOk returns a tuple with the BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolPool) GetBlockHeadsOk() ([]UuidpoolBlockRelationship, bool) {
	if o == nil || IsNil(o.BlockHeads) {
		return nil, false
	}
	return o.BlockHeads, true
}

// HasBlockHeads returns a boolean if a field has been set.
func (o *UuidpoolPool) HasBlockHeads() bool {
	if o != nil && !IsNil(o.BlockHeads) {
		return true
	}

	return false
}

// SetBlockHeads gets a reference to the given []UuidpoolBlockRelationship and assigns it to the BlockHeads field.
func (o *UuidpoolPool) SetBlockHeads(v []UuidpoolBlockRelationship) {
	o.BlockHeads = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolPool) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *UuidpoolPool) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *UuidpoolPool) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *UuidpoolPool) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *UuidpoolPool) UnsetOrganization() {
	o.Organization.Unset()
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolPool) GetReservations() []UuidpoolReservationRelationship {
	if o == nil {
		var ret []UuidpoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolPool) GetReservationsOk() ([]UuidpoolReservationRelationship, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *UuidpoolPool) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []UuidpoolReservationRelationship and assigns it to the Reservations field.
func (o *UuidpoolPool) SetReservations(v []UuidpoolReservationRelationship) {
	o.Reservations = v
}

func (o UuidpoolPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UuidpoolPool) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Prefix) {
		toSerialize["Prefix"] = o.Prefix
	}
	if o.UuidSuffixBlocks != nil {
		toSerialize["UuidSuffixBlocks"] = o.UuidSuffixBlocks
	}
	if o.BlockHeads != nil {
		toSerialize["BlockHeads"] = o.BlockHeads
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UuidpoolPool) UnmarshalJSON(data []byte) (err error) {
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
	type UuidpoolPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx.
		Prefix           *string             `json:"Prefix,omitempty" validate:"regexp=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}$"`
		UuidSuffixBlocks []UuidpoolUuidBlock `json:"UuidSuffixBlocks,omitempty"`
		// An array of relationships to uuidpoolBlock resources.
		BlockHeads   []UuidpoolBlockRelationship                  `json:"BlockHeads,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to uuidpoolReservation resources.
		Reservations []UuidpoolReservationRelationship `json:"Reservations,omitempty"`
	}

	varUuidpoolPoolWithoutEmbeddedStruct := UuidpoolPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUuidpoolPoolWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolPool := _UuidpoolPool{}
		varUuidpoolPool.ClassId = varUuidpoolPoolWithoutEmbeddedStruct.ClassId
		varUuidpoolPool.ObjectType = varUuidpoolPoolWithoutEmbeddedStruct.ObjectType
		varUuidpoolPool.Prefix = varUuidpoolPoolWithoutEmbeddedStruct.Prefix
		varUuidpoolPool.UuidSuffixBlocks = varUuidpoolPoolWithoutEmbeddedStruct.UuidSuffixBlocks
		varUuidpoolPool.BlockHeads = varUuidpoolPoolWithoutEmbeddedStruct.BlockHeads
		varUuidpoolPool.Organization = varUuidpoolPoolWithoutEmbeddedStruct.Organization
		varUuidpoolPool.Reservations = varUuidpoolPoolWithoutEmbeddedStruct.Reservations
		*o = UuidpoolPool(varUuidpoolPool)
	} else {
		return err
	}

	varUuidpoolPool := _UuidpoolPool{}

	err = json.Unmarshal(data, &varUuidpoolPool)
	if err == nil {
		o.PoolAbstractPool = varUuidpoolPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Prefix")
		delete(additionalProperties, "UuidSuffixBlocks")
		delete(additionalProperties, "BlockHeads")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Reservations")

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

type NullableUuidpoolPool struct {
	value *UuidpoolPool
	isSet bool
}

func (v NullableUuidpoolPool) Get() *UuidpoolPool {
	return v.value
}

func (v *NullableUuidpoolPool) Set(val *UuidpoolPool) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolPool) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolPool(val *UuidpoolPool) *NullableUuidpoolPool {
	return &NullableUuidpoolPool{value: val, isSet: true}
}

func (v NullableUuidpoolPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
