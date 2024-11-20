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

// checks if the HyperflexCapabilityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexCapabilityInfo{}

// HyperflexCapabilityInfo A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, \"minUcsVersion\" for HyperFlex version \"4.0.1a\" corresponds to \"3.2.3\" or \"minHxdpVersion\" for HyperFlex Upgrade operation is \"4.0.1a\" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
type HyperflexCapabilityInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string          `json:"ObjectType"`
	CapabilityConstraints []HclConstraint `json:"CapabilityConstraints,omitempty"`
	// Name of the capability or feature set consisting of a collection of constraint rules and value.
	Name *string `json:"Name,omitempty"`
	// Capability Value which is valid only iff all specified constraints match.
	Value                *string                                 `json:"Value,omitempty"`
	AppCatalog           NullableHyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexCapabilityInfo HyperflexCapabilityInfo

// NewHyperflexCapabilityInfo instantiates a new HyperflexCapabilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexCapabilityInfo(classId string, objectType string) *HyperflexCapabilityInfo {
	this := HyperflexCapabilityInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexCapabilityInfoWithDefaults instantiates a new HyperflexCapabilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexCapabilityInfoWithDefaults() *HyperflexCapabilityInfo {
	this := HyperflexCapabilityInfo{}
	var classId string = "hyperflex.CapabilityInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.CapabilityInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexCapabilityInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexCapabilityInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.CapabilityInfo" of the ClassId field.
func (o *HyperflexCapabilityInfo) GetDefaultClassId() interface{} {
	return "hyperflex.CapabilityInfo"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexCapabilityInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexCapabilityInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.CapabilityInfo" of the ObjectType field.
func (o *HyperflexCapabilityInfo) GetDefaultObjectType() interface{} {
	return "hyperflex.CapabilityInfo"
}

// GetCapabilityConstraints returns the CapabilityConstraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCapabilityInfo) GetCapabilityConstraints() []HclConstraint {
	if o == nil {
		var ret []HclConstraint
		return ret
	}
	return o.CapabilityConstraints
}

// GetCapabilityConstraintsOk returns a tuple with the CapabilityConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCapabilityInfo) GetCapabilityConstraintsOk() ([]HclConstraint, bool) {
	if o == nil || IsNil(o.CapabilityConstraints) {
		return nil, false
	}
	return o.CapabilityConstraints, true
}

// HasCapabilityConstraints returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfo) HasCapabilityConstraints() bool {
	if o != nil && !IsNil(o.CapabilityConstraints) {
		return true
	}

	return false
}

// SetCapabilityConstraints gets a reference to the given []HclConstraint and assigns it to the CapabilityConstraints field.
func (o *HyperflexCapabilityInfo) SetCapabilityConstraints(v []HclConstraint) {
	o.CapabilityConstraints = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexCapabilityInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexCapabilityInfo) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HyperflexCapabilityInfo) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfo) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfo) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HyperflexCapabilityInfo) SetValue(v string) {
	o.Value = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCapabilityInfo) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || IsNil(o.AppCatalog.Get()) {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog.Get()
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCapabilityInfo) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppCatalog.Get(), o.AppCatalog.IsSet()
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfo) HasAppCatalog() bool {
	if o != nil && o.AppCatalog.IsSet() {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given NullableHyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexCapabilityInfo) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog.Set(&v)
}

// SetAppCatalogNil sets the value for AppCatalog to be an explicit nil
func (o *HyperflexCapabilityInfo) SetAppCatalogNil() {
	o.AppCatalog.Set(nil)
}

// UnsetAppCatalog ensures that no value is present for AppCatalog, not even an explicit nil
func (o *HyperflexCapabilityInfo) UnsetAppCatalog() {
	o.AppCatalog.Unset()
}

func (o HyperflexCapabilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexCapabilityInfo) ToMap() (map[string]interface{}, error) {
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
	if o.CapabilityConstraints != nil {
		toSerialize["CapabilityConstraints"] = o.CapabilityConstraints
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	if o.AppCatalog.IsSet() {
		toSerialize["AppCatalog"] = o.AppCatalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexCapabilityInfo) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexCapabilityInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType            string          `json:"ObjectType"`
		CapabilityConstraints []HclConstraint `json:"CapabilityConstraints,omitempty"`
		// Name of the capability or feature set consisting of a collection of constraint rules and value.
		Name *string `json:"Name,omitempty"`
		// Capability Value which is valid only iff all specified constraints match.
		Value      *string                                 `json:"Value,omitempty"`
		AppCatalog NullableHyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHyperflexCapabilityInfoWithoutEmbeddedStruct := HyperflexCapabilityInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexCapabilityInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexCapabilityInfo := _HyperflexCapabilityInfo{}
		varHyperflexCapabilityInfo.ClassId = varHyperflexCapabilityInfoWithoutEmbeddedStruct.ClassId
		varHyperflexCapabilityInfo.ObjectType = varHyperflexCapabilityInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexCapabilityInfo.CapabilityConstraints = varHyperflexCapabilityInfoWithoutEmbeddedStruct.CapabilityConstraints
		varHyperflexCapabilityInfo.Name = varHyperflexCapabilityInfoWithoutEmbeddedStruct.Name
		varHyperflexCapabilityInfo.Value = varHyperflexCapabilityInfoWithoutEmbeddedStruct.Value
		varHyperflexCapabilityInfo.AppCatalog = varHyperflexCapabilityInfoWithoutEmbeddedStruct.AppCatalog
		*o = HyperflexCapabilityInfo(varHyperflexCapabilityInfo)
	} else {
		return err
	}

	varHyperflexCapabilityInfo := _HyperflexCapabilityInfo{}

	err = json.Unmarshal(data, &varHyperflexCapabilityInfo)
	if err == nil {
		o.MoBaseMo = varHyperflexCapabilityInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CapabilityConstraints")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "AppCatalog")

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

type NullableHyperflexCapabilityInfo struct {
	value *HyperflexCapabilityInfo
	isSet bool
}

func (v NullableHyperflexCapabilityInfo) Get() *HyperflexCapabilityInfo {
	return v.value
}

func (v *NullableHyperflexCapabilityInfo) Set(val *HyperflexCapabilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexCapabilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexCapabilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexCapabilityInfo(val *HyperflexCapabilityInfo) *NullableHyperflexCapabilityInfo {
	return &NullableHyperflexCapabilityInfo{value: val, isSet: true}
}

func (v NullableHyperflexCapabilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexCapabilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
