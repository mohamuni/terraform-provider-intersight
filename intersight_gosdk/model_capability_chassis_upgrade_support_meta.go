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

// checks if the CapabilityChassisUpgradeSupportMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityChassisUpgradeSupportMeta{}

// CapabilityChassisUpgradeSupportMeta Internal meta-data to enable chassis upgrade related decision making.
type CapabilityChassisUpgradeSupportMeta struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, indicates that adapters in servers within this chassis would be upgraded by HSU.
	AdaptersUpgradedViaHsu *bool `json:"AdaptersUpgradedViaHsu,omitempty"`
	// Verbose description regarding this group of chassis.
	Description *string `json:"Description,omitempty"`
	// Classification of a set of chassis models.
	SeriesId             *string  `json:"SeriesId,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityChassisUpgradeSupportMeta CapabilityChassisUpgradeSupportMeta

// NewCapabilityChassisUpgradeSupportMeta instantiates a new CapabilityChassisUpgradeSupportMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityChassisUpgradeSupportMeta(classId string, objectType string) *CapabilityChassisUpgradeSupportMeta {
	this := CapabilityChassisUpgradeSupportMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityChassisUpgradeSupportMetaWithDefaults instantiates a new CapabilityChassisUpgradeSupportMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityChassisUpgradeSupportMetaWithDefaults() *CapabilityChassisUpgradeSupportMeta {
	this := CapabilityChassisUpgradeSupportMeta{}
	var classId string = "capability.ChassisUpgradeSupportMeta"
	this.ClassId = classId
	var objectType string = "capability.ChassisUpgradeSupportMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityChassisUpgradeSupportMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityChassisUpgradeSupportMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityChassisUpgradeSupportMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.ChassisUpgradeSupportMeta" of the ClassId field.
func (o *CapabilityChassisUpgradeSupportMeta) GetDefaultClassId() interface{} {
	return "capability.ChassisUpgradeSupportMeta"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityChassisUpgradeSupportMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityChassisUpgradeSupportMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityChassisUpgradeSupportMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.ChassisUpgradeSupportMeta" of the ObjectType field.
func (o *CapabilityChassisUpgradeSupportMeta) GetDefaultObjectType() interface{} {
	return "capability.ChassisUpgradeSupportMeta"
}

// GetAdaptersUpgradedViaHsu returns the AdaptersUpgradedViaHsu field value if set, zero value otherwise.
func (o *CapabilityChassisUpgradeSupportMeta) GetAdaptersUpgradedViaHsu() bool {
	if o == nil || IsNil(o.AdaptersUpgradedViaHsu) {
		var ret bool
		return ret
	}
	return *o.AdaptersUpgradedViaHsu
}

// GetAdaptersUpgradedViaHsuOk returns a tuple with the AdaptersUpgradedViaHsu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityChassisUpgradeSupportMeta) GetAdaptersUpgradedViaHsuOk() (*bool, bool) {
	if o == nil || IsNil(o.AdaptersUpgradedViaHsu) {
		return nil, false
	}
	return o.AdaptersUpgradedViaHsu, true
}

// HasAdaptersUpgradedViaHsu returns a boolean if a field has been set.
func (o *CapabilityChassisUpgradeSupportMeta) HasAdaptersUpgradedViaHsu() bool {
	if o != nil && !IsNil(o.AdaptersUpgradedViaHsu) {
		return true
	}

	return false
}

// SetAdaptersUpgradedViaHsu gets a reference to the given bool and assigns it to the AdaptersUpgradedViaHsu field.
func (o *CapabilityChassisUpgradeSupportMeta) SetAdaptersUpgradedViaHsu(v bool) {
	o.AdaptersUpgradedViaHsu = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityChassisUpgradeSupportMeta) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityChassisUpgradeSupportMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityChassisUpgradeSupportMeta) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityChassisUpgradeSupportMeta) SetDescription(v string) {
	o.Description = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *CapabilityChassisUpgradeSupportMeta) GetSeriesId() string {
	if o == nil || IsNil(o.SeriesId) {
		var ret string
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityChassisUpgradeSupportMeta) GetSeriesIdOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *CapabilityChassisUpgradeSupportMeta) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given string and assigns it to the SeriesId field.
func (o *CapabilityChassisUpgradeSupportMeta) SetSeriesId(v string) {
	o.SeriesId = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityChassisUpgradeSupportMeta) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityChassisUpgradeSupportMeta) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *CapabilityChassisUpgradeSupportMeta) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *CapabilityChassisUpgradeSupportMeta) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o CapabilityChassisUpgradeSupportMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityChassisUpgradeSupportMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdaptersUpgradedViaHsu) {
		toSerialize["AdaptersUpgradedViaHsu"] = o.AdaptersUpgradedViaHsu
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.SeriesId) {
		toSerialize["SeriesId"] = o.SeriesId
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityChassisUpgradeSupportMeta) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If enabled, indicates that adapters in servers within this chassis would be upgraded by HSU.
		AdaptersUpgradedViaHsu *bool `json:"AdaptersUpgradedViaHsu,omitempty"`
		// Verbose description regarding this group of chassis.
		Description *string `json:"Description,omitempty"`
		// Classification of a set of chassis models.
		SeriesId        *string  `json:"SeriesId,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
	}

	varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct := CapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityChassisUpgradeSupportMeta := _CapabilityChassisUpgradeSupportMeta{}
		varCapabilityChassisUpgradeSupportMeta.ClassId = varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct.ClassId
		varCapabilityChassisUpgradeSupportMeta.ObjectType = varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct.ObjectType
		varCapabilityChassisUpgradeSupportMeta.AdaptersUpgradedViaHsu = varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct.AdaptersUpgradedViaHsu
		varCapabilityChassisUpgradeSupportMeta.Description = varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct.Description
		varCapabilityChassisUpgradeSupportMeta.SeriesId = varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct.SeriesId
		varCapabilityChassisUpgradeSupportMeta.SupportedModels = varCapabilityChassisUpgradeSupportMetaWithoutEmbeddedStruct.SupportedModels
		*o = CapabilityChassisUpgradeSupportMeta(varCapabilityChassisUpgradeSupportMeta)
	} else {
		return err
	}

	varCapabilityChassisUpgradeSupportMeta := _CapabilityChassisUpgradeSupportMeta{}

	err = json.Unmarshal(data, &varCapabilityChassisUpgradeSupportMeta)
	if err == nil {
		o.CapabilityCapability = varCapabilityChassisUpgradeSupportMeta.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdaptersUpgradedViaHsu")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "SeriesId")
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

type NullableCapabilityChassisUpgradeSupportMeta struct {
	value *CapabilityChassisUpgradeSupportMeta
	isSet bool
}

func (v NullableCapabilityChassisUpgradeSupportMeta) Get() *CapabilityChassisUpgradeSupportMeta {
	return v.value
}

func (v *NullableCapabilityChassisUpgradeSupportMeta) Set(val *CapabilityChassisUpgradeSupportMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityChassisUpgradeSupportMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityChassisUpgradeSupportMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityChassisUpgradeSupportMeta(val *CapabilityChassisUpgradeSupportMeta) *NullableCapabilityChassisUpgradeSupportMeta {
	return &NullableCapabilityChassisUpgradeSupportMeta{value: val, isSet: true}
}

func (v NullableCapabilityChassisUpgradeSupportMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityChassisUpgradeSupportMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
