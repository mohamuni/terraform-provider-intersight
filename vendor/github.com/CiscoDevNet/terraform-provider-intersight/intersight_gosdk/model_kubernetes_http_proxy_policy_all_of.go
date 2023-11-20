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

// KubernetesHttpProxyPolicyAllOf Definition of the list of properties defined in 'kubernetes.HttpProxyPolicy', excluding properties defined in parent classes.
type KubernetesHttpProxyPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	HttpProxy            NullableKubernetesProxyConfig         `json:"HttpProxy,omitempty"`
	HttpsProxy           NullableKubernetesProxyConfig         `json:"HttpsProxy,omitempty"`
	NoProxy              []string                              `json:"NoProxy,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesHttpProxyPolicyAllOf KubernetesHttpProxyPolicyAllOf

// NewKubernetesHttpProxyPolicyAllOf instantiates a new KubernetesHttpProxyPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesHttpProxyPolicyAllOf(classId string, objectType string) *KubernetesHttpProxyPolicyAllOf {
	this := KubernetesHttpProxyPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesHttpProxyPolicyAllOfWithDefaults instantiates a new KubernetesHttpProxyPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesHttpProxyPolicyAllOfWithDefaults() *KubernetesHttpProxyPolicyAllOf {
	this := KubernetesHttpProxyPolicyAllOf{}
	var classId string = "kubernetes.HttpProxyPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.HttpProxyPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesHttpProxyPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesHttpProxyPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesHttpProxyPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesHttpProxyPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesHttpProxyPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesHttpProxyPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHttpProxy returns the HttpProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesHttpProxyPolicyAllOf) GetHttpProxy() KubernetesProxyConfig {
	if o == nil || o.HttpProxy.Get() == nil {
		var ret KubernetesProxyConfig
		return ret
	}
	return *o.HttpProxy.Get()
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesHttpProxyPolicyAllOf) GetHttpProxyOk() (*KubernetesProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpProxy.Get(), o.HttpProxy.IsSet()
}

// HasHttpProxy returns a boolean if a field has been set.
func (o *KubernetesHttpProxyPolicyAllOf) HasHttpProxy() bool {
	if o != nil && o.HttpProxy.IsSet() {
		return true
	}

	return false
}

// SetHttpProxy gets a reference to the given NullableKubernetesProxyConfig and assigns it to the HttpProxy field.
func (o *KubernetesHttpProxyPolicyAllOf) SetHttpProxy(v KubernetesProxyConfig) {
	o.HttpProxy.Set(&v)
}

// SetHttpProxyNil sets the value for HttpProxy to be an explicit nil
func (o *KubernetesHttpProxyPolicyAllOf) SetHttpProxyNil() {
	o.HttpProxy.Set(nil)
}

// UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
func (o *KubernetesHttpProxyPolicyAllOf) UnsetHttpProxy() {
	o.HttpProxy.Unset()
}

// GetHttpsProxy returns the HttpsProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesHttpProxyPolicyAllOf) GetHttpsProxy() KubernetesProxyConfig {
	if o == nil || o.HttpsProxy.Get() == nil {
		var ret KubernetesProxyConfig
		return ret
	}
	return *o.HttpsProxy.Get()
}

// GetHttpsProxyOk returns a tuple with the HttpsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesHttpProxyPolicyAllOf) GetHttpsProxyOk() (*KubernetesProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpsProxy.Get(), o.HttpsProxy.IsSet()
}

// HasHttpsProxy returns a boolean if a field has been set.
func (o *KubernetesHttpProxyPolicyAllOf) HasHttpsProxy() bool {
	if o != nil && o.HttpsProxy.IsSet() {
		return true
	}

	return false
}

// SetHttpsProxy gets a reference to the given NullableKubernetesProxyConfig and assigns it to the HttpsProxy field.
func (o *KubernetesHttpProxyPolicyAllOf) SetHttpsProxy(v KubernetesProxyConfig) {
	o.HttpsProxy.Set(&v)
}

// SetHttpsProxyNil sets the value for HttpsProxy to be an explicit nil
func (o *KubernetesHttpProxyPolicyAllOf) SetHttpsProxyNil() {
	o.HttpsProxy.Set(nil)
}

// UnsetHttpsProxy ensures that no value is present for HttpsProxy, not even an explicit nil
func (o *KubernetesHttpProxyPolicyAllOf) UnsetHttpsProxy() {
	o.HttpsProxy.Unset()
}

// GetNoProxy returns the NoProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesHttpProxyPolicyAllOf) GetNoProxy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NoProxy
}

// GetNoProxyOk returns a tuple with the NoProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesHttpProxyPolicyAllOf) GetNoProxyOk() ([]string, bool) {
	if o == nil || o.NoProxy == nil {
		return nil, false
	}
	return o.NoProxy, true
}

// HasNoProxy returns a boolean if a field has been set.
func (o *KubernetesHttpProxyPolicyAllOf) HasNoProxy() bool {
	if o != nil && o.NoProxy != nil {
		return true
	}

	return false
}

// SetNoProxy gets a reference to the given []string and assigns it to the NoProxy field.
func (o *KubernetesHttpProxyPolicyAllOf) SetNoProxy(v []string) {
	o.NoProxy = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesHttpProxyPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesHttpProxyPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesHttpProxyPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesHttpProxyPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesHttpProxyPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HttpProxy.IsSet() {
		toSerialize["HttpProxy"] = o.HttpProxy.Get()
	}
	if o.HttpsProxy.IsSet() {
		toSerialize["HttpsProxy"] = o.HttpsProxy.Get()
	}
	if o.NoProxy != nil {
		toSerialize["NoProxy"] = o.NoProxy
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesHttpProxyPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesHttpProxyPolicyAllOf := _KubernetesHttpProxyPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesHttpProxyPolicyAllOf); err == nil {
		*o = KubernetesHttpProxyPolicyAllOf(varKubernetesHttpProxyPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HttpProxy")
		delete(additionalProperties, "HttpsProxy")
		delete(additionalProperties, "NoProxy")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesHttpProxyPolicyAllOf struct {
	value *KubernetesHttpProxyPolicyAllOf
	isSet bool
}

func (v NullableKubernetesHttpProxyPolicyAllOf) Get() *KubernetesHttpProxyPolicyAllOf {
	return v.value
}

func (v *NullableKubernetesHttpProxyPolicyAllOf) Set(val *KubernetesHttpProxyPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesHttpProxyPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesHttpProxyPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesHttpProxyPolicyAllOf(val *KubernetesHttpProxyPolicyAllOf) *NullableKubernetesHttpProxyPolicyAllOf {
	return &NullableKubernetesHttpProxyPolicyAllOf{value: val, isSet: true}
}

func (v NullableKubernetesHttpProxyPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesHttpProxyPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
