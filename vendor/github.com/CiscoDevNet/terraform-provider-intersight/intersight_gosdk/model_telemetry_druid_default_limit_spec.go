/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidDefaultLimitSpec The default limit spec takes a limit and the list of columns to do an orderBy operation over.
type TelemetryDruidDefaultLimitSpec struct {
	// The limit spec type.
	Type string `json:"type"`
	// How many rows to return. If not specified, all rows will be returned.
	Limit                int32                             `json:"limit"`
	Columns              []TelemetryDruidOrderByColumnSpec `json:"columns"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidDefaultLimitSpec TelemetryDruidDefaultLimitSpec

// NewTelemetryDruidDefaultLimitSpec instantiates a new TelemetryDruidDefaultLimitSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDefaultLimitSpec(type_ string, limit int32, columns []TelemetryDruidOrderByColumnSpec) *TelemetryDruidDefaultLimitSpec {
	this := TelemetryDruidDefaultLimitSpec{}
	this.Type = type_
	this.Limit = limit
	this.Columns = columns
	return &this
}

// NewTelemetryDruidDefaultLimitSpecWithDefaults instantiates a new TelemetryDruidDefaultLimitSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDefaultLimitSpecWithDefaults() *TelemetryDruidDefaultLimitSpec {
	this := TelemetryDruidDefaultLimitSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidDefaultLimitSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidDefaultLimitSpec) SetType(v string) {
	o.Type = v
}

// GetLimit returns the Limit field value
func (o *TelemetryDruidDefaultLimitSpec) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpec) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *TelemetryDruidDefaultLimitSpec) SetLimit(v int32) {
	o.Limit = v
}

// GetColumns returns the Columns field value
func (o *TelemetryDruidDefaultLimitSpec) GetColumns() []TelemetryDruidOrderByColumnSpec {
	if o == nil {
		var ret []TelemetryDruidOrderByColumnSpec
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpec) GetColumnsOk() ([]TelemetryDruidOrderByColumnSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *TelemetryDruidDefaultLimitSpec) SetColumns(v []TelemetryDruidOrderByColumnSpec) {
	o.Columns = v
}

func (o TelemetryDruidDefaultLimitSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["columns"] = o.Columns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidDefaultLimitSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidDefaultLimitSpec := _TelemetryDruidDefaultLimitSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidDefaultLimitSpec); err == nil {
		*o = TelemetryDruidDefaultLimitSpec(varTelemetryDruidDefaultLimitSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "columns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidDefaultLimitSpec struct {
	value *TelemetryDruidDefaultLimitSpec
	isSet bool
}

func (v NullableTelemetryDruidDefaultLimitSpec) Get() *TelemetryDruidDefaultLimitSpec {
	return v.value
}

func (v *NullableTelemetryDruidDefaultLimitSpec) Set(val *TelemetryDruidDefaultLimitSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDefaultLimitSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDefaultLimitSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDefaultLimitSpec(val *TelemetryDruidDefaultLimitSpec) *NullableTelemetryDruidDefaultLimitSpec {
	return &NullableTelemetryDruidDefaultLimitSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidDefaultLimitSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDefaultLimitSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
