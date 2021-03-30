/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// BulkRestSubRequest An individual REST API request to be executed as part of a bulk request.
type BulkRestSubRequest struct {
	BulkSubRequest
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string    `json:"ObjectType"`
	Body       *MoBaseMo `json:"Body,omitempty"`
	// Used with PATCH & DELETE actions. The moid of an existing object instance.
	TargetMoid           *string `json:"TargetMoid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkRestSubRequest BulkRestSubRequest

// NewBulkRestSubRequest instantiates a new BulkRestSubRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRestSubRequest(classId string, objectType string) *BulkRestSubRequest {
	this := BulkRestSubRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkRestSubRequestWithDefaults instantiates a new BulkRestSubRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRestSubRequestWithDefaults() *BulkRestSubRequest {
	this := BulkRestSubRequest{}
	var classId string = "bulk.RestSubRequest"
	this.ClassId = classId
	var objectType string = "bulk.RestSubRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkRestSubRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkRestSubRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkRestSubRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkRestSubRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BulkRestSubRequest) GetBody() MoBaseMo {
	if o == nil || o.Body == nil {
		var ret MoBaseMo
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequest) GetBodyOk() (*MoBaseMo, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BulkRestSubRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given MoBaseMo and assigns it to the Body field.
func (o *BulkRestSubRequest) SetBody(v MoBaseMo) {
	o.Body = &v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *BulkRestSubRequest) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequest) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *BulkRestSubRequest) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *BulkRestSubRequest) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

func (o BulkRestSubRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBulkSubRequest, errBulkSubRequest := json.Marshal(o.BulkSubRequest)
	if errBulkSubRequest != nil {
		return []byte{}, errBulkSubRequest
	}
	errBulkSubRequest = json.Unmarshal([]byte(serializedBulkSubRequest), &toSerialize)
	if errBulkSubRequest != nil {
		return []byte{}, errBulkSubRequest
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkRestSubRequest) UnmarshalJSON(bytes []byte) (err error) {
	type BulkRestSubRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string    `json:"ObjectType"`
		Body       *MoBaseMo `json:"Body,omitempty"`
		// Used with PATCH & DELETE actions. The moid of an existing object instance.
		TargetMoid *string `json:"TargetMoid,omitempty"`
	}

	varBulkRestSubRequestWithoutEmbeddedStruct := BulkRestSubRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkRestSubRequestWithoutEmbeddedStruct)
	if err == nil {
		varBulkRestSubRequest := _BulkRestSubRequest{}
		varBulkRestSubRequest.ClassId = varBulkRestSubRequestWithoutEmbeddedStruct.ClassId
		varBulkRestSubRequest.ObjectType = varBulkRestSubRequestWithoutEmbeddedStruct.ObjectType
		varBulkRestSubRequest.Body = varBulkRestSubRequestWithoutEmbeddedStruct.Body
		varBulkRestSubRequest.TargetMoid = varBulkRestSubRequestWithoutEmbeddedStruct.TargetMoid
		*o = BulkRestSubRequest(varBulkRestSubRequest)
	} else {
		return err
	}

	varBulkRestSubRequest := _BulkRestSubRequest{}

	err = json.Unmarshal(bytes, &varBulkRestSubRequest)
	if err == nil {
		o.BulkSubRequest = varBulkRestSubRequest.BulkSubRequest
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "TargetMoid")

		// remove fields from embedded structs
		reflectBulkSubRequest := reflect.ValueOf(o.BulkSubRequest)
		for i := 0; i < reflectBulkSubRequest.Type().NumField(); i++ {
			t := reflectBulkSubRequest.Type().Field(i)

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

type NullableBulkRestSubRequest struct {
	value *BulkRestSubRequest
	isSet bool
}

func (v NullableBulkRestSubRequest) Get() *BulkRestSubRequest {
	return v.value
}

func (v *NullableBulkRestSubRequest) Set(val *BulkRestSubRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRestSubRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRestSubRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRestSubRequest(val *BulkRestSubRequest) *NullableBulkRestSubRequest {
	return &NullableBulkRestSubRequest{value: val, isSet: true}
}

func (v NullableBulkRestSubRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRestSubRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
