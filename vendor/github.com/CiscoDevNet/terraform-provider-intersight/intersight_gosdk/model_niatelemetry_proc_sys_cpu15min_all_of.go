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

// NiatelemetryProcSysCpu15minAllOf Definition of the list of properties defined in 'niatelemetry.ProcSysCpu15min', excluding properties defined in parent classes.
type NiatelemetryProcSysCpu15minAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the kernal average value.
	KernalAvg *string `json:"KernalAvg,omitempty"`
	// Returns the user average value.
	UserAvg              *string `json:"UserAvg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryProcSysCpu15minAllOf NiatelemetryProcSysCpu15minAllOf

// NewNiatelemetryProcSysCpu15minAllOf instantiates a new NiatelemetryProcSysCpu15minAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryProcSysCpu15minAllOf(classId string, objectType string) *NiatelemetryProcSysCpu15minAllOf {
	this := NiatelemetryProcSysCpu15minAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryProcSysCpu15minAllOfWithDefaults instantiates a new NiatelemetryProcSysCpu15minAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryProcSysCpu15minAllOfWithDefaults() *NiatelemetryProcSysCpu15minAllOf {
	this := NiatelemetryProcSysCpu15minAllOf{}
	var classId string = "niatelemetry.ProcSysCpu15min"
	this.ClassId = classId
	var objectType string = "niatelemetry.ProcSysCpu15min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryProcSysCpu15minAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu15minAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryProcSysCpu15minAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryProcSysCpu15minAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu15minAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryProcSysCpu15minAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKernalAvg returns the KernalAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysCpu15minAllOf) GetKernalAvg() string {
	if o == nil || o.KernalAvg == nil {
		var ret string
		return ret
	}
	return *o.KernalAvg
}

// GetKernalAvgOk returns a tuple with the KernalAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu15minAllOf) GetKernalAvgOk() (*string, bool) {
	if o == nil || o.KernalAvg == nil {
		return nil, false
	}
	return o.KernalAvg, true
}

// HasKernalAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysCpu15minAllOf) HasKernalAvg() bool {
	if o != nil && o.KernalAvg != nil {
		return true
	}

	return false
}

// SetKernalAvg gets a reference to the given string and assigns it to the KernalAvg field.
func (o *NiatelemetryProcSysCpu15minAllOf) SetKernalAvg(v string) {
	o.KernalAvg = &v
}

// GetUserAvg returns the UserAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysCpu15minAllOf) GetUserAvg() string {
	if o == nil || o.UserAvg == nil {
		var ret string
		return ret
	}
	return *o.UserAvg
}

// GetUserAvgOk returns a tuple with the UserAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu15minAllOf) GetUserAvgOk() (*string, bool) {
	if o == nil || o.UserAvg == nil {
		return nil, false
	}
	return o.UserAvg, true
}

// HasUserAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysCpu15minAllOf) HasUserAvg() bool {
	if o != nil && o.UserAvg != nil {
		return true
	}

	return false
}

// SetUserAvg gets a reference to the given string and assigns it to the UserAvg field.
func (o *NiatelemetryProcSysCpu15minAllOf) SetUserAvg(v string) {
	o.UserAvg = &v
}

func (o NiatelemetryProcSysCpu15minAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KernalAvg != nil {
		toSerialize["KernalAvg"] = o.KernalAvg
	}
	if o.UserAvg != nil {
		toSerialize["UserAvg"] = o.UserAvg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryProcSysCpu15minAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryProcSysCpu15minAllOf := _NiatelemetryProcSysCpu15minAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryProcSysCpu15minAllOf); err == nil {
		*o = NiatelemetryProcSysCpu15minAllOf(varNiatelemetryProcSysCpu15minAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KernalAvg")
		delete(additionalProperties, "UserAvg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryProcSysCpu15minAllOf struct {
	value *NiatelemetryProcSysCpu15minAllOf
	isSet bool
}

func (v NullableNiatelemetryProcSysCpu15minAllOf) Get() *NiatelemetryProcSysCpu15minAllOf {
	return v.value
}

func (v *NullableNiatelemetryProcSysCpu15minAllOf) Set(val *NiatelemetryProcSysCpu15minAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryProcSysCpu15minAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryProcSysCpu15minAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryProcSysCpu15minAllOf(val *NiatelemetryProcSysCpu15minAllOf) *NullableNiatelemetryProcSysCpu15minAllOf {
	return &NullableNiatelemetryProcSysCpu15minAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryProcSysCpu15minAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryProcSysCpu15minAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
