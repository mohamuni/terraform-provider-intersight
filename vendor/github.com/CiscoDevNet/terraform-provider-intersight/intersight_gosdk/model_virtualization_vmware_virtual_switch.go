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

// checks if the VirtualizationVmwareVirtualSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareVirtualSwitch{}

// VirtualizationVmwareVirtualSwitch The VMware Virtual Switch object is represented here.
type VirtualizationVmwareVirtualSwitch struct {
	VirtualizationBaseVirtualSwitch
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, the switch does not perform filtering and permits all outbound frames. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
	// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, the switch drops all inbound frames to the adapter. If property value is set to accept and the MAC address is changed, the switch allows frames to the new MAC address to pass. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	MacAddressChanges *string `json:"MacAddressChanges,omitempty"`
	// Maximum transmission unit configured on a virtual switch.
	Mtu                   *int64                                         `json:"Mtu,omitempty"`
	NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
	// Number of networks available on this virtual switch.
	NumNetworks *int64 `json:"NumNetworks,omitempty"`
	// Number of physical network interfaces connected with this virtual switch.
	NumPhysicalNetworkInterfaces *int64 `json:"NumPhysicalNetworkInterfaces,omitempty"`
	// If promiscuousMode property value is set to reject, the virtual switch forwards only frames that are addressed to the adapter. If property value is set to accept, the virtual switch forwards all frames to the adapter in compliance with the active VLAN policy for the port to which it is connected. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	PromiscuousMode      *string                                      `json:"PromiscuousMode,omitempty"`
	Host                 NullableVirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVirtualSwitch VirtualizationVmwareVirtualSwitch

// NewVirtualizationVmwareVirtualSwitch instantiates a new VirtualizationVmwareVirtualSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVirtualSwitch(classId string, objectType string) *VirtualizationVmwareVirtualSwitch {
	this := VirtualizationVmwareVirtualSwitch{}
	this.ClassId = classId
	this.ObjectType = objectType
	var forgedTransmits string = "Reject"
	this.ForgedTransmits = &forgedTransmits
	var macAddressChanges string = "Reject"
	this.MacAddressChanges = &macAddressChanges
	var promiscuousMode string = "Reject"
	this.PromiscuousMode = &promiscuousMode
	return &this
}

// NewVirtualizationVmwareVirtualSwitchWithDefaults instantiates a new VirtualizationVmwareVirtualSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVirtualSwitchWithDefaults() *VirtualizationVmwareVirtualSwitch {
	this := VirtualizationVmwareVirtualSwitch{}
	var classId string = "virtualization.VmwareVirtualSwitch"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualSwitch"
	this.ObjectType = objectType
	var forgedTransmits string = "Reject"
	this.ForgedTransmits = &forgedTransmits
	var macAddressChanges string = "Reject"
	this.MacAddressChanges = &macAddressChanges
	var promiscuousMode string = "Reject"
	this.PromiscuousMode = &promiscuousMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVirtualSwitch) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVirtualSwitch) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareVirtualSwitch" of the ClassId field.
func (o *VirtualizationVmwareVirtualSwitch) GetDefaultClassId() interface{} {
	return "virtualization.VmwareVirtualSwitch"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVirtualSwitch) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVirtualSwitch) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareVirtualSwitch" of the ObjectType field.
func (o *VirtualizationVmwareVirtualSwitch) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareVirtualSwitch"
}

// GetForgedTransmits returns the ForgedTransmits field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitch) GetForgedTransmits() string {
	if o == nil || IsNil(o.ForgedTransmits) {
		var ret string
		return ret
	}
	return *o.ForgedTransmits
}

// GetForgedTransmitsOk returns a tuple with the ForgedTransmits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetForgedTransmitsOk() (*string, bool) {
	if o == nil || IsNil(o.ForgedTransmits) {
		return nil, false
	}
	return o.ForgedTransmits, true
}

// HasForgedTransmits returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasForgedTransmits() bool {
	if o != nil && !IsNil(o.ForgedTransmits) {
		return true
	}

	return false
}

// SetForgedTransmits gets a reference to the given string and assigns it to the ForgedTransmits field.
func (o *VirtualizationVmwareVirtualSwitch) SetForgedTransmits(v string) {
	o.ForgedTransmits = &v
}

// GetMacAddressChanges returns the MacAddressChanges field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitch) GetMacAddressChanges() string {
	if o == nil || IsNil(o.MacAddressChanges) {
		var ret string
		return ret
	}
	return *o.MacAddressChanges
}

// GetMacAddressChangesOk returns a tuple with the MacAddressChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetMacAddressChangesOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddressChanges) {
		return nil, false
	}
	return o.MacAddressChanges, true
}

// HasMacAddressChanges returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasMacAddressChanges() bool {
	if o != nil && !IsNil(o.MacAddressChanges) {
		return true
	}

	return false
}

// SetMacAddressChanges gets a reference to the given string and assigns it to the MacAddressChanges field.
func (o *VirtualizationVmwareVirtualSwitch) SetMacAddressChanges(v string) {
	o.MacAddressChanges = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitch) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationVmwareVirtualSwitch) SetMtu(v int64) {
	o.Mtu = &v
}

// GetNicTeamingAndFailover returns the NicTeamingAndFailover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareVirtualSwitch) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover {
	if o == nil || IsNil(o.NicTeamingAndFailover.Get()) {
		var ret VirtualizationVmwareTeamingAndFailover
		return ret
	}
	return *o.NicTeamingAndFailover.Get()
}

// GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareVirtualSwitch) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicTeamingAndFailover.Get(), o.NicTeamingAndFailover.IsSet()
}

// HasNicTeamingAndFailover returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasNicTeamingAndFailover() bool {
	if o != nil && o.NicTeamingAndFailover.IsSet() {
		return true
	}

	return false
}

// SetNicTeamingAndFailover gets a reference to the given NullableVirtualizationVmwareTeamingAndFailover and assigns it to the NicTeamingAndFailover field.
func (o *VirtualizationVmwareVirtualSwitch) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover) {
	o.NicTeamingAndFailover.Set(&v)
}

// SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil
func (o *VirtualizationVmwareVirtualSwitch) SetNicTeamingAndFailoverNil() {
	o.NicTeamingAndFailover.Set(nil)
}

// UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
func (o *VirtualizationVmwareVirtualSwitch) UnsetNicTeamingAndFailover() {
	o.NicTeamingAndFailover.Unset()
}

// GetNumNetworks returns the NumNetworks field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitch) GetNumNetworks() int64 {
	if o == nil || IsNil(o.NumNetworks) {
		var ret int64
		return ret
	}
	return *o.NumNetworks
}

// GetNumNetworksOk returns a tuple with the NumNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetNumNetworksOk() (*int64, bool) {
	if o == nil || IsNil(o.NumNetworks) {
		return nil, false
	}
	return o.NumNetworks, true
}

// HasNumNetworks returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasNumNetworks() bool {
	if o != nil && !IsNil(o.NumNetworks) {
		return true
	}

	return false
}

// SetNumNetworks gets a reference to the given int64 and assigns it to the NumNetworks field.
func (o *VirtualizationVmwareVirtualSwitch) SetNumNetworks(v int64) {
	o.NumNetworks = &v
}

// GetNumPhysicalNetworkInterfaces returns the NumPhysicalNetworkInterfaces field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitch) GetNumPhysicalNetworkInterfaces() int64 {
	if o == nil || IsNil(o.NumPhysicalNetworkInterfaces) {
		var ret int64
		return ret
	}
	return *o.NumPhysicalNetworkInterfaces
}

// GetNumPhysicalNetworkInterfacesOk returns a tuple with the NumPhysicalNetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetNumPhysicalNetworkInterfacesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumPhysicalNetworkInterfaces) {
		return nil, false
	}
	return o.NumPhysicalNetworkInterfaces, true
}

// HasNumPhysicalNetworkInterfaces returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasNumPhysicalNetworkInterfaces() bool {
	if o != nil && !IsNil(o.NumPhysicalNetworkInterfaces) {
		return true
	}

	return false
}

// SetNumPhysicalNetworkInterfaces gets a reference to the given int64 and assigns it to the NumPhysicalNetworkInterfaces field.
func (o *VirtualizationVmwareVirtualSwitch) SetNumPhysicalNetworkInterfaces(v int64) {
	o.NumPhysicalNetworkInterfaces = &v
}

// GetPromiscuousMode returns the PromiscuousMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitch) GetPromiscuousMode() string {
	if o == nil || IsNil(o.PromiscuousMode) {
		var ret string
		return ret
	}
	return *o.PromiscuousMode
}

// GetPromiscuousModeOk returns a tuple with the PromiscuousMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitch) GetPromiscuousModeOk() (*string, bool) {
	if o == nil || IsNil(o.PromiscuousMode) {
		return nil, false
	}
	return o.PromiscuousMode, true
}

// HasPromiscuousMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasPromiscuousMode() bool {
	if o != nil && !IsNil(o.PromiscuousMode) {
		return true
	}

	return false
}

// SetPromiscuousMode gets a reference to the given string and assigns it to the PromiscuousMode field.
func (o *VirtualizationVmwareVirtualSwitch) SetPromiscuousMode(v string) {
	o.PromiscuousMode = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareVirtualSwitch) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || IsNil(o.Host.Get()) {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareVirtualSwitch) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitch) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableVirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwareVirtualSwitch) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *VirtualizationVmwareVirtualSwitch) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *VirtualizationVmwareVirtualSwitch) UnsetHost() {
	o.Host.Unset()
}

func (o VirtualizationVmwareVirtualSwitch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareVirtualSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualSwitch, errVirtualizationBaseVirtualSwitch := json.Marshal(o.VirtualizationBaseVirtualSwitch)
	if errVirtualizationBaseVirtualSwitch != nil {
		return map[string]interface{}{}, errVirtualizationBaseVirtualSwitch
	}
	errVirtualizationBaseVirtualSwitch = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualSwitch), &toSerialize)
	if errVirtualizationBaseVirtualSwitch != nil {
		return map[string]interface{}{}, errVirtualizationBaseVirtualSwitch
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ForgedTransmits) {
		toSerialize["ForgedTransmits"] = o.ForgedTransmits
	}
	if !IsNil(o.MacAddressChanges) {
		toSerialize["MacAddressChanges"] = o.MacAddressChanges
	}
	if !IsNil(o.Mtu) {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.NicTeamingAndFailover.IsSet() {
		toSerialize["NicTeamingAndFailover"] = o.NicTeamingAndFailover.Get()
	}
	if !IsNil(o.NumNetworks) {
		toSerialize["NumNetworks"] = o.NumNetworks
	}
	if !IsNil(o.NumPhysicalNetworkInterfaces) {
		toSerialize["NumPhysicalNetworkInterfaces"] = o.NumPhysicalNetworkInterfaces
	}
	if !IsNil(o.PromiscuousMode) {
		toSerialize["PromiscuousMode"] = o.PromiscuousMode
	}
	if o.Host.IsSet() {
		toSerialize["Host"] = o.Host.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareVirtualSwitch) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, the switch does not perform filtering and permits all outbound frames. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
		// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, the switch drops all inbound frames to the adapter. If property value is set to accept and the MAC address is changed, the switch allows frames to the new MAC address to pass. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		MacAddressChanges *string `json:"MacAddressChanges,omitempty"`
		// Maximum transmission unit configured on a virtual switch.
		Mtu                   *int64                                         `json:"Mtu,omitempty"`
		NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
		// Number of networks available on this virtual switch.
		NumNetworks *int64 `json:"NumNetworks,omitempty"`
		// Number of physical network interfaces connected with this virtual switch.
		NumPhysicalNetworkInterfaces *int64 `json:"NumPhysicalNetworkInterfaces,omitempty"`
		// If promiscuousMode property value is set to reject, the virtual switch forwards only frames that are addressed to the adapter. If property value is set to accept, the virtual switch forwards all frames to the adapter in compliance with the active VLAN policy for the port to which it is connected. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		PromiscuousMode *string                                      `json:"PromiscuousMode,omitempty"`
		Host            NullableVirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	}

	varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct := VirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVirtualSwitch := _VirtualizationVmwareVirtualSwitch{}
		varVirtualizationVmwareVirtualSwitch.ClassId = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVirtualSwitch.ObjectType = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVirtualSwitch.ForgedTransmits = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.ForgedTransmits
		varVirtualizationVmwareVirtualSwitch.MacAddressChanges = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.MacAddressChanges
		varVirtualizationVmwareVirtualSwitch.Mtu = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.Mtu
		varVirtualizationVmwareVirtualSwitch.NicTeamingAndFailover = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.NicTeamingAndFailover
		varVirtualizationVmwareVirtualSwitch.NumNetworks = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.NumNetworks
		varVirtualizationVmwareVirtualSwitch.NumPhysicalNetworkInterfaces = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.NumPhysicalNetworkInterfaces
		varVirtualizationVmwareVirtualSwitch.PromiscuousMode = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.PromiscuousMode
		varVirtualizationVmwareVirtualSwitch.Host = varVirtualizationVmwareVirtualSwitchWithoutEmbeddedStruct.Host
		*o = VirtualizationVmwareVirtualSwitch(varVirtualizationVmwareVirtualSwitch)
	} else {
		return err
	}

	varVirtualizationVmwareVirtualSwitch := _VirtualizationVmwareVirtualSwitch{}

	err = json.Unmarshal(data, &varVirtualizationVmwareVirtualSwitch)
	if err == nil {
		o.VirtualizationBaseVirtualSwitch = varVirtualizationVmwareVirtualSwitch.VirtualizationBaseVirtualSwitch
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ForgedTransmits")
		delete(additionalProperties, "MacAddressChanges")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "NicTeamingAndFailover")
		delete(additionalProperties, "NumNetworks")
		delete(additionalProperties, "NumPhysicalNetworkInterfaces")
		delete(additionalProperties, "PromiscuousMode")
		delete(additionalProperties, "Host")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualSwitch := reflect.ValueOf(o.VirtualizationBaseVirtualSwitch)
		for i := 0; i < reflectVirtualizationBaseVirtualSwitch.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualSwitch.Type().Field(i)

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

type NullableVirtualizationVmwareVirtualSwitch struct {
	value *VirtualizationVmwareVirtualSwitch
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualSwitch) Get() *VirtualizationVmwareVirtualSwitch {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualSwitch) Set(val *VirtualizationVmwareVirtualSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualSwitch(val *VirtualizationVmwareVirtualSwitch) *NullableVirtualizationVmwareVirtualSwitch {
	return &NullableVirtualizationVmwareVirtualSwitch{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
