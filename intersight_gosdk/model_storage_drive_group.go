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

// checks if the StorageDriveGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageDriveGroup{}

// StorageDriveGroup A reusable RAID drive group configuration that specifies a pool of drives and a set of virtual drives that are to be created using this pool of drives.
type StorageDriveGroup struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType          string                             `json:"ObjectType"`
	AutomaticDriveGroup NullableStorageAutomaticDriveGroup `json:"AutomaticDriveGroup,omitempty"`
	ManualDriveGroup    NullableStorageManualDriveGroup    `json:"ManualDriveGroup,omitempty"`
	// The name of the drive group. The name can be between 1 and 64 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,64}$"`
	// The supported RAID level for the disk group. * `Raid0` - RAID 0 Stripe Raid Level. * `Raid1` - RAID 1 Mirror Raid Level. * `Raid5` - RAID 5 Mirror Raid Level. * `Raid6` - RAID 6 Mirror Raid Level. * `Raid10` - RAID 10 Mirror Raid Level. * `Raid50` - RAID 50 Mirror Raid Level. * `Raid60` - RAID 60 Mirror Raid Level.
	RaidLevel *string `json:"RaidLevel,omitempty"`
	// Enables/disables the drive security on all the drives used in this policy. This flag just enables the drive security and only after Remote/Manual key setting configured, the actual security will be applied.
	SecureDriveGroup *bool `json:"SecureDriveGroup,omitempty"`
	// Type of drive selection to be used for this drive group. * `0` - Drives are selected manually by the user. * `1` - Drives are selected automatically based on the RAID and virtual drive configuration.
	Type                 *int32                                   `json:"Type,omitempty"`
	VirtualDrives        []StorageVirtualDriveConfiguration       `json:"VirtualDrives,omitempty"`
	StoragePolicy        NullableStorageStoragePolicyRelationship `json:"StoragePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageDriveGroup StorageDriveGroup

// NewStorageDriveGroup instantiates a new StorageDriveGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDriveGroup(classId string, objectType string) *StorageDriveGroup {
	this := StorageDriveGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// NewStorageDriveGroupWithDefaults instantiates a new StorageDriveGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDriveGroupWithDefaults() *StorageDriveGroup {
	this := StorageDriveGroup{}
	var classId string = "storage.DriveGroup"
	this.ClassId = classId
	var objectType string = "storage.DriveGroup"
	this.ObjectType = objectType
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageDriveGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageDriveGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageDriveGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.DriveGroup" of the ClassId field.
func (o *StorageDriveGroup) GetDefaultClassId() interface{} {
	return "storage.DriveGroup"
}

// GetObjectType returns the ObjectType field value
func (o *StorageDriveGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageDriveGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageDriveGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.DriveGroup" of the ObjectType field.
func (o *StorageDriveGroup) GetDefaultObjectType() interface{} {
	return "storage.DriveGroup"
}

// GetAutomaticDriveGroup returns the AutomaticDriveGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroup) GetAutomaticDriveGroup() StorageAutomaticDriveGroup {
	if o == nil || IsNil(o.AutomaticDriveGroup.Get()) {
		var ret StorageAutomaticDriveGroup
		return ret
	}
	return *o.AutomaticDriveGroup.Get()
}

// GetAutomaticDriveGroupOk returns a tuple with the AutomaticDriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroup) GetAutomaticDriveGroupOk() (*StorageAutomaticDriveGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomaticDriveGroup.Get(), o.AutomaticDriveGroup.IsSet()
}

// HasAutomaticDriveGroup returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasAutomaticDriveGroup() bool {
	if o != nil && o.AutomaticDriveGroup.IsSet() {
		return true
	}

	return false
}

// SetAutomaticDriveGroup gets a reference to the given NullableStorageAutomaticDriveGroup and assigns it to the AutomaticDriveGroup field.
func (o *StorageDriveGroup) SetAutomaticDriveGroup(v StorageAutomaticDriveGroup) {
	o.AutomaticDriveGroup.Set(&v)
}

// SetAutomaticDriveGroupNil sets the value for AutomaticDriveGroup to be an explicit nil
func (o *StorageDriveGroup) SetAutomaticDriveGroupNil() {
	o.AutomaticDriveGroup.Set(nil)
}

// UnsetAutomaticDriveGroup ensures that no value is present for AutomaticDriveGroup, not even an explicit nil
func (o *StorageDriveGroup) UnsetAutomaticDriveGroup() {
	o.AutomaticDriveGroup.Unset()
}

// GetManualDriveGroup returns the ManualDriveGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroup) GetManualDriveGroup() StorageManualDriveGroup {
	if o == nil || IsNil(o.ManualDriveGroup.Get()) {
		var ret StorageManualDriveGroup
		return ret
	}
	return *o.ManualDriveGroup.Get()
}

// GetManualDriveGroupOk returns a tuple with the ManualDriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroup) GetManualDriveGroupOk() (*StorageManualDriveGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManualDriveGroup.Get(), o.ManualDriveGroup.IsSet()
}

// HasManualDriveGroup returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasManualDriveGroup() bool {
	if o != nil && o.ManualDriveGroup.IsSet() {
		return true
	}

	return false
}

// SetManualDriveGroup gets a reference to the given NullableStorageManualDriveGroup and assigns it to the ManualDriveGroup field.
func (o *StorageDriveGroup) SetManualDriveGroup(v StorageManualDriveGroup) {
	o.ManualDriveGroup.Set(&v)
}

// SetManualDriveGroupNil sets the value for ManualDriveGroup to be an explicit nil
func (o *StorageDriveGroup) SetManualDriveGroupNil() {
	o.ManualDriveGroup.Set(nil)
}

// UnsetManualDriveGroup ensures that no value is present for ManualDriveGroup, not even an explicit nil
func (o *StorageDriveGroup) UnsetManualDriveGroup() {
	o.ManualDriveGroup.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageDriveGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageDriveGroup) SetName(v string) {
	o.Name = &v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise.
func (o *StorageDriveGroup) GetRaidLevel() string {
	if o == nil || IsNil(o.RaidLevel) {
		var ret string
		return ret
	}
	return *o.RaidLevel
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroup) GetRaidLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RaidLevel) {
		return nil, false
	}
	return o.RaidLevel, true
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasRaidLevel() bool {
	if o != nil && !IsNil(o.RaidLevel) {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given string and assigns it to the RaidLevel field.
func (o *StorageDriveGroup) SetRaidLevel(v string) {
	o.RaidLevel = &v
}

// GetSecureDriveGroup returns the SecureDriveGroup field value if set, zero value otherwise.
func (o *StorageDriveGroup) GetSecureDriveGroup() bool {
	if o == nil || IsNil(o.SecureDriveGroup) {
		var ret bool
		return ret
	}
	return *o.SecureDriveGroup
}

// GetSecureDriveGroupOk returns a tuple with the SecureDriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroup) GetSecureDriveGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureDriveGroup) {
		return nil, false
	}
	return o.SecureDriveGroup, true
}

// HasSecureDriveGroup returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasSecureDriveGroup() bool {
	if o != nil && !IsNil(o.SecureDriveGroup) {
		return true
	}

	return false
}

// SetSecureDriveGroup gets a reference to the given bool and assigns it to the SecureDriveGroup field.
func (o *StorageDriveGroup) SetSecureDriveGroup(v bool) {
	o.SecureDriveGroup = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageDriveGroup) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroup) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *StorageDriveGroup) SetType(v int32) {
	o.Type = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroup) GetVirtualDrives() []StorageVirtualDriveConfiguration {
	if o == nil {
		var ret []StorageVirtualDriveConfiguration
		return ret
	}
	return o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroup) GetVirtualDrivesOk() ([]StorageVirtualDriveConfiguration, bool) {
	if o == nil || IsNil(o.VirtualDrives) {
		return nil, false
	}
	return o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasVirtualDrives() bool {
	if o != nil && !IsNil(o.VirtualDrives) {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []StorageVirtualDriveConfiguration and assigns it to the VirtualDrives field.
func (o *StorageDriveGroup) SetVirtualDrives(v []StorageVirtualDriveConfiguration) {
	o.VirtualDrives = v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroup) GetStoragePolicy() StorageStoragePolicyRelationship {
	if o == nil || IsNil(o.StoragePolicy.Get()) {
		var ret StorageStoragePolicyRelationship
		return ret
	}
	return *o.StoragePolicy.Get()
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroup) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoragePolicy.Get(), o.StoragePolicy.IsSet()
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *StorageDriveGroup) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy.IsSet() {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given NullableStorageStoragePolicyRelationship and assigns it to the StoragePolicy field.
func (o *StorageDriveGroup) SetStoragePolicy(v StorageStoragePolicyRelationship) {
	o.StoragePolicy.Set(&v)
}

// SetStoragePolicyNil sets the value for StoragePolicy to be an explicit nil
func (o *StorageDriveGroup) SetStoragePolicyNil() {
	o.StoragePolicy.Set(nil)
}

// UnsetStoragePolicy ensures that no value is present for StoragePolicy, not even an explicit nil
func (o *StorageDriveGroup) UnsetStoragePolicy() {
	o.StoragePolicy.Unset()
}

func (o StorageDriveGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageDriveGroup) ToMap() (map[string]interface{}, error) {
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
	if o.AutomaticDriveGroup.IsSet() {
		toSerialize["AutomaticDriveGroup"] = o.AutomaticDriveGroup.Get()
	}
	if o.ManualDriveGroup.IsSet() {
		toSerialize["ManualDriveGroup"] = o.ManualDriveGroup.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.RaidLevel) {
		toSerialize["RaidLevel"] = o.RaidLevel
	}
	if !IsNil(o.SecureDriveGroup) {
		toSerialize["SecureDriveGroup"] = o.SecureDriveGroup
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}
	if o.StoragePolicy.IsSet() {
		toSerialize["StoragePolicy"] = o.StoragePolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageDriveGroup) UnmarshalJSON(data []byte) (err error) {
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
	type StorageDriveGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                             `json:"ObjectType"`
		AutomaticDriveGroup NullableStorageAutomaticDriveGroup `json:"AutomaticDriveGroup,omitempty"`
		ManualDriveGroup    NullableStorageManualDriveGroup    `json:"ManualDriveGroup,omitempty"`
		// The name of the drive group. The name can be between 1 and 64 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,64}$"`
		// The supported RAID level for the disk group. * `Raid0` - RAID 0 Stripe Raid Level. * `Raid1` - RAID 1 Mirror Raid Level. * `Raid5` - RAID 5 Mirror Raid Level. * `Raid6` - RAID 6 Mirror Raid Level. * `Raid10` - RAID 10 Mirror Raid Level. * `Raid50` - RAID 50 Mirror Raid Level. * `Raid60` - RAID 60 Mirror Raid Level.
		RaidLevel *string `json:"RaidLevel,omitempty"`
		// Enables/disables the drive security on all the drives used in this policy. This flag just enables the drive security and only after Remote/Manual key setting configured, the actual security will be applied.
		SecureDriveGroup *bool `json:"SecureDriveGroup,omitempty"`
		// Type of drive selection to be used for this drive group. * `0` - Drives are selected manually by the user. * `1` - Drives are selected automatically based on the RAID and virtual drive configuration.
		Type          *int32                                   `json:"Type,omitempty"`
		VirtualDrives []StorageVirtualDriveConfiguration       `json:"VirtualDrives,omitempty"`
		StoragePolicy NullableStorageStoragePolicyRelationship `json:"StoragePolicy,omitempty"`
	}

	varStorageDriveGroupWithoutEmbeddedStruct := StorageDriveGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageDriveGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageDriveGroup := _StorageDriveGroup{}
		varStorageDriveGroup.ClassId = varStorageDriveGroupWithoutEmbeddedStruct.ClassId
		varStorageDriveGroup.ObjectType = varStorageDriveGroupWithoutEmbeddedStruct.ObjectType
		varStorageDriveGroup.AutomaticDriveGroup = varStorageDriveGroupWithoutEmbeddedStruct.AutomaticDriveGroup
		varStorageDriveGroup.ManualDriveGroup = varStorageDriveGroupWithoutEmbeddedStruct.ManualDriveGroup
		varStorageDriveGroup.Name = varStorageDriveGroupWithoutEmbeddedStruct.Name
		varStorageDriveGroup.RaidLevel = varStorageDriveGroupWithoutEmbeddedStruct.RaidLevel
		varStorageDriveGroup.SecureDriveGroup = varStorageDriveGroupWithoutEmbeddedStruct.SecureDriveGroup
		varStorageDriveGroup.Type = varStorageDriveGroupWithoutEmbeddedStruct.Type
		varStorageDriveGroup.VirtualDrives = varStorageDriveGroupWithoutEmbeddedStruct.VirtualDrives
		varStorageDriveGroup.StoragePolicy = varStorageDriveGroupWithoutEmbeddedStruct.StoragePolicy
		*o = StorageDriveGroup(varStorageDriveGroup)
	} else {
		return err
	}

	varStorageDriveGroup := _StorageDriveGroup{}

	err = json.Unmarshal(data, &varStorageDriveGroup)
	if err == nil {
		o.MoBaseMo = varStorageDriveGroup.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutomaticDriveGroup")
		delete(additionalProperties, "ManualDriveGroup")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RaidLevel")
		delete(additionalProperties, "SecureDriveGroup")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VirtualDrives")
		delete(additionalProperties, "StoragePolicy")

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

type NullableStorageDriveGroup struct {
	value *StorageDriveGroup
	isSet bool
}

func (v NullableStorageDriveGroup) Get() *StorageDriveGroup {
	return v.value
}

func (v *NullableStorageDriveGroup) Set(val *StorageDriveGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDriveGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDriveGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDriveGroup(val *StorageDriveGroup) *NullableStorageDriveGroup {
	return &NullableStorageDriveGroup{value: val, isSet: true}
}

func (v NullableStorageDriveGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDriveGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
