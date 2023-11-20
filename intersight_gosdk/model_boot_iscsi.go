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

// BootIscsi Device type used when booting from iSCSI boot device.
type BootIscsi struct {
	BootDeviceBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                 `json:"ObjectType"`
	Bootloader NullableBootBootloader `json:"Bootloader,omitempty"`
	// The name of the underlying virtual ethernet interface used by the iSCSI boot device.
	InterfaceName *string `json:"InterfaceName,omitempty"`
	// Port ID of the ISCSI boot device.
	Port *int64 `json:"Port,omitempty"`
	// The slot id of the device. Supported values are (1 - 255, \"MLOM\", \"L\", \"L1\", \"L2\", \"OCP\").
	Slot                 *string `json:"Slot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootIscsi BootIscsi

// NewBootIscsi instantiates a new BootIscsi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootIscsi(classId string, objectType string) *BootIscsi {
	this := BootIscsi{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var port int64 = 0
	this.Port = &port
	return &this
}

// NewBootIscsiWithDefaults instantiates a new BootIscsi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootIscsiWithDefaults() *BootIscsi {
	this := BootIscsi{}
	var classId string = "boot.Iscsi"
	this.ClassId = classId
	var objectType string = "boot.Iscsi"
	this.ObjectType = objectType
	var port int64 = 0
	this.Port = &port
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootIscsi) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootIscsi) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootIscsi) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootIscsi) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootIscsi) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootIscsi) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootIscsi) GetBootloader() BootBootloader {
	if o == nil || o.Bootloader.Get() == nil {
		var ret BootBootloader
		return ret
	}
	return *o.Bootloader.Get()
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootIscsi) GetBootloaderOk() (*BootBootloader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bootloader.Get(), o.Bootloader.IsSet()
}

// HasBootloader returns a boolean if a field has been set.
func (o *BootIscsi) HasBootloader() bool {
	if o != nil && o.Bootloader.IsSet() {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given NullableBootBootloader and assigns it to the Bootloader field.
func (o *BootIscsi) SetBootloader(v BootBootloader) {
	o.Bootloader.Set(&v)
}

// SetBootloaderNil sets the value for Bootloader to be an explicit nil
func (o *BootIscsi) SetBootloaderNil() {
	o.Bootloader.Set(nil)
}

// UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil
func (o *BootIscsi) UnsetBootloader() {
	o.Bootloader.Unset()
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *BootIscsi) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootIscsi) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *BootIscsi) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *BootIscsi) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *BootIscsi) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootIscsi) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *BootIscsi) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *BootIscsi) SetPort(v int64) {
	o.Port = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *BootIscsi) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootIscsi) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *BootIscsi) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *BootIscsi) SetSlot(v string) {
	o.Slot = &v
}

func (o BootIscsi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBootDeviceBase, errBootDeviceBase := json.Marshal(o.BootDeviceBase)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	errBootDeviceBase = json.Unmarshal([]byte(serializedBootDeviceBase), &toSerialize)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootloader.IsSet() {
		toSerialize["Bootloader"] = o.Bootloader.Get()
	}
	if o.InterfaceName != nil {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootIscsi) UnmarshalJSON(bytes []byte) (err error) {
	type BootIscsiWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                 `json:"ObjectType"`
		Bootloader NullableBootBootloader `json:"Bootloader,omitempty"`
		// The name of the underlying virtual ethernet interface used by the iSCSI boot device.
		InterfaceName *string `json:"InterfaceName,omitempty"`
		// Port ID of the ISCSI boot device.
		Port *int64 `json:"Port,omitempty"`
		// The slot id of the device. Supported values are (1 - 255, \"MLOM\", \"L\", \"L1\", \"L2\", \"OCP\").
		Slot *string `json:"Slot,omitempty"`
	}

	varBootIscsiWithoutEmbeddedStruct := BootIscsiWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootIscsiWithoutEmbeddedStruct)
	if err == nil {
		varBootIscsi := _BootIscsi{}
		varBootIscsi.ClassId = varBootIscsiWithoutEmbeddedStruct.ClassId
		varBootIscsi.ObjectType = varBootIscsiWithoutEmbeddedStruct.ObjectType
		varBootIscsi.Bootloader = varBootIscsiWithoutEmbeddedStruct.Bootloader
		varBootIscsi.InterfaceName = varBootIscsiWithoutEmbeddedStruct.InterfaceName
		varBootIscsi.Port = varBootIscsiWithoutEmbeddedStruct.Port
		varBootIscsi.Slot = varBootIscsiWithoutEmbeddedStruct.Slot
		*o = BootIscsi(varBootIscsi)
	} else {
		return err
	}

	varBootIscsi := _BootIscsi{}

	err = json.Unmarshal(bytes, &varBootIscsi)
	if err == nil {
		o.BootDeviceBase = varBootIscsi.BootDeviceBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootloader")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Slot")

		// remove fields from embedded structs
		reflectBootDeviceBase := reflect.ValueOf(o.BootDeviceBase)
		for i := 0; i < reflectBootDeviceBase.Type().NumField(); i++ {
			t := reflectBootDeviceBase.Type().Field(i)

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

type NullableBootIscsi struct {
	value *BootIscsi
	isSet bool
}

func (v NullableBootIscsi) Get() *BootIscsi {
	return v.value
}

func (v *NullableBootIscsi) Set(val *BootIscsi) {
	v.value = val
	v.isSet = true
}

func (v NullableBootIscsi) IsSet() bool {
	return v.isSet
}

func (v *NullableBootIscsi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootIscsi(val *BootIscsi) *NullableBootIscsi {
	return &NullableBootIscsi{value: val, isSet: true}
}

func (v NullableBootIscsi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootIscsi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
