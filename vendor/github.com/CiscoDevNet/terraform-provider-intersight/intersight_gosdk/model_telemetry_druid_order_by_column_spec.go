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
)

// checks if the TelemetryDruidOrderByColumnSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidOrderByColumnSpec{}

// TelemetryDruidOrderByColumnSpec OrderByColumnSpecs indicate how to do order by operations. Each order-by condition can be a jsonString or a map.
type TelemetryDruidOrderByColumnSpec struct {
	// Any dimension or metric name.
	Dimension            *string `json:"dimension,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	DimensionOrder       *string `json:"dimensionOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidOrderByColumnSpec TelemetryDruidOrderByColumnSpec

// NewTelemetryDruidOrderByColumnSpec instantiates a new TelemetryDruidOrderByColumnSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidOrderByColumnSpec() *TelemetryDruidOrderByColumnSpec {
	this := TelemetryDruidOrderByColumnSpec{}
	return &this
}

// NewTelemetryDruidOrderByColumnSpecWithDefaults instantiates a new TelemetryDruidOrderByColumnSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidOrderByColumnSpecWithDefaults() *TelemetryDruidOrderByColumnSpec {
	this := TelemetryDruidOrderByColumnSpec{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *TelemetryDruidOrderByColumnSpec) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidOrderByColumnSpec) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *TelemetryDruidOrderByColumnSpec) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *TelemetryDruidOrderByColumnSpec) SetDimension(v string) {
	o.Dimension = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *TelemetryDruidOrderByColumnSpec) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidOrderByColumnSpec) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *TelemetryDruidOrderByColumnSpec) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *TelemetryDruidOrderByColumnSpec) SetDirection(v string) {
	o.Direction = &v
}

// GetDimensionOrder returns the DimensionOrder field value if set, zero value otherwise.
func (o *TelemetryDruidOrderByColumnSpec) GetDimensionOrder() string {
	if o == nil || IsNil(o.DimensionOrder) {
		var ret string
		return ret
	}
	return *o.DimensionOrder
}

// GetDimensionOrderOk returns a tuple with the DimensionOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidOrderByColumnSpec) GetDimensionOrderOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionOrder) {
		return nil, false
	}
	return o.DimensionOrder, true
}

// HasDimensionOrder returns a boolean if a field has been set.
func (o *TelemetryDruidOrderByColumnSpec) HasDimensionOrder() bool {
	if o != nil && !IsNil(o.DimensionOrder) {
		return true
	}

	return false
}

// SetDimensionOrder gets a reference to the given string and assigns it to the DimensionOrder field.
func (o *TelemetryDruidOrderByColumnSpec) SetDimensionOrder(v string) {
	o.DimensionOrder = &v
}

func (o TelemetryDruidOrderByColumnSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidOrderByColumnSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.DimensionOrder) {
		toSerialize["dimensionOrder"] = o.DimensionOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidOrderByColumnSpec) UnmarshalJSON(data []byte) (err error) {
	varTelemetryDruidOrderByColumnSpec := _TelemetryDruidOrderByColumnSpec{}

	err = json.Unmarshal(data, &varTelemetryDruidOrderByColumnSpec)

	if err != nil {
		return err
	}

	*o = TelemetryDruidOrderByColumnSpec(varTelemetryDruidOrderByColumnSpec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "dimensionOrder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidOrderByColumnSpec struct {
	value *TelemetryDruidOrderByColumnSpec
	isSet bool
}

func (v NullableTelemetryDruidOrderByColumnSpec) Get() *TelemetryDruidOrderByColumnSpec {
	return v.value
}

func (v *NullableTelemetryDruidOrderByColumnSpec) Set(val *TelemetryDruidOrderByColumnSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidOrderByColumnSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidOrderByColumnSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidOrderByColumnSpec(val *TelemetryDruidOrderByColumnSpec) *NullableTelemetryDruidOrderByColumnSpec {
	return &NullableTelemetryDruidOrderByColumnSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidOrderByColumnSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidOrderByColumnSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
