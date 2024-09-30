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

// checks if the NiatelemetryLc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryLc{}

// NiatelemetryLc Object is available at Line Card scope.
type NiatelemetryLc struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the line cards present.
	Description *string `json:"Description,omitempty"`
	// Dn value for the line cards present.
	Dn *string `json:"Dn,omitempty"`
	// Hardware version of the line cards present.
	HardwareVersion *string `json:"HardwareVersion,omitempty"`
	// Model of the line cards present.
	Model *string `json:"Model,omitempty"`
	// Node Id of the line card present.
	NodeId *int64 `json:"NodeId,omitempty"`
	// Opretaional state of the line cards present.
	OperationalState *string `json:"OperationalState,omitempty"`
	// Power state of the line cards present.
	PowerState *string `json:"PowerState,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Redundancy state of the line cards present.
	RedundancyState *string `json:"RedundancyState,omitempty"`
	// Serial number of the line card present.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.
	SiteName *string `json:"SiteName,omitempty"`
	// VID for the line card in the inventory.
	Vid                  *string                                     `json:"Vid,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryLc NiatelemetryLc

// NewNiatelemetryLc instantiates a new NiatelemetryLc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryLc(classId string, objectType string) *NiatelemetryLc {
	this := NiatelemetryLc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryLcWithDefaults instantiates a new NiatelemetryLc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryLcWithDefaults() *NiatelemetryLc {
	this := NiatelemetryLc{}
	var classId string = "niatelemetry.Lc"
	this.ClassId = classId
	var objectType string = "niatelemetry.Lc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryLc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryLc) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.Lc" of the ClassId field.
func (o *NiatelemetryLc) GetDefaultClassId() interface{} {
	return "niatelemetry.Lc"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryLc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryLc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.Lc" of the ObjectType field.
func (o *NiatelemetryLc) GetDefaultObjectType() interface{} {
	return "niatelemetry.Lc"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NiatelemetryLc) SetDescription(v string) {
	o.Description = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryLc) SetDn(v string) {
	o.Dn = &v
}

// GetHardwareVersion returns the HardwareVersion field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetHardwareVersion() string {
	if o == nil || IsNil(o.HardwareVersion) {
		var ret string
		return ret
	}
	return *o.HardwareVersion
}

// GetHardwareVersionOk returns a tuple with the HardwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetHardwareVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareVersion) {
		return nil, false
	}
	return o.HardwareVersion, true
}

// HasHardwareVersion returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasHardwareVersion() bool {
	if o != nil && !IsNil(o.HardwareVersion) {
		return true
	}

	return false
}

// SetHardwareVersion gets a reference to the given string and assigns it to the HardwareVersion field.
func (o *NiatelemetryLc) SetHardwareVersion(v string) {
	o.HardwareVersion = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *NiatelemetryLc) SetModel(v string) {
	o.Model = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetNodeId() int64 {
	if o == nil || IsNil(o.NodeId) {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetNodeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *NiatelemetryLc) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetOperationalState() string {
	if o == nil || IsNil(o.OperationalState) {
		var ret string
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetOperationalStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given string and assigns it to the OperationalState field.
func (o *NiatelemetryLc) SetOperationalState(v string) {
	o.OperationalState = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetPowerState() string {
	if o == nil || IsNil(o.PowerState) {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetPowerStateOk() (*string, bool) {
	if o == nil || IsNil(o.PowerState) {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasPowerState() bool {
	if o != nil && !IsNil(o.PowerState) {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *NiatelemetryLc) SetPowerState(v string) {
	o.PowerState = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryLc) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryLc) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetRedundancyState returns the RedundancyState field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetRedundancyState() string {
	if o == nil || IsNil(o.RedundancyState) {
		var ret string
		return ret
	}
	return *o.RedundancyState
}

// GetRedundancyStateOk returns a tuple with the RedundancyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetRedundancyStateOk() (*string, bool) {
	if o == nil || IsNil(o.RedundancyState) {
		return nil, false
	}
	return o.RedundancyState, true
}

// HasRedundancyState returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasRedundancyState() bool {
	if o != nil && !IsNil(o.RedundancyState) {
		return true
	}

	return false
}

// SetRedundancyState gets a reference to the given string and assigns it to the RedundancyState field.
func (o *NiatelemetryLc) SetRedundancyState(v string) {
	o.RedundancyState = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryLc) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryLc) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *NiatelemetryLc) GetVid() string {
	if o == nil || IsNil(o.Vid) {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLc) GetVidOk() (*string, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *NiatelemetryLc) SetVid(v string) {
	o.Vid = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryLc) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryLc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryLc) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryLc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryLc) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryLc) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryLc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryLc) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.HardwareVersion) {
		toSerialize["HardwareVersion"] = o.HardwareVersion
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.NodeId) {
		toSerialize["NodeId"] = o.NodeId
	}
	if !IsNil(o.OperationalState) {
		toSerialize["OperationalState"] = o.OperationalState
	}
	if !IsNil(o.PowerState) {
		toSerialize["PowerState"] = o.PowerState
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if !IsNil(o.RedundancyState) {
		toSerialize["RedundancyState"] = o.RedundancyState
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if !IsNil(o.SiteName) {
		toSerialize["SiteName"] = o.SiteName
	}
	if !IsNil(o.Vid) {
		toSerialize["Vid"] = o.Vid
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryLc) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryLcWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the line cards present.
		Description *string `json:"Description,omitempty"`
		// Dn value for the line cards present.
		Dn *string `json:"Dn,omitempty"`
		// Hardware version of the line cards present.
		HardwareVersion *string `json:"HardwareVersion,omitempty"`
		// Model of the line cards present.
		Model *string `json:"Model,omitempty"`
		// Node Id of the line card present.
		NodeId *int64 `json:"NodeId,omitempty"`
		// Opretaional state of the line cards present.
		OperationalState *string `json:"OperationalState,omitempty"`
		// Power state of the line cards present.
		PowerState *string `json:"PowerState,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Redundancy state of the line cards present.
		RedundancyState *string `json:"RedundancyState,omitempty"`
		// Serial number of the line card present.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.
		SiteName *string `json:"SiteName,omitempty"`
		// VID for the line card in the inventory.
		Vid              *string                                     `json:"Vid,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryLcWithoutEmbeddedStruct := NiatelemetryLcWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryLcWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryLc := _NiatelemetryLc{}
		varNiatelemetryLc.ClassId = varNiatelemetryLcWithoutEmbeddedStruct.ClassId
		varNiatelemetryLc.ObjectType = varNiatelemetryLcWithoutEmbeddedStruct.ObjectType
		varNiatelemetryLc.Description = varNiatelemetryLcWithoutEmbeddedStruct.Description
		varNiatelemetryLc.Dn = varNiatelemetryLcWithoutEmbeddedStruct.Dn
		varNiatelemetryLc.HardwareVersion = varNiatelemetryLcWithoutEmbeddedStruct.HardwareVersion
		varNiatelemetryLc.Model = varNiatelemetryLcWithoutEmbeddedStruct.Model
		varNiatelemetryLc.NodeId = varNiatelemetryLcWithoutEmbeddedStruct.NodeId
		varNiatelemetryLc.OperationalState = varNiatelemetryLcWithoutEmbeddedStruct.OperationalState
		varNiatelemetryLc.PowerState = varNiatelemetryLcWithoutEmbeddedStruct.PowerState
		varNiatelemetryLc.RecordType = varNiatelemetryLcWithoutEmbeddedStruct.RecordType
		varNiatelemetryLc.RecordVersion = varNiatelemetryLcWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryLc.RedundancyState = varNiatelemetryLcWithoutEmbeddedStruct.RedundancyState
		varNiatelemetryLc.SerialNumber = varNiatelemetryLcWithoutEmbeddedStruct.SerialNumber
		varNiatelemetryLc.SiteName = varNiatelemetryLcWithoutEmbeddedStruct.SiteName
		varNiatelemetryLc.Vid = varNiatelemetryLcWithoutEmbeddedStruct.Vid
		varNiatelemetryLc.RegisteredDevice = varNiatelemetryLcWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryLc(varNiatelemetryLc)
	} else {
		return err
	}

	varNiatelemetryLc := _NiatelemetryLc{}

	err = json.Unmarshal(data, &varNiatelemetryLc)
	if err == nil {
		o.MoBaseMo = varNiatelemetryLc.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "HardwareVersion")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "OperationalState")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "RedundancyState")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryLc struct {
	value *NiatelemetryLc
	isSet bool
}

func (v NullableNiatelemetryLc) Get() *NiatelemetryLc {
	return v.value
}

func (v *NullableNiatelemetryLc) Set(val *NiatelemetryLc) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryLc) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryLc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryLc(val *NiatelemetryLc) *NullableNiatelemetryLc {
	return &NullableNiatelemetryLc{value: val, isSet: true}
}

func (v NullableNiatelemetryLc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryLc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
