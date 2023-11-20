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

// ApplianceStatusCheckAllOf Definition of the list of properties defined in 'appliance.StatusCheck', excluding properties defined in parent classes.
type ApplianceStatusCheckAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identifier of the status check.
	Code *string `json:"Code,omitempty"`
	// Result of this status check. * `OK` - Result of the check is OK. * `Warning` - Result of the check is Warning. * `Critical` - Result of the check is Critical. * `Info` - Result of the check is low.
	Result               *string `json:"Result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceStatusCheckAllOf ApplianceStatusCheckAllOf

// NewApplianceStatusCheckAllOf instantiates a new ApplianceStatusCheckAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceStatusCheckAllOf(classId string, objectType string) *ApplianceStatusCheckAllOf {
	this := ApplianceStatusCheckAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var result string = "OK"
	this.Result = &result
	return &this
}

// NewApplianceStatusCheckAllOfWithDefaults instantiates a new ApplianceStatusCheckAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceStatusCheckAllOfWithDefaults() *ApplianceStatusCheckAllOf {
	this := ApplianceStatusCheckAllOf{}
	var classId string = "appliance.StatusCheck"
	this.ClassId = classId
	var objectType string = "appliance.StatusCheck"
	this.ObjectType = objectType
	var result string = "OK"
	this.Result = &result
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceStatusCheckAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheckAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceStatusCheckAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceStatusCheckAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheckAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceStatusCheckAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApplianceStatusCheckAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheckAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApplianceStatusCheckAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApplianceStatusCheckAllOf) SetCode(v string) {
	o.Code = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ApplianceStatusCheckAllOf) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheckAllOf) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ApplianceStatusCheckAllOf) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ApplianceStatusCheckAllOf) SetResult(v string) {
	o.Result = &v
}

func (o ApplianceStatusCheckAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceStatusCheckAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceStatusCheckAllOf := _ApplianceStatusCheckAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceStatusCheckAllOf); err == nil {
		*o = ApplianceStatusCheckAllOf(varApplianceStatusCheckAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Code")
		delete(additionalProperties, "Result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceStatusCheckAllOf struct {
	value *ApplianceStatusCheckAllOf
	isSet bool
}

func (v NullableApplianceStatusCheckAllOf) Get() *ApplianceStatusCheckAllOf {
	return v.value
}

func (v *NullableApplianceStatusCheckAllOf) Set(val *ApplianceStatusCheckAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceStatusCheckAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceStatusCheckAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceStatusCheckAllOf(val *ApplianceStatusCheckAllOf) *NullableApplianceStatusCheckAllOf {
	return &NullableApplianceStatusCheckAllOf{value: val, isSet: true}
}

func (v NullableApplianceStatusCheckAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceStatusCheckAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
