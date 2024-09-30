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

// checks if the StorageNvmeDedicatedHotSpareConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNvmeDedicatedHotSpareConfiguration{}

// StorageNvmeDedicatedHotSpareConfiguration Dedicated Hot Spare type models a single dedicated hot spare that needs to be created for NVMe RAID drive group on reboot. Used in activation workflow.
type StorageNvmeDedicatedHotSpareConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This defines if the vd does not exists at endpoint for specific storage controller per drive group. Only if it's false we will create dedicated hot spares for the existing vds.
	IsNewVd *bool `json:"IsNewVd,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed.
	Name *string `json:"Name,omitempty"`
	// Physical Disk Slot that is used as dedicated hot spare.
	Slot *string `json:"Slot,omitempty"`
	// The volume dn of the dedicated hot spare, this will be unique for each dedicated hot spare.
	VolumeDn             *string `json:"VolumeDn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNvmeDedicatedHotSpareConfiguration StorageNvmeDedicatedHotSpareConfiguration

// NewStorageNvmeDedicatedHotSpareConfiguration instantiates a new StorageNvmeDedicatedHotSpareConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNvmeDedicatedHotSpareConfiguration(classId string, objectType string) *StorageNvmeDedicatedHotSpareConfiguration {
	this := StorageNvmeDedicatedHotSpareConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNvmeDedicatedHotSpareConfigurationWithDefaults instantiates a new StorageNvmeDedicatedHotSpareConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNvmeDedicatedHotSpareConfigurationWithDefaults() *StorageNvmeDedicatedHotSpareConfiguration {
	this := StorageNvmeDedicatedHotSpareConfiguration{}
	var classId string = "storage.NvmeDedicatedHotSpareConfiguration"
	this.ClassId = classId
	var objectType string = "storage.NvmeDedicatedHotSpareConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNvmeDedicatedHotSpareConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NvmeDedicatedHotSpareConfiguration" of the ClassId field.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetDefaultClassId() interface{} {
	return "storage.NvmeDedicatedHotSpareConfiguration"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNvmeDedicatedHotSpareConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NvmeDedicatedHotSpareConfiguration" of the ObjectType field.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetDefaultObjectType() interface{} {
	return "storage.NvmeDedicatedHotSpareConfiguration"
}

// GetIsNewVd returns the IsNewVd field value if set, zero value otherwise.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetIsNewVd() bool {
	if o == nil || IsNil(o.IsNewVd) {
		var ret bool
		return ret
	}
	return *o.IsNewVd
}

// GetIsNewVdOk returns a tuple with the IsNewVd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetIsNewVdOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNewVd) {
		return nil, false
	}
	return o.IsNewVd, true
}

// HasIsNewVd returns a boolean if a field has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) HasIsNewVd() bool {
	if o != nil && !IsNil(o.IsNewVd) {
		return true
	}

	return false
}

// SetIsNewVd gets a reference to the given bool and assigns it to the IsNewVd field.
func (o *StorageNvmeDedicatedHotSpareConfiguration) SetIsNewVd(v bool) {
	o.IsNewVd = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNvmeDedicatedHotSpareConfiguration) SetName(v string) {
	o.Name = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *StorageNvmeDedicatedHotSpareConfiguration) SetSlot(v string) {
	o.Slot = &v
}

// GetVolumeDn returns the VolumeDn field value if set, zero value otherwise.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetVolumeDn() string {
	if o == nil || IsNil(o.VolumeDn) {
		var ret string
		return ret
	}
	return *o.VolumeDn
}

// GetVolumeDnOk returns a tuple with the VolumeDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) GetVolumeDnOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeDn) {
		return nil, false
	}
	return o.VolumeDn, true
}

// HasVolumeDn returns a boolean if a field has been set.
func (o *StorageNvmeDedicatedHotSpareConfiguration) HasVolumeDn() bool {
	if o != nil && !IsNil(o.VolumeDn) {
		return true
	}

	return false
}

// SetVolumeDn gets a reference to the given string and assigns it to the VolumeDn field.
func (o *StorageNvmeDedicatedHotSpareConfiguration) SetVolumeDn(v string) {
	o.VolumeDn = &v
}

func (o StorageNvmeDedicatedHotSpareConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNvmeDedicatedHotSpareConfiguration) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsNewVd) {
		toSerialize["IsNewVd"] = o.IsNewVd
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Slot) {
		toSerialize["Slot"] = o.Slot
	}
	if !IsNil(o.VolumeDn) {
		toSerialize["VolumeDn"] = o.VolumeDn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNvmeDedicatedHotSpareConfiguration) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This defines if the vd does not exists at endpoint for specific storage controller per drive group. Only if it's false we will create dedicated hot spares for the existing vds.
		IsNewVd *bool `json:"IsNewVd,omitempty"`
		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed.
		Name *string `json:"Name,omitempty"`
		// Physical Disk Slot that is used as dedicated hot spare.
		Slot *string `json:"Slot,omitempty"`
		// The volume dn of the dedicated hot spare, this will be unique for each dedicated hot spare.
		VolumeDn *string `json:"VolumeDn,omitempty"`
	}

	varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct := StorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varStorageNvmeDedicatedHotSpareConfiguration := _StorageNvmeDedicatedHotSpareConfiguration{}
		varStorageNvmeDedicatedHotSpareConfiguration.ClassId = varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct.ClassId
		varStorageNvmeDedicatedHotSpareConfiguration.ObjectType = varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct.ObjectType
		varStorageNvmeDedicatedHotSpareConfiguration.IsNewVd = varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct.IsNewVd
		varStorageNvmeDedicatedHotSpareConfiguration.Name = varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct.Name
		varStorageNvmeDedicatedHotSpareConfiguration.Slot = varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct.Slot
		varStorageNvmeDedicatedHotSpareConfiguration.VolumeDn = varStorageNvmeDedicatedHotSpareConfigurationWithoutEmbeddedStruct.VolumeDn
		*o = StorageNvmeDedicatedHotSpareConfiguration(varStorageNvmeDedicatedHotSpareConfiguration)
	} else {
		return err
	}

	varStorageNvmeDedicatedHotSpareConfiguration := _StorageNvmeDedicatedHotSpareConfiguration{}

	err = json.Unmarshal(data, &varStorageNvmeDedicatedHotSpareConfiguration)
	if err == nil {
		o.MoBaseComplexType = varStorageNvmeDedicatedHotSpareConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsNewVd")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Slot")
		delete(additionalProperties, "VolumeDn")

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

type NullableStorageNvmeDedicatedHotSpareConfiguration struct {
	value *StorageNvmeDedicatedHotSpareConfiguration
	isSet bool
}

func (v NullableStorageNvmeDedicatedHotSpareConfiguration) Get() *StorageNvmeDedicatedHotSpareConfiguration {
	return v.value
}

func (v *NullableStorageNvmeDedicatedHotSpareConfiguration) Set(val *StorageNvmeDedicatedHotSpareConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNvmeDedicatedHotSpareConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNvmeDedicatedHotSpareConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNvmeDedicatedHotSpareConfiguration(val *StorageNvmeDedicatedHotSpareConfiguration) *NullableStorageNvmeDedicatedHotSpareConfiguration {
	return &NullableStorageNvmeDedicatedHotSpareConfiguration{value: val, isSet: true}
}

func (v NullableStorageNvmeDedicatedHotSpareConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNvmeDedicatedHotSpareConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
