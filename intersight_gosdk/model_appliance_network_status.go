/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the ApplianceNetworkStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceNetworkStatus{}

// ApplianceNetworkStatus Time to ping a target endpoint.
type ApplianceNetworkStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Hostname of the destination endpoint.
	DestinationHostname *string `json:"DestinationHostname,omitempty"`
	// Time to reach the destination endpoint in milliseconds from the source endpoint.
	PingTime *float32 `json:"PingTime,omitempty"`
	// Hostname of the source endpoint.
	SourceHostname       *string `json:"SourceHostname,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNetworkStatus ApplianceNetworkStatus

// NewApplianceNetworkStatus instantiates a new ApplianceNetworkStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNetworkStatus(classId string, objectType string) *ApplianceNetworkStatus {
	this := ApplianceNetworkStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNetworkStatusWithDefaults instantiates a new ApplianceNetworkStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNetworkStatusWithDefaults() *ApplianceNetworkStatus {
	this := ApplianceNetworkStatus{}
	var classId string = "appliance.NetworkStatus"
	this.ClassId = classId
	var objectType string = "appliance.NetworkStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNetworkStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNetworkStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.NetworkStatus" of the ClassId field.
func (o *ApplianceNetworkStatus) GetDefaultClassId() interface{} {
	return "appliance.NetworkStatus"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNetworkStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNetworkStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.NetworkStatus" of the ObjectType field.
func (o *ApplianceNetworkStatus) GetDefaultObjectType() interface{} {
	return "appliance.NetworkStatus"
}

// GetDestinationHostname returns the DestinationHostname field value if set, zero value otherwise.
func (o *ApplianceNetworkStatus) GetDestinationHostname() string {
	if o == nil || IsNil(o.DestinationHostname) {
		var ret string
		return ret
	}
	return *o.DestinationHostname
}

// GetDestinationHostnameOk returns a tuple with the DestinationHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkStatus) GetDestinationHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationHostname) {
		return nil, false
	}
	return o.DestinationHostname, true
}

// HasDestinationHostname returns a boolean if a field has been set.
func (o *ApplianceNetworkStatus) HasDestinationHostname() bool {
	if o != nil && !IsNil(o.DestinationHostname) {
		return true
	}

	return false
}

// SetDestinationHostname gets a reference to the given string and assigns it to the DestinationHostname field.
func (o *ApplianceNetworkStatus) SetDestinationHostname(v string) {
	o.DestinationHostname = &v
}

// GetPingTime returns the PingTime field value if set, zero value otherwise.
func (o *ApplianceNetworkStatus) GetPingTime() float32 {
	if o == nil || IsNil(o.PingTime) {
		var ret float32
		return ret
	}
	return *o.PingTime
}

// GetPingTimeOk returns a tuple with the PingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkStatus) GetPingTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.PingTime) {
		return nil, false
	}
	return o.PingTime, true
}

// HasPingTime returns a boolean if a field has been set.
func (o *ApplianceNetworkStatus) HasPingTime() bool {
	if o != nil && !IsNil(o.PingTime) {
		return true
	}

	return false
}

// SetPingTime gets a reference to the given float32 and assigns it to the PingTime field.
func (o *ApplianceNetworkStatus) SetPingTime(v float32) {
	o.PingTime = &v
}

// GetSourceHostname returns the SourceHostname field value if set, zero value otherwise.
func (o *ApplianceNetworkStatus) GetSourceHostname() string {
	if o == nil || IsNil(o.SourceHostname) {
		var ret string
		return ret
	}
	return *o.SourceHostname
}

// GetSourceHostnameOk returns a tuple with the SourceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkStatus) GetSourceHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceHostname) {
		return nil, false
	}
	return o.SourceHostname, true
}

// HasSourceHostname returns a boolean if a field has been set.
func (o *ApplianceNetworkStatus) HasSourceHostname() bool {
	if o != nil && !IsNil(o.SourceHostname) {
		return true
	}

	return false
}

// SetSourceHostname gets a reference to the given string and assigns it to the SourceHostname field.
func (o *ApplianceNetworkStatus) SetSourceHostname(v string) {
	o.SourceHostname = &v
}

func (o ApplianceNetworkStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceNetworkStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DestinationHostname) {
		toSerialize["DestinationHostname"] = o.DestinationHostname
	}
	if !IsNil(o.PingTime) {
		toSerialize["PingTime"] = o.PingTime
	}
	if !IsNil(o.SourceHostname) {
		toSerialize["SourceHostname"] = o.SourceHostname
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceNetworkStatus) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceNetworkStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Hostname of the destination endpoint.
		DestinationHostname *string `json:"DestinationHostname,omitempty"`
		// Time to reach the destination endpoint in milliseconds from the source endpoint.
		PingTime *float32 `json:"PingTime,omitempty"`
		// Hostname of the source endpoint.
		SourceHostname *string `json:"SourceHostname,omitempty"`
	}

	varApplianceNetworkStatusWithoutEmbeddedStruct := ApplianceNetworkStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceNetworkStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceNetworkStatus := _ApplianceNetworkStatus{}
		varApplianceNetworkStatus.ClassId = varApplianceNetworkStatusWithoutEmbeddedStruct.ClassId
		varApplianceNetworkStatus.ObjectType = varApplianceNetworkStatusWithoutEmbeddedStruct.ObjectType
		varApplianceNetworkStatus.DestinationHostname = varApplianceNetworkStatusWithoutEmbeddedStruct.DestinationHostname
		varApplianceNetworkStatus.PingTime = varApplianceNetworkStatusWithoutEmbeddedStruct.PingTime
		varApplianceNetworkStatus.SourceHostname = varApplianceNetworkStatusWithoutEmbeddedStruct.SourceHostname
		*o = ApplianceNetworkStatus(varApplianceNetworkStatus)
	} else {
		return err
	}

	varApplianceNetworkStatus := _ApplianceNetworkStatus{}

	err = json.Unmarshal(data, &varApplianceNetworkStatus)
	if err == nil {
		o.MoBaseComplexType = varApplianceNetworkStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestinationHostname")
		delete(additionalProperties, "PingTime")
		delete(additionalProperties, "SourceHostname")

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

type NullableApplianceNetworkStatus struct {
	value *ApplianceNetworkStatus
	isSet bool
}

func (v NullableApplianceNetworkStatus) Get() *ApplianceNetworkStatus {
	return v.value
}

func (v *NullableApplianceNetworkStatus) Set(val *ApplianceNetworkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNetworkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNetworkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNetworkStatus(val *ApplianceNetworkStatus) *NullableApplianceNetworkStatus {
	return &NullableApplianceNetworkStatus{value: val, isSet: true}
}

func (v NullableApplianceNetworkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNetworkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
