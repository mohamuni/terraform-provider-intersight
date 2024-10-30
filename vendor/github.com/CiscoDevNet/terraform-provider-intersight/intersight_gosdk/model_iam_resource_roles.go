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

// checks if the IamResourceRoles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamResourceRoles{}

// IamResourceRoles ResourceRoles provides a way to specify the roles associated with a resource like organization in a permission which can be assigned to a user or user group.
type IamResourceRoles struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRoles []IamEndPointRoleRelationship     `json:"EndPointRoles,omitempty"`
	Permission    NullableIamPermissionRelationship `json:"Permission,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
	Resource      NullableMoBaseMoRelationship  `json:"Resource,omitempty"`
	// An array of relationships to iamRole resources.
	Roles                []IamRoleRelationship `json:"Roles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamResourceRoles IamResourceRoles

// NewIamResourceRoles instantiates a new IamResourceRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamResourceRoles(classId string, objectType string) *IamResourceRoles {
	this := IamResourceRoles{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamResourceRolesWithDefaults instantiates a new IamResourceRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamResourceRolesWithDefaults() *IamResourceRoles {
	this := IamResourceRoles{}
	var classId string = "iam.ResourceRoles"
	this.ClassId = classId
	var objectType string = "iam.ResourceRoles"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamResourceRoles) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamResourceRoles) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamResourceRoles) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.ResourceRoles" of the ClassId field.
func (o *IamResourceRoles) GetDefaultClassId() interface{} {
	return "iam.ResourceRoles"
}

// GetObjectType returns the ObjectType field value
func (o *IamResourceRoles) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamResourceRoles) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamResourceRoles) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.ResourceRoles" of the ObjectType field.
func (o *IamResourceRoles) GetDefaultObjectType() interface{} {
	return "iam.ResourceRoles"
}

// GetEndPointRoles returns the EndPointRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourceRoles) GetEndPointRoles() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRoles
}

// GetEndPointRolesOk returns a tuple with the EndPointRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourceRoles) GetEndPointRolesOk() ([]IamEndPointRoleRelationship, bool) {
	if o == nil || IsNil(o.EndPointRoles) {
		return nil, false
	}
	return o.EndPointRoles, true
}

// HasEndPointRoles returns a boolean if a field has been set.
func (o *IamResourceRoles) HasEndPointRoles() bool {
	if o != nil && !IsNil(o.EndPointRoles) {
		return true
	}

	return false
}

// SetEndPointRoles gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRoles field.
func (o *IamResourceRoles) SetEndPointRoles(v []IamEndPointRoleRelationship) {
	o.EndPointRoles = v
}

// GetPermission returns the Permission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourceRoles) GetPermission() IamPermissionRelationship {
	if o == nil || IsNil(o.Permission.Get()) {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourceRoles) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// HasPermission returns a boolean if a field has been set.
func (o *IamResourceRoles) HasPermission() bool {
	if o != nil && o.Permission.IsSet() {
		return true
	}

	return false
}

// SetPermission gets a reference to the given NullableIamPermissionRelationship and assigns it to the Permission field.
func (o *IamResourceRoles) SetPermission(v IamPermissionRelationship) {
	o.Permission.Set(&v)
}

// SetPermissionNil sets the value for Permission to be an explicit nil
func (o *IamResourceRoles) SetPermissionNil() {
	o.Permission.Set(nil)
}

// UnsetPermission ensures that no value is present for Permission, not even an explicit nil
func (o *IamResourceRoles) UnsetPermission() {
	o.Permission.Unset()
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourceRoles) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourceRoles) GetPrivilegeSetsOk() ([]IamPrivilegeSetRelationship, bool) {
	if o == nil || IsNil(o.PrivilegeSets) {
		return nil, false
	}
	return o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamResourceRoles) HasPrivilegeSets() bool {
	if o != nil && !IsNil(o.PrivilegeSets) {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamResourceRoles) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourceRoles) GetResource() MoBaseMoRelationship {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourceRoles) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *IamResourceRoles) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableMoBaseMoRelationship and assigns it to the Resource field.
func (o *IamResourceRoles) SetResource(v MoBaseMoRelationship) {
	o.Resource.Set(&v)
}

// SetResourceNil sets the value for Resource to be an explicit nil
func (o *IamResourceRoles) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *IamResourceRoles) UnsetResource() {
	o.Resource.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourceRoles) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourceRoles) GetRolesOk() ([]IamRoleRelationship, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamResourceRoles) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamResourceRoles) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

func (o IamResourceRoles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamResourceRoles) ToMap() (map[string]interface{}, error) {
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
	if o.EndPointRoles != nil {
		toSerialize["EndPointRoles"] = o.EndPointRoles
	}
	if o.Permission.IsSet() {
		toSerialize["Permission"] = o.Permission.Get()
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.Resource.IsSet() {
		toSerialize["Resource"] = o.Resource.Get()
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamResourceRoles) UnmarshalJSON(data []byte) (err error) {
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
	type IamResourceRolesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to iamEndPointRole resources.
		EndPointRoles []IamEndPointRoleRelationship     `json:"EndPointRoles,omitempty"`
		Permission    NullableIamPermissionRelationship `json:"Permission,omitempty"`
		// An array of relationships to iamPrivilegeSet resources.
		PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
		Resource      NullableMoBaseMoRelationship  `json:"Resource,omitempty"`
		// An array of relationships to iamRole resources.
		Roles []IamRoleRelationship `json:"Roles,omitempty"`
	}

	varIamResourceRolesWithoutEmbeddedStruct := IamResourceRolesWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamResourceRolesWithoutEmbeddedStruct)
	if err == nil {
		varIamResourceRoles := _IamResourceRoles{}
		varIamResourceRoles.ClassId = varIamResourceRolesWithoutEmbeddedStruct.ClassId
		varIamResourceRoles.ObjectType = varIamResourceRolesWithoutEmbeddedStruct.ObjectType
		varIamResourceRoles.EndPointRoles = varIamResourceRolesWithoutEmbeddedStruct.EndPointRoles
		varIamResourceRoles.Permission = varIamResourceRolesWithoutEmbeddedStruct.Permission
		varIamResourceRoles.PrivilegeSets = varIamResourceRolesWithoutEmbeddedStruct.PrivilegeSets
		varIamResourceRoles.Resource = varIamResourceRolesWithoutEmbeddedStruct.Resource
		varIamResourceRoles.Roles = varIamResourceRolesWithoutEmbeddedStruct.Roles
		*o = IamResourceRoles(varIamResourceRoles)
	} else {
		return err
	}

	varIamResourceRoles := _IamResourceRoles{}

	err = json.Unmarshal(data, &varIamResourceRoles)
	if err == nil {
		o.MoBaseMo = varIamResourceRoles.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndPointRoles")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "PrivilegeSets")
		delete(additionalProperties, "Resource")
		delete(additionalProperties, "Roles")

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

type NullableIamResourceRoles struct {
	value *IamResourceRoles
	isSet bool
}

func (v NullableIamResourceRoles) Get() *IamResourceRoles {
	return v.value
}

func (v *NullableIamResourceRoles) Set(val *IamResourceRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourceRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourceRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourceRoles(val *IamResourceRoles) *NullableIamResourceRoles {
	return &NullableIamResourceRoles{value: val, isSet: true}
}

func (v NullableIamResourceRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourceRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
