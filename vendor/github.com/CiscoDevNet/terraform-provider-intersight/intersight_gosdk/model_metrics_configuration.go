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

// checks if the MetricsConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsConfiguration{}

// MetricsConfiguration Controls the behavior of metric collection for an account, configuration includes options to enable or disable collection as well how new devices claimed to the account are handled.
type MetricsConfiguration struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The behavior of the system when new resources are added, controls whether metric collection are automatically enabled for the new resources. * `AutoEnable` - Automatically enable metric collection for new resources, up to the limit of resource collection. * `Disabled` - Metrics will not be enabled on new resources, to enable collection requires an explicit user enable.
	CollectNewDevices *string `json:"CollectNewDevices,omitempty"`
	// Enables metric collection for the account, if disabled metrics will be stopped for all resources in the account.
	Enabled *bool `json:"Enabled,omitempty"`
	// The total number of resources that can be enabled for metric collection in this account.
	Limit                *int64                         `json:"Limit,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetricsConfiguration MetricsConfiguration

// NewMetricsConfiguration instantiates a new MetricsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsConfiguration(classId string, objectType string) *MetricsConfiguration {
	this := MetricsConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	var collectNewDevices string = "AutoEnable"
	this.CollectNewDevices = &collectNewDevices
	return &this
}

// NewMetricsConfigurationWithDefaults instantiates a new MetricsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsConfigurationWithDefaults() *MetricsConfiguration {
	this := MetricsConfiguration{}
	var classId string = "metrics.Configuration"
	this.ClassId = classId
	var objectType string = "metrics.Configuration"
	this.ObjectType = objectType
	var collectNewDevices string = "AutoEnable"
	this.CollectNewDevices = &collectNewDevices
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetricsConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetricsConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetricsConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "metrics.Configuration" of the ClassId field.
func (o *MetricsConfiguration) GetDefaultClassId() interface{} {
	return "metrics.Configuration"
}

// GetObjectType returns the ObjectType field value
func (o *MetricsConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetricsConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetricsConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "metrics.Configuration" of the ObjectType field.
func (o *MetricsConfiguration) GetDefaultObjectType() interface{} {
	return "metrics.Configuration"
}

// GetCollectNewDevices returns the CollectNewDevices field value if set, zero value otherwise.
func (o *MetricsConfiguration) GetCollectNewDevices() string {
	if o == nil || IsNil(o.CollectNewDevices) {
		var ret string
		return ret
	}
	return *o.CollectNewDevices
}

// GetCollectNewDevicesOk returns a tuple with the CollectNewDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsConfiguration) GetCollectNewDevicesOk() (*string, bool) {
	if o == nil || IsNil(o.CollectNewDevices) {
		return nil, false
	}
	return o.CollectNewDevices, true
}

// HasCollectNewDevices returns a boolean if a field has been set.
func (o *MetricsConfiguration) HasCollectNewDevices() bool {
	if o != nil && !IsNil(o.CollectNewDevices) {
		return true
	}

	return false
}

// SetCollectNewDevices gets a reference to the given string and assigns it to the CollectNewDevices field.
func (o *MetricsConfiguration) SetCollectNewDevices(v string) {
	o.CollectNewDevices = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MetricsConfiguration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MetricsConfiguration) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MetricsConfiguration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MetricsConfiguration) GetLimit() int64 {
	if o == nil || IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsConfiguration) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MetricsConfiguration) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *MetricsConfiguration) SetLimit(v int64) {
	o.Limit = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsConfiguration) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsConfiguration) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *MetricsConfiguration) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *MetricsConfiguration) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *MetricsConfiguration) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *MetricsConfiguration) UnsetAccount() {
	o.Account.Unset()
}

func (o MetricsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsConfiguration) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CollectNewDevices) {
		toSerialize["CollectNewDevices"] = o.CollectNewDevices
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.Limit) {
		toSerialize["Limit"] = o.Limit
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricsConfiguration) UnmarshalJSON(data []byte) (err error) {
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
	type MetricsConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The behavior of the system when new resources are added, controls whether metric collection are automatically enabled for the new resources. * `AutoEnable` - Automatically enable metric collection for new resources, up to the limit of resource collection. * `Disabled` - Metrics will not be enabled on new resources, to enable collection requires an explicit user enable.
		CollectNewDevices *string `json:"CollectNewDevices,omitempty"`
		// Enables metric collection for the account, if disabled metrics will be stopped for all resources in the account.
		Enabled *bool `json:"Enabled,omitempty"`
		// The total number of resources that can be enabled for metric collection in this account.
		Limit   *int64                         `json:"Limit,omitempty"`
		Account NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varMetricsConfigurationWithoutEmbeddedStruct := MetricsConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMetricsConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varMetricsConfiguration := _MetricsConfiguration{}
		varMetricsConfiguration.ClassId = varMetricsConfigurationWithoutEmbeddedStruct.ClassId
		varMetricsConfiguration.ObjectType = varMetricsConfigurationWithoutEmbeddedStruct.ObjectType
		varMetricsConfiguration.CollectNewDevices = varMetricsConfigurationWithoutEmbeddedStruct.CollectNewDevices
		varMetricsConfiguration.Enabled = varMetricsConfigurationWithoutEmbeddedStruct.Enabled
		varMetricsConfiguration.Limit = varMetricsConfigurationWithoutEmbeddedStruct.Limit
		varMetricsConfiguration.Account = varMetricsConfigurationWithoutEmbeddedStruct.Account
		*o = MetricsConfiguration(varMetricsConfiguration)
	} else {
		return err
	}

	varMetricsConfiguration := _MetricsConfiguration{}

	err = json.Unmarshal(data, &varMetricsConfiguration)
	if err == nil {
		o.MoBaseMo = varMetricsConfiguration.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CollectNewDevices")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Limit")
		delete(additionalProperties, "Account")

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

type NullableMetricsConfiguration struct {
	value *MetricsConfiguration
	isSet bool
}

func (v NullableMetricsConfiguration) Get() *MetricsConfiguration {
	return v.value
}

func (v *NullableMetricsConfiguration) Set(val *MetricsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsConfiguration(val *MetricsConfiguration) *NullableMetricsConfiguration {
	return &NullableMetricsConfiguration{value: val, isSet: true}
}

func (v NullableMetricsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
