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

// checks if the DnacSite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacSite{}

// DnacSite Information about the Site.
type DnacSite struct {
	DnacInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Fabric site name hierarchy.
	FabricSiteNameHierarchy *string `json:"FabricSiteNameHierarchy,omitempty"`
	// Identification for the Site.
	SiteId               *string `json:"SiteId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacSite DnacSite

// NewDnacSite instantiates a new DnacSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacSite(classId string, objectType string) *DnacSite {
	this := DnacSite{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewDnacSiteWithDefaults instantiates a new DnacSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacSiteWithDefaults() *DnacSite {
	this := DnacSite{}
	var classId string = "dnac.Site"
	this.ClassId = classId
	var objectType string = "dnac.Site"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *DnacSite) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DnacSite) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DnacSite) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "dnac.Site" of the ClassId field.
func (o *DnacSite) GetDefaultClassId() interface{} {
	return "dnac.Site"
}

// GetObjectType returns the ObjectType field value
func (o *DnacSite) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DnacSite) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DnacSite) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "dnac.Site" of the ObjectType field.
func (o *DnacSite) GetDefaultObjectType() interface{} {
	return "dnac.Site"
}

// GetFabricSiteNameHierarchy returns the FabricSiteNameHierarchy field value if set, zero value otherwise.
func (o *DnacSite) GetFabricSiteNameHierarchy() string {
	if o == nil || IsNil(o.FabricSiteNameHierarchy) {
		var ret string
		return ret
	}
	return *o.FabricSiteNameHierarchy
}

// GetFabricSiteNameHierarchyOk returns a tuple with the FabricSiteNameHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacSite) GetFabricSiteNameHierarchyOk() (*string, bool) {
	if o == nil || IsNil(o.FabricSiteNameHierarchy) {
		return nil, false
	}
	return o.FabricSiteNameHierarchy, true
}

// HasFabricSiteNameHierarchy returns a boolean if a field has been set.
func (o *DnacSite) HasFabricSiteNameHierarchy() bool {
	if o != nil && !IsNil(o.FabricSiteNameHierarchy) {
		return true
	}

	return false
}

// SetFabricSiteNameHierarchy gets a reference to the given string and assigns it to the FabricSiteNameHierarchy field.
func (o *DnacSite) SetFabricSiteNameHierarchy(v string) {
	o.FabricSiteNameHierarchy = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DnacSite) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacSite) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DnacSite) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DnacSite) SetSiteId(v string) {
	o.SiteId = &v
}

func (o DnacSite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacSite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDnacInventoryEntity, errDnacInventoryEntity := json.Marshal(o.DnacInventoryEntity)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	errDnacInventoryEntity = json.Unmarshal([]byte(serializedDnacInventoryEntity), &toSerialize)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FabricSiteNameHierarchy) {
		toSerialize["FabricSiteNameHierarchy"] = o.FabricSiteNameHierarchy
	}
	if !IsNil(o.SiteId) {
		toSerialize["SiteId"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacSite) UnmarshalJSON(data []byte) (err error) {
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
	type DnacSiteWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Fabric site name hierarchy.
		FabricSiteNameHierarchy *string `json:"FabricSiteNameHierarchy,omitempty"`
		// Identification for the Site.
		SiteId *string `json:"SiteId,omitempty"`
	}

	varDnacSiteWithoutEmbeddedStruct := DnacSiteWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacSiteWithoutEmbeddedStruct)
	if err == nil {
		varDnacSite := _DnacSite{}
		varDnacSite.ClassId = varDnacSiteWithoutEmbeddedStruct.ClassId
		varDnacSite.ObjectType = varDnacSiteWithoutEmbeddedStruct.ObjectType
		varDnacSite.FabricSiteNameHierarchy = varDnacSiteWithoutEmbeddedStruct.FabricSiteNameHierarchy
		varDnacSite.SiteId = varDnacSiteWithoutEmbeddedStruct.SiteId
		*o = DnacSite(varDnacSite)
	} else {
		return err
	}

	varDnacSite := _DnacSite{}

	err = json.Unmarshal(data, &varDnacSite)
	if err == nil {
		o.DnacInventoryEntity = varDnacSite.DnacInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FabricSiteNameHierarchy")
		delete(additionalProperties, "SiteId")

		// remove fields from embedded structs
		reflectDnacInventoryEntity := reflect.ValueOf(o.DnacInventoryEntity)
		for i := 0; i < reflectDnacInventoryEntity.Type().NumField(); i++ {
			t := reflectDnacInventoryEntity.Type().Field(i)

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

type NullableDnacSite struct {
	value *DnacSite
	isSet bool
}

func (v NullableDnacSite) Get() *DnacSite {
	return v.value
}

func (v *NullableDnacSite) Set(val *DnacSite) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacSite) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacSite(val *DnacSite) *NullableDnacSite {
	return &NullableDnacSite{value: val, isSet: true}
}

func (v NullableDnacSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
