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

// InventoryInterface Abstract base class for all interfaces.
type InventoryInterface struct {
	InventoryBase
	AdditionalProperties map[string]interface{}
}

type _InventoryInterface InventoryInterface

// NewInventoryInterface instantiates a new InventoryInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryInterface(classId string, objectType string) *InventoryInterface {
	this := InventoryInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryInterfaceWithDefaults instantiates a new InventoryInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryInterfaceWithDefaults() *InventoryInterface {
	this := InventoryInterface{}
	return &this
}

func (o InventoryInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryInterface) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryInterfaceWithoutEmbeddedStruct struct {
	}

	varInventoryInterfaceWithoutEmbeddedStruct := InventoryInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varInventoryInterface := _InventoryInterface{}
		*o = InventoryInterface(varInventoryInterface)
	} else {
		return err
	}

	varInventoryInterface := _InventoryInterface{}

	err = json.Unmarshal(bytes, &varInventoryInterface)
	if err == nil {
		o.InventoryBase = varInventoryInterface.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

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

type NullableInventoryInterface struct {
	value *InventoryInterface
	isSet bool
}

func (v NullableInventoryInterface) Get() *InventoryInterface {
	return v.value
}

func (v *NullableInventoryInterface) Set(val *InventoryInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryInterface(val *InventoryInterface) *NullableInventoryInterface {
	return &NullableInventoryInterface{value: val, isSet: true}
}

func (v NullableInventoryInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
