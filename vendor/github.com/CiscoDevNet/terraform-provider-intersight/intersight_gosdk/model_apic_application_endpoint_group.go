/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-17057
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

// checks if the ApicApplicationEndpointGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApicApplicationEndpointGroup{}

// ApicApplicationEndpointGroup APIC Application Endpoint Group.
type ApicApplicationEndpointGroup struct {
	ApicInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Dn *string `json:"Dn,omitempty"`
	// Name of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Name                 *string                             `json:"Name,omitempty"`
	Application          NullableApicApplicationRelationship `json:"Application,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApicApplicationEndpointGroup ApicApplicationEndpointGroup

// NewApicApplicationEndpointGroup instantiates a new ApicApplicationEndpointGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApicApplicationEndpointGroup(classId string, objectType string) *ApicApplicationEndpointGroup {
	this := ApicApplicationEndpointGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApicApplicationEndpointGroupWithDefaults instantiates a new ApicApplicationEndpointGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApicApplicationEndpointGroupWithDefaults() *ApicApplicationEndpointGroup {
	this := ApicApplicationEndpointGroup{}
	var classId string = "apic.ApplicationEndpointGroup"
	this.ClassId = classId
	var objectType string = "apic.ApplicationEndpointGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApicApplicationEndpointGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApicApplicationEndpointGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApicApplicationEndpointGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApicApplicationEndpointGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApicApplicationEndpointGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApicApplicationEndpointGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *ApicApplicationEndpointGroup) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicApplicationEndpointGroup) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ApicApplicationEndpointGroup) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *ApicApplicationEndpointGroup) SetDn(v string) {
	o.Dn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApicApplicationEndpointGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicApplicationEndpointGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApicApplicationEndpointGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApicApplicationEndpointGroup) SetName(v string) {
	o.Name = &v
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApicApplicationEndpointGroup) GetApplication() ApicApplicationRelationship {
	if o == nil || IsNil(o.Application.Get()) {
		var ret ApicApplicationRelationship
		return ret
	}
	return *o.Application.Get()
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApicApplicationEndpointGroup) GetApplicationOk() (*ApicApplicationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Application.Get(), o.Application.IsSet()
}

// HasApplication returns a boolean if a field has been set.
func (o *ApicApplicationEndpointGroup) HasApplication() bool {
	if o != nil && o.Application.IsSet() {
		return true
	}

	return false
}

// SetApplication gets a reference to the given NullableApicApplicationRelationship and assigns it to the Application field.
func (o *ApicApplicationEndpointGroup) SetApplication(v ApicApplicationRelationship) {
	o.Application.Set(&v)
}

// SetApplicationNil sets the value for Application to be an explicit nil
func (o *ApicApplicationEndpointGroup) SetApplicationNil() {
	o.Application.Set(nil)
}

// UnsetApplication ensures that no value is present for Application, not even an explicit nil
func (o *ApicApplicationEndpointGroup) UnsetApplication() {
	o.Application.Unset()
}

func (o ApicApplicationEndpointGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApicApplicationEndpointGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedApicInventoryEntity, errApicInventoryEntity := json.Marshal(o.ApicInventoryEntity)
	if errApicInventoryEntity != nil {
		return map[string]interface{}{}, errApicInventoryEntity
	}
	errApicInventoryEntity = json.Unmarshal([]byte(serializedApicInventoryEntity), &toSerialize)
	if errApicInventoryEntity != nil {
		return map[string]interface{}{}, errApicInventoryEntity
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Application.IsSet() {
		toSerialize["Application"] = o.Application.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApicApplicationEndpointGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type ApicApplicationEndpointGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Dn *string `json:"Dn,omitempty"`
		// Name of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Name        *string                             `json:"Name,omitempty"`
		Application NullableApicApplicationRelationship `json:"Application,omitempty"`
	}

	varApicApplicationEndpointGroupWithoutEmbeddedStruct := ApicApplicationEndpointGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApicApplicationEndpointGroupWithoutEmbeddedStruct)
	if err == nil {
		varApicApplicationEndpointGroup := _ApicApplicationEndpointGroup{}
		varApicApplicationEndpointGroup.ClassId = varApicApplicationEndpointGroupWithoutEmbeddedStruct.ClassId
		varApicApplicationEndpointGroup.ObjectType = varApicApplicationEndpointGroupWithoutEmbeddedStruct.ObjectType
		varApicApplicationEndpointGroup.Dn = varApicApplicationEndpointGroupWithoutEmbeddedStruct.Dn
		varApicApplicationEndpointGroup.Name = varApicApplicationEndpointGroupWithoutEmbeddedStruct.Name
		varApicApplicationEndpointGroup.Application = varApicApplicationEndpointGroupWithoutEmbeddedStruct.Application
		*o = ApicApplicationEndpointGroup(varApicApplicationEndpointGroup)
	} else {
		return err
	}

	varApicApplicationEndpointGroup := _ApicApplicationEndpointGroup{}

	err = json.Unmarshal(data, &varApicApplicationEndpointGroup)
	if err == nil {
		o.ApicInventoryEntity = varApicApplicationEndpointGroup.ApicInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Application")

		// remove fields from embedded structs
		reflectApicInventoryEntity := reflect.ValueOf(o.ApicInventoryEntity)
		for i := 0; i < reflectApicInventoryEntity.Type().NumField(); i++ {
			t := reflectApicInventoryEntity.Type().Field(i)

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

type NullableApicApplicationEndpointGroup struct {
	value *ApicApplicationEndpointGroup
	isSet bool
}

func (v NullableApicApplicationEndpointGroup) Get() *ApicApplicationEndpointGroup {
	return v.value
}

func (v *NullableApicApplicationEndpointGroup) Set(val *ApicApplicationEndpointGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableApicApplicationEndpointGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableApicApplicationEndpointGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicApplicationEndpointGroup(val *ApicApplicationEndpointGroup) *NullableApicApplicationEndpointGroup {
	return &NullableApicApplicationEndpointGroup{value: val, isSet: true}
}

func (v NullableApicApplicationEndpointGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicApplicationEndpointGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}