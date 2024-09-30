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

// checks if the CloudBaseNetworkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudBaseNetworkInterface{}

// CloudBaseNetworkInterface A base network interface object that is extended by Cloud Network Interface objects.
type CloudBaseNetworkInterface struct {
	VirtualizationBaseVirtualNetworkInterfaceCard
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                   `json:"ObjectType"`
	BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
	// CIDR scheme for defining an IP block.
	Cidr *string `json:"Cidr,omitempty"`
	// User friendly description of network interface.
	Description *string `json:"Description,omitempty"`
	// Internally generated identity of this network interface.
	Identity             *string                       `json:"Identity,omitempty"`
	RegionInfo           NullableCloudCloudRegion      `json:"RegionInfo,omitempty"`
	ZoneInfo             NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBaseNetworkInterface CloudBaseNetworkInterface

// NewCloudBaseNetworkInterface instantiates a new CloudBaseNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBaseNetworkInterface(classId string, objectType string) *CloudBaseNetworkInterface {
	this := CloudBaseNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adapterType string = "Unknown"
	this.AdapterType = &adapterType
	return &this
}

// NewCloudBaseNetworkInterfaceWithDefaults instantiates a new CloudBaseNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBaseNetworkInterfaceWithDefaults() *CloudBaseNetworkInterface {
	this := CloudBaseNetworkInterface{}
	var classId string = "cloud.AwsNetworkInterface"
	this.ClassId = classId
	var objectType string = "cloud.AwsNetworkInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBaseNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBaseNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBaseNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cloud.AwsNetworkInterface" of the ClassId field.
func (o *CloudBaseNetworkInterface) GetDefaultClassId() interface{} {
	return "cloud.AwsNetworkInterface"
}

// GetObjectType returns the ObjectType field value
func (o *CloudBaseNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBaseNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBaseNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cloud.AwsNetworkInterface" of the ObjectType field.
func (o *CloudBaseNetworkInterface) GetDefaultObjectType() interface{} {
	return "cloud.AwsNetworkInterface"
}

// GetBillingUnit returns the BillingUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseNetworkInterface) GetBillingUnit() CloudBillingUnit {
	if o == nil || IsNil(o.BillingUnit.Get()) {
		var ret CloudBillingUnit
		return ret
	}
	return *o.BillingUnit.Get()
}

// GetBillingUnitOk returns a tuple with the BillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseNetworkInterface) GetBillingUnitOk() (*CloudBillingUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingUnit.Get(), o.BillingUnit.IsSet()
}

// HasBillingUnit returns a boolean if a field has been set.
func (o *CloudBaseNetworkInterface) HasBillingUnit() bool {
	if o != nil && o.BillingUnit.IsSet() {
		return true
	}

	return false
}

// SetBillingUnit gets a reference to the given NullableCloudBillingUnit and assigns it to the BillingUnit field.
func (o *CloudBaseNetworkInterface) SetBillingUnit(v CloudBillingUnit) {
	o.BillingUnit.Set(&v)
}

// SetBillingUnitNil sets the value for BillingUnit to be an explicit nil
func (o *CloudBaseNetworkInterface) SetBillingUnitNil() {
	o.BillingUnit.Set(nil)
}

// UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
func (o *CloudBaseNetworkInterface) UnsetBillingUnit() {
	o.BillingUnit.Unset()
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CloudBaseNetworkInterface) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseNetworkInterface) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CloudBaseNetworkInterface) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CloudBaseNetworkInterface) SetCidr(v string) {
	o.Cidr = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudBaseNetworkInterface) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseNetworkInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudBaseNetworkInterface) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudBaseNetworkInterface) SetDescription(v string) {
	o.Description = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudBaseNetworkInterface) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseNetworkInterface) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudBaseNetworkInterface) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudBaseNetworkInterface) SetIdentity(v string) {
	o.Identity = &v
}

// GetRegionInfo returns the RegionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseNetworkInterface) GetRegionInfo() CloudCloudRegion {
	if o == nil || IsNil(o.RegionInfo.Get()) {
		var ret CloudCloudRegion
		return ret
	}
	return *o.RegionInfo.Get()
}

// GetRegionInfoOk returns a tuple with the RegionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseNetworkInterface) GetRegionInfoOk() (*CloudCloudRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegionInfo.Get(), o.RegionInfo.IsSet()
}

// HasRegionInfo returns a boolean if a field has been set.
func (o *CloudBaseNetworkInterface) HasRegionInfo() bool {
	if o != nil && o.RegionInfo.IsSet() {
		return true
	}

	return false
}

// SetRegionInfo gets a reference to the given NullableCloudCloudRegion and assigns it to the RegionInfo field.
func (o *CloudBaseNetworkInterface) SetRegionInfo(v CloudCloudRegion) {
	o.RegionInfo.Set(&v)
}

// SetRegionInfoNil sets the value for RegionInfo to be an explicit nil
func (o *CloudBaseNetworkInterface) SetRegionInfoNil() {
	o.RegionInfo.Set(nil)
}

// UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
func (o *CloudBaseNetworkInterface) UnsetRegionInfo() {
	o.RegionInfo.Unset()
}

// GetZoneInfo returns the ZoneInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseNetworkInterface) GetZoneInfo() CloudAvailabilityZone {
	if o == nil || IsNil(o.ZoneInfo.Get()) {
		var ret CloudAvailabilityZone
		return ret
	}
	return *o.ZoneInfo.Get()
}

// GetZoneInfoOk returns a tuple with the ZoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseNetworkInterface) GetZoneInfoOk() (*CloudAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneInfo.Get(), o.ZoneInfo.IsSet()
}

// HasZoneInfo returns a boolean if a field has been set.
func (o *CloudBaseNetworkInterface) HasZoneInfo() bool {
	if o != nil && o.ZoneInfo.IsSet() {
		return true
	}

	return false
}

// SetZoneInfo gets a reference to the given NullableCloudAvailabilityZone and assigns it to the ZoneInfo field.
func (o *CloudBaseNetworkInterface) SetZoneInfo(v CloudAvailabilityZone) {
	o.ZoneInfo.Set(&v)
}

// SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil
func (o *CloudBaseNetworkInterface) SetZoneInfoNil() {
	o.ZoneInfo.Set(nil)
}

// UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil
func (o *CloudBaseNetworkInterface) UnsetZoneInfo() {
	o.ZoneInfo.Unset()
}

func (o CloudBaseNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudBaseNetworkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualNetworkInterfaceCard, errVirtualizationBaseVirtualNetworkInterfaceCard := json.Marshal(o.VirtualizationBaseVirtualNetworkInterfaceCard)
	if errVirtualizationBaseVirtualNetworkInterfaceCard != nil {
		return map[string]interface{}{}, errVirtualizationBaseVirtualNetworkInterfaceCard
	}
	errVirtualizationBaseVirtualNetworkInterfaceCard = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualNetworkInterfaceCard), &toSerialize)
	if errVirtualizationBaseVirtualNetworkInterfaceCard != nil {
		return map[string]interface{}{}, errVirtualizationBaseVirtualNetworkInterfaceCard
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.BillingUnit.IsSet() {
		toSerialize["BillingUnit"] = o.BillingUnit.Get()
	}
	if !IsNil(o.Cidr) {
		toSerialize["Cidr"] = o.Cidr
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Identity) {
		toSerialize["Identity"] = o.Identity
	}
	if o.RegionInfo.IsSet() {
		toSerialize["RegionInfo"] = o.RegionInfo.Get()
	}
	if o.ZoneInfo.IsSet() {
		toSerialize["ZoneInfo"] = o.ZoneInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudBaseNetworkInterface) UnmarshalJSON(data []byte) (err error) {
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
	type CloudBaseNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType  string                   `json:"ObjectType"`
		BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
		// CIDR scheme for defining an IP block.
		Cidr *string `json:"Cidr,omitempty"`
		// User friendly description of network interface.
		Description *string `json:"Description,omitempty"`
		// Internally generated identity of this network interface.
		Identity   *string                       `json:"Identity,omitempty"`
		RegionInfo NullableCloudCloudRegion      `json:"RegionInfo,omitempty"`
		ZoneInfo   NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	}

	varCloudBaseNetworkInterfaceWithoutEmbeddedStruct := CloudBaseNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCloudBaseNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varCloudBaseNetworkInterface := _CloudBaseNetworkInterface{}
		varCloudBaseNetworkInterface.ClassId = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varCloudBaseNetworkInterface.ObjectType = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varCloudBaseNetworkInterface.BillingUnit = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.BillingUnit
		varCloudBaseNetworkInterface.Cidr = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.Cidr
		varCloudBaseNetworkInterface.Description = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.Description
		varCloudBaseNetworkInterface.Identity = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.Identity
		varCloudBaseNetworkInterface.RegionInfo = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.RegionInfo
		varCloudBaseNetworkInterface.ZoneInfo = varCloudBaseNetworkInterfaceWithoutEmbeddedStruct.ZoneInfo
		*o = CloudBaseNetworkInterface(varCloudBaseNetworkInterface)
	} else {
		return err
	}

	varCloudBaseNetworkInterface := _CloudBaseNetworkInterface{}

	err = json.Unmarshal(data, &varCloudBaseNetworkInterface)
	if err == nil {
		o.VirtualizationBaseVirtualNetworkInterfaceCard = varCloudBaseNetworkInterface.VirtualizationBaseVirtualNetworkInterfaceCard
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingUnit")
		delete(additionalProperties, "Cidr")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "RegionInfo")
		delete(additionalProperties, "ZoneInfo")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualNetworkInterfaceCard := reflect.ValueOf(o.VirtualizationBaseVirtualNetworkInterfaceCard)
		for i := 0; i < reflectVirtualizationBaseVirtualNetworkInterfaceCard.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualNetworkInterfaceCard.Type().Field(i)

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

type NullableCloudBaseNetworkInterface struct {
	value *CloudBaseNetworkInterface
	isSet bool
}

func (v NullableCloudBaseNetworkInterface) Get() *CloudBaseNetworkInterface {
	return v.value
}

func (v *NullableCloudBaseNetworkInterface) Set(val *CloudBaseNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBaseNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBaseNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBaseNetworkInterface(val *CloudBaseNetworkInterface) *NullableCloudBaseNetworkInterface {
	return &NullableCloudBaseNetworkInterface{value: val, isSet: true}
}

func (v NullableCloudBaseNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBaseNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
