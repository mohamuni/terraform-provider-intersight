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

// HyperflexVirtualMachine Virtual Machine from StDataServiceManager.
type HyperflexVirtualMachine struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                                     `json:"ObjectType"`
	RunTimeInfo NullableHyperflexVirtualMachineRuntimeInfo `json:"RunTimeInfo,omitempty"`
	// Virtual machine status code. * `VM_ACCESSIBLE` - This virtual machine is accessible. * `VM_INACCESSIBLE` - This virtual machine is not accessible. * `VM_NOT_SUPPORTED` - This virtual machine is not supported. * `NONE` - This virtual machine does not have a status code.
	StatusCode *string `json:"StatusCode,omitempty"`
	// Virtual machine's current UUID.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVirtualMachine HyperflexVirtualMachine

// NewHyperflexVirtualMachine instantiates a new HyperflexVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVirtualMachine(classId string, objectType string) *HyperflexVirtualMachine {
	this := HyperflexVirtualMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVirtualMachineWithDefaults instantiates a new HyperflexVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVirtualMachineWithDefaults() *HyperflexVirtualMachine {
	this := HyperflexVirtualMachine{}
	var classId string = "hyperflex.VirtualMachine"
	this.ClassId = classId
	var objectType string = "hyperflex.VirtualMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVirtualMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVirtualMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRunTimeInfo returns the RunTimeInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVirtualMachine) GetRunTimeInfo() HyperflexVirtualMachineRuntimeInfo {
	if o == nil || o.RunTimeInfo.Get() == nil {
		var ret HyperflexVirtualMachineRuntimeInfo
		return ret
	}
	return *o.RunTimeInfo.Get()
}

// GetRunTimeInfoOk returns a tuple with the RunTimeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVirtualMachine) GetRunTimeInfoOk() (*HyperflexVirtualMachineRuntimeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunTimeInfo.Get(), o.RunTimeInfo.IsSet()
}

// HasRunTimeInfo returns a boolean if a field has been set.
func (o *HyperflexVirtualMachine) HasRunTimeInfo() bool {
	if o != nil && o.RunTimeInfo.IsSet() {
		return true
	}

	return false
}

// SetRunTimeInfo gets a reference to the given NullableHyperflexVirtualMachineRuntimeInfo and assigns it to the RunTimeInfo field.
func (o *HyperflexVirtualMachine) SetRunTimeInfo(v HyperflexVirtualMachineRuntimeInfo) {
	o.RunTimeInfo.Set(&v)
}

// SetRunTimeInfoNil sets the value for RunTimeInfo to be an explicit nil
func (o *HyperflexVirtualMachine) SetRunTimeInfoNil() {
	o.RunTimeInfo.Set(nil)
}

// UnsetRunTimeInfo ensures that no value is present for RunTimeInfo, not even an explicit nil
func (o *HyperflexVirtualMachine) UnsetRunTimeInfo() {
	o.RunTimeInfo.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *HyperflexVirtualMachine) GetStatusCode() string {
	if o == nil || o.StatusCode == nil {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachine) GetStatusCodeOk() (*string, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *HyperflexVirtualMachine) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *HyperflexVirtualMachine) SetStatusCode(v string) {
	o.StatusCode = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexVirtualMachine) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachine) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexVirtualMachine) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexVirtualMachine) SetUuid(v string) {
	o.Uuid = &v
}

func (o HyperflexVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RunTimeInfo.IsSet() {
		toSerialize["RunTimeInfo"] = o.RunTimeInfo.Get()
	}
	if o.StatusCode != nil {
		toSerialize["StatusCode"] = o.StatusCode
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexVirtualMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                                     `json:"ObjectType"`
		RunTimeInfo NullableHyperflexVirtualMachineRuntimeInfo `json:"RunTimeInfo,omitempty"`
		// Virtual machine status code. * `VM_ACCESSIBLE` - This virtual machine is accessible. * `VM_INACCESSIBLE` - This virtual machine is not accessible. * `VM_NOT_SUPPORTED` - This virtual machine is not supported. * `NONE` - This virtual machine does not have a status code.
		StatusCode *string `json:"StatusCode,omitempty"`
		// Virtual machine's current UUID.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varHyperflexVirtualMachineWithoutEmbeddedStruct := HyperflexVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexVirtualMachine := _HyperflexVirtualMachine{}
		varHyperflexVirtualMachine.ClassId = varHyperflexVirtualMachineWithoutEmbeddedStruct.ClassId
		varHyperflexVirtualMachine.ObjectType = varHyperflexVirtualMachineWithoutEmbeddedStruct.ObjectType
		varHyperflexVirtualMachine.RunTimeInfo = varHyperflexVirtualMachineWithoutEmbeddedStruct.RunTimeInfo
		varHyperflexVirtualMachine.StatusCode = varHyperflexVirtualMachineWithoutEmbeddedStruct.StatusCode
		varHyperflexVirtualMachine.Uuid = varHyperflexVirtualMachineWithoutEmbeddedStruct.Uuid
		*o = HyperflexVirtualMachine(varHyperflexVirtualMachine)
	} else {
		return err
	}

	varHyperflexVirtualMachine := _HyperflexVirtualMachine{}

	err = json.Unmarshal(bytes, &varHyperflexVirtualMachine)
	if err == nil {
		o.MoBaseComplexType = varHyperflexVirtualMachine.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RunTimeInfo")
		delete(additionalProperties, "StatusCode")
		delete(additionalProperties, "Uuid")

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

type NullableHyperflexVirtualMachine struct {
	value *HyperflexVirtualMachine
	isSet bool
}

func (v NullableHyperflexVirtualMachine) Get() *HyperflexVirtualMachine {
	return v.value
}

func (v *NullableHyperflexVirtualMachine) Set(val *HyperflexVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVirtualMachine(val *HyperflexVirtualMachine) *NullableHyperflexVirtualMachine {
	return &NullableHyperflexVirtualMachine{value: val, isSet: true}
}

func (v NullableHyperflexVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
