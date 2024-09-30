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

// checks if the CapabilitySwitchManufacturingDef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitySwitchManufacturingDef{}

// CapabilitySwitchManufacturingDef Switch/Fabric-Interconnect manufacturing def properties.
type CapabilitySwitchManufacturingDef struct {
	CapabilitySwitchCapabilityDef
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Caption for Switch/Fabric-Interconnect.
	Caption *string `json:"Caption,omitempty"`
	// Description for Switch/Fabric-Interconnect.
	Description *string `json:"Description,omitempty"`
	// Part Number for Switch/Fabric-Interconnect.
	PartNumber *string `json:"PartNumber,omitempty"`
	// Product Name for Switch/Fabric-Interconnect.
	ProductName          *string `json:"ProductName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchManufacturingDef CapabilitySwitchManufacturingDef

// NewCapabilitySwitchManufacturingDef instantiates a new CapabilitySwitchManufacturingDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchManufacturingDef(classId string, objectType string) *CapabilitySwitchManufacturingDef {
	this := CapabilitySwitchManufacturingDef{}
	this.ClassId = classId
	this.ObjectType = objectType
	var pid string = "UCS-FI-6454"
	this.Pid = &pid
	return &this
}

// NewCapabilitySwitchManufacturingDefWithDefaults instantiates a new CapabilitySwitchManufacturingDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchManufacturingDefWithDefaults() *CapabilitySwitchManufacturingDef {
	this := CapabilitySwitchManufacturingDef{}
	var classId string = "capability.SwitchManufacturingDef"
	this.ClassId = classId
	var objectType string = "capability.SwitchManufacturingDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchManufacturingDef) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchManufacturingDef) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.SwitchManufacturingDef" of the ClassId field.
func (o *CapabilitySwitchManufacturingDef) GetDefaultClassId() interface{} {
	return "capability.SwitchManufacturingDef"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchManufacturingDef) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchManufacturingDef) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.SwitchManufacturingDef" of the ObjectType field.
func (o *CapabilitySwitchManufacturingDef) GetDefaultObjectType() interface{} {
	return "capability.SwitchManufacturingDef"
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *CapabilitySwitchManufacturingDef) SetCaption(v string) {
	o.Caption = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilitySwitchManufacturingDef) SetDescription(v string) {
	o.Description = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *CapabilitySwitchManufacturingDef) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *CapabilitySwitchManufacturingDef) SetProductName(v string) {
	o.ProductName = &v
}

func (o CapabilitySwitchManufacturingDef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitySwitchManufacturingDef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilitySwitchCapabilityDef, errCapabilitySwitchCapabilityDef := json.Marshal(o.CapabilitySwitchCapabilityDef)
	if errCapabilitySwitchCapabilityDef != nil {
		return map[string]interface{}{}, errCapabilitySwitchCapabilityDef
	}
	errCapabilitySwitchCapabilityDef = json.Unmarshal([]byte(serializedCapabilitySwitchCapabilityDef), &toSerialize)
	if errCapabilitySwitchCapabilityDef != nil {
		return map[string]interface{}{}, errCapabilitySwitchCapabilityDef
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Caption) {
		toSerialize["Caption"] = o.Caption
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.ProductName) {
		toSerialize["ProductName"] = o.ProductName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitySwitchManufacturingDef) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilitySwitchManufacturingDefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Caption for Switch/Fabric-Interconnect.
		Caption *string `json:"Caption,omitempty"`
		// Description for Switch/Fabric-Interconnect.
		Description *string `json:"Description,omitempty"`
		// Part Number for Switch/Fabric-Interconnect.
		PartNumber *string `json:"PartNumber,omitempty"`
		// Product Name for Switch/Fabric-Interconnect.
		ProductName *string `json:"ProductName,omitempty"`
	}

	varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct := CapabilitySwitchManufacturingDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchManufacturingDef := _CapabilitySwitchManufacturingDef{}
		varCapabilitySwitchManufacturingDef.ClassId = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.ClassId
		varCapabilitySwitchManufacturingDef.ObjectType = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.ObjectType
		varCapabilitySwitchManufacturingDef.Caption = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.Caption
		varCapabilitySwitchManufacturingDef.Description = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.Description
		varCapabilitySwitchManufacturingDef.PartNumber = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.PartNumber
		varCapabilitySwitchManufacturingDef.ProductName = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.ProductName
		*o = CapabilitySwitchManufacturingDef(varCapabilitySwitchManufacturingDef)
	} else {
		return err
	}

	varCapabilitySwitchManufacturingDef := _CapabilitySwitchManufacturingDef{}

	err = json.Unmarshal(data, &varCapabilitySwitchManufacturingDef)
	if err == nil {
		o.CapabilitySwitchCapabilityDef = varCapabilitySwitchManufacturingDef.CapabilitySwitchCapabilityDef
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Caption")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "ProductName")

		// remove fields from embedded structs
		reflectCapabilitySwitchCapabilityDef := reflect.ValueOf(o.CapabilitySwitchCapabilityDef)
		for i := 0; i < reflectCapabilitySwitchCapabilityDef.Type().NumField(); i++ {
			t := reflectCapabilitySwitchCapabilityDef.Type().Field(i)

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

type NullableCapabilitySwitchManufacturingDef struct {
	value *CapabilitySwitchManufacturingDef
	isSet bool
}

func (v NullableCapabilitySwitchManufacturingDef) Get() *CapabilitySwitchManufacturingDef {
	return v.value
}

func (v *NullableCapabilitySwitchManufacturingDef) Set(val *CapabilitySwitchManufacturingDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchManufacturingDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchManufacturingDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchManufacturingDef(val *CapabilitySwitchManufacturingDef) *NullableCapabilitySwitchManufacturingDef {
	return &NullableCapabilitySwitchManufacturingDef{value: val, isSet: true}
}

func (v NullableCapabilitySwitchManufacturingDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchManufacturingDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
