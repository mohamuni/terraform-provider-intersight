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

// checks if the DnacTransit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacTransit{}

// DnacTransit Details for the transits.
type DnacTransit struct {
	DnacInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Identification for the Transit.
	TransitId *string `json:"TransitId,omitempty"`
	// Name identifier for the Transit.
	TransitName *string `json:"TransitName,omitempty"`
	// Transit type for the transit.
	TransitType          *string `json:"TransitType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacTransit DnacTransit

// NewDnacTransit instantiates a new DnacTransit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacTransit(classId string, objectType string) *DnacTransit {
	this := DnacTransit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewDnacTransitWithDefaults instantiates a new DnacTransit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacTransitWithDefaults() *DnacTransit {
	this := DnacTransit{}
	var classId string = "dnac.Transit"
	this.ClassId = classId
	var objectType string = "dnac.Transit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *DnacTransit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DnacTransit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DnacTransit) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "dnac.Transit" of the ClassId field.
func (o *DnacTransit) GetDefaultClassId() interface{} {
	return "dnac.Transit"
}

// GetObjectType returns the ObjectType field value
func (o *DnacTransit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DnacTransit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DnacTransit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "dnac.Transit" of the ObjectType field.
func (o *DnacTransit) GetDefaultObjectType() interface{} {
	return "dnac.Transit"
}

// GetTransitId returns the TransitId field value if set, zero value otherwise.
func (o *DnacTransit) GetTransitId() string {
	if o == nil || IsNil(o.TransitId) {
		var ret string
		return ret
	}
	return *o.TransitId
}

// GetTransitIdOk returns a tuple with the TransitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTransit) GetTransitIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransitId) {
		return nil, false
	}
	return o.TransitId, true
}

// HasTransitId returns a boolean if a field has been set.
func (o *DnacTransit) HasTransitId() bool {
	if o != nil && !IsNil(o.TransitId) {
		return true
	}

	return false
}

// SetTransitId gets a reference to the given string and assigns it to the TransitId field.
func (o *DnacTransit) SetTransitId(v string) {
	o.TransitId = &v
}

// GetTransitName returns the TransitName field value if set, zero value otherwise.
func (o *DnacTransit) GetTransitName() string {
	if o == nil || IsNil(o.TransitName) {
		var ret string
		return ret
	}
	return *o.TransitName
}

// GetTransitNameOk returns a tuple with the TransitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTransit) GetTransitNameOk() (*string, bool) {
	if o == nil || IsNil(o.TransitName) {
		return nil, false
	}
	return o.TransitName, true
}

// HasTransitName returns a boolean if a field has been set.
func (o *DnacTransit) HasTransitName() bool {
	if o != nil && !IsNil(o.TransitName) {
		return true
	}

	return false
}

// SetTransitName gets a reference to the given string and assigns it to the TransitName field.
func (o *DnacTransit) SetTransitName(v string) {
	o.TransitName = &v
}

// GetTransitType returns the TransitType field value if set, zero value otherwise.
func (o *DnacTransit) GetTransitType() string {
	if o == nil || IsNil(o.TransitType) {
		var ret string
		return ret
	}
	return *o.TransitType
}

// GetTransitTypeOk returns a tuple with the TransitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTransit) GetTransitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransitType) {
		return nil, false
	}
	return o.TransitType, true
}

// HasTransitType returns a boolean if a field has been set.
func (o *DnacTransit) HasTransitType() bool {
	if o != nil && !IsNil(o.TransitType) {
		return true
	}

	return false
}

// SetTransitType gets a reference to the given string and assigns it to the TransitType field.
func (o *DnacTransit) SetTransitType(v string) {
	o.TransitType = &v
}

func (o DnacTransit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacTransit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDnacInventoryEntity, errDnacInventoryEntity := json.Marshal(o.DnacInventoryEntity)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	errDnacInventoryEntity = json.Unmarshal([]byte(serializedDnacInventoryEntity), &toSerialize)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.TransitId) {
		toSerialize["TransitId"] = o.TransitId
	}
	if !IsNil(o.TransitName) {
		toSerialize["TransitName"] = o.TransitName
	}
	if !IsNil(o.TransitType) {
		toSerialize["TransitType"] = o.TransitType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacTransit) UnmarshalJSON(data []byte) (err error) {
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
	type DnacTransitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Identification for the Transit.
		TransitId *string `json:"TransitId,omitempty"`
		// Name identifier for the Transit.
		TransitName *string `json:"TransitName,omitempty"`
		// Transit type for the transit.
		TransitType *string `json:"TransitType,omitempty"`
	}

	varDnacTransitWithoutEmbeddedStruct := DnacTransitWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacTransitWithoutEmbeddedStruct)
	if err == nil {
		varDnacTransit := _DnacTransit{}
		varDnacTransit.ClassId = varDnacTransitWithoutEmbeddedStruct.ClassId
		varDnacTransit.ObjectType = varDnacTransitWithoutEmbeddedStruct.ObjectType
		varDnacTransit.TransitId = varDnacTransitWithoutEmbeddedStruct.TransitId
		varDnacTransit.TransitName = varDnacTransitWithoutEmbeddedStruct.TransitName
		varDnacTransit.TransitType = varDnacTransitWithoutEmbeddedStruct.TransitType
		*o = DnacTransit(varDnacTransit)
	} else {
		return err
	}

	varDnacTransit := _DnacTransit{}

	err = json.Unmarshal(data, &varDnacTransit)
	if err == nil {
		o.DnacInventoryEntity = varDnacTransit.DnacInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TransitId")
		delete(additionalProperties, "TransitName")
		delete(additionalProperties, "TransitType")

		// remove fields from embedded structs
		reflectDnacInventoryEntity := reflect.ValueOf(o.DnacInventoryEntity)
		for i := 0; i < reflectDnacInventoryEntity.Type().NumField(); i++ {
			t := reflectDnacInventoryEntity.Type().Field(i)

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

type NullableDnacTransit struct {
	value *DnacTransit
	isSet bool
}

func (v NullableDnacTransit) Get() *DnacTransit {
	return v.value
}

func (v *NullableDnacTransit) Set(val *DnacTransit) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacTransit) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacTransit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacTransit(val *DnacTransit) *NullableDnacTransit {
	return &NullableDnacTransit{value: val, isSet: true}
}

func (v NullableDnacTransit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacTransit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
