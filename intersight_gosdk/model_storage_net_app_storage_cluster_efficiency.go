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

// checks if the StorageNetAppStorageClusterEfficiency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppStorageClusterEfficiency{}

// StorageNetAppStorageClusterEfficiency Cluster space efficiency information.
type StorageNetAppStorageClusterEfficiency struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The logical space used for the cluster.
	LogicalUsed *int64 `json:"LogicalUsed,omitempty"`
	// Data reduction ratio (logical_used / used).
	Ratio *float32 `json:"Ratio,omitempty"`
	// Space saved by storage efficiencies (logical_used - used).
	Savings              *int64 `json:"Savings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppStorageClusterEfficiency StorageNetAppStorageClusterEfficiency

// NewStorageNetAppStorageClusterEfficiency instantiates a new StorageNetAppStorageClusterEfficiency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppStorageClusterEfficiency(classId string, objectType string) *StorageNetAppStorageClusterEfficiency {
	this := StorageNetAppStorageClusterEfficiency{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppStorageClusterEfficiencyWithDefaults instantiates a new StorageNetAppStorageClusterEfficiency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppStorageClusterEfficiencyWithDefaults() *StorageNetAppStorageClusterEfficiency {
	this := StorageNetAppStorageClusterEfficiency{}
	var classId string = "storage.NetAppStorageClusterEfficiency"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageClusterEfficiency"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppStorageClusterEfficiency) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageClusterEfficiency) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppStorageClusterEfficiency) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppStorageClusterEfficiency" of the ClassId field.
func (o *StorageNetAppStorageClusterEfficiency) GetDefaultClassId() interface{} {
	return "storage.NetAppStorageClusterEfficiency"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppStorageClusterEfficiency) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageClusterEfficiency) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppStorageClusterEfficiency) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppStorageClusterEfficiency" of the ObjectType field.
func (o *StorageNetAppStorageClusterEfficiency) GetDefaultObjectType() interface{} {
	return "storage.NetAppStorageClusterEfficiency"
}

// GetLogicalUsed returns the LogicalUsed field value if set, zero value otherwise.
func (o *StorageNetAppStorageClusterEfficiency) GetLogicalUsed() int64 {
	if o == nil || IsNil(o.LogicalUsed) {
		var ret int64
		return ret
	}
	return *o.LogicalUsed
}

// GetLogicalUsedOk returns a tuple with the LogicalUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageClusterEfficiency) GetLogicalUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.LogicalUsed) {
		return nil, false
	}
	return o.LogicalUsed, true
}

// HasLogicalUsed returns a boolean if a field has been set.
func (o *StorageNetAppStorageClusterEfficiency) HasLogicalUsed() bool {
	if o != nil && !IsNil(o.LogicalUsed) {
		return true
	}

	return false
}

// SetLogicalUsed gets a reference to the given int64 and assigns it to the LogicalUsed field.
func (o *StorageNetAppStorageClusterEfficiency) SetLogicalUsed(v int64) {
	o.LogicalUsed = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *StorageNetAppStorageClusterEfficiency) GetRatio() float32 {
	if o == nil || IsNil(o.Ratio) {
		var ret float32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageClusterEfficiency) GetRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *StorageNetAppStorageClusterEfficiency) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given float32 and assigns it to the Ratio field.
func (o *StorageNetAppStorageClusterEfficiency) SetRatio(v float32) {
	o.Ratio = &v
}

// GetSavings returns the Savings field value if set, zero value otherwise.
func (o *StorageNetAppStorageClusterEfficiency) GetSavings() int64 {
	if o == nil || IsNil(o.Savings) {
		var ret int64
		return ret
	}
	return *o.Savings
}

// GetSavingsOk returns a tuple with the Savings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageClusterEfficiency) GetSavingsOk() (*int64, bool) {
	if o == nil || IsNil(o.Savings) {
		return nil, false
	}
	return o.Savings, true
}

// HasSavings returns a boolean if a field has been set.
func (o *StorageNetAppStorageClusterEfficiency) HasSavings() bool {
	if o != nil && !IsNil(o.Savings) {
		return true
	}

	return false
}

// SetSavings gets a reference to the given int64 and assigns it to the Savings field.
func (o *StorageNetAppStorageClusterEfficiency) SetSavings(v int64) {
	o.Savings = &v
}

func (o StorageNetAppStorageClusterEfficiency) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppStorageClusterEfficiency) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LogicalUsed) {
		toSerialize["LogicalUsed"] = o.LogicalUsed
	}
	if !IsNil(o.Ratio) {
		toSerialize["Ratio"] = o.Ratio
	}
	if !IsNil(o.Savings) {
		toSerialize["Savings"] = o.Savings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppStorageClusterEfficiency) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The logical space used for the cluster.
		LogicalUsed *int64 `json:"LogicalUsed,omitempty"`
		// Data reduction ratio (logical_used / used).
		Ratio *float32 `json:"Ratio,omitempty"`
		// Space saved by storage efficiencies (logical_used - used).
		Savings *int64 `json:"Savings,omitempty"`
	}

	varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct := StorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppStorageClusterEfficiency := _StorageNetAppStorageClusterEfficiency{}
		varStorageNetAppStorageClusterEfficiency.ClassId = varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct.ClassId
		varStorageNetAppStorageClusterEfficiency.ObjectType = varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct.ObjectType
		varStorageNetAppStorageClusterEfficiency.LogicalUsed = varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct.LogicalUsed
		varStorageNetAppStorageClusterEfficiency.Ratio = varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct.Ratio
		varStorageNetAppStorageClusterEfficiency.Savings = varStorageNetAppStorageClusterEfficiencyWithoutEmbeddedStruct.Savings
		*o = StorageNetAppStorageClusterEfficiency(varStorageNetAppStorageClusterEfficiency)
	} else {
		return err
	}

	varStorageNetAppStorageClusterEfficiency := _StorageNetAppStorageClusterEfficiency{}

	err = json.Unmarshal(data, &varStorageNetAppStorageClusterEfficiency)
	if err == nil {
		o.MoBaseComplexType = varStorageNetAppStorageClusterEfficiency.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LogicalUsed")
		delete(additionalProperties, "Ratio")
		delete(additionalProperties, "Savings")

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

type NullableStorageNetAppStorageClusterEfficiency struct {
	value *StorageNetAppStorageClusterEfficiency
	isSet bool
}

func (v NullableStorageNetAppStorageClusterEfficiency) Get() *StorageNetAppStorageClusterEfficiency {
	return v.value
}

func (v *NullableStorageNetAppStorageClusterEfficiency) Set(val *StorageNetAppStorageClusterEfficiency) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppStorageClusterEfficiency) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppStorageClusterEfficiency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppStorageClusterEfficiency(val *StorageNetAppStorageClusterEfficiency) *NullableStorageNetAppStorageClusterEfficiency {
	return &NullableStorageNetAppStorageClusterEfficiency{value: val, isSet: true}
}

func (v NullableStorageNetAppStorageClusterEfficiency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppStorageClusterEfficiency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
