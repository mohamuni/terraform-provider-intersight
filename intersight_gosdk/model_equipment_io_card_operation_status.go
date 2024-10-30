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

// checks if the EquipmentIoCardOperationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentIoCardOperationStatus{}

// EquipmentIoCardOperationStatus The status of completed and in-progress operations performed on the IO Card.
type EquipmentIoCardOperationStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The configured state of the settings in the target IO Card. The value is any one of Applied, Applying or Failed. Applied - The state denotes that the settings are applied successfully in the target IO Card. Applying - The state denotes that the settings are being applied in the target IO Card. Failed - The state denotes that the settings could not be applied in the target IO Card. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// The workflow Id of the IO Card operations workflow.
	WorkflowId *string `json:"WorkflowId,omitempty"`
	// The workflow type of the IO Card operation workflow. This can be used to distinguish different IO Card operations.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardOperationStatus EquipmentIoCardOperationStatus

// NewEquipmentIoCardOperationStatus instantiates a new EquipmentIoCardOperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardOperationStatus(classId string, objectType string) *EquipmentIoCardOperationStatus {
	this := EquipmentIoCardOperationStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardOperationStatusWithDefaults instantiates a new EquipmentIoCardOperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardOperationStatusWithDefaults() *EquipmentIoCardOperationStatus {
	this := EquipmentIoCardOperationStatus{}
	var classId string = "equipment.IoCardOperationStatus"
	this.ClassId = classId
	var objectType string = "equipment.IoCardOperationStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardOperationStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardOperationStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.IoCardOperationStatus" of the ClassId field.
func (o *EquipmentIoCardOperationStatus) GetDefaultClassId() interface{} {
	return "equipment.IoCardOperationStatus"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardOperationStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardOperationStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.IoCardOperationStatus" of the ObjectType field.
func (o *EquipmentIoCardOperationStatus) GetDefaultObjectType() interface{} {
	return "equipment.IoCardOperationStatus"
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationStatus) GetConfigState() string {
	if o == nil || IsNil(o.ConfigState) {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationStatus) GetConfigStateOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigState) {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationStatus) HasConfigState() bool {
	if o != nil && !IsNil(o.ConfigState) {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentIoCardOperationStatus) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationStatus) GetWorkflowId() string {
	if o == nil || IsNil(o.WorkflowId) {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationStatus) GetWorkflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowId) {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationStatus) HasWorkflowId() bool {
	if o != nil && !IsNil(o.WorkflowId) {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *EquipmentIoCardOperationStatus) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationStatus) GetWorkflowType() string {
	if o == nil || IsNil(o.WorkflowType) {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationStatus) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowType) {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationStatus) HasWorkflowType() bool {
	if o != nil && !IsNil(o.WorkflowType) {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *EquipmentIoCardOperationStatus) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o EquipmentIoCardOperationStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentIoCardOperationStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConfigState) {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if !IsNil(o.WorkflowId) {
		toSerialize["WorkflowId"] = o.WorkflowId
	}
	if !IsNil(o.WorkflowType) {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentIoCardOperationStatus) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentIoCardOperationStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The configured state of the settings in the target IO Card. The value is any one of Applied, Applying or Failed. Applied - The state denotes that the settings are applied successfully in the target IO Card. Applying - The state denotes that the settings are being applied in the target IO Card. Failed - The state denotes that the settings could not be applied in the target IO Card. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState *string `json:"ConfigState,omitempty"`
		// The workflow Id of the IO Card operations workflow.
		WorkflowId *string `json:"WorkflowId,omitempty"`
		// The workflow type of the IO Card operation workflow. This can be used to distinguish different IO Card operations.
		WorkflowType *string `json:"WorkflowType,omitempty"`
	}

	varEquipmentIoCardOperationStatusWithoutEmbeddedStruct := EquipmentIoCardOperationStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentIoCardOperationStatusWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIoCardOperationStatus := _EquipmentIoCardOperationStatus{}
		varEquipmentIoCardOperationStatus.ClassId = varEquipmentIoCardOperationStatusWithoutEmbeddedStruct.ClassId
		varEquipmentIoCardOperationStatus.ObjectType = varEquipmentIoCardOperationStatusWithoutEmbeddedStruct.ObjectType
		varEquipmentIoCardOperationStatus.ConfigState = varEquipmentIoCardOperationStatusWithoutEmbeddedStruct.ConfigState
		varEquipmentIoCardOperationStatus.WorkflowId = varEquipmentIoCardOperationStatusWithoutEmbeddedStruct.WorkflowId
		varEquipmentIoCardOperationStatus.WorkflowType = varEquipmentIoCardOperationStatusWithoutEmbeddedStruct.WorkflowType
		*o = EquipmentIoCardOperationStatus(varEquipmentIoCardOperationStatus)
	} else {
		return err
	}

	varEquipmentIoCardOperationStatus := _EquipmentIoCardOperationStatus{}

	err = json.Unmarshal(data, &varEquipmentIoCardOperationStatus)
	if err == nil {
		o.MoBaseComplexType = varEquipmentIoCardOperationStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "WorkflowId")
		delete(additionalProperties, "WorkflowType")

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

type NullableEquipmentIoCardOperationStatus struct {
	value *EquipmentIoCardOperationStatus
	isSet bool
}

func (v NullableEquipmentIoCardOperationStatus) Get() *EquipmentIoCardOperationStatus {
	return v.value
}

func (v *NullableEquipmentIoCardOperationStatus) Set(val *EquipmentIoCardOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardOperationStatus(val *EquipmentIoCardOperationStatus) *NullableEquipmentIoCardOperationStatus {
	return &NullableEquipmentIoCardOperationStatus{value: val, isSet: true}
}

func (v NullableEquipmentIoCardOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
