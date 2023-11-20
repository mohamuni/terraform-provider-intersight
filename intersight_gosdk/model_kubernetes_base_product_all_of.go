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

// KubernetesBaseProductAllOf Definition of the list of properties defined in 'kubernetes.BaseProduct', excluding properties defined in parent classes.
type KubernetesBaseProductAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Optional description of a product.
	Description *string `json:"Description,omitempty"`
	// Device Id of a product, which is unique within a vendor.
	DeviceId *int64 `json:"DeviceId,omitempty"`
	// Display Name of a product.
	Name *string `json:"Name,omitempty"`
	// Vendor Id of a product. Each vendor has a globally unique Id, for example 0x10DE for Nvidia.
	VendorId             *int64                         `json:"VendorId,omitempty"`
	Catalog              *KubernetesCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaseProductAllOf KubernetesBaseProductAllOf

// NewKubernetesBaseProductAllOf instantiates a new KubernetesBaseProductAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaseProductAllOf(classId string, objectType string) *KubernetesBaseProductAllOf {
	this := KubernetesBaseProductAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesBaseProductAllOfWithDefaults instantiates a new KubernetesBaseProductAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaseProductAllOfWithDefaults() *KubernetesBaseProductAllOf {
	this := KubernetesBaseProductAllOf{}
	var classId string = "kubernetes.NvidiaGpuProduct"
	this.ClassId = classId
	var objectType string = "kubernetes.NvidiaGpuProduct"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaseProductAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaseProductAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaseProductAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaseProductAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KubernetesBaseProductAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KubernetesBaseProductAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KubernetesBaseProductAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *KubernetesBaseProductAllOf) GetDeviceId() int64 {
	if o == nil || o.DeviceId == nil {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetDeviceIdOk() (*int64, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *KubernetesBaseProductAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *KubernetesBaseProductAllOf) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesBaseProductAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesBaseProductAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesBaseProductAllOf) SetName(v string) {
	o.Name = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *KubernetesBaseProductAllOf) GetVendorId() int64 {
	if o == nil || o.VendorId == nil {
		var ret int64
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetVendorIdOk() (*int64, bool) {
	if o == nil || o.VendorId == nil {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *KubernetesBaseProductAllOf) HasVendorId() bool {
	if o != nil && o.VendorId != nil {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given int64 and assigns it to the VendorId field.
func (o *KubernetesBaseProductAllOf) SetVendorId(v int64) {
	o.VendorId = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *KubernetesBaseProductAllOf) GetCatalog() KubernetesCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret KubernetesCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseProductAllOf) GetCatalogOk() (*KubernetesCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *KubernetesBaseProductAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given KubernetesCatalogRelationship and assigns it to the Catalog field.
func (o *KubernetesBaseProductAllOf) SetCatalog(v KubernetesCatalogRelationship) {
	o.Catalog = &v
}

func (o KubernetesBaseProductAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VendorId != nil {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesBaseProductAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesBaseProductAllOf := _KubernetesBaseProductAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesBaseProductAllOf); err == nil {
		*o = KubernetesBaseProductAllOf(varKubernetesBaseProductAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VendorId")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesBaseProductAllOf struct {
	value *KubernetesBaseProductAllOf
	isSet bool
}

func (v NullableKubernetesBaseProductAllOf) Get() *KubernetesBaseProductAllOf {
	return v.value
}

func (v *NullableKubernetesBaseProductAllOf) Set(val *KubernetesBaseProductAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaseProductAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaseProductAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaseProductAllOf(val *KubernetesBaseProductAllOf) *NullableKubernetesBaseProductAllOf {
	return &NullableKubernetesBaseProductAllOf{value: val, isSet: true}
}

func (v NullableKubernetesBaseProductAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaseProductAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
