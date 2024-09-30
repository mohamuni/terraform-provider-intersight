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

// checks if the SoftwarerepositoryCategoryMapperModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwarerepositoryCategoryMapperModel{}

// SoftwarerepositoryCategoryMapperModel Maps a Cisco hardware model Series to its applicable hardware models.
type SoftwarerepositoryCategoryMapperModel struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// The distributable tag value of the model series.
	DistTag *string `json:"DistTag,omitempty"`
	// The type of image based on the endpoint it can upgrade. For example, ucs-bundle-6400-infra.4.1.2a.bin can upgrade ucs managed fabric interconnects, so the image type is UCS Managed Fabric Interconnect.
	ImageType *string `json:"ImageType,omitempty"`
	// Defines whether NFS firmware upgrade is supported with this image type.
	IsNfsUpgradeSupported *bool `json:"IsNfsUpgradeSupported,omitempty"`
	// The regex that all images of this model follow.
	RegexPattern *string `json:"RegexPattern,omitempty"`
	// Cisco hardware model series.
	SeriesId             *string  `json:"SeriesId,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategoryMapperModel SoftwarerepositoryCategoryMapperModel

// NewSoftwarerepositoryCategoryMapperModel instantiates a new SoftwarerepositoryCategoryMapperModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapperModel(classId string, objectType string) *SoftwarerepositoryCategoryMapperModel {
	this := SoftwarerepositoryCategoryMapperModel{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryCategoryMapperModelWithDefaults instantiates a new SoftwarerepositoryCategoryMapperModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperModelWithDefaults() *SoftwarerepositoryCategoryMapperModel {
	this := SoftwarerepositoryCategoryMapperModel{}
	var classId string = "softwarerepository.CategoryMapperModel"
	this.ClassId = classId
	var objectType string = "softwarerepository.CategoryMapperModel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCategoryMapperModel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCategoryMapperModel) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "softwarerepository.CategoryMapperModel" of the ClassId field.
func (o *SoftwarerepositoryCategoryMapperModel) GetDefaultClassId() interface{} {
	return "softwarerepository.CategoryMapperModel"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCategoryMapperModel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCategoryMapperModel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "softwarerepository.CategoryMapperModel" of the ObjectType field.
func (o *SoftwarerepositoryCategoryMapperModel) GetDefaultObjectType() interface{} {
	return "softwarerepository.CategoryMapperModel"
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapperModel) SetCategory(v string) {
	o.Category = &v
}

// GetDistTag returns the DistTag field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetDistTag() string {
	if o == nil || IsNil(o.DistTag) {
		var ret string
		return ret
	}
	return *o.DistTag
}

// GetDistTagOk returns a tuple with the DistTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetDistTagOk() (*string, bool) {
	if o == nil || IsNil(o.DistTag) {
		return nil, false
	}
	return o.DistTag, true
}

// HasDistTag returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasDistTag() bool {
	if o != nil && !IsNil(o.DistTag) {
		return true
	}

	return false
}

// SetDistTag gets a reference to the given string and assigns it to the DistTag field.
func (o *SoftwarerepositoryCategoryMapperModel) SetDistTag(v string) {
	o.DistTag = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *SoftwarerepositoryCategoryMapperModel) SetImageType(v string) {
	o.ImageType = &v
}

// GetIsNfsUpgradeSupported returns the IsNfsUpgradeSupported field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetIsNfsUpgradeSupported() bool {
	if o == nil || IsNil(o.IsNfsUpgradeSupported) {
		var ret bool
		return ret
	}
	return *o.IsNfsUpgradeSupported
}

// GetIsNfsUpgradeSupportedOk returns a tuple with the IsNfsUpgradeSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetIsNfsUpgradeSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNfsUpgradeSupported) {
		return nil, false
	}
	return o.IsNfsUpgradeSupported, true
}

// HasIsNfsUpgradeSupported returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasIsNfsUpgradeSupported() bool {
	if o != nil && !IsNil(o.IsNfsUpgradeSupported) {
		return true
	}

	return false
}

// SetIsNfsUpgradeSupported gets a reference to the given bool and assigns it to the IsNfsUpgradeSupported field.
func (o *SoftwarerepositoryCategoryMapperModel) SetIsNfsUpgradeSupported(v bool) {
	o.IsNfsUpgradeSupported = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetRegexPattern() string {
	if o == nil || IsNil(o.RegexPattern) {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetRegexPatternOk() (*string, bool) {
	if o == nil || IsNil(o.RegexPattern) {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasRegexPattern() bool {
	if o != nil && !IsNil(o.RegexPattern) {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapperModel) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetSeriesId() string {
	if o == nil || IsNil(o.SeriesId) {
		var ret string
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetSeriesIdOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given string and assigns it to the SeriesId field.
func (o *SoftwarerepositoryCategoryMapperModel) SetSeriesId(v string) {
	o.SeriesId = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapperModel) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapperModel) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapperModel) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o SoftwarerepositoryCategoryMapperModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwarerepositoryCategoryMapperModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.DistTag) {
		toSerialize["DistTag"] = o.DistTag
	}
	if !IsNil(o.ImageType) {
		toSerialize["ImageType"] = o.ImageType
	}
	if !IsNil(o.IsNfsUpgradeSupported) {
		toSerialize["IsNfsUpgradeSupported"] = o.IsNfsUpgradeSupported
	}
	if !IsNil(o.RegexPattern) {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if !IsNil(o.SeriesId) {
		toSerialize["SeriesId"] = o.SeriesId
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwarerepositoryCategoryMapperModel) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The category of the model series.
		Category *string `json:"Category,omitempty"`
		// The distributable tag value of the model series.
		DistTag *string `json:"DistTag,omitempty"`
		// The type of image based on the endpoint it can upgrade. For example, ucs-bundle-6400-infra.4.1.2a.bin can upgrade ucs managed fabric interconnects, so the image type is UCS Managed Fabric Interconnect.
		ImageType *string `json:"ImageType,omitempty"`
		// Defines whether NFS firmware upgrade is supported with this image type.
		IsNfsUpgradeSupported *bool `json:"IsNfsUpgradeSupported,omitempty"`
		// The regex that all images of this model follow.
		RegexPattern *string `json:"RegexPattern,omitempty"`
		// Cisco hardware model series.
		SeriesId        *string  `json:"SeriesId,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
	}

	varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct := SoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryCategoryMapperModel := _SoftwarerepositoryCategoryMapperModel{}
		varSoftwarerepositoryCategoryMapperModel.ClassId = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryCategoryMapperModel.ObjectType = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryCategoryMapperModel.Category = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.Category
		varSoftwarerepositoryCategoryMapperModel.DistTag = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.DistTag
		varSoftwarerepositoryCategoryMapperModel.ImageType = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.ImageType
		varSoftwarerepositoryCategoryMapperModel.IsNfsUpgradeSupported = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.IsNfsUpgradeSupported
		varSoftwarerepositoryCategoryMapperModel.RegexPattern = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.RegexPattern
		varSoftwarerepositoryCategoryMapperModel.SeriesId = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.SeriesId
		varSoftwarerepositoryCategoryMapperModel.SupportedModels = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.SupportedModels
		*o = SoftwarerepositoryCategoryMapperModel(varSoftwarerepositoryCategoryMapperModel)
	} else {
		return err
	}

	varSoftwarerepositoryCategoryMapperModel := _SoftwarerepositoryCategoryMapperModel{}

	err = json.Unmarshal(data, &varSoftwarerepositoryCategoryMapperModel)
	if err == nil {
		o.CapabilityCapability = varSoftwarerepositoryCategoryMapperModel.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "DistTag")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "IsNfsUpgradeSupported")
		delete(additionalProperties, "RegexPattern")
		delete(additionalProperties, "SeriesId")
		delete(additionalProperties, "SupportedModels")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableSoftwarerepositoryCategoryMapperModel struct {
	value *SoftwarerepositoryCategoryMapperModel
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapperModel) Get() *SoftwarerepositoryCategoryMapperModel {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapperModel) Set(val *SoftwarerepositoryCategoryMapperModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapperModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapperModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapperModel(val *SoftwarerepositoryCategoryMapperModel) *NullableSoftwarerepositoryCategoryMapperModel {
	return &NullableSoftwarerepositoryCategoryMapperModel{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapperModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapperModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
