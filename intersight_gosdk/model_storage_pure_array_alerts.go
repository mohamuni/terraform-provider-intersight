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
	"time"
)

// checks if the StoragePureArrayAlerts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePureArrayAlerts{}

// StoragePureArrayAlerts Alert Events in PureStorage FlashArray.
type StoragePureArrayAlerts struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reason for the event generation.
	Cause *string `json:"Cause,omitempty"`
	// Component for which the event generation.
	Component *string `json:"Component,omitempty"`
	// Date on which the event was generated on FlashArrays.
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// ID of the alert related to the event.
	Name *string `json:"Name,omitempty"`
	// Type of the severity of the event it could be Critical or Warning.
	Severity             *string                              `json:"Severity,omitempty"`
	Array                NullableStoragePureArrayRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureArrayAlerts StoragePureArrayAlerts

// NewStoragePureArrayAlerts instantiates a new StoragePureArrayAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureArrayAlerts(classId string, objectType string) *StoragePureArrayAlerts {
	this := StoragePureArrayAlerts{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureArrayAlertsWithDefaults instantiates a new StoragePureArrayAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureArrayAlertsWithDefaults() *StoragePureArrayAlerts {
	this := StoragePureArrayAlerts{}
	var classId string = "storage.PureArrayAlerts"
	this.ClassId = classId
	var objectType string = "storage.PureArrayAlerts"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureArrayAlerts) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureArrayAlerts) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureArrayAlerts" of the ClassId field.
func (o *StoragePureArrayAlerts) GetDefaultClassId() interface{} {
	return "storage.PureArrayAlerts"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureArrayAlerts) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureArrayAlerts) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureArrayAlerts" of the ObjectType field.
func (o *StoragePureArrayAlerts) GetDefaultObjectType() interface{} {
	return "storage.PureArrayAlerts"
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *StoragePureArrayAlerts) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *StoragePureArrayAlerts) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *StoragePureArrayAlerts) SetCause(v string) {
	o.Cause = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *StoragePureArrayAlerts) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *StoragePureArrayAlerts) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *StoragePureArrayAlerts) SetComponent(v string) {
	o.Component = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StoragePureArrayAlerts) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StoragePureArrayAlerts) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *StoragePureArrayAlerts) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoragePureArrayAlerts) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoragePureArrayAlerts) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoragePureArrayAlerts) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *StoragePureArrayAlerts) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureArrayAlerts) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *StoragePureArrayAlerts) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *StoragePureArrayAlerts) SetSeverity(v string) {
	o.Severity = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureArrayAlerts) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureArrayAlerts) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureArrayAlerts) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureArrayAlerts) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePureArrayAlerts) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePureArrayAlerts) UnsetArray() {
	o.Array.Unset()
}

func (o StoragePureArrayAlerts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePureArrayAlerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Cause) {
		toSerialize["Cause"] = o.Cause
	}
	if !IsNil(o.Component) {
		toSerialize["Component"] = o.Component
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoragePureArrayAlerts) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePureArrayAlertsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Reason for the event generation.
		Cause *string `json:"Cause,omitempty"`
		// Component for which the event generation.
		Component *string `json:"Component,omitempty"`
		// Date on which the event was generated on FlashArrays.
		CreatedAt *time.Time `json:"CreatedAt,omitempty"`
		// ID of the alert related to the event.
		Name *string `json:"Name,omitempty"`
		// Type of the severity of the event it could be Critical or Warning.
		Severity *string                              `json:"Severity,omitempty"`
		Array    NullableStoragePureArrayRelationship `json:"Array,omitempty"`
	}

	varStoragePureArrayAlertsWithoutEmbeddedStruct := StoragePureArrayAlertsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePureArrayAlertsWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureArrayAlerts := _StoragePureArrayAlerts{}
		varStoragePureArrayAlerts.ClassId = varStoragePureArrayAlertsWithoutEmbeddedStruct.ClassId
		varStoragePureArrayAlerts.ObjectType = varStoragePureArrayAlertsWithoutEmbeddedStruct.ObjectType
		varStoragePureArrayAlerts.Cause = varStoragePureArrayAlertsWithoutEmbeddedStruct.Cause
		varStoragePureArrayAlerts.Component = varStoragePureArrayAlertsWithoutEmbeddedStruct.Component
		varStoragePureArrayAlerts.CreatedAt = varStoragePureArrayAlertsWithoutEmbeddedStruct.CreatedAt
		varStoragePureArrayAlerts.Name = varStoragePureArrayAlertsWithoutEmbeddedStruct.Name
		varStoragePureArrayAlerts.Severity = varStoragePureArrayAlertsWithoutEmbeddedStruct.Severity
		varStoragePureArrayAlerts.Array = varStoragePureArrayAlertsWithoutEmbeddedStruct.Array
		*o = StoragePureArrayAlerts(varStoragePureArrayAlerts)
	} else {
		return err
	}

	varStoragePureArrayAlerts := _StoragePureArrayAlerts{}

	err = json.Unmarshal(data, &varStoragePureArrayAlerts)
	if err == nil {
		o.MoBaseMo = varStoragePureArrayAlerts.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cause")
		delete(additionalProperties, "Component")
		delete(additionalProperties, "CreatedAt")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "Array")

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

type NullableStoragePureArrayAlerts struct {
	value *StoragePureArrayAlerts
	isSet bool
}

func (v NullableStoragePureArrayAlerts) Get() *StoragePureArrayAlerts {
	return v.value
}

func (v *NullableStoragePureArrayAlerts) Set(val *StoragePureArrayAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureArrayAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureArrayAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureArrayAlerts(val *StoragePureArrayAlerts) *NullableStoragePureArrayAlerts {
	return &NullableStoragePureArrayAlerts{value: val, isSet: true}
}

func (v NullableStoragePureArrayAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureArrayAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
