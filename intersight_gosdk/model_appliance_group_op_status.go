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

// checks if the ApplianceGroupOpStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceGroupOpStatus{}

// ApplianceGroupOpStatus Status of a group of applications.
type ApplianceGroupOpStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the group.
	Description *string `json:"Description,omitempty"`
	// The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Infra and Appliance.
	GroupName *string `json:"GroupName,omitempty"`
	// The overall API status from this group's applications.
	OverallStatus *string `json:"OverallStatus,omitempty"`
	// An array of relationships to applianceAppOpStatus resources.
	Apps                 []ApplianceAppOpStatusRelationship          `json:"Apps,omitempty"`
	SystemOpStatus       NullableApplianceSystemOpStatusRelationship `json:"SystemOpStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceGroupOpStatus ApplianceGroupOpStatus

// NewApplianceGroupOpStatus instantiates a new ApplianceGroupOpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceGroupOpStatus(classId string, objectType string) *ApplianceGroupOpStatus {
	this := ApplianceGroupOpStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceGroupOpStatusWithDefaults instantiates a new ApplianceGroupOpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceGroupOpStatusWithDefaults() *ApplianceGroupOpStatus {
	this := ApplianceGroupOpStatus{}
	var classId string = "appliance.GroupOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.GroupOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceGroupOpStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceGroupOpStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.GroupOpStatus" of the ClassId field.
func (o *ApplianceGroupOpStatus) GetDefaultClassId() interface{} {
	return "appliance.GroupOpStatus"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceGroupOpStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceGroupOpStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.GroupOpStatus" of the ObjectType field.
func (o *ApplianceGroupOpStatus) GetDefaultObjectType() interface{} {
	return "appliance.GroupOpStatus"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatus) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatus) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplianceGroupOpStatus) SetDescription(v string) {
	o.Description = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatus) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatus) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatus) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ApplianceGroupOpStatus) SetGroupName(v string) {
	o.GroupName = &v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatus) GetOverallStatus() string {
	if o == nil || IsNil(o.OverallStatus) {
		var ret string
		return ret
	}
	return *o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatus) GetOverallStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OverallStatus) {
		return nil, false
	}
	return o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatus) HasOverallStatus() bool {
	if o != nil && !IsNil(o.OverallStatus) {
		return true
	}

	return false
}

// SetOverallStatus gets a reference to the given string and assigns it to the OverallStatus field.
func (o *ApplianceGroupOpStatus) SetOverallStatus(v string) {
	o.OverallStatus = &v
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceGroupOpStatus) GetApps() []ApplianceAppOpStatusRelationship {
	if o == nil {
		var ret []ApplianceAppOpStatusRelationship
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceGroupOpStatus) GetAppsOk() ([]ApplianceAppOpStatusRelationship, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatus) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []ApplianceAppOpStatusRelationship and assigns it to the Apps field.
func (o *ApplianceGroupOpStatus) SetApps(v []ApplianceAppOpStatusRelationship) {
	o.Apps = v
}

// GetSystemOpStatus returns the SystemOpStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceGroupOpStatus) GetSystemOpStatus() ApplianceSystemOpStatusRelationship {
	if o == nil || IsNil(o.SystemOpStatus.Get()) {
		var ret ApplianceSystemOpStatusRelationship
		return ret
	}
	return *o.SystemOpStatus.Get()
}

// GetSystemOpStatusOk returns a tuple with the SystemOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceGroupOpStatus) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemOpStatus.Get(), o.SystemOpStatus.IsSet()
}

// HasSystemOpStatus returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatus) HasSystemOpStatus() bool {
	if o != nil && o.SystemOpStatus.IsSet() {
		return true
	}

	return false
}

// SetSystemOpStatus gets a reference to the given NullableApplianceSystemOpStatusRelationship and assigns it to the SystemOpStatus field.
func (o *ApplianceGroupOpStatus) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship) {
	o.SystemOpStatus.Set(&v)
}

// SetSystemOpStatusNil sets the value for SystemOpStatus to be an explicit nil
func (o *ApplianceGroupOpStatus) SetSystemOpStatusNil() {
	o.SystemOpStatus.Set(nil)
}

// UnsetSystemOpStatus ensures that no value is present for SystemOpStatus, not even an explicit nil
func (o *ApplianceGroupOpStatus) UnsetSystemOpStatus() {
	o.SystemOpStatus.Unset()
}

func (o ApplianceGroupOpStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceGroupOpStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GroupName) {
		toSerialize["GroupName"] = o.GroupName
	}
	if !IsNil(o.OverallStatus) {
		toSerialize["OverallStatus"] = o.OverallStatus
	}
	if o.Apps != nil {
		toSerialize["Apps"] = o.Apps
	}
	if o.SystemOpStatus.IsSet() {
		toSerialize["SystemOpStatus"] = o.SystemOpStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceGroupOpStatus) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceGroupOpStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the group.
		Description *string `json:"Description,omitempty"`
		// The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Infra and Appliance.
		GroupName *string `json:"GroupName,omitempty"`
		// The overall API status from this group's applications.
		OverallStatus *string `json:"OverallStatus,omitempty"`
		// An array of relationships to applianceAppOpStatus resources.
		Apps           []ApplianceAppOpStatusRelationship          `json:"Apps,omitempty"`
		SystemOpStatus NullableApplianceSystemOpStatusRelationship `json:"SystemOpStatus,omitempty"`
	}

	varApplianceGroupOpStatusWithoutEmbeddedStruct := ApplianceGroupOpStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceGroupOpStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceGroupOpStatus := _ApplianceGroupOpStatus{}
		varApplianceGroupOpStatus.ClassId = varApplianceGroupOpStatusWithoutEmbeddedStruct.ClassId
		varApplianceGroupOpStatus.ObjectType = varApplianceGroupOpStatusWithoutEmbeddedStruct.ObjectType
		varApplianceGroupOpStatus.Description = varApplianceGroupOpStatusWithoutEmbeddedStruct.Description
		varApplianceGroupOpStatus.GroupName = varApplianceGroupOpStatusWithoutEmbeddedStruct.GroupName
		varApplianceGroupOpStatus.OverallStatus = varApplianceGroupOpStatusWithoutEmbeddedStruct.OverallStatus
		varApplianceGroupOpStatus.Apps = varApplianceGroupOpStatusWithoutEmbeddedStruct.Apps
		varApplianceGroupOpStatus.SystemOpStatus = varApplianceGroupOpStatusWithoutEmbeddedStruct.SystemOpStatus
		*o = ApplianceGroupOpStatus(varApplianceGroupOpStatus)
	} else {
		return err
	}

	varApplianceGroupOpStatus := _ApplianceGroupOpStatus{}

	err = json.Unmarshal(data, &varApplianceGroupOpStatus)
	if err == nil {
		o.MoBaseMo = varApplianceGroupOpStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "GroupName")
		delete(additionalProperties, "OverallStatus")
		delete(additionalProperties, "Apps")
		delete(additionalProperties, "SystemOpStatus")

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

type NullableApplianceGroupOpStatus struct {
	value *ApplianceGroupOpStatus
	isSet bool
}

func (v NullableApplianceGroupOpStatus) Get() *ApplianceGroupOpStatus {
	return v.value
}

func (v *NullableApplianceGroupOpStatus) Set(val *ApplianceGroupOpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceGroupOpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceGroupOpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceGroupOpStatus(val *ApplianceGroupOpStatus) *NullableApplianceGroupOpStatus {
	return &NullableApplianceGroupOpStatus{value: val, isSet: true}
}

func (v NullableApplianceGroupOpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceGroupOpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
