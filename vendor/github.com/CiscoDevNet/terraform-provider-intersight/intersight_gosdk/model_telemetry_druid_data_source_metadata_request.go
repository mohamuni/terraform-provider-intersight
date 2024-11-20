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

// checks if the TelemetryDruidDataSourceMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidDataSourceMetadataRequest{}

// TelemetryDruidDataSourceMetadataRequest Data Source Metadata queries return metadata information for a dataSource. These queries return information about the timestamp of latest ingested event for the dataSource. This is the ingested event without any consideration of rollup.
type TelemetryDruidDataSourceMetadataRequest struct {
	QueryType            string                      `json:"queryType"`
	DataSource           TelemetryDruidDataSource    `json:"dataSource"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidDataSourceMetadataRequest TelemetryDruidDataSourceMetadataRequest

// NewTelemetryDruidDataSourceMetadataRequest instantiates a new TelemetryDruidDataSourceMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDataSourceMetadataRequest(queryType string, dataSource TelemetryDruidDataSource) *TelemetryDruidDataSourceMetadataRequest {
	this := TelemetryDruidDataSourceMetadataRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	return &this
}

// NewTelemetryDruidDataSourceMetadataRequestWithDefaults instantiates a new TelemetryDruidDataSourceMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDataSourceMetadataRequestWithDefaults() *TelemetryDruidDataSourceMetadataRequest {
	this := TelemetryDruidDataSourceMetadataRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidDataSourceMetadataRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDataSourceMetadataRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidDataSourceMetadataRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidDataSourceMetadataRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDataSourceMetadataRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidDataSourceMetadataRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidDataSourceMetadataRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || IsNil(o.Context) {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDataSourceMetadataRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidDataSourceMetadataRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidDataSourceMetadataRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidDataSourceMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidDataSourceMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryType"] = o.QueryType
	toSerialize["dataSource"] = o.DataSource
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidDataSourceMetadataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryType",
		"dataSource",
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
	varTelemetryDruidDataSourceMetadataRequest := _TelemetryDruidDataSourceMetadataRequest{}

	err = json.Unmarshal(data, &varTelemetryDruidDataSourceMetadataRequest)

	if err != nil {
		return err
	}

	*o = TelemetryDruidDataSourceMetadataRequest(varTelemetryDruidDataSourceMetadataRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidDataSourceMetadataRequest struct {
	value *TelemetryDruidDataSourceMetadataRequest
	isSet bool
}

func (v NullableTelemetryDruidDataSourceMetadataRequest) Get() *TelemetryDruidDataSourceMetadataRequest {
	return v.value
}

func (v *NullableTelemetryDruidDataSourceMetadataRequest) Set(val *TelemetryDruidDataSourceMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDataSourceMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDataSourceMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDataSourceMetadataRequest(val *TelemetryDruidDataSourceMetadataRequest) *NullableTelemetryDruidDataSourceMetadataRequest {
	return &NullableTelemetryDruidDataSourceMetadataRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidDataSourceMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDataSourceMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
