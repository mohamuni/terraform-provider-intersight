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

// checks if the MotemplateActionEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MotemplateActionEntry{}

// MotemplateActionEntry An action entry that contains the action type and its associated action parameters as a key:value pair.
type MotemplateActionEntry struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                  `json:"ObjectType"`
	Params     []MotemplateActionParam `json:"Params,omitempty"`
	// The action type to be executed. * `Sync` - The action to merge values from the template to its derived objects. * `Deploy` - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types. * `Detach` - The action to detach the current derived object from its attached template. * `Attach` - The action to attach the current object to the specified template.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MotemplateActionEntry MotemplateActionEntry

// NewMotemplateActionEntry instantiates a new MotemplateActionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMotemplateActionEntry(classId string, objectType string) *MotemplateActionEntry {
	this := MotemplateActionEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "Sync"
	this.Type = &type_
	return &this
}

// NewMotemplateActionEntryWithDefaults instantiates a new MotemplateActionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMotemplateActionEntryWithDefaults() *MotemplateActionEntry {
	this := MotemplateActionEntry{}
	var classId string = "motemplate.ActionEntry"
	this.ClassId = classId
	var objectType string = "motemplate.ActionEntry"
	this.ObjectType = objectType
	var type_ string = "Sync"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *MotemplateActionEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MotemplateActionEntry) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MotemplateActionEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "motemplate.ActionEntry" of the ClassId field.
func (o *MotemplateActionEntry) GetDefaultClassId() interface{} {
	return "motemplate.ActionEntry"
}

// GetObjectType returns the ObjectType field value
func (o *MotemplateActionEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MotemplateActionEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MotemplateActionEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "motemplate.ActionEntry" of the ObjectType field.
func (o *MotemplateActionEntry) GetDefaultObjectType() interface{} {
	return "motemplate.ActionEntry"
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MotemplateActionEntry) GetParams() []MotemplateActionParam {
	if o == nil {
		var ret []MotemplateActionParam
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MotemplateActionEntry) GetParamsOk() ([]MotemplateActionParam, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *MotemplateActionEntry) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []MotemplateActionParam and assigns it to the Params field.
func (o *MotemplateActionEntry) SetParams(v []MotemplateActionParam) {
	o.Params = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MotemplateActionEntry) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotemplateActionEntry) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MotemplateActionEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MotemplateActionEntry) SetType(v string) {
	o.Type = &v
}

func (o MotemplateActionEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MotemplateActionEntry) ToMap() (map[string]interface{}, error) {
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
	if o.Params != nil {
		toSerialize["Params"] = o.Params
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MotemplateActionEntry) UnmarshalJSON(data []byte) (err error) {
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
	type MotemplateActionEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                  `json:"ObjectType"`
		Params     []MotemplateActionParam `json:"Params,omitempty"`
		// The action type to be executed. * `Sync` - The action to merge values from the template to its derived objects. * `Deploy` - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types. * `Detach` - The action to detach the current derived object from its attached template. * `Attach` - The action to attach the current object to the specified template.
		Type *string `json:"Type,omitempty"`
	}

	varMotemplateActionEntryWithoutEmbeddedStruct := MotemplateActionEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMotemplateActionEntryWithoutEmbeddedStruct)
	if err == nil {
		varMotemplateActionEntry := _MotemplateActionEntry{}
		varMotemplateActionEntry.ClassId = varMotemplateActionEntryWithoutEmbeddedStruct.ClassId
		varMotemplateActionEntry.ObjectType = varMotemplateActionEntryWithoutEmbeddedStruct.ObjectType
		varMotemplateActionEntry.Params = varMotemplateActionEntryWithoutEmbeddedStruct.Params
		varMotemplateActionEntry.Type = varMotemplateActionEntryWithoutEmbeddedStruct.Type
		*o = MotemplateActionEntry(varMotemplateActionEntry)
	} else {
		return err
	}

	varMotemplateActionEntry := _MotemplateActionEntry{}

	err = json.Unmarshal(data, &varMotemplateActionEntry)
	if err == nil {
		o.MoBaseComplexType = varMotemplateActionEntry.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Params")
		delete(additionalProperties, "Type")

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

type NullableMotemplateActionEntry struct {
	value *MotemplateActionEntry
	isSet bool
}

func (v NullableMotemplateActionEntry) Get() *MotemplateActionEntry {
	return v.value
}

func (v *NullableMotemplateActionEntry) Set(val *MotemplateActionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableMotemplateActionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableMotemplateActionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMotemplateActionEntry(val *MotemplateActionEntry) *NullableMotemplateActionEntry {
	return &NullableMotemplateActionEntry{value: val, isSet: true}
}

func (v NullableMotemplateActionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMotemplateActionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
