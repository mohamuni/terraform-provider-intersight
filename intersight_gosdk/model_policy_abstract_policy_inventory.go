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

// checks if the PolicyAbstractPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAbstractPolicyInventory{}

// PolicyAbstractPolicyInventory A base abstract class for all Inventoried policies that can be applied to profiles.
type PolicyAbstractPolicyInventory struct {
	PolicyAbstractInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Description of the policy.
	Description *string `json:"Description,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]+[\\\\sa-zA-Z0-9_'.:-]*$"`
	// Name of the inventoried policy object.
	Name                 *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,512}$"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractPolicyInventory PolicyAbstractPolicyInventory

// NewPolicyAbstractPolicyInventory instantiates a new PolicyAbstractPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractPolicyInventory(classId string, objectType string) *PolicyAbstractPolicyInventory {
	this := PolicyAbstractPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyAbstractPolicyInventoryWithDefaults instantiates a new PolicyAbstractPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractPolicyInventoryWithDefaults() *PolicyAbstractPolicyInventory {
	this := PolicyAbstractPolicyInventory{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyAbstractPolicyInventory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractPolicyInventory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyAbstractPolicyInventory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyAbstractPolicyInventory) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyAbstractPolicyInventory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractPolicyInventory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyAbstractPolicyInventory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyAbstractPolicyInventory) SetName(v string) {
	o.Name = &v
}

func (o PolicyAbstractPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAbstractPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractInventory, errPolicyAbstractInventory := json.Marshal(o.PolicyAbstractInventory)
	if errPolicyAbstractInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractInventory
	}
	errPolicyAbstractInventory = json.Unmarshal([]byte(serializedPolicyAbstractInventory), &toSerialize)
	if errPolicyAbstractInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractInventory
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAbstractPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type PolicyAbstractPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Description of the policy.
		Description *string `json:"Description,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]+[\\\\sa-zA-Z0-9_'.:-]*$"`
		// Name of the inventoried policy object.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,512}$"`
	}

	varPolicyAbstractPolicyInventoryWithoutEmbeddedStruct := PolicyAbstractPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPolicyAbstractPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractPolicyInventory := _PolicyAbstractPolicyInventory{}
		varPolicyAbstractPolicyInventory.ClassId = varPolicyAbstractPolicyInventoryWithoutEmbeddedStruct.ClassId
		varPolicyAbstractPolicyInventory.ObjectType = varPolicyAbstractPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varPolicyAbstractPolicyInventory.Description = varPolicyAbstractPolicyInventoryWithoutEmbeddedStruct.Description
		varPolicyAbstractPolicyInventory.Name = varPolicyAbstractPolicyInventoryWithoutEmbeddedStruct.Name
		*o = PolicyAbstractPolicyInventory(varPolicyAbstractPolicyInventory)
	} else {
		return err
	}

	varPolicyAbstractPolicyInventory := _PolicyAbstractPolicyInventory{}

	err = json.Unmarshal(data, &varPolicyAbstractPolicyInventory)
	if err == nil {
		o.PolicyAbstractInventory = varPolicyAbstractPolicyInventory.PolicyAbstractInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectPolicyAbstractInventory := reflect.ValueOf(o.PolicyAbstractInventory)
		for i := 0; i < reflectPolicyAbstractInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractInventory.Type().Field(i)

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

type NullablePolicyAbstractPolicyInventory struct {
	value *PolicyAbstractPolicyInventory
	isSet bool
}

func (v NullablePolicyAbstractPolicyInventory) Get() *PolicyAbstractPolicyInventory {
	return v.value
}

func (v *NullablePolicyAbstractPolicyInventory) Set(val *PolicyAbstractPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractPolicyInventory(val *PolicyAbstractPolicyInventory) *NullablePolicyAbstractPolicyInventory {
	return &NullablePolicyAbstractPolicyInventory{value: val, isSet: true}
}

func (v NullablePolicyAbstractPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
