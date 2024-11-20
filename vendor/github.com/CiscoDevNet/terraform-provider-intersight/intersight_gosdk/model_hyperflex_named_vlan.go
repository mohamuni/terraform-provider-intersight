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

// checks if the HyperflexNamedVlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexNamedVlan{}

// HyperflexNamedVlan A VLAN with a name and ID. Named VLANs are used for defining the network and iSCSI external storage policies for the cluster.
type HyperflexNamedVlan struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the VLAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
	Name *string `json:"Name,omitempty" validate:"regexp=^$|^[a-zA-Z0-9-_.]{1,32}$"`
	// The ID of the named VLAN. An ID of 0 means the traffic is untagged. The ID can be any number between 0 and 4095, inclusive.
	VlanId               *int64 `json:"VlanId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNamedVlan HyperflexNamedVlan

// NewHyperflexNamedVlan instantiates a new HyperflexNamedVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNamedVlan(classId string, objectType string) *HyperflexNamedVlan {
	this := HyperflexNamedVlan{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexNamedVlanWithDefaults instantiates a new HyperflexNamedVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNamedVlanWithDefaults() *HyperflexNamedVlan {
	this := HyperflexNamedVlan{}
	var classId string = "hyperflex.NamedVlan"
	this.ClassId = classId
	var objectType string = "hyperflex.NamedVlan"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNamedVlan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVlan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNamedVlan) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.NamedVlan" of the ClassId field.
func (o *HyperflexNamedVlan) GetDefaultClassId() interface{} {
	return "hyperflex.NamedVlan"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNamedVlan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVlan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNamedVlan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.NamedVlan" of the ObjectType field.
func (o *HyperflexNamedVlan) GetDefaultObjectType() interface{} {
	return "hyperflex.NamedVlan"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexNamedVlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexNamedVlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexNamedVlan) SetName(v string) {
	o.Name = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *HyperflexNamedVlan) GetVlanId() int64 {
	if o == nil || IsNil(o.VlanId) {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVlan) GetVlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *HyperflexNamedVlan) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *HyperflexNamedVlan) SetVlanId(v int64) {
	o.VlanId = &v
}

func (o HyperflexNamedVlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexNamedVlan) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.VlanId) {
		toSerialize["VlanId"] = o.VlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexNamedVlan) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexNamedVlanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the VLAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
		Name *string `json:"Name,omitempty" validate:"regexp=^$|^[a-zA-Z0-9-_.]{1,32}$"`
		// The ID of the named VLAN. An ID of 0 means the traffic is untagged. The ID can be any number between 0 and 4095, inclusive.
		VlanId *int64 `json:"VlanId,omitempty"`
	}

	varHyperflexNamedVlanWithoutEmbeddedStruct := HyperflexNamedVlanWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexNamedVlanWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexNamedVlan := _HyperflexNamedVlan{}
		varHyperflexNamedVlan.ClassId = varHyperflexNamedVlanWithoutEmbeddedStruct.ClassId
		varHyperflexNamedVlan.ObjectType = varHyperflexNamedVlanWithoutEmbeddedStruct.ObjectType
		varHyperflexNamedVlan.Name = varHyperflexNamedVlanWithoutEmbeddedStruct.Name
		varHyperflexNamedVlan.VlanId = varHyperflexNamedVlanWithoutEmbeddedStruct.VlanId
		*o = HyperflexNamedVlan(varHyperflexNamedVlan)
	} else {
		return err
	}

	varHyperflexNamedVlan := _HyperflexNamedVlan{}

	err = json.Unmarshal(data, &varHyperflexNamedVlan)
	if err == nil {
		o.MoBaseComplexType = varHyperflexNamedVlan.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VlanId")

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

type NullableHyperflexNamedVlan struct {
	value *HyperflexNamedVlan
	isSet bool
}

func (v NullableHyperflexNamedVlan) Get() *HyperflexNamedVlan {
	return v.value
}

func (v *NullableHyperflexNamedVlan) Set(val *HyperflexNamedVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNamedVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNamedVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNamedVlan(val *HyperflexNamedVlan) *NullableHyperflexNamedVlan {
	return &NullableHyperflexNamedVlan{value: val, isSet: true}
}

func (v NullableHyperflexNamedVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNamedVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
