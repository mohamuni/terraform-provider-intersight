/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the StorageNetAppFcPortEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppFcPortEvent{}

// StorageNetAppFcPortEvent An event where the impacted resource type is a FC port.
type StorageNetAppFcPortEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                  `json:"ObjectType"`
	FcPort               NullableStorageNetAppFcPortRelationship `json:"FcPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppFcPortEvent StorageNetAppFcPortEvent

// NewStorageNetAppFcPortEvent instantiates a new StorageNetAppFcPortEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppFcPortEvent(classId string, objectType string) *StorageNetAppFcPortEvent {
	this := StorageNetAppFcPortEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppFcPortEventWithDefaults instantiates a new StorageNetAppFcPortEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppFcPortEventWithDefaults() *StorageNetAppFcPortEvent {
	this := StorageNetAppFcPortEvent{}
	var classId string = "storage.NetAppFcPortEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppFcPortEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppFcPortEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppFcPortEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppFcPortEvent" of the ClassId field.
func (o *StorageNetAppFcPortEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppFcPortEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppFcPortEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppFcPortEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppFcPortEvent" of the ObjectType field.
func (o *StorageNetAppFcPortEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppFcPortEvent"
}

// GetFcPort returns the FcPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcPortEvent) GetFcPort() StorageNetAppFcPortRelationship {
	if o == nil || IsNil(o.FcPort.Get()) {
		var ret StorageNetAppFcPortRelationship
		return ret
	}
	return *o.FcPort.Get()
}

// GetFcPortOk returns a tuple with the FcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcPortEvent) GetFcPortOk() (*StorageNetAppFcPortRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.FcPort.Get(), o.FcPort.IsSet()
}

// HasFcPort returns a boolean if a field has been set.
func (o *StorageNetAppFcPortEvent) HasFcPort() bool {
	if o != nil && o.FcPort.IsSet() {
		return true
	}

	return false
}

// SetFcPort gets a reference to the given NullableStorageNetAppFcPortRelationship and assigns it to the FcPort field.
func (o *StorageNetAppFcPortEvent) SetFcPort(v StorageNetAppFcPortRelationship) {
	o.FcPort.Set(&v)
}

// SetFcPortNil sets the value for FcPort to be an explicit nil
func (o *StorageNetAppFcPortEvent) SetFcPortNil() {
	o.FcPort.Set(nil)
}

// UnsetFcPort ensures that no value is present for FcPort, not even an explicit nil
func (o *StorageNetAppFcPortEvent) UnsetFcPort() {
	o.FcPort.Unset()
}

func (o StorageNetAppFcPortEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppFcPortEvent) ToMap() (map[string]interface{}, error) {
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
	if o.FcPort.IsSet() {
		toSerialize["FcPort"] = o.FcPort.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppFcPortEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppFcPortEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                  `json:"ObjectType"`
		FcPort     NullableStorageNetAppFcPortRelationship `json:"FcPort,omitempty"`
	}

	varStorageNetAppFcPortEventWithoutEmbeddedStruct := StorageNetAppFcPortEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppFcPortEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppFcPortEvent := _StorageNetAppFcPortEvent{}
		varStorageNetAppFcPortEvent.ClassId = varStorageNetAppFcPortEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppFcPortEvent.ObjectType = varStorageNetAppFcPortEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppFcPortEvent.FcPort = varStorageNetAppFcPortEventWithoutEmbeddedStruct.FcPort
		*o = StorageNetAppFcPortEvent(varStorageNetAppFcPortEvent)
	} else {
		return err
	}

	varStorageNetAppFcPortEvent := _StorageNetAppFcPortEvent{}

	err = json.Unmarshal(data, &varStorageNetAppFcPortEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppFcPortEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FcPort")

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

type NullableStorageNetAppFcPortEvent struct {
	value *StorageNetAppFcPortEvent
	isSet bool
}

func (v NullableStorageNetAppFcPortEvent) Get() *StorageNetAppFcPortEvent {
	return v.value
}

func (v *NullableStorageNetAppFcPortEvent) Set(val *StorageNetAppFcPortEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppFcPortEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppFcPortEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppFcPortEvent(val *StorageNetAppFcPortEvent) *NullableStorageNetAppFcPortEvent {
	return &NullableStorageNetAppFcPortEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppFcPortEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppFcPortEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
