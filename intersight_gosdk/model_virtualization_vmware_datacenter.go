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

// checks if the VirtualizationVmwareDatacenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareDatacenter{}

// VirtualizationVmwareDatacenter Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
type VirtualizationVmwareDatacenter struct {
	VirtualizationBaseDatacenter
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Count of all clusters associated with this DC.
	ClusterCount *int64 `json:"ClusterCount,omitempty"`
	// Count of all datastores associated with this DC.
	DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
	// Count of all distributed networks associated with this datacenter (DC).
	DistributedNetworkCount *int64 `json:"DistributedNetworkCount,omitempty"`
	// Count of all distributed virtual switches associated with this datacenter (DC).
	DistributedVirtualSwitchCount *int64 `json:"DistributedVirtualSwitchCount,omitempty"`
	// Count of all hosts associated with this DC.
	HostCount *int64 `json:"HostCount,omitempty"`
	// Inventory path of the DC.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Count of all networks associated with this datacenter (DC).
	NetworkCount *int64 `json:"NetworkCount,omitempty"`
	// Count of all standard networks associated with this datacenter (DC).
	StandardNetworkCount *int64 `json:"StandardNetworkCount,omitempty"`
	// Count of all virtual machines (VMs) associated with this DC.
	VmCount *int64 `json:"VmCount,omitempty"`
	// Count of all virtual machines templates associated with this DC.
	VmTemplateCount      *int64                                          `json:"VmTemplateCount,omitempty"`
	HypervisorManager    NullableVirtualizationVmwareVcenterRelationship `json:"HypervisorManager,omitempty"`
	ParentFolder         NullableVirtualizationVmwareFolderRelationship  `json:"ParentFolder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDatacenter VirtualizationVmwareDatacenter

// NewVirtualizationVmwareDatacenter instantiates a new VirtualizationVmwareDatacenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDatacenter(classId string, objectType string) *VirtualizationVmwareDatacenter {
	this := VirtualizationVmwareDatacenter{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareDatacenterWithDefaults instantiates a new VirtualizationVmwareDatacenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDatacenterWithDefaults() *VirtualizationVmwareDatacenter {
	this := VirtualizationVmwareDatacenter{}
	var classId string = "virtualization.VmwareDatacenter"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatacenter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDatacenter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDatacenter) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareDatacenter" of the ClassId field.
func (o *VirtualizationVmwareDatacenter) GetDefaultClassId() interface{} {
	return "virtualization.VmwareDatacenter"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDatacenter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDatacenter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareDatacenter" of the ObjectType field.
func (o *VirtualizationVmwareDatacenter) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareDatacenter"
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetClusterCount() int64 {
	if o == nil || IsNil(o.ClusterCount) {
		var ret int64
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetClusterCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ClusterCount) {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasClusterCount() bool {
	if o != nil && !IsNil(o.ClusterCount) {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int64 and assigns it to the ClusterCount field.
func (o *VirtualizationVmwareDatacenter) SetClusterCount(v int64) {
	o.ClusterCount = &v
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetDatastoreCount() int64 {
	if o == nil || IsNil(o.DatastoreCount) {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DatastoreCount) {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasDatastoreCount() bool {
	if o != nil && !IsNil(o.DatastoreCount) {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareDatacenter) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetDistributedNetworkCount returns the DistributedNetworkCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetDistributedNetworkCount() int64 {
	if o == nil || IsNil(o.DistributedNetworkCount) {
		var ret int64
		return ret
	}
	return *o.DistributedNetworkCount
}

// GetDistributedNetworkCountOk returns a tuple with the DistributedNetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetDistributedNetworkCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DistributedNetworkCount) {
		return nil, false
	}
	return o.DistributedNetworkCount, true
}

// HasDistributedNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasDistributedNetworkCount() bool {
	if o != nil && !IsNil(o.DistributedNetworkCount) {
		return true
	}

	return false
}

// SetDistributedNetworkCount gets a reference to the given int64 and assigns it to the DistributedNetworkCount field.
func (o *VirtualizationVmwareDatacenter) SetDistributedNetworkCount(v int64) {
	o.DistributedNetworkCount = &v
}

// GetDistributedVirtualSwitchCount returns the DistributedVirtualSwitchCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetDistributedVirtualSwitchCount() int64 {
	if o == nil || IsNil(o.DistributedVirtualSwitchCount) {
		var ret int64
		return ret
	}
	return *o.DistributedVirtualSwitchCount
}

// GetDistributedVirtualSwitchCountOk returns a tuple with the DistributedVirtualSwitchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetDistributedVirtualSwitchCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DistributedVirtualSwitchCount) {
		return nil, false
	}
	return o.DistributedVirtualSwitchCount, true
}

// HasDistributedVirtualSwitchCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasDistributedVirtualSwitchCount() bool {
	if o != nil && !IsNil(o.DistributedVirtualSwitchCount) {
		return true
	}

	return false
}

// SetDistributedVirtualSwitchCount gets a reference to the given int64 and assigns it to the DistributedVirtualSwitchCount field.
func (o *VirtualizationVmwareDatacenter) SetDistributedVirtualSwitchCount(v int64) {
	o.DistributedVirtualSwitchCount = &v
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetHostCount() int64 {
	if o == nil || IsNil(o.HostCount) {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetHostCountOk() (*int64, bool) {
	if o == nil || IsNil(o.HostCount) {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasHostCount() bool {
	if o != nil && !IsNil(o.HostCount) {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationVmwareDatacenter) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetInventoryPath() string {
	if o == nil || IsNil(o.InventoryPath) {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetInventoryPathOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryPath) {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasInventoryPath() bool {
	if o != nil && !IsNil(o.InventoryPath) {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareDatacenter) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetNetworkCount returns the NetworkCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetNetworkCount() int64 {
	if o == nil || IsNil(o.NetworkCount) {
		var ret int64
		return ret
	}
	return *o.NetworkCount
}

// GetNetworkCountOk returns a tuple with the NetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetNetworkCountOk() (*int64, bool) {
	if o == nil || IsNil(o.NetworkCount) {
		return nil, false
	}
	return o.NetworkCount, true
}

// HasNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasNetworkCount() bool {
	if o != nil && !IsNil(o.NetworkCount) {
		return true
	}

	return false
}

// SetNetworkCount gets a reference to the given int64 and assigns it to the NetworkCount field.
func (o *VirtualizationVmwareDatacenter) SetNetworkCount(v int64) {
	o.NetworkCount = &v
}

// GetStandardNetworkCount returns the StandardNetworkCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetStandardNetworkCount() int64 {
	if o == nil || IsNil(o.StandardNetworkCount) {
		var ret int64
		return ret
	}
	return *o.StandardNetworkCount
}

// GetStandardNetworkCountOk returns a tuple with the StandardNetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetStandardNetworkCountOk() (*int64, bool) {
	if o == nil || IsNil(o.StandardNetworkCount) {
		return nil, false
	}
	return o.StandardNetworkCount, true
}

// HasStandardNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasStandardNetworkCount() bool {
	if o != nil && !IsNil(o.StandardNetworkCount) {
		return true
	}

	return false
}

// SetStandardNetworkCount gets a reference to the given int64 and assigns it to the StandardNetworkCount field.
func (o *VirtualizationVmwareDatacenter) SetStandardNetworkCount(v int64) {
	o.StandardNetworkCount = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetVmCount() int64 {
	if o == nil || IsNil(o.VmCount) {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetVmCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VmCount) {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasVmCount() bool {
	if o != nil && !IsNil(o.VmCount) {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationVmwareDatacenter) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetVmTemplateCount returns the VmTemplateCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetVmTemplateCount() int64 {
	if o == nil || IsNil(o.VmTemplateCount) {
		var ret int64
		return ret
	}
	return *o.VmTemplateCount
}

// GetVmTemplateCountOk returns a tuple with the VmTemplateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetVmTemplateCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VmTemplateCount) {
		return nil, false
	}
	return o.VmTemplateCount, true
}

// HasVmTemplateCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasVmTemplateCount() bool {
	if o != nil && !IsNil(o.VmTemplateCount) {
		return true
	}

	return false
}

// SetVmTemplateCount gets a reference to the given int64 and assigns it to the VmTemplateCount field.
func (o *VirtualizationVmwareDatacenter) SetVmTemplateCount(v int64) {
	o.VmTemplateCount = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatacenter) GetHypervisorManager() VirtualizationVmwareVcenterRelationship {
	if o == nil || IsNil(o.HypervisorManager.Get()) {
		var ret VirtualizationVmwareVcenterRelationship
		return ret
	}
	return *o.HypervisorManager.Get()
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatacenter) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.HypervisorManager.Get(), o.HypervisorManager.IsSet()
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager.IsSet() {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given NullableVirtualizationVmwareVcenterRelationship and assigns it to the HypervisorManager field.
func (o *VirtualizationVmwareDatacenter) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship) {
	o.HypervisorManager.Set(&v)
}

// SetHypervisorManagerNil sets the value for HypervisorManager to be an explicit nil
func (o *VirtualizationVmwareDatacenter) SetHypervisorManagerNil() {
	o.HypervisorManager.Set(nil)
}

// UnsetHypervisorManager ensures that no value is present for HypervisorManager, not even an explicit nil
func (o *VirtualizationVmwareDatacenter) UnsetHypervisorManager() {
	o.HypervisorManager.Unset()
}

// GetParentFolder returns the ParentFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatacenter) GetParentFolder() VirtualizationVmwareFolderRelationship {
	if o == nil || IsNil(o.ParentFolder.Get()) {
		var ret VirtualizationVmwareFolderRelationship
		return ret
	}
	return *o.ParentFolder.Get()
}

// GetParentFolderOk returns a tuple with the ParentFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatacenter) GetParentFolderOk() (*VirtualizationVmwareFolderRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFolder.Get(), o.ParentFolder.IsSet()
}

// HasParentFolder returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasParentFolder() bool {
	if o != nil && o.ParentFolder.IsSet() {
		return true
	}

	return false
}

// SetParentFolder gets a reference to the given NullableVirtualizationVmwareFolderRelationship and assigns it to the ParentFolder field.
func (o *VirtualizationVmwareDatacenter) SetParentFolder(v VirtualizationVmwareFolderRelationship) {
	o.ParentFolder.Set(&v)
}

// SetParentFolderNil sets the value for ParentFolder to be an explicit nil
func (o *VirtualizationVmwareDatacenter) SetParentFolderNil() {
	o.ParentFolder.Set(nil)
}

// UnsetParentFolder ensures that no value is present for ParentFolder, not even an explicit nil
func (o *VirtualizationVmwareDatacenter) UnsetParentFolder() {
	o.ParentFolder.Unset()
}

func (o VirtualizationVmwareDatacenter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareDatacenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDatacenter, errVirtualizationBaseDatacenter := json.Marshal(o.VirtualizationBaseDatacenter)
	if errVirtualizationBaseDatacenter != nil {
		return map[string]interface{}{}, errVirtualizationBaseDatacenter
	}
	errVirtualizationBaseDatacenter = json.Unmarshal([]byte(serializedVirtualizationBaseDatacenter), &toSerialize)
	if errVirtualizationBaseDatacenter != nil {
		return map[string]interface{}{}, errVirtualizationBaseDatacenter
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ClusterCount) {
		toSerialize["ClusterCount"] = o.ClusterCount
	}
	if !IsNil(o.DatastoreCount) {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if !IsNil(o.DistributedNetworkCount) {
		toSerialize["DistributedNetworkCount"] = o.DistributedNetworkCount
	}
	if !IsNil(o.DistributedVirtualSwitchCount) {
		toSerialize["DistributedVirtualSwitchCount"] = o.DistributedVirtualSwitchCount
	}
	if !IsNil(o.HostCount) {
		toSerialize["HostCount"] = o.HostCount
	}
	if !IsNil(o.InventoryPath) {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if !IsNil(o.NetworkCount) {
		toSerialize["NetworkCount"] = o.NetworkCount
	}
	if !IsNil(o.StandardNetworkCount) {
		toSerialize["StandardNetworkCount"] = o.StandardNetworkCount
	}
	if !IsNil(o.VmCount) {
		toSerialize["VmCount"] = o.VmCount
	}
	if !IsNil(o.VmTemplateCount) {
		toSerialize["VmTemplateCount"] = o.VmTemplateCount
	}
	if o.HypervisorManager.IsSet() {
		toSerialize["HypervisorManager"] = o.HypervisorManager.Get()
	}
	if o.ParentFolder.IsSet() {
		toSerialize["ParentFolder"] = o.ParentFolder.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareDatacenter) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareDatacenterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Count of all clusters associated with this DC.
		ClusterCount *int64 `json:"ClusterCount,omitempty"`
		// Count of all datastores associated with this DC.
		DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
		// Count of all distributed networks associated with this datacenter (DC).
		DistributedNetworkCount *int64 `json:"DistributedNetworkCount,omitempty"`
		// Count of all distributed virtual switches associated with this datacenter (DC).
		DistributedVirtualSwitchCount *int64 `json:"DistributedVirtualSwitchCount,omitempty"`
		// Count of all hosts associated with this DC.
		HostCount *int64 `json:"HostCount,omitempty"`
		// Inventory path of the DC.
		InventoryPath *string `json:"InventoryPath,omitempty"`
		// Count of all networks associated with this datacenter (DC).
		NetworkCount *int64 `json:"NetworkCount,omitempty"`
		// Count of all standard networks associated with this datacenter (DC).
		StandardNetworkCount *int64 `json:"StandardNetworkCount,omitempty"`
		// Count of all virtual machines (VMs) associated with this DC.
		VmCount *int64 `json:"VmCount,omitempty"`
		// Count of all virtual machines templates associated with this DC.
		VmTemplateCount   *int64                                          `json:"VmTemplateCount,omitempty"`
		HypervisorManager NullableVirtualizationVmwareVcenterRelationship `json:"HypervisorManager,omitempty"`
		ParentFolder      NullableVirtualizationVmwareFolderRelationship  `json:"ParentFolder,omitempty"`
	}

	varVirtualizationVmwareDatacenterWithoutEmbeddedStruct := VirtualizationVmwareDatacenterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareDatacenterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareDatacenter := _VirtualizationVmwareDatacenter{}
		varVirtualizationVmwareDatacenter.ClassId = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareDatacenter.ObjectType = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareDatacenter.ClusterCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ClusterCount
		varVirtualizationVmwareDatacenter.DatastoreCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.DatastoreCount
		varVirtualizationVmwareDatacenter.DistributedNetworkCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.DistributedNetworkCount
		varVirtualizationVmwareDatacenter.DistributedVirtualSwitchCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.DistributedVirtualSwitchCount
		varVirtualizationVmwareDatacenter.HostCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.HostCount
		varVirtualizationVmwareDatacenter.InventoryPath = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.InventoryPath
		varVirtualizationVmwareDatacenter.NetworkCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.NetworkCount
		varVirtualizationVmwareDatacenter.StandardNetworkCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.StandardNetworkCount
		varVirtualizationVmwareDatacenter.VmCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.VmCount
		varVirtualizationVmwareDatacenter.VmTemplateCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.VmTemplateCount
		varVirtualizationVmwareDatacenter.HypervisorManager = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.HypervisorManager
		varVirtualizationVmwareDatacenter.ParentFolder = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ParentFolder
		*o = VirtualizationVmwareDatacenter(varVirtualizationVmwareDatacenter)
	} else {
		return err
	}

	varVirtualizationVmwareDatacenter := _VirtualizationVmwareDatacenter{}

	err = json.Unmarshal(data, &varVirtualizationVmwareDatacenter)
	if err == nil {
		o.VirtualizationBaseDatacenter = varVirtualizationVmwareDatacenter.VirtualizationBaseDatacenter
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterCount")
		delete(additionalProperties, "DatastoreCount")
		delete(additionalProperties, "DistributedNetworkCount")
		delete(additionalProperties, "DistributedVirtualSwitchCount")
		delete(additionalProperties, "HostCount")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "NetworkCount")
		delete(additionalProperties, "StandardNetworkCount")
		delete(additionalProperties, "VmCount")
		delete(additionalProperties, "VmTemplateCount")
		delete(additionalProperties, "HypervisorManager")
		delete(additionalProperties, "ParentFolder")

		// remove fields from embedded structs
		reflectVirtualizationBaseDatacenter := reflect.ValueOf(o.VirtualizationBaseDatacenter)
		for i := 0; i < reflectVirtualizationBaseDatacenter.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDatacenter.Type().Field(i)

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

type NullableVirtualizationVmwareDatacenter struct {
	value *VirtualizationVmwareDatacenter
	isSet bool
}

func (v NullableVirtualizationVmwareDatacenter) Get() *VirtualizationVmwareDatacenter {
	return v.value
}

func (v *NullableVirtualizationVmwareDatacenter) Set(val *VirtualizationVmwareDatacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatacenter(val *VirtualizationVmwareDatacenter) *NullableVirtualizationVmwareDatacenter {
	return &NullableVirtualizationVmwareDatacenter{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
