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
	"reflect"
	"strings"
)

// checks if the MetricsMetricCriterion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsMetricCriterion{}

// MetricsMetricCriterion Contains all the data and parameters used to configure a metric.
type MetricsMetricCriterion struct {
	MoBaseComplexType
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

type _MetricsMetricCriterion MetricsMetricCriterion

// NewMetricsMetricCriterion instantiates a new MetricsMetricCriterion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMetricCriterion(classId string, objectType string) *MetricsMetricCriterion {
	this := MetricsMetricCriterion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetricsMetricCriterionWithDefaults instantiates a new MetricsMetricCriterion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMetricCriterionWithDefaults() *MetricsMetricCriterion {
	this := MetricsMetricCriterion{}
	var classId string = "metrics.MetricCriterion"
	this.ClassId = classId
	var objectType string = "metrics.MetricCriterion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetricsMetricCriterion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetricsMetricCriterion) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "metrics.MetricCriterion" of the ClassId field.
func (o *MetricsMetricCriterion) GetDefaultClassId() interface{} {
	return "metrics.MetricCriterion"
}

// GetObjectType returns the ObjectType field value
func (o *MetricsMetricCriterion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetricsMetricCriterion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "metrics.MetricCriterion" of the ObjectType field.
func (o *MetricsMetricCriterion) GetDefaultObjectType() interface{} {
	return "metrics.MetricCriterion"
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetAggregation() string {
	if o == nil || IsNil(o.Aggregation) {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetAggregationOk() (*string, bool) {
	if o == nil || IsNil(o.Aggregation) {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasAggregation() bool {
	if o != nil && !IsNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *MetricsMetricCriterion) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsMetricCriterion) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsMetricCriterion) GetFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *MetricsMetricCriterion) SetFilters(v []string) {
	o.Filters = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsMetricCriterion) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsMetricCriterion) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *MetricsMetricCriterion) SetGroups(v []string) {
	o.Groups = v
}

// GetInstrument returns the Instrument field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetInstrument() string {
	if o == nil || IsNil(o.Instrument) {
		var ret string
		return ret
	}
	return *o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetInstrumentOk() (*string, bool) {
	if o == nil || IsNil(o.Instrument) {
		return nil, false
	}
	return o.Instrument, true
}

// HasInstrument returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasInstrument() bool {
	if o != nil && !IsNil(o.Instrument) {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given string and assigns it to the Instrument field.
func (o *MetricsMetricCriterion) SetInstrument(v string) {
	o.Instrument = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *MetricsMetricCriterion) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetMetric() string {
	if o == nil || IsNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetMetricOk() (*string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *MetricsMetricCriterion) SetMetric(v string) {
	o.Metric = &v
}

// GetMetricAggregation returns the MetricAggregation field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetMetricAggregation() string {
	if o == nil || IsNil(o.MetricAggregation) {
		var ret string
		return ret
	}
	return *o.MetricAggregation
}

// GetMetricAggregationOk returns a tuple with the MetricAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetMetricAggregationOk() (*string, bool) {
	if o == nil || IsNil(o.MetricAggregation) {
		return nil, false
	}
	return o.MetricAggregation, true
}

// HasMetricAggregation returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasMetricAggregation() bool {
	if o != nil && !IsNil(o.MetricAggregation) {
		return true
	}

	return false
}

// SetMetricAggregation gets a reference to the given string and assigns it to the MetricAggregation field.
func (o *MetricsMetricCriterion) SetMetricAggregation(v string) {
	o.MetricAggregation = &v
}

// GetTopLimit returns the TopLimit field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetTopLimit() int64 {
	if o == nil || IsNil(o.TopLimit) {
		var ret int64
		return ret
	}
	return *o.TopLimit
}

// GetTopLimitOk returns a tuple with the TopLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetTopLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.TopLimit) {
		return nil, false
	}
	return o.TopLimit, true
}

// HasTopLimit returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasTopLimit() bool {
	if o != nil && !IsNil(o.TopLimit) {
		return true
	}

	return false
}

// SetTopLimit gets a reference to the given int64 and assigns it to the TopLimit field.
func (o *MetricsMetricCriterion) SetTopLimit(v int64) {
	o.TopLimit = &v
}

// GetTopSort returns the TopSort field value if set, zero value otherwise.
func (o *MetricsMetricCriterion) GetTopSort() string {
	if o == nil || IsNil(o.TopSort) {
		var ret string
		return ret
	}
	return *o.TopSort
}

// GetTopSortOk returns a tuple with the TopSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetricCriterion) GetTopSortOk() (*string, bool) {
	if o == nil || IsNil(o.TopSort) {
		return nil, false
	}
	return o.TopSort, true
}

// HasTopSort returns a boolean if a field has been set.
func (o *MetricsMetricCriterion) HasTopSort() bool {
	if o != nil && !IsNil(o.TopSort) {
		return true
	}

	return false
}

// SetTopSort gets a reference to the given string and assigns it to the TopSort field.
func (o *MetricsMetricCriterion) SetTopSort(v string) {
	o.TopSort = &v
}

func (o MetricsMetricCriterion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsMetricCriterion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Aggregation) {
		toSerialize["Aggregation"] = o.Aggregation
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.Groups != nil {
		toSerialize["Groups"] = o.Groups
	}
	if !IsNil(o.Instrument) {
		toSerialize["Instrument"] = o.Instrument
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["IsEnabled"] = o.IsEnabled
	}
	if !IsNil(o.Metric) {
		toSerialize["Metric"] = o.Metric
	}
	if !IsNil(o.MetricAggregation) {
		toSerialize["MetricAggregation"] = o.MetricAggregation
	}
	if !IsNil(o.TopLimit) {
		toSerialize["TopLimit"] = o.TopLimit
	}
	if !IsNil(o.TopSort) {
		toSerialize["TopSort"] = o.TopSort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricsMetricCriterion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
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
	type MetricsMetricCriterionWithoutEmbeddedStruct struct {
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
		TopSort *string `json:"TopSort,omitempty"`
	}

	varMetricsMetricCriterionWithoutEmbeddedStruct := MetricsMetricCriterionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMetricsMetricCriterionWithoutEmbeddedStruct)
	if err == nil {
		varMetricsMetricCriterion := _MetricsMetricCriterion{}
		varMetricsMetricCriterion.ClassId = varMetricsMetricCriterionWithoutEmbeddedStruct.ClassId
		varMetricsMetricCriterion.ObjectType = varMetricsMetricCriterionWithoutEmbeddedStruct.ObjectType
		varMetricsMetricCriterion.Aggregation = varMetricsMetricCriterionWithoutEmbeddedStruct.Aggregation
		varMetricsMetricCriterion.Filters = varMetricsMetricCriterionWithoutEmbeddedStruct.Filters
		varMetricsMetricCriterion.Groups = varMetricsMetricCriterionWithoutEmbeddedStruct.Groups
		varMetricsMetricCriterion.Instrument = varMetricsMetricCriterionWithoutEmbeddedStruct.Instrument
		varMetricsMetricCriterion.IsEnabled = varMetricsMetricCriterionWithoutEmbeddedStruct.IsEnabled
		varMetricsMetricCriterion.Metric = varMetricsMetricCriterionWithoutEmbeddedStruct.Metric
		varMetricsMetricCriterion.MetricAggregation = varMetricsMetricCriterionWithoutEmbeddedStruct.MetricAggregation
		varMetricsMetricCriterion.TopLimit = varMetricsMetricCriterionWithoutEmbeddedStruct.TopLimit
		varMetricsMetricCriterion.TopSort = varMetricsMetricCriterionWithoutEmbeddedStruct.TopSort
		*o = MetricsMetricCriterion(varMetricsMetricCriterion)
	} else {
		return err
	}

	varMetricsMetricCriterion := _MetricsMetricCriterion{}

	err = json.Unmarshal(data, &varMetricsMetricCriterion)
	if err == nil {
		o.MoBaseComplexType = varMetricsMetricCriterion.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricsMetricCriterion struct {
	value *MetricsMetricCriterion
	isSet bool
}

func (v NullableMetricsMetricCriterion) Get() *MetricsMetricCriterion {
	return v.value
}

func (v *NullableMetricsMetricCriterion) Set(val *MetricsMetricCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMetricCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMetricCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMetricCriterion(val *MetricsMetricCriterion) *NullableMetricsMetricCriterion {
	return &NullableMetricsMetricCriterion{value: val, isSet: true}
}

func (v NullableMetricsMetricCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMetricCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
