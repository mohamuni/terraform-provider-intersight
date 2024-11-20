/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the StorageNvmeVirtualDriveConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNvmeVirtualDriveConfiguration{}

// StorageNvmeVirtualDriveConfiguration Virtual Drive type models a single virtual drive that needs to be created for NVMe Raid on reboot. Used in activation workflow.
type StorageNvmeVirtualDriveConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This defines the characteristics of a specific virtual drive.
	AccessPolicy *string `json:"AccessPolicy,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	Bootable *bool `json:"Bootable,omitempty"`
	// This defines the characteristics of a specific storage controller.
	ControllerDn *string `json:"ControllerDn,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	DedicatedHotSpare *string `json:"DedicatedHotSpare,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	DiskCachePolicy *string `json:"DiskCachePolicy,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed.
	Name *string `json:"Name,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	RaidLevel *string `json:"RaidLevel,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	ReadPolicy *string `json:"ReadPolicy,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	SelfEncrypt *string `json:"SelfEncrypt,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	Size      *string  `json:"Size,omitempty"`
	SpanDisks []string `json:"SpanDisks,omitempty"`
	// Virtual drive strip size.
	StripSize *string `json:"StripSize,omitempty"`
	// This defines the characteristics of a specific virtual drive.
	WritePolicy          *string `json:"WritePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNvmeVirtualDriveConfiguration StorageNvmeVirtualDriveConfiguration

// NewStorageNvmeVirtualDriveConfiguration instantiates a new StorageNvmeVirtualDriveConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNvmeVirtualDriveConfiguration(classId string, objectType string) *StorageNvmeVirtualDriveConfiguration {
	this := StorageNvmeVirtualDriveConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNvmeVirtualDriveConfigurationWithDefaults instantiates a new StorageNvmeVirtualDriveConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNvmeVirtualDriveConfigurationWithDefaults() *StorageNvmeVirtualDriveConfiguration {
	this := StorageNvmeVirtualDriveConfiguration{}
	var classId string = "storage.NvmeVirtualDriveConfiguration"
	this.ClassId = classId
	var objectType string = "storage.NvmeVirtualDriveConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNvmeVirtualDriveConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNvmeVirtualDriveConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NvmeVirtualDriveConfiguration" of the ClassId field.
func (o *StorageNvmeVirtualDriveConfiguration) GetDefaultClassId() interface{} {
	return "storage.NvmeVirtualDriveConfiguration"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNvmeVirtualDriveConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNvmeVirtualDriveConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NvmeVirtualDriveConfiguration" of the ObjectType field.
func (o *StorageNvmeVirtualDriveConfiguration) GetDefaultObjectType() interface{} {
	return "storage.NvmeVirtualDriveConfiguration"
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetAccessPolicy() string {
	if o == nil || IsNil(o.AccessPolicy) {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetAccessPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicy) {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasAccessPolicy() bool {
	if o != nil && !IsNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *StorageNvmeVirtualDriveConfiguration) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetBootable() bool {
	if o == nil || IsNil(o.Bootable) {
		var ret bool
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetBootableOk() (*bool, bool) {
	if o == nil || IsNil(o.Bootable) {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasBootable() bool {
	if o != nil && !IsNil(o.Bootable) {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *StorageNvmeVirtualDriveConfiguration) SetBootable(v bool) {
	o.Bootable = &v
}

// GetControllerDn returns the ControllerDn field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetControllerDn() string {
	if o == nil || IsNil(o.ControllerDn) {
		var ret string
		return ret
	}
	return *o.ControllerDn
}

// GetControllerDnOk returns a tuple with the ControllerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetControllerDnOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerDn) {
		return nil, false
	}
	return o.ControllerDn, true
}

// HasControllerDn returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasControllerDn() bool {
	if o != nil && !IsNil(o.ControllerDn) {
		return true
	}

	return false
}

// SetControllerDn gets a reference to the given string and assigns it to the ControllerDn field.
func (o *StorageNvmeVirtualDriveConfiguration) SetControllerDn(v string) {
	o.ControllerDn = &v
}

// GetDedicatedHotSpare returns the DedicatedHotSpare field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetDedicatedHotSpare() string {
	if o == nil || IsNil(o.DedicatedHotSpare) {
		var ret string
		return ret
	}
	return *o.DedicatedHotSpare
}

// GetDedicatedHotSpareOk returns a tuple with the DedicatedHotSpare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetDedicatedHotSpareOk() (*string, bool) {
	if o == nil || IsNil(o.DedicatedHotSpare) {
		return nil, false
	}
	return o.DedicatedHotSpare, true
}

// HasDedicatedHotSpare returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasDedicatedHotSpare() bool {
	if o != nil && !IsNil(o.DedicatedHotSpare) {
		return true
	}

	return false
}

// SetDedicatedHotSpare gets a reference to the given string and assigns it to the DedicatedHotSpare field.
func (o *StorageNvmeVirtualDriveConfiguration) SetDedicatedHotSpare(v string) {
	o.DedicatedHotSpare = &v
}

// GetDiskCachePolicy returns the DiskCachePolicy field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetDiskCachePolicy() string {
	if o == nil || IsNil(o.DiskCachePolicy) {
		var ret string
		return ret
	}
	return *o.DiskCachePolicy
}

// GetDiskCachePolicyOk returns a tuple with the DiskCachePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetDiskCachePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DiskCachePolicy) {
		return nil, false
	}
	return o.DiskCachePolicy, true
}

// HasDiskCachePolicy returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasDiskCachePolicy() bool {
	if o != nil && !IsNil(o.DiskCachePolicy) {
		return true
	}

	return false
}

// SetDiskCachePolicy gets a reference to the given string and assigns it to the DiskCachePolicy field.
func (o *StorageNvmeVirtualDriveConfiguration) SetDiskCachePolicy(v string) {
	o.DiskCachePolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNvmeVirtualDriveConfiguration) SetName(v string) {
	o.Name = &v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetRaidLevel() string {
	if o == nil || IsNil(o.RaidLevel) {
		var ret string
		return ret
	}
	return *o.RaidLevel
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetRaidLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RaidLevel) {
		return nil, false
	}
	return o.RaidLevel, true
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasRaidLevel() bool {
	if o != nil && !IsNil(o.RaidLevel) {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given string and assigns it to the RaidLevel field.
func (o *StorageNvmeVirtualDriveConfiguration) SetRaidLevel(v string) {
	o.RaidLevel = &v
}

// GetReadPolicy returns the ReadPolicy field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetReadPolicy() string {
	if o == nil || IsNil(o.ReadPolicy) {
		var ret string
		return ret
	}
	return *o.ReadPolicy
}

// GetReadPolicyOk returns a tuple with the ReadPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetReadPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ReadPolicy) {
		return nil, false
	}
	return o.ReadPolicy, true
}

// HasReadPolicy returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasReadPolicy() bool {
	if o != nil && !IsNil(o.ReadPolicy) {
		return true
	}

	return false
}

// SetReadPolicy gets a reference to the given string and assigns it to the ReadPolicy field.
func (o *StorageNvmeVirtualDriveConfiguration) SetReadPolicy(v string) {
	o.ReadPolicy = &v
}

// GetSelfEncrypt returns the SelfEncrypt field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetSelfEncrypt() string {
	if o == nil || IsNil(o.SelfEncrypt) {
		var ret string
		return ret
	}
	return *o.SelfEncrypt
}

// GetSelfEncryptOk returns a tuple with the SelfEncrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetSelfEncryptOk() (*string, bool) {
	if o == nil || IsNil(o.SelfEncrypt) {
		return nil, false
	}
	return o.SelfEncrypt, true
}

// HasSelfEncrypt returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasSelfEncrypt() bool {
	if o != nil && !IsNil(o.SelfEncrypt) {
		return true
	}

	return false
}

// SetSelfEncrypt gets a reference to the given string and assigns it to the SelfEncrypt field.
func (o *StorageNvmeVirtualDriveConfiguration) SetSelfEncrypt(v string) {
	o.SelfEncrypt = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageNvmeVirtualDriveConfiguration) SetSize(v string) {
	o.Size = &v
}

// GetSpanDisks returns the SpanDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNvmeVirtualDriveConfiguration) GetSpanDisks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SpanDisks
}

// GetSpanDisksOk returns a tuple with the SpanDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNvmeVirtualDriveConfiguration) GetSpanDisksOk() ([]string, bool) {
	if o == nil || IsNil(o.SpanDisks) {
		return nil, false
	}
	return o.SpanDisks, true
}

// HasSpanDisks returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasSpanDisks() bool {
	if o != nil && !IsNil(o.SpanDisks) {
		return true
	}

	return false
}

// SetSpanDisks gets a reference to the given []string and assigns it to the SpanDisks field.
func (o *StorageNvmeVirtualDriveConfiguration) SetSpanDisks(v []string) {
	o.SpanDisks = v
}

// GetStripSize returns the StripSize field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetStripSize() string {
	if o == nil || IsNil(o.StripSize) {
		var ret string
		return ret
	}
	return *o.StripSize
}

// GetStripSizeOk returns a tuple with the StripSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetStripSizeOk() (*string, bool) {
	if o == nil || IsNil(o.StripSize) {
		return nil, false
	}
	return o.StripSize, true
}

// HasStripSize returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasStripSize() bool {
	if o != nil && !IsNil(o.StripSize) {
		return true
	}

	return false
}

// SetStripSize gets a reference to the given string and assigns it to the StripSize field.
func (o *StorageNvmeVirtualDriveConfiguration) SetStripSize(v string) {
	o.StripSize = &v
}

// GetWritePolicy returns the WritePolicy field value if set, zero value otherwise.
func (o *StorageNvmeVirtualDriveConfiguration) GetWritePolicy() string {
	if o == nil || IsNil(o.WritePolicy) {
		var ret string
		return ret
	}
	return *o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeVirtualDriveConfiguration) GetWritePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.WritePolicy) {
		return nil, false
	}
	return o.WritePolicy, true
}

// HasWritePolicy returns a boolean if a field has been set.
func (o *StorageNvmeVirtualDriveConfiguration) HasWritePolicy() bool {
	if o != nil && !IsNil(o.WritePolicy) {
		return true
	}

	return false
}

// SetWritePolicy gets a reference to the given string and assigns it to the WritePolicy field.
func (o *StorageNvmeVirtualDriveConfiguration) SetWritePolicy(v string) {
	o.WritePolicy = &v
}

func (o StorageNvmeVirtualDriveConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNvmeVirtualDriveConfiguration) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AccessPolicy) {
		toSerialize["AccessPolicy"] = o.AccessPolicy
	}
	if !IsNil(o.Bootable) {
		toSerialize["Bootable"] = o.Bootable
	}
	if !IsNil(o.ControllerDn) {
		toSerialize["ControllerDn"] = o.ControllerDn
	}
	if !IsNil(o.DedicatedHotSpare) {
		toSerialize["DedicatedHotSpare"] = o.DedicatedHotSpare
	}
	if !IsNil(o.DiskCachePolicy) {
		toSerialize["DiskCachePolicy"] = o.DiskCachePolicy
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.RaidLevel) {
		toSerialize["RaidLevel"] = o.RaidLevel
	}
	if !IsNil(o.ReadPolicy) {
		toSerialize["ReadPolicy"] = o.ReadPolicy
	}
	if !IsNil(o.SelfEncrypt) {
		toSerialize["SelfEncrypt"] = o.SelfEncrypt
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if o.SpanDisks != nil {
		toSerialize["SpanDisks"] = o.SpanDisks
	}
	if !IsNil(o.StripSize) {
		toSerialize["StripSize"] = o.StripSize
	}
	if !IsNil(o.WritePolicy) {
		toSerialize["WritePolicy"] = o.WritePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNvmeVirtualDriveConfiguration) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This defines the characteristics of a specific virtual drive.
		AccessPolicy *string `json:"AccessPolicy,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		Bootable *bool `json:"Bootable,omitempty"`
		// This defines the characteristics of a specific storage controller.
		ControllerDn *string `json:"ControllerDn,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		DedicatedHotSpare *string `json:"DedicatedHotSpare,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		DiskCachePolicy *string `json:"DiskCachePolicy,omitempty"`
		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed.
		Name *string `json:"Name,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		RaidLevel *string `json:"RaidLevel,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		ReadPolicy *string `json:"ReadPolicy,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		SelfEncrypt *string `json:"SelfEncrypt,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		Size      *string  `json:"Size,omitempty"`
		SpanDisks []string `json:"SpanDisks,omitempty"`
		// Virtual drive strip size.
		StripSize *string `json:"StripSize,omitempty"`
		// This defines the characteristics of a specific virtual drive.
		WritePolicy *string `json:"WritePolicy,omitempty"`
	}

	varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct := StorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varStorageNvmeVirtualDriveConfiguration := _StorageNvmeVirtualDriveConfiguration{}
		varStorageNvmeVirtualDriveConfiguration.ClassId = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.ClassId
		varStorageNvmeVirtualDriveConfiguration.ObjectType = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.ObjectType
		varStorageNvmeVirtualDriveConfiguration.AccessPolicy = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.AccessPolicy
		varStorageNvmeVirtualDriveConfiguration.Bootable = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.Bootable
		varStorageNvmeVirtualDriveConfiguration.ControllerDn = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.ControllerDn
		varStorageNvmeVirtualDriveConfiguration.DedicatedHotSpare = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.DedicatedHotSpare
		varStorageNvmeVirtualDriveConfiguration.DiskCachePolicy = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.DiskCachePolicy
		varStorageNvmeVirtualDriveConfiguration.Name = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.Name
		varStorageNvmeVirtualDriveConfiguration.RaidLevel = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.RaidLevel
		varStorageNvmeVirtualDriveConfiguration.ReadPolicy = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.ReadPolicy
		varStorageNvmeVirtualDriveConfiguration.SelfEncrypt = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.SelfEncrypt
		varStorageNvmeVirtualDriveConfiguration.Size = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.Size
		varStorageNvmeVirtualDriveConfiguration.SpanDisks = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.SpanDisks
		varStorageNvmeVirtualDriveConfiguration.StripSize = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.StripSize
		varStorageNvmeVirtualDriveConfiguration.WritePolicy = varStorageNvmeVirtualDriveConfigurationWithoutEmbeddedStruct.WritePolicy
		*o = StorageNvmeVirtualDriveConfiguration(varStorageNvmeVirtualDriveConfiguration)
	} else {
		return err
	}

	varStorageNvmeVirtualDriveConfiguration := _StorageNvmeVirtualDriveConfiguration{}

	err = json.Unmarshal(data, &varStorageNvmeVirtualDriveConfiguration)
	if err == nil {
		o.MoBaseComplexType = varStorageNvmeVirtualDriveConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessPolicy")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "ControllerDn")
		delete(additionalProperties, "DedicatedHotSpare")
		delete(additionalProperties, "DiskCachePolicy")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RaidLevel")
		delete(additionalProperties, "ReadPolicy")
		delete(additionalProperties, "SelfEncrypt")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "SpanDisks")
		delete(additionalProperties, "StripSize")
		delete(additionalProperties, "WritePolicy")

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

type NullableStorageNvmeVirtualDriveConfiguration struct {
	value *StorageNvmeVirtualDriveConfiguration
	isSet bool
}

func (v NullableStorageNvmeVirtualDriveConfiguration) Get() *StorageNvmeVirtualDriveConfiguration {
	return v.value
}

func (v *NullableStorageNvmeVirtualDriveConfiguration) Set(val *StorageNvmeVirtualDriveConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNvmeVirtualDriveConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNvmeVirtualDriveConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNvmeVirtualDriveConfiguration(val *StorageNvmeVirtualDriveConfiguration) *NullableStorageNvmeVirtualDriveConfiguration {
	return &NullableStorageNvmeVirtualDriveConfiguration{value: val, isSet: true}
}

func (v NullableStorageNvmeVirtualDriveConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNvmeVirtualDriveConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
