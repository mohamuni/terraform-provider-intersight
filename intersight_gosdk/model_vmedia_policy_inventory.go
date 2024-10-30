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

// checks if the VmediaPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmediaPolicyInventory{}

// VmediaPolicyInventory Policy to configure virtual media settings on endpoint.
type VmediaPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the Virtual Media service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// If enabled, allows encryption of all Virtual Media communications. Please note that this can no longer be disabled for servers running versions 4.2 and above.
	Encryption *bool `json:"Encryption,omitempty"`
	// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
	LowPowerUsb          *bool                        `json:"LowPowerUsb,omitempty"`
	Mappings             []VmediaMapping              `json:"Mappings,omitempty"`
	TargetMo             NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VmediaPolicyInventory VmediaPolicyInventory

// NewVmediaPolicyInventory instantiates a new VmediaPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmediaPolicyInventory(classId string, objectType string) *VmediaPolicyInventory {
	this := VmediaPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVmediaPolicyInventoryWithDefaults instantiates a new VmediaPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmediaPolicyInventoryWithDefaults() *VmediaPolicyInventory {
	this := VmediaPolicyInventory{}
	var classId string = "vmedia.PolicyInventory"
	this.ClassId = classId
	var objectType string = "vmedia.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VmediaPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VmediaPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VmediaPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vmedia.PolicyInventory" of the ClassId field.
func (o *VmediaPolicyInventory) GetDefaultClassId() interface{} {
	return "vmedia.PolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *VmediaPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VmediaPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VmediaPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vmedia.PolicyInventory" of the ObjectType field.
func (o *VmediaPolicyInventory) GetDefaultObjectType() interface{} {
	return "vmedia.PolicyInventory"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VmediaPolicyInventory) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicyInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VmediaPolicyInventory) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VmediaPolicyInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *VmediaPolicyInventory) GetEncryption() bool {
	if o == nil || IsNil(o.Encryption) {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicyInventory) GetEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *VmediaPolicyInventory) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *VmediaPolicyInventory) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetLowPowerUsb returns the LowPowerUsb field value if set, zero value otherwise.
func (o *VmediaPolicyInventory) GetLowPowerUsb() bool {
	if o == nil || IsNil(o.LowPowerUsb) {
		var ret bool
		return ret
	}
	return *o.LowPowerUsb
}

// GetLowPowerUsbOk returns a tuple with the LowPowerUsb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicyInventory) GetLowPowerUsbOk() (*bool, bool) {
	if o == nil || IsNil(o.LowPowerUsb) {
		return nil, false
	}
	return o.LowPowerUsb, true
}

// HasLowPowerUsb returns a boolean if a field has been set.
func (o *VmediaPolicyInventory) HasLowPowerUsb() bool {
	if o != nil && !IsNil(o.LowPowerUsb) {
		return true
	}

	return false
}

// SetLowPowerUsb gets a reference to the given bool and assigns it to the LowPowerUsb field.
func (o *VmediaPolicyInventory) SetLowPowerUsb(v bool) {
	o.LowPowerUsb = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmediaPolicyInventory) GetMappings() []VmediaMapping {
	if o == nil {
		var ret []VmediaMapping
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmediaPolicyInventory) GetMappingsOk() ([]VmediaMapping, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *VmediaPolicyInventory) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []VmediaMapping and assigns it to the Mappings field.
func (o *VmediaPolicyInventory) SetMappings(v []VmediaMapping) {
	o.Mappings = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmediaPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmediaPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VmediaPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VmediaPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *VmediaPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *VmediaPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o VmediaPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmediaPolicyInventory) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Encryption) {
		toSerialize["Encryption"] = o.Encryption
	}
	if !IsNil(o.LowPowerUsb) {
		toSerialize["LowPowerUsb"] = o.LowPowerUsb
	}
	if o.Mappings != nil {
		toSerialize["Mappings"] = o.Mappings
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VmediaPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type VmediaPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of the Virtual Media service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// If enabled, allows encryption of all Virtual Media communications. Please note that this can no longer be disabled for servers running versions 4.2 and above.
		Encryption *bool `json:"Encryption,omitempty"`
		// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
		LowPowerUsb *bool                        `json:"LowPowerUsb,omitempty"`
		Mappings    []VmediaMapping              `json:"Mappings,omitempty"`
		TargetMo    NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVmediaPolicyInventoryWithoutEmbeddedStruct := VmediaPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVmediaPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVmediaPolicyInventory := _VmediaPolicyInventory{}
		varVmediaPolicyInventory.ClassId = varVmediaPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVmediaPolicyInventory.ObjectType = varVmediaPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVmediaPolicyInventory.Enabled = varVmediaPolicyInventoryWithoutEmbeddedStruct.Enabled
		varVmediaPolicyInventory.Encryption = varVmediaPolicyInventoryWithoutEmbeddedStruct.Encryption
		varVmediaPolicyInventory.LowPowerUsb = varVmediaPolicyInventoryWithoutEmbeddedStruct.LowPowerUsb
		varVmediaPolicyInventory.Mappings = varVmediaPolicyInventoryWithoutEmbeddedStruct.Mappings
		varVmediaPolicyInventory.TargetMo = varVmediaPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VmediaPolicyInventory(varVmediaPolicyInventory)
	} else {
		return err
	}

	varVmediaPolicyInventory := _VmediaPolicyInventory{}

	err = json.Unmarshal(data, &varVmediaPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVmediaPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "LowPowerUsb")
		delete(additionalProperties, "Mappings")
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

type NullableVmediaPolicyInventory struct {
	value *VmediaPolicyInventory
	isSet bool
}

func (v NullableVmediaPolicyInventory) Get() *VmediaPolicyInventory {
	return v.value
}

func (v *NullableVmediaPolicyInventory) Set(val *VmediaPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVmediaPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVmediaPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmediaPolicyInventory(val *VmediaPolicyInventory) *NullableVmediaPolicyInventory {
	return &NullableVmediaPolicyInventory{value: val, isSet: true}
}

func (v NullableVmediaPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmediaPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
