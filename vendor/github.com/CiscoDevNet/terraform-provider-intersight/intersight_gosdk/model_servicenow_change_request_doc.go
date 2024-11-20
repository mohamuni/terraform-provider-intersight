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

// checks if the ServicenowChangeRequestDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicenowChangeRequestDoc{}

// ServicenowChangeRequestDoc Change Request schema document on ServiceNow.
type ServicenowChangeRequestDoc struct {
	ServicenowInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Example value for Change request property.
	ExampleValue *string `json:"ExampleValue,omitempty"`
	// Internal type for Change request property.
	InternalType *string `json:"InternalType,omitempty"`
	// Label for Change request property.
	Label *string `json:"Label,omitempty"`
	// Name for Change request property.
	PropertyName         *string `json:"PropertyName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicenowChangeRequestDoc ServicenowChangeRequestDoc

// NewServicenowChangeRequestDoc instantiates a new ServicenowChangeRequestDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicenowChangeRequestDoc(classId string, objectType string) *ServicenowChangeRequestDoc {
	this := ServicenowChangeRequestDoc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServicenowChangeRequestDocWithDefaults instantiates a new ServicenowChangeRequestDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicenowChangeRequestDocWithDefaults() *ServicenowChangeRequestDoc {
	this := ServicenowChangeRequestDoc{}
	var classId string = "servicenow.ChangeRequestDoc"
	this.ClassId = classId
	var objectType string = "servicenow.ChangeRequestDoc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServicenowChangeRequestDoc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequestDoc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServicenowChangeRequestDoc) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "servicenow.ChangeRequestDoc" of the ClassId field.
func (o *ServicenowChangeRequestDoc) GetDefaultClassId() interface{} {
	return "servicenow.ChangeRequestDoc"
}

// GetObjectType returns the ObjectType field value
func (o *ServicenowChangeRequestDoc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequestDoc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServicenowChangeRequestDoc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "servicenow.ChangeRequestDoc" of the ObjectType field.
func (o *ServicenowChangeRequestDoc) GetDefaultObjectType() interface{} {
	return "servicenow.ChangeRequestDoc"
}

// GetExampleValue returns the ExampleValue field value if set, zero value otherwise.
func (o *ServicenowChangeRequestDoc) GetExampleValue() string {
	if o == nil || IsNil(o.ExampleValue) {
		var ret string
		return ret
	}
	return *o.ExampleValue
}

// GetExampleValueOk returns a tuple with the ExampleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequestDoc) GetExampleValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExampleValue) {
		return nil, false
	}
	return o.ExampleValue, true
}

// HasExampleValue returns a boolean if a field has been set.
func (o *ServicenowChangeRequestDoc) HasExampleValue() bool {
	if o != nil && !IsNil(o.ExampleValue) {
		return true
	}

	return false
}

// SetExampleValue gets a reference to the given string and assigns it to the ExampleValue field.
func (o *ServicenowChangeRequestDoc) SetExampleValue(v string) {
	o.ExampleValue = &v
}

// GetInternalType returns the InternalType field value if set, zero value otherwise.
func (o *ServicenowChangeRequestDoc) GetInternalType() string {
	if o == nil || IsNil(o.InternalType) {
		var ret string
		return ret
	}
	return *o.InternalType
}

// GetInternalTypeOk returns a tuple with the InternalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequestDoc) GetInternalTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InternalType) {
		return nil, false
	}
	return o.InternalType, true
}

// HasInternalType returns a boolean if a field has been set.
func (o *ServicenowChangeRequestDoc) HasInternalType() bool {
	if o != nil && !IsNil(o.InternalType) {
		return true
	}

	return false
}

// SetInternalType gets a reference to the given string and assigns it to the InternalType field.
func (o *ServicenowChangeRequestDoc) SetInternalType(v string) {
	o.InternalType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServicenowChangeRequestDoc) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequestDoc) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServicenowChangeRequestDoc) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServicenowChangeRequestDoc) SetLabel(v string) {
	o.Label = &v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *ServicenowChangeRequestDoc) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequestDoc) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *ServicenowChangeRequestDoc) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *ServicenowChangeRequestDoc) SetPropertyName(v string) {
	o.PropertyName = &v
}

func (o ServicenowChangeRequestDoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicenowChangeRequestDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedServicenowInventoryEntity, errServicenowInventoryEntity := json.Marshal(o.ServicenowInventoryEntity)
	if errServicenowInventoryEntity != nil {
		return map[string]interface{}{}, errServicenowInventoryEntity
	}
	errServicenowInventoryEntity = json.Unmarshal([]byte(serializedServicenowInventoryEntity), &toSerialize)
	if errServicenowInventoryEntity != nil {
		return map[string]interface{}{}, errServicenowInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ExampleValue) {
		toSerialize["ExampleValue"] = o.ExampleValue
	}
	if !IsNil(o.InternalType) {
		toSerialize["InternalType"] = o.InternalType
	}
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.PropertyName) {
		toSerialize["PropertyName"] = o.PropertyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicenowChangeRequestDoc) UnmarshalJSON(data []byte) (err error) {
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
	type ServicenowChangeRequestDocWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Example value for Change request property.
		ExampleValue *string `json:"ExampleValue,omitempty"`
		// Internal type for Change request property.
		InternalType *string `json:"InternalType,omitempty"`
		// Label for Change request property.
		Label *string `json:"Label,omitempty"`
		// Name for Change request property.
		PropertyName *string `json:"PropertyName,omitempty"`
	}

	varServicenowChangeRequestDocWithoutEmbeddedStruct := ServicenowChangeRequestDocWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServicenowChangeRequestDocWithoutEmbeddedStruct)
	if err == nil {
		varServicenowChangeRequestDoc := _ServicenowChangeRequestDoc{}
		varServicenowChangeRequestDoc.ClassId = varServicenowChangeRequestDocWithoutEmbeddedStruct.ClassId
		varServicenowChangeRequestDoc.ObjectType = varServicenowChangeRequestDocWithoutEmbeddedStruct.ObjectType
		varServicenowChangeRequestDoc.ExampleValue = varServicenowChangeRequestDocWithoutEmbeddedStruct.ExampleValue
		varServicenowChangeRequestDoc.InternalType = varServicenowChangeRequestDocWithoutEmbeddedStruct.InternalType
		varServicenowChangeRequestDoc.Label = varServicenowChangeRequestDocWithoutEmbeddedStruct.Label
		varServicenowChangeRequestDoc.PropertyName = varServicenowChangeRequestDocWithoutEmbeddedStruct.PropertyName
		*o = ServicenowChangeRequestDoc(varServicenowChangeRequestDoc)
	} else {
		return err
	}

	varServicenowChangeRequestDoc := _ServicenowChangeRequestDoc{}

	err = json.Unmarshal(data, &varServicenowChangeRequestDoc)
	if err == nil {
		o.ServicenowInventoryEntity = varServicenowChangeRequestDoc.ServicenowInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExampleValue")
		delete(additionalProperties, "InternalType")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "PropertyName")

		// remove fields from embedded structs
		reflectServicenowInventoryEntity := reflect.ValueOf(o.ServicenowInventoryEntity)
		for i := 0; i < reflectServicenowInventoryEntity.Type().NumField(); i++ {
			t := reflectServicenowInventoryEntity.Type().Field(i)

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

type NullableServicenowChangeRequestDoc struct {
	value *ServicenowChangeRequestDoc
	isSet bool
}

func (v NullableServicenowChangeRequestDoc) Get() *ServicenowChangeRequestDoc {
	return v.value
}

func (v *NullableServicenowChangeRequestDoc) Set(val *ServicenowChangeRequestDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableServicenowChangeRequestDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableServicenowChangeRequestDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicenowChangeRequestDoc(val *ServicenowChangeRequestDoc) *NullableServicenowChangeRequestDoc {
	return &NullableServicenowChangeRequestDoc{value: val, isSet: true}
}

func (v NullableServicenowChangeRequestDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicenowChangeRequestDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
