/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidUnionDataSource This data source unions two or more table data sources. Note that the data sources being unioned should have the same schema.
type TelemetryDruidUnionDataSource struct {
	// The type of data source.
	Type string `json:"type"`
	// A list of data sources.
	DataSources          []string `json:"dataSources"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidUnionDataSource TelemetryDruidUnionDataSource

// NewTelemetryDruidUnionDataSource instantiates a new TelemetryDruidUnionDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidUnionDataSource(type_ string, dataSources []string) *TelemetryDruidUnionDataSource {
	this := TelemetryDruidUnionDataSource{}
	this.Type = type_
	this.DataSources = dataSources
	return &this
}

// NewTelemetryDruidUnionDataSourceWithDefaults instantiates a new TelemetryDruidUnionDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidUnionDataSourceWithDefaults() *TelemetryDruidUnionDataSource {
	this := TelemetryDruidUnionDataSource{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidUnionDataSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidUnionDataSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidUnionDataSource) SetType(v string) {
	o.Type = v
}

// GetDataSources returns the DataSources field value
func (o *TelemetryDruidUnionDataSource) GetDataSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DataSources
}

// GetDataSourcesOk returns a tuple with the DataSources field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidUnionDataSource) GetDataSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSources, true
}

// SetDataSources sets field value
func (o *TelemetryDruidUnionDataSource) SetDataSources(v []string) {
	o.DataSources = v
}

func (o TelemetryDruidUnionDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["dataSources"] = o.DataSources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidUnionDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidUnionDataSource := _TelemetryDruidUnionDataSource{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidUnionDataSource); err == nil {
		*o = TelemetryDruidUnionDataSource(varTelemetryDruidUnionDataSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "dataSources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidUnionDataSource struct {
	value *TelemetryDruidUnionDataSource
	isSet bool
}

func (v NullableTelemetryDruidUnionDataSource) Get() *TelemetryDruidUnionDataSource {
	return v.value
}

func (v *NullableTelemetryDruidUnionDataSource) Set(val *TelemetryDruidUnionDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidUnionDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidUnionDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidUnionDataSource(val *TelemetryDruidUnionDataSource) *NullableTelemetryDruidUnionDataSource {
	return &NullableTelemetryDruidUnionDataSource{value: val, isSet: true}
}

func (v NullableTelemetryDruidUnionDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidUnionDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
