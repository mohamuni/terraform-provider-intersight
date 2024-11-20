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

// checks if the TelemetryDruidQuerySpecFragment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidQuerySpecFragment{}

// TelemetryDruidQuerySpecFragment A 'fragment' query spec.
type TelemetryDruidQuerySpecFragment struct {
	Type string `json:"type"`
	// A JSON array of string values to search.
	Values []string `json:"values"`
	// Whether the string comparison is case-sensitive or not.
	CaseSensitive        *bool `json:"caseSensitive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidQuerySpecFragment TelemetryDruidQuerySpecFragment

// NewTelemetryDruidQuerySpecFragment instantiates a new TelemetryDruidQuerySpecFragment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidQuerySpecFragment(type_ string, values []string) *TelemetryDruidQuerySpecFragment {
	this := TelemetryDruidQuerySpecFragment{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewTelemetryDruidQuerySpecFragmentWithDefaults instantiates a new TelemetryDruidQuerySpecFragment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidQuerySpecFragmentWithDefaults() *TelemetryDruidQuerySpecFragment {
	this := TelemetryDruidQuerySpecFragment{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidQuerySpecFragment) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQuerySpecFragment) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidQuerySpecFragment) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *TelemetryDruidQuerySpecFragment) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQuerySpecFragment) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *TelemetryDruidQuerySpecFragment) SetValues(v []string) {
	o.Values = v
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *TelemetryDruidQuerySpecFragment) GetCaseSensitive() bool {
	if o == nil || IsNil(o.CaseSensitive) {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQuerySpecFragment) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseSensitive) {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *TelemetryDruidQuerySpecFragment) HasCaseSensitive() bool {
	if o != nil && !IsNil(o.CaseSensitive) {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *TelemetryDruidQuerySpecFragment) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

func (o TelemetryDruidQuerySpecFragment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidQuerySpecFragment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values
	if !IsNil(o.CaseSensitive) {
		toSerialize["caseSensitive"] = o.CaseSensitive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidQuerySpecFragment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"values",
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
	varTelemetryDruidQuerySpecFragment := _TelemetryDruidQuerySpecFragment{}

	err = json.Unmarshal(data, &varTelemetryDruidQuerySpecFragment)

	if err != nil {
		return err
	}

	*o = TelemetryDruidQuerySpecFragment(varTelemetryDruidQuerySpecFragment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		delete(additionalProperties, "caseSensitive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidQuerySpecFragment struct {
	value *TelemetryDruidQuerySpecFragment
	isSet bool
}

func (v NullableTelemetryDruidQuerySpecFragment) Get() *TelemetryDruidQuerySpecFragment {
	return v.value
}

func (v *NullableTelemetryDruidQuerySpecFragment) Set(val *TelemetryDruidQuerySpecFragment) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidQuerySpecFragment) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidQuerySpecFragment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidQuerySpecFragment(val *TelemetryDruidQuerySpecFragment) *NullableTelemetryDruidQuerySpecFragment {
	return &NullableTelemetryDruidQuerySpecFragment{value: val, isSet: true}
}

func (v NullableTelemetryDruidQuerySpecFragment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidQuerySpecFragment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
