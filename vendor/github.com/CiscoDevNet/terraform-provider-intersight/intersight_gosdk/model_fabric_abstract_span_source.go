/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18369
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

// checks if the FabricAbstractSpanSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricAbstractSpanSource{}

// FabricAbstractSpanSource Abstract class for all SPAN sources.
type FabricAbstractSpanSource struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Direction of the source SPAN. * `Receive` - SPAN incoming traffic on the SPAN source interface. * `Transmit` - SPAN outgoing traffic on the SPAN source interface. * `Both` - SPAN incoming and outgoing traffic on the SPAN source interface.
	Direction            *string `json:"Direction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricAbstractSpanSource FabricAbstractSpanSource

// NewFabricAbstractSpanSource instantiates a new FabricAbstractSpanSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAbstractSpanSource(classId string, objectType string) *FabricAbstractSpanSource {
	this := FabricAbstractSpanSource{}
	this.ClassId = classId
	this.ObjectType = objectType
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// NewFabricAbstractSpanSourceWithDefaults instantiates a new FabricAbstractSpanSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAbstractSpanSourceWithDefaults() *FabricAbstractSpanSource {
	this := FabricAbstractSpanSource{}
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAbstractSpanSource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAbstractSpanSource) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricAbstractSpanSource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAbstractSpanSource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *FabricAbstractSpanSource) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSource) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *FabricAbstractSpanSource) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *FabricAbstractSpanSource) SetDirection(v string) {
	o.Direction = &v
}

func (o FabricAbstractSpanSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricAbstractSpanSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Direction) {
		toSerialize["Direction"] = o.Direction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricAbstractSpanSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type FabricAbstractSpanSourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Direction of the source SPAN. * `Receive` - SPAN incoming traffic on the SPAN source interface. * `Transmit` - SPAN outgoing traffic on the SPAN source interface. * `Both` - SPAN incoming and outgoing traffic on the SPAN source interface.
		Direction *string `json:"Direction,omitempty"`
	}

	varFabricAbstractSpanSourceWithoutEmbeddedStruct := FabricAbstractSpanSourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourceWithoutEmbeddedStruct)
	if err == nil {
		varFabricAbstractSpanSource := _FabricAbstractSpanSource{}
		varFabricAbstractSpanSource.ClassId = varFabricAbstractSpanSourceWithoutEmbeddedStruct.ClassId
		varFabricAbstractSpanSource.ObjectType = varFabricAbstractSpanSourceWithoutEmbeddedStruct.ObjectType
		varFabricAbstractSpanSource.Direction = varFabricAbstractSpanSourceWithoutEmbeddedStruct.Direction
		*o = FabricAbstractSpanSource(varFabricAbstractSpanSource)
	} else {
		return err
	}

	varFabricAbstractSpanSource := _FabricAbstractSpanSource{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSource)
	if err == nil {
		o.MoBaseMo = varFabricAbstractSpanSource.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Direction")

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

type NullableFabricAbstractSpanSource struct {
	value *FabricAbstractSpanSource
	isSet bool
}

func (v NullableFabricAbstractSpanSource) Get() *FabricAbstractSpanSource {
	return v.value
}

func (v *NullableFabricAbstractSpanSource) Set(val *FabricAbstractSpanSource) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractSpanSource) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractSpanSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractSpanSource(val *FabricAbstractSpanSource) *NullableFabricAbstractSpanSource {
	return &NullableFabricAbstractSpanSource{value: val, isSet: true}
}

func (v NullableFabricAbstractSpanSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractSpanSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}