/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
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

// checks if the NiatelemetrySmartLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetrySmartLicense{}

// NiatelemetrySmartLicense Object that carries details about smart license of a fabric.
type NiatelemetrySmartLicense struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicate the mode smart license is curerntly running.
	ActiveMode *string `json:"ActiveMode,omitempty"`
	// Authorization status of the smart license.
	AuthStatus *string `json:"AuthStatus,omitempty"`
	// License Udi of the smart license.
	LicenseUdi *string `json:"LicenseUdi,omitempty"`
	// Smart licensing account name in CSSM and is retrieved from CSSM after regsitration.
	SmartAccount         *string `json:"SmartAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySmartLicense NiatelemetrySmartLicense

// NewNiatelemetrySmartLicense instantiates a new NiatelemetrySmartLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySmartLicense(classId string, objectType string) *NiatelemetrySmartLicense {
	this := NiatelemetrySmartLicense{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySmartLicenseWithDefaults instantiates a new NiatelemetrySmartLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySmartLicenseWithDefaults() *NiatelemetrySmartLicense {
	this := NiatelemetrySmartLicense{}
	var classId string = "niatelemetry.SmartLicense"
	this.ClassId = classId
	var objectType string = "niatelemetry.SmartLicense"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySmartLicense) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicense) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySmartLicense) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.SmartLicense" of the ClassId field.
func (o *NiatelemetrySmartLicense) GetDefaultClassId() interface{} {
	return "niatelemetry.SmartLicense"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySmartLicense) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicense) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySmartLicense) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.SmartLicense" of the ObjectType field.
func (o *NiatelemetrySmartLicense) GetDefaultObjectType() interface{} {
	return "niatelemetry.SmartLicense"
}

// GetActiveMode returns the ActiveMode field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicense) GetActiveMode() string {
	if o == nil || IsNil(o.ActiveMode) {
		var ret string
		return ret
	}
	return *o.ActiveMode
}

// GetActiveModeOk returns a tuple with the ActiveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicense) GetActiveModeOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveMode) {
		return nil, false
	}
	return o.ActiveMode, true
}

// HasActiveMode returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicense) HasActiveMode() bool {
	if o != nil && !IsNil(o.ActiveMode) {
		return true
	}

	return false
}

// SetActiveMode gets a reference to the given string and assigns it to the ActiveMode field.
func (o *NiatelemetrySmartLicense) SetActiveMode(v string) {
	o.ActiveMode = &v
}

// GetAuthStatus returns the AuthStatus field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicense) GetAuthStatus() string {
	if o == nil || IsNil(o.AuthStatus) {
		var ret string
		return ret
	}
	return *o.AuthStatus
}

// GetAuthStatusOk returns a tuple with the AuthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicense) GetAuthStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuthStatus) {
		return nil, false
	}
	return o.AuthStatus, true
}

// HasAuthStatus returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicense) HasAuthStatus() bool {
	if o != nil && !IsNil(o.AuthStatus) {
		return true
	}

	return false
}

// SetAuthStatus gets a reference to the given string and assigns it to the AuthStatus field.
func (o *NiatelemetrySmartLicense) SetAuthStatus(v string) {
	o.AuthStatus = &v
}

// GetLicenseUdi returns the LicenseUdi field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicense) GetLicenseUdi() string {
	if o == nil || IsNil(o.LicenseUdi) {
		var ret string
		return ret
	}
	return *o.LicenseUdi
}

// GetLicenseUdiOk returns a tuple with the LicenseUdi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicense) GetLicenseUdiOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseUdi) {
		return nil, false
	}
	return o.LicenseUdi, true
}

// HasLicenseUdi returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicense) HasLicenseUdi() bool {
	if o != nil && !IsNil(o.LicenseUdi) {
		return true
	}

	return false
}

// SetLicenseUdi gets a reference to the given string and assigns it to the LicenseUdi field.
func (o *NiatelemetrySmartLicense) SetLicenseUdi(v string) {
	o.LicenseUdi = &v
}

// GetSmartAccount returns the SmartAccount field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicense) GetSmartAccount() string {
	if o == nil || IsNil(o.SmartAccount) {
		var ret string
		return ret
	}
	return *o.SmartAccount
}

// GetSmartAccountOk returns a tuple with the SmartAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicense) GetSmartAccountOk() (*string, bool) {
	if o == nil || IsNil(o.SmartAccount) {
		return nil, false
	}
	return o.SmartAccount, true
}

// HasSmartAccount returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicense) HasSmartAccount() bool {
	if o != nil && !IsNil(o.SmartAccount) {
		return true
	}

	return false
}

// SetSmartAccount gets a reference to the given string and assigns it to the SmartAccount field.
func (o *NiatelemetrySmartLicense) SetSmartAccount(v string) {
	o.SmartAccount = &v
}

func (o NiatelemetrySmartLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetrySmartLicense) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ActiveMode) {
		toSerialize["ActiveMode"] = o.ActiveMode
	}
	if !IsNil(o.AuthStatus) {
		toSerialize["AuthStatus"] = o.AuthStatus
	}
	if !IsNil(o.LicenseUdi) {
		toSerialize["LicenseUdi"] = o.LicenseUdi
	}
	if !IsNil(o.SmartAccount) {
		toSerialize["SmartAccount"] = o.SmartAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetrySmartLicense) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetrySmartLicenseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicate the mode smart license is curerntly running.
		ActiveMode *string `json:"ActiveMode,omitempty"`
		// Authorization status of the smart license.
		AuthStatus *string `json:"AuthStatus,omitempty"`
		// License Udi of the smart license.
		LicenseUdi *string `json:"LicenseUdi,omitempty"`
		// Smart licensing account name in CSSM and is retrieved from CSSM after regsitration.
		SmartAccount *string `json:"SmartAccount,omitempty"`
	}

	varNiatelemetrySmartLicenseWithoutEmbeddedStruct := NiatelemetrySmartLicenseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetrySmartLicenseWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySmartLicense := _NiatelemetrySmartLicense{}
		varNiatelemetrySmartLicense.ClassId = varNiatelemetrySmartLicenseWithoutEmbeddedStruct.ClassId
		varNiatelemetrySmartLicense.ObjectType = varNiatelemetrySmartLicenseWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySmartLicense.ActiveMode = varNiatelemetrySmartLicenseWithoutEmbeddedStruct.ActiveMode
		varNiatelemetrySmartLicense.AuthStatus = varNiatelemetrySmartLicenseWithoutEmbeddedStruct.AuthStatus
		varNiatelemetrySmartLicense.LicenseUdi = varNiatelemetrySmartLicenseWithoutEmbeddedStruct.LicenseUdi
		varNiatelemetrySmartLicense.SmartAccount = varNiatelemetrySmartLicenseWithoutEmbeddedStruct.SmartAccount
		*o = NiatelemetrySmartLicense(varNiatelemetrySmartLicense)
	} else {
		return err
	}

	varNiatelemetrySmartLicense := _NiatelemetrySmartLicense{}

	err = json.Unmarshal(data, &varNiatelemetrySmartLicense)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetrySmartLicense.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveMode")
		delete(additionalProperties, "AuthStatus")
		delete(additionalProperties, "LicenseUdi")
		delete(additionalProperties, "SmartAccount")

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

type NullableNiatelemetrySmartLicense struct {
	value *NiatelemetrySmartLicense
	isSet bool
}

func (v NullableNiatelemetrySmartLicense) Get() *NiatelemetrySmartLicense {
	return v.value
}

func (v *NullableNiatelemetrySmartLicense) Set(val *NiatelemetrySmartLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySmartLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySmartLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySmartLicense(val *NiatelemetrySmartLicense) *NullableNiatelemetrySmartLicense {
	return &NullableNiatelemetrySmartLicense{value: val, isSet: true}
}

func (v NullableNiatelemetrySmartLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySmartLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
