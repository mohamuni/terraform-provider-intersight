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

// checks if the SwIdPoolBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwIdPoolBase{}

// SwIdPoolBase Container representation of the Idpools defined for fabric.
type SwIdPoolBase struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType      string  `json:"ObjectType"`
	GapAvailableIds []int64 `json:"GapAvailableIds,omitempty"`
	// Shows the next available Chassis ID to be allocated.
	NextAvailableId      *int64 `json:"NextAvailableId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwIdPoolBase SwIdPoolBase

// NewSwIdPoolBase instantiates a new SwIdPoolBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwIdPoolBase(classId string, objectType string) *SwIdPoolBase {
	this := SwIdPoolBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSwIdPoolBaseWithDefaults instantiates a new SwIdPoolBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwIdPoolBaseWithDefaults() *SwIdPoolBase {
	this := SwIdPoolBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *SwIdPoolBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SwIdPoolBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SwIdPoolBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SwIdPoolBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SwIdPoolBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SwIdPoolBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGapAvailableIds returns the GapAvailableIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwIdPoolBase) GetGapAvailableIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GapAvailableIds
}

// GetGapAvailableIdsOk returns a tuple with the GapAvailableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwIdPoolBase) GetGapAvailableIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GapAvailableIds) {
		return nil, false
	}
	return o.GapAvailableIds, true
}

// HasGapAvailableIds returns a boolean if a field has been set.
func (o *SwIdPoolBase) HasGapAvailableIds() bool {
	if o != nil && !IsNil(o.GapAvailableIds) {
		return true
	}

	return false
}

// SetGapAvailableIds gets a reference to the given []int64 and assigns it to the GapAvailableIds field.
func (o *SwIdPoolBase) SetGapAvailableIds(v []int64) {
	o.GapAvailableIds = v
}

// GetNextAvailableId returns the NextAvailableId field value if set, zero value otherwise.
func (o *SwIdPoolBase) GetNextAvailableId() int64 {
	if o == nil || IsNil(o.NextAvailableId) {
		var ret int64
		return ret
	}
	return *o.NextAvailableId
}

// GetNextAvailableIdOk returns a tuple with the NextAvailableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwIdPoolBase) GetNextAvailableIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NextAvailableId) {
		return nil, false
	}
	return o.NextAvailableId, true
}

// HasNextAvailableId returns a boolean if a field has been set.
func (o *SwIdPoolBase) HasNextAvailableId() bool {
	if o != nil && !IsNil(o.NextAvailableId) {
		return true
	}

	return false
}

// SetNextAvailableId gets a reference to the given int64 and assigns it to the NextAvailableId field.
func (o *SwIdPoolBase) SetNextAvailableId(v int64) {
	o.NextAvailableId = &v
}

func (o SwIdPoolBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwIdPoolBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.GapAvailableIds != nil {
		toSerialize["GapAvailableIds"] = o.GapAvailableIds
	}
	if !IsNil(o.NextAvailableId) {
		toSerialize["NextAvailableId"] = o.NextAvailableId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwIdPoolBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type SwIdPoolBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType      string  `json:"ObjectType"`
		GapAvailableIds []int64 `json:"GapAvailableIds,omitempty"`
		// Shows the next available Chassis ID to be allocated.
		NextAvailableId *int64 `json:"NextAvailableId,omitempty"`
	}

	varSwIdPoolBaseWithoutEmbeddedStruct := SwIdPoolBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSwIdPoolBaseWithoutEmbeddedStruct)
	if err == nil {
		varSwIdPoolBase := _SwIdPoolBase{}
		varSwIdPoolBase.ClassId = varSwIdPoolBaseWithoutEmbeddedStruct.ClassId
		varSwIdPoolBase.ObjectType = varSwIdPoolBaseWithoutEmbeddedStruct.ObjectType
		varSwIdPoolBase.GapAvailableIds = varSwIdPoolBaseWithoutEmbeddedStruct.GapAvailableIds
		varSwIdPoolBase.NextAvailableId = varSwIdPoolBaseWithoutEmbeddedStruct.NextAvailableId
		*o = SwIdPoolBase(varSwIdPoolBase)
	} else {
		return err
	}

	varSwIdPoolBase := _SwIdPoolBase{}

	err = json.Unmarshal(data, &varSwIdPoolBase)
	if err == nil {
		o.MoBaseMo = varSwIdPoolBase.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GapAvailableIds")
		delete(additionalProperties, "NextAvailableId")

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

type NullableSwIdPoolBase struct {
	value *SwIdPoolBase
	isSet bool
}

func (v NullableSwIdPoolBase) Get() *SwIdPoolBase {
	return v.value
}

func (v *NullableSwIdPoolBase) Set(val *SwIdPoolBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSwIdPoolBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSwIdPoolBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwIdPoolBase(val *SwIdPoolBase) *NullableSwIdPoolBase {
	return &NullableSwIdPoolBase{value: val, isSet: true}
}

func (v NullableSwIdPoolBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwIdPoolBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
