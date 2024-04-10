/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ThermalPolicyInventory Thermal Management policy models a configuration that can be applied to Chassis or Server to manage Thermal Features.
type ThermalPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Sets the Aggressive Cooling to Enable/Disable. * `Disabled` - Set the value to Disabled. * `Enabled` - Set the value to Enabled.
	AggressiveCooling *string `json:"AggressiveCooling,omitempty"`
	// Sets the Fan Control Mode. High Power, Maximum Power and Acoustic modes are supported only on the Cisco UCS C-Series servers and on the X-Series Chassis. * `Balanced` - The fans run faster when needed based on the heat generated by the server. When possible, the fans returns to the minimum required speed. * `LowPower` - The Fans run at the minimum speed required to keep the server cool. * `HighPower` - The fans are kept at higher speed to emphasizes performance over power consumption. This Mode is only supported for UCS X series Chassis. * `MaximumPower` - The fans are always kept at maximum speed. This option provides the most cooling and consumes the most power. This Mode is only supported for UCS X series Chassis. * `Acoustic` - The fan speed is reduced to reduce noise levels in acoustic-sensitive environments. This Mode is only supported for UCS X series Chassis.
	FanControlMode       *string               `json:"FanControlMode,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThermalPolicyInventory ThermalPolicyInventory

// NewThermalPolicyInventory instantiates a new ThermalPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThermalPolicyInventory(classId string, objectType string) *ThermalPolicyInventory {
	this := ThermalPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewThermalPolicyInventoryWithDefaults instantiates a new ThermalPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThermalPolicyInventoryWithDefaults() *ThermalPolicyInventory {
	this := ThermalPolicyInventory{}
	var classId string = "thermal.PolicyInventory"
	this.ClassId = classId
	var objectType string = "thermal.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ThermalPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ThermalPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ThermalPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ThermalPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ThermalPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ThermalPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggressiveCooling returns the AggressiveCooling field value if set, zero value otherwise.
func (o *ThermalPolicyInventory) GetAggressiveCooling() string {
	if o == nil || o.AggressiveCooling == nil {
		var ret string
		return ret
	}
	return *o.AggressiveCooling
}

// GetAggressiveCoolingOk returns a tuple with the AggressiveCooling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThermalPolicyInventory) GetAggressiveCoolingOk() (*string, bool) {
	if o == nil || o.AggressiveCooling == nil {
		return nil, false
	}
	return o.AggressiveCooling, true
}

// HasAggressiveCooling returns a boolean if a field has been set.
func (o *ThermalPolicyInventory) HasAggressiveCooling() bool {
	if o != nil && o.AggressiveCooling != nil {
		return true
	}

	return false
}

// SetAggressiveCooling gets a reference to the given string and assigns it to the AggressiveCooling field.
func (o *ThermalPolicyInventory) SetAggressiveCooling(v string) {
	o.AggressiveCooling = &v
}

// GetFanControlMode returns the FanControlMode field value if set, zero value otherwise.
func (o *ThermalPolicyInventory) GetFanControlMode() string {
	if o == nil || o.FanControlMode == nil {
		var ret string
		return ret
	}
	return *o.FanControlMode
}

// GetFanControlModeOk returns a tuple with the FanControlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThermalPolicyInventory) GetFanControlModeOk() (*string, bool) {
	if o == nil || o.FanControlMode == nil {
		return nil, false
	}
	return o.FanControlMode, true
}

// HasFanControlMode returns a boolean if a field has been set.
func (o *ThermalPolicyInventory) HasFanControlMode() bool {
	if o != nil && o.FanControlMode != nil {
		return true
	}

	return false
}

// SetFanControlMode gets a reference to the given string and assigns it to the FanControlMode field.
func (o *ThermalPolicyInventory) SetFanControlMode(v string) {
	o.FanControlMode = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *ThermalPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThermalPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *ThermalPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *ThermalPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o ThermalPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggressiveCooling != nil {
		toSerialize["AggressiveCooling"] = o.AggressiveCooling
	}
	if o.FanControlMode != nil {
		toSerialize["FanControlMode"] = o.FanControlMode
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ThermalPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type ThermalPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Sets the Aggressive Cooling to Enable/Disable. * `Disabled` - Set the value to Disabled. * `Enabled` - Set the value to Enabled.
		AggressiveCooling *string `json:"AggressiveCooling,omitempty"`
		// Sets the Fan Control Mode. High Power, Maximum Power and Acoustic modes are supported only on the Cisco UCS C-Series servers and on the X-Series Chassis. * `Balanced` - The fans run faster when needed based on the heat generated by the server. When possible, the fans returns to the minimum required speed. * `LowPower` - The Fans run at the minimum speed required to keep the server cool. * `HighPower` - The fans are kept at higher speed to emphasizes performance over power consumption. This Mode is only supported for UCS X series Chassis. * `MaximumPower` - The fans are always kept at maximum speed. This option provides the most cooling and consumes the most power. This Mode is only supported for UCS X series Chassis. * `Acoustic` - The fan speed is reduced to reduce noise levels in acoustic-sensitive environments. This Mode is only supported for UCS X series Chassis.
		FanControlMode *string               `json:"FanControlMode,omitempty"`
		TargetMo       *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varThermalPolicyInventoryWithoutEmbeddedStruct := ThermalPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varThermalPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varThermalPolicyInventory := _ThermalPolicyInventory{}
		varThermalPolicyInventory.ClassId = varThermalPolicyInventoryWithoutEmbeddedStruct.ClassId
		varThermalPolicyInventory.ObjectType = varThermalPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varThermalPolicyInventory.AggressiveCooling = varThermalPolicyInventoryWithoutEmbeddedStruct.AggressiveCooling
		varThermalPolicyInventory.FanControlMode = varThermalPolicyInventoryWithoutEmbeddedStruct.FanControlMode
		varThermalPolicyInventory.TargetMo = varThermalPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = ThermalPolicyInventory(varThermalPolicyInventory)
	} else {
		return err
	}

	varThermalPolicyInventory := _ThermalPolicyInventory{}

	err = json.Unmarshal(bytes, &varThermalPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varThermalPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggressiveCooling")
		delete(additionalProperties, "FanControlMode")
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

type NullableThermalPolicyInventory struct {
	value *ThermalPolicyInventory
	isSet bool
}

func (v NullableThermalPolicyInventory) Get() *ThermalPolicyInventory {
	return v.value
}

func (v *NullableThermalPolicyInventory) Set(val *ThermalPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableThermalPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableThermalPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThermalPolicyInventory(val *ThermalPolicyInventory) *NullableThermalPolicyInventory {
	return &NullableThermalPolicyInventory{value: val, isSet: true}
}

func (v NullableThermalPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThermalPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}