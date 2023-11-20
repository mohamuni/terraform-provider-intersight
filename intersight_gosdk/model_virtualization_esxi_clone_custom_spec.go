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

// VirtualizationEsxiCloneCustomSpec ESXi virtual machine clone specification.
type VirtualizationEsxiCloneCustomSpec struct {
	VirtualizationBaseCustomSpec
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the Extra Config specification which can be configured on virtual machine.
	ExtraConfig interface{} `json:"ExtraConfig,omitempty"`
	// Specify the OVA Environment specification which can be configured on virtual machine.
	OvaEnvSpec           interface{} `json:"OvaEnvSpec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiCloneCustomSpec VirtualizationEsxiCloneCustomSpec

// NewVirtualizationEsxiCloneCustomSpec instantiates a new VirtualizationEsxiCloneCustomSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiCloneCustomSpec(classId string, objectType string) *VirtualizationEsxiCloneCustomSpec {
	this := VirtualizationEsxiCloneCustomSpec{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiCloneCustomSpecWithDefaults instantiates a new VirtualizationEsxiCloneCustomSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiCloneCustomSpecWithDefaults() *VirtualizationEsxiCloneCustomSpec {
	this := VirtualizationEsxiCloneCustomSpec{}
	var classId string = "virtualization.EsxiCloneCustomSpec"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiCloneCustomSpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiCloneCustomSpec) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiCloneCustomSpec) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiCloneCustomSpec) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiCloneCustomSpec) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiCloneCustomSpec) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiCloneCustomSpec) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiCloneCustomSpec) GetExtraConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiCloneCustomSpec) GetExtraConfigOk() (*interface{}, bool) {
	if o == nil || o.ExtraConfig == nil {
		return nil, false
	}
	return &o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *VirtualizationEsxiCloneCustomSpec) HasExtraConfig() bool {
	if o != nil && o.ExtraConfig != nil {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given interface{} and assigns it to the ExtraConfig field.
func (o *VirtualizationEsxiCloneCustomSpec) SetExtraConfig(v interface{}) {
	o.ExtraConfig = v
}

// GetOvaEnvSpec returns the OvaEnvSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiCloneCustomSpec) GetOvaEnvSpec() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OvaEnvSpec
}

// GetOvaEnvSpecOk returns a tuple with the OvaEnvSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiCloneCustomSpec) GetOvaEnvSpecOk() (*interface{}, bool) {
	if o == nil || o.OvaEnvSpec == nil {
		return nil, false
	}
	return &o.OvaEnvSpec, true
}

// HasOvaEnvSpec returns a boolean if a field has been set.
func (o *VirtualizationEsxiCloneCustomSpec) HasOvaEnvSpec() bool {
	if o != nil && o.OvaEnvSpec != nil {
		return true
	}

	return false
}

// SetOvaEnvSpec gets a reference to the given interface{} and assigns it to the OvaEnvSpec field.
func (o *VirtualizationEsxiCloneCustomSpec) SetOvaEnvSpec(v interface{}) {
	o.OvaEnvSpec = v
}

func (o VirtualizationEsxiCloneCustomSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseCustomSpec, errVirtualizationBaseCustomSpec := json.Marshal(o.VirtualizationBaseCustomSpec)
	if errVirtualizationBaseCustomSpec != nil {
		return []byte{}, errVirtualizationBaseCustomSpec
	}
	errVirtualizationBaseCustomSpec = json.Unmarshal([]byte(serializedVirtualizationBaseCustomSpec), &toSerialize)
	if errVirtualizationBaseCustomSpec != nil {
		return []byte{}, errVirtualizationBaseCustomSpec
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExtraConfig != nil {
		toSerialize["ExtraConfig"] = o.ExtraConfig
	}
	if o.OvaEnvSpec != nil {
		toSerialize["OvaEnvSpec"] = o.OvaEnvSpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiCloneCustomSpec) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specify the Extra Config specification which can be configured on virtual machine.
		ExtraConfig interface{} `json:"ExtraConfig,omitempty"`
		// Specify the OVA Environment specification which can be configured on virtual machine.
		OvaEnvSpec interface{} `json:"OvaEnvSpec,omitempty"`
	}

	varVirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct := VirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationEsxiCloneCustomSpec := _VirtualizationEsxiCloneCustomSpec{}
		varVirtualizationEsxiCloneCustomSpec.ClassId = varVirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct.ClassId
		varVirtualizationEsxiCloneCustomSpec.ObjectType = varVirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct.ObjectType
		varVirtualizationEsxiCloneCustomSpec.ExtraConfig = varVirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct.ExtraConfig
		varVirtualizationEsxiCloneCustomSpec.OvaEnvSpec = varVirtualizationEsxiCloneCustomSpecWithoutEmbeddedStruct.OvaEnvSpec
		*o = VirtualizationEsxiCloneCustomSpec(varVirtualizationEsxiCloneCustomSpec)
	} else {
		return err
	}

	varVirtualizationEsxiCloneCustomSpec := _VirtualizationEsxiCloneCustomSpec{}

	err = json.Unmarshal(bytes, &varVirtualizationEsxiCloneCustomSpec)
	if err == nil {
		o.VirtualizationBaseCustomSpec = varVirtualizationEsxiCloneCustomSpec.VirtualizationBaseCustomSpec
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExtraConfig")
		delete(additionalProperties, "OvaEnvSpec")

		// remove fields from embedded structs
		reflectVirtualizationBaseCustomSpec := reflect.ValueOf(o.VirtualizationBaseCustomSpec)
		for i := 0; i < reflectVirtualizationBaseCustomSpec.Type().NumField(); i++ {
			t := reflectVirtualizationBaseCustomSpec.Type().Field(i)

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

type NullableVirtualizationEsxiCloneCustomSpec struct {
	value *VirtualizationEsxiCloneCustomSpec
	isSet bool
}

func (v NullableVirtualizationEsxiCloneCustomSpec) Get() *VirtualizationEsxiCloneCustomSpec {
	return v.value
}

func (v *NullableVirtualizationEsxiCloneCustomSpec) Set(val *VirtualizationEsxiCloneCustomSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiCloneCustomSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiCloneCustomSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiCloneCustomSpec(val *VirtualizationEsxiCloneCustomSpec) *NullableVirtualizationEsxiCloneCustomSpec {
	return &NullableVirtualizationEsxiCloneCustomSpec{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiCloneCustomSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiCloneCustomSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
