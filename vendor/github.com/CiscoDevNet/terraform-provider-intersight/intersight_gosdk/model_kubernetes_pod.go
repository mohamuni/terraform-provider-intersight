/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the KubernetesPod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesPod{}

// KubernetesPod Pod represents a Kubernetes Pod. In Intersight, it is a read only model that represent Kubernetes Pod information. In Kubernetes, A Pod is the basic execution unit of a Kubernetes application the smallest and simplest unit in the Kubernetes object model that you create or deploy. A Pod represents processes running on your Cluster.
type KubernetesPod struct {
	KubernetesAbstractPod
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	Status               NullableKubernetesPodStatus                 `json:"Status,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesPod KubernetesPod

// NewKubernetesPod instantiates a new KubernetesPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesPod(classId string, objectType string) *KubernetesPod {
	this := KubernetesPod{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesPodWithDefaults instantiates a new KubernetesPod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesPodWithDefaults() *KubernetesPod {
	this := KubernetesPod{}
	var classId string = "kubernetes.Pod"
	this.ClassId = classId
	var objectType string = "kubernetes.Pod"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesPod) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesPod) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.Pod" of the ClassId field.
func (o *KubernetesPod) GetDefaultClassId() interface{} {
	return "kubernetes.Pod"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesPod) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesPod) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.Pod" of the ObjectType field.
func (o *KubernetesPod) GetDefaultObjectType() interface{} {
	return "kubernetes.Pod"
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesPod) GetStatus() KubernetesPodStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret KubernetesPodStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesPod) GetStatusOk() (*KubernetesPodStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesPod) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableKubernetesPodStatus and assigns it to the Status field.
func (o *KubernetesPod) SetStatus(v KubernetesPodStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesPod) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesPod) UnsetStatus() {
	o.Status.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesPod) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesPod) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesPod) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesPod) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *KubernetesPod) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *KubernetesPod) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o KubernetesPod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesPod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesAbstractPod, errKubernetesAbstractPod := json.Marshal(o.KubernetesAbstractPod)
	if errKubernetesAbstractPod != nil {
		return map[string]interface{}{}, errKubernetesAbstractPod
	}
	errKubernetesAbstractPod = json.Unmarshal([]byte(serializedKubernetesAbstractPod), &toSerialize)
	if errKubernetesAbstractPod != nil {
		return map[string]interface{}{}, errKubernetesAbstractPod
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesPod) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesPodWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                                      `json:"ObjectType"`
		Status           NullableKubernetesPodStatus                 `json:"Status,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesPodWithoutEmbeddedStruct := KubernetesPodWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesPodWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesPod := _KubernetesPod{}
		varKubernetesPod.ClassId = varKubernetesPodWithoutEmbeddedStruct.ClassId
		varKubernetesPod.ObjectType = varKubernetesPodWithoutEmbeddedStruct.ObjectType
		varKubernetesPod.Status = varKubernetesPodWithoutEmbeddedStruct.Status
		varKubernetesPod.RegisteredDevice = varKubernetesPodWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesPod(varKubernetesPod)
	} else {
		return err
	}

	varKubernetesPod := _KubernetesPod{}

	err = json.Unmarshal(data, &varKubernetesPod)
	if err == nil {
		o.KubernetesAbstractPod = varKubernetesPod.KubernetesAbstractPod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectKubernetesAbstractPod := reflect.ValueOf(o.KubernetesAbstractPod)
		for i := 0; i < reflectKubernetesAbstractPod.Type().NumField(); i++ {
			t := reflectKubernetesAbstractPod.Type().Field(i)

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

type NullableKubernetesPod struct {
	value *KubernetesPod
	isSet bool
}

func (v NullableKubernetesPod) Get() *KubernetesPod {
	return v.value
}

func (v *NullableKubernetesPod) Set(val *KubernetesPod) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesPod) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesPod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesPod(val *KubernetesPod) *NullableKubernetesPod {
	return &NullableKubernetesPod{value: val, isSet: true}
}

func (v NullableKubernetesPod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesPod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
