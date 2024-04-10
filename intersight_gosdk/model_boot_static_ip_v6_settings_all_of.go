/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// BootStaticIpV6SettingsAllOf Definition of the list of properties defined in 'boot.StaticIpV6Settings', excluding properties defined in parent classes.
type BootStaticIpV6SettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of DNS server.
	DnsIp *string `json:"DnsIp,omitempty"`
	// IP address of default gateway.
	GatewayIp *string `json:"GatewayIp,omitempty"`
	// Ipv6 static Internet Protocol address.
	Ip *string `json:"Ip,omitempty"`
	// A prefix length which masks the IP address and divides the IP address into network address and host address.
	PrefixLength         *int64 `json:"PrefixLength,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootStaticIpV6SettingsAllOf BootStaticIpV6SettingsAllOf

// NewBootStaticIpV6SettingsAllOf instantiates a new BootStaticIpV6SettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootStaticIpV6SettingsAllOf(classId string, objectType string) *BootStaticIpV6SettingsAllOf {
	this := BootStaticIpV6SettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var prefixLength int64 = 1
	this.PrefixLength = &prefixLength
	return &this
}

// NewBootStaticIpV6SettingsAllOfWithDefaults instantiates a new BootStaticIpV6SettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootStaticIpV6SettingsAllOfWithDefaults() *BootStaticIpV6SettingsAllOf {
	this := BootStaticIpV6SettingsAllOf{}
	var classId string = "boot.StaticIpV6Settings"
	this.ClassId = classId
	var objectType string = "boot.StaticIpV6Settings"
	this.ObjectType = objectType
	var prefixLength int64 = 1
	this.PrefixLength = &prefixLength
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootStaticIpV6SettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6SettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootStaticIpV6SettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootStaticIpV6SettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6SettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootStaticIpV6SettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDnsIp returns the DnsIp field value if set, zero value otherwise.
func (o *BootStaticIpV6SettingsAllOf) GetDnsIp() string {
	if o == nil || o.DnsIp == nil {
		var ret string
		return ret
	}
	return *o.DnsIp
}

// GetDnsIpOk returns a tuple with the DnsIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6SettingsAllOf) GetDnsIpOk() (*string, bool) {
	if o == nil || o.DnsIp == nil {
		return nil, false
	}
	return o.DnsIp, true
}

// HasDnsIp returns a boolean if a field has been set.
func (o *BootStaticIpV6SettingsAllOf) HasDnsIp() bool {
	if o != nil && o.DnsIp != nil {
		return true
	}

	return false
}

// SetDnsIp gets a reference to the given string and assigns it to the DnsIp field.
func (o *BootStaticIpV6SettingsAllOf) SetDnsIp(v string) {
	o.DnsIp = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *BootStaticIpV6SettingsAllOf) GetGatewayIp() string {
	if o == nil || o.GatewayIp == nil {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6SettingsAllOf) GetGatewayIpOk() (*string, bool) {
	if o == nil || o.GatewayIp == nil {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *BootStaticIpV6SettingsAllOf) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *BootStaticIpV6SettingsAllOf) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *BootStaticIpV6SettingsAllOf) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6SettingsAllOf) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *BootStaticIpV6SettingsAllOf) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *BootStaticIpV6SettingsAllOf) SetIp(v string) {
	o.Ip = &v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *BootStaticIpV6SettingsAllOf) GetPrefixLength() int64 {
	if o == nil || o.PrefixLength == nil {
		var ret int64
		return ret
	}
	return *o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6SettingsAllOf) GetPrefixLengthOk() (*int64, bool) {
	if o == nil || o.PrefixLength == nil {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *BootStaticIpV6SettingsAllOf) HasPrefixLength() bool {
	if o != nil && o.PrefixLength != nil {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given int64 and assigns it to the PrefixLength field.
func (o *BootStaticIpV6SettingsAllOf) SetPrefixLength(v int64) {
	o.PrefixLength = &v
}

func (o BootStaticIpV6SettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DnsIp != nil {
		toSerialize["DnsIp"] = o.DnsIp
	}
	if o.GatewayIp != nil {
		toSerialize["GatewayIp"] = o.GatewayIp
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.PrefixLength != nil {
		toSerialize["PrefixLength"] = o.PrefixLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootStaticIpV6SettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootStaticIpV6SettingsAllOf := _BootStaticIpV6SettingsAllOf{}

	if err = json.Unmarshal(bytes, &varBootStaticIpV6SettingsAllOf); err == nil {
		*o = BootStaticIpV6SettingsAllOf(varBootStaticIpV6SettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DnsIp")
		delete(additionalProperties, "GatewayIp")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "PrefixLength")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootStaticIpV6SettingsAllOf struct {
	value *BootStaticIpV6SettingsAllOf
	isSet bool
}

func (v NullableBootStaticIpV6SettingsAllOf) Get() *BootStaticIpV6SettingsAllOf {
	return v.value
}

func (v *NullableBootStaticIpV6SettingsAllOf) Set(val *BootStaticIpV6SettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootStaticIpV6SettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootStaticIpV6SettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootStaticIpV6SettingsAllOf(val *BootStaticIpV6SettingsAllOf) *NullableBootStaticIpV6SettingsAllOf {
	return &NullableBootStaticIpV6SettingsAllOf{value: val, isSet: true}
}

func (v NullableBootStaticIpV6SettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootStaticIpV6SettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}