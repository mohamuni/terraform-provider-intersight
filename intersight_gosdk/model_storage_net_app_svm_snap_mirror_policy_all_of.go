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

// StorageNetAppSvmSnapMirrorPolicyAllOf Definition of the list of properties defined in 'storage.NetAppSvmSnapMirrorPolicy', excluding properties defined in parent classes.
type StorageNetAppSvmSnapMirrorPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The storage virtual machine name for the policy.
	SvmName              *string                             `json:"SvmName,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppSvmSnapMirrorPolicyAllOf StorageNetAppSvmSnapMirrorPolicyAllOf

// NewStorageNetAppSvmSnapMirrorPolicyAllOf instantiates a new StorageNetAppSvmSnapMirrorPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppSvmSnapMirrorPolicyAllOf(classId string, objectType string) *StorageNetAppSvmSnapMirrorPolicyAllOf {
	this := StorageNetAppSvmSnapMirrorPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppSvmSnapMirrorPolicyAllOfWithDefaults instantiates a new StorageNetAppSvmSnapMirrorPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppSvmSnapMirrorPolicyAllOfWithDefaults() *StorageNetAppSvmSnapMirrorPolicyAllOf {
	this := StorageNetAppSvmSnapMirrorPolicyAllOf{}
	var classId string = "storage.NetAppSvmSnapMirrorPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppSvmSnapMirrorPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetSvmName() string {
	if o == nil || o.SvmName == nil {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetSvmNameOk() (*string, bool) {
	if o == nil || o.SvmName == nil {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) HasSvmName() bool {
	if o != nil && o.SvmName != nil {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) SetSvmName(v string) {
	o.SvmName = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppSvmSnapMirrorPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SvmName != nil {
		toSerialize["SvmName"] = o.SvmName
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppSvmSnapMirrorPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppSvmSnapMirrorPolicyAllOf := _StorageNetAppSvmSnapMirrorPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppSvmSnapMirrorPolicyAllOf); err == nil {
		*o = StorageNetAppSvmSnapMirrorPolicyAllOf(varStorageNetAppSvmSnapMirrorPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppSvmSnapMirrorPolicyAllOf struct {
	value *StorageNetAppSvmSnapMirrorPolicyAllOf
	isSet bool
}

func (v NullableStorageNetAppSvmSnapMirrorPolicyAllOf) Get() *StorageNetAppSvmSnapMirrorPolicyAllOf {
	return v.value
}

func (v *NullableStorageNetAppSvmSnapMirrorPolicyAllOf) Set(val *StorageNetAppSvmSnapMirrorPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppSvmSnapMirrorPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppSvmSnapMirrorPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppSvmSnapMirrorPolicyAllOf(val *StorageNetAppSvmSnapMirrorPolicyAllOf) *NullableStorageNetAppSvmSnapMirrorPolicyAllOf {
	return &NullableStorageNetAppSvmSnapMirrorPolicyAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppSvmSnapMirrorPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppSvmSnapMirrorPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
