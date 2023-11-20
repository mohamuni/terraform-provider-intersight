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
	"reflect"
	"strings"
)

// NiatelemetryMdsNeighbors MdsNeighbors object available per device scope for neighbor discovery.
type NiatelemetryMdsNeighbors struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Ip address of device being inventoried.
	DeviceIp *string `json:"DeviceIp,omitempty"`
	// Device name of device being inventoried.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Device wwn of device being inventoried.
	DeviceWwn    *string                       `json:"DeviceWwn,omitempty"`
	NeighborInfo []NiatelemetryMdsNeighborInfo `json:"NeighborInfo,omitempty"`
	// Type of record MDS. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of device being inventoried. The serial number is unique per device.
	Serial               *string                              `json:"Serial,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMdsNeighbors NiatelemetryMdsNeighbors

// NewNiatelemetryMdsNeighbors instantiates a new NiatelemetryMdsNeighbors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMdsNeighbors(classId string, objectType string) *NiatelemetryMdsNeighbors {
	this := NiatelemetryMdsNeighbors{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMdsNeighborsWithDefaults instantiates a new NiatelemetryMdsNeighbors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMdsNeighborsWithDefaults() *NiatelemetryMdsNeighbors {
	this := NiatelemetryMdsNeighbors{}
	var classId string = "niatelemetry.MdsNeighbors"
	this.ClassId = classId
	var objectType string = "niatelemetry.MdsNeighbors"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMdsNeighbors) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMdsNeighbors) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMdsNeighbors) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMdsNeighbors) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceIp returns the DeviceIp field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetDeviceIp() string {
	if o == nil || o.DeviceIp == nil {
		var ret string
		return ret
	}
	return *o.DeviceIp
}

// GetDeviceIpOk returns a tuple with the DeviceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetDeviceIpOk() (*string, bool) {
	if o == nil || o.DeviceIp == nil {
		return nil, false
	}
	return o.DeviceIp, true
}

// HasDeviceIp returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasDeviceIp() bool {
	if o != nil && o.DeviceIp != nil {
		return true
	}

	return false
}

// SetDeviceIp gets a reference to the given string and assigns it to the DeviceIp field.
func (o *NiatelemetryMdsNeighbors) SetDeviceIp(v string) {
	o.DeviceIp = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *NiatelemetryMdsNeighbors) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceWwn returns the DeviceWwn field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetDeviceWwn() string {
	if o == nil || o.DeviceWwn == nil {
		var ret string
		return ret
	}
	return *o.DeviceWwn
}

// GetDeviceWwnOk returns a tuple with the DeviceWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetDeviceWwnOk() (*string, bool) {
	if o == nil || o.DeviceWwn == nil {
		return nil, false
	}
	return o.DeviceWwn, true
}

// HasDeviceWwn returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasDeviceWwn() bool {
	if o != nil && o.DeviceWwn != nil {
		return true
	}

	return false
}

// SetDeviceWwn gets a reference to the given string and assigns it to the DeviceWwn field.
func (o *NiatelemetryMdsNeighbors) SetDeviceWwn(v string) {
	o.DeviceWwn = &v
}

// GetNeighborInfo returns the NeighborInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryMdsNeighbors) GetNeighborInfo() []NiatelemetryMdsNeighborInfo {
	if o == nil {
		var ret []NiatelemetryMdsNeighborInfo
		return ret
	}
	return o.NeighborInfo
}

// GetNeighborInfoOk returns a tuple with the NeighborInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryMdsNeighbors) GetNeighborInfoOk() ([]NiatelemetryMdsNeighborInfo, bool) {
	if o == nil || o.NeighborInfo == nil {
		return nil, false
	}
	return o.NeighborInfo, true
}

// HasNeighborInfo returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasNeighborInfo() bool {
	if o != nil && o.NeighborInfo != nil {
		return true
	}

	return false
}

// SetNeighborInfo gets a reference to the given []NiatelemetryMdsNeighborInfo and assigns it to the NeighborInfo field.
func (o *NiatelemetryMdsNeighbors) SetNeighborInfo(v []NiatelemetryMdsNeighborInfo) {
	o.NeighborInfo = v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryMdsNeighbors) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryMdsNeighbors) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryMdsNeighbors) SetSerial(v string) {
	o.Serial = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighbors) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighbors) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighbors) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMdsNeighbors) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMdsNeighbors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceIp != nil {
		toSerialize["DeviceIp"] = o.DeviceIp
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.DeviceWwn != nil {
		toSerialize["DeviceWwn"] = o.DeviceWwn
	}
	if o.NeighborInfo != nil {
		toSerialize["NeighborInfo"] = o.NeighborInfo
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryMdsNeighbors) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryMdsNeighborsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Ip address of device being inventoried.
		DeviceIp *string `json:"DeviceIp,omitempty"`
		// Device name of device being inventoried.
		DeviceName *string `json:"DeviceName,omitempty"`
		// Device wwn of device being inventoried.
		DeviceWwn    *string                       `json:"DeviceWwn,omitempty"`
		NeighborInfo []NiatelemetryMdsNeighborInfo `json:"NeighborInfo,omitempty"`
		// Type of record MDS. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of device being inventoried. The serial number is unique per device.
		Serial           *string                              `json:"Serial,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryMdsNeighborsWithoutEmbeddedStruct := NiatelemetryMdsNeighborsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryMdsNeighborsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMdsNeighbors := _NiatelemetryMdsNeighbors{}
		varNiatelemetryMdsNeighbors.ClassId = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.ClassId
		varNiatelemetryMdsNeighbors.ObjectType = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMdsNeighbors.DeviceIp = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.DeviceIp
		varNiatelemetryMdsNeighbors.DeviceName = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.DeviceName
		varNiatelemetryMdsNeighbors.DeviceWwn = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.DeviceWwn
		varNiatelemetryMdsNeighbors.NeighborInfo = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.NeighborInfo
		varNiatelemetryMdsNeighbors.RecordType = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.RecordType
		varNiatelemetryMdsNeighbors.RecordVersion = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryMdsNeighbors.Serial = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.Serial
		varNiatelemetryMdsNeighbors.RegisteredDevice = varNiatelemetryMdsNeighborsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryMdsNeighbors(varNiatelemetryMdsNeighbors)
	} else {
		return err
	}

	varNiatelemetryMdsNeighbors := _NiatelemetryMdsNeighbors{}

	err = json.Unmarshal(bytes, &varNiatelemetryMdsNeighbors)
	if err == nil {
		o.MoBaseMo = varNiatelemetryMdsNeighbors.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceIp")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "DeviceWwn")
		delete(additionalProperties, "NeighborInfo")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "Serial")
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

type NullableNiatelemetryMdsNeighbors struct {
	value *NiatelemetryMdsNeighbors
	isSet bool
}

func (v NullableNiatelemetryMdsNeighbors) Get() *NiatelemetryMdsNeighbors {
	return v.value
}

func (v *NullableNiatelemetryMdsNeighbors) Set(val *NiatelemetryMdsNeighbors) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMdsNeighbors) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMdsNeighbors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMdsNeighbors(val *NiatelemetryMdsNeighbors) *NullableNiatelemetryMdsNeighbors {
	return &NullableNiatelemetryMdsNeighbors{value: val, isSet: true}
}

func (v NullableNiatelemetryMdsNeighbors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMdsNeighbors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
