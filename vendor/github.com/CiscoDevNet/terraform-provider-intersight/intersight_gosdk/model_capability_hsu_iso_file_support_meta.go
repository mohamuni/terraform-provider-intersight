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

// checks if the CapabilityHsuIsoFileSupportMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityHsuIsoFileSupportMeta{}

// CapabilityHsuIsoFileSupportMeta Internal meta-data to check the HSU support to accept an iso file path in the upgrade request.
type CapabilityHsuIsoFileSupportMeta struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the symbolic link to the actual iso file in the default case.
	DefaultFileName *string `json:"DefaultFileName,omitempty"`
	// Firmware version from which the HSU capability is present in the default case.
	DefaultMinVersion       *string                                   `json:"DefaultMinVersion,omitempty"`
	ModelSpecificConstraint []CapabilityHsuIsoModelSpecificConstraint `json:"ModelSpecificConstraint,omitempty"`
	// Series name for the capability group.
	Series               *string `json:"Series,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityHsuIsoFileSupportMeta CapabilityHsuIsoFileSupportMeta

// NewCapabilityHsuIsoFileSupportMeta instantiates a new CapabilityHsuIsoFileSupportMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityHsuIsoFileSupportMeta(classId string, objectType string) *CapabilityHsuIsoFileSupportMeta {
	this := CapabilityHsuIsoFileSupportMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityHsuIsoFileSupportMetaWithDefaults instantiates a new CapabilityHsuIsoFileSupportMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityHsuIsoFileSupportMetaWithDefaults() *CapabilityHsuIsoFileSupportMeta {
	this := CapabilityHsuIsoFileSupportMeta{}
	var classId string = "capability.HsuIsoFileSupportMeta"
	this.ClassId = classId
	var objectType string = "capability.HsuIsoFileSupportMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityHsuIsoFileSupportMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityHsuIsoFileSupportMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityHsuIsoFileSupportMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.HsuIsoFileSupportMeta" of the ClassId field.
func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultClassId() interface{} {
	return "capability.HsuIsoFileSupportMeta"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityHsuIsoFileSupportMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityHsuIsoFileSupportMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityHsuIsoFileSupportMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.HsuIsoFileSupportMeta" of the ObjectType field.
func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultObjectType() interface{} {
	return "capability.HsuIsoFileSupportMeta"
}

// GetDefaultFileName returns the DefaultFileName field value if set, zero value otherwise.
func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultFileName() string {
	if o == nil || IsNil(o.DefaultFileName) {
		var ret string
		return ret
	}
	return *o.DefaultFileName
}

// GetDefaultFileNameOk returns a tuple with the DefaultFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultFileName) {
		return nil, false
	}
	return o.DefaultFileName, true
}

// HasDefaultFileName returns a boolean if a field has been set.
func (o *CapabilityHsuIsoFileSupportMeta) HasDefaultFileName() bool {
	if o != nil && !IsNil(o.DefaultFileName) {
		return true
	}

	return false
}

// SetDefaultFileName gets a reference to the given string and assigns it to the DefaultFileName field.
func (o *CapabilityHsuIsoFileSupportMeta) SetDefaultFileName(v string) {
	o.DefaultFileName = &v
}

// GetDefaultMinVersion returns the DefaultMinVersion field value if set, zero value otherwise.
func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultMinVersion() string {
	if o == nil || IsNil(o.DefaultMinVersion) {
		var ret string
		return ret
	}
	return *o.DefaultMinVersion
}

// GetDefaultMinVersionOk returns a tuple with the DefaultMinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityHsuIsoFileSupportMeta) GetDefaultMinVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultMinVersion) {
		return nil, false
	}
	return o.DefaultMinVersion, true
}

// HasDefaultMinVersion returns a boolean if a field has been set.
func (o *CapabilityHsuIsoFileSupportMeta) HasDefaultMinVersion() bool {
	if o != nil && !IsNil(o.DefaultMinVersion) {
		return true
	}

	return false
}

// SetDefaultMinVersion gets a reference to the given string and assigns it to the DefaultMinVersion field.
func (o *CapabilityHsuIsoFileSupportMeta) SetDefaultMinVersion(v string) {
	o.DefaultMinVersion = &v
}

// GetModelSpecificConstraint returns the ModelSpecificConstraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityHsuIsoFileSupportMeta) GetModelSpecificConstraint() []CapabilityHsuIsoModelSpecificConstraint {
	if o == nil {
		var ret []CapabilityHsuIsoModelSpecificConstraint
		return ret
	}
	return o.ModelSpecificConstraint
}

// GetModelSpecificConstraintOk returns a tuple with the ModelSpecificConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityHsuIsoFileSupportMeta) GetModelSpecificConstraintOk() ([]CapabilityHsuIsoModelSpecificConstraint, bool) {
	if o == nil || IsNil(o.ModelSpecificConstraint) {
		return nil, false
	}
	return o.ModelSpecificConstraint, true
}

// HasModelSpecificConstraint returns a boolean if a field has been set.
func (o *CapabilityHsuIsoFileSupportMeta) HasModelSpecificConstraint() bool {
	if o != nil && !IsNil(o.ModelSpecificConstraint) {
		return true
	}

	return false
}

// SetModelSpecificConstraint gets a reference to the given []CapabilityHsuIsoModelSpecificConstraint and assigns it to the ModelSpecificConstraint field.
func (o *CapabilityHsuIsoFileSupportMeta) SetModelSpecificConstraint(v []CapabilityHsuIsoModelSpecificConstraint) {
	o.ModelSpecificConstraint = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *CapabilityHsuIsoFileSupportMeta) GetSeries() string {
	if o == nil || IsNil(o.Series) {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityHsuIsoFileSupportMeta) GetSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *CapabilityHsuIsoFileSupportMeta) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *CapabilityHsuIsoFileSupportMeta) SetSeries(v string) {
	o.Series = &v
}

func (o CapabilityHsuIsoFileSupportMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityHsuIsoFileSupportMeta) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DefaultFileName) {
		toSerialize["DefaultFileName"] = o.DefaultFileName
	}
	if !IsNil(o.DefaultMinVersion) {
		toSerialize["DefaultMinVersion"] = o.DefaultMinVersion
	}
	if o.ModelSpecificConstraint != nil {
		toSerialize["ModelSpecificConstraint"] = o.ModelSpecificConstraint
	}
	if !IsNil(o.Series) {
		toSerialize["Series"] = o.Series
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityHsuIsoFileSupportMeta) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the symbolic link to the actual iso file in the default case.
		DefaultFileName *string `json:"DefaultFileName,omitempty"`
		// Firmware version from which the HSU capability is present in the default case.
		DefaultMinVersion       *string                                   `json:"DefaultMinVersion,omitempty"`
		ModelSpecificConstraint []CapabilityHsuIsoModelSpecificConstraint `json:"ModelSpecificConstraint,omitempty"`
		// Series name for the capability group.
		Series *string `json:"Series,omitempty"`
	}

	varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct := CapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityHsuIsoFileSupportMeta := _CapabilityHsuIsoFileSupportMeta{}
		varCapabilityHsuIsoFileSupportMeta.ClassId = varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct.ClassId
		varCapabilityHsuIsoFileSupportMeta.ObjectType = varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct.ObjectType
		varCapabilityHsuIsoFileSupportMeta.DefaultFileName = varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct.DefaultFileName
		varCapabilityHsuIsoFileSupportMeta.DefaultMinVersion = varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct.DefaultMinVersion
		varCapabilityHsuIsoFileSupportMeta.ModelSpecificConstraint = varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct.ModelSpecificConstraint
		varCapabilityHsuIsoFileSupportMeta.Series = varCapabilityHsuIsoFileSupportMetaWithoutEmbeddedStruct.Series
		*o = CapabilityHsuIsoFileSupportMeta(varCapabilityHsuIsoFileSupportMeta)
	} else {
		return err
	}

	varCapabilityHsuIsoFileSupportMeta := _CapabilityHsuIsoFileSupportMeta{}

	err = json.Unmarshal(data, &varCapabilityHsuIsoFileSupportMeta)
	if err == nil {
		o.CapabilityCapability = varCapabilityHsuIsoFileSupportMeta.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DefaultFileName")
		delete(additionalProperties, "DefaultMinVersion")
		delete(additionalProperties, "ModelSpecificConstraint")
		delete(additionalProperties, "Series")

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

type NullableCapabilityHsuIsoFileSupportMeta struct {
	value *CapabilityHsuIsoFileSupportMeta
	isSet bool
}

func (v NullableCapabilityHsuIsoFileSupportMeta) Get() *CapabilityHsuIsoFileSupportMeta {
	return v.value
}

func (v *NullableCapabilityHsuIsoFileSupportMeta) Set(val *CapabilityHsuIsoFileSupportMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityHsuIsoFileSupportMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityHsuIsoFileSupportMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityHsuIsoFileSupportMeta(val *CapabilityHsuIsoFileSupportMeta) *NullableCapabilityHsuIsoFileSupportMeta {
	return &NullableCapabilityHsuIsoFileSupportMeta{value: val, isSet: true}
}

func (v NullableCapabilityHsuIsoFileSupportMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityHsuIsoFileSupportMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
