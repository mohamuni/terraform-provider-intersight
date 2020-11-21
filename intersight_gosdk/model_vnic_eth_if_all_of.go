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
)

// VnicEthIfAllOf Definition of the list of properties defined in 'vnic.EthIf', excluding properties defined in parent classes.
type VnicEthIfAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string          `json:"ObjectType"`
	Cdn        NullableVnicCdn `json:"Cdn,omitempty"`
	// Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover.
	FailoverEnabled *bool `json:"FailoverEnabled,omitempty"`
	// The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the virtual ethernet interface.
	Name *string `json:"Name,omitempty"`
	// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
	Order     *int64                        `json:"Order,omitempty"`
	Placement NullableVnicPlacementSettings `json:"Placement,omitempty"`
	// The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path.
	StandbyVifId  *int64                    `json:"StandbyVifId,omitempty"`
	UsnicSettings NullableVnicUsnicSettings `json:"UsnicSettings,omitempty"`
	// The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC.
	VifId                         *int64                                     `json:"VifId,omitempty"`
	VmqSettings                   NullableVnicVmqSettings                    `json:"VmqSettings,omitempty"`
	EthAdapterPolicy              *VnicEthAdapterPolicyRelationship          `json:"EthAdapterPolicy,omitempty"`
	EthNetworkPolicy              *VnicEthNetworkPolicyRelationship          `json:"EthNetworkPolicy,omitempty"`
	EthQosPolicy                  *VnicEthQosPolicyRelationship              `json:"EthQosPolicy,omitempty"`
	FabricEthNetworkControlPolicy *FabricEthNetworkControlPolicyRelationship `json:"FabricEthNetworkControlPolicy,omitempty"`
	// An array of relationships to fabricEthNetworkGroupPolicy resources.
	FabricEthNetworkGroupPolicy []FabricEthNetworkGroupPolicyRelationship `json:"FabricEthNetworkGroupPolicy,omitempty"`
	LanConnectivityPolicy       *VnicLanConnectivityPolicyRelationship    `json:"LanConnectivityPolicy,omitempty"`
	LcpVnic                     *VnicEthIfRelationship                    `json:"LcpVnic,omitempty"`
	MacLease                    *MacpoolLeaseRelationship                 `json:"MacLease,omitempty"`
	MacPool                     *MacpoolPoolRelationship                  `json:"MacPool,omitempty"`
	Profile                     *PolicyAbstractConfigProfileRelationship  `json:"Profile,omitempty"`
	// An array of relationships to vnicEthIf resources.
	SpVnics              []VnicEthIfRelationship `json:"SpVnics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthIfAllOf VnicEthIfAllOf

// NewVnicEthIfAllOf instantiates a new VnicEthIfAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthIfAllOf(classId string, objectType string) *VnicEthIfAllOf {
	this := VnicEthIfAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicEthIfAllOfWithDefaults instantiates a new VnicEthIfAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthIfAllOfWithDefaults() *VnicEthIfAllOf {
	this := VnicEthIfAllOf{}
	var classId string = "vnic.EthIf"
	this.ClassId = classId
	var objectType string = "vnic.EthIf"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthIfAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthIfAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthIfAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthIfAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCdn returns the Cdn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIfAllOf) GetCdn() VnicCdn {
	if o == nil || o.Cdn.Get() == nil {
		var ret VnicCdn
		return ret
	}
	return *o.Cdn.Get()
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIfAllOf) GetCdnOk() (*VnicCdn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cdn.Get(), o.Cdn.IsSet()
}

// HasCdn returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasCdn() bool {
	if o != nil && o.Cdn.IsSet() {
		return true
	}

	return false
}

// SetCdn gets a reference to the given NullableVnicCdn and assigns it to the Cdn field.
func (o *VnicEthIfAllOf) SetCdn(v VnicCdn) {
	o.Cdn.Set(&v)
}

// SetCdnNil sets the value for Cdn to be an explicit nil
func (o *VnicEthIfAllOf) SetCdnNil() {
	o.Cdn.Set(nil)
}

// UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
func (o *VnicEthIfAllOf) UnsetCdn() {
	o.Cdn.Unset()
}

// GetFailoverEnabled returns the FailoverEnabled field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetFailoverEnabled() bool {
	if o == nil || o.FailoverEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FailoverEnabled
}

// GetFailoverEnabledOk returns a tuple with the FailoverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetFailoverEnabledOk() (*bool, bool) {
	if o == nil || o.FailoverEnabled == nil {
		return nil, false
	}
	return o.FailoverEnabled, true
}

// HasFailoverEnabled returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasFailoverEnabled() bool {
	if o != nil && o.FailoverEnabled != nil {
		return true
	}

	return false
}

// SetFailoverEnabled gets a reference to the given bool and assigns it to the FailoverEnabled field.
func (o *VnicEthIfAllOf) SetFailoverEnabled(v bool) {
	o.FailoverEnabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VnicEthIfAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicEthIfAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VnicEthIfAllOf) SetOrder(v int64) {
	o.Order = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIfAllOf) GetPlacement() VnicPlacementSettings {
	if o == nil || o.Placement.Get() == nil {
		var ret VnicPlacementSettings
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIfAllOf) GetPlacementOk() (*VnicPlacementSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableVnicPlacementSettings and assigns it to the Placement field.
func (o *VnicEthIfAllOf) SetPlacement(v VnicPlacementSettings) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *VnicEthIfAllOf) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *VnicEthIfAllOf) UnsetPlacement() {
	o.Placement.Unset()
}

// GetStandbyVifId returns the StandbyVifId field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetStandbyVifId() int64 {
	if o == nil || o.StandbyVifId == nil {
		var ret int64
		return ret
	}
	return *o.StandbyVifId
}

// GetStandbyVifIdOk returns a tuple with the StandbyVifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetStandbyVifIdOk() (*int64, bool) {
	if o == nil || o.StandbyVifId == nil {
		return nil, false
	}
	return o.StandbyVifId, true
}

// HasStandbyVifId returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasStandbyVifId() bool {
	if o != nil && o.StandbyVifId != nil {
		return true
	}

	return false
}

// SetStandbyVifId gets a reference to the given int64 and assigns it to the StandbyVifId field.
func (o *VnicEthIfAllOf) SetStandbyVifId(v int64) {
	o.StandbyVifId = &v
}

// GetUsnicSettings returns the UsnicSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIfAllOf) GetUsnicSettings() VnicUsnicSettings {
	if o == nil || o.UsnicSettings.Get() == nil {
		var ret VnicUsnicSettings
		return ret
	}
	return *o.UsnicSettings.Get()
}

// GetUsnicSettingsOk returns a tuple with the UsnicSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIfAllOf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsnicSettings.Get(), o.UsnicSettings.IsSet()
}

// HasUsnicSettings returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasUsnicSettings() bool {
	if o != nil && o.UsnicSettings.IsSet() {
		return true
	}

	return false
}

// SetUsnicSettings gets a reference to the given NullableVnicUsnicSettings and assigns it to the UsnicSettings field.
func (o *VnicEthIfAllOf) SetUsnicSettings(v VnicUsnicSettings) {
	o.UsnicSettings.Set(&v)
}

// SetUsnicSettingsNil sets the value for UsnicSettings to be an explicit nil
func (o *VnicEthIfAllOf) SetUsnicSettingsNil() {
	o.UsnicSettings.Set(nil)
}

// UnsetUsnicSettings ensures that no value is present for UsnicSettings, not even an explicit nil
func (o *VnicEthIfAllOf) UnsetUsnicSettings() {
	o.UsnicSettings.Unset()
}

// GetVifId returns the VifId field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetVifId() int64 {
	if o == nil || o.VifId == nil {
		var ret int64
		return ret
	}
	return *o.VifId
}

// GetVifIdOk returns a tuple with the VifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetVifIdOk() (*int64, bool) {
	if o == nil || o.VifId == nil {
		return nil, false
	}
	return o.VifId, true
}

// HasVifId returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasVifId() bool {
	if o != nil && o.VifId != nil {
		return true
	}

	return false
}

// SetVifId gets a reference to the given int64 and assigns it to the VifId field.
func (o *VnicEthIfAllOf) SetVifId(v int64) {
	o.VifId = &v
}

// GetVmqSettings returns the VmqSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIfAllOf) GetVmqSettings() VnicVmqSettings {
	if o == nil || o.VmqSettings.Get() == nil {
		var ret VnicVmqSettings
		return ret
	}
	return *o.VmqSettings.Get()
}

// GetVmqSettingsOk returns a tuple with the VmqSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIfAllOf) GetVmqSettingsOk() (*VnicVmqSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmqSettings.Get(), o.VmqSettings.IsSet()
}

// HasVmqSettings returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasVmqSettings() bool {
	if o != nil && o.VmqSettings.IsSet() {
		return true
	}

	return false
}

// SetVmqSettings gets a reference to the given NullableVnicVmqSettings and assigns it to the VmqSettings field.
func (o *VnicEthIfAllOf) SetVmqSettings(v VnicVmqSettings) {
	o.VmqSettings.Set(&v)
}

// SetVmqSettingsNil sets the value for VmqSettings to be an explicit nil
func (o *VnicEthIfAllOf) SetVmqSettingsNil() {
	o.VmqSettings.Set(nil)
}

// UnsetVmqSettings ensures that no value is present for VmqSettings, not even an explicit nil
func (o *VnicEthIfAllOf) UnsetVmqSettings() {
	o.VmqSettings.Unset()
}

// GetEthAdapterPolicy returns the EthAdapterPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship {
	if o == nil || o.EthAdapterPolicy == nil {
		var ret VnicEthAdapterPolicyRelationship
		return ret
	}
	return *o.EthAdapterPolicy
}

// GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool) {
	if o == nil || o.EthAdapterPolicy == nil {
		return nil, false
	}
	return o.EthAdapterPolicy, true
}

// HasEthAdapterPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasEthAdapterPolicy() bool {
	if o != nil && o.EthAdapterPolicy != nil {
		return true
	}

	return false
}

// SetEthAdapterPolicy gets a reference to the given VnicEthAdapterPolicyRelationship and assigns it to the EthAdapterPolicy field.
func (o *VnicEthIfAllOf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship) {
	o.EthAdapterPolicy = &v
}

// GetEthNetworkPolicy returns the EthNetworkPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship {
	if o == nil || o.EthNetworkPolicy == nil {
		var ret VnicEthNetworkPolicyRelationship
		return ret
	}
	return *o.EthNetworkPolicy
}

// GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool) {
	if o == nil || o.EthNetworkPolicy == nil {
		return nil, false
	}
	return o.EthNetworkPolicy, true
}

// HasEthNetworkPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasEthNetworkPolicy() bool {
	if o != nil && o.EthNetworkPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkPolicy gets a reference to the given VnicEthNetworkPolicyRelationship and assigns it to the EthNetworkPolicy field.
func (o *VnicEthIfAllOf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship) {
	o.EthNetworkPolicy = &v
}

// GetEthQosPolicy returns the EthQosPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetEthQosPolicy() VnicEthQosPolicyRelationship {
	if o == nil || o.EthQosPolicy == nil {
		var ret VnicEthQosPolicyRelationship
		return ret
	}
	return *o.EthQosPolicy
}

// GetEthQosPolicyOk returns a tuple with the EthQosPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool) {
	if o == nil || o.EthQosPolicy == nil {
		return nil, false
	}
	return o.EthQosPolicy, true
}

// HasEthQosPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasEthQosPolicy() bool {
	if o != nil && o.EthQosPolicy != nil {
		return true
	}

	return false
}

// SetEthQosPolicy gets a reference to the given VnicEthQosPolicyRelationship and assigns it to the EthQosPolicy field.
func (o *VnicEthIfAllOf) SetEthQosPolicy(v VnicEthQosPolicyRelationship) {
	o.EthQosPolicy = &v
}

// GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship {
	if o == nil || o.FabricEthNetworkControlPolicy == nil {
		var ret FabricEthNetworkControlPolicyRelationship
		return ret
	}
	return *o.FabricEthNetworkControlPolicy
}

// GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool) {
	if o == nil || o.FabricEthNetworkControlPolicy == nil {
		return nil, false
	}
	return o.FabricEthNetworkControlPolicy, true
}

// HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasFabricEthNetworkControlPolicy() bool {
	if o != nil && o.FabricEthNetworkControlPolicy != nil {
		return true
	}

	return false
}

// SetFabricEthNetworkControlPolicy gets a reference to the given FabricEthNetworkControlPolicyRelationship and assigns it to the FabricEthNetworkControlPolicy field.
func (o *VnicEthIfAllOf) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship) {
	o.FabricEthNetworkControlPolicy = &v
}

// GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIfAllOf) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship {
	if o == nil {
		var ret []FabricEthNetworkGroupPolicyRelationship
		return ret
	}
	return o.FabricEthNetworkGroupPolicy
}

// GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIfAllOf) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool) {
	if o == nil || o.FabricEthNetworkGroupPolicy == nil {
		return nil, false
	}
	return &o.FabricEthNetworkGroupPolicy, true
}

// HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasFabricEthNetworkGroupPolicy() bool {
	if o != nil && o.FabricEthNetworkGroupPolicy != nil {
		return true
	}

	return false
}

// SetFabricEthNetworkGroupPolicy gets a reference to the given []FabricEthNetworkGroupPolicyRelationship and assigns it to the FabricEthNetworkGroupPolicy field.
func (o *VnicEthIfAllOf) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship) {
	o.FabricEthNetworkGroupPolicy = v
}

// GetLanConnectivityPolicy returns the LanConnectivityPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetLanConnectivityPolicy() VnicLanConnectivityPolicyRelationship {
	if o == nil || o.LanConnectivityPolicy == nil {
		var ret VnicLanConnectivityPolicyRelationship
		return ret
	}
	return *o.LanConnectivityPolicy
}

// GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyRelationship, bool) {
	if o == nil || o.LanConnectivityPolicy == nil {
		return nil, false
	}
	return o.LanConnectivityPolicy, true
}

// HasLanConnectivityPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasLanConnectivityPolicy() bool {
	if o != nil && o.LanConnectivityPolicy != nil {
		return true
	}

	return false
}

// SetLanConnectivityPolicy gets a reference to the given VnicLanConnectivityPolicyRelationship and assigns it to the LanConnectivityPolicy field.
func (o *VnicEthIfAllOf) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyRelationship) {
	o.LanConnectivityPolicy = &v
}

// GetLcpVnic returns the LcpVnic field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetLcpVnic() VnicEthIfRelationship {
	if o == nil || o.LcpVnic == nil {
		var ret VnicEthIfRelationship
		return ret
	}
	return *o.LcpVnic
}

// GetLcpVnicOk returns a tuple with the LcpVnic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetLcpVnicOk() (*VnicEthIfRelationship, bool) {
	if o == nil || o.LcpVnic == nil {
		return nil, false
	}
	return o.LcpVnic, true
}

// HasLcpVnic returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasLcpVnic() bool {
	if o != nil && o.LcpVnic != nil {
		return true
	}

	return false
}

// SetLcpVnic gets a reference to the given VnicEthIfRelationship and assigns it to the LcpVnic field.
func (o *VnicEthIfAllOf) SetLcpVnic(v VnicEthIfRelationship) {
	o.LcpVnic = &v
}

// GetMacLease returns the MacLease field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetMacLease() MacpoolLeaseRelationship {
	if o == nil || o.MacLease == nil {
		var ret MacpoolLeaseRelationship
		return ret
	}
	return *o.MacLease
}

// GetMacLeaseOk returns a tuple with the MacLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool) {
	if o == nil || o.MacLease == nil {
		return nil, false
	}
	return o.MacLease, true
}

// HasMacLease returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasMacLease() bool {
	if o != nil && o.MacLease != nil {
		return true
	}

	return false
}

// SetMacLease gets a reference to the given MacpoolLeaseRelationship and assigns it to the MacLease field.
func (o *VnicEthIfAllOf) SetMacLease(v MacpoolLeaseRelationship) {
	o.MacLease = &v
}

// GetMacPool returns the MacPool field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetMacPool() MacpoolPoolRelationship {
	if o == nil || o.MacPool == nil {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.MacPool
}

// GetMacPoolOk returns a tuple with the MacPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetMacPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil || o.MacPool == nil {
		return nil, false
	}
	return o.MacPool, true
}

// HasMacPool returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasMacPool() bool {
	if o != nil && o.MacPool != nil {
		return true
	}

	return false
}

// SetMacPool gets a reference to the given MacpoolPoolRelationship and assigns it to the MacPool field.
func (o *VnicEthIfAllOf) SetMacPool(v MacpoolPoolRelationship) {
	o.MacPool = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicEthIfAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile = &v
}

// GetSpVnics returns the SpVnics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIfAllOf) GetSpVnics() []VnicEthIfRelationship {
	if o == nil {
		var ret []VnicEthIfRelationship
		return ret
	}
	return o.SpVnics
}

// GetSpVnicsOk returns a tuple with the SpVnics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIfAllOf) GetSpVnicsOk() (*[]VnicEthIfRelationship, bool) {
	if o == nil || o.SpVnics == nil {
		return nil, false
	}
	return &o.SpVnics, true
}

// HasSpVnics returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasSpVnics() bool {
	if o != nil && o.SpVnics != nil {
		return true
	}

	return false
}

// SetSpVnics gets a reference to the given []VnicEthIfRelationship and assigns it to the SpVnics field.
func (o *VnicEthIfAllOf) SetSpVnics(v []VnicEthIfRelationship) {
	o.SpVnics = v
}

func (o VnicEthIfAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cdn.IsSet() {
		toSerialize["Cdn"] = o.Cdn.Get()
	}
	if o.FailoverEnabled != nil {
		toSerialize["FailoverEnabled"] = o.FailoverEnabled
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.Placement.IsSet() {
		toSerialize["Placement"] = o.Placement.Get()
	}
	if o.StandbyVifId != nil {
		toSerialize["StandbyVifId"] = o.StandbyVifId
	}
	if o.UsnicSettings.IsSet() {
		toSerialize["UsnicSettings"] = o.UsnicSettings.Get()
	}
	if o.VifId != nil {
		toSerialize["VifId"] = o.VifId
	}
	if o.VmqSettings.IsSet() {
		toSerialize["VmqSettings"] = o.VmqSettings.Get()
	}
	if o.EthAdapterPolicy != nil {
		toSerialize["EthAdapterPolicy"] = o.EthAdapterPolicy
	}
	if o.EthNetworkPolicy != nil {
		toSerialize["EthNetworkPolicy"] = o.EthNetworkPolicy
	}
	if o.EthQosPolicy != nil {
		toSerialize["EthQosPolicy"] = o.EthQosPolicy
	}
	if o.FabricEthNetworkControlPolicy != nil {
		toSerialize["FabricEthNetworkControlPolicy"] = o.FabricEthNetworkControlPolicy
	}
	if o.FabricEthNetworkGroupPolicy != nil {
		toSerialize["FabricEthNetworkGroupPolicy"] = o.FabricEthNetworkGroupPolicy
	}
	if o.LanConnectivityPolicy != nil {
		toSerialize["LanConnectivityPolicy"] = o.LanConnectivityPolicy
	}
	if o.LcpVnic != nil {
		toSerialize["LcpVnic"] = o.LcpVnic
	}
	if o.MacLease != nil {
		toSerialize["MacLease"] = o.MacLease
	}
	if o.MacPool != nil {
		toSerialize["MacPool"] = o.MacPool
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.SpVnics != nil {
		toSerialize["SpVnics"] = o.SpVnics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthIfAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicEthIfAllOf := _VnicEthIfAllOf{}

	if err = json.Unmarshal(bytes, &varVnicEthIfAllOf); err == nil {
		*o = VnicEthIfAllOf(varVnicEthIfAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cdn")
		delete(additionalProperties, "FailoverEnabled")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "Placement")
		delete(additionalProperties, "StandbyVifId")
		delete(additionalProperties, "UsnicSettings")
		delete(additionalProperties, "VifId")
		delete(additionalProperties, "VmqSettings")
		delete(additionalProperties, "EthAdapterPolicy")
		delete(additionalProperties, "EthNetworkPolicy")
		delete(additionalProperties, "EthQosPolicy")
		delete(additionalProperties, "FabricEthNetworkControlPolicy")
		delete(additionalProperties, "FabricEthNetworkGroupPolicy")
		delete(additionalProperties, "LanConnectivityPolicy")
		delete(additionalProperties, "LcpVnic")
		delete(additionalProperties, "MacLease")
		delete(additionalProperties, "MacPool")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "SpVnics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthIfAllOf struct {
	value *VnicEthIfAllOf
	isSet bool
}

func (v NullableVnicEthIfAllOf) Get() *VnicEthIfAllOf {
	return v.value
}

func (v *NullableVnicEthIfAllOf) Set(val *VnicEthIfAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthIfAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthIfAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthIfAllOf(val *VnicEthIfAllOf) *NullableVnicEthIfAllOf {
	return &NullableVnicEthIfAllOf{value: val, isSet: true}
}

func (v NullableVnicEthIfAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthIfAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}