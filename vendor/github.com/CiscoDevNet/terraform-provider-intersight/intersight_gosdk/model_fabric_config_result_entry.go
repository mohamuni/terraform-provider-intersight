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

// checks if the FabricConfigResultEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricConfigResultEntry{}

// FabricConfigResultEntry This provides detailed information for the deploy and validation profile configuration results.
type FabricConfigResultEntry struct {
	PolicyAbstractConfigResultEntry
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                 `json:"ObjectType"`
	ConfigResult         NullableFabricConfigResultRelationship `json:"ConfigResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricConfigResultEntry FabricConfigResultEntry

// NewFabricConfigResultEntry instantiates a new FabricConfigResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricConfigResultEntry(classId string, objectType string) *FabricConfigResultEntry {
	this := FabricConfigResultEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricConfigResultEntryWithDefaults instantiates a new FabricConfigResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricConfigResultEntryWithDefaults() *FabricConfigResultEntry {
	this := FabricConfigResultEntry{}
	var classId string = "fabric.ConfigResultEntry"
	this.ClassId = classId
	var objectType string = "fabric.ConfigResultEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricConfigResultEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricConfigResultEntry) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricConfigResultEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.ConfigResultEntry" of the ClassId field.
func (o *FabricConfigResultEntry) GetDefaultClassId() interface{} {
	return "fabric.ConfigResultEntry"
}

// GetObjectType returns the ObjectType field value
func (o *FabricConfigResultEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricConfigResultEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricConfigResultEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.ConfigResultEntry" of the ObjectType field.
func (o *FabricConfigResultEntry) GetDefaultObjectType() interface{} {
	return "fabric.ConfigResultEntry"
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricConfigResultEntry) GetConfigResult() FabricConfigResultRelationship {
	if o == nil || IsNil(o.ConfigResult.Get()) {
		var ret FabricConfigResultRelationship
		return ret
	}
	return *o.ConfigResult.Get()
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricConfigResultEntry) GetConfigResultOk() (*FabricConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigResult.Get(), o.ConfigResult.IsSet()
}

// HasConfigResult returns a boolean if a field has been set.
func (o *FabricConfigResultEntry) HasConfigResult() bool {
	if o != nil && o.ConfigResult.IsSet() {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given NullableFabricConfigResultRelationship and assigns it to the ConfigResult field.
func (o *FabricConfigResultEntry) SetConfigResult(v FabricConfigResultRelationship) {
	o.ConfigResult.Set(&v)
}

// SetConfigResultNil sets the value for ConfigResult to be an explicit nil
func (o *FabricConfigResultEntry) SetConfigResultNil() {
	o.ConfigResult.Set(nil)
}

// UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil
func (o *FabricConfigResultEntry) UnsetConfigResult() {
	o.ConfigResult.Unset()
}

func (o FabricConfigResultEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricConfigResultEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResultEntry, errPolicyAbstractConfigResultEntry := json.Marshal(o.PolicyAbstractConfigResultEntry)
	if errPolicyAbstractConfigResultEntry != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigResultEntry
	}
	errPolicyAbstractConfigResultEntry = json.Unmarshal([]byte(serializedPolicyAbstractConfigResultEntry), &toSerialize)
	if errPolicyAbstractConfigResultEntry != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigResultEntry
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ConfigResult.IsSet() {
		toSerialize["ConfigResult"] = o.ConfigResult.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricConfigResultEntry) UnmarshalJSON(data []byte) (err error) {
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
	type FabricConfigResultEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                 `json:"ObjectType"`
		ConfigResult NullableFabricConfigResultRelationship `json:"ConfigResult,omitempty"`
	}

	varFabricConfigResultEntryWithoutEmbeddedStruct := FabricConfigResultEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricConfigResultEntryWithoutEmbeddedStruct)
	if err == nil {
		varFabricConfigResultEntry := _FabricConfigResultEntry{}
		varFabricConfigResultEntry.ClassId = varFabricConfigResultEntryWithoutEmbeddedStruct.ClassId
		varFabricConfigResultEntry.ObjectType = varFabricConfigResultEntryWithoutEmbeddedStruct.ObjectType
		varFabricConfigResultEntry.ConfigResult = varFabricConfigResultEntryWithoutEmbeddedStruct.ConfigResult
		*o = FabricConfigResultEntry(varFabricConfigResultEntry)
	} else {
		return err
	}

	varFabricConfigResultEntry := _FabricConfigResultEntry{}

	err = json.Unmarshal(data, &varFabricConfigResultEntry)
	if err == nil {
		o.PolicyAbstractConfigResultEntry = varFabricConfigResultEntry.PolicyAbstractConfigResultEntry
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigResult")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResultEntry := reflect.ValueOf(o.PolicyAbstractConfigResultEntry)
		for i := 0; i < reflectPolicyAbstractConfigResultEntry.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResultEntry.Type().Field(i)

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

type NullableFabricConfigResultEntry struct {
	value *FabricConfigResultEntry
	isSet bool
}

func (v NullableFabricConfigResultEntry) Get() *FabricConfigResultEntry {
	return v.value
}

func (v *NullableFabricConfigResultEntry) Set(val *FabricConfigResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricConfigResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricConfigResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricConfigResultEntry(val *FabricConfigResultEntry) *NullableFabricConfigResultEntry {
	return &NullableFabricConfigResultEntry{value: val, isSet: true}
}

func (v NullableFabricConfigResultEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricConfigResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
