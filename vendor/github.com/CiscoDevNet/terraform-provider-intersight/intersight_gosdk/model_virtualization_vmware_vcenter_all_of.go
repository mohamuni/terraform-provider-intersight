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
)

// VirtualizationVmwareVcenterAllOf Definition of the list of properties defined in 'virtualization.VmwareVcenter', excluding properties defined in parent classes.
type VirtualizationVmwareVcenterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Count of all Clusters associated with the vcenter.
	ClusterCount *int64 `json:"ClusterCount,omitempty"`
	// Count of all Datacenters in the vcenter.
	DatacenterCount *int64 `json:"DatacenterCount,omitempty"`
	// Count of all Datastores Templates associated with the vcenter.
	DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
	// Count of all Distributed Virtual Switches associated with vcenter.
	DistributedVirtualSwitchCount *int64 `json:"DistributedVirtualSwitchCount,omitempty"`
	// Count of all Datastore cluster associated with the vcenter.
	DsClusterCount *int64 `json:"DsClusterCount,omitempty"`
	// External Ip address fot the vcenter.
	ExternalIp *string `json:"ExternalIp,omitempty"`
	// Count of all Hosts associated with the vcenter.
	HostCount *int64   `json:"HostCount,omitempty"`
	IpAddress []string `json:"IpAddress,omitempty"`
	// Name of th Target with which the vcenter was claimed.
	TargetName *string `json:"TargetName,omitempty"`
	// Count of all Virtual Machines associated with the vcenter.
	VmCount *int64 `json:"VmCount,omitempty"`
	// Count of all VM Templates associated with the vcenter.
	VmTemplatesCount     *int64 `json:"VmTemplatesCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVcenterAllOf VirtualizationVmwareVcenterAllOf

// NewVirtualizationVmwareVcenterAllOf instantiates a new VirtualizationVmwareVcenterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVcenterAllOf(classId string, objectType string) *VirtualizationVmwareVcenterAllOf {
	this := VirtualizationVmwareVcenterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVcenterAllOfWithDefaults instantiates a new VirtualizationVmwareVcenterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVcenterAllOfWithDefaults() *VirtualizationVmwareVcenterAllOf {
	this := VirtualizationVmwareVcenterAllOf{}
	var classId string = "virtualization.VmwareVcenter"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVcenter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVcenterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVcenterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVcenterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVcenterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetClusterCount() int64 {
	if o == nil || o.ClusterCount == nil {
		var ret int64
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetClusterCountOk() (*int64, bool) {
	if o == nil || o.ClusterCount == nil {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasClusterCount() bool {
	if o != nil && o.ClusterCount != nil {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int64 and assigns it to the ClusterCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetClusterCount(v int64) {
	o.ClusterCount = &v
}

// GetDatacenterCount returns the DatacenterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetDatacenterCount() int64 {
	if o == nil || o.DatacenterCount == nil {
		var ret int64
		return ret
	}
	return *o.DatacenterCount
}

// GetDatacenterCountOk returns a tuple with the DatacenterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetDatacenterCountOk() (*int64, bool) {
	if o == nil || o.DatacenterCount == nil {
		return nil, false
	}
	return o.DatacenterCount, true
}

// HasDatacenterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasDatacenterCount() bool {
	if o != nil && o.DatacenterCount != nil {
		return true
	}

	return false
}

// SetDatacenterCount gets a reference to the given int64 and assigns it to the DatacenterCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetDatacenterCount(v int64) {
	o.DatacenterCount = &v
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetDatastoreCount() int64 {
	if o == nil || o.DatastoreCount == nil {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || o.DatastoreCount == nil {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasDatastoreCount() bool {
	if o != nil && o.DatastoreCount != nil {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetDistributedVirtualSwitchCount returns the DistributedVirtualSwitchCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetDistributedVirtualSwitchCount() int64 {
	if o == nil || o.DistributedVirtualSwitchCount == nil {
		var ret int64
		return ret
	}
	return *o.DistributedVirtualSwitchCount
}

// GetDistributedVirtualSwitchCountOk returns a tuple with the DistributedVirtualSwitchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetDistributedVirtualSwitchCountOk() (*int64, bool) {
	if o == nil || o.DistributedVirtualSwitchCount == nil {
		return nil, false
	}
	return o.DistributedVirtualSwitchCount, true
}

// HasDistributedVirtualSwitchCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasDistributedVirtualSwitchCount() bool {
	if o != nil && o.DistributedVirtualSwitchCount != nil {
		return true
	}

	return false
}

// SetDistributedVirtualSwitchCount gets a reference to the given int64 and assigns it to the DistributedVirtualSwitchCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetDistributedVirtualSwitchCount(v int64) {
	o.DistributedVirtualSwitchCount = &v
}

// GetDsClusterCount returns the DsClusterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetDsClusterCount() int64 {
	if o == nil || o.DsClusterCount == nil {
		var ret int64
		return ret
	}
	return *o.DsClusterCount
}

// GetDsClusterCountOk returns a tuple with the DsClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetDsClusterCountOk() (*int64, bool) {
	if o == nil || o.DsClusterCount == nil {
		return nil, false
	}
	return o.DsClusterCount, true
}

// HasDsClusterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasDsClusterCount() bool {
	if o != nil && o.DsClusterCount != nil {
		return true
	}

	return false
}

// SetDsClusterCount gets a reference to the given int64 and assigns it to the DsClusterCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetDsClusterCount(v int64) {
	o.DsClusterCount = &v
}

// GetExternalIp returns the ExternalIp field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetExternalIp() string {
	if o == nil || o.ExternalIp == nil {
		var ret string
		return ret
	}
	return *o.ExternalIp
}

// GetExternalIpOk returns a tuple with the ExternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetExternalIpOk() (*string, bool) {
	if o == nil || o.ExternalIp == nil {
		return nil, false
	}
	return o.ExternalIp, true
}

// HasExternalIp returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasExternalIp() bool {
	if o != nil && o.ExternalIp != nil {
		return true
	}

	return false
}

// SetExternalIp gets a reference to the given string and assigns it to the ExternalIp field.
func (o *VirtualizationVmwareVcenterAllOf) SetExternalIp(v string) {
	o.ExternalIp = &v
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareVcenterAllOf) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareVcenterAllOf) GetIpAddressOk() ([]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *VirtualizationVmwareVcenterAllOf) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *VirtualizationVmwareVcenterAllOf) SetTargetName(v string) {
	o.TargetName = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetVmTemplatesCount returns the VmTemplatesCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareVcenterAllOf) GetVmTemplatesCount() int64 {
	if o == nil || o.VmTemplatesCount == nil {
		var ret int64
		return ret
	}
	return *o.VmTemplatesCount
}

// GetVmTemplatesCountOk returns a tuple with the VmTemplatesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVcenterAllOf) GetVmTemplatesCountOk() (*int64, bool) {
	if o == nil || o.VmTemplatesCount == nil {
		return nil, false
	}
	return o.VmTemplatesCount, true
}

// HasVmTemplatesCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareVcenterAllOf) HasVmTemplatesCount() bool {
	if o != nil && o.VmTemplatesCount != nil {
		return true
	}

	return false
}

// SetVmTemplatesCount gets a reference to the given int64 and assigns it to the VmTemplatesCount field.
func (o *VirtualizationVmwareVcenterAllOf) SetVmTemplatesCount(v int64) {
	o.VmTemplatesCount = &v
}

func (o VirtualizationVmwareVcenterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterCount != nil {
		toSerialize["ClusterCount"] = o.ClusterCount
	}
	if o.DatacenterCount != nil {
		toSerialize["DatacenterCount"] = o.DatacenterCount
	}
	if o.DatastoreCount != nil {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if o.DistributedVirtualSwitchCount != nil {
		toSerialize["DistributedVirtualSwitchCount"] = o.DistributedVirtualSwitchCount
	}
	if o.DsClusterCount != nil {
		toSerialize["DsClusterCount"] = o.DsClusterCount
	}
	if o.ExternalIp != nil {
		toSerialize["ExternalIp"] = o.ExternalIp
	}
	if o.HostCount != nil {
		toSerialize["HostCount"] = o.HostCount
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.VmTemplatesCount != nil {
		toSerialize["VmTemplatesCount"] = o.VmTemplatesCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVcenterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareVcenterAllOf := _VirtualizationVmwareVcenterAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareVcenterAllOf); err == nil {
		*o = VirtualizationVmwareVcenterAllOf(varVirtualizationVmwareVcenterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterCount")
		delete(additionalProperties, "DatacenterCount")
		delete(additionalProperties, "DatastoreCount")
		delete(additionalProperties, "DistributedVirtualSwitchCount")
		delete(additionalProperties, "DsClusterCount")
		delete(additionalProperties, "ExternalIp")
		delete(additionalProperties, "HostCount")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "VmCount")
		delete(additionalProperties, "VmTemplatesCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVcenterAllOf struct {
	value *VirtualizationVmwareVcenterAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVcenterAllOf) Get() *VirtualizationVmwareVcenterAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVcenterAllOf) Set(val *VirtualizationVmwareVcenterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVcenterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVcenterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVcenterAllOf(val *VirtualizationVmwareVcenterAllOf) *NullableVirtualizationVmwareVcenterAllOf {
	return &NullableVirtualizationVmwareVcenterAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVcenterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVcenterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}