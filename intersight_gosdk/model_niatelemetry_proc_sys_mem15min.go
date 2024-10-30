/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the NiatelemetryProcSysMem15min type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryProcSysMem15min{}

// NiatelemetryProcSysMem15min Aci node memory usage info in last 15 mintutes.
type NiatelemetryProcSysMem15min struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the free average value.
	FreeAvg *string `json:"FreeAvg,omitempty"`
	// Returns the total average value.
	TotalAvg             *string `json:"TotalAvg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryProcSysMem15min NiatelemetryProcSysMem15min

// NewNiatelemetryProcSysMem15min instantiates a new NiatelemetryProcSysMem15min object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryProcSysMem15min(classId string, objectType string) *NiatelemetryProcSysMem15min {
	this := NiatelemetryProcSysMem15min{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryProcSysMem15minWithDefaults instantiates a new NiatelemetryProcSysMem15min object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryProcSysMem15minWithDefaults() *NiatelemetryProcSysMem15min {
	this := NiatelemetryProcSysMem15min{}
	var classId string = "niatelemetry.ProcSysMem15min"
	this.ClassId = classId
	var objectType string = "niatelemetry.ProcSysMem15min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryProcSysMem15min) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15min) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryProcSysMem15min) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.ProcSysMem15min" of the ClassId field.
func (o *NiatelemetryProcSysMem15min) GetDefaultClassId() interface{} {
	return "niatelemetry.ProcSysMem15min"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryProcSysMem15min) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15min) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryProcSysMem15min) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.ProcSysMem15min" of the ObjectType field.
func (o *NiatelemetryProcSysMem15min) GetDefaultObjectType() interface{} {
	return "niatelemetry.ProcSysMem15min"
}

// GetFreeAvg returns the FreeAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysMem15min) GetFreeAvg() string {
	if o == nil || IsNil(o.FreeAvg) {
		var ret string
		return ret
	}
	return *o.FreeAvg
}

// GetFreeAvgOk returns a tuple with the FreeAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15min) GetFreeAvgOk() (*string, bool) {
	if o == nil || IsNil(o.FreeAvg) {
		return nil, false
	}
	return o.FreeAvg, true
}

// HasFreeAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysMem15min) HasFreeAvg() bool {
	if o != nil && !IsNil(o.FreeAvg) {
		return true
	}

	return false
}

// SetFreeAvg gets a reference to the given string and assigns it to the FreeAvg field.
func (o *NiatelemetryProcSysMem15min) SetFreeAvg(v string) {
	o.FreeAvg = &v
}

// GetTotalAvg returns the TotalAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysMem15min) GetTotalAvg() string {
	if o == nil || IsNil(o.TotalAvg) {
		var ret string
		return ret
	}
	return *o.TotalAvg
}

// GetTotalAvgOk returns a tuple with the TotalAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15min) GetTotalAvgOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAvg) {
		return nil, false
	}
	return o.TotalAvg, true
}

// HasTotalAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysMem15min) HasTotalAvg() bool {
	if o != nil && !IsNil(o.TotalAvg) {
		return true
	}

	return false
}

// SetTotalAvg gets a reference to the given string and assigns it to the TotalAvg field.
func (o *NiatelemetryProcSysMem15min) SetTotalAvg(v string) {
	o.TotalAvg = &v
}

func (o NiatelemetryProcSysMem15min) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryProcSysMem15min) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FreeAvg) {
		toSerialize["FreeAvg"] = o.FreeAvg
	}
	if !IsNil(o.TotalAvg) {
		toSerialize["TotalAvg"] = o.TotalAvg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryProcSysMem15min) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryProcSysMem15minWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns the free average value.
		FreeAvg *string `json:"FreeAvg,omitempty"`
		// Returns the total average value.
		TotalAvg *string `json:"TotalAvg,omitempty"`
	}

	varNiatelemetryProcSysMem15minWithoutEmbeddedStruct := NiatelemetryProcSysMem15minWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryProcSysMem15minWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryProcSysMem15min := _NiatelemetryProcSysMem15min{}
		varNiatelemetryProcSysMem15min.ClassId = varNiatelemetryProcSysMem15minWithoutEmbeddedStruct.ClassId
		varNiatelemetryProcSysMem15min.ObjectType = varNiatelemetryProcSysMem15minWithoutEmbeddedStruct.ObjectType
		varNiatelemetryProcSysMem15min.FreeAvg = varNiatelemetryProcSysMem15minWithoutEmbeddedStruct.FreeAvg
		varNiatelemetryProcSysMem15min.TotalAvg = varNiatelemetryProcSysMem15minWithoutEmbeddedStruct.TotalAvg
		*o = NiatelemetryProcSysMem15min(varNiatelemetryProcSysMem15min)
	} else {
		return err
	}

	varNiatelemetryProcSysMem15min := _NiatelemetryProcSysMem15min{}

	err = json.Unmarshal(data, &varNiatelemetryProcSysMem15min)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryProcSysMem15min.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FreeAvg")
		delete(additionalProperties, "TotalAvg")

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

type NullableNiatelemetryProcSysMem15min struct {
	value *NiatelemetryProcSysMem15min
	isSet bool
}

func (v NullableNiatelemetryProcSysMem15min) Get() *NiatelemetryProcSysMem15min {
	return v.value
}

func (v *NullableNiatelemetryProcSysMem15min) Set(val *NiatelemetryProcSysMem15min) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryProcSysMem15min) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryProcSysMem15min) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryProcSysMem15min(val *NiatelemetryProcSysMem15min) *NullableNiatelemetryProcSysMem15min {
	return &NullableNiatelemetryProcSysMem15min{value: val, isSet: true}
}

func (v NullableNiatelemetryProcSysMem15min) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryProcSysMem15min) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
