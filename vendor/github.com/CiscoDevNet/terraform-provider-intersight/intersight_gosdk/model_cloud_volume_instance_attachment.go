/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18775
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// checks if the CloudVolumeInstanceAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudVolumeInstanceAttachment{}

// CloudVolumeInstanceAttachment Information about the volume attachments.
type CloudVolumeInstanceAttachment struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Time stamp when the attachment initiated.
	AttachTime *time.Time `json:"AttachTime,omitempty"`
	// If true, volume is deleted on virtual machine termination.
	AutoDelete *bool `json:"AutoDelete,omitempty"`
	// Device name assigned to the volume.
	DeviceName *string `json:"DeviceName,omitempty"`
	// ID of the virtual machine, the volume is attached to.
	InstanceId *string `json:"InstanceId,omitempty"`
	// The attachment state of the volume. * `UnRecognized` - Volume is in unrecognized state. * `Attached` - Volume is attached to the virtual machine. * `Attaching` - Volume is being attached to the virtual machine. * `Detaching` - Volume is being detached from the virtual machine.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudVolumeInstanceAttachment CloudVolumeInstanceAttachment

// NewCloudVolumeInstanceAttachment instantiates a new CloudVolumeInstanceAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeInstanceAttachment(classId string, objectType string) *CloudVolumeInstanceAttachment {
	this := CloudVolumeInstanceAttachment{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudVolumeInstanceAttachmentWithDefaults instantiates a new CloudVolumeInstanceAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeInstanceAttachmentWithDefaults() *CloudVolumeInstanceAttachment {
	this := CloudVolumeInstanceAttachment{}
	var classId string = "cloud.VolumeInstanceAttachment"
	this.ClassId = classId
	var objectType string = "cloud.VolumeInstanceAttachment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudVolumeInstanceAttachment) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudVolumeInstanceAttachment) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cloud.VolumeInstanceAttachment" of the ClassId field.
func (o *CloudVolumeInstanceAttachment) GetDefaultClassId() interface{} {
	return "cloud.VolumeInstanceAttachment"
}

// GetObjectType returns the ObjectType field value
func (o *CloudVolumeInstanceAttachment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudVolumeInstanceAttachment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cloud.VolumeInstanceAttachment" of the ObjectType field.
func (o *CloudVolumeInstanceAttachment) GetDefaultObjectType() interface{} {
	return "cloud.VolumeInstanceAttachment"
}

// GetAttachTime returns the AttachTime field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachment) GetAttachTime() time.Time {
	if o == nil || IsNil(o.AttachTime) {
		var ret time.Time
		return ret
	}
	return *o.AttachTime
}

// GetAttachTimeOk returns a tuple with the AttachTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetAttachTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AttachTime) {
		return nil, false
	}
	return o.AttachTime, true
}

// HasAttachTime returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachment) HasAttachTime() bool {
	if o != nil && !IsNil(o.AttachTime) {
		return true
	}

	return false
}

// SetAttachTime gets a reference to the given time.Time and assigns it to the AttachTime field.
func (o *CloudVolumeInstanceAttachment) SetAttachTime(v time.Time) {
	o.AttachTime = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachment) GetAutoDelete() bool {
	if o == nil || IsNil(o.AutoDelete) {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachment) HasAutoDelete() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *CloudVolumeInstanceAttachment) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachment) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachment) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *CloudVolumeInstanceAttachment) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachment) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachment) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *CloudVolumeInstanceAttachment) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachment) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachment) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachment) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudVolumeInstanceAttachment) SetState(v string) {
	o.State = &v
}

func (o CloudVolumeInstanceAttachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudVolumeInstanceAttachment) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AttachTime) {
		toSerialize["AttachTime"] = o.AttachTime
	}
	if !IsNil(o.AutoDelete) {
		toSerialize["AutoDelete"] = o.AutoDelete
	}
	if !IsNil(o.DeviceName) {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if !IsNil(o.InstanceId) {
		toSerialize["InstanceId"] = o.InstanceId
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudVolumeInstanceAttachment) UnmarshalJSON(data []byte) (err error) {
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
	type CloudVolumeInstanceAttachmentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Time stamp when the attachment initiated.
		AttachTime *time.Time `json:"AttachTime,omitempty"`
		// If true, volume is deleted on virtual machine termination.
		AutoDelete *bool `json:"AutoDelete,omitempty"`
		// Device name assigned to the volume.
		DeviceName *string `json:"DeviceName,omitempty"`
		// ID of the virtual machine, the volume is attached to.
		InstanceId *string `json:"InstanceId,omitempty"`
		// The attachment state of the volume. * `UnRecognized` - Volume is in unrecognized state. * `Attached` - Volume is attached to the virtual machine. * `Attaching` - Volume is being attached to the virtual machine. * `Detaching` - Volume is being detached from the virtual machine.
		State *string `json:"State,omitempty"`
	}

	varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct := CloudVolumeInstanceAttachmentWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct)
	if err == nil {
		varCloudVolumeInstanceAttachment := _CloudVolumeInstanceAttachment{}
		varCloudVolumeInstanceAttachment.ClassId = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.ClassId
		varCloudVolumeInstanceAttachment.ObjectType = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.ObjectType
		varCloudVolumeInstanceAttachment.AttachTime = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.AttachTime
		varCloudVolumeInstanceAttachment.AutoDelete = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.AutoDelete
		varCloudVolumeInstanceAttachment.DeviceName = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.DeviceName
		varCloudVolumeInstanceAttachment.InstanceId = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.InstanceId
		varCloudVolumeInstanceAttachment.State = varCloudVolumeInstanceAttachmentWithoutEmbeddedStruct.State
		*o = CloudVolumeInstanceAttachment(varCloudVolumeInstanceAttachment)
	} else {
		return err
	}

	varCloudVolumeInstanceAttachment := _CloudVolumeInstanceAttachment{}

	err = json.Unmarshal(data, &varCloudVolumeInstanceAttachment)
	if err == nil {
		o.MoBaseComplexType = varCloudVolumeInstanceAttachment.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttachTime")
		delete(additionalProperties, "AutoDelete")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "InstanceId")
		delete(additionalProperties, "State")

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

type NullableCloudVolumeInstanceAttachment struct {
	value *CloudVolumeInstanceAttachment
	isSet bool
}

func (v NullableCloudVolumeInstanceAttachment) Get() *CloudVolumeInstanceAttachment {
	return v.value
}

func (v *NullableCloudVolumeInstanceAttachment) Set(val *CloudVolumeInstanceAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeInstanceAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeInstanceAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeInstanceAttachment(val *CloudVolumeInstanceAttachment) *NullableCloudVolumeInstanceAttachment {
	return &NullableCloudVolumeInstanceAttachment{value: val, isSet: true}
}

func (v NullableCloudVolumeInstanceAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeInstanceAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
