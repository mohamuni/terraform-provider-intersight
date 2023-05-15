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
)

// WorkflowResourceInlineSelectorAllOf Definition of the list of properties defined in 'workflow.ResourceInlineSelector', excluding properties defined in parent classes.
type WorkflowResourceInlineSelectorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The odata based filter URL to select the resources.
	Selector             *string `json:"Selector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowResourceInlineSelectorAllOf WorkflowResourceInlineSelectorAllOf

// NewWorkflowResourceInlineSelectorAllOf instantiates a new WorkflowResourceInlineSelectorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowResourceInlineSelectorAllOf(classId string, objectType string) *WorkflowResourceInlineSelectorAllOf {
	this := WorkflowResourceInlineSelectorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowResourceInlineSelectorAllOfWithDefaults instantiates a new WorkflowResourceInlineSelectorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowResourceInlineSelectorAllOfWithDefaults() *WorkflowResourceInlineSelectorAllOf {
	this := WorkflowResourceInlineSelectorAllOf{}
	var classId string = "workflow.ResourceInlineSelector"
	this.ClassId = classId
	var objectType string = "workflow.ResourceInlineSelector"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowResourceInlineSelectorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowResourceInlineSelectorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowResourceInlineSelectorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowResourceInlineSelectorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowResourceInlineSelectorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowResourceInlineSelectorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *WorkflowResourceInlineSelectorAllOf) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowResourceInlineSelectorAllOf) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *WorkflowResourceInlineSelectorAllOf) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *WorkflowResourceInlineSelectorAllOf) SetSelector(v string) {
	o.Selector = &v
}

func (o WorkflowResourceInlineSelectorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowResourceInlineSelectorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowResourceInlineSelectorAllOf := _WorkflowResourceInlineSelectorAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowResourceInlineSelectorAllOf); err == nil {
		*o = WorkflowResourceInlineSelectorAllOf(varWorkflowResourceInlineSelectorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Selector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowResourceInlineSelectorAllOf struct {
	value *WorkflowResourceInlineSelectorAllOf
	isSet bool
}

func (v NullableWorkflowResourceInlineSelectorAllOf) Get() *WorkflowResourceInlineSelectorAllOf {
	return v.value
}

func (v *NullableWorkflowResourceInlineSelectorAllOf) Set(val *WorkflowResourceInlineSelectorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowResourceInlineSelectorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowResourceInlineSelectorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowResourceInlineSelectorAllOf(val *WorkflowResourceInlineSelectorAllOf) *NullableWorkflowResourceInlineSelectorAllOf {
	return &NullableWorkflowResourceInlineSelectorAllOf{value: val, isSet: true}
}

func (v NullableWorkflowResourceInlineSelectorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowResourceInlineSelectorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}