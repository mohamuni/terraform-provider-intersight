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

// checks if the HyperflexVmImportOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexVmImportOperation{}

// HyperflexVmImportOperation Invoke Virtual Machine import operation.
type HyperflexVmImportOperation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                       `json:"ObjectType"`
	DeviceMoid           NullableAssetDeviceRegistrationRelationship  `json:"DeviceMoid,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVmImportOperation HyperflexVmImportOperation

// NewHyperflexVmImportOperation instantiates a new HyperflexVmImportOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmImportOperation(classId string, objectType string) *HyperflexVmImportOperation {
	this := HyperflexVmImportOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVmImportOperationWithDefaults instantiates a new HyperflexVmImportOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmImportOperationWithDefaults() *HyperflexVmImportOperation {
	this := HyperflexVmImportOperation{}
	var classId string = "hyperflex.VmImportOperation"
	this.ClassId = classId
	var objectType string = "hyperflex.VmImportOperation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmImportOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmImportOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmImportOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.VmImportOperation" of the ClassId field.
func (o *HyperflexVmImportOperation) GetDefaultClassId() interface{} {
	return "hyperflex.VmImportOperation"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmImportOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmImportOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmImportOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.VmImportOperation" of the ObjectType field.
func (o *HyperflexVmImportOperation) GetDefaultObjectType() interface{} {
	return "hyperflex.VmImportOperation"
}

// GetDeviceMoid returns the DeviceMoid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmImportOperation) GetDeviceMoid() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.DeviceMoid.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceMoid.Get()
}

// GetDeviceMoidOk returns a tuple with the DeviceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmImportOperation) GetDeviceMoidOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceMoid.Get(), o.DeviceMoid.IsSet()
}

// HasDeviceMoid returns a boolean if a field has been set.
func (o *HyperflexVmImportOperation) HasDeviceMoid() bool {
	if o != nil && o.DeviceMoid.IsSet() {
		return true
	}

	return false
}

// SetDeviceMoid gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the DeviceMoid field.
func (o *HyperflexVmImportOperation) SetDeviceMoid(v AssetDeviceRegistrationRelationship) {
	o.DeviceMoid.Set(&v)
}

// SetDeviceMoidNil sets the value for DeviceMoid to be an explicit nil
func (o *HyperflexVmImportOperation) SetDeviceMoidNil() {
	o.DeviceMoid.Set(nil)
}

// UnsetDeviceMoid ensures that no value is present for DeviceMoid, not even an explicit nil
func (o *HyperflexVmImportOperation) UnsetDeviceMoid() {
	o.DeviceMoid.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmImportOperation) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmImportOperation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexVmImportOperation) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexVmImportOperation) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *HyperflexVmImportOperation) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *HyperflexVmImportOperation) UnsetOrganization() {
	o.Organization.Unset()
}

func (o HyperflexVmImportOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexVmImportOperation) ToMap() (map[string]interface{}, error) {
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
	if o.DeviceMoid.IsSet() {
		toSerialize["DeviceMoid"] = o.DeviceMoid.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexVmImportOperation) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexVmImportOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                       `json:"ObjectType"`
		DeviceMoid   NullableAssetDeviceRegistrationRelationship  `json:"DeviceMoid,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexVmImportOperationWithoutEmbeddedStruct := HyperflexVmImportOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexVmImportOperationWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexVmImportOperation := _HyperflexVmImportOperation{}
		varHyperflexVmImportOperation.ClassId = varHyperflexVmImportOperationWithoutEmbeddedStruct.ClassId
		varHyperflexVmImportOperation.ObjectType = varHyperflexVmImportOperationWithoutEmbeddedStruct.ObjectType
		varHyperflexVmImportOperation.DeviceMoid = varHyperflexVmImportOperationWithoutEmbeddedStruct.DeviceMoid
		varHyperflexVmImportOperation.Organization = varHyperflexVmImportOperationWithoutEmbeddedStruct.Organization
		*o = HyperflexVmImportOperation(varHyperflexVmImportOperation)
	} else {
		return err
	}

	varHyperflexVmImportOperation := _HyperflexVmImportOperation{}

	err = json.Unmarshal(data, &varHyperflexVmImportOperation)
	if err == nil {
		o.MoBaseMo = varHyperflexVmImportOperation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceMoid")
		delete(additionalProperties, "Organization")

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

type NullableHyperflexVmImportOperation struct {
	value *HyperflexVmImportOperation
	isSet bool
}

func (v NullableHyperflexVmImportOperation) Get() *HyperflexVmImportOperation {
	return v.value
}

func (v *NullableHyperflexVmImportOperation) Set(val *HyperflexVmImportOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmImportOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmImportOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmImportOperation(val *HyperflexVmImportOperation) *NullableHyperflexVmImportOperation {
	return &NullableHyperflexVmImportOperation{value: val, isSet: true}
}

func (v NullableHyperflexVmImportOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmImportOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
