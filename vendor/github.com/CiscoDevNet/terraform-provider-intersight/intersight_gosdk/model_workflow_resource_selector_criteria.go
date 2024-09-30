/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
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

// checks if the WorkflowResourceSelectorCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowResourceSelectorCriteria{}

// WorkflowResourceSelectorCriteria The resource selector holds one or more Resource Selection conditions and the runtime parameters.
type WorkflowResourceSelectorCriteria struct {
	WorkflowAbstractResourceSelector
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The runtime properties and the value can be stored in this property.
	Parameters                interface{} `json:"Parameters,omitempty"`
	ResourceSelectionCriteria *MoMoRef    `json:"ResourceSelectionCriteria,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _WorkflowResourceSelectorCriteria WorkflowResourceSelectorCriteria

// NewWorkflowResourceSelectorCriteria instantiates a new WorkflowResourceSelectorCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowResourceSelectorCriteria(classId string, objectType string) *WorkflowResourceSelectorCriteria {
	this := WorkflowResourceSelectorCriteria{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowResourceSelectorCriteriaWithDefaults instantiates a new WorkflowResourceSelectorCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowResourceSelectorCriteriaWithDefaults() *WorkflowResourceSelectorCriteria {
	this := WorkflowResourceSelectorCriteria{}
	var classId string = "workflow.ResourceSelectorCriteria"
	this.ClassId = classId
	var objectType string = "workflow.ResourceSelectorCriteria"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowResourceSelectorCriteria) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowResourceSelectorCriteria) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowResourceSelectorCriteria) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.ResourceSelectorCriteria" of the ClassId field.
func (o *WorkflowResourceSelectorCriteria) GetDefaultClassId() interface{} {
	return "workflow.ResourceSelectorCriteria"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowResourceSelectorCriteria) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowResourceSelectorCriteria) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowResourceSelectorCriteria) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.ResourceSelectorCriteria" of the ObjectType field.
func (o *WorkflowResourceSelectorCriteria) GetDefaultObjectType() interface{} {
	return "workflow.ResourceSelectorCriteria"
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowResourceSelectorCriteria) GetParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowResourceSelectorCriteria) GetParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowResourceSelectorCriteria) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given interface{} and assigns it to the Parameters field.
func (o *WorkflowResourceSelectorCriteria) SetParameters(v interface{}) {
	o.Parameters = v
}

// GetResourceSelectionCriteria returns the ResourceSelectionCriteria field value if set, zero value otherwise.
func (o *WorkflowResourceSelectorCriteria) GetResourceSelectionCriteria() MoMoRef {
	if o == nil || IsNil(o.ResourceSelectionCriteria) {
		var ret MoMoRef
		return ret
	}
	return *o.ResourceSelectionCriteria
}

// GetResourceSelectionCriteriaOk returns a tuple with the ResourceSelectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowResourceSelectorCriteria) GetResourceSelectionCriteriaOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.ResourceSelectionCriteria) {
		return nil, false
	}
	return o.ResourceSelectionCriteria, true
}

// HasResourceSelectionCriteria returns a boolean if a field has been set.
func (o *WorkflowResourceSelectorCriteria) HasResourceSelectionCriteria() bool {
	if o != nil && !IsNil(o.ResourceSelectionCriteria) {
		return true
	}

	return false
}

// SetResourceSelectionCriteria gets a reference to the given MoMoRef and assigns it to the ResourceSelectionCriteria field.
func (o *WorkflowResourceSelectorCriteria) SetResourceSelectionCriteria(v MoMoRef) {
	o.ResourceSelectionCriteria = &v
}

func (o WorkflowResourceSelectorCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowResourceSelectorCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowAbstractResourceSelector, errWorkflowAbstractResourceSelector := json.Marshal(o.WorkflowAbstractResourceSelector)
	if errWorkflowAbstractResourceSelector != nil {
		return map[string]interface{}{}, errWorkflowAbstractResourceSelector
	}
	errWorkflowAbstractResourceSelector = json.Unmarshal([]byte(serializedWorkflowAbstractResourceSelector), &toSerialize)
	if errWorkflowAbstractResourceSelector != nil {
		return map[string]interface{}{}, errWorkflowAbstractResourceSelector
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if !IsNil(o.ResourceSelectionCriteria) {
		toSerialize["ResourceSelectionCriteria"] = o.ResourceSelectionCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowResourceSelectorCriteria) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowResourceSelectorCriteriaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The runtime properties and the value can be stored in this property.
		Parameters                interface{} `json:"Parameters,omitempty"`
		ResourceSelectionCriteria *MoMoRef    `json:"ResourceSelectionCriteria,omitempty"`
	}

	varWorkflowResourceSelectorCriteriaWithoutEmbeddedStruct := WorkflowResourceSelectorCriteriaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowResourceSelectorCriteriaWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowResourceSelectorCriteria := _WorkflowResourceSelectorCriteria{}
		varWorkflowResourceSelectorCriteria.ClassId = varWorkflowResourceSelectorCriteriaWithoutEmbeddedStruct.ClassId
		varWorkflowResourceSelectorCriteria.ObjectType = varWorkflowResourceSelectorCriteriaWithoutEmbeddedStruct.ObjectType
		varWorkflowResourceSelectorCriteria.Parameters = varWorkflowResourceSelectorCriteriaWithoutEmbeddedStruct.Parameters
		varWorkflowResourceSelectorCriteria.ResourceSelectionCriteria = varWorkflowResourceSelectorCriteriaWithoutEmbeddedStruct.ResourceSelectionCriteria
		*o = WorkflowResourceSelectorCriteria(varWorkflowResourceSelectorCriteria)
	} else {
		return err
	}

	varWorkflowResourceSelectorCriteria := _WorkflowResourceSelectorCriteria{}

	err = json.Unmarshal(data, &varWorkflowResourceSelectorCriteria)
	if err == nil {
		o.WorkflowAbstractResourceSelector = varWorkflowResourceSelectorCriteria.WorkflowAbstractResourceSelector
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Parameters")
		delete(additionalProperties, "ResourceSelectionCriteria")

		// remove fields from embedded structs
		reflectWorkflowAbstractResourceSelector := reflect.ValueOf(o.WorkflowAbstractResourceSelector)
		for i := 0; i < reflectWorkflowAbstractResourceSelector.Type().NumField(); i++ {
			t := reflectWorkflowAbstractResourceSelector.Type().Field(i)

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

type NullableWorkflowResourceSelectorCriteria struct {
	value *WorkflowResourceSelectorCriteria
	isSet bool
}

func (v NullableWorkflowResourceSelectorCriteria) Get() *WorkflowResourceSelectorCriteria {
	return v.value
}

func (v *NullableWorkflowResourceSelectorCriteria) Set(val *WorkflowResourceSelectorCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowResourceSelectorCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowResourceSelectorCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowResourceSelectorCriteria(val *WorkflowResourceSelectorCriteria) *NullableWorkflowResourceSelectorCriteria {
	return &NullableWorkflowResourceSelectorCriteria{value: val, isSet: true}
}

func (v NullableWorkflowResourceSelectorCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowResourceSelectorCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
