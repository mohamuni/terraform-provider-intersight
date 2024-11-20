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
)

// checks if the StorageInternalMoPhysicalDiskConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageInternalMoPhysicalDiskConfig{}

// StorageInternalMoPhysicalDiskConfig InternalMoPhysicalDiskConfig models physical disk config that needs to be created for Secure Drives Config on reboot. Used in activation workflow.
type StorageInternalMoPhysicalDiskConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Physical Disk Encryption operation that is to be set at endpoint.
	EncryptionOp *string `json:"EncryptionOp,omitempty"`
	// Physical Disk Slot that is to be configured.
	Slot *string `json:"Slot,omitempty"`
	// Physical Disk State that is to be set at endpoint.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageInternalMoPhysicalDiskConfig StorageInternalMoPhysicalDiskConfig

// NewStorageInternalMoPhysicalDiskConfig instantiates a new StorageInternalMoPhysicalDiskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageInternalMoPhysicalDiskConfig(classId string, objectType string) *StorageInternalMoPhysicalDiskConfig {
	this := StorageInternalMoPhysicalDiskConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageInternalMoPhysicalDiskConfigWithDefaults instantiates a new StorageInternalMoPhysicalDiskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageInternalMoPhysicalDiskConfigWithDefaults() *StorageInternalMoPhysicalDiskConfig {
	this := StorageInternalMoPhysicalDiskConfig{}
	var classId string = "storage.InternalMoPhysicalDiskConfig"
	this.ClassId = classId
	var objectType string = "storage.InternalMoPhysicalDiskConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageInternalMoPhysicalDiskConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageInternalMoPhysicalDiskConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageInternalMoPhysicalDiskConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.InternalMoPhysicalDiskConfig" of the ClassId field.
func (o *StorageInternalMoPhysicalDiskConfig) GetDefaultClassId() interface{} {
	return "storage.InternalMoPhysicalDiskConfig"
}

// GetObjectType returns the ObjectType field value
func (o *StorageInternalMoPhysicalDiskConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageInternalMoPhysicalDiskConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageInternalMoPhysicalDiskConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.InternalMoPhysicalDiskConfig" of the ObjectType field.
func (o *StorageInternalMoPhysicalDiskConfig) GetDefaultObjectType() interface{} {
	return "storage.InternalMoPhysicalDiskConfig"
}

// GetEncryptionOp returns the EncryptionOp field value if set, zero value otherwise.
func (o *StorageInternalMoPhysicalDiskConfig) GetEncryptionOp() string {
	if o == nil || IsNil(o.EncryptionOp) {
		var ret string
		return ret
	}
	return *o.EncryptionOp
}

// GetEncryptionOpOk returns a tuple with the EncryptionOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageInternalMoPhysicalDiskConfig) GetEncryptionOpOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionOp) {
		return nil, false
	}
	return o.EncryptionOp, true
}

// HasEncryptionOp returns a boolean if a field has been set.
func (o *StorageInternalMoPhysicalDiskConfig) HasEncryptionOp() bool {
	if o != nil && !IsNil(o.EncryptionOp) {
		return true
	}

	return false
}

// SetEncryptionOp gets a reference to the given string and assigns it to the EncryptionOp field.
func (o *StorageInternalMoPhysicalDiskConfig) SetEncryptionOp(v string) {
	o.EncryptionOp = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *StorageInternalMoPhysicalDiskConfig) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageInternalMoPhysicalDiskConfig) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *StorageInternalMoPhysicalDiskConfig) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *StorageInternalMoPhysicalDiskConfig) SetSlot(v string) {
	o.Slot = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageInternalMoPhysicalDiskConfig) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageInternalMoPhysicalDiskConfig) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageInternalMoPhysicalDiskConfig) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageInternalMoPhysicalDiskConfig) SetState(v string) {
	o.State = &v
}

func (o StorageInternalMoPhysicalDiskConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageInternalMoPhysicalDiskConfig) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EncryptionOp) {
		toSerialize["EncryptionOp"] = o.EncryptionOp
	}
	if !IsNil(o.Slot) {
		toSerialize["Slot"] = o.Slot
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageInternalMoPhysicalDiskConfig) UnmarshalJSON(data []byte) (err error) {
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
	type StorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Physical Disk Encryption operation that is to be set at endpoint.
		EncryptionOp *string `json:"EncryptionOp,omitempty"`
		// Physical Disk Slot that is to be configured.
		Slot *string `json:"Slot,omitempty"`
		// Physical Disk State that is to be set at endpoint.
		State *string `json:"State,omitempty"`
	}

	varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct := StorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct)
	if err == nil {
		varStorageInternalMoPhysicalDiskConfig := _StorageInternalMoPhysicalDiskConfig{}
		varStorageInternalMoPhysicalDiskConfig.ClassId = varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct.ClassId
		varStorageInternalMoPhysicalDiskConfig.ObjectType = varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct.ObjectType
		varStorageInternalMoPhysicalDiskConfig.EncryptionOp = varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct.EncryptionOp
		varStorageInternalMoPhysicalDiskConfig.Slot = varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct.Slot
		varStorageInternalMoPhysicalDiskConfig.State = varStorageInternalMoPhysicalDiskConfigWithoutEmbeddedStruct.State
		*o = StorageInternalMoPhysicalDiskConfig(varStorageInternalMoPhysicalDiskConfig)
	} else {
		return err
	}

	varStorageInternalMoPhysicalDiskConfig := _StorageInternalMoPhysicalDiskConfig{}

	err = json.Unmarshal(data, &varStorageInternalMoPhysicalDiskConfig)
	if err == nil {
		o.MoBaseComplexType = varStorageInternalMoPhysicalDiskConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EncryptionOp")
		delete(additionalProperties, "Slot")
		delete(additionalProperties, "State")

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

type NullableStorageInternalMoPhysicalDiskConfig struct {
	value *StorageInternalMoPhysicalDiskConfig
	isSet bool
}

func (v NullableStorageInternalMoPhysicalDiskConfig) Get() *StorageInternalMoPhysicalDiskConfig {
	return v.value
}

func (v *NullableStorageInternalMoPhysicalDiskConfig) Set(val *StorageInternalMoPhysicalDiskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageInternalMoPhysicalDiskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageInternalMoPhysicalDiskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageInternalMoPhysicalDiskConfig(val *StorageInternalMoPhysicalDiskConfig) *NullableStorageInternalMoPhysicalDiskConfig {
	return &NullableStorageInternalMoPhysicalDiskConfig{value: val, isSet: true}
}

func (v NullableStorageInternalMoPhysicalDiskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageInternalMoPhysicalDiskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
