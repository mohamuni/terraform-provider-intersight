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

// checks if the TelemetryDruidExtractionFunctionPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidExtractionFunctionPartial{}

// TelemetryDruidExtractionFunctionPartial Returns the dimension value unchanged if the regular expression matches, otherwise returns null.
type TelemetryDruidExtractionFunctionPartial struct {
	Type                 string `json:"type"`
	Expr                 string `json:"expr"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionPartial TelemetryDruidExtractionFunctionPartial

// NewTelemetryDruidExtractionFunctionPartial instantiates a new TelemetryDruidExtractionFunctionPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionPartial(type_ string, expr string) *TelemetryDruidExtractionFunctionPartial {
	this := TelemetryDruidExtractionFunctionPartial{}
	this.Type = type_
	this.Expr = expr
	return &this
}

// NewTelemetryDruidExtractionFunctionPartialWithDefaults instantiates a new TelemetryDruidExtractionFunctionPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionPartialWithDefaults() *TelemetryDruidExtractionFunctionPartial {
	this := TelemetryDruidExtractionFunctionPartial{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionPartial) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionPartial) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionPartial) SetType(v string) {
	o.Type = v
}

// GetExpr returns the Expr field value
func (o *TelemetryDruidExtractionFunctionPartial) GetExpr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expr
}

// GetExprOk returns a tuple with the Expr field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionPartial) GetExprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expr, true
}

// SetExpr sets field value
func (o *TelemetryDruidExtractionFunctionPartial) SetExpr(v string) {
	o.Expr = v
}

func (o TelemetryDruidExtractionFunctionPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidExtractionFunctionPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["expr"] = o.Expr

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidExtractionFunctionPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"expr",
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
	varTelemetryDruidExtractionFunctionPartial := _TelemetryDruidExtractionFunctionPartial{}

	err = json.Unmarshal(data, &varTelemetryDruidExtractionFunctionPartial)

	if err != nil {
		return err
	}

	*o = TelemetryDruidExtractionFunctionPartial(varTelemetryDruidExtractionFunctionPartial)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "expr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionPartial struct {
	value *TelemetryDruidExtractionFunctionPartial
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionPartial) Get() *TelemetryDruidExtractionFunctionPartial {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionPartial) Set(val *TelemetryDruidExtractionFunctionPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionPartial(val *TelemetryDruidExtractionFunctionPartial) *NullableTelemetryDruidExtractionFunctionPartial {
	return &NullableTelemetryDruidExtractionFunctionPartial{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
