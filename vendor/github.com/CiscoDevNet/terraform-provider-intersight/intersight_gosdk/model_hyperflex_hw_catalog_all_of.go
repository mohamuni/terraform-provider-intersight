/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexHwCatalogAllOf Definition of the list of properties defined in 'hyperflex.HwCatalog', excluding properties defined in parent classes.
type HyperflexHwCatalogAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Hardware catalog version for HyperFlex hardware catalog.
	Version *string `json:"Version,omitempty"`
	// An array of relationships to hclHwCatalogInfo resources.
	CatalogInfos         []HclHwCatalogInfoRelationship `json:"CatalogInfos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHwCatalogAllOf HyperflexHwCatalogAllOf

// NewHyperflexHwCatalogAllOf instantiates a new HyperflexHwCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHwCatalogAllOf(classId string, objectType string) *HyperflexHwCatalogAllOf {
	this := HyperflexHwCatalogAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHwCatalogAllOfWithDefaults instantiates a new HyperflexHwCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHwCatalogAllOfWithDefaults() *HyperflexHwCatalogAllOf {
	this := HyperflexHwCatalogAllOf{}
	var classId string = "hyperflex.HwCatalog"
	this.ClassId = classId
	var objectType string = "hyperflex.HwCatalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHwCatalogAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHwCatalogAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHwCatalogAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHwCatalogAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHwCatalogAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHwCatalogAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHwCatalogAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHwCatalogAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHwCatalogAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHwCatalogAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCatalogInfos returns the CatalogInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHwCatalogAllOf) GetCatalogInfos() []HclHwCatalogInfoRelationship {
	if o == nil {
		var ret []HclHwCatalogInfoRelationship
		return ret
	}
	return o.CatalogInfos
}

// GetCatalogInfosOk returns a tuple with the CatalogInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHwCatalogAllOf) GetCatalogInfosOk() ([]HclHwCatalogInfoRelationship, bool) {
	if o == nil || o.CatalogInfos == nil {
		return nil, false
	}
	return o.CatalogInfos, true
}

// HasCatalogInfos returns a boolean if a field has been set.
func (o *HyperflexHwCatalogAllOf) HasCatalogInfos() bool {
	if o != nil && o.CatalogInfos != nil {
		return true
	}

	return false
}

// SetCatalogInfos gets a reference to the given []HclHwCatalogInfoRelationship and assigns it to the CatalogInfos field.
func (o *HyperflexHwCatalogAllOf) SetCatalogInfos(v []HclHwCatalogInfoRelationship) {
	o.CatalogInfos = v
}

func (o HyperflexHwCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.CatalogInfos != nil {
		toSerialize["CatalogInfos"] = o.CatalogInfos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHwCatalogAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHwCatalogAllOf := _HyperflexHwCatalogAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHwCatalogAllOf); err == nil {
		*o = HyperflexHwCatalogAllOf(varHyperflexHwCatalogAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "CatalogInfos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHwCatalogAllOf struct {
	value *HyperflexHwCatalogAllOf
	isSet bool
}

func (v NullableHyperflexHwCatalogAllOf) Get() *HyperflexHwCatalogAllOf {
	return v.value
}

func (v *NullableHyperflexHwCatalogAllOf) Set(val *HyperflexHwCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHwCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHwCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHwCatalogAllOf(val *HyperflexHwCatalogAllOf) *NullableHyperflexHwCatalogAllOf {
	return &NullableHyperflexHwCatalogAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHwCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHwCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
