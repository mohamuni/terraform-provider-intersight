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

// checks if the CapabilityTemplateCatalog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityTemplateCatalog{}

// CapabilityTemplateCatalog Catalog of override allowed configurations for vNIC and vHBA templates.
type CapabilityTemplateCatalog struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	AllowedOverrideList  []string `json:"AllowedOverrideList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityTemplateCatalog CapabilityTemplateCatalog

// NewCapabilityTemplateCatalog instantiates a new CapabilityTemplateCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityTemplateCatalog(classId string, objectType string) *CapabilityTemplateCatalog {
	this := CapabilityTemplateCatalog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityTemplateCatalogWithDefaults instantiates a new CapabilityTemplateCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityTemplateCatalogWithDefaults() *CapabilityTemplateCatalog {
	this := CapabilityTemplateCatalog{}
	var classId string = "capability.TemplateCatalog"
	this.ClassId = classId
	var objectType string = "capability.TemplateCatalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityTemplateCatalog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityTemplateCatalog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityTemplateCatalog) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.TemplateCatalog" of the ClassId field.
func (o *CapabilityTemplateCatalog) GetDefaultClassId() interface{} {
	return "capability.TemplateCatalog"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityTemplateCatalog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityTemplateCatalog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityTemplateCatalog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.TemplateCatalog" of the ObjectType field.
func (o *CapabilityTemplateCatalog) GetDefaultObjectType() interface{} {
	return "capability.TemplateCatalog"
}

// GetAllowedOverrideList returns the AllowedOverrideList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityTemplateCatalog) GetAllowedOverrideList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedOverrideList
}

// GetAllowedOverrideListOk returns a tuple with the AllowedOverrideList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityTemplateCatalog) GetAllowedOverrideListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedOverrideList) {
		return nil, false
	}
	return o.AllowedOverrideList, true
}

// HasAllowedOverrideList returns a boolean if a field has been set.
func (o *CapabilityTemplateCatalog) HasAllowedOverrideList() bool {
	if o != nil && !IsNil(o.AllowedOverrideList) {
		return true
	}

	return false
}

// SetAllowedOverrideList gets a reference to the given []string and assigns it to the AllowedOverrideList field.
func (o *CapabilityTemplateCatalog) SetAllowedOverrideList(v []string) {
	o.AllowedOverrideList = v
}

func (o CapabilityTemplateCatalog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityTemplateCatalog) ToMap() (map[string]interface{}, error) {
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
	if o.AllowedOverrideList != nil {
		toSerialize["AllowedOverrideList"] = o.AllowedOverrideList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityTemplateCatalog) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityTemplateCatalogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string   `json:"ObjectType"`
		AllowedOverrideList []string `json:"AllowedOverrideList,omitempty"`
	}

	varCapabilityTemplateCatalogWithoutEmbeddedStruct := CapabilityTemplateCatalogWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityTemplateCatalogWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityTemplateCatalog := _CapabilityTemplateCatalog{}
		varCapabilityTemplateCatalog.ClassId = varCapabilityTemplateCatalogWithoutEmbeddedStruct.ClassId
		varCapabilityTemplateCatalog.ObjectType = varCapabilityTemplateCatalogWithoutEmbeddedStruct.ObjectType
		varCapabilityTemplateCatalog.AllowedOverrideList = varCapabilityTemplateCatalogWithoutEmbeddedStruct.AllowedOverrideList
		*o = CapabilityTemplateCatalog(varCapabilityTemplateCatalog)
	} else {
		return err
	}

	varCapabilityTemplateCatalog := _CapabilityTemplateCatalog{}

	err = json.Unmarshal(data, &varCapabilityTemplateCatalog)
	if err == nil {
		o.CapabilityCapability = varCapabilityTemplateCatalog.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowedOverrideList")

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

type NullableCapabilityTemplateCatalog struct {
	value *CapabilityTemplateCatalog
	isSet bool
}

func (v NullableCapabilityTemplateCatalog) Get() *CapabilityTemplateCatalog {
	return v.value
}

func (v *NullableCapabilityTemplateCatalog) Set(val *CapabilityTemplateCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityTemplateCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityTemplateCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityTemplateCatalog(val *CapabilityTemplateCatalog) *NullableCapabilityTemplateCatalog {
	return &NullableCapabilityTemplateCatalog{value: val, isSet: true}
}

func (v NullableCapabilityTemplateCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityTemplateCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
