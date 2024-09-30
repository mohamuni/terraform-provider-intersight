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

// checks if the IpmioverlanPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpmioverlanPolicyInventory{}

// IpmioverlanPolicyInventory Intelligent Platform Management Interface Over LAN Policy.
type IpmioverlanPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the IPMI Over LAN service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. Use “00” to disable encryption key use. This configuration is supported by all Standalone C-Series servers. FI-attached C-Series servers with firmware at minimum of 4.2.3a support this configuration. B/X-Series servers with firmware at minimum of 5.1.0.x support this configuration. IPMI commands using this key should append zeroes to the key to achieve a length of 40 characters.
	EncryptionKey *string `json:"EncryptionKey,omitempty" validate:"regexp=^[a-fA-F0-9]*$"`
	// Indicates whether the value of the 'encryptionKey' property has been set.
	IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`
	// The highest privilege level that can be assigned to an IPMI session on a server. This configuration is supported by all Standalone C-Series servers. FI-attached C-Series servers with firmware at minimum of 4.2.3a support this configuration. B/X-Series servers with firmware at minimum of 5.1.0.x support this configuration. Privilege level user is not supported for B/X-Series servers. * `admin` - Privilege to perform all actions available through IPMI. * `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks. * `read-only` - Privilege to view information throught IPMI but restriction on making any changes.
	Privilege            *string                      `json:"Privilege,omitempty"`
	TargetMo             NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IpmioverlanPolicyInventory IpmioverlanPolicyInventory

// NewIpmioverlanPolicyInventory instantiates a new IpmioverlanPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmioverlanPolicyInventory(classId string, objectType string) *IpmioverlanPolicyInventory {
	this := IpmioverlanPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIpmioverlanPolicyInventoryWithDefaults instantiates a new IpmioverlanPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmioverlanPolicyInventoryWithDefaults() *IpmioverlanPolicyInventory {
	this := IpmioverlanPolicyInventory{}
	var classId string = "ipmioverlan.PolicyInventory"
	this.ClassId = classId
	var objectType string = "ipmioverlan.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IpmioverlanPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IpmioverlanPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ipmioverlan.PolicyInventory" of the ClassId field.
func (o *IpmioverlanPolicyInventory) GetDefaultClassId() interface{} {
	return "ipmioverlan.PolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *IpmioverlanPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IpmioverlanPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ipmioverlan.PolicyInventory" of the ObjectType field.
func (o *IpmioverlanPolicyInventory) GetDefaultObjectType() interface{} {
	return "ipmioverlan.PolicyInventory"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IpmioverlanPolicyInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey) {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasEncryptionKey() bool {
	if o != nil && !IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *IpmioverlanPolicyInventory) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetIsEncryptionKeySet returns the IsEncryptionKeySet field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetIsEncryptionKeySet() bool {
	if o == nil || IsNil(o.IsEncryptionKeySet) {
		var ret bool
		return ret
	}
	return *o.IsEncryptionKeySet
}

// GetIsEncryptionKeySetOk returns a tuple with the IsEncryptionKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetIsEncryptionKeySetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncryptionKeySet) {
		return nil, false
	}
	return o.IsEncryptionKeySet, true
}

// HasIsEncryptionKeySet returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasIsEncryptionKeySet() bool {
	if o != nil && !IsNil(o.IsEncryptionKeySet) {
		return true
	}

	return false
}

// SetIsEncryptionKeySet gets a reference to the given bool and assigns it to the IsEncryptionKeySet field.
func (o *IpmioverlanPolicyInventory) SetIsEncryptionKeySet(v bool) {
	o.IsEncryptionKeySet = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetPrivilege() string {
	if o == nil || IsNil(o.Privilege) {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetPrivilegeOk() (*string, bool) {
	if o == nil || IsNil(o.Privilege) {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasPrivilege() bool {
	if o != nil && !IsNil(o.Privilege) {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *IpmioverlanPolicyInventory) SetPrivilege(v string) {
	o.Privilege = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmioverlanPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmioverlanPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IpmioverlanPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *IpmioverlanPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *IpmioverlanPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o IpmioverlanPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpmioverlanPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.EncryptionKey) {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if !IsNil(o.IsEncryptionKeySet) {
		toSerialize["IsEncryptionKeySet"] = o.IsEncryptionKeySet
	}
	if !IsNil(o.Privilege) {
		toSerialize["Privilege"] = o.Privilege
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IpmioverlanPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type IpmioverlanPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of the IPMI Over LAN service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. Use “00” to disable encryption key use. This configuration is supported by all Standalone C-Series servers. FI-attached C-Series servers with firmware at minimum of 4.2.3a support this configuration. B/X-Series servers with firmware at minimum of 5.1.0.x support this configuration. IPMI commands using this key should append zeroes to the key to achieve a length of 40 characters.
		EncryptionKey *string `json:"EncryptionKey,omitempty" validate:"regexp=^[a-fA-F0-9]*$"`
		// Indicates whether the value of the 'encryptionKey' property has been set.
		IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`
		// The highest privilege level that can be assigned to an IPMI session on a server. This configuration is supported by all Standalone C-Series servers. FI-attached C-Series servers with firmware at minimum of 4.2.3a support this configuration. B/X-Series servers with firmware at minimum of 5.1.0.x support this configuration. Privilege level user is not supported for B/X-Series servers. * `admin` - Privilege to perform all actions available through IPMI. * `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks. * `read-only` - Privilege to view information throught IPMI but restriction on making any changes.
		Privilege *string                      `json:"Privilege,omitempty"`
		TargetMo  NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varIpmioverlanPolicyInventoryWithoutEmbeddedStruct := IpmioverlanPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIpmioverlanPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varIpmioverlanPolicyInventory := _IpmioverlanPolicyInventory{}
		varIpmioverlanPolicyInventory.ClassId = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.ClassId
		varIpmioverlanPolicyInventory.ObjectType = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varIpmioverlanPolicyInventory.Enabled = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.Enabled
		varIpmioverlanPolicyInventory.EncryptionKey = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.EncryptionKey
		varIpmioverlanPolicyInventory.IsEncryptionKeySet = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.IsEncryptionKeySet
		varIpmioverlanPolicyInventory.Privilege = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.Privilege
		varIpmioverlanPolicyInventory.TargetMo = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = IpmioverlanPolicyInventory(varIpmioverlanPolicyInventory)
	} else {
		return err
	}

	varIpmioverlanPolicyInventory := _IpmioverlanPolicyInventory{}

	err = json.Unmarshal(data, &varIpmioverlanPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varIpmioverlanPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "EncryptionKey")
		delete(additionalProperties, "IsEncryptionKeySet")
		delete(additionalProperties, "Privilege")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableIpmioverlanPolicyInventory struct {
	value *IpmioverlanPolicyInventory
	isSet bool
}

func (v NullableIpmioverlanPolicyInventory) Get() *IpmioverlanPolicyInventory {
	return v.value
}

func (v *NullableIpmioverlanPolicyInventory) Set(val *IpmioverlanPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmioverlanPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmioverlanPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmioverlanPolicyInventory(val *IpmioverlanPolicyInventory) *NullableIpmioverlanPolicyInventory {
	return &NullableIpmioverlanPolicyInventory{value: val, isSet: true}
}

func (v NullableIpmioverlanPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmioverlanPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
