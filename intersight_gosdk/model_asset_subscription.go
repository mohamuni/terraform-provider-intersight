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

// checks if the AssetSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetSubscription{}

// AssetSubscription Contains information about consumption-based Subscriptions related to the Cisco devices associated. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
type AssetSubscription struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Application name reported by Cisco Install Base.
	ApplicationName *string `json:"ApplicationName,omitempty"`
	// Identifies the consumption-based subscription.
	SubscriptionRefId *string `json:"SubscriptionRefId,omitempty"`
	// An array of relationships to assetDeployment resources.
	Deployments          []AssetDeploymentRelationship                `json:"Deployments,omitempty"`
	SubscriptionAccount  NullableAssetSubscriptionAccountRelationship `json:"SubscriptionAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSubscription AssetSubscription

// NewAssetSubscription instantiates a new AssetSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSubscription(classId string, objectType string) *AssetSubscription {
	this := AssetSubscription{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetSubscriptionWithDefaults instantiates a new AssetSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSubscriptionWithDefaults() *AssetSubscription {
	this := AssetSubscription{}
	var classId string = "asset.Subscription"
	this.ClassId = classId
	var objectType string = "asset.Subscription"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSubscription) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSubscription) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSubscription) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.Subscription" of the ClassId field.
func (o *AssetSubscription) GetDefaultClassId() interface{} {
	return "asset.Subscription"
}

// GetObjectType returns the ObjectType field value
func (o *AssetSubscription) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSubscription) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSubscription) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.Subscription" of the ObjectType field.
func (o *AssetSubscription) GetDefaultObjectType() interface{} {
	return "asset.Subscription"
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AssetSubscription) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscription) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AssetSubscription) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AssetSubscription) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetSubscriptionRefId returns the SubscriptionRefId field value if set, zero value otherwise.
func (o *AssetSubscription) GetSubscriptionRefId() string {
	if o == nil || IsNil(o.SubscriptionRefId) {
		var ret string
		return ret
	}
	return *o.SubscriptionRefId
}

// GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscription) GetSubscriptionRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionRefId) {
		return nil, false
	}
	return o.SubscriptionRefId, true
}

// HasSubscriptionRefId returns a boolean if a field has been set.
func (o *AssetSubscription) HasSubscriptionRefId() bool {
	if o != nil && !IsNil(o.SubscriptionRefId) {
		return true
	}

	return false
}

// SetSubscriptionRefId gets a reference to the given string and assigns it to the SubscriptionRefId field.
func (o *AssetSubscription) SetSubscriptionRefId(v string) {
	o.SubscriptionRefId = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSubscription) GetDeployments() []AssetDeploymentRelationship {
	if o == nil {
		var ret []AssetDeploymentRelationship
		return ret
	}
	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSubscription) GetDeploymentsOk() ([]AssetDeploymentRelationship, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *AssetSubscription) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []AssetDeploymentRelationship and assigns it to the Deployments field.
func (o *AssetSubscription) SetDeployments(v []AssetDeploymentRelationship) {
	o.Deployments = v
}

// GetSubscriptionAccount returns the SubscriptionAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSubscription) GetSubscriptionAccount() AssetSubscriptionAccountRelationship {
	if o == nil || IsNil(o.SubscriptionAccount.Get()) {
		var ret AssetSubscriptionAccountRelationship
		return ret
	}
	return *o.SubscriptionAccount.Get()
}

// GetSubscriptionAccountOk returns a tuple with the SubscriptionAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSubscription) GetSubscriptionAccountOk() (*AssetSubscriptionAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionAccount.Get(), o.SubscriptionAccount.IsSet()
}

// HasSubscriptionAccount returns a boolean if a field has been set.
func (o *AssetSubscription) HasSubscriptionAccount() bool {
	if o != nil && o.SubscriptionAccount.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionAccount gets a reference to the given NullableAssetSubscriptionAccountRelationship and assigns it to the SubscriptionAccount field.
func (o *AssetSubscription) SetSubscriptionAccount(v AssetSubscriptionAccountRelationship) {
	o.SubscriptionAccount.Set(&v)
}

// SetSubscriptionAccountNil sets the value for SubscriptionAccount to be an explicit nil
func (o *AssetSubscription) SetSubscriptionAccountNil() {
	o.SubscriptionAccount.Set(nil)
}

// UnsetSubscriptionAccount ensures that no value is present for SubscriptionAccount, not even an explicit nil
func (o *AssetSubscription) UnsetSubscriptionAccount() {
	o.SubscriptionAccount.Unset()
}

func (o AssetSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ApplicationName) {
		toSerialize["ApplicationName"] = o.ApplicationName
	}
	if !IsNil(o.SubscriptionRefId) {
		toSerialize["SubscriptionRefId"] = o.SubscriptionRefId
	}
	if o.Deployments != nil {
		toSerialize["Deployments"] = o.Deployments
	}
	if o.SubscriptionAccount.IsSet() {
		toSerialize["SubscriptionAccount"] = o.SubscriptionAccount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetSubscription) UnmarshalJSON(data []byte) (err error) {
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
	type AssetSubscriptionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Application name reported by Cisco Install Base.
		ApplicationName *string `json:"ApplicationName,omitempty"`
		// Identifies the consumption-based subscription.
		SubscriptionRefId *string `json:"SubscriptionRefId,omitempty"`
		// An array of relationships to assetDeployment resources.
		Deployments         []AssetDeploymentRelationship                `json:"Deployments,omitempty"`
		SubscriptionAccount NullableAssetSubscriptionAccountRelationship `json:"SubscriptionAccount,omitempty"`
	}

	varAssetSubscriptionWithoutEmbeddedStruct := AssetSubscriptionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetSubscriptionWithoutEmbeddedStruct)
	if err == nil {
		varAssetSubscription := _AssetSubscription{}
		varAssetSubscription.ClassId = varAssetSubscriptionWithoutEmbeddedStruct.ClassId
		varAssetSubscription.ObjectType = varAssetSubscriptionWithoutEmbeddedStruct.ObjectType
		varAssetSubscription.ApplicationName = varAssetSubscriptionWithoutEmbeddedStruct.ApplicationName
		varAssetSubscription.SubscriptionRefId = varAssetSubscriptionWithoutEmbeddedStruct.SubscriptionRefId
		varAssetSubscription.Deployments = varAssetSubscriptionWithoutEmbeddedStruct.Deployments
		varAssetSubscription.SubscriptionAccount = varAssetSubscriptionWithoutEmbeddedStruct.SubscriptionAccount
		*o = AssetSubscription(varAssetSubscription)
	} else {
		return err
	}

	varAssetSubscription := _AssetSubscription{}

	err = json.Unmarshal(data, &varAssetSubscription)
	if err == nil {
		o.MoBaseMo = varAssetSubscription.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApplicationName")
		delete(additionalProperties, "SubscriptionRefId")
		delete(additionalProperties, "Deployments")
		delete(additionalProperties, "SubscriptionAccount")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableAssetSubscription struct {
	value *AssetSubscription
	isSet bool
}

func (v NullableAssetSubscription) Get() *AssetSubscription {
	return v.value
}

func (v *NullableAssetSubscription) Set(val *AssetSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSubscription(val *AssetSubscription) *NullableAssetSubscription {
	return &NullableAssetSubscription{value: val, isSet: true}
}

func (v NullableAssetSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
