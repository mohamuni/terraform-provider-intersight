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

// checks if the StorageNetAppNfsService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppNfsService{}

// StorageNetAppNfsService An NFS service retrieves the NFS configuration of a storage virtual machine.
type StorageNetAppNfsService struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether NFSv3 protocol is enabled.
	NfsV3Enabled *bool `json:"NfsV3Enabled,omitempty"`
	// Specifies whether NFSv4.1 protocol is enabled.
	NfsV41Enabled *bool `json:"NfsV41Enabled,omitempty"`
	// Specifies whether NFSv4.0 protocol is enabled.
	NfsV4Enabled *bool `json:"NfsV4Enabled,omitempty"`
	// Unique identifier for the NetApp Storage Virtual Machine.
	SvmUuid              *string                                    `json:"SvmUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	Tenant               NullableStorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNfsService StorageNetAppNfsService

// NewStorageNetAppNfsService instantiates a new StorageNetAppNfsService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNfsService(classId string, objectType string) *StorageNetAppNfsService {
	this := StorageNetAppNfsService{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNfsServiceWithDefaults instantiates a new StorageNetAppNfsService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNfsServiceWithDefaults() *StorageNetAppNfsService {
	this := StorageNetAppNfsService{}
	var classId string = "storage.NetAppNfsService"
	this.ClassId = classId
	var objectType string = "storage.NetAppNfsService"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNfsService) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsService) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNfsService) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppNfsService" of the ClassId field.
func (o *StorageNetAppNfsService) GetDefaultClassId() interface{} {
	return "storage.NetAppNfsService"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNfsService) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsService) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNfsService) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppNfsService" of the ObjectType field.
func (o *StorageNetAppNfsService) GetDefaultObjectType() interface{} {
	return "storage.NetAppNfsService"
}

// GetNfsV3Enabled returns the NfsV3Enabled field value if set, zero value otherwise.
func (o *StorageNetAppNfsService) GetNfsV3Enabled() bool {
	if o == nil || IsNil(o.NfsV3Enabled) {
		var ret bool
		return ret
	}
	return *o.NfsV3Enabled
}

// GetNfsV3EnabledOk returns a tuple with the NfsV3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsService) GetNfsV3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NfsV3Enabled) {
		return nil, false
	}
	return o.NfsV3Enabled, true
}

// HasNfsV3Enabled returns a boolean if a field has been set.
func (o *StorageNetAppNfsService) HasNfsV3Enabled() bool {
	if o != nil && !IsNil(o.NfsV3Enabled) {
		return true
	}

	return false
}

// SetNfsV3Enabled gets a reference to the given bool and assigns it to the NfsV3Enabled field.
func (o *StorageNetAppNfsService) SetNfsV3Enabled(v bool) {
	o.NfsV3Enabled = &v
}

// GetNfsV41Enabled returns the NfsV41Enabled field value if set, zero value otherwise.
func (o *StorageNetAppNfsService) GetNfsV41Enabled() bool {
	if o == nil || IsNil(o.NfsV41Enabled) {
		var ret bool
		return ret
	}
	return *o.NfsV41Enabled
}

// GetNfsV41EnabledOk returns a tuple with the NfsV41Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsService) GetNfsV41EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NfsV41Enabled) {
		return nil, false
	}
	return o.NfsV41Enabled, true
}

// HasNfsV41Enabled returns a boolean if a field has been set.
func (o *StorageNetAppNfsService) HasNfsV41Enabled() bool {
	if o != nil && !IsNil(o.NfsV41Enabled) {
		return true
	}

	return false
}

// SetNfsV41Enabled gets a reference to the given bool and assigns it to the NfsV41Enabled field.
func (o *StorageNetAppNfsService) SetNfsV41Enabled(v bool) {
	o.NfsV41Enabled = &v
}

// GetNfsV4Enabled returns the NfsV4Enabled field value if set, zero value otherwise.
func (o *StorageNetAppNfsService) GetNfsV4Enabled() bool {
	if o == nil || IsNil(o.NfsV4Enabled) {
		var ret bool
		return ret
	}
	return *o.NfsV4Enabled
}

// GetNfsV4EnabledOk returns a tuple with the NfsV4Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsService) GetNfsV4EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NfsV4Enabled) {
		return nil, false
	}
	return o.NfsV4Enabled, true
}

// HasNfsV4Enabled returns a boolean if a field has been set.
func (o *StorageNetAppNfsService) HasNfsV4Enabled() bool {
	if o != nil && !IsNil(o.NfsV4Enabled) {
		return true
	}

	return false
}

// SetNfsV4Enabled gets a reference to the given bool and assigns it to the NfsV4Enabled field.
func (o *StorageNetAppNfsService) SetNfsV4Enabled(v bool) {
	o.NfsV4Enabled = &v
}

// GetSvmUuid returns the SvmUuid field value if set, zero value otherwise.
func (o *StorageNetAppNfsService) GetSvmUuid() string {
	if o == nil || IsNil(o.SvmUuid) {
		var ret string
		return ret
	}
	return *o.SvmUuid
}

// GetSvmUuidOk returns a tuple with the SvmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsService) GetSvmUuidOk() (*string, bool) {
	if o == nil || IsNil(o.SvmUuid) {
		return nil, false
	}
	return o.SvmUuid, true
}

// HasSvmUuid returns a boolean if a field has been set.
func (o *StorageNetAppNfsService) HasSvmUuid() bool {
	if o != nil && !IsNil(o.SvmUuid) {
		return true
	}

	return false
}

// SetSvmUuid gets a reference to the given string and assigns it to the SvmUuid field.
func (o *StorageNetAppNfsService) SetSvmUuid(v string) {
	o.SvmUuid = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNfsService) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNfsService) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppNfsService) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableStorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppNfsService) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *StorageNetAppNfsService) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *StorageNetAppNfsService) UnsetTenant() {
	o.Tenant.Unset()
}

func (o StorageNetAppNfsService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppNfsService) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NfsV3Enabled) {
		toSerialize["NfsV3Enabled"] = o.NfsV3Enabled
	}
	if !IsNil(o.NfsV41Enabled) {
		toSerialize["NfsV41Enabled"] = o.NfsV41Enabled
	}
	if !IsNil(o.NfsV4Enabled) {
		toSerialize["NfsV4Enabled"] = o.NfsV4Enabled
	}
	if !IsNil(o.SvmUuid) {
		toSerialize["SvmUuid"] = o.SvmUuid
	}
	if o.Tenant.IsSet() {
		toSerialize["Tenant"] = o.Tenant.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppNfsService) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppNfsServiceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specifies whether NFSv3 protocol is enabled.
		NfsV3Enabled *bool `json:"NfsV3Enabled,omitempty"`
		// Specifies whether NFSv4.1 protocol is enabled.
		NfsV41Enabled *bool `json:"NfsV41Enabled,omitempty"`
		// Specifies whether NFSv4.0 protocol is enabled.
		NfsV4Enabled *bool `json:"NfsV4Enabled,omitempty"`
		// Unique identifier for the NetApp Storage Virtual Machine.
		SvmUuid *string                                    `json:"SvmUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		Tenant  NullableStorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppNfsServiceWithoutEmbeddedStruct := StorageNetAppNfsServiceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppNfsServiceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppNfsService := _StorageNetAppNfsService{}
		varStorageNetAppNfsService.ClassId = varStorageNetAppNfsServiceWithoutEmbeddedStruct.ClassId
		varStorageNetAppNfsService.ObjectType = varStorageNetAppNfsServiceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppNfsService.NfsV3Enabled = varStorageNetAppNfsServiceWithoutEmbeddedStruct.NfsV3Enabled
		varStorageNetAppNfsService.NfsV41Enabled = varStorageNetAppNfsServiceWithoutEmbeddedStruct.NfsV41Enabled
		varStorageNetAppNfsService.NfsV4Enabled = varStorageNetAppNfsServiceWithoutEmbeddedStruct.NfsV4Enabled
		varStorageNetAppNfsService.SvmUuid = varStorageNetAppNfsServiceWithoutEmbeddedStruct.SvmUuid
		varStorageNetAppNfsService.Tenant = varStorageNetAppNfsServiceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppNfsService(varStorageNetAppNfsService)
	} else {
		return err
	}

	varStorageNetAppNfsService := _StorageNetAppNfsService{}

	err = json.Unmarshal(data, &varStorageNetAppNfsService)
	if err == nil {
		o.MoBaseMo = varStorageNetAppNfsService.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NfsV3Enabled")
		delete(additionalProperties, "NfsV41Enabled")
		delete(additionalProperties, "NfsV4Enabled")
		delete(additionalProperties, "SvmUuid")
		delete(additionalProperties, "Tenant")

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

type NullableStorageNetAppNfsService struct {
	value *StorageNetAppNfsService
	isSet bool
}

func (v NullableStorageNetAppNfsService) Get() *StorageNetAppNfsService {
	return v.value
}

func (v *NullableStorageNetAppNfsService) Set(val *StorageNetAppNfsService) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNfsService) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNfsService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNfsService(val *StorageNetAppNfsService) *NullableStorageNetAppNfsService {
	return &NullableStorageNetAppNfsService{value: val, isSet: true}
}

func (v NullableStorageNetAppNfsService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNfsService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
