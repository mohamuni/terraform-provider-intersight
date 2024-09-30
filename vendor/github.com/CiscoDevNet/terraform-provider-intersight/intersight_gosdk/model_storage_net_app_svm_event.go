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

// checks if the StorageNetAppSvmEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppSvmEvent{}

// StorageNetAppSvmEvent An event where the impacted resource type is a storage vm.
type StorageNetAppSvmEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                     `json:"ObjectType"`
	Tenant               NullableStorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppSvmEvent StorageNetAppSvmEvent

// NewStorageNetAppSvmEvent instantiates a new StorageNetAppSvmEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppSvmEvent(classId string, objectType string) *StorageNetAppSvmEvent {
	this := StorageNetAppSvmEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppSvmEventWithDefaults instantiates a new StorageNetAppSvmEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppSvmEventWithDefaults() *StorageNetAppSvmEvent {
	this := StorageNetAppSvmEvent{}
	var classId string = "storage.NetAppSvmEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppSvmEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppSvmEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppSvmEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppSvmEvent" of the ClassId field.
func (o *StorageNetAppSvmEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppSvmEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppSvmEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppSvmEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppSvmEvent" of the ObjectType field.
func (o *StorageNetAppSvmEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppSvmEvent"
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppSvmEvent) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppSvmEvent) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppSvmEvent) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableStorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppSvmEvent) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *StorageNetAppSvmEvent) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *StorageNetAppSvmEvent) UnsetTenant() {
	o.Tenant.Unset()
}

func (o StorageNetAppSvmEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppSvmEvent) ToMap() (map[string]interface{}, error) {
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
	if o.Tenant.IsSet() {
		toSerialize["Tenant"] = o.Tenant.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppSvmEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppSvmEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                     `json:"ObjectType"`
		Tenant     NullableStorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppSvmEventWithoutEmbeddedStruct := StorageNetAppSvmEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppSvmEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppSvmEvent := _StorageNetAppSvmEvent{}
		varStorageNetAppSvmEvent.ClassId = varStorageNetAppSvmEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppSvmEvent.ObjectType = varStorageNetAppSvmEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppSvmEvent.Tenant = varStorageNetAppSvmEventWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppSvmEvent(varStorageNetAppSvmEvent)
	} else {
		return err
	}

	varStorageNetAppSvmEvent := _StorageNetAppSvmEvent{}

	err = json.Unmarshal(data, &varStorageNetAppSvmEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppSvmEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Tenant")

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

type NullableStorageNetAppSvmEvent struct {
	value *StorageNetAppSvmEvent
	isSet bool
}

func (v NullableStorageNetAppSvmEvent) Get() *StorageNetAppSvmEvent {
	return v.value
}

func (v *NullableStorageNetAppSvmEvent) Set(val *StorageNetAppSvmEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppSvmEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppSvmEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppSvmEvent(val *StorageNetAppSvmEvent) *NullableStorageNetAppSvmEvent {
	return &NullableStorageNetAppSvmEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppSvmEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppSvmEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
