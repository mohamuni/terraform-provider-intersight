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

// AssetTargetKeyAllOf Definition of the list of properties defined in 'asset.TargetKey', excluding properties defined in parent classes.
type AssetTargetKeyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'privateKey' property has been set.
	IsPrivateKeySet      *bool `json:"IsPrivateKeySet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetTargetKeyAllOf AssetTargetKeyAllOf

// NewAssetTargetKeyAllOf instantiates a new AssetTargetKeyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTargetKeyAllOf(classId string, objectType string) *AssetTargetKeyAllOf {
	this := AssetTargetKeyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTargetKeyAllOfWithDefaults instantiates a new AssetTargetKeyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetKeyAllOfWithDefaults() *AssetTargetKeyAllOf {
	this := AssetTargetKeyAllOf{}
	var classId string = "asset.TargetKey"
	this.ClassId = classId
	var objectType string = "asset.TargetKey"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTargetKeyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTargetKeyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTargetKeyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTargetKeyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTargetKeyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTargetKeyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPrivateKeySet returns the IsPrivateKeySet field value if set, zero value otherwise.
func (o *AssetTargetKeyAllOf) GetIsPrivateKeySet() bool {
	if o == nil || o.IsPrivateKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivateKeySet
}

// GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetKeyAllOf) GetIsPrivateKeySetOk() (*bool, bool) {
	if o == nil || o.IsPrivateKeySet == nil {
		return nil, false
	}
	return o.IsPrivateKeySet, true
}

// HasIsPrivateKeySet returns a boolean if a field has been set.
func (o *AssetTargetKeyAllOf) HasIsPrivateKeySet() bool {
	if o != nil && o.IsPrivateKeySet != nil {
		return true
	}

	return false
}

// SetIsPrivateKeySet gets a reference to the given bool and assigns it to the IsPrivateKeySet field.
func (o *AssetTargetKeyAllOf) SetIsPrivateKeySet(v bool) {
	o.IsPrivateKeySet = &v
}

func (o AssetTargetKeyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPrivateKeySet != nil {
		toSerialize["IsPrivateKeySet"] = o.IsPrivateKeySet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTargetKeyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTargetKeyAllOf := _AssetTargetKeyAllOf{}

	if err = json.Unmarshal(bytes, &varAssetTargetKeyAllOf); err == nil {
		*o = AssetTargetKeyAllOf(varAssetTargetKeyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPrivateKeySet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTargetKeyAllOf struct {
	value *AssetTargetKeyAllOf
	isSet bool
}

func (v NullableAssetTargetKeyAllOf) Get() *AssetTargetKeyAllOf {
	return v.value
}

func (v *NullableAssetTargetKeyAllOf) Set(val *AssetTargetKeyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTargetKeyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTargetKeyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTargetKeyAllOf(val *AssetTargetKeyAllOf) *NullableAssetTargetKeyAllOf {
	return &NullableAssetTargetKeyAllOf{value: val, isSet: true}
}

func (v NullableAssetTargetKeyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTargetKeyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
