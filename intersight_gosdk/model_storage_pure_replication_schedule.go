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

// checks if the StoragePureReplicationSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePureReplicationSchedule{}

// StoragePureReplicationSchedule Pure snapshot replication schedule entity.
type StoragePureReplicationSchedule struct {
	StorageBaseReplicationSchedule
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
	DailyLimit                   *int64                           `json:"DailyLimit,omitempty"`
	ReplicationBlackoutIntervals []StoragePureReplicationBlackout `json:"ReplicationBlackoutIntervals,omitempty"`
	// Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.
	SnapshotExpiryTime   *string                                        `json:"SnapshotExpiryTime,omitempty"`
	Array                NullableStoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      NullableStoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureReplicationSchedule StoragePureReplicationSchedule

// NewStoragePureReplicationSchedule instantiates a new StoragePureReplicationSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureReplicationSchedule(classId string, objectType string) *StoragePureReplicationSchedule {
	this := StoragePureReplicationSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureReplicationScheduleWithDefaults instantiates a new StoragePureReplicationSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureReplicationScheduleWithDefaults() *StoragePureReplicationSchedule {
	this := StoragePureReplicationSchedule{}
	var classId string = "storage.PureReplicationSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureReplicationSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureReplicationSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureReplicationSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureReplicationSchedule" of the ClassId field.
func (o *StoragePureReplicationSchedule) GetDefaultClassId() interface{} {
	return "storage.PureReplicationSchedule"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureReplicationSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureReplicationSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureReplicationSchedule" of the ObjectType field.
func (o *StoragePureReplicationSchedule) GetDefaultObjectType() interface{} {
	return "storage.PureReplicationSchedule"
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetDailyLimit() int64 {
	if o == nil || IsNil(o.DailyLimit) {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetDailyLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DailyLimit) {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasDailyLimit() bool {
	if o != nil && !IsNil(o.DailyLimit) {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *StoragePureReplicationSchedule) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetReplicationBlackoutIntervals returns the ReplicationBlackoutIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureReplicationSchedule) GetReplicationBlackoutIntervals() []StoragePureReplicationBlackout {
	if o == nil {
		var ret []StoragePureReplicationBlackout
		return ret
	}
	return o.ReplicationBlackoutIntervals
}

// GetReplicationBlackoutIntervalsOk returns a tuple with the ReplicationBlackoutIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureReplicationSchedule) GetReplicationBlackoutIntervalsOk() ([]StoragePureReplicationBlackout, bool) {
	if o == nil || IsNil(o.ReplicationBlackoutIntervals) {
		return nil, false
	}
	return o.ReplicationBlackoutIntervals, true
}

// HasReplicationBlackoutIntervals returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasReplicationBlackoutIntervals() bool {
	if o != nil && !IsNil(o.ReplicationBlackoutIntervals) {
		return true
	}

	return false
}

// SetReplicationBlackoutIntervals gets a reference to the given []StoragePureReplicationBlackout and assigns it to the ReplicationBlackoutIntervals field.
func (o *StoragePureReplicationSchedule) SetReplicationBlackoutIntervals(v []StoragePureReplicationBlackout) {
	o.ReplicationBlackoutIntervals = v
}

// GetSnapshotExpiryTime returns the SnapshotExpiryTime field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetSnapshotExpiryTime() string {
	if o == nil || IsNil(o.SnapshotExpiryTime) {
		var ret string
		return ret
	}
	return *o.SnapshotExpiryTime
}

// GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetSnapshotExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotExpiryTime) {
		return nil, false
	}
	return o.SnapshotExpiryTime, true
}

// HasSnapshotExpiryTime returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasSnapshotExpiryTime() bool {
	if o != nil && !IsNil(o.SnapshotExpiryTime) {
		return true
	}

	return false
}

// SetSnapshotExpiryTime gets a reference to the given string and assigns it to the SnapshotExpiryTime field.
func (o *StoragePureReplicationSchedule) SetSnapshotExpiryTime(v string) {
	o.SnapshotExpiryTime = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureReplicationSchedule) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureReplicationSchedule) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureReplicationSchedule) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePureReplicationSchedule) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePureReplicationSchedule) UnsetArray() {
	o.Array.Unset()
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureReplicationSchedule) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || IsNil(o.ProtectionGroup.Get()) {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup.Get()
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureReplicationSchedule) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroup.Get(), o.ProtectionGroup.IsSet()
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given NullableStoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureReplicationSchedule) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup.Set(&v)
}

// SetProtectionGroupNil sets the value for ProtectionGroup to be an explicit nil
func (o *StoragePureReplicationSchedule) SetProtectionGroupNil() {
	o.ProtectionGroup.Set(nil)
}

// UnsetProtectionGroup ensures that no value is present for ProtectionGroup, not even an explicit nil
func (o *StoragePureReplicationSchedule) UnsetProtectionGroup() {
	o.ProtectionGroup.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureReplicationSchedule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureReplicationSchedule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureReplicationSchedule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StoragePureReplicationSchedule) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StoragePureReplicationSchedule) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StoragePureReplicationSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePureReplicationSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseReplicationSchedule, errStorageBaseReplicationSchedule := json.Marshal(o.StorageBaseReplicationSchedule)
	if errStorageBaseReplicationSchedule != nil {
		return map[string]interface{}{}, errStorageBaseReplicationSchedule
	}
	errStorageBaseReplicationSchedule = json.Unmarshal([]byte(serializedStorageBaseReplicationSchedule), &toSerialize)
	if errStorageBaseReplicationSchedule != nil {
		return map[string]interface{}{}, errStorageBaseReplicationSchedule
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
	if o.ReplicationBlackoutIntervals != nil {
		toSerialize["ReplicationBlackoutIntervals"] = o.ReplicationBlackoutIntervals
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

func (o *StoragePureReplicationSchedule) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePureReplicationScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
		DailyLimit                   *int64                           `json:"DailyLimit,omitempty"`
		ReplicationBlackoutIntervals []StoragePureReplicationBlackout `json:"ReplicationBlackoutIntervals,omitempty"`
		// Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.
		SnapshotExpiryTime *string                                        `json:"SnapshotExpiryTime,omitempty"`
		Array              NullableStoragePureArrayRelationship           `json:"Array,omitempty"`
		ProtectionGroup    NullableStoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
		RegisteredDevice   NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureReplicationScheduleWithoutEmbeddedStruct := StoragePureReplicationScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePureReplicationScheduleWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureReplicationSchedule := _StoragePureReplicationSchedule{}
		varStoragePureReplicationSchedule.ClassId = varStoragePureReplicationScheduleWithoutEmbeddedStruct.ClassId
		varStoragePureReplicationSchedule.ObjectType = varStoragePureReplicationScheduleWithoutEmbeddedStruct.ObjectType
		varStoragePureReplicationSchedule.DailyLimit = varStoragePureReplicationScheduleWithoutEmbeddedStruct.DailyLimit
		varStoragePureReplicationSchedule.ReplicationBlackoutIntervals = varStoragePureReplicationScheduleWithoutEmbeddedStruct.ReplicationBlackoutIntervals
		varStoragePureReplicationSchedule.SnapshotExpiryTime = varStoragePureReplicationScheduleWithoutEmbeddedStruct.SnapshotExpiryTime
		varStoragePureReplicationSchedule.Array = varStoragePureReplicationScheduleWithoutEmbeddedStruct.Array
		varStoragePureReplicationSchedule.ProtectionGroup = varStoragePureReplicationScheduleWithoutEmbeddedStruct.ProtectionGroup
		varStoragePureReplicationSchedule.RegisteredDevice = varStoragePureReplicationScheduleWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureReplicationSchedule(varStoragePureReplicationSchedule)
	} else {
		return err
	}

	varStoragePureReplicationSchedule := _StoragePureReplicationSchedule{}

	err = json.Unmarshal(data, &varStoragePureReplicationSchedule)
	if err == nil {
		o.StorageBaseReplicationSchedule = varStoragePureReplicationSchedule.StorageBaseReplicationSchedule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DailyLimit")
		delete(additionalProperties, "ReplicationBlackoutIntervals")
		delete(additionalProperties, "SnapshotExpiryTime")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseReplicationSchedule := reflect.ValueOf(o.StorageBaseReplicationSchedule)
		for i := 0; i < reflectStorageBaseReplicationSchedule.Type().NumField(); i++ {
			t := reflectStorageBaseReplicationSchedule.Type().Field(i)

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

type NullableStoragePureReplicationSchedule struct {
	value *StoragePureReplicationSchedule
	isSet bool
}

func (v NullableStoragePureReplicationSchedule) Get() *StoragePureReplicationSchedule {
	return v.value
}

func (v *NullableStoragePureReplicationSchedule) Set(val *StoragePureReplicationSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureReplicationSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureReplicationSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureReplicationSchedule(val *StoragePureReplicationSchedule) *NullableStoragePureReplicationSchedule {
	return &NullableStoragePureReplicationSchedule{value: val, isSet: true}
}

func (v NullableStoragePureReplicationSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureReplicationSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
