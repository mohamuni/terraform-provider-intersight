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

// checks if the StoragePureHostLun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePureHostLun{}

// StoragePureHostLun A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
type StoragePureHostLun struct {
	StorageBaseHostLun
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the host group associated with LUN.
	HostGroupName *string `json:"HostGroupName,omitempty"`
	// Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection.
	Shared               *bool                                       `json:"Shared,omitempty"`
	Array                NullableStoragePureArrayRelationship        `json:"Array,omitempty"`
	Host                 NullableStoragePureHostRelationship         `json:"Host,omitempty"`
	HostGroup            NullableStoragePureHostGroupRelationship    `json:"HostGroup,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Volume               NullableStoragePureVolumeRelationship       `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureHostLun StoragePureHostLun

// NewStoragePureHostLun instantiates a new StoragePureHostLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostLun(classId string, objectType string) *StoragePureHostLun {
	this := StoragePureHostLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureHostLunWithDefaults instantiates a new StoragePureHostLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostLunWithDefaults() *StoragePureHostLun {
	this := StoragePureHostLun{}
	var classId string = "storage.PureHostLun"
	this.ClassId = classId
	var objectType string = "storage.PureHostLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureHostLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureHostLun) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureHostLun" of the ClassId field.
func (o *StoragePureHostLun) GetDefaultClassId() interface{} {
	return "storage.PureHostLun"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureHostLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureHostLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureHostLun" of the ObjectType field.
func (o *StoragePureHostLun) GetDefaultObjectType() interface{} {
	return "storage.PureHostLun"
}

// GetHostGroupName returns the HostGroupName field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetHostGroupName() string {
	if o == nil || IsNil(o.HostGroupName) {
		var ret string
		return ret
	}
	return *o.HostGroupName
}

// GetHostGroupNameOk returns a tuple with the HostGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetHostGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostGroupName) {
		return nil, false
	}
	return o.HostGroupName, true
}

// HasHostGroupName returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasHostGroupName() bool {
	if o != nil && !IsNil(o.HostGroupName) {
		return true
	}

	return false
}

// SetHostGroupName gets a reference to the given string and assigns it to the HostGroupName field.
func (o *StoragePureHostLun) SetHostGroupName(v string) {
	o.HostGroupName = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *StoragePureHostLun) SetShared(v bool) {
	o.Shared = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostLun) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostLun) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHostLun) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePureHostLun) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePureHostLun) UnsetArray() {
	o.Array.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostLun) GetHost() StoragePureHostRelationship {
	if o == nil || IsNil(o.Host.Get()) {
		var ret StoragePureHostRelationship
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostLun) GetHostOk() (*StoragePureHostRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableStoragePureHostRelationship and assigns it to the Host field.
func (o *StoragePureHostLun) SetHost(v StoragePureHostRelationship) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *StoragePureHostLun) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *StoragePureHostLun) UnsetHost() {
	o.Host.Unset()
}

// GetHostGroup returns the HostGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostLun) GetHostGroup() StoragePureHostGroupRelationship {
	if o == nil || IsNil(o.HostGroup.Get()) {
		var ret StoragePureHostGroupRelationship
		return ret
	}
	return *o.HostGroup.Get()
}

// GetHostGroupOk returns a tuple with the HostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostLun) GetHostGroupOk() (*StoragePureHostGroupRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostGroup.Get(), o.HostGroup.IsSet()
}

// HasHostGroup returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasHostGroup() bool {
	if o != nil && o.HostGroup.IsSet() {
		return true
	}

	return false
}

// SetHostGroup gets a reference to the given NullableStoragePureHostGroupRelationship and assigns it to the HostGroup field.
func (o *StoragePureHostLun) SetHostGroup(v StoragePureHostGroupRelationship) {
	o.HostGroup.Set(&v)
}

// SetHostGroupNil sets the value for HostGroup to be an explicit nil
func (o *StoragePureHostLun) SetHostGroupNil() {
	o.HostGroup.Set(nil)
}

// UnsetHostGroup ensures that no value is present for HostGroup, not even an explicit nil
func (o *StoragePureHostLun) UnsetHostGroup() {
	o.HostGroup.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostLun) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostLun) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHostLun) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StoragePureHostLun) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StoragePureHostLun) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetVolume returns the Volume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostLun) GetVolume() StoragePureVolumeRelationship {
	if o == nil || IsNil(o.Volume.Get()) {
		var ret StoragePureVolumeRelationship
		return ret
	}
	return *o.Volume.Get()
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostLun) GetVolumeOk() (*StoragePureVolumeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volume.Get(), o.Volume.IsSet()
}

// HasVolume returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasVolume() bool {
	if o != nil && o.Volume.IsSet() {
		return true
	}

	return false
}

// SetVolume gets a reference to the given NullableStoragePureVolumeRelationship and assigns it to the Volume field.
func (o *StoragePureHostLun) SetVolume(v StoragePureVolumeRelationship) {
	o.Volume.Set(&v)
}

// SetVolumeNil sets the value for Volume to be an explicit nil
func (o *StoragePureHostLun) SetVolumeNil() {
	o.Volume.Set(nil)
}

// UnsetVolume ensures that no value is present for Volume, not even an explicit nil
func (o *StoragePureHostLun) UnsetVolume() {
	o.Volume.Unset()
}

func (o StoragePureHostLun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePureHostLun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostLun, errStorageBaseHostLun := json.Marshal(o.StorageBaseHostLun)
	if errStorageBaseHostLun != nil {
		return map[string]interface{}{}, errStorageBaseHostLun
	}
	errStorageBaseHostLun = json.Unmarshal([]byte(serializedStorageBaseHostLun), &toSerialize)
	if errStorageBaseHostLun != nil {
		return map[string]interface{}{}, errStorageBaseHostLun
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.HostGroupName) {
		toSerialize["HostGroupName"] = o.HostGroupName
	}
	if !IsNil(o.Shared) {
		toSerialize["Shared"] = o.Shared
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.Host.IsSet() {
		toSerialize["Host"] = o.Host.Get()
	}
	if o.HostGroup.IsSet() {
		toSerialize["HostGroup"] = o.HostGroup.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.Volume.IsSet() {
		toSerialize["Volume"] = o.Volume.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoragePureHostLun) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePureHostLunWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the host group associated with LUN.
		HostGroupName *string `json:"HostGroupName,omitempty"`
		// Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection.
		Shared           *bool                                       `json:"Shared,omitempty"`
		Array            NullableStoragePureArrayRelationship        `json:"Array,omitempty"`
		Host             NullableStoragePureHostRelationship         `json:"Host,omitempty"`
		HostGroup        NullableStoragePureHostGroupRelationship    `json:"HostGroup,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		Volume           NullableStoragePureVolumeRelationship       `json:"Volume,omitempty"`
	}

	varStoragePureHostLunWithoutEmbeddedStruct := StoragePureHostLunWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePureHostLunWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureHostLun := _StoragePureHostLun{}
		varStoragePureHostLun.ClassId = varStoragePureHostLunWithoutEmbeddedStruct.ClassId
		varStoragePureHostLun.ObjectType = varStoragePureHostLunWithoutEmbeddedStruct.ObjectType
		varStoragePureHostLun.HostGroupName = varStoragePureHostLunWithoutEmbeddedStruct.HostGroupName
		varStoragePureHostLun.Shared = varStoragePureHostLunWithoutEmbeddedStruct.Shared
		varStoragePureHostLun.Array = varStoragePureHostLunWithoutEmbeddedStruct.Array
		varStoragePureHostLun.Host = varStoragePureHostLunWithoutEmbeddedStruct.Host
		varStoragePureHostLun.HostGroup = varStoragePureHostLunWithoutEmbeddedStruct.HostGroup
		varStoragePureHostLun.RegisteredDevice = varStoragePureHostLunWithoutEmbeddedStruct.RegisteredDevice
		varStoragePureHostLun.Volume = varStoragePureHostLunWithoutEmbeddedStruct.Volume
		*o = StoragePureHostLun(varStoragePureHostLun)
	} else {
		return err
	}

	varStoragePureHostLun := _StoragePureHostLun{}

	err = json.Unmarshal(data, &varStoragePureHostLun)
	if err == nil {
		o.StorageBaseHostLun = varStoragePureHostLun.StorageBaseHostLun
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostGroupName")
		delete(additionalProperties, "Shared")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "HostGroup")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Volume")

		// remove fields from embedded structs
		reflectStorageBaseHostLun := reflect.ValueOf(o.StorageBaseHostLun)
		for i := 0; i < reflectStorageBaseHostLun.Type().NumField(); i++ {
			t := reflectStorageBaseHostLun.Type().Field(i)

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

type NullableStoragePureHostLun struct {
	value *StoragePureHostLun
	isSet bool
}

func (v NullableStoragePureHostLun) Get() *StoragePureHostLun {
	return v.value
}

func (v *NullableStoragePureHostLun) Set(val *StoragePureHostLun) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostLun) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostLun(val *StoragePureHostLun) *NullableStoragePureHostLun {
	return &NullableStoragePureHostLun{value: val, isSet: true}
}

func (v NullableStoragePureHostLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
