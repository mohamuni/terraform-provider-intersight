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

// checks if the NetworkFeatureControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkFeatureControl{}

// NetworkFeatureControl List of features available on a switch along with the index and admin state. These features will allow the user to perform certain set of actions on the switch and get a view of the status of the sub-set feature names.
type NetworkFeatureControl struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The admin state of the feature.
	AdminState *string `json:"AdminState,omitempty"`
	// The number of instances of the feature.
	Instance *int64 `json:"Instance,omitempty"`
	// The name to identify the feature.
	Name *string `json:"Name,omitempty"`
	// The operational state of the feature.
	OperationalState *string `json:"OperationalState,omitempty"`
	// The status message to capture admin state detailed information.
	StatusMsg            *string                                     `json:"StatusMsg,omitempty"`
	NetworkElement       NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkFeatureControl NetworkFeatureControl

// NewNetworkFeatureControl instantiates a new NetworkFeatureControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeatureControl(classId string, objectType string) *NetworkFeatureControl {
	this := NetworkFeatureControl{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkFeatureControlWithDefaults instantiates a new NetworkFeatureControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeatureControlWithDefaults() *NetworkFeatureControl {
	this := NetworkFeatureControl{}
	var classId string = "network.FeatureControl"
	this.ClassId = classId
	var objectType string = "network.FeatureControl"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkFeatureControl) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkFeatureControl) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "network.FeatureControl" of the ClassId field.
func (o *NetworkFeatureControl) GetDefaultClassId() interface{} {
	return "network.FeatureControl"
}

// GetObjectType returns the ObjectType field value
func (o *NetworkFeatureControl) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkFeatureControl) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "network.FeatureControl" of the ObjectType field.
func (o *NetworkFeatureControl) GetDefaultObjectType() interface{} {
	return "network.FeatureControl"
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetAdminState() string {
	if o == nil || IsNil(o.AdminState) {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetAdminStateOk() (*string, bool) {
	if o == nil || IsNil(o.AdminState) {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasAdminState() bool {
	if o != nil && !IsNil(o.AdminState) {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NetworkFeatureControl) SetAdminState(v string) {
	o.AdminState = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetInstance() int64 {
	if o == nil || IsNil(o.Instance) {
		var ret int64
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetInstanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given int64 and assigns it to the Instance field.
func (o *NetworkFeatureControl) SetInstance(v int64) {
	o.Instance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkFeatureControl) SetName(v string) {
	o.Name = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetOperationalState() string {
	if o == nil || IsNil(o.OperationalState) {
		var ret string
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetOperationalStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given string and assigns it to the OperationalState field.
func (o *NetworkFeatureControl) SetOperationalState(v string) {
	o.OperationalState = &v
}

// GetStatusMsg returns the StatusMsg field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetStatusMsg() string {
	if o == nil || IsNil(o.StatusMsg) {
		var ret string
		return ret
	}
	return *o.StatusMsg
}

// GetStatusMsgOk returns a tuple with the StatusMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetStatusMsgOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMsg) {
		return nil, false
	}
	return o.StatusMsg, true
}

// HasStatusMsg returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasStatusMsg() bool {
	if o != nil && !IsNil(o.StatusMsg) {
		return true
	}

	return false
}

// SetStatusMsg gets a reference to the given string and assigns it to the StatusMsg field.
func (o *NetworkFeatureControl) SetStatusMsg(v string) {
	o.StatusMsg = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkFeatureControl) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkFeatureControl) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkFeatureControl) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *NetworkFeatureControl) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *NetworkFeatureControl) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkFeatureControl) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkFeatureControl) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkFeatureControl) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NetworkFeatureControl) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NetworkFeatureControl) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NetworkFeatureControl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkFeatureControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdminState) {
		toSerialize["AdminState"] = o.AdminState
	}
	if !IsNil(o.Instance) {
		toSerialize["Instance"] = o.Instance
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.OperationalState) {
		toSerialize["OperationalState"] = o.OperationalState
	}
	if !IsNil(o.StatusMsg) {
		toSerialize["StatusMsg"] = o.StatusMsg
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkFeatureControl) UnmarshalJSON(data []byte) (err error) {
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
	type NetworkFeatureControlWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The admin state of the feature.
		AdminState *string `json:"AdminState,omitempty"`
		// The number of instances of the feature.
		Instance *int64 `json:"Instance,omitempty"`
		// The name to identify the feature.
		Name *string `json:"Name,omitempty"`
		// The operational state of the feature.
		OperationalState *string `json:"OperationalState,omitempty"`
		// The status message to capture admin state detailed information.
		StatusMsg        *string                                     `json:"StatusMsg,omitempty"`
		NetworkElement   NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkFeatureControlWithoutEmbeddedStruct := NetworkFeatureControlWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNetworkFeatureControlWithoutEmbeddedStruct)
	if err == nil {
		varNetworkFeatureControl := _NetworkFeatureControl{}
		varNetworkFeatureControl.ClassId = varNetworkFeatureControlWithoutEmbeddedStruct.ClassId
		varNetworkFeatureControl.ObjectType = varNetworkFeatureControlWithoutEmbeddedStruct.ObjectType
		varNetworkFeatureControl.AdminState = varNetworkFeatureControlWithoutEmbeddedStruct.AdminState
		varNetworkFeatureControl.Instance = varNetworkFeatureControlWithoutEmbeddedStruct.Instance
		varNetworkFeatureControl.Name = varNetworkFeatureControlWithoutEmbeddedStruct.Name
		varNetworkFeatureControl.OperationalState = varNetworkFeatureControlWithoutEmbeddedStruct.OperationalState
		varNetworkFeatureControl.StatusMsg = varNetworkFeatureControlWithoutEmbeddedStruct.StatusMsg
		varNetworkFeatureControl.NetworkElement = varNetworkFeatureControlWithoutEmbeddedStruct.NetworkElement
		varNetworkFeatureControl.RegisteredDevice = varNetworkFeatureControlWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkFeatureControl(varNetworkFeatureControl)
	} else {
		return err
	}

	varNetworkFeatureControl := _NetworkFeatureControl{}

	err = json.Unmarshal(data, &varNetworkFeatureControl)
	if err == nil {
		o.InventoryBase = varNetworkFeatureControl.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "Instance")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperationalState")
		delete(additionalProperties, "StatusMsg")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkFeatureControl struct {
	value *NetworkFeatureControl
	isSet bool
}

func (v NullableNetworkFeatureControl) Get() *NetworkFeatureControl {
	return v.value
}

func (v *NullableNetworkFeatureControl) Set(val *NetworkFeatureControl) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeatureControl) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeatureControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeatureControl(val *NetworkFeatureControl) *NullableNetworkFeatureControl {
	return &NullableNetworkFeatureControl{value: val, isSet: true}
}

func (v NullableNetworkFeatureControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeatureControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
