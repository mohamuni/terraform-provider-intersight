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
	"reflect"
	"strings"
)

// IamUserSetting Holder for UI Settings such as preference of the user for Session Recording.
type IamUserSetting struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UI preference of the user for Session Recording.
	AllowUiSessionRecording *bool `json:"AllowUiSessionRecording,omitempty"`
	// UserID or email as configured in the IdP.
	UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
	// Unique id of the user used by the identity provider to store the user.
	UserUniqueIdentifier *string                      `json:"UserUniqueIdentifier,omitempty"`
	Idp                  *IamIdpRelationship          `json:"Idp,omitempty"`
	IdpReference         *IamIdpReferenceRelationship `json:"IdpReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamUserSetting IamUserSetting

// NewIamUserSetting instantiates a new IamUserSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserSetting(classId string, objectType string) *IamUserSetting {
	this := IamUserSetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allowUiSessionRecording bool = true
	this.AllowUiSessionRecording = &allowUiSessionRecording
	return &this
}

// NewIamUserSettingWithDefaults instantiates a new IamUserSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserSettingWithDefaults() *IamUserSetting {
	this := IamUserSetting{}
	var classId string = "iam.UserSetting"
	this.ClassId = classId
	var objectType string = "iam.UserSetting"
	this.ObjectType = objectType
	var allowUiSessionRecording bool = true
	this.AllowUiSessionRecording = &allowUiSessionRecording
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamUserSetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamUserSetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamUserSetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamUserSetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllowUiSessionRecording returns the AllowUiSessionRecording field value if set, zero value otherwise.
func (o *IamUserSetting) GetAllowUiSessionRecording() bool {
	if o == nil || o.AllowUiSessionRecording == nil {
		var ret bool
		return ret
	}
	return *o.AllowUiSessionRecording
}

// GetAllowUiSessionRecordingOk returns a tuple with the AllowUiSessionRecording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetAllowUiSessionRecordingOk() (*bool, bool) {
	if o == nil || o.AllowUiSessionRecording == nil {
		return nil, false
	}
	return o.AllowUiSessionRecording, true
}

// HasAllowUiSessionRecording returns a boolean if a field has been set.
func (o *IamUserSetting) HasAllowUiSessionRecording() bool {
	if o != nil && o.AllowUiSessionRecording != nil {
		return true
	}

	return false
}

// SetAllowUiSessionRecording gets a reference to the given bool and assigns it to the AllowUiSessionRecording field.
func (o *IamUserSetting) SetAllowUiSessionRecording(v bool) {
	o.AllowUiSessionRecording = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *IamUserSetting) GetUserIdOrEmail() string {
	if o == nil || o.UserIdOrEmail == nil {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || o.UserIdOrEmail == nil {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *IamUserSetting) HasUserIdOrEmail() bool {
	if o != nil && o.UserIdOrEmail != nil {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *IamUserSetting) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetUserUniqueIdentifier returns the UserUniqueIdentifier field value if set, zero value otherwise.
func (o *IamUserSetting) GetUserUniqueIdentifier() string {
	if o == nil || o.UserUniqueIdentifier == nil {
		var ret string
		return ret
	}
	return *o.UserUniqueIdentifier
}

// GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetUserUniqueIdentifierOk() (*string, bool) {
	if o == nil || o.UserUniqueIdentifier == nil {
		return nil, false
	}
	return o.UserUniqueIdentifier, true
}

// HasUserUniqueIdentifier returns a boolean if a field has been set.
func (o *IamUserSetting) HasUserUniqueIdentifier() bool {
	if o != nil && o.UserUniqueIdentifier != nil {
		return true
	}

	return false
}

// SetUserUniqueIdentifier gets a reference to the given string and assigns it to the UserUniqueIdentifier field.
func (o *IamUserSetting) SetUserUniqueIdentifier(v string) {
	o.UserUniqueIdentifier = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamUserSetting) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUserSetting) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamUserSetting) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpReference returns the IdpReference field value if set, zero value otherwise.
func (o *IamUserSetting) GetIdpReference() IamIdpReferenceRelationship {
	if o == nil || o.IdpReference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.IdpReference
}

// GetIdpReferenceOk returns a tuple with the IdpReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserSetting) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.IdpReference == nil {
		return nil, false
	}
	return o.IdpReference, true
}

// HasIdpReference returns a boolean if a field has been set.
func (o *IamUserSetting) HasIdpReference() bool {
	if o != nil && o.IdpReference != nil {
		return true
	}

	return false
}

// SetIdpReference gets a reference to the given IamIdpReferenceRelationship and assigns it to the IdpReference field.
func (o *IamUserSetting) SetIdpReference(v IamIdpReferenceRelationship) {
	o.IdpReference = &v
}

func (o IamUserSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllowUiSessionRecording != nil {
		toSerialize["AllowUiSessionRecording"] = o.AllowUiSessionRecording
	}
	if o.UserIdOrEmail != nil {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if o.UserUniqueIdentifier != nil {
		toSerialize["UserUniqueIdentifier"] = o.UserUniqueIdentifier
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.IdpReference != nil {
		toSerialize["IdpReference"] = o.IdpReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamUserSetting) UnmarshalJSON(bytes []byte) (err error) {
	type IamUserSettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UI preference of the user for Session Recording.
		AllowUiSessionRecording *bool `json:"AllowUiSessionRecording,omitempty"`
		// UserID or email as configured in the IdP.
		UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
		// Unique id of the user used by the identity provider to store the user.
		UserUniqueIdentifier *string                      `json:"UserUniqueIdentifier,omitempty"`
		Idp                  *IamIdpRelationship          `json:"Idp,omitempty"`
		IdpReference         *IamIdpReferenceRelationship `json:"IdpReference,omitempty"`
	}

	varIamUserSettingWithoutEmbeddedStruct := IamUserSettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamUserSettingWithoutEmbeddedStruct)
	if err == nil {
		varIamUserSetting := _IamUserSetting{}
		varIamUserSetting.ClassId = varIamUserSettingWithoutEmbeddedStruct.ClassId
		varIamUserSetting.ObjectType = varIamUserSettingWithoutEmbeddedStruct.ObjectType
		varIamUserSetting.AllowUiSessionRecording = varIamUserSettingWithoutEmbeddedStruct.AllowUiSessionRecording
		varIamUserSetting.UserIdOrEmail = varIamUserSettingWithoutEmbeddedStruct.UserIdOrEmail
		varIamUserSetting.UserUniqueIdentifier = varIamUserSettingWithoutEmbeddedStruct.UserUniqueIdentifier
		varIamUserSetting.Idp = varIamUserSettingWithoutEmbeddedStruct.Idp
		varIamUserSetting.IdpReference = varIamUserSettingWithoutEmbeddedStruct.IdpReference
		*o = IamUserSetting(varIamUserSetting)
	} else {
		return err
	}

	varIamUserSetting := _IamUserSetting{}

	err = json.Unmarshal(bytes, &varIamUserSetting)
	if err == nil {
		o.MoBaseMo = varIamUserSetting.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowUiSessionRecording")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "UserUniqueIdentifier")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "IdpReference")

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

type NullableIamUserSetting struct {
	value *IamUserSetting
	isSet bool
}

func (v NullableIamUserSetting) Get() *IamUserSetting {
	return v.value
}

func (v *NullableIamUserSetting) Set(val *IamUserSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserSetting(val *IamUserSetting) *NullableIamUserSetting {
	return &NullableIamUserSetting{value: val, isSet: true}
}

func (v NullableIamUserSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}