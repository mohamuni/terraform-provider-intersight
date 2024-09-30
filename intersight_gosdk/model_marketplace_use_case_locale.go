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

// checks if the MarketplaceUseCaseLocale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceUseCaseLocale{}

// MarketplaceUseCaseLocale Locale object for UseCase.
type MarketplaceUseCaseLocale struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                         `json:"ObjectType"`
	Automations []MarketplaceUseCaseAutomation `json:"Automations,omitempty"`
	// The string field to hold the contents value.
	Contents *string `json:"Contents,omitempty"`
	// The string field to hold the description value.
	Description *string `json:"Description,omitempty"`
	// A base64-encoded image for the use case.
	Icon *string `json:"Icon,omitempty"`
	// The string field to hold the locale.
	Locale *string `json:"Locale,omitempty" validate:"regexp=^$|^[a-z]{2,4}(-[A-Z][a-z]{3})?(-([A-Za-z]{2}|[0-9]{3}))?$"`
	// The string field to hold the name value.
	Name *string `json:"Name,omitempty"`
	// The string field to hold the summary value.
	Summary              *string `json:"Summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseLocale MarketplaceUseCaseLocale

// NewMarketplaceUseCaseLocale instantiates a new MarketplaceUseCaseLocale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseLocale(classId string, objectType string) *MarketplaceUseCaseLocale {
	this := MarketplaceUseCaseLocale{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseLocaleWithDefaults instantiates a new MarketplaceUseCaseLocale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseLocaleWithDefaults() *MarketplaceUseCaseLocale {
	this := MarketplaceUseCaseLocale{}
	var classId string = "marketplace.UseCaseLocale"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseLocale"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseLocale) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseLocale) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "marketplace.UseCaseLocale" of the ClassId field.
func (o *MarketplaceUseCaseLocale) GetDefaultClassId() interface{} {
	return "marketplace.UseCaseLocale"
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseLocale) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseLocale) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "marketplace.UseCaseLocale" of the ObjectType field.
func (o *MarketplaceUseCaseLocale) GetDefaultObjectType() interface{} {
	return "marketplace.UseCaseLocale"
}

// GetAutomations returns the Automations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarketplaceUseCaseLocale) GetAutomations() []MarketplaceUseCaseAutomation {
	if o == nil {
		var ret []MarketplaceUseCaseAutomation
		return ret
	}
	return o.Automations
}

// GetAutomationsOk returns a tuple with the Automations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarketplaceUseCaseLocale) GetAutomationsOk() ([]MarketplaceUseCaseAutomation, bool) {
	if o == nil || IsNil(o.Automations) {
		return nil, false
	}
	return o.Automations, true
}

// HasAutomations returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasAutomations() bool {
	if o != nil && !IsNil(o.Automations) {
		return true
	}

	return false
}

// SetAutomations gets a reference to the given []MarketplaceUseCaseAutomation and assigns it to the Automations field.
func (o *MarketplaceUseCaseLocale) SetAutomations(v []MarketplaceUseCaseAutomation) {
	o.Automations = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocale) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *MarketplaceUseCaseLocale) SetContents(v string) {
	o.Contents = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocale) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MarketplaceUseCaseLocale) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocale) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *MarketplaceUseCaseLocale) SetIcon(v string) {
	o.Icon = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocale) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *MarketplaceUseCaseLocale) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocale) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketplaceUseCaseLocale) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocale) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocale) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocale) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *MarketplaceUseCaseLocale) SetSummary(v string) {
	o.Summary = &v
}

func (o MarketplaceUseCaseLocale) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceUseCaseLocale) ToMap() (map[string]interface{}, error) {
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
	if o.Automations != nil {
		toSerialize["Automations"] = o.Automations
	}
	if !IsNil(o.Contents) {
		toSerialize["Contents"] = o.Contents
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["Icon"] = o.Icon
	}
	if !IsNil(o.Locale) {
		toSerialize["Locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Summary) {
		toSerialize["Summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarketplaceUseCaseLocale) UnmarshalJSON(data []byte) (err error) {
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
	type MarketplaceUseCaseLocaleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                         `json:"ObjectType"`
		Automations []MarketplaceUseCaseAutomation `json:"Automations,omitempty"`
		// The string field to hold the contents value.
		Contents *string `json:"Contents,omitempty"`
		// The string field to hold the description value.
		Description *string `json:"Description,omitempty"`
		// A base64-encoded image for the use case.
		Icon *string `json:"Icon,omitempty"`
		// The string field to hold the locale.
		Locale *string `json:"Locale,omitempty" validate:"regexp=^$|^[a-z]{2,4}(-[A-Z][a-z]{3})?(-([A-Za-z]{2}|[0-9]{3}))?$"`
		// The string field to hold the name value.
		Name *string `json:"Name,omitempty"`
		// The string field to hold the summary value.
		Summary *string `json:"Summary,omitempty"`
	}

	varMarketplaceUseCaseLocaleWithoutEmbeddedStruct := MarketplaceUseCaseLocaleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMarketplaceUseCaseLocaleWithoutEmbeddedStruct)
	if err == nil {
		varMarketplaceUseCaseLocale := _MarketplaceUseCaseLocale{}
		varMarketplaceUseCaseLocale.ClassId = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.ClassId
		varMarketplaceUseCaseLocale.ObjectType = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.ObjectType
		varMarketplaceUseCaseLocale.Automations = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Automations
		varMarketplaceUseCaseLocale.Contents = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Contents
		varMarketplaceUseCaseLocale.Description = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Description
		varMarketplaceUseCaseLocale.Icon = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Icon
		varMarketplaceUseCaseLocale.Locale = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Locale
		varMarketplaceUseCaseLocale.Name = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Name
		varMarketplaceUseCaseLocale.Summary = varMarketplaceUseCaseLocaleWithoutEmbeddedStruct.Summary
		*o = MarketplaceUseCaseLocale(varMarketplaceUseCaseLocale)
	} else {
		return err
	}

	varMarketplaceUseCaseLocale := _MarketplaceUseCaseLocale{}

	err = json.Unmarshal(data, &varMarketplaceUseCaseLocale)
	if err == nil {
		o.MoBaseComplexType = varMarketplaceUseCaseLocale.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Automations")
		delete(additionalProperties, "Contents")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Icon")
		delete(additionalProperties, "Locale")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Summary")

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

type NullableMarketplaceUseCaseLocale struct {
	value *MarketplaceUseCaseLocale
	isSet bool
}

func (v NullableMarketplaceUseCaseLocale) Get() *MarketplaceUseCaseLocale {
	return v.value
}

func (v *NullableMarketplaceUseCaseLocale) Set(val *MarketplaceUseCaseLocale) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseLocale(val *MarketplaceUseCaseLocale) *NullableMarketplaceUseCaseLocale {
	return &NullableMarketplaceUseCaseLocale{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseLocale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
