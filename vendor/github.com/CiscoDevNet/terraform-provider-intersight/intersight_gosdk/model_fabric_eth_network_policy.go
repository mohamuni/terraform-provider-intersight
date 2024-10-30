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

// checks if the FabricEthNetworkPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricEthNetworkPolicy{}

// FabricEthNetworkPolicy A policy for all the Virtual LAN networks to be deployed on the Fabric Interconnect.
type FabricEthNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                       `json:"ObjectType"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricBaseSwitchProfile resources.
	Profiles             []FabricBaseSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEthNetworkPolicy FabricEthNetworkPolicy

// NewFabricEthNetworkPolicy instantiates a new FabricEthNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEthNetworkPolicy(classId string, objectType string) *FabricEthNetworkPolicy {
	this := FabricEthNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricEthNetworkPolicyWithDefaults instantiates a new FabricEthNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEthNetworkPolicyWithDefaults() *FabricEthNetworkPolicy {
	this := FabricEthNetworkPolicy{}
	var classId string = "fabric.EthNetworkPolicy"
	this.ClassId = classId
	var objectType string = "fabric.EthNetworkPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricEthNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricEthNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.EthNetworkPolicy" of the ClassId field.
func (o *FabricEthNetworkPolicy) GetDefaultClassId() interface{} {
	return "fabric.EthNetworkPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *FabricEthNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricEthNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.EthNetworkPolicy" of the ObjectType field.
func (o *FabricEthNetworkPolicy) GetDefaultObjectType() interface{} {
	return "fabric.EthNetworkPolicy"
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricEthNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricEthNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *FabricEthNetworkPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *FabricEthNetworkPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkPolicy) GetProfiles() []FabricBaseSwitchProfileRelationship {
	if o == nil {
		var ret []FabricBaseSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkPolicy) GetProfilesOk() ([]FabricBaseSwitchProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricEthNetworkPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricBaseSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricEthNetworkPolicy) SetProfiles(v []FabricBaseSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricEthNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricEthNetworkPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricEthNetworkPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type FabricEthNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                       `json:"ObjectType"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricBaseSwitchProfile resources.
		Profiles []FabricBaseSwitchProfileRelationship `json:"Profiles,omitempty"`
	}

	varFabricEthNetworkPolicyWithoutEmbeddedStruct := FabricEthNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricEthNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricEthNetworkPolicy := _FabricEthNetworkPolicy{}
		varFabricEthNetworkPolicy.ClassId = varFabricEthNetworkPolicyWithoutEmbeddedStruct.ClassId
		varFabricEthNetworkPolicy.ObjectType = varFabricEthNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varFabricEthNetworkPolicy.Organization = varFabricEthNetworkPolicyWithoutEmbeddedStruct.Organization
		varFabricEthNetworkPolicy.Profiles = varFabricEthNetworkPolicyWithoutEmbeddedStruct.Profiles
		*o = FabricEthNetworkPolicy(varFabricEthNetworkPolicy)
	} else {
		return err
	}

	varFabricEthNetworkPolicy := _FabricEthNetworkPolicy{}

	err = json.Unmarshal(data, &varFabricEthNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricEthNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableFabricEthNetworkPolicy struct {
	value *FabricEthNetworkPolicy
	isSet bool
}

func (v NullableFabricEthNetworkPolicy) Get() *FabricEthNetworkPolicy {
	return v.value
}

func (v *NullableFabricEthNetworkPolicy) Set(val *FabricEthNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEthNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEthNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEthNetworkPolicy(val *FabricEthNetworkPolicy) *NullableFabricEthNetworkPolicy {
	return &NullableFabricEthNetworkPolicy{value: val, isSet: true}
}

func (v NullableFabricEthNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEthNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
