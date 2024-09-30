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

// checks if the StoragePureSnapshotSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePureSnapshotSchedule{}

// StoragePureSnapshotSchedule PureStorage FlashArray snapshot schedule configuration entity.
type StoragePureSnapshotSchedule struct {
	StorageBaseSnapshotSchedule
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
	DailyLimit *int64 `json:"DailyLimit,omitempty"`
	// Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period.
	SnapshotExpiryTime   *string                                        `json:"SnapshotExpiryTime,omitempty"`
	Array                NullableStoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      NullableStoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureSnapshotSchedule StoragePureSnapshotSchedule

// NewStoragePureSnapshotSchedule instantiates a new StoragePureSnapshotSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureSnapshotSchedule(classId string, objectType string) *StoragePureSnapshotSchedule {
	this := StoragePureSnapshotSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureSnapshotScheduleWithDefaults instantiates a new StoragePureSnapshotSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureSnapshotScheduleWithDefaults() *StoragePureSnapshotSchedule {
	this := StoragePureSnapshotSchedule{}
	var classId string = "storage.PureSnapshotSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureSnapshotSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureSnapshotSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureSnapshotSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureSnapshotSchedule" of the ClassId field.
func (o *StoragePureSnapshotSchedule) GetDefaultClassId() interface{} {
	return "storage.PureSnapshotSchedule"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureSnapshotSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureSnapshotSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureSnapshotSchedule" of the ObjectType field.
func (o *StoragePureSnapshotSchedule) GetDefaultObjectType() interface{} {
	return "storage.PureSnapshotSchedule"
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetDailyLimit() int64 {
	if o == nil || IsNil(o.DailyLimit) {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetDailyLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DailyLimit) {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasDailyLimit() bool {
	if o != nil && !IsNil(o.DailyLimit) {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *StoragePureSnapshotSchedule) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetSnapshotExpiryTime returns the SnapshotExpiryTime field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetSnapshotExpiryTime() string {
	if o == nil || IsNil(o.SnapshotExpiryTime) {
		var ret string
		return ret
	}
	return *o.SnapshotExpiryTime
}

// GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetSnapshotExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotExpiryTime) {
		return nil, false
	}
	return o.SnapshotExpiryTime, true
}

// HasSnapshotExpiryTime returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasSnapshotExpiryTime() bool {
	if o != nil && !IsNil(o.SnapshotExpiryTime) {
		return true
	}

	return false
}

// SetSnapshotExpiryTime gets a reference to the given string and assigns it to the SnapshotExpiryTime field.
func (o *StoragePureSnapshotSchedule) SetSnapshotExpiryTime(v string) {
	o.SnapshotExpiryTime = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureSnapshotSchedule) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureSnapshotSchedule) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureSnapshotSchedule) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePureSnapshotSchedule) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePureSnapshotSchedule) UnsetArray() {
	o.Array.Unset()
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureSnapshotSchedule) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || IsNil(o.ProtectionGroup.Get()) {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup.Get()
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureSnapshotSchedule) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroup.Get(), o.ProtectionGroup.IsSet()
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given NullableStoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureSnapshotSchedule) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup.Set(&v)
}

// SetProtectionGroupNil sets the value for ProtectionGroup to be an explicit nil
func (o *StoragePureSnapshotSchedule) SetProtectionGroupNil() {
	o.ProtectionGroup.Set(nil)
}

// UnsetProtectionGroup ensures that no value is present for ProtectionGroup, not even an explicit nil
func (o *StoragePureSnapshotSchedule) UnsetProtectionGroup() {
	o.ProtectionGroup.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureSnapshotSchedule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureSnapshotSchedule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureSnapshotSchedule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StoragePureSnapshotSchedule) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StoragePureSnapshotSchedule) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StoragePureSnapshotSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePureSnapshotSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseSnapshotSchedule, errStorageBaseSnapshotSchedule := json.Marshal(o.StorageBaseSnapshotSchedule)
	if errStorageBaseSnapshotSchedule != nil {
		return map[string]interface{}{}, errStorageBaseSnapshotSchedule
	}
	errStorageBaseSnapshotSchedule = json.Unmarshal([]byte(serializedStorageBaseSnapshotSchedule), &toSerialize)
	if errStorageBaseSnapshotSchedule != nil {
		return map[string]interface{}{}, errStorageBaseSnapshotSchedule
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DailyLimit) {
		toSerialize["DailyLimit"] = o.DailyLimit
	}
	if !IsNil(o.SnapshotExpiryTime) {
		toSerialize["SnapshotExpiryTime"] = o.SnapshotExpiryTime
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.ProtectionGroup.IsSet() {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoragePureSnapshotSchedule) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePureSnapshotScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
		DailyLimit *int64 `json:"DailyLimit,omitempty"`
		// Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period.
		SnapshotExpiryTime *string                                        `json:"SnapshotExpiryTime,omitempty"`
		Array              NullableStoragePureArrayRelationship           `json:"Array,omitempty"`
		ProtectionGroup    NullableStoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
		RegisteredDevice   NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureSnapshotScheduleWithoutEmbeddedStruct := StoragePureSnapshotScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePureSnapshotScheduleWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureSnapshotSchedule := _StoragePureSnapshotSchedule{}
		varStoragePureSnapshotSchedule.ClassId = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.ClassId
		varStoragePureSnapshotSchedule.ObjectType = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.ObjectType
		varStoragePureSnapshotSchedule.DailyLimit = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.DailyLimit
		varStoragePureSnapshotSchedule.SnapshotExpiryTime = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.SnapshotExpiryTime
		varStoragePureSnapshotSchedule.Array = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.Array
		varStoragePureSnapshotSchedule.ProtectionGroup = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.ProtectionGroup
		varStoragePureSnapshotSchedule.RegisteredDevice = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureSnapshotSchedule(varStoragePureSnapshotSchedule)
	} else {
		return err
	}

	varStoragePureSnapshotSchedule := _StoragePureSnapshotSchedule{}

	err = json.Unmarshal(data, &varStoragePureSnapshotSchedule)
	if err == nil {
		o.StorageBaseSnapshotSchedule = varStoragePureSnapshotSchedule.StorageBaseSnapshotSchedule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DailyLimit")
		delete(additionalProperties, "SnapshotExpiryTime")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseSnapshotSchedule := reflect.ValueOf(o.StorageBaseSnapshotSchedule)
		for i := 0; i < reflectStorageBaseSnapshotSchedule.Type().NumField(); i++ {
			t := reflectStorageBaseSnapshotSchedule.Type().Field(i)

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

type NullableStoragePureSnapshotSchedule struct {
	value *StoragePureSnapshotSchedule
	isSet bool
}

func (v NullableStoragePureSnapshotSchedule) Get() *StoragePureSnapshotSchedule {
	return v.value
}

func (v *NullableStoragePureSnapshotSchedule) Set(val *StoragePureSnapshotSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureSnapshotSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureSnapshotSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureSnapshotSchedule(val *StoragePureSnapshotSchedule) *NullableStoragePureSnapshotSchedule {
	return &NullableStoragePureSnapshotSchedule{value: val, isSet: true}
}

func (v NullableStoragePureSnapshotSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureSnapshotSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
