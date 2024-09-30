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

// checks if the WorkflowBatchExecutor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowBatchExecutor{}

// WorkflowBatchExecutor Intersight allows generic API tasks to be created by taking the API request body and a response parser specification in the form of content.Grammar object. Batch API associates the list of API requests to be executed as part of single task execution. Each API request takes the request body and a response parser specification.
type WorkflowBatchExecutor struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType   string                          `json:"ObjectType"`
	Batch        []WorkflowApi                   `json:"Batch,omitempty"`
	CancelAction []WorkflowApi                   `json:"CancelAction,omitempty"`
	Constraints  NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
	// A detailed description about the batch APIs.
	Description *string `json:"Description,omitempty"`
	// Name for the batch API task.
	Name *string `json:"Name,omitempty"`
	// All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
	Outcomes interface{} `json:"Outcomes,omitempty"`
	// Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property.
	Output interface{} `json:"Output,omitempty"`
	// When an execution of a nth API in the Batch fails, Retry from failed API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false.
	RetryFromFailedApi *bool `json:"RetryFromFailedApi,omitempty"`
	// The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
	SkipOnCondition *string `json:"SkipOnCondition,omitempty"`
	// This will hold the data needed for task to be rendered in the user interface.
	UiRenderingData      interface{} `json:"UiRenderingData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBatchExecutor WorkflowBatchExecutor

// NewWorkflowBatchExecutor instantiates a new WorkflowBatchExecutor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBatchExecutor(classId string, objectType string) *WorkflowBatchExecutor {
	this := WorkflowBatchExecutor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBatchExecutorWithDefaults instantiates a new WorkflowBatchExecutor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBatchExecutorWithDefaults() *WorkflowBatchExecutor {
	this := WorkflowBatchExecutor{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBatchExecutor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBatchExecutor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBatchExecutor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBatchExecutor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBatch returns the Batch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutor) GetBatch() []WorkflowApi {
	if o == nil {
		var ret []WorkflowApi
		return ret
	}
	return o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutor) GetBatchOk() ([]WorkflowApi, bool) {
	if o == nil || IsNil(o.Batch) {
		return nil, false
	}
	return o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasBatch() bool {
	if o != nil && !IsNil(o.Batch) {
		return true
	}

	return false
}

// SetBatch gets a reference to the given []WorkflowApi and assigns it to the Batch field.
func (o *WorkflowBatchExecutor) SetBatch(v []WorkflowApi) {
	o.Batch = v
}

// GetCancelAction returns the CancelAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutor) GetCancelAction() []WorkflowApi {
	if o == nil {
		var ret []WorkflowApi
		return ret
	}
	return o.CancelAction
}

// GetCancelActionOk returns a tuple with the CancelAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutor) GetCancelActionOk() ([]WorkflowApi, bool) {
	if o == nil || IsNil(o.CancelAction) {
		return nil, false
	}
	return o.CancelAction, true
}

// HasCancelAction returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasCancelAction() bool {
	if o != nil && !IsNil(o.CancelAction) {
		return true
	}

	return false
}

// SetCancelAction gets a reference to the given []WorkflowApi and assigns it to the CancelAction field.
func (o *WorkflowBatchExecutor) SetCancelAction(v []WorkflowApi) {
	o.CancelAction = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutor) GetConstraints() WorkflowTaskConstraints {
	if o == nil || IsNil(o.Constraints.Get()) {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutor) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableWorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowBatchExecutor) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *WorkflowBatchExecutor) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *WorkflowBatchExecutor) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowBatchExecutor) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutor) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowBatchExecutor) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBatchExecutor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBatchExecutor) SetName(v string) {
	o.Name = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutor) GetOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutor) GetOutcomesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Outcomes) {
		return nil, false
	}
	return &o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasOutcomes() bool {
	if o != nil && !IsNil(o.Outcomes) {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given interface{} and assigns it to the Outcomes field.
func (o *WorkflowBatchExecutor) SetOutcomes(v interface{}) {
	o.Outcomes = v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutor) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutor) GetOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowBatchExecutor) SetOutput(v interface{}) {
	o.Output = v
}

// GetRetryFromFailedApi returns the RetryFromFailedApi field value if set, zero value otherwise.
func (o *WorkflowBatchExecutor) GetRetryFromFailedApi() bool {
	if o == nil || IsNil(o.RetryFromFailedApi) {
		var ret bool
		return ret
	}
	return *o.RetryFromFailedApi
}

// GetRetryFromFailedApiOk returns a tuple with the RetryFromFailedApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutor) GetRetryFromFailedApiOk() (*bool, bool) {
	if o == nil || IsNil(o.RetryFromFailedApi) {
		return nil, false
	}
	return o.RetryFromFailedApi, true
}

// HasRetryFromFailedApi returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasRetryFromFailedApi() bool {
	if o != nil && !IsNil(o.RetryFromFailedApi) {
		return true
	}

	return false
}

// SetRetryFromFailedApi gets a reference to the given bool and assigns it to the RetryFromFailedApi field.
func (o *WorkflowBatchExecutor) SetRetryFromFailedApi(v bool) {
	o.RetryFromFailedApi = &v
}

// GetSkipOnCondition returns the SkipOnCondition field value if set, zero value otherwise.
func (o *WorkflowBatchExecutor) GetSkipOnCondition() string {
	if o == nil || IsNil(o.SkipOnCondition) {
		var ret string
		return ret
	}
	return *o.SkipOnCondition
}

// GetSkipOnConditionOk returns a tuple with the SkipOnCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutor) GetSkipOnConditionOk() (*string, bool) {
	if o == nil || IsNil(o.SkipOnCondition) {
		return nil, false
	}
	return o.SkipOnCondition, true
}

// HasSkipOnCondition returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasSkipOnCondition() bool {
	if o != nil && !IsNil(o.SkipOnCondition) {
		return true
	}

	return false
}

// SetSkipOnCondition gets a reference to the given string and assigns it to the SkipOnCondition field.
func (o *WorkflowBatchExecutor) SetSkipOnCondition(v string) {
	o.SkipOnCondition = &v
}

// GetUiRenderingData returns the UiRenderingData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutor) GetUiRenderingData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiRenderingData
}

// GetUiRenderingDataOk returns a tuple with the UiRenderingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutor) GetUiRenderingDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UiRenderingData) {
		return nil, false
	}
	return &o.UiRenderingData, true
}

// HasUiRenderingData returns a boolean if a field has been set.
func (o *WorkflowBatchExecutor) HasUiRenderingData() bool {
	if o != nil && !IsNil(o.UiRenderingData) {
		return true
	}

	return false
}

// SetUiRenderingData gets a reference to the given interface{} and assigns it to the UiRenderingData field.
func (o *WorkflowBatchExecutor) SetUiRenderingData(v interface{}) {
	o.UiRenderingData = v
}

func (o WorkflowBatchExecutor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowBatchExecutor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.Batch != nil {
		toSerialize["Batch"] = o.Batch
	}
	if o.CancelAction != nil {
		toSerialize["CancelAction"] = o.CancelAction
	}
	if o.Constraints.IsSet() {
		toSerialize["Constraints"] = o.Constraints.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Outcomes != nil {
		toSerialize["Outcomes"] = o.Outcomes
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if !IsNil(o.RetryFromFailedApi) {
		toSerialize["RetryFromFailedApi"] = o.RetryFromFailedApi
	}
	if !IsNil(o.SkipOnCondition) {
		toSerialize["SkipOnCondition"] = o.SkipOnCondition
	}
	if o.UiRenderingData != nil {
		toSerialize["UiRenderingData"] = o.UiRenderingData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowBatchExecutor) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowBatchExecutorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType   string                          `json:"ObjectType"`
		Batch        []WorkflowApi                   `json:"Batch,omitempty"`
		CancelAction []WorkflowApi                   `json:"CancelAction,omitempty"`
		Constraints  NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
		// A detailed description about the batch APIs.
		Description *string `json:"Description,omitempty"`
		// Name for the batch API task.
		Name *string `json:"Name,omitempty"`
		// All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
		Outcomes interface{} `json:"Outcomes,omitempty"`
		// Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property.
		Output interface{} `json:"Output,omitempty"`
		// When an execution of a nth API in the Batch fails, Retry from failed API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false.
		RetryFromFailedApi *bool `json:"RetryFromFailedApi,omitempty"`
		// The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
		SkipOnCondition *string `json:"SkipOnCondition,omitempty"`
		// This will hold the data needed for task to be rendered in the user interface.
		UiRenderingData interface{} `json:"UiRenderingData,omitempty"`
	}

	varWorkflowBatchExecutorWithoutEmbeddedStruct := WorkflowBatchExecutorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowBatchExecutorWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowBatchExecutor := _WorkflowBatchExecutor{}
		varWorkflowBatchExecutor.ClassId = varWorkflowBatchExecutorWithoutEmbeddedStruct.ClassId
		varWorkflowBatchExecutor.ObjectType = varWorkflowBatchExecutorWithoutEmbeddedStruct.ObjectType
		varWorkflowBatchExecutor.Batch = varWorkflowBatchExecutorWithoutEmbeddedStruct.Batch
		varWorkflowBatchExecutor.CancelAction = varWorkflowBatchExecutorWithoutEmbeddedStruct.CancelAction
		varWorkflowBatchExecutor.Constraints = varWorkflowBatchExecutorWithoutEmbeddedStruct.Constraints
		varWorkflowBatchExecutor.Description = varWorkflowBatchExecutorWithoutEmbeddedStruct.Description
		varWorkflowBatchExecutor.Name = varWorkflowBatchExecutorWithoutEmbeddedStruct.Name
		varWorkflowBatchExecutor.Outcomes = varWorkflowBatchExecutorWithoutEmbeddedStruct.Outcomes
		varWorkflowBatchExecutor.Output = varWorkflowBatchExecutorWithoutEmbeddedStruct.Output
		varWorkflowBatchExecutor.RetryFromFailedApi = varWorkflowBatchExecutorWithoutEmbeddedStruct.RetryFromFailedApi
		varWorkflowBatchExecutor.SkipOnCondition = varWorkflowBatchExecutorWithoutEmbeddedStruct.SkipOnCondition
		varWorkflowBatchExecutor.UiRenderingData = varWorkflowBatchExecutorWithoutEmbeddedStruct.UiRenderingData
		*o = WorkflowBatchExecutor(varWorkflowBatchExecutor)
	} else {
		return err
	}

	varWorkflowBatchExecutor := _WorkflowBatchExecutor{}

	err = json.Unmarshal(data, &varWorkflowBatchExecutor)
	if err == nil {
		o.MoBaseMo = varWorkflowBatchExecutor.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Batch")
		delete(additionalProperties, "CancelAction")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outcomes")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "RetryFromFailedApi")
		delete(additionalProperties, "SkipOnCondition")
		delete(additionalProperties, "UiRenderingData")

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

type NullableWorkflowBatchExecutor struct {
	value *WorkflowBatchExecutor
	isSet bool
}

func (v NullableWorkflowBatchExecutor) Get() *WorkflowBatchExecutor {
	return v.value
}

func (v *NullableWorkflowBatchExecutor) Set(val *WorkflowBatchExecutor) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBatchExecutor) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBatchExecutor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBatchExecutor(val *WorkflowBatchExecutor) *NullableWorkflowBatchExecutor {
	return &NullableWorkflowBatchExecutor{value: val, isSet: true}
}

func (v NullableWorkflowBatchExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBatchExecutor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
