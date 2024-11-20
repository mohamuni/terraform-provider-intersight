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

// checks if the InventoryJobInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryJobInfo{}

// InventoryJobInfo The runtime information about a job scheduled on a managed device.
type InventoryJobInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The execution status of the job scheduled. * `Unknown` - Operational job status on a managed device is not known. * `Scheduled` - Job on a managed device is scheduled. * `Completed` - Job on a managed device is completed. * `Error` - Job on a managed device is errored out.
	ExecutionStatus *string `json:"ExecutionStatus,omitempty"`
	// The name of the job scheduled.
	JobName *string `json:"JobName,omitempty"`
	// Timestamp to indicate the last processed time for this job.
	LastProcessedTime *time.Time `json:"LastProcessedTime,omitempty"`
	// Last timestamp when this inventory was scheduled.
	LastScheduledTime    *time.Time `json:"LastScheduledTime,omitempty"`
	Properties           []string   `json:"Properties,omitempty"`
	Regex                []string   `json:"Regex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryJobInfo InventoryJobInfo

// NewInventoryJobInfo instantiates a new InventoryJobInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryJobInfo(classId string, objectType string) *InventoryJobInfo {
	this := InventoryJobInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryJobInfoWithDefaults instantiates a new InventoryJobInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryJobInfoWithDefaults() *InventoryJobInfo {
	this := InventoryJobInfo{}
	var classId string = "inventory.JobInfo"
	this.ClassId = classId
	var objectType string = "inventory.JobInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryJobInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryJobInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryJobInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "inventory.JobInfo" of the ClassId field.
func (o *InventoryJobInfo) GetDefaultClassId() interface{} {
	return "inventory.JobInfo"
}

// GetObjectType returns the ObjectType field value
func (o *InventoryJobInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryJobInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryJobInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "inventory.JobInfo" of the ObjectType field.
func (o *InventoryJobInfo) GetDefaultObjectType() interface{} {
	return "inventory.JobInfo"
}

// GetExecutionStatus returns the ExecutionStatus field value if set, zero value otherwise.
func (o *InventoryJobInfo) GetExecutionStatus() string {
	if o == nil || IsNil(o.ExecutionStatus) {
		var ret string
		return ret
	}
	return *o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryJobInfo) GetExecutionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionStatus) {
		return nil, false
	}
	return o.ExecutionStatus, true
}

// HasExecutionStatus returns a boolean if a field has been set.
func (o *InventoryJobInfo) HasExecutionStatus() bool {
	if o != nil && !IsNil(o.ExecutionStatus) {
		return true
	}

	return false
}

// SetExecutionStatus gets a reference to the given string and assigns it to the ExecutionStatus field.
func (o *InventoryJobInfo) SetExecutionStatus(v string) {
	o.ExecutionStatus = &v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *InventoryJobInfo) GetJobName() string {
	if o == nil || IsNil(o.JobName) {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryJobInfo) GetJobNameOk() (*string, bool) {
	if o == nil || IsNil(o.JobName) {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *InventoryJobInfo) HasJobName() bool {
	if o != nil && !IsNil(o.JobName) {
		return true
	}

	return false
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *InventoryJobInfo) SetJobName(v string) {
	o.JobName = &v
}

// GetLastProcessedTime returns the LastProcessedTime field value if set, zero value otherwise.
func (o *InventoryJobInfo) GetLastProcessedTime() time.Time {
	if o == nil || IsNil(o.LastProcessedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastProcessedTime
}

// GetLastProcessedTimeOk returns a tuple with the LastProcessedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryJobInfo) GetLastProcessedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastProcessedTime) {
		return nil, false
	}
	return o.LastProcessedTime, true
}

// HasLastProcessedTime returns a boolean if a field has been set.
func (o *InventoryJobInfo) HasLastProcessedTime() bool {
	if o != nil && !IsNil(o.LastProcessedTime) {
		return true
	}

	return false
}

// SetLastProcessedTime gets a reference to the given time.Time and assigns it to the LastProcessedTime field.
func (o *InventoryJobInfo) SetLastProcessedTime(v time.Time) {
	o.LastProcessedTime = &v
}

// GetLastScheduledTime returns the LastScheduledTime field value if set, zero value otherwise.
func (o *InventoryJobInfo) GetLastScheduledTime() time.Time {
	if o == nil || IsNil(o.LastScheduledTime) {
		var ret time.Time
		return ret
	}
	return *o.LastScheduledTime
}

// GetLastScheduledTimeOk returns a tuple with the LastScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryJobInfo) GetLastScheduledTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastScheduledTime) {
		return nil, false
	}
	return o.LastScheduledTime, true
}

// HasLastScheduledTime returns a boolean if a field has been set.
func (o *InventoryJobInfo) HasLastScheduledTime() bool {
	if o != nil && !IsNil(o.LastScheduledTime) {
		return true
	}

	return false
}

// SetLastScheduledTime gets a reference to the given time.Time and assigns it to the LastScheduledTime field.
func (o *InventoryJobInfo) SetLastScheduledTime(v time.Time) {
	o.LastScheduledTime = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryJobInfo) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryJobInfo) GetPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *InventoryJobInfo) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []string and assigns it to the Properties field.
func (o *InventoryJobInfo) SetProperties(v []string) {
	o.Properties = v
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryJobInfo) GetRegex() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryJobInfo) GetRegexOk() ([]string, bool) {
	if o == nil || IsNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *InventoryJobInfo) HasRegex() bool {
	if o != nil && !IsNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given []string and assigns it to the Regex field.
func (o *InventoryJobInfo) SetRegex(v []string) {
	o.Regex = v
}

func (o InventoryJobInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryJobInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExecutionStatus) {
		toSerialize["ExecutionStatus"] = o.ExecutionStatus
	}
	if !IsNil(o.JobName) {
		toSerialize["JobName"] = o.JobName
	}
	if !IsNil(o.LastProcessedTime) {
		toSerialize["LastProcessedTime"] = o.LastProcessedTime
	}
	if !IsNil(o.LastScheduledTime) {
		toSerialize["LastScheduledTime"] = o.LastScheduledTime
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	if o.Regex != nil {
		toSerialize["Regex"] = o.Regex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryJobInfo) UnmarshalJSON(data []byte) (err error) {
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
	type InventoryJobInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The execution status of the job scheduled. * `Unknown` - Operational job status on a managed device is not known. * `Scheduled` - Job on a managed device is scheduled. * `Completed` - Job on a managed device is completed. * `Error` - Job on a managed device is errored out.
		ExecutionStatus *string `json:"ExecutionStatus,omitempty"`
		// The name of the job scheduled.
		JobName *string `json:"JobName,omitempty"`
		// Timestamp to indicate the last processed time for this job.
		LastProcessedTime *time.Time `json:"LastProcessedTime,omitempty"`
		// Last timestamp when this inventory was scheduled.
		LastScheduledTime *time.Time `json:"LastScheduledTime,omitempty"`
		Properties        []string   `json:"Properties,omitempty"`
		Regex             []string   `json:"Regex,omitempty"`
	}

	varInventoryJobInfoWithoutEmbeddedStruct := InventoryJobInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varInventoryJobInfoWithoutEmbeddedStruct)
	if err == nil {
		varInventoryJobInfo := _InventoryJobInfo{}
		varInventoryJobInfo.ClassId = varInventoryJobInfoWithoutEmbeddedStruct.ClassId
		varInventoryJobInfo.ObjectType = varInventoryJobInfoWithoutEmbeddedStruct.ObjectType
		varInventoryJobInfo.ExecutionStatus = varInventoryJobInfoWithoutEmbeddedStruct.ExecutionStatus
		varInventoryJobInfo.JobName = varInventoryJobInfoWithoutEmbeddedStruct.JobName
		varInventoryJobInfo.LastProcessedTime = varInventoryJobInfoWithoutEmbeddedStruct.LastProcessedTime
		varInventoryJobInfo.LastScheduledTime = varInventoryJobInfoWithoutEmbeddedStruct.LastScheduledTime
		varInventoryJobInfo.Properties = varInventoryJobInfoWithoutEmbeddedStruct.Properties
		varInventoryJobInfo.Regex = varInventoryJobInfoWithoutEmbeddedStruct.Regex
		*o = InventoryJobInfo(varInventoryJobInfo)
	} else {
		return err
	}

	varInventoryJobInfo := _InventoryJobInfo{}

	err = json.Unmarshal(data, &varInventoryJobInfo)
	if err == nil {
		o.MoBaseComplexType = varInventoryJobInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExecutionStatus")
		delete(additionalProperties, "JobName")
		delete(additionalProperties, "LastProcessedTime")
		delete(additionalProperties, "LastScheduledTime")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "Regex")

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

type NullableInventoryJobInfo struct {
	value *InventoryJobInfo
	isSet bool
}

func (v NullableInventoryJobInfo) Get() *InventoryJobInfo {
	return v.value
}

func (v *NullableInventoryJobInfo) Set(val *InventoryJobInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryJobInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryJobInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryJobInfo(val *InventoryJobInfo) *NullableInventoryJobInfo {
	return &NullableInventoryJobInfo{value: val, isSet: true}
}

func (v NullableInventoryJobInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryJobInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
