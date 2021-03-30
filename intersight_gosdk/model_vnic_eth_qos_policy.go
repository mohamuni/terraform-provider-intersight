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

// VnicEthQosPolicy An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.
type VnicEthQosPolicy struct {
	PolicyAbstractPolicy
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
	// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority *string `json:"Priority,omitempty"`
	// The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	RateLimit *int64 `json:"RateLimit,omitempty"`
	// Enables usage of the Class of Service provided by the operating system.
	TrustHostCos         *bool                                 `json:"TrustHostCos,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthQosPolicy VnicEthQosPolicy

// NewVnicEthQosPolicy instantiates a new VnicEthQosPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthQosPolicy(classId string, objectType string) *VnicEthQosPolicy {
	this := VnicEthQosPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var burst int64 = 1024
	this.Burst = &burst
	var cos int64 = 0
	this.Cos = &cos
	var mtu int64 = 1500
	this.Mtu = &mtu
	var priority string = "Best Effort"
	this.Priority = &priority
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	var trustHostCos bool = false
	this.TrustHostCos = &trustHostCos
	return &this
}

// NewVnicEthQosPolicyWithDefaults instantiates a new VnicEthQosPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthQosPolicyWithDefaults() *VnicEthQosPolicy {
	this := VnicEthQosPolicy{}
	var classId string = "vnic.EthQosPolicy"
	this.ClassId = classId
	var objectType string = "vnic.EthQosPolicy"
	this.ObjectType = objectType
	var burst int64 = 1024
	this.Burst = &burst
	var cos int64 = 0
	this.Cos = &cos
	var mtu int64 = 1500
	this.Mtu = &mtu
	var priority string = "Best Effort"
	this.Priority = &priority
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	var trustHostCos bool = false
	this.TrustHostCos = &trustHostCos
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthQosPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthQosPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthQosPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthQosPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBurst returns the Burst field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetBurst() int64 {
	if o == nil || o.Burst == nil {
		var ret int64
		return ret
	}
	return *o.Burst
}

// GetBurstOk returns a tuple with the Burst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetBurstOk() (*int64, bool) {
	if o == nil || o.Burst == nil {
		return nil, false
	}
	return o.Burst, true
}

// HasBurst returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasBurst() bool {
	if o != nil && o.Burst != nil {
		return true
	}

	return false
}

// SetBurst gets a reference to the given int64 and assigns it to the Burst field.
func (o *VnicEthQosPolicy) SetBurst(v int64) {
	o.Burst = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicEthQosPolicy) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VnicEthQosPolicy) SetMtu(v int64) {
	o.Mtu = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *VnicEthQosPolicy) SetPriority(v string) {
	o.Priority = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicEthQosPolicy) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetTrustHostCos returns the TrustHostCos field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetTrustHostCos() bool {
	if o == nil || o.TrustHostCos == nil {
		var ret bool
		return ret
	}
	return *o.TrustHostCos
}

// GetTrustHostCosOk returns a tuple with the TrustHostCos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetTrustHostCosOk() (*bool, bool) {
	if o == nil || o.TrustHostCos == nil {
		return nil, false
	}
	return o.TrustHostCos, true
}

// HasTrustHostCos returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasTrustHostCos() bool {
	if o != nil && o.TrustHostCos != nil {
		return true
	}

	return false
}

// SetTrustHostCos gets a reference to the given bool and assigns it to the TrustHostCos field.
func (o *VnicEthQosPolicy) SetTrustHostCos(v bool) {
	o.TrustHostCos = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthQosPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthQosPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthQosPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthQosPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
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
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthQosPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicEthQosPolicyWithoutEmbeddedStruct struct {
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
		// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority *string `json:"Priority,omitempty"`
		// The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
		RateLimit *int64 `json:"RateLimit,omitempty"`
		// Enables usage of the Class of Service provided by the operating system.
		TrustHostCos *bool                                 `json:"TrustHostCos,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicEthQosPolicyWithoutEmbeddedStruct := VnicEthQosPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicEthQosPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthQosPolicy := _VnicEthQosPolicy{}
		varVnicEthQosPolicy.ClassId = varVnicEthQosPolicyWithoutEmbeddedStruct.ClassId
		varVnicEthQosPolicy.ObjectType = varVnicEthQosPolicyWithoutEmbeddedStruct.ObjectType
		varVnicEthQosPolicy.Burst = varVnicEthQosPolicyWithoutEmbeddedStruct.Burst
		varVnicEthQosPolicy.Cos = varVnicEthQosPolicyWithoutEmbeddedStruct.Cos
		varVnicEthQosPolicy.Mtu = varVnicEthQosPolicyWithoutEmbeddedStruct.Mtu
		varVnicEthQosPolicy.Priority = varVnicEthQosPolicyWithoutEmbeddedStruct.Priority
		varVnicEthQosPolicy.RateLimit = varVnicEthQosPolicyWithoutEmbeddedStruct.RateLimit
		varVnicEthQosPolicy.TrustHostCos = varVnicEthQosPolicyWithoutEmbeddedStruct.TrustHostCos
		varVnicEthQosPolicy.Organization = varVnicEthQosPolicyWithoutEmbeddedStruct.Organization
		*o = VnicEthQosPolicy(varVnicEthQosPolicy)
	} else {
		return err
	}

	varVnicEthQosPolicy := _VnicEthQosPolicy{}

	err = json.Unmarshal(bytes, &varVnicEthQosPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicEthQosPolicy.PolicyAbstractPolicy
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
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableVnicEthQosPolicy struct {
	value *VnicEthQosPolicy
	isSet bool
}

func (v NullableVnicEthQosPolicy) Get() *VnicEthQosPolicy {
	return v.value
}

func (v *NullableVnicEthQosPolicy) Set(val *VnicEthQosPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthQosPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthQosPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthQosPolicy(val *VnicEthQosPolicy) *NullableVnicEthQosPolicy {
	return &NullableVnicEthQosPolicy{value: val, isSet: true}
}

func (v NullableVnicEthQosPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthQosPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
