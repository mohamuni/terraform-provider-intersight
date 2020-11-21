/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationStorageCapacity Captures the common attributes of a storage entity.
type VirtualizationStorageCapacity struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The total capacity of the entity (bytes).
	Capacity *int64 `json:"Capacity,omitempty"`
	// Free storage space remaining in the entity (bytes) as a point-in-time snapshot. The available space is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used storage capacity is also reported.
	Free *int64 `json:"Free,omitempty"`
	// Spaced already used by this entity (bytes), as a point-in-time snapshot.
	Used                 *int64 `json:"Used,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationStorageCapacity VirtualizationStorageCapacity

// NewVirtualizationStorageCapacity instantiates a new VirtualizationStorageCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationStorageCapacity(classId string, objectType string) *VirtualizationStorageCapacity {
	this := VirtualizationStorageCapacity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationStorageCapacityWithDefaults instantiates a new VirtualizationStorageCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationStorageCapacityWithDefaults() *VirtualizationStorageCapacity {
	this := VirtualizationStorageCapacity{}
	var classId string = "virtualization.StorageCapacity"
	this.ClassId = classId
	var objectType string = "virtualization.StorageCapacity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationStorageCapacity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageCapacity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationStorageCapacity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationStorageCapacity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageCapacity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationStorageCapacity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationStorageCapacity) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageCapacity) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationStorageCapacity) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *VirtualizationStorageCapacity) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *VirtualizationStorageCapacity) GetFree() int64 {
	if o == nil || o.Free == nil {
		var ret int64
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageCapacity) GetFreeOk() (*int64, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *VirtualizationStorageCapacity) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given int64 and assigns it to the Free field.
func (o *VirtualizationStorageCapacity) SetFree(v int64) {
	o.Free = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *VirtualizationStorageCapacity) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageCapacity) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *VirtualizationStorageCapacity) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *VirtualizationStorageCapacity) SetUsed(v int64) {
	o.Used = &v
}

func (o VirtualizationStorageCapacity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Free != nil {
		toSerialize["Free"] = o.Free
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationStorageCapacity) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationStorageCapacityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The total capacity of the entity (bytes).
		Capacity *int64 `json:"Capacity,omitempty"`
		// Free storage space remaining in the entity (bytes) as a point-in-time snapshot. The available space is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used storage capacity is also reported.
		Free *int64 `json:"Free,omitempty"`
		// Spaced already used by this entity (bytes), as a point-in-time snapshot.
		Used *int64 `json:"Used,omitempty"`
	}

	varVirtualizationStorageCapacityWithoutEmbeddedStruct := VirtualizationStorageCapacityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationStorageCapacityWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationStorageCapacity := _VirtualizationStorageCapacity{}
		varVirtualizationStorageCapacity.ClassId = varVirtualizationStorageCapacityWithoutEmbeddedStruct.ClassId
		varVirtualizationStorageCapacity.ObjectType = varVirtualizationStorageCapacityWithoutEmbeddedStruct.ObjectType
		varVirtualizationStorageCapacity.Capacity = varVirtualizationStorageCapacityWithoutEmbeddedStruct.Capacity
		varVirtualizationStorageCapacity.Free = varVirtualizationStorageCapacityWithoutEmbeddedStruct.Free
		varVirtualizationStorageCapacity.Used = varVirtualizationStorageCapacityWithoutEmbeddedStruct.Used
		*o = VirtualizationStorageCapacity(varVirtualizationStorageCapacity)
	} else {
		return err
	}

	varVirtualizationStorageCapacity := _VirtualizationStorageCapacity{}

	err = json.Unmarshal(bytes, &varVirtualizationStorageCapacity)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationStorageCapacity.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Free")
		delete(additionalProperties, "Used")

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

type NullableVirtualizationStorageCapacity struct {
	value *VirtualizationStorageCapacity
	isSet bool
}

func (v NullableVirtualizationStorageCapacity) Get() *VirtualizationStorageCapacity {
	return v.value
}

func (v *NullableVirtualizationStorageCapacity) Set(val *VirtualizationStorageCapacity) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationStorageCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationStorageCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationStorageCapacity(val *VirtualizationStorageCapacity) *NullableVirtualizationStorageCapacity {
	return &NullableVirtualizationStorageCapacity{value: val, isSet: true}
}

func (v NullableVirtualizationStorageCapacity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationStorageCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}