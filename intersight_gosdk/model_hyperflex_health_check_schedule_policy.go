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
	"time"
)

// checks if the HyperflexHealthCheckSchedulePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexHealthCheckSchedulePolicy{}

// HyperflexHealthCheckSchedulePolicy Continuous health check schedule policy of a HyperFlex cluster.
type HyperflexHealthCheckSchedulePolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The date and time when this HealthCheck Policy was last enabled.
	LastScheduledOn *time.Time `json:"LastScheduledOn,omitempty"`
	// The date and time when this HealthCheck Policy was last disabled.
	LastUnscheduledOn *time.Time `json:"LastUnscheduledOn,omitempty"`
	// The date and time when the next health check execution is expected.
	NextExpectedExecution *time.Time `json:"NextExpectedExecution,omitempty"`
	// Indicates whether HealthCheck schedule policy is enabled on the HyperFlex cluster.
	PolicyEnabled *bool `json:"PolicyEnabled,omitempty"`
	// The frequency at which the health checks are run on the HyperFlex cluster. * `86400` - Execute the health check every 24 hours. * `43200` - Execute the health check every 12 hours. * `21600` - Execute the health check every 6 hours. * `10800` - Execute the health check every 3 hours. * `3600` - Execute the health check every 1 hours. * `300` - Execute the health check every 5 minutes. * `0` - Disable the continuous health check.
	ScheduleInterval *int32 `json:"ScheduleInterval,omitempty"`
	// The unique identifier of the health check policy.
	Uuid                 *string                                     `json:"Uuid,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHealthCheckSchedulePolicy HyperflexHealthCheckSchedulePolicy

// NewHyperflexHealthCheckSchedulePolicy instantiates a new HyperflexHealthCheckSchedulePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckSchedulePolicy(classId string, objectType string) *HyperflexHealthCheckSchedulePolicy {
	this := HyperflexHealthCheckSchedulePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var scheduleInterval int32 = 86400
	this.ScheduleInterval = &scheduleInterval
	return &this
}

// NewHyperflexHealthCheckSchedulePolicyWithDefaults instantiates a new HyperflexHealthCheckSchedulePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckSchedulePolicyWithDefaults() *HyperflexHealthCheckSchedulePolicy {
	this := HyperflexHealthCheckSchedulePolicy{}
	var classId string = "hyperflex.HealthCheckSchedulePolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckSchedulePolicy"
	this.ObjectType = objectType
	var scheduleInterval int32 = 86400
	this.ScheduleInterval = &scheduleInterval
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckSchedulePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckSchedulePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.HealthCheckSchedulePolicy" of the ClassId field.
func (o *HyperflexHealthCheckSchedulePolicy) GetDefaultClassId() interface{} {
	return "hyperflex.HealthCheckSchedulePolicy"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckSchedulePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckSchedulePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.HealthCheckSchedulePolicy" of the ObjectType field.
func (o *HyperflexHealthCheckSchedulePolicy) GetDefaultObjectType() interface{} {
	return "hyperflex.HealthCheckSchedulePolicy"
}

// GetLastScheduledOn returns the LastScheduledOn field value if set, zero value otherwise.
func (o *HyperflexHealthCheckSchedulePolicy) GetLastScheduledOn() time.Time {
	if o == nil || IsNil(o.LastScheduledOn) {
		var ret time.Time
		return ret
	}
	return *o.LastScheduledOn
}

// GetLastScheduledOnOk returns a tuple with the LastScheduledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetLastScheduledOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastScheduledOn) {
		return nil, false
	}
	return o.LastScheduledOn, true
}

// HasLastScheduledOn returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasLastScheduledOn() bool {
	if o != nil && !IsNil(o.LastScheduledOn) {
		return true
	}

	return false
}

// SetLastScheduledOn gets a reference to the given time.Time and assigns it to the LastScheduledOn field.
func (o *HyperflexHealthCheckSchedulePolicy) SetLastScheduledOn(v time.Time) {
	o.LastScheduledOn = &v
}

// GetLastUnscheduledOn returns the LastUnscheduledOn field value if set, zero value otherwise.
func (o *HyperflexHealthCheckSchedulePolicy) GetLastUnscheduledOn() time.Time {
	if o == nil || IsNil(o.LastUnscheduledOn) {
		var ret time.Time
		return ret
	}
	return *o.LastUnscheduledOn
}

// GetLastUnscheduledOnOk returns a tuple with the LastUnscheduledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetLastUnscheduledOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUnscheduledOn) {
		return nil, false
	}
	return o.LastUnscheduledOn, true
}

// HasLastUnscheduledOn returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasLastUnscheduledOn() bool {
	if o != nil && !IsNil(o.LastUnscheduledOn) {
		return true
	}

	return false
}

// SetLastUnscheduledOn gets a reference to the given time.Time and assigns it to the LastUnscheduledOn field.
func (o *HyperflexHealthCheckSchedulePolicy) SetLastUnscheduledOn(v time.Time) {
	o.LastUnscheduledOn = &v
}

// GetNextExpectedExecution returns the NextExpectedExecution field value if set, zero value otherwise.
func (o *HyperflexHealthCheckSchedulePolicy) GetNextExpectedExecution() time.Time {
	if o == nil || IsNil(o.NextExpectedExecution) {
		var ret time.Time
		return ret
	}
	return *o.NextExpectedExecution
}

// GetNextExpectedExecutionOk returns a tuple with the NextExpectedExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetNextExpectedExecutionOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextExpectedExecution) {
		return nil, false
	}
	return o.NextExpectedExecution, true
}

// HasNextExpectedExecution returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasNextExpectedExecution() bool {
	if o != nil && !IsNil(o.NextExpectedExecution) {
		return true
	}

	return false
}

// SetNextExpectedExecution gets a reference to the given time.Time and assigns it to the NextExpectedExecution field.
func (o *HyperflexHealthCheckSchedulePolicy) SetNextExpectedExecution(v time.Time) {
	o.NextExpectedExecution = &v
}

// GetPolicyEnabled returns the PolicyEnabled field value if set, zero value otherwise.
func (o *HyperflexHealthCheckSchedulePolicy) GetPolicyEnabled() bool {
	if o == nil || IsNil(o.PolicyEnabled) {
		var ret bool
		return ret
	}
	return *o.PolicyEnabled
}

// GetPolicyEnabledOk returns a tuple with the PolicyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetPolicyEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PolicyEnabled) {
		return nil, false
	}
	return o.PolicyEnabled, true
}

// HasPolicyEnabled returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasPolicyEnabled() bool {
	if o != nil && !IsNil(o.PolicyEnabled) {
		return true
	}

	return false
}

// SetPolicyEnabled gets a reference to the given bool and assigns it to the PolicyEnabled field.
func (o *HyperflexHealthCheckSchedulePolicy) SetPolicyEnabled(v bool) {
	o.PolicyEnabled = &v
}

// GetScheduleInterval returns the ScheduleInterval field value if set, zero value otherwise.
func (o *HyperflexHealthCheckSchedulePolicy) GetScheduleInterval() int32 {
	if o == nil || IsNil(o.ScheduleInterval) {
		var ret int32
		return ret
	}
	return *o.ScheduleInterval
}

// GetScheduleIntervalOk returns a tuple with the ScheduleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetScheduleIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.ScheduleInterval) {
		return nil, false
	}
	return o.ScheduleInterval, true
}

// HasScheduleInterval returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasScheduleInterval() bool {
	if o != nil && !IsNil(o.ScheduleInterval) {
		return true
	}

	return false
}

// SetScheduleInterval gets a reference to the given int32 and assigns it to the ScheduleInterval field.
func (o *HyperflexHealthCheckSchedulePolicy) SetScheduleInterval(v int32) {
	o.ScheduleInterval = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHealthCheckSchedulePolicy) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckSchedulePolicy) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHealthCheckSchedulePolicy) SetUuid(v string) {
	o.Uuid = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealthCheckSchedulePolicy) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealthCheckSchedulePolicy) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexHealthCheckSchedulePolicy) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexHealthCheckSchedulePolicy) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *HyperflexHealthCheckSchedulePolicy) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *HyperflexHealthCheckSchedulePolicy) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o HyperflexHealthCheckSchedulePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexHealthCheckSchedulePolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LastScheduledOn) {
		toSerialize["LastScheduledOn"] = o.LastScheduledOn
	}
	if !IsNil(o.LastUnscheduledOn) {
		toSerialize["LastUnscheduledOn"] = o.LastUnscheduledOn
	}
	if !IsNil(o.NextExpectedExecution) {
		toSerialize["NextExpectedExecution"] = o.NextExpectedExecution
	}
	if !IsNil(o.PolicyEnabled) {
		toSerialize["PolicyEnabled"] = o.PolicyEnabled
	}
	if !IsNil(o.ScheduleInterval) {
		toSerialize["ScheduleInterval"] = o.ScheduleInterval
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexHealthCheckSchedulePolicy) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The date and time when this HealthCheck Policy was last enabled.
		LastScheduledOn *time.Time `json:"LastScheduledOn,omitempty"`
		// The date and time when this HealthCheck Policy was last disabled.
		LastUnscheduledOn *time.Time `json:"LastUnscheduledOn,omitempty"`
		// The date and time when the next health check execution is expected.
		NextExpectedExecution *time.Time `json:"NextExpectedExecution,omitempty"`
		// Indicates whether HealthCheck schedule policy is enabled on the HyperFlex cluster.
		PolicyEnabled *bool `json:"PolicyEnabled,omitempty"`
		// The frequency at which the health checks are run on the HyperFlex cluster. * `86400` - Execute the health check every 24 hours. * `43200` - Execute the health check every 12 hours. * `21600` - Execute the health check every 6 hours. * `10800` - Execute the health check every 3 hours. * `3600` - Execute the health check every 1 hours. * `300` - Execute the health check every 5 minutes. * `0` - Disable the continuous health check.
		ScheduleInterval *int32 `json:"ScheduleInterval,omitempty"`
		// The unique identifier of the health check policy.
		Uuid             *string                                     `json:"Uuid,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct := HyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHealthCheckSchedulePolicy := _HyperflexHealthCheckSchedulePolicy{}
		varHyperflexHealthCheckSchedulePolicy.ClassId = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.ClassId
		varHyperflexHealthCheckSchedulePolicy.ObjectType = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexHealthCheckSchedulePolicy.LastScheduledOn = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.LastScheduledOn
		varHyperflexHealthCheckSchedulePolicy.LastUnscheduledOn = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.LastUnscheduledOn
		varHyperflexHealthCheckSchedulePolicy.NextExpectedExecution = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.NextExpectedExecution
		varHyperflexHealthCheckSchedulePolicy.PolicyEnabled = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.PolicyEnabled
		varHyperflexHealthCheckSchedulePolicy.ScheduleInterval = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.ScheduleInterval
		varHyperflexHealthCheckSchedulePolicy.Uuid = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.Uuid
		varHyperflexHealthCheckSchedulePolicy.RegisteredDevice = varHyperflexHealthCheckSchedulePolicyWithoutEmbeddedStruct.RegisteredDevice
		*o = HyperflexHealthCheckSchedulePolicy(varHyperflexHealthCheckSchedulePolicy)
	} else {
		return err
	}

	varHyperflexHealthCheckSchedulePolicy := _HyperflexHealthCheckSchedulePolicy{}

	err = json.Unmarshal(data, &varHyperflexHealthCheckSchedulePolicy)
	if err == nil {
		o.MoBaseMo = varHyperflexHealthCheckSchedulePolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LastScheduledOn")
		delete(additionalProperties, "LastUnscheduledOn")
		delete(additionalProperties, "NextExpectedExecution")
		delete(additionalProperties, "PolicyEnabled")
		delete(additionalProperties, "ScheduleInterval")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableHyperflexHealthCheckSchedulePolicy struct {
	value *HyperflexHealthCheckSchedulePolicy
	isSet bool
}

func (v NullableHyperflexHealthCheckSchedulePolicy) Get() *HyperflexHealthCheckSchedulePolicy {
	return v.value
}

func (v *NullableHyperflexHealthCheckSchedulePolicy) Set(val *HyperflexHealthCheckSchedulePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckSchedulePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckSchedulePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckSchedulePolicy(val *HyperflexHealthCheckSchedulePolicy) *NullableHyperflexHealthCheckSchedulePolicy {
	return &NullableHyperflexHealthCheckSchedulePolicy{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckSchedulePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckSchedulePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
