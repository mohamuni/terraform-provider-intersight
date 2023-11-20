/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ApplianceNodeOpStatusAllOf Definition of the list of properties defined in 'appliance.NodeOpStatus', excluding properties defined in parent classes.
type ApplianceNodeOpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	ClusterNetworkStatus []ApplianceNetworkStatus `json:"ClusterNetworkStatus,omitempty"`
	// Percentage of CPU currently in use.
	CpuUsage *float32 `json:"CpuUsage,omitempty"`
	// Percentage of memory currently in use.
	MemUsage *float32 `json:"MemUsage,omitempty"`
	// Hostname of the Intersight Appliance node.
	NodeHostname *string `json:"NodeHostname,omitempty"`
	// System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
	NodeId *int64 `json:"NodeId,omitempty"`
	// State of the node in terms of its readiness to host Kubernetes pods. * `Down` - The node is yet to come up and join as a member of theKubernetes cluster. * `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods. * `Ready` - The node is ready to host Kubernetes pods.
	NodeState *string `json:"NodeState,omitempty"`
	// Operational status of the Intersight Appliance node. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
	OperationalStatus *string                 `json:"OperationalStatus,omitempty"`
	Account           *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to applianceFileSystemOpStatus resources.
	FileSystemOpStatuses []ApplianceFileSystemOpStatusRelationship `json:"FileSystemOpStatuses,omitempty"`
	// An array of relationships to applianceNetworkLinkStatus resources.
	NetworkLinkStatuses  []ApplianceNetworkLinkStatusRelationship `json:"NetworkLinkStatuses,omitempty"`
	NodeInfo             *ApplianceNodeInfoRelationship           `json:"NodeInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty"`
	SystemOpStatus       *ApplianceSystemOpStatusRelationship     `json:"SystemOpStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNodeOpStatusAllOf ApplianceNodeOpStatusAllOf

// NewApplianceNodeOpStatusAllOf instantiates a new ApplianceNodeOpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNodeOpStatusAllOf(classId string, objectType string) *ApplianceNodeOpStatusAllOf {
	this := ApplianceNodeOpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNodeOpStatusAllOfWithDefaults instantiates a new ApplianceNodeOpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNodeOpStatusAllOfWithDefaults() *ApplianceNodeOpStatusAllOf {
	this := ApplianceNodeOpStatusAllOf{}
	var classId string = "appliance.NodeOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.NodeOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNodeOpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNodeOpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNodeOpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNodeOpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterNetworkStatus returns the ClusterNetworkStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeOpStatusAllOf) GetClusterNetworkStatus() []ApplianceNetworkStatus {
	if o == nil {
		var ret []ApplianceNetworkStatus
		return ret
	}
	return o.ClusterNetworkStatus
}

// GetClusterNetworkStatusOk returns a tuple with the ClusterNetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeOpStatusAllOf) GetClusterNetworkStatusOk() ([]ApplianceNetworkStatus, bool) {
	if o == nil || o.ClusterNetworkStatus == nil {
		return nil, false
	}
	return o.ClusterNetworkStatus, true
}

// HasClusterNetworkStatus returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasClusterNetworkStatus() bool {
	if o != nil && o.ClusterNetworkStatus != nil {
		return true
	}

	return false
}

// SetClusterNetworkStatus gets a reference to the given []ApplianceNetworkStatus and assigns it to the ClusterNetworkStatus field.
func (o *ApplianceNodeOpStatusAllOf) SetClusterNetworkStatus(v []ApplianceNetworkStatus) {
	o.ClusterNetworkStatus = v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetCpuUsage() float32 {
	if o == nil || o.CpuUsage == nil {
		var ret float32
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetCpuUsageOk() (*float32, bool) {
	if o == nil || o.CpuUsage == nil {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasCpuUsage() bool {
	if o != nil && o.CpuUsage != nil {
		return true
	}

	return false
}

// SetCpuUsage gets a reference to the given float32 and assigns it to the CpuUsage field.
func (o *ApplianceNodeOpStatusAllOf) SetCpuUsage(v float32) {
	o.CpuUsage = &v
}

// GetMemUsage returns the MemUsage field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetMemUsage() float32 {
	if o == nil || o.MemUsage == nil {
		var ret float32
		return ret
	}
	return *o.MemUsage
}

// GetMemUsageOk returns a tuple with the MemUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetMemUsageOk() (*float32, bool) {
	if o == nil || o.MemUsage == nil {
		return nil, false
	}
	return o.MemUsage, true
}

// HasMemUsage returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasMemUsage() bool {
	if o != nil && o.MemUsage != nil {
		return true
	}

	return false
}

// SetMemUsage gets a reference to the given float32 and assigns it to the MemUsage field.
func (o *ApplianceNodeOpStatusAllOf) SetMemUsage(v float32) {
	o.MemUsage = &v
}

// GetNodeHostname returns the NodeHostname field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetNodeHostname() string {
	if o == nil || o.NodeHostname == nil {
		var ret string
		return ret
	}
	return *o.NodeHostname
}

// GetNodeHostnameOk returns a tuple with the NodeHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetNodeHostnameOk() (*string, bool) {
	if o == nil || o.NodeHostname == nil {
		return nil, false
	}
	return o.NodeHostname, true
}

// HasNodeHostname returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasNodeHostname() bool {
	if o != nil && o.NodeHostname != nil {
		return true
	}

	return false
}

// SetNodeHostname gets a reference to the given string and assigns it to the NodeHostname field.
func (o *ApplianceNodeOpStatusAllOf) SetNodeHostname(v string) {
	o.NodeHostname = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceNodeOpStatusAllOf) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeState returns the NodeState field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetNodeState() string {
	if o == nil || o.NodeState == nil {
		var ret string
		return ret
	}
	return *o.NodeState
}

// GetNodeStateOk returns a tuple with the NodeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetNodeStateOk() (*string, bool) {
	if o == nil || o.NodeState == nil {
		return nil, false
	}
	return o.NodeState, true
}

// HasNodeState returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasNodeState() bool {
	if o != nil && o.NodeState != nil {
		return true
	}

	return false
}

// SetNodeState gets a reference to the given string and assigns it to the NodeState field.
func (o *ApplianceNodeOpStatusAllOf) SetNodeState(v string) {
	o.NodeState = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceNodeOpStatusAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceNodeOpStatusAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetFileSystemOpStatuses returns the FileSystemOpStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeOpStatusAllOf) GetFileSystemOpStatuses() []ApplianceFileSystemOpStatusRelationship {
	if o == nil {
		var ret []ApplianceFileSystemOpStatusRelationship
		return ret
	}
	return o.FileSystemOpStatuses
}

// GetFileSystemOpStatusesOk returns a tuple with the FileSystemOpStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeOpStatusAllOf) GetFileSystemOpStatusesOk() ([]ApplianceFileSystemOpStatusRelationship, bool) {
	if o == nil || o.FileSystemOpStatuses == nil {
		return nil, false
	}
	return o.FileSystemOpStatuses, true
}

// HasFileSystemOpStatuses returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasFileSystemOpStatuses() bool {
	if o != nil && o.FileSystemOpStatuses != nil {
		return true
	}

	return false
}

// SetFileSystemOpStatuses gets a reference to the given []ApplianceFileSystemOpStatusRelationship and assigns it to the FileSystemOpStatuses field.
func (o *ApplianceNodeOpStatusAllOf) SetFileSystemOpStatuses(v []ApplianceFileSystemOpStatusRelationship) {
	o.FileSystemOpStatuses = v
}

// GetNetworkLinkStatuses returns the NetworkLinkStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeOpStatusAllOf) GetNetworkLinkStatuses() []ApplianceNetworkLinkStatusRelationship {
	if o == nil {
		var ret []ApplianceNetworkLinkStatusRelationship
		return ret
	}
	return o.NetworkLinkStatuses
}

// GetNetworkLinkStatusesOk returns a tuple with the NetworkLinkStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeOpStatusAllOf) GetNetworkLinkStatusesOk() ([]ApplianceNetworkLinkStatusRelationship, bool) {
	if o == nil || o.NetworkLinkStatuses == nil {
		return nil, false
	}
	return o.NetworkLinkStatuses, true
}

// HasNetworkLinkStatuses returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasNetworkLinkStatuses() bool {
	if o != nil && o.NetworkLinkStatuses != nil {
		return true
	}

	return false
}

// SetNetworkLinkStatuses gets a reference to the given []ApplianceNetworkLinkStatusRelationship and assigns it to the NetworkLinkStatuses field.
func (o *ApplianceNodeOpStatusAllOf) SetNetworkLinkStatuses(v []ApplianceNetworkLinkStatusRelationship) {
	o.NetworkLinkStatuses = v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetNodeInfo() ApplianceNodeInfoRelationship {
	if o == nil || o.NodeInfo == nil {
		var ret ApplianceNodeInfoRelationship
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetNodeInfoOk() (*ApplianceNodeInfoRelationship, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given ApplianceNodeInfoRelationship and assigns it to the NodeInfo field.
func (o *ApplianceNodeOpStatusAllOf) SetNodeInfo(v ApplianceNodeInfoRelationship) {
	o.NodeInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceNodeOpStatusAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSystemOpStatus returns the SystemOpStatus field value if set, zero value otherwise.
func (o *ApplianceNodeOpStatusAllOf) GetSystemOpStatus() ApplianceSystemOpStatusRelationship {
	if o == nil || o.SystemOpStatus == nil {
		var ret ApplianceSystemOpStatusRelationship
		return ret
	}
	return *o.SystemOpStatus
}

// GetSystemOpStatusOk returns a tuple with the SystemOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeOpStatusAllOf) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool) {
	if o == nil || o.SystemOpStatus == nil {
		return nil, false
	}
	return o.SystemOpStatus, true
}

// HasSystemOpStatus returns a boolean if a field has been set.
func (o *ApplianceNodeOpStatusAllOf) HasSystemOpStatus() bool {
	if o != nil && o.SystemOpStatus != nil {
		return true
	}

	return false
}

// SetSystemOpStatus gets a reference to the given ApplianceSystemOpStatusRelationship and assigns it to the SystemOpStatus field.
func (o *ApplianceNodeOpStatusAllOf) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship) {
	o.SystemOpStatus = &v
}

func (o ApplianceNodeOpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterNetworkStatus != nil {
		toSerialize["ClusterNetworkStatus"] = o.ClusterNetworkStatus
	}
	if o.CpuUsage != nil {
		toSerialize["CpuUsage"] = o.CpuUsage
	}
	if o.MemUsage != nil {
		toSerialize["MemUsage"] = o.MemUsage
	}
	if o.NodeHostname != nil {
		toSerialize["NodeHostname"] = o.NodeHostname
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.NodeState != nil {
		toSerialize["NodeState"] = o.NodeState
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.FileSystemOpStatuses != nil {
		toSerialize["FileSystemOpStatuses"] = o.FileSystemOpStatuses
	}
	if o.NetworkLinkStatuses != nil {
		toSerialize["NetworkLinkStatuses"] = o.NetworkLinkStatuses
	}
	if o.NodeInfo != nil {
		toSerialize["NodeInfo"] = o.NodeInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SystemOpStatus != nil {
		toSerialize["SystemOpStatus"] = o.SystemOpStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceNodeOpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceNodeOpStatusAllOf := _ApplianceNodeOpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceNodeOpStatusAllOf); err == nil {
		*o = ApplianceNodeOpStatusAllOf(varApplianceNodeOpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterNetworkStatus")
		delete(additionalProperties, "CpuUsage")
		delete(additionalProperties, "MemUsage")
		delete(additionalProperties, "NodeHostname")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "NodeState")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "FileSystemOpStatuses")
		delete(additionalProperties, "NetworkLinkStatuses")
		delete(additionalProperties, "NodeInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SystemOpStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceNodeOpStatusAllOf struct {
	value *ApplianceNodeOpStatusAllOf
	isSet bool
}

func (v NullableApplianceNodeOpStatusAllOf) Get() *ApplianceNodeOpStatusAllOf {
	return v.value
}

func (v *NullableApplianceNodeOpStatusAllOf) Set(val *ApplianceNodeOpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNodeOpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNodeOpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNodeOpStatusAllOf(val *ApplianceNodeOpStatusAllOf) *NullableApplianceNodeOpStatusAllOf {
	return &NullableApplianceNodeOpStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceNodeOpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNodeOpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
