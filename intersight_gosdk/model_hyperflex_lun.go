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

// checks if the HyperflexLun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexLun{}

// HyperflexLun A HyperFlex iSCSI logical unit number (LUN) entity. Contains detailed information about the iSCSI LUN which includes the identity and capacity information, and the iSCSI target to which it is associated.
type HyperflexLun struct {
	StorageBaseHostLun
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description of iSCSI LUN.
	Description *string `json:"Description,omitempty"`
	// The datastore capacity in bytes.
	DsCapacityInBytes *int64 `json:"DsCapacityInBytes,omitempty"`
	// The HyperFlex datastore name.
	DsName *string `json:"DsName,omitempty"`
	// The HyperFlex datastore UUID.
	DsUuid *string `json:"DsUuid,omitempty"`
	// Source of the lun inventory. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
	InventorySource *string `json:"InventorySource,omitempty"`
	// Indicates if the datastore on which LUN resides is encrypted or un-encrypted.
	IsEncrypted *bool `json:"IsEncrypted,omitempty"`
	// The identity of HyperFlex iSCSI LUN.
	LunId *string `json:"LunId,omitempty"`
	// Name of the HyperFlex iSCSI LUN.
	Name *string `json:"Name,omitempty"`
	// Serial number of HyperFlex iSCSI LUN.
	SerialNo *string `json:"SerialNo,omitempty"`
	// Total capacity of iSCSI LUN in bytes.
	TotalCapacityInBytes *int64 `json:"TotalCapacityInBytes,omitempty"`
	// HyperFlex iSCSI Target associated with the HyperFlex iSCSI LUN.
	Tuuid *string `json:"Tuuid,omitempty"`
	// Used capacity of iSCSI LUN in bytes.
	UsedCapacityInBytes *int64 `json:"UsedCapacityInBytes,omitempty"`
	// UUID of the HyperFlex iSCSI LUN.
	Uuid *string `json:"Uuid,omitempty"`
	// Version of the HyperFlex iSCSI lun.
	Version              *int64                              `json:"Version,omitempty"`
	Target               NullableHyperflexTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexLun HyperflexLun

// NewHyperflexLun instantiates a new HyperflexLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexLun(classId string, objectType string) *HyperflexLun {
	this := HyperflexLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexLunWithDefaults instantiates a new HyperflexLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexLunWithDefaults() *HyperflexLun {
	this := HyperflexLun{}
	var classId string = "hyperflex.Lun"
	this.ClassId = classId
	var objectType string = "hyperflex.Lun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexLun) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.Lun" of the ClassId field.
func (o *HyperflexLun) GetDefaultClassId() interface{} {
	return "hyperflex.Lun"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.Lun" of the ObjectType field.
func (o *HyperflexLun) GetDefaultObjectType() interface{} {
	return "hyperflex.Lun"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexLun) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexLun) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexLun) SetDescription(v string) {
	o.Description = &v
}

// GetDsCapacityInBytes returns the DsCapacityInBytes field value if set, zero value otherwise.
func (o *HyperflexLun) GetDsCapacityInBytes() int64 {
	if o == nil || IsNil(o.DsCapacityInBytes) {
		var ret int64
		return ret
	}
	return *o.DsCapacityInBytes
}

// GetDsCapacityInBytesOk returns a tuple with the DsCapacityInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetDsCapacityInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DsCapacityInBytes) {
		return nil, false
	}
	return o.DsCapacityInBytes, true
}

// HasDsCapacityInBytes returns a boolean if a field has been set.
func (o *HyperflexLun) HasDsCapacityInBytes() bool {
	if o != nil && !IsNil(o.DsCapacityInBytes) {
		return true
	}

	return false
}

// SetDsCapacityInBytes gets a reference to the given int64 and assigns it to the DsCapacityInBytes field.
func (o *HyperflexLun) SetDsCapacityInBytes(v int64) {
	o.DsCapacityInBytes = &v
}

// GetDsName returns the DsName field value if set, zero value otherwise.
func (o *HyperflexLun) GetDsName() string {
	if o == nil || IsNil(o.DsName) {
		var ret string
		return ret
	}
	return *o.DsName
}

// GetDsNameOk returns a tuple with the DsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetDsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DsName) {
		return nil, false
	}
	return o.DsName, true
}

// HasDsName returns a boolean if a field has been set.
func (o *HyperflexLun) HasDsName() bool {
	if o != nil && !IsNil(o.DsName) {
		return true
	}

	return false
}

// SetDsName gets a reference to the given string and assigns it to the DsName field.
func (o *HyperflexLun) SetDsName(v string) {
	o.DsName = &v
}

// GetDsUuid returns the DsUuid field value if set, zero value otherwise.
func (o *HyperflexLun) GetDsUuid() string {
	if o == nil || IsNil(o.DsUuid) {
		var ret string
		return ret
	}
	return *o.DsUuid
}

// GetDsUuidOk returns a tuple with the DsUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetDsUuidOk() (*string, bool) {
	if o == nil || IsNil(o.DsUuid) {
		return nil, false
	}
	return o.DsUuid, true
}

// HasDsUuid returns a boolean if a field has been set.
func (o *HyperflexLun) HasDsUuid() bool {
	if o != nil && !IsNil(o.DsUuid) {
		return true
	}

	return false
}

// SetDsUuid gets a reference to the given string and assigns it to the DsUuid field.
func (o *HyperflexLun) SetDsUuid(v string) {
	o.DsUuid = &v
}

// GetInventorySource returns the InventorySource field value if set, zero value otherwise.
func (o *HyperflexLun) GetInventorySource() string {
	if o == nil || IsNil(o.InventorySource) {
		var ret string
		return ret
	}
	return *o.InventorySource
}

// GetInventorySourceOk returns a tuple with the InventorySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetInventorySourceOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySource) {
		return nil, false
	}
	return o.InventorySource, true
}

// HasInventorySource returns a boolean if a field has been set.
func (o *HyperflexLun) HasInventorySource() bool {
	if o != nil && !IsNil(o.InventorySource) {
		return true
	}

	return false
}

// SetInventorySource gets a reference to the given string and assigns it to the InventorySource field.
func (o *HyperflexLun) SetInventorySource(v string) {
	o.InventorySource = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *HyperflexLun) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncrypted) {
		return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *HyperflexLun) HasIsEncrypted() bool {
	if o != nil && !IsNil(o.IsEncrypted) {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *HyperflexLun) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *HyperflexLun) GetLunId() string {
	if o == nil || IsNil(o.LunId) {
		var ret string
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetLunIdOk() (*string, bool) {
	if o == nil || IsNil(o.LunId) {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *HyperflexLun) HasLunId() bool {
	if o != nil && !IsNil(o.LunId) {
		return true
	}

	return false
}

// SetLunId gets a reference to the given string and assigns it to the LunId field.
func (o *HyperflexLun) SetLunId(v string) {
	o.LunId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexLun) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexLun) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexLun) SetName(v string) {
	o.Name = &v
}

// GetSerialNo returns the SerialNo field value if set, zero value otherwise.
func (o *HyperflexLun) GetSerialNo() string {
	if o == nil || IsNil(o.SerialNo) {
		var ret string
		return ret
	}
	return *o.SerialNo
}

// GetSerialNoOk returns a tuple with the SerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetSerialNoOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNo) {
		return nil, false
	}
	return o.SerialNo, true
}

// HasSerialNo returns a boolean if a field has been set.
func (o *HyperflexLun) HasSerialNo() bool {
	if o != nil && !IsNil(o.SerialNo) {
		return true
	}

	return false
}

// SetSerialNo gets a reference to the given string and assigns it to the SerialNo field.
func (o *HyperflexLun) SetSerialNo(v string) {
	o.SerialNo = &v
}

// GetTotalCapacityInBytes returns the TotalCapacityInBytes field value if set, zero value otherwise.
func (o *HyperflexLun) GetTotalCapacityInBytes() int64 {
	if o == nil || IsNil(o.TotalCapacityInBytes) {
		var ret int64
		return ret
	}
	return *o.TotalCapacityInBytes
}

// GetTotalCapacityInBytesOk returns a tuple with the TotalCapacityInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetTotalCapacityInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCapacityInBytes) {
		return nil, false
	}
	return o.TotalCapacityInBytes, true
}

// HasTotalCapacityInBytes returns a boolean if a field has been set.
func (o *HyperflexLun) HasTotalCapacityInBytes() bool {
	if o != nil && !IsNil(o.TotalCapacityInBytes) {
		return true
	}

	return false
}

// SetTotalCapacityInBytes gets a reference to the given int64 and assigns it to the TotalCapacityInBytes field.
func (o *HyperflexLun) SetTotalCapacityInBytes(v int64) {
	o.TotalCapacityInBytes = &v
}

// GetTuuid returns the Tuuid field value if set, zero value otherwise.
func (o *HyperflexLun) GetTuuid() string {
	if o == nil || IsNil(o.Tuuid) {
		var ret string
		return ret
	}
	return *o.Tuuid
}

// GetTuuidOk returns a tuple with the Tuuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetTuuidOk() (*string, bool) {
	if o == nil || IsNil(o.Tuuid) {
		return nil, false
	}
	return o.Tuuid, true
}

// HasTuuid returns a boolean if a field has been set.
func (o *HyperflexLun) HasTuuid() bool {
	if o != nil && !IsNil(o.Tuuid) {
		return true
	}

	return false
}

// SetTuuid gets a reference to the given string and assigns it to the Tuuid field.
func (o *HyperflexLun) SetTuuid(v string) {
	o.Tuuid = &v
}

// GetUsedCapacityInBytes returns the UsedCapacityInBytes field value if set, zero value otherwise.
func (o *HyperflexLun) GetUsedCapacityInBytes() int64 {
	if o == nil || IsNil(o.UsedCapacityInBytes) {
		var ret int64
		return ret
	}
	return *o.UsedCapacityInBytes
}

// GetUsedCapacityInBytesOk returns a tuple with the UsedCapacityInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetUsedCapacityInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedCapacityInBytes) {
		return nil, false
	}
	return o.UsedCapacityInBytes, true
}

// HasUsedCapacityInBytes returns a boolean if a field has been set.
func (o *HyperflexLun) HasUsedCapacityInBytes() bool {
	if o != nil && !IsNil(o.UsedCapacityInBytes) {
		return true
	}

	return false
}

// SetUsedCapacityInBytes gets a reference to the given int64 and assigns it to the UsedCapacityInBytes field.
func (o *HyperflexLun) SetUsedCapacityInBytes(v int64) {
	o.UsedCapacityInBytes = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexLun) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexLun) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexLun) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexLun) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLun) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexLun) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HyperflexLun) SetVersion(v int64) {
	o.Version = &v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexLun) GetTarget() HyperflexTargetRelationship {
	if o == nil || IsNil(o.Target.Get()) {
		var ret HyperflexTargetRelationship
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexLun) GetTargetOk() (*HyperflexTargetRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *HyperflexLun) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableHyperflexTargetRelationship and assigns it to the Target field.
func (o *HyperflexLun) SetTarget(v HyperflexTargetRelationship) {
	o.Target.Set(&v)
}

// SetTargetNil sets the value for Target to be an explicit nil
func (o *HyperflexLun) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *HyperflexLun) UnsetTarget() {
	o.Target.Unset()
}

func (o HyperflexLun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexLun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostLun, errStorageBaseHostLun := json.Marshal(o.StorageBaseHostLun)
	if errStorageBaseHostLun != nil {
		return map[string]interface{}{}, errStorageBaseHostLun
	}
	errStorageBaseHostLun = json.Unmarshal([]byte(serializedStorageBaseHostLun), &toSerialize)
	if errStorageBaseHostLun != nil {
		return map[string]interface{}{}, errStorageBaseHostLun
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DsCapacityInBytes) {
		toSerialize["DsCapacityInBytes"] = o.DsCapacityInBytes
	}
	if !IsNil(o.DsName) {
		toSerialize["DsName"] = o.DsName
	}
	if !IsNil(o.DsUuid) {
		toSerialize["DsUuid"] = o.DsUuid
	}
	if !IsNil(o.InventorySource) {
		toSerialize["InventorySource"] = o.InventorySource
	}
	if !IsNil(o.IsEncrypted) {
		toSerialize["IsEncrypted"] = o.IsEncrypted
	}
	if !IsNil(o.LunId) {
		toSerialize["LunId"] = o.LunId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SerialNo) {
		toSerialize["SerialNo"] = o.SerialNo
	}
	if !IsNil(o.TotalCapacityInBytes) {
		toSerialize["TotalCapacityInBytes"] = o.TotalCapacityInBytes
	}
	if !IsNil(o.Tuuid) {
		toSerialize["Tuuid"] = o.Tuuid
	}
	if !IsNil(o.UsedCapacityInBytes) {
		toSerialize["UsedCapacityInBytes"] = o.UsedCapacityInBytes
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.Target.IsSet() {
		toSerialize["Target"] = o.Target.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexLun) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexLunWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description of iSCSI LUN.
		Description *string `json:"Description,omitempty"`
		// The datastore capacity in bytes.
		DsCapacityInBytes *int64 `json:"DsCapacityInBytes,omitempty"`
		// The HyperFlex datastore name.
		DsName *string `json:"DsName,omitempty"`
		// The HyperFlex datastore UUID.
		DsUuid *string `json:"DsUuid,omitempty"`
		// Source of the lun inventory. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
		InventorySource *string `json:"InventorySource,omitempty"`
		// Indicates if the datastore on which LUN resides is encrypted or un-encrypted.
		IsEncrypted *bool `json:"IsEncrypted,omitempty"`
		// The identity of HyperFlex iSCSI LUN.
		LunId *string `json:"LunId,omitempty"`
		// Name of the HyperFlex iSCSI LUN.
		Name *string `json:"Name,omitempty"`
		// Serial number of HyperFlex iSCSI LUN.
		SerialNo *string `json:"SerialNo,omitempty"`
		// Total capacity of iSCSI LUN in bytes.
		TotalCapacityInBytes *int64 `json:"TotalCapacityInBytes,omitempty"`
		// HyperFlex iSCSI Target associated with the HyperFlex iSCSI LUN.
		Tuuid *string `json:"Tuuid,omitempty"`
		// Used capacity of iSCSI LUN in bytes.
		UsedCapacityInBytes *int64 `json:"UsedCapacityInBytes,omitempty"`
		// UUID of the HyperFlex iSCSI LUN.
		Uuid *string `json:"Uuid,omitempty"`
		// Version of the HyperFlex iSCSI lun.
		Version *int64                              `json:"Version,omitempty"`
		Target  NullableHyperflexTargetRelationship `json:"Target,omitempty"`
	}

	varHyperflexLunWithoutEmbeddedStruct := HyperflexLunWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexLunWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexLun := _HyperflexLun{}
		varHyperflexLun.ClassId = varHyperflexLunWithoutEmbeddedStruct.ClassId
		varHyperflexLun.ObjectType = varHyperflexLunWithoutEmbeddedStruct.ObjectType
		varHyperflexLun.Description = varHyperflexLunWithoutEmbeddedStruct.Description
		varHyperflexLun.DsCapacityInBytes = varHyperflexLunWithoutEmbeddedStruct.DsCapacityInBytes
		varHyperflexLun.DsName = varHyperflexLunWithoutEmbeddedStruct.DsName
		varHyperflexLun.DsUuid = varHyperflexLunWithoutEmbeddedStruct.DsUuid
		varHyperflexLun.InventorySource = varHyperflexLunWithoutEmbeddedStruct.InventorySource
		varHyperflexLun.IsEncrypted = varHyperflexLunWithoutEmbeddedStruct.IsEncrypted
		varHyperflexLun.LunId = varHyperflexLunWithoutEmbeddedStruct.LunId
		varHyperflexLun.Name = varHyperflexLunWithoutEmbeddedStruct.Name
		varHyperflexLun.SerialNo = varHyperflexLunWithoutEmbeddedStruct.SerialNo
		varHyperflexLun.TotalCapacityInBytes = varHyperflexLunWithoutEmbeddedStruct.TotalCapacityInBytes
		varHyperflexLun.Tuuid = varHyperflexLunWithoutEmbeddedStruct.Tuuid
		varHyperflexLun.UsedCapacityInBytes = varHyperflexLunWithoutEmbeddedStruct.UsedCapacityInBytes
		varHyperflexLun.Uuid = varHyperflexLunWithoutEmbeddedStruct.Uuid
		varHyperflexLun.Version = varHyperflexLunWithoutEmbeddedStruct.Version
		varHyperflexLun.Target = varHyperflexLunWithoutEmbeddedStruct.Target
		*o = HyperflexLun(varHyperflexLun)
	} else {
		return err
	}

	varHyperflexLun := _HyperflexLun{}

	err = json.Unmarshal(data, &varHyperflexLun)
	if err == nil {
		o.StorageBaseHostLun = varHyperflexLun.StorageBaseHostLun
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DsCapacityInBytes")
		delete(additionalProperties, "DsName")
		delete(additionalProperties, "DsUuid")
		delete(additionalProperties, "InventorySource")
		delete(additionalProperties, "IsEncrypted")
		delete(additionalProperties, "LunId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SerialNo")
		delete(additionalProperties, "TotalCapacityInBytes")
		delete(additionalProperties, "Tuuid")
		delete(additionalProperties, "UsedCapacityInBytes")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Target")

		// remove fields from embedded structs
		reflectStorageBaseHostLun := reflect.ValueOf(o.StorageBaseHostLun)
		for i := 0; i < reflectStorageBaseHostLun.Type().NumField(); i++ {
			t := reflectStorageBaseHostLun.Type().Field(i)

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

type NullableHyperflexLun struct {
	value *HyperflexLun
	isSet bool
}

func (v NullableHyperflexLun) Get() *HyperflexLun {
	return v.value
}

func (v *NullableHyperflexLun) Set(val *HyperflexLun) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexLun) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexLun(val *HyperflexLun) *NullableHyperflexLun {
	return &NullableHyperflexLun{value: val, isSet: true}
}

func (v NullableHyperflexLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
