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

// ConnectorStreamInput Write input to a running stream. Cloud services can send input to a running stream. e.g. input to a running pseudoterminal If the requested stream is not running and error will be returned.
type ConnectorStreamInput struct {
	ConnectorStreamMessage
	AdditionalProperties map[string]interface{}
}

type _ConnectorStreamInput ConnectorStreamInput

// NewConnectorStreamInput instantiates a new ConnectorStreamInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamInput(classId string, objectType string) *ConnectorStreamInput {
	this := ConnectorStreamInput{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStreamInputWithDefaults instantiates a new ConnectorStreamInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamInputWithDefaults() *ConnectorStreamInput {
	this := ConnectorStreamInput{}
	return &this
}

func (o ConnectorStreamInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStreamInput) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorStreamInputWithoutEmbeddedStruct struct {
	}

	varConnectorStreamInputWithoutEmbeddedStruct := ConnectorStreamInputWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorStreamInputWithoutEmbeddedStruct)
	if err == nil {
		varConnectorStreamInput := _ConnectorStreamInput{}
		*o = ConnectorStreamInput(varConnectorStreamInput)
	} else {
		return err
	}

	varConnectorStreamInput := _ConnectorStreamInput{}

	err = json.Unmarshal(bytes, &varConnectorStreamInput)
	if err == nil {
		o.ConnectorStreamMessage = varConnectorStreamInput.ConnectorStreamMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectConnectorStreamMessage := reflect.ValueOf(o.ConnectorStreamMessage)
		for i := 0; i < reflectConnectorStreamMessage.Type().NumField(); i++ {
			t := reflectConnectorStreamMessage.Type().Field(i)

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

type NullableConnectorStreamInput struct {
	value *ConnectorStreamInput
	isSet bool
}

func (v NullableConnectorStreamInput) Get() *ConnectorStreamInput {
	return v.value
}

func (v *NullableConnectorStreamInput) Set(val *ConnectorStreamInput) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamInput) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamInput(val *ConnectorStreamInput) *NullableConnectorStreamInput {
	return &NullableConnectorStreamInput{value: val, isSet: true}
}

func (v NullableConnectorStreamInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
