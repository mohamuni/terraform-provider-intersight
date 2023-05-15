/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NetworkLicenseFileAllOf Definition of the list of properties defined in 'network.LicenseFile', excluding properties defined in parent classes.
type NetworkLicenseFileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The expiry date of the license feature.
	ExpiryDate *string `json:"ExpiryDate,omitempty"`
	// The name of the license feature.
	FeatureName *string `json:"FeatureName,omitempty"`
	// The file Id of the license file.
	FileId *string `json:"FileId,omitempty"`
	// The identifier to identify license host Id.
	HostId *string `json:"HostId,omitempty"`
	// The vendor of the license.
	Vendor               *string                              `json:"Vendor,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkLicenseFileAllOf NetworkLicenseFileAllOf

// NewNetworkLicenseFileAllOf instantiates a new NetworkLicenseFileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLicenseFileAllOf(classId string, objectType string) *NetworkLicenseFileAllOf {
	this := NetworkLicenseFileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkLicenseFileAllOfWithDefaults instantiates a new NetworkLicenseFileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLicenseFileAllOfWithDefaults() *NetworkLicenseFileAllOf {
	this := NetworkLicenseFileAllOf{}
	var classId string = "network.LicenseFile"
	this.ClassId = classId
	var objectType string = "network.LicenseFile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkLicenseFileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkLicenseFileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkLicenseFileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkLicenseFileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetExpiryDate() string {
	if o == nil || o.ExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetExpiryDateOk() (*string, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *NetworkLicenseFileAllOf) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetFeatureName returns the FeatureName field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetFeatureName() string {
	if o == nil || o.FeatureName == nil {
		var ret string
		return ret
	}
	return *o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetFeatureNameOk() (*string, bool) {
	if o == nil || o.FeatureName == nil {
		return nil, false
	}
	return o.FeatureName, true
}

// HasFeatureName returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasFeatureName() bool {
	if o != nil && o.FeatureName != nil {
		return true
	}

	return false
}

// SetFeatureName gets a reference to the given string and assigns it to the FeatureName field.
func (o *NetworkLicenseFileAllOf) SetFeatureName(v string) {
	o.FeatureName = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetFileId() string {
	if o == nil || o.FileId == nil {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetFileIdOk() (*string, bool) {
	if o == nil || o.FileId == nil {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasFileId() bool {
	if o != nil && o.FileId != nil {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *NetworkLicenseFileAllOf) SetFileId(v string) {
	o.FileId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *NetworkLicenseFileAllOf) SetHostId(v string) {
	o.HostId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *NetworkLicenseFileAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkLicenseFileAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkLicenseFileAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLicenseFileAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkLicenseFileAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkLicenseFileAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkLicenseFileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExpiryDate != nil {
		toSerialize["ExpiryDate"] = o.ExpiryDate
	}
	if o.FeatureName != nil {
		toSerialize["FeatureName"] = o.FeatureName
	}
	if o.FileId != nil {
		toSerialize["FileId"] = o.FileId
	}
	if o.HostId != nil {
		toSerialize["HostId"] = o.HostId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkLicenseFileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkLicenseFileAllOf := _NetworkLicenseFileAllOf{}

	if err = json.Unmarshal(bytes, &varNetworkLicenseFileAllOf); err == nil {
		*o = NetworkLicenseFileAllOf(varNetworkLicenseFileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExpiryDate")
		delete(additionalProperties, "FeatureName")
		delete(additionalProperties, "FileId")
		delete(additionalProperties, "HostId")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkLicenseFileAllOf struct {
	value *NetworkLicenseFileAllOf
	isSet bool
}

func (v NullableNetworkLicenseFileAllOf) Get() *NetworkLicenseFileAllOf {
	return v.value
}

func (v *NullableNetworkLicenseFileAllOf) Set(val *NetworkLicenseFileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLicenseFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLicenseFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLicenseFileAllOf(val *NetworkLicenseFileAllOf) *NullableNetworkLicenseFileAllOf {
	return &NullableNetworkLicenseFileAllOf{value: val, isSet: true}
}

func (v NullableNetworkLicenseFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLicenseFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}