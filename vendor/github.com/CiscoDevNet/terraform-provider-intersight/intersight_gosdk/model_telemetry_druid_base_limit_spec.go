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

// checks if the TelemetryDruidBaseLimitSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidBaseLimitSpec{}

// TelemetryDruidBaseLimitSpec The limitSpec field provides the functionality to sort and limit the set of results from a groupBy query. If you group by a single dimension and are ordering by a single metric, we highly recommend using TopN Queries instead. The performance will be substantially better.
type TelemetryDruidBaseLimitSpec struct {
	// The limit spec type.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidBaseLimitSpec TelemetryDruidBaseLimitSpec

// NewTelemetryDruidBaseLimitSpec instantiates a new TelemetryDruidBaseLimitSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidBaseLimitSpec(type_ string) *TelemetryDruidBaseLimitSpec {
	this := TelemetryDruidBaseLimitSpec{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidBaseLimitSpecWithDefaults instantiates a new TelemetryDruidBaseLimitSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidBaseLimitSpecWithDefaults() *TelemetryDruidBaseLimitSpec {
	this := TelemetryDruidBaseLimitSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidBaseLimitSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidBaseLimitSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidBaseLimitSpec) SetType(v string) {
	o.Type = v
}

func (o TelemetryDruidBaseLimitSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidBaseLimitSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidBaseLimitSpec) UnmarshalJSON(data []byte) (err error) {
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
	varTelemetryDruidBaseLimitSpec := _TelemetryDruidBaseLimitSpec{}

	err = json.Unmarshal(data, &varTelemetryDruidBaseLimitSpec)

	if err != nil {
		return err
	}

	*o = TelemetryDruidBaseLimitSpec(varTelemetryDruidBaseLimitSpec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidBaseLimitSpec struct {
	value *TelemetryDruidBaseLimitSpec
	isSet bool
}

func (v NullableTelemetryDruidBaseLimitSpec) Get() *TelemetryDruidBaseLimitSpec {
	return v.value
}

func (v *NullableTelemetryDruidBaseLimitSpec) Set(val *TelemetryDruidBaseLimitSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidBaseLimitSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidBaseLimitSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidBaseLimitSpec(val *TelemetryDruidBaseLimitSpec) *NullableTelemetryDruidBaseLimitSpec {
	return &NullableTelemetryDruidBaseLimitSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidBaseLimitSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidBaseLimitSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
