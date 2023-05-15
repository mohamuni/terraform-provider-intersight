/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// InfraBaseGpuConfigurationAllOf Definition of the list of properties defined in 'infra.BaseGpuConfiguration', excluding properties defined in parent classes.
type InfraBaseGpuConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The amount of memory on the GPU (GBs).
	MemorySize           *int64 `json:"MemorySize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfraBaseGpuConfigurationAllOf InfraBaseGpuConfigurationAllOf

// NewInfraBaseGpuConfigurationAllOf instantiates a new InfraBaseGpuConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraBaseGpuConfigurationAllOf(classId string, objectType string) *InfraBaseGpuConfigurationAllOf {
	this := InfraBaseGpuConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInfraBaseGpuConfigurationAllOfWithDefaults instantiates a new InfraBaseGpuConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraBaseGpuConfigurationAllOfWithDefaults() *InfraBaseGpuConfigurationAllOf {
	this := InfraBaseGpuConfigurationAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *InfraBaseGpuConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InfraBaseGpuConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InfraBaseGpuConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InfraBaseGpuConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InfraBaseGpuConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InfraBaseGpuConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *InfraBaseGpuConfigurationAllOf) GetMemorySize() int64 {
	if o == nil || o.MemorySize == nil {
		var ret int64
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraBaseGpuConfigurationAllOf) GetMemorySizeOk() (*int64, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *InfraBaseGpuConfigurationAllOf) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given int64 and assigns it to the MemorySize field.
func (o *InfraBaseGpuConfigurationAllOf) SetMemorySize(v int64) {
	o.MemorySize = &v
}

func (o InfraBaseGpuConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MemorySize != nil {
		toSerialize["MemorySize"] = o.MemorySize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InfraBaseGpuConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInfraBaseGpuConfigurationAllOf := _InfraBaseGpuConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varInfraBaseGpuConfigurationAllOf); err == nil {
		*o = InfraBaseGpuConfigurationAllOf(varInfraBaseGpuConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemorySize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfraBaseGpuConfigurationAllOf struct {
	value *InfraBaseGpuConfigurationAllOf
	isSet bool
}

func (v NullableInfraBaseGpuConfigurationAllOf) Get() *InfraBaseGpuConfigurationAllOf {
	return v.value
}

func (v *NullableInfraBaseGpuConfigurationAllOf) Set(val *InfraBaseGpuConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraBaseGpuConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraBaseGpuConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraBaseGpuConfigurationAllOf(val *InfraBaseGpuConfigurationAllOf) *NullableInfraBaseGpuConfigurationAllOf {
	return &NullableInfraBaseGpuConfigurationAllOf{value: val, isSet: true}
}

func (v NullableInfraBaseGpuConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraBaseGpuConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}