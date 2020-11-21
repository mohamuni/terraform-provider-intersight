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

// ProcessorUnit The CPU present on a server.
type ProcessorUnit struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The architecture of the installed processor.
	Architecture *string `json:"Architecture,omitempty"`
	// The number of cores present in a given processor.
	NumCores *int64 `json:"NumCores,omitempty"`
	// The number of enabled cores in the installed processor.
	NumCoresEnabled *string `json:"NumCoresEnabled,omitempty"`
	// The maximum number of threads available in the installed processor.
	NumThreads *string `json:"NumThreads,omitempty"`
	// The power state of the processor.
	OperPowerState *string `json:"OperPowerState,omitempty"`
	// The health indicator of the processor, 'OK' indicates the processor is operatinal.
	OperState *string `json:"OperState,omitempty"`
	// Operability state of the central processing unit.
	Operability *string `json:"Operability,omitempty"`
	// The valid values are 'equipped' and 'absent'.
	Presence *string `json:"Presence,omitempty"`
	// The ID number of a given processor.
	ProcessorId *int64 `json:"ProcessorId,omitempty"`
	// The socket ID of the installed processor.
	SocketDesignation *string `json:"SocketDesignation,omitempty"`
	// The maximum speed of the installed processor in GHz.
	Speed *float32 `json:"Speed,omitempty"`
	// The CPU stepping of the installed processor.
	Stepping *string `json:"Stepping,omitempty"`
	// The temperature of the processor in centigrade.
	Thermal              *string                              `json:"Thermal,omitempty"`
	ComputeBlade         *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeBoard         *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorUnit ProcessorUnit

// NewProcessorUnit instantiates a new ProcessorUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorUnit(classId string, objectType string) *ProcessorUnit {
	this := ProcessorUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewProcessorUnitWithDefaults instantiates a new ProcessorUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorUnitWithDefaults() *ProcessorUnit {
	this := ProcessorUnit{}
	var classId string = "processor.Unit"
	this.ClassId = classId
	var objectType string = "processor.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ProcessorUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ProcessorUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ProcessorUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ProcessorUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *ProcessorUnit) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *ProcessorUnit) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *ProcessorUnit) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetNumCores returns the NumCores field value if set, zero value otherwise.
func (o *ProcessorUnit) GetNumCores() int64 {
	if o == nil || o.NumCores == nil {
		var ret int64
		return ret
	}
	return *o.NumCores
}

// GetNumCoresOk returns a tuple with the NumCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetNumCoresOk() (*int64, bool) {
	if o == nil || o.NumCores == nil {
		return nil, false
	}
	return o.NumCores, true
}

// HasNumCores returns a boolean if a field has been set.
func (o *ProcessorUnit) HasNumCores() bool {
	if o != nil && o.NumCores != nil {
		return true
	}

	return false
}

// SetNumCores gets a reference to the given int64 and assigns it to the NumCores field.
func (o *ProcessorUnit) SetNumCores(v int64) {
	o.NumCores = &v
}

// GetNumCoresEnabled returns the NumCoresEnabled field value if set, zero value otherwise.
func (o *ProcessorUnit) GetNumCoresEnabled() string {
	if o == nil || o.NumCoresEnabled == nil {
		var ret string
		return ret
	}
	return *o.NumCoresEnabled
}

// GetNumCoresEnabledOk returns a tuple with the NumCoresEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetNumCoresEnabledOk() (*string, bool) {
	if o == nil || o.NumCoresEnabled == nil {
		return nil, false
	}
	return o.NumCoresEnabled, true
}

// HasNumCoresEnabled returns a boolean if a field has been set.
func (o *ProcessorUnit) HasNumCoresEnabled() bool {
	if o != nil && o.NumCoresEnabled != nil {
		return true
	}

	return false
}

// SetNumCoresEnabled gets a reference to the given string and assigns it to the NumCoresEnabled field.
func (o *ProcessorUnit) SetNumCoresEnabled(v string) {
	o.NumCoresEnabled = &v
}

// GetNumThreads returns the NumThreads field value if set, zero value otherwise.
func (o *ProcessorUnit) GetNumThreads() string {
	if o == nil || o.NumThreads == nil {
		var ret string
		return ret
	}
	return *o.NumThreads
}

// GetNumThreadsOk returns a tuple with the NumThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetNumThreadsOk() (*string, bool) {
	if o == nil || o.NumThreads == nil {
		return nil, false
	}
	return o.NumThreads, true
}

// HasNumThreads returns a boolean if a field has been set.
func (o *ProcessorUnit) HasNumThreads() bool {
	if o != nil && o.NumThreads != nil {
		return true
	}

	return false
}

// SetNumThreads gets a reference to the given string and assigns it to the NumThreads field.
func (o *ProcessorUnit) SetNumThreads(v string) {
	o.NumThreads = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *ProcessorUnit) GetOperPowerState() string {
	if o == nil || o.OperPowerState == nil {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetOperPowerStateOk() (*string, bool) {
	if o == nil || o.OperPowerState == nil {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *ProcessorUnit) HasOperPowerState() bool {
	if o != nil && o.OperPowerState != nil {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *ProcessorUnit) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *ProcessorUnit) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *ProcessorUnit) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *ProcessorUnit) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *ProcessorUnit) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *ProcessorUnit) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *ProcessorUnit) SetOperability(v string) {
	o.Operability = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *ProcessorUnit) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *ProcessorUnit) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *ProcessorUnit) SetPresence(v string) {
	o.Presence = &v
}

// GetProcessorId returns the ProcessorId field value if set, zero value otherwise.
func (o *ProcessorUnit) GetProcessorId() int64 {
	if o == nil || o.ProcessorId == nil {
		var ret int64
		return ret
	}
	return *o.ProcessorId
}

// GetProcessorIdOk returns a tuple with the ProcessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetProcessorIdOk() (*int64, bool) {
	if o == nil || o.ProcessorId == nil {
		return nil, false
	}
	return o.ProcessorId, true
}

// HasProcessorId returns a boolean if a field has been set.
func (o *ProcessorUnit) HasProcessorId() bool {
	if o != nil && o.ProcessorId != nil {
		return true
	}

	return false
}

// SetProcessorId gets a reference to the given int64 and assigns it to the ProcessorId field.
func (o *ProcessorUnit) SetProcessorId(v int64) {
	o.ProcessorId = &v
}

// GetSocketDesignation returns the SocketDesignation field value if set, zero value otherwise.
func (o *ProcessorUnit) GetSocketDesignation() string {
	if o == nil || o.SocketDesignation == nil {
		var ret string
		return ret
	}
	return *o.SocketDesignation
}

// GetSocketDesignationOk returns a tuple with the SocketDesignation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetSocketDesignationOk() (*string, bool) {
	if o == nil || o.SocketDesignation == nil {
		return nil, false
	}
	return o.SocketDesignation, true
}

// HasSocketDesignation returns a boolean if a field has been set.
func (o *ProcessorUnit) HasSocketDesignation() bool {
	if o != nil && o.SocketDesignation != nil {
		return true
	}

	return false
}

// SetSocketDesignation gets a reference to the given string and assigns it to the SocketDesignation field.
func (o *ProcessorUnit) SetSocketDesignation(v string) {
	o.SocketDesignation = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *ProcessorUnit) GetSpeed() float32 {
	if o == nil || o.Speed == nil {
		var ret float32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetSpeedOk() (*float32, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *ProcessorUnit) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given float32 and assigns it to the Speed field.
func (o *ProcessorUnit) SetSpeed(v float32) {
	o.Speed = &v
}

// GetStepping returns the Stepping field value if set, zero value otherwise.
func (o *ProcessorUnit) GetStepping() string {
	if o == nil || o.Stepping == nil {
		var ret string
		return ret
	}
	return *o.Stepping
}

// GetSteppingOk returns a tuple with the Stepping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetSteppingOk() (*string, bool) {
	if o == nil || o.Stepping == nil {
		return nil, false
	}
	return o.Stepping, true
}

// HasStepping returns a boolean if a field has been set.
func (o *ProcessorUnit) HasStepping() bool {
	if o != nil && o.Stepping != nil {
		return true
	}

	return false
}

// SetStepping gets a reference to the given string and assigns it to the Stepping field.
func (o *ProcessorUnit) SetStepping(v string) {
	o.Stepping = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *ProcessorUnit) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *ProcessorUnit) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *ProcessorUnit) SetThermal(v string) {
	o.Thermal = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *ProcessorUnit) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *ProcessorUnit) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *ProcessorUnit) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *ProcessorUnit) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *ProcessorUnit) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *ProcessorUnit) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *ProcessorUnit) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *ProcessorUnit) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *ProcessorUnit) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ProcessorUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ProcessorUnit) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ProcessorUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ProcessorUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ProcessorUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ProcessorUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ProcessorUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Architecture != nil {
		toSerialize["Architecture"] = o.Architecture
	}
	if o.NumCores != nil {
		toSerialize["NumCores"] = o.NumCores
	}
	if o.NumCoresEnabled != nil {
		toSerialize["NumCoresEnabled"] = o.NumCoresEnabled
	}
	if o.NumThreads != nil {
		toSerialize["NumThreads"] = o.NumThreads
	}
	if o.OperPowerState != nil {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.ProcessorId != nil {
		toSerialize["ProcessorId"] = o.ProcessorId
	}
	if o.SocketDesignation != nil {
		toSerialize["SocketDesignation"] = o.SocketDesignation
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Stepping != nil {
		toSerialize["Stepping"] = o.Stepping
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorUnit) UnmarshalJSON(bytes []byte) (err error) {
	type ProcessorUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The architecture of the installed processor.
		Architecture *string `json:"Architecture,omitempty"`
		// The number of cores present in a given processor.
		NumCores *int64 `json:"NumCores,omitempty"`
		// The number of enabled cores in the installed processor.
		NumCoresEnabled *string `json:"NumCoresEnabled,omitempty"`
		// The maximum number of threads available in the installed processor.
		NumThreads *string `json:"NumThreads,omitempty"`
		// The power state of the processor.
		OperPowerState *string `json:"OperPowerState,omitempty"`
		// The health indicator of the processor, 'OK' indicates the processor is operatinal.
		OperState *string `json:"OperState,omitempty"`
		// Operability state of the central processing unit.
		Operability *string `json:"Operability,omitempty"`
		// The valid values are 'equipped' and 'absent'.
		Presence *string `json:"Presence,omitempty"`
		// The ID number of a given processor.
		ProcessorId *int64 `json:"ProcessorId,omitempty"`
		// The socket ID of the installed processor.
		SocketDesignation *string `json:"SocketDesignation,omitempty"`
		// The maximum speed of the installed processor in GHz.
		Speed *float32 `json:"Speed,omitempty"`
		// The CPU stepping of the installed processor.
		Stepping *string `json:"Stepping,omitempty"`
		// The temperature of the processor in centigrade.
		Thermal             *string                              `json:"Thermal,omitempty"`
		ComputeBlade        *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeBoard        *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
		ComputeRackUnit     *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varProcessorUnitWithoutEmbeddedStruct := ProcessorUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varProcessorUnitWithoutEmbeddedStruct)
	if err == nil {
		varProcessorUnit := _ProcessorUnit{}
		varProcessorUnit.ClassId = varProcessorUnitWithoutEmbeddedStruct.ClassId
		varProcessorUnit.ObjectType = varProcessorUnitWithoutEmbeddedStruct.ObjectType
		varProcessorUnit.Architecture = varProcessorUnitWithoutEmbeddedStruct.Architecture
		varProcessorUnit.NumCores = varProcessorUnitWithoutEmbeddedStruct.NumCores
		varProcessorUnit.NumCoresEnabled = varProcessorUnitWithoutEmbeddedStruct.NumCoresEnabled
		varProcessorUnit.NumThreads = varProcessorUnitWithoutEmbeddedStruct.NumThreads
		varProcessorUnit.OperPowerState = varProcessorUnitWithoutEmbeddedStruct.OperPowerState
		varProcessorUnit.OperState = varProcessorUnitWithoutEmbeddedStruct.OperState
		varProcessorUnit.Operability = varProcessorUnitWithoutEmbeddedStruct.Operability
		varProcessorUnit.Presence = varProcessorUnitWithoutEmbeddedStruct.Presence
		varProcessorUnit.ProcessorId = varProcessorUnitWithoutEmbeddedStruct.ProcessorId
		varProcessorUnit.SocketDesignation = varProcessorUnitWithoutEmbeddedStruct.SocketDesignation
		varProcessorUnit.Speed = varProcessorUnitWithoutEmbeddedStruct.Speed
		varProcessorUnit.Stepping = varProcessorUnitWithoutEmbeddedStruct.Stepping
		varProcessorUnit.Thermal = varProcessorUnitWithoutEmbeddedStruct.Thermal
		varProcessorUnit.ComputeBlade = varProcessorUnitWithoutEmbeddedStruct.ComputeBlade
		varProcessorUnit.ComputeBoard = varProcessorUnitWithoutEmbeddedStruct.ComputeBoard
		varProcessorUnit.ComputeRackUnit = varProcessorUnitWithoutEmbeddedStruct.ComputeRackUnit
		varProcessorUnit.InventoryDeviceInfo = varProcessorUnitWithoutEmbeddedStruct.InventoryDeviceInfo
		varProcessorUnit.RegisteredDevice = varProcessorUnitWithoutEmbeddedStruct.RegisteredDevice
		*o = ProcessorUnit(varProcessorUnit)
	} else {
		return err
	}

	varProcessorUnit := _ProcessorUnit{}

	err = json.Unmarshal(bytes, &varProcessorUnit)
	if err == nil {
		o.EquipmentBase = varProcessorUnit.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Architecture")
		delete(additionalProperties, "NumCores")
		delete(additionalProperties, "NumCoresEnabled")
		delete(additionalProperties, "NumThreads")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "ProcessorId")
		delete(additionalProperties, "SocketDesignation")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Stepping")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableProcessorUnit struct {
	value *ProcessorUnit
	isSet bool
}

func (v NullableProcessorUnit) Get() *ProcessorUnit {
	return v.value
}

func (v *NullableProcessorUnit) Set(val *ProcessorUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorUnit(val *ProcessorUnit) *NullableProcessorUnit {
	return &NullableProcessorUnit{value: val, isSet: true}
}

func (v NullableProcessorUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}