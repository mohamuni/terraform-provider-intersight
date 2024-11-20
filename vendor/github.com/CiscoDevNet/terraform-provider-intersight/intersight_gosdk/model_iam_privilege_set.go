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

// checks if the IamPrivilegeSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamPrivilegeSet{}

// IamPrivilegeSet A collection of privileges.
type IamPrivilegeSet struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the privilege set.
	Description *string `json:"Description,omitempty"`
	// Name of the privilege set.
	Name           *string                        `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_ .:-]{1,64}$"`
	PrivilegeNames []string                       `json:"PrivilegeNames,omitempty"`
	Account        NullableIamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	AssociatedPrivilegeSets []IamPrivilegeSetRelationship `json:"AssociatedPrivilegeSets,omitempty"`
	// An array of relationships to iamPrivilege resources.
	Privileges           []IamPrivilegeRelationship    `json:"Privileges,omitempty"`
	System               NullableIamSystemRelationship `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPrivilegeSet IamPrivilegeSet

// NewIamPrivilegeSet instantiates a new IamPrivilegeSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivilegeSet(classId string, objectType string) *IamPrivilegeSet {
	this := IamPrivilegeSet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPrivilegeSetWithDefaults instantiates a new IamPrivilegeSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivilegeSetWithDefaults() *IamPrivilegeSet {
	this := IamPrivilegeSet{}
	var classId string = "iam.PrivilegeSet"
	this.ClassId = classId
	var objectType string = "iam.PrivilegeSet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPrivilegeSet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPrivilegeSet) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.PrivilegeSet" of the ClassId field.
func (o *IamPrivilegeSet) GetDefaultClassId() interface{} {
	return "iam.PrivilegeSet"
}

// GetObjectType returns the ObjectType field value
func (o *IamPrivilegeSet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPrivilegeSet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.PrivilegeSet" of the ObjectType field.
func (o *IamPrivilegeSet) GetDefaultObjectType() interface{} {
	return "iam.PrivilegeSet"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamPrivilegeSet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamPrivilegeSet) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamPrivilegeSet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamPrivilegeSet) SetName(v string) {
	o.Name = &v
}

// GetPrivilegeNames returns the PrivilegeNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetPrivilegeNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PrivilegeNames
}

// GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetPrivilegeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.PrivilegeNames) {
		return nil, false
	}
	return o.PrivilegeNames, true
}

// HasPrivilegeNames returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasPrivilegeNames() bool {
	if o != nil && !IsNil(o.PrivilegeNames) {
		return true
	}

	return false
}

// SetPrivilegeNames gets a reference to the given []string and assigns it to the PrivilegeNames field.
func (o *IamPrivilegeSet) SetPrivilegeNames(v []string) {
	o.PrivilegeNames = v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IamPrivilegeSet) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IamPrivilegeSet) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IamPrivilegeSet) UnsetAccount() {
	o.Account.Unset()
}

// GetAssociatedPrivilegeSets returns the AssociatedPrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetAssociatedPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.AssociatedPrivilegeSets
}

// GetAssociatedPrivilegeSetsOk returns a tuple with the AssociatedPrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetAssociatedPrivilegeSetsOk() ([]IamPrivilegeSetRelationship, bool) {
	if o == nil || IsNil(o.AssociatedPrivilegeSets) {
		return nil, false
	}
	return o.AssociatedPrivilegeSets, true
}

// HasAssociatedPrivilegeSets returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasAssociatedPrivilegeSets() bool {
	if o != nil && !IsNil(o.AssociatedPrivilegeSets) {
		return true
	}

	return false
}

// SetAssociatedPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the AssociatedPrivilegeSets field.
func (o *IamPrivilegeSet) SetAssociatedPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.AssociatedPrivilegeSets = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetPrivileges() []IamPrivilegeRelationship {
	if o == nil {
		var ret []IamPrivilegeRelationship
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetPrivilegesOk() ([]IamPrivilegeRelationship, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []IamPrivilegeRelationship and assigns it to the Privileges field.
func (o *IamPrivilegeSet) SetPrivileges(v []IamPrivilegeRelationship) {
	o.Privileges = v
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetSystem() IamSystemRelationship {
	if o == nil || IsNil(o.System.Get()) {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System.Get()
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.System.Get(), o.System.IsSet()
}

// HasSystem returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasSystem() bool {
	if o != nil && o.System.IsSet() {
		return true
	}

	return false
}

// SetSystem gets a reference to the given NullableIamSystemRelationship and assigns it to the System field.
func (o *IamPrivilegeSet) SetSystem(v IamSystemRelationship) {
	o.System.Set(&v)
}

// SetSystemNil sets the value for System to be an explicit nil
func (o *IamPrivilegeSet) SetSystemNil() {
	o.System.Set(nil)
}

// UnsetSystem ensures that no value is present for System, not even an explicit nil
func (o *IamPrivilegeSet) UnsetSystem() {
	o.System.Unset()
}

func (o IamPrivilegeSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamPrivilegeSet) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.PrivilegeNames != nil {
		toSerialize["PrivilegeNames"] = o.PrivilegeNames
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.AssociatedPrivilegeSets != nil {
		toSerialize["AssociatedPrivilegeSets"] = o.AssociatedPrivilegeSets
	}
	if o.Privileges != nil {
		toSerialize["Privileges"] = o.Privileges
	}
	if o.System.IsSet() {
		toSerialize["System"] = o.System.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamPrivilegeSet) UnmarshalJSON(data []byte) (err error) {
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
	type IamPrivilegeSetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the privilege set.
		Description *string `json:"Description,omitempty"`
		// Name of the privilege set.
		Name           *string                        `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_ .:-]{1,64}$"`
		PrivilegeNames []string                       `json:"PrivilegeNames,omitempty"`
		Account        NullableIamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to iamPrivilegeSet resources.
		AssociatedPrivilegeSets []IamPrivilegeSetRelationship `json:"AssociatedPrivilegeSets,omitempty"`
		// An array of relationships to iamPrivilege resources.
		Privileges []IamPrivilegeRelationship    `json:"Privileges,omitempty"`
		System     NullableIamSystemRelationship `json:"System,omitempty"`
	}

	varIamPrivilegeSetWithoutEmbeddedStruct := IamPrivilegeSetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamPrivilegeSetWithoutEmbeddedStruct)
	if err == nil {
		varIamPrivilegeSet := _IamPrivilegeSet{}
		varIamPrivilegeSet.ClassId = varIamPrivilegeSetWithoutEmbeddedStruct.ClassId
		varIamPrivilegeSet.ObjectType = varIamPrivilegeSetWithoutEmbeddedStruct.ObjectType
		varIamPrivilegeSet.Description = varIamPrivilegeSetWithoutEmbeddedStruct.Description
		varIamPrivilegeSet.Name = varIamPrivilegeSetWithoutEmbeddedStruct.Name
		varIamPrivilegeSet.PrivilegeNames = varIamPrivilegeSetWithoutEmbeddedStruct.PrivilegeNames
		varIamPrivilegeSet.Account = varIamPrivilegeSetWithoutEmbeddedStruct.Account
		varIamPrivilegeSet.AssociatedPrivilegeSets = varIamPrivilegeSetWithoutEmbeddedStruct.AssociatedPrivilegeSets
		varIamPrivilegeSet.Privileges = varIamPrivilegeSetWithoutEmbeddedStruct.Privileges
		varIamPrivilegeSet.System = varIamPrivilegeSetWithoutEmbeddedStruct.System
		*o = IamPrivilegeSet(varIamPrivilegeSet)
	} else {
		return err
	}

	varIamPrivilegeSet := _IamPrivilegeSet{}

	err = json.Unmarshal(data, &varIamPrivilegeSet)
	if err == nil {
		o.MoBaseMo = varIamPrivilegeSet.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PrivilegeNames")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "AssociatedPrivilegeSets")
		delete(additionalProperties, "Privileges")
		delete(additionalProperties, "System")

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

type NullableIamPrivilegeSet struct {
	value *IamPrivilegeSet
	isSet bool
}

func (v NullableIamPrivilegeSet) Get() *IamPrivilegeSet {
	return v.value
}

func (v *NullableIamPrivilegeSet) Set(val *IamPrivilegeSet) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivilegeSet) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivilegeSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivilegeSet(val *IamPrivilegeSet) *NullableIamPrivilegeSet {
	return &NullableIamPrivilegeSet{value: val, isSet: true}
}

func (v NullableIamPrivilegeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivilegeSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
