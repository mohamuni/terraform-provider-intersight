/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KvmVmConsoleAllOf Definition of the list of properties defined in 'kvm.VmConsole', excluding properties defined in parent classes.
type KvmVmConsoleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The URL of the KVM MUX to connect to.
	KvmMuxUrl            *string                                  `json:"KvmMuxUrl,omitempty"`
	VirtualMachine       *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmVmConsoleAllOf KvmVmConsoleAllOf

// NewKvmVmConsoleAllOf instantiates a new KvmVmConsoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmVmConsoleAllOf(classId string, objectType string) *KvmVmConsoleAllOf {
	this := KvmVmConsoleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmVmConsoleAllOfWithDefaults instantiates a new KvmVmConsoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmVmConsoleAllOfWithDefaults() *KvmVmConsoleAllOf {
	this := KvmVmConsoleAllOf{}
	var classId string = "kvm.VmConsole"
	this.ClassId = classId
	var objectType string = "kvm.VmConsole"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmVmConsoleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmVmConsoleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmVmConsoleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmVmConsoleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmVmConsoleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmVmConsoleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKvmMuxUrl returns the KvmMuxUrl field value if set, zero value otherwise.
func (o *KvmVmConsoleAllOf) GetKvmMuxUrl() string {
	if o == nil || o.KvmMuxUrl == nil {
		var ret string
		return ret
	}
	return *o.KvmMuxUrl
}

// GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmVmConsoleAllOf) GetKvmMuxUrlOk() (*string, bool) {
	if o == nil || o.KvmMuxUrl == nil {
		return nil, false
	}
	return o.KvmMuxUrl, true
}

// HasKvmMuxUrl returns a boolean if a field has been set.
func (o *KvmVmConsoleAllOf) HasKvmMuxUrl() bool {
	if o != nil && o.KvmMuxUrl != nil {
		return true
	}

	return false
}

// SetKvmMuxUrl gets a reference to the given string and assigns it to the KvmMuxUrl field.
func (o *KvmVmConsoleAllOf) SetKvmMuxUrl(v string) {
	o.KvmMuxUrl = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *KvmVmConsoleAllOf) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret HyperflexHxapVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmVmConsoleAllOf) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *KvmVmConsoleAllOf) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given HyperflexHxapVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *KvmVmConsoleAllOf) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o KvmVmConsoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KvmMuxUrl != nil {
		toSerialize["KvmMuxUrl"] = o.KvmMuxUrl
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmVmConsoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKvmVmConsoleAllOf := _KvmVmConsoleAllOf{}

	if err = json.Unmarshal(bytes, &varKvmVmConsoleAllOf); err == nil {
		*o = KvmVmConsoleAllOf(varKvmVmConsoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KvmMuxUrl")
		delete(additionalProperties, "VirtualMachine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKvmVmConsoleAllOf struct {
	value *KvmVmConsoleAllOf
	isSet bool
}

func (v NullableKvmVmConsoleAllOf) Get() *KvmVmConsoleAllOf {
	return v.value
}

func (v *NullableKvmVmConsoleAllOf) Set(val *KvmVmConsoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmVmConsoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmVmConsoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmVmConsoleAllOf(val *KvmVmConsoleAllOf) *NullableKvmVmConsoleAllOf {
	return &NullableKvmVmConsoleAllOf{value: val, isSet: true}
}

func (v NullableKvmVmConsoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmVmConsoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
