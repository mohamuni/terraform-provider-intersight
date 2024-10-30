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

// checks if the StorageBaseProtectionGroupSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageBaseProtectionGroupSnapshot{}

// StorageBaseProtectionGroupSnapshot Protection group snapshot entity in protection group.
type StorageBaseProtectionGroupSnapshot struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Protection group snapshot creation time.
	CreatedTime *time.Time `json:"CreatedTime,omitempty"`
	// Protection group snapshot name which represents point-in-time copy of all members in protection group.
	Name *string `json:"Name,omitempty"`
	// Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set.
	Size *int64 `json:"Size,omitempty"`
	// Source protection group name on which the snapshot is created.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseProtectionGroupSnapshot StorageBaseProtectionGroupSnapshot

// NewStorageBaseProtectionGroupSnapshot instantiates a new StorageBaseProtectionGroupSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseProtectionGroupSnapshot(classId string, objectType string) *StorageBaseProtectionGroupSnapshot {
	this := StorageBaseProtectionGroupSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseProtectionGroupSnapshotWithDefaults instantiates a new StorageBaseProtectionGroupSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseProtectionGroupSnapshotWithDefaults() *StorageBaseProtectionGroupSnapshot {
	this := StorageBaseProtectionGroupSnapshot{}
	var classId string = "storage.PureProtectionGroupSnapshot"
	this.ClassId = classId
	var objectType string = "storage.PureProtectionGroupSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseProtectionGroupSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseProtectionGroupSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureProtectionGroupSnapshot" of the ClassId field.
func (o *StorageBaseProtectionGroupSnapshot) GetDefaultClassId() interface{} {
	return "storage.PureProtectionGroupSnapshot"
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseProtectionGroupSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseProtectionGroupSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureProtectionGroupSnapshot" of the ObjectType field.
func (o *StorageBaseProtectionGroupSnapshot) GetDefaultObjectType() interface{} {
	return "storage.PureProtectionGroupSnapshot"
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshot) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshot) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshot) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *StorageBaseProtectionGroupSnapshot) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshot) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshot) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshot) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseProtectionGroupSnapshot) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshot) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshot) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshot) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseProtectionGroupSnapshot) SetSize(v int64) {
	o.Size = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshot) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshot) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshot) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StorageBaseProtectionGroupSnapshot) SetSource(v string) {
	o.Source = &v
}

func (o StorageBaseProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageBaseProtectionGroupSnapshot) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreatedTime) {
		toSerialize["CreatedTime"] = o.CreatedTime
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.Source) {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageBaseProtectionGroupSnapshot) UnmarshalJSON(data []byte) (err error) {
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
	type StorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Protection group snapshot creation time.
		CreatedTime *time.Time `json:"CreatedTime,omitempty"`
		// Protection group snapshot name which represents point-in-time copy of all members in protection group.
		Name *string `json:"Name,omitempty"`
		// Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set.
		Size *int64 `json:"Size,omitempty"`
		// Source protection group name on which the snapshot is created.
		Source *string `json:"Source,omitempty"`
	}

	varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct := StorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseProtectionGroupSnapshot := _StorageBaseProtectionGroupSnapshot{}
		varStorageBaseProtectionGroupSnapshot.ClassId = varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct.ClassId
		varStorageBaseProtectionGroupSnapshot.ObjectType = varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct.ObjectType
		varStorageBaseProtectionGroupSnapshot.CreatedTime = varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct.CreatedTime
		varStorageBaseProtectionGroupSnapshot.Name = varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct.Name
		varStorageBaseProtectionGroupSnapshot.Size = varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct.Size
		varStorageBaseProtectionGroupSnapshot.Source = varStorageBaseProtectionGroupSnapshotWithoutEmbeddedStruct.Source
		*o = StorageBaseProtectionGroupSnapshot(varStorageBaseProtectionGroupSnapshot)
	} else {
		return err
	}

	varStorageBaseProtectionGroupSnapshot := _StorageBaseProtectionGroupSnapshot{}

	err = json.Unmarshal(data, &varStorageBaseProtectionGroupSnapshot)
	if err == nil {
		o.MoBaseMo = varStorageBaseProtectionGroupSnapshot.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Source")

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

type NullableStorageBaseProtectionGroupSnapshot struct {
	value *StorageBaseProtectionGroupSnapshot
	isSet bool
}

func (v NullableStorageBaseProtectionGroupSnapshot) Get() *StorageBaseProtectionGroupSnapshot {
	return v.value
}

func (v *NullableStorageBaseProtectionGroupSnapshot) Set(val *StorageBaseProtectionGroupSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseProtectionGroupSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseProtectionGroupSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseProtectionGroupSnapshot(val *StorageBaseProtectionGroupSnapshot) *NullableStorageBaseProtectionGroupSnapshot {
	return &NullableStorageBaseProtectionGroupSnapshot{value: val, isSet: true}
}

func (v NullableStorageBaseProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseProtectionGroupSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
