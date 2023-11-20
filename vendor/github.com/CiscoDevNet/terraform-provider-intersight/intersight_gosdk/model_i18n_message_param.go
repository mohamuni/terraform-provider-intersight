/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// I18nMessageParam A key, value pair used to substitute variables when translating a message in a particular locale for internationalization purpose.
type I18nMessageParam struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of a variable which is referenced in a i18n text template.
	Name *string `json:"Name,omitempty"`
	// The value of a variable which is substituted in a i18n text template.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _I18nMessageParam I18nMessageParam

// NewI18nMessageParam instantiates a new I18nMessageParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewI18nMessageParam(classId string, objectType string) *I18nMessageParam {
	this := I18nMessageParam{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewI18nMessageParamWithDefaults instantiates a new I18nMessageParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewI18nMessageParamWithDefaults() *I18nMessageParam {
	this := I18nMessageParam{}
	var classId string = "i18n.MessageParam"
	this.ClassId = classId
	var objectType string = "i18n.MessageParam"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *I18nMessageParam) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *I18nMessageParam) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *I18nMessageParam) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *I18nMessageParam) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *I18nMessageParam) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *I18nMessageParam) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *I18nMessageParam) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *I18nMessageParam) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *I18nMessageParam) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *I18nMessageParam) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *I18nMessageParam) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *I18nMessageParam) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *I18nMessageParam) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *I18nMessageParam) SetValue(v string) {
	o.Value = &v
}

func (o I18nMessageParam) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *I18nMessageParam) UnmarshalJSON(bytes []byte) (err error) {
	type I18nMessageParamWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of a variable which is referenced in a i18n text template.
		Name *string `json:"Name,omitempty"`
		// The value of a variable which is substituted in a i18n text template.
		Value *string `json:"Value,omitempty"`
	}

	varI18nMessageParamWithoutEmbeddedStruct := I18nMessageParamWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varI18nMessageParamWithoutEmbeddedStruct)
	if err == nil {
		varI18nMessageParam := _I18nMessageParam{}
		varI18nMessageParam.ClassId = varI18nMessageParamWithoutEmbeddedStruct.ClassId
		varI18nMessageParam.ObjectType = varI18nMessageParamWithoutEmbeddedStruct.ObjectType
		varI18nMessageParam.Name = varI18nMessageParamWithoutEmbeddedStruct.Name
		varI18nMessageParam.Value = varI18nMessageParamWithoutEmbeddedStruct.Value
		*o = I18nMessageParam(varI18nMessageParam)
	} else {
		return err
	}

	varI18nMessageParam := _I18nMessageParam{}

	err = json.Unmarshal(bytes, &varI18nMessageParam)
	if err == nil {
		o.MoBaseComplexType = varI18nMessageParam.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")

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

type NullableI18nMessageParam struct {
	value *I18nMessageParam
	isSet bool
}

func (v NullableI18nMessageParam) Get() *I18nMessageParam {
	return v.value
}

func (v *NullableI18nMessageParam) Set(val *I18nMessageParam) {
	v.value = val
	v.isSet = true
}

func (v NullableI18nMessageParam) IsSet() bool {
	return v.isSet
}

func (v *NullableI18nMessageParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableI18nMessageParam(val *I18nMessageParam) *NullableI18nMessageParam {
	return &NullableI18nMessageParam{value: val, isSet: true}
}

func (v NullableI18nMessageParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableI18nMessageParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
