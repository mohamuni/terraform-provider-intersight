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
	"fmt"
)

// checks if the TelemetryDruidSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidSearchRequest{}

// TelemetryDruidSearchRequest These types of queries returns dimension values that match the search specification
type TelemetryDruidSearchRequest struct {
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals   []string                  `json:"intervals"`
	Granularity TelemetryDruidGranularity `json:"granularity"`
	Filter      *TelemetryDruidFilter     `json:"filter,omitempty"`
	// Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket.
	Aggregations []TelemetryDruidAggregator `json:"aggregations,omitempty"`
	// The list of dimensions to run the search over. Excluding this means the search is run over all dimensions.
	SearchDimensions []string                           `json:"searchDimensions,omitempty"`
	Query            *TelemetryDruidAggregateSearchSpec `json:"query,omitempty"`
	// An integer that limits the number of results. The default is unlimited.
	Limit                *int32                      `json:"limit,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidSearchRequest TelemetryDruidSearchRequest

// NewTelemetryDruidSearchRequest instantiates a new TelemetryDruidSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidSearchRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity) *TelemetryDruidSearchRequest {
	this := TelemetryDruidSearchRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	this.Intervals = intervals
	this.Granularity = granularity
	return &this
}

// NewTelemetryDruidSearchRequestWithDefaults instantiates a new TelemetryDruidSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidSearchRequestWithDefaults() *TelemetryDruidSearchRequest {
	this := TelemetryDruidSearchRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidSearchRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidSearchRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidSearchRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidSearchRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidSearchRequest) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetIntervalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidSearchRequest) SetIntervals(v []string) {
	o.Intervals = v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidSearchRequest) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidSearchRequest) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || IsNil(o.Filter) {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidSearchRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequest) GetAggregations() []TelemetryDruidAggregator {
	if o == nil || IsNil(o.Aggregations) {
		var ret []TelemetryDruidAggregator
		return ret
	}
	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetAggregationsOk() ([]TelemetryDruidAggregator, bool) {
	if o == nil || IsNil(o.Aggregations) {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequest) HasAggregations() bool {
	if o != nil && !IsNil(o.Aggregations) {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidSearchRequest) SetAggregations(v []TelemetryDruidAggregator) {
	o.Aggregations = v
}

// GetSearchDimensions returns the SearchDimensions field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequest) GetSearchDimensions() []string {
	if o == nil || IsNil(o.SearchDimensions) {
		var ret []string
		return ret
	}
	return o.SearchDimensions
}

// GetSearchDimensionsOk returns a tuple with the SearchDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetSearchDimensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchDimensions) {
		return nil, false
	}
	return o.SearchDimensions, true
}

// HasSearchDimensions returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequest) HasSearchDimensions() bool {
	if o != nil && !IsNil(o.SearchDimensions) {
		return true
	}

	return false
}

// SetSearchDimensions gets a reference to the given []string and assigns it to the SearchDimensions field.
func (o *TelemetryDruidSearchRequest) SetSearchDimensions(v []string) {
	o.SearchDimensions = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequest) GetQuery() TelemetryDruidAggregateSearchSpec {
	if o == nil || IsNil(o.Query) {
		var ret TelemetryDruidAggregateSearchSpec
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetQueryOk() (*TelemetryDruidAggregateSearchSpec, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given TelemetryDruidAggregateSearchSpec and assigns it to the Query field.
func (o *TelemetryDruidSearchRequest) SetQuery(v TelemetryDruidAggregateSearchSpec) {
	o.Query = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TelemetryDruidSearchRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || IsNil(o.Context) {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidSearchRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryType"] = o.QueryType
	toSerialize["dataSource"] = o.DataSource
	toSerialize["intervals"] = o.Intervals
	toSerialize["granularity"] = o.Granularity
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Aggregations) {
		toSerialize["aggregations"] = o.Aggregations
	}
	if !IsNil(o.SearchDimensions) {
		toSerialize["searchDimensions"] = o.SearchDimensions
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidSearchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryType",
		"dataSource",
		"intervals",
		"granularity",
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
	varTelemetryDruidSearchRequest := _TelemetryDruidSearchRequest{}

	err = json.Unmarshal(data, &varTelemetryDruidSearchRequest)

	if err != nil {
		return err
	}

	*o = TelemetryDruidSearchRequest(varTelemetryDruidSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "searchDimensions")
		delete(additionalProperties, "query")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidSearchRequest struct {
	value *TelemetryDruidSearchRequest
	isSet bool
}

func (v NullableTelemetryDruidSearchRequest) Get() *TelemetryDruidSearchRequest {
	return v.value
}

func (v *NullableTelemetryDruidSearchRequest) Set(val *TelemetryDruidSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidSearchRequest(val *TelemetryDruidSearchRequest) *NullableTelemetryDruidSearchRequest {
	return &NullableTelemetryDruidSearchRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
