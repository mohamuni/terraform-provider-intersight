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

// TelemetryDruidGroupByRequest These types of Apache Druid queries take a groupBy query object and return an array of JSON objects where each object represents a grouping asked for by the query.
type TelemetryDruidGroupByRequest struct {
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON list of dimensions to do the groupBy over; or see DimensionSpec for ways to extract dimensions..
	Dimensions  []TelemetryDruidDimensionSpec   `json:"dimensions"`
	LimitSpec   *TelemetryDruidDefaultLimitSpec `json:"limitSpec,omitempty"`
	Having      *TelemetryDruidHavingFilter     `json:"having,omitempty"`
	Granularity TelemetryDruidGranularity       `json:"granularity"`
	Filter      *TelemetryDruidFilter           `json:"filter,omitempty"`
	// Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket.
	Aggregations []TelemetryDruidAggregator `json:"aggregations,omitempty"`
	// Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires.
	PostAggregations []TelemetryDruidPostAggregator `json:"postAggregations,omitempty"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals []string `json:"intervals"`
	// A JSON array of arrays to return additional result sets for groupings of subsets of top level dimensions. The subtotals feature allows computation of multiple sub-groupings in a single query. To use this feature, add a \"subtotalsSpec\" to your query, which should be a list of subgroup dimension sets. It should contain the \"outputName\" from dimensions in your \"dimensions\" attribute, in the same order as they appear in the \"dimensions\" attribute.
	SubtotalsSpec        map[string]interface{}      `json:"subtotalsSpec,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidGroupByRequest TelemetryDruidGroupByRequest

// NewTelemetryDruidGroupByRequest instantiates a new TelemetryDruidGroupByRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidGroupByRequest(queryType string, dataSource TelemetryDruidDataSource, dimensions []TelemetryDruidDimensionSpec, granularity TelemetryDruidGranularity, intervals []string) *TelemetryDruidGroupByRequest {
	this := TelemetryDruidGroupByRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	this.Dimensions = dimensions
	this.Granularity = granularity
	this.Intervals = intervals
	return &this
}

// NewTelemetryDruidGroupByRequestWithDefaults instantiates a new TelemetryDruidGroupByRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidGroupByRequestWithDefaults() *TelemetryDruidGroupByRequest {
	this := TelemetryDruidGroupByRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidGroupByRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidGroupByRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidGroupByRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidGroupByRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetDimensions returns the Dimensions field value
func (o *TelemetryDruidGroupByRequest) GetDimensions() []TelemetryDruidDimensionSpec {
	if o == nil {
		var ret []TelemetryDruidDimensionSpec
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetDimensionsOk() ([]TelemetryDruidDimensionSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// SetDimensions sets field value
func (o *TelemetryDruidGroupByRequest) SetDimensions(v []TelemetryDruidDimensionSpec) {
	o.Dimensions = v
}

// GetLimitSpec returns the LimitSpec field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetLimitSpec() TelemetryDruidDefaultLimitSpec {
	if o == nil || o.LimitSpec == nil {
		var ret TelemetryDruidDefaultLimitSpec
		return ret
	}
	return *o.LimitSpec
}

// GetLimitSpecOk returns a tuple with the LimitSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetLimitSpecOk() (*TelemetryDruidDefaultLimitSpec, bool) {
	if o == nil || o.LimitSpec == nil {
		return nil, false
	}
	return o.LimitSpec, true
}

// HasLimitSpec returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasLimitSpec() bool {
	if o != nil && o.LimitSpec != nil {
		return true
	}

	return false
}

// SetLimitSpec gets a reference to the given TelemetryDruidDefaultLimitSpec and assigns it to the LimitSpec field.
func (o *TelemetryDruidGroupByRequest) SetLimitSpec(v TelemetryDruidDefaultLimitSpec) {
	o.LimitSpec = &v
}

// GetHaving returns the Having field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetHaving() TelemetryDruidHavingFilter {
	if o == nil || o.Having == nil {
		var ret TelemetryDruidHavingFilter
		return ret
	}
	return *o.Having
}

// GetHavingOk returns a tuple with the Having field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetHavingOk() (*TelemetryDruidHavingFilter, bool) {
	if o == nil || o.Having == nil {
		return nil, false
	}
	return o.Having, true
}

// HasHaving returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasHaving() bool {
	if o != nil && o.Having != nil {
		return true
	}

	return false
}

// SetHaving gets a reference to the given TelemetryDruidHavingFilter and assigns it to the Having field.
func (o *TelemetryDruidGroupByRequest) SetHaving(v TelemetryDruidHavingFilter) {
	o.Having = &v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidGroupByRequest) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidGroupByRequest) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidGroupByRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetAggregations() []TelemetryDruidAggregator {
	if o == nil || o.Aggregations == nil {
		var ret []TelemetryDruidAggregator
		return ret
	}
	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetAggregationsOk() ([]TelemetryDruidAggregator, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidGroupByRequest) SetAggregations(v []TelemetryDruidAggregator) {
	o.Aggregations = v
}

// GetPostAggregations returns the PostAggregations field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetPostAggregations() []TelemetryDruidPostAggregator {
	if o == nil || o.PostAggregations == nil {
		var ret []TelemetryDruidPostAggregator
		return ret
	}
	return o.PostAggregations
}

// GetPostAggregationsOk returns a tuple with the PostAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetPostAggregationsOk() ([]TelemetryDruidPostAggregator, bool) {
	if o == nil || o.PostAggregations == nil {
		return nil, false
	}
	return o.PostAggregations, true
}

// HasPostAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasPostAggregations() bool {
	if o != nil && o.PostAggregations != nil {
		return true
	}

	return false
}

// SetPostAggregations gets a reference to the given []TelemetryDruidPostAggregator and assigns it to the PostAggregations field.
func (o *TelemetryDruidGroupByRequest) SetPostAggregations(v []TelemetryDruidPostAggregator) {
	o.PostAggregations = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidGroupByRequest) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetIntervalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidGroupByRequest) SetIntervals(v []string) {
	o.Intervals = v
}

// GetSubtotalsSpec returns the SubtotalsSpec field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetSubtotalsSpec() map[string]interface{} {
	if o == nil || o.SubtotalsSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SubtotalsSpec
}

// GetSubtotalsSpecOk returns a tuple with the SubtotalsSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetSubtotalsSpecOk() (map[string]interface{}, bool) {
	if o == nil || o.SubtotalsSpec == nil {
		return nil, false
	}
	return o.SubtotalsSpec, true
}

// HasSubtotalsSpec returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasSubtotalsSpec() bool {
	if o != nil && o.SubtotalsSpec != nil {
		return true
	}

	return false
}

// SetSubtotalsSpec gets a reference to the given map[string]interface{} and assigns it to the SubtotalsSpec field.
func (o *TelemetryDruidGroupByRequest) SetSubtotalsSpec(v map[string]interface{}) {
	o.SubtotalsSpec = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidGroupByRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidGroupByRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryType"] = o.QueryType
	}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if true {
		toSerialize["dimensions"] = o.Dimensions
	}
	if o.LimitSpec != nil {
		toSerialize["limitSpec"] = o.LimitSpec
	}
	if o.Having != nil {
		toSerialize["having"] = o.Having
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.PostAggregations != nil {
		toSerialize["postAggregations"] = o.PostAggregations
	}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if o.SubtotalsSpec != nil {
		toSerialize["subtotalsSpec"] = o.SubtotalsSpec
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidGroupByRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidGroupByRequest := _TelemetryDruidGroupByRequest{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidGroupByRequest); err == nil {
		*o = TelemetryDruidGroupByRequest(varTelemetryDruidGroupByRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "dimensions")
		delete(additionalProperties, "limitSpec")
		delete(additionalProperties, "having")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "postAggregations")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "subtotalsSpec")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidGroupByRequest struct {
	value *TelemetryDruidGroupByRequest
	isSet bool
}

func (v NullableTelemetryDruidGroupByRequest) Get() *TelemetryDruidGroupByRequest {
	return v.value
}

func (v *NullableTelemetryDruidGroupByRequest) Set(val *TelemetryDruidGroupByRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidGroupByRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidGroupByRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidGroupByRequest(val *TelemetryDruidGroupByRequest) *NullableTelemetryDruidGroupByRequest {
	return &NullableTelemetryDruidGroupByRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidGroupByRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidGroupByRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
