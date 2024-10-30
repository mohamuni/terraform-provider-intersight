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

// checks if the FirmwareChassisUpgradeImpact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareChassisUpgradeImpact{}

// FirmwareChassisUpgradeImpact Impact of the chassis endpoint during the upgrade operation.
type FirmwareChassisUpgradeImpact struct {
	FirmwareBaseImpact
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                    `json:"ObjectType"`
	ImpactDetail []FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
	// Name of the chassis that will be affected by the upgrade.
	Name *string `json:"Name,omitempty"`
	// Details for the chassis that will be impacted by the upgrade.
	UserLabel            *string `json:"UserLabel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareChassisUpgradeImpact FirmwareChassisUpgradeImpact

// NewFirmwareChassisUpgradeImpact instantiates a new FirmwareChassisUpgradeImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareChassisUpgradeImpact(classId string, objectType string) *FirmwareChassisUpgradeImpact {
	this := FirmwareChassisUpgradeImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	var computationStatusDetail string = "Inprogress"
	this.ComputationStatusDetail = &computationStatusDetail
	var impactType string = "NoReboot"
	this.ImpactType = &impactType
	var versionImpact string = "None"
	this.VersionImpact = &versionImpact
	return &this
}

// NewFirmwareChassisUpgradeImpactWithDefaults instantiates a new FirmwareChassisUpgradeImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareChassisUpgradeImpactWithDefaults() *FirmwareChassisUpgradeImpact {
	this := FirmwareChassisUpgradeImpact{}
	var classId string = "firmware.ChassisUpgradeImpact"
	this.ClassId = classId
	var objectType string = "firmware.ChassisUpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareChassisUpgradeImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpact) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareChassisUpgradeImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.ChassisUpgradeImpact" of the ClassId field.
func (o *FirmwareChassisUpgradeImpact) GetDefaultClassId() interface{} {
	return "firmware.ChassisUpgradeImpact"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareChassisUpgradeImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareChassisUpgradeImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.ChassisUpgradeImpact" of the ObjectType field.
func (o *FirmwareChassisUpgradeImpact) GetDefaultObjectType() interface{} {
	return "firmware.ChassisUpgradeImpact"
}

// GetImpactDetail returns the ImpactDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareChassisUpgradeImpact) GetImpactDetail() []FirmwareComponentImpact {
	if o == nil {
		var ret []FirmwareComponentImpact
		return ret
	}
	return o.ImpactDetail
}

// GetImpactDetailOk returns a tuple with the ImpactDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareChassisUpgradeImpact) GetImpactDetailOk() ([]FirmwareComponentImpact, bool) {
	if o == nil || IsNil(o.ImpactDetail) {
		return nil, false
	}
	return o.ImpactDetail, true
}

// HasImpactDetail returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeImpact) HasImpactDetail() bool {
	if o != nil && !IsNil(o.ImpactDetail) {
		return true
	}

	return false
}

// SetImpactDetail gets a reference to the given []FirmwareComponentImpact and assigns it to the ImpactDetail field.
func (o *FirmwareChassisUpgradeImpact) SetImpactDetail(v []FirmwareComponentImpact) {
	o.ImpactDetail = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeImpact) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpact) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeImpact) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirmwareChassisUpgradeImpact) SetName(v string) {
	o.Name = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeImpact) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpact) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeImpact) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *FirmwareChassisUpgradeImpact) SetUserLabel(v string) {
	o.UserLabel = &v
}

func (o FirmwareChassisUpgradeImpact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareChassisUpgradeImpact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseImpact, errFirmwareBaseImpact := json.Marshal(o.FirmwareBaseImpact)
	if errFirmwareBaseImpact != nil {
		return map[string]interface{}{}, errFirmwareBaseImpact
	}
	errFirmwareBaseImpact = json.Unmarshal([]byte(serializedFirmwareBaseImpact), &toSerialize)
	if errFirmwareBaseImpact != nil {
		return map[string]interface{}{}, errFirmwareBaseImpact
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ImpactDetail != nil {
		toSerialize["ImpactDetail"] = o.ImpactDetail
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.UserLabel) {
		toSerialize["UserLabel"] = o.UserLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareChassisUpgradeImpact) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareChassisUpgradeImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                    `json:"ObjectType"`
		ImpactDetail []FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
		// Name of the chassis that will be affected by the upgrade.
		Name *string `json:"Name,omitempty"`
		// Details for the chassis that will be impacted by the upgrade.
		UserLabel *string `json:"UserLabel,omitempty"`
	}

	varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct := FirmwareChassisUpgradeImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareChassisUpgradeImpact := _FirmwareChassisUpgradeImpact{}
		varFirmwareChassisUpgradeImpact.ClassId = varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct.ClassId
		varFirmwareChassisUpgradeImpact.ObjectType = varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct.ObjectType
		varFirmwareChassisUpgradeImpact.ImpactDetail = varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct.ImpactDetail
		varFirmwareChassisUpgradeImpact.Name = varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct.Name
		varFirmwareChassisUpgradeImpact.UserLabel = varFirmwareChassisUpgradeImpactWithoutEmbeddedStruct.UserLabel
		*o = FirmwareChassisUpgradeImpact(varFirmwareChassisUpgradeImpact)
	} else {
		return err
	}

	varFirmwareChassisUpgradeImpact := _FirmwareChassisUpgradeImpact{}

	err = json.Unmarshal(data, &varFirmwareChassisUpgradeImpact)
	if err == nil {
		o.FirmwareBaseImpact = varFirmwareChassisUpgradeImpact.FirmwareBaseImpact
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ImpactDetail")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserLabel")

		// remove fields from embedded structs
		reflectFirmwareBaseImpact := reflect.ValueOf(o.FirmwareBaseImpact)
		for i := 0; i < reflectFirmwareBaseImpact.Type().NumField(); i++ {
			t := reflectFirmwareBaseImpact.Type().Field(i)

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

type NullableFirmwareChassisUpgradeImpact struct {
	value *FirmwareChassisUpgradeImpact
	isSet bool
}

func (v NullableFirmwareChassisUpgradeImpact) Get() *FirmwareChassisUpgradeImpact {
	return v.value
}

func (v *NullableFirmwareChassisUpgradeImpact) Set(val *FirmwareChassisUpgradeImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareChassisUpgradeImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareChassisUpgradeImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareChassisUpgradeImpact(val *FirmwareChassisUpgradeImpact) *NullableFirmwareChassisUpgradeImpact {
	return &NullableFirmwareChassisUpgradeImpact{value: val, isSet: true}
}

func (v NullableFirmwareChassisUpgradeImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareChassisUpgradeImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
