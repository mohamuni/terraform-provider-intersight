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

// checks if the StorageExternalLun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageExternalLun{}

// StorageExternalLun The following properties are restored for each LU on the external storage system.
type StorageExternalLun struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// LUN within the ports of the external storage system.
	ExternalLun *int64 `json:"ExternalLun,omitempty"`
	// WWN of the external storage system.
	ExternalWwn *string `json:"ExternalWwn,omitempty"`
	// Status of the external path.
	PathStatus *string `json:"PathStatus,omitempty"`
	// Port number of the local storage system.
	PortId *string `json:"PortId,omitempty"`
	// Priority within the external path group.
	Priority             *int64 `json:"Priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageExternalLun StorageExternalLun

// NewStorageExternalLun instantiates a new StorageExternalLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageExternalLun(classId string, objectType string) *StorageExternalLun {
	this := StorageExternalLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageExternalLunWithDefaults instantiates a new StorageExternalLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageExternalLunWithDefaults() *StorageExternalLun {
	this := StorageExternalLun{}
	var classId string = "storage.ExternalLun"
	this.ClassId = classId
	var objectType string = "storage.ExternalLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageExternalLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageExternalLun) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.ExternalLun" of the ClassId field.
func (o *StorageExternalLun) GetDefaultClassId() interface{} {
	return "storage.ExternalLun"
}

// GetObjectType returns the ObjectType field value
func (o *StorageExternalLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageExternalLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.ExternalLun" of the ObjectType field.
func (o *StorageExternalLun) GetDefaultObjectType() interface{} {
	return "storage.ExternalLun"
}

// GetExternalLun returns the ExternalLun field value if set, zero value otherwise.
func (o *StorageExternalLun) GetExternalLun() int64 {
	if o == nil || IsNil(o.ExternalLun) {
		var ret int64
		return ret
	}
	return *o.ExternalLun
}

// GetExternalLunOk returns a tuple with the ExternalLun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetExternalLunOk() (*int64, bool) {
	if o == nil || IsNil(o.ExternalLun) {
		return nil, false
	}
	return o.ExternalLun, true
}

// HasExternalLun returns a boolean if a field has been set.
func (o *StorageExternalLun) HasExternalLun() bool {
	if o != nil && !IsNil(o.ExternalLun) {
		return true
	}

	return false
}

// SetExternalLun gets a reference to the given int64 and assigns it to the ExternalLun field.
func (o *StorageExternalLun) SetExternalLun(v int64) {
	o.ExternalLun = &v
}

// GetExternalWwn returns the ExternalWwn field value if set, zero value otherwise.
func (o *StorageExternalLun) GetExternalWwn() string {
	if o == nil || IsNil(o.ExternalWwn) {
		var ret string
		return ret
	}
	return *o.ExternalWwn
}

// GetExternalWwnOk returns a tuple with the ExternalWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetExternalWwnOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalWwn) {
		return nil, false
	}
	return o.ExternalWwn, true
}

// HasExternalWwn returns a boolean if a field has been set.
func (o *StorageExternalLun) HasExternalWwn() bool {
	if o != nil && !IsNil(o.ExternalWwn) {
		return true
	}

	return false
}

// SetExternalWwn gets a reference to the given string and assigns it to the ExternalWwn field.
func (o *StorageExternalLun) SetExternalWwn(v string) {
	o.ExternalWwn = &v
}

// GetPathStatus returns the PathStatus field value if set, zero value otherwise.
func (o *StorageExternalLun) GetPathStatus() string {
	if o == nil || IsNil(o.PathStatus) {
		var ret string
		return ret
	}
	return *o.PathStatus
}

// GetPathStatusOk returns a tuple with the PathStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetPathStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PathStatus) {
		return nil, false
	}
	return o.PathStatus, true
}

// HasPathStatus returns a boolean if a field has been set.
func (o *StorageExternalLun) HasPathStatus() bool {
	if o != nil && !IsNil(o.PathStatus) {
		return true
	}

	return false
}

// SetPathStatus gets a reference to the given string and assigns it to the PathStatus field.
func (o *StorageExternalLun) SetPathStatus(v string) {
	o.PathStatus = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageExternalLun) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageExternalLun) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageExternalLun) SetPortId(v string) {
	o.PortId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *StorageExternalLun) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLun) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *StorageExternalLun) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *StorageExternalLun) SetPriority(v int64) {
	o.Priority = &v
}

func (o StorageExternalLun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageExternalLun) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExternalLun) {
		toSerialize["ExternalLun"] = o.ExternalLun
	}
	if !IsNil(o.ExternalWwn) {
		toSerialize["ExternalWwn"] = o.ExternalWwn
	}
	if !IsNil(o.PathStatus) {
		toSerialize["PathStatus"] = o.PathStatus
	}
	if !IsNil(o.PortId) {
		toSerialize["PortId"] = o.PortId
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageExternalLun) UnmarshalJSON(data []byte) (err error) {
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
	type StorageExternalLunWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// LUN within the ports of the external storage system.
		ExternalLun *int64 `json:"ExternalLun,omitempty"`
		// WWN of the external storage system.
		ExternalWwn *string `json:"ExternalWwn,omitempty"`
		// Status of the external path.
		PathStatus *string `json:"PathStatus,omitempty"`
		// Port number of the local storage system.
		PortId *string `json:"PortId,omitempty"`
		// Priority within the external path group.
		Priority *int64 `json:"Priority,omitempty"`
	}

	varStorageExternalLunWithoutEmbeddedStruct := StorageExternalLunWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageExternalLunWithoutEmbeddedStruct)
	if err == nil {
		varStorageExternalLun := _StorageExternalLun{}
		varStorageExternalLun.ClassId = varStorageExternalLunWithoutEmbeddedStruct.ClassId
		varStorageExternalLun.ObjectType = varStorageExternalLunWithoutEmbeddedStruct.ObjectType
		varStorageExternalLun.ExternalLun = varStorageExternalLunWithoutEmbeddedStruct.ExternalLun
		varStorageExternalLun.ExternalWwn = varStorageExternalLunWithoutEmbeddedStruct.ExternalWwn
		varStorageExternalLun.PathStatus = varStorageExternalLunWithoutEmbeddedStruct.PathStatus
		varStorageExternalLun.PortId = varStorageExternalLunWithoutEmbeddedStruct.PortId
		varStorageExternalLun.Priority = varStorageExternalLunWithoutEmbeddedStruct.Priority
		*o = StorageExternalLun(varStorageExternalLun)
	} else {
		return err
	}

	varStorageExternalLun := _StorageExternalLun{}

	err = json.Unmarshal(data, &varStorageExternalLun)
	if err == nil {
		o.MoBaseComplexType = varStorageExternalLun.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalLun")
		delete(additionalProperties, "ExternalWwn")
		delete(additionalProperties, "PathStatus")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "Priority")

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

type NullableStorageExternalLun struct {
	value *StorageExternalLun
	isSet bool
}

func (v NullableStorageExternalLun) Get() *StorageExternalLun {
	return v.value
}

func (v *NullableStorageExternalLun) Set(val *StorageExternalLun) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageExternalLun) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageExternalLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageExternalLun(val *StorageExternalLun) *NullableStorageExternalLun {
	return &NullableStorageExternalLun{value: val, isSet: true}
}

func (v NullableStorageExternalLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageExternalLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
