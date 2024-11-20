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

// checks if the KubernetesIngress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesIngress{}

// KubernetesIngress Ingress inventory represent a Kubernetes Ingress. In Kubernetes, Ingress exposes HTTP and HTTPS routes from outside the cluster to services within the cluster. Traffic routing is controlled by rules defined on the Ingress resource.
type KubernetesIngress struct {
	KubernetesAbstractIngress
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	Status               NullableKubernetesIngressStatus             `json:"Status,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesIngress KubernetesIngress

// NewKubernetesIngress instantiates a new KubernetesIngress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesIngress(classId string, objectType string) *KubernetesIngress {
	this := KubernetesIngress{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesIngressWithDefaults instantiates a new KubernetesIngress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesIngressWithDefaults() *KubernetesIngress {
	this := KubernetesIngress{}
	var classId string = "kubernetes.Ingress"
	this.ClassId = classId
	var objectType string = "kubernetes.Ingress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesIngress) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesIngress) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesIngress) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.Ingress" of the ClassId field.
func (o *KubernetesIngress) GetDefaultClassId() interface{} {
	return "kubernetes.Ingress"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesIngress) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesIngress) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesIngress) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.Ingress" of the ObjectType field.
func (o *KubernetesIngress) GetDefaultObjectType() interface{} {
	return "kubernetes.Ingress"
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesIngress) GetStatus() KubernetesIngressStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret KubernetesIngressStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesIngress) GetStatusOk() (*KubernetesIngressStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesIngress) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableKubernetesIngressStatus and assigns it to the Status field.
func (o *KubernetesIngress) SetStatus(v KubernetesIngressStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesIngress) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesIngress) UnsetStatus() {
	o.Status.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesIngress) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesIngress) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesIngress) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesIngress) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *KubernetesIngress) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *KubernetesIngress) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o KubernetesIngress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesIngress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesAbstractIngress, errKubernetesAbstractIngress := json.Marshal(o.KubernetesAbstractIngress)
	if errKubernetesAbstractIngress != nil {
		return map[string]interface{}{}, errKubernetesAbstractIngress
	}
	errKubernetesAbstractIngress = json.Unmarshal([]byte(serializedKubernetesAbstractIngress), &toSerialize)
	if errKubernetesAbstractIngress != nil {
		return map[string]interface{}{}, errKubernetesAbstractIngress
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

func (o *KubernetesIngress) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesIngressWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                                      `json:"ObjectType"`
		Status           NullableKubernetesIngressStatus             `json:"Status,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesIngressWithoutEmbeddedStruct := KubernetesIngressWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesIngressWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesIngress := _KubernetesIngress{}
		varKubernetesIngress.ClassId = varKubernetesIngressWithoutEmbeddedStruct.ClassId
		varKubernetesIngress.ObjectType = varKubernetesIngressWithoutEmbeddedStruct.ObjectType
		varKubernetesIngress.Status = varKubernetesIngressWithoutEmbeddedStruct.Status
		varKubernetesIngress.RegisteredDevice = varKubernetesIngressWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesIngress(varKubernetesIngress)
	} else {
		return err
	}

	varKubernetesIngress := _KubernetesIngress{}

	err = json.Unmarshal(data, &varKubernetesIngress)
	if err == nil {
		o.KubernetesAbstractIngress = varKubernetesIngress.KubernetesAbstractIngress
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
		reflectKubernetesAbstractIngress := reflect.ValueOf(o.KubernetesAbstractIngress)
		for i := 0; i < reflectKubernetesAbstractIngress.Type().NumField(); i++ {
			t := reflectKubernetesAbstractIngress.Type().Field(i)

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

type NullableKubernetesIngress struct {
	value *KubernetesIngress
	isSet bool
}

func (v NullableKubernetesIngress) Get() *KubernetesIngress {
	return v.value
}

func (v *NullableKubernetesIngress) Set(val *KubernetesIngress) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesIngress) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesIngress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesIngress(val *KubernetesIngress) *NullableKubernetesIngress {
	return &NullableKubernetesIngress{value: val, isSet: true}
}

func (v NullableKubernetesIngress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesIngress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
