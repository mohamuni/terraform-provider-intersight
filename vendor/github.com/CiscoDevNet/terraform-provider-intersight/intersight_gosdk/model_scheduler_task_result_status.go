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

// checks if the SchedulerTaskResultStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerTaskResultStatus{}

// SchedulerTaskResultStatus The status of a single instance of the scheduled task.
type SchedulerTaskResultStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The status reason, used when a scheduled task fails or was suspended by the system.
	Reason *string `json:"Reason,omitempty"`
	// The status of the current scheduled task. * `None` - No status is set (default). * `Scheduled` - The status is set when a task is scheduled. * `Running` - The status is set when a task is running. * `Completed` - The status is set when a task is complete. * `Failed` - The status is set when a task fails. * `Suspended` - The status is set when a task is suspended. * `Skipped` - The status is set when a task is skipped because the previous task is still running.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerTaskResultStatus SchedulerTaskResultStatus

// NewSchedulerTaskResultStatus instantiates a new SchedulerTaskResultStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerTaskResultStatus(classId string, objectType string) *SchedulerTaskResultStatus {
	this := SchedulerTaskResultStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSchedulerTaskResultStatusWithDefaults instantiates a new SchedulerTaskResultStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerTaskResultStatusWithDefaults() *SchedulerTaskResultStatus {
	this := SchedulerTaskResultStatus{}
	var classId string = "scheduler.TaskResultStatus"
	this.ClassId = classId
	var objectType string = "scheduler.TaskResultStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerTaskResultStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResultStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerTaskResultStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "scheduler.TaskResultStatus" of the ClassId field.
func (o *SchedulerTaskResultStatus) GetDefaultClassId() interface{} {
	return "scheduler.TaskResultStatus"
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerTaskResultStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResultStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerTaskResultStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "scheduler.TaskResultStatus" of the ObjectType field.
func (o *SchedulerTaskResultStatus) GetDefaultObjectType() interface{} {
	return "scheduler.TaskResultStatus"
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SchedulerTaskResultStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResultStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SchedulerTaskResultStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SchedulerTaskResultStatus) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SchedulerTaskResultStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResultStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SchedulerTaskResultStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SchedulerTaskResultStatus) SetStatus(v string) {
	o.Status = &v
}

func (o SchedulerTaskResultStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerTaskResultStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerTaskResultStatus) UnmarshalJSON(data []byte) (err error) {
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
	type SchedulerTaskResultStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The status reason, used when a scheduled task fails or was suspended by the system.
		Reason *string `json:"Reason,omitempty"`
		// The status of the current scheduled task. * `None` - No status is set (default). * `Scheduled` - The status is set when a task is scheduled. * `Running` - The status is set when a task is running. * `Completed` - The status is set when a task is complete. * `Failed` - The status is set when a task fails. * `Suspended` - The status is set when a task is suspended. * `Skipped` - The status is set when a task is skipped because the previous task is still running.
		Status *string `json:"Status,omitempty"`
	}

	varSchedulerTaskResultStatusWithoutEmbeddedStruct := SchedulerTaskResultStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerTaskResultStatusWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerTaskResultStatus := _SchedulerTaskResultStatus{}
		varSchedulerTaskResultStatus.ClassId = varSchedulerTaskResultStatusWithoutEmbeddedStruct.ClassId
		varSchedulerTaskResultStatus.ObjectType = varSchedulerTaskResultStatusWithoutEmbeddedStruct.ObjectType
		varSchedulerTaskResultStatus.Reason = varSchedulerTaskResultStatusWithoutEmbeddedStruct.Reason
		varSchedulerTaskResultStatus.Status = varSchedulerTaskResultStatusWithoutEmbeddedStruct.Status
		*o = SchedulerTaskResultStatus(varSchedulerTaskResultStatus)
	} else {
		return err
	}

	varSchedulerTaskResultStatus := _SchedulerTaskResultStatus{}

	err = json.Unmarshal(data, &varSchedulerTaskResultStatus)
	if err == nil {
		o.MoBaseComplexType = varSchedulerTaskResultStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")

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

type NullableSchedulerTaskResultStatus struct {
	value *SchedulerTaskResultStatus
	isSet bool
}

func (v NullableSchedulerTaskResultStatus) Get() *SchedulerTaskResultStatus {
	return v.value
}

func (v *NullableSchedulerTaskResultStatus) Set(val *SchedulerTaskResultStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerTaskResultStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerTaskResultStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerTaskResultStatus(val *SchedulerTaskResultStatus) *NullableSchedulerTaskResultStatus {
	return &NullableSchedulerTaskResultStatus{value: val, isSet: true}
}

func (v NullableSchedulerTaskResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerTaskResultStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
