/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
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

// checks if the HyperflexExtFcStoragePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexExtFcStoragePolicy{}

// HyperflexExtFcStoragePolicy A policy specifying external storage connectivity information via Fabric attached FC storage.
type HyperflexExtFcStoragePolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables or disables external FC storage configuration.
	AdminState      *bool                            `json:"AdminState,omitempty"`
	ExtaTraffic     NullableHyperflexNamedVsan       `json:"ExtaTraffic,omitempty"`
	ExtbTraffic     NullableHyperflexNamedVsan       `json:"ExtbTraffic,omitempty"`
	WwxnPrefixRange NullableHyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexExtFcStoragePolicy HyperflexExtFcStoragePolicy

// NewHyperflexExtFcStoragePolicy instantiates a new HyperflexExtFcStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexExtFcStoragePolicy(classId string, objectType string) *HyperflexExtFcStoragePolicy {
	this := HyperflexExtFcStoragePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexExtFcStoragePolicyWithDefaults instantiates a new HyperflexExtFcStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexExtFcStoragePolicyWithDefaults() *HyperflexExtFcStoragePolicy {
	this := HyperflexExtFcStoragePolicy{}
	var classId string = "hyperflex.ExtFcStoragePolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ExtFcStoragePolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexExtFcStoragePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexExtFcStoragePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ExtFcStoragePolicy" of the ClassId field.
func (o *HyperflexExtFcStoragePolicy) GetDefaultClassId() interface{} {
	return "hyperflex.ExtFcStoragePolicy"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexExtFcStoragePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexExtFcStoragePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ExtFcStoragePolicy" of the ObjectType field.
func (o *HyperflexExtFcStoragePolicy) GetDefaultObjectType() interface{} {
	return "hyperflex.ExtFcStoragePolicy"
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicy) GetAdminState() bool {
	if o == nil || IsNil(o.AdminState) {
		var ret bool
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicy) GetAdminStateOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminState) {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicy) HasAdminState() bool {
	if o != nil && !IsNil(o.AdminState) {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given bool and assigns it to the AdminState field.
func (o *HyperflexExtFcStoragePolicy) SetAdminState(v bool) {
	o.AdminState = &v
}

// GetExtaTraffic returns the ExtaTraffic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicy) GetExtaTraffic() HyperflexNamedVsan {
	if o == nil || IsNil(o.ExtaTraffic.Get()) {
		var ret HyperflexNamedVsan
		return ret
	}
	return *o.ExtaTraffic.Get()
}

// GetExtaTrafficOk returns a tuple with the ExtaTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicy) GetExtaTrafficOk() (*HyperflexNamedVsan, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtaTraffic.Get(), o.ExtaTraffic.IsSet()
}

// HasExtaTraffic returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicy) HasExtaTraffic() bool {
	if o != nil && o.ExtaTraffic.IsSet() {
		return true
	}

	return false
}

// SetExtaTraffic gets a reference to the given NullableHyperflexNamedVsan and assigns it to the ExtaTraffic field.
func (o *HyperflexExtFcStoragePolicy) SetExtaTraffic(v HyperflexNamedVsan) {
	o.ExtaTraffic.Set(&v)
}

// SetExtaTrafficNil sets the value for ExtaTraffic to be an explicit nil
func (o *HyperflexExtFcStoragePolicy) SetExtaTrafficNil() {
	o.ExtaTraffic.Set(nil)
}

// UnsetExtaTraffic ensures that no value is present for ExtaTraffic, not even an explicit nil
func (o *HyperflexExtFcStoragePolicy) UnsetExtaTraffic() {
	o.ExtaTraffic.Unset()
}

// GetExtbTraffic returns the ExtbTraffic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicy) GetExtbTraffic() HyperflexNamedVsan {
	if o == nil || IsNil(o.ExtbTraffic.Get()) {
		var ret HyperflexNamedVsan
		return ret
	}
	return *o.ExtbTraffic.Get()
}

// GetExtbTrafficOk returns a tuple with the ExtbTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicy) GetExtbTrafficOk() (*HyperflexNamedVsan, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtbTraffic.Get(), o.ExtbTraffic.IsSet()
}

// HasExtbTraffic returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicy) HasExtbTraffic() bool {
	if o != nil && o.ExtbTraffic.IsSet() {
		return true
	}

	return false
}

// SetExtbTraffic gets a reference to the given NullableHyperflexNamedVsan and assigns it to the ExtbTraffic field.
func (o *HyperflexExtFcStoragePolicy) SetExtbTraffic(v HyperflexNamedVsan) {
	o.ExtbTraffic.Set(&v)
}

// SetExtbTrafficNil sets the value for ExtbTraffic to be an explicit nil
func (o *HyperflexExtFcStoragePolicy) SetExtbTrafficNil() {
	o.ExtbTraffic.Set(nil)
}

// UnsetExtbTraffic ensures that no value is present for ExtbTraffic, not even an explicit nil
func (o *HyperflexExtFcStoragePolicy) UnsetExtbTraffic() {
	o.ExtbTraffic.Unset()
}

// GetWwxnPrefixRange returns the WwxnPrefixRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicy) GetWwxnPrefixRange() HyperflexWwxnPrefixRange {
	if o == nil || IsNil(o.WwxnPrefixRange.Get()) {
		var ret HyperflexWwxnPrefixRange
		return ret
	}
	return *o.WwxnPrefixRange.Get()
}

// GetWwxnPrefixRangeOk returns a tuple with the WwxnPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicy) GetWwxnPrefixRangeOk() (*HyperflexWwxnPrefixRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.WwxnPrefixRange.Get(), o.WwxnPrefixRange.IsSet()
}

// HasWwxnPrefixRange returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicy) HasWwxnPrefixRange() bool {
	if o != nil && o.WwxnPrefixRange.IsSet() {
		return true
	}

	return false
}

// SetWwxnPrefixRange gets a reference to the given NullableHyperflexWwxnPrefixRange and assigns it to the WwxnPrefixRange field.
func (o *HyperflexExtFcStoragePolicy) SetWwxnPrefixRange(v HyperflexWwxnPrefixRange) {
	o.WwxnPrefixRange.Set(&v)
}

// SetWwxnPrefixRangeNil sets the value for WwxnPrefixRange to be an explicit nil
func (o *HyperflexExtFcStoragePolicy) SetWwxnPrefixRangeNil() {
	o.WwxnPrefixRange.Set(nil)
}

// UnsetWwxnPrefixRange ensures that no value is present for WwxnPrefixRange, not even an explicit nil
func (o *HyperflexExtFcStoragePolicy) UnsetWwxnPrefixRange() {
	o.WwxnPrefixRange.Unset()
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicy) GetClusterProfilesOk() ([]HyperflexClusterProfileRelationship, bool) {
	if o == nil || IsNil(o.ClusterProfiles) {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicy) HasClusterProfiles() bool {
	if o != nil && !IsNil(o.ClusterProfiles) {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexExtFcStoragePolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexExtFcStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *HyperflexExtFcStoragePolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *HyperflexExtFcStoragePolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o HyperflexExtFcStoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexExtFcStoragePolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdminState) {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ExtaTraffic.IsSet() {
		toSerialize["ExtaTraffic"] = o.ExtaTraffic.Get()
	}
	if o.ExtbTraffic.IsSet() {
		toSerialize["ExtbTraffic"] = o.ExtbTraffic.Get()
	}
	if o.WwxnPrefixRange.IsSet() {
		toSerialize["WwxnPrefixRange"] = o.WwxnPrefixRange.Get()
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

func (o *HyperflexExtFcStoragePolicy) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexExtFcStoragePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables or disables external FC storage configuration.
		AdminState      *bool                            `json:"AdminState,omitempty"`
		ExtaTraffic     NullableHyperflexNamedVsan       `json:"ExtaTraffic,omitempty"`
		ExtbTraffic     NullableHyperflexNamedVsan       `json:"ExtbTraffic,omitempty"`
		WwxnPrefixRange NullableHyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
		Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct := HyperflexExtFcStoragePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexExtFcStoragePolicy := _HyperflexExtFcStoragePolicy{}
		varHyperflexExtFcStoragePolicy.ClassId = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.ClassId
		varHyperflexExtFcStoragePolicy.ObjectType = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexExtFcStoragePolicy.AdminState = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.AdminState
		varHyperflexExtFcStoragePolicy.ExtaTraffic = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.ExtaTraffic
		varHyperflexExtFcStoragePolicy.ExtbTraffic = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.ExtbTraffic
		varHyperflexExtFcStoragePolicy.WwxnPrefixRange = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.WwxnPrefixRange
		varHyperflexExtFcStoragePolicy.ClusterProfiles = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexExtFcStoragePolicy.Organization = varHyperflexExtFcStoragePolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexExtFcStoragePolicy(varHyperflexExtFcStoragePolicy)
	} else {
		return err
	}

	varHyperflexExtFcStoragePolicy := _HyperflexExtFcStoragePolicy{}

	err = json.Unmarshal(data, &varHyperflexExtFcStoragePolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexExtFcStoragePolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ExtaTraffic")
		delete(additionalProperties, "ExtbTraffic")
		delete(additionalProperties, "WwxnPrefixRange")
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

type NullableHyperflexExtFcStoragePolicy struct {
	value *HyperflexExtFcStoragePolicy
	isSet bool
}

func (v NullableHyperflexExtFcStoragePolicy) Get() *HyperflexExtFcStoragePolicy {
	return v.value
}

func (v *NullableHyperflexExtFcStoragePolicy) Set(val *HyperflexExtFcStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexExtFcStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexExtFcStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexExtFcStoragePolicy(val *HyperflexExtFcStoragePolicy) *NullableHyperflexExtFcStoragePolicy {
	return &NullableHyperflexExtFcStoragePolicy{value: val, isSet: true}
}

func (v NullableHyperflexExtFcStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexExtFcStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
