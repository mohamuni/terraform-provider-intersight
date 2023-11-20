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
	"reflect"
	"strings"
)

// IpmioverlanPolicyInventory Intelligent Platform Management Interface Over LAN Policy.
type IpmioverlanPolicyInventory struct {
	PolicyAbstractPolicyInventory
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

type _IpmioverlanPolicyInventory IpmioverlanPolicyInventory

// NewIpmioverlanPolicyInventory instantiates a new IpmioverlanPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmioverlanPolicyInventory(classId string, objectType string) *IpmioverlanPolicyInventory {
	this := IpmioverlanPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIpmioverlanPolicyInventoryWithDefaults instantiates a new IpmioverlanPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmioverlanPolicyInventoryWithDefaults() *IpmioverlanPolicyInventory {
	this := IpmioverlanPolicyInventory{}
	var classId string = "ipmioverlan.PolicyInventory"
	this.ClassId = classId
	var objectType string = "ipmioverlan.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IpmioverlanPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IpmioverlanPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IpmioverlanPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IpmioverlanPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IpmioverlanPolicyInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *IpmioverlanPolicyInventory) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetIsEncryptionKeySet returns the IsEncryptionKeySet field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetIsEncryptionKeySet() bool {
	if o == nil || o.IsEncryptionKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsEncryptionKeySet
}

// GetIsEncryptionKeySetOk returns a tuple with the IsEncryptionKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetIsEncryptionKeySetOk() (*bool, bool) {
	if o == nil || o.IsEncryptionKeySet == nil {
		return nil, false
	}
	return o.IsEncryptionKeySet, true
}

// HasIsEncryptionKeySet returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasIsEncryptionKeySet() bool {
	if o != nil && o.IsEncryptionKeySet != nil {
		return true
	}

	return false
}

// SetIsEncryptionKeySet gets a reference to the given bool and assigns it to the IsEncryptionKeySet field.
func (o *IpmioverlanPolicyInventory) SetIsEncryptionKeySet(v bool) {
	o.IsEncryptionKeySet = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetPrivilege() string {
	if o == nil || o.Privilege == nil {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetPrivilegeOk() (*string, bool) {
	if o == nil || o.Privilege == nil {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasPrivilege() bool {
	if o != nil && o.Privilege != nil {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *IpmioverlanPolicyInventory) SetPrivilege(v string) {
	o.Privilege = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *IpmioverlanPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IpmioverlanPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IpmioverlanPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o IpmioverlanPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
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

func (o *IpmioverlanPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type IpmioverlanPolicyInventoryWithoutEmbeddedStruct struct {
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
		Privilege *string               `json:"Privilege,omitempty"`
		TargetMo  *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varIpmioverlanPolicyInventoryWithoutEmbeddedStruct := IpmioverlanPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIpmioverlanPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varIpmioverlanPolicyInventory := _IpmioverlanPolicyInventory{}
		varIpmioverlanPolicyInventory.ClassId = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.ClassId
		varIpmioverlanPolicyInventory.ObjectType = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varIpmioverlanPolicyInventory.Enabled = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.Enabled
		varIpmioverlanPolicyInventory.EncryptionKey = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.EncryptionKey
		varIpmioverlanPolicyInventory.IsEncryptionKeySet = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.IsEncryptionKeySet
		varIpmioverlanPolicyInventory.Privilege = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.Privilege
		varIpmioverlanPolicyInventory.TargetMo = varIpmioverlanPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = IpmioverlanPolicyInventory(varIpmioverlanPolicyInventory)
	} else {
		return err
	}

	varIpmioverlanPolicyInventory := _IpmioverlanPolicyInventory{}

	err = json.Unmarshal(bytes, &varIpmioverlanPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varIpmioverlanPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
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

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableIpmioverlanPolicyInventory struct {
	value *IpmioverlanPolicyInventory
	isSet bool
}

func (v NullableIpmioverlanPolicyInventory) Get() *IpmioverlanPolicyInventory {
	return v.value
}

func (v *NullableIpmioverlanPolicyInventory) Set(val *IpmioverlanPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmioverlanPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmioverlanPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmioverlanPolicyInventory(val *IpmioverlanPolicyInventory) *NullableIpmioverlanPolicyInventory {
	return &NullableIpmioverlanPolicyInventory{value: val, isSet: true}
}

func (v NullableIpmioverlanPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmioverlanPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
