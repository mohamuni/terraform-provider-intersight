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

// checks if the ThermalPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThermalPolicy{}

// ThermalPolicy Thermal Management policy models a configuration that can be applied to Chassis or Server to manage Thermal Features.
type ThermalPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Sets the Fan Control Mode. High Power, Maximum Power and Acoustic modes are supported only on the Cisco UCS C-Series servers and on the X-Series Chassis. * `Balanced` - The fans run faster when needed based on the heat generated by the server. When possible, the fans returns to the minimum required speed. * `LowPower` - The Fans run at the minimum speed required to keep the server cool. * `HighPower` - The fans are kept at higher speed to emphasizes performance over power consumption. This Mode is only supported for UCS X series Chassis. * `MaximumPower` - The fans are always kept at maximum speed. This option provides the most cooling and consumes the most power. This Mode is only supported for UCS X series Chassis. * `Acoustic` - The fan speed is reduced to reduce noise levels in acoustic-sensitive environments. This Mode is only supported for UCS X series Chassis.
	FanControlMode *string                                      `json:"FanControlMode,omitempty"`
	Organization   NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThermalPolicy ThermalPolicy

// NewThermalPolicy instantiates a new ThermalPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThermalPolicy(classId string, objectType string) *ThermalPolicy {
	this := ThermalPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fanControlMode string = "Balanced"
	this.FanControlMode = &fanControlMode
	return &this
}

// NewThermalPolicyWithDefaults instantiates a new ThermalPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThermalPolicyWithDefaults() *ThermalPolicy {
	this := ThermalPolicy{}
	var classId string = "thermal.Policy"
	this.ClassId = classId
	var objectType string = "thermal.Policy"
	this.ObjectType = objectType
	var fanControlMode string = "Balanced"
	this.FanControlMode = &fanControlMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *ThermalPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ThermalPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ThermalPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "thermal.Policy" of the ClassId field.
func (o *ThermalPolicy) GetDefaultClassId() interface{} {
	return "thermal.Policy"
}

// GetObjectType returns the ObjectType field value
func (o *ThermalPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ThermalPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ThermalPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "thermal.Policy" of the ObjectType field.
func (o *ThermalPolicy) GetDefaultObjectType() interface{} {
	return "thermal.Policy"
}

// GetFanControlMode returns the FanControlMode field value if set, zero value otherwise.
func (o *ThermalPolicy) GetFanControlMode() string {
	if o == nil || IsNil(o.FanControlMode) {
		var ret string
		return ret
	}
	return *o.FanControlMode
}

// GetFanControlModeOk returns a tuple with the FanControlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThermalPolicy) GetFanControlModeOk() (*string, bool) {
	if o == nil || IsNil(o.FanControlMode) {
		return nil, false
	}
	return o.FanControlMode, true
}

// HasFanControlMode returns a boolean if a field has been set.
func (o *ThermalPolicy) HasFanControlMode() bool {
	if o != nil && !IsNil(o.FanControlMode) {
		return true
	}

	return false
}

// SetFanControlMode gets a reference to the given string and assigns it to the FanControlMode field.
func (o *ThermalPolicy) SetFanControlMode(v string) {
	o.FanControlMode = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThermalPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThermalPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *ThermalPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ThermalPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *ThermalPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *ThermalPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThermalPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThermalPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ThermalPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *ThermalPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o ThermalPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThermalPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FanControlMode) {
		toSerialize["FanControlMode"] = o.FanControlMode
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThermalPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type ThermalPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Sets the Fan Control Mode. High Power, Maximum Power and Acoustic modes are supported only on the Cisco UCS C-Series servers and on the X-Series Chassis. * `Balanced` - The fans run faster when needed based on the heat generated by the server. When possible, the fans returns to the minimum required speed. * `LowPower` - The Fans run at the minimum speed required to keep the server cool. * `HighPower` - The fans are kept at higher speed to emphasizes performance over power consumption. This Mode is only supported for UCS X series Chassis. * `MaximumPower` - The fans are always kept at maximum speed. This option provides the most cooling and consumes the most power. This Mode is only supported for UCS X series Chassis. * `Acoustic` - The fan speed is reduced to reduce noise levels in acoustic-sensitive environments. This Mode is only supported for UCS X series Chassis.
		FanControlMode *string                                      `json:"FanControlMode,omitempty"`
		Organization   NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varThermalPolicyWithoutEmbeddedStruct := ThermalPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varThermalPolicyWithoutEmbeddedStruct)
	if err == nil {
		varThermalPolicy := _ThermalPolicy{}
		varThermalPolicy.ClassId = varThermalPolicyWithoutEmbeddedStruct.ClassId
		varThermalPolicy.ObjectType = varThermalPolicyWithoutEmbeddedStruct.ObjectType
		varThermalPolicy.FanControlMode = varThermalPolicyWithoutEmbeddedStruct.FanControlMode
		varThermalPolicy.Organization = varThermalPolicyWithoutEmbeddedStruct.Organization
		varThermalPolicy.Profiles = varThermalPolicyWithoutEmbeddedStruct.Profiles
		*o = ThermalPolicy(varThermalPolicy)
	} else {
		return err
	}

	varThermalPolicy := _ThermalPolicy{}

	err = json.Unmarshal(data, &varThermalPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varThermalPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FanControlMode")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableThermalPolicy struct {
	value *ThermalPolicy
	isSet bool
}

func (v NullableThermalPolicy) Get() *ThermalPolicy {
	return v.value
}

func (v *NullableThermalPolicy) Set(val *ThermalPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableThermalPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableThermalPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThermalPolicy(val *ThermalPolicy) *NullableThermalPolicy {
	return &NullableThermalPolicy{value: val, isSet: true}
}

func (v NullableThermalPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThermalPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
