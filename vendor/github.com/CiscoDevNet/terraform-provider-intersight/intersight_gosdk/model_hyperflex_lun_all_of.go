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

// HyperflexLunAllOf Definition of the list of properties defined in 'hyperflex.Lun', excluding properties defined in parent classes.
type HyperflexLunAllOf struct {
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
	Version              *int64                       `json:"Version,omitempty"`
	Target               *HyperflexTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexLunAllOf HyperflexLunAllOf

// NewHyperflexLunAllOf instantiates a new HyperflexLunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexLunAllOf(classId string, objectType string) *HyperflexLunAllOf {
	this := HyperflexLunAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexLunAllOfWithDefaults instantiates a new HyperflexLunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexLunAllOfWithDefaults() *HyperflexLunAllOf {
	this := HyperflexLunAllOf{}
	var classId string = "hyperflex.Lun"
	this.ClassId = classId
	var objectType string = "hyperflex.Lun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexLunAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexLunAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexLunAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexLunAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexLunAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDsCapacityInBytes returns the DsCapacityInBytes field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetDsCapacityInBytes() int64 {
	if o == nil || o.DsCapacityInBytes == nil {
		var ret int64
		return ret
	}
	return *o.DsCapacityInBytes
}

// GetDsCapacityInBytesOk returns a tuple with the DsCapacityInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetDsCapacityInBytesOk() (*int64, bool) {
	if o == nil || o.DsCapacityInBytes == nil {
		return nil, false
	}
	return o.DsCapacityInBytes, true
}

// HasDsCapacityInBytes returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasDsCapacityInBytes() bool {
	if o != nil && o.DsCapacityInBytes != nil {
		return true
	}

	return false
}

// SetDsCapacityInBytes gets a reference to the given int64 and assigns it to the DsCapacityInBytes field.
func (o *HyperflexLunAllOf) SetDsCapacityInBytes(v int64) {
	o.DsCapacityInBytes = &v
}

// GetDsName returns the DsName field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetDsName() string {
	if o == nil || o.DsName == nil {
		var ret string
		return ret
	}
	return *o.DsName
}

// GetDsNameOk returns a tuple with the DsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetDsNameOk() (*string, bool) {
	if o == nil || o.DsName == nil {
		return nil, false
	}
	return o.DsName, true
}

// HasDsName returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasDsName() bool {
	if o != nil && o.DsName != nil {
		return true
	}

	return false
}

// SetDsName gets a reference to the given string and assigns it to the DsName field.
func (o *HyperflexLunAllOf) SetDsName(v string) {
	o.DsName = &v
}

// GetDsUuid returns the DsUuid field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetDsUuid() string {
	if o == nil || o.DsUuid == nil {
		var ret string
		return ret
	}
	return *o.DsUuid
}

// GetDsUuidOk returns a tuple with the DsUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetDsUuidOk() (*string, bool) {
	if o == nil || o.DsUuid == nil {
		return nil, false
	}
	return o.DsUuid, true
}

// HasDsUuid returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasDsUuid() bool {
	if o != nil && o.DsUuid != nil {
		return true
	}

	return false
}

// SetDsUuid gets a reference to the given string and assigns it to the DsUuid field.
func (o *HyperflexLunAllOf) SetDsUuid(v string) {
	o.DsUuid = &v
}

// GetInventorySource returns the InventorySource field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetInventorySource() string {
	if o == nil || o.InventorySource == nil {
		var ret string
		return ret
	}
	return *o.InventorySource
}

// GetInventorySourceOk returns a tuple with the InventorySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetInventorySourceOk() (*string, bool) {
	if o == nil || o.InventorySource == nil {
		return nil, false
	}
	return o.InventorySource, true
}

// HasInventorySource returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasInventorySource() bool {
	if o != nil && o.InventorySource != nil {
		return true
	}

	return false
}

// SetInventorySource gets a reference to the given string and assigns it to the InventorySource field.
func (o *HyperflexLunAllOf) SetInventorySource(v string) {
	o.InventorySource = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetIsEncrypted() bool {
	if o == nil || o.IsEncrypted == nil {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || o.IsEncrypted == nil {
		return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted != nil {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *HyperflexLunAllOf) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetLunId() string {
	if o == nil || o.LunId == nil {
		var ret string
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetLunIdOk() (*string, bool) {
	if o == nil || o.LunId == nil {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasLunId() bool {
	if o != nil && o.LunId != nil {
		return true
	}

	return false
}

// SetLunId gets a reference to the given string and assigns it to the LunId field.
func (o *HyperflexLunAllOf) SetLunId(v string) {
	o.LunId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexLunAllOf) SetName(v string) {
	o.Name = &v
}

// GetSerialNo returns the SerialNo field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetSerialNo() string {
	if o == nil || o.SerialNo == nil {
		var ret string
		return ret
	}
	return *o.SerialNo
}

// GetSerialNoOk returns a tuple with the SerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetSerialNoOk() (*string, bool) {
	if o == nil || o.SerialNo == nil {
		return nil, false
	}
	return o.SerialNo, true
}

// HasSerialNo returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasSerialNo() bool {
	if o != nil && o.SerialNo != nil {
		return true
	}

	return false
}

// SetSerialNo gets a reference to the given string and assigns it to the SerialNo field.
func (o *HyperflexLunAllOf) SetSerialNo(v string) {
	o.SerialNo = &v
}

// GetTotalCapacityInBytes returns the TotalCapacityInBytes field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetTotalCapacityInBytes() int64 {
	if o == nil || o.TotalCapacityInBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacityInBytes
}

// GetTotalCapacityInBytesOk returns a tuple with the TotalCapacityInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetTotalCapacityInBytesOk() (*int64, bool) {
	if o == nil || o.TotalCapacityInBytes == nil {
		return nil, false
	}
	return o.TotalCapacityInBytes, true
}

// HasTotalCapacityInBytes returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasTotalCapacityInBytes() bool {
	if o != nil && o.TotalCapacityInBytes != nil {
		return true
	}

	return false
}

// SetTotalCapacityInBytes gets a reference to the given int64 and assigns it to the TotalCapacityInBytes field.
func (o *HyperflexLunAllOf) SetTotalCapacityInBytes(v int64) {
	o.TotalCapacityInBytes = &v
}

// GetTuuid returns the Tuuid field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetTuuid() string {
	if o == nil || o.Tuuid == nil {
		var ret string
		return ret
	}
	return *o.Tuuid
}

// GetTuuidOk returns a tuple with the Tuuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetTuuidOk() (*string, bool) {
	if o == nil || o.Tuuid == nil {
		return nil, false
	}
	return o.Tuuid, true
}

// HasTuuid returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasTuuid() bool {
	if o != nil && o.Tuuid != nil {
		return true
	}

	return false
}

// SetTuuid gets a reference to the given string and assigns it to the Tuuid field.
func (o *HyperflexLunAllOf) SetTuuid(v string) {
	o.Tuuid = &v
}

// GetUsedCapacityInBytes returns the UsedCapacityInBytes field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetUsedCapacityInBytes() int64 {
	if o == nil || o.UsedCapacityInBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedCapacityInBytes
}

// GetUsedCapacityInBytesOk returns a tuple with the UsedCapacityInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetUsedCapacityInBytesOk() (*int64, bool) {
	if o == nil || o.UsedCapacityInBytes == nil {
		return nil, false
	}
	return o.UsedCapacityInBytes, true
}

// HasUsedCapacityInBytes returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasUsedCapacityInBytes() bool {
	if o != nil && o.UsedCapacityInBytes != nil {
		return true
	}

	return false
}

// SetUsedCapacityInBytes gets a reference to the given int64 and assigns it to the UsedCapacityInBytes field.
func (o *HyperflexLunAllOf) SetUsedCapacityInBytes(v int64) {
	o.UsedCapacityInBytes = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexLunAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HyperflexLunAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *HyperflexLunAllOf) GetTarget() HyperflexTargetRelationship {
	if o == nil || o.Target == nil {
		var ret HyperflexTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLunAllOf) GetTargetOk() (*HyperflexTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *HyperflexLunAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given HyperflexTargetRelationship and assigns it to the Target field.
func (o *HyperflexLunAllOf) SetTarget(v HyperflexTargetRelationship) {
	o.Target = &v
}

func (o HyperflexLunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DsCapacityInBytes != nil {
		toSerialize["DsCapacityInBytes"] = o.DsCapacityInBytes
	}
	if o.DsName != nil {
		toSerialize["DsName"] = o.DsName
	}
	if o.DsUuid != nil {
		toSerialize["DsUuid"] = o.DsUuid
	}
	if o.InventorySource != nil {
		toSerialize["InventorySource"] = o.InventorySource
	}
	if o.IsEncrypted != nil {
		toSerialize["IsEncrypted"] = o.IsEncrypted
	}
	if o.LunId != nil {
		toSerialize["LunId"] = o.LunId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SerialNo != nil {
		toSerialize["SerialNo"] = o.SerialNo
	}
	if o.TotalCapacityInBytes != nil {
		toSerialize["TotalCapacityInBytes"] = o.TotalCapacityInBytes
	}
	if o.Tuuid != nil {
		toSerialize["Tuuid"] = o.Tuuid
	}
	if o.UsedCapacityInBytes != nil {
		toSerialize["UsedCapacityInBytes"] = o.UsedCapacityInBytes
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexLunAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexLunAllOf := _HyperflexLunAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexLunAllOf); err == nil {
		*o = HyperflexLunAllOf(varHyperflexLunAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexLunAllOf struct {
	value *HyperflexLunAllOf
	isSet bool
}

func (v NullableHyperflexLunAllOf) Get() *HyperflexLunAllOf {
	return v.value
}

func (v *NullableHyperflexLunAllOf) Set(val *HyperflexLunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexLunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexLunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexLunAllOf(val *HyperflexLunAllOf) *NullableHyperflexLunAllOf {
	return &NullableHyperflexLunAllOf{value: val, isSet: true}
}

func (v NullableHyperflexLunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexLunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
