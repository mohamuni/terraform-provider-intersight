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

// checks if the StorageNetAppClusterEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppClusterEvent{}

// StorageNetAppClusterEvent An event where the impacted resource type is a cluster.
type StorageNetAppClusterEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                   `json:"ObjectType"`
	Array                NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppClusterEvent StorageNetAppClusterEvent

// NewStorageNetAppClusterEvent instantiates a new StorageNetAppClusterEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppClusterEvent(classId string, objectType string) *StorageNetAppClusterEvent {
	this := StorageNetAppClusterEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterEventWithDefaults instantiates a new StorageNetAppClusterEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterEventWithDefaults() *StorageNetAppClusterEvent {
	this := StorageNetAppClusterEvent{}
	var classId string = "storage.NetAppClusterEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppClusterEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppClusterEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppClusterEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppClusterEvent" of the ClassId field.
func (o *StorageNetAppClusterEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppClusterEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppClusterEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppClusterEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppClusterEvent" of the ObjectType field.
func (o *StorageNetAppClusterEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppClusterEvent"
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterEvent) GetArray() StorageNetAppClusterRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterEvent) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppClusterEvent) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppClusterEvent) SetArray(v StorageNetAppClusterRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageNetAppClusterEvent) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageNetAppClusterEvent) UnsetArray() {
	o.Array.Unset()
}

func (o StorageNetAppClusterEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppClusterEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseEvent, errStorageNetAppBaseEvent := json.Marshal(o.StorageNetAppBaseEvent)
	if errStorageNetAppBaseEvent != nil {
		return map[string]interface{}{}, errStorageNetAppBaseEvent
	}
	errStorageNetAppBaseEvent = json.Unmarshal([]byte(serializedStorageNetAppBaseEvent), &toSerialize)
	if errStorageNetAppBaseEvent != nil {
		return map[string]interface{}{}, errStorageNetAppBaseEvent
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppClusterEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppClusterEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                   `json:"ObjectType"`
		Array      NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppClusterEventWithoutEmbeddedStruct := StorageNetAppClusterEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppClusterEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppClusterEvent := _StorageNetAppClusterEvent{}
		varStorageNetAppClusterEvent.ClassId = varStorageNetAppClusterEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppClusterEvent.ObjectType = varStorageNetAppClusterEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppClusterEvent.Array = varStorageNetAppClusterEventWithoutEmbeddedStruct.Array
		*o = StorageNetAppClusterEvent(varStorageNetAppClusterEvent)
	} else {
		return err
	}

	varStorageNetAppClusterEvent := _StorageNetAppClusterEvent{}

	err = json.Unmarshal(data, &varStorageNetAppClusterEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppClusterEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")

		// remove fields from embedded structs
		reflectStorageNetAppBaseEvent := reflect.ValueOf(o.StorageNetAppBaseEvent)
		for i := 0; i < reflectStorageNetAppBaseEvent.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseEvent.Type().Field(i)

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

type NullableStorageNetAppClusterEvent struct {
	value *StorageNetAppClusterEvent
	isSet bool
}

func (v NullableStorageNetAppClusterEvent) Get() *StorageNetAppClusterEvent {
	return v.value
}

func (v *NullableStorageNetAppClusterEvent) Set(val *StorageNetAppClusterEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppClusterEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppClusterEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppClusterEvent(val *StorageNetAppClusterEvent) *NullableStorageNetAppClusterEvent {
	return &NullableStorageNetAppClusterEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppClusterEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppClusterEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
