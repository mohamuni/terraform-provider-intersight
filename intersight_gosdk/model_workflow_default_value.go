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

// checks if the WorkflowDefaultValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowDefaultValue{}

// WorkflowDefaultValue Captures default vales for a data type.
type WorkflowDefaultValue struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.
	IsValueSet *bool `json:"IsValueSet,omitempty"`
	// Override the default value provided for the data type. When true, allow the user to enter value for the data type.
	Override *bool `json:"Override,omitempty"`
	// Default value for the data type. If default value was provided and the input was required the default value will be used as the input.
	Value                interface{} `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDefaultValue WorkflowDefaultValue

// NewWorkflowDefaultValue instantiates a new WorkflowDefaultValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDefaultValue(classId string, objectType string) *WorkflowDefaultValue {
	this := WorkflowDefaultValue{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowDefaultValueWithDefaults instantiates a new WorkflowDefaultValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDefaultValueWithDefaults() *WorkflowDefaultValue {
	this := WorkflowDefaultValue{}
	var classId string = "workflow.DefaultValue"
	this.ClassId = classId
	var objectType string = "workflow.DefaultValue"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDefaultValue) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValue) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDefaultValue) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.DefaultValue" of the ClassId field.
func (o *WorkflowDefaultValue) GetDefaultClassId() interface{} {
	return "workflow.DefaultValue"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDefaultValue) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValue) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDefaultValue) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.DefaultValue" of the ObjectType field.
func (o *WorkflowDefaultValue) GetDefaultObjectType() interface{} {
	return "workflow.DefaultValue"
}

// GetIsValueSet returns the IsValueSet field value if set, zero value otherwise.
func (o *WorkflowDefaultValue) GetIsValueSet() bool {
	if o == nil || IsNil(o.IsValueSet) {
		var ret bool
		return ret
	}
	return *o.IsValueSet
}

// GetIsValueSetOk returns a tuple with the IsValueSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValue) GetIsValueSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValueSet) {
		return nil, false
	}
	return o.IsValueSet, true
}

// HasIsValueSet returns a boolean if a field has been set.
func (o *WorkflowDefaultValue) HasIsValueSet() bool {
	if o != nil && !IsNil(o.IsValueSet) {
		return true
	}

	return false
}

// SetIsValueSet gets a reference to the given bool and assigns it to the IsValueSet field.
func (o *WorkflowDefaultValue) SetIsValueSet(v bool) {
	o.IsValueSet = &v
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *WorkflowDefaultValue) GetOverride() bool {
	if o == nil || IsNil(o.Override) {
		var ret bool
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValue) GetOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.Override) {
		return nil, false
	}
	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *WorkflowDefaultValue) HasOverride() bool {
	if o != nil && !IsNil(o.Override) {
		return true
	}

	return false
}

// SetOverride gets a reference to the given bool and assigns it to the Override field.
func (o *WorkflowDefaultValue) SetOverride(v bool) {
	o.Override = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDefaultValue) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDefaultValue) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowDefaultValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *WorkflowDefaultValue) SetValue(v interface{}) {
	o.Value = v
}

func (o WorkflowDefaultValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowDefaultValue) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsValueSet) {
		toSerialize["IsValueSet"] = o.IsValueSet
	}
	if !IsNil(o.Override) {
		toSerialize["Override"] = o.Override
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowDefaultValue) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowDefaultValueWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.
		IsValueSet *bool `json:"IsValueSet,omitempty"`
		// Override the default value provided for the data type. When true, allow the user to enter value for the data type.
		Override *bool `json:"Override,omitempty"`
		// Default value for the data type. If default value was provided and the input was required the default value will be used as the input.
		Value interface{} `json:"Value,omitempty"`
	}

	varWorkflowDefaultValueWithoutEmbeddedStruct := WorkflowDefaultValueWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowDefaultValueWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowDefaultValue := _WorkflowDefaultValue{}
		varWorkflowDefaultValue.ClassId = varWorkflowDefaultValueWithoutEmbeddedStruct.ClassId
		varWorkflowDefaultValue.ObjectType = varWorkflowDefaultValueWithoutEmbeddedStruct.ObjectType
		varWorkflowDefaultValue.IsValueSet = varWorkflowDefaultValueWithoutEmbeddedStruct.IsValueSet
		varWorkflowDefaultValue.Override = varWorkflowDefaultValueWithoutEmbeddedStruct.Override
		varWorkflowDefaultValue.Value = varWorkflowDefaultValueWithoutEmbeddedStruct.Value
		*o = WorkflowDefaultValue(varWorkflowDefaultValue)
	} else {
		return err
	}

	varWorkflowDefaultValue := _WorkflowDefaultValue{}

	err = json.Unmarshal(data, &varWorkflowDefaultValue)
	if err == nil {
		o.MoBaseComplexType = varWorkflowDefaultValue.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsValueSet")
		delete(additionalProperties, "Override")
		delete(additionalProperties, "Value")

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

type NullableWorkflowDefaultValue struct {
	value *WorkflowDefaultValue
	isSet bool
}

func (v NullableWorkflowDefaultValue) Get() *WorkflowDefaultValue {
	return v.value
}

func (v *NullableWorkflowDefaultValue) Set(val *WorkflowDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDefaultValue(val *WorkflowDefaultValue) *NullableWorkflowDefaultValue {
	return &NullableWorkflowDefaultValue{value: val, isSet: true}
}

func (v NullableWorkflowDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
