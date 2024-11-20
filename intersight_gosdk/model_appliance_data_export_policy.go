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

// checks if the ApplianceDataExportPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceDataExportPolicy{}

// ApplianceDataExportPolicy Data Export Policy is a category-based data collection policy that enables or disables data export (data collection) from the Intersight Appliance to the Intersight. The Data Export Policy configuration is organized hierarchically as follows.   Global:      Inventory:         Network         Storage      TechSupport When the DataExportPolicy for a category is enabled/disabled, all the sub-category configurations are enabled/disabled as well. For example, if you enable/disable Inventory, all its sub-category configurations (ie. Network and Storage) are also enabled/disabled.
type ApplianceDataExportPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the data collection mode. If the value is 'true', then data collection is enabled.
	Enable *bool `json:"Enable,omitempty"`
	// Name of the Data Export Policy.
	Name         *string                                       `json:"Name,omitempty"`
	Account      NullableIamAccountRelationship                `json:"Account,omitempty"`
	ParentConfig NullableApplianceDataExportPolicyRelationship `json:"ParentConfig,omitempty"`
	// An array of relationships to applianceDataExportPolicy resources.
	SubConfigs           []ApplianceDataExportPolicyRelationship `json:"SubConfigs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDataExportPolicy ApplianceDataExportPolicy

// NewApplianceDataExportPolicy instantiates a new ApplianceDataExportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDataExportPolicy(classId string, objectType string) *ApplianceDataExportPolicy {
	this := ApplianceDataExportPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceDataExportPolicyWithDefaults instantiates a new ApplianceDataExportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDataExportPolicyWithDefaults() *ApplianceDataExportPolicy {
	this := ApplianceDataExportPolicy{}
	var classId string = "appliance.DataExportPolicy"
	this.ClassId = classId
	var objectType string = "appliance.DataExportPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDataExportPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDataExportPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.DataExportPolicy" of the ClassId field.
func (o *ApplianceDataExportPolicy) GetDefaultClassId() interface{} {
	return "appliance.DataExportPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDataExportPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDataExportPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.DataExportPolicy" of the ObjectType field.
func (o *ApplianceDataExportPolicy) GetDefaultObjectType() interface{} {
	return "appliance.DataExportPolicy"
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ApplianceDataExportPolicy) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicy) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicy) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ApplianceDataExportPolicy) SetEnable(v bool) {
	o.Enable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceDataExportPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceDataExportPolicy) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDataExportPolicy) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDataExportPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicy) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *ApplianceDataExportPolicy) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *ApplianceDataExportPolicy) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ApplianceDataExportPolicy) UnsetAccount() {
	o.Account.Unset()
}

// GetParentConfig returns the ParentConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDataExportPolicy) GetParentConfig() ApplianceDataExportPolicyRelationship {
	if o == nil || IsNil(o.ParentConfig.Get()) {
		var ret ApplianceDataExportPolicyRelationship
		return ret
	}
	return *o.ParentConfig.Get()
}

// GetParentConfigOk returns a tuple with the ParentConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDataExportPolicy) GetParentConfigOk() (*ApplianceDataExportPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentConfig.Get(), o.ParentConfig.IsSet()
}

// HasParentConfig returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicy) HasParentConfig() bool {
	if o != nil && o.ParentConfig.IsSet() {
		return true
	}

	return false
}

// SetParentConfig gets a reference to the given NullableApplianceDataExportPolicyRelationship and assigns it to the ParentConfig field.
func (o *ApplianceDataExportPolicy) SetParentConfig(v ApplianceDataExportPolicyRelationship) {
	o.ParentConfig.Set(&v)
}

// SetParentConfigNil sets the value for ParentConfig to be an explicit nil
func (o *ApplianceDataExportPolicy) SetParentConfigNil() {
	o.ParentConfig.Set(nil)
}

// UnsetParentConfig ensures that no value is present for ParentConfig, not even an explicit nil
func (o *ApplianceDataExportPolicy) UnsetParentConfig() {
	o.ParentConfig.Unset()
}

// GetSubConfigs returns the SubConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDataExportPolicy) GetSubConfigs() []ApplianceDataExportPolicyRelationship {
	if o == nil {
		var ret []ApplianceDataExportPolicyRelationship
		return ret
	}
	return o.SubConfigs
}

// GetSubConfigsOk returns a tuple with the SubConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDataExportPolicy) GetSubConfigsOk() ([]ApplianceDataExportPolicyRelationship, bool) {
	if o == nil || IsNil(o.SubConfigs) {
		return nil, false
	}
	return o.SubConfigs, true
}

// HasSubConfigs returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicy) HasSubConfigs() bool {
	if o != nil && !IsNil(o.SubConfigs) {
		return true
	}

	return false
}

// SetSubConfigs gets a reference to the given []ApplianceDataExportPolicyRelationship and assigns it to the SubConfigs field.
func (o *ApplianceDataExportPolicy) SetSubConfigs(v []ApplianceDataExportPolicyRelationship) {
	o.SubConfigs = v
}

func (o ApplianceDataExportPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceDataExportPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Enable) {
		toSerialize["Enable"] = o.Enable
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.ParentConfig.IsSet() {
		toSerialize["ParentConfig"] = o.ParentConfig.Get()
	}
	if o.SubConfigs != nil {
		toSerialize["SubConfigs"] = o.SubConfigs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceDataExportPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceDataExportPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the data collection mode. If the value is 'true', then data collection is enabled.
		Enable *bool `json:"Enable,omitempty"`
		// Name of the Data Export Policy.
		Name         *string                                       `json:"Name,omitempty"`
		Account      NullableIamAccountRelationship                `json:"Account,omitempty"`
		ParentConfig NullableApplianceDataExportPolicyRelationship `json:"ParentConfig,omitempty"`
		// An array of relationships to applianceDataExportPolicy resources.
		SubConfigs []ApplianceDataExportPolicyRelationship `json:"SubConfigs,omitempty"`
	}

	varApplianceDataExportPolicyWithoutEmbeddedStruct := ApplianceDataExportPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceDataExportPolicyWithoutEmbeddedStruct)
	if err == nil {
		varApplianceDataExportPolicy := _ApplianceDataExportPolicy{}
		varApplianceDataExportPolicy.ClassId = varApplianceDataExportPolicyWithoutEmbeddedStruct.ClassId
		varApplianceDataExportPolicy.ObjectType = varApplianceDataExportPolicyWithoutEmbeddedStruct.ObjectType
		varApplianceDataExportPolicy.Enable = varApplianceDataExportPolicyWithoutEmbeddedStruct.Enable
		varApplianceDataExportPolicy.Name = varApplianceDataExportPolicyWithoutEmbeddedStruct.Name
		varApplianceDataExportPolicy.Account = varApplianceDataExportPolicyWithoutEmbeddedStruct.Account
		varApplianceDataExportPolicy.ParentConfig = varApplianceDataExportPolicyWithoutEmbeddedStruct.ParentConfig
		varApplianceDataExportPolicy.SubConfigs = varApplianceDataExportPolicyWithoutEmbeddedStruct.SubConfigs
		*o = ApplianceDataExportPolicy(varApplianceDataExportPolicy)
	} else {
		return err
	}

	varApplianceDataExportPolicy := _ApplianceDataExportPolicy{}

	err = json.Unmarshal(data, &varApplianceDataExportPolicy)
	if err == nil {
		o.MoBaseMo = varApplianceDataExportPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "ParentConfig")
		delete(additionalProperties, "SubConfigs")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableApplianceDataExportPolicy struct {
	value *ApplianceDataExportPolicy
	isSet bool
}

func (v NullableApplianceDataExportPolicy) Get() *ApplianceDataExportPolicy {
	return v.value
}

func (v *NullableApplianceDataExportPolicy) Set(val *ApplianceDataExportPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDataExportPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDataExportPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDataExportPolicy(val *ApplianceDataExportPolicy) *NullableApplianceDataExportPolicy {
	return &NullableApplianceDataExportPolicy{value: val, isSet: true}
}

func (v NullableApplianceDataExportPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDataExportPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
