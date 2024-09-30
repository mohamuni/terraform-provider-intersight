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

// checks if the TelemetryDruidExtractionFunctionTimeFormat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidExtractionFunctionTimeFormat{}

// TelemetryDruidExtractionFunctionTimeFormat Returns the dimension value formatted according to the given format string, time zone, and locale. For __time dimension values, this formats the time value bucketed by the aggregation granularity For a regular dimension, it assumes the string is formatted in ISO-8601 date and time format.
type TelemetryDruidExtractionFunctionTimeFormat struct {
	Type string `json:"type"`
	// Date time format for the resulting dimension value, in Joda Time DateTimeFormat, or null to use the default ISO8601 format.
	Format *string `json:"format,omitempty"`
	// Locale (language and country) to use, given as a IETF BCP 47 language tag, e.g. en-US, en-GB, fr-FR, fr-CA, etc.
	Locale *string `json:"locale,omitempty"`
	// Time zone to use in IANA tz database format, e.g. Europe/Berlin (this can possibly be different than the aggregation time-zone)
	TimeZone    *string                    `json:"timeZone,omitempty"`
	Granularity *TelemetryDruidGranularity `json:"granularity,omitempty"`
	// boolean value, set to true to treat input strings as millis rather than ISO8601 strings. Additionally, if format is null or not specified, output will be in millis rather than ISO8601.
	AsMillis             *bool `json:"asMillis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionTimeFormat TelemetryDruidExtractionFunctionTimeFormat

// NewTelemetryDruidExtractionFunctionTimeFormat instantiates a new TelemetryDruidExtractionFunctionTimeFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionTimeFormat(type_ string) *TelemetryDruidExtractionFunctionTimeFormat {
	this := TelemetryDruidExtractionFunctionTimeFormat{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidExtractionFunctionTimeFormatWithDefaults instantiates a new TelemetryDruidExtractionFunctionTimeFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionTimeFormatWithDefaults() *TelemetryDruidExtractionFunctionTimeFormat {
	this := TelemetryDruidExtractionFunctionTimeFormat{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionTimeFormat) SetType(v string) {
	o.Type = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *TelemetryDruidExtractionFunctionTimeFormat) SetFormat(v string) {
	o.Format = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *TelemetryDruidExtractionFunctionTimeFormat) SetLocale(v string) {
	o.Locale = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *TelemetryDruidExtractionFunctionTimeFormat) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetGranularity() TelemetryDruidGranularity {
	if o == nil || IsNil(o.Granularity) {
		var ret TelemetryDruidGranularity
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given TelemetryDruidGranularity and assigns it to the Granularity field.
func (o *TelemetryDruidExtractionFunctionTimeFormat) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = &v
}

// GetAsMillis returns the AsMillis field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetAsMillis() bool {
	if o == nil || IsNil(o.AsMillis) {
		var ret bool
		return ret
	}
	return *o.AsMillis
}

// GetAsMillisOk returns a tuple with the AsMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) GetAsMillisOk() (*bool, bool) {
	if o == nil || IsNil(o.AsMillis) {
		return nil, false
	}
	return o.AsMillis, true
}

// HasAsMillis returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeFormat) HasAsMillis() bool {
	if o != nil && !IsNil(o.AsMillis) {
		return true
	}

	return false
}

// SetAsMillis gets a reference to the given bool and assigns it to the AsMillis field.
func (o *TelemetryDruidExtractionFunctionTimeFormat) SetAsMillis(v bool) {
	o.AsMillis = &v
}

func (o TelemetryDruidExtractionFunctionTimeFormat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidExtractionFunctionTimeFormat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !IsNil(o.AsMillis) {
		toSerialize["asMillis"] = o.AsMillis
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidExtractionFunctionTimeFormat) UnmarshalJSON(data []byte) (err error) {
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
	varTelemetryDruidExtractionFunctionTimeFormat := _TelemetryDruidExtractionFunctionTimeFormat{}

	err = json.Unmarshal(data, &varTelemetryDruidExtractionFunctionTimeFormat)

	if err != nil {
		return err
	}

	*o = TelemetryDruidExtractionFunctionTimeFormat(varTelemetryDruidExtractionFunctionTimeFormat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "format")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "asMillis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionTimeFormat struct {
	value *TelemetryDruidExtractionFunctionTimeFormat
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionTimeFormat) Get() *TelemetryDruidExtractionFunctionTimeFormat {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionTimeFormat) Set(val *TelemetryDruidExtractionFunctionTimeFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionTimeFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionTimeFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionTimeFormat(val *TelemetryDruidExtractionFunctionTimeFormat) *NullableTelemetryDruidExtractionFunctionTimeFormat {
	return &NullableTelemetryDruidExtractionFunctionTimeFormat{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionTimeFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionTimeFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
