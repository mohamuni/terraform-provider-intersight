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

// NiatelemetryBootflashDetails Stores the boot flash details per switch.
type NiatelemetryBootflashDetails struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return firmware revision in boot flash details.
	FwRev *string `json:"FwRev,omitempty"`
	// Return model type in boot flash details.
	ModelType *string `json:"ModelType,omitempty"`
	// Return serial id in boot flash details.
	Serial               *string `json:"Serial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryBootflashDetails NiatelemetryBootflashDetails

// NewNiatelemetryBootflashDetails instantiates a new NiatelemetryBootflashDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryBootflashDetails(classId string, objectType string) *NiatelemetryBootflashDetails {
	this := NiatelemetryBootflashDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryBootflashDetailsWithDefaults instantiates a new NiatelemetryBootflashDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryBootflashDetailsWithDefaults() *NiatelemetryBootflashDetails {
	this := NiatelemetryBootflashDetails{}
	var classId string = "niatelemetry.BootflashDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.BootflashDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryBootflashDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryBootflashDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryBootflashDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryBootflashDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFwRev returns the FwRev field value if set, zero value otherwise.
func (o *NiatelemetryBootflashDetails) GetFwRev() string {
	if o == nil || o.FwRev == nil {
		var ret string
		return ret
	}
	return *o.FwRev
}

// GetFwRevOk returns a tuple with the FwRev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetails) GetFwRevOk() (*string, bool) {
	if o == nil || o.FwRev == nil {
		return nil, false
	}
	return o.FwRev, true
}

// HasFwRev returns a boolean if a field has been set.
func (o *NiatelemetryBootflashDetails) HasFwRev() bool {
	if o != nil && o.FwRev != nil {
		return true
	}

	return false
}

// SetFwRev gets a reference to the given string and assigns it to the FwRev field.
func (o *NiatelemetryBootflashDetails) SetFwRev(v string) {
	o.FwRev = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *NiatelemetryBootflashDetails) GetModelType() string {
	if o == nil || o.ModelType == nil {
		var ret string
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetails) GetModelTypeOk() (*string, bool) {
	if o == nil || o.ModelType == nil {
		return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *NiatelemetryBootflashDetails) HasModelType() bool {
	if o != nil && o.ModelType != nil {
		return true
	}

	return false
}

// SetModelType gets a reference to the given string and assigns it to the ModelType field.
func (o *NiatelemetryBootflashDetails) SetModelType(v string) {
	o.ModelType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryBootflashDetails) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetails) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryBootflashDetails) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryBootflashDetails) SetSerial(v string) {
	o.Serial = &v
}

func (o NiatelemetryBootflashDetails) MarshalJSON() ([]byte, error) {
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
	if o.FwRev != nil {
		toSerialize["FwRev"] = o.FwRev
	}
	if o.ModelType != nil {
		toSerialize["ModelType"] = o.ModelType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryBootflashDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryBootflashDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return firmware revision in boot flash details.
		FwRev *string `json:"FwRev,omitempty"`
		// Return model type in boot flash details.
		ModelType *string `json:"ModelType,omitempty"`
		// Return serial id in boot flash details.
		Serial *string `json:"Serial,omitempty"`
	}

	varNiatelemetryBootflashDetailsWithoutEmbeddedStruct := NiatelemetryBootflashDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryBootflashDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryBootflashDetails := _NiatelemetryBootflashDetails{}
		varNiatelemetryBootflashDetails.ClassId = varNiatelemetryBootflashDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryBootflashDetails.ObjectType = varNiatelemetryBootflashDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryBootflashDetails.FwRev = varNiatelemetryBootflashDetailsWithoutEmbeddedStruct.FwRev
		varNiatelemetryBootflashDetails.ModelType = varNiatelemetryBootflashDetailsWithoutEmbeddedStruct.ModelType
		varNiatelemetryBootflashDetails.Serial = varNiatelemetryBootflashDetailsWithoutEmbeddedStruct.Serial
		*o = NiatelemetryBootflashDetails(varNiatelemetryBootflashDetails)
	} else {
		return err
	}

	varNiatelemetryBootflashDetails := _NiatelemetryBootflashDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryBootflashDetails)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryBootflashDetails.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FwRev")
		delete(additionalProperties, "ModelType")
		delete(additionalProperties, "Serial")

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

type NullableNiatelemetryBootflashDetails struct {
	value *NiatelemetryBootflashDetails
	isSet bool
}

func (v NullableNiatelemetryBootflashDetails) Get() *NiatelemetryBootflashDetails {
	return v.value
}

func (v *NullableNiatelemetryBootflashDetails) Set(val *NiatelemetryBootflashDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryBootflashDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryBootflashDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryBootflashDetails(val *NiatelemetryBootflashDetails) *NullableNiatelemetryBootflashDetails {
	return &NullableNiatelemetryBootflashDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryBootflashDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryBootflashDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
