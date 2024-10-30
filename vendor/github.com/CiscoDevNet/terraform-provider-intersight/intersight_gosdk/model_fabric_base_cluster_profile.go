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

// checks if the FabricBaseClusterProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricBaseClusterProfile{}

// FabricBaseClusterProfile Abstract class of the Switch Cluster Profile and Template.
type FabricBaseClusterProfile struct {
	PolicyAbstractProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Number of switch profiles that are part of this cluster profile.
	SwitchProfilesCount  *int64 `json:"SwitchProfilesCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricBaseClusterProfile FabricBaseClusterProfile

// NewFabricBaseClusterProfile instantiates a new FabricBaseClusterProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricBaseClusterProfile(classId string, objectType string) *FabricBaseClusterProfile {
	this := FabricBaseClusterProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// NewFabricBaseClusterProfileWithDefaults instantiates a new FabricBaseClusterProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricBaseClusterProfileWithDefaults() *FabricBaseClusterProfile {
	this := FabricBaseClusterProfile{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricBaseClusterProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricBaseClusterProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricBaseClusterProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricBaseClusterProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricBaseClusterProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricBaseClusterProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSwitchProfilesCount returns the SwitchProfilesCount field value if set, zero value otherwise.
func (o *FabricBaseClusterProfile) GetSwitchProfilesCount() int64 {
	if o == nil || IsNil(o.SwitchProfilesCount) {
		var ret int64
		return ret
	}
	return *o.SwitchProfilesCount
}

// GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricBaseClusterProfile) GetSwitchProfilesCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SwitchProfilesCount) {
		return nil, false
	}
	return o.SwitchProfilesCount, true
}

// HasSwitchProfilesCount returns a boolean if a field has been set.
func (o *FabricBaseClusterProfile) HasSwitchProfilesCount() bool {
	if o != nil && !IsNil(o.SwitchProfilesCount) {
		return true
	}

	return false
}

// SetSwitchProfilesCount gets a reference to the given int64 and assigns it to the SwitchProfilesCount field.
func (o *FabricBaseClusterProfile) SetSwitchProfilesCount(v int64) {
	o.SwitchProfilesCount = &v
}

func (o FabricBaseClusterProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricBaseClusterProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractProfile
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.SwitchProfilesCount) {
		toSerialize["SwitchProfilesCount"] = o.SwitchProfilesCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricBaseClusterProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type FabricBaseClusterProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Number of switch profiles that are part of this cluster profile.
		SwitchProfilesCount *int64 `json:"SwitchProfilesCount,omitempty"`
	}

	varFabricBaseClusterProfileWithoutEmbeddedStruct := FabricBaseClusterProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricBaseClusterProfileWithoutEmbeddedStruct)
	if err == nil {
		varFabricBaseClusterProfile := _FabricBaseClusterProfile{}
		varFabricBaseClusterProfile.ClassId = varFabricBaseClusterProfileWithoutEmbeddedStruct.ClassId
		varFabricBaseClusterProfile.ObjectType = varFabricBaseClusterProfileWithoutEmbeddedStruct.ObjectType
		varFabricBaseClusterProfile.SwitchProfilesCount = varFabricBaseClusterProfileWithoutEmbeddedStruct.SwitchProfilesCount
		*o = FabricBaseClusterProfile(varFabricBaseClusterProfile)
	} else {
		return err
	}

	varFabricBaseClusterProfile := _FabricBaseClusterProfile{}

	err = json.Unmarshal(data, &varFabricBaseClusterProfile)
	if err == nil {
		o.PolicyAbstractProfile = varFabricBaseClusterProfile.PolicyAbstractProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SwitchProfilesCount")

		// remove fields from embedded structs
		reflectPolicyAbstractProfile := reflect.ValueOf(o.PolicyAbstractProfile)
		for i := 0; i < reflectPolicyAbstractProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractProfile.Type().Field(i)

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

type NullableFabricBaseClusterProfile struct {
	value *FabricBaseClusterProfile
	isSet bool
}

func (v NullableFabricBaseClusterProfile) Get() *FabricBaseClusterProfile {
	return v.value
}

func (v *NullableFabricBaseClusterProfile) Set(val *FabricBaseClusterProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricBaseClusterProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricBaseClusterProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricBaseClusterProfile(val *FabricBaseClusterProfile) *NullableFabricBaseClusterProfile {
	return &NullableFabricBaseClusterProfile{value: val, isSet: true}
}

func (v NullableFabricBaseClusterProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricBaseClusterProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
