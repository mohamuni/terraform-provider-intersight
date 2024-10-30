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
	"time"
)

// checks if the OnpremClusterNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnpremClusterNode{}

// OnpremClusterNode ClusterNode encapsulates the properties of an Intersight Appliance node (virtual machine). ClusterNode and ClusterInfo data structures are shared between the upgrade services running in the Intersight and the Intersight Appliance.
type OnpremClusterNode struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Start time of the device connector in the Intersight Appliance node.
	BootTime *time.Time `json:"BootTime,omitempty"`
	// Mode of the appliance cluster.
	ClusterMode *string `json:"ClusterMode,omitempty"`
	// Number of CPUs configured for the Intersight Appliance node.
	CpuCount *int64 `json:"CpuCount,omitempty"`
	// Deployment type of the Intersight Appliance node.
	DeploymentType *string              `json:"DeploymentType,omitempty"`
	Disks          []OnpremResourceInfo `json:"Disks,omitempty"`
	// The hostname or FQDN of the Intersight Appliance node.
	Hostname *string                    `json:"Hostname,omitempty"`
	Memory   NullableOnpremResourceInfo `json:"Memory,omitempty"`
	// Identifier of the Intersight Appliance node (one based).
	NodeId         *int64   `json:"NodeId,omitempty"`
	PingErrorNodes []string `json:"PingErrorNodes,omitempty"`
	// Indicates if the node can ping other nodes in the Intersight Appliance cluster. The Ping operation is a high level application specific status check operation, not an ICMP ping between the hosts.
	PingOk *bool `json:"PingOk,omitempty"`
	// Indicates if this node is the primary node.
	PrimaryNode     *bool    `json:"PrimaryNode,omitempty"`
	RsyncErrorNodes []string `json:"RsyncErrorNodes,omitempty"`
	// Indicates if this node can rsync to other nodes.
	RsyncOk *bool `json:"RsyncOk,omitempty"`
	// Current version of the device connector in the Intersight Appliance node.
	Version *string `json:"Version,omitempty"`
	// Virtual Env type of the Intersight Appliance node (ESXi, Hyper-V or KVM).
	VirtualEnvType       *string `json:"VirtualEnvType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnpremClusterNode OnpremClusterNode

// NewOnpremClusterNode instantiates a new OnpremClusterNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremClusterNode(classId string, objectType string) *OnpremClusterNode {
	this := OnpremClusterNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOnpremClusterNodeWithDefaults instantiates a new OnpremClusterNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremClusterNodeWithDefaults() *OnpremClusterNode {
	this := OnpremClusterNode{}
	var classId string = "onprem.ClusterNode"
	this.ClassId = classId
	var objectType string = "onprem.ClusterNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OnpremClusterNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OnpremClusterNode) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "onprem.ClusterNode" of the ClassId field.
func (o *OnpremClusterNode) GetDefaultClassId() interface{} {
	return "onprem.ClusterNode"
}

// GetObjectType returns the ObjectType field value
func (o *OnpremClusterNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OnpremClusterNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "onprem.ClusterNode" of the ObjectType field.
func (o *OnpremClusterNode) GetDefaultObjectType() interface{} {
	return "onprem.ClusterNode"
}

// GetBootTime returns the BootTime field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetBootTime() time.Time {
	if o == nil || IsNil(o.BootTime) {
		var ret time.Time
		return ret
	}
	return *o.BootTime
}

// GetBootTimeOk returns a tuple with the BootTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetBootTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BootTime) {
		return nil, false
	}
	return o.BootTime, true
}

// HasBootTime returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasBootTime() bool {
	if o != nil && !IsNil(o.BootTime) {
		return true
	}

	return false
}

// SetBootTime gets a reference to the given time.Time and assigns it to the BootTime field.
func (o *OnpremClusterNode) SetBootTime(v time.Time) {
	o.BootTime = &v
}

// GetClusterMode returns the ClusterMode field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetClusterMode() string {
	if o == nil || IsNil(o.ClusterMode) {
		var ret string
		return ret
	}
	return *o.ClusterMode
}

// GetClusterModeOk returns a tuple with the ClusterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetClusterModeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterMode) {
		return nil, false
	}
	return o.ClusterMode, true
}

// HasClusterMode returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasClusterMode() bool {
	if o != nil && !IsNil(o.ClusterMode) {
		return true
	}

	return false
}

// SetClusterMode gets a reference to the given string and assigns it to the ClusterMode field.
func (o *OnpremClusterNode) SetClusterMode(v string) {
	o.ClusterMode = &v
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetCpuCount() int64 {
	if o == nil || IsNil(o.CpuCount) {
		var ret int64
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetCpuCountOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuCount) {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasCpuCount() bool {
	if o != nil && !IsNil(o.CpuCount) {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int64 and assigns it to the CpuCount field.
func (o *OnpremClusterNode) SetCpuCount(v int64) {
	o.CpuCount = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetDeploymentType() string {
	if o == nil || IsNil(o.DeploymentType) {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentType) {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasDeploymentType() bool {
	if o != nil && !IsNil(o.DeploymentType) {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *OnpremClusterNode) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnpremClusterNode) GetDisks() []OnpremResourceInfo {
	if o == nil {
		var ret []OnpremResourceInfo
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnpremClusterNode) GetDisksOk() ([]OnpremResourceInfo, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []OnpremResourceInfo and assigns it to the Disks field.
func (o *OnpremClusterNode) SetDisks(v []OnpremResourceInfo) {
	o.Disks = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *OnpremClusterNode) SetHostname(v string) {
	o.Hostname = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnpremClusterNode) GetMemory() OnpremResourceInfo {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret OnpremResourceInfo
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnpremClusterNode) GetMemoryOk() (*OnpremResourceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableOnpremResourceInfo and assigns it to the Memory field.
func (o *OnpremClusterNode) SetMemory(v OnpremResourceInfo) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *OnpremClusterNode) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *OnpremClusterNode) UnsetMemory() {
	o.Memory.Unset()
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetNodeId() int64 {
	if o == nil || IsNil(o.NodeId) {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetNodeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *OnpremClusterNode) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetPingErrorNodes returns the PingErrorNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnpremClusterNode) GetPingErrorNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PingErrorNodes
}

// GetPingErrorNodesOk returns a tuple with the PingErrorNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnpremClusterNode) GetPingErrorNodesOk() ([]string, bool) {
	if o == nil || IsNil(o.PingErrorNodes) {
		return nil, false
	}
	return o.PingErrorNodes, true
}

// HasPingErrorNodes returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasPingErrorNodes() bool {
	if o != nil && !IsNil(o.PingErrorNodes) {
		return true
	}

	return false
}

// SetPingErrorNodes gets a reference to the given []string and assigns it to the PingErrorNodes field.
func (o *OnpremClusterNode) SetPingErrorNodes(v []string) {
	o.PingErrorNodes = v
}

// GetPingOk returns the PingOk field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetPingOk() bool {
	if o == nil || IsNil(o.PingOk) {
		var ret bool
		return ret
	}
	return *o.PingOk
}

// GetPingOkOk returns a tuple with the PingOk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetPingOkOk() (*bool, bool) {
	if o == nil || IsNil(o.PingOk) {
		return nil, false
	}
	return o.PingOk, true
}

// HasPingOk returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasPingOk() bool {
	if o != nil && !IsNil(o.PingOk) {
		return true
	}

	return false
}

// SetPingOk gets a reference to the given bool and assigns it to the PingOk field.
func (o *OnpremClusterNode) SetPingOk(v bool) {
	o.PingOk = &v
}

// GetPrimaryNode returns the PrimaryNode field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetPrimaryNode() bool {
	if o == nil || IsNil(o.PrimaryNode) {
		var ret bool
		return ret
	}
	return *o.PrimaryNode
}

// GetPrimaryNodeOk returns a tuple with the PrimaryNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetPrimaryNodeOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryNode) {
		return nil, false
	}
	return o.PrimaryNode, true
}

// HasPrimaryNode returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasPrimaryNode() bool {
	if o != nil && !IsNil(o.PrimaryNode) {
		return true
	}

	return false
}

// SetPrimaryNode gets a reference to the given bool and assigns it to the PrimaryNode field.
func (o *OnpremClusterNode) SetPrimaryNode(v bool) {
	o.PrimaryNode = &v
}

// GetRsyncErrorNodes returns the RsyncErrorNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnpremClusterNode) GetRsyncErrorNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RsyncErrorNodes
}

// GetRsyncErrorNodesOk returns a tuple with the RsyncErrorNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnpremClusterNode) GetRsyncErrorNodesOk() ([]string, bool) {
	if o == nil || IsNil(o.RsyncErrorNodes) {
		return nil, false
	}
	return o.RsyncErrorNodes, true
}

// HasRsyncErrorNodes returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasRsyncErrorNodes() bool {
	if o != nil && !IsNil(o.RsyncErrorNodes) {
		return true
	}

	return false
}

// SetRsyncErrorNodes gets a reference to the given []string and assigns it to the RsyncErrorNodes field.
func (o *OnpremClusterNode) SetRsyncErrorNodes(v []string) {
	o.RsyncErrorNodes = v
}

// GetRsyncOk returns the RsyncOk field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetRsyncOk() bool {
	if o == nil || IsNil(o.RsyncOk) {
		var ret bool
		return ret
	}
	return *o.RsyncOk
}

// GetRsyncOkOk returns a tuple with the RsyncOk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetRsyncOkOk() (*bool, bool) {
	if o == nil || IsNil(o.RsyncOk) {
		return nil, false
	}
	return o.RsyncOk, true
}

// HasRsyncOk returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasRsyncOk() bool {
	if o != nil && !IsNil(o.RsyncOk) {
		return true
	}

	return false
}

// SetRsyncOk gets a reference to the given bool and assigns it to the RsyncOk field.
func (o *OnpremClusterNode) SetRsyncOk(v bool) {
	o.RsyncOk = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OnpremClusterNode) SetVersion(v string) {
	o.Version = &v
}

// GetVirtualEnvType returns the VirtualEnvType field value if set, zero value otherwise.
func (o *OnpremClusterNode) GetVirtualEnvType() string {
	if o == nil || IsNil(o.VirtualEnvType) {
		var ret string
		return ret
	}
	return *o.VirtualEnvType
}

// GetVirtualEnvTypeOk returns a tuple with the VirtualEnvType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremClusterNode) GetVirtualEnvTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualEnvType) {
		return nil, false
	}
	return o.VirtualEnvType, true
}

// HasVirtualEnvType returns a boolean if a field has been set.
func (o *OnpremClusterNode) HasVirtualEnvType() bool {
	if o != nil && !IsNil(o.VirtualEnvType) {
		return true
	}

	return false
}

// SetVirtualEnvType gets a reference to the given string and assigns it to the VirtualEnvType field.
func (o *OnpremClusterNode) SetVirtualEnvType(v string) {
	o.VirtualEnvType = &v
}

func (o OnpremClusterNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnpremClusterNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BootTime) {
		toSerialize["BootTime"] = o.BootTime
	}
	if !IsNil(o.ClusterMode) {
		toSerialize["ClusterMode"] = o.ClusterMode
	}
	if !IsNil(o.CpuCount) {
		toSerialize["CpuCount"] = o.CpuCount
	}
	if !IsNil(o.DeploymentType) {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}
	if !IsNil(o.Hostname) {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.Memory.IsSet() {
		toSerialize["Memory"] = o.Memory.Get()
	}
	if !IsNil(o.NodeId) {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.PingErrorNodes != nil {
		toSerialize["PingErrorNodes"] = o.PingErrorNodes
	}
	if !IsNil(o.PingOk) {
		toSerialize["PingOk"] = o.PingOk
	}
	if !IsNil(o.PrimaryNode) {
		toSerialize["PrimaryNode"] = o.PrimaryNode
	}
	if o.RsyncErrorNodes != nil {
		toSerialize["RsyncErrorNodes"] = o.RsyncErrorNodes
	}
	if !IsNil(o.RsyncOk) {
		toSerialize["RsyncOk"] = o.RsyncOk
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.VirtualEnvType) {
		toSerialize["VirtualEnvType"] = o.VirtualEnvType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnpremClusterNode) UnmarshalJSON(data []byte) (err error) {
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
	type OnpremClusterNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Start time of the device connector in the Intersight Appliance node.
		BootTime *time.Time `json:"BootTime,omitempty"`
		// Mode of the appliance cluster.
		ClusterMode *string `json:"ClusterMode,omitempty"`
		// Number of CPUs configured for the Intersight Appliance node.
		CpuCount *int64 `json:"CpuCount,omitempty"`
		// Deployment type of the Intersight Appliance node.
		DeploymentType *string              `json:"DeploymentType,omitempty"`
		Disks          []OnpremResourceInfo `json:"Disks,omitempty"`
		// The hostname or FQDN of the Intersight Appliance node.
		Hostname *string                    `json:"Hostname,omitempty"`
		Memory   NullableOnpremResourceInfo `json:"Memory,omitempty"`
		// Identifier of the Intersight Appliance node (one based).
		NodeId         *int64   `json:"NodeId,omitempty"`
		PingErrorNodes []string `json:"PingErrorNodes,omitempty"`
		// Indicates if the node can ping other nodes in the Intersight Appliance cluster. The Ping operation is a high level application specific status check operation, not an ICMP ping between the hosts.
		PingOk *bool `json:"PingOk,omitempty"`
		// Indicates if this node is the primary node.
		PrimaryNode     *bool    `json:"PrimaryNode,omitempty"`
		RsyncErrorNodes []string `json:"RsyncErrorNodes,omitempty"`
		// Indicates if this node can rsync to other nodes.
		RsyncOk *bool `json:"RsyncOk,omitempty"`
		// Current version of the device connector in the Intersight Appliance node.
		Version *string `json:"Version,omitempty"`
		// Virtual Env type of the Intersight Appliance node (ESXi, Hyper-V or KVM).
		VirtualEnvType *string `json:"VirtualEnvType,omitempty"`
	}

	varOnpremClusterNodeWithoutEmbeddedStruct := OnpremClusterNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOnpremClusterNodeWithoutEmbeddedStruct)
	if err == nil {
		varOnpremClusterNode := _OnpremClusterNode{}
		varOnpremClusterNode.ClassId = varOnpremClusterNodeWithoutEmbeddedStruct.ClassId
		varOnpremClusterNode.ObjectType = varOnpremClusterNodeWithoutEmbeddedStruct.ObjectType
		varOnpremClusterNode.BootTime = varOnpremClusterNodeWithoutEmbeddedStruct.BootTime
		varOnpremClusterNode.ClusterMode = varOnpremClusterNodeWithoutEmbeddedStruct.ClusterMode
		varOnpremClusterNode.CpuCount = varOnpremClusterNodeWithoutEmbeddedStruct.CpuCount
		varOnpremClusterNode.DeploymentType = varOnpremClusterNodeWithoutEmbeddedStruct.DeploymentType
		varOnpremClusterNode.Disks = varOnpremClusterNodeWithoutEmbeddedStruct.Disks
		varOnpremClusterNode.Hostname = varOnpremClusterNodeWithoutEmbeddedStruct.Hostname
		varOnpremClusterNode.Memory = varOnpremClusterNodeWithoutEmbeddedStruct.Memory
		varOnpremClusterNode.NodeId = varOnpremClusterNodeWithoutEmbeddedStruct.NodeId
		varOnpremClusterNode.PingErrorNodes = varOnpremClusterNodeWithoutEmbeddedStruct.PingErrorNodes
		varOnpremClusterNode.PingOk = varOnpremClusterNodeWithoutEmbeddedStruct.PingOk
		varOnpremClusterNode.PrimaryNode = varOnpremClusterNodeWithoutEmbeddedStruct.PrimaryNode
		varOnpremClusterNode.RsyncErrorNodes = varOnpremClusterNodeWithoutEmbeddedStruct.RsyncErrorNodes
		varOnpremClusterNode.RsyncOk = varOnpremClusterNodeWithoutEmbeddedStruct.RsyncOk
		varOnpremClusterNode.Version = varOnpremClusterNodeWithoutEmbeddedStruct.Version
		varOnpremClusterNode.VirtualEnvType = varOnpremClusterNodeWithoutEmbeddedStruct.VirtualEnvType
		*o = OnpremClusterNode(varOnpremClusterNode)
	} else {
		return err
	}

	varOnpremClusterNode := _OnpremClusterNode{}

	err = json.Unmarshal(data, &varOnpremClusterNode)
	if err == nil {
		o.MoBaseComplexType = varOnpremClusterNode.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootTime")
		delete(additionalProperties, "ClusterMode")
		delete(additionalProperties, "CpuCount")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "Disks")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "Memory")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "PingErrorNodes")
		delete(additionalProperties, "PingOk")
		delete(additionalProperties, "PrimaryNode")
		delete(additionalProperties, "RsyncErrorNodes")
		delete(additionalProperties, "RsyncOk")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "VirtualEnvType")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableOnpremClusterNode struct {
	value *OnpremClusterNode
	isSet bool
}

func (v NullableOnpremClusterNode) Get() *OnpremClusterNode {
	return v.value
}

func (v *NullableOnpremClusterNode) Set(val *OnpremClusterNode) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremClusterNode) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremClusterNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremClusterNode(val *OnpremClusterNode) *NullableOnpremClusterNode {
	return &NullableOnpremClusterNode{value: val, isSet: true}
}

func (v NullableOnpremClusterNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremClusterNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
