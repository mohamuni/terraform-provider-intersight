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
	"reflect"
	"strings"
)

// StorageExternalParityGroup The following properties are stored for each external parity group.
type StorageExternalParityGroup struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Cache mode. Either E (enabled) or D (disabled) is displayed.
	CacheMode    *string              `json:"CacheMode,omitempty"`
	ExternalLuns []StorageExternalLun `json:"ExternalLuns,omitempty"`
	// External parity group number.
	ExternalParityGroupId *string `json:"ExternalParityGroupId,omitempty"`
	// Status of the external parity group.
	ExternalParityGroupStatus *string `json:"ExternalParityGroupStatus,omitempty"`
	// Whether the data direct mapping attribute is enabled.
	IsDataDirectMapping *bool `json:"IsDataDirectMapping,omitempty"`
	// Inflow cache control. Either true (enabled) or false (disabled) is displayed.
	IsInflowControlEnabled *bool `json:"IsInflowControlEnabled,omitempty"`
	// The load balancing method for I-O operations for the external storage system.
	LoadBalanceMode *string `json:"LoadBalanceMode,omitempty"`
	// The blade number of the MP blade assigned to the external parity group.
	MpBladeId *int64 `json:"MpBladeId,omitempty"`
	// Path mode of the external storage system.
	PathMode             *string `json:"PathMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageExternalParityGroup StorageExternalParityGroup

// NewStorageExternalParityGroup instantiates a new StorageExternalParityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageExternalParityGroup(classId string, objectType string) *StorageExternalParityGroup {
	this := StorageExternalParityGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageExternalParityGroupWithDefaults instantiates a new StorageExternalParityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageExternalParityGroupWithDefaults() *StorageExternalParityGroup {
	this := StorageExternalParityGroup{}
	var classId string = "storage.ExternalParityGroup"
	this.ClassId = classId
	var objectType string = "storage.ExternalParityGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageExternalParityGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageExternalParityGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageExternalParityGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageExternalParityGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCacheMode returns the CacheMode field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetCacheMode() string {
	if o == nil || o.CacheMode == nil {
		var ret string
		return ret
	}
	return *o.CacheMode
}

// GetCacheModeOk returns a tuple with the CacheMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetCacheModeOk() (*string, bool) {
	if o == nil || o.CacheMode == nil {
		return nil, false
	}
	return o.CacheMode, true
}

// HasCacheMode returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasCacheMode() bool {
	if o != nil && o.CacheMode != nil {
		return true
	}

	return false
}

// SetCacheMode gets a reference to the given string and assigns it to the CacheMode field.
func (o *StorageExternalParityGroup) SetCacheMode(v string) {
	o.CacheMode = &v
}

// GetExternalLuns returns the ExternalLuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageExternalParityGroup) GetExternalLuns() []StorageExternalLun {
	if o == nil {
		var ret []StorageExternalLun
		return ret
	}
	return o.ExternalLuns
}

// GetExternalLunsOk returns a tuple with the ExternalLuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageExternalParityGroup) GetExternalLunsOk() ([]StorageExternalLun, bool) {
	if o == nil || o.ExternalLuns == nil {
		return nil, false
	}
	return o.ExternalLuns, true
}

// HasExternalLuns returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasExternalLuns() bool {
	if o != nil && o.ExternalLuns != nil {
		return true
	}

	return false
}

// SetExternalLuns gets a reference to the given []StorageExternalLun and assigns it to the ExternalLuns field.
func (o *StorageExternalParityGroup) SetExternalLuns(v []StorageExternalLun) {
	o.ExternalLuns = v
}

// GetExternalParityGroupId returns the ExternalParityGroupId field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetExternalParityGroupId() string {
	if o == nil || o.ExternalParityGroupId == nil {
		var ret string
		return ret
	}
	return *o.ExternalParityGroupId
}

// GetExternalParityGroupIdOk returns a tuple with the ExternalParityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetExternalParityGroupIdOk() (*string, bool) {
	if o == nil || o.ExternalParityGroupId == nil {
		return nil, false
	}
	return o.ExternalParityGroupId, true
}

// HasExternalParityGroupId returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasExternalParityGroupId() bool {
	if o != nil && o.ExternalParityGroupId != nil {
		return true
	}

	return false
}

// SetExternalParityGroupId gets a reference to the given string and assigns it to the ExternalParityGroupId field.
func (o *StorageExternalParityGroup) SetExternalParityGroupId(v string) {
	o.ExternalParityGroupId = &v
}

// GetExternalParityGroupStatus returns the ExternalParityGroupStatus field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetExternalParityGroupStatus() string {
	if o == nil || o.ExternalParityGroupStatus == nil {
		var ret string
		return ret
	}
	return *o.ExternalParityGroupStatus
}

// GetExternalParityGroupStatusOk returns a tuple with the ExternalParityGroupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetExternalParityGroupStatusOk() (*string, bool) {
	if o == nil || o.ExternalParityGroupStatus == nil {
		return nil, false
	}
	return o.ExternalParityGroupStatus, true
}

// HasExternalParityGroupStatus returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasExternalParityGroupStatus() bool {
	if o != nil && o.ExternalParityGroupStatus != nil {
		return true
	}

	return false
}

// SetExternalParityGroupStatus gets a reference to the given string and assigns it to the ExternalParityGroupStatus field.
func (o *StorageExternalParityGroup) SetExternalParityGroupStatus(v string) {
	o.ExternalParityGroupStatus = &v
}

// GetIsDataDirectMapping returns the IsDataDirectMapping field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetIsDataDirectMapping() bool {
	if o == nil || o.IsDataDirectMapping == nil {
		var ret bool
		return ret
	}
	return *o.IsDataDirectMapping
}

// GetIsDataDirectMappingOk returns a tuple with the IsDataDirectMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetIsDataDirectMappingOk() (*bool, bool) {
	if o == nil || o.IsDataDirectMapping == nil {
		return nil, false
	}
	return o.IsDataDirectMapping, true
}

// HasIsDataDirectMapping returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasIsDataDirectMapping() bool {
	if o != nil && o.IsDataDirectMapping != nil {
		return true
	}

	return false
}

// SetIsDataDirectMapping gets a reference to the given bool and assigns it to the IsDataDirectMapping field.
func (o *StorageExternalParityGroup) SetIsDataDirectMapping(v bool) {
	o.IsDataDirectMapping = &v
}

// GetIsInflowControlEnabled returns the IsInflowControlEnabled field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetIsInflowControlEnabled() bool {
	if o == nil || o.IsInflowControlEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsInflowControlEnabled
}

// GetIsInflowControlEnabledOk returns a tuple with the IsInflowControlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetIsInflowControlEnabledOk() (*bool, bool) {
	if o == nil || o.IsInflowControlEnabled == nil {
		return nil, false
	}
	return o.IsInflowControlEnabled, true
}

// HasIsInflowControlEnabled returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasIsInflowControlEnabled() bool {
	if o != nil && o.IsInflowControlEnabled != nil {
		return true
	}

	return false
}

// SetIsInflowControlEnabled gets a reference to the given bool and assigns it to the IsInflowControlEnabled field.
func (o *StorageExternalParityGroup) SetIsInflowControlEnabled(v bool) {
	o.IsInflowControlEnabled = &v
}

// GetLoadBalanceMode returns the LoadBalanceMode field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetLoadBalanceMode() string {
	if o == nil || o.LoadBalanceMode == nil {
		var ret string
		return ret
	}
	return *o.LoadBalanceMode
}

// GetLoadBalanceModeOk returns a tuple with the LoadBalanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetLoadBalanceModeOk() (*string, bool) {
	if o == nil || o.LoadBalanceMode == nil {
		return nil, false
	}
	return o.LoadBalanceMode, true
}

// HasLoadBalanceMode returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasLoadBalanceMode() bool {
	if o != nil && o.LoadBalanceMode != nil {
		return true
	}

	return false
}

// SetLoadBalanceMode gets a reference to the given string and assigns it to the LoadBalanceMode field.
func (o *StorageExternalParityGroup) SetLoadBalanceMode(v string) {
	o.LoadBalanceMode = &v
}

// GetMpBladeId returns the MpBladeId field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetMpBladeId() int64 {
	if o == nil || o.MpBladeId == nil {
		var ret int64
		return ret
	}
	return *o.MpBladeId
}

// GetMpBladeIdOk returns a tuple with the MpBladeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetMpBladeIdOk() (*int64, bool) {
	if o == nil || o.MpBladeId == nil {
		return nil, false
	}
	return o.MpBladeId, true
}

// HasMpBladeId returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasMpBladeId() bool {
	if o != nil && o.MpBladeId != nil {
		return true
	}

	return false
}

// SetMpBladeId gets a reference to the given int64 and assigns it to the MpBladeId field.
func (o *StorageExternalParityGroup) SetMpBladeId(v int64) {
	o.MpBladeId = &v
}

// GetPathMode returns the PathMode field value if set, zero value otherwise.
func (o *StorageExternalParityGroup) GetPathMode() string {
	if o == nil || o.PathMode == nil {
		var ret string
		return ret
	}
	return *o.PathMode
}

// GetPathModeOk returns a tuple with the PathMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalParityGroup) GetPathModeOk() (*string, bool) {
	if o == nil || o.PathMode == nil {
		return nil, false
	}
	return o.PathMode, true
}

// HasPathMode returns a boolean if a field has been set.
func (o *StorageExternalParityGroup) HasPathMode() bool {
	if o != nil && o.PathMode != nil {
		return true
	}

	return false
}

// SetPathMode gets a reference to the given string and assigns it to the PathMode field.
func (o *StorageExternalParityGroup) SetPathMode(v string) {
	o.PathMode = &v
}

func (o StorageExternalParityGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CacheMode != nil {
		toSerialize["CacheMode"] = o.CacheMode
	}
	if o.ExternalLuns != nil {
		toSerialize["ExternalLuns"] = o.ExternalLuns
	}
	if o.ExternalParityGroupId != nil {
		toSerialize["ExternalParityGroupId"] = o.ExternalParityGroupId
	}
	if o.ExternalParityGroupStatus != nil {
		toSerialize["ExternalParityGroupStatus"] = o.ExternalParityGroupStatus
	}
	if o.IsDataDirectMapping != nil {
		toSerialize["IsDataDirectMapping"] = o.IsDataDirectMapping
	}
	if o.IsInflowControlEnabled != nil {
		toSerialize["IsInflowControlEnabled"] = o.IsInflowControlEnabled
	}
	if o.LoadBalanceMode != nil {
		toSerialize["LoadBalanceMode"] = o.LoadBalanceMode
	}
	if o.MpBladeId != nil {
		toSerialize["MpBladeId"] = o.MpBladeId
	}
	if o.PathMode != nil {
		toSerialize["PathMode"] = o.PathMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageExternalParityGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageExternalParityGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Cache mode. Either E (enabled) or D (disabled) is displayed.
		CacheMode    *string              `json:"CacheMode,omitempty"`
		ExternalLuns []StorageExternalLun `json:"ExternalLuns,omitempty"`
		// External parity group number.
		ExternalParityGroupId *string `json:"ExternalParityGroupId,omitempty"`
		// Status of the external parity group.
		ExternalParityGroupStatus *string `json:"ExternalParityGroupStatus,omitempty"`
		// Whether the data direct mapping attribute is enabled.
		IsDataDirectMapping *bool `json:"IsDataDirectMapping,omitempty"`
		// Inflow cache control. Either true (enabled) or false (disabled) is displayed.
		IsInflowControlEnabled *bool `json:"IsInflowControlEnabled,omitempty"`
		// The load balancing method for I-O operations for the external storage system.
		LoadBalanceMode *string `json:"LoadBalanceMode,omitempty"`
		// The blade number of the MP blade assigned to the external parity group.
		MpBladeId *int64 `json:"MpBladeId,omitempty"`
		// Path mode of the external storage system.
		PathMode *string `json:"PathMode,omitempty"`
	}

	varStorageExternalParityGroupWithoutEmbeddedStruct := StorageExternalParityGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageExternalParityGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageExternalParityGroup := _StorageExternalParityGroup{}
		varStorageExternalParityGroup.ClassId = varStorageExternalParityGroupWithoutEmbeddedStruct.ClassId
		varStorageExternalParityGroup.ObjectType = varStorageExternalParityGroupWithoutEmbeddedStruct.ObjectType
		varStorageExternalParityGroup.CacheMode = varStorageExternalParityGroupWithoutEmbeddedStruct.CacheMode
		varStorageExternalParityGroup.ExternalLuns = varStorageExternalParityGroupWithoutEmbeddedStruct.ExternalLuns
		varStorageExternalParityGroup.ExternalParityGroupId = varStorageExternalParityGroupWithoutEmbeddedStruct.ExternalParityGroupId
		varStorageExternalParityGroup.ExternalParityGroupStatus = varStorageExternalParityGroupWithoutEmbeddedStruct.ExternalParityGroupStatus
		varStorageExternalParityGroup.IsDataDirectMapping = varStorageExternalParityGroupWithoutEmbeddedStruct.IsDataDirectMapping
		varStorageExternalParityGroup.IsInflowControlEnabled = varStorageExternalParityGroupWithoutEmbeddedStruct.IsInflowControlEnabled
		varStorageExternalParityGroup.LoadBalanceMode = varStorageExternalParityGroupWithoutEmbeddedStruct.LoadBalanceMode
		varStorageExternalParityGroup.MpBladeId = varStorageExternalParityGroupWithoutEmbeddedStruct.MpBladeId
		varStorageExternalParityGroup.PathMode = varStorageExternalParityGroupWithoutEmbeddedStruct.PathMode
		*o = StorageExternalParityGroup(varStorageExternalParityGroup)
	} else {
		return err
	}

	varStorageExternalParityGroup := _StorageExternalParityGroup{}

	err = json.Unmarshal(bytes, &varStorageExternalParityGroup)
	if err == nil {
		o.MoBaseComplexType = varStorageExternalParityGroup.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CacheMode")
		delete(additionalProperties, "ExternalLuns")
		delete(additionalProperties, "ExternalParityGroupId")
		delete(additionalProperties, "ExternalParityGroupStatus")
		delete(additionalProperties, "IsDataDirectMapping")
		delete(additionalProperties, "IsInflowControlEnabled")
		delete(additionalProperties, "LoadBalanceMode")
		delete(additionalProperties, "MpBladeId")
		delete(additionalProperties, "PathMode")

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

type NullableStorageExternalParityGroup struct {
	value *StorageExternalParityGroup
	isSet bool
}

func (v NullableStorageExternalParityGroup) Get() *StorageExternalParityGroup {
	return v.value
}

func (v *NullableStorageExternalParityGroup) Set(val *StorageExternalParityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageExternalParityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageExternalParityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageExternalParityGroup(val *StorageExternalParityGroup) *NullableStorageExternalParityGroup {
	return &NullableStorageExternalParityGroup{value: val, isSet: true}
}

func (v NullableStorageExternalParityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageExternalParityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
