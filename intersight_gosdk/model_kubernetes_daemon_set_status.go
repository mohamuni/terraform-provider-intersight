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

// checks if the KubernetesDaemonSetStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesDaemonSetStatus{}

// KubernetesDaemonSetStatus The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only.
type KubernetesDaemonSetStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod.
	CurrentNumberScheduled *int64 `json:"CurrentNumberScheduled,omitempty"`
	// The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod).
	DesiredNumberScheduled *int64 `json:"DesiredNumberScheduled,omitempty"`
	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds).
	NumberAvailable *string `json:"NumberAvailable,omitempty"`
	// The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod.
	NumberMisscheduled *int64 `json:"NumberMisscheduled,omitempty"`
	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	NumberReady *int64 `json:"NumberReady,omitempty"`
	// The most recent generation observed by the daemon set controller.
	ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
	// The total number of nodes that are running updated daemon pod.
	UpdatedNumberScheduled *string `json:"UpdatedNumberScheduled,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _KubernetesDaemonSetStatus KubernetesDaemonSetStatus

// NewKubernetesDaemonSetStatus instantiates a new KubernetesDaemonSetStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesDaemonSetStatus(classId string, objectType string) *KubernetesDaemonSetStatus {
	this := KubernetesDaemonSetStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var currentNumberScheduled int64 = 0
	this.CurrentNumberScheduled = &currentNumberScheduled
	var desiredNumberScheduled int64 = 0
	this.DesiredNumberScheduled = &desiredNumberScheduled
	var numberMisscheduled int64 = 0
	this.NumberMisscheduled = &numberMisscheduled
	var numberReady int64 = 0
	this.NumberReady = &numberReady
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	return &this
}

// NewKubernetesDaemonSetStatusWithDefaults instantiates a new KubernetesDaemonSetStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesDaemonSetStatusWithDefaults() *KubernetesDaemonSetStatus {
	this := KubernetesDaemonSetStatus{}
	var classId string = "kubernetes.DaemonSetStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.DaemonSetStatus"
	this.ObjectType = objectType
	var currentNumberScheduled int64 = 0
	this.CurrentNumberScheduled = &currentNumberScheduled
	var desiredNumberScheduled int64 = 0
	this.DesiredNumberScheduled = &desiredNumberScheduled
	var numberMisscheduled int64 = 0
	this.NumberMisscheduled = &numberMisscheduled
	var numberReady int64 = 0
	this.NumberReady = &numberReady
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesDaemonSetStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesDaemonSetStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.DaemonSetStatus" of the ClassId field.
func (o *KubernetesDaemonSetStatus) GetDefaultClassId() interface{} {
	return "kubernetes.DaemonSetStatus"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesDaemonSetStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesDaemonSetStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.DaemonSetStatus" of the ObjectType field.
func (o *KubernetesDaemonSetStatus) GetDefaultObjectType() interface{} {
	return "kubernetes.DaemonSetStatus"
}

// GetCurrentNumberScheduled returns the CurrentNumberScheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetCurrentNumberScheduled() int64 {
	if o == nil || IsNil(o.CurrentNumberScheduled) {
		var ret int64
		return ret
	}
	return *o.CurrentNumberScheduled
}

// GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetCurrentNumberScheduledOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentNumberScheduled) {
		return nil, false
	}
	return o.CurrentNumberScheduled, true
}

// HasCurrentNumberScheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasCurrentNumberScheduled() bool {
	if o != nil && !IsNil(o.CurrentNumberScheduled) {
		return true
	}

	return false
}

// SetCurrentNumberScheduled gets a reference to the given int64 and assigns it to the CurrentNumberScheduled field.
func (o *KubernetesDaemonSetStatus) SetCurrentNumberScheduled(v int64) {
	o.CurrentNumberScheduled = &v
}

// GetDesiredNumberScheduled returns the DesiredNumberScheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetDesiredNumberScheduled() int64 {
	if o == nil || IsNil(o.DesiredNumberScheduled) {
		var ret int64
		return ret
	}
	return *o.DesiredNumberScheduled
}

// GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetDesiredNumberScheduledOk() (*int64, bool) {
	if o == nil || IsNil(o.DesiredNumberScheduled) {
		return nil, false
	}
	return o.DesiredNumberScheduled, true
}

// HasDesiredNumberScheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasDesiredNumberScheduled() bool {
	if o != nil && !IsNil(o.DesiredNumberScheduled) {
		return true
	}

	return false
}

// SetDesiredNumberScheduled gets a reference to the given int64 and assigns it to the DesiredNumberScheduled field.
func (o *KubernetesDaemonSetStatus) SetDesiredNumberScheduled(v int64) {
	o.DesiredNumberScheduled = &v
}

// GetNumberAvailable returns the NumberAvailable field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetNumberAvailable() string {
	if o == nil || IsNil(o.NumberAvailable) {
		var ret string
		return ret
	}
	return *o.NumberAvailable
}

// GetNumberAvailableOk returns a tuple with the NumberAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetNumberAvailableOk() (*string, bool) {
	if o == nil || IsNil(o.NumberAvailable) {
		return nil, false
	}
	return o.NumberAvailable, true
}

// HasNumberAvailable returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasNumberAvailable() bool {
	if o != nil && !IsNil(o.NumberAvailable) {
		return true
	}

	return false
}

// SetNumberAvailable gets a reference to the given string and assigns it to the NumberAvailable field.
func (o *KubernetesDaemonSetStatus) SetNumberAvailable(v string) {
	o.NumberAvailable = &v
}

// GetNumberMisscheduled returns the NumberMisscheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetNumberMisscheduled() int64 {
	if o == nil || IsNil(o.NumberMisscheduled) {
		var ret int64
		return ret
	}
	return *o.NumberMisscheduled
}

// GetNumberMisscheduledOk returns a tuple with the NumberMisscheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetNumberMisscheduledOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberMisscheduled) {
		return nil, false
	}
	return o.NumberMisscheduled, true
}

// HasNumberMisscheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasNumberMisscheduled() bool {
	if o != nil && !IsNil(o.NumberMisscheduled) {
		return true
	}

	return false
}

// SetNumberMisscheduled gets a reference to the given int64 and assigns it to the NumberMisscheduled field.
func (o *KubernetesDaemonSetStatus) SetNumberMisscheduled(v int64) {
	o.NumberMisscheduled = &v
}

// GetNumberReady returns the NumberReady field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetNumberReady() int64 {
	if o == nil || IsNil(o.NumberReady) {
		var ret int64
		return ret
	}
	return *o.NumberReady
}

// GetNumberReadyOk returns a tuple with the NumberReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetNumberReadyOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberReady) {
		return nil, false
	}
	return o.NumberReady, true
}

// HasNumberReady returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasNumberReady() bool {
	if o != nil && !IsNil(o.NumberReady) {
		return true
	}

	return false
}

// SetNumberReady gets a reference to the given int64 and assigns it to the NumberReady field.
func (o *KubernetesDaemonSetStatus) SetNumberReady(v int64) {
	o.NumberReady = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetObservedGeneration() int64 {
	if o == nil || IsNil(o.ObservedGeneration) {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || IsNil(o.ObservedGeneration) {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasObservedGeneration() bool {
	if o != nil && !IsNil(o.ObservedGeneration) {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *KubernetesDaemonSetStatus) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatus) GetUpdatedNumberScheduled() string {
	if o == nil || IsNil(o.UpdatedNumberScheduled) {
		var ret string
		return ret
	}
	return *o.UpdatedNumberScheduled
}

// GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatus) GetUpdatedNumberScheduledOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedNumberScheduled) {
		return nil, false
	}
	return o.UpdatedNumberScheduled, true
}

// HasUpdatedNumberScheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatus) HasUpdatedNumberScheduled() bool {
	if o != nil && !IsNil(o.UpdatedNumberScheduled) {
		return true
	}

	return false
}

// SetUpdatedNumberScheduled gets a reference to the given string and assigns it to the UpdatedNumberScheduled field.
func (o *KubernetesDaemonSetStatus) SetUpdatedNumberScheduled(v string) {
	o.UpdatedNumberScheduled = &v
}

func (o KubernetesDaemonSetStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesDaemonSetStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CurrentNumberScheduled) {
		toSerialize["CurrentNumberScheduled"] = o.CurrentNumberScheduled
	}
	if !IsNil(o.DesiredNumberScheduled) {
		toSerialize["DesiredNumberScheduled"] = o.DesiredNumberScheduled
	}
	if !IsNil(o.NumberAvailable) {
		toSerialize["NumberAvailable"] = o.NumberAvailable
	}
	if !IsNil(o.NumberMisscheduled) {
		toSerialize["NumberMisscheduled"] = o.NumberMisscheduled
	}
	if !IsNil(o.NumberReady) {
		toSerialize["NumberReady"] = o.NumberReady
	}
	if !IsNil(o.ObservedGeneration) {
		toSerialize["ObservedGeneration"] = o.ObservedGeneration
	}
	if !IsNil(o.UpdatedNumberScheduled) {
		toSerialize["UpdatedNumberScheduled"] = o.UpdatedNumberScheduled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesDaemonSetStatus) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesDaemonSetStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod.
		CurrentNumberScheduled *int64 `json:"CurrentNumberScheduled,omitempty"`
		// The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod).
		DesiredNumberScheduled *int64 `json:"DesiredNumberScheduled,omitempty"`
		// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds).
		NumberAvailable *string `json:"NumberAvailable,omitempty"`
		// The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod.
		NumberMisscheduled *int64 `json:"NumberMisscheduled,omitempty"`
		// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
		NumberReady *int64 `json:"NumberReady,omitempty"`
		// The most recent generation observed by the daemon set controller.
		ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
		// The total number of nodes that are running updated daemon pod.
		UpdatedNumberScheduled *string `json:"UpdatedNumberScheduled,omitempty"`
	}

	varKubernetesDaemonSetStatusWithoutEmbeddedStruct := KubernetesDaemonSetStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesDaemonSetStatusWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesDaemonSetStatus := _KubernetesDaemonSetStatus{}
		varKubernetesDaemonSetStatus.ClassId = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.ClassId
		varKubernetesDaemonSetStatus.ObjectType = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.ObjectType
		varKubernetesDaemonSetStatus.CurrentNumberScheduled = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.CurrentNumberScheduled
		varKubernetesDaemonSetStatus.DesiredNumberScheduled = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.DesiredNumberScheduled
		varKubernetesDaemonSetStatus.NumberAvailable = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.NumberAvailable
		varKubernetesDaemonSetStatus.NumberMisscheduled = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.NumberMisscheduled
		varKubernetesDaemonSetStatus.NumberReady = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.NumberReady
		varKubernetesDaemonSetStatus.ObservedGeneration = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.ObservedGeneration
		varKubernetesDaemonSetStatus.UpdatedNumberScheduled = varKubernetesDaemonSetStatusWithoutEmbeddedStruct.UpdatedNumberScheduled
		*o = KubernetesDaemonSetStatus(varKubernetesDaemonSetStatus)
	} else {
		return err
	}

	varKubernetesDaemonSetStatus := _KubernetesDaemonSetStatus{}

	err = json.Unmarshal(data, &varKubernetesDaemonSetStatus)
	if err == nil {
		o.MoBaseComplexType = varKubernetesDaemonSetStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentNumberScheduled")
		delete(additionalProperties, "DesiredNumberScheduled")
		delete(additionalProperties, "NumberAvailable")
		delete(additionalProperties, "NumberMisscheduled")
		delete(additionalProperties, "NumberReady")
		delete(additionalProperties, "ObservedGeneration")
		delete(additionalProperties, "UpdatedNumberScheduled")

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

type NullableKubernetesDaemonSetStatus struct {
	value *KubernetesDaemonSetStatus
	isSet bool
}

func (v NullableKubernetesDaemonSetStatus) Get() *KubernetesDaemonSetStatus {
	return v.value
}

func (v *NullableKubernetesDaemonSetStatus) Set(val *KubernetesDaemonSetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesDaemonSetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesDaemonSetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesDaemonSetStatus(val *KubernetesDaemonSetStatus) *NullableKubernetesDaemonSetStatus {
	return &NullableKubernetesDaemonSetStatus{value: val, isSet: true}
}

func (v NullableKubernetesDaemonSetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesDaemonSetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
