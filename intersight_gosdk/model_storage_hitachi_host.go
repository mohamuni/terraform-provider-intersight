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

// StorageHitachiHost A host group entity in Hitachi storage array. It is an abstraction used by Hitachi storage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
type StorageHitachiHost struct {
	StorageBaseHost
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authentication mode for the iSCSI target. * `N/A` - Authentication Mode is not available. * `CHAP` - CHAP-authentication mode. * `NONE` - Authentication mode is not set. * `BOTH` - Comply with Host Setting.
	AuthenticationMode *string `json:"AuthenticationMode,omitempty"`
	// ID of the host group for this host.
	HostGroupId *string `json:"HostGroupId,omitempty"`
	// Host group number for this host.
	HostGroupNumber *int64  `json:"HostGroupNumber,omitempty"`
	HostModeOptions []int64 `json:"HostModeOptions,omitempty"`
	// Mutual CHAP setting that is Enabled or Disabled.
	IsChapMutual *bool `json:"IsChapMutual,omitempty"`
	// IQN (iSCSI qualified name). Can be up to 255 characters long.
	IscsiName *string `json:"IscsiName,omitempty"`
	// Port ID of the host group.
	PortId *string `json:"PortId,omitempty"`
	// LUN security setting for the port.
	PortLunSecurity *bool `json:"PortLunSecurity,omitempty"`
	// Host Group type, it can be FC or iSCSI. * `FC` - Port supports fibre channel protocol. * `iSCSI` - Port supports iSCSI protocol. * `FCoE` - Port supports fibre channel over ethernet protocol.
	Type *string `json:"Type,omitempty"`
	// World wide port name, 64 bit address represented in hexa decimal notation.
	Wwn                  *string                              `json:"Wwn,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiHost StorageHitachiHost

// NewStorageHitachiHost instantiates a new StorageHitachiHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiHost(classId string, objectType string) *StorageHitachiHost {
	this := StorageHitachiHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiHostWithDefaults instantiates a new StorageHitachiHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiHostWithDefaults() *StorageHitachiHost {
	this := StorageHitachiHost{}
	var classId string = "storage.HitachiHost"
	this.ClassId = classId
	var objectType string = "storage.HitachiHost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthenticationMode returns the AuthenticationMode field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetAuthenticationMode() string {
	if o == nil || o.AuthenticationMode == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationMode
}

// GetAuthenticationModeOk returns a tuple with the AuthenticationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetAuthenticationModeOk() (*string, bool) {
	if o == nil || o.AuthenticationMode == nil {
		return nil, false
	}
	return o.AuthenticationMode, true
}

// HasAuthenticationMode returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasAuthenticationMode() bool {
	if o != nil && o.AuthenticationMode != nil {
		return true
	}

	return false
}

// SetAuthenticationMode gets a reference to the given string and assigns it to the AuthenticationMode field.
func (o *StorageHitachiHost) SetAuthenticationMode(v string) {
	o.AuthenticationMode = &v
}

// GetHostGroupId returns the HostGroupId field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetHostGroupId() string {
	if o == nil || o.HostGroupId == nil {
		var ret string
		return ret
	}
	return *o.HostGroupId
}

// GetHostGroupIdOk returns a tuple with the HostGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetHostGroupIdOk() (*string, bool) {
	if o == nil || o.HostGroupId == nil {
		return nil, false
	}
	return o.HostGroupId, true
}

// HasHostGroupId returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasHostGroupId() bool {
	if o != nil && o.HostGroupId != nil {
		return true
	}

	return false
}

// SetHostGroupId gets a reference to the given string and assigns it to the HostGroupId field.
func (o *StorageHitachiHost) SetHostGroupId(v string) {
	o.HostGroupId = &v
}

// GetHostGroupNumber returns the HostGroupNumber field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetHostGroupNumber() int64 {
	if o == nil || o.HostGroupNumber == nil {
		var ret int64
		return ret
	}
	return *o.HostGroupNumber
}

// GetHostGroupNumberOk returns a tuple with the HostGroupNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetHostGroupNumberOk() (*int64, bool) {
	if o == nil || o.HostGroupNumber == nil {
		return nil, false
	}
	return o.HostGroupNumber, true
}

// HasHostGroupNumber returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasHostGroupNumber() bool {
	if o != nil && o.HostGroupNumber != nil {
		return true
	}

	return false
}

// SetHostGroupNumber gets a reference to the given int64 and assigns it to the HostGroupNumber field.
func (o *StorageHitachiHost) SetHostGroupNumber(v int64) {
	o.HostGroupNumber = &v
}

// GetHostModeOptions returns the HostModeOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiHost) GetHostModeOptions() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.HostModeOptions
}

// GetHostModeOptionsOk returns a tuple with the HostModeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiHost) GetHostModeOptionsOk() ([]int64, bool) {
	if o == nil || o.HostModeOptions == nil {
		return nil, false
	}
	return o.HostModeOptions, true
}

// HasHostModeOptions returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasHostModeOptions() bool {
	if o != nil && o.HostModeOptions != nil {
		return true
	}

	return false
}

// SetHostModeOptions gets a reference to the given []int64 and assigns it to the HostModeOptions field.
func (o *StorageHitachiHost) SetHostModeOptions(v []int64) {
	o.HostModeOptions = v
}

// GetIsChapMutual returns the IsChapMutual field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetIsChapMutual() bool {
	if o == nil || o.IsChapMutual == nil {
		var ret bool
		return ret
	}
	return *o.IsChapMutual
}

// GetIsChapMutualOk returns a tuple with the IsChapMutual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetIsChapMutualOk() (*bool, bool) {
	if o == nil || o.IsChapMutual == nil {
		return nil, false
	}
	return o.IsChapMutual, true
}

// HasIsChapMutual returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasIsChapMutual() bool {
	if o != nil && o.IsChapMutual != nil {
		return true
	}

	return false
}

// SetIsChapMutual gets a reference to the given bool and assigns it to the IsChapMutual field.
func (o *StorageHitachiHost) SetIsChapMutual(v bool) {
	o.IsChapMutual = &v
}

// GetIscsiName returns the IscsiName field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetIscsiName() string {
	if o == nil || o.IscsiName == nil {
		var ret string
		return ret
	}
	return *o.IscsiName
}

// GetIscsiNameOk returns a tuple with the IscsiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetIscsiNameOk() (*string, bool) {
	if o == nil || o.IscsiName == nil {
		return nil, false
	}
	return o.IscsiName, true
}

// HasIscsiName returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasIscsiName() bool {
	if o != nil && o.IscsiName != nil {
		return true
	}

	return false
}

// SetIscsiName gets a reference to the given string and assigns it to the IscsiName field.
func (o *StorageHitachiHost) SetIscsiName(v string) {
	o.IscsiName = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetPortId() string {
	if o == nil || o.PortId == nil {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetPortIdOk() (*string, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageHitachiHost) SetPortId(v string) {
	o.PortId = &v
}

// GetPortLunSecurity returns the PortLunSecurity field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetPortLunSecurity() bool {
	if o == nil || o.PortLunSecurity == nil {
		var ret bool
		return ret
	}
	return *o.PortLunSecurity
}

// GetPortLunSecurityOk returns a tuple with the PortLunSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetPortLunSecurityOk() (*bool, bool) {
	if o == nil || o.PortLunSecurity == nil {
		return nil, false
	}
	return o.PortLunSecurity, true
}

// HasPortLunSecurity returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasPortLunSecurity() bool {
	if o != nil && o.PortLunSecurity != nil {
		return true
	}

	return false
}

// SetPortLunSecurity gets a reference to the given bool and assigns it to the PortLunSecurity field.
func (o *StorageHitachiHost) SetPortLunSecurity(v bool) {
	o.PortLunSecurity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageHitachiHost) SetType(v string) {
	o.Type = &v
}

// GetWwn returns the Wwn field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetWwn() string {
	if o == nil || o.Wwn == nil {
		var ret string
		return ret
	}
	return *o.Wwn
}

// GetWwnOk returns a tuple with the Wwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetWwnOk() (*string, bool) {
	if o == nil || o.Wwn == nil {
		return nil, false
	}
	return o.Wwn, true
}

// HasWwn returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasWwn() bool {
	if o != nil && o.Wwn != nil {
		return true
	}

	return false
}

// SetWwn gets a reference to the given string and assigns it to the Wwn field.
func (o *StorageHitachiHost) SetWwn(v string) {
	o.Wwn = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiHost) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiHost) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHost) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiHost) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiHost) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHost, errStorageBaseHost := json.Marshal(o.StorageBaseHost)
	if errStorageBaseHost != nil {
		return []byte{}, errStorageBaseHost
	}
	errStorageBaseHost = json.Unmarshal([]byte(serializedStorageBaseHost), &toSerialize)
	if errStorageBaseHost != nil {
		return []byte{}, errStorageBaseHost
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthenticationMode != nil {
		toSerialize["AuthenticationMode"] = o.AuthenticationMode
	}
	if o.HostGroupId != nil {
		toSerialize["HostGroupId"] = o.HostGroupId
	}
	if o.HostGroupNumber != nil {
		toSerialize["HostGroupNumber"] = o.HostGroupNumber
	}
	if o.HostModeOptions != nil {
		toSerialize["HostModeOptions"] = o.HostModeOptions
	}
	if o.IsChapMutual != nil {
		toSerialize["IsChapMutual"] = o.IsChapMutual
	}
	if o.IscsiName != nil {
		toSerialize["IscsiName"] = o.IscsiName
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.PortLunSecurity != nil {
		toSerialize["PortLunSecurity"] = o.PortLunSecurity
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Wwn != nil {
		toSerialize["Wwn"] = o.Wwn
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiHost) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Authentication mode for the iSCSI target. * `N/A` - Authentication Mode is not available. * `CHAP` - CHAP-authentication mode. * `NONE` - Authentication mode is not set. * `BOTH` - Comply with Host Setting.
		AuthenticationMode *string `json:"AuthenticationMode,omitempty"`
		// ID of the host group for this host.
		HostGroupId *string `json:"HostGroupId,omitempty"`
		// Host group number for this host.
		HostGroupNumber *int64  `json:"HostGroupNumber,omitempty"`
		HostModeOptions []int64 `json:"HostModeOptions,omitempty"`
		// Mutual CHAP setting that is Enabled or Disabled.
		IsChapMutual *bool `json:"IsChapMutual,omitempty"`
		// IQN (iSCSI qualified name). Can be up to 255 characters long.
		IscsiName *string `json:"IscsiName,omitempty"`
		// Port ID of the host group.
		PortId *string `json:"PortId,omitempty"`
		// LUN security setting for the port.
		PortLunSecurity *bool `json:"PortLunSecurity,omitempty"`
		// Host Group type, it can be FC or iSCSI. * `FC` - Port supports fibre channel protocol. * `iSCSI` - Port supports iSCSI protocol. * `FCoE` - Port supports fibre channel over ethernet protocol.
		Type *string `json:"Type,omitempty"`
		// World wide port name, 64 bit address represented in hexa decimal notation.
		Wwn              *string                              `json:"Wwn,omitempty"`
		Array            *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiHostWithoutEmbeddedStruct := StorageHitachiHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiHostWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiHost := _StorageHitachiHost{}
		varStorageHitachiHost.ClassId = varStorageHitachiHostWithoutEmbeddedStruct.ClassId
		varStorageHitachiHost.ObjectType = varStorageHitachiHostWithoutEmbeddedStruct.ObjectType
		varStorageHitachiHost.AuthenticationMode = varStorageHitachiHostWithoutEmbeddedStruct.AuthenticationMode
		varStorageHitachiHost.HostGroupId = varStorageHitachiHostWithoutEmbeddedStruct.HostGroupId
		varStorageHitachiHost.HostGroupNumber = varStorageHitachiHostWithoutEmbeddedStruct.HostGroupNumber
		varStorageHitachiHost.HostModeOptions = varStorageHitachiHostWithoutEmbeddedStruct.HostModeOptions
		varStorageHitachiHost.IsChapMutual = varStorageHitachiHostWithoutEmbeddedStruct.IsChapMutual
		varStorageHitachiHost.IscsiName = varStorageHitachiHostWithoutEmbeddedStruct.IscsiName
		varStorageHitachiHost.PortId = varStorageHitachiHostWithoutEmbeddedStruct.PortId
		varStorageHitachiHost.PortLunSecurity = varStorageHitachiHostWithoutEmbeddedStruct.PortLunSecurity
		varStorageHitachiHost.Type = varStorageHitachiHostWithoutEmbeddedStruct.Type
		varStorageHitachiHost.Wwn = varStorageHitachiHostWithoutEmbeddedStruct.Wwn
		varStorageHitachiHost.Array = varStorageHitachiHostWithoutEmbeddedStruct.Array
		varStorageHitachiHost.RegisteredDevice = varStorageHitachiHostWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiHost(varStorageHitachiHost)
	} else {
		return err
	}

	varStorageHitachiHost := _StorageHitachiHost{}

	err = json.Unmarshal(bytes, &varStorageHitachiHost)
	if err == nil {
		o.StorageBaseHost = varStorageHitachiHost.StorageBaseHost
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthenticationMode")
		delete(additionalProperties, "HostGroupId")
		delete(additionalProperties, "HostGroupNumber")
		delete(additionalProperties, "HostModeOptions")
		delete(additionalProperties, "IsChapMutual")
		delete(additionalProperties, "IscsiName")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "PortLunSecurity")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Wwn")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseHost := reflect.ValueOf(o.StorageBaseHost)
		for i := 0; i < reflectStorageBaseHost.Type().NumField(); i++ {
			t := reflectStorageBaseHost.Type().Field(i)

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

type NullableStorageHitachiHost struct {
	value *StorageHitachiHost
	isSet bool
}

func (v NullableStorageHitachiHost) Get() *StorageHitachiHost {
	return v.value
}

func (v *NullableStorageHitachiHost) Set(val *StorageHitachiHost) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiHost) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiHost(val *StorageHitachiHost) *NullableStorageHitachiHost {
	return &NullableStorageHitachiHost{value: val, isSet: true}
}

func (v NullableStorageHitachiHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
