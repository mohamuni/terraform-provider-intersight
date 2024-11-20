/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the LicenseSmartlicenseToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseSmartlicenseToken{}

// LicenseSmartlicenseToken SmartlicenseToken collection stores license registration tokens.
type LicenseSmartlicenseToken struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Smart license registration token.
	Token                *string                                       `json:"Token,omitempty"`
	AccountLicenseData   NullableLicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseSmartlicenseToken LicenseSmartlicenseToken

// NewLicenseSmartlicenseToken instantiates a new LicenseSmartlicenseToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSmartlicenseToken(classId string, objectType string) *LicenseSmartlicenseToken {
	this := LicenseSmartlicenseToken{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseSmartlicenseTokenWithDefaults instantiates a new LicenseSmartlicenseToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseSmartlicenseTokenWithDefaults() *LicenseSmartlicenseToken {
	this := LicenseSmartlicenseToken{}
	var classId string = "license.SmartlicenseToken"
	this.ClassId = classId
	var objectType string = "license.SmartlicenseToken"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseSmartlicenseToken) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseSmartlicenseToken) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseSmartlicenseToken) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "license.SmartlicenseToken" of the ClassId field.
func (o *LicenseSmartlicenseToken) GetDefaultClassId() interface{} {
	return "license.SmartlicenseToken"
}

// GetObjectType returns the ObjectType field value
func (o *LicenseSmartlicenseToken) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseSmartlicenseToken) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseSmartlicenseToken) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "license.SmartlicenseToken" of the ObjectType field.
func (o *LicenseSmartlicenseToken) GetDefaultObjectType() interface{} {
	return "license.SmartlicenseToken"
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LicenseSmartlicenseToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSmartlicenseToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LicenseSmartlicenseToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LicenseSmartlicenseToken) SetToken(v string) {
	o.Token = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseSmartlicenseToken) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || IsNil(o.AccountLicenseData.Get()) {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData.Get()
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseSmartlicenseToken) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountLicenseData.Get(), o.AccountLicenseData.IsSet()
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseSmartlicenseToken) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData.IsSet() {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given NullableLicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseSmartlicenseToken) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData.Set(&v)
}

// SetAccountLicenseDataNil sets the value for AccountLicenseData to be an explicit nil
func (o *LicenseSmartlicenseToken) SetAccountLicenseDataNil() {
	o.AccountLicenseData.Set(nil)
}

// UnsetAccountLicenseData ensures that no value is present for AccountLicenseData, not even an explicit nil
func (o *LicenseSmartlicenseToken) UnsetAccountLicenseData() {
	o.AccountLicenseData.Unset()
}

func (o LicenseSmartlicenseToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseSmartlicenseToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Token) {
		toSerialize["Token"] = o.Token
	}
	if o.AccountLicenseData.IsSet() {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LicenseSmartlicenseToken) UnmarshalJSON(data []byte) (err error) {
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
	type LicenseSmartlicenseTokenWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Smart license registration token.
		Token              *string                                       `json:"Token,omitempty"`
		AccountLicenseData NullableLicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	}

	varLicenseSmartlicenseTokenWithoutEmbeddedStruct := LicenseSmartlicenseTokenWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varLicenseSmartlicenseTokenWithoutEmbeddedStruct)
	if err == nil {
		varLicenseSmartlicenseToken := _LicenseSmartlicenseToken{}
		varLicenseSmartlicenseToken.ClassId = varLicenseSmartlicenseTokenWithoutEmbeddedStruct.ClassId
		varLicenseSmartlicenseToken.ObjectType = varLicenseSmartlicenseTokenWithoutEmbeddedStruct.ObjectType
		varLicenseSmartlicenseToken.Token = varLicenseSmartlicenseTokenWithoutEmbeddedStruct.Token
		varLicenseSmartlicenseToken.AccountLicenseData = varLicenseSmartlicenseTokenWithoutEmbeddedStruct.AccountLicenseData
		*o = LicenseSmartlicenseToken(varLicenseSmartlicenseToken)
	} else {
		return err
	}

	varLicenseSmartlicenseToken := _LicenseSmartlicenseToken{}

	err = json.Unmarshal(data, &varLicenseSmartlicenseToken)
	if err == nil {
		o.MoBaseMo = varLicenseSmartlicenseToken.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Token")
		delete(additionalProperties, "AccountLicenseData")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableLicenseSmartlicenseToken struct {
	value *LicenseSmartlicenseToken
	isSet bool
}

func (v NullableLicenseSmartlicenseToken) Get() *LicenseSmartlicenseToken {
	return v.value
}

func (v *NullableLicenseSmartlicenseToken) Set(val *LicenseSmartlicenseToken) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSmartlicenseToken) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSmartlicenseToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSmartlicenseToken(val *LicenseSmartlicenseToken) *NullableLicenseSmartlicenseToken {
	return &NullableLicenseSmartlicenseToken{value: val, isSet: true}
}

func (v NullableLicenseSmartlicenseToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSmartlicenseToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
