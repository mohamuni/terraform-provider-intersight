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

// checks if the KubernetesNodeProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeProfile{}

// KubernetesNodeProfile A profile specifying configuration settings for a Kubernetes node.
type KubernetesNodeProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Cloud provider for this node profile. * `noProvider` - Enables the use of no cloud provider. * `external` - Out of tree cloud provider, e.g. CPI for vsphere.
	CloudProvider        *string                                        `json:"CloudProvider,omitempty"`
	ConfigResult         NullableKubernetesConfigResultRelationship     `json:"ConfigResult,omitempty"`
	NodeGroup            NullableKubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
	Target               NullableAssetDeviceRegistrationRelationship    `json:"Target,omitempty"`
	Version              NullableKubernetesVersionRelationship          `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeProfile KubernetesNodeProfile

// NewKubernetesNodeProfile instantiates a new KubernetesNodeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeProfile(classId string, objectType string) *KubernetesNodeProfile {
	this := KubernetesNodeProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// NewKubernetesNodeProfileWithDefaults instantiates a new KubernetesNodeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeProfileWithDefaults() *KubernetesNodeProfile {
	this := KubernetesNodeProfile{}
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KubernetesNodeProfile) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KubernetesNodeProfile) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeProfile) GetConfigResult() KubernetesConfigResultRelationship {
	if o == nil || IsNil(o.ConfigResult.Get()) {
		var ret KubernetesConfigResultRelationship
		return ret
	}
	return *o.ConfigResult.Get()
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProfile) GetConfigResultOk() (*KubernetesConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigResult.Get(), o.ConfigResult.IsSet()
}

// HasConfigResult returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult.IsSet() {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given NullableKubernetesConfigResultRelationship and assigns it to the ConfigResult field.
func (o *KubernetesNodeProfile) SetConfigResult(v KubernetesConfigResultRelationship) {
	o.ConfigResult.Set(&v)
}

// SetConfigResultNil sets the value for ConfigResult to be an explicit nil
func (o *KubernetesNodeProfile) SetConfigResultNil() {
	o.ConfigResult.Set(nil)
}

// UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil
func (o *KubernetesNodeProfile) UnsetConfigResult() {
	o.ConfigResult.Unset()
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeProfile) GetNodeGroup() KubernetesNodeGroupProfileRelationship {
	if o == nil || IsNil(o.NodeGroup.Get()) {
		var ret KubernetesNodeGroupProfileRelationship
		return ret
	}
	return *o.NodeGroup.Get()
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProfile) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeGroup.Get(), o.NodeGroup.IsSet()
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasNodeGroup() bool {
	if o != nil && o.NodeGroup.IsSet() {
		return true
	}

	return false
}

// SetNodeGroup gets a reference to the given NullableKubernetesNodeGroupProfileRelationship and assigns it to the NodeGroup field.
func (o *KubernetesNodeProfile) SetNodeGroup(v KubernetesNodeGroupProfileRelationship) {
	o.NodeGroup.Set(&v)
}

// SetNodeGroupNil sets the value for NodeGroup to be an explicit nil
func (o *KubernetesNodeProfile) SetNodeGroupNil() {
	o.NodeGroup.Set(nil)
}

// UnsetNodeGroup ensures that no value is present for NodeGroup, not even an explicit nil
func (o *KubernetesNodeProfile) UnsetNodeGroup() {
	o.NodeGroup.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeProfile) GetTarget() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Target.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProfile) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Target field.
func (o *KubernetesNodeProfile) SetTarget(v AssetDeviceRegistrationRelationship) {
	o.Target.Set(&v)
}

// SetTargetNil sets the value for Target to be an explicit nil
func (o *KubernetesNodeProfile) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *KubernetesNodeProfile) UnsetTarget() {
	o.Target.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeProfile) GetVersion() KubernetesVersionRelationship {
	if o == nil || IsNil(o.Version.Get()) {
		var ret KubernetesVersionRelationship
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProfile) GetVersionOk() (*KubernetesVersionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableKubernetesVersionRelationship and assigns it to the Version field.
func (o *KubernetesNodeProfile) SetVersion(v KubernetesVersionRelationship) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *KubernetesNodeProfile) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *KubernetesNodeProfile) UnsetVersion() {
	o.Version.Unset()
}

func (o KubernetesNodeProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigProfile
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CloudProvider) {
		toSerialize["CloudProvider"] = o.CloudProvider
	}
	if o.ConfigResult.IsSet() {
		toSerialize["ConfigResult"] = o.ConfigResult.Get()
	}
	if o.NodeGroup.IsSet() {
		toSerialize["NodeGroup"] = o.NodeGroup.Get()
	}
	if o.Target.IsSet() {
		toSerialize["Target"] = o.Target.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesNodeProfile) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesNodeProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Cloud provider for this node profile. * `noProvider` - Enables the use of no cloud provider. * `external` - Out of tree cloud provider, e.g. CPI for vsphere.
		CloudProvider *string                                        `json:"CloudProvider,omitempty"`
		ConfigResult  NullableKubernetesConfigResultRelationship     `json:"ConfigResult,omitempty"`
		NodeGroup     NullableKubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
		Target        NullableAssetDeviceRegistrationRelationship    `json:"Target,omitempty"`
		Version       NullableKubernetesVersionRelationship          `json:"Version,omitempty"`
	}

	varKubernetesNodeProfileWithoutEmbeddedStruct := KubernetesNodeProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesNodeProfileWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNodeProfile := _KubernetesNodeProfile{}
		varKubernetesNodeProfile.ClassId = varKubernetesNodeProfileWithoutEmbeddedStruct.ClassId
		varKubernetesNodeProfile.ObjectType = varKubernetesNodeProfileWithoutEmbeddedStruct.ObjectType
		varKubernetesNodeProfile.CloudProvider = varKubernetesNodeProfileWithoutEmbeddedStruct.CloudProvider
		varKubernetesNodeProfile.ConfigResult = varKubernetesNodeProfileWithoutEmbeddedStruct.ConfigResult
		varKubernetesNodeProfile.NodeGroup = varKubernetesNodeProfileWithoutEmbeddedStruct.NodeGroup
		varKubernetesNodeProfile.Target = varKubernetesNodeProfileWithoutEmbeddedStruct.Target
		varKubernetesNodeProfile.Version = varKubernetesNodeProfileWithoutEmbeddedStruct.Version
		*o = KubernetesNodeProfile(varKubernetesNodeProfile)
	} else {
		return err
	}

	varKubernetesNodeProfile := _KubernetesNodeProfile{}

	err = json.Unmarshal(data, &varKubernetesNodeProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varKubernetesNodeProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CloudProvider")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "NodeGroup")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "Version")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableKubernetesNodeProfile struct {
	value *KubernetesNodeProfile
	isSet bool
}

func (v NullableKubernetesNodeProfile) Get() *KubernetesNodeProfile {
	return v.value
}

func (v *NullableKubernetesNodeProfile) Set(val *KubernetesNodeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeProfile(val *KubernetesNodeProfile) *NullableKubernetesNodeProfile {
	return &NullableKubernetesNodeProfile{value: val, isSet: true}
}

func (v NullableKubernetesNodeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
