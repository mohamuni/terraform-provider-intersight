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

// StorageExternalLunAllOf Definition of the list of properties defined in 'storage.ExternalLun', excluding properties defined in parent classes.
type StorageExternalLunAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// LUN within the ports of the external storage system.
	ExternalLun *int64 `json:"ExternalLun,omitempty"`
	// WWN of the external storage system.
	ExternalWwn *string `json:"ExternalWwn,omitempty"`
	// Status of the external path.
	PathStatus *string `json:"PathStatus,omitempty"`
	// Port number of the local storage system.
	PortId *string `json:"PortId,omitempty"`
	// Priority within the external path group.
	Priority             *int64 `json:"Priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageExternalLunAllOf StorageExternalLunAllOf

// NewStorageExternalLunAllOf instantiates a new StorageExternalLunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageExternalLunAllOf(classId string, objectType string) *StorageExternalLunAllOf {
	this := StorageExternalLunAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageExternalLunAllOfWithDefaults instantiates a new StorageExternalLunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageExternalLunAllOfWithDefaults() *StorageExternalLunAllOf {
	this := StorageExternalLunAllOf{}
	var classId string = "storage.ExternalLun"
	this.ClassId = classId
	var objectType string = "storage.ExternalLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageExternalLunAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageExternalLunAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageExternalLunAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageExternalLunAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalLun returns the ExternalLun field value if set, zero value otherwise.
func (o *StorageExternalLunAllOf) GetExternalLun() int64 {
	if o == nil || o.ExternalLun == nil {
		var ret int64
		return ret
	}
	return *o.ExternalLun
}

// GetExternalLunOk returns a tuple with the ExternalLun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetExternalLunOk() (*int64, bool) {
	if o == nil || o.ExternalLun == nil {
		return nil, false
	}
	return o.ExternalLun, true
}

// HasExternalLun returns a boolean if a field has been set.
func (o *StorageExternalLunAllOf) HasExternalLun() bool {
	if o != nil && o.ExternalLun != nil {
		return true
	}

	return false
}

// SetExternalLun gets a reference to the given int64 and assigns it to the ExternalLun field.
func (o *StorageExternalLunAllOf) SetExternalLun(v int64) {
	o.ExternalLun = &v
}

// GetExternalWwn returns the ExternalWwn field value if set, zero value otherwise.
func (o *StorageExternalLunAllOf) GetExternalWwn() string {
	if o == nil || o.ExternalWwn == nil {
		var ret string
		return ret
	}
	return *o.ExternalWwn
}

// GetExternalWwnOk returns a tuple with the ExternalWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetExternalWwnOk() (*string, bool) {
	if o == nil || o.ExternalWwn == nil {
		return nil, false
	}
	return o.ExternalWwn, true
}

// HasExternalWwn returns a boolean if a field has been set.
func (o *StorageExternalLunAllOf) HasExternalWwn() bool {
	if o != nil && o.ExternalWwn != nil {
		return true
	}

	return false
}

// SetExternalWwn gets a reference to the given string and assigns it to the ExternalWwn field.
func (o *StorageExternalLunAllOf) SetExternalWwn(v string) {
	o.ExternalWwn = &v
}

// GetPathStatus returns the PathStatus field value if set, zero value otherwise.
func (o *StorageExternalLunAllOf) GetPathStatus() string {
	if o == nil || o.PathStatus == nil {
		var ret string
		return ret
	}
	return *o.PathStatus
}

// GetPathStatusOk returns a tuple with the PathStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetPathStatusOk() (*string, bool) {
	if o == nil || o.PathStatus == nil {
		return nil, false
	}
	return o.PathStatus, true
}

// HasPathStatus returns a boolean if a field has been set.
func (o *StorageExternalLunAllOf) HasPathStatus() bool {
	if o != nil && o.PathStatus != nil {
		return true
	}

	return false
}

// SetPathStatus gets a reference to the given string and assigns it to the PathStatus field.
func (o *StorageExternalLunAllOf) SetPathStatus(v string) {
	o.PathStatus = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageExternalLunAllOf) GetPortId() string {
	if o == nil || o.PortId == nil {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetPortIdOk() (*string, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageExternalLunAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageExternalLunAllOf) SetPortId(v string) {
	o.PortId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *StorageExternalLunAllOf) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalLunAllOf) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *StorageExternalLunAllOf) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *StorageExternalLunAllOf) SetPriority(v int64) {
	o.Priority = &v
}

func (o StorageExternalLunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExternalLun != nil {
		toSerialize["ExternalLun"] = o.ExternalLun
	}
	if o.ExternalWwn != nil {
		toSerialize["ExternalWwn"] = o.ExternalWwn
	}
	if o.PathStatus != nil {
		toSerialize["PathStatus"] = o.PathStatus
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageExternalLunAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageExternalLunAllOf := _StorageExternalLunAllOf{}

	if err = json.Unmarshal(bytes, &varStorageExternalLunAllOf); err == nil {
		*o = StorageExternalLunAllOf(varStorageExternalLunAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalLun")
		delete(additionalProperties, "ExternalWwn")
		delete(additionalProperties, "PathStatus")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "Priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageExternalLunAllOf struct {
	value *StorageExternalLunAllOf
	isSet bool
}

func (v NullableStorageExternalLunAllOf) Get() *StorageExternalLunAllOf {
	return v.value
}

func (v *NullableStorageExternalLunAllOf) Set(val *StorageExternalLunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageExternalLunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageExternalLunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageExternalLunAllOf(val *StorageExternalLunAllOf) *NullableStorageExternalLunAllOf {
	return &NullableStorageExternalLunAllOf{value: val, isSet: true}
}

func (v NullableStorageExternalLunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageExternalLunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
