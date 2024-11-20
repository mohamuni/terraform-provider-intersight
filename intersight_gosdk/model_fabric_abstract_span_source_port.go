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

// checks if the FabricAbstractSpanSourcePort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricAbstractSpanSourcePort{}

// FabricAbstractSpanSourcePort Abstract class for all SPAN Source Ports including Ethernet and FC.
type FabricAbstractSpanSourcePort struct {
	FabricAbstractSpanSource
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment, e.g. the id of the port on the switch.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Port Identifier of the Switch Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.
	PortId *int64 `json:"PortId,omitempty"`
	// Slot Identifier of the Switch Interface.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Role of the port configured as SPAN source. * `Uplink` - Uplink Role corresponding to PortRole in PortPolicy. * `FcoeUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `FcUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `Appliance` - FcoeUplink Role corresponding to PortRole in PortPolicy.
	SourceRole           *string                                `json:"SourceRole,omitempty"`
	PhysicalPort         NullableInventoryInterfaceRelationship `json:"PhysicalPort,omitempty"`
	SpanSession          NullableFabricSpanSessionRelationship  `json:"SpanSession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricAbstractSpanSourcePort FabricAbstractSpanSourcePort

// NewFabricAbstractSpanSourcePort instantiates a new FabricAbstractSpanSourcePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAbstractSpanSourcePort(classId string, objectType string) *FabricAbstractSpanSourcePort {
	this := FabricAbstractSpanSourcePort{}
	this.ClassId = classId
	this.ObjectType = objectType
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// NewFabricAbstractSpanSourcePortWithDefaults instantiates a new FabricAbstractSpanSourcePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAbstractSpanSourcePortWithDefaults() *FabricAbstractSpanSourcePort {
	this := FabricAbstractSpanSourcePort{}
	var classId string = "fabric.SpanSourceEthPort"
	this.ClassId = classId
	var objectType string = "fabric.SpanSourceEthPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAbstractSpanSourcePort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAbstractSpanSourcePort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SpanSourceEthPort" of the ClassId field.
func (o *FabricAbstractSpanSourcePort) GetDefaultClassId() interface{} {
	return "fabric.SpanSourceEthPort"
}

// GetObjectType returns the ObjectType field value
func (o *FabricAbstractSpanSourcePort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAbstractSpanSourcePort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SpanSourceEthPort" of the ObjectType field.
func (o *FabricAbstractSpanSourcePort) GetDefaultObjectType() interface{} {
	return "fabric.SpanSourceEthPort"
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourcePort) GetAggregatePortId() int64 {
	if o == nil || IsNil(o.AggregatePortId) {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePort) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AggregatePortId) {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePort) HasAggregatePortId() bool {
	if o != nil && !IsNil(o.AggregatePortId) {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *FabricAbstractSpanSourcePort) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourcePort) GetPortId() int64 {
	if o == nil || IsNil(o.PortId) {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePort) GetPortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePort) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *FabricAbstractSpanSourcePort) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourcePort) GetSlotId() int64 {
	if o == nil || IsNil(o.SlotId) {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePort) GetSlotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SlotId) {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePort) HasSlotId() bool {
	if o != nil && !IsNil(o.SlotId) {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *FabricAbstractSpanSourcePort) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSourceRole returns the SourceRole field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourcePort) GetSourceRole() string {
	if o == nil || IsNil(o.SourceRole) {
		var ret string
		return ret
	}
	return *o.SourceRole
}

// GetSourceRoleOk returns a tuple with the SourceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePort) GetSourceRoleOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRole) {
		return nil, false
	}
	return o.SourceRole, true
}

// HasSourceRole returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePort) HasSourceRole() bool {
	if o != nil && !IsNil(o.SourceRole) {
		return true
	}

	return false
}

// SetSourceRole gets a reference to the given string and assigns it to the SourceRole field.
func (o *FabricAbstractSpanSourcePort) SetSourceRole(v string) {
	o.SourceRole = &v
}

// GetPhysicalPort returns the PhysicalPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanSourcePort) GetPhysicalPort() InventoryInterfaceRelationship {
	if o == nil || IsNil(o.PhysicalPort.Get()) {
		var ret InventoryInterfaceRelationship
		return ret
	}
	return *o.PhysicalPort.Get()
}

// GetPhysicalPortOk returns a tuple with the PhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanSourcePort) GetPhysicalPortOk() (*InventoryInterfaceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhysicalPort.Get(), o.PhysicalPort.IsSet()
}

// HasPhysicalPort returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePort) HasPhysicalPort() bool {
	if o != nil && o.PhysicalPort.IsSet() {
		return true
	}

	return false
}

// SetPhysicalPort gets a reference to the given NullableInventoryInterfaceRelationship and assigns it to the PhysicalPort field.
func (o *FabricAbstractSpanSourcePort) SetPhysicalPort(v InventoryInterfaceRelationship) {
	o.PhysicalPort.Set(&v)
}

// SetPhysicalPortNil sets the value for PhysicalPort to be an explicit nil
func (o *FabricAbstractSpanSourcePort) SetPhysicalPortNil() {
	o.PhysicalPort.Set(nil)
}

// UnsetPhysicalPort ensures that no value is present for PhysicalPort, not even an explicit nil
func (o *FabricAbstractSpanSourcePort) UnsetPhysicalPort() {
	o.PhysicalPort.Unset()
}

// GetSpanSession returns the SpanSession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanSourcePort) GetSpanSession() FabricSpanSessionRelationship {
	if o == nil || IsNil(o.SpanSession.Get()) {
		var ret FabricSpanSessionRelationship
		return ret
	}
	return *o.SpanSession.Get()
}

// GetSpanSessionOk returns a tuple with the SpanSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanSourcePort) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpanSession.Get(), o.SpanSession.IsSet()
}

// HasSpanSession returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePort) HasSpanSession() bool {
	if o != nil && o.SpanSession.IsSet() {
		return true
	}

	return false
}

// SetSpanSession gets a reference to the given NullableFabricSpanSessionRelationship and assigns it to the SpanSession field.
func (o *FabricAbstractSpanSourcePort) SetSpanSession(v FabricSpanSessionRelationship) {
	o.SpanSession.Set(&v)
}

// SetSpanSessionNil sets the value for SpanSession to be an explicit nil
func (o *FabricAbstractSpanSourcePort) SetSpanSessionNil() {
	o.SpanSession.Set(nil)
}

// UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil
func (o *FabricAbstractSpanSourcePort) UnsetSpanSession() {
	o.SpanSession.Unset()
}

func (o FabricAbstractSpanSourcePort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricAbstractSpanSourcePort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractSpanSource, errFabricAbstractSpanSource := json.Marshal(o.FabricAbstractSpanSource)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	errFabricAbstractSpanSource = json.Unmarshal([]byte(serializedFabricAbstractSpanSource), &toSerialize)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AggregatePortId) {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if !IsNil(o.PortId) {
		toSerialize["PortId"] = o.PortId
	}
	if !IsNil(o.SlotId) {
		toSerialize["SlotId"] = o.SlotId
	}
	if !IsNil(o.SourceRole) {
		toSerialize["SourceRole"] = o.SourceRole
	}
	if o.PhysicalPort.IsSet() {
		toSerialize["PhysicalPort"] = o.PhysicalPort.Get()
	}
	if o.SpanSession.IsSet() {
		toSerialize["SpanSession"] = o.SpanSession.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricAbstractSpanSourcePort) UnmarshalJSON(data []byte) (err error) {
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
	type FabricAbstractSpanSourcePortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment, e.g. the id of the port on the switch.
		AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
		// Port Identifier of the Switch Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.
		PortId *int64 `json:"PortId,omitempty"`
		// Slot Identifier of the Switch Interface.
		SlotId *int64 `json:"SlotId,omitempty"`
		// Role of the port configured as SPAN source. * `Uplink` - Uplink Role corresponding to PortRole in PortPolicy. * `FcoeUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `FcUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `Appliance` - FcoeUplink Role corresponding to PortRole in PortPolicy.
		SourceRole   *string                                `json:"SourceRole,omitempty"`
		PhysicalPort NullableInventoryInterfaceRelationship `json:"PhysicalPort,omitempty"`
		SpanSession  NullableFabricSpanSessionRelationship  `json:"SpanSession,omitempty"`
	}

	varFabricAbstractSpanSourcePortWithoutEmbeddedStruct := FabricAbstractSpanSourcePortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourcePortWithoutEmbeddedStruct)
	if err == nil {
		varFabricAbstractSpanSourcePort := _FabricAbstractSpanSourcePort{}
		varFabricAbstractSpanSourcePort.ClassId = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.ClassId
		varFabricAbstractSpanSourcePort.ObjectType = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.ObjectType
		varFabricAbstractSpanSourcePort.AggregatePortId = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.AggregatePortId
		varFabricAbstractSpanSourcePort.PortId = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.PortId
		varFabricAbstractSpanSourcePort.SlotId = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.SlotId
		varFabricAbstractSpanSourcePort.SourceRole = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.SourceRole
		varFabricAbstractSpanSourcePort.PhysicalPort = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.PhysicalPort
		varFabricAbstractSpanSourcePort.SpanSession = varFabricAbstractSpanSourcePortWithoutEmbeddedStruct.SpanSession
		*o = FabricAbstractSpanSourcePort(varFabricAbstractSpanSourcePort)
	} else {
		return err
	}

	varFabricAbstractSpanSourcePort := _FabricAbstractSpanSourcePort{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourcePort)
	if err == nil {
		o.FabricAbstractSpanSource = varFabricAbstractSpanSourcePort.FabricAbstractSpanSource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "SourceRole")
		delete(additionalProperties, "PhysicalPort")
		delete(additionalProperties, "SpanSession")

		// remove fields from embedded structs
		reflectFabricAbstractSpanSource := reflect.ValueOf(o.FabricAbstractSpanSource)
		for i := 0; i < reflectFabricAbstractSpanSource.Type().NumField(); i++ {
			t := reflectFabricAbstractSpanSource.Type().Field(i)

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

type NullableFabricAbstractSpanSourcePort struct {
	value *FabricAbstractSpanSourcePort
	isSet bool
}

func (v NullableFabricAbstractSpanSourcePort) Get() *FabricAbstractSpanSourcePort {
	return v.value
}

func (v *NullableFabricAbstractSpanSourcePort) Set(val *FabricAbstractSpanSourcePort) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractSpanSourcePort) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractSpanSourcePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractSpanSourcePort(val *FabricAbstractSpanSourcePort) *NullableFabricAbstractSpanSourcePort {
	return &NullableFabricAbstractSpanSourcePort{value: val, isSet: true}
}

func (v NullableFabricAbstractSpanSourcePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractSpanSourcePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
