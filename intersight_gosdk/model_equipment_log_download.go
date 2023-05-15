/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EquipmentLogDownload Download the log collected from end point.
type EquipmentLogDownload struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	Server               *ComputePhysicalRelationship `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentLogDownload EquipmentLogDownload

// NewEquipmentLogDownload instantiates a new EquipmentLogDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentLogDownload(classId string, objectType string) *EquipmentLogDownload {
	this := EquipmentLogDownload{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentLogDownloadWithDefaults instantiates a new EquipmentLogDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentLogDownloadWithDefaults() *EquipmentLogDownload {
	this := EquipmentLogDownload{}
	var classId string = "equipment.LogDownload"
	this.ClassId = classId
	var objectType string = "equipment.LogDownload"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentLogDownload) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentLogDownload) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentLogDownload) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentLogDownload) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentLogDownload) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentLogDownload) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *EquipmentLogDownload) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLogDownload) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *EquipmentLogDownload) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *EquipmentLogDownload) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o EquipmentLogDownload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentLogDownload) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentLogDownloadWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                       `json:"ObjectType"`
		Server     *ComputePhysicalRelationship `json:"Server,omitempty"`
	}

	varEquipmentLogDownloadWithoutEmbeddedStruct := EquipmentLogDownloadWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentLogDownloadWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentLogDownload := _EquipmentLogDownload{}
		varEquipmentLogDownload.ClassId = varEquipmentLogDownloadWithoutEmbeddedStruct.ClassId
		varEquipmentLogDownload.ObjectType = varEquipmentLogDownloadWithoutEmbeddedStruct.ObjectType
		varEquipmentLogDownload.Server = varEquipmentLogDownloadWithoutEmbeddedStruct.Server
		*o = EquipmentLogDownload(varEquipmentLogDownload)
	} else {
		return err
	}

	varEquipmentLogDownload := _EquipmentLogDownload{}

	err = json.Unmarshal(bytes, &varEquipmentLogDownload)
	if err == nil {
		o.MoBaseMo = varEquipmentLogDownload.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableEquipmentLogDownload struct {
	value *EquipmentLogDownload
	isSet bool
}

func (v NullableEquipmentLogDownload) Get() *EquipmentLogDownload {
	return v.value
}

func (v *NullableEquipmentLogDownload) Set(val *EquipmentLogDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentLogDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentLogDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentLogDownload(val *EquipmentLogDownload) *NullableEquipmentLogDownload {
	return &NullableEquipmentLogDownload{value: val, isSet: true}
}

func (v NullableEquipmentLogDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentLogDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}