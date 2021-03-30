/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationVmEsxiDiskAllOf Definition of the list of properties defined in 'virtualization.VmEsxiDisk', excluding properties defined in parent classes.
type VirtualizationVmEsxiDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk capacity to be provided with units example - 10Gi.
	Capacity *string `json:"Capacity,omitempty"`
	// Datastore of the disk from the space is allocated.
	Datastore *string `json:"Datastore,omitempty"`
	// Mode of the disk, like persistent, independent persistent.
	Diskmode *string `json:"Diskmode,omitempty"`
	// Name of the virtual disk.
	Name *string `json:"Name,omitempty"`
	// Specify the allocation type as thin or thick.
	StorageAllocation *string `json:"StorageAllocation,omitempty"`
	// Controller name of the disk, if not specified uses the default one.
	StorageController    *string `json:"StorageController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmEsxiDiskAllOf VirtualizationVmEsxiDiskAllOf

// NewVirtualizationVmEsxiDiskAllOf instantiates a new VirtualizationVmEsxiDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmEsxiDiskAllOf(classId string, objectType string) *VirtualizationVmEsxiDiskAllOf {
	this := VirtualizationVmEsxiDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmEsxiDiskAllOfWithDefaults instantiates a new VirtualizationVmEsxiDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmEsxiDiskAllOfWithDefaults() *VirtualizationVmEsxiDiskAllOf {
	this := VirtualizationVmEsxiDiskAllOf{}
	var classId string = "virtualization.VmEsxiDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VmEsxiDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmEsxiDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmEsxiDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmEsxiDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmEsxiDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationVmEsxiDiskAllOf) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationVmEsxiDiskAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *VirtualizationVmEsxiDiskAllOf) SetCapacity(v string) {
	o.Capacity = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VirtualizationVmEsxiDiskAllOf) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VirtualizationVmEsxiDiskAllOf) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *VirtualizationVmEsxiDiskAllOf) SetDatastore(v string) {
	o.Datastore = &v
}

// GetDiskmode returns the Diskmode field value if set, zero value otherwise.
func (o *VirtualizationVmEsxiDiskAllOf) GetDiskmode() string {
	if o == nil || o.Diskmode == nil {
		var ret string
		return ret
	}
	return *o.Diskmode
}

// GetDiskmodeOk returns a tuple with the Diskmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetDiskmodeOk() (*string, bool) {
	if o == nil || o.Diskmode == nil {
		return nil, false
	}
	return o.Diskmode, true
}

// HasDiskmode returns a boolean if a field has been set.
func (o *VirtualizationVmEsxiDiskAllOf) HasDiskmode() bool {
	if o != nil && o.Diskmode != nil {
		return true
	}

	return false
}

// SetDiskmode gets a reference to the given string and assigns it to the Diskmode field.
func (o *VirtualizationVmEsxiDiskAllOf) SetDiskmode(v string) {
	o.Diskmode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVmEsxiDiskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVmEsxiDiskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVmEsxiDiskAllOf) SetName(v string) {
	o.Name = &v
}

// GetStorageAllocation returns the StorageAllocation field value if set, zero value otherwise.
func (o *VirtualizationVmEsxiDiskAllOf) GetStorageAllocation() string {
	if o == nil || o.StorageAllocation == nil {
		var ret string
		return ret
	}
	return *o.StorageAllocation
}

// GetStorageAllocationOk returns a tuple with the StorageAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetStorageAllocationOk() (*string, bool) {
	if o == nil || o.StorageAllocation == nil {
		return nil, false
	}
	return o.StorageAllocation, true
}

// HasStorageAllocation returns a boolean if a field has been set.
func (o *VirtualizationVmEsxiDiskAllOf) HasStorageAllocation() bool {
	if o != nil && o.StorageAllocation != nil {
		return true
	}

	return false
}

// SetStorageAllocation gets a reference to the given string and assigns it to the StorageAllocation field.
func (o *VirtualizationVmEsxiDiskAllOf) SetStorageAllocation(v string) {
	o.StorageAllocation = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *VirtualizationVmEsxiDiskAllOf) GetStorageController() string {
	if o == nil || o.StorageController == nil {
		var ret string
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmEsxiDiskAllOf) GetStorageControllerOk() (*string, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *VirtualizationVmEsxiDiskAllOf) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given string and assigns it to the StorageController field.
func (o *VirtualizationVmEsxiDiskAllOf) SetStorageController(v string) {
	o.StorageController = &v
}

func (o VirtualizationVmEsxiDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Datastore != nil {
		toSerialize["Datastore"] = o.Datastore
	}
	if o.Diskmode != nil {
		toSerialize["Diskmode"] = o.Diskmode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StorageAllocation != nil {
		toSerialize["StorageAllocation"] = o.StorageAllocation
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmEsxiDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmEsxiDiskAllOf := _VirtualizationVmEsxiDiskAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmEsxiDiskAllOf); err == nil {
		*o = VirtualizationVmEsxiDiskAllOf(varVirtualizationVmEsxiDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Datastore")
		delete(additionalProperties, "Diskmode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StorageAllocation")
		delete(additionalProperties, "StorageController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmEsxiDiskAllOf struct {
	value *VirtualizationVmEsxiDiskAllOf
	isSet bool
}

func (v NullableVirtualizationVmEsxiDiskAllOf) Get() *VirtualizationVmEsxiDiskAllOf {
	return v.value
}

func (v *NullableVirtualizationVmEsxiDiskAllOf) Set(val *VirtualizationVmEsxiDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmEsxiDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmEsxiDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmEsxiDiskAllOf(val *VirtualizationVmEsxiDiskAllOf) *NullableVirtualizationVmEsxiDiskAllOf {
	return &NullableVirtualizationVmEsxiDiskAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmEsxiDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmEsxiDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
