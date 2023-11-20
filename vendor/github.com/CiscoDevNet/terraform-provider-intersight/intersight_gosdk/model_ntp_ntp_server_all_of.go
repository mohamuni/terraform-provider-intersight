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

// NtpNtpServerAllOf Definition of the list of properties defined in 'ntp.NtpServer', excluding properties defined in parent classes.
type NtpNtpServerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The default polling interval of the NTP server in seconds.
	Poll *int64 `json:"Poll,omitempty"`
	// The IP address of the NTP server. Supports both IPv4 or IPv6 address.
	ServerIpAddress *string `json:"ServerIpAddress,omitempty"`
	// The stratum level of the NTP server.
	Stratum *int64 `json:"Stratum,omitempty"`
	// It determines whether the IP address configured is server or peer. * `Server` - NTP configured is server type. * `Peer` - NTP configured is peer type.
	Type *string `json:"Type,omitempty"`
	// VRF name to be used by NTP Server.
	VrfName              *string                              `json:"VrfName,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NtpNtpServerAllOf NtpNtpServerAllOf

// NewNtpNtpServerAllOf instantiates a new NtpNtpServerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpNtpServerAllOf(classId string, objectType string) *NtpNtpServerAllOf {
	this := NtpNtpServerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNtpNtpServerAllOfWithDefaults instantiates a new NtpNtpServerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpNtpServerAllOfWithDefaults() *NtpNtpServerAllOf {
	this := NtpNtpServerAllOf{}
	var classId string = "ntp.NtpServer"
	this.ClassId = classId
	var objectType string = "ntp.NtpServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NtpNtpServerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NtpNtpServerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NtpNtpServerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NtpNtpServerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetPoll() int64 {
	if o == nil || o.Poll == nil {
		var ret int64
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetPollOk() (*int64, bool) {
	if o == nil || o.Poll == nil {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasPoll() bool {
	if o != nil && o.Poll != nil {
		return true
	}

	return false
}

// SetPoll gets a reference to the given int64 and assigns it to the Poll field.
func (o *NtpNtpServerAllOf) SetPoll(v int64) {
	o.Poll = &v
}

// GetServerIpAddress returns the ServerIpAddress field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetServerIpAddress() string {
	if o == nil || o.ServerIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ServerIpAddress
}

// GetServerIpAddressOk returns a tuple with the ServerIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetServerIpAddressOk() (*string, bool) {
	if o == nil || o.ServerIpAddress == nil {
		return nil, false
	}
	return o.ServerIpAddress, true
}

// HasServerIpAddress returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasServerIpAddress() bool {
	if o != nil && o.ServerIpAddress != nil {
		return true
	}

	return false
}

// SetServerIpAddress gets a reference to the given string and assigns it to the ServerIpAddress field.
func (o *NtpNtpServerAllOf) SetServerIpAddress(v string) {
	o.ServerIpAddress = &v
}

// GetStratum returns the Stratum field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetStratum() int64 {
	if o == nil || o.Stratum == nil {
		var ret int64
		return ret
	}
	return *o.Stratum
}

// GetStratumOk returns a tuple with the Stratum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetStratumOk() (*int64, bool) {
	if o == nil || o.Stratum == nil {
		return nil, false
	}
	return o.Stratum, true
}

// HasStratum returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasStratum() bool {
	if o != nil && o.Stratum != nil {
		return true
	}

	return false
}

// SetStratum gets a reference to the given int64 and assigns it to the Stratum field.
func (o *NtpNtpServerAllOf) SetStratum(v int64) {
	o.Stratum = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NtpNtpServerAllOf) SetType(v string) {
	o.Type = &v
}

// GetVrfName returns the VrfName field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetVrfName() string {
	if o == nil || o.VrfName == nil {
		var ret string
		return ret
	}
	return *o.VrfName
}

// GetVrfNameOk returns a tuple with the VrfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetVrfNameOk() (*string, bool) {
	if o == nil || o.VrfName == nil {
		return nil, false
	}
	return o.VrfName, true
}

// HasVrfName returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasVrfName() bool {
	if o != nil && o.VrfName != nil {
		return true
	}

	return false
}

// SetVrfName gets a reference to the given string and assigns it to the VrfName field.
func (o *NtpNtpServerAllOf) SetVrfName(v string) {
	o.VrfName = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NtpNtpServerAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NtpNtpServerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpNtpServerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NtpNtpServerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NtpNtpServerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NtpNtpServerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Poll != nil {
		toSerialize["Poll"] = o.Poll
	}
	if o.ServerIpAddress != nil {
		toSerialize["ServerIpAddress"] = o.ServerIpAddress
	}
	if o.Stratum != nil {
		toSerialize["Stratum"] = o.Stratum
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VrfName != nil {
		toSerialize["VrfName"] = o.VrfName
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

func (o *NtpNtpServerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNtpNtpServerAllOf := _NtpNtpServerAllOf{}

	if err = json.Unmarshal(bytes, &varNtpNtpServerAllOf); err == nil {
		*o = NtpNtpServerAllOf(varNtpNtpServerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Poll")
		delete(additionalProperties, "ServerIpAddress")
		delete(additionalProperties, "Stratum")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VrfName")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNtpNtpServerAllOf struct {
	value *NtpNtpServerAllOf
	isSet bool
}

func (v NullableNtpNtpServerAllOf) Get() *NtpNtpServerAllOf {
	return v.value
}

func (v *NullableNtpNtpServerAllOf) Set(val *NtpNtpServerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNtpNtpServerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNtpNtpServerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtpNtpServerAllOf(val *NtpNtpServerAllOf) *NullableNtpNtpServerAllOf {
	return &NullableNtpNtpServerAllOf{value: val, isSet: true}
}

func (v NullableNtpNtpServerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtpNtpServerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
