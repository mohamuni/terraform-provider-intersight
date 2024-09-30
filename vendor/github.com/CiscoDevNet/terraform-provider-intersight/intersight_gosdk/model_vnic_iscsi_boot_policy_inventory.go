/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
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

// checks if the VnicIscsiBootPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicIscsiBootPolicyInventory{}

// VnicIscsiBootPolicyInventory Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
type VnicIscsiBootPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters.
	AutoTargetvendorName *string                      `json:"AutoTargetvendorName,omitempty"`
	Chap                 NullableVnicIscsiAuthProfile `json:"Chap,omitempty"`
	// Source Type of Initiator IP Address - Auto/Static/Pool. * `DHCP` - The IP address is assigned using DHCP, if available. * `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.
	InitiatorIpSource *string `json:"InitiatorIpSource,omitempty"`
	// Static IP address provided for iSCSI Initiator.
	InitiatorStaticIpV4Address *string                      `json:"InitiatorStaticIpV4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	InitiatorStaticIpV4Config  NullableIppoolIpV4Config     `json:"InitiatorStaticIpV4Config,omitempty"`
	MutualChap                 NullableVnicIscsiAuthProfile `json:"MutualChap,omitempty"`
	// Source Type of Targets - Auto/Static. * `Static` - Type indicates that static target interface is assigned to iSCSI boot. * `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot.
	TargetSourceType      *string                                                  `json:"TargetSourceType,omitempty"`
	InitiatorIpPool       NullableIppoolPoolRelationship                           `json:"InitiatorIpPool,omitempty"`
	IscsiAdapterPolicy    NullableVnicIscsiAdapterPolicyInventoryRelationship      `json:"IscsiAdapterPolicy,omitempty"`
	PrimaryTargetPolicy   NullableVnicIscsiStaticTargetPolicyInventoryRelationship `json:"PrimaryTargetPolicy,omitempty"`
	SecondaryTargetPolicy NullableVnicIscsiStaticTargetPolicyInventoryRelationship `json:"SecondaryTargetPolicy,omitempty"`
	TargetMo              NullableMoBaseMoRelationship                             `json:"TargetMo,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _VnicIscsiBootPolicyInventory VnicIscsiBootPolicyInventory

// NewVnicIscsiBootPolicyInventory instantiates a new VnicIscsiBootPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiBootPolicyInventory(classId string, objectType string) *VnicIscsiBootPolicyInventory {
	this := VnicIscsiBootPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicIscsiBootPolicyInventoryWithDefaults instantiates a new VnicIscsiBootPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiBootPolicyInventoryWithDefaults() *VnicIscsiBootPolicyInventory {
	this := VnicIscsiBootPolicyInventory{}
	var classId string = "vnic.IscsiBootPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.IscsiBootPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiBootPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiBootPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.IscsiBootPolicyInventory" of the ClassId field.
func (o *VnicIscsiBootPolicyInventory) GetDefaultClassId() interface{} {
	return "vnic.IscsiBootPolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiBootPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiBootPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.IscsiBootPolicyInventory" of the ObjectType field.
func (o *VnicIscsiBootPolicyInventory) GetDefaultObjectType() interface{} {
	return "vnic.IscsiBootPolicyInventory"
}

// GetAutoTargetvendorName returns the AutoTargetvendorName field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyInventory) GetAutoTargetvendorName() string {
	if o == nil || IsNil(o.AutoTargetvendorName) {
		var ret string
		return ret
	}
	return *o.AutoTargetvendorName
}

// GetAutoTargetvendorNameOk returns a tuple with the AutoTargetvendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyInventory) GetAutoTargetvendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AutoTargetvendorName) {
		return nil, false
	}
	return o.AutoTargetvendorName, true
}

// HasAutoTargetvendorName returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasAutoTargetvendorName() bool {
	if o != nil && !IsNil(o.AutoTargetvendorName) {
		return true
	}

	return false
}

// SetAutoTargetvendorName gets a reference to the given string and assigns it to the AutoTargetvendorName field.
func (o *VnicIscsiBootPolicyInventory) SetAutoTargetvendorName(v string) {
	o.AutoTargetvendorName = &v
}

// GetChap returns the Chap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetChap() VnicIscsiAuthProfile {
	if o == nil || IsNil(o.Chap.Get()) {
		var ret VnicIscsiAuthProfile
		return ret
	}
	return *o.Chap.Get()
}

// GetChapOk returns a tuple with the Chap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetChapOk() (*VnicIscsiAuthProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chap.Get(), o.Chap.IsSet()
}

// HasChap returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasChap() bool {
	if o != nil && o.Chap.IsSet() {
		return true
	}

	return false
}

// SetChap gets a reference to the given NullableVnicIscsiAuthProfile and assigns it to the Chap field.
func (o *VnicIscsiBootPolicyInventory) SetChap(v VnicIscsiAuthProfile) {
	o.Chap.Set(&v)
}

// SetChapNil sets the value for Chap to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetChapNil() {
	o.Chap.Set(nil)
}

// UnsetChap ensures that no value is present for Chap, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetChap() {
	o.Chap.Unset()
}

// GetInitiatorIpSource returns the InitiatorIpSource field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpSource() string {
	if o == nil || IsNil(o.InitiatorIpSource) {
		var ret string
		return ret
	}
	return *o.InitiatorIpSource
}

// GetInitiatorIpSourceOk returns a tuple with the InitiatorIpSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpSourceOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorIpSource) {
		return nil, false
	}
	return o.InitiatorIpSource, true
}

// HasInitiatorIpSource returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasInitiatorIpSource() bool {
	if o != nil && !IsNil(o.InitiatorIpSource) {
		return true
	}

	return false
}

// SetInitiatorIpSource gets a reference to the given string and assigns it to the InitiatorIpSource field.
func (o *VnicIscsiBootPolicyInventory) SetInitiatorIpSource(v string) {
	o.InitiatorIpSource = &v
}

// GetInitiatorStaticIpV4Address returns the InitiatorStaticIpV4Address field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4Address() string {
	if o == nil || IsNil(o.InitiatorStaticIpV4Address) {
		var ret string
		return ret
	}
	return *o.InitiatorStaticIpV4Address
}

// GetInitiatorStaticIpV4AddressOk returns a tuple with the InitiatorStaticIpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorStaticIpV4Address) {
		return nil, false
	}
	return o.InitiatorStaticIpV4Address, true
}

// HasInitiatorStaticIpV4Address returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasInitiatorStaticIpV4Address() bool {
	if o != nil && !IsNil(o.InitiatorStaticIpV4Address) {
		return true
	}

	return false
}

// SetInitiatorStaticIpV4Address gets a reference to the given string and assigns it to the InitiatorStaticIpV4Address field.
func (o *VnicIscsiBootPolicyInventory) SetInitiatorStaticIpV4Address(v string) {
	o.InitiatorStaticIpV4Address = &v
}

// GetInitiatorStaticIpV4Config returns the InitiatorStaticIpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4Config() IppoolIpV4Config {
	if o == nil || IsNil(o.InitiatorStaticIpV4Config.Get()) {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.InitiatorStaticIpV4Config.Get()
}

// GetInitiatorStaticIpV4ConfigOk returns a tuple with the InitiatorStaticIpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorStaticIpV4Config.Get(), o.InitiatorStaticIpV4Config.IsSet()
}

// HasInitiatorStaticIpV4Config returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasInitiatorStaticIpV4Config() bool {
	if o != nil && o.InitiatorStaticIpV4Config.IsSet() {
		return true
	}

	return false
}

// SetInitiatorStaticIpV4Config gets a reference to the given NullableIppoolIpV4Config and assigns it to the InitiatorStaticIpV4Config field.
func (o *VnicIscsiBootPolicyInventory) SetInitiatorStaticIpV4Config(v IppoolIpV4Config) {
	o.InitiatorStaticIpV4Config.Set(&v)
}

// SetInitiatorStaticIpV4ConfigNil sets the value for InitiatorStaticIpV4Config to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetInitiatorStaticIpV4ConfigNil() {
	o.InitiatorStaticIpV4Config.Set(nil)
}

// UnsetInitiatorStaticIpV4Config ensures that no value is present for InitiatorStaticIpV4Config, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetInitiatorStaticIpV4Config() {
	o.InitiatorStaticIpV4Config.Unset()
}

// GetMutualChap returns the MutualChap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetMutualChap() VnicIscsiAuthProfile {
	if o == nil || IsNil(o.MutualChap.Get()) {
		var ret VnicIscsiAuthProfile
		return ret
	}
	return *o.MutualChap.Get()
}

// GetMutualChapOk returns a tuple with the MutualChap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetMutualChapOk() (*VnicIscsiAuthProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.MutualChap.Get(), o.MutualChap.IsSet()
}

// HasMutualChap returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasMutualChap() bool {
	if o != nil && o.MutualChap.IsSet() {
		return true
	}

	return false
}

// SetMutualChap gets a reference to the given NullableVnicIscsiAuthProfile and assigns it to the MutualChap field.
func (o *VnicIscsiBootPolicyInventory) SetMutualChap(v VnicIscsiAuthProfile) {
	o.MutualChap.Set(&v)
}

// SetMutualChapNil sets the value for MutualChap to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetMutualChapNil() {
	o.MutualChap.Set(nil)
}

// UnsetMutualChap ensures that no value is present for MutualChap, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetMutualChap() {
	o.MutualChap.Unset()
}

// GetTargetSourceType returns the TargetSourceType field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyInventory) GetTargetSourceType() string {
	if o == nil || IsNil(o.TargetSourceType) {
		var ret string
		return ret
	}
	return *o.TargetSourceType
}

// GetTargetSourceTypeOk returns a tuple with the TargetSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyInventory) GetTargetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetSourceType) {
		return nil, false
	}
	return o.TargetSourceType, true
}

// HasTargetSourceType returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasTargetSourceType() bool {
	if o != nil && !IsNil(o.TargetSourceType) {
		return true
	}

	return false
}

// SetTargetSourceType gets a reference to the given string and assigns it to the TargetSourceType field.
func (o *VnicIscsiBootPolicyInventory) SetTargetSourceType(v string) {
	o.TargetSourceType = &v
}

// GetInitiatorIpPool returns the InitiatorIpPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.InitiatorIpPool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InitiatorIpPool.Get()
}

// GetInitiatorIpPoolOk returns a tuple with the InitiatorIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorIpPool.Get(), o.InitiatorIpPool.IsSet()
}

// HasInitiatorIpPool returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasInitiatorIpPool() bool {
	if o != nil && o.InitiatorIpPool.IsSet() {
		return true
	}

	return false
}

// SetInitiatorIpPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the InitiatorIpPool field.
func (o *VnicIscsiBootPolicyInventory) SetInitiatorIpPool(v IppoolPoolRelationship) {
	o.InitiatorIpPool.Set(&v)
}

// SetInitiatorIpPoolNil sets the value for InitiatorIpPool to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetInitiatorIpPoolNil() {
	o.InitiatorIpPool.Set(nil)
}

// UnsetInitiatorIpPool ensures that no value is present for InitiatorIpPool, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetInitiatorIpPool() {
	o.InitiatorIpPool.Unset()
}

// GetIscsiAdapterPolicy returns the IscsiAdapterPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetIscsiAdapterPolicy() VnicIscsiAdapterPolicyInventoryRelationship {
	if o == nil || IsNil(o.IscsiAdapterPolicy.Get()) {
		var ret VnicIscsiAdapterPolicyInventoryRelationship
		return ret
	}
	return *o.IscsiAdapterPolicy.Get()
}

// GetIscsiAdapterPolicyOk returns a tuple with the IscsiAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetIscsiAdapterPolicyOk() (*VnicIscsiAdapterPolicyInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.IscsiAdapterPolicy.Get(), o.IscsiAdapterPolicy.IsSet()
}

// HasIscsiAdapterPolicy returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasIscsiAdapterPolicy() bool {
	if o != nil && o.IscsiAdapterPolicy.IsSet() {
		return true
	}

	return false
}

// SetIscsiAdapterPolicy gets a reference to the given NullableVnicIscsiAdapterPolicyInventoryRelationship and assigns it to the IscsiAdapterPolicy field.
func (o *VnicIscsiBootPolicyInventory) SetIscsiAdapterPolicy(v VnicIscsiAdapterPolicyInventoryRelationship) {
	o.IscsiAdapterPolicy.Set(&v)
}

// SetIscsiAdapterPolicyNil sets the value for IscsiAdapterPolicy to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetIscsiAdapterPolicyNil() {
	o.IscsiAdapterPolicy.Set(nil)
}

// UnsetIscsiAdapterPolicy ensures that no value is present for IscsiAdapterPolicy, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetIscsiAdapterPolicy() {
	o.IscsiAdapterPolicy.Unset()
}

// GetPrimaryTargetPolicy returns the PrimaryTargetPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetPrimaryTargetPolicy() VnicIscsiStaticTargetPolicyInventoryRelationship {
	if o == nil || IsNil(o.PrimaryTargetPolicy.Get()) {
		var ret VnicIscsiStaticTargetPolicyInventoryRelationship
		return ret
	}
	return *o.PrimaryTargetPolicy.Get()
}

// GetPrimaryTargetPolicyOk returns a tuple with the PrimaryTargetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetPrimaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryTargetPolicy.Get(), o.PrimaryTargetPolicy.IsSet()
}

// HasPrimaryTargetPolicy returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasPrimaryTargetPolicy() bool {
	if o != nil && o.PrimaryTargetPolicy.IsSet() {
		return true
	}

	return false
}

// SetPrimaryTargetPolicy gets a reference to the given NullableVnicIscsiStaticTargetPolicyInventoryRelationship and assigns it to the PrimaryTargetPolicy field.
func (o *VnicIscsiBootPolicyInventory) SetPrimaryTargetPolicy(v VnicIscsiStaticTargetPolicyInventoryRelationship) {
	o.PrimaryTargetPolicy.Set(&v)
}

// SetPrimaryTargetPolicyNil sets the value for PrimaryTargetPolicy to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetPrimaryTargetPolicyNil() {
	o.PrimaryTargetPolicy.Set(nil)
}

// UnsetPrimaryTargetPolicy ensures that no value is present for PrimaryTargetPolicy, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetPrimaryTargetPolicy() {
	o.PrimaryTargetPolicy.Unset()
}

// GetSecondaryTargetPolicy returns the SecondaryTargetPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetSecondaryTargetPolicy() VnicIscsiStaticTargetPolicyInventoryRelationship {
	if o == nil || IsNil(o.SecondaryTargetPolicy.Get()) {
		var ret VnicIscsiStaticTargetPolicyInventoryRelationship
		return ret
	}
	return *o.SecondaryTargetPolicy.Get()
}

// GetSecondaryTargetPolicyOk returns a tuple with the SecondaryTargetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetSecondaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryTargetPolicy.Get(), o.SecondaryTargetPolicy.IsSet()
}

// HasSecondaryTargetPolicy returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasSecondaryTargetPolicy() bool {
	if o != nil && o.SecondaryTargetPolicy.IsSet() {
		return true
	}

	return false
}

// SetSecondaryTargetPolicy gets a reference to the given NullableVnicIscsiStaticTargetPolicyInventoryRelationship and assigns it to the SecondaryTargetPolicy field.
func (o *VnicIscsiBootPolicyInventory) SetSecondaryTargetPolicy(v VnicIscsiStaticTargetPolicyInventoryRelationship) {
	o.SecondaryTargetPolicy.Set(&v)
}

// SetSecondaryTargetPolicyNil sets the value for SecondaryTargetPolicy to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetSecondaryTargetPolicyNil() {
	o.SecondaryTargetPolicy.Set(nil)
}

// UnsetSecondaryTargetPolicy ensures that no value is present for SecondaryTargetPolicy, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetSecondaryTargetPolicy() {
	o.SecondaryTargetPolicy.Unset()
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicIscsiBootPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *VnicIscsiBootPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *VnicIscsiBootPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o VnicIscsiBootPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicIscsiBootPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AutoTargetvendorName) {
		toSerialize["AutoTargetvendorName"] = o.AutoTargetvendorName
	}
	if o.Chap.IsSet() {
		toSerialize["Chap"] = o.Chap.Get()
	}
	if !IsNil(o.InitiatorIpSource) {
		toSerialize["InitiatorIpSource"] = o.InitiatorIpSource
	}
	if !IsNil(o.InitiatorStaticIpV4Address) {
		toSerialize["InitiatorStaticIpV4Address"] = o.InitiatorStaticIpV4Address
	}
	if o.InitiatorStaticIpV4Config.IsSet() {
		toSerialize["InitiatorStaticIpV4Config"] = o.InitiatorStaticIpV4Config.Get()
	}
	if o.MutualChap.IsSet() {
		toSerialize["MutualChap"] = o.MutualChap.Get()
	}
	if !IsNil(o.TargetSourceType) {
		toSerialize["TargetSourceType"] = o.TargetSourceType
	}
	if o.InitiatorIpPool.IsSet() {
		toSerialize["InitiatorIpPool"] = o.InitiatorIpPool.Get()
	}
	if o.IscsiAdapterPolicy.IsSet() {
		toSerialize["IscsiAdapterPolicy"] = o.IscsiAdapterPolicy.Get()
	}
	if o.PrimaryTargetPolicy.IsSet() {
		toSerialize["PrimaryTargetPolicy"] = o.PrimaryTargetPolicy.Get()
	}
	if o.SecondaryTargetPolicy.IsSet() {
		toSerialize["SecondaryTargetPolicy"] = o.SecondaryTargetPolicy.Get()
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicIscsiBootPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type VnicIscsiBootPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters.
		AutoTargetvendorName *string                      `json:"AutoTargetvendorName,omitempty"`
		Chap                 NullableVnicIscsiAuthProfile `json:"Chap,omitempty"`
		// Source Type of Initiator IP Address - Auto/Static/Pool. * `DHCP` - The IP address is assigned using DHCP, if available. * `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.
		InitiatorIpSource *string `json:"InitiatorIpSource,omitempty"`
		// Static IP address provided for iSCSI Initiator.
		InitiatorStaticIpV4Address *string                      `json:"InitiatorStaticIpV4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		InitiatorStaticIpV4Config  NullableIppoolIpV4Config     `json:"InitiatorStaticIpV4Config,omitempty"`
		MutualChap                 NullableVnicIscsiAuthProfile `json:"MutualChap,omitempty"`
		// Source Type of Targets - Auto/Static. * `Static` - Type indicates that static target interface is assigned to iSCSI boot. * `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot.
		TargetSourceType      *string                                                  `json:"TargetSourceType,omitempty"`
		InitiatorIpPool       NullableIppoolPoolRelationship                           `json:"InitiatorIpPool,omitempty"`
		IscsiAdapterPolicy    NullableVnicIscsiAdapterPolicyInventoryRelationship      `json:"IscsiAdapterPolicy,omitempty"`
		PrimaryTargetPolicy   NullableVnicIscsiStaticTargetPolicyInventoryRelationship `json:"PrimaryTargetPolicy,omitempty"`
		SecondaryTargetPolicy NullableVnicIscsiStaticTargetPolicyInventoryRelationship `json:"SecondaryTargetPolicy,omitempty"`
		TargetMo              NullableMoBaseMoRelationship                             `json:"TargetMo,omitempty"`
	}

	varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct := VnicIscsiBootPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicIscsiBootPolicyInventory := _VnicIscsiBootPolicyInventory{}
		varVnicIscsiBootPolicyInventory.ClassId = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicIscsiBootPolicyInventory.ObjectType = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicIscsiBootPolicyInventory.AutoTargetvendorName = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.AutoTargetvendorName
		varVnicIscsiBootPolicyInventory.Chap = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.Chap
		varVnicIscsiBootPolicyInventory.InitiatorIpSource = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.InitiatorIpSource
		varVnicIscsiBootPolicyInventory.InitiatorStaticIpV4Address = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.InitiatorStaticIpV4Address
		varVnicIscsiBootPolicyInventory.InitiatorStaticIpV4Config = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.InitiatorStaticIpV4Config
		varVnicIscsiBootPolicyInventory.MutualChap = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.MutualChap
		varVnicIscsiBootPolicyInventory.TargetSourceType = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.TargetSourceType
		varVnicIscsiBootPolicyInventory.InitiatorIpPool = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.InitiatorIpPool
		varVnicIscsiBootPolicyInventory.IscsiAdapterPolicy = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.IscsiAdapterPolicy
		varVnicIscsiBootPolicyInventory.PrimaryTargetPolicy = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.PrimaryTargetPolicy
		varVnicIscsiBootPolicyInventory.SecondaryTargetPolicy = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.SecondaryTargetPolicy
		varVnicIscsiBootPolicyInventory.TargetMo = varVnicIscsiBootPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicIscsiBootPolicyInventory(varVnicIscsiBootPolicyInventory)
	} else {
		return err
	}

	varVnicIscsiBootPolicyInventory := _VnicIscsiBootPolicyInventory{}

	err = json.Unmarshal(data, &varVnicIscsiBootPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicIscsiBootPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoTargetvendorName")
		delete(additionalProperties, "Chap")
		delete(additionalProperties, "InitiatorIpSource")
		delete(additionalProperties, "InitiatorStaticIpV4Address")
		delete(additionalProperties, "InitiatorStaticIpV4Config")
		delete(additionalProperties, "MutualChap")
		delete(additionalProperties, "TargetSourceType")
		delete(additionalProperties, "InitiatorIpPool")
		delete(additionalProperties, "IscsiAdapterPolicy")
		delete(additionalProperties, "PrimaryTargetPolicy")
		delete(additionalProperties, "SecondaryTargetPolicy")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableVnicIscsiBootPolicyInventory struct {
	value *VnicIscsiBootPolicyInventory
	isSet bool
}

func (v NullableVnicIscsiBootPolicyInventory) Get() *VnicIscsiBootPolicyInventory {
	return v.value
}

func (v *NullableVnicIscsiBootPolicyInventory) Set(val *VnicIscsiBootPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiBootPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiBootPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiBootPolicyInventory(val *VnicIscsiBootPolicyInventory) *NullableVnicIscsiBootPolicyInventory {
	return &NullableVnicIscsiBootPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicIscsiBootPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiBootPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
