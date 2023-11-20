/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ThermalPolicyAllOf Definition of the list of properties defined in 'thermal.Policy', excluding properties defined in parent classes.
type ThermalPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Sets the Fan Control Mode. High Power, Maximum Power and Acoustic modes are supported only on the Cisco UCS C-Series servers and on the X-Series Chassis. * `Balanced` - The fans run faster when needed based on the heat generated by the server. When possible, the fans returns to the minimum required speed. * `LowPower` - The Fans run at the minimum speed required to keep the server cool. * `HighPower` - The fans are kept at higher speed to emphasizes performance over power consumption. This Mode is only supported for UCS X series Chassis. * `MaximumPower` - The fans are always kept at maximum speed. This option provides the most cooling and consumes the most power. This Mode is only supported for UCS X series Chassis. * `Acoustic` - The fan speed is reduced to reduce noise levels in acoustic-sensitive environments. This Mode is only supported for UCS X series Chassis.
	FanControlMode *string                               `json:"FanControlMode,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThermalPolicyAllOf ThermalPolicyAllOf

// NewThermalPolicyAllOf instantiates a new ThermalPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThermalPolicyAllOf(classId string, objectType string) *ThermalPolicyAllOf {
	this := ThermalPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fanControlMode string = "Balanced"
	this.FanControlMode = &fanControlMode
	return &this
}

// NewThermalPolicyAllOfWithDefaults instantiates a new ThermalPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThermalPolicyAllOfWithDefaults() *ThermalPolicyAllOf {
	this := ThermalPolicyAllOf{}
	var classId string = "thermal.Policy"
	this.ClassId = classId
	var objectType string = "thermal.Policy"
	this.ObjectType = objectType
	var fanControlMode string = "Balanced"
	this.FanControlMode = &fanControlMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *ThermalPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ThermalPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ThermalPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ThermalPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ThermalPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ThermalPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFanControlMode returns the FanControlMode field value if set, zero value otherwise.
func (o *ThermalPolicyAllOf) GetFanControlMode() string {
	if o == nil || o.FanControlMode == nil {
		var ret string
		return ret
	}
	return *o.FanControlMode
}

// GetFanControlModeOk returns a tuple with the FanControlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThermalPolicyAllOf) GetFanControlModeOk() (*string, bool) {
	if o == nil || o.FanControlMode == nil {
		return nil, false
	}
	return o.FanControlMode, true
}

// HasFanControlMode returns a boolean if a field has been set.
func (o *ThermalPolicyAllOf) HasFanControlMode() bool {
	if o != nil && o.FanControlMode != nil {
		return true
	}

	return false
}

// SetFanControlMode gets a reference to the given string and assigns it to the FanControlMode field.
func (o *ThermalPolicyAllOf) SetFanControlMode(v string) {
	o.FanControlMode = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ThermalPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThermalPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ThermalPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ThermalPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThermalPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThermalPolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ThermalPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *ThermalPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o ThermalPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FanControlMode != nil {
		toSerialize["FanControlMode"] = o.FanControlMode
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ThermalPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varThermalPolicyAllOf := _ThermalPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varThermalPolicyAllOf); err == nil {
		*o = ThermalPolicyAllOf(varThermalPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FanControlMode")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThermalPolicyAllOf struct {
	value *ThermalPolicyAllOf
	isSet bool
}

func (v NullableThermalPolicyAllOf) Get() *ThermalPolicyAllOf {
	return v.value
}

func (v *NullableThermalPolicyAllOf) Set(val *ThermalPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableThermalPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableThermalPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThermalPolicyAllOf(val *ThermalPolicyAllOf) *NullableThermalPolicyAllOf {
	return &NullableThermalPolicyAllOf{value: val, isSet: true}
}

func (v NullableThermalPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThermalPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
