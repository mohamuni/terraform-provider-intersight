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

// checks if the StorageNetAppEthernetPortEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppEthernetPortEvent{}

// StorageNetAppEthernetPortEvent An event where the impacted resource type is an ethernet port.
type StorageNetAppEthernetPortEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                        `json:"ObjectType"`
	EthernetPort         NullableStorageNetAppEthernetPortRelationship `json:"EthernetPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPortEvent StorageNetAppEthernetPortEvent

// NewStorageNetAppEthernetPortEvent instantiates a new StorageNetAppEthernetPortEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPortEvent(classId string, objectType string) *StorageNetAppEthernetPortEvent {
	this := StorageNetAppEthernetPortEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppEthernetPortEventWithDefaults instantiates a new StorageNetAppEthernetPortEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortEventWithDefaults() *StorageNetAppEthernetPortEvent {
	this := StorageNetAppEthernetPortEvent{}
	var classId string = "storage.NetAppEthernetPortEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPortEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPortEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPortEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppEthernetPortEvent" of the ClassId field.
func (o *StorageNetAppEthernetPortEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppEthernetPortEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPortEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPortEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppEthernetPortEvent" of the ObjectType field.
func (o *StorageNetAppEthernetPortEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppEthernetPortEvent"
}

// GetEthernetPort returns the EthernetPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPortEvent) GetEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || IsNil(o.EthernetPort.Get()) {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.EthernetPort.Get()
}

// GetEthernetPortOk returns a tuple with the EthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPortEvent) GetEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EthernetPort.Get(), o.EthernetPort.IsSet()
}

// HasEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortEvent) HasEthernetPort() bool {
	if o != nil && o.EthernetPort.IsSet() {
		return true
	}

	return false
}

// SetEthernetPort gets a reference to the given NullableStorageNetAppEthernetPortRelationship and assigns it to the EthernetPort field.
func (o *StorageNetAppEthernetPortEvent) SetEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.EthernetPort.Set(&v)
}

// SetEthernetPortNil sets the value for EthernetPort to be an explicit nil
func (o *StorageNetAppEthernetPortEvent) SetEthernetPortNil() {
	o.EthernetPort.Set(nil)
}

// UnsetEthernetPort ensures that no value is present for EthernetPort, not even an explicit nil
func (o *StorageNetAppEthernetPortEvent) UnsetEthernetPort() {
	o.EthernetPort.Unset()
}

func (o StorageNetAppEthernetPortEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppEthernetPortEvent) ToMap() (map[string]interface{}, error) {
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
	if o.EthernetPort.IsSet() {
		toSerialize["EthernetPort"] = o.EthernetPort.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppEthernetPortEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppEthernetPortEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                        `json:"ObjectType"`
		EthernetPort NullableStorageNetAppEthernetPortRelationship `json:"EthernetPort,omitempty"`
	}

	varStorageNetAppEthernetPortEventWithoutEmbeddedStruct := StorageNetAppEthernetPortEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppEthernetPortEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppEthernetPortEvent := _StorageNetAppEthernetPortEvent{}
		varStorageNetAppEthernetPortEvent.ClassId = varStorageNetAppEthernetPortEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppEthernetPortEvent.ObjectType = varStorageNetAppEthernetPortEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppEthernetPortEvent.EthernetPort = varStorageNetAppEthernetPortEventWithoutEmbeddedStruct.EthernetPort
		*o = StorageNetAppEthernetPortEvent(varStorageNetAppEthernetPortEvent)
	} else {
		return err
	}

	varStorageNetAppEthernetPortEvent := _StorageNetAppEthernetPortEvent{}

	err = json.Unmarshal(data, &varStorageNetAppEthernetPortEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppEthernetPortEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EthernetPort")

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

type NullableStorageNetAppEthernetPortEvent struct {
	value *StorageNetAppEthernetPortEvent
	isSet bool
}

func (v NullableStorageNetAppEthernetPortEvent) Get() *StorageNetAppEthernetPortEvent {
	return v.value
}

func (v *NullableStorageNetAppEthernetPortEvent) Set(val *StorageNetAppEthernetPortEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPortEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPortEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPortEvent(val *StorageNetAppEthernetPortEvent) *NullableStorageNetAppEthernetPortEvent {
	return &NullableStorageNetAppEthernetPortEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPortEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPortEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
