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
)

// IamSharingRuleAllOf Definition of the list of properties defined in 'iam.SharingRule', excluding properties defined in parent classes.
type IamSharingRuleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                  `json:"ObjectType"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	SharedResource       *MoBaseMoRelationship   `json:"SharedResource,omitempty"`
	SharedWithResource   *MoBaseMoRelationship   `json:"SharedWithResource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSharingRuleAllOf IamSharingRuleAllOf

// NewIamSharingRuleAllOf instantiates a new IamSharingRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSharingRuleAllOf(classId string, objectType string) *IamSharingRuleAllOf {
	this := IamSharingRuleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSharingRuleAllOfWithDefaults instantiates a new IamSharingRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSharingRuleAllOfWithDefaults() *IamSharingRuleAllOf {
	this := IamSharingRuleAllOf{}
	var classId string = "iam.SharingRule"
	this.ClassId = classId
	var objectType string = "iam.SharingRule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSharingRuleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSharingRuleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSharingRuleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSharingRuleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSharingRuleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSharingRuleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamSharingRuleAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSharingRuleAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamSharingRuleAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamSharingRuleAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetSharedResource returns the SharedResource field value if set, zero value otherwise.
func (o *IamSharingRuleAllOf) GetSharedResource() MoBaseMoRelationship {
	if o == nil || o.SharedResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.SharedResource
}

// GetSharedResourceOk returns a tuple with the SharedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSharingRuleAllOf) GetSharedResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.SharedResource == nil {
		return nil, false
	}
	return o.SharedResource, true
}

// HasSharedResource returns a boolean if a field has been set.
func (o *IamSharingRuleAllOf) HasSharedResource() bool {
	if o != nil && o.SharedResource != nil {
		return true
	}

	return false
}

// SetSharedResource gets a reference to the given MoBaseMoRelationship and assigns it to the SharedResource field.
func (o *IamSharingRuleAllOf) SetSharedResource(v MoBaseMoRelationship) {
	o.SharedResource = &v
}

// GetSharedWithResource returns the SharedWithResource field value if set, zero value otherwise.
func (o *IamSharingRuleAllOf) GetSharedWithResource() MoBaseMoRelationship {
	if o == nil || o.SharedWithResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.SharedWithResource
}

// GetSharedWithResourceOk returns a tuple with the SharedWithResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSharingRuleAllOf) GetSharedWithResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.SharedWithResource == nil {
		return nil, false
	}
	return o.SharedWithResource, true
}

// HasSharedWithResource returns a boolean if a field has been set.
func (o *IamSharingRuleAllOf) HasSharedWithResource() bool {
	if o != nil && o.SharedWithResource != nil {
		return true
	}

	return false
}

// SetSharedWithResource gets a reference to the given MoBaseMoRelationship and assigns it to the SharedWithResource field.
func (o *IamSharingRuleAllOf) SetSharedWithResource(v MoBaseMoRelationship) {
	o.SharedWithResource = &v
}

func (o IamSharingRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.SharedResource != nil {
		toSerialize["SharedResource"] = o.SharedResource
	}
	if o.SharedWithResource != nil {
		toSerialize["SharedWithResource"] = o.SharedWithResource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSharingRuleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamSharingRuleAllOf := _IamSharingRuleAllOf{}

	if err = json.Unmarshal(bytes, &varIamSharingRuleAllOf); err == nil {
		*o = IamSharingRuleAllOf(varIamSharingRuleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "SharedResource")
		delete(additionalProperties, "SharedWithResource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamSharingRuleAllOf struct {
	value *IamSharingRuleAllOf
	isSet bool
}

func (v NullableIamSharingRuleAllOf) Get() *IamSharingRuleAllOf {
	return v.value
}

func (v *NullableIamSharingRuleAllOf) Set(val *IamSharingRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSharingRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSharingRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSharingRuleAllOf(val *IamSharingRuleAllOf) *NullableIamSharingRuleAllOf {
	return &NullableIamSharingRuleAllOf{value: val, isSet: true}
}

func (v NullableIamSharingRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSharingRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
