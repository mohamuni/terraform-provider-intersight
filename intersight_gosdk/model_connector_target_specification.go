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

// checks if the ConnectorTargetSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorTargetSpecification{}

// ConnectorTargetSpecification Target deployment specification details to deploy target as a kube service. This allows Intersight to define supported docker image tag and customize resources for each target deployment.
type ConnectorTargetSpecification struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CPU limit assigned to the docker container. It is total amount of CPU time that a container can use every 100ms. A container cannot use more than its share of CPU time during this interval.
	CpuLimit *string `json:"CpuLimit,omitempty"`
	// Requested CPU value for a docker container to run in Intersight Assist.
	CpuRequest *string `json:"CpuRequest,omitempty"`
	// Docker image tag used to define kubernetes deployment for each target. Image tag should be the complete URL. This image can be found locally in case of Intersight Appliance or can be pulled from Intersight cloud in Intersight Assist deployment.
	ImageTag *string `json:"ImageTag,omitempty"`
	// Intersight Assist prevents the docker container from using more than the configured memory limit. If a Container exceeds its memory limit, it might be terminated. If it is restartable, the kubelet will restart it, as with any other type of runtime failure.
	MemoryLimit *string `json:"MemoryLimit,omitempty"`
	// Requested memory value for a docker container to run in Intersight Assist.
	MemoryRequest        *string `json:"MemoryRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorTargetSpecification ConnectorTargetSpecification

// NewConnectorTargetSpecification instantiates a new ConnectorTargetSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTargetSpecification(classId string, objectType string) *ConnectorTargetSpecification {
	this := ConnectorTargetSpecification{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorTargetSpecificationWithDefaults instantiates a new ConnectorTargetSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTargetSpecificationWithDefaults() *ConnectorTargetSpecification {
	this := ConnectorTargetSpecification{}
	var classId string = "connector.TargetSpecification"
	this.ClassId = classId
	var objectType string = "connector.TargetSpecification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorTargetSpecification) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorTargetSpecification) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "connector.TargetSpecification" of the ClassId field.
func (o *ConnectorTargetSpecification) GetDefaultClassId() interface{} {
	return "connector.TargetSpecification"
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorTargetSpecification) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorTargetSpecification) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "connector.TargetSpecification" of the ObjectType field.
func (o *ConnectorTargetSpecification) GetDefaultObjectType() interface{} {
	return "connector.TargetSpecification"
}

// GetCpuLimit returns the CpuLimit field value if set, zero value otherwise.
func (o *ConnectorTargetSpecification) GetCpuLimit() string {
	if o == nil || IsNil(o.CpuLimit) {
		var ret string
		return ret
	}
	return *o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetCpuLimitOk() (*string, bool) {
	if o == nil || IsNil(o.CpuLimit) {
		return nil, false
	}
	return o.CpuLimit, true
}

// HasCpuLimit returns a boolean if a field has been set.
func (o *ConnectorTargetSpecification) HasCpuLimit() bool {
	if o != nil && !IsNil(o.CpuLimit) {
		return true
	}

	return false
}

// SetCpuLimit gets a reference to the given string and assigns it to the CpuLimit field.
func (o *ConnectorTargetSpecification) SetCpuLimit(v string) {
	o.CpuLimit = &v
}

// GetCpuRequest returns the CpuRequest field value if set, zero value otherwise.
func (o *ConnectorTargetSpecification) GetCpuRequest() string {
	if o == nil || IsNil(o.CpuRequest) {
		var ret string
		return ret
	}
	return *o.CpuRequest
}

// GetCpuRequestOk returns a tuple with the CpuRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetCpuRequestOk() (*string, bool) {
	if o == nil || IsNil(o.CpuRequest) {
		return nil, false
	}
	return o.CpuRequest, true
}

// HasCpuRequest returns a boolean if a field has been set.
func (o *ConnectorTargetSpecification) HasCpuRequest() bool {
	if o != nil && !IsNil(o.CpuRequest) {
		return true
	}

	return false
}

// SetCpuRequest gets a reference to the given string and assigns it to the CpuRequest field.
func (o *ConnectorTargetSpecification) SetCpuRequest(v string) {
	o.CpuRequest = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *ConnectorTargetSpecification) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag) {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.ImageTag) {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *ConnectorTargetSpecification) HasImageTag() bool {
	if o != nil && !IsNil(o.ImageTag) {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *ConnectorTargetSpecification) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetMemoryLimit returns the MemoryLimit field value if set, zero value otherwise.
func (o *ConnectorTargetSpecification) GetMemoryLimit() string {
	if o == nil || IsNil(o.MemoryLimit) {
		var ret string
		return ret
	}
	return *o.MemoryLimit
}

// GetMemoryLimitOk returns a tuple with the MemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetMemoryLimitOk() (*string, bool) {
	if o == nil || IsNil(o.MemoryLimit) {
		return nil, false
	}
	return o.MemoryLimit, true
}

// HasMemoryLimit returns a boolean if a field has been set.
func (o *ConnectorTargetSpecification) HasMemoryLimit() bool {
	if o != nil && !IsNil(o.MemoryLimit) {
		return true
	}

	return false
}

// SetMemoryLimit gets a reference to the given string and assigns it to the MemoryLimit field.
func (o *ConnectorTargetSpecification) SetMemoryLimit(v string) {
	o.MemoryLimit = &v
}

// GetMemoryRequest returns the MemoryRequest field value if set, zero value otherwise.
func (o *ConnectorTargetSpecification) GetMemoryRequest() string {
	if o == nil || IsNil(o.MemoryRequest) {
		var ret string
		return ret
	}
	return *o.MemoryRequest
}

// GetMemoryRequestOk returns a tuple with the MemoryRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetSpecification) GetMemoryRequestOk() (*string, bool) {
	if o == nil || IsNil(o.MemoryRequest) {
		return nil, false
	}
	return o.MemoryRequest, true
}

// HasMemoryRequest returns a boolean if a field has been set.
func (o *ConnectorTargetSpecification) HasMemoryRequest() bool {
	if o != nil && !IsNil(o.MemoryRequest) {
		return true
	}

	return false
}

// SetMemoryRequest gets a reference to the given string and assigns it to the MemoryRequest field.
func (o *ConnectorTargetSpecification) SetMemoryRequest(v string) {
	o.MemoryRequest = &v
}

func (o ConnectorTargetSpecification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorTargetSpecification) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CpuLimit) {
		toSerialize["CpuLimit"] = o.CpuLimit
	}
	if !IsNil(o.CpuRequest) {
		toSerialize["CpuRequest"] = o.CpuRequest
	}
	if !IsNil(o.ImageTag) {
		toSerialize["ImageTag"] = o.ImageTag
	}
	if !IsNil(o.MemoryLimit) {
		toSerialize["MemoryLimit"] = o.MemoryLimit
	}
	if !IsNil(o.MemoryRequest) {
		toSerialize["MemoryRequest"] = o.MemoryRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorTargetSpecification) UnmarshalJSON(data []byte) (err error) {
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
	type ConnectorTargetSpecificationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// CPU limit assigned to the docker container. It is total amount of CPU time that a container can use every 100ms. A container cannot use more than its share of CPU time during this interval.
		CpuLimit *string `json:"CpuLimit,omitempty"`
		// Requested CPU value for a docker container to run in Intersight Assist.
		CpuRequest *string `json:"CpuRequest,omitempty"`
		// Docker image tag used to define kubernetes deployment for each target. Image tag should be the complete URL. This image can be found locally in case of Intersight Appliance or can be pulled from Intersight cloud in Intersight Assist deployment.
		ImageTag *string `json:"ImageTag,omitempty"`
		// Intersight Assist prevents the docker container from using more than the configured memory limit. If a Container exceeds its memory limit, it might be terminated. If it is restartable, the kubelet will restart it, as with any other type of runtime failure.
		MemoryLimit *string `json:"MemoryLimit,omitempty"`
		// Requested memory value for a docker container to run in Intersight Assist.
		MemoryRequest *string `json:"MemoryRequest,omitempty"`
	}

	varConnectorTargetSpecificationWithoutEmbeddedStruct := ConnectorTargetSpecificationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConnectorTargetSpecificationWithoutEmbeddedStruct)
	if err == nil {
		varConnectorTargetSpecification := _ConnectorTargetSpecification{}
		varConnectorTargetSpecification.ClassId = varConnectorTargetSpecificationWithoutEmbeddedStruct.ClassId
		varConnectorTargetSpecification.ObjectType = varConnectorTargetSpecificationWithoutEmbeddedStruct.ObjectType
		varConnectorTargetSpecification.CpuLimit = varConnectorTargetSpecificationWithoutEmbeddedStruct.CpuLimit
		varConnectorTargetSpecification.CpuRequest = varConnectorTargetSpecificationWithoutEmbeddedStruct.CpuRequest
		varConnectorTargetSpecification.ImageTag = varConnectorTargetSpecificationWithoutEmbeddedStruct.ImageTag
		varConnectorTargetSpecification.MemoryLimit = varConnectorTargetSpecificationWithoutEmbeddedStruct.MemoryLimit
		varConnectorTargetSpecification.MemoryRequest = varConnectorTargetSpecificationWithoutEmbeddedStruct.MemoryRequest
		*o = ConnectorTargetSpecification(varConnectorTargetSpecification)
	} else {
		return err
	}

	varConnectorTargetSpecification := _ConnectorTargetSpecification{}

	err = json.Unmarshal(data, &varConnectorTargetSpecification)
	if err == nil {
		o.MoBaseComplexType = varConnectorTargetSpecification.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuLimit")
		delete(additionalProperties, "CpuRequest")
		delete(additionalProperties, "ImageTag")
		delete(additionalProperties, "MemoryLimit")
		delete(additionalProperties, "MemoryRequest")

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

type NullableConnectorTargetSpecification struct {
	value *ConnectorTargetSpecification
	isSet bool
}

func (v NullableConnectorTargetSpecification) Get() *ConnectorTargetSpecification {
	return v.value
}

func (v *NullableConnectorTargetSpecification) Set(val *ConnectorTargetSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTargetSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTargetSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTargetSpecification(val *ConnectorTargetSpecification) *NullableConnectorTargetSpecification {
	return &NullableConnectorTargetSpecification{value: val, isSet: true}
}

func (v NullableConnectorTargetSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTargetSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
