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

// checks if the KvmSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvmSession{}

// KvmSession Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
type KvmSession struct {
	SessionAbstractSubSession
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// One time URL that is used to launch the vKVM console.
	KvmLaunchUrlPath *string `json:"KvmLaunchUrlPath,omitempty"`
	// Unique ID of the KVM Session URI.
	KvmSessionId *string `json:"KvmSessionId,omitempty"`
	// Temporary one-time password for vKVM access.
	OneTimePassword *string `json:"OneTimePassword,omitempty"`
	// Indicates if vKVM SSO is supported on the server.
	SsoSupported *bool `json:"SsoSupported,omitempty"`
	// Username used for vKVM access.
	Username             *string                                     `json:"Username,omitempty"`
	Device               NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Server               NullableComputePhysicalRelationship         `json:"Server,omitempty"`
	Tunnel               NullableKvmTunnelRelationship               `json:"Tunnel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmSession KvmSession

// NewKvmSession instantiates a new KvmSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmSession(classId string, objectType string) *KvmSession {
	this := KvmSession{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewKvmSessionWithDefaults instantiates a new KvmSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmSessionWithDefaults() *KvmSession {
	this := KvmSession{}
	var classId string = "kvm.Session"
	this.ClassId = classId
	var objectType string = "kvm.Session"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmSession) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmSession) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmSession) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kvm.Session" of the ClassId field.
func (o *KvmSession) GetDefaultClassId() interface{} {
	return "kvm.Session"
}

// GetObjectType returns the ObjectType field value
func (o *KvmSession) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmSession) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmSession) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kvm.Session" of the ObjectType field.
func (o *KvmSession) GetDefaultObjectType() interface{} {
	return "kvm.Session"
}

// GetKvmLaunchUrlPath returns the KvmLaunchUrlPath field value if set, zero value otherwise.
func (o *KvmSession) GetKvmLaunchUrlPath() string {
	if o == nil || IsNil(o.KvmLaunchUrlPath) {
		var ret string
		return ret
	}
	return *o.KvmLaunchUrlPath
}

// GetKvmLaunchUrlPathOk returns a tuple with the KvmLaunchUrlPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetKvmLaunchUrlPathOk() (*string, bool) {
	if o == nil || IsNil(o.KvmLaunchUrlPath) {
		return nil, false
	}
	return o.KvmLaunchUrlPath, true
}

// HasKvmLaunchUrlPath returns a boolean if a field has been set.
func (o *KvmSession) HasKvmLaunchUrlPath() bool {
	if o != nil && !IsNil(o.KvmLaunchUrlPath) {
		return true
	}

	return false
}

// SetKvmLaunchUrlPath gets a reference to the given string and assigns it to the KvmLaunchUrlPath field.
func (o *KvmSession) SetKvmLaunchUrlPath(v string) {
	o.KvmLaunchUrlPath = &v
}

// GetKvmSessionId returns the KvmSessionId field value if set, zero value otherwise.
func (o *KvmSession) GetKvmSessionId() string {
	if o == nil || IsNil(o.KvmSessionId) {
		var ret string
		return ret
	}
	return *o.KvmSessionId
}

// GetKvmSessionIdOk returns a tuple with the KvmSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetKvmSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.KvmSessionId) {
		return nil, false
	}
	return o.KvmSessionId, true
}

// HasKvmSessionId returns a boolean if a field has been set.
func (o *KvmSession) HasKvmSessionId() bool {
	if o != nil && !IsNil(o.KvmSessionId) {
		return true
	}

	return false
}

// SetKvmSessionId gets a reference to the given string and assigns it to the KvmSessionId field.
func (o *KvmSession) SetKvmSessionId(v string) {
	o.KvmSessionId = &v
}

// GetOneTimePassword returns the OneTimePassword field value if set, zero value otherwise.
func (o *KvmSession) GetOneTimePassword() string {
	if o == nil || IsNil(o.OneTimePassword) {
		var ret string
		return ret
	}
	return *o.OneTimePassword
}

// GetOneTimePasswordOk returns a tuple with the OneTimePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetOneTimePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OneTimePassword) {
		return nil, false
	}
	return o.OneTimePassword, true
}

// HasOneTimePassword returns a boolean if a field has been set.
func (o *KvmSession) HasOneTimePassword() bool {
	if o != nil && !IsNil(o.OneTimePassword) {
		return true
	}

	return false
}

// SetOneTimePassword gets a reference to the given string and assigns it to the OneTimePassword field.
func (o *KvmSession) SetOneTimePassword(v string) {
	o.OneTimePassword = &v
}

// GetSsoSupported returns the SsoSupported field value if set, zero value otherwise.
func (o *KvmSession) GetSsoSupported() bool {
	if o == nil || IsNil(o.SsoSupported) {
		var ret bool
		return ret
	}
	return *o.SsoSupported
}

// GetSsoSupportedOk returns a tuple with the SsoSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetSsoSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.SsoSupported) {
		return nil, false
	}
	return o.SsoSupported, true
}

// HasSsoSupported returns a boolean if a field has been set.
func (o *KvmSession) HasSsoSupported() bool {
	if o != nil && !IsNil(o.SsoSupported) {
		return true
	}

	return false
}

// SetSsoSupported gets a reference to the given bool and assigns it to the SsoSupported field.
func (o *KvmSession) SetSsoSupported(v bool) {
	o.SsoSupported = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *KvmSession) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *KvmSession) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *KvmSession) SetUsername(v string) {
	o.Username = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmSession) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Device.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmSession) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmSession) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmSession) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *KvmSession) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *KvmSession) UnsetDevice() {
	o.Device.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmSession) GetServer() ComputePhysicalRelationship {
	if o == nil || IsNil(o.Server.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmSession) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *KvmSession) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableComputePhysicalRelationship and assigns it to the Server field.
func (o *KvmSession) SetServer(v ComputePhysicalRelationship) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *KvmSession) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *KvmSession) UnsetServer() {
	o.Server.Unset()
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmSession) GetTunnel() KvmTunnelRelationship {
	if o == nil || IsNil(o.Tunnel.Get()) {
		var ret KvmTunnelRelationship
		return ret
	}
	return *o.Tunnel.Get()
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmSession) GetTunnelOk() (*KvmTunnelRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tunnel.Get(), o.Tunnel.IsSet()
}

// HasTunnel returns a boolean if a field has been set.
func (o *KvmSession) HasTunnel() bool {
	if o != nil && o.Tunnel.IsSet() {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given NullableKvmTunnelRelationship and assigns it to the Tunnel field.
func (o *KvmSession) SetTunnel(v KvmTunnelRelationship) {
	o.Tunnel.Set(&v)
}

// SetTunnelNil sets the value for Tunnel to be an explicit nil
func (o *KvmSession) SetTunnelNil() {
	o.Tunnel.Set(nil)
}

// UnsetTunnel ensures that no value is present for Tunnel, not even an explicit nil
func (o *KvmSession) UnsetTunnel() {
	o.Tunnel.Unset()
}

func (o KvmSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvmSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSessionAbstractSubSession, errSessionAbstractSubSession := json.Marshal(o.SessionAbstractSubSession)
	if errSessionAbstractSubSession != nil {
		return map[string]interface{}{}, errSessionAbstractSubSession
	}
	errSessionAbstractSubSession = json.Unmarshal([]byte(serializedSessionAbstractSubSession), &toSerialize)
	if errSessionAbstractSubSession != nil {
		return map[string]interface{}{}, errSessionAbstractSubSession
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.KvmLaunchUrlPath) {
		toSerialize["KvmLaunchUrlPath"] = o.KvmLaunchUrlPath
	}
	if !IsNil(o.KvmSessionId) {
		toSerialize["KvmSessionId"] = o.KvmSessionId
	}
	if !IsNil(o.OneTimePassword) {
		toSerialize["OneTimePassword"] = o.OneTimePassword
	}
	if !IsNil(o.SsoSupported) {
		toSerialize["SsoSupported"] = o.SsoSupported
	}
	if !IsNil(o.Username) {
		toSerialize["Username"] = o.Username
	}
	if o.Device.IsSet() {
		toSerialize["Device"] = o.Device.Get()
	}
	if o.Server.IsSet() {
		toSerialize["Server"] = o.Server.Get()
	}
	if o.Tunnel.IsSet() {
		toSerialize["Tunnel"] = o.Tunnel.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KvmSession) UnmarshalJSON(data []byte) (err error) {
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
	type KvmSessionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// One time URL that is used to launch the vKVM console.
		KvmLaunchUrlPath *string `json:"KvmLaunchUrlPath,omitempty"`
		// Unique ID of the KVM Session URI.
		KvmSessionId *string `json:"KvmSessionId,omitempty"`
		// Temporary one-time password for vKVM access.
		OneTimePassword *string `json:"OneTimePassword,omitempty"`
		// Indicates if vKVM SSO is supported on the server.
		SsoSupported *bool `json:"SsoSupported,omitempty"`
		// Username used for vKVM access.
		Username *string                                     `json:"Username,omitempty"`
		Device   NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Server   NullableComputePhysicalRelationship         `json:"Server,omitempty"`
		Tunnel   NullableKvmTunnelRelationship               `json:"Tunnel,omitempty"`
	}

	varKvmSessionWithoutEmbeddedStruct := KvmSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKvmSessionWithoutEmbeddedStruct)
	if err == nil {
		varKvmSession := _KvmSession{}
		varKvmSession.ClassId = varKvmSessionWithoutEmbeddedStruct.ClassId
		varKvmSession.ObjectType = varKvmSessionWithoutEmbeddedStruct.ObjectType
		varKvmSession.KvmLaunchUrlPath = varKvmSessionWithoutEmbeddedStruct.KvmLaunchUrlPath
		varKvmSession.KvmSessionId = varKvmSessionWithoutEmbeddedStruct.KvmSessionId
		varKvmSession.OneTimePassword = varKvmSessionWithoutEmbeddedStruct.OneTimePassword
		varKvmSession.SsoSupported = varKvmSessionWithoutEmbeddedStruct.SsoSupported
		varKvmSession.Username = varKvmSessionWithoutEmbeddedStruct.Username
		varKvmSession.Device = varKvmSessionWithoutEmbeddedStruct.Device
		varKvmSession.Server = varKvmSessionWithoutEmbeddedStruct.Server
		varKvmSession.Tunnel = varKvmSessionWithoutEmbeddedStruct.Tunnel
		*o = KvmSession(varKvmSession)
	} else {
		return err
	}

	varKvmSession := _KvmSession{}

	err = json.Unmarshal(data, &varKvmSession)
	if err == nil {
		o.SessionAbstractSubSession = varKvmSession.SessionAbstractSubSession
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KvmLaunchUrlPath")
		delete(additionalProperties, "KvmSessionId")
		delete(additionalProperties, "OneTimePassword")
		delete(additionalProperties, "SsoSupported")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "Tunnel")

		// remove fields from embedded structs
		reflectSessionAbstractSubSession := reflect.ValueOf(o.SessionAbstractSubSession)
		for i := 0; i < reflectSessionAbstractSubSession.Type().NumField(); i++ {
			t := reflectSessionAbstractSubSession.Type().Field(i)

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

type NullableKvmSession struct {
	value *KvmSession
	isSet bool
}

func (v NullableKvmSession) Get() *KvmSession {
	return v.value
}

func (v *NullableKvmSession) Set(val *KvmSession) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmSession) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmSession(val *KvmSession) *NullableKvmSession {
	return &NullableKvmSession{value: val, isSet: true}
}

func (v NullableKvmSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
