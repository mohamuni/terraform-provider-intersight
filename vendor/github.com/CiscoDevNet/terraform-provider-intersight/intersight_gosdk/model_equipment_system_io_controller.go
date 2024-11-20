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

// checks if the EquipmentSystemIoController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentSystemIoController{}

// EquipmentSystemIoController I/O Controller on a chassis which provides the data path to S-series server.
type EquipmentSystemIoController struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The assigned identifier for a chassis.
	ChassisId *string `json:"ChassisId,omitempty"`
	// Connection Path identifies the data path available between IOModule and FI.
	ConnectionPath *string `json:"ConnectionPath,omitempty"`
	// Connection status identifies the status of data path.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// This field gives a brief information on systemIOController.
	Description *string `json:"Description,omitempty"`
	// This field identifies the CIMC that manages the controller.
	ManagingInstance *string `json:"ManagingInstance,omitempty"`
	// This field identifies the SIOC operational state.
	OperState *string `json:"OperState,omitempty"`
	// Part Number identifier for the IO module.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for systemIOController.
	Pid *string `json:"Pid,omitempty"`
	// This represents system I/O Controller identifier.
	SystemIoControllerId *int64                                      `json:"SystemIoControllerId,omitempty"`
	Cmc                  NullableManagementControllerRelationship    `json:"Cmc,omitempty"`
	EquipmentChassis     NullableEquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	SharedIoModule       NullableEquipmentSharedIoModuleRelationship `json:"SharedIoModule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSystemIoController EquipmentSystemIoController

// NewEquipmentSystemIoController instantiates a new EquipmentSystemIoController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSystemIoController(classId string, objectType string) *EquipmentSystemIoController {
	this := EquipmentSystemIoController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentSystemIoControllerWithDefaults instantiates a new EquipmentSystemIoController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSystemIoControllerWithDefaults() *EquipmentSystemIoController {
	this := EquipmentSystemIoController{}
	var classId string = "equipment.SystemIoController"
	this.ClassId = classId
	var objectType string = "equipment.SystemIoController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSystemIoController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSystemIoController) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.SystemIoController" of the ClassId field.
func (o *EquipmentSystemIoController) GetDefaultClassId() interface{} {
	return "equipment.SystemIoController"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSystemIoController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSystemIoController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.SystemIoController" of the ObjectType field.
func (o *EquipmentSystemIoController) GetDefaultObjectType() interface{} {
	return "equipment.SystemIoController"
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetChassisId() string {
	if o == nil || IsNil(o.ChassisId) {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetChassisIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *EquipmentSystemIoController) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetConnectionPath returns the ConnectionPath field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetConnectionPath() string {
	if o == nil || IsNil(o.ConnectionPath) {
		var ret string
		return ret
	}
	return *o.ConnectionPath
}

// GetConnectionPathOk returns a tuple with the ConnectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetConnectionPathOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionPath) {
		return nil, false
	}
	return o.ConnectionPath, true
}

// HasConnectionPath returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasConnectionPath() bool {
	if o != nil && !IsNil(o.ConnectionPath) {
		return true
	}

	return false
}

// SetConnectionPath gets a reference to the given string and assigns it to the ConnectionPath field.
func (o *EquipmentSystemIoController) SetConnectionPath(v string) {
	o.ConnectionPath = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetConnectionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentSystemIoController) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentSystemIoController) SetDescription(v string) {
	o.Description = &v
}

// GetManagingInstance returns the ManagingInstance field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetManagingInstance() string {
	if o == nil || IsNil(o.ManagingInstance) {
		var ret string
		return ret
	}
	return *o.ManagingInstance
}

// GetManagingInstanceOk returns a tuple with the ManagingInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetManagingInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ManagingInstance) {
		return nil, false
	}
	return o.ManagingInstance, true
}

// HasManagingInstance returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasManagingInstance() bool {
	if o != nil && !IsNil(o.ManagingInstance) {
		return true
	}

	return false
}

// SetManagingInstance gets a reference to the given string and assigns it to the ManagingInstance field.
func (o *EquipmentSystemIoController) SetManagingInstance(v string) {
	o.ManagingInstance = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentSystemIoController) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentSystemIoController) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentSystemIoController) SetPid(v string) {
	o.Pid = &v
}

// GetSystemIoControllerId returns the SystemIoControllerId field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetSystemIoControllerId() int64 {
	if o == nil || IsNil(o.SystemIoControllerId) {
		var ret int64
		return ret
	}
	return *o.SystemIoControllerId
}

// GetSystemIoControllerIdOk returns a tuple with the SystemIoControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetSystemIoControllerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SystemIoControllerId) {
		return nil, false
	}
	return o.SystemIoControllerId, true
}

// HasSystemIoControllerId returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasSystemIoControllerId() bool {
	if o != nil && !IsNil(o.SystemIoControllerId) {
		return true
	}

	return false
}

// SetSystemIoControllerId gets a reference to the given int64 and assigns it to the SystemIoControllerId field.
func (o *EquipmentSystemIoController) SetSystemIoControllerId(v int64) {
	o.SystemIoControllerId = &v
}

// GetCmc returns the Cmc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSystemIoController) GetCmc() ManagementControllerRelationship {
	if o == nil || IsNil(o.Cmc.Get()) {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Cmc.Get()
}

// GetCmcOk returns a tuple with the Cmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSystemIoController) GetCmcOk() (*ManagementControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cmc.Get(), o.Cmc.IsSet()
}

// HasCmc returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasCmc() bool {
	if o != nil && o.Cmc.IsSet() {
		return true
	}

	return false
}

// SetCmc gets a reference to the given NullableManagementControllerRelationship and assigns it to the Cmc field.
func (o *EquipmentSystemIoController) SetCmc(v ManagementControllerRelationship) {
	o.Cmc.Set(&v)
}

// SetCmcNil sets the value for Cmc to be an explicit nil
func (o *EquipmentSystemIoController) SetCmcNil() {
	o.Cmc.Set(nil)
}

// UnsetCmc ensures that no value is present for Cmc, not even an explicit nil
func (o *EquipmentSystemIoController) UnsetCmc() {
	o.Cmc.Unset()
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSystemIoController) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || IsNil(o.EquipmentChassis.Get()) {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis.Get()
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSystemIoController) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipmentChassis.Get(), o.EquipmentChassis.IsSet()
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis.IsSet() {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given NullableEquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentSystemIoController) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis.Set(&v)
}

// SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil
func (o *EquipmentSystemIoController) SetEquipmentChassisNil() {
	o.EquipmentChassis.Set(nil)
}

// UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
func (o *EquipmentSystemIoController) UnsetEquipmentChassis() {
	o.EquipmentChassis.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSystemIoController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSystemIoController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSystemIoController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *EquipmentSystemIoController) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *EquipmentSystemIoController) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSystemIoController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSystemIoController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSystemIoController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EquipmentSystemIoController) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EquipmentSystemIoController) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetSharedIoModule returns the SharedIoModule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSystemIoController) GetSharedIoModule() EquipmentSharedIoModuleRelationship {
	if o == nil || IsNil(o.SharedIoModule.Get()) {
		var ret EquipmentSharedIoModuleRelationship
		return ret
	}
	return *o.SharedIoModule.Get()
}

// GetSharedIoModuleOk returns a tuple with the SharedIoModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSystemIoController) GetSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedIoModule.Get(), o.SharedIoModule.IsSet()
}

// HasSharedIoModule returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasSharedIoModule() bool {
	if o != nil && o.SharedIoModule.IsSet() {
		return true
	}

	return false
}

// SetSharedIoModule gets a reference to the given NullableEquipmentSharedIoModuleRelationship and assigns it to the SharedIoModule field.
func (o *EquipmentSystemIoController) SetSharedIoModule(v EquipmentSharedIoModuleRelationship) {
	o.SharedIoModule.Set(&v)
}

// SetSharedIoModuleNil sets the value for SharedIoModule to be an explicit nil
func (o *EquipmentSystemIoController) SetSharedIoModuleNil() {
	o.SharedIoModule.Set(nil)
}

// UnsetSharedIoModule ensures that no value is present for SharedIoModule, not even an explicit nil
func (o *EquipmentSystemIoController) UnsetSharedIoModule() {
	o.SharedIoModule.Unset()
}

func (o EquipmentSystemIoController) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentSystemIoController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ChassisId) {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if !IsNil(o.ConnectionPath) {
		toSerialize["ConnectionPath"] = o.ConnectionPath
	}
	if !IsNil(o.ConnectionStatus) {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.ManagingInstance) {
		toSerialize["ManagingInstance"] = o.ManagingInstance
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.Pid) {
		toSerialize["Pid"] = o.Pid
	}
	if !IsNil(o.SystemIoControllerId) {
		toSerialize["SystemIoControllerId"] = o.SystemIoControllerId
	}
	if o.Cmc.IsSet() {
		toSerialize["Cmc"] = o.Cmc.Get()
	}
	if o.EquipmentChassis.IsSet() {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.SharedIoModule.IsSet() {
		toSerialize["SharedIoModule"] = o.SharedIoModule.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentSystemIoController) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentSystemIoControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The assigned identifier for a chassis.
		ChassisId *string `json:"ChassisId,omitempty"`
		// Connection Path identifies the data path available between IOModule and FI.
		ConnectionPath *string `json:"ConnectionPath,omitempty"`
		// Connection status identifies the status of data path.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// This field gives a brief information on systemIOController.
		Description *string `json:"Description,omitempty"`
		// This field identifies the CIMC that manages the controller.
		ManagingInstance *string `json:"ManagingInstance,omitempty"`
		// This field identifies the SIOC operational state.
		OperState *string `json:"OperState,omitempty"`
		// Part Number identifier for the IO module.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for systemIOController.
		Pid *string `json:"Pid,omitempty"`
		// This represents system I/O Controller identifier.
		SystemIoControllerId *int64                                      `json:"SystemIoControllerId,omitempty"`
		Cmc                  NullableManagementControllerRelationship    `json:"Cmc,omitempty"`
		EquipmentChassis     NullableEquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		SharedIoModule       NullableEquipmentSharedIoModuleRelationship `json:"SharedIoModule,omitempty"`
	}

	varEquipmentSystemIoControllerWithoutEmbeddedStruct := EquipmentSystemIoControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentSystemIoControllerWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentSystemIoController := _EquipmentSystemIoController{}
		varEquipmentSystemIoController.ClassId = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ClassId
		varEquipmentSystemIoController.ObjectType = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ObjectType
		varEquipmentSystemIoController.ChassisId = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ChassisId
		varEquipmentSystemIoController.ConnectionPath = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ConnectionPath
		varEquipmentSystemIoController.ConnectionStatus = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ConnectionStatus
		varEquipmentSystemIoController.Description = varEquipmentSystemIoControllerWithoutEmbeddedStruct.Description
		varEquipmentSystemIoController.ManagingInstance = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ManagingInstance
		varEquipmentSystemIoController.OperState = varEquipmentSystemIoControllerWithoutEmbeddedStruct.OperState
		varEquipmentSystemIoController.PartNumber = varEquipmentSystemIoControllerWithoutEmbeddedStruct.PartNumber
		varEquipmentSystemIoController.Pid = varEquipmentSystemIoControllerWithoutEmbeddedStruct.Pid
		varEquipmentSystemIoController.SystemIoControllerId = varEquipmentSystemIoControllerWithoutEmbeddedStruct.SystemIoControllerId
		varEquipmentSystemIoController.Cmc = varEquipmentSystemIoControllerWithoutEmbeddedStruct.Cmc
		varEquipmentSystemIoController.EquipmentChassis = varEquipmentSystemIoControllerWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentSystemIoController.InventoryDeviceInfo = varEquipmentSystemIoControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentSystemIoController.RegisteredDevice = varEquipmentSystemIoControllerWithoutEmbeddedStruct.RegisteredDevice
		varEquipmentSystemIoController.SharedIoModule = varEquipmentSystemIoControllerWithoutEmbeddedStruct.SharedIoModule
		*o = EquipmentSystemIoController(varEquipmentSystemIoController)
	} else {
		return err
	}

	varEquipmentSystemIoController := _EquipmentSystemIoController{}

	err = json.Unmarshal(data, &varEquipmentSystemIoController)
	if err == nil {
		o.EquipmentBase = varEquipmentSystemIoController.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ConnectionPath")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ManagingInstance")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SystemIoControllerId")
		delete(additionalProperties, "Cmc")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SharedIoModule")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentSystemIoController struct {
	value *EquipmentSystemIoController
	isSet bool
}

func (v NullableEquipmentSystemIoController) Get() *EquipmentSystemIoController {
	return v.value
}

func (v *NullableEquipmentSystemIoController) Set(val *EquipmentSystemIoController) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSystemIoController) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSystemIoController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSystemIoController(val *EquipmentSystemIoController) *NullableEquipmentSystemIoController {
	return &NullableEquipmentSystemIoController{value: val, isSet: true}
}

func (v NullableEquipmentSystemIoController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSystemIoController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
