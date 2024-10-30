/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the NiatelemetryEqptcapacityPrefixEntries5min type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryEqptcapacityPrefixEntries5min{}

// NiatelemetryEqptcapacityPrefixEntries5min Aci node lpm utilization information in last 15 mintutes.
type NiatelemetryEqptcapacityPrefixEntries5min struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Aci node lpm utilized last.
	ExtNormalizedLast    *string `json:"ExtNormalizedLast,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryEqptcapacityPrefixEntries5min NiatelemetryEqptcapacityPrefixEntries5min

// NewNiatelemetryEqptcapacityPrefixEntries5min instantiates a new NiatelemetryEqptcapacityPrefixEntries5min object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryEqptcapacityPrefixEntries5min(classId string, objectType string) *NiatelemetryEqptcapacityPrefixEntries5min {
	this := NiatelemetryEqptcapacityPrefixEntries5min{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryEqptcapacityPrefixEntries5minWithDefaults instantiates a new NiatelemetryEqptcapacityPrefixEntries5min object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryEqptcapacityPrefixEntries5minWithDefaults() *NiatelemetryEqptcapacityPrefixEntries5min {
	this := NiatelemetryEqptcapacityPrefixEntries5min{}
	var classId string = "niatelemetry.EqptcapacityPrefixEntries5min"
	this.ClassId = classId
	var objectType string = "niatelemetry.EqptcapacityPrefixEntries5min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryEqptcapacityPrefixEntries5min) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.EqptcapacityPrefixEntries5min" of the ClassId field.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetDefaultClassId() interface{} {
	return "niatelemetry.EqptcapacityPrefixEntries5min"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryEqptcapacityPrefixEntries5min) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.EqptcapacityPrefixEntries5min" of the ObjectType field.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetDefaultObjectType() interface{} {
	return "niatelemetry.EqptcapacityPrefixEntries5min"
}

// GetExtNormalizedLast returns the ExtNormalizedLast field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetExtNormalizedLast() string {
	if o == nil || IsNil(o.ExtNormalizedLast) {
		var ret string
		return ret
	}
	return *o.ExtNormalizedLast
}

// GetExtNormalizedLastOk returns a tuple with the ExtNormalizedLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) GetExtNormalizedLastOk() (*string, bool) {
	if o == nil || IsNil(o.ExtNormalizedLast) {
		return nil, false
	}
	return o.ExtNormalizedLast, true
}

// HasExtNormalizedLast returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) HasExtNormalizedLast() bool {
	if o != nil && !IsNil(o.ExtNormalizedLast) {
		return true
	}

	return false
}

// SetExtNormalizedLast gets a reference to the given string and assigns it to the ExtNormalizedLast field.
func (o *NiatelemetryEqptcapacityPrefixEntries5min) SetExtNormalizedLast(v string) {
	o.ExtNormalizedLast = &v
}

func (o NiatelemetryEqptcapacityPrefixEntries5min) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryEqptcapacityPrefixEntries5min) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExtNormalizedLast) {
		toSerialize["ExtNormalizedLast"] = o.ExtNormalizedLast
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryEqptcapacityPrefixEntries5min) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Aci node lpm utilized last.
		ExtNormalizedLast *string `json:"ExtNormalizedLast,omitempty"`
	}

	varNiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct := NiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryEqptcapacityPrefixEntries5min := _NiatelemetryEqptcapacityPrefixEntries5min{}
		varNiatelemetryEqptcapacityPrefixEntries5min.ClassId = varNiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct.ClassId
		varNiatelemetryEqptcapacityPrefixEntries5min.ObjectType = varNiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct.ObjectType
		varNiatelemetryEqptcapacityPrefixEntries5min.ExtNormalizedLast = varNiatelemetryEqptcapacityPrefixEntries5minWithoutEmbeddedStruct.ExtNormalizedLast
		*o = NiatelemetryEqptcapacityPrefixEntries5min(varNiatelemetryEqptcapacityPrefixEntries5min)
	} else {
		return err
	}

	varNiatelemetryEqptcapacityPrefixEntries5min := _NiatelemetryEqptcapacityPrefixEntries5min{}

	err = json.Unmarshal(data, &varNiatelemetryEqptcapacityPrefixEntries5min)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryEqptcapacityPrefixEntries5min.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExtNormalizedLast")

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

type NullableNiatelemetryEqptcapacityPrefixEntries5min struct {
	value *NiatelemetryEqptcapacityPrefixEntries5min
	isSet bool
}

func (v NullableNiatelemetryEqptcapacityPrefixEntries5min) Get() *NiatelemetryEqptcapacityPrefixEntries5min {
	return v.value
}

func (v *NullableNiatelemetryEqptcapacityPrefixEntries5min) Set(val *NiatelemetryEqptcapacityPrefixEntries5min) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryEqptcapacityPrefixEntries5min) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryEqptcapacityPrefixEntries5min) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryEqptcapacityPrefixEntries5min(val *NiatelemetryEqptcapacityPrefixEntries5min) *NullableNiatelemetryEqptcapacityPrefixEntries5min {
	return &NullableNiatelemetryEqptcapacityPrefixEntries5min{value: val, isSet: true}
}

func (v NullableNiatelemetryEqptcapacityPrefixEntries5min) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryEqptcapacityPrefixEntries5min) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
