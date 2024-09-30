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

// checks if the ApplianceFileSystemOpStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceFileSystemOpStatus{}

// ApplianceFileSystemOpStatus Status of a file system on an Intersight Appliance node.
type ApplianceFileSystemOpStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Capacity of the file system in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Mount point of this file system.
	Mountpoint *string `json:"Mountpoint,omitempty"`
	// Operational status of the file system. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
	// Percentage of the file system capacity currently in use.
	Usage                *float32                                    `json:"Usage,omitempty"`
	NodeOpStatus         NullableApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceFileSystemOpStatus ApplianceFileSystemOpStatus

// NewApplianceFileSystemOpStatus instantiates a new ApplianceFileSystemOpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceFileSystemOpStatus(classId string, objectType string) *ApplianceFileSystemOpStatus {
	this := ApplianceFileSystemOpStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceFileSystemOpStatusWithDefaults instantiates a new ApplianceFileSystemOpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceFileSystemOpStatusWithDefaults() *ApplianceFileSystemOpStatus {
	this := ApplianceFileSystemOpStatus{}
	var classId string = "appliance.FileSystemOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.FileSystemOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceFileSystemOpStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceFileSystemOpStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.FileSystemOpStatus" of the ClassId field.
func (o *ApplianceFileSystemOpStatus) GetDefaultClassId() interface{} {
	return "appliance.FileSystemOpStatus"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceFileSystemOpStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceFileSystemOpStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.FileSystemOpStatus" of the ObjectType field.
func (o *ApplianceFileSystemOpStatus) GetDefaultObjectType() interface{} {
	return "appliance.FileSystemOpStatus"
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatus) GetCapacity() int64 {
	if o == nil || IsNil(o.Capacity) {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatus) GetCapacityOk() (*int64, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatus) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *ApplianceFileSystemOpStatus) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetMountpoint returns the Mountpoint field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatus) GetMountpoint() string {
	if o == nil || IsNil(o.Mountpoint) {
		var ret string
		return ret
	}
	return *o.Mountpoint
}

// GetMountpointOk returns a tuple with the Mountpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatus) GetMountpointOk() (*string, bool) {
	if o == nil || IsNil(o.Mountpoint) {
		return nil, false
	}
	return o.Mountpoint, true
}

// HasMountpoint returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatus) HasMountpoint() bool {
	if o != nil && !IsNil(o.Mountpoint) {
		return true
	}

	return false
}

// SetMountpoint gets a reference to the given string and assigns it to the Mountpoint field.
func (o *ApplianceFileSystemOpStatus) SetMountpoint(v string) {
	o.Mountpoint = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatus) GetOperationalStatus() string {
	if o == nil || IsNil(o.OperationalStatus) {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatus) GetOperationalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OperationalStatus) {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatus) HasOperationalStatus() bool {
	if o != nil && !IsNil(o.OperationalStatus) {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceFileSystemOpStatus) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatus) GetUsage() float32 {
	if o == nil || IsNil(o.Usage) {
		var ret float32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatus) GetUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatus) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given float32 and assigns it to the Usage field.
func (o *ApplianceFileSystemOpStatus) SetUsage(v float32) {
	o.Usage = &v
}

// GetNodeOpStatus returns the NodeOpStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceFileSystemOpStatus) GetNodeOpStatus() ApplianceNodeOpStatusRelationship {
	if o == nil || IsNil(o.NodeOpStatus.Get()) {
		var ret ApplianceNodeOpStatusRelationship
		return ret
	}
	return *o.NodeOpStatus.Get()
}

// GetNodeOpStatusOk returns a tuple with the NodeOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceFileSystemOpStatus) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeOpStatus.Get(), o.NodeOpStatus.IsSet()
}

// HasNodeOpStatus returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatus) HasNodeOpStatus() bool {
	if o != nil && o.NodeOpStatus.IsSet() {
		return true
	}

	return false
}

// SetNodeOpStatus gets a reference to the given NullableApplianceNodeOpStatusRelationship and assigns it to the NodeOpStatus field.
func (o *ApplianceFileSystemOpStatus) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship) {
	o.NodeOpStatus.Set(&v)
}

// SetNodeOpStatusNil sets the value for NodeOpStatus to be an explicit nil
func (o *ApplianceFileSystemOpStatus) SetNodeOpStatusNil() {
	o.NodeOpStatus.Set(nil)
}

// UnsetNodeOpStatus ensures that no value is present for NodeOpStatus, not even an explicit nil
func (o *ApplianceFileSystemOpStatus) UnsetNodeOpStatus() {
	o.NodeOpStatus.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceFileSystemOpStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceFileSystemOpStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatus) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceFileSystemOpStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *ApplianceFileSystemOpStatus) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *ApplianceFileSystemOpStatus) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o ApplianceFileSystemOpStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceFileSystemOpStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Capacity) {
		toSerialize["Capacity"] = o.Capacity
	}
	if !IsNil(o.Mountpoint) {
		toSerialize["Mountpoint"] = o.Mountpoint
	}
	if !IsNil(o.OperationalStatus) {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if !IsNil(o.Usage) {
		toSerialize["Usage"] = o.Usage
	}
	if o.NodeOpStatus.IsSet() {
		toSerialize["NodeOpStatus"] = o.NodeOpStatus.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceFileSystemOpStatus) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceFileSystemOpStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Capacity of the file system in bytes.
		Capacity *int64 `json:"Capacity,omitempty"`
		// Mount point of this file system.
		Mountpoint *string `json:"Mountpoint,omitempty"`
		// Operational status of the file system. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
		OperationalStatus *string `json:"OperationalStatus,omitempty"`
		// Percentage of the file system capacity currently in use.
		Usage            *float32                                    `json:"Usage,omitempty"`
		NodeOpStatus     NullableApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varApplianceFileSystemOpStatusWithoutEmbeddedStruct := ApplianceFileSystemOpStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceFileSystemOpStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceFileSystemOpStatus := _ApplianceFileSystemOpStatus{}
		varApplianceFileSystemOpStatus.ClassId = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.ClassId
		varApplianceFileSystemOpStatus.ObjectType = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.ObjectType
		varApplianceFileSystemOpStatus.Capacity = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.Capacity
		varApplianceFileSystemOpStatus.Mountpoint = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.Mountpoint
		varApplianceFileSystemOpStatus.OperationalStatus = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.OperationalStatus
		varApplianceFileSystemOpStatus.Usage = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.Usage
		varApplianceFileSystemOpStatus.NodeOpStatus = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.NodeOpStatus
		varApplianceFileSystemOpStatus.RegisteredDevice = varApplianceFileSystemOpStatusWithoutEmbeddedStruct.RegisteredDevice
		*o = ApplianceFileSystemOpStatus(varApplianceFileSystemOpStatus)
	} else {
		return err
	}

	varApplianceFileSystemOpStatus := _ApplianceFileSystemOpStatus{}

	err = json.Unmarshal(data, &varApplianceFileSystemOpStatus)
	if err == nil {
		o.MoBaseMo = varApplianceFileSystemOpStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mountpoint")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "NodeOpStatus")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableApplianceFileSystemOpStatus struct {
	value *ApplianceFileSystemOpStatus
	isSet bool
}

func (v NullableApplianceFileSystemOpStatus) Get() *ApplianceFileSystemOpStatus {
	return v.value
}

func (v *NullableApplianceFileSystemOpStatus) Set(val *ApplianceFileSystemOpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceFileSystemOpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceFileSystemOpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceFileSystemOpStatus(val *ApplianceFileSystemOpStatus) *NullableApplianceFileSystemOpStatus {
	return &NullableApplianceFileSystemOpStatus{value: val, isSet: true}
}

func (v NullableApplianceFileSystemOpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceFileSystemOpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
