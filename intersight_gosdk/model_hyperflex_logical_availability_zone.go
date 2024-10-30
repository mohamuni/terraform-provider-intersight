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

// checks if the HyperflexLogicalAvailabilityZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexLogicalAvailabilityZone{}

// HyperflexLogicalAvailabilityZone A configuration for the Logical Availability Zone. Logical Availability Zones (LAZ) allow for increased fault tolerance by dividing clusters into logical partitions where a given block of data is only written to a zone once. This allows replications of data to be distributed evenly across zones. LAZ configurations are compatible with HyperFlex clusters meeting all of the following criteria: 1. The HyperFlex cluster must be running HyperFlex Data Platform 3.0 or higher. 2. The HyperFlex cluster must have 8 or more converged nodes.
type HyperflexLogicalAvailabilityZone struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or disable Logical Availability Zones (LAZ). If enabled, HyperFlex Data Platform automatically selects and groups nodes into different availability zones. For HyperFlex Data Platform versions prior to 3.0 release, this setting does not apply. For HyperFlex Data Platform versions 3.0 or higher, this setting is only applicable to HyperFlex systems with 8 or more converged nodes.
	AutoConfig           *bool `json:"AutoConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexLogicalAvailabilityZone HyperflexLogicalAvailabilityZone

// NewHyperflexLogicalAvailabilityZone instantiates a new HyperflexLogicalAvailabilityZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexLogicalAvailabilityZone(classId string, objectType string) *HyperflexLogicalAvailabilityZone {
	this := HyperflexLogicalAvailabilityZone{}
	this.ClassId = classId
	this.ObjectType = objectType
	var autoConfig bool = false
	this.AutoConfig = &autoConfig
	return &this
}

// NewHyperflexLogicalAvailabilityZoneWithDefaults instantiates a new HyperflexLogicalAvailabilityZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexLogicalAvailabilityZoneWithDefaults() *HyperflexLogicalAvailabilityZone {
	this := HyperflexLogicalAvailabilityZone{}
	var classId string = "hyperflex.LogicalAvailabilityZone"
	this.ClassId = classId
	var objectType string = "hyperflex.LogicalAvailabilityZone"
	this.ObjectType = objectType
	var autoConfig bool = false
	this.AutoConfig = &autoConfig
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexLogicalAvailabilityZone) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexLogicalAvailabilityZone) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexLogicalAvailabilityZone) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.LogicalAvailabilityZone" of the ClassId field.
func (o *HyperflexLogicalAvailabilityZone) GetDefaultClassId() interface{} {
	return "hyperflex.LogicalAvailabilityZone"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexLogicalAvailabilityZone) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexLogicalAvailabilityZone) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexLogicalAvailabilityZone) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.LogicalAvailabilityZone" of the ObjectType field.
func (o *HyperflexLogicalAvailabilityZone) GetDefaultObjectType() interface{} {
	return "hyperflex.LogicalAvailabilityZone"
}

// GetAutoConfig returns the AutoConfig field value if set, zero value otherwise.
func (o *HyperflexLogicalAvailabilityZone) GetAutoConfig() bool {
	if o == nil || IsNil(o.AutoConfig) {
		var ret bool
		return ret
	}
	return *o.AutoConfig
}

// GetAutoConfigOk returns a tuple with the AutoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLogicalAvailabilityZone) GetAutoConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConfig) {
		return nil, false
	}
	return o.AutoConfig, true
}

// HasAutoConfig returns a boolean if a field has been set.
func (o *HyperflexLogicalAvailabilityZone) HasAutoConfig() bool {
	if o != nil && !IsNil(o.AutoConfig) {
		return true
	}

	return false
}

// SetAutoConfig gets a reference to the given bool and assigns it to the AutoConfig field.
func (o *HyperflexLogicalAvailabilityZone) SetAutoConfig(v bool) {
	o.AutoConfig = &v
}

func (o HyperflexLogicalAvailabilityZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexLogicalAvailabilityZone) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AutoConfig) {
		toSerialize["AutoConfig"] = o.AutoConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexLogicalAvailabilityZone) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or disable Logical Availability Zones (LAZ). If enabled, HyperFlex Data Platform automatically selects and groups nodes into different availability zones. For HyperFlex Data Platform versions prior to 3.0 release, this setting does not apply. For HyperFlex Data Platform versions 3.0 or higher, this setting is only applicable to HyperFlex systems with 8 or more converged nodes.
		AutoConfig *bool `json:"AutoConfig,omitempty"`
	}

	varHyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct := HyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexLogicalAvailabilityZone := _HyperflexLogicalAvailabilityZone{}
		varHyperflexLogicalAvailabilityZone.ClassId = varHyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct.ClassId
		varHyperflexLogicalAvailabilityZone.ObjectType = varHyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct.ObjectType
		varHyperflexLogicalAvailabilityZone.AutoConfig = varHyperflexLogicalAvailabilityZoneWithoutEmbeddedStruct.AutoConfig
		*o = HyperflexLogicalAvailabilityZone(varHyperflexLogicalAvailabilityZone)
	} else {
		return err
	}

	varHyperflexLogicalAvailabilityZone := _HyperflexLogicalAvailabilityZone{}

	err = json.Unmarshal(data, &varHyperflexLogicalAvailabilityZone)
	if err == nil {
		o.MoBaseComplexType = varHyperflexLogicalAvailabilityZone.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoConfig")

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

type NullableHyperflexLogicalAvailabilityZone struct {
	value *HyperflexLogicalAvailabilityZone
	isSet bool
}

func (v NullableHyperflexLogicalAvailabilityZone) Get() *HyperflexLogicalAvailabilityZone {
	return v.value
}

func (v *NullableHyperflexLogicalAvailabilityZone) Set(val *HyperflexLogicalAvailabilityZone) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexLogicalAvailabilityZone) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexLogicalAvailabilityZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexLogicalAvailabilityZone(val *HyperflexLogicalAvailabilityZone) *NullableHyperflexLogicalAvailabilityZone {
	return &NullableHyperflexLogicalAvailabilityZone{value: val, isSet: true}
}

func (v NullableHyperflexLogicalAvailabilityZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexLogicalAvailabilityZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
