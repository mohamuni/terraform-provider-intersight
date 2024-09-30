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

// checks if the NiatelemetryDcnmTransceiverDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryDcnmTransceiverDetails{}

// NiatelemetryDcnmTransceiverDetails Inventory Object available for DCNM transceiver details.
type NiatelemetryDcnmTransceiverDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Part number of the transceiver in the fabric inventory.
	PartNumber *string `json:"PartNumber,omitempty"`
	// Product Id of the transceiver in the fabric inventory.
	ProductId *string `json:"ProductId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the transceiver in the fabric inventory.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Vendor Id of the transceiver in the fabric inventory.
	VendorId             *string                                     `json:"VendorId,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDcnmTransceiverDetails NiatelemetryDcnmTransceiverDetails

// NewNiatelemetryDcnmTransceiverDetails instantiates a new NiatelemetryDcnmTransceiverDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDcnmTransceiverDetails(classId string, objectType string) *NiatelemetryDcnmTransceiverDetails {
	this := NiatelemetryDcnmTransceiverDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDcnmTransceiverDetailsWithDefaults instantiates a new NiatelemetryDcnmTransceiverDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDcnmTransceiverDetailsWithDefaults() *NiatelemetryDcnmTransceiverDetails {
	this := NiatelemetryDcnmTransceiverDetails{}
	var classId string = "niatelemetry.DcnmTransceiverDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.DcnmTransceiverDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDcnmTransceiverDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDcnmTransceiverDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.DcnmTransceiverDetails" of the ClassId field.
func (o *NiatelemetryDcnmTransceiverDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.DcnmTransceiverDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDcnmTransceiverDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDcnmTransceiverDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.DcnmTransceiverDetails" of the ObjectType field.
func (o *NiatelemetryDcnmTransceiverDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.DcnmTransceiverDetails"
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetails) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *NiatelemetryDcnmTransceiverDetails) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetails) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *NiatelemetryDcnmTransceiverDetails) SetProductId(v string) {
	o.ProductId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryDcnmTransceiverDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryDcnmTransceiverDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetails) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryDcnmTransceiverDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetails) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetails) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *NiatelemetryDcnmTransceiverDetails) SetVendorId(v string) {
	o.VendorId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryDcnmTransceiverDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryDcnmTransceiverDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryDcnmTransceiverDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryDcnmTransceiverDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryDcnmTransceiverDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryDcnmTransceiverDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryDcnmTransceiverDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.ProductId) {
		toSerialize["ProductId"] = o.ProductId
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
	if !IsNil(o.VendorId) {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryDcnmTransceiverDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Part number of the transceiver in the fabric inventory.
		PartNumber *string `json:"PartNumber,omitempty"`
		// Product Id of the transceiver in the fabric inventory.
		ProductId *string `json:"ProductId,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of the transceiver in the fabric inventory.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// Vendor Id of the transceiver in the fabric inventory.
		VendorId         *string                                     `json:"VendorId,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct := NiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryDcnmTransceiverDetails := _NiatelemetryDcnmTransceiverDetails{}
		varNiatelemetryDcnmTransceiverDetails.ClassId = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryDcnmTransceiverDetails.ObjectType = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryDcnmTransceiverDetails.PartNumber = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.PartNumber
		varNiatelemetryDcnmTransceiverDetails.ProductId = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.ProductId
		varNiatelemetryDcnmTransceiverDetails.RecordType = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryDcnmTransceiverDetails.RecordVersion = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryDcnmTransceiverDetails.SerialNumber = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.SerialNumber
		varNiatelemetryDcnmTransceiverDetails.VendorId = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.VendorId
		varNiatelemetryDcnmTransceiverDetails.RegisteredDevice = varNiatelemetryDcnmTransceiverDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryDcnmTransceiverDetails(varNiatelemetryDcnmTransceiverDetails)
	} else {
		return err
	}

	varNiatelemetryDcnmTransceiverDetails := _NiatelemetryDcnmTransceiverDetails{}

	err = json.Unmarshal(data, &varNiatelemetryDcnmTransceiverDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryDcnmTransceiverDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "ProductId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "VendorId")
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

type NullableNiatelemetryDcnmTransceiverDetails struct {
	value *NiatelemetryDcnmTransceiverDetails
	isSet bool
}

func (v NullableNiatelemetryDcnmTransceiverDetails) Get() *NiatelemetryDcnmTransceiverDetails {
	return v.value
}

func (v *NullableNiatelemetryDcnmTransceiverDetails) Set(val *NiatelemetryDcnmTransceiverDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDcnmTransceiverDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDcnmTransceiverDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDcnmTransceiverDetails(val *NiatelemetryDcnmTransceiverDetails) *NullableNiatelemetryDcnmTransceiverDetails {
	return &NullableNiatelemetryDcnmTransceiverDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryDcnmTransceiverDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDcnmTransceiverDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
