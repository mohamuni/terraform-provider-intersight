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

// checks if the ServerServerAssignTypeSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerServerAssignTypeSlot{}

// ServerServerAssignTypeSlot Server profile pre-assignment by slot.
type ServerServerAssignTypeSlot struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Chassis-id of the slot that would be assigned to this pre-assigned server profile.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// Domain name of the Fabric Interconnect to which the chassis is or to be connected. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters.
	DomainName *string `json:"DomainName,omitempty" validate:"regexp=^[a-zA-Z0-9_\\\\-]{0,30}$"`
	// Slot-id of the server that would be assigned to this pre-assigned server profile.
	SlotId               *int64 `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerServerAssignTypeSlot ServerServerAssignTypeSlot

// NewServerServerAssignTypeSlot instantiates a new ServerServerAssignTypeSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerServerAssignTypeSlot(classId string, objectType string) *ServerServerAssignTypeSlot {
	this := ServerServerAssignTypeSlot{}
	this.ClassId = classId
	this.ObjectType = objectType
	var chassisId int64 = 0
	this.ChassisId = &chassisId
	var slotId int64 = 0
	this.SlotId = &slotId
	return &this
}

// NewServerServerAssignTypeSlotWithDefaults instantiates a new ServerServerAssignTypeSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerServerAssignTypeSlotWithDefaults() *ServerServerAssignTypeSlot {
	this := ServerServerAssignTypeSlot{}
	var classId string = "server.ServerAssignTypeSlot"
	this.ClassId = classId
	var objectType string = "server.ServerAssignTypeSlot"
	this.ObjectType = objectType
	var chassisId int64 = 0
	this.ChassisId = &chassisId
	var slotId int64 = 0
	this.SlotId = &slotId
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerServerAssignTypeSlot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerServerAssignTypeSlot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "server.ServerAssignTypeSlot" of the ClassId field.
func (o *ServerServerAssignTypeSlot) GetDefaultClassId() interface{} {
	return "server.ServerAssignTypeSlot"
}

// GetObjectType returns the ObjectType field value
func (o *ServerServerAssignTypeSlot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerServerAssignTypeSlot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "server.ServerAssignTypeSlot" of the ObjectType field.
func (o *ServerServerAssignTypeSlot) GetDefaultObjectType() interface{} {
	return "server.ServerAssignTypeSlot"
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ServerServerAssignTypeSlot) GetChassisId() int64 {
	if o == nil || IsNil(o.ChassisId) {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlot) GetChassisIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ServerServerAssignTypeSlot) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *ServerServerAssignTypeSlot) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *ServerServerAssignTypeSlot) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlot) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *ServerServerAssignTypeSlot) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *ServerServerAssignTypeSlot) SetDomainName(v string) {
	o.DomainName = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ServerServerAssignTypeSlot) GetSlotId() int64 {
	if o == nil || IsNil(o.SlotId) {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlot) GetSlotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SlotId) {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ServerServerAssignTypeSlot) HasSlotId() bool {
	if o != nil && !IsNil(o.SlotId) {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ServerServerAssignTypeSlot) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o ServerServerAssignTypeSlot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerServerAssignTypeSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ChassisId) {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if !IsNil(o.DomainName) {
		toSerialize["DomainName"] = o.DomainName
	}
	if !IsNil(o.SlotId) {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerServerAssignTypeSlot) UnmarshalJSON(data []byte) (err error) {
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
	type ServerServerAssignTypeSlotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Chassis-id of the slot that would be assigned to this pre-assigned server profile.
		ChassisId *int64 `json:"ChassisId,omitempty"`
		// Domain name of the Fabric Interconnect to which the chassis is or to be connected. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters.
		DomainName *string `json:"DomainName,omitempty" validate:"regexp=^[a-zA-Z0-9_\\\\-]{0,30}$"`
		// Slot-id of the server that would be assigned to this pre-assigned server profile.
		SlotId *int64 `json:"SlotId,omitempty"`
	}

	varServerServerAssignTypeSlotWithoutEmbeddedStruct := ServerServerAssignTypeSlotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServerServerAssignTypeSlotWithoutEmbeddedStruct)
	if err == nil {
		varServerServerAssignTypeSlot := _ServerServerAssignTypeSlot{}
		varServerServerAssignTypeSlot.ClassId = varServerServerAssignTypeSlotWithoutEmbeddedStruct.ClassId
		varServerServerAssignTypeSlot.ObjectType = varServerServerAssignTypeSlotWithoutEmbeddedStruct.ObjectType
		varServerServerAssignTypeSlot.ChassisId = varServerServerAssignTypeSlotWithoutEmbeddedStruct.ChassisId
		varServerServerAssignTypeSlot.DomainName = varServerServerAssignTypeSlotWithoutEmbeddedStruct.DomainName
		varServerServerAssignTypeSlot.SlotId = varServerServerAssignTypeSlotWithoutEmbeddedStruct.SlotId
		*o = ServerServerAssignTypeSlot(varServerServerAssignTypeSlot)
	} else {
		return err
	}

	varServerServerAssignTypeSlot := _ServerServerAssignTypeSlot{}

	err = json.Unmarshal(data, &varServerServerAssignTypeSlot)
	if err == nil {
		o.MoBaseComplexType = varServerServerAssignTypeSlot.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "SlotId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableServerServerAssignTypeSlot struct {
	value *ServerServerAssignTypeSlot
	isSet bool
}

func (v NullableServerServerAssignTypeSlot) Get() *ServerServerAssignTypeSlot {
	return v.value
}

func (v *NullableServerServerAssignTypeSlot) Set(val *ServerServerAssignTypeSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableServerServerAssignTypeSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableServerServerAssignTypeSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerServerAssignTypeSlot(val *ServerServerAssignTypeSlot) *NullableServerServerAssignTypeSlot {
	return &NullableServerServerAssignTypeSlot{value: val, isSet: true}
}

func (v NullableServerServerAssignTypeSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerServerAssignTypeSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
