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

// checks if the FirmwarePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwarePolicy{}

// FirmwarePolicy Firmware policy on the endpoint.
type FirmwarePolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	ExcludeComponentList []string                     `json:"ExcludeComponentList,omitempty"`
	ModelBundleCombo     []FirmwareModelBundleVersion `json:"ModelBundleCombo,omitempty"`
	// The target platform on which the policy to be applied. Either standalone or connected. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string                                      `json:"TargetPlatform,omitempty"`
	Organization   NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwarePolicy FirmwarePolicy

// NewFirmwarePolicy instantiates a new FirmwarePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwarePolicy(classId string, objectType string) *FirmwarePolicy {
	this := FirmwarePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewFirmwarePolicyWithDefaults instantiates a new FirmwarePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwarePolicyWithDefaults() *FirmwarePolicy {
	this := FirmwarePolicy{}
	var classId string = "firmware.Policy"
	this.ClassId = classId
	var objectType string = "firmware.Policy"
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwarePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwarePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwarePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.Policy" of the ClassId field.
func (o *FirmwarePolicy) GetDefaultClassId() interface{} {
	return "firmware.Policy"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwarePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwarePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwarePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.Policy" of the ObjectType field.
func (o *FirmwarePolicy) GetDefaultObjectType() interface{} {
	return "firmware.Policy"
}

// GetExcludeComponentList returns the ExcludeComponentList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicy) GetExcludeComponentList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeComponentList
}

// GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicy) GetExcludeComponentListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeComponentList) {
		return nil, false
	}
	return o.ExcludeComponentList, true
}

// HasExcludeComponentList returns a boolean if a field has been set.
func (o *FirmwarePolicy) HasExcludeComponentList() bool {
	if o != nil && !IsNil(o.ExcludeComponentList) {
		return true
	}

	return false
}

// SetExcludeComponentList gets a reference to the given []string and assigns it to the ExcludeComponentList field.
func (o *FirmwarePolicy) SetExcludeComponentList(v []string) {
	o.ExcludeComponentList = v
}

// GetModelBundleCombo returns the ModelBundleCombo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicy) GetModelBundleCombo() []FirmwareModelBundleVersion {
	if o == nil {
		var ret []FirmwareModelBundleVersion
		return ret
	}
	return o.ModelBundleCombo
}

// GetModelBundleComboOk returns a tuple with the ModelBundleCombo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicy) GetModelBundleComboOk() ([]FirmwareModelBundleVersion, bool) {
	if o == nil || IsNil(o.ModelBundleCombo) {
		return nil, false
	}
	return o.ModelBundleCombo, true
}

// HasModelBundleCombo returns a boolean if a field has been set.
func (o *FirmwarePolicy) HasModelBundleCombo() bool {
	if o != nil && !IsNil(o.ModelBundleCombo) {
		return true
	}

	return false
}

// SetModelBundleCombo gets a reference to the given []FirmwareModelBundleVersion and assigns it to the ModelBundleCombo field.
func (o *FirmwarePolicy) SetModelBundleCombo(v []FirmwareModelBundleVersion) {
	o.ModelBundleCombo = v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *FirmwarePolicy) GetTargetPlatform() string {
	if o == nil || IsNil(o.TargetPlatform) {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwarePolicy) GetTargetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPlatform) {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *FirmwarePolicy) HasTargetPlatform() bool {
	if o != nil && !IsNil(o.TargetPlatform) {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *FirmwarePolicy) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *FirmwarePolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FirmwarePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *FirmwarePolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *FirmwarePolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FirmwarePolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *FirmwarePolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o FirmwarePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwarePolicy) ToMap() (map[string]interface{}, error) {
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
	if o.ExcludeComponentList != nil {
		toSerialize["ExcludeComponentList"] = o.ExcludeComponentList
	}
	if o.ModelBundleCombo != nil {
		toSerialize["ModelBundleCombo"] = o.ModelBundleCombo
	}
	if !IsNil(o.TargetPlatform) {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
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

func (o *FirmwarePolicy) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwarePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                       `json:"ObjectType"`
		ExcludeComponentList []string                     `json:"ExcludeComponentList,omitempty"`
		ModelBundleCombo     []FirmwareModelBundleVersion `json:"ModelBundleCombo,omitempty"`
		// The target platform on which the policy to be applied. Either standalone or connected. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform *string                                      `json:"TargetPlatform,omitempty"`
		Organization   NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varFirmwarePolicyWithoutEmbeddedStruct := FirmwarePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwarePolicyWithoutEmbeddedStruct)
	if err == nil {
		varFirmwarePolicy := _FirmwarePolicy{}
		varFirmwarePolicy.ClassId = varFirmwarePolicyWithoutEmbeddedStruct.ClassId
		varFirmwarePolicy.ObjectType = varFirmwarePolicyWithoutEmbeddedStruct.ObjectType
		varFirmwarePolicy.ExcludeComponentList = varFirmwarePolicyWithoutEmbeddedStruct.ExcludeComponentList
		varFirmwarePolicy.ModelBundleCombo = varFirmwarePolicyWithoutEmbeddedStruct.ModelBundleCombo
		varFirmwarePolicy.TargetPlatform = varFirmwarePolicyWithoutEmbeddedStruct.TargetPlatform
		varFirmwarePolicy.Organization = varFirmwarePolicyWithoutEmbeddedStruct.Organization
		varFirmwarePolicy.Profiles = varFirmwarePolicyWithoutEmbeddedStruct.Profiles
		*o = FirmwarePolicy(varFirmwarePolicy)
	} else {
		return err
	}

	varFirmwarePolicy := _FirmwarePolicy{}

	err = json.Unmarshal(data, &varFirmwarePolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFirmwarePolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeComponentList")
		delete(additionalProperties, "ModelBundleCombo")
		delete(additionalProperties, "TargetPlatform")
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

type NullableFirmwarePolicy struct {
	value *FirmwarePolicy
	isSet bool
}

func (v NullableFirmwarePolicy) Get() *FirmwarePolicy {
	return v.value
}

func (v *NullableFirmwarePolicy) Set(val *FirmwarePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwarePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwarePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwarePolicy(val *FirmwarePolicy) *NullableFirmwarePolicy {
	return &NullableFirmwarePolicy{value: val, isSet: true}
}

func (v NullableFirmwarePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwarePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
