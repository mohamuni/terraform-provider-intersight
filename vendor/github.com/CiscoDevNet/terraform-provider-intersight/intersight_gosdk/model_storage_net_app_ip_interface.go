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

// checks if the StorageNetAppIpInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppIpInterface{}

// StorageNetAppIpInterface NetApp IP interface is a logical interface.
type StorageNetAppIpInterface struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP interface is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// Name of home node of IP interface.
	HomeNode *string `json:"HomeNode,omitempty"`
	// Name of home port of IP interface.
	HomePort *string `json:"HomePort,omitempty"`
	// The IP address of interface.
	IpAddress *string `json:"IpAddress,omitempty"`
	// IP address family of interface. * `IPv4` - IP address family type is IPv4. * `IPv6` - IP address family type is IP6.
	IpFamily *string `json:"IpFamily,omitempty"`
	// The name of the IP interface.
	Name *string `json:"Name,omitempty"`
	// The netmask of the interface.
	Netmask *string `json:"Netmask,omitempty"`
	// Service policy name of IP interface.
	ServicePolicyName *string `json:"ServicePolicyName,omitempty"`
	// Service policy UUID of IP interface.
	ServicePolicyUuid *string  `json:"ServicePolicyUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	Services          []string `json:"Services,omitempty"`
	// The state of the IP interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Uuid of NetApp IP Interface.
	Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppIpInterfaceEvent resources.
	Events               []StorageNetAppIpInterfaceEventRelationship   `json:"Events,omitempty"`
	NetAppEthernetPort   NullableStorageNetAppEthernetPortRelationship `json:"NetAppEthernetPort,omitempty"`
	Tenant               NullableStorageNetAppStorageVmRelationship    `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppIpInterface StorageNetAppIpInterface

// NewStorageNetAppIpInterface instantiates a new StorageNetAppIpInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppIpInterface(classId string, objectType string) *StorageNetAppIpInterface {
	this := StorageNetAppIpInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppIpInterfaceWithDefaults instantiates a new StorageNetAppIpInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppIpInterfaceWithDefaults() *StorageNetAppIpInterface {
	this := StorageNetAppIpInterface{}
	var classId string = "storage.NetAppIpInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppIpInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppIpInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppIpInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppIpInterface" of the ClassId field.
func (o *StorageNetAppIpInterface) GetDefaultClassId() interface{} {
	return "storage.NetAppIpInterface"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppIpInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppIpInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppIpInterface" of the ObjectType field.
func (o *StorageNetAppIpInterface) GetDefaultObjectType() interface{} {
	return "storage.NetAppIpInterface"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetEnabled() string {
	if o == nil || IsNil(o.Enabled) {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppIpInterface) SetEnabled(v string) {
	o.Enabled = &v
}

// GetHomeNode returns the HomeNode field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetHomeNode() string {
	if o == nil || IsNil(o.HomeNode) {
		var ret string
		return ret
	}
	return *o.HomeNode
}

// GetHomeNodeOk returns a tuple with the HomeNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetHomeNodeOk() (*string, bool) {
	if o == nil || IsNil(o.HomeNode) {
		return nil, false
	}
	return o.HomeNode, true
}

// HasHomeNode returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasHomeNode() bool {
	if o != nil && !IsNil(o.HomeNode) {
		return true
	}

	return false
}

// SetHomeNode gets a reference to the given string and assigns it to the HomeNode field.
func (o *StorageNetAppIpInterface) SetHomeNode(v string) {
	o.HomeNode = &v
}

// GetHomePort returns the HomePort field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetHomePort() string {
	if o == nil || IsNil(o.HomePort) {
		var ret string
		return ret
	}
	return *o.HomePort
}

// GetHomePortOk returns a tuple with the HomePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetHomePortOk() (*string, bool) {
	if o == nil || IsNil(o.HomePort) {
		return nil, false
	}
	return o.HomePort, true
}

// HasHomePort returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasHomePort() bool {
	if o != nil && !IsNil(o.HomePort) {
		return true
	}

	return false
}

// SetHomePort gets a reference to the given string and assigns it to the HomePort field.
func (o *StorageNetAppIpInterface) SetHomePort(v string) {
	o.HomePort = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *StorageNetAppIpInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpFamily returns the IpFamily field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetIpFamily() string {
	if o == nil || IsNil(o.IpFamily) {
		var ret string
		return ret
	}
	return *o.IpFamily
}

// GetIpFamilyOk returns a tuple with the IpFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetIpFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.IpFamily) {
		return nil, false
	}
	return o.IpFamily, true
}

// HasIpFamily returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasIpFamily() bool {
	if o != nil && !IsNil(o.IpFamily) {
		return true
	}

	return false
}

// SetIpFamily gets a reference to the given string and assigns it to the IpFamily field.
func (o *StorageNetAppIpInterface) SetIpFamily(v string) {
	o.IpFamily = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppIpInterface) SetName(v string) {
	o.Name = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *StorageNetAppIpInterface) SetNetmask(v string) {
	o.Netmask = &v
}

// GetServicePolicyName returns the ServicePolicyName field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetServicePolicyName() string {
	if o == nil || IsNil(o.ServicePolicyName) {
		var ret string
		return ret
	}
	return *o.ServicePolicyName
}

// GetServicePolicyNameOk returns a tuple with the ServicePolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetServicePolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePolicyName) {
		return nil, false
	}
	return o.ServicePolicyName, true
}

// HasServicePolicyName returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasServicePolicyName() bool {
	if o != nil && !IsNil(o.ServicePolicyName) {
		return true
	}

	return false
}

// SetServicePolicyName gets a reference to the given string and assigns it to the ServicePolicyName field.
func (o *StorageNetAppIpInterface) SetServicePolicyName(v string) {
	o.ServicePolicyName = &v
}

// GetServicePolicyUuid returns the ServicePolicyUuid field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetServicePolicyUuid() string {
	if o == nil || IsNil(o.ServicePolicyUuid) {
		var ret string
		return ret
	}
	return *o.ServicePolicyUuid
}

// GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetServicePolicyUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePolicyUuid) {
		return nil, false
	}
	return o.ServicePolicyUuid, true
}

// HasServicePolicyUuid returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasServicePolicyUuid() bool {
	if o != nil && !IsNil(o.ServicePolicyUuid) {
		return true
	}

	return false
}

// SetServicePolicyUuid gets a reference to the given string and assigns it to the ServicePolicyUuid field.
func (o *StorageNetAppIpInterface) SetServicePolicyUuid(v string) {
	o.ServicePolicyUuid = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *StorageNetAppIpInterface) SetServices(v []string) {
	o.Services = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppIpInterface) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppIpInterface) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || IsNil(o.ArrayController.Get()) {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController.Get()
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayController.Get(), o.ArrayController.IsSet()
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasArrayController() bool {
	if o != nil && o.ArrayController.IsSet() {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given NullableStorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppIpInterface) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController.Set(&v)
}

// SetArrayControllerNil sets the value for ArrayController to be an explicit nil
func (o *StorageNetAppIpInterface) SetArrayControllerNil() {
	o.ArrayController.Set(nil)
}

// UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
func (o *StorageNetAppIpInterface) UnsetArrayController() {
	o.ArrayController.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetEvents() []StorageNetAppIpInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppIpInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetEventsOk() ([]StorageNetAppIpInterfaceEventRelationship, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppIpInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppIpInterface) SetEvents(v []StorageNetAppIpInterfaceEventRelationship) {
	o.Events = v
}

// GetNetAppEthernetPort returns the NetAppEthernetPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || IsNil(o.NetAppEthernetPort.Get()) {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.NetAppEthernetPort.Get()
}

// GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetAppEthernetPort.Get(), o.NetAppEthernetPort.IsSet()
}

// HasNetAppEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasNetAppEthernetPort() bool {
	if o != nil && o.NetAppEthernetPort.IsSet() {
		return true
	}

	return false
}

// SetNetAppEthernetPort gets a reference to the given NullableStorageNetAppEthernetPortRelationship and assigns it to the NetAppEthernetPort field.
func (o *StorageNetAppIpInterface) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.NetAppEthernetPort.Set(&v)
}

// SetNetAppEthernetPortNil sets the value for NetAppEthernetPort to be an explicit nil
func (o *StorageNetAppIpInterface) SetNetAppEthernetPortNil() {
	o.NetAppEthernetPort.Set(nil)
}

// UnsetNetAppEthernetPort ensures that no value is present for NetAppEthernetPort, not even an explicit nil
func (o *StorageNetAppIpInterface) UnsetNetAppEthernetPort() {
	o.NetAppEthernetPort.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableStorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppIpInterface) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *StorageNetAppIpInterface) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *StorageNetAppIpInterface) UnsetTenant() {
	o.Tenant.Unset()
}

func (o StorageNetAppIpInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppIpInterface) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.HomeNode) {
		toSerialize["HomeNode"] = o.HomeNode
	}
	if !IsNil(o.HomePort) {
		toSerialize["HomePort"] = o.HomePort
	}
	if !IsNil(o.IpAddress) {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if !IsNil(o.IpFamily) {
		toSerialize["IpFamily"] = o.IpFamily
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Netmask) {
		toSerialize["Netmask"] = o.Netmask
	}
	if !IsNil(o.ServicePolicyName) {
		toSerialize["ServicePolicyName"] = o.ServicePolicyName
	}
	if !IsNil(o.ServicePolicyUuid) {
		toSerialize["ServicePolicyUuid"] = o.ServicePolicyUuid
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
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
	if o.NetAppEthernetPort.IsSet() {
		toSerialize["NetAppEthernetPort"] = o.NetAppEthernetPort.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["Tenant"] = o.Tenant.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppIpInterface) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppIpInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP interface is enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// Name of home node of IP interface.
		HomeNode *string `json:"HomeNode,omitempty"`
		// Name of home port of IP interface.
		HomePort *string `json:"HomePort,omitempty"`
		// The IP address of interface.
		IpAddress *string `json:"IpAddress,omitempty"`
		// IP address family of interface. * `IPv4` - IP address family type is IPv4. * `IPv6` - IP address family type is IP6.
		IpFamily *string `json:"IpFamily,omitempty"`
		// The name of the IP interface.
		Name *string `json:"Name,omitempty"`
		// The netmask of the interface.
		Netmask *string `json:"Netmask,omitempty"`
		// Service policy name of IP interface.
		ServicePolicyName *string `json:"ServicePolicyName,omitempty"`
		// Service policy UUID of IP interface.
		ServicePolicyUuid *string  `json:"ServicePolicyUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		Services          []string `json:"Services,omitempty"`
		// The state of the IP interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		State *string `json:"State,omitempty"`
		// Uuid of NetApp IP Interface.
		Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppIpInterfaceEvent resources.
		Events             []StorageNetAppIpInterfaceEventRelationship   `json:"Events,omitempty"`
		NetAppEthernetPort NullableStorageNetAppEthernetPortRelationship `json:"NetAppEthernetPort,omitempty"`
		Tenant             NullableStorageNetAppStorageVmRelationship    `json:"Tenant,omitempty"`
	}

	varStorageNetAppIpInterfaceWithoutEmbeddedStruct := StorageNetAppIpInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppIpInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppIpInterface := _StorageNetAppIpInterface{}
		varStorageNetAppIpInterface.ClassId = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppIpInterface.ObjectType = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppIpInterface.Enabled = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Enabled
		varStorageNetAppIpInterface.HomeNode = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.HomeNode
		varStorageNetAppIpInterface.HomePort = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.HomePort
		varStorageNetAppIpInterface.IpAddress = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.IpAddress
		varStorageNetAppIpInterface.IpFamily = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.IpFamily
		varStorageNetAppIpInterface.Name = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Name
		varStorageNetAppIpInterface.Netmask = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Netmask
		varStorageNetAppIpInterface.ServicePolicyName = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ServicePolicyName
		varStorageNetAppIpInterface.ServicePolicyUuid = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ServicePolicyUuid
		varStorageNetAppIpInterface.Services = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Services
		varStorageNetAppIpInterface.State = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.State
		varStorageNetAppIpInterface.Uuid = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Uuid
		varStorageNetAppIpInterface.ArrayController = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ArrayController
		varStorageNetAppIpInterface.Events = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Events
		varStorageNetAppIpInterface.NetAppEthernetPort = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.NetAppEthernetPort
		varStorageNetAppIpInterface.Tenant = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppIpInterface(varStorageNetAppIpInterface)
	} else {
		return err
	}

	varStorageNetAppIpInterface := _StorageNetAppIpInterface{}

	err = json.Unmarshal(data, &varStorageNetAppIpInterface)
	if err == nil {
		o.MoBaseMo = varStorageNetAppIpInterface.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "HomeNode")
		delete(additionalProperties, "HomePort")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "IpFamily")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "ServicePolicyName")
		delete(additionalProperties, "ServicePolicyUuid")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "NetAppEthernetPort")
		delete(additionalProperties, "Tenant")

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

type NullableStorageNetAppIpInterface struct {
	value *StorageNetAppIpInterface
	isSet bool
}

func (v NullableStorageNetAppIpInterface) Get() *StorageNetAppIpInterface {
	return v.value
}

func (v *NullableStorageNetAppIpInterface) Set(val *StorageNetAppIpInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppIpInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppIpInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppIpInterface(val *StorageNetAppIpInterface) *NullableStorageNetAppIpInterface {
	return &NullableStorageNetAppIpInterface{value: val, isSet: true}
}

func (v NullableStorageNetAppIpInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppIpInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
