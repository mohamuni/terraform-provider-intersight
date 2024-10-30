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

// checks if the IssueMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueMessage{}

// IssueMessage A message which is optionally parameterized and which supports language localization.
type IssueMessage struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A parameterized message such as \"The temperature is %[1]v.\" where 'hot' could be substituted for %[1]v.
	Message              *string  `json:"Message,omitempty"`
	Parameters           []string `json:"Parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IssueMessage IssueMessage

// NewIssueMessage instantiates a new IssueMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueMessage(classId string, objectType string) *IssueMessage {
	this := IssueMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIssueMessageWithDefaults instantiates a new IssueMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueMessageWithDefaults() *IssueMessage {
	this := IssueMessage{}
	var classId string = "issue.Message"
	this.ClassId = classId
	var objectType string = "issue.Message"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IssueMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IssueMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IssueMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "issue.Message" of the ClassId field.
func (o *IssueMessage) GetDefaultClassId() interface{} {
	return "issue.Message"
}

// GetObjectType returns the ObjectType field value
func (o *IssueMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IssueMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IssueMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "issue.Message" of the ObjectType field.
func (o *IssueMessage) GetDefaultObjectType() interface{} {
	return "issue.Message"
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IssueMessage) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueMessage) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IssueMessage) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IssueMessage) SetMessage(v string) {
	o.Message = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueMessage) GetParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueMessage) GetParametersOk() ([]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IssueMessage) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []string and assigns it to the Parameters field.
func (o *IssueMessage) SetParameters(v []string) {
	o.Parameters = v
}

func (o IssueMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueMessage) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IssueMessage) UnmarshalJSON(data []byte) (err error) {
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
	type IssueMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A parameterized message such as \"The temperature is %[1]v.\" where 'hot' could be substituted for %[1]v.
		Message    *string  `json:"Message,omitempty"`
		Parameters []string `json:"Parameters,omitempty"`
	}

	varIssueMessageWithoutEmbeddedStruct := IssueMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIssueMessageWithoutEmbeddedStruct)
	if err == nil {
		varIssueMessage := _IssueMessage{}
		varIssueMessage.ClassId = varIssueMessageWithoutEmbeddedStruct.ClassId
		varIssueMessage.ObjectType = varIssueMessageWithoutEmbeddedStruct.ObjectType
		varIssueMessage.Message = varIssueMessageWithoutEmbeddedStruct.Message
		varIssueMessage.Parameters = varIssueMessageWithoutEmbeddedStruct.Parameters
		*o = IssueMessage(varIssueMessage)
	} else {
		return err
	}

	varIssueMessage := _IssueMessage{}

	err = json.Unmarshal(data, &varIssueMessage)
	if err == nil {
		o.MoBaseComplexType = varIssueMessage.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Parameters")

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

type NullableIssueMessage struct {
	value *IssueMessage
	isSet bool
}

func (v NullableIssueMessage) Get() *IssueMessage {
	return v.value
}

func (v *NullableIssueMessage) Set(val *IssueMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueMessage(val *IssueMessage) *NullableIssueMessage {
	return &NullableIssueMessage{value: val, isSet: true}
}

func (v NullableIssueMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
