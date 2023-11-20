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

// HyperflexHxZoneInfoDtAllOf Definition of the list of properties defined in 'hyperflex.HxZoneInfoDt', excluding properties defined in parent classes.
type HyperflexHxZoneInfoDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of nodes in the Zone.
	NumNodes *int64 `json:"NumNodes,omitempty"`
	// Zone UUID for this HyperFlex node.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxZoneInfoDtAllOf HyperflexHxZoneInfoDtAllOf

// NewHyperflexHxZoneInfoDtAllOf instantiates a new HyperflexHxZoneInfoDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxZoneInfoDtAllOf(classId string, objectType string) *HyperflexHxZoneInfoDtAllOf {
	this := HyperflexHxZoneInfoDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxZoneInfoDtAllOfWithDefaults instantiates a new HyperflexHxZoneInfoDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxZoneInfoDtAllOfWithDefaults() *HyperflexHxZoneInfoDtAllOf {
	this := HyperflexHxZoneInfoDtAllOf{}
	var classId string = "hyperflex.HxZoneInfoDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxZoneInfoDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxZoneInfoDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneInfoDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxZoneInfoDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxZoneInfoDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneInfoDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxZoneInfoDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNumNodes returns the NumNodes field value if set, zero value otherwise.
func (o *HyperflexHxZoneInfoDtAllOf) GetNumNodes() int64 {
	if o == nil || o.NumNodes == nil {
		var ret int64
		return ret
	}
	return *o.NumNodes
}

// GetNumNodesOk returns a tuple with the NumNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneInfoDtAllOf) GetNumNodesOk() (*int64, bool) {
	if o == nil || o.NumNodes == nil {
		return nil, false
	}
	return o.NumNodes, true
}

// HasNumNodes returns a boolean if a field has been set.
func (o *HyperflexHxZoneInfoDtAllOf) HasNumNodes() bool {
	if o != nil && o.NumNodes != nil {
		return true
	}

	return false
}

// SetNumNodes gets a reference to the given int64 and assigns it to the NumNodes field.
func (o *HyperflexHxZoneInfoDtAllOf) SetNumNodes(v int64) {
	o.NumNodes = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHxZoneInfoDtAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneInfoDtAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHxZoneInfoDtAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHxZoneInfoDtAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o HyperflexHxZoneInfoDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NumNodes != nil {
		toSerialize["NumNodes"] = o.NumNodes
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxZoneInfoDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxZoneInfoDtAllOf := _HyperflexHxZoneInfoDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxZoneInfoDtAllOf); err == nil {
		*o = HyperflexHxZoneInfoDtAllOf(varHyperflexHxZoneInfoDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NumNodes")
		delete(additionalProperties, "Uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxZoneInfoDtAllOf struct {
	value *HyperflexHxZoneInfoDtAllOf
	isSet bool
}

func (v NullableHyperflexHxZoneInfoDtAllOf) Get() *HyperflexHxZoneInfoDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxZoneInfoDtAllOf) Set(val *HyperflexHxZoneInfoDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxZoneInfoDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxZoneInfoDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxZoneInfoDtAllOf(val *HyperflexHxZoneInfoDtAllOf) *NullableHyperflexHxZoneInfoDtAllOf {
	return &NullableHyperflexHxZoneInfoDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxZoneInfoDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxZoneInfoDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
