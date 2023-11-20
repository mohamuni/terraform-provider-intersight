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

// AssetSubscriptionAccount It is a empty placeholder object for Account collection. Its purpose is to avoid data updates to Subscription related MOs on Account MO change event.
type AssetSubscriptionAccount struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                  `json:"ObjectType"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSubscriptionAccount AssetSubscriptionAccount

// NewAssetSubscriptionAccount instantiates a new AssetSubscriptionAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSubscriptionAccount(classId string, objectType string) *AssetSubscriptionAccount {
	this := AssetSubscriptionAccount{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetSubscriptionAccountWithDefaults instantiates a new AssetSubscriptionAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSubscriptionAccountWithDefaults() *AssetSubscriptionAccount {
	this := AssetSubscriptionAccount{}
	var classId string = "asset.SubscriptionAccount"
	this.ClassId = classId
	var objectType string = "asset.SubscriptionAccount"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSubscriptionAccount) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAccount) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSubscriptionAccount) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSubscriptionAccount) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAccount) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSubscriptionAccount) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetSubscriptionAccount) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAccount) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetSubscriptionAccount) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetSubscriptionAccount) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o AssetSubscriptionAccount) MarshalJSON() ([]byte, error) {
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
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSubscriptionAccount) UnmarshalJSON(bytes []byte) (err error) {
	type AssetSubscriptionAccountWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                  `json:"ObjectType"`
		Account    *IamAccountRelationship `json:"Account,omitempty"`
	}

	varAssetSubscriptionAccountWithoutEmbeddedStruct := AssetSubscriptionAccountWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetSubscriptionAccountWithoutEmbeddedStruct)
	if err == nil {
		varAssetSubscriptionAccount := _AssetSubscriptionAccount{}
		varAssetSubscriptionAccount.ClassId = varAssetSubscriptionAccountWithoutEmbeddedStruct.ClassId
		varAssetSubscriptionAccount.ObjectType = varAssetSubscriptionAccountWithoutEmbeddedStruct.ObjectType
		varAssetSubscriptionAccount.Account = varAssetSubscriptionAccountWithoutEmbeddedStruct.Account
		*o = AssetSubscriptionAccount(varAssetSubscriptionAccount)
	} else {
		return err
	}

	varAssetSubscriptionAccount := _AssetSubscriptionAccount{}

	err = json.Unmarshal(bytes, &varAssetSubscriptionAccount)
	if err == nil {
		o.MoBaseMo = varAssetSubscriptionAccount.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")

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

type NullableAssetSubscriptionAccount struct {
	value *AssetSubscriptionAccount
	isSet bool
}

func (v NullableAssetSubscriptionAccount) Get() *AssetSubscriptionAccount {
	return v.value
}

func (v *NullableAssetSubscriptionAccount) Set(val *AssetSubscriptionAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSubscriptionAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSubscriptionAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSubscriptionAccount(val *AssetSubscriptionAccount) *NullableAssetSubscriptionAccount {
	return &NullableAssetSubscriptionAccount{value: val, isSet: true}
}

func (v NullableAssetSubscriptionAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSubscriptionAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
