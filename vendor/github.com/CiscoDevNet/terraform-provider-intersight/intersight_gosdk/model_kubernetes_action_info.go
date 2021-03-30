/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// KubernetesActionInfo Action details like name of the action performed, status, failure reason, etc.
type KubernetesActionInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of failure i.e. derived from the workflow failure message.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Name of the Action performed on a resource like VM, Disk etc.
	Name *string `json:"Name,omitempty"`
	// Status of the Action like InProgress, Success, Failure etc. * `None` - A place holder for the default value. * `InProgress` - Action triggered on the resource is still running. * `Success` - Action triggered on the resource is completed successfully. * `Failure` - Action triggered on the resource is failed.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesActionInfo KubernetesActionInfo

// NewKubernetesActionInfo instantiates a new KubernetesActionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesActionInfo(classId string, objectType string) *KubernetesActionInfo {
	this := KubernetesActionInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "None"
	this.Status = &status
	return &this
}

// NewKubernetesActionInfoWithDefaults instantiates a new KubernetesActionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesActionInfoWithDefaults() *KubernetesActionInfo {
	this := KubernetesActionInfo{}
	var classId string = "kubernetes.ActionInfo"
	this.ClassId = classId
	var objectType string = "kubernetes.ActionInfo"
	this.ObjectType = objectType
	var status string = "None"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesActionInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesActionInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesActionInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesActionInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesActionInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesActionInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *KubernetesActionInfo) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesActionInfo) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *KubernetesActionInfo) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *KubernetesActionInfo) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesActionInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesActionInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesActionInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesActionInfo) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KubernetesActionInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesActionInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesActionInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KubernetesActionInfo) SetStatus(v string) {
	o.Status = &v
}

func (o KubernetesActionInfo) MarshalJSON() ([]byte, error) {
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
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesActionInfo) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesActionInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of failure i.e. derived from the workflow failure message.
		FailureReason *string `json:"FailureReason,omitempty"`
		// Name of the Action performed on a resource like VM, Disk etc.
		Name *string `json:"Name,omitempty"`
		// Status of the Action like InProgress, Success, Failure etc. * `None` - A place holder for the default value. * `InProgress` - Action triggered on the resource is still running. * `Success` - Action triggered on the resource is completed successfully. * `Failure` - Action triggered on the resource is failed.
		Status *string `json:"Status,omitempty"`
	}

	varKubernetesActionInfoWithoutEmbeddedStruct := KubernetesActionInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesActionInfoWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesActionInfo := _KubernetesActionInfo{}
		varKubernetesActionInfo.ClassId = varKubernetesActionInfoWithoutEmbeddedStruct.ClassId
		varKubernetesActionInfo.ObjectType = varKubernetesActionInfoWithoutEmbeddedStruct.ObjectType
		varKubernetesActionInfo.FailureReason = varKubernetesActionInfoWithoutEmbeddedStruct.FailureReason
		varKubernetesActionInfo.Name = varKubernetesActionInfoWithoutEmbeddedStruct.Name
		varKubernetesActionInfo.Status = varKubernetesActionInfoWithoutEmbeddedStruct.Status
		*o = KubernetesActionInfo(varKubernetesActionInfo)
	} else {
		return err
	}

	varKubernetesActionInfo := _KubernetesActionInfo{}

	err = json.Unmarshal(bytes, &varKubernetesActionInfo)
	if err == nil {
		o.MoBaseComplexType = varKubernetesActionInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")

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

type NullableKubernetesActionInfo struct {
	value *KubernetesActionInfo
	isSet bool
}

func (v NullableKubernetesActionInfo) Get() *KubernetesActionInfo {
	return v.value
}

func (v *NullableKubernetesActionInfo) Set(val *KubernetesActionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesActionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesActionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesActionInfo(val *KubernetesActionInfo) *NullableKubernetesActionInfo {
	return &NullableKubernetesActionInfo{value: val, isSet: true}
}

func (v NullableKubernetesActionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesActionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
