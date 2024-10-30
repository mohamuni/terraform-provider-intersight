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

// checks if the IamServiceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamServiceProvider{}

// IamServiceProvider SAML Service provider related information in Intersight.
type IamServiceProvider struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
	EntityId *string `json:"EntityId,omitempty"`
	// Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication.
	Metadata *string `json:"Metadata,omitempty"`
	// Name of the Intersight Service Provider.
	Name                 *string                       `json:"Name,omitempty"`
	System               NullableIamSystemRelationship `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamServiceProvider IamServiceProvider

// NewIamServiceProvider instantiates a new IamServiceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamServiceProvider(classId string, objectType string) *IamServiceProvider {
	this := IamServiceProvider{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamServiceProviderWithDefaults instantiates a new IamServiceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamServiceProviderWithDefaults() *IamServiceProvider {
	this := IamServiceProvider{}
	var classId string = "iam.ServiceProvider"
	this.ClassId = classId
	var objectType string = "iam.ServiceProvider"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamServiceProvider) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamServiceProvider) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.ServiceProvider" of the ClassId field.
func (o *IamServiceProvider) GetDefaultClassId() interface{} {
	return "iam.ServiceProvider"
}

// GetObjectType returns the ObjectType field value
func (o *IamServiceProvider) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamServiceProvider) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.ServiceProvider" of the ObjectType field.
func (o *IamServiceProvider) GetDefaultObjectType() interface{} {
	return "iam.ServiceProvider"
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *IamServiceProvider) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *IamServiceProvider) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *IamServiceProvider) SetEntityId(v string) {
	o.EntityId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamServiceProvider) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamServiceProvider) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *IamServiceProvider) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamServiceProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamServiceProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamServiceProvider) SetName(v string) {
	o.Name = &v
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamServiceProvider) GetSystem() IamSystemRelationship {
	if o == nil || IsNil(o.System.Get()) {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System.Get()
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamServiceProvider) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.System.Get(), o.System.IsSet()
}

// HasSystem returns a boolean if a field has been set.
func (o *IamServiceProvider) HasSystem() bool {
	if o != nil && o.System.IsSet() {
		return true
	}

	return false
}

// SetSystem gets a reference to the given NullableIamSystemRelationship and assigns it to the System field.
func (o *IamServiceProvider) SetSystem(v IamSystemRelationship) {
	o.System.Set(&v)
}

// SetSystemNil sets the value for System to be an explicit nil
func (o *IamServiceProvider) SetSystemNil() {
	o.System.Set(nil)
}

// UnsetSystem ensures that no value is present for System, not even an explicit nil
func (o *IamServiceProvider) UnsetSystem() {
	o.System.Unset()
}

func (o IamServiceProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamServiceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.EntityId) {
		toSerialize["EntityId"] = o.EntityId
	}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.System.IsSet() {
		toSerialize["System"] = o.System.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamServiceProvider) UnmarshalJSON(data []byte) (err error) {
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
	type IamServiceProviderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
		EntityId *string `json:"EntityId,omitempty"`
		// Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication.
		Metadata *string `json:"Metadata,omitempty"`
		// Name of the Intersight Service Provider.
		Name   *string                       `json:"Name,omitempty"`
		System NullableIamSystemRelationship `json:"System,omitempty"`
	}

	varIamServiceProviderWithoutEmbeddedStruct := IamServiceProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamServiceProviderWithoutEmbeddedStruct)
	if err == nil {
		varIamServiceProvider := _IamServiceProvider{}
		varIamServiceProvider.ClassId = varIamServiceProviderWithoutEmbeddedStruct.ClassId
		varIamServiceProvider.ObjectType = varIamServiceProviderWithoutEmbeddedStruct.ObjectType
		varIamServiceProvider.EntityId = varIamServiceProviderWithoutEmbeddedStruct.EntityId
		varIamServiceProvider.Metadata = varIamServiceProviderWithoutEmbeddedStruct.Metadata
		varIamServiceProvider.Name = varIamServiceProviderWithoutEmbeddedStruct.Name
		varIamServiceProvider.System = varIamServiceProviderWithoutEmbeddedStruct.System
		*o = IamServiceProvider(varIamServiceProvider)
	} else {
		return err
	}

	varIamServiceProvider := _IamServiceProvider{}

	err = json.Unmarshal(data, &varIamServiceProvider)
	if err == nil {
		o.MoBaseMo = varIamServiceProvider.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EntityId")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "System")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamServiceProvider struct {
	value *IamServiceProvider
	isSet bool
}

func (v NullableIamServiceProvider) Get() *IamServiceProvider {
	return v.value
}

func (v *NullableIamServiceProvider) Set(val *IamServiceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIamServiceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIamServiceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamServiceProvider(val *IamServiceProvider) *NullableIamServiceProvider {
	return &NullableIamServiceProvider{value: val, isSet: true}
}

func (v NullableIamServiceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamServiceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
