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

// checks if the WorkflowMoInventoryProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowMoInventoryProperty{}

// WorkflowMoInventoryProperty Captures the properties from the original Intersight resource that need to be reflected in the datatype.
type WorkflowMoInventoryProperty struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Attributes []string `json:"Attributes,omitempty"`
	// ObjectType for which the attributes need to be collected.
	ReferenceObjectType *string `json:"ReferenceObjectType,omitempty"`
	// Defines how the reference to the shadow resource is done. Base case is via an Moid, but reference via a selector is also possible. * `Moid` - The reference to the original resource is via an Moid. * `Selector` - The reference to the original resource is via a selector query. This can potentially lead to tracking data for multiple resources.
	ReferenceType        *string `json:"ReferenceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMoInventoryProperty WorkflowMoInventoryProperty

// NewWorkflowMoInventoryProperty instantiates a new WorkflowMoInventoryProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMoInventoryProperty(classId string, objectType string) *WorkflowMoInventoryProperty {
	this := WorkflowMoInventoryProperty{}
	this.ClassId = classId
	this.ObjectType = objectType
	var referenceType string = "Moid"
	this.ReferenceType = &referenceType
	return &this
}

// NewWorkflowMoInventoryPropertyWithDefaults instantiates a new WorkflowMoInventoryProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMoInventoryPropertyWithDefaults() *WorkflowMoInventoryProperty {
	this := WorkflowMoInventoryProperty{}
	var classId string = "workflow.MoInventoryProperty"
	this.ClassId = classId
	var objectType string = "workflow.MoInventoryProperty"
	this.ObjectType = objectType
	var referenceType string = "Moid"
	this.ReferenceType = &referenceType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMoInventoryProperty) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoInventoryProperty) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMoInventoryProperty) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.MoInventoryProperty" of the ClassId field.
func (o *WorkflowMoInventoryProperty) GetDefaultClassId() interface{} {
	return "workflow.MoInventoryProperty"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMoInventoryProperty) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoInventoryProperty) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMoInventoryProperty) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.MoInventoryProperty" of the ObjectType field.
func (o *WorkflowMoInventoryProperty) GetDefaultObjectType() interface{} {
	return "workflow.MoInventoryProperty"
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoInventoryProperty) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoInventoryProperty) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkflowMoInventoryProperty) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *WorkflowMoInventoryProperty) SetAttributes(v []string) {
	o.Attributes = v
}

// GetReferenceObjectType returns the ReferenceObjectType field value if set, zero value otherwise.
func (o *WorkflowMoInventoryProperty) GetReferenceObjectType() string {
	if o == nil || IsNil(o.ReferenceObjectType) {
		var ret string
		return ret
	}
	return *o.ReferenceObjectType
}

// GetReferenceObjectTypeOk returns a tuple with the ReferenceObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMoInventoryProperty) GetReferenceObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceObjectType) {
		return nil, false
	}
	return o.ReferenceObjectType, true
}

// HasReferenceObjectType returns a boolean if a field has been set.
func (o *WorkflowMoInventoryProperty) HasReferenceObjectType() bool {
	if o != nil && !IsNil(o.ReferenceObjectType) {
		return true
	}

	return false
}

// SetReferenceObjectType gets a reference to the given string and assigns it to the ReferenceObjectType field.
func (o *WorkflowMoInventoryProperty) SetReferenceObjectType(v string) {
	o.ReferenceObjectType = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *WorkflowMoInventoryProperty) GetReferenceType() string {
	if o == nil || IsNil(o.ReferenceType) {
		var ret string
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMoInventoryProperty) GetReferenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceType) {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *WorkflowMoInventoryProperty) HasReferenceType() bool {
	if o != nil && !IsNil(o.ReferenceType) {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given string and assigns it to the ReferenceType field.
func (o *WorkflowMoInventoryProperty) SetReferenceType(v string) {
	o.ReferenceType = &v
}

func (o WorkflowMoInventoryProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowMoInventoryProperty) ToMap() (map[string]interface{}, error) {
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
	if o.Attributes != nil {
		toSerialize["Attributes"] = o.Attributes
	}
	if !IsNil(o.ReferenceObjectType) {
		toSerialize["ReferenceObjectType"] = o.ReferenceObjectType
	}
	if !IsNil(o.ReferenceType) {
		toSerialize["ReferenceType"] = o.ReferenceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowMoInventoryProperty) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowMoInventoryPropertyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Attributes []string `json:"Attributes,omitempty"`
		// ObjectType for which the attributes need to be collected.
		ReferenceObjectType *string `json:"ReferenceObjectType,omitempty"`
		// Defines how the reference to the shadow resource is done. Base case is via an Moid, but reference via a selector is also possible. * `Moid` - The reference to the original resource is via an Moid. * `Selector` - The reference to the original resource is via a selector query. This can potentially lead to tracking data for multiple resources.
		ReferenceType *string `json:"ReferenceType,omitempty"`
	}

	varWorkflowMoInventoryPropertyWithoutEmbeddedStruct := WorkflowMoInventoryPropertyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowMoInventoryPropertyWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowMoInventoryProperty := _WorkflowMoInventoryProperty{}
		varWorkflowMoInventoryProperty.ClassId = varWorkflowMoInventoryPropertyWithoutEmbeddedStruct.ClassId
		varWorkflowMoInventoryProperty.ObjectType = varWorkflowMoInventoryPropertyWithoutEmbeddedStruct.ObjectType
		varWorkflowMoInventoryProperty.Attributes = varWorkflowMoInventoryPropertyWithoutEmbeddedStruct.Attributes
		varWorkflowMoInventoryProperty.ReferenceObjectType = varWorkflowMoInventoryPropertyWithoutEmbeddedStruct.ReferenceObjectType
		varWorkflowMoInventoryProperty.ReferenceType = varWorkflowMoInventoryPropertyWithoutEmbeddedStruct.ReferenceType
		*o = WorkflowMoInventoryProperty(varWorkflowMoInventoryProperty)
	} else {
		return err
	}

	varWorkflowMoInventoryProperty := _WorkflowMoInventoryProperty{}

	err = json.Unmarshal(data, &varWorkflowMoInventoryProperty)
	if err == nil {
		o.MoBaseComplexType = varWorkflowMoInventoryProperty.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Attributes")
		delete(additionalProperties, "ReferenceObjectType")
		delete(additionalProperties, "ReferenceType")

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

type NullableWorkflowMoInventoryProperty struct {
	value *WorkflowMoInventoryProperty
	isSet bool
}

func (v NullableWorkflowMoInventoryProperty) Get() *WorkflowMoInventoryProperty {
	return v.value
}

func (v *NullableWorkflowMoInventoryProperty) Set(val *WorkflowMoInventoryProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMoInventoryProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMoInventoryProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMoInventoryProperty(val *WorkflowMoInventoryProperty) *NullableWorkflowMoInventoryProperty {
	return &NullableWorkflowMoInventoryProperty{value: val, isSet: true}
}

func (v NullableWorkflowMoInventoryProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMoInventoryProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
