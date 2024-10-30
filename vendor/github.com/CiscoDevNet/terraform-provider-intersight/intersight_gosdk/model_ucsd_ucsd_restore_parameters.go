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

// checks if the UcsdUcsdRestoreParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UcsdUcsdRestoreParameters{}

// UcsdUcsdRestoreParameters Restore Configuration Parameters for UCS Director restore workflow.
type UcsdUcsdRestoreParameters struct {
	RecoveryConfigParams
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The complete location of the path on the server. The location should be specified in the following format- hostname-or-ipv4address<:port>/absolute-file-path.
	Location *string `json:"Location,omitempty"`
	// The password of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
	Password *string `json:"Password,omitempty"`
	// The protocol used to backup the UCS Director.
	Protocol *string `json:"Protocol,omitempty"`
	// Decides whether UCS Director property files should also be restored.
	RestoreConfigurationFiles *bool `json:"RestoreConfigurationFiles,omitempty"`
	// Decides whether license should also be restored.
	RestoreLicense *bool `json:"RestoreLicense,omitempty"`
	// The username of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcsdUcsdRestoreParameters UcsdUcsdRestoreParameters

// NewUcsdUcsdRestoreParameters instantiates a new UcsdUcsdRestoreParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdUcsdRestoreParameters(classId string, objectType string) *UcsdUcsdRestoreParameters {
	this := UcsdUcsdRestoreParameters{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUcsdUcsdRestoreParametersWithDefaults instantiates a new UcsdUcsdRestoreParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdUcsdRestoreParametersWithDefaults() *UcsdUcsdRestoreParameters {
	this := UcsdUcsdRestoreParameters{}
	var classId string = "ucsd.UcsdRestoreParameters"
	this.ClassId = classId
	var objectType string = "ucsd.UcsdRestoreParameters"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UcsdUcsdRestoreParameters) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UcsdUcsdRestoreParameters) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ucsd.UcsdRestoreParameters" of the ClassId field.
func (o *UcsdUcsdRestoreParameters) GetDefaultClassId() interface{} {
	return "ucsd.UcsdRestoreParameters"
}

// GetObjectType returns the ObjectType field value
func (o *UcsdUcsdRestoreParameters) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UcsdUcsdRestoreParameters) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ucsd.UcsdRestoreParameters" of the ObjectType field.
func (o *UcsdUcsdRestoreParameters) GetDefaultObjectType() interface{} {
	return "ucsd.UcsdRestoreParameters"
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetIsPasswordSet() bool {
	if o == nil || IsNil(o.IsPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasswordSet) {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasIsPasswordSet() bool {
	if o != nil && !IsNil(o.IsPasswordSet) {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *UcsdUcsdRestoreParameters) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UcsdUcsdRestoreParameters) SetLocation(v string) {
	o.Location = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UcsdUcsdRestoreParameters) SetPassword(v string) {
	o.Password = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UcsdUcsdRestoreParameters) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRestoreConfigurationFiles returns the RestoreConfigurationFiles field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetRestoreConfigurationFiles() bool {
	if o == nil || IsNil(o.RestoreConfigurationFiles) {
		var ret bool
		return ret
	}
	return *o.RestoreConfigurationFiles
}

// GetRestoreConfigurationFilesOk returns a tuple with the RestoreConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetRestoreConfigurationFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.RestoreConfigurationFiles) {
		return nil, false
	}
	return o.RestoreConfigurationFiles, true
}

// HasRestoreConfigurationFiles returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasRestoreConfigurationFiles() bool {
	if o != nil && !IsNil(o.RestoreConfigurationFiles) {
		return true
	}

	return false
}

// SetRestoreConfigurationFiles gets a reference to the given bool and assigns it to the RestoreConfigurationFiles field.
func (o *UcsdUcsdRestoreParameters) SetRestoreConfigurationFiles(v bool) {
	o.RestoreConfigurationFiles = &v
}

// GetRestoreLicense returns the RestoreLicense field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetRestoreLicense() bool {
	if o == nil || IsNil(o.RestoreLicense) {
		var ret bool
		return ret
	}
	return *o.RestoreLicense
}

// GetRestoreLicenseOk returns a tuple with the RestoreLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetRestoreLicenseOk() (*bool, bool) {
	if o == nil || IsNil(o.RestoreLicense) {
		return nil, false
	}
	return o.RestoreLicense, true
}

// HasRestoreLicense returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasRestoreLicense() bool {
	if o != nil && !IsNil(o.RestoreLicense) {
		return true
	}

	return false
}

// SetRestoreLicense gets a reference to the given bool and assigns it to the RestoreLicense field.
func (o *UcsdUcsdRestoreParameters) SetRestoreLicense(v bool) {
	o.RestoreLicense = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UcsdUcsdRestoreParameters) SetUsername(v string) {
	o.Username = &v
}

func (o UcsdUcsdRestoreParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UcsdUcsdRestoreParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedRecoveryConfigParams, errRecoveryConfigParams := json.Marshal(o.RecoveryConfigParams)
	if errRecoveryConfigParams != nil {
		return map[string]interface{}{}, errRecoveryConfigParams
	}
	errRecoveryConfigParams = json.Unmarshal([]byte(serializedRecoveryConfigParams), &toSerialize)
	if errRecoveryConfigParams != nil {
		return map[string]interface{}{}, errRecoveryConfigParams
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
	if !IsNil(o.Location) {
		toSerialize["Location"] = o.Location
	}
	if !IsNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.RestoreConfigurationFiles) {
		toSerialize["RestoreConfigurationFiles"] = o.RestoreConfigurationFiles
	}
	if !IsNil(o.RestoreLicense) {
		toSerialize["RestoreLicense"] = o.RestoreLicense
	}
	if !IsNil(o.Username) {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UcsdUcsdRestoreParameters) UnmarshalJSON(data []byte) (err error) {
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
	type UcsdUcsdRestoreParametersWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The complete location of the path on the server. The location should be specified in the following format- hostname-or-ipv4address<:port>/absolute-file-path.
		Location *string `json:"Location,omitempty"`
		// The password of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
		Password *string `json:"Password,omitempty"`
		// The protocol used to backup the UCS Director.
		Protocol *string `json:"Protocol,omitempty"`
		// Decides whether UCS Director property files should also be restored.
		RestoreConfigurationFiles *bool `json:"RestoreConfigurationFiles,omitempty"`
		// Decides whether license should also be restored.
		RestoreLicense *bool `json:"RestoreLicense,omitempty"`
		// The username of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
		Username *string `json:"Username,omitempty"`
	}

	varUcsdUcsdRestoreParametersWithoutEmbeddedStruct := UcsdUcsdRestoreParametersWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUcsdUcsdRestoreParametersWithoutEmbeddedStruct)
	if err == nil {
		varUcsdUcsdRestoreParameters := _UcsdUcsdRestoreParameters{}
		varUcsdUcsdRestoreParameters.ClassId = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.ClassId
		varUcsdUcsdRestoreParameters.ObjectType = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.ObjectType
		varUcsdUcsdRestoreParameters.IsPasswordSet = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.IsPasswordSet
		varUcsdUcsdRestoreParameters.Location = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.Location
		varUcsdUcsdRestoreParameters.Password = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.Password
		varUcsdUcsdRestoreParameters.Protocol = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.Protocol
		varUcsdUcsdRestoreParameters.RestoreConfigurationFiles = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.RestoreConfigurationFiles
		varUcsdUcsdRestoreParameters.RestoreLicense = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.RestoreLicense
		varUcsdUcsdRestoreParameters.Username = varUcsdUcsdRestoreParametersWithoutEmbeddedStruct.Username
		*o = UcsdUcsdRestoreParameters(varUcsdUcsdRestoreParameters)
	} else {
		return err
	}

	varUcsdUcsdRestoreParameters := _UcsdUcsdRestoreParameters{}

	err = json.Unmarshal(data, &varUcsdUcsdRestoreParameters)
	if err == nil {
		o.RecoveryConfigParams = varUcsdUcsdRestoreParameters.RecoveryConfigParams
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RestoreConfigurationFiles")
		delete(additionalProperties, "RestoreLicense")
		delete(additionalProperties, "Username")

		// remove fields from embedded structs
		reflectRecoveryConfigParams := reflect.ValueOf(o.RecoveryConfigParams)
		for i := 0; i < reflectRecoveryConfigParams.Type().NumField(); i++ {
			t := reflectRecoveryConfigParams.Type().Field(i)

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

type NullableUcsdUcsdRestoreParameters struct {
	value *UcsdUcsdRestoreParameters
	isSet bool
}

func (v NullableUcsdUcsdRestoreParameters) Get() *UcsdUcsdRestoreParameters {
	return v.value
}

func (v *NullableUcsdUcsdRestoreParameters) Set(val *UcsdUcsdRestoreParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdUcsdRestoreParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdUcsdRestoreParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdUcsdRestoreParameters(val *UcsdUcsdRestoreParameters) *NullableUcsdUcsdRestoreParameters {
	return &NullableUcsdUcsdRestoreParameters{value: val, isSet: true}
}

func (v NullableUcsdUcsdRestoreParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdUcsdRestoreParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
