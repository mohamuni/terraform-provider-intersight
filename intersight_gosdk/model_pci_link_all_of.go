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

// PciLinkAllOf Definition of the list of properties defined in 'pci.Link', excluding properties defined in parent classes.
type PciLinkAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the PCI device.
	Adapter *string `json:"Adapter,omitempty"`
	// The upstream link speed of the PCI device.
	LinkSpeed *string `json:"LinkSpeed,omitempty"`
	// The upstream link status of the PCI device.
	LinkStatus *string `json:"LinkStatus,omitempty"`
	// The upstream link width of the PCI device.
	LinkWidth *string `json:"LinkWidth,omitempty"`
	// The slot name of the PCI device.
	PciSlot *string `json:"PciSlot,omitempty"`
	// The health information of the PCI device.
	SlotStatus           *string                              `json:"SlotStatus,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PciSwitch            *PciSwitchRelationship               `json:"PciSwitch,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PciLinkAllOf PciLinkAllOf

// NewPciLinkAllOf instantiates a new PciLinkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciLinkAllOf(classId string, objectType string) *PciLinkAllOf {
	this := PciLinkAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPciLinkAllOfWithDefaults instantiates a new PciLinkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciLinkAllOfWithDefaults() *PciLinkAllOf {
	this := PciLinkAllOf{}
	var classId string = "pci.Link"
	this.ClassId = classId
	var objectType string = "pci.Link"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PciLinkAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PciLinkAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PciLinkAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PciLinkAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetAdapter() string {
	if o == nil || o.Adapter == nil {
		var ret string
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetAdapterOk() (*string, bool) {
	if o == nil || o.Adapter == nil {
		return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasAdapter() bool {
	if o != nil && o.Adapter != nil {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given string and assigns it to the Adapter field.
func (o *PciLinkAllOf) SetAdapter(v string) {
	o.Adapter = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetLinkSpeed() string {
	if o == nil || o.LinkSpeed == nil {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetLinkSpeedOk() (*string, bool) {
	if o == nil || o.LinkSpeed == nil {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasLinkSpeed() bool {
	if o != nil && o.LinkSpeed != nil {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *PciLinkAllOf) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetLinkStatus returns the LinkStatus field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetLinkStatus() string {
	if o == nil || o.LinkStatus == nil {
		var ret string
		return ret
	}
	return *o.LinkStatus
}

// GetLinkStatusOk returns a tuple with the LinkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetLinkStatusOk() (*string, bool) {
	if o == nil || o.LinkStatus == nil {
		return nil, false
	}
	return o.LinkStatus, true
}

// HasLinkStatus returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasLinkStatus() bool {
	if o != nil && o.LinkStatus != nil {
		return true
	}

	return false
}

// SetLinkStatus gets a reference to the given string and assigns it to the LinkStatus field.
func (o *PciLinkAllOf) SetLinkStatus(v string) {
	o.LinkStatus = &v
}

// GetLinkWidth returns the LinkWidth field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetLinkWidth() string {
	if o == nil || o.LinkWidth == nil {
		var ret string
		return ret
	}
	return *o.LinkWidth
}

// GetLinkWidthOk returns a tuple with the LinkWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetLinkWidthOk() (*string, bool) {
	if o == nil || o.LinkWidth == nil {
		return nil, false
	}
	return o.LinkWidth, true
}

// HasLinkWidth returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasLinkWidth() bool {
	if o != nil && o.LinkWidth != nil {
		return true
	}

	return false
}

// SetLinkWidth gets a reference to the given string and assigns it to the LinkWidth field.
func (o *PciLinkAllOf) SetLinkWidth(v string) {
	o.LinkWidth = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *PciLinkAllOf) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetSlotStatus returns the SlotStatus field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetSlotStatus() string {
	if o == nil || o.SlotStatus == nil {
		var ret string
		return ret
	}
	return *o.SlotStatus
}

// GetSlotStatusOk returns a tuple with the SlotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetSlotStatusOk() (*string, bool) {
	if o == nil || o.SlotStatus == nil {
		return nil, false
	}
	return o.SlotStatus, true
}

// HasSlotStatus returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasSlotStatus() bool {
	if o != nil && o.SlotStatus != nil {
		return true
	}

	return false
}

// SetSlotStatus gets a reference to the given string and assigns it to the SlotStatus field.
func (o *PciLinkAllOf) SetSlotStatus(v string) {
	o.SlotStatus = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PciLinkAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPciSwitch returns the PciSwitch field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetPciSwitch() PciSwitchRelationship {
	if o == nil || o.PciSwitch == nil {
		var ret PciSwitchRelationship
		return ret
	}
	return *o.PciSwitch
}

// GetPciSwitchOk returns a tuple with the PciSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetPciSwitchOk() (*PciSwitchRelationship, bool) {
	if o == nil || o.PciSwitch == nil {
		return nil, false
	}
	return o.PciSwitch, true
}

// HasPciSwitch returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasPciSwitch() bool {
	if o != nil && o.PciSwitch != nil {
		return true
	}

	return false
}

// SetPciSwitch gets a reference to the given PciSwitchRelationship and assigns it to the PciSwitch field.
func (o *PciLinkAllOf) SetPciSwitch(v PciSwitchRelationship) {
	o.PciSwitch = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PciLinkAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLinkAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PciLinkAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PciLinkAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PciLinkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Adapter != nil {
		toSerialize["Adapter"] = o.Adapter
	}
	if o.LinkSpeed != nil {
		toSerialize["LinkSpeed"] = o.LinkSpeed
	}
	if o.LinkStatus != nil {
		toSerialize["LinkStatus"] = o.LinkStatus
	}
	if o.LinkWidth != nil {
		toSerialize["LinkWidth"] = o.LinkWidth
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.SlotStatus != nil {
		toSerialize["SlotStatus"] = o.SlotStatus
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PciSwitch != nil {
		toSerialize["PciSwitch"] = o.PciSwitch
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PciLinkAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPciLinkAllOf := _PciLinkAllOf{}

	if err = json.Unmarshal(bytes, &varPciLinkAllOf); err == nil {
		*o = PciLinkAllOf(varPciLinkAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Adapter")
		delete(additionalProperties, "LinkSpeed")
		delete(additionalProperties, "LinkStatus")
		delete(additionalProperties, "LinkWidth")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "SlotStatus")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PciSwitch")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePciLinkAllOf struct {
	value *PciLinkAllOf
	isSet bool
}

func (v NullablePciLinkAllOf) Get() *PciLinkAllOf {
	return v.value
}

func (v *NullablePciLinkAllOf) Set(val *PciLinkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePciLinkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePciLinkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciLinkAllOf(val *PciLinkAllOf) *NullablePciLinkAllOf {
	return &NullablePciLinkAllOf{value: val, isSet: true}
}

func (v NullablePciLinkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciLinkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
