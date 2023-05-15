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
	"reflect"
	"strings"
)

// KubernetesBaseGpuProduct Common information of a GPU product.
type KubernetesBaseGpuProduct struct {
	KubernetesBaseProduct
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Memory size of a GPU product in GB.
	MemorySize           *int64 `json:"MemorySize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaseGpuProduct KubernetesBaseGpuProduct

// NewKubernetesBaseGpuProduct instantiates a new KubernetesBaseGpuProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaseGpuProduct(classId string, objectType string) *KubernetesBaseGpuProduct {
	this := KubernetesBaseGpuProduct{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesBaseGpuProductWithDefaults instantiates a new KubernetesBaseGpuProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaseGpuProductWithDefaults() *KubernetesBaseGpuProduct {
	this := KubernetesBaseGpuProduct{}
	var classId string = "kubernetes.NvidiaGpuProduct"
	this.ClassId = classId
	var objectType string = "kubernetes.NvidiaGpuProduct"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaseGpuProduct) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseGpuProduct) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaseGpuProduct) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaseGpuProduct) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseGpuProduct) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaseGpuProduct) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *KubernetesBaseGpuProduct) GetMemorySize() int64 {
	if o == nil || o.MemorySize == nil {
		var ret int64
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseGpuProduct) GetMemorySizeOk() (*int64, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *KubernetesBaseGpuProduct) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given int64 and assigns it to the MemorySize field.
func (o *KubernetesBaseGpuProduct) SetMemorySize(v int64) {
	o.MemorySize = &v
}

func (o KubernetesBaseGpuProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesBaseProduct, errKubernetesBaseProduct := json.Marshal(o.KubernetesBaseProduct)
	if errKubernetesBaseProduct != nil {
		return []byte{}, errKubernetesBaseProduct
	}
	errKubernetesBaseProduct = json.Unmarshal([]byte(serializedKubernetesBaseProduct), &toSerialize)
	if errKubernetesBaseProduct != nil {
		return []byte{}, errKubernetesBaseProduct
	}
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

func (o *KubernetesBaseGpuProduct) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesBaseGpuProductWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Memory size of a GPU product in GB.
		MemorySize *int64 `json:"MemorySize,omitempty"`
	}

	varKubernetesBaseGpuProductWithoutEmbeddedStruct := KubernetesBaseGpuProductWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesBaseGpuProductWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesBaseGpuProduct := _KubernetesBaseGpuProduct{}
		varKubernetesBaseGpuProduct.ClassId = varKubernetesBaseGpuProductWithoutEmbeddedStruct.ClassId
		varKubernetesBaseGpuProduct.ObjectType = varKubernetesBaseGpuProductWithoutEmbeddedStruct.ObjectType
		varKubernetesBaseGpuProduct.MemorySize = varKubernetesBaseGpuProductWithoutEmbeddedStruct.MemorySize
		*o = KubernetesBaseGpuProduct(varKubernetesBaseGpuProduct)
	} else {
		return err
	}

	varKubernetesBaseGpuProduct := _KubernetesBaseGpuProduct{}

	err = json.Unmarshal(bytes, &varKubernetesBaseGpuProduct)
	if err == nil {
		o.KubernetesBaseProduct = varKubernetesBaseGpuProduct.KubernetesBaseProduct
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemorySize")

		// remove fields from embedded structs
		reflectKubernetesBaseProduct := reflect.ValueOf(o.KubernetesBaseProduct)
		for i := 0; i < reflectKubernetesBaseProduct.Type().NumField(); i++ {
			t := reflectKubernetesBaseProduct.Type().Field(i)

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

type NullableKubernetesBaseGpuProduct struct {
	value *KubernetesBaseGpuProduct
	isSet bool
}

func (v NullableKubernetesBaseGpuProduct) Get() *KubernetesBaseGpuProduct {
	return v.value
}

func (v *NullableKubernetesBaseGpuProduct) Set(val *KubernetesBaseGpuProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaseGpuProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaseGpuProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaseGpuProduct(val *KubernetesBaseGpuProduct) *NullableKubernetesBaseGpuProduct {
	return &NullableKubernetesBaseGpuProduct{value: val, isSet: true}
}

func (v NullableKubernetesBaseGpuProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaseGpuProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}