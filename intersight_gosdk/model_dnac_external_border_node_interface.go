/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-17057
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

// checks if the DnacExternalBorderNodeInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacExternalBorderNodeInterface{}

// DnacExternalBorderNodeInterface The external border node interface serves as the gateway between a network and external entities.
type DnacExternalBorderNodeInterface struct {
	DnacInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrator status for external border node interface.
	AdminStatus *string `json:"AdminStatus,omitempty"`
	// External border node's id.
	ExternalBorderNodeId *string `json:"ExternalBorderNodeId,omitempty"`
	// The Moid for the interface in the external border node.
	InterfaceId *string `json:"InterfaceId,omitempty"`
	// The name for the external border node's interface.
	InterfaceName *string `json:"InterfaceName,omitempty"`
	// Interface type for external border node interface.
	InterfaceType *string `json:"InterfaceType,omitempty"`
	// Port type for external border node interface.
	PortType             *string `json:"PortType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacExternalBorderNodeInterface DnacExternalBorderNodeInterface

// NewDnacExternalBorderNodeInterface instantiates a new DnacExternalBorderNodeInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacExternalBorderNodeInterface(classId string, objectType string) *DnacExternalBorderNodeInterface {
	this := DnacExternalBorderNodeInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewDnacExternalBorderNodeInterfaceWithDefaults instantiates a new DnacExternalBorderNodeInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacExternalBorderNodeInterfaceWithDefaults() *DnacExternalBorderNodeInterface {
	this := DnacExternalBorderNodeInterface{}
	var classId string = "dnac.ExternalBorderNodeInterface"
	this.ClassId = classId
	var objectType string = "dnac.ExternalBorderNodeInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *DnacExternalBorderNodeInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DnacExternalBorderNodeInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *DnacExternalBorderNodeInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DnacExternalBorderNodeInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminStatus returns the AdminStatus field value if set, zero value otherwise.
func (o *DnacExternalBorderNodeInterface) GetAdminStatus() string {
	if o == nil || IsNil(o.AdminStatus) {
		var ret string
		return ret
	}
	return *o.AdminStatus
}

// GetAdminStatusOk returns a tuple with the AdminStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetAdminStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AdminStatus) {
		return nil, false
	}
	return o.AdminStatus, true
}

// HasAdminStatus returns a boolean if a field has been set.
func (o *DnacExternalBorderNodeInterface) HasAdminStatus() bool {
	if o != nil && !IsNil(o.AdminStatus) {
		return true
	}

	return false
}

// SetAdminStatus gets a reference to the given string and assigns it to the AdminStatus field.
func (o *DnacExternalBorderNodeInterface) SetAdminStatus(v string) {
	o.AdminStatus = &v
}

// GetExternalBorderNodeId returns the ExternalBorderNodeId field value if set, zero value otherwise.
func (o *DnacExternalBorderNodeInterface) GetExternalBorderNodeId() string {
	if o == nil || IsNil(o.ExternalBorderNodeId) {
		var ret string
		return ret
	}
	return *o.ExternalBorderNodeId
}

// GetExternalBorderNodeIdOk returns a tuple with the ExternalBorderNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetExternalBorderNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalBorderNodeId) {
		return nil, false
	}
	return o.ExternalBorderNodeId, true
}

// HasExternalBorderNodeId returns a boolean if a field has been set.
func (o *DnacExternalBorderNodeInterface) HasExternalBorderNodeId() bool {
	if o != nil && !IsNil(o.ExternalBorderNodeId) {
		return true
	}

	return false
}

// SetExternalBorderNodeId gets a reference to the given string and assigns it to the ExternalBorderNodeId field.
func (o *DnacExternalBorderNodeInterface) SetExternalBorderNodeId(v string) {
	o.ExternalBorderNodeId = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *DnacExternalBorderNodeInterface) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *DnacExternalBorderNodeInterface) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *DnacExternalBorderNodeInterface) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *DnacExternalBorderNodeInterface) GetInterfaceName() string {
	if o == nil || IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetInterfaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *DnacExternalBorderNodeInterface) HasInterfaceName() bool {
	if o != nil && !IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *DnacExternalBorderNodeInterface) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *DnacExternalBorderNodeInterface) GetInterfaceType() string {
	if o == nil || IsNil(o.InterfaceType) {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *DnacExternalBorderNodeInterface) HasInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *DnacExternalBorderNodeInterface) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *DnacExternalBorderNodeInterface) GetPortType() string {
	if o == nil || IsNil(o.PortType) {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacExternalBorderNodeInterface) GetPortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PortType) {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *DnacExternalBorderNodeInterface) HasPortType() bool {
	if o != nil && !IsNil(o.PortType) {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *DnacExternalBorderNodeInterface) SetPortType(v string) {
	o.PortType = &v
}

func (o DnacExternalBorderNodeInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacExternalBorderNodeInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDnacInventoryEntity, errDnacInventoryEntity := json.Marshal(o.DnacInventoryEntity)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	errDnacInventoryEntity = json.Unmarshal([]byte(serializedDnacInventoryEntity), &toSerialize)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdminStatus) {
		toSerialize["AdminStatus"] = o.AdminStatus
	}
	if !IsNil(o.ExternalBorderNodeId) {
		toSerialize["ExternalBorderNodeId"] = o.ExternalBorderNodeId
	}
	if !IsNil(o.InterfaceId) {
		toSerialize["InterfaceId"] = o.InterfaceId
	}
	if !IsNil(o.InterfaceName) {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["InterfaceType"] = o.InterfaceType
	}
	if !IsNil(o.PortType) {
		toSerialize["PortType"] = o.PortType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacExternalBorderNodeInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type DnacExternalBorderNodeInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Administrator status for external border node interface.
		AdminStatus *string `json:"AdminStatus,omitempty"`
		// External border node's id.
		ExternalBorderNodeId *string `json:"ExternalBorderNodeId,omitempty"`
		// The Moid for the interface in the external border node.
		InterfaceId *string `json:"InterfaceId,omitempty"`
		// The name for the external border node's interface.
		InterfaceName *string `json:"InterfaceName,omitempty"`
		// Interface type for external border node interface.
		InterfaceType *string `json:"InterfaceType,omitempty"`
		// Port type for external border node interface.
		PortType *string `json:"PortType,omitempty"`
	}

	varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct := DnacExternalBorderNodeInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varDnacExternalBorderNodeInterface := _DnacExternalBorderNodeInterface{}
		varDnacExternalBorderNodeInterface.ClassId = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.ClassId
		varDnacExternalBorderNodeInterface.ObjectType = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.ObjectType
		varDnacExternalBorderNodeInterface.AdminStatus = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.AdminStatus
		varDnacExternalBorderNodeInterface.ExternalBorderNodeId = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.ExternalBorderNodeId
		varDnacExternalBorderNodeInterface.InterfaceId = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.InterfaceId
		varDnacExternalBorderNodeInterface.InterfaceName = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.InterfaceName
		varDnacExternalBorderNodeInterface.InterfaceType = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.InterfaceType
		varDnacExternalBorderNodeInterface.PortType = varDnacExternalBorderNodeInterfaceWithoutEmbeddedStruct.PortType
		*o = DnacExternalBorderNodeInterface(varDnacExternalBorderNodeInterface)
	} else {
		return err
	}

	varDnacExternalBorderNodeInterface := _DnacExternalBorderNodeInterface{}

	err = json.Unmarshal(data, &varDnacExternalBorderNodeInterface)
	if err == nil {
		o.DnacInventoryEntity = varDnacExternalBorderNodeInterface.DnacInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminStatus")
		delete(additionalProperties, "ExternalBorderNodeId")
		delete(additionalProperties, "InterfaceId")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "InterfaceType")
		delete(additionalProperties, "PortType")

		// remove fields from embedded structs
		reflectDnacInventoryEntity := reflect.ValueOf(o.DnacInventoryEntity)
		for i := 0; i < reflectDnacInventoryEntity.Type().NumField(); i++ {
			t := reflectDnacInventoryEntity.Type().Field(i)

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

type NullableDnacExternalBorderNodeInterface struct {
	value *DnacExternalBorderNodeInterface
	isSet bool
}

func (v NullableDnacExternalBorderNodeInterface) Get() *DnacExternalBorderNodeInterface {
	return v.value
}

func (v *NullableDnacExternalBorderNodeInterface) Set(val *DnacExternalBorderNodeInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacExternalBorderNodeInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacExternalBorderNodeInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacExternalBorderNodeInterface(val *DnacExternalBorderNodeInterface) *NullableDnacExternalBorderNodeInterface {
	return &NullableDnacExternalBorderNodeInterface{value: val, isSet: true}
}

func (v NullableDnacExternalBorderNodeInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacExternalBorderNodeInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}