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

// checks if the NiatelemetryApicFanDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryApicFanDetails{}

// NiatelemetryApicFanDetails Object to capture the fan details in APIC.
type NiatelemetryApicFanDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn for the fan in the inventory.
	Dn *string `json:"Dn,omitempty"`
	// Model number of the fan in APIC.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// Node id for the fan in the inventory.
	NodeId *int64 `json:"NodeId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the fan in APIC.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Vendor name of the fan in APIC.
	VendorName *string `json:"VendorName,omitempty"`
	// VID for the fan in the inventory.
	Vid                  *string                                     `json:"Vid,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicFanDetails NiatelemetryApicFanDetails

// NewNiatelemetryApicFanDetails instantiates a new NiatelemetryApicFanDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicFanDetails(classId string, objectType string) *NiatelemetryApicFanDetails {
	this := NiatelemetryApicFanDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicFanDetailsWithDefaults instantiates a new NiatelemetryApicFanDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicFanDetailsWithDefaults() *NiatelemetryApicFanDetails {
	this := NiatelemetryApicFanDetails{}
	var classId string = "niatelemetry.ApicFanDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicFanDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicFanDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicFanDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.ApicFanDetails" of the ClassId field.
func (o *NiatelemetryApicFanDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.ApicFanDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicFanDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicFanDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.ApicFanDetails" of the ObjectType field.
func (o *NiatelemetryApicFanDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.ApicFanDetails"
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicFanDetails) SetDn(v string) {
	o.Dn = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetModelNumber() string {
	if o == nil || IsNil(o.ModelNumber) {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetModelNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ModelNumber) {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasModelNumber() bool {
	if o != nil && !IsNil(o.ModelNumber) {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *NiatelemetryApicFanDetails) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetNodeId() int64 {
	if o == nil || IsNil(o.NodeId) {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetNodeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *NiatelemetryApicFanDetails) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicFanDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicFanDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryApicFanDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicFanDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *NiatelemetryApicFanDetails) SetVendorName(v string) {
	o.VendorName = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetVid() string {
	if o == nil || IsNil(o.Vid) {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetVidOk() (*string, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *NiatelemetryApicFanDetails) SetVid(v string) {
	o.Vid = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicFanDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicFanDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicFanDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryApicFanDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryApicFanDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryApicFanDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryApicFanDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.ModelNumber) {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if !IsNil(o.NodeId) {
		toSerialize["NodeId"] = o.NodeId
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if !IsNil(o.SiteName) {
		toSerialize["SiteName"] = o.SiteName
	}
	if !IsNil(o.VendorName) {
		toSerialize["VendorName"] = o.VendorName
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

func (o *NiatelemetryApicFanDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryApicFanDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn for the fan in the inventory.
		Dn *string `json:"Dn,omitempty"`
		// Model number of the fan in APIC.
		ModelNumber *string `json:"ModelNumber,omitempty"`
		// Node id for the fan in the inventory.
		NodeId *int64 `json:"NodeId,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of the fan in APIC.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// Vendor name of the fan in APIC.
		VendorName *string `json:"VendorName,omitempty"`
		// VID for the fan in the inventory.
		Vid              *string                                     `json:"Vid,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicFanDetailsWithoutEmbeddedStruct := NiatelemetryApicFanDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryApicFanDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicFanDetails := _NiatelemetryApicFanDetails{}
		varNiatelemetryApicFanDetails.ClassId = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicFanDetails.ObjectType = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicFanDetails.Dn = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicFanDetails.ModelNumber = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.ModelNumber
		varNiatelemetryApicFanDetails.NodeId = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.NodeId
		varNiatelemetryApicFanDetails.RecordType = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicFanDetails.RecordVersion = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicFanDetails.SerialNumber = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.SerialNumber
		varNiatelemetryApicFanDetails.SiteName = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicFanDetails.VendorName = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.VendorName
		varNiatelemetryApicFanDetails.Vid = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.Vid
		varNiatelemetryApicFanDetails.RegisteredDevice = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicFanDetails(varNiatelemetryApicFanDetails)
	} else {
		return err
	}

	varNiatelemetryApicFanDetails := _NiatelemetryApicFanDetails{}

	err = json.Unmarshal(data, &varNiatelemetryApicFanDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicFanDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "ModelNumber")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "VendorName")
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

type NullableNiatelemetryApicFanDetails struct {
	value *NiatelemetryApicFanDetails
	isSet bool
}

func (v NullableNiatelemetryApicFanDetails) Get() *NiatelemetryApicFanDetails {
	return v.value
}

func (v *NullableNiatelemetryApicFanDetails) Set(val *NiatelemetryApicFanDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicFanDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicFanDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicFanDetails(val *NiatelemetryApicFanDetails) *NullableNiatelemetryApicFanDetails {
	return &NullableNiatelemetryApicFanDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicFanDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicFanDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
