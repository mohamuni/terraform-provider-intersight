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

// checks if the StorageHitachiVolumeMigrationPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageHitachiVolumeMigrationPair{}

// StorageHitachiVolumeMigrationPair A copy pair to be used for Volume Migration in Hitachi storage array.
type StorageHitachiVolumeMigrationPair struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Copy mode. NotSynchronized or VolumeMigration is stored.
	CopyMode *string `json:"CopyMode,omitempty"`
	// Object ID of the pair. The following informations of pair are output, separated by commas. <copy group name>, <device group name for the P-VOL (source volume)>, <device group name for the S-VOL (target volume)>, <name of the pair>.
	LocalCloneCopypairId *string `json:"LocalCloneCopypairId,omitempty"`
	// LDEV number of the P-VOL (source volume) with a decimal (base 10) number.
	PvolLdevId *int64 `json:"PvolLdevId,omitempty"`
	// Pair volume status of the P-VOL.
	PvolStatus *string `json:"PvolStatus,omitempty"`
	// LDEV number of the S-VOL (target volume) with a decimal (base 10) number.
	SvolLdevId *int64 `json:"SvolLdevId,omitempty"`
	// Pair volume status of the S-VOL.
	SvolStatus           *string                                     `json:"SvolStatus,omitempty"`
	Array                NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiVolumeMigrationPair StorageHitachiVolumeMigrationPair

// NewStorageHitachiVolumeMigrationPair instantiates a new StorageHitachiVolumeMigrationPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiVolumeMigrationPair(classId string, objectType string) *StorageHitachiVolumeMigrationPair {
	this := StorageHitachiVolumeMigrationPair{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiVolumeMigrationPairWithDefaults instantiates a new StorageHitachiVolumeMigrationPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiVolumeMigrationPairWithDefaults() *StorageHitachiVolumeMigrationPair {
	this := StorageHitachiVolumeMigrationPair{}
	var classId string = "storage.HitachiVolumeMigrationPair"
	this.ClassId = classId
	var objectType string = "storage.HitachiVolumeMigrationPair"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiVolumeMigrationPair) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiVolumeMigrationPair) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.HitachiVolumeMigrationPair" of the ClassId field.
func (o *StorageHitachiVolumeMigrationPair) GetDefaultClassId() interface{} {
	return "storage.HitachiVolumeMigrationPair"
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiVolumeMigrationPair) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiVolumeMigrationPair) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.HitachiVolumeMigrationPair" of the ObjectType field.
func (o *StorageHitachiVolumeMigrationPair) GetDefaultObjectType() interface{} {
	return "storage.HitachiVolumeMigrationPair"
}

// GetCopyMode returns the CopyMode field value if set, zero value otherwise.
func (o *StorageHitachiVolumeMigrationPair) GetCopyMode() string {
	if o == nil || IsNil(o.CopyMode) {
		var ret string
		return ret
	}
	return *o.CopyMode
}

// GetCopyModeOk returns a tuple with the CopyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetCopyModeOk() (*string, bool) {
	if o == nil || IsNil(o.CopyMode) {
		return nil, false
	}
	return o.CopyMode, true
}

// HasCopyMode returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasCopyMode() bool {
	if o != nil && !IsNil(o.CopyMode) {
		return true
	}

	return false
}

// SetCopyMode gets a reference to the given string and assigns it to the CopyMode field.
func (o *StorageHitachiVolumeMigrationPair) SetCopyMode(v string) {
	o.CopyMode = &v
}

// GetLocalCloneCopypairId returns the LocalCloneCopypairId field value if set, zero value otherwise.
func (o *StorageHitachiVolumeMigrationPair) GetLocalCloneCopypairId() string {
	if o == nil || IsNil(o.LocalCloneCopypairId) {
		var ret string
		return ret
	}
	return *o.LocalCloneCopypairId
}

// GetLocalCloneCopypairIdOk returns a tuple with the LocalCloneCopypairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetLocalCloneCopypairIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocalCloneCopypairId) {
		return nil, false
	}
	return o.LocalCloneCopypairId, true
}

// HasLocalCloneCopypairId returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasLocalCloneCopypairId() bool {
	if o != nil && !IsNil(o.LocalCloneCopypairId) {
		return true
	}

	return false
}

// SetLocalCloneCopypairId gets a reference to the given string and assigns it to the LocalCloneCopypairId field.
func (o *StorageHitachiVolumeMigrationPair) SetLocalCloneCopypairId(v string) {
	o.LocalCloneCopypairId = &v
}

// GetPvolLdevId returns the PvolLdevId field value if set, zero value otherwise.
func (o *StorageHitachiVolumeMigrationPair) GetPvolLdevId() int64 {
	if o == nil || IsNil(o.PvolLdevId) {
		var ret int64
		return ret
	}
	return *o.PvolLdevId
}

// GetPvolLdevIdOk returns a tuple with the PvolLdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetPvolLdevIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PvolLdevId) {
		return nil, false
	}
	return o.PvolLdevId, true
}

// HasPvolLdevId returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasPvolLdevId() bool {
	if o != nil && !IsNil(o.PvolLdevId) {
		return true
	}

	return false
}

// SetPvolLdevId gets a reference to the given int64 and assigns it to the PvolLdevId field.
func (o *StorageHitachiVolumeMigrationPair) SetPvolLdevId(v int64) {
	o.PvolLdevId = &v
}

// GetPvolStatus returns the PvolStatus field value if set, zero value otherwise.
func (o *StorageHitachiVolumeMigrationPair) GetPvolStatus() string {
	if o == nil || IsNil(o.PvolStatus) {
		var ret string
		return ret
	}
	return *o.PvolStatus
}

// GetPvolStatusOk returns a tuple with the PvolStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetPvolStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PvolStatus) {
		return nil, false
	}
	return o.PvolStatus, true
}

// HasPvolStatus returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasPvolStatus() bool {
	if o != nil && !IsNil(o.PvolStatus) {
		return true
	}

	return false
}

// SetPvolStatus gets a reference to the given string and assigns it to the PvolStatus field.
func (o *StorageHitachiVolumeMigrationPair) SetPvolStatus(v string) {
	o.PvolStatus = &v
}

// GetSvolLdevId returns the SvolLdevId field value if set, zero value otherwise.
func (o *StorageHitachiVolumeMigrationPair) GetSvolLdevId() int64 {
	if o == nil || IsNil(o.SvolLdevId) {
		var ret int64
		return ret
	}
	return *o.SvolLdevId
}

// GetSvolLdevIdOk returns a tuple with the SvolLdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetSvolLdevIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SvolLdevId) {
		return nil, false
	}
	return o.SvolLdevId, true
}

// HasSvolLdevId returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasSvolLdevId() bool {
	if o != nil && !IsNil(o.SvolLdevId) {
		return true
	}

	return false
}

// SetSvolLdevId gets a reference to the given int64 and assigns it to the SvolLdevId field.
func (o *StorageHitachiVolumeMigrationPair) SetSvolLdevId(v int64) {
	o.SvolLdevId = &v
}

// GetSvolStatus returns the SvolStatus field value if set, zero value otherwise.
func (o *StorageHitachiVolumeMigrationPair) GetSvolStatus() string {
	if o == nil || IsNil(o.SvolStatus) {
		var ret string
		return ret
	}
	return *o.SvolStatus
}

// GetSvolStatusOk returns a tuple with the SvolStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolumeMigrationPair) GetSvolStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SvolStatus) {
		return nil, false
	}
	return o.SvolStatus, true
}

// HasSvolStatus returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasSvolStatus() bool {
	if o != nil && !IsNil(o.SvolStatus) {
		return true
	}

	return false
}

// SetSvolStatus gets a reference to the given string and assigns it to the SvolStatus field.
func (o *StorageHitachiVolumeMigrationPair) SetSvolStatus(v string) {
	o.SvolStatus = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiVolumeMigrationPair) GetArray() StorageHitachiArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiVolumeMigrationPair) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiVolumeMigrationPair) SetArray(v StorageHitachiArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageHitachiVolumeMigrationPair) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageHitachiVolumeMigrationPair) UnsetArray() {
	o.Array.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiVolumeMigrationPair) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiVolumeMigrationPair) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiVolumeMigrationPair) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiVolumeMigrationPair) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageHitachiVolumeMigrationPair) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageHitachiVolumeMigrationPair) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageHitachiVolumeMigrationPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageHitachiVolumeMigrationPair) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CopyMode) {
		toSerialize["CopyMode"] = o.CopyMode
	}
	if !IsNil(o.LocalCloneCopypairId) {
		toSerialize["LocalCloneCopypairId"] = o.LocalCloneCopypairId
	}
	if !IsNil(o.PvolLdevId) {
		toSerialize["PvolLdevId"] = o.PvolLdevId
	}
	if !IsNil(o.PvolStatus) {
		toSerialize["PvolStatus"] = o.PvolStatus
	}
	if !IsNil(o.SvolLdevId) {
		toSerialize["SvolLdevId"] = o.SvolLdevId
	}
	if !IsNil(o.SvolStatus) {
		toSerialize["SvolStatus"] = o.SvolStatus
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageHitachiVolumeMigrationPair) UnmarshalJSON(data []byte) (err error) {
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
	type StorageHitachiVolumeMigrationPairWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Copy mode. NotSynchronized or VolumeMigration is stored.
		CopyMode *string `json:"CopyMode,omitempty"`
		// Object ID of the pair. The following informations of pair are output, separated by commas. <copy group name>, <device group name for the P-VOL (source volume)>, <device group name for the S-VOL (target volume)>, <name of the pair>.
		LocalCloneCopypairId *string `json:"LocalCloneCopypairId,omitempty"`
		// LDEV number of the P-VOL (source volume) with a decimal (base 10) number.
		PvolLdevId *int64 `json:"PvolLdevId,omitempty"`
		// Pair volume status of the P-VOL.
		PvolStatus *string `json:"PvolStatus,omitempty"`
		// LDEV number of the S-VOL (target volume) with a decimal (base 10) number.
		SvolLdevId *int64 `json:"SvolLdevId,omitempty"`
		// Pair volume status of the S-VOL.
		SvolStatus       *string                                     `json:"SvolStatus,omitempty"`
		Array            NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct := StorageHitachiVolumeMigrationPairWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiVolumeMigrationPair := _StorageHitachiVolumeMigrationPair{}
		varStorageHitachiVolumeMigrationPair.ClassId = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.ClassId
		varStorageHitachiVolumeMigrationPair.ObjectType = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.ObjectType
		varStorageHitachiVolumeMigrationPair.CopyMode = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.CopyMode
		varStorageHitachiVolumeMigrationPair.LocalCloneCopypairId = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.LocalCloneCopypairId
		varStorageHitachiVolumeMigrationPair.PvolLdevId = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.PvolLdevId
		varStorageHitachiVolumeMigrationPair.PvolStatus = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.PvolStatus
		varStorageHitachiVolumeMigrationPair.SvolLdevId = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.SvolLdevId
		varStorageHitachiVolumeMigrationPair.SvolStatus = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.SvolStatus
		varStorageHitachiVolumeMigrationPair.Array = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.Array
		varStorageHitachiVolumeMigrationPair.RegisteredDevice = varStorageHitachiVolumeMigrationPairWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiVolumeMigrationPair(varStorageHitachiVolumeMigrationPair)
	} else {
		return err
	}

	varStorageHitachiVolumeMigrationPair := _StorageHitachiVolumeMigrationPair{}

	err = json.Unmarshal(data, &varStorageHitachiVolumeMigrationPair)
	if err == nil {
		o.MoBaseMo = varStorageHitachiVolumeMigrationPair.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CopyMode")
		delete(additionalProperties, "LocalCloneCopypairId")
		delete(additionalProperties, "PvolLdevId")
		delete(additionalProperties, "PvolStatus")
		delete(additionalProperties, "SvolLdevId")
		delete(additionalProperties, "SvolStatus")
		delete(additionalProperties, "Array")
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

type NullableStorageHitachiVolumeMigrationPair struct {
	value *StorageHitachiVolumeMigrationPair
	isSet bool
}

func (v NullableStorageHitachiVolumeMigrationPair) Get() *StorageHitachiVolumeMigrationPair {
	return v.value
}

func (v *NullableStorageHitachiVolumeMigrationPair) Set(val *StorageHitachiVolumeMigrationPair) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiVolumeMigrationPair) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiVolumeMigrationPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiVolumeMigrationPair(val *StorageHitachiVolumeMigrationPair) *NullableStorageHitachiVolumeMigrationPair {
	return &NullableStorageHitachiVolumeMigrationPair{value: val, isSet: true}
}

func (v NullableStorageHitachiVolumeMigrationPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiVolumeMigrationPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
