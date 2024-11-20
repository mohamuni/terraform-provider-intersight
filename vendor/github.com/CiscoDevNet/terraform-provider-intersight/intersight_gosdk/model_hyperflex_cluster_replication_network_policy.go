/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the HyperflexClusterReplicationNetworkPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexClusterReplicationNetworkPolicy{}

// HyperflexClusterReplicationNetworkPolicy Specifies all replication network parameters for the cluster.
type HyperflexClusterReplicationNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bandwidth for the Replication network in Mbps.
	ReplicationBandwidthMbps *int64                 `json:"ReplicationBandwidthMbps,omitempty"`
	ReplicationIpranges      []HyperflexIpAddrRange `json:"ReplicationIpranges,omitempty"`
	// MTU for the Replication network.
	ReplicationMtu  *int64                     `json:"ReplicationMtu,omitempty"`
	ReplicationVlan NullableHyperflexNamedVlan `json:"ReplicationVlan,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterReplicationNetworkPolicy HyperflexClusterReplicationNetworkPolicy

// NewHyperflexClusterReplicationNetworkPolicy instantiates a new HyperflexClusterReplicationNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterReplicationNetworkPolicy(classId string, objectType string) *HyperflexClusterReplicationNetworkPolicy {
	this := HyperflexClusterReplicationNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var replicationBandwidthMbps int64 = 0
	this.ReplicationBandwidthMbps = &replicationBandwidthMbps
	var replicationMtu int64 = 1500
	this.ReplicationMtu = &replicationMtu
	return &this
}

// NewHyperflexClusterReplicationNetworkPolicyWithDefaults instantiates a new HyperflexClusterReplicationNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterReplicationNetworkPolicyWithDefaults() *HyperflexClusterReplicationNetworkPolicy {
	this := HyperflexClusterReplicationNetworkPolicy{}
	var classId string = "hyperflex.ClusterReplicationNetworkPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterReplicationNetworkPolicy"
	this.ObjectType = objectType
	var replicationBandwidthMbps int64 = 0
	this.ReplicationBandwidthMbps = &replicationBandwidthMbps
	var replicationMtu int64 = 1500
	this.ReplicationMtu = &replicationMtu
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterReplicationNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterReplicationNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ClusterReplicationNetworkPolicy" of the ClassId field.
func (o *HyperflexClusterReplicationNetworkPolicy) GetDefaultClassId() interface{} {
	return "hyperflex.ClusterReplicationNetworkPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterReplicationNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterReplicationNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ClusterReplicationNetworkPolicy" of the ObjectType field.
func (o *HyperflexClusterReplicationNetworkPolicy) GetDefaultObjectType() interface{} {
	return "hyperflex.ClusterReplicationNetworkPolicy"
}

// GetReplicationBandwidthMbps returns the ReplicationBandwidthMbps field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationBandwidthMbps() int64 {
	if o == nil || IsNil(o.ReplicationBandwidthMbps) {
		var ret int64
		return ret
	}
	return *o.ReplicationBandwidthMbps
}

// GetReplicationBandwidthMbpsOk returns a tuple with the ReplicationBandwidthMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationBandwidthMbpsOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplicationBandwidthMbps) {
		return nil, false
	}
	return o.ReplicationBandwidthMbps, true
}

// HasReplicationBandwidthMbps returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) HasReplicationBandwidthMbps() bool {
	if o != nil && !IsNil(o.ReplicationBandwidthMbps) {
		return true
	}

	return false
}

// SetReplicationBandwidthMbps gets a reference to the given int64 and assigns it to the ReplicationBandwidthMbps field.
func (o *HyperflexClusterReplicationNetworkPolicy) SetReplicationBandwidthMbps(v int64) {
	o.ReplicationBandwidthMbps = &v
}

// GetReplicationIpranges returns the ReplicationIpranges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationIpranges() []HyperflexIpAddrRange {
	if o == nil {
		var ret []HyperflexIpAddrRange
		return ret
	}
	return o.ReplicationIpranges
}

// GetReplicationIprangesOk returns a tuple with the ReplicationIpranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationIprangesOk() ([]HyperflexIpAddrRange, bool) {
	if o == nil || IsNil(o.ReplicationIpranges) {
		return nil, false
	}
	return o.ReplicationIpranges, true
}

// HasReplicationIpranges returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) HasReplicationIpranges() bool {
	if o != nil && !IsNil(o.ReplicationIpranges) {
		return true
	}

	return false
}

// SetReplicationIpranges gets a reference to the given []HyperflexIpAddrRange and assigns it to the ReplicationIpranges field.
func (o *HyperflexClusterReplicationNetworkPolicy) SetReplicationIpranges(v []HyperflexIpAddrRange) {
	o.ReplicationIpranges = v
}

// GetReplicationMtu returns the ReplicationMtu field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationMtu() int64 {
	if o == nil || IsNil(o.ReplicationMtu) {
		var ret int64
		return ret
	}
	return *o.ReplicationMtu
}

// GetReplicationMtuOk returns a tuple with the ReplicationMtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplicationMtu) {
		return nil, false
	}
	return o.ReplicationMtu, true
}

// HasReplicationMtu returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) HasReplicationMtu() bool {
	if o != nil && !IsNil(o.ReplicationMtu) {
		return true
	}

	return false
}

// SetReplicationMtu gets a reference to the given int64 and assigns it to the ReplicationMtu field.
func (o *HyperflexClusterReplicationNetworkPolicy) SetReplicationMtu(v int64) {
	o.ReplicationMtu = &v
}

// GetReplicationVlan returns the ReplicationVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationVlan() HyperflexNamedVlan {
	if o == nil || IsNil(o.ReplicationVlan.Get()) {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.ReplicationVlan.Get()
}

// GetReplicationVlanOk returns a tuple with the ReplicationVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicy) GetReplicationVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationVlan.Get(), o.ReplicationVlan.IsSet()
}

// HasReplicationVlan returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) HasReplicationVlan() bool {
	if o != nil && o.ReplicationVlan.IsSet() {
		return true
	}

	return false
}

// SetReplicationVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the ReplicationVlan field.
func (o *HyperflexClusterReplicationNetworkPolicy) SetReplicationVlan(v HyperflexNamedVlan) {
	o.ReplicationVlan.Set(&v)
}

// SetReplicationVlanNil sets the value for ReplicationVlan to be an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicy) SetReplicationVlanNil() {
	o.ReplicationVlan.Set(nil)
}

// UnsetReplicationVlan ensures that no value is present for ReplicationVlan, not even an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicy) UnsetReplicationVlan() {
	o.ReplicationVlan.Unset()
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicy) GetClusterProfilesOk() ([]HyperflexClusterProfileRelationship, bool) {
	if o == nil || IsNil(o.ClusterProfiles) {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) HasClusterProfiles() bool {
	if o != nil && !IsNil(o.ClusterProfiles) {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterReplicationNetworkPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterReplicationNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o HyperflexClusterReplicationNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexClusterReplicationNetworkPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ReplicationBandwidthMbps) {
		toSerialize["ReplicationBandwidthMbps"] = o.ReplicationBandwidthMbps
	}
	if o.ReplicationIpranges != nil {
		toSerialize["ReplicationIpranges"] = o.ReplicationIpranges
	}
	if !IsNil(o.ReplicationMtu) {
		toSerialize["ReplicationMtu"] = o.ReplicationMtu
	}
	if o.ReplicationVlan.IsSet() {
		toSerialize["ReplicationVlan"] = o.ReplicationVlan.Get()
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexClusterReplicationNetworkPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type HyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Bandwidth for the Replication network in Mbps.
		ReplicationBandwidthMbps *int64                 `json:"ReplicationBandwidthMbps,omitempty"`
		ReplicationIpranges      []HyperflexIpAddrRange `json:"ReplicationIpranges,omitempty"`
		// MTU for the Replication network.
		ReplicationMtu  *int64                     `json:"ReplicationMtu,omitempty"`
		ReplicationVlan NullableHyperflexNamedVlan `json:"ReplicationVlan,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
		Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct := HyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterReplicationNetworkPolicy := _HyperflexClusterReplicationNetworkPolicy{}
		varHyperflexClusterReplicationNetworkPolicy.ClassId = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ClassId
		varHyperflexClusterReplicationNetworkPolicy.ObjectType = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterReplicationNetworkPolicy.ReplicationBandwidthMbps = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ReplicationBandwidthMbps
		varHyperflexClusterReplicationNetworkPolicy.ReplicationIpranges = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ReplicationIpranges
		varHyperflexClusterReplicationNetworkPolicy.ReplicationMtu = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ReplicationMtu
		varHyperflexClusterReplicationNetworkPolicy.ReplicationVlan = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ReplicationVlan
		varHyperflexClusterReplicationNetworkPolicy.ClusterProfiles = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexClusterReplicationNetworkPolicy.Organization = varHyperflexClusterReplicationNetworkPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexClusterReplicationNetworkPolicy(varHyperflexClusterReplicationNetworkPolicy)
	} else {
		return err
	}

	varHyperflexClusterReplicationNetworkPolicy := _HyperflexClusterReplicationNetworkPolicy{}

	err = json.Unmarshal(data, &varHyperflexClusterReplicationNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexClusterReplicationNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ReplicationBandwidthMbps")
		delete(additionalProperties, "ReplicationIpranges")
		delete(additionalProperties, "ReplicationMtu")
		delete(additionalProperties, "ReplicationVlan")
		delete(additionalProperties, "ClusterProfiles")
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

type NullableHyperflexClusterReplicationNetworkPolicy struct {
	value *HyperflexClusterReplicationNetworkPolicy
	isSet bool
}

func (v NullableHyperflexClusterReplicationNetworkPolicy) Get() *HyperflexClusterReplicationNetworkPolicy {
	return v.value
}

func (v *NullableHyperflexClusterReplicationNetworkPolicy) Set(val *HyperflexClusterReplicationNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterReplicationNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterReplicationNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterReplicationNetworkPolicy(val *HyperflexClusterReplicationNetworkPolicy) *NullableHyperflexClusterReplicationNetworkPolicy {
	return &NullableHyperflexClusterReplicationNetworkPolicy{value: val, isSet: true}
}

func (v NullableHyperflexClusterReplicationNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterReplicationNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
