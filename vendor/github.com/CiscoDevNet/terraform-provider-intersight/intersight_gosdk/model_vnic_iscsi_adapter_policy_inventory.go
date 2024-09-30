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

// checks if the VnicIscsiAdapterPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicIscsiAdapterPolicyInventory{}

// VnicIscsiAdapterPolicyInventory Set of iSCSI properties that govern the host-side behavior of the adapter.
type VnicIscsiAdapterPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable.
	ConnectionTimeOut *int64 `json:"ConnectionTimeOut,omitempty"`
	// The number of seconds to wait before the initiator assumes that the DHCP server is unavailable.
	DhcpTimeout *int64 `json:"DhcpTimeout,omitempty"`
	// The number of times to retry the connection in case of a failure during iSCSI LUN discovery.
	LunBusyRetryCount    *int64                       `json:"LunBusyRetryCount,omitempty"`
	TargetMo             NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicIscsiAdapterPolicyInventory VnicIscsiAdapterPolicyInventory

// NewVnicIscsiAdapterPolicyInventory instantiates a new VnicIscsiAdapterPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiAdapterPolicyInventory(classId string, objectType string) *VnicIscsiAdapterPolicyInventory {
	this := VnicIscsiAdapterPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicIscsiAdapterPolicyInventoryWithDefaults instantiates a new VnicIscsiAdapterPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiAdapterPolicyInventoryWithDefaults() *VnicIscsiAdapterPolicyInventory {
	this := VnicIscsiAdapterPolicyInventory{}
	var classId string = "vnic.IscsiAdapterPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.IscsiAdapterPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiAdapterPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiAdapterPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.IscsiAdapterPolicyInventory" of the ClassId field.
func (o *VnicIscsiAdapterPolicyInventory) GetDefaultClassId() interface{} {
	return "vnic.IscsiAdapterPolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiAdapterPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiAdapterPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.IscsiAdapterPolicyInventory" of the ObjectType field.
func (o *VnicIscsiAdapterPolicyInventory) GetDefaultObjectType() interface{} {
	return "vnic.IscsiAdapterPolicyInventory"
}

// GetConnectionTimeOut returns the ConnectionTimeOut field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventory) GetConnectionTimeOut() int64 {
	if o == nil || IsNil(o.ConnectionTimeOut) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeOut
}

// GetConnectionTimeOutOk returns a tuple with the ConnectionTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventory) GetConnectionTimeOutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeOut) {
		return nil, false
	}
	return o.ConnectionTimeOut, true
}

// HasConnectionTimeOut returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventory) HasConnectionTimeOut() bool {
	if o != nil && !IsNil(o.ConnectionTimeOut) {
		return true
	}

	return false
}

// SetConnectionTimeOut gets a reference to the given int64 and assigns it to the ConnectionTimeOut field.
func (o *VnicIscsiAdapterPolicyInventory) SetConnectionTimeOut(v int64) {
	o.ConnectionTimeOut = &v
}

// GetDhcpTimeout returns the DhcpTimeout field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventory) GetDhcpTimeout() int64 {
	if o == nil || IsNil(o.DhcpTimeout) {
		var ret int64
		return ret
	}
	return *o.DhcpTimeout
}

// GetDhcpTimeoutOk returns a tuple with the DhcpTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventory) GetDhcpTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.DhcpTimeout) {
		return nil, false
	}
	return o.DhcpTimeout, true
}

// HasDhcpTimeout returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventory) HasDhcpTimeout() bool {
	if o != nil && !IsNil(o.DhcpTimeout) {
		return true
	}

	return false
}

// SetDhcpTimeout gets a reference to the given int64 and assigns it to the DhcpTimeout field.
func (o *VnicIscsiAdapterPolicyInventory) SetDhcpTimeout(v int64) {
	o.DhcpTimeout = &v
}

// GetLunBusyRetryCount returns the LunBusyRetryCount field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventory) GetLunBusyRetryCount() int64 {
	if o == nil || IsNil(o.LunBusyRetryCount) {
		var ret int64
		return ret
	}
	return *o.LunBusyRetryCount
}

// GetLunBusyRetryCountOk returns a tuple with the LunBusyRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventory) GetLunBusyRetryCountOk() (*int64, bool) {
	if o == nil || IsNil(o.LunBusyRetryCount) {
		return nil, false
	}
	return o.LunBusyRetryCount, true
}

// HasLunBusyRetryCount returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventory) HasLunBusyRetryCount() bool {
	if o != nil && !IsNil(o.LunBusyRetryCount) {
		return true
	}

	return false
}

// SetLunBusyRetryCount gets a reference to the given int64 and assigns it to the LunBusyRetryCount field.
func (o *VnicIscsiAdapterPolicyInventory) SetLunBusyRetryCount(v int64) {
	o.LunBusyRetryCount = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiAdapterPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiAdapterPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicIscsiAdapterPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *VnicIscsiAdapterPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *VnicIscsiAdapterPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o VnicIscsiAdapterPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicIscsiAdapterPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConnectionTimeOut) {
		toSerialize["ConnectionTimeOut"] = o.ConnectionTimeOut
	}
	if !IsNil(o.DhcpTimeout) {
		toSerialize["DhcpTimeout"] = o.DhcpTimeout
	}
	if !IsNil(o.LunBusyRetryCount) {
		toSerialize["LunBusyRetryCount"] = o.LunBusyRetryCount
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicIscsiAdapterPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type VnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable.
		ConnectionTimeOut *int64 `json:"ConnectionTimeOut,omitempty"`
		// The number of seconds to wait before the initiator assumes that the DHCP server is unavailable.
		DhcpTimeout *int64 `json:"DhcpTimeout,omitempty"`
		// The number of times to retry the connection in case of a failure during iSCSI LUN discovery.
		LunBusyRetryCount *int64                       `json:"LunBusyRetryCount,omitempty"`
		TargetMo          NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct := VnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicIscsiAdapterPolicyInventory := _VnicIscsiAdapterPolicyInventory{}
		varVnicIscsiAdapterPolicyInventory.ClassId = varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicIscsiAdapterPolicyInventory.ObjectType = varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicIscsiAdapterPolicyInventory.ConnectionTimeOut = varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct.ConnectionTimeOut
		varVnicIscsiAdapterPolicyInventory.DhcpTimeout = varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct.DhcpTimeout
		varVnicIscsiAdapterPolicyInventory.LunBusyRetryCount = varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct.LunBusyRetryCount
		varVnicIscsiAdapterPolicyInventory.TargetMo = varVnicIscsiAdapterPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicIscsiAdapterPolicyInventory(varVnicIscsiAdapterPolicyInventory)
	} else {
		return err
	}

	varVnicIscsiAdapterPolicyInventory := _VnicIscsiAdapterPolicyInventory{}

	err = json.Unmarshal(data, &varVnicIscsiAdapterPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicIscsiAdapterPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionTimeOut")
		delete(additionalProperties, "DhcpTimeout")
		delete(additionalProperties, "LunBusyRetryCount")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableVnicIscsiAdapterPolicyInventory struct {
	value *VnicIscsiAdapterPolicyInventory
	isSet bool
}

func (v NullableVnicIscsiAdapterPolicyInventory) Get() *VnicIscsiAdapterPolicyInventory {
	return v.value
}

func (v *NullableVnicIscsiAdapterPolicyInventory) Set(val *VnicIscsiAdapterPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiAdapterPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiAdapterPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiAdapterPolicyInventory(val *VnicIscsiAdapterPolicyInventory) *NullableVnicIscsiAdapterPolicyInventory {
	return &NullableVnicIscsiAdapterPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicIscsiAdapterPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiAdapterPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
