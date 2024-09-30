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

// checks if the SoftwareApplianceDistributable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareApplianceDistributable{}

// SoftwareApplianceDistributable Appliance image distributed by Cisco. This image is required to upgrade the on-premise Intersight Appliance. There are two use cases. In Intersight SaaS, the object represents a downloadable image, whereas on the Appliance the represents the image that is uploaded by the user and to be used for upgrade.
type SoftwareApplianceDistributable struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                        `json:"ObjectType"`
	Catalog              NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareApplianceDistributable SoftwareApplianceDistributable

// NewSoftwareApplianceDistributable instantiates a new SoftwareApplianceDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareApplianceDistributable(classId string, objectType string) *SoftwareApplianceDistributable {
	this := SoftwareApplianceDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var recommendedBuild string = "N"
	this.RecommendedBuild = &recommendedBuild
	var vendor string = "Cisco"
	this.Vendor = &vendor
	return &this
}

// NewSoftwareApplianceDistributableWithDefaults instantiates a new SoftwareApplianceDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareApplianceDistributableWithDefaults() *SoftwareApplianceDistributable {
	this := SoftwareApplianceDistributable{}
	var classId string = "software.ApplianceDistributable"
	this.ClassId = classId
	var objectType string = "software.ApplianceDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareApplianceDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareApplianceDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareApplianceDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "software.ApplianceDistributable" of the ClassId field.
func (o *SoftwareApplianceDistributable) GetDefaultClassId() interface{} {
	return "software.ApplianceDistributable"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareApplianceDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareApplianceDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareApplianceDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "software.ApplianceDistributable" of the ObjectType field.
func (o *SoftwareApplianceDistributable) GetDefaultObjectType() interface{} {
	return "software.ApplianceDistributable"
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareApplianceDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareApplianceDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareApplianceDistributable) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableSoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareApplianceDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *SoftwareApplianceDistributable) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *SoftwareApplianceDistributable) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o SoftwareApplianceDistributable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareApplianceDistributable) ToMap() (map[string]interface{}, error) {
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
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwareApplianceDistributable) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwareApplianceDistributableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                        `json:"ObjectType"`
		Catalog    NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varSoftwareApplianceDistributableWithoutEmbeddedStruct := SoftwareApplianceDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwareApplianceDistributableWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareApplianceDistributable := _SoftwareApplianceDistributable{}
		varSoftwareApplianceDistributable.ClassId = varSoftwareApplianceDistributableWithoutEmbeddedStruct.ClassId
		varSoftwareApplianceDistributable.ObjectType = varSoftwareApplianceDistributableWithoutEmbeddedStruct.ObjectType
		varSoftwareApplianceDistributable.Catalog = varSoftwareApplianceDistributableWithoutEmbeddedStruct.Catalog
		*o = SoftwareApplianceDistributable(varSoftwareApplianceDistributable)
	} else {
		return err
	}

	varSoftwareApplianceDistributable := _SoftwareApplianceDistributable{}

	err = json.Unmarshal(data, &varSoftwareApplianceDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varSoftwareApplianceDistributable.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
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

type NullableSoftwareApplianceDistributable struct {
	value *SoftwareApplianceDistributable
	isSet bool
}

func (v NullableSoftwareApplianceDistributable) Get() *SoftwareApplianceDistributable {
	return v.value
}

func (v *NullableSoftwareApplianceDistributable) Set(val *SoftwareApplianceDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareApplianceDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareApplianceDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareApplianceDistributable(val *SoftwareApplianceDistributable) *NullableSoftwareApplianceDistributable {
	return &NullableSoftwareApplianceDistributable{value: val, isSet: true}
}

func (v NullableSoftwareApplianceDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareApplianceDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
