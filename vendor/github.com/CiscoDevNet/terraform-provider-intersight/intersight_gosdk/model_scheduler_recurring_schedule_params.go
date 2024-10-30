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
	"time"
)

// checks if the SchedulerRecurringScheduleParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerRecurringScheduleParams{}

// SchedulerRecurringScheduleParams The parameters for configuring a recurring schedule.
type SchedulerRecurringScheduleParams struct {
	SchedulerBaseScheduleParams
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Allowed values for a recurring schedule cadence. * `None` - No value set for the cadence type (Enum value None). * `Every` - Use the 'Every' cadence for tasks that need to be run frequently and are relatively small or quick to execute. This could include tasks such as checking the status of a service every 15 minutes, or updating a counter. * `Daily` - A Daily cadence allows for a scheduled task to be run every day or every n-interval days. * `Weekly` - A Weekly cadence allows for a scheduled task to be run every week or every n-interval weeks on specific days. * `Monthly` - A Montly cadence allows for a scheduled task to be run every month on specific days.
	Cadence *string `json:"Cadence,omitempty"`
	// Specify the number of occurrences (instead of an end-time) for a recurring schedule.
	EndAfterOccurrences *int64 `json:"EndAfterOccurrences,omitempty"`
	// End time for the recurring schedule. The schedule will not run beyond this time. If using the endAfterOccurrences parameter instead, this field should be set to zero time, i.e, 0001-01-01T00:00:00Z.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The maximum number of consecutive failures until the recurring scheduled task is suspended by the system. The default is 1.
	FailureThreshold     *int64                             `json:"FailureThreshold,omitempty"`
	Params               NullableSchedulerBaseCadenceParams `json:"Params,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerRecurringScheduleParams SchedulerRecurringScheduleParams

// NewSchedulerRecurringScheduleParams instantiates a new SchedulerRecurringScheduleParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerRecurringScheduleParams(classId string, objectType string) *SchedulerRecurringScheduleParams {
	this := SchedulerRecurringScheduleParams{}
	this.ClassId = classId
	this.ObjectType = objectType
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	var cadence string = "None"
	this.Cadence = &cadence
	var failureThreshold int64 = 1
	this.FailureThreshold = &failureThreshold
	return &this
}

// NewSchedulerRecurringScheduleParamsWithDefaults instantiates a new SchedulerRecurringScheduleParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerRecurringScheduleParamsWithDefaults() *SchedulerRecurringScheduleParams {
	this := SchedulerRecurringScheduleParams{}
	var classId string = "scheduler.RecurringScheduleParams"
	this.ClassId = classId
	var objectType string = "scheduler.RecurringScheduleParams"
	this.ObjectType = objectType
	var cadence string = "None"
	this.Cadence = &cadence
	var failureThreshold int64 = 1
	this.FailureThreshold = &failureThreshold
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerRecurringScheduleParams) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerRecurringScheduleParams) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerRecurringScheduleParams) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "scheduler.RecurringScheduleParams" of the ClassId field.
func (o *SchedulerRecurringScheduleParams) GetDefaultClassId() interface{} {
	return "scheduler.RecurringScheduleParams"
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerRecurringScheduleParams) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerRecurringScheduleParams) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerRecurringScheduleParams) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "scheduler.RecurringScheduleParams" of the ObjectType field.
func (o *SchedulerRecurringScheduleParams) GetDefaultObjectType() interface{} {
	return "scheduler.RecurringScheduleParams"
}

// GetCadence returns the Cadence field value if set, zero value otherwise.
func (o *SchedulerRecurringScheduleParams) GetCadence() string {
	if o == nil || IsNil(o.Cadence) {
		var ret string
		return ret
	}
	return *o.Cadence
}

// GetCadenceOk returns a tuple with the Cadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerRecurringScheduleParams) GetCadenceOk() (*string, bool) {
	if o == nil || IsNil(o.Cadence) {
		return nil, false
	}
	return o.Cadence, true
}

// HasCadence returns a boolean if a field has been set.
func (o *SchedulerRecurringScheduleParams) HasCadence() bool {
	if o != nil && !IsNil(o.Cadence) {
		return true
	}

	return false
}

// SetCadence gets a reference to the given string and assigns it to the Cadence field.
func (o *SchedulerRecurringScheduleParams) SetCadence(v string) {
	o.Cadence = &v
}

// GetEndAfterOccurrences returns the EndAfterOccurrences field value if set, zero value otherwise.
func (o *SchedulerRecurringScheduleParams) GetEndAfterOccurrences() int64 {
	if o == nil || IsNil(o.EndAfterOccurrences) {
		var ret int64
		return ret
	}
	return *o.EndAfterOccurrences
}

// GetEndAfterOccurrencesOk returns a tuple with the EndAfterOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerRecurringScheduleParams) GetEndAfterOccurrencesOk() (*int64, bool) {
	if o == nil || IsNil(o.EndAfterOccurrences) {
		return nil, false
	}
	return o.EndAfterOccurrences, true
}

// HasEndAfterOccurrences returns a boolean if a field has been set.
func (o *SchedulerRecurringScheduleParams) HasEndAfterOccurrences() bool {
	if o != nil && !IsNil(o.EndAfterOccurrences) {
		return true
	}

	return false
}

// SetEndAfterOccurrences gets a reference to the given int64 and assigns it to the EndAfterOccurrences field.
func (o *SchedulerRecurringScheduleParams) SetEndAfterOccurrences(v int64) {
	o.EndAfterOccurrences = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SchedulerRecurringScheduleParams) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerRecurringScheduleParams) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SchedulerRecurringScheduleParams) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *SchedulerRecurringScheduleParams) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise.
func (o *SchedulerRecurringScheduleParams) GetFailureThreshold() int64 {
	if o == nil || IsNil(o.FailureThreshold) {
		var ret int64
		return ret
	}
	return *o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerRecurringScheduleParams) GetFailureThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.FailureThreshold) {
		return nil, false
	}
	return o.FailureThreshold, true
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *SchedulerRecurringScheduleParams) HasFailureThreshold() bool {
	if o != nil && !IsNil(o.FailureThreshold) {
		return true
	}

	return false
}

// SetFailureThreshold gets a reference to the given int64 and assigns it to the FailureThreshold field.
func (o *SchedulerRecurringScheduleParams) SetFailureThreshold(v int64) {
	o.FailureThreshold = &v
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerRecurringScheduleParams) GetParams() SchedulerBaseCadenceParams {
	if o == nil || IsNil(o.Params.Get()) {
		var ret SchedulerBaseCadenceParams
		return ret
	}
	return *o.Params.Get()
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerRecurringScheduleParams) GetParamsOk() (*SchedulerBaseCadenceParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Params.Get(), o.Params.IsSet()
}

// HasParams returns a boolean if a field has been set.
func (o *SchedulerRecurringScheduleParams) HasParams() bool {
	if o != nil && o.Params.IsSet() {
		return true
	}

	return false
}

// SetParams gets a reference to the given NullableSchedulerBaseCadenceParams and assigns it to the Params field.
func (o *SchedulerRecurringScheduleParams) SetParams(v SchedulerBaseCadenceParams) {
	o.Params.Set(&v)
}

// SetParamsNil sets the value for Params to be an explicit nil
func (o *SchedulerRecurringScheduleParams) SetParamsNil() {
	o.Params.Set(nil)
}

// UnsetParams ensures that no value is present for Params, not even an explicit nil
func (o *SchedulerRecurringScheduleParams) UnsetParams() {
	o.Params.Unset()
}

func (o SchedulerRecurringScheduleParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerRecurringScheduleParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSchedulerBaseScheduleParams, errSchedulerBaseScheduleParams := json.Marshal(o.SchedulerBaseScheduleParams)
	if errSchedulerBaseScheduleParams != nil {
		return map[string]interface{}{}, errSchedulerBaseScheduleParams
	}
	errSchedulerBaseScheduleParams = json.Unmarshal([]byte(serializedSchedulerBaseScheduleParams), &toSerialize)
	if errSchedulerBaseScheduleParams != nil {
		return map[string]interface{}{}, errSchedulerBaseScheduleParams
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Cadence) {
		toSerialize["Cadence"] = o.Cadence
	}
	if !IsNil(o.EndAfterOccurrences) {
		toSerialize["EndAfterOccurrences"] = o.EndAfterOccurrences
	}
	if !IsNil(o.EndTime) {
		toSerialize["EndTime"] = o.EndTime
	}
	if !IsNil(o.FailureThreshold) {
		toSerialize["FailureThreshold"] = o.FailureThreshold
	}
	if o.Params.IsSet() {
		toSerialize["Params"] = o.Params.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerRecurringScheduleParams) UnmarshalJSON(data []byte) (err error) {
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
	type SchedulerRecurringScheduleParamsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Allowed values for a recurring schedule cadence. * `None` - No value set for the cadence type (Enum value None). * `Every` - Use the 'Every' cadence for tasks that need to be run frequently and are relatively small or quick to execute. This could include tasks such as checking the status of a service every 15 minutes, or updating a counter. * `Daily` - A Daily cadence allows for a scheduled task to be run every day or every n-interval days. * `Weekly` - A Weekly cadence allows for a scheduled task to be run every week or every n-interval weeks on specific days. * `Monthly` - A Montly cadence allows for a scheduled task to be run every month on specific days.
		Cadence *string `json:"Cadence,omitempty"`
		// Specify the number of occurrences (instead of an end-time) for a recurring schedule.
		EndAfterOccurrences *int64 `json:"EndAfterOccurrences,omitempty"`
		// End time for the recurring schedule. The schedule will not run beyond this time. If using the endAfterOccurrences parameter instead, this field should be set to zero time, i.e, 0001-01-01T00:00:00Z.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// The maximum number of consecutive failures until the recurring scheduled task is suspended by the system. The default is 1.
		FailureThreshold *int64                             `json:"FailureThreshold,omitempty"`
		Params           NullableSchedulerBaseCadenceParams `json:"Params,omitempty"`
	}

	varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct := SchedulerRecurringScheduleParamsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerRecurringScheduleParams := _SchedulerRecurringScheduleParams{}
		varSchedulerRecurringScheduleParams.ClassId = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.ClassId
		varSchedulerRecurringScheduleParams.ObjectType = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.ObjectType
		varSchedulerRecurringScheduleParams.Cadence = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.Cadence
		varSchedulerRecurringScheduleParams.EndAfterOccurrences = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.EndAfterOccurrences
		varSchedulerRecurringScheduleParams.EndTime = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.EndTime
		varSchedulerRecurringScheduleParams.FailureThreshold = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.FailureThreshold
		varSchedulerRecurringScheduleParams.Params = varSchedulerRecurringScheduleParamsWithoutEmbeddedStruct.Params
		*o = SchedulerRecurringScheduleParams(varSchedulerRecurringScheduleParams)
	} else {
		return err
	}

	varSchedulerRecurringScheduleParams := _SchedulerRecurringScheduleParams{}

	err = json.Unmarshal(data, &varSchedulerRecurringScheduleParams)
	if err == nil {
		o.SchedulerBaseScheduleParams = varSchedulerRecurringScheduleParams.SchedulerBaseScheduleParams
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cadence")
		delete(additionalProperties, "EndAfterOccurrences")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "FailureThreshold")
		delete(additionalProperties, "Params")

		// remove fields from embedded structs
		reflectSchedulerBaseScheduleParams := reflect.ValueOf(o.SchedulerBaseScheduleParams)
		for i := 0; i < reflectSchedulerBaseScheduleParams.Type().NumField(); i++ {
			t := reflectSchedulerBaseScheduleParams.Type().Field(i)

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

type NullableSchedulerRecurringScheduleParams struct {
	value *SchedulerRecurringScheduleParams
	isSet bool
}

func (v NullableSchedulerRecurringScheduleParams) Get() *SchedulerRecurringScheduleParams {
	return v.value
}

func (v *NullableSchedulerRecurringScheduleParams) Set(val *SchedulerRecurringScheduleParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerRecurringScheduleParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerRecurringScheduleParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerRecurringScheduleParams(val *SchedulerRecurringScheduleParams) *NullableSchedulerRecurringScheduleParams {
	return &NullableSchedulerRecurringScheduleParams{value: val, isSet: true}
}

func (v NullableSchedulerRecurringScheduleParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerRecurringScheduleParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
