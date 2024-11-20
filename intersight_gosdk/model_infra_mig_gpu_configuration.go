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

// checks if the InfraMigGpuConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraMigGpuConfiguration{}

// InfraMigGpuConfiguration Nvidia MIG capable GPU configuration on a compute resource (BM or VM).
type InfraMigGpuConfiguration struct {
	InfraBaseGpuConfiguration
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The predefined MIG profile name, e.g. 1g.5gb, 2g.10gb, etc.
	MigProfileName       *string `json:"MigProfileName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfraMigGpuConfiguration InfraMigGpuConfiguration

// NewInfraMigGpuConfiguration instantiates a new InfraMigGpuConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraMigGpuConfiguration(classId string, objectType string) *InfraMigGpuConfiguration {
	this := InfraMigGpuConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInfraMigGpuConfigurationWithDefaults instantiates a new InfraMigGpuConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraMigGpuConfigurationWithDefaults() *InfraMigGpuConfiguration {
	this := InfraMigGpuConfiguration{}
	var classId string = "infra.MigGpuConfiguration"
	this.ClassId = classId
	var objectType string = "infra.MigGpuConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InfraMigGpuConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InfraMigGpuConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InfraMigGpuConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "infra.MigGpuConfiguration" of the ClassId field.
func (o *InfraMigGpuConfiguration) GetDefaultClassId() interface{} {
	return "infra.MigGpuConfiguration"
}

// GetObjectType returns the ObjectType field value
func (o *InfraMigGpuConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InfraMigGpuConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InfraMigGpuConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "infra.MigGpuConfiguration" of the ObjectType field.
func (o *InfraMigGpuConfiguration) GetDefaultObjectType() interface{} {
	return "infra.MigGpuConfiguration"
}

// GetMigProfileName returns the MigProfileName field value if set, zero value otherwise.
func (o *InfraMigGpuConfiguration) GetMigProfileName() string {
	if o == nil || IsNil(o.MigProfileName) {
		var ret string
		return ret
	}
	return *o.MigProfileName
}

// GetMigProfileNameOk returns a tuple with the MigProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraMigGpuConfiguration) GetMigProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.MigProfileName) {
		return nil, false
	}
	return o.MigProfileName, true
}

// HasMigProfileName returns a boolean if a field has been set.
func (o *InfraMigGpuConfiguration) HasMigProfileName() bool {
	if o != nil && !IsNil(o.MigProfileName) {
		return true
	}

	return false
}

// SetMigProfileName gets a reference to the given string and assigns it to the MigProfileName field.
func (o *InfraMigGpuConfiguration) SetMigProfileName(v string) {
	o.MigProfileName = &v
}

func (o InfraMigGpuConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraMigGpuConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInfraBaseGpuConfiguration, errInfraBaseGpuConfiguration := json.Marshal(o.InfraBaseGpuConfiguration)
	if errInfraBaseGpuConfiguration != nil {
		return map[string]interface{}{}, errInfraBaseGpuConfiguration
	}
	errInfraBaseGpuConfiguration = json.Unmarshal([]byte(serializedInfraBaseGpuConfiguration), &toSerialize)
	if errInfraBaseGpuConfiguration != nil {
		return map[string]interface{}{}, errInfraBaseGpuConfiguration
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MigProfileName) {
		toSerialize["MigProfileName"] = o.MigProfileName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InfraMigGpuConfiguration) UnmarshalJSON(data []byte) (err error) {
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
	type InfraMigGpuConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The predefined MIG profile name, e.g. 1g.5gb, 2g.10gb, etc.
		MigProfileName *string `json:"MigProfileName,omitempty"`
	}

	varInfraMigGpuConfigurationWithoutEmbeddedStruct := InfraMigGpuConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varInfraMigGpuConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varInfraMigGpuConfiguration := _InfraMigGpuConfiguration{}
		varInfraMigGpuConfiguration.ClassId = varInfraMigGpuConfigurationWithoutEmbeddedStruct.ClassId
		varInfraMigGpuConfiguration.ObjectType = varInfraMigGpuConfigurationWithoutEmbeddedStruct.ObjectType
		varInfraMigGpuConfiguration.MigProfileName = varInfraMigGpuConfigurationWithoutEmbeddedStruct.MigProfileName
		*o = InfraMigGpuConfiguration(varInfraMigGpuConfiguration)
	} else {
		return err
	}

	varInfraMigGpuConfiguration := _InfraMigGpuConfiguration{}

	err = json.Unmarshal(data, &varInfraMigGpuConfiguration)
	if err == nil {
		o.InfraBaseGpuConfiguration = varInfraMigGpuConfiguration.InfraBaseGpuConfiguration
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MigProfileName")

		// remove fields from embedded structs
		reflectInfraBaseGpuConfiguration := reflect.ValueOf(o.InfraBaseGpuConfiguration)
		for i := 0; i < reflectInfraBaseGpuConfiguration.Type().NumField(); i++ {
			t := reflectInfraBaseGpuConfiguration.Type().Field(i)

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

type NullableInfraMigGpuConfiguration struct {
	value *InfraMigGpuConfiguration
	isSet bool
}

func (v NullableInfraMigGpuConfiguration) Get() *InfraMigGpuConfiguration {
	return v.value
}

func (v *NullableInfraMigGpuConfiguration) Set(val *InfraMigGpuConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraMigGpuConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraMigGpuConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraMigGpuConfiguration(val *InfraMigGpuConfiguration) *NullableInfraMigGpuConfiguration {
	return &NullableInfraMigGpuConfiguration{value: val, isSet: true}
}

func (v NullableInfraMigGpuConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraMigGpuConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
