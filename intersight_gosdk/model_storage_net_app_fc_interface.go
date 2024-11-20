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

// checks if the StorageNetAppFcInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppFcInterface{}

// StorageNetAppFcInterface NetApp FC Interface is a logical interface.
type StorageNetAppFcInterface struct {
	StorageBasePhysicalPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// FC interface is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// The state of the FC interface. * `Down` - The state is set to down if the interface is not enabled. * `Up` - The state is set to up if the interface is enabled.
	InterfaceState *string `json:"InterfaceState,omitempty"`
	// The state of the FC interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	// Deprecated
	State *string `json:"State,omitempty"`
	// The storage virtual machine name for the interface.
	SvmName *string `json:"SvmName,omitempty"`
	// Uuid of NetApp FC Interface.
	Uuid *string `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	// The parent volume name for the interface.
	VolumeName      *string                               `json:"VolumeName,omitempty"`
	ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppFcInterfaceEvent resources.
	Events               []StorageNetAppFcInterfaceEventRelationship `json:"Events,omitempty"`
	PhysicalPort         NullableStorageNetAppFcPortRelationship     `json:"PhysicalPort,omitempty"`
	Tenant               NullableStorageNetAppStorageVmRelationship  `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppFcInterface StorageNetAppFcInterface

// NewStorageNetAppFcInterface instantiates a new StorageNetAppFcInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppFcInterface(classId string, objectType string) *StorageNetAppFcInterface {
	this := StorageNetAppFcInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppFcInterfaceWithDefaults instantiates a new StorageNetAppFcInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppFcInterfaceWithDefaults() *StorageNetAppFcInterface {
	this := StorageNetAppFcInterface{}
	var classId string = "storage.NetAppFcInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppFcInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppFcInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppFcInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppFcInterface" of the ClassId field.
func (o *StorageNetAppFcInterface) GetDefaultClassId() interface{} {
	return "storage.NetAppFcInterface"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppFcInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppFcInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppFcInterface" of the ObjectType field.
func (o *StorageNetAppFcInterface) GetDefaultObjectType() interface{} {
	return "storage.NetAppFcInterface"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetEnabled() string {
	if o == nil || IsNil(o.Enabled) {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppFcInterface) SetEnabled(v string) {
	o.Enabled = &v
}

// GetInterfaceState returns the InterfaceState field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetInterfaceState() string {
	if o == nil || IsNil(o.InterfaceState) {
		var ret string
		return ret
	}
	return *o.InterfaceState
}

// GetInterfaceStateOk returns a tuple with the InterfaceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetInterfaceStateOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceState) {
		return nil, false
	}
	return o.InterfaceState, true
}

// HasInterfaceState returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasInterfaceState() bool {
	if o != nil && !IsNil(o.InterfaceState) {
		return true
	}

	return false
}

// SetInterfaceState gets a reference to the given string and assigns it to the InterfaceState field.
func (o *StorageNetAppFcInterface) SetInterfaceState(v string) {
	o.InterfaceState = &v
}

// GetState returns the State field value if set, zero value otherwise.
// Deprecated
func (o *StorageNetAppFcInterface) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *StorageNetAppFcInterface) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
// Deprecated
func (o *StorageNetAppFcInterface) SetState(v string) {
	o.State = &v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetSvmName() string {
	if o == nil || IsNil(o.SvmName) {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetSvmNameOk() (*string, bool) {
	if o == nil || IsNil(o.SvmName) {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasSvmName() bool {
	if o != nil && !IsNil(o.SvmName) {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppFcInterface) SetSvmName(v string) {
	o.SvmName = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppFcInterface) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *StorageNetAppFcInterface) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcInterface) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || IsNil(o.ArrayController.Get()) {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController.Get()
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayController.Get(), o.ArrayController.IsSet()
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasArrayController() bool {
	if o != nil && o.ArrayController.IsSet() {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given NullableStorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppFcInterface) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController.Set(&v)
}

// SetArrayControllerNil sets the value for ArrayController to be an explicit nil
func (o *StorageNetAppFcInterface) SetArrayControllerNil() {
	o.ArrayController.Set(nil)
}

// UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
func (o *StorageNetAppFcInterface) UnsetArrayController() {
	o.ArrayController.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcInterface) GetEvents() []StorageNetAppFcInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppFcInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcInterface) GetEventsOk() ([]StorageNetAppFcInterfaceEventRelationship, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppFcInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppFcInterface) SetEvents(v []StorageNetAppFcInterfaceEventRelationship) {
	o.Events = v
}

// GetPhysicalPort returns the PhysicalPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcInterface) GetPhysicalPort() StorageNetAppFcPortRelationship {
	if o == nil || IsNil(o.PhysicalPort.Get()) {
		var ret StorageNetAppFcPortRelationship
		return ret
	}
	return *o.PhysicalPort.Get()
}

// GetPhysicalPortOk returns a tuple with the PhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcInterface) GetPhysicalPortOk() (*StorageNetAppFcPortRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhysicalPort.Get(), o.PhysicalPort.IsSet()
}

// HasPhysicalPort returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasPhysicalPort() bool {
	if o != nil && o.PhysicalPort.IsSet() {
		return true
	}

	return false
}

// SetPhysicalPort gets a reference to the given NullableStorageNetAppFcPortRelationship and assigns it to the PhysicalPort field.
func (o *StorageNetAppFcInterface) SetPhysicalPort(v StorageNetAppFcPortRelationship) {
	o.PhysicalPort.Set(&v)
}

// SetPhysicalPortNil sets the value for PhysicalPort to be an explicit nil
func (o *StorageNetAppFcInterface) SetPhysicalPortNil() {
	o.PhysicalPort.Set(nil)
}

// UnsetPhysicalPort ensures that no value is present for PhysicalPort, not even an explicit nil
func (o *StorageNetAppFcInterface) UnsetPhysicalPort() {
	o.PhysicalPort.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcInterface) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableStorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppFcInterface) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *StorageNetAppFcInterface) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *StorageNetAppFcInterface) UnsetTenant() {
	o.Tenant.Unset()
}

func (o StorageNetAppFcInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppFcInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBasePhysicalPort, errStorageBasePhysicalPort := json.Marshal(o.StorageBasePhysicalPort)
	if errStorageBasePhysicalPort != nil {
		return map[string]interface{}{}, errStorageBasePhysicalPort
	}
	errStorageBasePhysicalPort = json.Unmarshal([]byte(serializedStorageBasePhysicalPort), &toSerialize)
	if errStorageBasePhysicalPort != nil {
		return map[string]interface{}{}, errStorageBasePhysicalPort
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.InterfaceState) {
		toSerialize["InterfaceState"] = o.InterfaceState
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.SvmName) {
		toSerialize["SvmName"] = o.SvmName
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.VolumeName) {
		toSerialize["VolumeName"] = o.VolumeName
	}
	if o.ArrayController.IsSet() {
		toSerialize["ArrayController"] = o.ArrayController.Get()
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.PhysicalPort.IsSet() {
		toSerialize["PhysicalPort"] = o.PhysicalPort.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["Tenant"] = o.Tenant.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppFcInterface) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppFcInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// FC interface is enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// The state of the FC interface. * `Down` - The state is set to down if the interface is not enabled. * `Up` - The state is set to up if the interface is enabled.
		InterfaceState *string `json:"InterfaceState,omitempty"`
		// The state of the FC interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		// Deprecated
		State *string `json:"State,omitempty"`
		// The storage virtual machine name for the interface.
		SvmName *string `json:"SvmName,omitempty"`
		// Uuid of NetApp FC Interface.
		Uuid *string `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		// The parent volume name for the interface.
		VolumeName      *string                               `json:"VolumeName,omitempty"`
		ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppFcInterfaceEvent resources.
		Events       []StorageNetAppFcInterfaceEventRelationship `json:"Events,omitempty"`
		PhysicalPort NullableStorageNetAppFcPortRelationship     `json:"PhysicalPort,omitempty"`
		Tenant       NullableStorageNetAppStorageVmRelationship  `json:"Tenant,omitempty"`
	}

	varStorageNetAppFcInterfaceWithoutEmbeddedStruct := StorageNetAppFcInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppFcInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppFcInterface := _StorageNetAppFcInterface{}
		varStorageNetAppFcInterface.ClassId = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppFcInterface.ObjectType = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppFcInterface.Enabled = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Enabled
		varStorageNetAppFcInterface.InterfaceState = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.InterfaceState
		varStorageNetAppFcInterface.State = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.State
		varStorageNetAppFcInterface.SvmName = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.SvmName
		varStorageNetAppFcInterface.Uuid = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Uuid
		varStorageNetAppFcInterface.VolumeName = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.VolumeName
		varStorageNetAppFcInterface.ArrayController = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.ArrayController
		varStorageNetAppFcInterface.Events = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Events
		varStorageNetAppFcInterface.PhysicalPort = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.PhysicalPort
		varStorageNetAppFcInterface.Tenant = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppFcInterface(varStorageNetAppFcInterface)
	} else {
		return err
	}

	varStorageNetAppFcInterface := _StorageNetAppFcInterface{}

	err = json.Unmarshal(data, &varStorageNetAppFcInterface)
	if err == nil {
		o.StorageBasePhysicalPort = varStorageNetAppFcInterface.StorageBasePhysicalPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "InterfaceState")
		delete(additionalProperties, "State")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeName")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "PhysicalPort")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageBasePhysicalPort := reflect.ValueOf(o.StorageBasePhysicalPort)
		for i := 0; i < reflectStorageBasePhysicalPort.Type().NumField(); i++ {
			t := reflectStorageBasePhysicalPort.Type().Field(i)

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

type NullableStorageNetAppFcInterface struct {
	value *StorageNetAppFcInterface
	isSet bool
}

func (v NullableStorageNetAppFcInterface) Get() *StorageNetAppFcInterface {
	return v.value
}

func (v *NullableStorageNetAppFcInterface) Set(val *StorageNetAppFcInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppFcInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppFcInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppFcInterface(val *StorageNetAppFcInterface) *NullableStorageNetAppFcInterface {
	return &NullableStorageNetAppFcInterface{value: val, isSet: true}
}

func (v NullableStorageNetAppFcInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppFcInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
