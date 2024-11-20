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

// checks if the EquipmentResetAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentResetAction{}

// EquipmentResetAction An action to initiate a reboot of the Fabric Interconnect with an option to perform fabric evacuation beforehand.
type EquipmentResetAction struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The reboot behavior for the Fabric Interconnect. Reboot - An action to reset the Fabric Interconnect by initiating its reboot. ForceReboot - Forces an immediate reboot of the Fabric Interconnect, overriding normal reboot validation checks. None - No reboot action should be triggered on the the Fabric Interconnect. * `None` - No action to be triggered on the Fabric Interconnect. * `Reboot` - An action to reset the Fabric Interconnect by initiating its reboot. * `ForceReboot` - An action to enforce an immediate reboot of the Fabric Interconnect regardless of existing validation checks.By default, a reset action is not allowed on an Fabric Interconnect if Domain Profile deployment, Manual Data Evacuation, or a reset action on the peer FI is already in progress. The force option allows users to override this check and perform the reset action on the FI.
	Action *string `json:"Action,omitempty"`
	// The flag to enable or disable fabric evacuation before rebooting the Fabric Interconnect.
	EnableFabricEvacuation *bool `json:"EnableFabricEvacuation,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _EquipmentResetAction EquipmentResetAction

// NewEquipmentResetAction instantiates a new EquipmentResetAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentResetAction(classId string, objectType string) *EquipmentResetAction {
	this := EquipmentResetAction{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	return &this
}

// NewEquipmentResetActionWithDefaults instantiates a new EquipmentResetAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentResetActionWithDefaults() *EquipmentResetAction {
	this := EquipmentResetAction{}
	var classId string = "equipment.ResetAction"
	this.ClassId = classId
	var objectType string = "equipment.ResetAction"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentResetAction) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentResetAction) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentResetAction) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.ResetAction" of the ClassId field.
func (o *EquipmentResetAction) GetDefaultClassId() interface{} {
	return "equipment.ResetAction"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentResetAction) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentResetAction) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentResetAction) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.ResetAction" of the ObjectType field.
func (o *EquipmentResetAction) GetDefaultObjectType() interface{} {
	return "equipment.ResetAction"
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EquipmentResetAction) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentResetAction) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EquipmentResetAction) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EquipmentResetAction) SetAction(v string) {
	o.Action = &v
}

// GetEnableFabricEvacuation returns the EnableFabricEvacuation field value if set, zero value otherwise.
func (o *EquipmentResetAction) GetEnableFabricEvacuation() bool {
	if o == nil || IsNil(o.EnableFabricEvacuation) {
		var ret bool
		return ret
	}
	return *o.EnableFabricEvacuation
}

// GetEnableFabricEvacuationOk returns a tuple with the EnableFabricEvacuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentResetAction) GetEnableFabricEvacuationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableFabricEvacuation) {
		return nil, false
	}
	return o.EnableFabricEvacuation, true
}

// HasEnableFabricEvacuation returns a boolean if a field has been set.
func (o *EquipmentResetAction) HasEnableFabricEvacuation() bool {
	if o != nil && !IsNil(o.EnableFabricEvacuation) {
		return true
	}

	return false
}

// SetEnableFabricEvacuation gets a reference to the given bool and assigns it to the EnableFabricEvacuation field.
func (o *EquipmentResetAction) SetEnableFabricEvacuation(v bool) {
	o.EnableFabricEvacuation = &v
}

func (o EquipmentResetAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentResetAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if !IsNil(o.EnableFabricEvacuation) {
		toSerialize["EnableFabricEvacuation"] = o.EnableFabricEvacuation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentResetAction) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentResetActionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The reboot behavior for the Fabric Interconnect. Reboot - An action to reset the Fabric Interconnect by initiating its reboot. ForceReboot - Forces an immediate reboot of the Fabric Interconnect, overriding normal reboot validation checks. None - No reboot action should be triggered on the the Fabric Interconnect. * `None` - No action to be triggered on the Fabric Interconnect. * `Reboot` - An action to reset the Fabric Interconnect by initiating its reboot. * `ForceReboot` - An action to enforce an immediate reboot of the Fabric Interconnect regardless of existing validation checks.By default, a reset action is not allowed on an Fabric Interconnect if Domain Profile deployment, Manual Data Evacuation, or a reset action on the peer FI is already in progress. The force option allows users to override this check and perform the reset action on the FI.
		Action *string `json:"Action,omitempty"`
		// The flag to enable or disable fabric evacuation before rebooting the Fabric Interconnect.
		EnableFabricEvacuation *bool `json:"EnableFabricEvacuation,omitempty"`
	}

	varEquipmentResetActionWithoutEmbeddedStruct := EquipmentResetActionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentResetActionWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentResetAction := _EquipmentResetAction{}
		varEquipmentResetAction.ClassId = varEquipmentResetActionWithoutEmbeddedStruct.ClassId
		varEquipmentResetAction.ObjectType = varEquipmentResetActionWithoutEmbeddedStruct.ObjectType
		varEquipmentResetAction.Action = varEquipmentResetActionWithoutEmbeddedStruct.Action
		varEquipmentResetAction.EnableFabricEvacuation = varEquipmentResetActionWithoutEmbeddedStruct.EnableFabricEvacuation
		*o = EquipmentResetAction(varEquipmentResetAction)
	} else {
		return err
	}

	varEquipmentResetAction := _EquipmentResetAction{}

	err = json.Unmarshal(data, &varEquipmentResetAction)
	if err == nil {
		o.MoBaseComplexType = varEquipmentResetAction.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "EnableFabricEvacuation")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableEquipmentResetAction struct {
	value *EquipmentResetAction
	isSet bool
}

func (v NullableEquipmentResetAction) Get() *EquipmentResetAction {
	return v.value
}

func (v *NullableEquipmentResetAction) Set(val *EquipmentResetAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentResetAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentResetAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentResetAction(val *EquipmentResetAction) *NullableEquipmentResetAction {
	return &NullableEquipmentResetAction{value: val, isSet: true}
}

func (v NullableEquipmentResetAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentResetAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
