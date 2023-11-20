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
)

// StorageEnclosureDiskSlotEpAllOf Definition of the list of properties defined in 'storage.EnclosureDiskSlotEp', excluding properties defined in parent classes.
type StorageEnclosureDiskSlotEpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the zoning configuration applied to  this enclosure slot.
	DrivePath *string `json:"DrivePath,omitempty"`
	// This field identifies the health of the disk inserted in the slot.
	Health *string `json:"Health,omitempty"`
	// This field identifies the disk is present in the enclosure slot.
	Presence *string `json:"Presence,omitempty"`
	// This field represents the slot Id in the storage enclosure.
	Slot                 *string                              `json:"Slot,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageEnclosure     *StorageEnclosureRelationship        `json:"StorageEnclosure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageEnclosureDiskSlotEpAllOf StorageEnclosureDiskSlotEpAllOf

// NewStorageEnclosureDiskSlotEpAllOf instantiates a new StorageEnclosureDiskSlotEpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageEnclosureDiskSlotEpAllOf(classId string, objectType string) *StorageEnclosureDiskSlotEpAllOf {
	this := StorageEnclosureDiskSlotEpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageEnclosureDiskSlotEpAllOfWithDefaults instantiates a new StorageEnclosureDiskSlotEpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageEnclosureDiskSlotEpAllOfWithDefaults() *StorageEnclosureDiskSlotEpAllOf {
	this := StorageEnclosureDiskSlotEpAllOf{}
	var classId string = "storage.EnclosureDiskSlotEp"
	this.ClassId = classId
	var objectType string = "storage.EnclosureDiskSlotEp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageEnclosureDiskSlotEpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageEnclosureDiskSlotEpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageEnclosureDiskSlotEpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageEnclosureDiskSlotEpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDrivePath returns the DrivePath field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetDrivePath() string {
	if o == nil || o.DrivePath == nil {
		var ret string
		return ret
	}
	return *o.DrivePath
}

// GetDrivePathOk returns a tuple with the DrivePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetDrivePathOk() (*string, bool) {
	if o == nil || o.DrivePath == nil {
		return nil, false
	}
	return o.DrivePath, true
}

// HasDrivePath returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasDrivePath() bool {
	if o != nil && o.DrivePath != nil {
		return true
	}

	return false
}

// SetDrivePath gets a reference to the given string and assigns it to the DrivePath field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetDrivePath(v string) {
	o.DrivePath = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetHealth(v string) {
	o.Health = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetSlot(v string) {
	o.Slot = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageEnclosure returns the StorageEnclosure field value if set, zero value otherwise.
func (o *StorageEnclosureDiskSlotEpAllOf) GetStorageEnclosure() StorageEnclosureRelationship {
	if o == nil || o.StorageEnclosure == nil {
		var ret StorageEnclosureRelationship
		return ret
	}
	return *o.StorageEnclosure
}

// GetStorageEnclosureOk returns a tuple with the StorageEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosure == nil {
		return nil, false
	}
	return o.StorageEnclosure, true
}

// HasStorageEnclosure returns a boolean if a field has been set.
func (o *StorageEnclosureDiskSlotEpAllOf) HasStorageEnclosure() bool {
	if o != nil && o.StorageEnclosure != nil {
		return true
	}

	return false
}

// SetStorageEnclosure gets a reference to the given StorageEnclosureRelationship and assigns it to the StorageEnclosure field.
func (o *StorageEnclosureDiskSlotEpAllOf) SetStorageEnclosure(v StorageEnclosureRelationship) {
	o.StorageEnclosure = &v
}

func (o StorageEnclosureDiskSlotEpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DrivePath != nil {
		toSerialize["DrivePath"] = o.DrivePath
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageEnclosure != nil {
		toSerialize["StorageEnclosure"] = o.StorageEnclosure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageEnclosureDiskSlotEpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageEnclosureDiskSlotEpAllOf := _StorageEnclosureDiskSlotEpAllOf{}

	if err = json.Unmarshal(bytes, &varStorageEnclosureDiskSlotEpAllOf); err == nil {
		*o = StorageEnclosureDiskSlotEpAllOf(varStorageEnclosureDiskSlotEpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DrivePath")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Slot")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageEnclosure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageEnclosureDiskSlotEpAllOf struct {
	value *StorageEnclosureDiskSlotEpAllOf
	isSet bool
}

func (v NullableStorageEnclosureDiskSlotEpAllOf) Get() *StorageEnclosureDiskSlotEpAllOf {
	return v.value
}

func (v *NullableStorageEnclosureDiskSlotEpAllOf) Set(val *StorageEnclosureDiskSlotEpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageEnclosureDiskSlotEpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageEnclosureDiskSlotEpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageEnclosureDiskSlotEpAllOf(val *StorageEnclosureDiskSlotEpAllOf) *NullableStorageEnclosureDiskSlotEpAllOf {
	return &NullableStorageEnclosureDiskSlotEpAllOf{value: val, isSet: true}
}

func (v NullableStorageEnclosureDiskSlotEpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageEnclosureDiskSlotEpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
