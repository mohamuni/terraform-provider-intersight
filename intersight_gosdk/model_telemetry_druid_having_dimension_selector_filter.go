/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidHavingDimensionSelectorFilter The dimSelector filter will match rows with dimension values equal to the specified value.
type TelemetryDruidHavingDimensionSelectorFilter struct {
	// The having filter type.
	Type string `json:"type"`
	// dimension
	Dimension string `json:"dimension"`
	// null
	Value                float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidHavingDimensionSelectorFilter TelemetryDruidHavingDimensionSelectorFilter

// NewTelemetryDruidHavingDimensionSelectorFilter instantiates a new TelemetryDruidHavingDimensionSelectorFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidHavingDimensionSelectorFilter(type_ string, dimension string, value float64) *TelemetryDruidHavingDimensionSelectorFilter {
	this := TelemetryDruidHavingDimensionSelectorFilter{}
	this.Type = type_
	this.Dimension = dimension
	this.Value = value
	return &this
}

// NewTelemetryDruidHavingDimensionSelectorFilterWithDefaults instantiates a new TelemetryDruidHavingDimensionSelectorFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidHavingDimensionSelectorFilterWithDefaults() *TelemetryDruidHavingDimensionSelectorFilter {
	this := TelemetryDruidHavingDimensionSelectorFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidHavingDimensionSelectorFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingDimensionSelectorFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidHavingDimensionSelectorFilter) SetType(v string) {
	o.Type = v
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidHavingDimensionSelectorFilter) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingDimensionSelectorFilter) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidHavingDimensionSelectorFilter) SetDimension(v string) {
	o.Dimension = v
}

// GetValue returns the Value field value
func (o *TelemetryDruidHavingDimensionSelectorFilter) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingDimensionSelectorFilter) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidHavingDimensionSelectorFilter) SetValue(v float64) {
	o.Value = v
}

func (o TelemetryDruidHavingDimensionSelectorFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidHavingDimensionSelectorFilter) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidHavingDimensionSelectorFilter := _TelemetryDruidHavingDimensionSelectorFilter{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidHavingDimensionSelectorFilter); err == nil {
		*o = TelemetryDruidHavingDimensionSelectorFilter(varTelemetryDruidHavingDimensionSelectorFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidHavingDimensionSelectorFilter struct {
	value *TelemetryDruidHavingDimensionSelectorFilter
	isSet bool
}

func (v NullableTelemetryDruidHavingDimensionSelectorFilter) Get() *TelemetryDruidHavingDimensionSelectorFilter {
	return v.value
}

func (v *NullableTelemetryDruidHavingDimensionSelectorFilter) Set(val *TelemetryDruidHavingDimensionSelectorFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidHavingDimensionSelectorFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidHavingDimensionSelectorFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidHavingDimensionSelectorFilter(val *TelemetryDruidHavingDimensionSelectorFilter) *NullableTelemetryDruidHavingDimensionSelectorFilter {
	return &NullableTelemetryDruidHavingDimensionSelectorFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidHavingDimensionSelectorFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidHavingDimensionSelectorFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}