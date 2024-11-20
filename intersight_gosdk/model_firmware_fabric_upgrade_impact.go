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

// checks if the FirmwareFabricUpgradeImpact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareFabricUpgradeImpact{}

// FirmwareFabricUpgradeImpact Impact for the Fabric Interconnect endpoint during the upgrade operation.
type FirmwareFabricUpgradeImpact struct {
	FirmwareBaseImpact
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                    `json:"ObjectType"`
	ImpactDetail []FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
	// Details for the Fabric Interconnect that will be impacted by the upgrade.
	Serial               *string `json:"Serial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareFabricUpgradeImpact FirmwareFabricUpgradeImpact

// NewFirmwareFabricUpgradeImpact instantiates a new FirmwareFabricUpgradeImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareFabricUpgradeImpact(classId string, objectType string) *FirmwareFabricUpgradeImpact {
	this := FirmwareFabricUpgradeImpact{}
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

// NewFirmwareFabricUpgradeImpactWithDefaults instantiates a new FirmwareFabricUpgradeImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareFabricUpgradeImpactWithDefaults() *FirmwareFabricUpgradeImpact {
	this := FirmwareFabricUpgradeImpact{}
	var classId string = "firmware.FabricUpgradeImpact"
	this.ClassId = classId
	var objectType string = "firmware.FabricUpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareFabricUpgradeImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareFabricUpgradeImpact) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareFabricUpgradeImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.FabricUpgradeImpact" of the ClassId field.
func (o *FirmwareFabricUpgradeImpact) GetDefaultClassId() interface{} {
	return "firmware.FabricUpgradeImpact"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareFabricUpgradeImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareFabricUpgradeImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareFabricUpgradeImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.FabricUpgradeImpact" of the ObjectType field.
func (o *FirmwareFabricUpgradeImpact) GetDefaultObjectType() interface{} {
	return "firmware.FabricUpgradeImpact"
}

// GetImpactDetail returns the ImpactDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareFabricUpgradeImpact) GetImpactDetail() []FirmwareComponentImpact {
	if o == nil {
		var ret []FirmwareComponentImpact
		return ret
	}
	return o.ImpactDetail
}

// GetImpactDetailOk returns a tuple with the ImpactDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareFabricUpgradeImpact) GetImpactDetailOk() ([]FirmwareComponentImpact, bool) {
	if o == nil || IsNil(o.ImpactDetail) {
		return nil, false
	}
	return o.ImpactDetail, true
}

// HasImpactDetail returns a boolean if a field has been set.
func (o *FirmwareFabricUpgradeImpact) HasImpactDetail() bool {
	if o != nil && !IsNil(o.ImpactDetail) {
		return true
	}

	return false
}

// SetImpactDetail gets a reference to the given []FirmwareComponentImpact and assigns it to the ImpactDetail field.
func (o *FirmwareFabricUpgradeImpact) SetImpactDetail(v []FirmwareComponentImpact) {
	o.ImpactDetail = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *FirmwareFabricUpgradeImpact) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFabricUpgradeImpact) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *FirmwareFabricUpgradeImpact) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *FirmwareFabricUpgradeImpact) SetSerial(v string) {
	o.Serial = &v
}

func (o FirmwareFabricUpgradeImpact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareFabricUpgradeImpact) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Serial) {
		toSerialize["Serial"] = o.Serial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareFabricUpgradeImpact) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareFabricUpgradeImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                    `json:"ObjectType"`
		ImpactDetail []FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
		// Details for the Fabric Interconnect that will be impacted by the upgrade.
		Serial *string `json:"Serial,omitempty"`
	}

	varFirmwareFabricUpgradeImpactWithoutEmbeddedStruct := FirmwareFabricUpgradeImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareFabricUpgradeImpactWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareFabricUpgradeImpact := _FirmwareFabricUpgradeImpact{}
		varFirmwareFabricUpgradeImpact.ClassId = varFirmwareFabricUpgradeImpactWithoutEmbeddedStruct.ClassId
		varFirmwareFabricUpgradeImpact.ObjectType = varFirmwareFabricUpgradeImpactWithoutEmbeddedStruct.ObjectType
		varFirmwareFabricUpgradeImpact.ImpactDetail = varFirmwareFabricUpgradeImpactWithoutEmbeddedStruct.ImpactDetail
		varFirmwareFabricUpgradeImpact.Serial = varFirmwareFabricUpgradeImpactWithoutEmbeddedStruct.Serial
		*o = FirmwareFabricUpgradeImpact(varFirmwareFabricUpgradeImpact)
	} else {
		return err
	}

	varFirmwareFabricUpgradeImpact := _FirmwareFabricUpgradeImpact{}

	err = json.Unmarshal(data, &varFirmwareFabricUpgradeImpact)
	if err == nil {
		o.FirmwareBaseImpact = varFirmwareFabricUpgradeImpact.FirmwareBaseImpact
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ImpactDetail")
		delete(additionalProperties, "Serial")

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

type NullableFirmwareFabricUpgradeImpact struct {
	value *FirmwareFabricUpgradeImpact
	isSet bool
}

func (v NullableFirmwareFabricUpgradeImpact) Get() *FirmwareFabricUpgradeImpact {
	return v.value
}

func (v *NullableFirmwareFabricUpgradeImpact) Set(val *FirmwareFabricUpgradeImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareFabricUpgradeImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareFabricUpgradeImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareFabricUpgradeImpact(val *FirmwareFabricUpgradeImpact) *NullableFirmwareFabricUpgradeImpact {
	return &NullableFirmwareFabricUpgradeImpact{value: val, isSet: true}
}

func (v NullableFirmwareFabricUpgradeImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareFabricUpgradeImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
