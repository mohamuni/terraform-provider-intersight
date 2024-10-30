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

// checks if the NiatelemetryNexusCloudSite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryNexusCloudSite{}

// NiatelemetryNexusCloudSite Stores information of Nexus Cloud site devices.
type NiatelemetryNexusCloudSite struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Advisories setting status, based on license type.
	Advisories *bool `json:"Advisories,omitempty"`
	// Anomalies setting status, based on license type.
	Anomalies *bool `json:"Anomalies,omitempty"`
	// Capacity utilization setting status, based on license type.
	CapacityUtilization *bool `json:"CapacityUtilization,omitempty"`
	// Returns the license type of the device.
	LicenseType *string `json:"LicenseType,omitempty"`
	// Returns the name of the Nexus Cloud site.
	Name *string `json:"Name,omitempty"`
	// Returns the type of the Nexus Cloud site.
	SiteType *string `json:"SiteType,omitempty"`
	// Software management setting status, based on license type.
	SoftwareManagement *bool `json:"SoftwareManagement,omitempty"`
	// Returns the uuid of the Nexus Cloud site.
	Uuid                 *string                                           `json:"Uuid,omitempty"`
	NexusCloudAccount    NullableNiatelemetryNexusCloudAccountRelationship `json:"NexusCloudAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusCloudSite NiatelemetryNexusCloudSite

// NewNiatelemetryNexusCloudSite instantiates a new NiatelemetryNexusCloudSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusCloudSite(classId string, objectType string) *NiatelemetryNexusCloudSite {
	this := NiatelemetryNexusCloudSite{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusCloudSiteWithDefaults instantiates a new NiatelemetryNexusCloudSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusCloudSiteWithDefaults() *NiatelemetryNexusCloudSite {
	this := NiatelemetryNexusCloudSite{}
	var classId string = "niatelemetry.NexusCloudSite"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusCloudSite"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusCloudSite) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusCloudSite) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.NexusCloudSite" of the ClassId field.
func (o *NiatelemetryNexusCloudSite) GetDefaultClassId() interface{} {
	return "niatelemetry.NexusCloudSite"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusCloudSite) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusCloudSite) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.NexusCloudSite" of the ObjectType field.
func (o *NiatelemetryNexusCloudSite) GetDefaultObjectType() interface{} {
	return "niatelemetry.NexusCloudSite"
}

// GetAdvisories returns the Advisories field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetAdvisories() bool {
	if o == nil || IsNil(o.Advisories) {
		var ret bool
		return ret
	}
	return *o.Advisories
}

// GetAdvisoriesOk returns a tuple with the Advisories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetAdvisoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.Advisories) {
		return nil, false
	}
	return o.Advisories, true
}

// HasAdvisories returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasAdvisories() bool {
	if o != nil && !IsNil(o.Advisories) {
		return true
	}

	return false
}

// SetAdvisories gets a reference to the given bool and assigns it to the Advisories field.
func (o *NiatelemetryNexusCloudSite) SetAdvisories(v bool) {
	o.Advisories = &v
}

// GetAnomalies returns the Anomalies field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetAnomalies() bool {
	if o == nil || IsNil(o.Anomalies) {
		var ret bool
		return ret
	}
	return *o.Anomalies
}

// GetAnomaliesOk returns a tuple with the Anomalies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetAnomaliesOk() (*bool, bool) {
	if o == nil || IsNil(o.Anomalies) {
		return nil, false
	}
	return o.Anomalies, true
}

// HasAnomalies returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasAnomalies() bool {
	if o != nil && !IsNil(o.Anomalies) {
		return true
	}

	return false
}

// SetAnomalies gets a reference to the given bool and assigns it to the Anomalies field.
func (o *NiatelemetryNexusCloudSite) SetAnomalies(v bool) {
	o.Anomalies = &v
}

// GetCapacityUtilization returns the CapacityUtilization field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetCapacityUtilization() bool {
	if o == nil || IsNil(o.CapacityUtilization) {
		var ret bool
		return ret
	}
	return *o.CapacityUtilization
}

// GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetCapacityUtilizationOk() (*bool, bool) {
	if o == nil || IsNil(o.CapacityUtilization) {
		return nil, false
	}
	return o.CapacityUtilization, true
}

// HasCapacityUtilization returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasCapacityUtilization() bool {
	if o != nil && !IsNil(o.CapacityUtilization) {
		return true
	}

	return false
}

// SetCapacityUtilization gets a reference to the given bool and assigns it to the CapacityUtilization field.
func (o *NiatelemetryNexusCloudSite) SetCapacityUtilization(v bool) {
	o.CapacityUtilization = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetLicenseType() string {
	if o == nil || IsNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetLicenseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *NiatelemetryNexusCloudSite) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryNexusCloudSite) SetName(v string) {
	o.Name = &v
}

// GetSiteType returns the SiteType field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetSiteType() string {
	if o == nil || IsNil(o.SiteType) {
		var ret string
		return ret
	}
	return *o.SiteType
}

// GetSiteTypeOk returns a tuple with the SiteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetSiteTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SiteType) {
		return nil, false
	}
	return o.SiteType, true
}

// HasSiteType returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasSiteType() bool {
	if o != nil && !IsNil(o.SiteType) {
		return true
	}

	return false
}

// SetSiteType gets a reference to the given string and assigns it to the SiteType field.
func (o *NiatelemetryNexusCloudSite) SetSiteType(v string) {
	o.SiteType = &v
}

// GetSoftwareManagement returns the SoftwareManagement field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetSoftwareManagement() bool {
	if o == nil || IsNil(o.SoftwareManagement) {
		var ret bool
		return ret
	}
	return *o.SoftwareManagement
}

// GetSoftwareManagementOk returns a tuple with the SoftwareManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetSoftwareManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.SoftwareManagement) {
		return nil, false
	}
	return o.SoftwareManagement, true
}

// HasSoftwareManagement returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasSoftwareManagement() bool {
	if o != nil && !IsNil(o.SoftwareManagement) {
		return true
	}

	return false
}

// SetSoftwareManagement gets a reference to the given bool and assigns it to the SoftwareManagement field.
func (o *NiatelemetryNexusCloudSite) SetSoftwareManagement(v bool) {
	o.SoftwareManagement = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSite) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSite) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NiatelemetryNexusCloudSite) SetUuid(v string) {
	o.Uuid = &v
}

// GetNexusCloudAccount returns the NexusCloudAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNexusCloudSite) GetNexusCloudAccount() NiatelemetryNexusCloudAccountRelationship {
	if o == nil || IsNil(o.NexusCloudAccount.Get()) {
		var ret NiatelemetryNexusCloudAccountRelationship
		return ret
	}
	return *o.NexusCloudAccount.Get()
}

// GetNexusCloudAccountOk returns a tuple with the NexusCloudAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNexusCloudSite) GetNexusCloudAccountOk() (*NiatelemetryNexusCloudAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NexusCloudAccount.Get(), o.NexusCloudAccount.IsSet()
}

// HasNexusCloudAccount returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSite) HasNexusCloudAccount() bool {
	if o != nil && o.NexusCloudAccount.IsSet() {
		return true
	}

	return false
}

// SetNexusCloudAccount gets a reference to the given NullableNiatelemetryNexusCloudAccountRelationship and assigns it to the NexusCloudAccount field.
func (o *NiatelemetryNexusCloudSite) SetNexusCloudAccount(v NiatelemetryNexusCloudAccountRelationship) {
	o.NexusCloudAccount.Set(&v)
}

// SetNexusCloudAccountNil sets the value for NexusCloudAccount to be an explicit nil
func (o *NiatelemetryNexusCloudSite) SetNexusCloudAccountNil() {
	o.NexusCloudAccount.Set(nil)
}

// UnsetNexusCloudAccount ensures that no value is present for NexusCloudAccount, not even an explicit nil
func (o *NiatelemetryNexusCloudSite) UnsetNexusCloudAccount() {
	o.NexusCloudAccount.Unset()
}

func (o NiatelemetryNexusCloudSite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryNexusCloudSite) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Advisories) {
		toSerialize["Advisories"] = o.Advisories
	}
	if !IsNil(o.Anomalies) {
		toSerialize["Anomalies"] = o.Anomalies
	}
	if !IsNil(o.CapacityUtilization) {
		toSerialize["CapacityUtilization"] = o.CapacityUtilization
	}
	if !IsNil(o.LicenseType) {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SiteType) {
		toSerialize["SiteType"] = o.SiteType
	}
	if !IsNil(o.SoftwareManagement) {
		toSerialize["SoftwareManagement"] = o.SoftwareManagement
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.NexusCloudAccount.IsSet() {
		toSerialize["NexusCloudAccount"] = o.NexusCloudAccount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryNexusCloudSite) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryNexusCloudSiteWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Advisories setting status, based on license type.
		Advisories *bool `json:"Advisories,omitempty"`
		// Anomalies setting status, based on license type.
		Anomalies *bool `json:"Anomalies,omitempty"`
		// Capacity utilization setting status, based on license type.
		CapacityUtilization *bool `json:"CapacityUtilization,omitempty"`
		// Returns the license type of the device.
		LicenseType *string `json:"LicenseType,omitempty"`
		// Returns the name of the Nexus Cloud site.
		Name *string `json:"Name,omitempty"`
		// Returns the type of the Nexus Cloud site.
		SiteType *string `json:"SiteType,omitempty"`
		// Software management setting status, based on license type.
		SoftwareManagement *bool `json:"SoftwareManagement,omitempty"`
		// Returns the uuid of the Nexus Cloud site.
		Uuid              *string                                           `json:"Uuid,omitempty"`
		NexusCloudAccount NullableNiatelemetryNexusCloudAccountRelationship `json:"NexusCloudAccount,omitempty"`
	}

	varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct := NiatelemetryNexusCloudSiteWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNexusCloudSite := _NiatelemetryNexusCloudSite{}
		varNiatelemetryNexusCloudSite.ClassId = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.ClassId
		varNiatelemetryNexusCloudSite.ObjectType = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNexusCloudSite.Advisories = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.Advisories
		varNiatelemetryNexusCloudSite.Anomalies = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.Anomalies
		varNiatelemetryNexusCloudSite.CapacityUtilization = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.CapacityUtilization
		varNiatelemetryNexusCloudSite.LicenseType = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.LicenseType
		varNiatelemetryNexusCloudSite.Name = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.Name
		varNiatelemetryNexusCloudSite.SiteType = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.SiteType
		varNiatelemetryNexusCloudSite.SoftwareManagement = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.SoftwareManagement
		varNiatelemetryNexusCloudSite.Uuid = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.Uuid
		varNiatelemetryNexusCloudSite.NexusCloudAccount = varNiatelemetryNexusCloudSiteWithoutEmbeddedStruct.NexusCloudAccount
		*o = NiatelemetryNexusCloudSite(varNiatelemetryNexusCloudSite)
	} else {
		return err
	}

	varNiatelemetryNexusCloudSite := _NiatelemetryNexusCloudSite{}

	err = json.Unmarshal(data, &varNiatelemetryNexusCloudSite)
	if err == nil {
		o.MoBaseMo = varNiatelemetryNexusCloudSite.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Advisories")
		delete(additionalProperties, "Anomalies")
		delete(additionalProperties, "CapacityUtilization")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SiteType")
		delete(additionalProperties, "SoftwareManagement")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "NexusCloudAccount")

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

type NullableNiatelemetryNexusCloudSite struct {
	value *NiatelemetryNexusCloudSite
	isSet bool
}

func (v NullableNiatelemetryNexusCloudSite) Get() *NiatelemetryNexusCloudSite {
	return v.value
}

func (v *NullableNiatelemetryNexusCloudSite) Set(val *NiatelemetryNexusCloudSite) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusCloudSite) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusCloudSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusCloudSite(val *NiatelemetryNexusCloudSite) *NullableNiatelemetryNexusCloudSite {
	return &NullableNiatelemetryNexusCloudSite{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusCloudSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusCloudSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
