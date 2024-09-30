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

// checks if the StorageBaseVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageBaseVolume{}

// StorageBaseVolume Generic storage volume object. It is a provisioned storage identity which can be mapped to host to enable host access.
type StorageBaseVolume struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Short description about the volume.
	Description *string `json:"Description,omitempty"`
	// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
	NaaId *string `json:"NaaId,omitempty"`
	// Named entity of the volume.
	Name *string `json:"Name,omitempty"`
	// User provisioned volume size. It is the size exposed to host.
	Size                 *int64                      `json:"Size,omitempty"`
	StorageUtilization   NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseVolume StorageBaseVolume

// NewStorageBaseVolume instantiates a new StorageBaseVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseVolume(classId string, objectType string) *StorageBaseVolume {
	this := StorageBaseVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseVolumeWithDefaults instantiates a new StorageBaseVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseVolumeWithDefaults() *StorageBaseVolume {
	this := StorageBaseVolume{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageBaseVolume) SetDescription(v string) {
	o.Description = &v
}

// GetNaaId returns the NaaId field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetNaaId() string {
	if o == nil || IsNil(o.NaaId) {
		var ret string
		return ret
	}
	return *o.NaaId
}

// GetNaaIdOk returns a tuple with the NaaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetNaaIdOk() (*string, bool) {
	if o == nil || IsNil(o.NaaId) {
		return nil, false
	}
	return o.NaaId, true
}

// HasNaaId returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasNaaId() bool {
	if o != nil && !IsNil(o.NaaId) {
		return true
	}

	return false
}

// SetNaaId gets a reference to the given string and assigns it to the NaaId field.
func (o *StorageBaseVolume) SetNaaId(v string) {
	o.NaaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseVolume) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseVolume) SetSize(v int64) {
	o.Size = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseVolume) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || IsNil(o.StorageUtilization.Get()) {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseVolume) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseVolume) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseVolume) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseVolume) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

func (o StorageBaseVolume) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageBaseVolume) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.NaaId) {
		toSerialize["NaaId"] = o.NaaId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageBaseVolume) UnmarshalJSON(data []byte) (err error) {
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
	type StorageBaseVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Short description about the volume.
		Description *string `json:"Description,omitempty"`
		// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
		NaaId *string `json:"NaaId,omitempty"`
		// Named entity of the volume.
		Name *string `json:"Name,omitempty"`
		// User provisioned volume size. It is the size exposed to host.
		Size               *int64                      `json:"Size,omitempty"`
		StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	}

	varStorageBaseVolumeWithoutEmbeddedStruct := StorageBaseVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageBaseVolumeWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseVolume := _StorageBaseVolume{}
		varStorageBaseVolume.ClassId = varStorageBaseVolumeWithoutEmbeddedStruct.ClassId
		varStorageBaseVolume.ObjectType = varStorageBaseVolumeWithoutEmbeddedStruct.ObjectType
		varStorageBaseVolume.Description = varStorageBaseVolumeWithoutEmbeddedStruct.Description
		varStorageBaseVolume.NaaId = varStorageBaseVolumeWithoutEmbeddedStruct.NaaId
		varStorageBaseVolume.Name = varStorageBaseVolumeWithoutEmbeddedStruct.Name
		varStorageBaseVolume.Size = varStorageBaseVolumeWithoutEmbeddedStruct.Size
		varStorageBaseVolume.StorageUtilization = varStorageBaseVolumeWithoutEmbeddedStruct.StorageUtilization
		*o = StorageBaseVolume(varStorageBaseVolume)
	} else {
		return err
	}

	varStorageBaseVolume := _StorageBaseVolume{}

	err = json.Unmarshal(data, &varStorageBaseVolume)
	if err == nil {
		o.MoBaseMo = varStorageBaseVolume.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "NaaId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "StorageUtilization")

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

type NullableStorageBaseVolume struct {
	value *StorageBaseVolume
	isSet bool
}

func (v NullableStorageBaseVolume) Get() *StorageBaseVolume {
	return v.value
}

func (v *NullableStorageBaseVolume) Set(val *StorageBaseVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseVolume(val *StorageBaseVolume) *NullableStorageBaseVolume {
	return &NullableStorageBaseVolume{value: val, isSet: true}
}

func (v NullableStorageBaseVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
