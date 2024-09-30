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

// checks if the CapabilityFeatureConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityFeatureConfig{}

// CapabilityFeatureConfig Configuration specific to the adapter feature.
type CapabilityFeatureConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the feature that identifies the specific adapter configuration. * `RoCEv2` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2. * `RoCEv1` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1. * `VMQ` - Capability indicator of the Virtual Machine Queue (VMQ) feature. * `VMMQ` - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature. * `VMQInterrupts` - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature. * `NVGRE` - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature. * `ARFS` - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature. * `VXLAN` - Capability indicator of the Virtual Extensible LAN (VXLAN) feature. * `usNIC` - Capability indicator of the User Space NIC (usNIC) feature. * `Advanced Filter` - Capability indicator of the Advanced Filter feature. * `Azure Stack Host QOS` - Capability indicator of the Azure Stack Host QOS feature. * `GENEVE Offload` - Capability indicator of the Generic Network Virtualization Encapsulation (Geneve) Offload feature. * `QinQ` - Capability indicator of the QinQ feature. * `SRIOV` - Capability indicator of the Single Root Input Output Virtualization (SR-IOV).
	FeatureName *string `json:"FeatureName,omitempty"`
	// Firmware version of Adapter from which support for this feature is available.
	MinAdapterFwVersion *string `json:"MinAdapterFwVersion,omitempty"`
	// Firmware version of BMC from which support for this feature is available.
	MinFwVersion             *string                              `json:"MinFwVersion,omitempty"`
	SupportedFwVersions      []string                             `json:"SupportedFwVersions,omitempty"`
	SupportedInAdapters      []string                             `json:"SupportedInAdapters,omitempty"`
	SupportedInGenerations   []int32                              `json:"SupportedInGenerations,omitempty"`
	UnsupportedFeatureMatrix []CapabilityUnsupportedFeatureConfig `json:"UnsupportedFeatureMatrix,omitempty"`
	// Action to be taken when validation does not succeed. * `Error` - Stop workflow execution by throwing error. * `Skip` - Remove the feature from configuration and continue workflow execution.
	ValidationAction     *string `json:"ValidationAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityFeatureConfig CapabilityFeatureConfig

// NewCapabilityFeatureConfig instantiates a new CapabilityFeatureConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityFeatureConfig(classId string, objectType string) *CapabilityFeatureConfig {
	this := CapabilityFeatureConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var featureName string = "RoCEv2"
	this.FeatureName = &featureName
	var validationAction string = "Error"
	this.ValidationAction = &validationAction
	return &this
}

// NewCapabilityFeatureConfigWithDefaults instantiates a new CapabilityFeatureConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityFeatureConfigWithDefaults() *CapabilityFeatureConfig {
	this := CapabilityFeatureConfig{}
	var classId string = "capability.FeatureConfig"
	this.ClassId = classId
	var objectType string = "capability.FeatureConfig"
	this.ObjectType = objectType
	var featureName string = "RoCEv2"
	this.FeatureName = &featureName
	var validationAction string = "Error"
	this.ValidationAction = &validationAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityFeatureConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityFeatureConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityFeatureConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.FeatureConfig" of the ClassId field.
func (o *CapabilityFeatureConfig) GetDefaultClassId() interface{} {
	return "capability.FeatureConfig"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityFeatureConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityFeatureConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityFeatureConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.FeatureConfig" of the ObjectType field.
func (o *CapabilityFeatureConfig) GetDefaultObjectType() interface{} {
	return "capability.FeatureConfig"
}

// GetFeatureName returns the FeatureName field value if set, zero value otherwise.
func (o *CapabilityFeatureConfig) GetFeatureName() string {
	if o == nil || IsNil(o.FeatureName) {
		var ret string
		return ret
	}
	return *o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFeatureConfig) GetFeatureNameOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureName) {
		return nil, false
	}
	return o.FeatureName, true
}

// HasFeatureName returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasFeatureName() bool {
	if o != nil && !IsNil(o.FeatureName) {
		return true
	}

	return false
}

// SetFeatureName gets a reference to the given string and assigns it to the FeatureName field.
func (o *CapabilityFeatureConfig) SetFeatureName(v string) {
	o.FeatureName = &v
}

// GetMinAdapterFwVersion returns the MinAdapterFwVersion field value if set, zero value otherwise.
func (o *CapabilityFeatureConfig) GetMinAdapterFwVersion() string {
	if o == nil || IsNil(o.MinAdapterFwVersion) {
		var ret string
		return ret
	}
	return *o.MinAdapterFwVersion
}

// GetMinAdapterFwVersionOk returns a tuple with the MinAdapterFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFeatureConfig) GetMinAdapterFwVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinAdapterFwVersion) {
		return nil, false
	}
	return o.MinAdapterFwVersion, true
}

// HasMinAdapterFwVersion returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasMinAdapterFwVersion() bool {
	if o != nil && !IsNil(o.MinAdapterFwVersion) {
		return true
	}

	return false
}

// SetMinAdapterFwVersion gets a reference to the given string and assigns it to the MinAdapterFwVersion field.
func (o *CapabilityFeatureConfig) SetMinAdapterFwVersion(v string) {
	o.MinAdapterFwVersion = &v
}

// GetMinFwVersion returns the MinFwVersion field value if set, zero value otherwise.
func (o *CapabilityFeatureConfig) GetMinFwVersion() string {
	if o == nil || IsNil(o.MinFwVersion) {
		var ret string
		return ret
	}
	return *o.MinFwVersion
}

// GetMinFwVersionOk returns a tuple with the MinFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFeatureConfig) GetMinFwVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinFwVersion) {
		return nil, false
	}
	return o.MinFwVersion, true
}

// HasMinFwVersion returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasMinFwVersion() bool {
	if o != nil && !IsNil(o.MinFwVersion) {
		return true
	}

	return false
}

// SetMinFwVersion gets a reference to the given string and assigns it to the MinFwVersion field.
func (o *CapabilityFeatureConfig) SetMinFwVersion(v string) {
	o.MinFwVersion = &v
}

// GetSupportedFwVersions returns the SupportedFwVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityFeatureConfig) GetSupportedFwVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedFwVersions
}

// GetSupportedFwVersionsOk returns a tuple with the SupportedFwVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityFeatureConfig) GetSupportedFwVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedFwVersions) {
		return nil, false
	}
	return o.SupportedFwVersions, true
}

// HasSupportedFwVersions returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasSupportedFwVersions() bool {
	if o != nil && !IsNil(o.SupportedFwVersions) {
		return true
	}

	return false
}

// SetSupportedFwVersions gets a reference to the given []string and assigns it to the SupportedFwVersions field.
func (o *CapabilityFeatureConfig) SetSupportedFwVersions(v []string) {
	o.SupportedFwVersions = v
}

// GetSupportedInAdapters returns the SupportedInAdapters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityFeatureConfig) GetSupportedInAdapters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedInAdapters
}

// GetSupportedInAdaptersOk returns a tuple with the SupportedInAdapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityFeatureConfig) GetSupportedInAdaptersOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedInAdapters) {
		return nil, false
	}
	return o.SupportedInAdapters, true
}

// HasSupportedInAdapters returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasSupportedInAdapters() bool {
	if o != nil && !IsNil(o.SupportedInAdapters) {
		return true
	}

	return false
}

// SetSupportedInAdapters gets a reference to the given []string and assigns it to the SupportedInAdapters field.
func (o *CapabilityFeatureConfig) SetSupportedInAdapters(v []string) {
	o.SupportedInAdapters = v
}

// GetSupportedInGenerations returns the SupportedInGenerations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityFeatureConfig) GetSupportedInGenerations() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.SupportedInGenerations
}

// GetSupportedInGenerationsOk returns a tuple with the SupportedInGenerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityFeatureConfig) GetSupportedInGenerationsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SupportedInGenerations) {
		return nil, false
	}
	return o.SupportedInGenerations, true
}

// HasSupportedInGenerations returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasSupportedInGenerations() bool {
	if o != nil && !IsNil(o.SupportedInGenerations) {
		return true
	}

	return false
}

// SetSupportedInGenerations gets a reference to the given []int32 and assigns it to the SupportedInGenerations field.
func (o *CapabilityFeatureConfig) SetSupportedInGenerations(v []int32) {
	o.SupportedInGenerations = v
}

// GetUnsupportedFeatureMatrix returns the UnsupportedFeatureMatrix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityFeatureConfig) GetUnsupportedFeatureMatrix() []CapabilityUnsupportedFeatureConfig {
	if o == nil {
		var ret []CapabilityUnsupportedFeatureConfig
		return ret
	}
	return o.UnsupportedFeatureMatrix
}

// GetUnsupportedFeatureMatrixOk returns a tuple with the UnsupportedFeatureMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityFeatureConfig) GetUnsupportedFeatureMatrixOk() ([]CapabilityUnsupportedFeatureConfig, bool) {
	if o == nil || IsNil(o.UnsupportedFeatureMatrix) {
		return nil, false
	}
	return o.UnsupportedFeatureMatrix, true
}

// HasUnsupportedFeatureMatrix returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasUnsupportedFeatureMatrix() bool {
	if o != nil && !IsNil(o.UnsupportedFeatureMatrix) {
		return true
	}

	return false
}

// SetUnsupportedFeatureMatrix gets a reference to the given []CapabilityUnsupportedFeatureConfig and assigns it to the UnsupportedFeatureMatrix field.
func (o *CapabilityFeatureConfig) SetUnsupportedFeatureMatrix(v []CapabilityUnsupportedFeatureConfig) {
	o.UnsupportedFeatureMatrix = v
}

// GetValidationAction returns the ValidationAction field value if set, zero value otherwise.
func (o *CapabilityFeatureConfig) GetValidationAction() string {
	if o == nil || IsNil(o.ValidationAction) {
		var ret string
		return ret
	}
	return *o.ValidationAction
}

// GetValidationActionOk returns a tuple with the ValidationAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFeatureConfig) GetValidationActionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationAction) {
		return nil, false
	}
	return o.ValidationAction, true
}

// HasValidationAction returns a boolean if a field has been set.
func (o *CapabilityFeatureConfig) HasValidationAction() bool {
	if o != nil && !IsNil(o.ValidationAction) {
		return true
	}

	return false
}

// SetValidationAction gets a reference to the given string and assigns it to the ValidationAction field.
func (o *CapabilityFeatureConfig) SetValidationAction(v string) {
	o.ValidationAction = &v
}

func (o CapabilityFeatureConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityFeatureConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FeatureName) {
		toSerialize["FeatureName"] = o.FeatureName
	}
	if !IsNil(o.MinAdapterFwVersion) {
		toSerialize["MinAdapterFwVersion"] = o.MinAdapterFwVersion
	}
	if !IsNil(o.MinFwVersion) {
		toSerialize["MinFwVersion"] = o.MinFwVersion
	}
	if o.SupportedFwVersions != nil {
		toSerialize["SupportedFwVersions"] = o.SupportedFwVersions
	}
	if o.SupportedInAdapters != nil {
		toSerialize["SupportedInAdapters"] = o.SupportedInAdapters
	}
	if o.SupportedInGenerations != nil {
		toSerialize["SupportedInGenerations"] = o.SupportedInGenerations
	}
	if o.UnsupportedFeatureMatrix != nil {
		toSerialize["UnsupportedFeatureMatrix"] = o.UnsupportedFeatureMatrix
	}
	if !IsNil(o.ValidationAction) {
		toSerialize["ValidationAction"] = o.ValidationAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityFeatureConfig) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityFeatureConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the feature that identifies the specific adapter configuration. * `RoCEv2` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2. * `RoCEv1` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1. * `VMQ` - Capability indicator of the Virtual Machine Queue (VMQ) feature. * `VMMQ` - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature. * `VMQInterrupts` - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature. * `NVGRE` - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature. * `ARFS` - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature. * `VXLAN` - Capability indicator of the Virtual Extensible LAN (VXLAN) feature. * `usNIC` - Capability indicator of the User Space NIC (usNIC) feature. * `Advanced Filter` - Capability indicator of the Advanced Filter feature. * `Azure Stack Host QOS` - Capability indicator of the Azure Stack Host QOS feature. * `GENEVE Offload` - Capability indicator of the Generic Network Virtualization Encapsulation (Geneve) Offload feature. * `QinQ` - Capability indicator of the QinQ feature. * `SRIOV` - Capability indicator of the Single Root Input Output Virtualization (SR-IOV).
		FeatureName *string `json:"FeatureName,omitempty"`
		// Firmware version of Adapter from which support for this feature is available.
		MinAdapterFwVersion *string `json:"MinAdapterFwVersion,omitempty"`
		// Firmware version of BMC from which support for this feature is available.
		MinFwVersion             *string                              `json:"MinFwVersion,omitempty"`
		SupportedFwVersions      []string                             `json:"SupportedFwVersions,omitempty"`
		SupportedInAdapters      []string                             `json:"SupportedInAdapters,omitempty"`
		SupportedInGenerations   []int32                              `json:"SupportedInGenerations,omitempty"`
		UnsupportedFeatureMatrix []CapabilityUnsupportedFeatureConfig `json:"UnsupportedFeatureMatrix,omitempty"`
		// Action to be taken when validation does not succeed. * `Error` - Stop workflow execution by throwing error. * `Skip` - Remove the feature from configuration and continue workflow execution.
		ValidationAction *string `json:"ValidationAction,omitempty"`
	}

	varCapabilityFeatureConfigWithoutEmbeddedStruct := CapabilityFeatureConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityFeatureConfigWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityFeatureConfig := _CapabilityFeatureConfig{}
		varCapabilityFeatureConfig.ClassId = varCapabilityFeatureConfigWithoutEmbeddedStruct.ClassId
		varCapabilityFeatureConfig.ObjectType = varCapabilityFeatureConfigWithoutEmbeddedStruct.ObjectType
		varCapabilityFeatureConfig.FeatureName = varCapabilityFeatureConfigWithoutEmbeddedStruct.FeatureName
		varCapabilityFeatureConfig.MinAdapterFwVersion = varCapabilityFeatureConfigWithoutEmbeddedStruct.MinAdapterFwVersion
		varCapabilityFeatureConfig.MinFwVersion = varCapabilityFeatureConfigWithoutEmbeddedStruct.MinFwVersion
		varCapabilityFeatureConfig.SupportedFwVersions = varCapabilityFeatureConfigWithoutEmbeddedStruct.SupportedFwVersions
		varCapabilityFeatureConfig.SupportedInAdapters = varCapabilityFeatureConfigWithoutEmbeddedStruct.SupportedInAdapters
		varCapabilityFeatureConfig.SupportedInGenerations = varCapabilityFeatureConfigWithoutEmbeddedStruct.SupportedInGenerations
		varCapabilityFeatureConfig.UnsupportedFeatureMatrix = varCapabilityFeatureConfigWithoutEmbeddedStruct.UnsupportedFeatureMatrix
		varCapabilityFeatureConfig.ValidationAction = varCapabilityFeatureConfigWithoutEmbeddedStruct.ValidationAction
		*o = CapabilityFeatureConfig(varCapabilityFeatureConfig)
	} else {
		return err
	}

	varCapabilityFeatureConfig := _CapabilityFeatureConfig{}

	err = json.Unmarshal(data, &varCapabilityFeatureConfig)
	if err == nil {
		o.MoBaseComplexType = varCapabilityFeatureConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FeatureName")
		delete(additionalProperties, "MinAdapterFwVersion")
		delete(additionalProperties, "MinFwVersion")
		delete(additionalProperties, "SupportedFwVersions")
		delete(additionalProperties, "SupportedInAdapters")
		delete(additionalProperties, "SupportedInGenerations")
		delete(additionalProperties, "UnsupportedFeatureMatrix")
		delete(additionalProperties, "ValidationAction")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableCapabilityFeatureConfig struct {
	value *CapabilityFeatureConfig
	isSet bool
}

func (v NullableCapabilityFeatureConfig) Get() *CapabilityFeatureConfig {
	return v.value
}

func (v *NullableCapabilityFeatureConfig) Set(val *CapabilityFeatureConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityFeatureConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityFeatureConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityFeatureConfig(val *CapabilityFeatureConfig) *NullableCapabilityFeatureConfig {
	return &NullableCapabilityFeatureConfig{value: val, isSet: true}
}

func (v NullableCapabilityFeatureConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityFeatureConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
