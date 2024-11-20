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

// checks if the SoftwareSolutionDistributable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareSolutionDistributable{}

// SoftwareSolutionDistributable A solution image distributed by Cisco.
type SoftwareSolutionDistributable struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The path of the file in S3/minio bucket.
	FilePath *string `json:"FilePath,omitempty"`
	// The name of the solution in which the image belongs.
	SolutionName *string `json:"SolutionName,omitempty"`
	// The type of the file like OS image, Script etc. * `osimage` - The solution OS image for deployment. * `script` - The Python script for the solution VM configuration and deployment.
	SubType              *string                                       `json:"SubType,omitempty"`
	Catalog              NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareSolutionDistributable SoftwareSolutionDistributable

// NewSoftwareSolutionDistributable instantiates a new SoftwareSolutionDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareSolutionDistributable(classId string, objectType string) *SoftwareSolutionDistributable {
	this := SoftwareSolutionDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var vendor string = "Cisco"
	this.Vendor = &vendor
	var subType string = "osimage"
	this.SubType = &subType
	return &this
}

// NewSoftwareSolutionDistributableWithDefaults instantiates a new SoftwareSolutionDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareSolutionDistributableWithDefaults() *SoftwareSolutionDistributable {
	this := SoftwareSolutionDistributable{}
	var classId string = "software.SolutionDistributable"
	this.ClassId = classId
	var objectType string = "software.SolutionDistributable"
	this.ObjectType = objectType
	var subType string = "osimage"
	this.SubType = &subType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareSolutionDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareSolutionDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "software.SolutionDistributable" of the ClassId field.
func (o *SoftwareSolutionDistributable) GetDefaultClassId() interface{} {
	return "software.SolutionDistributable"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareSolutionDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareSolutionDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "software.SolutionDistributable" of the ObjectType field.
func (o *SoftwareSolutionDistributable) GetDefaultObjectType() interface{} {
	return "software.SolutionDistributable"
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributable) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributable) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributable) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *SoftwareSolutionDistributable) SetFilePath(v string) {
	o.FilePath = &v
}

// GetSolutionName returns the SolutionName field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributable) GetSolutionName() string {
	if o == nil || IsNil(o.SolutionName) {
		var ret string
		return ret
	}
	return *o.SolutionName
}

// GetSolutionNameOk returns a tuple with the SolutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributable) GetSolutionNameOk() (*string, bool) {
	if o == nil || IsNil(o.SolutionName) {
		return nil, false
	}
	return o.SolutionName, true
}

// HasSolutionName returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributable) HasSolutionName() bool {
	if o != nil && !IsNil(o.SolutionName) {
		return true
	}

	return false
}

// SetSolutionName gets a reference to the given string and assigns it to the SolutionName field.
func (o *SoftwareSolutionDistributable) SetSolutionName(v string) {
	o.SolutionName = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributable) GetSubType() string {
	if o == nil || IsNil(o.SubType) {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributable) GetSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributable) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *SoftwareSolutionDistributable) SetSubType(v string) {
	o.SubType = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareSolutionDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareSolutionDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributable) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableSoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareSolutionDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *SoftwareSolutionDistributable) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *SoftwareSolutionDistributable) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o SoftwareSolutionDistributable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareSolutionDistributable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return map[string]interface{}{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return map[string]interface{}{}, errFirmwareBaseDistributable
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FilePath) {
		toSerialize["FilePath"] = o.FilePath
	}
	if !IsNil(o.SolutionName) {
		toSerialize["SolutionName"] = o.SolutionName
	}
	if !IsNil(o.SubType) {
		toSerialize["SubType"] = o.SubType
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwareSolutionDistributable) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwareSolutionDistributableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The path of the file in S3/minio bucket.
		FilePath *string `json:"FilePath,omitempty"`
		// The name of the solution in which the image belongs.
		SolutionName *string `json:"SolutionName,omitempty"`
		// The type of the file like OS image, Script etc. * `osimage` - The solution OS image for deployment. * `script` - The Python script for the solution VM configuration and deployment.
		SubType *string                                       `json:"SubType,omitempty"`
		Catalog NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varSoftwareSolutionDistributableWithoutEmbeddedStruct := SoftwareSolutionDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwareSolutionDistributableWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareSolutionDistributable := _SoftwareSolutionDistributable{}
		varSoftwareSolutionDistributable.ClassId = varSoftwareSolutionDistributableWithoutEmbeddedStruct.ClassId
		varSoftwareSolutionDistributable.ObjectType = varSoftwareSolutionDistributableWithoutEmbeddedStruct.ObjectType
		varSoftwareSolutionDistributable.FilePath = varSoftwareSolutionDistributableWithoutEmbeddedStruct.FilePath
		varSoftwareSolutionDistributable.SolutionName = varSoftwareSolutionDistributableWithoutEmbeddedStruct.SolutionName
		varSoftwareSolutionDistributable.SubType = varSoftwareSolutionDistributableWithoutEmbeddedStruct.SubType
		varSoftwareSolutionDistributable.Catalog = varSoftwareSolutionDistributableWithoutEmbeddedStruct.Catalog
		*o = SoftwareSolutionDistributable(varSoftwareSolutionDistributable)
	} else {
		return err
	}

	varSoftwareSolutionDistributable := _SoftwareSolutionDistributable{}

	err = json.Unmarshal(data, &varSoftwareSolutionDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varSoftwareSolutionDistributable.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "SolutionName")
		delete(additionalProperties, "SubType")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableSoftwareSolutionDistributable struct {
	value *SoftwareSolutionDistributable
	isSet bool
}

func (v NullableSoftwareSolutionDistributable) Get() *SoftwareSolutionDistributable {
	return v.value
}

func (v *NullableSoftwareSolutionDistributable) Set(val *SoftwareSolutionDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareSolutionDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareSolutionDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareSolutionDistributable(val *SoftwareSolutionDistributable) *NullableSoftwareSolutionDistributable {
	return &NullableSoftwareSolutionDistributable{value: val, isSet: true}
}

func (v NullableSoftwareSolutionDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareSolutionDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
