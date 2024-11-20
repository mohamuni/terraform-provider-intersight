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

// checks if the FabricAbstractSpanDestPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricAbstractSpanDestPort{}

// FabricAbstractSpanDestPort Abstract class for all Ethernet SPAN sources.
type FabricAbstractSpanDestPort struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment, e.g. the id of the port on the switch.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Port Identifier of the Switch Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.
	PortId *int64 `json:"PortId,omitempty"`
	// Slot Identifier of the Switch Interface.
	SlotId               *int64                                 `json:"SlotId,omitempty"`
	PhysicalPort         NullableInventoryInterfaceRelationship `json:"PhysicalPort,omitempty"`
	SpanSession          NullableFabricSpanSessionRelationship  `json:"SpanSession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricAbstractSpanDestPort FabricAbstractSpanDestPort

// NewFabricAbstractSpanDestPort instantiates a new FabricAbstractSpanDestPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAbstractSpanDestPort(classId string, objectType string) *FabricAbstractSpanDestPort {
	this := FabricAbstractSpanDestPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricAbstractSpanDestPortWithDefaults instantiates a new FabricAbstractSpanDestPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAbstractSpanDestPortWithDefaults() *FabricAbstractSpanDestPort {
	this := FabricAbstractSpanDestPort{}
	var classId string = "fabric.SpanDestEthPort"
	this.ClassId = classId
	var objectType string = "fabric.SpanDestEthPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAbstractSpanDestPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanDestPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAbstractSpanDestPort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SpanDestEthPort" of the ClassId field.
func (o *FabricAbstractSpanDestPort) GetDefaultClassId() interface{} {
	return "fabric.SpanDestEthPort"
}

// GetObjectType returns the ObjectType field value
func (o *FabricAbstractSpanDestPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanDestPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAbstractSpanDestPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SpanDestEthPort" of the ObjectType field.
func (o *FabricAbstractSpanDestPort) GetDefaultObjectType() interface{} {
	return "fabric.SpanDestEthPort"
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *FabricAbstractSpanDestPort) GetAggregatePortId() int64 {
	if o == nil || IsNil(o.AggregatePortId) {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanDestPort) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AggregatePortId) {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *FabricAbstractSpanDestPort) HasAggregatePortId() bool {
	if o != nil && !IsNil(o.AggregatePortId) {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *FabricAbstractSpanDestPort) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FabricAbstractSpanDestPort) GetPortId() int64 {
	if o == nil || IsNil(o.PortId) {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanDestPort) GetPortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FabricAbstractSpanDestPort) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *FabricAbstractSpanDestPort) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *FabricAbstractSpanDestPort) GetSlotId() int64 {
	if o == nil || IsNil(o.SlotId) {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanDestPort) GetSlotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SlotId) {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *FabricAbstractSpanDestPort) HasSlotId() bool {
	if o != nil && !IsNil(o.SlotId) {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *FabricAbstractSpanDestPort) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetPhysicalPort returns the PhysicalPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanDestPort) GetPhysicalPort() InventoryInterfaceRelationship {
	if o == nil || IsNil(o.PhysicalPort.Get()) {
		var ret InventoryInterfaceRelationship
		return ret
	}
	return *o.PhysicalPort.Get()
}

// GetPhysicalPortOk returns a tuple with the PhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanDestPort) GetPhysicalPortOk() (*InventoryInterfaceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhysicalPort.Get(), o.PhysicalPort.IsSet()
}

// HasPhysicalPort returns a boolean if a field has been set.
func (o *FabricAbstractSpanDestPort) HasPhysicalPort() bool {
	if o != nil && o.PhysicalPort.IsSet() {
		return true
	}

	return false
}

// SetPhysicalPort gets a reference to the given NullableInventoryInterfaceRelationship and assigns it to the PhysicalPort field.
func (o *FabricAbstractSpanDestPort) SetPhysicalPort(v InventoryInterfaceRelationship) {
	o.PhysicalPort.Set(&v)
}

// SetPhysicalPortNil sets the value for PhysicalPort to be an explicit nil
func (o *FabricAbstractSpanDestPort) SetPhysicalPortNil() {
	o.PhysicalPort.Set(nil)
}

// UnsetPhysicalPort ensures that no value is present for PhysicalPort, not even an explicit nil
func (o *FabricAbstractSpanDestPort) UnsetPhysicalPort() {
	o.PhysicalPort.Unset()
}

// GetSpanSession returns the SpanSession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanDestPort) GetSpanSession() FabricSpanSessionRelationship {
	if o == nil || IsNil(o.SpanSession.Get()) {
		var ret FabricSpanSessionRelationship
		return ret
	}
	return *o.SpanSession.Get()
}

// GetSpanSessionOk returns a tuple with the SpanSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanDestPort) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpanSession.Get(), o.SpanSession.IsSet()
}

// HasSpanSession returns a boolean if a field has been set.
func (o *FabricAbstractSpanDestPort) HasSpanSession() bool {
	if o != nil && o.SpanSession.IsSet() {
		return true
	}

	return false
}

// SetSpanSession gets a reference to the given NullableFabricSpanSessionRelationship and assigns it to the SpanSession field.
func (o *FabricAbstractSpanDestPort) SetSpanSession(v FabricSpanSessionRelationship) {
	o.SpanSession.Set(&v)
}

// SetSpanSessionNil sets the value for SpanSession to be an explicit nil
func (o *FabricAbstractSpanDestPort) SetSpanSessionNil() {
	o.SpanSession.Set(nil)
}

// UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil
func (o *FabricAbstractSpanDestPort) UnsetSpanSession() {
	o.SpanSession.Unset()
}

func (o FabricAbstractSpanDestPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricAbstractSpanDestPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
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

func (o *FabricAbstractSpanDestPort) UnmarshalJSON(data []byte) (err error) {
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
	type FabricAbstractSpanDestPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment, e.g. the id of the port on the switch.
		AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
		// Port Identifier of the Switch Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.
		PortId *int64 `json:"PortId,omitempty"`
		// Slot Identifier of the Switch Interface.
		SlotId       *int64                                 `json:"SlotId,omitempty"`
		PhysicalPort NullableInventoryInterfaceRelationship `json:"PhysicalPort,omitempty"`
		SpanSession  NullableFabricSpanSessionRelationship  `json:"SpanSession,omitempty"`
	}

	varFabricAbstractSpanDestPortWithoutEmbeddedStruct := FabricAbstractSpanDestPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricAbstractSpanDestPortWithoutEmbeddedStruct)
	if err == nil {
		varFabricAbstractSpanDestPort := _FabricAbstractSpanDestPort{}
		varFabricAbstractSpanDestPort.ClassId = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.ClassId
		varFabricAbstractSpanDestPort.ObjectType = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.ObjectType
		varFabricAbstractSpanDestPort.AggregatePortId = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.AggregatePortId
		varFabricAbstractSpanDestPort.PortId = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.PortId
		varFabricAbstractSpanDestPort.SlotId = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.SlotId
		varFabricAbstractSpanDestPort.PhysicalPort = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.PhysicalPort
		varFabricAbstractSpanDestPort.SpanSession = varFabricAbstractSpanDestPortWithoutEmbeddedStruct.SpanSession
		*o = FabricAbstractSpanDestPort(varFabricAbstractSpanDestPort)
	} else {
		return err
	}

	varFabricAbstractSpanDestPort := _FabricAbstractSpanDestPort{}

	err = json.Unmarshal(data, &varFabricAbstractSpanDestPort)
	if err == nil {
		o.MoBaseMo = varFabricAbstractSpanDestPort.MoBaseMo
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
		delete(additionalProperties, "PhysicalPort")
		delete(additionalProperties, "SpanSession")

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

type NullableFabricAbstractSpanDestPort struct {
	value *FabricAbstractSpanDestPort
	isSet bool
}

func (v NullableFabricAbstractSpanDestPort) Get() *FabricAbstractSpanDestPort {
	return v.value
}

func (v *NullableFabricAbstractSpanDestPort) Set(val *FabricAbstractSpanDestPort) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractSpanDestPort) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractSpanDestPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractSpanDestPort(val *FabricAbstractSpanDestPort) *NullableFabricAbstractSpanDestPort {
	return &NullableFabricAbstractSpanDestPort{value: val, isSet: true}
}

func (v NullableFabricAbstractSpanDestPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractSpanDestPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
