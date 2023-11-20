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

// StorageHitachiExternalStoragePort Port entity in Hitachi storage array. A port for an external storage system that is connected to the local storage system.
type StorageHitachiExternalStoragePort struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Is Used of the external storage.
	ExternalIsUsed *bool `json:"ExternalIsUsed,omitempty"`
	// Path Mode of the external storage.
	ExternalPathMode *string `json:"ExternalPathMode,omitempty"`
	// Object ID of ports on an external storage system.
	ExternalPortId *string `json:"ExternalPortId,omitempty"`
	// Serial Number of the external storage.
	ExternalSerialNumber *string `json:"ExternalSerialNumber,omitempty"`
	// Storage Information of the external storage.
	ExternalStorageInfo *string `json:"ExternalStorageInfo,omitempty"`
	// WWN of the external storage port.
	ExternalWwn *string `json:"ExternalWwn,omitempty"`
	// The iSCSI IP Address of the external storage port.
	IscsiIpAddress *string `json:"IscsiIpAddress,omitempty"`
	// The iSCSI Name of the external storage port.
	IscsiName *string `json:"IscsiName,omitempty"`
	// Port ID of the local storage.
	PortId *string `json:"PortId,omitempty"`
	// Virtual port ID. This attribute is displayed when an iSCSI port is used and virtual port mode is enabled.
	VirtualPortId        *int64                               `json:"VirtualPortId,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiExternalStoragePort StorageHitachiExternalStoragePort

// NewStorageHitachiExternalStoragePort instantiates a new StorageHitachiExternalStoragePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiExternalStoragePort(classId string, objectType string) *StorageHitachiExternalStoragePort {
	this := StorageHitachiExternalStoragePort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiExternalStoragePortWithDefaults instantiates a new StorageHitachiExternalStoragePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiExternalStoragePortWithDefaults() *StorageHitachiExternalStoragePort {
	this := StorageHitachiExternalStoragePort{}
	var classId string = "storage.HitachiExternalStoragePort"
	this.ClassId = classId
	var objectType string = "storage.HitachiExternalStoragePort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiExternalStoragePort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiExternalStoragePort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiExternalStoragePort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiExternalStoragePort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalIsUsed returns the ExternalIsUsed field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetExternalIsUsed() bool {
	if o == nil || o.ExternalIsUsed == nil {
		var ret bool
		return ret
	}
	return *o.ExternalIsUsed
}

// GetExternalIsUsedOk returns a tuple with the ExternalIsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetExternalIsUsedOk() (*bool, bool) {
	if o == nil || o.ExternalIsUsed == nil {
		return nil, false
	}
	return o.ExternalIsUsed, true
}

// HasExternalIsUsed returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasExternalIsUsed() bool {
	if o != nil && o.ExternalIsUsed != nil {
		return true
	}

	return false
}

// SetExternalIsUsed gets a reference to the given bool and assigns it to the ExternalIsUsed field.
func (o *StorageHitachiExternalStoragePort) SetExternalIsUsed(v bool) {
	o.ExternalIsUsed = &v
}

// GetExternalPathMode returns the ExternalPathMode field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetExternalPathMode() string {
	if o == nil || o.ExternalPathMode == nil {
		var ret string
		return ret
	}
	return *o.ExternalPathMode
}

// GetExternalPathModeOk returns a tuple with the ExternalPathMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetExternalPathModeOk() (*string, bool) {
	if o == nil || o.ExternalPathMode == nil {
		return nil, false
	}
	return o.ExternalPathMode, true
}

// HasExternalPathMode returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasExternalPathMode() bool {
	if o != nil && o.ExternalPathMode != nil {
		return true
	}

	return false
}

// SetExternalPathMode gets a reference to the given string and assigns it to the ExternalPathMode field.
func (o *StorageHitachiExternalStoragePort) SetExternalPathMode(v string) {
	o.ExternalPathMode = &v
}

// GetExternalPortId returns the ExternalPortId field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetExternalPortId() string {
	if o == nil || o.ExternalPortId == nil {
		var ret string
		return ret
	}
	return *o.ExternalPortId
}

// GetExternalPortIdOk returns a tuple with the ExternalPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetExternalPortIdOk() (*string, bool) {
	if o == nil || o.ExternalPortId == nil {
		return nil, false
	}
	return o.ExternalPortId, true
}

// HasExternalPortId returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasExternalPortId() bool {
	if o != nil && o.ExternalPortId != nil {
		return true
	}

	return false
}

// SetExternalPortId gets a reference to the given string and assigns it to the ExternalPortId field.
func (o *StorageHitachiExternalStoragePort) SetExternalPortId(v string) {
	o.ExternalPortId = &v
}

// GetExternalSerialNumber returns the ExternalSerialNumber field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetExternalSerialNumber() string {
	if o == nil || o.ExternalSerialNumber == nil {
		var ret string
		return ret
	}
	return *o.ExternalSerialNumber
}

// GetExternalSerialNumberOk returns a tuple with the ExternalSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetExternalSerialNumberOk() (*string, bool) {
	if o == nil || o.ExternalSerialNumber == nil {
		return nil, false
	}
	return o.ExternalSerialNumber, true
}

// HasExternalSerialNumber returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasExternalSerialNumber() bool {
	if o != nil && o.ExternalSerialNumber != nil {
		return true
	}

	return false
}

// SetExternalSerialNumber gets a reference to the given string and assigns it to the ExternalSerialNumber field.
func (o *StorageHitachiExternalStoragePort) SetExternalSerialNumber(v string) {
	o.ExternalSerialNumber = &v
}

// GetExternalStorageInfo returns the ExternalStorageInfo field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetExternalStorageInfo() string {
	if o == nil || o.ExternalStorageInfo == nil {
		var ret string
		return ret
	}
	return *o.ExternalStorageInfo
}

// GetExternalStorageInfoOk returns a tuple with the ExternalStorageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetExternalStorageInfoOk() (*string, bool) {
	if o == nil || o.ExternalStorageInfo == nil {
		return nil, false
	}
	return o.ExternalStorageInfo, true
}

// HasExternalStorageInfo returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasExternalStorageInfo() bool {
	if o != nil && o.ExternalStorageInfo != nil {
		return true
	}

	return false
}

// SetExternalStorageInfo gets a reference to the given string and assigns it to the ExternalStorageInfo field.
func (o *StorageHitachiExternalStoragePort) SetExternalStorageInfo(v string) {
	o.ExternalStorageInfo = &v
}

// GetExternalWwn returns the ExternalWwn field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetExternalWwn() string {
	if o == nil || o.ExternalWwn == nil {
		var ret string
		return ret
	}
	return *o.ExternalWwn
}

// GetExternalWwnOk returns a tuple with the ExternalWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetExternalWwnOk() (*string, bool) {
	if o == nil || o.ExternalWwn == nil {
		return nil, false
	}
	return o.ExternalWwn, true
}

// HasExternalWwn returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasExternalWwn() bool {
	if o != nil && o.ExternalWwn != nil {
		return true
	}

	return false
}

// SetExternalWwn gets a reference to the given string and assigns it to the ExternalWwn field.
func (o *StorageHitachiExternalStoragePort) SetExternalWwn(v string) {
	o.ExternalWwn = &v
}

// GetIscsiIpAddress returns the IscsiIpAddress field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetIscsiIpAddress() string {
	if o == nil || o.IscsiIpAddress == nil {
		var ret string
		return ret
	}
	return *o.IscsiIpAddress
}

// GetIscsiIpAddressOk returns a tuple with the IscsiIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetIscsiIpAddressOk() (*string, bool) {
	if o == nil || o.IscsiIpAddress == nil {
		return nil, false
	}
	return o.IscsiIpAddress, true
}

// HasIscsiIpAddress returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasIscsiIpAddress() bool {
	if o != nil && o.IscsiIpAddress != nil {
		return true
	}

	return false
}

// SetIscsiIpAddress gets a reference to the given string and assigns it to the IscsiIpAddress field.
func (o *StorageHitachiExternalStoragePort) SetIscsiIpAddress(v string) {
	o.IscsiIpAddress = &v
}

// GetIscsiName returns the IscsiName field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetIscsiName() string {
	if o == nil || o.IscsiName == nil {
		var ret string
		return ret
	}
	return *o.IscsiName
}

// GetIscsiNameOk returns a tuple with the IscsiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetIscsiNameOk() (*string, bool) {
	if o == nil || o.IscsiName == nil {
		return nil, false
	}
	return o.IscsiName, true
}

// HasIscsiName returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasIscsiName() bool {
	if o != nil && o.IscsiName != nil {
		return true
	}

	return false
}

// SetIscsiName gets a reference to the given string and assigns it to the IscsiName field.
func (o *StorageHitachiExternalStoragePort) SetIscsiName(v string) {
	o.IscsiName = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetPortId() string {
	if o == nil || o.PortId == nil {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetPortIdOk() (*string, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageHitachiExternalStoragePort) SetPortId(v string) {
	o.PortId = &v
}

// GetVirtualPortId returns the VirtualPortId field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetVirtualPortId() int64 {
	if o == nil || o.VirtualPortId == nil {
		var ret int64
		return ret
	}
	return *o.VirtualPortId
}

// GetVirtualPortIdOk returns a tuple with the VirtualPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetVirtualPortIdOk() (*int64, bool) {
	if o == nil || o.VirtualPortId == nil {
		return nil, false
	}
	return o.VirtualPortId, true
}

// HasVirtualPortId returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasVirtualPortId() bool {
	if o != nil && o.VirtualPortId != nil {
		return true
	}

	return false
}

// SetVirtualPortId gets a reference to the given int64 and assigns it to the VirtualPortId field.
func (o *StorageHitachiExternalStoragePort) SetVirtualPortId(v int64) {
	o.VirtualPortId = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiExternalStoragePort) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiExternalStoragePort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStoragePort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiExternalStoragePort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiExternalStoragePort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiExternalStoragePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExternalIsUsed != nil {
		toSerialize["ExternalIsUsed"] = o.ExternalIsUsed
	}
	if o.ExternalPathMode != nil {
		toSerialize["ExternalPathMode"] = o.ExternalPathMode
	}
	if o.ExternalPortId != nil {
		toSerialize["ExternalPortId"] = o.ExternalPortId
	}
	if o.ExternalSerialNumber != nil {
		toSerialize["ExternalSerialNumber"] = o.ExternalSerialNumber
	}
	if o.ExternalStorageInfo != nil {
		toSerialize["ExternalStorageInfo"] = o.ExternalStorageInfo
	}
	if o.ExternalWwn != nil {
		toSerialize["ExternalWwn"] = o.ExternalWwn
	}
	if o.IscsiIpAddress != nil {
		toSerialize["IscsiIpAddress"] = o.IscsiIpAddress
	}
	if o.IscsiName != nil {
		toSerialize["IscsiName"] = o.IscsiName
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.VirtualPortId != nil {
		toSerialize["VirtualPortId"] = o.VirtualPortId
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

func (o *StorageHitachiExternalStoragePort) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiExternalStoragePortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Is Used of the external storage.
		ExternalIsUsed *bool `json:"ExternalIsUsed,omitempty"`
		// Path Mode of the external storage.
		ExternalPathMode *string `json:"ExternalPathMode,omitempty"`
		// Object ID of ports on an external storage system.
		ExternalPortId *string `json:"ExternalPortId,omitempty"`
		// Serial Number of the external storage.
		ExternalSerialNumber *string `json:"ExternalSerialNumber,omitempty"`
		// Storage Information of the external storage.
		ExternalStorageInfo *string `json:"ExternalStorageInfo,omitempty"`
		// WWN of the external storage port.
		ExternalWwn *string `json:"ExternalWwn,omitempty"`
		// The iSCSI IP Address of the external storage port.
		IscsiIpAddress *string `json:"IscsiIpAddress,omitempty"`
		// The iSCSI Name of the external storage port.
		IscsiName *string `json:"IscsiName,omitempty"`
		// Port ID of the local storage.
		PortId *string `json:"PortId,omitempty"`
		// Virtual port ID. This attribute is displayed when an iSCSI port is used and virtual port mode is enabled.
		VirtualPortId    *int64                               `json:"VirtualPortId,omitempty"`
		Array            *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiExternalStoragePortWithoutEmbeddedStruct := StorageHitachiExternalStoragePortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiExternalStoragePortWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiExternalStoragePort := _StorageHitachiExternalStoragePort{}
		varStorageHitachiExternalStoragePort.ClassId = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ClassId
		varStorageHitachiExternalStoragePort.ObjectType = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ObjectType
		varStorageHitachiExternalStoragePort.ExternalIsUsed = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ExternalIsUsed
		varStorageHitachiExternalStoragePort.ExternalPathMode = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ExternalPathMode
		varStorageHitachiExternalStoragePort.ExternalPortId = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ExternalPortId
		varStorageHitachiExternalStoragePort.ExternalSerialNumber = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ExternalSerialNumber
		varStorageHitachiExternalStoragePort.ExternalStorageInfo = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ExternalStorageInfo
		varStorageHitachiExternalStoragePort.ExternalWwn = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.ExternalWwn
		varStorageHitachiExternalStoragePort.IscsiIpAddress = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.IscsiIpAddress
		varStorageHitachiExternalStoragePort.IscsiName = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.IscsiName
		varStorageHitachiExternalStoragePort.PortId = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.PortId
		varStorageHitachiExternalStoragePort.VirtualPortId = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.VirtualPortId
		varStorageHitachiExternalStoragePort.Array = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.Array
		varStorageHitachiExternalStoragePort.RegisteredDevice = varStorageHitachiExternalStoragePortWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiExternalStoragePort(varStorageHitachiExternalStoragePort)
	} else {
		return err
	}

	varStorageHitachiExternalStoragePort := _StorageHitachiExternalStoragePort{}

	err = json.Unmarshal(bytes, &varStorageHitachiExternalStoragePort)
	if err == nil {
		o.MoBaseMo = varStorageHitachiExternalStoragePort.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalIsUsed")
		delete(additionalProperties, "ExternalPathMode")
		delete(additionalProperties, "ExternalPortId")
		delete(additionalProperties, "ExternalSerialNumber")
		delete(additionalProperties, "ExternalStorageInfo")
		delete(additionalProperties, "ExternalWwn")
		delete(additionalProperties, "IscsiIpAddress")
		delete(additionalProperties, "IscsiName")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "VirtualPortId")
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

type NullableStorageHitachiExternalStoragePort struct {
	value *StorageHitachiExternalStoragePort
	isSet bool
}

func (v NullableStorageHitachiExternalStoragePort) Get() *StorageHitachiExternalStoragePort {
	return v.value
}

func (v *NullableStorageHitachiExternalStoragePort) Set(val *StorageHitachiExternalStoragePort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiExternalStoragePort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiExternalStoragePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiExternalStoragePort(val *StorageHitachiExternalStoragePort) *NullableStorageHitachiExternalStoragePort {
	return &NullableStorageHitachiExternalStoragePort{value: val, isSet: true}
}

func (v NullableStorageHitachiExternalStoragePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiExternalStoragePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
