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

// checks if the BulkMoDeepCloner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkMoDeepCloner{}

// BulkMoDeepCloner The MODeepCloner interface facilitates making n number of deep copies of any resource instance which supports the CREATE operation. The MO to be cloned should be specified as an MoRef object in the \"Source\". The \"Targets\" array should contain n JSON documents each compliant to RFC 7386.  For each target MO to be created, the user can specify the following - - new values for the identity properties, if applicable - new values for specific properties or references of the source MO which need to be overridden in the cloned object.
type BulkMoDeepCloner struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string   `json:"ObjectType"`
	ExcludeProperties []string `json:"ExcludeProperties,omitempty"`
	// Name suffix to be applied to all the MOs being cloned when ReferencePolicy chosen is CreateNew. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_).
	ReferenceNameSuffix *string `json:"ReferenceNameSuffix,omitempty" validate:"regexp=^$|^[a-zA-Z0-9_-]{1,64}$"`
	// User selected reference clone behavior. Applies to all the MOs being cloned. * `ReuseAll` - Any policies in the destination organization whose name matches the policy referenced in the cloned policy will be attached. If no policyin the destination organization matches by name, a policy will be cloned with the same name.Pool references will always be matched by name. If not found, the pool will be cloned in the destination organization, but no identifierblocks will be created. * `CreateNew` - New policies will be created for the source and all the attached policies. If a policy of the same name and type already exists in thedestination organization or any organization from which it shares policies, a clone will be created with the provided suffix added to the name.Pool references will always be matched by name. If not found, the pool will be cloned in the destination organization, but no identifierblocks will be created.
	ReferencePolicy *string    `json:"ReferencePolicy,omitempty"`
	Source          *MoMoRef   `json:"Source,omitempty"`
	Targets         []MoBaseMo `json:"Targets,omitempty"`
	// A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_).
	WorkflowNameSuffix   *string                                      `json:"WorkflowNameSuffix,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]{1}[\\\\sa-zA-Z0-9_.\\\\,\\/:-]{0,63}$"`
	AsyncResult          NullableBulkResultRelationship               `json:"AsyncResult,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkMoDeepCloner BulkMoDeepCloner

// NewBulkMoDeepCloner instantiates a new BulkMoDeepCloner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMoDeepCloner(classId string, objectType string) *BulkMoDeepCloner {
	this := BulkMoDeepCloner{}
	this.ClassId = classId
	this.ObjectType = objectType
	var referencePolicy string = "ReuseAll"
	this.ReferencePolicy = &referencePolicy
	return &this
}

// NewBulkMoDeepClonerWithDefaults instantiates a new BulkMoDeepCloner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMoDeepClonerWithDefaults() *BulkMoDeepCloner {
	this := BulkMoDeepCloner{}
	var classId string = "bulk.MoDeepCloner"
	this.ClassId = classId
	var objectType string = "bulk.MoDeepCloner"
	this.ObjectType = objectType
	var referencePolicy string = "ReuseAll"
	this.ReferencePolicy = &referencePolicy
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkMoDeepCloner) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkMoDeepCloner) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.MoDeepCloner" of the ClassId field.
func (o *BulkMoDeepCloner) GetDefaultClassId() interface{} {
	return "bulk.MoDeepCloner"
}

// GetObjectType returns the ObjectType field value
func (o *BulkMoDeepCloner) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkMoDeepCloner) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.MoDeepCloner" of the ObjectType field.
func (o *BulkMoDeepCloner) GetDefaultObjectType() interface{} {
	return "bulk.MoDeepCloner"
}

// GetExcludeProperties returns the ExcludeProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoDeepCloner) GetExcludeProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeProperties
}

// GetExcludePropertiesOk returns a tuple with the ExcludeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoDeepCloner) GetExcludePropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeProperties) {
		return nil, false
	}
	return o.ExcludeProperties, true
}

// HasExcludeProperties returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasExcludeProperties() bool {
	if o != nil && !IsNil(o.ExcludeProperties) {
		return true
	}

	return false
}

// SetExcludeProperties gets a reference to the given []string and assigns it to the ExcludeProperties field.
func (o *BulkMoDeepCloner) SetExcludeProperties(v []string) {
	o.ExcludeProperties = v
}

// GetReferenceNameSuffix returns the ReferenceNameSuffix field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetReferenceNameSuffix() string {
	if o == nil || IsNil(o.ReferenceNameSuffix) {
		var ret string
		return ret
	}
	return *o.ReferenceNameSuffix
}

// GetReferenceNameSuffixOk returns a tuple with the ReferenceNameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetReferenceNameSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNameSuffix) {
		return nil, false
	}
	return o.ReferenceNameSuffix, true
}

// HasReferenceNameSuffix returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasReferenceNameSuffix() bool {
	if o != nil && !IsNil(o.ReferenceNameSuffix) {
		return true
	}

	return false
}

// SetReferenceNameSuffix gets a reference to the given string and assigns it to the ReferenceNameSuffix field.
func (o *BulkMoDeepCloner) SetReferenceNameSuffix(v string) {
	o.ReferenceNameSuffix = &v
}

// GetReferencePolicy returns the ReferencePolicy field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetReferencePolicy() string {
	if o == nil || IsNil(o.ReferencePolicy) {
		var ret string
		return ret
	}
	return *o.ReferencePolicy
}

// GetReferencePolicyOk returns a tuple with the ReferencePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetReferencePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ReferencePolicy) {
		return nil, false
	}
	return o.ReferencePolicy, true
}

// HasReferencePolicy returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasReferencePolicy() bool {
	if o != nil && !IsNil(o.ReferencePolicy) {
		return true
	}

	return false
}

// SetReferencePolicy gets a reference to the given string and assigns it to the ReferencePolicy field.
func (o *BulkMoDeepCloner) SetReferencePolicy(v string) {
	o.ReferencePolicy = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetSource() MoMoRef {
	if o == nil || IsNil(o.Source) {
		var ret MoMoRef
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetSourceOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given MoMoRef and assigns it to the Source field.
func (o *BulkMoDeepCloner) SetSource(v MoMoRef) {
	o.Source = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoDeepCloner) GetTargets() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoDeepCloner) GetTargetsOk() ([]MoBaseMo, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MoBaseMo and assigns it to the Targets field.
func (o *BulkMoDeepCloner) SetTargets(v []MoBaseMo) {
	o.Targets = v
}

// GetWorkflowNameSuffix returns the WorkflowNameSuffix field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetWorkflowNameSuffix() string {
	if o == nil || IsNil(o.WorkflowNameSuffix) {
		var ret string
		return ret
	}
	return *o.WorkflowNameSuffix
}

// GetWorkflowNameSuffixOk returns a tuple with the WorkflowNameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetWorkflowNameSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowNameSuffix) {
		return nil, false
	}
	return o.WorkflowNameSuffix, true
}

// HasWorkflowNameSuffix returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasWorkflowNameSuffix() bool {
	if o != nil && !IsNil(o.WorkflowNameSuffix) {
		return true
	}

	return false
}

// SetWorkflowNameSuffix gets a reference to the given string and assigns it to the WorkflowNameSuffix field.
func (o *BulkMoDeepCloner) SetWorkflowNameSuffix(v string) {
	o.WorkflowNameSuffix = &v
}

// GetAsyncResult returns the AsyncResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoDeepCloner) GetAsyncResult() BulkResultRelationship {
	if o == nil || IsNil(o.AsyncResult.Get()) {
		var ret BulkResultRelationship
		return ret
	}
	return *o.AsyncResult.Get()
}

// GetAsyncResultOk returns a tuple with the AsyncResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoDeepCloner) GetAsyncResultOk() (*BulkResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsyncResult.Get(), o.AsyncResult.IsSet()
}

// HasAsyncResult returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasAsyncResult() bool {
	if o != nil && o.AsyncResult.IsSet() {
		return true
	}

	return false
}

// SetAsyncResult gets a reference to the given NullableBulkResultRelationship and assigns it to the AsyncResult field.
func (o *BulkMoDeepCloner) SetAsyncResult(v BulkResultRelationship) {
	o.AsyncResult.Set(&v)
}

// SetAsyncResultNil sets the value for AsyncResult to be an explicit nil
func (o *BulkMoDeepCloner) SetAsyncResultNil() {
	o.AsyncResult.Set(nil)
}

// UnsetAsyncResult ensures that no value is present for AsyncResult, not even an explicit nil
func (o *BulkMoDeepCloner) UnsetAsyncResult() {
	o.AsyncResult.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoDeepCloner) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoDeepCloner) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkMoDeepCloner) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *BulkMoDeepCloner) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *BulkMoDeepCloner) UnsetOrganization() {
	o.Organization.Unset()
}

func (o BulkMoDeepCloner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkMoDeepCloner) ToMap() (map[string]interface{}, error) {
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
	if o.ExcludeProperties != nil {
		toSerialize["ExcludeProperties"] = o.ExcludeProperties
	}
	if !IsNil(o.ReferenceNameSuffix) {
		toSerialize["ReferenceNameSuffix"] = o.ReferenceNameSuffix
	}
	if !IsNil(o.ReferencePolicy) {
		toSerialize["ReferencePolicy"] = o.ReferencePolicy
	}
	if !IsNil(o.Source) {
		toSerialize["Source"] = o.Source
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if !IsNil(o.WorkflowNameSuffix) {
		toSerialize["WorkflowNameSuffix"] = o.WorkflowNameSuffix
	}
	if o.AsyncResult.IsSet() {
		toSerialize["AsyncResult"] = o.AsyncResult.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkMoDeepCloner) UnmarshalJSON(data []byte) (err error) {
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
	type BulkMoDeepClonerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string   `json:"ObjectType"`
		ExcludeProperties []string `json:"ExcludeProperties,omitempty"`
		// Name suffix to be applied to all the MOs being cloned when ReferencePolicy chosen is CreateNew. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_).
		ReferenceNameSuffix *string `json:"ReferenceNameSuffix,omitempty" validate:"regexp=^$|^[a-zA-Z0-9_-]{1,64}$"`
		// User selected reference clone behavior. Applies to all the MOs being cloned. * `ReuseAll` - Any policies in the destination organization whose name matches the policy referenced in the cloned policy will be attached. If no policyin the destination organization matches by name, a policy will be cloned with the same name.Pool references will always be matched by name. If not found, the pool will be cloned in the destination organization, but no identifierblocks will be created. * `CreateNew` - New policies will be created for the source and all the attached policies. If a policy of the same name and type already exists in thedestination organization or any organization from which it shares policies, a clone will be created with the provided suffix added to the name.Pool references will always be matched by name. If not found, the pool will be cloned in the destination organization, but no identifierblocks will be created.
		ReferencePolicy *string    `json:"ReferencePolicy,omitempty"`
		Source          *MoMoRef   `json:"Source,omitempty"`
		Targets         []MoBaseMo `json:"Targets,omitempty"`
		// A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_).
		WorkflowNameSuffix *string                                      `json:"WorkflowNameSuffix,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]{1}[\\\\sa-zA-Z0-9_.\\\\,\\/:-]{0,63}$"`
		AsyncResult        NullableBulkResultRelationship               `json:"AsyncResult,omitempty"`
		Organization       NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varBulkMoDeepClonerWithoutEmbeddedStruct := BulkMoDeepClonerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkMoDeepClonerWithoutEmbeddedStruct)
	if err == nil {
		varBulkMoDeepCloner := _BulkMoDeepCloner{}
		varBulkMoDeepCloner.ClassId = varBulkMoDeepClonerWithoutEmbeddedStruct.ClassId
		varBulkMoDeepCloner.ObjectType = varBulkMoDeepClonerWithoutEmbeddedStruct.ObjectType
		varBulkMoDeepCloner.ExcludeProperties = varBulkMoDeepClonerWithoutEmbeddedStruct.ExcludeProperties
		varBulkMoDeepCloner.ReferenceNameSuffix = varBulkMoDeepClonerWithoutEmbeddedStruct.ReferenceNameSuffix
		varBulkMoDeepCloner.ReferencePolicy = varBulkMoDeepClonerWithoutEmbeddedStruct.ReferencePolicy
		varBulkMoDeepCloner.Source = varBulkMoDeepClonerWithoutEmbeddedStruct.Source
		varBulkMoDeepCloner.Targets = varBulkMoDeepClonerWithoutEmbeddedStruct.Targets
		varBulkMoDeepCloner.WorkflowNameSuffix = varBulkMoDeepClonerWithoutEmbeddedStruct.WorkflowNameSuffix
		varBulkMoDeepCloner.AsyncResult = varBulkMoDeepClonerWithoutEmbeddedStruct.AsyncResult
		varBulkMoDeepCloner.Organization = varBulkMoDeepClonerWithoutEmbeddedStruct.Organization
		*o = BulkMoDeepCloner(varBulkMoDeepCloner)
	} else {
		return err
	}

	varBulkMoDeepCloner := _BulkMoDeepCloner{}

	err = json.Unmarshal(data, &varBulkMoDeepCloner)
	if err == nil {
		o.MoBaseMo = varBulkMoDeepCloner.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeProperties")
		delete(additionalProperties, "ReferenceNameSuffix")
		delete(additionalProperties, "ReferencePolicy")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "WorkflowNameSuffix")
		delete(additionalProperties, "AsyncResult")
		delete(additionalProperties, "Organization")

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

type NullableBulkMoDeepCloner struct {
	value *BulkMoDeepCloner
	isSet bool
}

func (v NullableBulkMoDeepCloner) Get() *BulkMoDeepCloner {
	return v.value
}

func (v *NullableBulkMoDeepCloner) Set(val *BulkMoDeepCloner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMoDeepCloner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMoDeepCloner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMoDeepCloner(val *BulkMoDeepCloner) *NullableBulkMoDeepCloner {
	return &NullableBulkMoDeepCloner{value: val, isSet: true}
}

func (v NullableBulkMoDeepCloner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMoDeepCloner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
