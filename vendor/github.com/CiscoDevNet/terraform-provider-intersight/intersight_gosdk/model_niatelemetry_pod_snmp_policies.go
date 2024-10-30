/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the NiatelemetryPodSnmpPolicies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryPodSnmpPolicies{}

// NiatelemetryPodSnmpPolicies Object to capture Pod SNMP Policy details.
type NiatelemetryPodSnmpPolicies struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin State of the Snmp Pol in APIC.
	AdminSt *string `json:"AdminSt,omitempty"`
	// Dn of the Snmp Pol in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// List of Dn of the SNMP Client grp in APIC.
	SnmpClientGrp *string `json:"SnmpClientGrp,omitempty"`
	// List of Dn of the SNMP Community in APIC.
	SnmpCommunity *string `json:"SnmpCommunity,omitempty"`
	// List of Dn of the SNMP Trap Fwd Server in APIC.
	SnmpTrapFwdServer *string `json:"SnmpTrapFwdServer,omitempty"`
	// List of Dn of the SNMP user in APIC.
	SnmpUser             *string                                     `json:"SnmpUser,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryPodSnmpPolicies NiatelemetryPodSnmpPolicies

// NewNiatelemetryPodSnmpPolicies instantiates a new NiatelemetryPodSnmpPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryPodSnmpPolicies(classId string, objectType string) *NiatelemetryPodSnmpPolicies {
	this := NiatelemetryPodSnmpPolicies{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryPodSnmpPoliciesWithDefaults instantiates a new NiatelemetryPodSnmpPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryPodSnmpPoliciesWithDefaults() *NiatelemetryPodSnmpPolicies {
	this := NiatelemetryPodSnmpPolicies{}
	var classId string = "niatelemetry.PodSnmpPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.PodSnmpPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryPodSnmpPolicies) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryPodSnmpPolicies) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.PodSnmpPolicies" of the ClassId field.
func (o *NiatelemetryPodSnmpPolicies) GetDefaultClassId() interface{} {
	return "niatelemetry.PodSnmpPolicies"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryPodSnmpPolicies) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryPodSnmpPolicies) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.PodSnmpPolicies" of the ObjectType field.
func (o *NiatelemetryPodSnmpPolicies) GetDefaultObjectType() interface{} {
	return "niatelemetry.PodSnmpPolicies"
}

// GetAdminSt returns the AdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetAdminSt() string {
	if o == nil || IsNil(o.AdminSt) {
		var ret string
		return ret
	}
	return *o.AdminSt
}

// GetAdminStOk returns a tuple with the AdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetAdminStOk() (*string, bool) {
	if o == nil || IsNil(o.AdminSt) {
		return nil, false
	}
	return o.AdminSt, true
}

// HasAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasAdminSt() bool {
	if o != nil && !IsNil(o.AdminSt) {
		return true
	}

	return false
}

// SetAdminSt gets a reference to the given string and assigns it to the AdminSt field.
func (o *NiatelemetryPodSnmpPolicies) SetAdminSt(v string) {
	o.AdminSt = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetPolDn() string {
	if o == nil || IsNil(o.PolDn) {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetPolDnOk() (*string, bool) {
	if o == nil || IsNil(o.PolDn) {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasPolDn() bool {
	if o != nil && !IsNil(o.PolDn) {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryPodSnmpPolicies) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryPodSnmpPolicies) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryPodSnmpPolicies) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryPodSnmpPolicies) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSnmpClientGrp returns the SnmpClientGrp field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpClientGrp() string {
	if o == nil || IsNil(o.SnmpClientGrp) {
		var ret string
		return ret
	}
	return *o.SnmpClientGrp
}

// GetSnmpClientGrpOk returns a tuple with the SnmpClientGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpClientGrpOk() (*string, bool) {
	if o == nil || IsNil(o.SnmpClientGrp) {
		return nil, false
	}
	return o.SnmpClientGrp, true
}

// HasSnmpClientGrp returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpClientGrp() bool {
	if o != nil && !IsNil(o.SnmpClientGrp) {
		return true
	}

	return false
}

// SetSnmpClientGrp gets a reference to the given string and assigns it to the SnmpClientGrp field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpClientGrp(v string) {
	o.SnmpClientGrp = &v
}

// GetSnmpCommunity returns the SnmpCommunity field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpCommunity() string {
	if o == nil || IsNil(o.SnmpCommunity) {
		var ret string
		return ret
	}
	return *o.SnmpCommunity
}

// GetSnmpCommunityOk returns a tuple with the SnmpCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpCommunityOk() (*string, bool) {
	if o == nil || IsNil(o.SnmpCommunity) {
		return nil, false
	}
	return o.SnmpCommunity, true
}

// HasSnmpCommunity returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpCommunity() bool {
	if o != nil && !IsNil(o.SnmpCommunity) {
		return true
	}

	return false
}

// SetSnmpCommunity gets a reference to the given string and assigns it to the SnmpCommunity field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpCommunity(v string) {
	o.SnmpCommunity = &v
}

// GetSnmpTrapFwdServer returns the SnmpTrapFwdServer field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpTrapFwdServer() string {
	if o == nil || IsNil(o.SnmpTrapFwdServer) {
		var ret string
		return ret
	}
	return *o.SnmpTrapFwdServer
}

// GetSnmpTrapFwdServerOk returns a tuple with the SnmpTrapFwdServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpTrapFwdServerOk() (*string, bool) {
	if o == nil || IsNil(o.SnmpTrapFwdServer) {
		return nil, false
	}
	return o.SnmpTrapFwdServer, true
}

// HasSnmpTrapFwdServer returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpTrapFwdServer() bool {
	if o != nil && !IsNil(o.SnmpTrapFwdServer) {
		return true
	}

	return false
}

// SetSnmpTrapFwdServer gets a reference to the given string and assigns it to the SnmpTrapFwdServer field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpTrapFwdServer(v string) {
	o.SnmpTrapFwdServer = &v
}

// GetSnmpUser returns the SnmpUser field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpUser() string {
	if o == nil || IsNil(o.SnmpUser) {
		var ret string
		return ret
	}
	return *o.SnmpUser
}

// GetSnmpUserOk returns a tuple with the SnmpUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpUserOk() (*string, bool) {
	if o == nil || IsNil(o.SnmpUser) {
		return nil, false
	}
	return o.SnmpUser, true
}

// HasSnmpUser returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpUser() bool {
	if o != nil && !IsNil(o.SnmpUser) {
		return true
	}

	return false
}

// SetSnmpUser gets a reference to the given string and assigns it to the SnmpUser field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpUser(v string) {
	o.SnmpUser = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryPodSnmpPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryPodSnmpPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryPodSnmpPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryPodSnmpPolicies) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryPodSnmpPolicies) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryPodSnmpPolicies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryPodSnmpPolicies) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdminSt) {
		toSerialize["AdminSt"] = o.AdminSt
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
	if !IsNil(o.SnmpClientGrp) {
		toSerialize["SnmpClientGrp"] = o.SnmpClientGrp
	}
	if !IsNil(o.SnmpCommunity) {
		toSerialize["SnmpCommunity"] = o.SnmpCommunity
	}
	if !IsNil(o.SnmpTrapFwdServer) {
		toSerialize["SnmpTrapFwdServer"] = o.SnmpTrapFwdServer
	}
	if !IsNil(o.SnmpUser) {
		toSerialize["SnmpUser"] = o.SnmpUser
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryPodSnmpPolicies) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin State of the Snmp Pol in APIC.
		AdminSt *string `json:"AdminSt,omitempty"`
		// Dn of the Snmp Pol in APIC.
		PolDn *string `json:"PolDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// List of Dn of the SNMP Client grp in APIC.
		SnmpClientGrp *string `json:"SnmpClientGrp,omitempty"`
		// List of Dn of the SNMP Community in APIC.
		SnmpCommunity *string `json:"SnmpCommunity,omitempty"`
		// List of Dn of the SNMP Trap Fwd Server in APIC.
		SnmpTrapFwdServer *string `json:"SnmpTrapFwdServer,omitempty"`
		// List of Dn of the SNMP user in APIC.
		SnmpUser         *string                                     `json:"SnmpUser,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct := NiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryPodSnmpPolicies := _NiatelemetryPodSnmpPolicies{}
		varNiatelemetryPodSnmpPolicies.ClassId = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.ClassId
		varNiatelemetryPodSnmpPolicies.ObjectType = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.ObjectType
		varNiatelemetryPodSnmpPolicies.AdminSt = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.AdminSt
		varNiatelemetryPodSnmpPolicies.PolDn = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.PolDn
		varNiatelemetryPodSnmpPolicies.RecordType = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.RecordType
		varNiatelemetryPodSnmpPolicies.RecordVersion = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryPodSnmpPolicies.SiteName = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SiteName
		varNiatelemetryPodSnmpPolicies.SnmpClientGrp = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpClientGrp
		varNiatelemetryPodSnmpPolicies.SnmpCommunity = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpCommunity
		varNiatelemetryPodSnmpPolicies.SnmpTrapFwdServer = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpTrapFwdServer
		varNiatelemetryPodSnmpPolicies.SnmpUser = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpUser
		varNiatelemetryPodSnmpPolicies.RegisteredDevice = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryPodSnmpPolicies(varNiatelemetryPodSnmpPolicies)
	} else {
		return err
	}

	varNiatelemetryPodSnmpPolicies := _NiatelemetryPodSnmpPolicies{}

	err = json.Unmarshal(data, &varNiatelemetryPodSnmpPolicies)
	if err == nil {
		o.MoBaseMo = varNiatelemetryPodSnmpPolicies.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSt")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SnmpClientGrp")
		delete(additionalProperties, "SnmpCommunity")
		delete(additionalProperties, "SnmpTrapFwdServer")
		delete(additionalProperties, "SnmpUser")
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

type NullableNiatelemetryPodSnmpPolicies struct {
	value *NiatelemetryPodSnmpPolicies
	isSet bool
}

func (v NullableNiatelemetryPodSnmpPolicies) Get() *NiatelemetryPodSnmpPolicies {
	return v.value
}

func (v *NullableNiatelemetryPodSnmpPolicies) Set(val *NiatelemetryPodSnmpPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryPodSnmpPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryPodSnmpPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryPodSnmpPolicies(val *NiatelemetryPodSnmpPolicies) *NullableNiatelemetryPodSnmpPolicies {
	return &NullableNiatelemetryPodSnmpPolicies{value: val, isSet: true}
}

func (v NullableNiatelemetryPodSnmpPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryPodSnmpPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
