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

// checks if the CondAlarmAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CondAlarmAggregation{}

// CondAlarmAggregation Object which represents alarm aggregation for a managed end point.
type CondAlarmAggregation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                   `json:"ObjectType"`
	AlarmSummary NullableCondAlarmSummary `json:"AlarmSummary,omitempty"`
	// Count of all alarms with severity Critical, irrespective of acknowledgement status.
	CriticalAlarmsCount *int64 `json:"CriticalAlarmsCount,omitempty"`
	// Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	// Deprecated
	Health *string `json:"Health,omitempty"`
	// Count of all alarms with severity Info, irrespective of acknowledgement status.
	InfoAlarmsCount *int64 `json:"InfoAlarmsCount,omitempty"`
	// Managed object type. For example, FI managed object type will be network.Element.
	MoType *string `json:"MoType,omitempty"`
	// Count of all alarms with severity Warning, irrespective of acknowledgement status.
	WarningAlarmsCount     *int64                       `json:"WarningAlarmsCount,omitempty"`
	AlarmAggregationSource NullableMoBaseMoRelationship `json:"AlarmAggregationSource,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _CondAlarmAggregation CondAlarmAggregation

// NewCondAlarmAggregation instantiates a new CondAlarmAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmAggregation(classId string, objectType string) *CondAlarmAggregation {
	this := CondAlarmAggregation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCondAlarmAggregationWithDefaults instantiates a new CondAlarmAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmAggregationWithDefaults() *CondAlarmAggregation {
	this := CondAlarmAggregation{}
	var classId string = "cond.AlarmAggregation"
	this.ClassId = classId
	var objectType string = "cond.AlarmAggregation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmAggregation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmAggregation) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cond.AlarmAggregation" of the ClassId field.
func (o *CondAlarmAggregation) GetDefaultClassId() interface{} {
	return "cond.AlarmAggregation"
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmAggregation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmAggregation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cond.AlarmAggregation" of the ObjectType field.
func (o *CondAlarmAggregation) GetDefaultObjectType() interface{} {
	return "cond.AlarmAggregation"
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmAggregation) GetAlarmSummary() CondAlarmSummary {
	if o == nil || IsNil(o.AlarmSummary.Get()) {
		var ret CondAlarmSummary
		return ret
	}
	return *o.AlarmSummary.Get()
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmAggregation) GetAlarmSummaryOk() (*CondAlarmSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmSummary.Get(), o.AlarmSummary.IsSet()
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary.IsSet() {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given NullableCondAlarmSummary and assigns it to the AlarmSummary field.
func (o *CondAlarmAggregation) SetAlarmSummary(v CondAlarmSummary) {
	o.AlarmSummary.Set(&v)
}

// SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil
func (o *CondAlarmAggregation) SetAlarmSummaryNil() {
	o.AlarmSummary.Set(nil)
}

// UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
func (o *CondAlarmAggregation) UnsetAlarmSummary() {
	o.AlarmSummary.Unset()
}

// GetCriticalAlarmsCount returns the CriticalAlarmsCount field value if set, zero value otherwise.
func (o *CondAlarmAggregation) GetCriticalAlarmsCount() int64 {
	if o == nil || IsNil(o.CriticalAlarmsCount) {
		var ret int64
		return ret
	}
	return *o.CriticalAlarmsCount
}

// GetCriticalAlarmsCountOk returns a tuple with the CriticalAlarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregation) GetCriticalAlarmsCountOk() (*int64, bool) {
	if o == nil || IsNil(o.CriticalAlarmsCount) {
		return nil, false
	}
	return o.CriticalAlarmsCount, true
}

// HasCriticalAlarmsCount returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasCriticalAlarmsCount() bool {
	if o != nil && !IsNil(o.CriticalAlarmsCount) {
		return true
	}

	return false
}

// SetCriticalAlarmsCount gets a reference to the given int64 and assigns it to the CriticalAlarmsCount field.
func (o *CondAlarmAggregation) SetCriticalAlarmsCount(v int64) {
	o.CriticalAlarmsCount = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
// Deprecated
func (o *CondAlarmAggregation) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CondAlarmAggregation) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
// Deprecated
func (o *CondAlarmAggregation) SetHealth(v string) {
	o.Health = &v
}

// GetInfoAlarmsCount returns the InfoAlarmsCount field value if set, zero value otherwise.
func (o *CondAlarmAggregation) GetInfoAlarmsCount() int64 {
	if o == nil || IsNil(o.InfoAlarmsCount) {
		var ret int64
		return ret
	}
	return *o.InfoAlarmsCount
}

// GetInfoAlarmsCountOk returns a tuple with the InfoAlarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregation) GetInfoAlarmsCountOk() (*int64, bool) {
	if o == nil || IsNil(o.InfoAlarmsCount) {
		return nil, false
	}
	return o.InfoAlarmsCount, true
}

// HasInfoAlarmsCount returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasInfoAlarmsCount() bool {
	if o != nil && !IsNil(o.InfoAlarmsCount) {
		return true
	}

	return false
}

// SetInfoAlarmsCount gets a reference to the given int64 and assigns it to the InfoAlarmsCount field.
func (o *CondAlarmAggregation) SetInfoAlarmsCount(v int64) {
	o.InfoAlarmsCount = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *CondAlarmAggregation) GetMoType() string {
	if o == nil || IsNil(o.MoType) {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregation) GetMoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MoType) {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasMoType() bool {
	if o != nil && !IsNil(o.MoType) {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *CondAlarmAggregation) SetMoType(v string) {
	o.MoType = &v
}

// GetWarningAlarmsCount returns the WarningAlarmsCount field value if set, zero value otherwise.
func (o *CondAlarmAggregation) GetWarningAlarmsCount() int64 {
	if o == nil || IsNil(o.WarningAlarmsCount) {
		var ret int64
		return ret
	}
	return *o.WarningAlarmsCount
}

// GetWarningAlarmsCountOk returns a tuple with the WarningAlarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregation) GetWarningAlarmsCountOk() (*int64, bool) {
	if o == nil || IsNil(o.WarningAlarmsCount) {
		return nil, false
	}
	return o.WarningAlarmsCount, true
}

// HasWarningAlarmsCount returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasWarningAlarmsCount() bool {
	if o != nil && !IsNil(o.WarningAlarmsCount) {
		return true
	}

	return false
}

// SetWarningAlarmsCount gets a reference to the given int64 and assigns it to the WarningAlarmsCount field.
func (o *CondAlarmAggregation) SetWarningAlarmsCount(v int64) {
	o.WarningAlarmsCount = &v
}

// GetAlarmAggregationSource returns the AlarmAggregationSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmAggregation) GetAlarmAggregationSource() MoBaseMoRelationship {
	if o == nil || IsNil(o.AlarmAggregationSource.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AlarmAggregationSource.Get()
}

// GetAlarmAggregationSourceOk returns a tuple with the AlarmAggregationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmAggregation) GetAlarmAggregationSourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmAggregationSource.Get(), o.AlarmAggregationSource.IsSet()
}

// HasAlarmAggregationSource returns a boolean if a field has been set.
func (o *CondAlarmAggregation) HasAlarmAggregationSource() bool {
	if o != nil && o.AlarmAggregationSource.IsSet() {
		return true
	}

	return false
}

// SetAlarmAggregationSource gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AlarmAggregationSource field.
func (o *CondAlarmAggregation) SetAlarmAggregationSource(v MoBaseMoRelationship) {
	o.AlarmAggregationSource.Set(&v)
}

// SetAlarmAggregationSourceNil sets the value for AlarmAggregationSource to be an explicit nil
func (o *CondAlarmAggregation) SetAlarmAggregationSourceNil() {
	o.AlarmAggregationSource.Set(nil)
}

// UnsetAlarmAggregationSource ensures that no value is present for AlarmAggregationSource, not even an explicit nil
func (o *CondAlarmAggregation) UnsetAlarmAggregationSource() {
	o.AlarmAggregationSource.Unset()
}

func (o CondAlarmAggregation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CondAlarmAggregation) ToMap() (map[string]interface{}, error) {
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
	if o.AlarmSummary.IsSet() {
		toSerialize["AlarmSummary"] = o.AlarmSummary.Get()
	}
	if !IsNil(o.CriticalAlarmsCount) {
		toSerialize["CriticalAlarmsCount"] = o.CriticalAlarmsCount
	}
	if !IsNil(o.Health) {
		toSerialize["Health"] = o.Health
	}
	if !IsNil(o.InfoAlarmsCount) {
		toSerialize["InfoAlarmsCount"] = o.InfoAlarmsCount
	}
	if !IsNil(o.MoType) {
		toSerialize["MoType"] = o.MoType
	}
	if !IsNil(o.WarningAlarmsCount) {
		toSerialize["WarningAlarmsCount"] = o.WarningAlarmsCount
	}
	if o.AlarmAggregationSource.IsSet() {
		toSerialize["AlarmAggregationSource"] = o.AlarmAggregationSource.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CondAlarmAggregation) UnmarshalJSON(data []byte) (err error) {
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
	type CondAlarmAggregationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                   `json:"ObjectType"`
		AlarmSummary NullableCondAlarmSummary `json:"AlarmSummary,omitempty"`
		// Count of all alarms with severity Critical, irrespective of acknowledgement status.
		CriticalAlarmsCount *int64 `json:"CriticalAlarmsCount,omitempty"`
		// Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
		// Deprecated
		Health *string `json:"Health,omitempty"`
		// Count of all alarms with severity Info, irrespective of acknowledgement status.
		InfoAlarmsCount *int64 `json:"InfoAlarmsCount,omitempty"`
		// Managed object type. For example, FI managed object type will be network.Element.
		MoType *string `json:"MoType,omitempty"`
		// Count of all alarms with severity Warning, irrespective of acknowledgement status.
		WarningAlarmsCount     *int64                       `json:"WarningAlarmsCount,omitempty"`
		AlarmAggregationSource NullableMoBaseMoRelationship `json:"AlarmAggregationSource,omitempty"`
	}

	varCondAlarmAggregationWithoutEmbeddedStruct := CondAlarmAggregationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCondAlarmAggregationWithoutEmbeddedStruct)
	if err == nil {
		varCondAlarmAggregation := _CondAlarmAggregation{}
		varCondAlarmAggregation.ClassId = varCondAlarmAggregationWithoutEmbeddedStruct.ClassId
		varCondAlarmAggregation.ObjectType = varCondAlarmAggregationWithoutEmbeddedStruct.ObjectType
		varCondAlarmAggregation.AlarmSummary = varCondAlarmAggregationWithoutEmbeddedStruct.AlarmSummary
		varCondAlarmAggregation.CriticalAlarmsCount = varCondAlarmAggregationWithoutEmbeddedStruct.CriticalAlarmsCount
		varCondAlarmAggregation.Health = varCondAlarmAggregationWithoutEmbeddedStruct.Health
		varCondAlarmAggregation.InfoAlarmsCount = varCondAlarmAggregationWithoutEmbeddedStruct.InfoAlarmsCount
		varCondAlarmAggregation.MoType = varCondAlarmAggregationWithoutEmbeddedStruct.MoType
		varCondAlarmAggregation.WarningAlarmsCount = varCondAlarmAggregationWithoutEmbeddedStruct.WarningAlarmsCount
		varCondAlarmAggregation.AlarmAggregationSource = varCondAlarmAggregationWithoutEmbeddedStruct.AlarmAggregationSource
		*o = CondAlarmAggregation(varCondAlarmAggregation)
	} else {
		return err
	}

	varCondAlarmAggregation := _CondAlarmAggregation{}

	err = json.Unmarshal(data, &varCondAlarmAggregation)
	if err == nil {
		o.MoBaseMo = varCondAlarmAggregation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "CriticalAlarmsCount")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "InfoAlarmsCount")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "WarningAlarmsCount")
		delete(additionalProperties, "AlarmAggregationSource")

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

type NullableCondAlarmAggregation struct {
	value *CondAlarmAggregation
	isSet bool
}

func (v NullableCondAlarmAggregation) Get() *CondAlarmAggregation {
	return v.value
}

func (v *NullableCondAlarmAggregation) Set(val *CondAlarmAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmAggregation(val *CondAlarmAggregation) *NullableCondAlarmAggregation {
	return &NullableCondAlarmAggregation{value: val, isSet: true}
}

func (v NullableCondAlarmAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
