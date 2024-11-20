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

// checks if the NiatelemetrySites type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetrySites{}

// NiatelemetrySites Stores information of sites that are in an insight group.
type NiatelemetrySites struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the name of the site.
	Name *string `json:"Name,omitempty"`
	// Returns the type of the site.
	SiteType *string `json:"SiteType,omitempty"`
	// Returns the uuid of the site.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySites NiatelemetrySites

// NewNiatelemetrySites instantiates a new NiatelemetrySites object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySites(classId string, objectType string) *NiatelemetrySites {
	this := NiatelemetrySites{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySitesWithDefaults instantiates a new NiatelemetrySites object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySitesWithDefaults() *NiatelemetrySites {
	this := NiatelemetrySites{}
	var classId string = "niatelemetry.Sites"
	this.ClassId = classId
	var objectType string = "niatelemetry.Sites"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySites) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySites) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySites) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.Sites" of the ClassId field.
func (o *NiatelemetrySites) GetDefaultClassId() interface{} {
	return "niatelemetry.Sites"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySites) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySites) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySites) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.Sites" of the ObjectType field.
func (o *NiatelemetrySites) GetDefaultObjectType() interface{} {
	return "niatelemetry.Sites"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetrySites) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySites) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetrySites) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetrySites) SetName(v string) {
	o.Name = &v
}

// GetSiteType returns the SiteType field value if set, zero value otherwise.
func (o *NiatelemetrySites) GetSiteType() string {
	if o == nil || IsNil(o.SiteType) {
		var ret string
		return ret
	}
	return *o.SiteType
}

// GetSiteTypeOk returns a tuple with the SiteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySites) GetSiteTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SiteType) {
		return nil, false
	}
	return o.SiteType, true
}

// HasSiteType returns a boolean if a field has been set.
func (o *NiatelemetrySites) HasSiteType() bool {
	if o != nil && !IsNil(o.SiteType) {
		return true
	}

	return false
}

// SetSiteType gets a reference to the given string and assigns it to the SiteType field.
func (o *NiatelemetrySites) SetSiteType(v string) {
	o.SiteType = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NiatelemetrySites) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySites) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NiatelemetrySites) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NiatelemetrySites) SetUuid(v string) {
	o.Uuid = &v
}

func (o NiatelemetrySites) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetrySites) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SiteType) {
		toSerialize["SiteType"] = o.SiteType
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetrySites) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetrySitesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns the name of the site.
		Name *string `json:"Name,omitempty"`
		// Returns the type of the site.
		SiteType *string `json:"SiteType,omitempty"`
		// Returns the uuid of the site.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varNiatelemetrySitesWithoutEmbeddedStruct := NiatelemetrySitesWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetrySitesWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySites := _NiatelemetrySites{}
		varNiatelemetrySites.ClassId = varNiatelemetrySitesWithoutEmbeddedStruct.ClassId
		varNiatelemetrySites.ObjectType = varNiatelemetrySitesWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySites.Name = varNiatelemetrySitesWithoutEmbeddedStruct.Name
		varNiatelemetrySites.SiteType = varNiatelemetrySitesWithoutEmbeddedStruct.SiteType
		varNiatelemetrySites.Uuid = varNiatelemetrySitesWithoutEmbeddedStruct.Uuid
		*o = NiatelemetrySites(varNiatelemetrySites)
	} else {
		return err
	}

	varNiatelemetrySites := _NiatelemetrySites{}

	err = json.Unmarshal(data, &varNiatelemetrySites)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetrySites.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SiteType")
		delete(additionalProperties, "Uuid")

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

type NullableNiatelemetrySites struct {
	value *NiatelemetrySites
	isSet bool
}

func (v NullableNiatelemetrySites) Get() *NiatelemetrySites {
	return v.value
}

func (v *NullableNiatelemetrySites) Set(val *NiatelemetrySites) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySites) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySites) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySites(val *NiatelemetrySites) *NullableNiatelemetrySites {
	return &NullableNiatelemetrySites{value: val, isSet: true}
}

func (v NullableNiatelemetrySites) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySites) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
