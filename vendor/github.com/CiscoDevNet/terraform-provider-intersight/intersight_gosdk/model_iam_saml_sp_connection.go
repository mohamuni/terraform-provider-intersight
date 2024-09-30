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

// checks if the IamSamlSpConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamSamlSpConnection{}

// IamSamlSpConnection SamlSpConnection stores the metadata elements.
type IamSamlSpConnection struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// SingleLogout Services binding in IdP Metadata.
	IdentityProviderSloBinding *string `json:"IdentityProviderSloBinding,omitempty"`
	// SingleLogOut Services location in IdP Metadata.
	IdentityProviderSloUrl *string `json:"IdentityProviderSloUrl,omitempty"`
	// SingleSignOn Services binding in IdP Metadata.
	IdentityProviderSsoBinding *string `json:"IdentityProviderSsoBinding,omitempty"`
	// SingleSignOn Services location in IdP Metadata.
	IdentityProviderSsoUrl *string `json:"IdentityProviderSsoUrl,omitempty"`
	// Decoded Certificate from IdP Metatdata.
	IdpCertificateStore interface{} `json:"IdpCertificateStore,omitempty"`
	// WantAuthnRequestsSigned from IdP Metadata.
	SignAuthnRequests    *bool `json:"SignAuthnRequests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSamlSpConnection IamSamlSpConnection

// NewIamSamlSpConnection instantiates a new IamSamlSpConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSamlSpConnection(classId string, objectType string) *IamSamlSpConnection {
	this := IamSamlSpConnection{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSamlSpConnectionWithDefaults instantiates a new IamSamlSpConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSamlSpConnectionWithDefaults() *IamSamlSpConnection {
	this := IamSamlSpConnection{}
	var classId string = "iam.SamlSpConnection"
	this.ClassId = classId
	var objectType string = "iam.SamlSpConnection"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSamlSpConnection) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSamlSpConnection) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.SamlSpConnection" of the ClassId field.
func (o *IamSamlSpConnection) GetDefaultClassId() interface{} {
	return "iam.SamlSpConnection"
}

// GetObjectType returns the ObjectType field value
func (o *IamSamlSpConnection) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSamlSpConnection) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.SamlSpConnection" of the ObjectType field.
func (o *IamSamlSpConnection) GetDefaultObjectType() interface{} {
	return "iam.SamlSpConnection"
}

// GetIdentityProviderSloBinding returns the IdentityProviderSloBinding field value if set, zero value otherwise.
func (o *IamSamlSpConnection) GetIdentityProviderSloBinding() string {
	if o == nil || IsNil(o.IdentityProviderSloBinding) {
		var ret string
		return ret
	}
	return *o.IdentityProviderSloBinding
}

// GetIdentityProviderSloBindingOk returns a tuple with the IdentityProviderSloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetIdentityProviderSloBindingOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityProviderSloBinding) {
		return nil, false
	}
	return o.IdentityProviderSloBinding, true
}

// HasIdentityProviderSloBinding returns a boolean if a field has been set.
func (o *IamSamlSpConnection) HasIdentityProviderSloBinding() bool {
	if o != nil && !IsNil(o.IdentityProviderSloBinding) {
		return true
	}

	return false
}

// SetIdentityProviderSloBinding gets a reference to the given string and assigns it to the IdentityProviderSloBinding field.
func (o *IamSamlSpConnection) SetIdentityProviderSloBinding(v string) {
	o.IdentityProviderSloBinding = &v
}

// GetIdentityProviderSloUrl returns the IdentityProviderSloUrl field value if set, zero value otherwise.
func (o *IamSamlSpConnection) GetIdentityProviderSloUrl() string {
	if o == nil || IsNil(o.IdentityProviderSloUrl) {
		var ret string
		return ret
	}
	return *o.IdentityProviderSloUrl
}

// GetIdentityProviderSloUrlOk returns a tuple with the IdentityProviderSloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetIdentityProviderSloUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityProviderSloUrl) {
		return nil, false
	}
	return o.IdentityProviderSloUrl, true
}

// HasIdentityProviderSloUrl returns a boolean if a field has been set.
func (o *IamSamlSpConnection) HasIdentityProviderSloUrl() bool {
	if o != nil && !IsNil(o.IdentityProviderSloUrl) {
		return true
	}

	return false
}

// SetIdentityProviderSloUrl gets a reference to the given string and assigns it to the IdentityProviderSloUrl field.
func (o *IamSamlSpConnection) SetIdentityProviderSloUrl(v string) {
	o.IdentityProviderSloUrl = &v
}

// GetIdentityProviderSsoBinding returns the IdentityProviderSsoBinding field value if set, zero value otherwise.
func (o *IamSamlSpConnection) GetIdentityProviderSsoBinding() string {
	if o == nil || IsNil(o.IdentityProviderSsoBinding) {
		var ret string
		return ret
	}
	return *o.IdentityProviderSsoBinding
}

// GetIdentityProviderSsoBindingOk returns a tuple with the IdentityProviderSsoBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetIdentityProviderSsoBindingOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityProviderSsoBinding) {
		return nil, false
	}
	return o.IdentityProviderSsoBinding, true
}

// HasIdentityProviderSsoBinding returns a boolean if a field has been set.
func (o *IamSamlSpConnection) HasIdentityProviderSsoBinding() bool {
	if o != nil && !IsNil(o.IdentityProviderSsoBinding) {
		return true
	}

	return false
}

// SetIdentityProviderSsoBinding gets a reference to the given string and assigns it to the IdentityProviderSsoBinding field.
func (o *IamSamlSpConnection) SetIdentityProviderSsoBinding(v string) {
	o.IdentityProviderSsoBinding = &v
}

// GetIdentityProviderSsoUrl returns the IdentityProviderSsoUrl field value if set, zero value otherwise.
func (o *IamSamlSpConnection) GetIdentityProviderSsoUrl() string {
	if o == nil || IsNil(o.IdentityProviderSsoUrl) {
		var ret string
		return ret
	}
	return *o.IdentityProviderSsoUrl
}

// GetIdentityProviderSsoUrlOk returns a tuple with the IdentityProviderSsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetIdentityProviderSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityProviderSsoUrl) {
		return nil, false
	}
	return o.IdentityProviderSsoUrl, true
}

// HasIdentityProviderSsoUrl returns a boolean if a field has been set.
func (o *IamSamlSpConnection) HasIdentityProviderSsoUrl() bool {
	if o != nil && !IsNil(o.IdentityProviderSsoUrl) {
		return true
	}

	return false
}

// SetIdentityProviderSsoUrl gets a reference to the given string and assigns it to the IdentityProviderSsoUrl field.
func (o *IamSamlSpConnection) SetIdentityProviderSsoUrl(v string) {
	o.IdentityProviderSsoUrl = &v
}

// GetIdpCertificateStore returns the IdpCertificateStore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSamlSpConnection) GetIdpCertificateStore() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IdpCertificateStore
}

// GetIdpCertificateStoreOk returns a tuple with the IdpCertificateStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSamlSpConnection) GetIdpCertificateStoreOk() (*interface{}, bool) {
	if o == nil || IsNil(o.IdpCertificateStore) {
		return nil, false
	}
	return &o.IdpCertificateStore, true
}

// HasIdpCertificateStore returns a boolean if a field has been set.
func (o *IamSamlSpConnection) HasIdpCertificateStore() bool {
	if o != nil && !IsNil(o.IdpCertificateStore) {
		return true
	}

	return false
}

// SetIdpCertificateStore gets a reference to the given interface{} and assigns it to the IdpCertificateStore field.
func (o *IamSamlSpConnection) SetIdpCertificateStore(v interface{}) {
	o.IdpCertificateStore = v
}

// GetSignAuthnRequests returns the SignAuthnRequests field value if set, zero value otherwise.
func (o *IamSamlSpConnection) GetSignAuthnRequests() bool {
	if o == nil || IsNil(o.SignAuthnRequests) {
		var ret bool
		return ret
	}
	return *o.SignAuthnRequests
}

// GetSignAuthnRequestsOk returns a tuple with the SignAuthnRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnection) GetSignAuthnRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAuthnRequests) {
		return nil, false
	}
	return o.SignAuthnRequests, true
}

// HasSignAuthnRequests returns a boolean if a field has been set.
func (o *IamSamlSpConnection) HasSignAuthnRequests() bool {
	if o != nil && !IsNil(o.SignAuthnRequests) {
		return true
	}

	return false
}

// SetSignAuthnRequests gets a reference to the given bool and assigns it to the SignAuthnRequests field.
func (o *IamSamlSpConnection) SetSignAuthnRequests(v bool) {
	o.SignAuthnRequests = &v
}

func (o IamSamlSpConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamSamlSpConnection) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IdentityProviderSloBinding) {
		toSerialize["IdentityProviderSloBinding"] = o.IdentityProviderSloBinding
	}
	if !IsNil(o.IdentityProviderSloUrl) {
		toSerialize["IdentityProviderSloUrl"] = o.IdentityProviderSloUrl
	}
	if !IsNil(o.IdentityProviderSsoBinding) {
		toSerialize["IdentityProviderSsoBinding"] = o.IdentityProviderSsoBinding
	}
	if !IsNil(o.IdentityProviderSsoUrl) {
		toSerialize["IdentityProviderSsoUrl"] = o.IdentityProviderSsoUrl
	}
	if o.IdpCertificateStore != nil {
		toSerialize["IdpCertificateStore"] = o.IdpCertificateStore
	}
	if !IsNil(o.SignAuthnRequests) {
		toSerialize["SignAuthnRequests"] = o.SignAuthnRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamSamlSpConnection) UnmarshalJSON(data []byte) (err error) {
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
	type IamSamlSpConnectionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// SingleLogout Services binding in IdP Metadata.
		IdentityProviderSloBinding *string `json:"IdentityProviderSloBinding,omitempty"`
		// SingleLogOut Services location in IdP Metadata.
		IdentityProviderSloUrl *string `json:"IdentityProviderSloUrl,omitempty"`
		// SingleSignOn Services binding in IdP Metadata.
		IdentityProviderSsoBinding *string `json:"IdentityProviderSsoBinding,omitempty"`
		// SingleSignOn Services location in IdP Metadata.
		IdentityProviderSsoUrl *string `json:"IdentityProviderSsoUrl,omitempty"`
		// Decoded Certificate from IdP Metatdata.
		IdpCertificateStore interface{} `json:"IdpCertificateStore,omitempty"`
		// WantAuthnRequestsSigned from IdP Metadata.
		SignAuthnRequests *bool `json:"SignAuthnRequests,omitempty"`
	}

	varIamSamlSpConnectionWithoutEmbeddedStruct := IamSamlSpConnectionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamSamlSpConnectionWithoutEmbeddedStruct)
	if err == nil {
		varIamSamlSpConnection := _IamSamlSpConnection{}
		varIamSamlSpConnection.ClassId = varIamSamlSpConnectionWithoutEmbeddedStruct.ClassId
		varIamSamlSpConnection.ObjectType = varIamSamlSpConnectionWithoutEmbeddedStruct.ObjectType
		varIamSamlSpConnection.IdentityProviderSloBinding = varIamSamlSpConnectionWithoutEmbeddedStruct.IdentityProviderSloBinding
		varIamSamlSpConnection.IdentityProviderSloUrl = varIamSamlSpConnectionWithoutEmbeddedStruct.IdentityProviderSloUrl
		varIamSamlSpConnection.IdentityProviderSsoBinding = varIamSamlSpConnectionWithoutEmbeddedStruct.IdentityProviderSsoBinding
		varIamSamlSpConnection.IdentityProviderSsoUrl = varIamSamlSpConnectionWithoutEmbeddedStruct.IdentityProviderSsoUrl
		varIamSamlSpConnection.IdpCertificateStore = varIamSamlSpConnectionWithoutEmbeddedStruct.IdpCertificateStore
		varIamSamlSpConnection.SignAuthnRequests = varIamSamlSpConnectionWithoutEmbeddedStruct.SignAuthnRequests
		*o = IamSamlSpConnection(varIamSamlSpConnection)
	} else {
		return err
	}

	varIamSamlSpConnection := _IamSamlSpConnection{}

	err = json.Unmarshal(data, &varIamSamlSpConnection)
	if err == nil {
		o.MoBaseComplexType = varIamSamlSpConnection.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdentityProviderSloBinding")
		delete(additionalProperties, "IdentityProviderSloUrl")
		delete(additionalProperties, "IdentityProviderSsoBinding")
		delete(additionalProperties, "IdentityProviderSsoUrl")
		delete(additionalProperties, "IdpCertificateStore")
		delete(additionalProperties, "SignAuthnRequests")

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

type NullableIamSamlSpConnection struct {
	value *IamSamlSpConnection
	isSet bool
}

func (v NullableIamSamlSpConnection) Get() *IamSamlSpConnection {
	return v.value
}

func (v *NullableIamSamlSpConnection) Set(val *IamSamlSpConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSamlSpConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSamlSpConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSamlSpConnection(val *IamSamlSpConnection) *NullableIamSamlSpConnection {
	return &NullableIamSamlSpConnection{value: val, isSet: true}
}

func (v NullableIamSamlSpConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSamlSpConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
