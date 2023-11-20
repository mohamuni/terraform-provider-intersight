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

// BootLocalCdd Device type used when booting from local CD drive.
type BootLocalCdd struct {
	BootDeviceBase
	AdditionalProperties map[string]interface{}
}

type _BootLocalCdd BootLocalCdd

// NewBootLocalCdd instantiates a new BootLocalCdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootLocalCdd(classId string, objectType string) *BootLocalCdd {
	this := BootLocalCdd{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewBootLocalCddWithDefaults instantiates a new BootLocalCdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootLocalCddWithDefaults() *BootLocalCdd {
	this := BootLocalCdd{}
	return &this
}

func (o BootLocalCdd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBootDeviceBase, errBootDeviceBase := json.Marshal(o.BootDeviceBase)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	errBootDeviceBase = json.Unmarshal([]byte(serializedBootDeviceBase), &toSerialize)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootLocalCdd) UnmarshalJSON(bytes []byte) (err error) {
	type BootLocalCddWithoutEmbeddedStruct struct {
	}

	varBootLocalCddWithoutEmbeddedStruct := BootLocalCddWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootLocalCddWithoutEmbeddedStruct)
	if err == nil {
		varBootLocalCdd := _BootLocalCdd{}
		*o = BootLocalCdd(varBootLocalCdd)
	} else {
		return err
	}

	varBootLocalCdd := _BootLocalCdd{}

	err = json.Unmarshal(bytes, &varBootLocalCdd)
	if err == nil {
		o.BootDeviceBase = varBootLocalCdd.BootDeviceBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectBootDeviceBase := reflect.ValueOf(o.BootDeviceBase)
		for i := 0; i < reflectBootDeviceBase.Type().NumField(); i++ {
			t := reflectBootDeviceBase.Type().Field(i)

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

type NullableBootLocalCdd struct {
	value *BootLocalCdd
	isSet bool
}

func (v NullableBootLocalCdd) Get() *BootLocalCdd {
	return v.value
}

func (v *NullableBootLocalCdd) Set(val *BootLocalCdd) {
	v.value = val
	v.isSet = true
}

func (v NullableBootLocalCdd) IsSet() bool {
	return v.isSet
}

func (v *NullableBootLocalCdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootLocalCdd(val *BootLocalCdd) *NullableBootLocalCdd {
	return &NullableBootLocalCdd{value: val, isSet: true}
}

func (v NullableBootLocalCdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootLocalCdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
