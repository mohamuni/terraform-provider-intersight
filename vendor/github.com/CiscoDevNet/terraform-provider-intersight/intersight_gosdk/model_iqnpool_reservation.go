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

// checks if the IqnpoolReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IqnpoolReservation{}

// IqnpoolReservation The IQN reservation object, used to hold reserved identity.
type IqnpoolReservation struct {
	PoolReservation
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IQN identity to be reserved.
	Identity *string `json:"Identity,omitempty" validate:"regexp=^$|^(?:iqn\\\\.[0-9]{4}-[0-9]{2}(?:\\\\.[A-Za-z](?:[A-Za-z0-9\\\\-]*[A-Za-z0-9])?)+(?::.*)?|eui\\\\.[0-9A-Fa-f]{16})"`
	// Number of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
	IqnNumber *int64 `json:"IqnNumber,omitempty"`
	// Prefix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
	IqnPrefix *string `json:"IqnPrefix,omitempty"`
	// Suffix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
	IqnSuffix            *string                                      `json:"IqnSuffix,omitempty"`
	Block                NullableIqnpoolBlockRelationship             `json:"Block,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Pool                 NullableIqnpoolPoolRelationship              `json:"Pool,omitempty"`
	PoolMember           NullableIqnpoolPoolMemberRelationship        `json:"PoolMember,omitempty"`
	Universe             NullableIqnpoolUniverseRelationship          `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolReservation IqnpoolReservation

// NewIqnpoolReservation instantiates a new IqnpoolReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolReservation(classId string, objectType string) *IqnpoolReservation {
	this := IqnpoolReservation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// NewIqnpoolReservationWithDefaults instantiates a new IqnpoolReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolReservationWithDefaults() *IqnpoolReservation {
	this := IqnpoolReservation{}
	var classId string = "iqnpool.Reservation"
	this.ClassId = classId
	var objectType string = "iqnpool.Reservation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolReservation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolReservation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolReservation) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iqnpool.Reservation" of the ClassId field.
func (o *IqnpoolReservation) GetDefaultClassId() interface{} {
	return "iqnpool.Reservation"
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolReservation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolReservation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolReservation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iqnpool.Reservation" of the ObjectType field.
func (o *IqnpoolReservation) GetDefaultObjectType() interface{} {
	return "iqnpool.Reservation"
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *IqnpoolReservation) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolReservation) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *IqnpoolReservation) SetIdentity(v string) {
	o.Identity = &v
}

// GetIqnNumber returns the IqnNumber field value if set, zero value otherwise.
func (o *IqnpoolReservation) GetIqnNumber() int64 {
	if o == nil || IsNil(o.IqnNumber) {
		var ret int64
		return ret
	}
	return *o.IqnNumber
}

// GetIqnNumberOk returns a tuple with the IqnNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolReservation) GetIqnNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.IqnNumber) {
		return nil, false
	}
	return o.IqnNumber, true
}

// HasIqnNumber returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasIqnNumber() bool {
	if o != nil && !IsNil(o.IqnNumber) {
		return true
	}

	return false
}

// SetIqnNumber gets a reference to the given int64 and assigns it to the IqnNumber field.
func (o *IqnpoolReservation) SetIqnNumber(v int64) {
	o.IqnNumber = &v
}

// GetIqnPrefix returns the IqnPrefix field value if set, zero value otherwise.
func (o *IqnpoolReservation) GetIqnPrefix() string {
	if o == nil || IsNil(o.IqnPrefix) {
		var ret string
		return ret
	}
	return *o.IqnPrefix
}

// GetIqnPrefixOk returns a tuple with the IqnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolReservation) GetIqnPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.IqnPrefix) {
		return nil, false
	}
	return o.IqnPrefix, true
}

// HasIqnPrefix returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasIqnPrefix() bool {
	if o != nil && !IsNil(o.IqnPrefix) {
		return true
	}

	return false
}

// SetIqnPrefix gets a reference to the given string and assigns it to the IqnPrefix field.
func (o *IqnpoolReservation) SetIqnPrefix(v string) {
	o.IqnPrefix = &v
}

// GetIqnSuffix returns the IqnSuffix field value if set, zero value otherwise.
func (o *IqnpoolReservation) GetIqnSuffix() string {
	if o == nil || IsNil(o.IqnSuffix) {
		var ret string
		return ret
	}
	return *o.IqnSuffix
}

// GetIqnSuffixOk returns a tuple with the IqnSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolReservation) GetIqnSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.IqnSuffix) {
		return nil, false
	}
	return o.IqnSuffix, true
}

// HasIqnSuffix returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasIqnSuffix() bool {
	if o != nil && !IsNil(o.IqnSuffix) {
		return true
	}

	return false
}

// SetIqnSuffix gets a reference to the given string and assigns it to the IqnSuffix field.
func (o *IqnpoolReservation) SetIqnSuffix(v string) {
	o.IqnSuffix = &v
}

// GetBlock returns the Block field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolReservation) GetBlock() IqnpoolBlockRelationship {
	if o == nil || IsNil(o.Block.Get()) {
		var ret IqnpoolBlockRelationship
		return ret
	}
	return *o.Block.Get()
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolReservation) GetBlockOk() (*IqnpoolBlockRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Block.Get(), o.Block.IsSet()
}

// HasBlock returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasBlock() bool {
	if o != nil && o.Block.IsSet() {
		return true
	}

	return false
}

// SetBlock gets a reference to the given NullableIqnpoolBlockRelationship and assigns it to the Block field.
func (o *IqnpoolReservation) SetBlock(v IqnpoolBlockRelationship) {
	o.Block.Set(&v)
}

// SetBlockNil sets the value for Block to be an explicit nil
func (o *IqnpoolReservation) SetBlockNil() {
	o.Block.Set(nil)
}

// UnsetBlock ensures that no value is present for Block, not even an explicit nil
func (o *IqnpoolReservation) UnsetBlock() {
	o.Block.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolReservation) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IqnpoolReservation) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *IqnpoolReservation) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *IqnpoolReservation) UnsetOrganization() {
	o.Organization.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolReservation) GetPool() IqnpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolReservation) GetPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableIqnpoolPoolRelationship and assigns it to the Pool field.
func (o *IqnpoolReservation) SetPool(v IqnpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *IqnpoolReservation) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *IqnpoolReservation) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolReservation) GetPoolMember() IqnpoolPoolMemberRelationship {
	if o == nil || IsNil(o.PoolMember.Get()) {
		var ret IqnpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember.Get()
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolReservation) GetPoolMemberOk() (*IqnpoolPoolMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolMember.Get(), o.PoolMember.IsSet()
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasPoolMember() bool {
	if o != nil && o.PoolMember.IsSet() {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given NullableIqnpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IqnpoolReservation) SetPoolMember(v IqnpoolPoolMemberRelationship) {
	o.PoolMember.Set(&v)
}

// SetPoolMemberNil sets the value for PoolMember to be an explicit nil
func (o *IqnpoolReservation) SetPoolMemberNil() {
	o.PoolMember.Set(nil)
}

// UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
func (o *IqnpoolReservation) UnsetPoolMember() {
	o.PoolMember.Unset()
}

// GetUniverse returns the Universe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolReservation) GetUniverse() IqnpoolUniverseRelationship {
	if o == nil || IsNil(o.Universe.Get()) {
		var ret IqnpoolUniverseRelationship
		return ret
	}
	return *o.Universe.Get()
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolReservation) GetUniverseOk() (*IqnpoolUniverseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Universe.Get(), o.Universe.IsSet()
}

// HasUniverse returns a boolean if a field has been set.
func (o *IqnpoolReservation) HasUniverse() bool {
	if o != nil && o.Universe.IsSet() {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given NullableIqnpoolUniverseRelationship and assigns it to the Universe field.
func (o *IqnpoolReservation) SetUniverse(v IqnpoolUniverseRelationship) {
	o.Universe.Set(&v)
}

// SetUniverseNil sets the value for Universe to be an explicit nil
func (o *IqnpoolReservation) SetUniverseNil() {
	o.Universe.Set(nil)
}

// UnsetUniverse ensures that no value is present for Universe, not even an explicit nil
func (o *IqnpoolReservation) UnsetUniverse() {
	o.Universe.Unset()
}

func (o IqnpoolReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IqnpoolReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservation, errPoolReservation := json.Marshal(o.PoolReservation)
	if errPoolReservation != nil {
		return map[string]interface{}{}, errPoolReservation
	}
	errPoolReservation = json.Unmarshal([]byte(serializedPoolReservation), &toSerialize)
	if errPoolReservation != nil {
		return map[string]interface{}{}, errPoolReservation
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Identity) {
		toSerialize["Identity"] = o.Identity
	}
	if !IsNil(o.IqnNumber) {
		toSerialize["IqnNumber"] = o.IqnNumber
	}
	if !IsNil(o.IqnPrefix) {
		toSerialize["IqnPrefix"] = o.IqnPrefix
	}
	if !IsNil(o.IqnSuffix) {
		toSerialize["IqnSuffix"] = o.IqnSuffix
	}
	if o.Block.IsSet() {
		toSerialize["Block"] = o.Block.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
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

func (o *IqnpoolReservation) UnmarshalJSON(data []byte) (err error) {
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
	type IqnpoolReservationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IQN identity to be reserved.
		Identity *string `json:"Identity,omitempty" validate:"regexp=^$|^(?:iqn\\\\.[0-9]{4}-[0-9]{2}(?:\\\\.[A-Za-z](?:[A-Za-z0-9\\\\-]*[A-Za-z0-9])?)+(?::.*)?|eui\\\\.[0-9A-Fa-f]{16})"`
		// Number of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
		IqnNumber *int64 `json:"IqnNumber,omitempty"`
		// Prefix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
		IqnPrefix *string `json:"IqnPrefix,omitempty"`
		// Suffix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
		IqnSuffix    *string                                      `json:"IqnSuffix,omitempty"`
		Block        NullableIqnpoolBlockRelationship             `json:"Block,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		Pool         NullableIqnpoolPoolRelationship              `json:"Pool,omitempty"`
		PoolMember   NullableIqnpoolPoolMemberRelationship        `json:"PoolMember,omitempty"`
		Universe     NullableIqnpoolUniverseRelationship          `json:"Universe,omitempty"`
	}

	varIqnpoolReservationWithoutEmbeddedStruct := IqnpoolReservationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIqnpoolReservationWithoutEmbeddedStruct)
	if err == nil {
		varIqnpoolReservation := _IqnpoolReservation{}
		varIqnpoolReservation.ClassId = varIqnpoolReservationWithoutEmbeddedStruct.ClassId
		varIqnpoolReservation.ObjectType = varIqnpoolReservationWithoutEmbeddedStruct.ObjectType
		varIqnpoolReservation.Identity = varIqnpoolReservationWithoutEmbeddedStruct.Identity
		varIqnpoolReservation.IqnNumber = varIqnpoolReservationWithoutEmbeddedStruct.IqnNumber
		varIqnpoolReservation.IqnPrefix = varIqnpoolReservationWithoutEmbeddedStruct.IqnPrefix
		varIqnpoolReservation.IqnSuffix = varIqnpoolReservationWithoutEmbeddedStruct.IqnSuffix
		varIqnpoolReservation.Block = varIqnpoolReservationWithoutEmbeddedStruct.Block
		varIqnpoolReservation.Organization = varIqnpoolReservationWithoutEmbeddedStruct.Organization
		varIqnpoolReservation.Pool = varIqnpoolReservationWithoutEmbeddedStruct.Pool
		varIqnpoolReservation.PoolMember = varIqnpoolReservationWithoutEmbeddedStruct.PoolMember
		varIqnpoolReservation.Universe = varIqnpoolReservationWithoutEmbeddedStruct.Universe
		*o = IqnpoolReservation(varIqnpoolReservation)
	} else {
		return err
	}

	varIqnpoolReservation := _IqnpoolReservation{}

	err = json.Unmarshal(data, &varIqnpoolReservation)
	if err == nil {
		o.PoolReservation = varIqnpoolReservation.PoolReservation
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "IqnNumber")
		delete(additionalProperties, "IqnPrefix")
		delete(additionalProperties, "IqnSuffix")
		delete(additionalProperties, "Block")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolReservation := reflect.ValueOf(o.PoolReservation)
		for i := 0; i < reflectPoolReservation.Type().NumField(); i++ {
			t := reflectPoolReservation.Type().Field(i)

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

type NullableIqnpoolReservation struct {
	value *IqnpoolReservation
	isSet bool
}

func (v NullableIqnpoolReservation) Get() *IqnpoolReservation {
	return v.value
}

func (v *NullableIqnpoolReservation) Set(val *IqnpoolReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolReservation(val *IqnpoolReservation) *NullableIqnpoolReservation {
	return &NullableIqnpoolReservation{value: val, isSet: true}
}

func (v NullableIqnpoolReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
