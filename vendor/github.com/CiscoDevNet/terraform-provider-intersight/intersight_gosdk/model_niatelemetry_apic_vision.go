/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryApicVision Object to capture ApicVision App properties.
type NiatelemetryApicVision struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ApicVision App state. apicVisionState checks the current operational state of ApicVision app on APIC.
	ApicVisionState *string `json:"ApicVisionState,omitempty"`
	// ApicVision App state last updated timestamp. It indicates the last updated timestamp for operational state of ApicVision app.
	ApicVisionStateLastUpdateTs *string `json:"ApicVisionStateLastUpdateTs,omitempty"`
	// ApicVision App version. apicVisionVersion is used to check compatibility with Nexus Cloud features.
	ApicVisionVersion *string `json:"ApicVisionVersion,omitempty"`
	// Configuration issues depicts the failures for ApicVision managed package upgrade on APIC.
	ConfigIssues         *string                              `json:"ConfigIssues,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicVision NiatelemetryApicVision

// NewNiatelemetryApicVision instantiates a new NiatelemetryApicVision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicVision(classId string, objectType string) *NiatelemetryApicVision {
	this := NiatelemetryApicVision{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicVisionWithDefaults instantiates a new NiatelemetryApicVision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicVisionWithDefaults() *NiatelemetryApicVision {
	this := NiatelemetryApicVision{}
	var classId string = "niatelemetry.ApicVision"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicVision"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicVision) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicVision) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicVision) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicVision) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApicVisionState returns the ApicVisionState field value if set, zero value otherwise.
func (o *NiatelemetryApicVision) GetApicVisionState() string {
	if o == nil || o.ApicVisionState == nil {
		var ret string
		return ret
	}
	return *o.ApicVisionState
}

// GetApicVisionStateOk returns a tuple with the ApicVisionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetApicVisionStateOk() (*string, bool) {
	if o == nil || o.ApicVisionState == nil {
		return nil, false
	}
	return o.ApicVisionState, true
}

// HasApicVisionState returns a boolean if a field has been set.
func (o *NiatelemetryApicVision) HasApicVisionState() bool {
	if o != nil && o.ApicVisionState != nil {
		return true
	}

	return false
}

// SetApicVisionState gets a reference to the given string and assigns it to the ApicVisionState field.
func (o *NiatelemetryApicVision) SetApicVisionState(v string) {
	o.ApicVisionState = &v
}

// GetApicVisionStateLastUpdateTs returns the ApicVisionStateLastUpdateTs field value if set, zero value otherwise.
func (o *NiatelemetryApicVision) GetApicVisionStateLastUpdateTs() string {
	if o == nil || o.ApicVisionStateLastUpdateTs == nil {
		var ret string
		return ret
	}
	return *o.ApicVisionStateLastUpdateTs
}

// GetApicVisionStateLastUpdateTsOk returns a tuple with the ApicVisionStateLastUpdateTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetApicVisionStateLastUpdateTsOk() (*string, bool) {
	if o == nil || o.ApicVisionStateLastUpdateTs == nil {
		return nil, false
	}
	return o.ApicVisionStateLastUpdateTs, true
}

// HasApicVisionStateLastUpdateTs returns a boolean if a field has been set.
func (o *NiatelemetryApicVision) HasApicVisionStateLastUpdateTs() bool {
	if o != nil && o.ApicVisionStateLastUpdateTs != nil {
		return true
	}

	return false
}

// SetApicVisionStateLastUpdateTs gets a reference to the given string and assigns it to the ApicVisionStateLastUpdateTs field.
func (o *NiatelemetryApicVision) SetApicVisionStateLastUpdateTs(v string) {
	o.ApicVisionStateLastUpdateTs = &v
}

// GetApicVisionVersion returns the ApicVisionVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicVision) GetApicVisionVersion() string {
	if o == nil || o.ApicVisionVersion == nil {
		var ret string
		return ret
	}
	return *o.ApicVisionVersion
}

// GetApicVisionVersionOk returns a tuple with the ApicVisionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetApicVisionVersionOk() (*string, bool) {
	if o == nil || o.ApicVisionVersion == nil {
		return nil, false
	}
	return o.ApicVisionVersion, true
}

// HasApicVisionVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicVision) HasApicVisionVersion() bool {
	if o != nil && o.ApicVisionVersion != nil {
		return true
	}

	return false
}

// SetApicVisionVersion gets a reference to the given string and assigns it to the ApicVisionVersion field.
func (o *NiatelemetryApicVision) SetApicVisionVersion(v string) {
	o.ApicVisionVersion = &v
}

// GetConfigIssues returns the ConfigIssues field value if set, zero value otherwise.
func (o *NiatelemetryApicVision) GetConfigIssues() string {
	if o == nil || o.ConfigIssues == nil {
		var ret string
		return ret
	}
	return *o.ConfigIssues
}

// GetConfigIssuesOk returns a tuple with the ConfigIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetConfigIssuesOk() (*string, bool) {
	if o == nil || o.ConfigIssues == nil {
		return nil, false
	}
	return o.ConfigIssues, true
}

// HasConfigIssues returns a boolean if a field has been set.
func (o *NiatelemetryApicVision) HasConfigIssues() bool {
	if o != nil && o.ConfigIssues != nil {
		return true
	}

	return false
}

// SetConfigIssues gets a reference to the given string and assigns it to the ConfigIssues field.
func (o *NiatelemetryApicVision) SetConfigIssues(v string) {
	o.ConfigIssues = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicVision) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicVision) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicVision) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicVision) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicVision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApicVisionState != nil {
		toSerialize["ApicVisionState"] = o.ApicVisionState
	}
	if o.ApicVisionStateLastUpdateTs != nil {
		toSerialize["ApicVisionStateLastUpdateTs"] = o.ApicVisionStateLastUpdateTs
	}
	if o.ApicVisionVersion != nil {
		toSerialize["ApicVisionVersion"] = o.ApicVisionVersion
	}
	if o.ConfigIssues != nil {
		toSerialize["ConfigIssues"] = o.ConfigIssues
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicVision) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicVisionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// ApicVision App state. apicVisionState checks the current operational state of ApicVision app on APIC.
		ApicVisionState *string `json:"ApicVisionState,omitempty"`
		// ApicVision App state last updated timestamp. It indicates the last updated timestamp for operational state of ApicVision app.
		ApicVisionStateLastUpdateTs *string `json:"ApicVisionStateLastUpdateTs,omitempty"`
		// ApicVision App version. apicVisionVersion is used to check compatibility with Nexus Cloud features.
		ApicVisionVersion *string `json:"ApicVisionVersion,omitempty"`
		// Configuration issues depicts the failures for ApicVision managed package upgrade on APIC.
		ConfigIssues     *string                              `json:"ConfigIssues,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicVisionWithoutEmbeddedStruct := NiatelemetryApicVisionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicVisionWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicVision := _NiatelemetryApicVision{}
		varNiatelemetryApicVision.ClassId = varNiatelemetryApicVisionWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicVision.ObjectType = varNiatelemetryApicVisionWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicVision.ApicVisionState = varNiatelemetryApicVisionWithoutEmbeddedStruct.ApicVisionState
		varNiatelemetryApicVision.ApicVisionStateLastUpdateTs = varNiatelemetryApicVisionWithoutEmbeddedStruct.ApicVisionStateLastUpdateTs
		varNiatelemetryApicVision.ApicVisionVersion = varNiatelemetryApicVisionWithoutEmbeddedStruct.ApicVisionVersion
		varNiatelemetryApicVision.ConfigIssues = varNiatelemetryApicVisionWithoutEmbeddedStruct.ConfigIssues
		varNiatelemetryApicVision.RegisteredDevice = varNiatelemetryApicVisionWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicVision(varNiatelemetryApicVision)
	} else {
		return err
	}

	varNiatelemetryApicVision := _NiatelemetryApicVision{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicVision)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicVision.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApicVisionState")
		delete(additionalProperties, "ApicVisionStateLastUpdateTs")
		delete(additionalProperties, "ApicVisionVersion")
		delete(additionalProperties, "ConfigIssues")
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

type NullableNiatelemetryApicVision struct {
	value *NiatelemetryApicVision
	isSet bool
}

func (v NullableNiatelemetryApicVision) Get() *NiatelemetryApicVision {
	return v.value
}

func (v *NullableNiatelemetryApicVision) Set(val *NiatelemetryApicVision) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicVision) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicVision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicVision(val *NiatelemetryApicVision) *NullableNiatelemetryApicVision {
	return &NullableNiatelemetryApicVision{value: val, isSet: true}
}

func (v NullableNiatelemetryApicVision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicVision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
