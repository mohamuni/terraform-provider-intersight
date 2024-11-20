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
	"time"
)

// checks if the ChangelogItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangelogItem{}

// ChangelogItem An API contract changelog item. It represents an item of contract changes between the version indicated by the attribute semanticVersion and the previous version.
type ChangelogItem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The date version for the API contract changelog item in the format rfc3339 with no fraction seconds set.  Note that there can be more than one item per DateVersion. Example: 2023-12-19T00:00:00Z .
	DateVersion *time.Time `json:"DateVersion,omitempty"`
	// The operationId of the endpoint for which changelog item is being generated.
	Entity *string `json:"Entity,omitempty"`
	// The semantic version for the API contract changelog item. Note that there can be more than one item per SemanticVersion.
	SemanticVersion *string `json:"SemanticVersion,omitempty" validate:"regexp=^(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)(?:-((?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\\\.(?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\\\+([0-9a-zA-Z-]+(?:\\\\.[0-9a-zA-Z-]+)*))?$"`
	// The value of the API contract changelog item.
	Value                *string                             `json:"Value,omitempty"`
	Catalog              NullableWorkflowCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangelogItem ChangelogItem

// NewChangelogItem instantiates a new ChangelogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangelogItem(classId string, objectType string) *ChangelogItem {
	this := ChangelogItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChangelogItemWithDefaults instantiates a new ChangelogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangelogItemWithDefaults() *ChangelogItem {
	this := ChangelogItem{}
	var classId string = "changelog.Item"
	this.ClassId = classId
	var objectType string = "changelog.Item"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChangelogItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChangelogItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChangelogItem) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "changelog.Item" of the ClassId field.
func (o *ChangelogItem) GetDefaultClassId() interface{} {
	return "changelog.Item"
}

// GetObjectType returns the ObjectType field value
func (o *ChangelogItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChangelogItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChangelogItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "changelog.Item" of the ObjectType field.
func (o *ChangelogItem) GetDefaultObjectType() interface{} {
	return "changelog.Item"
}

// GetDateVersion returns the DateVersion field value if set, zero value otherwise.
func (o *ChangelogItem) GetDateVersion() time.Time {
	if o == nil || IsNil(o.DateVersion) {
		var ret time.Time
		return ret
	}
	return *o.DateVersion
}

// GetDateVersionOk returns a tuple with the DateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogItem) GetDateVersionOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateVersion) {
		return nil, false
	}
	return o.DateVersion, true
}

// HasDateVersion returns a boolean if a field has been set.
func (o *ChangelogItem) HasDateVersion() bool {
	if o != nil && !IsNil(o.DateVersion) {
		return true
	}

	return false
}

// SetDateVersion gets a reference to the given time.Time and assigns it to the DateVersion field.
func (o *ChangelogItem) SetDateVersion(v time.Time) {
	o.DateVersion = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *ChangelogItem) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogItem) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *ChangelogItem) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *ChangelogItem) SetEntity(v string) {
	o.Entity = &v
}

// GetSemanticVersion returns the SemanticVersion field value if set, zero value otherwise.
func (o *ChangelogItem) GetSemanticVersion() string {
	if o == nil || IsNil(o.SemanticVersion) {
		var ret string
		return ret
	}
	return *o.SemanticVersion
}

// GetSemanticVersionOk returns a tuple with the SemanticVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogItem) GetSemanticVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SemanticVersion) {
		return nil, false
	}
	return o.SemanticVersion, true
}

// HasSemanticVersion returns a boolean if a field has been set.
func (o *ChangelogItem) HasSemanticVersion() bool {
	if o != nil && !IsNil(o.SemanticVersion) {
		return true
	}

	return false
}

// SetSemanticVersion gets a reference to the given string and assigns it to the SemanticVersion field.
func (o *ChangelogItem) SetSemanticVersion(v string) {
	o.SemanticVersion = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ChangelogItem) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogItem) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ChangelogItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ChangelogItem) SetValue(v string) {
	o.Value = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangelogItem) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangelogItem) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *ChangelogItem) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableWorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *ChangelogItem) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *ChangelogItem) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *ChangelogItem) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o ChangelogItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangelogItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DateVersion) {
		toSerialize["DateVersion"] = o.DateVersion
	}
	if !IsNil(o.Entity) {
		toSerialize["Entity"] = o.Entity
	}
	if !IsNil(o.SemanticVersion) {
		toSerialize["SemanticVersion"] = o.SemanticVersion
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangelogItem) UnmarshalJSON(data []byte) (err error) {
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
	type ChangelogItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The date version for the API contract changelog item in the format rfc3339 with no fraction seconds set.  Note that there can be more than one item per DateVersion. Example: 2023-12-19T00:00:00Z .
		DateVersion *time.Time `json:"DateVersion,omitempty"`
		// The operationId of the endpoint for which changelog item is being generated.
		Entity *string `json:"Entity,omitempty"`
		// The semantic version for the API contract changelog item. Note that there can be more than one item per SemanticVersion.
		SemanticVersion *string `json:"SemanticVersion,omitempty" validate:"regexp=^(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)(?:-((?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\\\.(?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\\\+([0-9a-zA-Z-]+(?:\\\\.[0-9a-zA-Z-]+)*))?$"`
		// The value of the API contract changelog item.
		Value   *string                             `json:"Value,omitempty"`
		Catalog NullableWorkflowCatalogRelationship `json:"Catalog,omitempty"`
	}

	varChangelogItemWithoutEmbeddedStruct := ChangelogItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varChangelogItemWithoutEmbeddedStruct)
	if err == nil {
		varChangelogItem := _ChangelogItem{}
		varChangelogItem.ClassId = varChangelogItemWithoutEmbeddedStruct.ClassId
		varChangelogItem.ObjectType = varChangelogItemWithoutEmbeddedStruct.ObjectType
		varChangelogItem.DateVersion = varChangelogItemWithoutEmbeddedStruct.DateVersion
		varChangelogItem.Entity = varChangelogItemWithoutEmbeddedStruct.Entity
		varChangelogItem.SemanticVersion = varChangelogItemWithoutEmbeddedStruct.SemanticVersion
		varChangelogItem.Value = varChangelogItemWithoutEmbeddedStruct.Value
		varChangelogItem.Catalog = varChangelogItemWithoutEmbeddedStruct.Catalog
		*o = ChangelogItem(varChangelogItem)
	} else {
		return err
	}

	varChangelogItem := _ChangelogItem{}

	err = json.Unmarshal(data, &varChangelogItem)
	if err == nil {
		o.MoBaseMo = varChangelogItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DateVersion")
		delete(additionalProperties, "Entity")
		delete(additionalProperties, "SemanticVersion")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "Catalog")

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

type NullableChangelogItem struct {
	value *ChangelogItem
	isSet bool
}

func (v NullableChangelogItem) Get() *ChangelogItem {
	return v.value
}

func (v *NullableChangelogItem) Set(val *ChangelogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableChangelogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableChangelogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangelogItem(val *ChangelogItem) *NullableChangelogItem {
	return &NullableChangelogItem{value: val, isSet: true}
}

func (v NullableChangelogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangelogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
