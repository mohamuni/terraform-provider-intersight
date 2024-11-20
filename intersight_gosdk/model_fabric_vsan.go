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

// checks if the FabricVsan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricVsan{}

// FabricVsan Configuration object sent by user to create VSAN configurations.
type FabricVsan struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables or Disables the default zoning state. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	// Deprecated
	DefaultZoning *string `json:"DefaultZoning,omitempty"`
	// Logical grouping mode for fc ports.
	// Deprecated
	FcZoneSharingMode *string `json:"FcZoneSharingMode,omitempty"`
	// FCOE Vlan associated to the VSAN configuration.
	FcoeVlan *int64 `json:"FcoeVlan,omitempty"`
	// User given name for the VSAN configuration.
	Name *string `json:"Name,omitempty"`
	// Virtual San Identifier in the switch.
	VsanId *int64 `json:"VsanId,omitempty"`
	// Used to indicate whether the VSAN Id is defined for storage or uplink or both traffics in FI. * `Uplink` - Vsan associated with uplink network. * `Storage` - Vsan associated with storage network. * `Common` - Vsan that is common for uplink and storage network.
	VsanScope            *string                                   `json:"VsanScope,omitempty"`
	FcNetworkPolicy      NullableFabricFcNetworkPolicyRelationship `json:"FcNetworkPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVsan FabricVsan

// NewFabricVsan instantiates a new FabricVsan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVsan(classId string, objectType string) *FabricVsan {
	this := FabricVsan{}
	this.ClassId = classId
	this.ObjectType = objectType
	var defaultZoning string = "Enabled"
	this.DefaultZoning = &defaultZoning
	var vsanScope string = "Uplink"
	this.VsanScope = &vsanScope
	return &this
}

// NewFabricVsanWithDefaults instantiates a new FabricVsan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVsanWithDefaults() *FabricVsan {
	this := FabricVsan{}
	var classId string = "fabric.Vsan"
	this.ClassId = classId
	var objectType string = "fabric.Vsan"
	this.ObjectType = objectType
	var defaultZoning string = "Enabled"
	this.DefaultZoning = &defaultZoning
	var vsanScope string = "Uplink"
	this.VsanScope = &vsanScope
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVsan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVsan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVsan) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.Vsan" of the ClassId field.
func (o *FabricVsan) GetDefaultClassId() interface{} {
	return "fabric.Vsan"
}

// GetObjectType returns the ObjectType field value
func (o *FabricVsan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVsan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVsan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.Vsan" of the ObjectType field.
func (o *FabricVsan) GetDefaultObjectType() interface{} {
	return "fabric.Vsan"
}

// GetDefaultZoning returns the DefaultZoning field value if set, zero value otherwise.
// Deprecated
func (o *FabricVsan) GetDefaultZoning() string {
	if o == nil || IsNil(o.DefaultZoning) {
		var ret string
		return ret
	}
	return *o.DefaultZoning
}

// GetDefaultZoningOk returns a tuple with the DefaultZoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FabricVsan) GetDefaultZoningOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultZoning) {
		return nil, false
	}
	return o.DefaultZoning, true
}

// HasDefaultZoning returns a boolean if a field has been set.
func (o *FabricVsan) HasDefaultZoning() bool {
	if o != nil && !IsNil(o.DefaultZoning) {
		return true
	}

	return false
}

// SetDefaultZoning gets a reference to the given string and assigns it to the DefaultZoning field.
// Deprecated
func (o *FabricVsan) SetDefaultZoning(v string) {
	o.DefaultZoning = &v
}

// GetFcZoneSharingMode returns the FcZoneSharingMode field value if set, zero value otherwise.
// Deprecated
func (o *FabricVsan) GetFcZoneSharingMode() string {
	if o == nil || IsNil(o.FcZoneSharingMode) {
		var ret string
		return ret
	}
	return *o.FcZoneSharingMode
}

// GetFcZoneSharingModeOk returns a tuple with the FcZoneSharingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FabricVsan) GetFcZoneSharingModeOk() (*string, bool) {
	if o == nil || IsNil(o.FcZoneSharingMode) {
		return nil, false
	}
	return o.FcZoneSharingMode, true
}

// HasFcZoneSharingMode returns a boolean if a field has been set.
func (o *FabricVsan) HasFcZoneSharingMode() bool {
	if o != nil && !IsNil(o.FcZoneSharingMode) {
		return true
	}

	return false
}

// SetFcZoneSharingMode gets a reference to the given string and assigns it to the FcZoneSharingMode field.
// Deprecated
func (o *FabricVsan) SetFcZoneSharingMode(v string) {
	o.FcZoneSharingMode = &v
}

// GetFcoeVlan returns the FcoeVlan field value if set, zero value otherwise.
func (o *FabricVsan) GetFcoeVlan() int64 {
	if o == nil || IsNil(o.FcoeVlan) {
		var ret int64
		return ret
	}
	return *o.FcoeVlan
}

// GetFcoeVlanOk returns a tuple with the FcoeVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsan) GetFcoeVlanOk() (*int64, bool) {
	if o == nil || IsNil(o.FcoeVlan) {
		return nil, false
	}
	return o.FcoeVlan, true
}

// HasFcoeVlan returns a boolean if a field has been set.
func (o *FabricVsan) HasFcoeVlan() bool {
	if o != nil && !IsNil(o.FcoeVlan) {
		return true
	}

	return false
}

// SetFcoeVlan gets a reference to the given int64 and assigns it to the FcoeVlan field.
func (o *FabricVsan) SetFcoeVlan(v int64) {
	o.FcoeVlan = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVsan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVsan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVsan) SetName(v string) {
	o.Name = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricVsan) GetVsanId() int64 {
	if o == nil || IsNil(o.VsanId) {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsan) GetVsanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VsanId) {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricVsan) HasVsanId() bool {
	if o != nil && !IsNil(o.VsanId) {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricVsan) SetVsanId(v int64) {
	o.VsanId = &v
}

// GetVsanScope returns the VsanScope field value if set, zero value otherwise.
func (o *FabricVsan) GetVsanScope() string {
	if o == nil || IsNil(o.VsanScope) {
		var ret string
		return ret
	}
	return *o.VsanScope
}

// GetVsanScopeOk returns a tuple with the VsanScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsan) GetVsanScopeOk() (*string, bool) {
	if o == nil || IsNil(o.VsanScope) {
		return nil, false
	}
	return o.VsanScope, true
}

// HasVsanScope returns a boolean if a field has been set.
func (o *FabricVsan) HasVsanScope() bool {
	if o != nil && !IsNil(o.VsanScope) {
		return true
	}

	return false
}

// SetVsanScope gets a reference to the given string and assigns it to the VsanScope field.
func (o *FabricVsan) SetVsanScope(v string) {
	o.VsanScope = &v
}

// GetFcNetworkPolicy returns the FcNetworkPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVsan) GetFcNetworkPolicy() FabricFcNetworkPolicyRelationship {
	if o == nil || IsNil(o.FcNetworkPolicy.Get()) {
		var ret FabricFcNetworkPolicyRelationship
		return ret
	}
	return *o.FcNetworkPolicy.Get()
}

// GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVsan) GetFcNetworkPolicyOk() (*FabricFcNetworkPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.FcNetworkPolicy.Get(), o.FcNetworkPolicy.IsSet()
}

// HasFcNetworkPolicy returns a boolean if a field has been set.
func (o *FabricVsan) HasFcNetworkPolicy() bool {
	if o != nil && o.FcNetworkPolicy.IsSet() {
		return true
	}

	return false
}

// SetFcNetworkPolicy gets a reference to the given NullableFabricFcNetworkPolicyRelationship and assigns it to the FcNetworkPolicy field.
func (o *FabricVsan) SetFcNetworkPolicy(v FabricFcNetworkPolicyRelationship) {
	o.FcNetworkPolicy.Set(&v)
}

// SetFcNetworkPolicyNil sets the value for FcNetworkPolicy to be an explicit nil
func (o *FabricVsan) SetFcNetworkPolicyNil() {
	o.FcNetworkPolicy.Set(nil)
}

// UnsetFcNetworkPolicy ensures that no value is present for FcNetworkPolicy, not even an explicit nil
func (o *FabricVsan) UnsetFcNetworkPolicy() {
	o.FcNetworkPolicy.Unset()
}

func (o FabricVsan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricVsan) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DefaultZoning) {
		toSerialize["DefaultZoning"] = o.DefaultZoning
	}
	if !IsNil(o.FcZoneSharingMode) {
		toSerialize["FcZoneSharingMode"] = o.FcZoneSharingMode
	}
	if !IsNil(o.FcoeVlan) {
		toSerialize["FcoeVlan"] = o.FcoeVlan
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.VsanId) {
		toSerialize["VsanId"] = o.VsanId
	}
	if !IsNil(o.VsanScope) {
		toSerialize["VsanScope"] = o.VsanScope
	}
	if o.FcNetworkPolicy.IsSet() {
		toSerialize["FcNetworkPolicy"] = o.FcNetworkPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricVsan) UnmarshalJSON(data []byte) (err error) {
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
	type FabricVsanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables or Disables the default zoning state. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
		// Deprecated
		DefaultZoning *string `json:"DefaultZoning,omitempty"`
		// Logical grouping mode for fc ports.
		// Deprecated
		FcZoneSharingMode *string `json:"FcZoneSharingMode,omitempty"`
		// FCOE Vlan associated to the VSAN configuration.
		FcoeVlan *int64 `json:"FcoeVlan,omitempty"`
		// User given name for the VSAN configuration.
		Name *string `json:"Name,omitempty"`
		// Virtual San Identifier in the switch.
		VsanId *int64 `json:"VsanId,omitempty"`
		// Used to indicate whether the VSAN Id is defined for storage or uplink or both traffics in FI. * `Uplink` - Vsan associated with uplink network. * `Storage` - Vsan associated with storage network. * `Common` - Vsan that is common for uplink and storage network.
		VsanScope       *string                                   `json:"VsanScope,omitempty"`
		FcNetworkPolicy NullableFabricFcNetworkPolicyRelationship `json:"FcNetworkPolicy,omitempty"`
	}

	varFabricVsanWithoutEmbeddedStruct := FabricVsanWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricVsanWithoutEmbeddedStruct)
	if err == nil {
		varFabricVsan := _FabricVsan{}
		varFabricVsan.ClassId = varFabricVsanWithoutEmbeddedStruct.ClassId
		varFabricVsan.ObjectType = varFabricVsanWithoutEmbeddedStruct.ObjectType
		varFabricVsan.DefaultZoning = varFabricVsanWithoutEmbeddedStruct.DefaultZoning
		varFabricVsan.FcZoneSharingMode = varFabricVsanWithoutEmbeddedStruct.FcZoneSharingMode
		varFabricVsan.FcoeVlan = varFabricVsanWithoutEmbeddedStruct.FcoeVlan
		varFabricVsan.Name = varFabricVsanWithoutEmbeddedStruct.Name
		varFabricVsan.VsanId = varFabricVsanWithoutEmbeddedStruct.VsanId
		varFabricVsan.VsanScope = varFabricVsanWithoutEmbeddedStruct.VsanScope
		varFabricVsan.FcNetworkPolicy = varFabricVsanWithoutEmbeddedStruct.FcNetworkPolicy
		*o = FabricVsan(varFabricVsan)
	} else {
		return err
	}

	varFabricVsan := _FabricVsan{}

	err = json.Unmarshal(data, &varFabricVsan)
	if err == nil {
		o.MoBaseMo = varFabricVsan.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DefaultZoning")
		delete(additionalProperties, "FcZoneSharingMode")
		delete(additionalProperties, "FcoeVlan")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VsanId")
		delete(additionalProperties, "VsanScope")
		delete(additionalProperties, "FcNetworkPolicy")

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

type NullableFabricVsan struct {
	value *FabricVsan
	isSet bool
}

func (v NullableFabricVsan) Get() *FabricVsan {
	return v.value
}

func (v *NullableFabricVsan) Set(val *FabricVsan) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVsan) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVsan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVsan(val *FabricVsan) *NullableFabricVsan {
	return &NullableFabricVsan{value: val, isSet: true}
}

func (v NullableFabricVsan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVsan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
