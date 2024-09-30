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

// checks if the VirtualizationVmwareVirtualMachineGpu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareVirtualMachineGpu{}

// VirtualizationVmwareVirtualMachineGpu Common attributes of virtual GPU device on a VMware virtual machine.
type VirtualizationVmwareVirtualMachineGpu struct {
	VirtualizationBaseVirtualMachineGpu
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The internally assigned key of this virtual GPU device. This entity is not manipulated by users.
	Key *int64 `json:"Key,omitempty"`
	// Identity of the virtual machine where the virtual gpu is created.
	VmIdentity           *string `json:"VmIdentity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVirtualMachineGpu VirtualizationVmwareVirtualMachineGpu

// NewVirtualizationVmwareVirtualMachineGpu instantiates a new VirtualizationVmwareVirtualMachineGpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVirtualMachineGpu(classId string, objectType string) *VirtualizationVmwareVirtualMachineGpu {
	this := VirtualizationVmwareVirtualMachineGpu{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVirtualMachineGpuWithDefaults instantiates a new VirtualizationVmwareVirtualMachineGpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVirtualMachineGpuWithDefaults() *VirtualizationVmwareVirtualMachineGpu {
	this := VirtualizationVmwareVirtualMachineGpu{}
	var classId string = "virtualization.VmwareVirtualMachineGpu"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualMachineGpu"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVirtualMachineGpu) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineGpu) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVirtualMachineGpu) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareVirtualMachineGpu" of the ClassId field.
func (o *VirtualizationVmwareVirtualMachineGpu) GetDefaultClassId() interface{} {
	return "virtualization.VmwareVirtualMachineGpu"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVirtualMachineGpu) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineGpu) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVirtualMachineGpu) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareVirtualMachineGpu" of the ObjectType field.
func (o *VirtualizationVmwareVirtualMachineGpu) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareVirtualMachineGpu"
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineGpu) GetKey() int64 {
	if o == nil || IsNil(o.Key) {
		var ret int64
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineGpu) GetKeyOk() (*int64, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineGpu) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given int64 and assigns it to the Key field.
func (o *VirtualizationVmwareVirtualMachineGpu) SetKey(v int64) {
	o.Key = &v
}

// GetVmIdentity returns the VmIdentity field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineGpu) GetVmIdentity() string {
	if o == nil || IsNil(o.VmIdentity) {
		var ret string
		return ret
	}
	return *o.VmIdentity
}

// GetVmIdentityOk returns a tuple with the VmIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineGpu) GetVmIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.VmIdentity) {
		return nil, false
	}
	return o.VmIdentity, true
}

// HasVmIdentity returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineGpu) HasVmIdentity() bool {
	if o != nil && !IsNil(o.VmIdentity) {
		return true
	}

	return false
}

// SetVmIdentity gets a reference to the given string and assigns it to the VmIdentity field.
func (o *VirtualizationVmwareVirtualMachineGpu) SetVmIdentity(v string) {
	o.VmIdentity = &v
}

func (o VirtualizationVmwareVirtualMachineGpu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareVirtualMachineGpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualMachineGpu, errVirtualizationBaseVirtualMachineGpu := json.Marshal(o.VirtualizationBaseVirtualMachineGpu)
	if errVirtualizationBaseVirtualMachineGpu != nil {
		return map[string]interface{}{}, errVirtualizationBaseVirtualMachineGpu
	}
	errVirtualizationBaseVirtualMachineGpu = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualMachineGpu), &toSerialize)
	if errVirtualizationBaseVirtualMachineGpu != nil {
		return map[string]interface{}{}, errVirtualizationBaseVirtualMachineGpu
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.VmIdentity) {
		toSerialize["VmIdentity"] = o.VmIdentity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareVirtualMachineGpu) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The internally assigned key of this virtual GPU device. This entity is not manipulated by users.
		Key *int64 `json:"Key,omitempty"`
		// Identity of the virtual machine where the virtual gpu is created.
		VmIdentity *string `json:"VmIdentity,omitempty"`
	}

	varVirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct := VirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVirtualMachineGpu := _VirtualizationVmwareVirtualMachineGpu{}
		varVirtualizationVmwareVirtualMachineGpu.ClassId = varVirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVirtualMachineGpu.ObjectType = varVirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVirtualMachineGpu.Key = varVirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct.Key
		varVirtualizationVmwareVirtualMachineGpu.VmIdentity = varVirtualizationVmwareVirtualMachineGpuWithoutEmbeddedStruct.VmIdentity
		*o = VirtualizationVmwareVirtualMachineGpu(varVirtualizationVmwareVirtualMachineGpu)
	} else {
		return err
	}

	varVirtualizationVmwareVirtualMachineGpu := _VirtualizationVmwareVirtualMachineGpu{}

	err = json.Unmarshal(data, &varVirtualizationVmwareVirtualMachineGpu)
	if err == nil {
		o.VirtualizationBaseVirtualMachineGpu = varVirtualizationVmwareVirtualMachineGpu.VirtualizationBaseVirtualMachineGpu
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "VmIdentity")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualMachineGpu := reflect.ValueOf(o.VirtualizationBaseVirtualMachineGpu)
		for i := 0; i < reflectVirtualizationBaseVirtualMachineGpu.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualMachineGpu.Type().Field(i)

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

type NullableVirtualizationVmwareVirtualMachineGpu struct {
	value *VirtualizationVmwareVirtualMachineGpu
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualMachineGpu) Get() *VirtualizationVmwareVirtualMachineGpu {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualMachineGpu) Set(val *VirtualizationVmwareVirtualMachineGpu) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualMachineGpu) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualMachineGpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualMachineGpu(val *VirtualizationVmwareVirtualMachineGpu) *NullableVirtualizationVmwareVirtualMachineGpu {
	return &NullableVirtualizationVmwareVirtualMachineGpu{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualMachineGpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualMachineGpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
