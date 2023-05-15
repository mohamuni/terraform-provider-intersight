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

// VnicEthVnicInventory QoS settings of server’s vNIC.
type VnicEthVnicInventory struct {
	PolicyAbstractInventory
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

type _VnicEthVnicInventory VnicEthVnicInventory

// NewVnicEthVnicInventory instantiates a new VnicEthVnicInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthVnicInventory(classId string, objectType string) *VnicEthVnicInventory {
	this := VnicEthVnicInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicEthVnicInventoryWithDefaults instantiates a new VnicEthVnicInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthVnicInventoryWithDefaults() *VnicEthVnicInventory {
	this := VnicEthVnicInventory{}
	var classId string = "vnic.EthVnicInventory"
	this.ClassId = classId
	var objectType string = "vnic.EthVnicInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthVnicInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthVnicInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthVnicInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthVnicInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicEthVnicInventory) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicEthVnicInventory) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicEthVnicInventory) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VnicEthVnicInventory) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VnicEthVnicInventory) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VnicEthVnicInventory) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicEthVnicInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicEthVnicInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicEthVnicInventory) SetName(v string) {
	o.Name = &v
}

// GetTrustHostCos returns the TrustHostCos field value if set, zero value otherwise.
func (o *VnicEthVnicInventory) GetTrustHostCos() bool {
	if o == nil || o.TrustHostCos == nil {
		var ret bool
		return ret
	}
	return *o.TrustHostCos
}

// GetTrustHostCosOk returns a tuple with the TrustHostCos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetTrustHostCosOk() (*bool, bool) {
	if o == nil || o.TrustHostCos == nil {
		return nil, false
	}
	return o.TrustHostCos, true
}

// HasTrustHostCos returns a boolean if a field has been set.
func (o *VnicEthVnicInventory) HasTrustHostCos() bool {
	if o != nil && o.TrustHostCos != nil {
		return true
	}

	return false
}

// SetTrustHostCos gets a reference to the given bool and assigns it to the TrustHostCos field.
func (o *VnicEthVnicInventory) SetTrustHostCos(v bool) {
	o.TrustHostCos = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicEthVnicInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVnicInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicEthVnicInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicEthVnicInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicEthVnicInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractInventory, errPolicyAbstractInventory := json.Marshal(o.PolicyAbstractInventory)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	errPolicyAbstractInventory = json.Unmarshal([]byte(serializedPolicyAbstractInventory), &toSerialize)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
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

func (o *VnicEthVnicInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicEthVnicInventoryWithoutEmbeddedStruct struct {
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
		TrustHostCos *bool                 `json:"TrustHostCos,omitempty"`
		TargetMo     *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicEthVnicInventoryWithoutEmbeddedStruct := VnicEthVnicInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicEthVnicInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthVnicInventory := _VnicEthVnicInventory{}
		varVnicEthVnicInventory.ClassId = varVnicEthVnicInventoryWithoutEmbeddedStruct.ClassId
		varVnicEthVnicInventory.ObjectType = varVnicEthVnicInventoryWithoutEmbeddedStruct.ObjectType
		varVnicEthVnicInventory.Cos = varVnicEthVnicInventoryWithoutEmbeddedStruct.Cos
		varVnicEthVnicInventory.Mtu = varVnicEthVnicInventoryWithoutEmbeddedStruct.Mtu
		varVnicEthVnicInventory.Name = varVnicEthVnicInventoryWithoutEmbeddedStruct.Name
		varVnicEthVnicInventory.TrustHostCos = varVnicEthVnicInventoryWithoutEmbeddedStruct.TrustHostCos
		varVnicEthVnicInventory.TargetMo = varVnicEthVnicInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicEthVnicInventory(varVnicEthVnicInventory)
	} else {
		return err
	}

	varVnicEthVnicInventory := _VnicEthVnicInventory{}

	err = json.Unmarshal(bytes, &varVnicEthVnicInventory)
	if err == nil {
		o.PolicyAbstractInventory = varVnicEthVnicInventory.PolicyAbstractInventory
	} else {
		return err
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

		// remove fields from embedded structs
		reflectPolicyAbstractInventory := reflect.ValueOf(o.PolicyAbstractInventory)
		for i := 0; i < reflectPolicyAbstractInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractInventory.Type().Field(i)

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

type NullableVnicEthVnicInventory struct {
	value *VnicEthVnicInventory
	isSet bool
}

func (v NullableVnicEthVnicInventory) Get() *VnicEthVnicInventory {
	return v.value
}

func (v *NullableVnicEthVnicInventory) Set(val *VnicEthVnicInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthVnicInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthVnicInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthVnicInventory(val *VnicEthVnicInventory) *NullableVnicEthVnicInventory {
	return &NullableVnicEthVnicInventory{value: val, isSet: true}
}

func (v NullableVnicEthVnicInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthVnicInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}