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

// checks if the StorageKmipAuthCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageKmipAuthCredentials{}

// StorageKmipAuthCredentials Models the KMIP Server Authentication configuration to be able to connect and communicate from Endpoint.
type StorageKmipAuthCredentials struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password for the KMIP server login.
	Password *string "json:\"Password,omitempty\" validate:\"regexp=[!\\\"#%&'\\\\(\\\\)\\\\*\\\\+,\\\\-\\\\.\\/:;<>@\\\\[\\\\\\\\\\\\]\\\\^_`\\\\{\\\\|\\\\}~a-zA-Z0-9]{0,80}\""
	// Enables/disables the authentication for communicating with KMIP server. This flag enables the authentication which makes authentication mandatory.
	UseAuthentication *bool `json:"UseAuthentication,omitempty"`
	// The user name for the KMIP server login.
	Username             *string `json:"Username,omitempty" validate:"regexp=^$|[a-zA-Z][a-zA-Z0-9_.-]{0,31}"`
	AdditionalProperties map[string]interface{}
}

type _StorageKmipAuthCredentials StorageKmipAuthCredentials

// NewStorageKmipAuthCredentials instantiates a new StorageKmipAuthCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageKmipAuthCredentials(classId string, objectType string) *StorageKmipAuthCredentials {
	this := StorageKmipAuthCredentials{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageKmipAuthCredentialsWithDefaults instantiates a new StorageKmipAuthCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageKmipAuthCredentialsWithDefaults() *StorageKmipAuthCredentials {
	this := StorageKmipAuthCredentials{}
	var classId string = "storage.KmipAuthCredentials"
	this.ClassId = classId
	var objectType string = "storage.KmipAuthCredentials"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageKmipAuthCredentials) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageKmipAuthCredentials) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageKmipAuthCredentials) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.KmipAuthCredentials" of the ClassId field.
func (o *StorageKmipAuthCredentials) GetDefaultClassId() interface{} {
	return "storage.KmipAuthCredentials"
}

// GetObjectType returns the ObjectType field value
func (o *StorageKmipAuthCredentials) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageKmipAuthCredentials) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageKmipAuthCredentials) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.KmipAuthCredentials" of the ObjectType field.
func (o *StorageKmipAuthCredentials) GetDefaultObjectType() interface{} {
	return "storage.KmipAuthCredentials"
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *StorageKmipAuthCredentials) GetIsPasswordSet() bool {
	if o == nil || IsNil(o.IsPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipAuthCredentials) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasswordSet) {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *StorageKmipAuthCredentials) HasIsPasswordSet() bool {
	if o != nil && !IsNil(o.IsPasswordSet) {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *StorageKmipAuthCredentials) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *StorageKmipAuthCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipAuthCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *StorageKmipAuthCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *StorageKmipAuthCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetUseAuthentication returns the UseAuthentication field value if set, zero value otherwise.
func (o *StorageKmipAuthCredentials) GetUseAuthentication() bool {
	if o == nil || IsNil(o.UseAuthentication) {
		var ret bool
		return ret
	}
	return *o.UseAuthentication
}

// GetUseAuthenticationOk returns a tuple with the UseAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipAuthCredentials) GetUseAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.UseAuthentication) {
		return nil, false
	}
	return o.UseAuthentication, true
}

// HasUseAuthentication returns a boolean if a field has been set.
func (o *StorageKmipAuthCredentials) HasUseAuthentication() bool {
	if o != nil && !IsNil(o.UseAuthentication) {
		return true
	}

	return false
}

// SetUseAuthentication gets a reference to the given bool and assigns it to the UseAuthentication field.
func (o *StorageKmipAuthCredentials) SetUseAuthentication(v bool) {
	o.UseAuthentication = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *StorageKmipAuthCredentials) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipAuthCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *StorageKmipAuthCredentials) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *StorageKmipAuthCredentials) SetUsername(v string) {
	o.Username = &v
}

func (o StorageKmipAuthCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageKmipAuthCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IsPasswordSet) {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if !IsNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !IsNil(o.UseAuthentication) {
		toSerialize["UseAuthentication"] = o.UseAuthentication
	}
	if !IsNil(o.Username) {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageKmipAuthCredentials) UnmarshalJSON(data []byte) (err error) {
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
	type StorageKmipAuthCredentialsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The password for the KMIP server login.
		Password *string "json:\"Password,omitempty\" validate:\"regexp=[!\\\"#%&'\\\\(\\\\)\\\\*\\\\+,\\\\-\\\\.\\/:;<>@\\\\[\\\\\\\\\\\\]\\\\^_`\\\\{\\\\|\\\\}~a-zA-Z0-9]{0,80}\""
		// Enables/disables the authentication for communicating with KMIP server. This flag enables the authentication which makes authentication mandatory.
		UseAuthentication *bool `json:"UseAuthentication,omitempty"`
		// The user name for the KMIP server login.
		Username *string `json:"Username,omitempty" validate:"regexp=^$|[a-zA-Z][a-zA-Z0-9_.-]{0,31}"`
	}

	varStorageKmipAuthCredentialsWithoutEmbeddedStruct := StorageKmipAuthCredentialsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageKmipAuthCredentialsWithoutEmbeddedStruct)
	if err == nil {
		varStorageKmipAuthCredentials := _StorageKmipAuthCredentials{}
		varStorageKmipAuthCredentials.ClassId = varStorageKmipAuthCredentialsWithoutEmbeddedStruct.ClassId
		varStorageKmipAuthCredentials.ObjectType = varStorageKmipAuthCredentialsWithoutEmbeddedStruct.ObjectType
		varStorageKmipAuthCredentials.IsPasswordSet = varStorageKmipAuthCredentialsWithoutEmbeddedStruct.IsPasswordSet
		varStorageKmipAuthCredentials.Password = varStorageKmipAuthCredentialsWithoutEmbeddedStruct.Password
		varStorageKmipAuthCredentials.UseAuthentication = varStorageKmipAuthCredentialsWithoutEmbeddedStruct.UseAuthentication
		varStorageKmipAuthCredentials.Username = varStorageKmipAuthCredentialsWithoutEmbeddedStruct.Username
		*o = StorageKmipAuthCredentials(varStorageKmipAuthCredentials)
	} else {
		return err
	}

	varStorageKmipAuthCredentials := _StorageKmipAuthCredentials{}

	err = json.Unmarshal(data, &varStorageKmipAuthCredentials)
	if err == nil {
		o.MoBaseComplexType = varStorageKmipAuthCredentials.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "UseAuthentication")
		delete(additionalProperties, "Username")

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

type NullableStorageKmipAuthCredentials struct {
	value *StorageKmipAuthCredentials
	isSet bool
}

func (v NullableStorageKmipAuthCredentials) Get() *StorageKmipAuthCredentials {
	return v.value
}

func (v *NullableStorageKmipAuthCredentials) Set(val *StorageKmipAuthCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageKmipAuthCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageKmipAuthCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageKmipAuthCredentials(val *StorageKmipAuthCredentials) *NullableStorageKmipAuthCredentials {
	return &NullableStorageKmipAuthCredentials{value: val, isSet: true}
}

func (v NullableStorageKmipAuthCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageKmipAuthCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
