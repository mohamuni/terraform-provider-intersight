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

// checks if the ApicFabricLeafNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApicFabricLeafNode{}

// ApicFabricLeafNode Cisco Application Policy Infrastructure Controller Fabric Leaf Nodes.
type ApicFabricLeafNode struct {
	ApicInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Address of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Address *string `json:"Address,omitempty"`
	// Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Dn                    *string                           `json:"Dn,omitempty"`
	FabricLeafNodeDetails NullableApicFabricLeafNodeDetails `json:"FabricLeafNodeDetails,omitempty"`
	// Name of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Name *string `json:"Name,omitempty"`
	// Pod of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
	Pod                  *string `json:"Pod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApicFabricLeafNode ApicFabricLeafNode

// NewApicFabricLeafNode instantiates a new ApicFabricLeafNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApicFabricLeafNode(classId string, objectType string) *ApicFabricLeafNode {
	this := ApicFabricLeafNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApicFabricLeafNodeWithDefaults instantiates a new ApicFabricLeafNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApicFabricLeafNodeWithDefaults() *ApicFabricLeafNode {
	this := ApicFabricLeafNode{}
	var classId string = "apic.FabricLeafNode"
	this.ClassId = classId
	var objectType string = "apic.FabricLeafNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApicFabricLeafNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApicFabricLeafNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApicFabricLeafNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApicFabricLeafNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApicFabricLeafNode) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNode) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApicFabricLeafNode) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ApicFabricLeafNode) SetAddress(v string) {
	o.Address = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *ApicFabricLeafNode) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNode) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ApicFabricLeafNode) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *ApicFabricLeafNode) SetDn(v string) {
	o.Dn = &v
}

// GetFabricLeafNodeDetails returns the FabricLeafNodeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApicFabricLeafNode) GetFabricLeafNodeDetails() ApicFabricLeafNodeDetails {
	if o == nil || IsNil(o.FabricLeafNodeDetails.Get()) {
		var ret ApicFabricLeafNodeDetails
		return ret
	}
	return *o.FabricLeafNodeDetails.Get()
}

// GetFabricLeafNodeDetailsOk returns a tuple with the FabricLeafNodeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApicFabricLeafNode) GetFabricLeafNodeDetailsOk() (*ApicFabricLeafNodeDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.FabricLeafNodeDetails.Get(), o.FabricLeafNodeDetails.IsSet()
}

// HasFabricLeafNodeDetails returns a boolean if a field has been set.
func (o *ApicFabricLeafNode) HasFabricLeafNodeDetails() bool {
	if o != nil && o.FabricLeafNodeDetails.IsSet() {
		return true
	}

	return false
}

// SetFabricLeafNodeDetails gets a reference to the given NullableApicFabricLeafNodeDetails and assigns it to the FabricLeafNodeDetails field.
func (o *ApicFabricLeafNode) SetFabricLeafNodeDetails(v ApicFabricLeafNodeDetails) {
	o.FabricLeafNodeDetails.Set(&v)
}

// SetFabricLeafNodeDetailsNil sets the value for FabricLeafNodeDetails to be an explicit nil
func (o *ApicFabricLeafNode) SetFabricLeafNodeDetailsNil() {
	o.FabricLeafNodeDetails.Set(nil)
}

// UnsetFabricLeafNodeDetails ensures that no value is present for FabricLeafNodeDetails, not even an explicit nil
func (o *ApicFabricLeafNode) UnsetFabricLeafNodeDetails() {
	o.FabricLeafNodeDetails.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApicFabricLeafNode) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApicFabricLeafNode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApicFabricLeafNode) SetName(v string) {
	o.Name = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *ApicFabricLeafNode) GetPod() string {
	if o == nil || IsNil(o.Pod) {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicFabricLeafNode) GetPodOk() (*string, bool) {
	if o == nil || IsNil(o.Pod) {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *ApicFabricLeafNode) HasPod() bool {
	if o != nil && !IsNil(o.Pod) {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *ApicFabricLeafNode) SetPod(v string) {
	o.Pod = &v
}

func (o ApicFabricLeafNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApicFabricLeafNode) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Address) {
		toSerialize["Address"] = o.Address
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if o.FabricLeafNodeDetails.IsSet() {
		toSerialize["FabricLeafNodeDetails"] = o.FabricLeafNodeDetails.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Pod) {
		toSerialize["Pod"] = o.Pod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApicFabricLeafNode) UnmarshalJSON(data []byte) (err error) {
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

	type ApicFabricLeafNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Address of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Address *string `json:"Address,omitempty"`
		// Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Dn                    *string                           `json:"Dn,omitempty"`
		FabricLeafNodeDetails NullableApicFabricLeafNodeDetails `json:"FabricLeafNodeDetails,omitempty"`
		// Name of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Name *string `json:"Name,omitempty"`
		// Pod of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI.
		Pod *string `json:"Pod,omitempty"`
	}

	varApicFabricLeafNodeWithoutEmbeddedStruct := ApicFabricLeafNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApicFabricLeafNodeWithoutEmbeddedStruct)
	if err == nil {
		varApicFabricLeafNode := _ApicFabricLeafNode{}
		varApicFabricLeafNode.ClassId = varApicFabricLeafNodeWithoutEmbeddedStruct.ClassId
		varApicFabricLeafNode.ObjectType = varApicFabricLeafNodeWithoutEmbeddedStruct.ObjectType
		varApicFabricLeafNode.Address = varApicFabricLeafNodeWithoutEmbeddedStruct.Address
		varApicFabricLeafNode.Dn = varApicFabricLeafNodeWithoutEmbeddedStruct.Dn
		varApicFabricLeafNode.FabricLeafNodeDetails = varApicFabricLeafNodeWithoutEmbeddedStruct.FabricLeafNodeDetails
		varApicFabricLeafNode.Name = varApicFabricLeafNodeWithoutEmbeddedStruct.Name
		varApicFabricLeafNode.Pod = varApicFabricLeafNodeWithoutEmbeddedStruct.Pod
		*o = ApicFabricLeafNode(varApicFabricLeafNode)
	} else {
		return err
	}

	varApicFabricLeafNode := _ApicFabricLeafNode{}

	err = json.Unmarshal(data, &varApicFabricLeafNode)
	if err == nil {
		o.ApicInventoryEntity = varApicFabricLeafNode.ApicInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FabricLeafNodeDetails")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Pod")

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

type NullableApicFabricLeafNode struct {
	value *ApicFabricLeafNode
	isSet bool
}

func (v NullableApicFabricLeafNode) Get() *ApicFabricLeafNode {
	return v.value
}

func (v *NullableApicFabricLeafNode) Set(val *ApicFabricLeafNode) {
	v.value = val
	v.isSet = true
}

func (v NullableApicFabricLeafNode) IsSet() bool {
	return v.isSet
}

func (v *NullableApicFabricLeafNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicFabricLeafNode(val *ApicFabricLeafNode) *NullableApicFabricLeafNode {
	return &NullableApicFabricLeafNode{value: val, isSet: true}
}

func (v NullableApicFabricLeafNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicFabricLeafNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}