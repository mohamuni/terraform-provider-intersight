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

// checks if the StorageNetAppExportPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppExportPolicy{}

// StorageNetAppExportPolicy NetApp export policies enable client access control to volumes. Clients that match specific IP addresses and/or specific authentication types are granted access.
type StorageNetAppExportPolicy struct {
	StorageBaseNfsExport
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identity of the device.
	ClusterUuid            *string                         `json:"ClusterUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	NetAppExportPolicyRule []StorageNetAppExportPolicyRule `json:"NetAppExportPolicyRule,omitempty"`
	// ID for the Export Policy.
	PolicyId *int64 `json:"PolicyId,omitempty"`
	// The storage virtual machine name for the export policy.
	SvmName              *string                                    `json:"SvmName,omitempty"`
	Array                NullableStorageNetAppClusterRelationship   `json:"Array,omitempty"`
	Tenant               NullableStorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppExportPolicy StorageNetAppExportPolicy

// NewStorageNetAppExportPolicy instantiates a new StorageNetAppExportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppExportPolicy(classId string, objectType string) *StorageNetAppExportPolicy {
	this := StorageNetAppExportPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppExportPolicyWithDefaults instantiates a new StorageNetAppExportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppExportPolicyWithDefaults() *StorageNetAppExportPolicy {
	this := StorageNetAppExportPolicy{}
	var classId string = "storage.NetAppExportPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppExportPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppExportPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppExportPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppExportPolicy" of the ClassId field.
func (o *StorageNetAppExportPolicy) GetDefaultClassId() interface{} {
	return "storage.NetAppExportPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppExportPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppExportPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppExportPolicy" of the ObjectType field.
func (o *StorageNetAppExportPolicy) GetDefaultObjectType() interface{} {
	return "storage.NetAppExportPolicy"
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetClusterUuid() string {
	if o == nil || IsNil(o.ClusterUuid) {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetClusterUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterUuid) {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasClusterUuid() bool {
	if o != nil && !IsNil(o.ClusterUuid) {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *StorageNetAppExportPolicy) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetNetAppExportPolicyRule returns the NetAppExportPolicyRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicy) GetNetAppExportPolicyRule() []StorageNetAppExportPolicyRule {
	if o == nil {
		var ret []StorageNetAppExportPolicyRule
		return ret
	}
	return o.NetAppExportPolicyRule
}

// GetNetAppExportPolicyRuleOk returns a tuple with the NetAppExportPolicyRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicy) GetNetAppExportPolicyRuleOk() ([]StorageNetAppExportPolicyRule, bool) {
	if o == nil || IsNil(o.NetAppExportPolicyRule) {
		return nil, false
	}
	return o.NetAppExportPolicyRule, true
}

// HasNetAppExportPolicyRule returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasNetAppExportPolicyRule() bool {
	if o != nil && !IsNil(o.NetAppExportPolicyRule) {
		return true
	}

	return false
}

// SetNetAppExportPolicyRule gets a reference to the given []StorageNetAppExportPolicyRule and assigns it to the NetAppExportPolicyRule field.
func (o *StorageNetAppExportPolicy) SetNetAppExportPolicyRule(v []StorageNetAppExportPolicyRule) {
	o.NetAppExportPolicyRule = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetPolicyId() int64 {
	if o == nil || IsNil(o.PolicyId) {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetPolicyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *StorageNetAppExportPolicy) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetSvmName() string {
	if o == nil || IsNil(o.SvmName) {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetSvmNameOk() (*string, bool) {
	if o == nil || IsNil(o.SvmName) {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasSvmName() bool {
	if o != nil && !IsNil(o.SvmName) {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppExportPolicy) SetSvmName(v string) {
	o.SvmName = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicy) GetArray() StorageNetAppClusterRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicy) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppExportPolicy) SetArray(v StorageNetAppClusterRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageNetAppExportPolicy) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageNetAppExportPolicy) UnsetArray() {
	o.Array.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicy) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicy) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableStorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppExportPolicy) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *StorageNetAppExportPolicy) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *StorageNetAppExportPolicy) UnsetTenant() {
	o.Tenant.Unset()
}

func (o StorageNetAppExportPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppExportPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseNfsExport, errStorageBaseNfsExport := json.Marshal(o.StorageBaseNfsExport)
	if errStorageBaseNfsExport != nil {
		return map[string]interface{}{}, errStorageBaseNfsExport
	}
	errStorageBaseNfsExport = json.Unmarshal([]byte(serializedStorageBaseNfsExport), &toSerialize)
	if errStorageBaseNfsExport != nil {
		return map[string]interface{}{}, errStorageBaseNfsExport
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ClusterUuid) {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.NetAppExportPolicyRule != nil {
		toSerialize["NetAppExportPolicyRule"] = o.NetAppExportPolicyRule
	}
	if !IsNil(o.PolicyId) {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if !IsNil(o.SvmName) {
		toSerialize["SvmName"] = o.SvmName
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["Tenant"] = o.Tenant.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppExportPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppExportPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique identity of the device.
		ClusterUuid            *string                         `json:"ClusterUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		NetAppExportPolicyRule []StorageNetAppExportPolicyRule `json:"NetAppExportPolicyRule,omitempty"`
		// ID for the Export Policy.
		PolicyId *int64 `json:"PolicyId,omitempty"`
		// The storage virtual machine name for the export policy.
		SvmName *string                                    `json:"SvmName,omitempty"`
		Array   NullableStorageNetAppClusterRelationship   `json:"Array,omitempty"`
		Tenant  NullableStorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppExportPolicyWithoutEmbeddedStruct := StorageNetAppExportPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppExportPolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppExportPolicy := _StorageNetAppExportPolicy{}
		varStorageNetAppExportPolicy.ClassId = varStorageNetAppExportPolicyWithoutEmbeddedStruct.ClassId
		varStorageNetAppExportPolicy.ObjectType = varStorageNetAppExportPolicyWithoutEmbeddedStruct.ObjectType
		varStorageNetAppExportPolicy.ClusterUuid = varStorageNetAppExportPolicyWithoutEmbeddedStruct.ClusterUuid
		varStorageNetAppExportPolicy.NetAppExportPolicyRule = varStorageNetAppExportPolicyWithoutEmbeddedStruct.NetAppExportPolicyRule
		varStorageNetAppExportPolicy.PolicyId = varStorageNetAppExportPolicyWithoutEmbeddedStruct.PolicyId
		varStorageNetAppExportPolicy.SvmName = varStorageNetAppExportPolicyWithoutEmbeddedStruct.SvmName
		varStorageNetAppExportPolicy.Array = varStorageNetAppExportPolicyWithoutEmbeddedStruct.Array
		varStorageNetAppExportPolicy.Tenant = varStorageNetAppExportPolicyWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppExportPolicy(varStorageNetAppExportPolicy)
	} else {
		return err
	}

	varStorageNetAppExportPolicy := _StorageNetAppExportPolicy{}

	err = json.Unmarshal(data, &varStorageNetAppExportPolicy)
	if err == nil {
		o.StorageBaseNfsExport = varStorageNetAppExportPolicy.StorageBaseNfsExport
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "NetAppExportPolicyRule")
		delete(additionalProperties, "PolicyId")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageBaseNfsExport := reflect.ValueOf(o.StorageBaseNfsExport)
		for i := 0; i < reflectStorageBaseNfsExport.Type().NumField(); i++ {
			t := reflectStorageBaseNfsExport.Type().Field(i)

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

type NullableStorageNetAppExportPolicy struct {
	value *StorageNetAppExportPolicy
	isSet bool
}

func (v NullableStorageNetAppExportPolicy) Get() *StorageNetAppExportPolicy {
	return v.value
}

func (v *NullableStorageNetAppExportPolicy) Set(val *StorageNetAppExportPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppExportPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppExportPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppExportPolicy(val *StorageNetAppExportPolicy) *NullableStorageNetAppExportPolicy {
	return &NullableStorageNetAppExportPolicy{value: val, isSet: true}
}

func (v NullableStorageNetAppExportPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppExportPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
