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

// checks if the CapabilityUpdateOrderListType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityUpdateOrderListType{}

// CapabilityUpdateOrderListType The list of update orders.
type CapabilityUpdateOrderListType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The interim version to be updated to in the update order.
	InterimVersion *string `json:"InterimVersion,omitempty"`
	// The source version of the update order.
	SourceVersion   *string  `json:"SourceVersion,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The target version of the update order.
	TargetVersion        *string `json:"TargetVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityUpdateOrderListType CapabilityUpdateOrderListType

// NewCapabilityUpdateOrderListType instantiates a new CapabilityUpdateOrderListType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityUpdateOrderListType(classId string, objectType string) *CapabilityUpdateOrderListType {
	this := CapabilityUpdateOrderListType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityUpdateOrderListTypeWithDefaults instantiates a new CapabilityUpdateOrderListType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityUpdateOrderListTypeWithDefaults() *CapabilityUpdateOrderListType {
	this := CapabilityUpdateOrderListType{}
	var classId string = "capability.UpdateOrderListType"
	this.ClassId = classId
	var objectType string = "capability.UpdateOrderListType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityUpdateOrderListType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityUpdateOrderListType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityUpdateOrderListType) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.UpdateOrderListType" of the ClassId field.
func (o *CapabilityUpdateOrderListType) GetDefaultClassId() interface{} {
	return "capability.UpdateOrderListType"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityUpdateOrderListType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityUpdateOrderListType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityUpdateOrderListType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.UpdateOrderListType" of the ObjectType field.
func (o *CapabilityUpdateOrderListType) GetDefaultObjectType() interface{} {
	return "capability.UpdateOrderListType"
}

// GetInterimVersion returns the InterimVersion field value if set, zero value otherwise.
func (o *CapabilityUpdateOrderListType) GetInterimVersion() string {
	if o == nil || IsNil(o.InterimVersion) {
		var ret string
		return ret
	}
	return *o.InterimVersion
}

// GetInterimVersionOk returns a tuple with the InterimVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityUpdateOrderListType) GetInterimVersionOk() (*string, bool) {
	if o == nil || IsNil(o.InterimVersion) {
		return nil, false
	}
	return o.InterimVersion, true
}

// HasInterimVersion returns a boolean if a field has been set.
func (o *CapabilityUpdateOrderListType) HasInterimVersion() bool {
	if o != nil && !IsNil(o.InterimVersion) {
		return true
	}

	return false
}

// SetInterimVersion gets a reference to the given string and assigns it to the InterimVersion field.
func (o *CapabilityUpdateOrderListType) SetInterimVersion(v string) {
	o.InterimVersion = &v
}

// GetSourceVersion returns the SourceVersion field value if set, zero value otherwise.
func (o *CapabilityUpdateOrderListType) GetSourceVersion() string {
	if o == nil || IsNil(o.SourceVersion) {
		var ret string
		return ret
	}
	return *o.SourceVersion
}

// GetSourceVersionOk returns a tuple with the SourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityUpdateOrderListType) GetSourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SourceVersion) {
		return nil, false
	}
	return o.SourceVersion, true
}

// HasSourceVersion returns a boolean if a field has been set.
func (o *CapabilityUpdateOrderListType) HasSourceVersion() bool {
	if o != nil && !IsNil(o.SourceVersion) {
		return true
	}

	return false
}

// SetSourceVersion gets a reference to the given string and assigns it to the SourceVersion field.
func (o *CapabilityUpdateOrderListType) SetSourceVersion(v string) {
	o.SourceVersion = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityUpdateOrderListType) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityUpdateOrderListType) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *CapabilityUpdateOrderListType) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *CapabilityUpdateOrderListType) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *CapabilityUpdateOrderListType) GetTargetVersion() string {
	if o == nil || IsNil(o.TargetVersion) {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityUpdateOrderListType) GetTargetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetVersion) {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *CapabilityUpdateOrderListType) HasTargetVersion() bool {
	if o != nil && !IsNil(o.TargetVersion) {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *CapabilityUpdateOrderListType) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

func (o CapabilityUpdateOrderListType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityUpdateOrderListType) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.InterimVersion) {
		toSerialize["InterimVersion"] = o.InterimVersion
	}
	if !IsNil(o.SourceVersion) {
		toSerialize["SourceVersion"] = o.SourceVersion
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if !IsNil(o.TargetVersion) {
		toSerialize["TargetVersion"] = o.TargetVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityUpdateOrderListType) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityUpdateOrderListTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The interim version to be updated to in the update order.
		InterimVersion *string `json:"InterimVersion,omitempty"`
		// The source version of the update order.
		SourceVersion   *string  `json:"SourceVersion,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
		// The target version of the update order.
		TargetVersion *string `json:"TargetVersion,omitempty"`
	}

	varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct := CapabilityUpdateOrderListTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityUpdateOrderListType := _CapabilityUpdateOrderListType{}
		varCapabilityUpdateOrderListType.ClassId = varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct.ClassId
		varCapabilityUpdateOrderListType.ObjectType = varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct.ObjectType
		varCapabilityUpdateOrderListType.InterimVersion = varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct.InterimVersion
		varCapabilityUpdateOrderListType.SourceVersion = varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct.SourceVersion
		varCapabilityUpdateOrderListType.SupportedModels = varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct.SupportedModels
		varCapabilityUpdateOrderListType.TargetVersion = varCapabilityUpdateOrderListTypeWithoutEmbeddedStruct.TargetVersion
		*o = CapabilityUpdateOrderListType(varCapabilityUpdateOrderListType)
	} else {
		return err
	}

	varCapabilityUpdateOrderListType := _CapabilityUpdateOrderListType{}

	err = json.Unmarshal(data, &varCapabilityUpdateOrderListType)
	if err == nil {
		o.MoBaseComplexType = varCapabilityUpdateOrderListType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterimVersion")
		delete(additionalProperties, "SourceVersion")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "TargetVersion")

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

type NullableCapabilityUpdateOrderListType struct {
	value *CapabilityUpdateOrderListType
	isSet bool
}

func (v NullableCapabilityUpdateOrderListType) Get() *CapabilityUpdateOrderListType {
	return v.value
}

func (v *NullableCapabilityUpdateOrderListType) Set(val *CapabilityUpdateOrderListType) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityUpdateOrderListType) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityUpdateOrderListType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityUpdateOrderListType(val *CapabilityUpdateOrderListType) *NullableCapabilityUpdateOrderListType {
	return &NullableCapabilityUpdateOrderListType{value: val, isSet: true}
}

func (v NullableCapabilityUpdateOrderListType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityUpdateOrderListType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
