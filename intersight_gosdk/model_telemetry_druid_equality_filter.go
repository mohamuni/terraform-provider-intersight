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
)

// checks if the TelemetryDruidEqualityFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidEqualityFilter{}

// TelemetryDruidEqualityFilter The equality filter is a replacement for the selector filter with the ability to match against any type of column. The equality filter is designed to have more SQL compatible behavior than the selector filter and so can not match null values. To match null values use the null filter.
type TelemetryDruidEqualityFilter struct {
	Type string `json:"type"`
	// Input column or virtual column name to filter.
	Column string `json:"column"`
	// String specifying the type of value to match. For example STRING, LONG, DOUBLE, FLOAT, ARRAY<STRING>, ARRAY<LONG>, or any other Druid type. The matchValueType determines how Druid interprets the matchValue to assist in converting to the type of the matched column.
	MatchValueType string `json:"matchValueType"`
	// Value to match, must not be null.
	MatchValue           interface{} `json:"matchValue"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidEqualityFilter TelemetryDruidEqualityFilter

// NewTelemetryDruidEqualityFilter instantiates a new TelemetryDruidEqualityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidEqualityFilter(type_ string, column string, matchValueType string, matchValue interface{}) *TelemetryDruidEqualityFilter {
	this := TelemetryDruidEqualityFilter{}
	this.Type = type_
	this.Column = column
	this.MatchValueType = matchValueType
	this.MatchValue = matchValue
	return &this
}

// NewTelemetryDruidEqualityFilterWithDefaults instantiates a new TelemetryDruidEqualityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidEqualityFilterWithDefaults() *TelemetryDruidEqualityFilter {
	this := TelemetryDruidEqualityFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidEqualityFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidEqualityFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidEqualityFilter) SetType(v string) {
	o.Type = v
}

// GetColumn returns the Column field value
func (o *TelemetryDruidEqualityFilter) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidEqualityFilter) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value
func (o *TelemetryDruidEqualityFilter) SetColumn(v string) {
	o.Column = v
}

// GetMatchValueType returns the MatchValueType field value
func (o *TelemetryDruidEqualityFilter) GetMatchValueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchValueType
}

// GetMatchValueTypeOk returns a tuple with the MatchValueType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidEqualityFilter) GetMatchValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchValueType, true
}

// SetMatchValueType sets field value
func (o *TelemetryDruidEqualityFilter) SetMatchValueType(v string) {
	o.MatchValueType = v
}

// GetMatchValue returns the MatchValue field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TelemetryDruidEqualityFilter) GetMatchValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MatchValue
}

// GetMatchValueOk returns a tuple with the MatchValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelemetryDruidEqualityFilter) GetMatchValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MatchValue) {
		return nil, false
	}
	return &o.MatchValue, true
}

// SetMatchValue sets field value
func (o *TelemetryDruidEqualityFilter) SetMatchValue(v interface{}) {
	o.MatchValue = v
}

func (o TelemetryDruidEqualityFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidEqualityFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["column"] = o.Column
	toSerialize["matchValueType"] = o.MatchValueType
	if o.MatchValue != nil {
		toSerialize["matchValue"] = o.MatchValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidEqualityFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"column",
		"matchValueType",
		"matchValue",
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
	varTelemetryDruidEqualityFilter := _TelemetryDruidEqualityFilter{}

	err = json.Unmarshal(data, &varTelemetryDruidEqualityFilter)

	if err != nil {
		return err
	}

	*o = TelemetryDruidEqualityFilter(varTelemetryDruidEqualityFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "column")
		delete(additionalProperties, "matchValueType")
		delete(additionalProperties, "matchValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidEqualityFilter struct {
	value *TelemetryDruidEqualityFilter
	isSet bool
}

func (v NullableTelemetryDruidEqualityFilter) Get() *TelemetryDruidEqualityFilter {
	return v.value
}

func (v *NullableTelemetryDruidEqualityFilter) Set(val *TelemetryDruidEqualityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidEqualityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidEqualityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidEqualityFilter(val *TelemetryDruidEqualityFilter) *NullableTelemetryDruidEqualityFilter {
	return &NullableTelemetryDruidEqualityFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidEqualityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidEqualityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
