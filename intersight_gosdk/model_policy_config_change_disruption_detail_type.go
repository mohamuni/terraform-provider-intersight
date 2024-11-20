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

// checks if the PolicyConfigChangeDisruptionDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyConfigChangeDisruptionDetailType{}

// PolicyConfigChangeDisruptionDetailType Details of the configuration change that includes type and description of the change and the type of disruption.
type PolicyConfigChangeDisruptionDetailType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType"`
	Disruptions []string `json:"Disruptions,omitempty"`
	// Name of the policy that, when modified, causes the disruption.
	PolicyName *string `json:"PolicyName,omitempty"`
	// Name of the action which is pending on this policy. Example, if policy is not yet activated we mark this field as not-activated. Currently we support two actions, not-deployed and not-activated.
	PolicyPendingAction  *string `json:"PolicyPendingAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigChangeDisruptionDetailType PolicyConfigChangeDisruptionDetailType

// NewPolicyConfigChangeDisruptionDetailType instantiates a new PolicyConfigChangeDisruptionDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigChangeDisruptionDetailType(classId string, objectType string) *PolicyConfigChangeDisruptionDetailType {
	this := PolicyConfigChangeDisruptionDetailType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyConfigChangeDisruptionDetailTypeWithDefaults instantiates a new PolicyConfigChangeDisruptionDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigChangeDisruptionDetailTypeWithDefaults() *PolicyConfigChangeDisruptionDetailType {
	this := PolicyConfigChangeDisruptionDetailType{}
	var classId string = "policy.ConfigChangeDisruptionDetailType"
	this.ClassId = classId
	var objectType string = "policy.ConfigChangeDisruptionDetailType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigChangeDisruptionDetailType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeDisruptionDetailType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigChangeDisruptionDetailType) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "policy.ConfigChangeDisruptionDetailType" of the ClassId field.
func (o *PolicyConfigChangeDisruptionDetailType) GetDefaultClassId() interface{} {
	return "policy.ConfigChangeDisruptionDetailType"
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigChangeDisruptionDetailType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeDisruptionDetailType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigChangeDisruptionDetailType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "policy.ConfigChangeDisruptionDetailType" of the ObjectType field.
func (o *PolicyConfigChangeDisruptionDetailType) GetDefaultObjectType() interface{} {
	return "policy.ConfigChangeDisruptionDetailType"
}

// GetDisruptions returns the Disruptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigChangeDisruptionDetailType) GetDisruptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Disruptions
}

// GetDisruptionsOk returns a tuple with the Disruptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigChangeDisruptionDetailType) GetDisruptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Disruptions) {
		return nil, false
	}
	return o.Disruptions, true
}

// HasDisruptions returns a boolean if a field has been set.
func (o *PolicyConfigChangeDisruptionDetailType) HasDisruptions() bool {
	if o != nil && !IsNil(o.Disruptions) {
		return true
	}

	return false
}

// SetDisruptions gets a reference to the given []string and assigns it to the Disruptions field.
func (o *PolicyConfigChangeDisruptionDetailType) SetDisruptions(v []string) {
	o.Disruptions = v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *PolicyConfigChangeDisruptionDetailType) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *PolicyConfigChangeDisruptionDetailType) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyPendingAction returns the PolicyPendingAction field value if set, zero value otherwise.
func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyPendingAction() string {
	if o == nil || IsNil(o.PolicyPendingAction) {
		var ret string
		return ret
	}
	return *o.PolicyPendingAction
}

// GetPolicyPendingActionOk returns a tuple with the PolicyPendingAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyPendingActionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyPendingAction) {
		return nil, false
	}
	return o.PolicyPendingAction, true
}

// HasPolicyPendingAction returns a boolean if a field has been set.
func (o *PolicyConfigChangeDisruptionDetailType) HasPolicyPendingAction() bool {
	if o != nil && !IsNil(o.PolicyPendingAction) {
		return true
	}

	return false
}

// SetPolicyPendingAction gets a reference to the given string and assigns it to the PolicyPendingAction field.
func (o *PolicyConfigChangeDisruptionDetailType) SetPolicyPendingAction(v string) {
	o.PolicyPendingAction = &v
}

func (o PolicyConfigChangeDisruptionDetailType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyConfigChangeDisruptionDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Disruptions != nil {
		toSerialize["Disruptions"] = o.Disruptions
	}
	if !IsNil(o.PolicyName) {
		toSerialize["PolicyName"] = o.PolicyName
	}
	if !IsNil(o.PolicyPendingAction) {
		toSerialize["PolicyPendingAction"] = o.PolicyPendingAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyConfigChangeDisruptionDetailType) UnmarshalJSON(data []byte) (err error) {
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
	type PolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string   `json:"ObjectType"`
		Disruptions []string `json:"Disruptions,omitempty"`
		// Name of the policy that, when modified, causes the disruption.
		PolicyName *string `json:"PolicyName,omitempty"`
		// Name of the action which is pending on this policy. Example, if policy is not yet activated we mark this field as not-activated. Currently we support two actions, not-deployed and not-activated.
		PolicyPendingAction *string `json:"PolicyPendingAction,omitempty"`
	}

	varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct := PolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct)
	if err == nil {
		varPolicyConfigChangeDisruptionDetailType := _PolicyConfigChangeDisruptionDetailType{}
		varPolicyConfigChangeDisruptionDetailType.ClassId = varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct.ClassId
		varPolicyConfigChangeDisruptionDetailType.ObjectType = varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct.ObjectType
		varPolicyConfigChangeDisruptionDetailType.Disruptions = varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct.Disruptions
		varPolicyConfigChangeDisruptionDetailType.PolicyName = varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct.PolicyName
		varPolicyConfigChangeDisruptionDetailType.PolicyPendingAction = varPolicyConfigChangeDisruptionDetailTypeWithoutEmbeddedStruct.PolicyPendingAction
		*o = PolicyConfigChangeDisruptionDetailType(varPolicyConfigChangeDisruptionDetailType)
	} else {
		return err
	}

	varPolicyConfigChangeDisruptionDetailType := _PolicyConfigChangeDisruptionDetailType{}

	err = json.Unmarshal(data, &varPolicyConfigChangeDisruptionDetailType)
	if err == nil {
		o.MoBaseComplexType = varPolicyConfigChangeDisruptionDetailType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Disruptions")
		delete(additionalProperties, "PolicyName")
		delete(additionalProperties, "PolicyPendingAction")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullablePolicyConfigChangeDisruptionDetailType struct {
	value *PolicyConfigChangeDisruptionDetailType
	isSet bool
}

func (v NullablePolicyConfigChangeDisruptionDetailType) Get() *PolicyConfigChangeDisruptionDetailType {
	return v.value
}

func (v *NullablePolicyConfigChangeDisruptionDetailType) Set(val *PolicyConfigChangeDisruptionDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigChangeDisruptionDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigChangeDisruptionDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigChangeDisruptionDetailType(val *PolicyConfigChangeDisruptionDetailType) *NullablePolicyConfigChangeDisruptionDetailType {
	return &NullablePolicyConfigChangeDisruptionDetailType{value: val, isSet: true}
}

func (v NullablePolicyConfigChangeDisruptionDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigChangeDisruptionDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
