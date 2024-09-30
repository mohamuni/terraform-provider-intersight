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

// checks if the FabricSpanDestEthPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricSpanDestEthPort{}

// FabricSpanDestEthPort Configures Ethernet SPAN Destination Port for a given SPAN session.
type FabricSpanDestEthPort struct {
	FabricAbstractSpanDestPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin speed of the SPAN Destination Port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps. * `NegAuto25Gbps` - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Forward error correction setting of the SPAN Destination Port. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
	Fec                  *string `json:"Fec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSpanDestEthPort FabricSpanDestEthPort

// NewFabricSpanDestEthPort instantiates a new FabricSpanDestEthPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSpanDestEthPort(classId string, objectType string) *FabricSpanDestEthPort {
	this := FabricSpanDestEthPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// NewFabricSpanDestEthPortWithDefaults instantiates a new FabricSpanDestEthPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSpanDestEthPortWithDefaults() *FabricSpanDestEthPort {
	this := FabricSpanDestEthPort{}
	var classId string = "fabric.SpanDestEthPort"
	this.ClassId = classId
	var objectType string = "fabric.SpanDestEthPort"
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSpanDestEthPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSpanDestEthPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSpanDestEthPort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SpanDestEthPort" of the ClassId field.
func (o *FabricSpanDestEthPort) GetDefaultClassId() interface{} {
	return "fabric.SpanDestEthPort"
}

// GetObjectType returns the ObjectType field value
func (o *FabricSpanDestEthPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSpanDestEthPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSpanDestEthPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SpanDestEthPort" of the ObjectType field.
func (o *FabricSpanDestEthPort) GetDefaultObjectType() interface{} {
	return "fabric.SpanDestEthPort"
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricSpanDestEthPort) GetAdminSpeed() string {
	if o == nil || IsNil(o.AdminSpeed) {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSpanDestEthPort) GetAdminSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.AdminSpeed) {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricSpanDestEthPort) HasAdminSpeed() bool {
	if o != nil && !IsNil(o.AdminSpeed) {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricSpanDestEthPort) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetFec returns the Fec field value if set, zero value otherwise.
func (o *FabricSpanDestEthPort) GetFec() string {
	if o == nil || IsNil(o.Fec) {
		var ret string
		return ret
	}
	return *o.Fec
}

// GetFecOk returns a tuple with the Fec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSpanDestEthPort) GetFecOk() (*string, bool) {
	if o == nil || IsNil(o.Fec) {
		return nil, false
	}
	return o.Fec, true
}

// HasFec returns a boolean if a field has been set.
func (o *FabricSpanDestEthPort) HasFec() bool {
	if o != nil && !IsNil(o.Fec) {
		return true
	}

	return false
}

// SetFec gets a reference to the given string and assigns it to the Fec field.
func (o *FabricSpanDestEthPort) SetFec(v string) {
	o.Fec = &v
}

func (o FabricSpanDestEthPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricSpanDestEthPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractSpanDestPort, errFabricAbstractSpanDestPort := json.Marshal(o.FabricAbstractSpanDestPort)
	if errFabricAbstractSpanDestPort != nil {
		return map[string]interface{}{}, errFabricAbstractSpanDestPort
	}
	errFabricAbstractSpanDestPort = json.Unmarshal([]byte(serializedFabricAbstractSpanDestPort), &toSerialize)
	if errFabricAbstractSpanDestPort != nil {
		return map[string]interface{}{}, errFabricAbstractSpanDestPort
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdminSpeed) {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if !IsNil(o.Fec) {
		toSerialize["Fec"] = o.Fec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricSpanDestEthPort) UnmarshalJSON(data []byte) (err error) {
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
	type FabricSpanDestEthPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin speed of the SPAN Destination Port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps. * `NegAuto25Gbps` - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Forward error correction setting of the SPAN Destination Port. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
		Fec *string `json:"Fec,omitempty"`
	}

	varFabricSpanDestEthPortWithoutEmbeddedStruct := FabricSpanDestEthPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricSpanDestEthPortWithoutEmbeddedStruct)
	if err == nil {
		varFabricSpanDestEthPort := _FabricSpanDestEthPort{}
		varFabricSpanDestEthPort.ClassId = varFabricSpanDestEthPortWithoutEmbeddedStruct.ClassId
		varFabricSpanDestEthPort.ObjectType = varFabricSpanDestEthPortWithoutEmbeddedStruct.ObjectType
		varFabricSpanDestEthPort.AdminSpeed = varFabricSpanDestEthPortWithoutEmbeddedStruct.AdminSpeed
		varFabricSpanDestEthPort.Fec = varFabricSpanDestEthPortWithoutEmbeddedStruct.Fec
		*o = FabricSpanDestEthPort(varFabricSpanDestEthPort)
	} else {
		return err
	}

	varFabricSpanDestEthPort := _FabricSpanDestEthPort{}

	err = json.Unmarshal(data, &varFabricSpanDestEthPort)
	if err == nil {
		o.FabricAbstractSpanDestPort = varFabricSpanDestEthPort.FabricAbstractSpanDestPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "Fec")

		// remove fields from embedded structs
		reflectFabricAbstractSpanDestPort := reflect.ValueOf(o.FabricAbstractSpanDestPort)
		for i := 0; i < reflectFabricAbstractSpanDestPort.Type().NumField(); i++ {
			t := reflectFabricAbstractSpanDestPort.Type().Field(i)

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

type NullableFabricSpanDestEthPort struct {
	value *FabricSpanDestEthPort
	isSet bool
}

func (v NullableFabricSpanDestEthPort) Get() *FabricSpanDestEthPort {
	return v.value
}

func (v *NullableFabricSpanDestEthPort) Set(val *FabricSpanDestEthPort) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSpanDestEthPort) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSpanDestEthPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSpanDestEthPort(val *FabricSpanDestEthPort) *NullableFabricSpanDestEthPort {
	return &NullableFabricSpanDestEthPort{value: val, isSet: true}
}

func (v NullableFabricSpanDestEthPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSpanDestEthPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
