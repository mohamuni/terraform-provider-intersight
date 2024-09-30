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

// checks if the CapabilityDrivesEndpointDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityDrivesEndpointDescriptor{}

// CapabilityDrivesEndpointDescriptor Descriptor that uniquely identifies a drive.
type CapabilityDrivesEndpointDescriptor struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field is to provide alias model of the drive.
	AliasModel *string `json:"AliasModel,omitempty"`
	// This field is to provide description of the drive.
	Description *string `json:"Description,omitempty"`
	// This field is to provide model of the drive.
	Model *string `json:"Model,omitempty"`
	// This field is to provide partNumber of the drive.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field is to provide pid of the drive.
	Pid                    *string  `json:"Pid,omitempty"`
	SupportedPlatformsPids []string `json:"SupportedPlatformsPids,omitempty"`
	// This field is to provide vendor of the drive.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityDrivesEndpointDescriptor CapabilityDrivesEndpointDescriptor

// NewCapabilityDrivesEndpointDescriptor instantiates a new CapabilityDrivesEndpointDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityDrivesEndpointDescriptor(classId string, objectType string) *CapabilityDrivesEndpointDescriptor {
	this := CapabilityDrivesEndpointDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityDrivesEndpointDescriptorWithDefaults instantiates a new CapabilityDrivesEndpointDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityDrivesEndpointDescriptorWithDefaults() *CapabilityDrivesEndpointDescriptor {
	this := CapabilityDrivesEndpointDescriptor{}
	var classId string = "capability.DrivesEndpointDescriptor"
	this.ClassId = classId
	var objectType string = "capability.DrivesEndpointDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityDrivesEndpointDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityDrivesEndpointDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.DrivesEndpointDescriptor" of the ClassId field.
func (o *CapabilityDrivesEndpointDescriptor) GetDefaultClassId() interface{} {
	return "capability.DrivesEndpointDescriptor"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityDrivesEndpointDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityDrivesEndpointDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.DrivesEndpointDescriptor" of the ObjectType field.
func (o *CapabilityDrivesEndpointDescriptor) GetDefaultObjectType() interface{} {
	return "capability.DrivesEndpointDescriptor"
}

// GetAliasModel returns the AliasModel field value if set, zero value otherwise.
func (o *CapabilityDrivesEndpointDescriptor) GetAliasModel() string {
	if o == nil || IsNil(o.AliasModel) {
		var ret string
		return ret
	}
	return *o.AliasModel
}

// GetAliasModelOk returns a tuple with the AliasModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetAliasModelOk() (*string, bool) {
	if o == nil || IsNil(o.AliasModel) {
		return nil, false
	}
	return o.AliasModel, true
}

// HasAliasModel returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasAliasModel() bool {
	if o != nil && !IsNil(o.AliasModel) {
		return true
	}

	return false
}

// SetAliasModel gets a reference to the given string and assigns it to the AliasModel field.
func (o *CapabilityDrivesEndpointDescriptor) SetAliasModel(v string) {
	o.AliasModel = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityDrivesEndpointDescriptor) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityDrivesEndpointDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityDrivesEndpointDescriptor) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityDrivesEndpointDescriptor) SetModel(v string) {
	o.Model = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *CapabilityDrivesEndpointDescriptor) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *CapabilityDrivesEndpointDescriptor) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *CapabilityDrivesEndpointDescriptor) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *CapabilityDrivesEndpointDescriptor) SetPid(v string) {
	o.Pid = &v
}

// GetSupportedPlatformsPids returns the SupportedPlatformsPids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDrivesEndpointDescriptor) GetSupportedPlatformsPids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedPlatformsPids
}

// GetSupportedPlatformsPidsOk returns a tuple with the SupportedPlatformsPids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDrivesEndpointDescriptor) GetSupportedPlatformsPidsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedPlatformsPids) {
		return nil, false
	}
	return o.SupportedPlatformsPids, true
}

// HasSupportedPlatformsPids returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasSupportedPlatformsPids() bool {
	if o != nil && !IsNil(o.SupportedPlatformsPids) {
		return true
	}

	return false
}

// SetSupportedPlatformsPids gets a reference to the given []string and assigns it to the SupportedPlatformsPids field.
func (o *CapabilityDrivesEndpointDescriptor) SetSupportedPlatformsPids(v []string) {
	o.SupportedPlatformsPids = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityDrivesEndpointDescriptor) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDrivesEndpointDescriptor) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityDrivesEndpointDescriptor) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityDrivesEndpointDescriptor) SetVendor(v string) {
	o.Vendor = &v
}

func (o CapabilityDrivesEndpointDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityDrivesEndpointDescriptor) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AliasModel) {
		toSerialize["AliasModel"] = o.AliasModel
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.Pid) {
		toSerialize["Pid"] = o.Pid
	}
	if o.SupportedPlatformsPids != nil {
		toSerialize["SupportedPlatformsPids"] = o.SupportedPlatformsPids
	}
	if !IsNil(o.Vendor) {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityDrivesEndpointDescriptor) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field is to provide alias model of the drive.
		AliasModel *string `json:"AliasModel,omitempty"`
		// This field is to provide description of the drive.
		Description *string `json:"Description,omitempty"`
		// This field is to provide model of the drive.
		Model *string `json:"Model,omitempty"`
		// This field is to provide partNumber of the drive.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field is to provide pid of the drive.
		Pid                    *string  `json:"Pid,omitempty"`
		SupportedPlatformsPids []string `json:"SupportedPlatformsPids,omitempty"`
		// This field is to provide vendor of the drive.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct := CapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityDrivesEndpointDescriptor := _CapabilityDrivesEndpointDescriptor{}
		varCapabilityDrivesEndpointDescriptor.ClassId = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilityDrivesEndpointDescriptor.ObjectType = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilityDrivesEndpointDescriptor.AliasModel = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.AliasModel
		varCapabilityDrivesEndpointDescriptor.Description = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.Description
		varCapabilityDrivesEndpointDescriptor.Model = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.Model
		varCapabilityDrivesEndpointDescriptor.PartNumber = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.PartNumber
		varCapabilityDrivesEndpointDescriptor.Pid = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.Pid
		varCapabilityDrivesEndpointDescriptor.SupportedPlatformsPids = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.SupportedPlatformsPids
		varCapabilityDrivesEndpointDescriptor.Vendor = varCapabilityDrivesEndpointDescriptorWithoutEmbeddedStruct.Vendor
		*o = CapabilityDrivesEndpointDescriptor(varCapabilityDrivesEndpointDescriptor)
	} else {
		return err
	}

	varCapabilityDrivesEndpointDescriptor := _CapabilityDrivesEndpointDescriptor{}

	err = json.Unmarshal(data, &varCapabilityDrivesEndpointDescriptor)
	if err == nil {
		o.MoBaseMo = varCapabilityDrivesEndpointDescriptor.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AliasModel")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SupportedPlatformsPids")
		delete(additionalProperties, "Vendor")

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

type NullableCapabilityDrivesEndpointDescriptor struct {
	value *CapabilityDrivesEndpointDescriptor
	isSet bool
}

func (v NullableCapabilityDrivesEndpointDescriptor) Get() *CapabilityDrivesEndpointDescriptor {
	return v.value
}

func (v *NullableCapabilityDrivesEndpointDescriptor) Set(val *CapabilityDrivesEndpointDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityDrivesEndpointDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityDrivesEndpointDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityDrivesEndpointDescriptor(val *CapabilityDrivesEndpointDescriptor) *NullableCapabilityDrivesEndpointDescriptor {
	return &NullableCapabilityDrivesEndpointDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityDrivesEndpointDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityDrivesEndpointDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
