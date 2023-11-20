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

// VnicEthVnicInventoryAllOf Definition of the list of properties defined in 'vnic.EthVnicInventory', excluding properties defined in parent classes.
type VnicEthVnicInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Name of the virtual ethernet interface.
	Name *string `json:"Name,omitempty"`
	// Enables usage of the Class of Service provided by the operating system.
	TrustHostCos         *bool                 `json:"TrustHostCos,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthVnicInventoryAllOf VnicEthVnicInventoryAllOf

// NewVnicEthVnicInventoryAllOf instantiates a new VnicEthVnicInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthVnicInventoryAllOf(classId string, objectType string) *VnicEthVnicInventoryAllOf {
	this := VnicEthVnicInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicEthVnicInventoryAllOfWithDefaults instantiates a new VnicEthVnicInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthVnicInventoryAllOfWithDefaults() *VnicEthVnicInventoryAllOf {
	this := VnicEthVnicInventoryAllOf{}
	var classId string = "vnic.EthVnicInventory"
	this.ClassId = classId
	var objectType string = "vnic.EthVnicInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthVnicInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthVnicInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthVnicInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthVnicInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicEthVnicInventoryAllOf) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicEthVnicInventoryAllOf) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicEthVnicInventoryAllOf) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VnicEthVnicInventoryAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VnicEthVnicInventoryAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VnicEthVnicInventoryAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicEthVnicInventoryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicEthVnicInventoryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicEthVnicInventoryAllOf) SetName(v string) {
	o.Name = &v
}

// GetTrustHostCos returns the TrustHostCos field value if set, zero value otherwise.
func (o *VnicEthVnicInventoryAllOf) GetTrustHostCos() bool {
	if o == nil || o.TrustHostCos == nil {
		var ret bool
		return ret
	}
	return *o.TrustHostCos
}

// GetTrustHostCosOk returns a tuple with the TrustHostCos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetTrustHostCosOk() (*bool, bool) {
	if o == nil || o.TrustHostCos == nil {
		return nil, false
	}
	return o.TrustHostCos, true
}

// HasTrustHostCos returns a boolean if a field has been set.
func (o *VnicEthVnicInventoryAllOf) HasTrustHostCos() bool {
	if o != nil && o.TrustHostCos != nil {
		return true
	}

	return false
}

// SetTrustHostCos gets a reference to the given bool and assigns it to the TrustHostCos field.
func (o *VnicEthVnicInventoryAllOf) SetTrustHostCos(v bool) {
	o.TrustHostCos = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicEthVnicInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicEthVnicInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicEthVnicInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicEthVnicInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TrustHostCos != nil {
		toSerialize["TrustHostCos"] = o.TrustHostCos
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthVnicInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicEthVnicInventoryAllOf := _VnicEthVnicInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varVnicEthVnicInventoryAllOf); err == nil {
		*o = VnicEthVnicInventoryAllOf(varVnicEthVnicInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "TrustHostCos")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthVnicInventoryAllOf struct {
	value *VnicEthVnicInventoryAllOf
	isSet bool
}

func (v NullableVnicEthVnicInventoryAllOf) Get() *VnicEthVnicInventoryAllOf {
	return v.value
}

func (v *NullableVnicEthVnicInventoryAllOf) Set(val *VnicEthVnicInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthVnicInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthVnicInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthVnicInventoryAllOf(val *VnicEthVnicInventoryAllOf) *NullableVnicEthVnicInventoryAllOf {
	return &NullableVnicEthVnicInventoryAllOf{value: val, isSet: true}
}

func (v NullableVnicEthVnicInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthVnicInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
