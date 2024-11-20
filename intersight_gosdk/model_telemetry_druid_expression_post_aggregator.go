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

// checks if the TelemetryDruidExpressionPostAggregator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidExpressionPostAggregator{}

// TelemetryDruidExpressionPostAggregator A post-aggregator defined as a Druid expression. Supported Druid expressions are documented at https://druid.apache.org/docs/latest/misc/math-expr.html.
type TelemetryDruidExpressionPostAggregator struct {
	// The post-aggregator type.
	Type string `json:"type"`
	// Fields processed by post aggregator
	Fields []TelemetryDruidPostAggregator `json:"fields,omitempty"`
	// Output name for the return value of the expression.
	Name *string `json:"name,omitempty"`
	// The Druid expression.
	Expression *string `json:"expression,omitempty"`
	// Expression post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last.
	Ordering             *string `json:"ordering,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExpressionPostAggregator TelemetryDruidExpressionPostAggregator

// NewTelemetryDruidExpressionPostAggregator instantiates a new TelemetryDruidExpressionPostAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExpressionPostAggregator(type_ string) *TelemetryDruidExpressionPostAggregator {
	this := TelemetryDruidExpressionPostAggregator{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidExpressionPostAggregatorWithDefaults instantiates a new TelemetryDruidExpressionPostAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExpressionPostAggregatorWithDefaults() *TelemetryDruidExpressionPostAggregator {
	this := TelemetryDruidExpressionPostAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExpressionPostAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExpressionPostAggregator) SetType(v string) {
	o.Type = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregator) GetFields() []TelemetryDruidPostAggregator {
	if o == nil || IsNil(o.Fields) {
		var ret []TelemetryDruidPostAggregator
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregator) GetFieldsOk() ([]TelemetryDruidPostAggregator, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregator) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []TelemetryDruidPostAggregator and assigns it to the Fields field.
func (o *TelemetryDruidExpressionPostAggregator) SetFields(v []TelemetryDruidPostAggregator) {
	o.Fields = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidExpressionPostAggregator) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregator) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregator) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregator) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *TelemetryDruidExpressionPostAggregator) SetExpression(v string) {
	o.Expression = &v
}

// GetOrdering returns the Ordering field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregator) GetOrdering() string {
	if o == nil || IsNil(o.Ordering) {
		var ret string
		return ret
	}
	return *o.Ordering
}

// GetOrderingOk returns a tuple with the Ordering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregator) GetOrderingOk() (*string, bool) {
	if o == nil || IsNil(o.Ordering) {
		return nil, false
	}
	return o.Ordering, true
}

// HasOrdering returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregator) HasOrdering() bool {
	if o != nil && !IsNil(o.Ordering) {
		return true
	}

	return false
}

// SetOrdering gets a reference to the given string and assigns it to the Ordering field.
func (o *TelemetryDruidExpressionPostAggregator) SetOrdering(v string) {
	o.Ordering = &v
}

func (o TelemetryDruidExpressionPostAggregator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidExpressionPostAggregator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Ordering) {
		toSerialize["ordering"] = o.Ordering
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidExpressionPostAggregator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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
	varTelemetryDruidExpressionPostAggregator := _TelemetryDruidExpressionPostAggregator{}

	err = json.Unmarshal(data, &varTelemetryDruidExpressionPostAggregator)

	if err != nil {
		return err
	}

	*o = TelemetryDruidExpressionPostAggregator(varTelemetryDruidExpressionPostAggregator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "name")
		delete(additionalProperties, "expression")
		delete(additionalProperties, "ordering")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExpressionPostAggregator struct {
	value *TelemetryDruidExpressionPostAggregator
	isSet bool
}

func (v NullableTelemetryDruidExpressionPostAggregator) Get() *TelemetryDruidExpressionPostAggregator {
	return v.value
}

func (v *NullableTelemetryDruidExpressionPostAggregator) Set(val *TelemetryDruidExpressionPostAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExpressionPostAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExpressionPostAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExpressionPostAggregator(val *TelemetryDruidExpressionPostAggregator) *NullableTelemetryDruidExpressionPostAggregator {
	return &NullableTelemetryDruidExpressionPostAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidExpressionPostAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExpressionPostAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
