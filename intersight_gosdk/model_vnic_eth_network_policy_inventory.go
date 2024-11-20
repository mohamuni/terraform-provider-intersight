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

// checks if the VnicEthNetworkPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicEthNetworkPolicyInventory{}

// VnicEthNetworkPolicyInventory An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.
type VnicEthNetworkPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	// Deprecated
	TargetPlatform       *string                      `json:"TargetPlatform,omitempty"`
	VlanSettings         NullableVnicVlanSettings     `json:"VlanSettings,omitempty"`
	TargetMo             NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthNetworkPolicyInventory VnicEthNetworkPolicyInventory

// NewVnicEthNetworkPolicyInventory instantiates a new VnicEthNetworkPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthNetworkPolicyInventory(classId string, objectType string) *VnicEthNetworkPolicyInventory {
	this := VnicEthNetworkPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicEthNetworkPolicyInventoryWithDefaults instantiates a new VnicEthNetworkPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthNetworkPolicyInventoryWithDefaults() *VnicEthNetworkPolicyInventory {
	this := VnicEthNetworkPolicyInventory{}
	var classId string = "vnic.EthNetworkPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.EthNetworkPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthNetworkPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthNetworkPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.EthNetworkPolicyInventory" of the ClassId field.
func (o *VnicEthNetworkPolicyInventory) GetDefaultClassId() interface{} {
	return "vnic.EthNetworkPolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthNetworkPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthNetworkPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.EthNetworkPolicyInventory" of the ObjectType field.
func (o *VnicEthNetworkPolicyInventory) GetDefaultObjectType() interface{} {
	return "vnic.EthNetworkPolicyInventory"
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
// Deprecated
func (o *VnicEthNetworkPolicyInventory) GetTargetPlatform() string {
	if o == nil || IsNil(o.TargetPlatform) {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VnicEthNetworkPolicyInventory) GetTargetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPlatform) {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyInventory) HasTargetPlatform() bool {
	if o != nil && !IsNil(o.TargetPlatform) {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
// Deprecated
func (o *VnicEthNetworkPolicyInventory) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetVlanSettings returns the VlanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthNetworkPolicyInventory) GetVlanSettings() VnicVlanSettings {
	if o == nil || IsNil(o.VlanSettings.Get()) {
		var ret VnicVlanSettings
		return ret
	}
	return *o.VlanSettings.Get()
}

// GetVlanSettingsOk returns a tuple with the VlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthNetworkPolicyInventory) GetVlanSettingsOk() (*VnicVlanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanSettings.Get(), o.VlanSettings.IsSet()
}

// HasVlanSettings returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyInventory) HasVlanSettings() bool {
	if o != nil && o.VlanSettings.IsSet() {
		return true
	}

	return false
}

// SetVlanSettings gets a reference to the given NullableVnicVlanSettings and assigns it to the VlanSettings field.
func (o *VnicEthNetworkPolicyInventory) SetVlanSettings(v VnicVlanSettings) {
	o.VlanSettings.Set(&v)
}

// SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil
func (o *VnicEthNetworkPolicyInventory) SetVlanSettingsNil() {
	o.VlanSettings.Set(nil)
}

// UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
func (o *VnicEthNetworkPolicyInventory) UnsetVlanSettings() {
	o.VlanSettings.Unset()
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthNetworkPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthNetworkPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicEthNetworkPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *VnicEthNetworkPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *VnicEthNetworkPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o VnicEthNetworkPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicEthNetworkPolicyInventory) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TargetPlatform) {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.VlanSettings.IsSet() {
		toSerialize["VlanSettings"] = o.VlanSettings.Get()
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicEthNetworkPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type VnicEthNetworkPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		// Deprecated
		TargetPlatform *string                      `json:"TargetPlatform,omitempty"`
		VlanSettings   NullableVnicVlanSettings     `json:"VlanSettings,omitempty"`
		TargetMo       NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct := VnicEthNetworkPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthNetworkPolicyInventory := _VnicEthNetworkPolicyInventory{}
		varVnicEthNetworkPolicyInventory.ClassId = varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicEthNetworkPolicyInventory.ObjectType = varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicEthNetworkPolicyInventory.TargetPlatform = varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct.TargetPlatform
		varVnicEthNetworkPolicyInventory.VlanSettings = varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct.VlanSettings
		varVnicEthNetworkPolicyInventory.TargetMo = varVnicEthNetworkPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicEthNetworkPolicyInventory(varVnicEthNetworkPolicyInventory)
	} else {
		return err
	}

	varVnicEthNetworkPolicyInventory := _VnicEthNetworkPolicyInventory{}

	err = json.Unmarshal(data, &varVnicEthNetworkPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicEthNetworkPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "VlanSettings")
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

type NullableVnicEthNetworkPolicyInventory struct {
	value *VnicEthNetworkPolicyInventory
	isSet bool
}

func (v NullableVnicEthNetworkPolicyInventory) Get() *VnicEthNetworkPolicyInventory {
	return v.value
}

func (v *NullableVnicEthNetworkPolicyInventory) Set(val *VnicEthNetworkPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthNetworkPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthNetworkPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthNetworkPolicyInventory(val *VnicEthNetworkPolicyInventory) *NullableVnicEthNetworkPolicyInventory {
	return &NullableVnicEthNetworkPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicEthNetworkPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthNetworkPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
