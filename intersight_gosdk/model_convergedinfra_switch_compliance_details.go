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

// checks if the ConvergedinfraSwitchComplianceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvergedinfraSwitchComplianceDetails{}

// ConvergedinfraSwitchComplianceDetails The compliance details for a switch (Fabric Interconnect/Nexus/MDS) which is part of a converged infrastructure pod.
type ConvergedinfraSwitchComplianceDetails struct {
	ConvergedinfraBaseComplianceDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The firmware of the switch as received from inventory.
	Firmware *string `json:"Firmware,omitempty"`
	// The model information of the switch.
	Model *string `json:"Model,omitempty"`
	// The type of switch component. It must be set to either Fabric Interconnect, Nexus or MDS. * `FabricInterconnect` - The default Switch type of UCSM and IMM mode devices. * `NexusDevice` - Switch type of Nexus devices. * `MDSDevice` - Switch type of Nexus MDS devices.
	Type          *string                                             `json:"Type,omitempty"`
	PodCompliance NullableConvergedinfraPodComplianceInfoRelationship `json:"PodCompliance,omitempty"`
	// An array of relationships to convergedinfraStorageComplianceDetails resources.
	StorageCompliances   []ConvergedinfraStorageComplianceDetailsRelationship `json:"StorageCompliances,omitempty"`
	Switch               NullableNetworkElementSummaryRelationship            `json:"Switch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraSwitchComplianceDetails ConvergedinfraSwitchComplianceDetails

// NewConvergedinfraSwitchComplianceDetails instantiates a new ConvergedinfraSwitchComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraSwitchComplianceDetails(classId string, objectType string) *ConvergedinfraSwitchComplianceDetails {
	this := ConvergedinfraSwitchComplianceDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraSwitchComplianceDetailsWithDefaults instantiates a new ConvergedinfraSwitchComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraSwitchComplianceDetailsWithDefaults() *ConvergedinfraSwitchComplianceDetails {
	this := ConvergedinfraSwitchComplianceDetails{}
	var classId string = "convergedinfra.SwitchComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.SwitchComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraSwitchComplianceDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraSwitchComplianceDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "convergedinfra.SwitchComplianceDetails" of the ClassId field.
func (o *ConvergedinfraSwitchComplianceDetails) GetDefaultClassId() interface{} {
	return "convergedinfra.SwitchComplianceDetails"
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraSwitchComplianceDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraSwitchComplianceDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "convergedinfra.SwitchComplianceDetails" of the ObjectType field.
func (o *ConvergedinfraSwitchComplianceDetails) GetDefaultObjectType() interface{} {
	return "convergedinfra.SwitchComplianceDetails"
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetails) GetFirmware() string {
	if o == nil || IsNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetails) GetFirmwareOk() (*string, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetails) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *ConvergedinfraSwitchComplianceDetails) SetFirmware(v string) {
	o.Firmware = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetails) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetails) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetails) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConvergedinfraSwitchComplianceDetails) SetModel(v string) {
	o.Model = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConvergedinfraSwitchComplianceDetails) SetType(v string) {
	o.Type = &v
}

// GetPodCompliance returns the PodCompliance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraSwitchComplianceDetails) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship {
	if o == nil || IsNil(o.PodCompliance.Get()) {
		var ret ConvergedinfraPodComplianceInfoRelationship
		return ret
	}
	return *o.PodCompliance.Get()
}

// GetPodComplianceOk returns a tuple with the PodCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraSwitchComplianceDetails) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PodCompliance.Get(), o.PodCompliance.IsSet()
}

// HasPodCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetails) HasPodCompliance() bool {
	if o != nil && o.PodCompliance.IsSet() {
		return true
	}

	return false
}

// SetPodCompliance gets a reference to the given NullableConvergedinfraPodComplianceInfoRelationship and assigns it to the PodCompliance field.
func (o *ConvergedinfraSwitchComplianceDetails) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship) {
	o.PodCompliance.Set(&v)
}

// SetPodComplianceNil sets the value for PodCompliance to be an explicit nil
func (o *ConvergedinfraSwitchComplianceDetails) SetPodComplianceNil() {
	o.PodCompliance.Set(nil)
}

// UnsetPodCompliance ensures that no value is present for PodCompliance, not even an explicit nil
func (o *ConvergedinfraSwitchComplianceDetails) UnsetPodCompliance() {
	o.PodCompliance.Unset()
}

// GetStorageCompliances returns the StorageCompliances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraSwitchComplianceDetails) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship {
	if o == nil {
		var ret []ConvergedinfraStorageComplianceDetailsRelationship
		return ret
	}
	return o.StorageCompliances
}

// GetStorageCompliancesOk returns a tuple with the StorageCompliances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraSwitchComplianceDetails) GetStorageCompliancesOk() ([]ConvergedinfraStorageComplianceDetailsRelationship, bool) {
	if o == nil || IsNil(o.StorageCompliances) {
		return nil, false
	}
	return o.StorageCompliances, true
}

// HasStorageCompliances returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetails) HasStorageCompliances() bool {
	if o != nil && !IsNil(o.StorageCompliances) {
		return true
	}

	return false
}

// SetStorageCompliances gets a reference to the given []ConvergedinfraStorageComplianceDetailsRelationship and assigns it to the StorageCompliances field.
func (o *ConvergedinfraSwitchComplianceDetails) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship) {
	o.StorageCompliances = v
}

// GetSwitch returns the Switch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraSwitchComplianceDetails) GetSwitch() NetworkElementSummaryRelationship {
	if o == nil || IsNil(o.Switch.Get()) {
		var ret NetworkElementSummaryRelationship
		return ret
	}
	return *o.Switch.Get()
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraSwitchComplianceDetails) GetSwitchOk() (*NetworkElementSummaryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Switch.Get(), o.Switch.IsSet()
}

// HasSwitch returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetails) HasSwitch() bool {
	if o != nil && o.Switch.IsSet() {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given NullableNetworkElementSummaryRelationship and assigns it to the Switch field.
func (o *ConvergedinfraSwitchComplianceDetails) SetSwitch(v NetworkElementSummaryRelationship) {
	o.Switch.Set(&v)
}

// SetSwitchNil sets the value for Switch to be an explicit nil
func (o *ConvergedinfraSwitchComplianceDetails) SetSwitchNil() {
	o.Switch.Set(nil)
}

// UnsetSwitch ensures that no value is present for Switch, not even an explicit nil
func (o *ConvergedinfraSwitchComplianceDetails) UnsetSwitch() {
	o.Switch.Unset()
}

func (o ConvergedinfraSwitchComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvergedinfraSwitchComplianceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConvergedinfraBaseComplianceDetails, errConvergedinfraBaseComplianceDetails := json.Marshal(o.ConvergedinfraBaseComplianceDetails)
	if errConvergedinfraBaseComplianceDetails != nil {
		return map[string]interface{}{}, errConvergedinfraBaseComplianceDetails
	}
	errConvergedinfraBaseComplianceDetails = json.Unmarshal([]byte(serializedConvergedinfraBaseComplianceDetails), &toSerialize)
	if errConvergedinfraBaseComplianceDetails != nil {
		return map[string]interface{}{}, errConvergedinfraBaseComplianceDetails
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Firmware) {
		toSerialize["Firmware"] = o.Firmware
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.PodCompliance.IsSet() {
		toSerialize["PodCompliance"] = o.PodCompliance.Get()
	}
	if o.StorageCompliances != nil {
		toSerialize["StorageCompliances"] = o.StorageCompliances
	}
	if o.Switch.IsSet() {
		toSerialize["Switch"] = o.Switch.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConvergedinfraSwitchComplianceDetails) UnmarshalJSON(data []byte) (err error) {
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
	type ConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The firmware of the switch as received from inventory.
		Firmware *string `json:"Firmware,omitempty"`
		// The model information of the switch.
		Model *string `json:"Model,omitempty"`
		// The type of switch component. It must be set to either Fabric Interconnect, Nexus or MDS. * `FabricInterconnect` - The default Switch type of UCSM and IMM mode devices. * `NexusDevice` - Switch type of Nexus devices. * `MDSDevice` - Switch type of Nexus MDS devices.
		Type          *string                                             `json:"Type,omitempty"`
		PodCompliance NullableConvergedinfraPodComplianceInfoRelationship `json:"PodCompliance,omitempty"`
		// An array of relationships to convergedinfraStorageComplianceDetails resources.
		StorageCompliances []ConvergedinfraStorageComplianceDetailsRelationship `json:"StorageCompliances,omitempty"`
		Switch             NullableNetworkElementSummaryRelationship            `json:"Switch,omitempty"`
	}

	varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct := ConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraSwitchComplianceDetails := _ConvergedinfraSwitchComplianceDetails{}
		varConvergedinfraSwitchComplianceDetails.ClassId = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.ClassId
		varConvergedinfraSwitchComplianceDetails.ObjectType = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.ObjectType
		varConvergedinfraSwitchComplianceDetails.Firmware = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.Firmware
		varConvergedinfraSwitchComplianceDetails.Model = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.Model
		varConvergedinfraSwitchComplianceDetails.Type = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.Type
		varConvergedinfraSwitchComplianceDetails.PodCompliance = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.PodCompliance
		varConvergedinfraSwitchComplianceDetails.StorageCompliances = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.StorageCompliances
		varConvergedinfraSwitchComplianceDetails.Switch = varConvergedinfraSwitchComplianceDetailsWithoutEmbeddedStruct.Switch
		*o = ConvergedinfraSwitchComplianceDetails(varConvergedinfraSwitchComplianceDetails)
	} else {
		return err
	}

	varConvergedinfraSwitchComplianceDetails := _ConvergedinfraSwitchComplianceDetails{}

	err = json.Unmarshal(data, &varConvergedinfraSwitchComplianceDetails)
	if err == nil {
		o.ConvergedinfraBaseComplianceDetails = varConvergedinfraSwitchComplianceDetails.ConvergedinfraBaseComplianceDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Firmware")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "PodCompliance")
		delete(additionalProperties, "StorageCompliances")
		delete(additionalProperties, "Switch")

		// remove fields from embedded structs
		reflectConvergedinfraBaseComplianceDetails := reflect.ValueOf(o.ConvergedinfraBaseComplianceDetails)
		for i := 0; i < reflectConvergedinfraBaseComplianceDetails.Type().NumField(); i++ {
			t := reflectConvergedinfraBaseComplianceDetails.Type().Field(i)

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

type NullableConvergedinfraSwitchComplianceDetails struct {
	value *ConvergedinfraSwitchComplianceDetails
	isSet bool
}

func (v NullableConvergedinfraSwitchComplianceDetails) Get() *ConvergedinfraSwitchComplianceDetails {
	return v.value
}

func (v *NullableConvergedinfraSwitchComplianceDetails) Set(val *ConvergedinfraSwitchComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraSwitchComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraSwitchComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraSwitchComplianceDetails(val *ConvergedinfraSwitchComplianceDetails) *NullableConvergedinfraSwitchComplianceDetails {
	return &NullableConvergedinfraSwitchComplianceDetails{value: val, isSet: true}
}

func (v NullableConvergedinfraSwitchComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraSwitchComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
