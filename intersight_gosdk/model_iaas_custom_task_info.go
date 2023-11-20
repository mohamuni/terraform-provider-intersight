/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IaasCustomTaskInfo List out the execution of the Custom Tasks with Names.
type IaasCustomTaskInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Task decription or Comment of the Custom task.
	TaskDescription *string `json:"TaskDescription,omitempty"`
	// Number of times this task has executed.
	TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
	// Task Label in the Workflow.
	TaskLabel *string `json:"TaskLabel,omitempty"`
	// Name of the Custom Task in UCSD.
	TaskName             *string                   `json:"TaskName,omitempty"`
	Guid                 *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasCustomTaskInfo IaasCustomTaskInfo

// NewIaasCustomTaskInfo instantiates a new IaasCustomTaskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasCustomTaskInfo(classId string, objectType string) *IaasCustomTaskInfo {
	this := IaasCustomTaskInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasCustomTaskInfoWithDefaults instantiates a new IaasCustomTaskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasCustomTaskInfoWithDefaults() *IaasCustomTaskInfo {
	this := IaasCustomTaskInfo{}
	var classId string = "iaas.CustomTaskInfo"
	this.ClassId = classId
	var objectType string = "iaas.CustomTaskInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasCustomTaskInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasCustomTaskInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasCustomTaskInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasCustomTaskInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTaskDescription returns the TaskDescription field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskDescription() string {
	if o == nil || o.TaskDescription == nil {
		var ret string
		return ret
	}
	return *o.TaskDescription
}

// GetTaskDescriptionOk returns a tuple with the TaskDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskDescriptionOk() (*string, bool) {
	if o == nil || o.TaskDescription == nil {
		return nil, false
	}
	return o.TaskDescription, true
}

// HasTaskDescription returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskDescription() bool {
	if o != nil && o.TaskDescription != nil {
		return true
	}

	return false
}

// SetTaskDescription gets a reference to the given string and assigns it to the TaskDescription field.
func (o *IaasCustomTaskInfo) SetTaskDescription(v string) {
	o.TaskDescription = &v
}

// GetTaskExecutionCount returns the TaskExecutionCount field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskExecutionCount() int64 {
	if o == nil || o.TaskExecutionCount == nil {
		var ret int64
		return ret
	}
	return *o.TaskExecutionCount
}

// GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskExecutionCountOk() (*int64, bool) {
	if o == nil || o.TaskExecutionCount == nil {
		return nil, false
	}
	return o.TaskExecutionCount, true
}

// HasTaskExecutionCount returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskExecutionCount() bool {
	if o != nil && o.TaskExecutionCount != nil {
		return true
	}

	return false
}

// SetTaskExecutionCount gets a reference to the given int64 and assigns it to the TaskExecutionCount field.
func (o *IaasCustomTaskInfo) SetTaskExecutionCount(v int64) {
	o.TaskExecutionCount = &v
}

// GetTaskLabel returns the TaskLabel field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskLabel() string {
	if o == nil || o.TaskLabel == nil {
		var ret string
		return ret
	}
	return *o.TaskLabel
}

// GetTaskLabelOk returns a tuple with the TaskLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskLabelOk() (*string, bool) {
	if o == nil || o.TaskLabel == nil {
		return nil, false
	}
	return o.TaskLabel, true
}

// HasTaskLabel returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskLabel() bool {
	if o != nil && o.TaskLabel != nil {
		return true
	}

	return false
}

// SetTaskLabel gets a reference to the given string and assigns it to the TaskLabel field.
func (o *IaasCustomTaskInfo) SetTaskLabel(v string) {
	o.TaskLabel = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskName() string {
	if o == nil || o.TaskName == nil {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskNameOk() (*string, bool) {
	if o == nil || o.TaskName == nil {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskName() bool {
	if o != nil && o.TaskName != nil {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *IaasCustomTaskInfo) SetTaskName(v string) {
	o.TaskName = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasCustomTaskInfo) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasCustomTaskInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TaskDescription != nil {
		toSerialize["TaskDescription"] = o.TaskDescription
	}
	if o.TaskExecutionCount != nil {
		toSerialize["TaskExecutionCount"] = o.TaskExecutionCount
	}
	if o.TaskLabel != nil {
		toSerialize["TaskLabel"] = o.TaskLabel
	}
	if o.TaskName != nil {
		toSerialize["TaskName"] = o.TaskName
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasCustomTaskInfo) UnmarshalJSON(bytes []byte) (err error) {
	type IaasCustomTaskInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Task decription or Comment of the Custom task.
		TaskDescription *string `json:"TaskDescription,omitempty"`
		// Number of times this task has executed.
		TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
		// Task Label in the Workflow.
		TaskLabel *string `json:"TaskLabel,omitempty"`
		// Name of the Custom Task in UCSD.
		TaskName *string                   `json:"TaskName,omitempty"`
		Guid     *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	}

	varIaasCustomTaskInfoWithoutEmbeddedStruct := IaasCustomTaskInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIaasCustomTaskInfoWithoutEmbeddedStruct)
	if err == nil {
		varIaasCustomTaskInfo := _IaasCustomTaskInfo{}
		varIaasCustomTaskInfo.ClassId = varIaasCustomTaskInfoWithoutEmbeddedStruct.ClassId
		varIaasCustomTaskInfo.ObjectType = varIaasCustomTaskInfoWithoutEmbeddedStruct.ObjectType
		varIaasCustomTaskInfo.TaskDescription = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskDescription
		varIaasCustomTaskInfo.TaskExecutionCount = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskExecutionCount
		varIaasCustomTaskInfo.TaskLabel = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskLabel
		varIaasCustomTaskInfo.TaskName = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskName
		varIaasCustomTaskInfo.Guid = varIaasCustomTaskInfoWithoutEmbeddedStruct.Guid
		*o = IaasCustomTaskInfo(varIaasCustomTaskInfo)
	} else {
		return err
	}

	varIaasCustomTaskInfo := _IaasCustomTaskInfo{}

	err = json.Unmarshal(bytes, &varIaasCustomTaskInfo)
	if err == nil {
		o.MoBaseMo = varIaasCustomTaskInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TaskDescription")
		delete(additionalProperties, "TaskExecutionCount")
		delete(additionalProperties, "TaskLabel")
		delete(additionalProperties, "TaskName")
		delete(additionalProperties, "Guid")

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

type NullableIaasCustomTaskInfo struct {
	value *IaasCustomTaskInfo
	isSet bool
}

func (v NullableIaasCustomTaskInfo) Get() *IaasCustomTaskInfo {
	return v.value
}

func (v *NullableIaasCustomTaskInfo) Set(val *IaasCustomTaskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasCustomTaskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasCustomTaskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasCustomTaskInfo(val *IaasCustomTaskInfo) *NullableIaasCustomTaskInfo {
	return &NullableIaasCustomTaskInfo{value: val, isSet: true}
}

func (v NullableIaasCustomTaskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasCustomTaskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
