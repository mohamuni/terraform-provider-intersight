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

// StorageRemoteKeySetting Models the remote key configuration required for disk encryption.
type StorageRemoteKeySetting struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                             `json:"ObjectType"`
	AuthCredentials NullableStorageKmipAuthCredentials `json:"AuthCredentials,omitempty"`
	// Indicates whether the value of the 'existingKey' property has been set.
	IsExistingKeySet *bool                     `json:"IsExistingKeySet,omitempty"`
	PrimaryServer    NullableStorageKmipServer `json:"PrimaryServer,omitempty"`
	SecondaryServer  NullableStorageKmipServer `json:"SecondaryServer,omitempty"`
	// The certificate/ public key of the KMIP server. It is required for initiating secure communication with the server.
	ServerCertificate    *string `json:"ServerCertificate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageRemoteKeySetting StorageRemoteKeySetting

// NewStorageRemoteKeySetting instantiates a new StorageRemoteKeySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRemoteKeySetting(classId string, objectType string) *StorageRemoteKeySetting {
	this := StorageRemoteKeySetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageRemoteKeySettingWithDefaults instantiates a new StorageRemoteKeySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRemoteKeySettingWithDefaults() *StorageRemoteKeySetting {
	this := StorageRemoteKeySetting{}
	var classId string = "storage.RemoteKeySetting"
	this.ClassId = classId
	var objectType string = "storage.RemoteKeySetting"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageRemoteKeySetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageRemoteKeySetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageRemoteKeySetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageRemoteKeySetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthCredentials returns the AuthCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageRemoteKeySetting) GetAuthCredentials() StorageKmipAuthCredentials {
	if o == nil || o.AuthCredentials.Get() == nil {
		var ret StorageKmipAuthCredentials
		return ret
	}
	return *o.AuthCredentials.Get()
}

// GetAuthCredentialsOk returns a tuple with the AuthCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRemoteKeySetting) GetAuthCredentialsOk() (*StorageKmipAuthCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthCredentials.Get(), o.AuthCredentials.IsSet()
}

// HasAuthCredentials returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasAuthCredentials() bool {
	if o != nil && o.AuthCredentials.IsSet() {
		return true
	}

	return false
}

// SetAuthCredentials gets a reference to the given NullableStorageKmipAuthCredentials and assigns it to the AuthCredentials field.
func (o *StorageRemoteKeySetting) SetAuthCredentials(v StorageKmipAuthCredentials) {
	o.AuthCredentials.Set(&v)
}

// SetAuthCredentialsNil sets the value for AuthCredentials to be an explicit nil
func (o *StorageRemoteKeySetting) SetAuthCredentialsNil() {
	o.AuthCredentials.Set(nil)
}

// UnsetAuthCredentials ensures that no value is present for AuthCredentials, not even an explicit nil
func (o *StorageRemoteKeySetting) UnsetAuthCredentials() {
	o.AuthCredentials.Unset()
}

// GetIsExistingKeySet returns the IsExistingKeySet field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetIsExistingKeySet() bool {
	if o == nil || o.IsExistingKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsExistingKeySet
}

// GetIsExistingKeySetOk returns a tuple with the IsExistingKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetIsExistingKeySetOk() (*bool, bool) {
	if o == nil || o.IsExistingKeySet == nil {
		return nil, false
	}
	return o.IsExistingKeySet, true
}

// HasIsExistingKeySet returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasIsExistingKeySet() bool {
	if o != nil && o.IsExistingKeySet != nil {
		return true
	}

	return false
}

// SetIsExistingKeySet gets a reference to the given bool and assigns it to the IsExistingKeySet field.
func (o *StorageRemoteKeySetting) SetIsExistingKeySet(v bool) {
	o.IsExistingKeySet = &v
}

// GetPrimaryServer returns the PrimaryServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageRemoteKeySetting) GetPrimaryServer() StorageKmipServer {
	if o == nil || o.PrimaryServer.Get() == nil {
		var ret StorageKmipServer
		return ret
	}
	return *o.PrimaryServer.Get()
}

// GetPrimaryServerOk returns a tuple with the PrimaryServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRemoteKeySetting) GetPrimaryServerOk() (*StorageKmipServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryServer.Get(), o.PrimaryServer.IsSet()
}

// HasPrimaryServer returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasPrimaryServer() bool {
	if o != nil && o.PrimaryServer.IsSet() {
		return true
	}

	return false
}

// SetPrimaryServer gets a reference to the given NullableStorageKmipServer and assigns it to the PrimaryServer field.
func (o *StorageRemoteKeySetting) SetPrimaryServer(v StorageKmipServer) {
	o.PrimaryServer.Set(&v)
}

// SetPrimaryServerNil sets the value for PrimaryServer to be an explicit nil
func (o *StorageRemoteKeySetting) SetPrimaryServerNil() {
	o.PrimaryServer.Set(nil)
}

// UnsetPrimaryServer ensures that no value is present for PrimaryServer, not even an explicit nil
func (o *StorageRemoteKeySetting) UnsetPrimaryServer() {
	o.PrimaryServer.Unset()
}

// GetSecondaryServer returns the SecondaryServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageRemoteKeySetting) GetSecondaryServer() StorageKmipServer {
	if o == nil || o.SecondaryServer.Get() == nil {
		var ret StorageKmipServer
		return ret
	}
	return *o.SecondaryServer.Get()
}

// GetSecondaryServerOk returns a tuple with the SecondaryServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRemoteKeySetting) GetSecondaryServerOk() (*StorageKmipServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryServer.Get(), o.SecondaryServer.IsSet()
}

// HasSecondaryServer returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasSecondaryServer() bool {
	if o != nil && o.SecondaryServer.IsSet() {
		return true
	}

	return false
}

// SetSecondaryServer gets a reference to the given NullableStorageKmipServer and assigns it to the SecondaryServer field.
func (o *StorageRemoteKeySetting) SetSecondaryServer(v StorageKmipServer) {
	o.SecondaryServer.Set(&v)
}

// SetSecondaryServerNil sets the value for SecondaryServer to be an explicit nil
func (o *StorageRemoteKeySetting) SetSecondaryServerNil() {
	o.SecondaryServer.Set(nil)
}

// UnsetSecondaryServer ensures that no value is present for SecondaryServer, not even an explicit nil
func (o *StorageRemoteKeySetting) UnsetSecondaryServer() {
	o.SecondaryServer.Unset()
}

// GetServerCertificate returns the ServerCertificate field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetServerCertificate() string {
	if o == nil || o.ServerCertificate == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificate
}

// GetServerCertificateOk returns a tuple with the ServerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetServerCertificateOk() (*string, bool) {
	if o == nil || o.ServerCertificate == nil {
		return nil, false
	}
	return o.ServerCertificate, true
}

// HasServerCertificate returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasServerCertificate() bool {
	if o != nil && o.ServerCertificate != nil {
		return true
	}

	return false
}

// SetServerCertificate gets a reference to the given string and assigns it to the ServerCertificate field.
func (o *StorageRemoteKeySetting) SetServerCertificate(v string) {
	o.ServerCertificate = &v
}

func (o StorageRemoteKeySetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthCredentials.IsSet() {
		toSerialize["AuthCredentials"] = o.AuthCredentials.Get()
	}
	if o.IsExistingKeySet != nil {
		toSerialize["IsExistingKeySet"] = o.IsExistingKeySet
	}
	if o.PrimaryServer.IsSet() {
		toSerialize["PrimaryServer"] = o.PrimaryServer.Get()
	}
	if o.SecondaryServer.IsSet() {
		toSerialize["SecondaryServer"] = o.SecondaryServer.Get()
	}
	if o.ServerCertificate != nil {
		toSerialize["ServerCertificate"] = o.ServerCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageRemoteKeySetting) UnmarshalJSON(bytes []byte) (err error) {
	type StorageRemoteKeySettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string                             `json:"ObjectType"`
		AuthCredentials NullableStorageKmipAuthCredentials `json:"AuthCredentials,omitempty"`
		// Indicates whether the value of the 'existingKey' property has been set.
		IsExistingKeySet *bool                     `json:"IsExistingKeySet,omitempty"`
		PrimaryServer    NullableStorageKmipServer `json:"PrimaryServer,omitempty"`
		SecondaryServer  NullableStorageKmipServer `json:"SecondaryServer,omitempty"`
		// The certificate/ public key of the KMIP server. It is required for initiating secure communication with the server.
		ServerCertificate *string `json:"ServerCertificate,omitempty"`
	}

	varStorageRemoteKeySettingWithoutEmbeddedStruct := StorageRemoteKeySettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageRemoteKeySettingWithoutEmbeddedStruct)
	if err == nil {
		varStorageRemoteKeySetting := _StorageRemoteKeySetting{}
		varStorageRemoteKeySetting.ClassId = varStorageRemoteKeySettingWithoutEmbeddedStruct.ClassId
		varStorageRemoteKeySetting.ObjectType = varStorageRemoteKeySettingWithoutEmbeddedStruct.ObjectType
		varStorageRemoteKeySetting.AuthCredentials = varStorageRemoteKeySettingWithoutEmbeddedStruct.AuthCredentials
		varStorageRemoteKeySetting.IsExistingKeySet = varStorageRemoteKeySettingWithoutEmbeddedStruct.IsExistingKeySet
		varStorageRemoteKeySetting.PrimaryServer = varStorageRemoteKeySettingWithoutEmbeddedStruct.PrimaryServer
		varStorageRemoteKeySetting.SecondaryServer = varStorageRemoteKeySettingWithoutEmbeddedStruct.SecondaryServer
		varStorageRemoteKeySetting.ServerCertificate = varStorageRemoteKeySettingWithoutEmbeddedStruct.ServerCertificate
		*o = StorageRemoteKeySetting(varStorageRemoteKeySetting)
	} else {
		return err
	}

	varStorageRemoteKeySetting := _StorageRemoteKeySetting{}

	err = json.Unmarshal(bytes, &varStorageRemoteKeySetting)
	if err == nil {
		o.MoBaseComplexType = varStorageRemoteKeySetting.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthCredentials")
		delete(additionalProperties, "IsExistingKeySet")
		delete(additionalProperties, "PrimaryServer")
		delete(additionalProperties, "SecondaryServer")
		delete(additionalProperties, "ServerCertificate")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableStorageRemoteKeySetting struct {
	value *StorageRemoteKeySetting
	isSet bool
}

func (v NullableStorageRemoteKeySetting) Get() *StorageRemoteKeySetting {
	return v.value
}

func (v *NullableStorageRemoteKeySetting) Set(val *StorageRemoteKeySetting) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRemoteKeySetting) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRemoteKeySetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRemoteKeySetting(val *StorageRemoteKeySetting) *NullableStorageRemoteKeySetting {
	return &NullableStorageRemoteKeySetting{value: val, isSet: true}
}

func (v NullableStorageRemoteKeySetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRemoteKeySetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
