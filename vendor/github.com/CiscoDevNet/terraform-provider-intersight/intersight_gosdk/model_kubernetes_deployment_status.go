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
)

// checks if the KubernetesDeploymentStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesDeploymentStatus{}

// KubernetesDeploymentStatus Most recently observed status of the Deployment.
type KubernetesDeploymentStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
	// The generation observed by the deployment controller.
	ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
	// Total number of ready pods targeted by this deployment.
	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty"`
	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Replicas *int64 `json:"Replicas,omitempty"`
	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	UpdatedReplicas      *int64 `json:"UpdatedReplicas,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesDeploymentStatus KubernetesDeploymentStatus

// NewKubernetesDeploymentStatus instantiates a new KubernetesDeploymentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesDeploymentStatus(classId string, objectType string) *KubernetesDeploymentStatus {
	this := KubernetesDeploymentStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var availableReplicas int64 = 0
	this.AvailableReplicas = &availableReplicas
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	var readyReplicas int64 = 0
	this.ReadyReplicas = &readyReplicas
	var replicas int64 = 0
	this.Replicas = &replicas
	var updatedReplicas int64 = 0
	this.UpdatedReplicas = &updatedReplicas
	return &this
}

// NewKubernetesDeploymentStatusWithDefaults instantiates a new KubernetesDeploymentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesDeploymentStatusWithDefaults() *KubernetesDeploymentStatus {
	this := KubernetesDeploymentStatus{}
	var classId string = "kubernetes.DeploymentStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.DeploymentStatus"
	this.ObjectType = objectType
	var availableReplicas int64 = 0
	this.AvailableReplicas = &availableReplicas
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	var readyReplicas int64 = 0
	this.ReadyReplicas = &readyReplicas
	var replicas int64 = 0
	this.Replicas = &replicas
	var updatedReplicas int64 = 0
	this.UpdatedReplicas = &updatedReplicas
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesDeploymentStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesDeploymentStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.DeploymentStatus" of the ClassId field.
func (o *KubernetesDeploymentStatus) GetDefaultClassId() interface{} {
	return "kubernetes.DeploymentStatus"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesDeploymentStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesDeploymentStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.DeploymentStatus" of the ObjectType field.
func (o *KubernetesDeploymentStatus) GetDefaultObjectType() interface{} {
	return "kubernetes.DeploymentStatus"
}

// GetAvailableReplicas returns the AvailableReplicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatus) GetAvailableReplicas() int64 {
	if o == nil || IsNil(o.AvailableReplicas) {
		var ret int64
		return ret
	}
	return *o.AvailableReplicas
}

// GetAvailableReplicasOk returns a tuple with the AvailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetAvailableReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.AvailableReplicas) {
		return nil, false
	}
	return o.AvailableReplicas, true
}

// HasAvailableReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatus) HasAvailableReplicas() bool {
	if o != nil && !IsNil(o.AvailableReplicas) {
		return true
	}

	return false
}

// SetAvailableReplicas gets a reference to the given int64 and assigns it to the AvailableReplicas field.
func (o *KubernetesDeploymentStatus) SetAvailableReplicas(v int64) {
	o.AvailableReplicas = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatus) GetObservedGeneration() int64 {
	if o == nil || IsNil(o.ObservedGeneration) {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || IsNil(o.ObservedGeneration) {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatus) HasObservedGeneration() bool {
	if o != nil && !IsNil(o.ObservedGeneration) {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *KubernetesDeploymentStatus) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetReadyReplicas returns the ReadyReplicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatus) GetReadyReplicas() int64 {
	if o == nil || IsNil(o.ReadyReplicas) {
		var ret int64
		return ret
	}
	return *o.ReadyReplicas
}

// GetReadyReplicasOk returns a tuple with the ReadyReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetReadyReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadyReplicas) {
		return nil, false
	}
	return o.ReadyReplicas, true
}

// HasReadyReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatus) HasReadyReplicas() bool {
	if o != nil && !IsNil(o.ReadyReplicas) {
		return true
	}

	return false
}

// SetReadyReplicas gets a reference to the given int64 and assigns it to the ReadyReplicas field.
func (o *KubernetesDeploymentStatus) SetReadyReplicas(v int64) {
	o.ReadyReplicas = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatus) GetReplicas() int64 {
	if o == nil || IsNil(o.Replicas) {
		var ret int64
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatus) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *KubernetesDeploymentStatus) SetReplicas(v int64) {
	o.Replicas = &v
}

// GetUpdatedReplicas returns the UpdatedReplicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatus) GetUpdatedReplicas() int64 {
	if o == nil || IsNil(o.UpdatedReplicas) {
		var ret int64
		return ret
	}
	return *o.UpdatedReplicas
}

// GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatus) GetUpdatedReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedReplicas) {
		return nil, false
	}
	return o.UpdatedReplicas, true
}

// HasUpdatedReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatus) HasUpdatedReplicas() bool {
	if o != nil && !IsNil(o.UpdatedReplicas) {
		return true
	}

	return false
}

// SetUpdatedReplicas gets a reference to the given int64 and assigns it to the UpdatedReplicas field.
func (o *KubernetesDeploymentStatus) SetUpdatedReplicas(v int64) {
	o.UpdatedReplicas = &v
}

func (o KubernetesDeploymentStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesDeploymentStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AvailableReplicas) {
		toSerialize["AvailableReplicas"] = o.AvailableReplicas
	}
	if !IsNil(o.ObservedGeneration) {
		toSerialize["ObservedGeneration"] = o.ObservedGeneration
	}
	if !IsNil(o.ReadyReplicas) {
		toSerialize["ReadyReplicas"] = o.ReadyReplicas
	}
	if !IsNil(o.Replicas) {
		toSerialize["Replicas"] = o.Replicas
	}
	if !IsNil(o.UpdatedReplicas) {
		toSerialize["UpdatedReplicas"] = o.UpdatedReplicas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesDeploymentStatus) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesDeploymentStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
		AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
		// The generation observed by the deployment controller.
		ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
		// Total number of ready pods targeted by this deployment.
		ReadyReplicas *int64 `json:"ReadyReplicas,omitempty"`
		// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
		Replicas *int64 `json:"Replicas,omitempty"`
		// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
		UpdatedReplicas *int64 `json:"UpdatedReplicas,omitempty"`
	}

	varKubernetesDeploymentStatusWithoutEmbeddedStruct := KubernetesDeploymentStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesDeploymentStatusWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesDeploymentStatus := _KubernetesDeploymentStatus{}
		varKubernetesDeploymentStatus.ClassId = varKubernetesDeploymentStatusWithoutEmbeddedStruct.ClassId
		varKubernetesDeploymentStatus.ObjectType = varKubernetesDeploymentStatusWithoutEmbeddedStruct.ObjectType
		varKubernetesDeploymentStatus.AvailableReplicas = varKubernetesDeploymentStatusWithoutEmbeddedStruct.AvailableReplicas
		varKubernetesDeploymentStatus.ObservedGeneration = varKubernetesDeploymentStatusWithoutEmbeddedStruct.ObservedGeneration
		varKubernetesDeploymentStatus.ReadyReplicas = varKubernetesDeploymentStatusWithoutEmbeddedStruct.ReadyReplicas
		varKubernetesDeploymentStatus.Replicas = varKubernetesDeploymentStatusWithoutEmbeddedStruct.Replicas
		varKubernetesDeploymentStatus.UpdatedReplicas = varKubernetesDeploymentStatusWithoutEmbeddedStruct.UpdatedReplicas
		*o = KubernetesDeploymentStatus(varKubernetesDeploymentStatus)
	} else {
		return err
	}

	varKubernetesDeploymentStatus := _KubernetesDeploymentStatus{}

	err = json.Unmarshal(data, &varKubernetesDeploymentStatus)
	if err == nil {
		o.MoBaseComplexType = varKubernetesDeploymentStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvailableReplicas")
		delete(additionalProperties, "ObservedGeneration")
		delete(additionalProperties, "ReadyReplicas")
		delete(additionalProperties, "Replicas")
		delete(additionalProperties, "UpdatedReplicas")

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

type NullableKubernetesDeploymentStatus struct {
	value *KubernetesDeploymentStatus
	isSet bool
}

func (v NullableKubernetesDeploymentStatus) Get() *KubernetesDeploymentStatus {
	return v.value
}

func (v *NullableKubernetesDeploymentStatus) Set(val *KubernetesDeploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesDeploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesDeploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesDeploymentStatus(val *KubernetesDeploymentStatus) *NullableKubernetesDeploymentStatus {
	return &NullableKubernetesDeploymentStatus{value: val, isSet: true}
}

func (v NullableKubernetesDeploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesDeploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
