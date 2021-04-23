/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-15T06:27:08Z.
 *
 * API version: 1.0.9-4247
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryNxosVtp Stores vtp data per switch.
type NiatelemetryNxosVtp struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the status of operational mode of vtp.
	OperMode *string `json:"OperMode,omitempty"`
	// Returns the status pruning mode of vtp.
	PruningMode *string `json:"PruningMode,omitempty"`
	// Returns the running version of vtp.
	RunningVersion *string `json:"RunningVersion,omitempty"`
	// Returns the status of trap in vtp.
	TrapEnabled *string `json:"TrapEnabled,omitempty"`
	// Returns the status of v2 mode of vtp.
	V2Mode *string `json:"V2Mode,omitempty"`
	// Returns version of vtp running.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNxosVtp NiatelemetryNxosVtp

// NewNiatelemetryNxosVtp instantiates a new NiatelemetryNxosVtp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNxosVtp(classId string, objectType string) *NiatelemetryNxosVtp {
	this := NiatelemetryNxosVtp{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNxosVtpWithDefaults instantiates a new NiatelemetryNxosVtp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNxosVtpWithDefaults() *NiatelemetryNxosVtp {
	this := NiatelemetryNxosVtp{}
	var classId string = "niatelemetry.NxosVtp"
	this.ClassId = classId
	var objectType string = "niatelemetry.NxosVtp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNxosVtp) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNxosVtp) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNxosVtp) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNxosVtp) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperMode returns the OperMode field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtp) GetOperMode() string {
	if o == nil || o.OperMode == nil {
		var ret string
		return ret
	}
	return *o.OperMode
}

// GetOperModeOk returns a tuple with the OperMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetOperModeOk() (*string, bool) {
	if o == nil || o.OperMode == nil {
		return nil, false
	}
	return o.OperMode, true
}

// HasOperMode returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtp) HasOperMode() bool {
	if o != nil && o.OperMode != nil {
		return true
	}

	return false
}

// SetOperMode gets a reference to the given string and assigns it to the OperMode field.
func (o *NiatelemetryNxosVtp) SetOperMode(v string) {
	o.OperMode = &v
}

// GetPruningMode returns the PruningMode field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtp) GetPruningMode() string {
	if o == nil || o.PruningMode == nil {
		var ret string
		return ret
	}
	return *o.PruningMode
}

// GetPruningModeOk returns a tuple with the PruningMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetPruningModeOk() (*string, bool) {
	if o == nil || o.PruningMode == nil {
		return nil, false
	}
	return o.PruningMode, true
}

// HasPruningMode returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtp) HasPruningMode() bool {
	if o != nil && o.PruningMode != nil {
		return true
	}

	return false
}

// SetPruningMode gets a reference to the given string and assigns it to the PruningMode field.
func (o *NiatelemetryNxosVtp) SetPruningMode(v string) {
	o.PruningMode = &v
}

// GetRunningVersion returns the RunningVersion field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtp) GetRunningVersion() string {
	if o == nil || o.RunningVersion == nil {
		var ret string
		return ret
	}
	return *o.RunningVersion
}

// GetRunningVersionOk returns a tuple with the RunningVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetRunningVersionOk() (*string, bool) {
	if o == nil || o.RunningVersion == nil {
		return nil, false
	}
	return o.RunningVersion, true
}

// HasRunningVersion returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtp) HasRunningVersion() bool {
	if o != nil && o.RunningVersion != nil {
		return true
	}

	return false
}

// SetRunningVersion gets a reference to the given string and assigns it to the RunningVersion field.
func (o *NiatelemetryNxosVtp) SetRunningVersion(v string) {
	o.RunningVersion = &v
}

// GetTrapEnabled returns the TrapEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtp) GetTrapEnabled() string {
	if o == nil || o.TrapEnabled == nil {
		var ret string
		return ret
	}
	return *o.TrapEnabled
}

// GetTrapEnabledOk returns a tuple with the TrapEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetTrapEnabledOk() (*string, bool) {
	if o == nil || o.TrapEnabled == nil {
		return nil, false
	}
	return o.TrapEnabled, true
}

// HasTrapEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtp) HasTrapEnabled() bool {
	if o != nil && o.TrapEnabled != nil {
		return true
	}

	return false
}

// SetTrapEnabled gets a reference to the given string and assigns it to the TrapEnabled field.
func (o *NiatelemetryNxosVtp) SetTrapEnabled(v string) {
	o.TrapEnabled = &v
}

// GetV2Mode returns the V2Mode field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtp) GetV2Mode() string {
	if o == nil || o.V2Mode == nil {
		var ret string
		return ret
	}
	return *o.V2Mode
}

// GetV2ModeOk returns a tuple with the V2Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetV2ModeOk() (*string, bool) {
	if o == nil || o.V2Mode == nil {
		return nil, false
	}
	return o.V2Mode, true
}

// HasV2Mode returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtp) HasV2Mode() bool {
	if o != nil && o.V2Mode != nil {
		return true
	}

	return false
}

// SetV2Mode gets a reference to the given string and assigns it to the V2Mode field.
func (o *NiatelemetryNxosVtp) SetV2Mode(v string) {
	o.V2Mode = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtp) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtp) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtp) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *NiatelemetryNxosVtp) SetVersion(v int64) {
	o.Version = &v
}

func (o NiatelemetryNxosVtp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperMode != nil {
		toSerialize["OperMode"] = o.OperMode
	}
	if o.PruningMode != nil {
		toSerialize["PruningMode"] = o.PruningMode
	}
	if o.RunningVersion != nil {
		toSerialize["RunningVersion"] = o.RunningVersion
	}
	if o.TrapEnabled != nil {
		toSerialize["TrapEnabled"] = o.TrapEnabled
	}
	if o.V2Mode != nil {
		toSerialize["V2Mode"] = o.V2Mode
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNxosVtp) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNxosVtpWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns the status of operational mode of vtp.
		OperMode *string `json:"OperMode,omitempty"`
		// Returns the status pruning mode of vtp.
		PruningMode *string `json:"PruningMode,omitempty"`
		// Returns the running version of vtp.
		RunningVersion *string `json:"RunningVersion,omitempty"`
		// Returns the status of trap in vtp.
		TrapEnabled *string `json:"TrapEnabled,omitempty"`
		// Returns the status of v2 mode of vtp.
		V2Mode *string `json:"V2Mode,omitempty"`
		// Returns version of vtp running.
		Version *int64 `json:"Version,omitempty"`
	}

	varNiatelemetryNxosVtpWithoutEmbeddedStruct := NiatelemetryNxosVtpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNxosVtpWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNxosVtp := _NiatelemetryNxosVtp{}
		varNiatelemetryNxosVtp.ClassId = varNiatelemetryNxosVtpWithoutEmbeddedStruct.ClassId
		varNiatelemetryNxosVtp.ObjectType = varNiatelemetryNxosVtpWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNxosVtp.OperMode = varNiatelemetryNxosVtpWithoutEmbeddedStruct.OperMode
		varNiatelemetryNxosVtp.PruningMode = varNiatelemetryNxosVtpWithoutEmbeddedStruct.PruningMode
		varNiatelemetryNxosVtp.RunningVersion = varNiatelemetryNxosVtpWithoutEmbeddedStruct.RunningVersion
		varNiatelemetryNxosVtp.TrapEnabled = varNiatelemetryNxosVtpWithoutEmbeddedStruct.TrapEnabled
		varNiatelemetryNxosVtp.V2Mode = varNiatelemetryNxosVtpWithoutEmbeddedStruct.V2Mode
		varNiatelemetryNxosVtp.Version = varNiatelemetryNxosVtpWithoutEmbeddedStruct.Version
		*o = NiatelemetryNxosVtp(varNiatelemetryNxosVtp)
	} else {
		return err
	}

	varNiatelemetryNxosVtp := _NiatelemetryNxosVtp{}

	err = json.Unmarshal(bytes, &varNiatelemetryNxosVtp)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryNxosVtp.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperMode")
		delete(additionalProperties, "PruningMode")
		delete(additionalProperties, "RunningVersion")
		delete(additionalProperties, "TrapEnabled")
		delete(additionalProperties, "V2Mode")
		delete(additionalProperties, "Version")

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

type NullableNiatelemetryNxosVtp struct {
	value *NiatelemetryNxosVtp
	isSet bool
}

func (v NullableNiatelemetryNxosVtp) Get() *NiatelemetryNxosVtp {
	return v.value
}

func (v *NullableNiatelemetryNxosVtp) Set(val *NiatelemetryNxosVtp) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNxosVtp) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNxosVtp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNxosVtp(val *NiatelemetryNxosVtp) *NullableNiatelemetryNxosVtp {
	return &NullableNiatelemetryNxosVtp{value: val, isSet: true}
}

func (v NullableNiatelemetryNxosVtp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNxosVtp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}