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

// CommIpV4AddressBlockAllOf Definition of the list of properties defined in 'comm.IpV4AddressBlock', excluding properties defined in parent classes.
type CommIpV4AddressBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The end address of the IPv4 block.
	EndAddress *string `json:"EndAddress,omitempty"`
	// The start address of the IPv4 block.
	StartAddress         *string `json:"StartAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommIpV4AddressBlockAllOf CommIpV4AddressBlockAllOf

// NewCommIpV4AddressBlockAllOf instantiates a new CommIpV4AddressBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommIpV4AddressBlockAllOf(classId string, objectType string) *CommIpV4AddressBlockAllOf {
	this := CommIpV4AddressBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCommIpV4AddressBlockAllOfWithDefaults instantiates a new CommIpV4AddressBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommIpV4AddressBlockAllOfWithDefaults() *CommIpV4AddressBlockAllOf {
	this := CommIpV4AddressBlockAllOf{}
	var classId string = "comm.IpV4AddressBlock"
	this.ClassId = classId
	var objectType string = "comm.IpV4AddressBlock"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommIpV4AddressBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommIpV4AddressBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommIpV4AddressBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommIpV4AddressBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommIpV4AddressBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommIpV4AddressBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *CommIpV4AddressBlockAllOf) GetEndAddress() string {
	if o == nil || o.EndAddress == nil {
		var ret string
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV4AddressBlockAllOf) GetEndAddressOk() (*string, bool) {
	if o == nil || o.EndAddress == nil {
		return nil, false
	}
	return o.EndAddress, true
}

// HasEndAddress returns a boolean if a field has been set.
func (o *CommIpV4AddressBlockAllOf) HasEndAddress() bool {
	if o != nil && o.EndAddress != nil {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given string and assigns it to the EndAddress field.
func (o *CommIpV4AddressBlockAllOf) SetEndAddress(v string) {
	o.EndAddress = &v
}

// GetStartAddress returns the StartAddress field value if set, zero value otherwise.
func (o *CommIpV4AddressBlockAllOf) GetStartAddress() string {
	if o == nil || o.StartAddress == nil {
		var ret string
		return ret
	}
	return *o.StartAddress
}

// GetStartAddressOk returns a tuple with the StartAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV4AddressBlockAllOf) GetStartAddressOk() (*string, bool) {
	if o == nil || o.StartAddress == nil {
		return nil, false
	}
	return o.StartAddress, true
}

// HasStartAddress returns a boolean if a field has been set.
func (o *CommIpV4AddressBlockAllOf) HasStartAddress() bool {
	if o != nil && o.StartAddress != nil {
		return true
	}

	return false
}

// SetStartAddress gets a reference to the given string and assigns it to the StartAddress field.
func (o *CommIpV4AddressBlockAllOf) SetStartAddress(v string) {
	o.StartAddress = &v
}

func (o CommIpV4AddressBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EndAddress != nil {
		toSerialize["EndAddress"] = o.EndAddress
	}
	if o.StartAddress != nil {
		toSerialize["StartAddress"] = o.StartAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommIpV4AddressBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCommIpV4AddressBlockAllOf := _CommIpV4AddressBlockAllOf{}

	if err = json.Unmarshal(bytes, &varCommIpV4AddressBlockAllOf); err == nil {
		*o = CommIpV4AddressBlockAllOf(varCommIpV4AddressBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndAddress")
		delete(additionalProperties, "StartAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommIpV4AddressBlockAllOf struct {
	value *CommIpV4AddressBlockAllOf
	isSet bool
}

func (v NullableCommIpV4AddressBlockAllOf) Get() *CommIpV4AddressBlockAllOf {
	return v.value
}

func (v *NullableCommIpV4AddressBlockAllOf) Set(val *CommIpV4AddressBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommIpV4AddressBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommIpV4AddressBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommIpV4AddressBlockAllOf(val *CommIpV4AddressBlockAllOf) *NullableCommIpV4AddressBlockAllOf {
	return &NullableCommIpV4AddressBlockAllOf{value: val, isSet: true}
}

func (v NullableCommIpV4AddressBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommIpV4AddressBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
