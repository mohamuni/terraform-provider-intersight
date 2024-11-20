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

// checks if the PartnerintegrationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerintegrationModel{}

// PartnerintegrationModel Model definition for the endpoint.
type PartnerintegrationModel struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Endpoint model in yaml format.
	Data interface{} `json:"Data,omitempty"`
	// Placeholder name for the model.
	Name                 *string                                         `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.-]{1,64}$"`
	Inventory            NullablePartnerintegrationInventoryRelationship `json:"Inventory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationModel PartnerintegrationModel

// NewPartnerintegrationModel instantiates a new PartnerintegrationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationModel(classId string, objectType string) *PartnerintegrationModel {
	this := PartnerintegrationModel{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPartnerintegrationModelWithDefaults instantiates a new PartnerintegrationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationModelWithDefaults() *PartnerintegrationModel {
	this := PartnerintegrationModel{}
	var classId string = "partnerintegration.Model"
	this.ClassId = classId
	var objectType string = "partnerintegration.Model"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationModel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationModel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationModel) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "partnerintegration.Model" of the ClassId field.
func (o *PartnerintegrationModel) GetDefaultClassId() interface{} {
	return "partnerintegration.Model"
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationModel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationModel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationModel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "partnerintegration.Model" of the ObjectType field.
func (o *PartnerintegrationModel) GetDefaultObjectType() interface{} {
	return "partnerintegration.Model"
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationModel) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationModel) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PartnerintegrationModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *PartnerintegrationModel) SetData(v interface{}) {
	o.Data = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartnerintegrationModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartnerintegrationModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartnerintegrationModel) SetName(v string) {
	o.Name = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationModel) GetInventory() PartnerintegrationInventoryRelationship {
	if o == nil || IsNil(o.Inventory.Get()) {
		var ret PartnerintegrationInventoryRelationship
		return ret
	}
	return *o.Inventory.Get()
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationModel) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inventory.Get(), o.Inventory.IsSet()
}

// HasInventory returns a boolean if a field has been set.
func (o *PartnerintegrationModel) HasInventory() bool {
	if o != nil && o.Inventory.IsSet() {
		return true
	}

	return false
}

// SetInventory gets a reference to the given NullablePartnerintegrationInventoryRelationship and assigns it to the Inventory field.
func (o *PartnerintegrationModel) SetInventory(v PartnerintegrationInventoryRelationship) {
	o.Inventory.Set(&v)
}

// SetInventoryNil sets the value for Inventory to be an explicit nil
func (o *PartnerintegrationModel) SetInventoryNil() {
	o.Inventory.Set(nil)
}

// UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
func (o *PartnerintegrationModel) UnsetInventory() {
	o.Inventory.Unset()
}

func (o PartnerintegrationModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerintegrationModel) ToMap() (map[string]interface{}, error) {
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
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Inventory.IsSet() {
		toSerialize["Inventory"] = o.Inventory.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PartnerintegrationModel) UnmarshalJSON(data []byte) (err error) {
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
	type PartnerintegrationModelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Endpoint model in yaml format.
		Data interface{} `json:"Data,omitempty"`
		// Placeholder name for the model.
		Name      *string                                         `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.-]{1,64}$"`
		Inventory NullablePartnerintegrationInventoryRelationship `json:"Inventory,omitempty"`
	}

	varPartnerintegrationModelWithoutEmbeddedStruct := PartnerintegrationModelWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPartnerintegrationModelWithoutEmbeddedStruct)
	if err == nil {
		varPartnerintegrationModel := _PartnerintegrationModel{}
		varPartnerintegrationModel.ClassId = varPartnerintegrationModelWithoutEmbeddedStruct.ClassId
		varPartnerintegrationModel.ObjectType = varPartnerintegrationModelWithoutEmbeddedStruct.ObjectType
		varPartnerintegrationModel.Data = varPartnerintegrationModelWithoutEmbeddedStruct.Data
		varPartnerintegrationModel.Name = varPartnerintegrationModelWithoutEmbeddedStruct.Name
		varPartnerintegrationModel.Inventory = varPartnerintegrationModelWithoutEmbeddedStruct.Inventory
		*o = PartnerintegrationModel(varPartnerintegrationModel)
	} else {
		return err
	}

	varPartnerintegrationModel := _PartnerintegrationModel{}

	err = json.Unmarshal(data, &varPartnerintegrationModel)
	if err == nil {
		o.MoBaseMo = varPartnerintegrationModel.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Data")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Inventory")

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

type NullablePartnerintegrationModel struct {
	value *PartnerintegrationModel
	isSet bool
}

func (v NullablePartnerintegrationModel) Get() *PartnerintegrationModel {
	return v.value
}

func (v *NullablePartnerintegrationModel) Set(val *PartnerintegrationModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationModel(val *PartnerintegrationModel) *NullablePartnerintegrationModel {
	return &NullablePartnerintegrationModel{value: val, isSet: true}
}

func (v NullablePartnerintegrationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
