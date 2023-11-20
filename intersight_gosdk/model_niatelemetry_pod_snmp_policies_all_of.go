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

// NiatelemetryPodSnmpPoliciesAllOf Definition of the list of properties defined in 'niatelemetry.PodSnmpPolicies', excluding properties defined in parent classes.
type NiatelemetryPodSnmpPoliciesAllOf struct {
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
	SnmpUser             *string                              `json:"SnmpUser,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryPodSnmpPoliciesAllOf NiatelemetryPodSnmpPoliciesAllOf

// NewNiatelemetryPodSnmpPoliciesAllOf instantiates a new NiatelemetryPodSnmpPoliciesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryPodSnmpPoliciesAllOf(classId string, objectType string) *NiatelemetryPodSnmpPoliciesAllOf {
	this := NiatelemetryPodSnmpPoliciesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryPodSnmpPoliciesAllOfWithDefaults instantiates a new NiatelemetryPodSnmpPoliciesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryPodSnmpPoliciesAllOfWithDefaults() *NiatelemetryPodSnmpPoliciesAllOf {
	this := NiatelemetryPodSnmpPoliciesAllOf{}
	var classId string = "niatelemetry.PodSnmpPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.PodSnmpPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSt returns the AdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetAdminSt() string {
	if o == nil || o.AdminSt == nil {
		var ret string
		return ret
	}
	return *o.AdminSt
}

// GetAdminStOk returns a tuple with the AdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetAdminStOk() (*string, bool) {
	if o == nil || o.AdminSt == nil {
		return nil, false
	}
	return o.AdminSt, true
}

// HasAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasAdminSt() bool {
	if o != nil && o.AdminSt != nil {
		return true
	}

	return false
}

// SetAdminSt gets a reference to the given string and assigns it to the AdminSt field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetAdminSt(v string) {
	o.AdminSt = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSnmpClientGrp returns the SnmpClientGrp field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpClientGrp() string {
	if o == nil || o.SnmpClientGrp == nil {
		var ret string
		return ret
	}
	return *o.SnmpClientGrp
}

// GetSnmpClientGrpOk returns a tuple with the SnmpClientGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpClientGrpOk() (*string, bool) {
	if o == nil || o.SnmpClientGrp == nil {
		return nil, false
	}
	return o.SnmpClientGrp, true
}

// HasSnmpClientGrp returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpClientGrp() bool {
	if o != nil && o.SnmpClientGrp != nil {
		return true
	}

	return false
}

// SetSnmpClientGrp gets a reference to the given string and assigns it to the SnmpClientGrp field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpClientGrp(v string) {
	o.SnmpClientGrp = &v
}

// GetSnmpCommunity returns the SnmpCommunity field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpCommunity() string {
	if o == nil || o.SnmpCommunity == nil {
		var ret string
		return ret
	}
	return *o.SnmpCommunity
}

// GetSnmpCommunityOk returns a tuple with the SnmpCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpCommunityOk() (*string, bool) {
	if o == nil || o.SnmpCommunity == nil {
		return nil, false
	}
	return o.SnmpCommunity, true
}

// HasSnmpCommunity returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpCommunity() bool {
	if o != nil && o.SnmpCommunity != nil {
		return true
	}

	return false
}

// SetSnmpCommunity gets a reference to the given string and assigns it to the SnmpCommunity field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpCommunity(v string) {
	o.SnmpCommunity = &v
}

// GetSnmpTrapFwdServer returns the SnmpTrapFwdServer field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpTrapFwdServer() string {
	if o == nil || o.SnmpTrapFwdServer == nil {
		var ret string
		return ret
	}
	return *o.SnmpTrapFwdServer
}

// GetSnmpTrapFwdServerOk returns a tuple with the SnmpTrapFwdServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpTrapFwdServerOk() (*string, bool) {
	if o == nil || o.SnmpTrapFwdServer == nil {
		return nil, false
	}
	return o.SnmpTrapFwdServer, true
}

// HasSnmpTrapFwdServer returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpTrapFwdServer() bool {
	if o != nil && o.SnmpTrapFwdServer != nil {
		return true
	}

	return false
}

// SetSnmpTrapFwdServer gets a reference to the given string and assigns it to the SnmpTrapFwdServer field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpTrapFwdServer(v string) {
	o.SnmpTrapFwdServer = &v
}

// GetSnmpUser returns the SnmpUser field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpUser() string {
	if o == nil || o.SnmpUser == nil {
		var ret string
		return ret
	}
	return *o.SnmpUser
}

// GetSnmpUserOk returns a tuple with the SnmpUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpUserOk() (*string, bool) {
	if o == nil || o.SnmpUser == nil {
		return nil, false
	}
	return o.SnmpUser, true
}

// HasSnmpUser returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpUser() bool {
	if o != nil && o.SnmpUser != nil {
		return true
	}

	return false
}

// SetSnmpUser gets a reference to the given string and assigns it to the SnmpUser field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpUser(v string) {
	o.SnmpUser = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPoliciesAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryPodSnmpPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryPodSnmpPoliciesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSt != nil {
		toSerialize["AdminSt"] = o.AdminSt
	}
	if o.PolDn != nil {
		toSerialize["PolDn"] = o.PolDn
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.SnmpClientGrp != nil {
		toSerialize["SnmpClientGrp"] = o.SnmpClientGrp
	}
	if o.SnmpCommunity != nil {
		toSerialize["SnmpCommunity"] = o.SnmpCommunity
	}
	if o.SnmpTrapFwdServer != nil {
		toSerialize["SnmpTrapFwdServer"] = o.SnmpTrapFwdServer
	}
	if o.SnmpUser != nil {
		toSerialize["SnmpUser"] = o.SnmpUser
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryPodSnmpPoliciesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryPodSnmpPoliciesAllOf := _NiatelemetryPodSnmpPoliciesAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryPodSnmpPoliciesAllOf); err == nil {
		*o = NiatelemetryPodSnmpPoliciesAllOf(varNiatelemetryPodSnmpPoliciesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryPodSnmpPoliciesAllOf struct {
	value *NiatelemetryPodSnmpPoliciesAllOf
	isSet bool
}

func (v NullableNiatelemetryPodSnmpPoliciesAllOf) Get() *NiatelemetryPodSnmpPoliciesAllOf {
	return v.value
}

func (v *NullableNiatelemetryPodSnmpPoliciesAllOf) Set(val *NiatelemetryPodSnmpPoliciesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryPodSnmpPoliciesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryPodSnmpPoliciesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryPodSnmpPoliciesAllOf(val *NiatelemetryPodSnmpPoliciesAllOf) *NullableNiatelemetryPodSnmpPoliciesAllOf {
	return &NullableNiatelemetryPodSnmpPoliciesAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryPodSnmpPoliciesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryPodSnmpPoliciesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
