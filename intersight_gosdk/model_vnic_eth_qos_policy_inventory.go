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
	"reflect"
	"strings"
)

// VnicEthQosPolicyInventory An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.
type VnicEthQosPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The burst traffic, in bytes, allowed on the vNIC.
	Burst *int64 `json:"Burst,omitempty"`
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts.
	Mtu *int64 `json:"Mtu,omitempty"`
	// The priority matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority *string `json:"Priority,omitempty"`
	// The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface.
	RateLimit *int64 `json:"RateLimit,omitempty"`
	// Enables usage of the Class of Service provided by the operating system.
	TrustHostCos         *bool                 `json:"TrustHostCos,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthQosPolicyInventory VnicEthQosPolicyInventory

// NewVnicEthQosPolicyInventory instantiates a new VnicEthQosPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthQosPolicyInventory(classId string, objectType string) *VnicEthQosPolicyInventory {
	this := VnicEthQosPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicEthQosPolicyInventoryWithDefaults instantiates a new VnicEthQosPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthQosPolicyInventoryWithDefaults() *VnicEthQosPolicyInventory {
	this := VnicEthQosPolicyInventory{}
	var classId string = "vnic.EthQosPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.EthQosPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthQosPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthQosPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthQosPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthQosPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBurst returns the Burst field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetBurst() int64 {
	if o == nil || o.Burst == nil {
		var ret int64
		return ret
	}
	return *o.Burst
}

// GetBurstOk returns a tuple with the Burst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetBurstOk() (*int64, bool) {
	if o == nil || o.Burst == nil {
		return nil, false
	}
	return o.Burst, true
}

// HasBurst returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasBurst() bool {
	if o != nil && o.Burst != nil {
		return true
	}

	return false
}

// SetBurst gets a reference to the given int64 and assigns it to the Burst field.
func (o *VnicEthQosPolicyInventory) SetBurst(v int64) {
	o.Burst = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicEthQosPolicyInventory) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VnicEthQosPolicyInventory) SetMtu(v int64) {
	o.Mtu = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *VnicEthQosPolicyInventory) SetPriority(v string) {
	o.Priority = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicEthQosPolicyInventory) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetTrustHostCos returns the TrustHostCos field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetTrustHostCos() bool {
	if o == nil || o.TrustHostCos == nil {
		var ret bool
		return ret
	}
	return *o.TrustHostCos
}

// GetTrustHostCosOk returns a tuple with the TrustHostCos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetTrustHostCosOk() (*bool, bool) {
	if o == nil || o.TrustHostCos == nil {
		return nil, false
	}
	return o.TrustHostCos, true
}

// HasTrustHostCos returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasTrustHostCos() bool {
	if o != nil && o.TrustHostCos != nil {
		return true
	}

	return false
}

// SetTrustHostCos gets a reference to the given bool and assigns it to the TrustHostCos field.
func (o *VnicEthQosPolicyInventory) SetTrustHostCos(v bool) {
	o.TrustHostCos = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicEthQosPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicEthQosPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicEthQosPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicEthQosPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Burst != nil {
		toSerialize["Burst"] = o.Burst
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.RateLimit != nil {
		toSerialize["RateLimit"] = o.RateLimit
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

func (o *VnicEthQosPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicEthQosPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The burst traffic, in bytes, allowed on the vNIC.
		Burst *int64 `json:"Burst,omitempty"`
		// Class of Service to be associated to the traffic on the virtual interface.
		Cos *int64 `json:"Cos,omitempty"`
		// The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts.
		Mtu *int64 `json:"Mtu,omitempty"`
		// The priority matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority *string `json:"Priority,omitempty"`
		// The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface.
		RateLimit *int64 `json:"RateLimit,omitempty"`
		// Enables usage of the Class of Service provided by the operating system.
		TrustHostCos *bool                 `json:"TrustHostCos,omitempty"`
		TargetMo     *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicEthQosPolicyInventoryWithoutEmbeddedStruct := VnicEthQosPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicEthQosPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthQosPolicyInventory := _VnicEthQosPolicyInventory{}
		varVnicEthQosPolicyInventory.ClassId = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicEthQosPolicyInventory.ObjectType = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicEthQosPolicyInventory.Burst = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.Burst
		varVnicEthQosPolicyInventory.Cos = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.Cos
		varVnicEthQosPolicyInventory.Mtu = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.Mtu
		varVnicEthQosPolicyInventory.Priority = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.Priority
		varVnicEthQosPolicyInventory.RateLimit = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.RateLimit
		varVnicEthQosPolicyInventory.TrustHostCos = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.TrustHostCos
		varVnicEthQosPolicyInventory.TargetMo = varVnicEthQosPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicEthQosPolicyInventory(varVnicEthQosPolicyInventory)
	} else {
		return err
	}

	varVnicEthQosPolicyInventory := _VnicEthQosPolicyInventory{}

	err = json.Unmarshal(bytes, &varVnicEthQosPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicEthQosPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Burst")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "RateLimit")
		delete(additionalProperties, "TrustHostCos")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableVnicEthQosPolicyInventory struct {
	value *VnicEthQosPolicyInventory
	isSet bool
}

func (v NullableVnicEthQosPolicyInventory) Get() *VnicEthQosPolicyInventory {
	return v.value
}

func (v *NullableVnicEthQosPolicyInventory) Set(val *VnicEthQosPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthQosPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthQosPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthQosPolicyInventory(val *VnicEthQosPolicyInventory) *NullableVnicEthQosPolicyInventory {
	return &NullableVnicEthQosPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicEthQosPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthQosPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
