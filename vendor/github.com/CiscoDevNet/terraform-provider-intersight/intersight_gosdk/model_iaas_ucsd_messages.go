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

// checks if the IaasUcsdMessages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IaasUcsdMessages{}

// IaasUcsdMessages Gets ucsd messages from UCSD.
type IaasUcsdMessages struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Last checked time of the alerts.
	StatusId             *string                                     `json:"StatusId,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasUcsdMessages IaasUcsdMessages

// NewIaasUcsdMessages instantiates a new IaasUcsdMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasUcsdMessages(classId string, objectType string) *IaasUcsdMessages {
	this := IaasUcsdMessages{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasUcsdMessagesWithDefaults instantiates a new IaasUcsdMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasUcsdMessagesWithDefaults() *IaasUcsdMessages {
	this := IaasUcsdMessages{}
	var classId string = "iaas.UcsdMessages"
	this.ClassId = classId
	var objectType string = "iaas.UcsdMessages"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasUcsdMessages) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdMessages) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasUcsdMessages) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iaas.UcsdMessages" of the ClassId field.
func (o *IaasUcsdMessages) GetDefaultClassId() interface{} {
	return "iaas.UcsdMessages"
}

// GetObjectType returns the ObjectType field value
func (o *IaasUcsdMessages) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdMessages) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasUcsdMessages) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iaas.UcsdMessages" of the ObjectType field.
func (o *IaasUcsdMessages) GetDefaultObjectType() interface{} {
	return "iaas.UcsdMessages"
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *IaasUcsdMessages) GetStatusId() string {
	if o == nil || IsNil(o.StatusId) {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdMessages) GetStatusIdOk() (*string, bool) {
	if o == nil || IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *IaasUcsdMessages) HasStatusId() bool {
	if o != nil && !IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *IaasUcsdMessages) SetStatusId(v string) {
	o.StatusId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdMessages) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdMessages) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasUcsdMessages) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasUcsdMessages) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *IaasUcsdMessages) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *IaasUcsdMessages) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o IaasUcsdMessages) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IaasUcsdMessages) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StatusId) {
		toSerialize["StatusId"] = o.StatusId
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IaasUcsdMessages) UnmarshalJSON(data []byte) (err error) {
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
	type IaasUcsdMessagesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Last checked time of the alerts.
		StatusId         *string                                     `json:"StatusId,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varIaasUcsdMessagesWithoutEmbeddedStruct := IaasUcsdMessagesWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIaasUcsdMessagesWithoutEmbeddedStruct)
	if err == nil {
		varIaasUcsdMessages := _IaasUcsdMessages{}
		varIaasUcsdMessages.ClassId = varIaasUcsdMessagesWithoutEmbeddedStruct.ClassId
		varIaasUcsdMessages.ObjectType = varIaasUcsdMessagesWithoutEmbeddedStruct.ObjectType
		varIaasUcsdMessages.StatusId = varIaasUcsdMessagesWithoutEmbeddedStruct.StatusId
		varIaasUcsdMessages.RegisteredDevice = varIaasUcsdMessagesWithoutEmbeddedStruct.RegisteredDevice
		*o = IaasUcsdMessages(varIaasUcsdMessages)
	} else {
		return err
	}

	varIaasUcsdMessages := _IaasUcsdMessages{}

	err = json.Unmarshal(data, &varIaasUcsdMessages)
	if err == nil {
		o.MoBaseMo = varIaasUcsdMessages.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "StatusId")
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

type NullableIaasUcsdMessages struct {
	value *IaasUcsdMessages
	isSet bool
}

func (v NullableIaasUcsdMessages) Get() *IaasUcsdMessages {
	return v.value
}

func (v *NullableIaasUcsdMessages) Set(val *IaasUcsdMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasUcsdMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasUcsdMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasUcsdMessages(val *IaasUcsdMessages) *NullableIaasUcsdMessages {
	return &NullableIaasUcsdMessages{value: val, isSet: true}
}

func (v NullableIaasUcsdMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasUcsdMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
