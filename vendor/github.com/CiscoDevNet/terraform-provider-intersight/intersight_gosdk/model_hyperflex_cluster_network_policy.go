/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the HyperflexClusterNetworkPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexClusterNetworkPolicy{}

// HyperflexClusterNetworkPolicy A policy specifying VLANs for management, VM migration, and VM traffic.
type HyperflexClusterNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The mode configured on the CIMC management interface. * `OutOfBand` - The server uses out-of-band management, i.e. management traffic traverses through the management interfaces on the UCS Fabric Interconnects. * `InBand` - The server uses in-band management, i.e. management traffic traverses through the data uplink ports on the UCS Fabric Interconnects.
	CimcManagementMode *string `json:"CimcManagementMode,omitempty"`
	// Enable or disable Jumbo Frames (MTU=9000). Jumbo Frames are used by Storage Network, VM Migration Network.
	JumboFrame     *bool                               `json:"JumboFrame,omitempty"`
	KvmIpRange     NullableHyperflexIpAddrRange        `json:"KvmIpRange,omitempty"`
	MacPrefixRange NullableHyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty"`
	MgmtVlan       NullableHyperflexNamedVlan          `json:"MgmtVlan,omitempty"`
	// Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only. * `default` - Current default value set on the hardware platform. * `1G` - A link speed of 1 gigabit per second. * `10G` - A link speed of 10 gigabits per second or above.
	UplinkSpeed     *string                    `json:"UplinkSpeed,omitempty"`
	VmMigrationVlan NullableHyperflexNamedVlan `json:"VmMigrationVlan,omitempty"`
	VmNetworkVlans  []HyperflexNamedVlan       `json:"VmNetworkVlans,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterNetworkPolicy HyperflexClusterNetworkPolicy

// NewHyperflexClusterNetworkPolicy instantiates a new HyperflexClusterNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterNetworkPolicy(classId string, objectType string) *HyperflexClusterNetworkPolicy {
	this := HyperflexClusterNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var uplinkSpeed string = "default"
	this.UplinkSpeed = &uplinkSpeed
	return &this
}

// NewHyperflexClusterNetworkPolicyWithDefaults instantiates a new HyperflexClusterNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterNetworkPolicyWithDefaults() *HyperflexClusterNetworkPolicy {
	this := HyperflexClusterNetworkPolicy{}
	var classId string = "hyperflex.ClusterNetworkPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterNetworkPolicy"
	this.ObjectType = objectType
	var uplinkSpeed string = "default"
	this.UplinkSpeed = &uplinkSpeed
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ClusterNetworkPolicy" of the ClassId field.
func (o *HyperflexClusterNetworkPolicy) GetDefaultClassId() interface{} {
	return "hyperflex.ClusterNetworkPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ClusterNetworkPolicy" of the ObjectType field.
func (o *HyperflexClusterNetworkPolicy) GetDefaultObjectType() interface{} {
	return "hyperflex.ClusterNetworkPolicy"
}

// GetCimcManagementMode returns the CimcManagementMode field value if set, zero value otherwise.
func (o *HyperflexClusterNetworkPolicy) GetCimcManagementMode() string {
	if o == nil || IsNil(o.CimcManagementMode) {
		var ret string
		return ret
	}
	return *o.CimcManagementMode
}

// GetCimcManagementModeOk returns a tuple with the CimcManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetCimcManagementModeOk() (*string, bool) {
	if o == nil || IsNil(o.CimcManagementMode) {
		return nil, false
	}
	return o.CimcManagementMode, true
}

// HasCimcManagementMode returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasCimcManagementMode() bool {
	if o != nil && !IsNil(o.CimcManagementMode) {
		return true
	}

	return false
}

// SetCimcManagementMode gets a reference to the given string and assigns it to the CimcManagementMode field.
func (o *HyperflexClusterNetworkPolicy) SetCimcManagementMode(v string) {
	o.CimcManagementMode = &v
}

// GetJumboFrame returns the JumboFrame field value if set, zero value otherwise.
func (o *HyperflexClusterNetworkPolicy) GetJumboFrame() bool {
	if o == nil || IsNil(o.JumboFrame) {
		var ret bool
		return ret
	}
	return *o.JumboFrame
}

// GetJumboFrameOk returns a tuple with the JumboFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetJumboFrameOk() (*bool, bool) {
	if o == nil || IsNil(o.JumboFrame) {
		return nil, false
	}
	return o.JumboFrame, true
}

// HasJumboFrame returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasJumboFrame() bool {
	if o != nil && !IsNil(o.JumboFrame) {
		return true
	}

	return false
}

// SetJumboFrame gets a reference to the given bool and assigns it to the JumboFrame field.
func (o *HyperflexClusterNetworkPolicy) SetJumboFrame(v bool) {
	o.JumboFrame = &v
}

// GetKvmIpRange returns the KvmIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetKvmIpRange() HyperflexIpAddrRange {
	if o == nil || IsNil(o.KvmIpRange.Get()) {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.KvmIpRange.Get()
}

// GetKvmIpRangeOk returns a tuple with the KvmIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.KvmIpRange.Get(), o.KvmIpRange.IsSet()
}

// HasKvmIpRange returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasKvmIpRange() bool {
	if o != nil && o.KvmIpRange.IsSet() {
		return true
	}

	return false
}

// SetKvmIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the KvmIpRange field.
func (o *HyperflexClusterNetworkPolicy) SetKvmIpRange(v HyperflexIpAddrRange) {
	o.KvmIpRange.Set(&v)
}

// SetKvmIpRangeNil sets the value for KvmIpRange to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetKvmIpRangeNil() {
	o.KvmIpRange.Set(nil)
}

// UnsetKvmIpRange ensures that no value is present for KvmIpRange, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetKvmIpRange() {
	o.KvmIpRange.Unset()
}

// GetMacPrefixRange returns the MacPrefixRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetMacPrefixRange() HyperflexMacAddrPrefixRange {
	if o == nil || IsNil(o.MacPrefixRange.Get()) {
		var ret HyperflexMacAddrPrefixRange
		return ret
	}
	return *o.MacPrefixRange.Get()
}

// GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacPrefixRange.Get(), o.MacPrefixRange.IsSet()
}

// HasMacPrefixRange returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasMacPrefixRange() bool {
	if o != nil && o.MacPrefixRange.IsSet() {
		return true
	}

	return false
}

// SetMacPrefixRange gets a reference to the given NullableHyperflexMacAddrPrefixRange and assigns it to the MacPrefixRange field.
func (o *HyperflexClusterNetworkPolicy) SetMacPrefixRange(v HyperflexMacAddrPrefixRange) {
	o.MacPrefixRange.Set(&v)
}

// SetMacPrefixRangeNil sets the value for MacPrefixRange to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetMacPrefixRangeNil() {
	o.MacPrefixRange.Set(nil)
}

// UnsetMacPrefixRange ensures that no value is present for MacPrefixRange, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetMacPrefixRange() {
	o.MacPrefixRange.Unset()
}

// GetMgmtVlan returns the MgmtVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetMgmtVlan() HyperflexNamedVlan {
	if o == nil || IsNil(o.MgmtVlan.Get()) {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.MgmtVlan.Get()
}

// GetMgmtVlanOk returns a tuple with the MgmtVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetMgmtVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.MgmtVlan.Get(), o.MgmtVlan.IsSet()
}

// HasMgmtVlan returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasMgmtVlan() bool {
	if o != nil && o.MgmtVlan.IsSet() {
		return true
	}

	return false
}

// SetMgmtVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the MgmtVlan field.
func (o *HyperflexClusterNetworkPolicy) SetMgmtVlan(v HyperflexNamedVlan) {
	o.MgmtVlan.Set(&v)
}

// SetMgmtVlanNil sets the value for MgmtVlan to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetMgmtVlanNil() {
	o.MgmtVlan.Set(nil)
}

// UnsetMgmtVlan ensures that no value is present for MgmtVlan, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetMgmtVlan() {
	o.MgmtVlan.Unset()
}

// GetUplinkSpeed returns the UplinkSpeed field value if set, zero value otherwise.
func (o *HyperflexClusterNetworkPolicy) GetUplinkSpeed() string {
	if o == nil || IsNil(o.UplinkSpeed) {
		var ret string
		return ret
	}
	return *o.UplinkSpeed
}

// GetUplinkSpeedOk returns a tuple with the UplinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetUplinkSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.UplinkSpeed) {
		return nil, false
	}
	return o.UplinkSpeed, true
}

// HasUplinkSpeed returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasUplinkSpeed() bool {
	if o != nil && !IsNil(o.UplinkSpeed) {
		return true
	}

	return false
}

// SetUplinkSpeed gets a reference to the given string and assigns it to the UplinkSpeed field.
func (o *HyperflexClusterNetworkPolicy) SetUplinkSpeed(v string) {
	o.UplinkSpeed = &v
}

// GetVmMigrationVlan returns the VmMigrationVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetVmMigrationVlan() HyperflexNamedVlan {
	if o == nil || IsNil(o.VmMigrationVlan.Get()) {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.VmMigrationVlan.Get()
}

// GetVmMigrationVlanOk returns a tuple with the VmMigrationVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetVmMigrationVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmMigrationVlan.Get(), o.VmMigrationVlan.IsSet()
}

// HasVmMigrationVlan returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasVmMigrationVlan() bool {
	if o != nil && o.VmMigrationVlan.IsSet() {
		return true
	}

	return false
}

// SetVmMigrationVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the VmMigrationVlan field.
func (o *HyperflexClusterNetworkPolicy) SetVmMigrationVlan(v HyperflexNamedVlan) {
	o.VmMigrationVlan.Set(&v)
}

// SetVmMigrationVlanNil sets the value for VmMigrationVlan to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetVmMigrationVlanNil() {
	o.VmMigrationVlan.Set(nil)
}

// UnsetVmMigrationVlan ensures that no value is present for VmMigrationVlan, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetVmMigrationVlan() {
	o.VmMigrationVlan.Unset()
}

// GetVmNetworkVlans returns the VmNetworkVlans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetVmNetworkVlans() []HyperflexNamedVlan {
	if o == nil {
		var ret []HyperflexNamedVlan
		return ret
	}
	return o.VmNetworkVlans
}

// GetVmNetworkVlansOk returns a tuple with the VmNetworkVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetVmNetworkVlansOk() ([]HyperflexNamedVlan, bool) {
	if o == nil || IsNil(o.VmNetworkVlans) {
		return nil, false
	}
	return o.VmNetworkVlans, true
}

// HasVmNetworkVlans returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasVmNetworkVlans() bool {
	if o != nil && !IsNil(o.VmNetworkVlans) {
		return true
	}

	return false
}

// SetVmNetworkVlans gets a reference to the given []HyperflexNamedVlan and assigns it to the VmNetworkVlans field.
func (o *HyperflexClusterNetworkPolicy) SetVmNetworkVlans(v []HyperflexNamedVlan) {
	o.VmNetworkVlans = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetClusterProfilesOk() ([]HyperflexClusterProfileRelationship, bool) {
	if o == nil || IsNil(o.ClusterProfiles) {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasClusterProfiles() bool {
	if o != nil && !IsNil(o.ClusterProfiles) {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterNetworkPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o HyperflexClusterNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexClusterNetworkPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CimcManagementMode) {
		toSerialize["CimcManagementMode"] = o.CimcManagementMode
	}
	if !IsNil(o.JumboFrame) {
		toSerialize["JumboFrame"] = o.JumboFrame
	}
	if o.KvmIpRange.IsSet() {
		toSerialize["KvmIpRange"] = o.KvmIpRange.Get()
	}
	if o.MacPrefixRange.IsSet() {
		toSerialize["MacPrefixRange"] = o.MacPrefixRange.Get()
	}
	if o.MgmtVlan.IsSet() {
		toSerialize["MgmtVlan"] = o.MgmtVlan.Get()
	}
	if !IsNil(o.UplinkSpeed) {
		toSerialize["UplinkSpeed"] = o.UplinkSpeed
	}
	if o.VmMigrationVlan.IsSet() {
		toSerialize["VmMigrationVlan"] = o.VmMigrationVlan.Get()
	}
	if o.VmNetworkVlans != nil {
		toSerialize["VmNetworkVlans"] = o.VmNetworkVlans
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

func (o *HyperflexClusterNetworkPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexClusterNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The mode configured on the CIMC management interface. * `OutOfBand` - The server uses out-of-band management, i.e. management traffic traverses through the management interfaces on the UCS Fabric Interconnects. * `InBand` - The server uses in-band management, i.e. management traffic traverses through the data uplink ports on the UCS Fabric Interconnects.
		CimcManagementMode *string `json:"CimcManagementMode,omitempty"`
		// Enable or disable Jumbo Frames (MTU=9000). Jumbo Frames are used by Storage Network, VM Migration Network.
		JumboFrame     *bool                               `json:"JumboFrame,omitempty"`
		KvmIpRange     NullableHyperflexIpAddrRange        `json:"KvmIpRange,omitempty"`
		MacPrefixRange NullableHyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty"`
		MgmtVlan       NullableHyperflexNamedVlan          `json:"MgmtVlan,omitempty"`
		// Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only. * `default` - Current default value set on the hardware platform. * `1G` - A link speed of 1 gigabit per second. * `10G` - A link speed of 10 gigabits per second or above.
		UplinkSpeed     *string                    `json:"UplinkSpeed,omitempty"`
		VmMigrationVlan NullableHyperflexNamedVlan `json:"VmMigrationVlan,omitempty"`
		VmNetworkVlans  []HyperflexNamedVlan       `json:"VmNetworkVlans,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
		Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct := HyperflexClusterNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterNetworkPolicy := _HyperflexClusterNetworkPolicy{}
		varHyperflexClusterNetworkPolicy.ClassId = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.ClassId
		varHyperflexClusterNetworkPolicy.ObjectType = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterNetworkPolicy.CimcManagementMode = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.CimcManagementMode
		varHyperflexClusterNetworkPolicy.JumboFrame = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.JumboFrame
		varHyperflexClusterNetworkPolicy.KvmIpRange = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.KvmIpRange
		varHyperflexClusterNetworkPolicy.MacPrefixRange = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.MacPrefixRange
		varHyperflexClusterNetworkPolicy.MgmtVlan = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.MgmtVlan
		varHyperflexClusterNetworkPolicy.UplinkSpeed = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.UplinkSpeed
		varHyperflexClusterNetworkPolicy.VmMigrationVlan = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.VmMigrationVlan
		varHyperflexClusterNetworkPolicy.VmNetworkVlans = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.VmNetworkVlans
		varHyperflexClusterNetworkPolicy.ClusterProfiles = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexClusterNetworkPolicy.Organization = varHyperflexClusterNetworkPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexClusterNetworkPolicy(varHyperflexClusterNetworkPolicy)
	} else {
		return err
	}

	varHyperflexClusterNetworkPolicy := _HyperflexClusterNetworkPolicy{}

	err = json.Unmarshal(data, &varHyperflexClusterNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexClusterNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CimcManagementMode")
		delete(additionalProperties, "JumboFrame")
		delete(additionalProperties, "KvmIpRange")
		delete(additionalProperties, "MacPrefixRange")
		delete(additionalProperties, "MgmtVlan")
		delete(additionalProperties, "UplinkSpeed")
		delete(additionalProperties, "VmMigrationVlan")
		delete(additionalProperties, "VmNetworkVlans")
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

type NullableHyperflexClusterNetworkPolicy struct {
	value *HyperflexClusterNetworkPolicy
	isSet bool
}

func (v NullableHyperflexClusterNetworkPolicy) Get() *HyperflexClusterNetworkPolicy {
	return v.value
}

func (v *NullableHyperflexClusterNetworkPolicy) Set(val *HyperflexClusterNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterNetworkPolicy(val *HyperflexClusterNetworkPolicy) *NullableHyperflexClusterNetworkPolicy {
	return &NullableHyperflexClusterNetworkPolicy{value: val, isSet: true}
}

func (v NullableHyperflexClusterNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
