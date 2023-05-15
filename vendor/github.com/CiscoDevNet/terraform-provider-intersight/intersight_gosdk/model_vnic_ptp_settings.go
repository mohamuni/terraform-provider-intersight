/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VnicPtpSettings Settings for Precision Time Protocol which can provide precise time of day information and time-stampted inputs, as well as scheduled and/or synchronized outputs for a variety of systems.
type VnicPtpSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of Precision Time Protocol (PTP) on the virtual ethernet interface. PTP can be enabled only on one vNIC on an adapter.
	Enabled              *bool `json:"Enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicPtpSettings VnicPtpSettings

// NewVnicPtpSettings instantiates a new VnicPtpSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicPtpSettings(classId string, objectType string) *VnicPtpSettings {
	this := VnicPtpSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewVnicPtpSettingsWithDefaults instantiates a new VnicPtpSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicPtpSettingsWithDefaults() *VnicPtpSettings {
	this := VnicPtpSettings{}
	var classId string = "vnic.PtpSettings"
	this.ClassId = classId
	var objectType string = "vnic.PtpSettings"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicPtpSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicPtpSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicPtpSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicPtpSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicPtpSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicPtpSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicPtpSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPtpSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicPtpSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicPtpSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o VnicPtpSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicPtpSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicPtpSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of Precision Time Protocol (PTP) on the virtual ethernet interface. PTP can be enabled only on one vNIC on an adapter.
		Enabled *bool `json:"Enabled,omitempty"`
	}

	varVnicPtpSettingsWithoutEmbeddedStruct := VnicPtpSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicPtpSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicPtpSettings := _VnicPtpSettings{}
		varVnicPtpSettings.ClassId = varVnicPtpSettingsWithoutEmbeddedStruct.ClassId
		varVnicPtpSettings.ObjectType = varVnicPtpSettingsWithoutEmbeddedStruct.ObjectType
		varVnicPtpSettings.Enabled = varVnicPtpSettingsWithoutEmbeddedStruct.Enabled
		*o = VnicPtpSettings(varVnicPtpSettings)
	} else {
		return err
	}

	varVnicPtpSettings := _VnicPtpSettings{}

	err = json.Unmarshal(bytes, &varVnicPtpSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicPtpSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")

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

type NullableVnicPtpSettings struct {
	value *VnicPtpSettings
	isSet bool
}

func (v NullableVnicPtpSettings) Get() *VnicPtpSettings {
	return v.value
}

func (v *NullableVnicPtpSettings) Set(val *VnicPtpSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicPtpSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicPtpSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicPtpSettings(val *VnicPtpSettings) *NullableVnicPtpSettings {
	return &NullableVnicPtpSettings{value: val, isSet: true}
}

func (v NullableVnicPtpSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicPtpSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}