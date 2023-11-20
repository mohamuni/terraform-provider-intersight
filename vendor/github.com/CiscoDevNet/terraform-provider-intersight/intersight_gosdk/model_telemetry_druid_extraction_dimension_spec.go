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

// TelemetryDruidExtractionDimensionSpec Returns dimension values transformed using the given extraction function.
type TelemetryDruidExtractionDimensionSpec struct {
	// the dimension spec type.
	Type       string `json:"type"`
	Dimension  string `json:"dimension"`
	OutputName string `json:"outputName"`
	OutputType string `json:"outputType"`
	// All filters except the \"spatial\" filter support extraction functions. An extraction function is defined by setting the \"extractionFn\" field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \"bar_1\".
	ExtractionFn         map[string]interface{} `json:"extractionFn"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionDimensionSpec TelemetryDruidExtractionDimensionSpec

// NewTelemetryDruidExtractionDimensionSpec instantiates a new TelemetryDruidExtractionDimensionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionDimensionSpec(type_ string, dimension string, outputName string, outputType string, extractionFn map[string]interface{}) *TelemetryDruidExtractionDimensionSpec {
	this := TelemetryDruidExtractionDimensionSpec{}
	this.Type = type_
	this.Dimension = dimension
	this.OutputName = outputName
	this.OutputType = outputType
	this.ExtractionFn = extractionFn
	return &this
}

// NewTelemetryDruidExtractionDimensionSpecWithDefaults instantiates a new TelemetryDruidExtractionDimensionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionDimensionSpecWithDefaults() *TelemetryDruidExtractionDimensionSpec {
	this := TelemetryDruidExtractionDimensionSpec{}
	var outputType string = "STRING"
	this.OutputType = outputType
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionDimensionSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionDimensionSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionDimensionSpec) SetType(v string) {
	o.Type = v
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidExtractionDimensionSpec) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionDimensionSpec) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidExtractionDimensionSpec) SetDimension(v string) {
	o.Dimension = v
}

// GetOutputName returns the OutputName field value
func (o *TelemetryDruidExtractionDimensionSpec) GetOutputName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputName
}

// GetOutputNameOk returns a tuple with the OutputName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionDimensionSpec) GetOutputNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputName, true
}

// SetOutputName sets field value
func (o *TelemetryDruidExtractionDimensionSpec) SetOutputName(v string) {
	o.OutputName = v
}

// GetOutputType returns the OutputType field value
func (o *TelemetryDruidExtractionDimensionSpec) GetOutputType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputType
}

// GetOutputTypeOk returns a tuple with the OutputType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionDimensionSpec) GetOutputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputType, true
}

// SetOutputType sets field value
func (o *TelemetryDruidExtractionDimensionSpec) SetOutputType(v string) {
	o.OutputType = v
}

// GetExtractionFn returns the ExtractionFn field value
func (o *TelemetryDruidExtractionDimensionSpec) GetExtractionFn() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtractionFn
}

// GetExtractionFnOk returns a tuple with the ExtractionFn field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionDimensionSpec) GetExtractionFnOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtractionFn, true
}

// SetExtractionFn sets field value
func (o *TelemetryDruidExtractionDimensionSpec) SetExtractionFn(v map[string]interface{}) {
	o.ExtractionFn = v
}

func (o TelemetryDruidExtractionDimensionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["outputName"] = o.OutputName
	}
	if true {
		toSerialize["outputType"] = o.OutputType
	}
	if true {
		toSerialize["extractionFn"] = o.ExtractionFn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidExtractionDimensionSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidExtractionDimensionSpec := _TelemetryDruidExtractionDimensionSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidExtractionDimensionSpec); err == nil {
		*o = TelemetryDruidExtractionDimensionSpec(varTelemetryDruidExtractionDimensionSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "outputName")
		delete(additionalProperties, "outputType")
		delete(additionalProperties, "extractionFn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionDimensionSpec struct {
	value *TelemetryDruidExtractionDimensionSpec
	isSet bool
}

func (v NullableTelemetryDruidExtractionDimensionSpec) Get() *TelemetryDruidExtractionDimensionSpec {
	return v.value
}

func (v *NullableTelemetryDruidExtractionDimensionSpec) Set(val *TelemetryDruidExtractionDimensionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionDimensionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionDimensionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionDimensionSpec(val *TelemetryDruidExtractionDimensionSpec) *NullableTelemetryDruidExtractionDimensionSpec {
	return &NullableTelemetryDruidExtractionDimensionSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionDimensionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionDimensionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
