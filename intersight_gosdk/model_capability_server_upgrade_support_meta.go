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

// CapabilityServerUpgradeSupportMeta Internal meta-data to map server family classification from server model, used in f/w policy also.
type CapabilityServerUpgradeSupportMeta struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Verbose description regarding this group of server.
	Description *string `json:"Description,omitempty"`
	// The target platform for which the mapping is defined.
	Platform *string `json:"Platform,omitempty"`
	// Classification of a set of server models.
	ServerFamily         *string  `json:"ServerFamily,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityServerUpgradeSupportMeta CapabilityServerUpgradeSupportMeta

// NewCapabilityServerUpgradeSupportMeta instantiates a new CapabilityServerUpgradeSupportMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityServerUpgradeSupportMeta(classId string, objectType string) *CapabilityServerUpgradeSupportMeta {
	this := CapabilityServerUpgradeSupportMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityServerUpgradeSupportMetaWithDefaults instantiates a new CapabilityServerUpgradeSupportMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityServerUpgradeSupportMetaWithDefaults() *CapabilityServerUpgradeSupportMeta {
	this := CapabilityServerUpgradeSupportMeta{}
	var classId string = "capability.ServerUpgradeSupportMeta"
	this.ClassId = classId
	var objectType string = "capability.ServerUpgradeSupportMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityServerUpgradeSupportMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityServerUpgradeSupportMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityServerUpgradeSupportMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityServerUpgradeSupportMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityServerUpgradeSupportMeta) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMeta) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityServerUpgradeSupportMeta) SetDescription(v string) {
	o.Description = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CapabilityServerUpgradeSupportMeta) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMeta) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMeta) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *CapabilityServerUpgradeSupportMeta) SetPlatform(v string) {
	o.Platform = &v
}

// GetServerFamily returns the ServerFamily field value if set, zero value otherwise.
func (o *CapabilityServerUpgradeSupportMeta) GetServerFamily() string {
	if o == nil || o.ServerFamily == nil {
		var ret string
		return ret
	}
	return *o.ServerFamily
}

// GetServerFamilyOk returns a tuple with the ServerFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMeta) GetServerFamilyOk() (*string, bool) {
	if o == nil || o.ServerFamily == nil {
		return nil, false
	}
	return o.ServerFamily, true
}

// HasServerFamily returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMeta) HasServerFamily() bool {
	if o != nil && o.ServerFamily != nil {
		return true
	}

	return false
}

// SetServerFamily gets a reference to the given string and assigns it to the ServerFamily field.
func (o *CapabilityServerUpgradeSupportMeta) SetServerFamily(v string) {
	o.ServerFamily = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityServerUpgradeSupportMeta) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityServerUpgradeSupportMeta) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMeta) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *CapabilityServerUpgradeSupportMeta) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o CapabilityServerUpgradeSupportMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}
	if o.ServerFamily != nil {
		toSerialize["ServerFamily"] = o.ServerFamily
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityServerUpgradeSupportMeta) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Verbose description regarding this group of server.
		Description *string `json:"Description,omitempty"`
		// The target platform for which the mapping is defined.
		Platform *string `json:"Platform,omitempty"`
		// Classification of a set of server models.
		ServerFamily    *string  `json:"ServerFamily,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
	}

	varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct := CapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityServerUpgradeSupportMeta := _CapabilityServerUpgradeSupportMeta{}
		varCapabilityServerUpgradeSupportMeta.ClassId = varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct.ClassId
		varCapabilityServerUpgradeSupportMeta.ObjectType = varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct.ObjectType
		varCapabilityServerUpgradeSupportMeta.Description = varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct.Description
		varCapabilityServerUpgradeSupportMeta.Platform = varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct.Platform
		varCapabilityServerUpgradeSupportMeta.ServerFamily = varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct.ServerFamily
		varCapabilityServerUpgradeSupportMeta.SupportedModels = varCapabilityServerUpgradeSupportMetaWithoutEmbeddedStruct.SupportedModels
		*o = CapabilityServerUpgradeSupportMeta(varCapabilityServerUpgradeSupportMeta)
	} else {
		return err
	}

	varCapabilityServerUpgradeSupportMeta := _CapabilityServerUpgradeSupportMeta{}

	err = json.Unmarshal(bytes, &varCapabilityServerUpgradeSupportMeta)
	if err == nil {
		o.CapabilityCapability = varCapabilityServerUpgradeSupportMeta.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Platform")
		delete(additionalProperties, "ServerFamily")
		delete(additionalProperties, "SupportedModels")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityServerUpgradeSupportMeta struct {
	value *CapabilityServerUpgradeSupportMeta
	isSet bool
}

func (v NullableCapabilityServerUpgradeSupportMeta) Get() *CapabilityServerUpgradeSupportMeta {
	return v.value
}

func (v *NullableCapabilityServerUpgradeSupportMeta) Set(val *CapabilityServerUpgradeSupportMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityServerUpgradeSupportMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityServerUpgradeSupportMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityServerUpgradeSupportMeta(val *CapabilityServerUpgradeSupportMeta) *NullableCapabilityServerUpgradeSupportMeta {
	return &NullableCapabilityServerUpgradeSupportMeta{value: val, isSet: true}
}

func (v NullableCapabilityServerUpgradeSupportMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityServerUpgradeSupportMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
