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

// IpmioverlanPolicyInventoryAllOf Definition of the list of properties defined in 'ipmioverlan.PolicyInventory', excluding properties defined in parent classes.
type IpmioverlanPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the IPMI Over LAN service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. Use “00” to disable encryption key use. This configuration is supported by all standalone rack servers. FI-attached rack servers with firmware at minimum of 4.2.3a support this configuration. FI-attached blade servers do not support an encryption key. IPMI commands using this key should append zeroes to the key to achieve a length of 40 characters.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// Indicates whether the value of the 'encryptionKey' property has been set.
	IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`
	// The highest privilege level that can be assigned to an IPMI session on a server. This configuration is supported by all standalone rack servers. FI-attached rack servers with firmware at minimum of 4.2.3a support this configuration. FI-attached blade servers do not support privilege level. Privilege level will be ignored for unsupported servers. * `admin` - Privilege to perform all actions available through IPMI. * `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks. * `read-only` - Privilege to view information throught IPMI but restriction on making any changes.
	Privilege            *string               `json:"Privilege,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IpmioverlanPolicyInventoryAllOf IpmioverlanPolicyInventoryAllOf

// NewIpmioverlanPolicyInventoryAllOf instantiates a new IpmioverlanPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmioverlanPolicyInventoryAllOf(classId string, objectType string) *IpmioverlanPolicyInventoryAllOf {
	this := IpmioverlanPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIpmioverlanPolicyInventoryAllOfWithDefaults instantiates a new IpmioverlanPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmioverlanPolicyInventoryAllOfWithDefaults() *IpmioverlanPolicyInventoryAllOf {
	this := IpmioverlanPolicyInventoryAllOf{}
	var classId string = "ipmioverlan.PolicyInventory"
	this.ClassId = classId
	var objectType string = "ipmioverlan.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IpmioverlanPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IpmioverlanPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IpmioverlanPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IpmioverlanPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventoryAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventoryAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IpmioverlanPolicyInventoryAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventoryAllOf) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventoryAllOf) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *IpmioverlanPolicyInventoryAllOf) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetIsEncryptionKeySet returns the IsEncryptionKeySet field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventoryAllOf) GetIsEncryptionKeySet() bool {
	if o == nil || o.IsEncryptionKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsEncryptionKeySet
}

// GetIsEncryptionKeySetOk returns a tuple with the IsEncryptionKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetIsEncryptionKeySetOk() (*bool, bool) {
	if o == nil || o.IsEncryptionKeySet == nil {
		return nil, false
	}
	return o.IsEncryptionKeySet, true
}

// HasIsEncryptionKeySet returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventoryAllOf) HasIsEncryptionKeySet() bool {
	if o != nil && o.IsEncryptionKeySet != nil {
		return true
	}

	return false
}

// SetIsEncryptionKeySet gets a reference to the given bool and assigns it to the IsEncryptionKeySet field.
func (o *IpmioverlanPolicyInventoryAllOf) SetIsEncryptionKeySet(v bool) {
	o.IsEncryptionKeySet = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventoryAllOf) GetPrivilege() string {
	if o == nil || o.Privilege == nil {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetPrivilegeOk() (*string, bool) {
	if o == nil || o.Privilege == nil {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventoryAllOf) HasPrivilege() bool {
	if o != nil && o.Privilege != nil {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *IpmioverlanPolicyInventoryAllOf) SetPrivilege(v string) {
	o.Privilege = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IpmioverlanPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o IpmioverlanPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.EncryptionKey != nil {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.IsEncryptionKeySet != nil {
		toSerialize["IsEncryptionKeySet"] = o.IsEncryptionKeySet
	}
	if o.Privilege != nil {
		toSerialize["Privilege"] = o.Privilege
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IpmioverlanPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIpmioverlanPolicyInventoryAllOf := _IpmioverlanPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varIpmioverlanPolicyInventoryAllOf); err == nil {
		*o = IpmioverlanPolicyInventoryAllOf(varIpmioverlanPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "EncryptionKey")
		delete(additionalProperties, "IsEncryptionKeySet")
		delete(additionalProperties, "Privilege")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpmioverlanPolicyInventoryAllOf struct {
	value *IpmioverlanPolicyInventoryAllOf
	isSet bool
}

func (v NullableIpmioverlanPolicyInventoryAllOf) Get() *IpmioverlanPolicyInventoryAllOf {
	return v.value
}

func (v *NullableIpmioverlanPolicyInventoryAllOf) Set(val *IpmioverlanPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmioverlanPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmioverlanPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmioverlanPolicyInventoryAllOf(val *IpmioverlanPolicyInventoryAllOf) *NullableIpmioverlanPolicyInventoryAllOf {
	return &NullableIpmioverlanPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableIpmioverlanPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmioverlanPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}