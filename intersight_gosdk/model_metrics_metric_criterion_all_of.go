/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// MetricsMetricCriterionAllOf Definition of the list of properties defined in 'metrics.MetricCriterion', excluding properties defined in parent classes.
type MetricsMetricCriterionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Function name which used to combine the group buckets into a single timeseries.
	Aggregation *string  `json:"Aggregation,omitempty"`
	Filters     []string `json:"Filters,omitempty"`
	Groups      []string `json:"Groups,omitempty"`
	// Instrument name used to collect measurements for the query.
	Instrument *string `json:"Instrument,omitempty"`
	// Indicates if this criterion should be used for the query.
	IsEnabled *bool `json:"IsEnabled,omitempty"`
	// Measurement name that is collected by the instrument for the query.
	Metric *string `json:"Metric,omitempty"`
	// Function name which used to combine the metrics into granularity buckets.
	MetricAggregation *string `json:"MetricAggregation,omitempty"`
	// The maximum number of result rows.
	TopLimit *int64 `json:"TopLimit,omitempty"`
	// Method on how to sort the result rows.
	TopSort              *string `json:"TopSort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetricsMetricCriterionAllOf MetricsMetricCriterionAllOf

// NewMetricsMetricCriterionAllOf instantiates a new MetricsMetricCriterionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMetricCriterionAllOf(classId string, objectType string) *MetricsMetricCriterionAllOf {
	this := MetricsMetricCriterionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetricsMetricCriterionAllOfWithDefaults instantiates a new MetricsMetricCriterionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMetricCriterionAllOfWithDefaults() *MetricsMetricCriterionAllOf {
	this := MetricsMetricCriterionAllOf{}
	var classId string = "metrics.MetricCriterion"
	this.ClassId = classId
	var objectType string = "metrics.MetricCriterion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetricsMetricCriterionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetricsMetricCriterionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetricsMetricCriterionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetricsMetricCriterionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *MetricsMetricCriterionAllOf) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsMetricCriterionAllOf) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsMetricCriterionAllOf) GetFiltersOk() ([]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *MetricsMetricCriterionAllOf) SetFilters(v []string) {
	o.Filters = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsMetricCriterionAllOf) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsMetricCriterionAllOf) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *MetricsMetricCriterionAllOf) SetGroups(v []string) {
	o.Groups = v
}

// GetInstrument returns the Instrument field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetInstrument() string {
	if o == nil || o.Instrument == nil {
		var ret string
		return ret
	}
	return *o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetInstrumentOk() (*string, bool) {
	if o == nil || o.Instrument == nil {
		return nil, false
	}
	return o.Instrument, true
}

// HasInstrument returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasInstrument() bool {
	if o != nil && o.Instrument != nil {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given string and assigns it to the Instrument field.
func (o *MetricsMetricCriterionAllOf) SetInstrument(v string) {
	o.Instrument = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *MetricsMetricCriterionAllOf) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *MetricsMetricCriterionAllOf) SetMetric(v string) {
	o.Metric = &v
}

// GetMetricAggregation returns the MetricAggregation field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetMetricAggregation() string {
	if o == nil || o.MetricAggregation == nil {
		var ret string
		return ret
	}
	return *o.MetricAggregation
}

// GetMetricAggregationOk returns a tuple with the MetricAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetMetricAggregationOk() (*string, bool) {
	if o == nil || o.MetricAggregation == nil {
		return nil, false
	}
	return o.MetricAggregation, true
}

// HasMetricAggregation returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasMetricAggregation() bool {
	if o != nil && o.MetricAggregation != nil {
		return true
	}

	return false
}

// SetMetricAggregation gets a reference to the given string and assigns it to the MetricAggregation field.
func (o *MetricsMetricCriterionAllOf) SetMetricAggregation(v string) {
	o.MetricAggregation = &v
}

// GetTopLimit returns the TopLimit field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetTopLimit() int64 {
	if o == nil || o.TopLimit == nil {
		var ret int64
		return ret
	}
	return *o.TopLimit
}

// GetTopLimitOk returns a tuple with the TopLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetTopLimitOk() (*int64, bool) {
	if o == nil || o.TopLimit == nil {
		return nil, false
	}
	return o.TopLimit, true
}

// HasTopLimit returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasTopLimit() bool {
	if o != nil && o.TopLimit != nil {
		return true
	}

	return false
}

// SetTopLimit gets a reference to the given int64 and assigns it to the TopLimit field.
func (o *MetricsMetricCriterionAllOf) SetTopLimit(v int64) {
	o.TopLimit = &v
}

// GetTopSort returns the TopSort field value if set, zero value otherwise.
func (o *MetricsMetricCriterionAllOf) GetTopSort() string {
	if o == nil || o.TopSort == nil {
		var ret string
		return ret
	}
	return *o.TopSort
}

// GetTopSortOk returns a tuple with the TopSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterionAllOf) GetTopSortOk() (*string, bool) {
	if o == nil || o.TopSort == nil {
		return nil, false
	}
	return o.TopSort, true
}

// HasTopSort returns a boolean if a field has been set.
func (o *MetricsMetricCriterionAllOf) HasTopSort() bool {
	if o != nil && o.TopSort != nil {
		return true
	}

	return false
}

// SetTopSort gets a reference to the given string and assigns it to the TopSort field.
func (o *MetricsMetricCriterionAllOf) SetTopSort(v string) {
	o.TopSort = &v
}

func (o MetricsMetricCriterionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Aggregation != nil {
		toSerialize["Aggregation"] = o.Aggregation
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.Groups != nil {
		toSerialize["Groups"] = o.Groups
	}
	if o.Instrument != nil {
		toSerialize["Instrument"] = o.Instrument
	}
	if o.IsEnabled != nil {
		toSerialize["IsEnabled"] = o.IsEnabled
	}
	if o.Metric != nil {
		toSerialize["Metric"] = o.Metric
	}
	if o.MetricAggregation != nil {
		toSerialize["MetricAggregation"] = o.MetricAggregation
	}
	if o.TopLimit != nil {
		toSerialize["TopLimit"] = o.TopLimit
	}
	if o.TopSort != nil {
		toSerialize["TopSort"] = o.TopSort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetricsMetricCriterionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMetricsMetricCriterionAllOf := _MetricsMetricCriterionAllOf{}

	if err = json.Unmarshal(bytes, &varMetricsMetricCriterionAllOf); err == nil {
		*o = MetricsMetricCriterionAllOf(varMetricsMetricCriterionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Aggregation")
		delete(additionalProperties, "Filters")
		delete(additionalProperties, "Groups")
		delete(additionalProperties, "Instrument")
		delete(additionalProperties, "IsEnabled")
		delete(additionalProperties, "Metric")
		delete(additionalProperties, "MetricAggregation")
		delete(additionalProperties, "TopLimit")
		delete(additionalProperties, "TopSort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricsMetricCriterionAllOf struct {
	value *MetricsMetricCriterionAllOf
	isSet bool
}

func (v NullableMetricsMetricCriterionAllOf) Get() *MetricsMetricCriterionAllOf {
	return v.value
}

func (v *NullableMetricsMetricCriterionAllOf) Set(val *MetricsMetricCriterionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMetricCriterionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMetricCriterionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMetricCriterionAllOf(val *MetricsMetricCriterionAllOf) *NullableMetricsMetricCriterionAllOf {
	return &NullableMetricsMetricCriterionAllOf{value: val, isSet: true}
}

func (v NullableMetricsMetricCriterionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMetricCriterionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}