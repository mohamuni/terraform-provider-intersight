/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexIscsiNetwork The HyperFlex iSCSI Network configuration. Contains detailed information about the iSCSI network which includes subnet, gateway, Virtual local area network (VLAN) name, VLAN identity, maximum transmission unit (MTU), ranges of the IP addresses belonging to the HyperFlex iSCSI network.
type HyperflexIscsiNetwork struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The gateway of the HyperFlex iSCSI network.
	Gateway *string `json:"Gateway,omitempty"`
	// Source of this inventory object. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
	InventorySource *string                                        `json:"InventorySource,omitempty"`
	IpRanges        []NetworkHyperFlexInternetProtocolAddressRange `json:"IpRanges,omitempty"`
	// An IP within the iSCSI IP Address Range which is CIP for iSCSI network.
	IscsiCip *string `json:"IscsiCip,omitempty"`
	// The maximum transmission unit of the HyperFlex iSCSI network. * `UNKNOWN` - The maximum transmission unit of the HyperFlex iSCSI network is unknown. * `MTU_1500` - The maximum transmission unit of the HyperFlex iSCSI network is 1500 bytes. * `MTU_9000` - The maximum transmission unit of the HyperFlex iSCSI network is 9000 bytes.
	Mtu *string `json:"Mtu,omitempty"`
	// Name of the HyperFlex iSCSI network configuration.
	Name *string `json:"Name,omitempty"`
	// Subnet of the HyperFlex iSCSI network. Subnet is in a.b.c.d/e notation.
	Subnet *string `json:"Subnet,omitempty"`
	// UCS Manager Host IP or FQDN.
	UcsHost *string `json:"UcsHost,omitempty"`
	// UUID of the HyperFlex iSCSI network configuration.
	Uuid *string `json:"Uuid,omitempty"`
	// Version of Network configuration in Inventory.
	Version *int64 `json:"Version,omitempty"`
	// The Virtual local area network (VLAN) name of the HyperFlex iSCSI network.
	VlanName *string `json:"VlanName,omitempty"`
	// The VLAN ID of the HyperFlex iSCSI network.
	Vlanid               *int64                        `json:"Vlanid,omitempty"`
	Cluster              *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexIscsiNetwork HyperflexIscsiNetwork

// NewHyperflexIscsiNetwork instantiates a new HyperflexIscsiNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexIscsiNetwork(classId string, objectType string) *HyperflexIscsiNetwork {
	this := HyperflexIscsiNetwork{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexIscsiNetworkWithDefaults instantiates a new HyperflexIscsiNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexIscsiNetworkWithDefaults() *HyperflexIscsiNetwork {
	this := HyperflexIscsiNetwork{}
	var classId string = "hyperflex.IscsiNetwork"
	this.ClassId = classId
	var objectType string = "hyperflex.IscsiNetwork"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexIscsiNetwork) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexIscsiNetwork) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexIscsiNetwork) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexIscsiNetwork) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *HyperflexIscsiNetwork) SetGateway(v string) {
	o.Gateway = &v
}

// GetInventorySource returns the InventorySource field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetInventorySource() string {
	if o == nil || o.InventorySource == nil {
		var ret string
		return ret
	}
	return *o.InventorySource
}

// GetInventorySourceOk returns a tuple with the InventorySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetInventorySourceOk() (*string, bool) {
	if o == nil || o.InventorySource == nil {
		return nil, false
	}
	return o.InventorySource, true
}

// HasInventorySource returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasInventorySource() bool {
	if o != nil && o.InventorySource != nil {
		return true
	}

	return false
}

// SetInventorySource gets a reference to the given string and assigns it to the InventorySource field.
func (o *HyperflexIscsiNetwork) SetInventorySource(v string) {
	o.InventorySource = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexIscsiNetwork) GetIpRanges() []NetworkHyperFlexInternetProtocolAddressRange {
	if o == nil {
		var ret []NetworkHyperFlexInternetProtocolAddressRange
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexIscsiNetwork) GetIpRangesOk() ([]NetworkHyperFlexInternetProtocolAddressRange, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []NetworkHyperFlexInternetProtocolAddressRange and assigns it to the IpRanges field.
func (o *HyperflexIscsiNetwork) SetIpRanges(v []NetworkHyperFlexInternetProtocolAddressRange) {
	o.IpRanges = v
}

// GetIscsiCip returns the IscsiCip field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetIscsiCip() string {
	if o == nil || o.IscsiCip == nil {
		var ret string
		return ret
	}
	return *o.IscsiCip
}

// GetIscsiCipOk returns a tuple with the IscsiCip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetIscsiCipOk() (*string, bool) {
	if o == nil || o.IscsiCip == nil {
		return nil, false
	}
	return o.IscsiCip, true
}

// HasIscsiCip returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasIscsiCip() bool {
	if o != nil && o.IscsiCip != nil {
		return true
	}

	return false
}

// SetIscsiCip gets a reference to the given string and assigns it to the IscsiCip field.
func (o *HyperflexIscsiNetwork) SetIscsiCip(v string) {
	o.IscsiCip = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetMtu() string {
	if o == nil || o.Mtu == nil {
		var ret string
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetMtuOk() (*string, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given string and assigns it to the Mtu field.
func (o *HyperflexIscsiNetwork) SetMtu(v string) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexIscsiNetwork) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *HyperflexIscsiNetwork) SetSubnet(v string) {
	o.Subnet = &v
}

// GetUcsHost returns the UcsHost field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetUcsHost() string {
	if o == nil || o.UcsHost == nil {
		var ret string
		return ret
	}
	return *o.UcsHost
}

// GetUcsHostOk returns a tuple with the UcsHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetUcsHostOk() (*string, bool) {
	if o == nil || o.UcsHost == nil {
		return nil, false
	}
	return o.UcsHost, true
}

// HasUcsHost returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasUcsHost() bool {
	if o != nil && o.UcsHost != nil {
		return true
	}

	return false
}

// SetUcsHost gets a reference to the given string and assigns it to the UcsHost field.
func (o *HyperflexIscsiNetwork) SetUcsHost(v string) {
	o.UcsHost = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexIscsiNetwork) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HyperflexIscsiNetwork) SetVersion(v int64) {
	o.Version = &v
}

// GetVlanName returns the VlanName field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetVlanName() string {
	if o == nil || o.VlanName == nil {
		var ret string
		return ret
	}
	return *o.VlanName
}

// GetVlanNameOk returns a tuple with the VlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetVlanNameOk() (*string, bool) {
	if o == nil || o.VlanName == nil {
		return nil, false
	}
	return o.VlanName, true
}

// HasVlanName returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasVlanName() bool {
	if o != nil && o.VlanName != nil {
		return true
	}

	return false
}

// SetVlanName gets a reference to the given string and assigns it to the VlanName field.
func (o *HyperflexIscsiNetwork) SetVlanName(v string) {
	o.VlanName = &v
}

// GetVlanid returns the Vlanid field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetVlanid() int64 {
	if o == nil || o.Vlanid == nil {
		var ret int64
		return ret
	}
	return *o.Vlanid
}

// GetVlanidOk returns a tuple with the Vlanid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetVlanidOk() (*int64, bool) {
	if o == nil || o.Vlanid == nil {
		return nil, false
	}
	return o.Vlanid, true
}

// HasVlanid returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasVlanid() bool {
	if o != nil && o.Vlanid != nil {
		return true
	}

	return false
}

// SetVlanid gets a reference to the given int64 and assigns it to the Vlanid field.
func (o *HyperflexIscsiNetwork) SetVlanid(v int64) {
	o.Vlanid = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexIscsiNetwork) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIscsiNetwork) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexIscsiNetwork) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexIscsiNetwork) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

func (o HyperflexIscsiNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.InventorySource != nil {
		toSerialize["InventorySource"] = o.InventorySource
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	if o.IscsiCip != nil {
		toSerialize["IscsiCip"] = o.IscsiCip
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["Subnet"] = o.Subnet
	}
	if o.UcsHost != nil {
		toSerialize["UcsHost"] = o.UcsHost
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.VlanName != nil {
		toSerialize["VlanName"] = o.VlanName
	}
	if o.Vlanid != nil {
		toSerialize["Vlanid"] = o.Vlanid
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexIscsiNetwork) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexIscsiNetworkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The gateway of the HyperFlex iSCSI network.
		Gateway *string `json:"Gateway,omitempty"`
		// Source of this inventory object. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
		InventorySource *string                                        `json:"InventorySource,omitempty"`
		IpRanges        []NetworkHyperFlexInternetProtocolAddressRange `json:"IpRanges,omitempty"`
		// An IP within the iSCSI IP Address Range which is CIP for iSCSI network.
		IscsiCip *string `json:"IscsiCip,omitempty"`
		// The maximum transmission unit of the HyperFlex iSCSI network. * `UNKNOWN` - The maximum transmission unit of the HyperFlex iSCSI network is unknown. * `MTU_1500` - The maximum transmission unit of the HyperFlex iSCSI network is 1500 bytes. * `MTU_9000` - The maximum transmission unit of the HyperFlex iSCSI network is 9000 bytes.
		Mtu *string `json:"Mtu,omitempty"`
		// Name of the HyperFlex iSCSI network configuration.
		Name *string `json:"Name,omitempty"`
		// Subnet of the HyperFlex iSCSI network. Subnet is in a.b.c.d/e notation.
		Subnet *string `json:"Subnet,omitempty"`
		// UCS Manager Host IP or FQDN.
		UcsHost *string `json:"UcsHost,omitempty"`
		// UUID of the HyperFlex iSCSI network configuration.
		Uuid *string `json:"Uuid,omitempty"`
		// Version of Network configuration in Inventory.
		Version *int64 `json:"Version,omitempty"`
		// The Virtual local area network (VLAN) name of the HyperFlex iSCSI network.
		VlanName *string `json:"VlanName,omitempty"`
		// The VLAN ID of the HyperFlex iSCSI network.
		Vlanid  *int64                        `json:"Vlanid,omitempty"`
		Cluster *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	}

	varHyperflexIscsiNetworkWithoutEmbeddedStruct := HyperflexIscsiNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexIscsiNetworkWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexIscsiNetwork := _HyperflexIscsiNetwork{}
		varHyperflexIscsiNetwork.ClassId = varHyperflexIscsiNetworkWithoutEmbeddedStruct.ClassId
		varHyperflexIscsiNetwork.ObjectType = varHyperflexIscsiNetworkWithoutEmbeddedStruct.ObjectType
		varHyperflexIscsiNetwork.Gateway = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Gateway
		varHyperflexIscsiNetwork.InventorySource = varHyperflexIscsiNetworkWithoutEmbeddedStruct.InventorySource
		varHyperflexIscsiNetwork.IpRanges = varHyperflexIscsiNetworkWithoutEmbeddedStruct.IpRanges
		varHyperflexIscsiNetwork.IscsiCip = varHyperflexIscsiNetworkWithoutEmbeddedStruct.IscsiCip
		varHyperflexIscsiNetwork.Mtu = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Mtu
		varHyperflexIscsiNetwork.Name = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Name
		varHyperflexIscsiNetwork.Subnet = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Subnet
		varHyperflexIscsiNetwork.UcsHost = varHyperflexIscsiNetworkWithoutEmbeddedStruct.UcsHost
		varHyperflexIscsiNetwork.Uuid = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Uuid
		varHyperflexIscsiNetwork.Version = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Version
		varHyperflexIscsiNetwork.VlanName = varHyperflexIscsiNetworkWithoutEmbeddedStruct.VlanName
		varHyperflexIscsiNetwork.Vlanid = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Vlanid
		varHyperflexIscsiNetwork.Cluster = varHyperflexIscsiNetworkWithoutEmbeddedStruct.Cluster
		*o = HyperflexIscsiNetwork(varHyperflexIscsiNetwork)
	} else {
		return err
	}

	varHyperflexIscsiNetwork := _HyperflexIscsiNetwork{}

	err = json.Unmarshal(bytes, &varHyperflexIscsiNetwork)
	if err == nil {
		o.MoBaseMo = varHyperflexIscsiNetwork.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "InventorySource")
		delete(additionalProperties, "IpRanges")
		delete(additionalProperties, "IscsiCip")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Subnet")
		delete(additionalProperties, "UcsHost")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "VlanName")
		delete(additionalProperties, "Vlanid")
		delete(additionalProperties, "Cluster")

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

type NullableHyperflexIscsiNetwork struct {
	value *HyperflexIscsiNetwork
	isSet bool
}

func (v NullableHyperflexIscsiNetwork) Get() *HyperflexIscsiNetwork {
	return v.value
}

func (v *NullableHyperflexIscsiNetwork) Set(val *HyperflexIscsiNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexIscsiNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexIscsiNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexIscsiNetwork(val *HyperflexIscsiNetwork) *NullableHyperflexIscsiNetwork {
	return &NullableHyperflexIscsiNetwork{value: val, isSet: true}
}

func (v NullableHyperflexIscsiNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexIscsiNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}