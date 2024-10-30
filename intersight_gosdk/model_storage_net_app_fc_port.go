/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the StorageNetAppFcPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppFcPort{}

// StorageNetAppFcPort Fibre Channel (FC) port is a port on a node in a storage array.
type StorageNetAppFcPort struct {
	StorageBasePhysicalPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The node name for the port.
	NodeName *string `json:"NodeName,omitempty"`
	// Status of storage array port.
	PortStatus *string `json:"PortStatus,omitempty"`
	// The configured speed of the FC port in gigabits per second.
	SpeedConfigured *string `json:"SpeedConfigured,omitempty"`
	// The maximum speed supported by the FC port in gigabits per second.
	SpeedMaximum *string `json:"SpeedMaximum,omitempty"`
	// State of the port available in storage array. * `Unknown` - Default unknown port state. * `StartUp` - The port in the storage array is booting up. * `LinkNotConnected` - The port has finished initialization, but a link with the fabric is not established. * `Online` - The port is initialized and a link with the fabric has been established. * `LinkDisconnected` - The link on this port is currently not established. * `OfflineUser` - The port is administratively disabled. * `OfflineSystem` - The port is set to offline by the system. This happens when the port encounters too many errors. * `NodeOffline` - The state information for the port cannot be retrieved. The node is offline or inaccessible.
	State *string `json:"State,omitempty"`
	// Universally unique identifier of the FC port.
	Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppFcPortEvent resources.
	Events               []StorageNetAppFcPortEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppFcPort StorageNetAppFcPort

// NewStorageNetAppFcPort instantiates a new StorageNetAppFcPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppFcPort(classId string, objectType string) *StorageNetAppFcPort {
	this := StorageNetAppFcPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppFcPortWithDefaults instantiates a new StorageNetAppFcPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppFcPortWithDefaults() *StorageNetAppFcPort {
	this := StorageNetAppFcPort{}
	var classId string = "storage.NetAppFcPort"
	this.ClassId = classId
	var objectType string = "storage.NetAppFcPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppFcPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppFcPort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppFcPort" of the ClassId field.
func (o *StorageNetAppFcPort) GetDefaultClassId() interface{} {
	return "storage.NetAppFcPort"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppFcPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppFcPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppFcPort" of the ObjectType field.
func (o *StorageNetAppFcPort) GetDefaultObjectType() interface{} {
	return "storage.NetAppFcPort"
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *StorageNetAppFcPort) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *StorageNetAppFcPort) SetNodeName(v string) {
	o.NodeName = &v
}

// GetPortStatus returns the PortStatus field value if set, zero value otherwise.
func (o *StorageNetAppFcPort) GetPortStatus() string {
	if o == nil || IsNil(o.PortStatus) {
		var ret string
		return ret
	}
	return *o.PortStatus
}

// GetPortStatusOk returns a tuple with the PortStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetPortStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PortStatus) {
		return nil, false
	}
	return o.PortStatus, true
}

// HasPortStatus returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasPortStatus() bool {
	if o != nil && !IsNil(o.PortStatus) {
		return true
	}

	return false
}

// SetPortStatus gets a reference to the given string and assigns it to the PortStatus field.
func (o *StorageNetAppFcPort) SetPortStatus(v string) {
	o.PortStatus = &v
}

// GetSpeedConfigured returns the SpeedConfigured field value if set, zero value otherwise.
func (o *StorageNetAppFcPort) GetSpeedConfigured() string {
	if o == nil || IsNil(o.SpeedConfigured) {
		var ret string
		return ret
	}
	return *o.SpeedConfigured
}

// GetSpeedConfiguredOk returns a tuple with the SpeedConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetSpeedConfiguredOk() (*string, bool) {
	if o == nil || IsNil(o.SpeedConfigured) {
		return nil, false
	}
	return o.SpeedConfigured, true
}

// HasSpeedConfigured returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasSpeedConfigured() bool {
	if o != nil && !IsNil(o.SpeedConfigured) {
		return true
	}

	return false
}

// SetSpeedConfigured gets a reference to the given string and assigns it to the SpeedConfigured field.
func (o *StorageNetAppFcPort) SetSpeedConfigured(v string) {
	o.SpeedConfigured = &v
}

// GetSpeedMaximum returns the SpeedMaximum field value if set, zero value otherwise.
func (o *StorageNetAppFcPort) GetSpeedMaximum() string {
	if o == nil || IsNil(o.SpeedMaximum) {
		var ret string
		return ret
	}
	return *o.SpeedMaximum
}

// GetSpeedMaximumOk returns a tuple with the SpeedMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetSpeedMaximumOk() (*string, bool) {
	if o == nil || IsNil(o.SpeedMaximum) {
		return nil, false
	}
	return o.SpeedMaximum, true
}

// HasSpeedMaximum returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasSpeedMaximum() bool {
	if o != nil && !IsNil(o.SpeedMaximum) {
		return true
	}

	return false
}

// SetSpeedMaximum gets a reference to the given string and assigns it to the SpeedMaximum field.
func (o *StorageNetAppFcPort) SetSpeedMaximum(v string) {
	o.SpeedMaximum = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppFcPort) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppFcPort) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppFcPort) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPort) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppFcPort) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcPort) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || IsNil(o.ArrayController.Get()) {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController.Get()
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcPort) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayController.Get(), o.ArrayController.IsSet()
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasArrayController() bool {
	if o != nil && o.ArrayController.IsSet() {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given NullableStorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppFcPort) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController.Set(&v)
}

// SetArrayControllerNil sets the value for ArrayController to be an explicit nil
func (o *StorageNetAppFcPort) SetArrayControllerNil() {
	o.ArrayController.Set(nil)
}

// UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
func (o *StorageNetAppFcPort) UnsetArrayController() {
	o.ArrayController.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcPort) GetEvents() []StorageNetAppFcPortEventRelationship {
	if o == nil {
		var ret []StorageNetAppFcPortEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcPort) GetEventsOk() ([]StorageNetAppFcPortEventRelationship, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppFcPort) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppFcPortEventRelationship and assigns it to the Events field.
func (o *StorageNetAppFcPort) SetEvents(v []StorageNetAppFcPortEventRelationship) {
	o.Events = v
}

func (o StorageNetAppFcPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppFcPort) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NodeName) {
		toSerialize["NodeName"] = o.NodeName
	}
	if !IsNil(o.PortStatus) {
		toSerialize["PortStatus"] = o.PortStatus
	}
	if !IsNil(o.SpeedConfigured) {
		toSerialize["SpeedConfigured"] = o.SpeedConfigured
	}
	if !IsNil(o.SpeedMaximum) {
		toSerialize["SpeedMaximum"] = o.SpeedMaximum
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController.IsSet() {
		toSerialize["ArrayController"] = o.ArrayController.Get()
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppFcPort) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppFcPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The node name for the port.
		NodeName *string `json:"NodeName,omitempty"`
		// Status of storage array port.
		PortStatus *string `json:"PortStatus,omitempty"`
		// The configured speed of the FC port in gigabits per second.
		SpeedConfigured *string `json:"SpeedConfigured,omitempty"`
		// The maximum speed supported by the FC port in gigabits per second.
		SpeedMaximum *string `json:"SpeedMaximum,omitempty"`
		// State of the port available in storage array. * `Unknown` - Default unknown port state. * `StartUp` - The port in the storage array is booting up. * `LinkNotConnected` - The port has finished initialization, but a link with the fabric is not established. * `Online` - The port is initialized and a link with the fabric has been established. * `LinkDisconnected` - The link on this port is currently not established. * `OfflineUser` - The port is administratively disabled. * `OfflineSystem` - The port is set to offline by the system. This happens when the port encounters too many errors. * `NodeOffline` - The state information for the port cannot be retrieved. The node is offline or inaccessible.
		State *string `json:"State,omitempty"`
		// Universally unique identifier of the FC port.
		Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppFcPortEvent resources.
		Events []StorageNetAppFcPortEventRelationship `json:"Events,omitempty"`
	}

	varStorageNetAppFcPortWithoutEmbeddedStruct := StorageNetAppFcPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppFcPortWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppFcPort := _StorageNetAppFcPort{}
		varStorageNetAppFcPort.ClassId = varStorageNetAppFcPortWithoutEmbeddedStruct.ClassId
		varStorageNetAppFcPort.ObjectType = varStorageNetAppFcPortWithoutEmbeddedStruct.ObjectType
		varStorageNetAppFcPort.NodeName = varStorageNetAppFcPortWithoutEmbeddedStruct.NodeName
		varStorageNetAppFcPort.PortStatus = varStorageNetAppFcPortWithoutEmbeddedStruct.PortStatus
		varStorageNetAppFcPort.SpeedConfigured = varStorageNetAppFcPortWithoutEmbeddedStruct.SpeedConfigured
		varStorageNetAppFcPort.SpeedMaximum = varStorageNetAppFcPortWithoutEmbeddedStruct.SpeedMaximum
		varStorageNetAppFcPort.State = varStorageNetAppFcPortWithoutEmbeddedStruct.State
		varStorageNetAppFcPort.Uuid = varStorageNetAppFcPortWithoutEmbeddedStruct.Uuid
		varStorageNetAppFcPort.ArrayController = varStorageNetAppFcPortWithoutEmbeddedStruct.ArrayController
		varStorageNetAppFcPort.Events = varStorageNetAppFcPortWithoutEmbeddedStruct.Events
		*o = StorageNetAppFcPort(varStorageNetAppFcPort)
	} else {
		return err
	}

	varStorageNetAppFcPort := _StorageNetAppFcPort{}

	err = json.Unmarshal(data, &varStorageNetAppFcPort)
	if err == nil {
		o.StorageBasePhysicalPort = varStorageNetAppFcPort.StorageBasePhysicalPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NodeName")
		delete(additionalProperties, "PortStatus")
		delete(additionalProperties, "SpeedConfigured")
		delete(additionalProperties, "SpeedMaximum")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")

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

type NullableStorageNetAppFcPort struct {
	value *StorageNetAppFcPort
	isSet bool
}

func (v NullableStorageNetAppFcPort) Get() *StorageNetAppFcPort {
	return v.value
}

func (v *NullableStorageNetAppFcPort) Set(val *StorageNetAppFcPort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppFcPort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppFcPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppFcPort(val *StorageNetAppFcPort) *NullableStorageNetAppFcPort {
	return &NullableStorageNetAppFcPort{value: val, isSet: true}
}

func (v NullableStorageNetAppFcPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppFcPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
