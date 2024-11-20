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

// checks if the HyperflexProxySettingPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexProxySettingPolicy{}

// HyperflexProxySettingPolicy A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
type HyperflexProxySettingPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// HTTP Proxy server FQDN or IP.
	Hostname *string `json:"Hostname,omitempty" validate:"regexp=^[A-Za-z]([A-Za-z0-9-]*[A-Za-z0-9])?$|^[A-Za-z0-9]([A-Za-z0-9-]*[A-Za-z0-9])?(\\\\.[A-Za-z0-9]([A-Za-z0-9-]*[A-Za-z0-9])?)*(\\\\.[A-Za-z]([A-Za-z0-9-]*[A-Za-z0-9])?)$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password for the HTTP Proxy.
	Password *string `json:"Password,omitempty"`
	// The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive.
	Port *int64 `json:"Port,omitempty"`
	// The username for the HTTP Proxy.
	Username *string `json:"Username,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexProxySettingPolicy HyperflexProxySettingPolicy

// NewHyperflexProxySettingPolicy instantiates a new HyperflexProxySettingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexProxySettingPolicy(classId string, objectType string) *HyperflexProxySettingPolicy {
	this := HyperflexProxySettingPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexProxySettingPolicyWithDefaults instantiates a new HyperflexProxySettingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexProxySettingPolicyWithDefaults() *HyperflexProxySettingPolicy {
	this := HyperflexProxySettingPolicy{}
	var classId string = "hyperflex.ProxySettingPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ProxySettingPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexProxySettingPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexProxySettingPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ProxySettingPolicy" of the ClassId field.
func (o *HyperflexProxySettingPolicy) GetDefaultClassId() interface{} {
	return "hyperflex.ProxySettingPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexProxySettingPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexProxySettingPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ProxySettingPolicy" of the ObjectType field.
func (o *HyperflexProxySettingPolicy) GetDefaultObjectType() interface{} {
	return "hyperflex.ProxySettingPolicy"
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HyperflexProxySettingPolicy) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HyperflexProxySettingPolicy) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *HyperflexProxySettingPolicy) GetIsPasswordSet() bool {
	if o == nil || IsNil(o.IsPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasswordSet) {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasIsPasswordSet() bool {
	if o != nil && !IsNil(o.IsPasswordSet) {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *HyperflexProxySettingPolicy) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HyperflexProxySettingPolicy) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *HyperflexProxySettingPolicy) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *HyperflexProxySettingPolicy) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *HyperflexProxySettingPolicy) SetPort(v int64) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HyperflexProxySettingPolicy) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProxySettingPolicy) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HyperflexProxySettingPolicy) SetUsername(v string) {
	o.Username = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProxySettingPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProxySettingPolicy) GetClusterProfilesOk() ([]HyperflexClusterProfileRelationship, bool) {
	if o == nil || IsNil(o.ClusterProfiles) {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasClusterProfiles() bool {
	if o != nil && !IsNil(o.ClusterProfiles) {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexProxySettingPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProxySettingPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProxySettingPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexProxySettingPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexProxySettingPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *HyperflexProxySettingPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *HyperflexProxySettingPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o HyperflexProxySettingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexProxySettingPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Hostname) {
		toSerialize["Hostname"] = o.Hostname
	}
	if !IsNil(o.IsPasswordSet) {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if !IsNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	if !IsNil(o.Username) {
		toSerialize["Username"] = o.Username
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

func (o *HyperflexProxySettingPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexProxySettingPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// HTTP Proxy server FQDN or IP.
		Hostname *string `json:"Hostname,omitempty" validate:"regexp=^[A-Za-z]([A-Za-z0-9-]*[A-Za-z0-9])?$|^[A-Za-z0-9]([A-Za-z0-9-]*[A-Za-z0-9])?(\\\\.[A-Za-z0-9]([A-Za-z0-9-]*[A-Za-z0-9])?)*(\\\\.[A-Za-z]([A-Za-z0-9-]*[A-Za-z0-9])?)$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The password for the HTTP Proxy.
		Password *string `json:"Password,omitempty"`
		// The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive.
		Port *int64 `json:"Port,omitempty"`
		// The username for the HTTP Proxy.
		Username *string `json:"Username,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship        `json:"ClusterProfiles,omitempty"`
		Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexProxySettingPolicyWithoutEmbeddedStruct := HyperflexProxySettingPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexProxySettingPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexProxySettingPolicy := _HyperflexProxySettingPolicy{}
		varHyperflexProxySettingPolicy.ClassId = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.ClassId
		varHyperflexProxySettingPolicy.ObjectType = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexProxySettingPolicy.Hostname = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.Hostname
		varHyperflexProxySettingPolicy.IsPasswordSet = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.IsPasswordSet
		varHyperflexProxySettingPolicy.Password = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.Password
		varHyperflexProxySettingPolicy.Port = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.Port
		varHyperflexProxySettingPolicy.Username = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.Username
		varHyperflexProxySettingPolicy.ClusterProfiles = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexProxySettingPolicy.Organization = varHyperflexProxySettingPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexProxySettingPolicy(varHyperflexProxySettingPolicy)
	} else {
		return err
	}

	varHyperflexProxySettingPolicy := _HyperflexProxySettingPolicy{}

	err = json.Unmarshal(data, &varHyperflexProxySettingPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexProxySettingPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Username")
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

type NullableHyperflexProxySettingPolicy struct {
	value *HyperflexProxySettingPolicy
	isSet bool
}

func (v NullableHyperflexProxySettingPolicy) Get() *HyperflexProxySettingPolicy {
	return v.value
}

func (v *NullableHyperflexProxySettingPolicy) Set(val *HyperflexProxySettingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexProxySettingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexProxySettingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexProxySettingPolicy(val *HyperflexProxySettingPolicy) *NullableHyperflexProxySettingPolicy {
	return &NullableHyperflexProxySettingPolicy{value: val, isSet: true}
}

func (v NullableHyperflexProxySettingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexProxySettingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
