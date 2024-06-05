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

// checks if the ApicFabricLeafNodeInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApicFabricLeafNodeInterface{}

// ApicFabricLeafNodeInterface APIC Fabric Leaf Node Interfaces.
type ApicFabricLeafNodeInterface struct {
	ApicInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Dn *string `json:"Dn,omitempty"`
	// Fabric Leaf Node Distinguished Name.
	FabricLeafNodeDn *string `json:"FabricLeafNodeDn,omitempty"`
	// Fabric Leaf Node Identification Number.
	FabricLeafNodeId *string `json:"FabricLeafNodeId,omitempty"`
	// Name of an object within the Cisco Application Policy Infrastructure Controller.
	Name                 *string                                `json:"Name,omitempty"`
	FabricLeafNode       NullableApicFabricLeafNodeRelationship `json:"FabricLeafNode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApicFabricLeafNodeInterface ApicFabricLeafNodeInterface

// NewApicFabricLeafNodeInterface instantiates a new ApicFabricLeafNodeInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApicFabricLeafNodeInterface(classId string, objectType string) *ApicFabricLeafNodeInterface {
	this := ApicFabricLeafNodeInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApicFabricLeafNodeInterfaceWithDefaults instantiates a new ApicFabricLeafNodeInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApicFabricLeafNodeInterfaceWithDefaults() *ApicFabricLeafNodeInterface {
	this := ApicFabricLeafNodeInterface{}
	var classId string = "apic.FabricLeafNodeInterface"
	this.ClassId = classId
	var objectType string = "apic.FabricLeafNodeInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApicFabricLeafNodeInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNodeInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApicFabricLeafNodeInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApicFabricLeafNodeInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNodeInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApicFabricLeafNodeInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *ApicFabricLeafNodeInterface) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNodeInterface) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ApicFabricLeafNodeInterface) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *ApicFabricLeafNodeInterface) SetDn(v string) {
	o.Dn = &v
}

// GetFabricLeafNodeDn returns the FabricLeafNodeDn field value if set, zero value otherwise.
func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeDn() string {
	if o == nil || IsNil(o.FabricLeafNodeDn) {
		var ret string
		return ret
	}
	return *o.FabricLeafNodeDn
}

// GetFabricLeafNodeDnOk returns a tuple with the FabricLeafNodeDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeDnOk() (*string, bool) {
	if o == nil || IsNil(o.FabricLeafNodeDn) {
		return nil, false
	}
	return o.FabricLeafNodeDn, true
}

// HasFabricLeafNodeDn returns a boolean if a field has been set.
func (o *ApicFabricLeafNodeInterface) HasFabricLeafNodeDn() bool {
	if o != nil && !IsNil(o.FabricLeafNodeDn) {
		return true
	}

	return false
}

// SetFabricLeafNodeDn gets a reference to the given string and assigns it to the FabricLeafNodeDn field.
func (o *ApicFabricLeafNodeInterface) SetFabricLeafNodeDn(v string) {
	o.FabricLeafNodeDn = &v
}

// GetFabricLeafNodeId returns the FabricLeafNodeId field value if set, zero value otherwise.
func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeId() string {
	if o == nil || IsNil(o.FabricLeafNodeId) {
		var ret string
		return ret
	}
	return *o.FabricLeafNodeId
}

// GetFabricLeafNodeIdOk returns a tuple with the FabricLeafNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.FabricLeafNodeId) {
		return nil, false
	}
	return o.FabricLeafNodeId, true
}

// HasFabricLeafNodeId returns a boolean if a field has been set.
func (o *ApicFabricLeafNodeInterface) HasFabricLeafNodeId() bool {
	if o != nil && !IsNil(o.FabricLeafNodeId) {
		return true
	}

	return false
}

// SetFabricLeafNodeId gets a reference to the given string and assigns it to the FabricLeafNodeId field.
func (o *ApicFabricLeafNodeInterface) SetFabricLeafNodeId(v string) {
	o.FabricLeafNodeId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApicFabricLeafNodeInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNodeInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApicFabricLeafNodeInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApicFabricLeafNodeInterface) SetName(v string) {
	o.Name = &v
}

// GetFabricLeafNode returns the FabricLeafNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApicFabricLeafNodeInterface) GetFabricLeafNode() ApicFabricLeafNodeRelationship {
	if o == nil || IsNil(o.FabricLeafNode.Get()) {
		var ret ApicFabricLeafNodeRelationship
		return ret
	}
	return *o.FabricLeafNode.Get()
}

// GetFabricLeafNodeOk returns a tuple with the FabricLeafNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeOk() (*ApicFabricLeafNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.FabricLeafNode.Get(), o.FabricLeafNode.IsSet()
}

// HasFabricLeafNode returns a boolean if a field has been set.
func (o *ApicFabricLeafNodeInterface) HasFabricLeafNode() bool {
	if o != nil && o.FabricLeafNode.IsSet() {
		return true
	}

	return false
}

// SetFabricLeafNode gets a reference to the given NullableApicFabricLeafNodeRelationship and assigns it to the FabricLeafNode field.
func (o *ApicFabricLeafNodeInterface) SetFabricLeafNode(v ApicFabricLeafNodeRelationship) {
	o.FabricLeafNode.Set(&v)
}

// SetFabricLeafNodeNil sets the value for FabricLeafNode to be an explicit nil
func (o *ApicFabricLeafNodeInterface) SetFabricLeafNodeNil() {
	o.FabricLeafNode.Set(nil)
}

// UnsetFabricLeafNode ensures that no value is present for FabricLeafNode, not even an explicit nil
func (o *ApicFabricLeafNodeInterface) UnsetFabricLeafNode() {
	o.FabricLeafNode.Unset()
}

func (o ApicFabricLeafNodeInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApicFabricLeafNodeInterface) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FabricLeafNodeDn) {
		toSerialize["FabricLeafNodeDn"] = o.FabricLeafNodeDn
	}
	if !IsNil(o.FabricLeafNodeId) {
		toSerialize["FabricLeafNodeId"] = o.FabricLeafNodeId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.FabricLeafNode.IsSet() {
		toSerialize["FabricLeafNode"] = o.FabricLeafNode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApicFabricLeafNodeInterface) UnmarshalJSON(data []byte) (err error) {
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

	type ApicFabricLeafNodeInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Dn *string `json:"Dn,omitempty"`
		// Fabric Leaf Node Distinguished Name.
		FabricLeafNodeDn *string `json:"FabricLeafNodeDn,omitempty"`
		// Fabric Leaf Node Identification Number.
		FabricLeafNodeId *string `json:"FabricLeafNodeId,omitempty"`
		// Name of an object within the Cisco Application Policy Infrastructure Controller.
		Name           *string                                `json:"Name,omitempty"`
		FabricLeafNode NullableApicFabricLeafNodeRelationship `json:"FabricLeafNode,omitempty"`
	}

	varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct := ApicFabricLeafNodeInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varApicFabricLeafNodeInterface := _ApicFabricLeafNodeInterface{}
		varApicFabricLeafNodeInterface.ClassId = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.ClassId
		varApicFabricLeafNodeInterface.ObjectType = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.ObjectType
		varApicFabricLeafNodeInterface.Dn = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.Dn
		varApicFabricLeafNodeInterface.FabricLeafNodeDn = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.FabricLeafNodeDn
		varApicFabricLeafNodeInterface.FabricLeafNodeId = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.FabricLeafNodeId
		varApicFabricLeafNodeInterface.Name = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.Name
		varApicFabricLeafNodeInterface.FabricLeafNode = varApicFabricLeafNodeInterfaceWithoutEmbeddedStruct.FabricLeafNode
		*o = ApicFabricLeafNodeInterface(varApicFabricLeafNodeInterface)
	} else {
		return err
	}

	varApicFabricLeafNodeInterface := _ApicFabricLeafNodeInterface{}

	err = json.Unmarshal(data, &varApicFabricLeafNodeInterface)
	if err == nil {
		o.ApicInventoryEntity = varApicFabricLeafNodeInterface.ApicInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FabricLeafNodeDn")
		delete(additionalProperties, "FabricLeafNodeId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "FabricLeafNode")

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

type NullableApicFabricLeafNodeInterface struct {
	value *ApicFabricLeafNodeInterface
	isSet bool
}

func (v NullableApicFabricLeafNodeInterface) Get() *ApicFabricLeafNodeInterface {
	return v.value
}

func (v *NullableApicFabricLeafNodeInterface) Set(val *ApicFabricLeafNodeInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableApicFabricLeafNodeInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableApicFabricLeafNodeInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicFabricLeafNodeInterface(val *ApicFabricLeafNodeInterface) *NullableApicFabricLeafNodeInterface {
	return &NullableApicFabricLeafNodeInterface{value: val, isSet: true}
}

func (v NullableApicFabricLeafNodeInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicFabricLeafNodeInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}