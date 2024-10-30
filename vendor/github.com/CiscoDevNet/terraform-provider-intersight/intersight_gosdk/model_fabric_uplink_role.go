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

// checks if the FabricUplinkRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricUplinkRole{}

// FabricUplinkRole Configuration object sent by user to create a uplink port.
type FabricUplinkRole struct {
	FabricTransceiverRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to fabricEthNetworkGroupPolicy resources.
	EthNetworkGroupPolicy []FabricEthNetworkGroupPolicyRelationship   `json:"EthNetworkGroupPolicy,omitempty"`
	FlowControlPolicy     NullableFabricFlowControlPolicyRelationship `json:"FlowControlPolicy,omitempty"`
	LinkControlPolicy     NullableFabricLinkControlPolicyRelationship `json:"LinkControlPolicy,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _FabricUplinkRole FabricUplinkRole

// NewFabricUplinkRole instantiates a new FabricUplinkRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUplinkRole(classId string, objectType string) *FabricUplinkRole {
	this := FabricUplinkRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// NewFabricUplinkRoleWithDefaults instantiates a new FabricUplinkRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUplinkRoleWithDefaults() *FabricUplinkRole {
	this := FabricUplinkRole{}
	var classId string = "fabric.UplinkRole"
	this.ClassId = classId
	var objectType string = "fabric.UplinkRole"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricUplinkRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricUplinkRole) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.UplinkRole" of the ClassId field.
func (o *FabricUplinkRole) GetDefaultClassId() interface{} {
	return "fabric.UplinkRole"
}

// GetObjectType returns the ObjectType field value
func (o *FabricUplinkRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricUplinkRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.UplinkRole" of the ObjectType field.
func (o *FabricUplinkRole) GetDefaultObjectType() interface{} {
	return "fabric.UplinkRole"
}

// GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkRole) GetEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship {
	if o == nil {
		var ret []FabricEthNetworkGroupPolicyRelationship
		return ret
	}
	return o.EthNetworkGroupPolicy
}

// GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkRole) GetEthNetworkGroupPolicyOk() ([]FabricEthNetworkGroupPolicyRelationship, bool) {
	if o == nil || IsNil(o.EthNetworkGroupPolicy) {
		return nil, false
	}
	return o.EthNetworkGroupPolicy, true
}

// HasEthNetworkGroupPolicy returns a boolean if a field has been set.
func (o *FabricUplinkRole) HasEthNetworkGroupPolicy() bool {
	if o != nil && !IsNil(o.EthNetworkGroupPolicy) {
		return true
	}

	return false
}

// SetEthNetworkGroupPolicy gets a reference to the given []FabricEthNetworkGroupPolicyRelationship and assigns it to the EthNetworkGroupPolicy field.
func (o *FabricUplinkRole) SetEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship) {
	o.EthNetworkGroupPolicy = v
}

// GetFlowControlPolicy returns the FlowControlPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkRole) GetFlowControlPolicy() FabricFlowControlPolicyRelationship {
	if o == nil || IsNil(o.FlowControlPolicy.Get()) {
		var ret FabricFlowControlPolicyRelationship
		return ret
	}
	return *o.FlowControlPolicy.Get()
}

// GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkRole) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowControlPolicy.Get(), o.FlowControlPolicy.IsSet()
}

// HasFlowControlPolicy returns a boolean if a field has been set.
func (o *FabricUplinkRole) HasFlowControlPolicy() bool {
	if o != nil && o.FlowControlPolicy.IsSet() {
		return true
	}

	return false
}

// SetFlowControlPolicy gets a reference to the given NullableFabricFlowControlPolicyRelationship and assigns it to the FlowControlPolicy field.
func (o *FabricUplinkRole) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship) {
	o.FlowControlPolicy.Set(&v)
}

// SetFlowControlPolicyNil sets the value for FlowControlPolicy to be an explicit nil
func (o *FabricUplinkRole) SetFlowControlPolicyNil() {
	o.FlowControlPolicy.Set(nil)
}

// UnsetFlowControlPolicy ensures that no value is present for FlowControlPolicy, not even an explicit nil
func (o *FabricUplinkRole) UnsetFlowControlPolicy() {
	o.FlowControlPolicy.Unset()
}

// GetLinkControlPolicy returns the LinkControlPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkRole) GetLinkControlPolicy() FabricLinkControlPolicyRelationship {
	if o == nil || IsNil(o.LinkControlPolicy.Get()) {
		var ret FabricLinkControlPolicyRelationship
		return ret
	}
	return *o.LinkControlPolicy.Get()
}

// GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkRole) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkControlPolicy.Get(), o.LinkControlPolicy.IsSet()
}

// HasLinkControlPolicy returns a boolean if a field has been set.
func (o *FabricUplinkRole) HasLinkControlPolicy() bool {
	if o != nil && o.LinkControlPolicy.IsSet() {
		return true
	}

	return false
}

// SetLinkControlPolicy gets a reference to the given NullableFabricLinkControlPolicyRelationship and assigns it to the LinkControlPolicy field.
func (o *FabricUplinkRole) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship) {
	o.LinkControlPolicy.Set(&v)
}

// SetLinkControlPolicyNil sets the value for LinkControlPolicy to be an explicit nil
func (o *FabricUplinkRole) SetLinkControlPolicyNil() {
	o.LinkControlPolicy.Set(nil)
}

// UnsetLinkControlPolicy ensures that no value is present for LinkControlPolicy, not even an explicit nil
func (o *FabricUplinkRole) UnsetLinkControlPolicy() {
	o.LinkControlPolicy.Unset()
}

func (o FabricUplinkRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricUplinkRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricTransceiverRole, errFabricTransceiverRole := json.Marshal(o.FabricTransceiverRole)
	if errFabricTransceiverRole != nil {
		return map[string]interface{}{}, errFabricTransceiverRole
	}
	errFabricTransceiverRole = json.Unmarshal([]byte(serializedFabricTransceiverRole), &toSerialize)
	if errFabricTransceiverRole != nil {
		return map[string]interface{}{}, errFabricTransceiverRole
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.EthNetworkGroupPolicy != nil {
		toSerialize["EthNetworkGroupPolicy"] = o.EthNetworkGroupPolicy
	}
	if o.FlowControlPolicy.IsSet() {
		toSerialize["FlowControlPolicy"] = o.FlowControlPolicy.Get()
	}
	if o.LinkControlPolicy.IsSet() {
		toSerialize["LinkControlPolicy"] = o.LinkControlPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricUplinkRole) UnmarshalJSON(data []byte) (err error) {
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
	type FabricUplinkRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to fabricEthNetworkGroupPolicy resources.
		EthNetworkGroupPolicy []FabricEthNetworkGroupPolicyRelationship   `json:"EthNetworkGroupPolicy,omitempty"`
		FlowControlPolicy     NullableFabricFlowControlPolicyRelationship `json:"FlowControlPolicy,omitempty"`
		LinkControlPolicy     NullableFabricLinkControlPolicyRelationship `json:"LinkControlPolicy,omitempty"`
	}

	varFabricUplinkRoleWithoutEmbeddedStruct := FabricUplinkRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricUplinkRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricUplinkRole := _FabricUplinkRole{}
		varFabricUplinkRole.ClassId = varFabricUplinkRoleWithoutEmbeddedStruct.ClassId
		varFabricUplinkRole.ObjectType = varFabricUplinkRoleWithoutEmbeddedStruct.ObjectType
		varFabricUplinkRole.EthNetworkGroupPolicy = varFabricUplinkRoleWithoutEmbeddedStruct.EthNetworkGroupPolicy
		varFabricUplinkRole.FlowControlPolicy = varFabricUplinkRoleWithoutEmbeddedStruct.FlowControlPolicy
		varFabricUplinkRole.LinkControlPolicy = varFabricUplinkRoleWithoutEmbeddedStruct.LinkControlPolicy
		*o = FabricUplinkRole(varFabricUplinkRole)
	} else {
		return err
	}

	varFabricUplinkRole := _FabricUplinkRole{}

	err = json.Unmarshal(data, &varFabricUplinkRole)
	if err == nil {
		o.FabricTransceiverRole = varFabricUplinkRole.FabricTransceiverRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EthNetworkGroupPolicy")
		delete(additionalProperties, "FlowControlPolicy")
		delete(additionalProperties, "LinkControlPolicy")

		// remove fields from embedded structs
		reflectFabricTransceiverRole := reflect.ValueOf(o.FabricTransceiverRole)
		for i := 0; i < reflectFabricTransceiverRole.Type().NumField(); i++ {
			t := reflectFabricTransceiverRole.Type().Field(i)

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

type NullableFabricUplinkRole struct {
	value *FabricUplinkRole
	isSet bool
}

func (v NullableFabricUplinkRole) Get() *FabricUplinkRole {
	return v.value
}

func (v *NullableFabricUplinkRole) Set(val *FabricUplinkRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUplinkRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUplinkRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUplinkRole(val *FabricUplinkRole) *NullableFabricUplinkRole {
	return &NullableFabricUplinkRole{value: val, isSet: true}
}

func (v NullableFabricUplinkRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUplinkRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
