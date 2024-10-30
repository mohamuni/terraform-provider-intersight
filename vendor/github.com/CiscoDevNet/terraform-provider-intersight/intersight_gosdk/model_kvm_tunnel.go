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

// checks if the KvmTunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvmTunnel{}

// KvmTunnel Tunneled Virtual KVM access to the vKVM console of a server. This must be specified while creating the vKVM session to gain tunneled access.
type KvmTunnel struct {
	TunnelingTunnel
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	Device               NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	KvmSession           NullableKvmSessionRelationship              `json:"KvmSession,omitempty"`
	Server               NullableComputePhysicalRelationship         `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmTunnel KvmTunnel

// NewKvmTunnel instantiates a new KvmTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmTunnel(classId string, objectType string) *KvmTunnel {
	this := KvmTunnel{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewKvmTunnelWithDefaults instantiates a new KvmTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmTunnelWithDefaults() *KvmTunnel {
	this := KvmTunnel{}
	var classId string = "kvm.Tunnel"
	this.ClassId = classId
	var objectType string = "kvm.Tunnel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmTunnel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmTunnel) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kvm.Tunnel" of the ClassId field.
func (o *KvmTunnel) GetDefaultClassId() interface{} {
	return "kvm.Tunnel"
}

// GetObjectType returns the ObjectType field value
func (o *KvmTunnel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmTunnel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kvm.Tunnel" of the ObjectType field.
func (o *KvmTunnel) GetDefaultObjectType() interface{} {
	return "kvm.Tunnel"
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmTunnel) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Device.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmTunnel) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmTunnel) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmTunnel) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *KvmTunnel) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *KvmTunnel) UnsetDevice() {
	o.Device.Unset()
}

// GetKvmSession returns the KvmSession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmTunnel) GetKvmSession() KvmSessionRelationship {
	if o == nil || IsNil(o.KvmSession.Get()) {
		var ret KvmSessionRelationship
		return ret
	}
	return *o.KvmSession.Get()
}

// GetKvmSessionOk returns a tuple with the KvmSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmTunnel) GetKvmSessionOk() (*KvmSessionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.KvmSession.Get(), o.KvmSession.IsSet()
}

// HasKvmSession returns a boolean if a field has been set.
func (o *KvmTunnel) HasKvmSession() bool {
	if o != nil && o.KvmSession.IsSet() {
		return true
	}

	return false
}

// SetKvmSession gets a reference to the given NullableKvmSessionRelationship and assigns it to the KvmSession field.
func (o *KvmTunnel) SetKvmSession(v KvmSessionRelationship) {
	o.KvmSession.Set(&v)
}

// SetKvmSessionNil sets the value for KvmSession to be an explicit nil
func (o *KvmTunnel) SetKvmSessionNil() {
	o.KvmSession.Set(nil)
}

// UnsetKvmSession ensures that no value is present for KvmSession, not even an explicit nil
func (o *KvmTunnel) UnsetKvmSession() {
	o.KvmSession.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmTunnel) GetServer() ComputePhysicalRelationship {
	if o == nil || IsNil(o.Server.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmTunnel) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *KvmTunnel) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableComputePhysicalRelationship and assigns it to the Server field.
func (o *KvmTunnel) SetServer(v ComputePhysicalRelationship) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *KvmTunnel) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *KvmTunnel) UnsetServer() {
	o.Server.Unset()
}

func (o KvmTunnel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvmTunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTunnelingTunnel, errTunnelingTunnel := json.Marshal(o.TunnelingTunnel)
	if errTunnelingTunnel != nil {
		return map[string]interface{}{}, errTunnelingTunnel
	}
	errTunnelingTunnel = json.Unmarshal([]byte(serializedTunnelingTunnel), &toSerialize)
	if errTunnelingTunnel != nil {
		return map[string]interface{}{}, errTunnelingTunnel
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Device.IsSet() {
		toSerialize["Device"] = o.Device.Get()
	}
	if o.KvmSession.IsSet() {
		toSerialize["KvmSession"] = o.KvmSession.Get()
	}
	if o.Server.IsSet() {
		toSerialize["Server"] = o.Server.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KvmTunnel) UnmarshalJSON(data []byte) (err error) {
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
	type KvmTunnelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                      `json:"ObjectType"`
		Device     NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		KvmSession NullableKvmSessionRelationship              `json:"KvmSession,omitempty"`
		Server     NullableComputePhysicalRelationship         `json:"Server,omitempty"`
	}

	varKvmTunnelWithoutEmbeddedStruct := KvmTunnelWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKvmTunnelWithoutEmbeddedStruct)
	if err == nil {
		varKvmTunnel := _KvmTunnel{}
		varKvmTunnel.ClassId = varKvmTunnelWithoutEmbeddedStruct.ClassId
		varKvmTunnel.ObjectType = varKvmTunnelWithoutEmbeddedStruct.ObjectType
		varKvmTunnel.Device = varKvmTunnelWithoutEmbeddedStruct.Device
		varKvmTunnel.KvmSession = varKvmTunnelWithoutEmbeddedStruct.KvmSession
		varKvmTunnel.Server = varKvmTunnelWithoutEmbeddedStruct.Server
		*o = KvmTunnel(varKvmTunnel)
	} else {
		return err
	}

	varKvmTunnel := _KvmTunnel{}

	err = json.Unmarshal(data, &varKvmTunnel)
	if err == nil {
		o.TunnelingTunnel = varKvmTunnel.TunnelingTunnel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "KvmSession")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectTunnelingTunnel := reflect.ValueOf(o.TunnelingTunnel)
		for i := 0; i < reflectTunnelingTunnel.Type().NumField(); i++ {
			t := reflectTunnelingTunnel.Type().Field(i)

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

type NullableKvmTunnel struct {
	value *KvmTunnel
	isSet bool
}

func (v NullableKvmTunnel) Get() *KvmTunnel {
	return v.value
}

func (v *NullableKvmTunnel) Set(val *KvmTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmTunnel(val *KvmTunnel) *NullableKvmTunnel {
	return &NullableKvmTunnel{value: val, isSet: true}
}

func (v NullableKvmTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
