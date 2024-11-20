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

// checks if the KubernetesAciCniApic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAciCniApic{}

// KubernetesAciCniApic Internally generated object of claimed APIC device known to Razor.
type KubernetesAciCniApic struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Moid of the apic device under the asset service.
	AssetApicMoid *string `json:"AssetApicMoid,omitempty"`
	// The number of ACI CNI profiles configured for this APIC.
	NumAciCniProfiles *int64 `json:"NumAciCniProfiles,omitempty"`
	// An array of relationships to kubernetesAciCniProfile resources.
	AciCniProfiles       []KubernetesAciCniProfileRelationship        `json:"AciCniProfiles,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship  `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAciCniApic KubernetesAciCniApic

// NewKubernetesAciCniApic instantiates a new KubernetesAciCniApic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAciCniApic(classId string, objectType string) *KubernetesAciCniApic {
	this := KubernetesAciCniApic{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAciCniApicWithDefaults instantiates a new KubernetesAciCniApic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAciCniApicWithDefaults() *KubernetesAciCniApic {
	this := KubernetesAciCniApic{}
	var classId string = "kubernetes.AciCniApic"
	this.ClassId = classId
	var objectType string = "kubernetes.AciCniApic"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAciCniApic) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAciCniApic) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.AciCniApic" of the ClassId field.
func (o *KubernetesAciCniApic) GetDefaultClassId() interface{} {
	return "kubernetes.AciCniApic"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAciCniApic) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAciCniApic) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.AciCniApic" of the ObjectType field.
func (o *KubernetesAciCniApic) GetDefaultObjectType() interface{} {
	return "kubernetes.AciCniApic"
}

// GetAssetApicMoid returns the AssetApicMoid field value if set, zero value otherwise.
func (o *KubernetesAciCniApic) GetAssetApicMoid() string {
	if o == nil || IsNil(o.AssetApicMoid) {
		var ret string
		return ret
	}
	return *o.AssetApicMoid
}

// GetAssetApicMoidOk returns a tuple with the AssetApicMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetAssetApicMoidOk() (*string, bool) {
	if o == nil || IsNil(o.AssetApicMoid) {
		return nil, false
	}
	return o.AssetApicMoid, true
}

// HasAssetApicMoid returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasAssetApicMoid() bool {
	if o != nil && !IsNil(o.AssetApicMoid) {
		return true
	}

	return false
}

// SetAssetApicMoid gets a reference to the given string and assigns it to the AssetApicMoid field.
func (o *KubernetesAciCniApic) SetAssetApicMoid(v string) {
	o.AssetApicMoid = &v
}

// GetNumAciCniProfiles returns the NumAciCniProfiles field value if set, zero value otherwise.
func (o *KubernetesAciCniApic) GetNumAciCniProfiles() int64 {
	if o == nil || IsNil(o.NumAciCniProfiles) {
		var ret int64
		return ret
	}
	return *o.NumAciCniProfiles
}

// GetNumAciCniProfilesOk returns a tuple with the NumAciCniProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetNumAciCniProfilesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumAciCniProfiles) {
		return nil, false
	}
	return o.NumAciCniProfiles, true
}

// HasNumAciCniProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasNumAciCniProfiles() bool {
	if o != nil && !IsNil(o.NumAciCniProfiles) {
		return true
	}

	return false
}

// SetNumAciCniProfiles gets a reference to the given int64 and assigns it to the NumAciCniProfiles field.
func (o *KubernetesAciCniApic) SetNumAciCniProfiles(v int64) {
	o.NumAciCniProfiles = &v
}

// GetAciCniProfiles returns the AciCniProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniApic) GetAciCniProfiles() []KubernetesAciCniProfileRelationship {
	if o == nil {
		var ret []KubernetesAciCniProfileRelationship
		return ret
	}
	return o.AciCniProfiles
}

// GetAciCniProfilesOk returns a tuple with the AciCniProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniApic) GetAciCniProfilesOk() ([]KubernetesAciCniProfileRelationship, bool) {
	if o == nil || IsNil(o.AciCniProfiles) {
		return nil, false
	}
	return o.AciCniProfiles, true
}

// HasAciCniProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasAciCniProfiles() bool {
	if o != nil && !IsNil(o.AciCniProfiles) {
		return true
	}

	return false
}

// SetAciCniProfiles gets a reference to the given []KubernetesAciCniProfileRelationship and assigns it to the AciCniProfiles field.
func (o *KubernetesAciCniApic) SetAciCniProfiles(v []KubernetesAciCniProfileRelationship) {
	o.AciCniProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniApic) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniApic) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAciCniApic) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *KubernetesAciCniApic) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *KubernetesAciCniApic) UnsetOrganization() {
	o.Organization.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniApic) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniApic) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesAciCniApic) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *KubernetesAciCniApic) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *KubernetesAciCniApic) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o KubernetesAciCniApic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAciCniApic) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AssetApicMoid) {
		toSerialize["AssetApicMoid"] = o.AssetApicMoid
	}
	if !IsNil(o.NumAciCniProfiles) {
		toSerialize["NumAciCniProfiles"] = o.NumAciCniProfiles
	}
	if o.AciCniProfiles != nil {
		toSerialize["AciCniProfiles"] = o.AciCniProfiles
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesAciCniApic) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesAciCniApicWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Moid of the apic device under the asset service.
		AssetApicMoid *string `json:"AssetApicMoid,omitempty"`
		// The number of ACI CNI profiles configured for this APIC.
		NumAciCniProfiles *int64 `json:"NumAciCniProfiles,omitempty"`
		// An array of relationships to kubernetesAciCniProfile resources.
		AciCniProfiles   []KubernetesAciCniProfileRelationship        `json:"AciCniProfiles,omitempty"`
		Organization     NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship  `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesAciCniApicWithoutEmbeddedStruct := KubernetesAciCniApicWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesAciCniApicWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAciCniApic := _KubernetesAciCniApic{}
		varKubernetesAciCniApic.ClassId = varKubernetesAciCniApicWithoutEmbeddedStruct.ClassId
		varKubernetesAciCniApic.ObjectType = varKubernetesAciCniApicWithoutEmbeddedStruct.ObjectType
		varKubernetesAciCniApic.AssetApicMoid = varKubernetesAciCniApicWithoutEmbeddedStruct.AssetApicMoid
		varKubernetesAciCniApic.NumAciCniProfiles = varKubernetesAciCniApicWithoutEmbeddedStruct.NumAciCniProfiles
		varKubernetesAciCniApic.AciCniProfiles = varKubernetesAciCniApicWithoutEmbeddedStruct.AciCniProfiles
		varKubernetesAciCniApic.Organization = varKubernetesAciCniApicWithoutEmbeddedStruct.Organization
		varKubernetesAciCniApic.RegisteredDevice = varKubernetesAciCniApicWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesAciCniApic(varKubernetesAciCniApic)
	} else {
		return err
	}

	varKubernetesAciCniApic := _KubernetesAciCniApic{}

	err = json.Unmarshal(data, &varKubernetesAciCniApic)
	if err == nil {
		o.MoBaseMo = varKubernetesAciCniApic.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetApicMoid")
		delete(additionalProperties, "NumAciCniProfiles")
		delete(additionalProperties, "AciCniProfiles")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableKubernetesAciCniApic struct {
	value *KubernetesAciCniApic
	isSet bool
}

func (v NullableKubernetesAciCniApic) Get() *KubernetesAciCniApic {
	return v.value
}

func (v *NullableKubernetesAciCniApic) Set(val *KubernetesAciCniApic) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAciCniApic) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAciCniApic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAciCniApic(val *KubernetesAciCniApic) *NullableKubernetesAciCniApic {
	return &NullableKubernetesAciCniApic{value: val, isSet: true}
}

func (v NullableKubernetesAciCniApic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAciCniApic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
