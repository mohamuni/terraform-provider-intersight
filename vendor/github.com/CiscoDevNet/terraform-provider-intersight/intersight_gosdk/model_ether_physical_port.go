/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EtherPhysicalPort Physical ethernet port present on a FI.
type EtherPhysicalPort struct {
	EtherPhysicalPortBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administratively configured speed for this port.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Administratively configured state (enabled/disabled) for this port.
	AdminState *string `json:"AdminState,omitempty"`
	// Breakout port member in the Fabric Interconnect.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// The number of days this port's license has been in Grace Period for.
	LicenseGrace *string `json:"LicenseGrace,omitempty"`
	// The state of the port's licensing.
	LicenseState         *string                              `json:"LicenseState,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PortGroup            *PortGroupRelationship               `json:"PortGroup,omitempty"`
	PortSubGroup         *PortSubGroupRelationship            `json:"PortSubGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EtherPhysicalPort EtherPhysicalPort

// NewEtherPhysicalPort instantiates a new EtherPhysicalPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherPhysicalPort(classId string, objectType string) *EtherPhysicalPort {
	this := EtherPhysicalPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherPhysicalPortWithDefaults instantiates a new EtherPhysicalPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherPhysicalPortWithDefaults() *EtherPhysicalPort {
	this := EtherPhysicalPort{}
	var classId string = "ether.PhysicalPort"
	this.ClassId = classId
	var objectType string = "ether.PhysicalPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherPhysicalPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherPhysicalPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EtherPhysicalPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherPhysicalPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *EtherPhysicalPort) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *EtherPhysicalPort) SetAdminState(v string) {
	o.AdminState = &v
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetAggregatePortId() int64 {
	if o == nil || o.AggregatePortId == nil {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || o.AggregatePortId == nil {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasAggregatePortId() bool {
	if o != nil && o.AggregatePortId != nil {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *EtherPhysicalPort) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetLicenseGrace returns the LicenseGrace field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetLicenseGrace() string {
	if o == nil || o.LicenseGrace == nil {
		var ret string
		return ret
	}
	return *o.LicenseGrace
}

// GetLicenseGraceOk returns a tuple with the LicenseGrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetLicenseGraceOk() (*string, bool) {
	if o == nil || o.LicenseGrace == nil {
		return nil, false
	}
	return o.LicenseGrace, true
}

// HasLicenseGrace returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasLicenseGrace() bool {
	if o != nil && o.LicenseGrace != nil {
		return true
	}

	return false
}

// SetLicenseGrace gets a reference to the given string and assigns it to the LicenseGrace field.
func (o *EtherPhysicalPort) SetLicenseGrace(v string) {
	o.LicenseGrace = &v
}

// GetLicenseState returns the LicenseState field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetLicenseState() string {
	if o == nil || o.LicenseState == nil {
		var ret string
		return ret
	}
	return *o.LicenseState
}

// GetLicenseStateOk returns a tuple with the LicenseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetLicenseStateOk() (*string, bool) {
	if o == nil || o.LicenseState == nil {
		return nil, false
	}
	return o.LicenseState, true
}

// HasLicenseState returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasLicenseState() bool {
	if o != nil && o.LicenseState != nil {
		return true
	}

	return false
}

// SetLicenseState gets a reference to the given string and assigns it to the LicenseState field.
func (o *EtherPhysicalPort) SetLicenseState(v string) {
	o.LicenseState = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EtherPhysicalPort) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPortGroup returns the PortGroup field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetPortGroup() PortGroupRelationship {
	if o == nil || o.PortGroup == nil {
		var ret PortGroupRelationship
		return ret
	}
	return *o.PortGroup
}

// GetPortGroupOk returns a tuple with the PortGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetPortGroupOk() (*PortGroupRelationship, bool) {
	if o == nil || o.PortGroup == nil {
		return nil, false
	}
	return o.PortGroup, true
}

// HasPortGroup returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasPortGroup() bool {
	if o != nil && o.PortGroup != nil {
		return true
	}

	return false
}

// SetPortGroup gets a reference to the given PortGroupRelationship and assigns it to the PortGroup field.
func (o *EtherPhysicalPort) SetPortGroup(v PortGroupRelationship) {
	o.PortGroup = &v
}

// GetPortSubGroup returns the PortSubGroup field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetPortSubGroup() PortSubGroupRelationship {
	if o == nil || o.PortSubGroup == nil {
		var ret PortSubGroupRelationship
		return ret
	}
	return *o.PortSubGroup
}

// GetPortSubGroupOk returns a tuple with the PortSubGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetPortSubGroupOk() (*PortSubGroupRelationship, bool) {
	if o == nil || o.PortSubGroup == nil {
		return nil, false
	}
	return o.PortSubGroup, true
}

// HasPortSubGroup returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasPortSubGroup() bool {
	if o != nil && o.PortSubGroup != nil {
		return true
	}

	return false
}

// SetPortSubGroup gets a reference to the given PortSubGroupRelationship and assigns it to the PortSubGroup field.
func (o *EtherPhysicalPort) SetPortSubGroup(v PortSubGroupRelationship) {
	o.PortSubGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EtherPhysicalPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EtherPhysicalPort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EtherPhysicalPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EtherPhysicalPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEtherPhysicalPortBase, errEtherPhysicalPortBase := json.Marshal(o.EtherPhysicalPortBase)
	if errEtherPhysicalPortBase != nil {
		return []byte{}, errEtherPhysicalPortBase
	}
	errEtherPhysicalPortBase = json.Unmarshal([]byte(serializedEtherPhysicalPortBase), &toSerialize)
	if errEtherPhysicalPortBase != nil {
		return []byte{}, errEtherPhysicalPortBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.AggregatePortId != nil {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if o.LicenseGrace != nil {
		toSerialize["LicenseGrace"] = o.LicenseGrace
	}
	if o.LicenseState != nil {
		toSerialize["LicenseState"] = o.LicenseState
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PortGroup != nil {
		toSerialize["PortGroup"] = o.PortGroup
	}
	if o.PortSubGroup != nil {
		toSerialize["PortSubGroup"] = o.PortSubGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EtherPhysicalPort) UnmarshalJSON(bytes []byte) (err error) {
	type EtherPhysicalPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Administratively configured speed for this port.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Administratively configured state (enabled/disabled) for this port.
		AdminState *string `json:"AdminState,omitempty"`
		// Breakout port member in the Fabric Interconnect.
		AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
		// The number of days this port's license has been in Grace Period for.
		LicenseGrace *string `json:"LicenseGrace,omitempty"`
		// The state of the port's licensing.
		LicenseState        *string                              `json:"LicenseState,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		PortGroup           *PortGroupRelationship               `json:"PortGroup,omitempty"`
		PortSubGroup        *PortSubGroupRelationship            `json:"PortSubGroup,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEtherPhysicalPortWithoutEmbeddedStruct := EtherPhysicalPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEtherPhysicalPortWithoutEmbeddedStruct)
	if err == nil {
		varEtherPhysicalPort := _EtherPhysicalPort{}
		varEtherPhysicalPort.ClassId = varEtherPhysicalPortWithoutEmbeddedStruct.ClassId
		varEtherPhysicalPort.ObjectType = varEtherPhysicalPortWithoutEmbeddedStruct.ObjectType
		varEtherPhysicalPort.AdminSpeed = varEtherPhysicalPortWithoutEmbeddedStruct.AdminSpeed
		varEtherPhysicalPort.AdminState = varEtherPhysicalPortWithoutEmbeddedStruct.AdminState
		varEtherPhysicalPort.AggregatePortId = varEtherPhysicalPortWithoutEmbeddedStruct.AggregatePortId
		varEtherPhysicalPort.LicenseGrace = varEtherPhysicalPortWithoutEmbeddedStruct.LicenseGrace
		varEtherPhysicalPort.LicenseState = varEtherPhysicalPortWithoutEmbeddedStruct.LicenseState
		varEtherPhysicalPort.InventoryDeviceInfo = varEtherPhysicalPortWithoutEmbeddedStruct.InventoryDeviceInfo
		varEtherPhysicalPort.PortGroup = varEtherPhysicalPortWithoutEmbeddedStruct.PortGroup
		varEtherPhysicalPort.PortSubGroup = varEtherPhysicalPortWithoutEmbeddedStruct.PortSubGroup
		varEtherPhysicalPort.RegisteredDevice = varEtherPhysicalPortWithoutEmbeddedStruct.RegisteredDevice
		*o = EtherPhysicalPort(varEtherPhysicalPort)
	} else {
		return err
	}

	varEtherPhysicalPort := _EtherPhysicalPort{}

	err = json.Unmarshal(bytes, &varEtherPhysicalPort)
	if err == nil {
		o.EtherPhysicalPortBase = varEtherPhysicalPort.EtherPhysicalPortBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "LicenseGrace")
		delete(additionalProperties, "LicenseState")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PortGroup")
		delete(additionalProperties, "PortSubGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEtherPhysicalPortBase := reflect.ValueOf(o.EtherPhysicalPortBase)
		for i := 0; i < reflectEtherPhysicalPortBase.Type().NumField(); i++ {
			t := reflectEtherPhysicalPortBase.Type().Field(i)

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

type NullableEtherPhysicalPort struct {
	value *EtherPhysicalPort
	isSet bool
}

func (v NullableEtherPhysicalPort) Get() *EtherPhysicalPort {
	return v.value
}

func (v *NullableEtherPhysicalPort) Set(val *EtherPhysicalPort) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherPhysicalPort) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherPhysicalPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherPhysicalPort(val *EtherPhysicalPort) *NullableEtherPhysicalPort {
	return &NullableEtherPhysicalPort{value: val, isSet: true}
}

func (v NullableEtherPhysicalPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherPhysicalPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}