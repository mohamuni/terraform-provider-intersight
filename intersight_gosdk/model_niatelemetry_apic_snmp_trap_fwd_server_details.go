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

// checks if the NiatelemetryApicSnmpTrapFwdServerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryApicSnmpTrapFwdServerDetails{}

// NiatelemetryApicSnmpTrapFwdServerDetails Object to capture the SNMP Trap Fwd Server details in APIC.
type NiatelemetryApicSnmpTrapFwdServerDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Address of SNMP Trap Fwd Server in APIC.
	Address *string `json:"Address,omitempty"`
	// Dn of the SNMP Trap Fwd Server details in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Dn of the parent SNMP Policy in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                                     `json:"SiteName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSnmpTrapFwdServerDetails NiatelemetryApicSnmpTrapFwdServerDetails

// NewNiatelemetryApicSnmpTrapFwdServerDetails instantiates a new NiatelemetryApicSnmpTrapFwdServerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSnmpTrapFwdServerDetails(classId string, objectType string) *NiatelemetryApicSnmpTrapFwdServerDetails {
	this := NiatelemetryApicSnmpTrapFwdServerDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSnmpTrapFwdServerDetailsWithDefaults instantiates a new NiatelemetryApicSnmpTrapFwdServerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSnmpTrapFwdServerDetailsWithDefaults() *NiatelemetryApicSnmpTrapFwdServerDetails {
	this := NiatelemetryApicSnmpTrapFwdServerDetails{}
	var classId string = "niatelemetry.ApicSnmpTrapFwdServerDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSnmpTrapFwdServerDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.ApicSnmpTrapFwdServerDetails" of the ClassId field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.ApicSnmpTrapFwdServerDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.ApicSnmpTrapFwdServerDetails" of the ObjectType field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.ApicSnmpTrapFwdServerDetails"
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetAddress(v string) {
	o.Address = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetDn(v string) {
	o.Dn = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetPolDn() string {
	if o == nil || IsNil(o.PolDn) {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetPolDnOk() (*string, bool) {
	if o == nil || IsNil(o.PolDn) {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasPolDn() bool {
	if o != nil && !IsNil(o.PolDn) {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryApicSnmpTrapFwdServerDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryApicSnmpTrapFwdServerDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryApicSnmpTrapFwdServerDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Address) {
		toSerialize["Address"] = o.Address
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.PolDn) {
		toSerialize["PolDn"] = o.PolDn
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if !IsNil(o.SiteName) {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryApicSnmpTrapFwdServerDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Address of SNMP Trap Fwd Server in APIC.
		Address *string `json:"Address,omitempty"`
		// Dn of the SNMP Trap Fwd Server details in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Dn of the parent SNMP Policy in APIC.
		PolDn *string `json:"PolDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                                     `json:"SiteName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct := NiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicSnmpTrapFwdServerDetails := _NiatelemetryApicSnmpTrapFwdServerDetails{}
		varNiatelemetryApicSnmpTrapFwdServerDetails.ClassId = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicSnmpTrapFwdServerDetails.ObjectType = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicSnmpTrapFwdServerDetails.Address = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.Address
		varNiatelemetryApicSnmpTrapFwdServerDetails.Dn = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicSnmpTrapFwdServerDetails.PolDn = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.PolDn
		varNiatelemetryApicSnmpTrapFwdServerDetails.RecordType = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicSnmpTrapFwdServerDetails.RecordVersion = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicSnmpTrapFwdServerDetails.SiteName = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicSnmpTrapFwdServerDetails.RegisteredDevice = varNiatelemetryApicSnmpTrapFwdServerDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicSnmpTrapFwdServerDetails(varNiatelemetryApicSnmpTrapFwdServerDetails)
	} else {
		return err
	}

	varNiatelemetryApicSnmpTrapFwdServerDetails := _NiatelemetryApicSnmpTrapFwdServerDetails{}

	err = json.Unmarshal(data, &varNiatelemetryApicSnmpTrapFwdServerDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicSnmpTrapFwdServerDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
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

type NullableNiatelemetryApicSnmpTrapFwdServerDetails struct {
	value *NiatelemetryApicSnmpTrapFwdServerDetails
	isSet bool
}

func (v NullableNiatelemetryApicSnmpTrapFwdServerDetails) Get() *NiatelemetryApicSnmpTrapFwdServerDetails {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpTrapFwdServerDetails) Set(val *NiatelemetryApicSnmpTrapFwdServerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpTrapFwdServerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpTrapFwdServerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpTrapFwdServerDetails(val *NiatelemetryApicSnmpTrapFwdServerDetails) *NullableNiatelemetryApicSnmpTrapFwdServerDetails {
	return &NullableNiatelemetryApicSnmpTrapFwdServerDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpTrapFwdServerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpTrapFwdServerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
