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

// checks if the OsDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsDistribution{}

// OsDistribution Intersight has the distribution details for all the Intersight supported OS distributions. There will be a Distribution object for each supported OS.
type OsDistribution struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An internal property that is used to denote if the OS Distribution is supported by Intersight for Automated Installation.
	IsSupported *bool `json:"IsSupported,omitempty"`
	// The label of the OS distribution such as ESXi, CentOS to be displayed.
	Label *string `json:"Label,omitempty"`
	// The name of the OS distribution such as ESXi, CentOS.
	Name *string `json:"Name,omitempty"`
	// An internal property that is used to denote if the OS Distribution is supported by the Server Configuration Utility.
	ScuSupported         *bool                                        `json:"ScuSupported,omitempty"`
	SupportedEditions    []string                                     `json:"SupportedEditions,omitempty"`
	Catalog              NullableOsCatalogRelationship                `json:"Catalog,omitempty"`
	Vendor               NullableHclOperatingSystemVendorRelationship `json:"Vendor,omitempty"`
	Version              NullableHclOperatingSystemRelationship       `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsDistribution OsDistribution

// NewOsDistribution instantiates a new OsDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsDistribution(classId string, objectType string) *OsDistribution {
	this := OsDistribution{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsDistributionWithDefaults instantiates a new OsDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsDistributionWithDefaults() *OsDistribution {
	this := OsDistribution{}
	var classId string = "os.Distribution"
	this.ClassId = classId
	var objectType string = "os.Distribution"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsDistribution) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsDistribution) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.Distribution" of the ClassId field.
func (o *OsDistribution) GetDefaultClassId() interface{} {
	return "os.Distribution"
}

// GetObjectType returns the ObjectType field value
func (o *OsDistribution) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsDistribution) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.Distribution" of the ObjectType field.
func (o *OsDistribution) GetDefaultObjectType() interface{} {
	return "os.Distribution"
}

// GetIsSupported returns the IsSupported field value if set, zero value otherwise.
func (o *OsDistribution) GetIsSupported() bool {
	if o == nil || IsNil(o.IsSupported) {
		var ret bool
		return ret
	}
	return *o.IsSupported
}

// GetIsSupportedOk returns a tuple with the IsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetIsSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupported) {
		return nil, false
	}
	return o.IsSupported, true
}

// HasIsSupported returns a boolean if a field has been set.
func (o *OsDistribution) HasIsSupported() bool {
	if o != nil && !IsNil(o.IsSupported) {
		return true
	}

	return false
}

// SetIsSupported gets a reference to the given bool and assigns it to the IsSupported field.
func (o *OsDistribution) SetIsSupported(v bool) {
	o.IsSupported = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OsDistribution) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OsDistribution) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OsDistribution) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsDistribution) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsDistribution) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsDistribution) SetName(v string) {
	o.Name = &v
}

// GetScuSupported returns the ScuSupported field value if set, zero value otherwise.
func (o *OsDistribution) GetScuSupported() bool {
	if o == nil || IsNil(o.ScuSupported) {
		var ret bool
		return ret
	}
	return *o.ScuSupported
}

// GetScuSupportedOk returns a tuple with the ScuSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetScuSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.ScuSupported) {
		return nil, false
	}
	return o.ScuSupported, true
}

// HasScuSupported returns a boolean if a field has been set.
func (o *OsDistribution) HasScuSupported() bool {
	if o != nil && !IsNil(o.ScuSupported) {
		return true
	}

	return false
}

// SetScuSupported gets a reference to the given bool and assigns it to the ScuSupported field.
func (o *OsDistribution) SetScuSupported(v bool) {
	o.ScuSupported = &v
}

// GetSupportedEditions returns the SupportedEditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsDistribution) GetSupportedEditions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedEditions
}

// GetSupportedEditionsOk returns a tuple with the SupportedEditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsDistribution) GetSupportedEditionsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedEditions) {
		return nil, false
	}
	return o.SupportedEditions, true
}

// HasSupportedEditions returns a boolean if a field has been set.
func (o *OsDistribution) HasSupportedEditions() bool {
	if o != nil && !IsNil(o.SupportedEditions) {
		return true
	}

	return false
}

// SetSupportedEditions gets a reference to the given []string and assigns it to the SupportedEditions field.
func (o *OsDistribution) SetSupportedEditions(v []string) {
	o.SupportedEditions = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsDistribution) GetCatalog() OsCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret OsCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsDistribution) GetCatalogOk() (*OsCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *OsDistribution) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableOsCatalogRelationship and assigns it to the Catalog field.
func (o *OsDistribution) SetCatalog(v OsCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *OsDistribution) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *OsDistribution) UnsetCatalog() {
	o.Catalog.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsDistribution) GetVendor() HclOperatingSystemVendorRelationship {
	if o == nil || IsNil(o.Vendor.Get()) {
		var ret HclOperatingSystemVendorRelationship
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsDistribution) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *OsDistribution) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableHclOperatingSystemVendorRelationship and assigns it to the Vendor field.
func (o *OsDistribution) SetVendor(v HclOperatingSystemVendorRelationship) {
	o.Vendor.Set(&v)
}

// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *OsDistribution) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *OsDistribution) UnsetVendor() {
	o.Vendor.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsDistribution) GetVersion() HclOperatingSystemRelationship {
	if o == nil || IsNil(o.Version.Get()) {
		var ret HclOperatingSystemRelationship
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsDistribution) GetVersionOk() (*HclOperatingSystemRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *OsDistribution) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableHclOperatingSystemRelationship and assigns it to the Version field.
func (o *OsDistribution) SetVersion(v HclOperatingSystemRelationship) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *OsDistribution) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *OsDistribution) UnsetVersion() {
	o.Version.Unset()
}

func (o OsDistribution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsDistribution) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsSupported) {
		toSerialize["IsSupported"] = o.IsSupported
	}
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.ScuSupported) {
		toSerialize["ScuSupported"] = o.ScuSupported
	}
	if o.SupportedEditions != nil {
		toSerialize["SupportedEditions"] = o.SupportedEditions
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["Vendor"] = o.Vendor.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsDistribution) UnmarshalJSON(data []byte) (err error) {
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
	type OsDistributionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An internal property that is used to denote if the OS Distribution is supported by Intersight for Automated Installation.
		IsSupported *bool `json:"IsSupported,omitempty"`
		// The label of the OS distribution such as ESXi, CentOS to be displayed.
		Label *string `json:"Label,omitempty"`
		// The name of the OS distribution such as ESXi, CentOS.
		Name *string `json:"Name,omitempty"`
		// An internal property that is used to denote if the OS Distribution is supported by the Server Configuration Utility.
		ScuSupported      *bool                                        `json:"ScuSupported,omitempty"`
		SupportedEditions []string                                     `json:"SupportedEditions,omitempty"`
		Catalog           NullableOsCatalogRelationship                `json:"Catalog,omitempty"`
		Vendor            NullableHclOperatingSystemVendorRelationship `json:"Vendor,omitempty"`
		Version           NullableHclOperatingSystemRelationship       `json:"Version,omitempty"`
	}

	varOsDistributionWithoutEmbeddedStruct := OsDistributionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsDistributionWithoutEmbeddedStruct)
	if err == nil {
		varOsDistribution := _OsDistribution{}
		varOsDistribution.ClassId = varOsDistributionWithoutEmbeddedStruct.ClassId
		varOsDistribution.ObjectType = varOsDistributionWithoutEmbeddedStruct.ObjectType
		varOsDistribution.IsSupported = varOsDistributionWithoutEmbeddedStruct.IsSupported
		varOsDistribution.Label = varOsDistributionWithoutEmbeddedStruct.Label
		varOsDistribution.Name = varOsDistributionWithoutEmbeddedStruct.Name
		varOsDistribution.ScuSupported = varOsDistributionWithoutEmbeddedStruct.ScuSupported
		varOsDistribution.SupportedEditions = varOsDistributionWithoutEmbeddedStruct.SupportedEditions
		varOsDistribution.Catalog = varOsDistributionWithoutEmbeddedStruct.Catalog
		varOsDistribution.Vendor = varOsDistributionWithoutEmbeddedStruct.Vendor
		varOsDistribution.Version = varOsDistributionWithoutEmbeddedStruct.Version
		*o = OsDistribution(varOsDistribution)
	} else {
		return err
	}

	varOsDistribution := _OsDistribution{}

	err = json.Unmarshal(data, &varOsDistribution)
	if err == nil {
		o.MoBaseMo = varOsDistribution.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsSupported")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ScuSupported")
		delete(additionalProperties, "SupportedEditions")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "Version")

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

type NullableOsDistribution struct {
	value *OsDistribution
	isSet bool
}

func (v NullableOsDistribution) Get() *OsDistribution {
	return v.value
}

func (v *NullableOsDistribution) Set(val *OsDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableOsDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableOsDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsDistribution(val *OsDistribution) *NullableOsDistribution {
	return &NullableOsDistribution{value: val, isSet: true}
}

func (v NullableOsDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
