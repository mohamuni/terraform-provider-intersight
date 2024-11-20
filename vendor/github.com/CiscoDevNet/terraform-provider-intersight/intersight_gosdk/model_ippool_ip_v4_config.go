/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the IppoolIpV4Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IppoolIpV4Config{}

// IppoolIpV4Config Network interface configuration data for IPv4 interfaces. Netmask, Gateway and DNS settings.
type IppoolIpV4Config struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of the default IPv4 gateway.
	Gateway *string `json:"Gateway,omitempty" validate:"regexp=^$|^([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.
	Netmask *string `json:"Netmask,omitempty" validate:"regexp=^$|^(((255\\\\.){3}(255|254|252|248|240|224|192|128|0+))|((255\\\\.){2}(255|254|252|248|240|224|192|128|0+)\\\\.0)|((255\\\\.)(255|254|252|248|240|224|192|128|0+)(\\\\.0+){2})|((255|254|252|248|240|224|192|128|0+)(\\\\.0+){3}))$"`
	// IP Address of the primary Domain Name System (DNS) server.
	PrimaryDns *string `json:"PrimaryDns,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// IP Address of the secondary Domain Name System (DNS) server.
	SecondaryDns         *string `json:"SecondaryDns,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	AdditionalProperties map[string]interface{}
}

type _IppoolIpV4Config IppoolIpV4Config

// NewIppoolIpV4Config instantiates a new IppoolIpV4Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpV4Config(classId string, objectType string) *IppoolIpV4Config {
	this := IppoolIpV4Config{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIppoolIpV4ConfigWithDefaults instantiates a new IppoolIpV4Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpV4ConfigWithDefaults() *IppoolIpV4Config {
	this := IppoolIpV4Config{}
	var classId string = "ippool.IpV4Config"
	this.ClassId = classId
	var objectType string = "ippool.IpV4Config"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolIpV4Config) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Config) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolIpV4Config) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ippool.IpV4Config" of the ClassId field.
func (o *IppoolIpV4Config) GetDefaultClassId() interface{} {
	return "ippool.IpV4Config"
}

// GetObjectType returns the ObjectType field value
func (o *IppoolIpV4Config) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Config) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolIpV4Config) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ippool.IpV4Config" of the ObjectType field.
func (o *IppoolIpV4Config) GetDefaultObjectType() interface{} {
	return "ippool.IpV4Config"
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *IppoolIpV4Config) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Config) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *IppoolIpV4Config) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *IppoolIpV4Config) SetGateway(v string) {
	o.Gateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *IppoolIpV4Config) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Config) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *IppoolIpV4Config) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *IppoolIpV4Config) SetNetmask(v string) {
	o.Netmask = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *IppoolIpV4Config) GetPrimaryDns() string {
	if o == nil || IsNil(o.PrimaryDns) {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Config) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryDns) {
		return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *IppoolIpV4Config) HasPrimaryDns() bool {
	if o != nil && !IsNil(o.PrimaryDns) {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *IppoolIpV4Config) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *IppoolIpV4Config) GetSecondaryDns() string {
	if o == nil || IsNil(o.SecondaryDns) {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Config) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryDns) {
		return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *IppoolIpV4Config) HasSecondaryDns() bool {
	if o != nil && !IsNil(o.SecondaryDns) {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *IppoolIpV4Config) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

func (o IppoolIpV4Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IppoolIpV4Config) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Gateway) {
		toSerialize["Gateway"] = o.Gateway
	}
	if !IsNil(o.Netmask) {
		toSerialize["Netmask"] = o.Netmask
	}
	if !IsNil(o.PrimaryDns) {
		toSerialize["PrimaryDns"] = o.PrimaryDns
	}
	if !IsNil(o.SecondaryDns) {
		toSerialize["SecondaryDns"] = o.SecondaryDns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IppoolIpV4Config) UnmarshalJSON(data []byte) (err error) {
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
	type IppoolIpV4ConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address of the default IPv4 gateway.
		Gateway *string `json:"Gateway,omitempty" validate:"regexp=^$|^([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.
		Netmask *string `json:"Netmask,omitempty" validate:"regexp=^$|^(((255\\\\.){3}(255|254|252|248|240|224|192|128|0+))|((255\\\\.){2}(255|254|252|248|240|224|192|128|0+)\\\\.0)|((255\\\\.)(255|254|252|248|240|224|192|128|0+)(\\\\.0+){2})|((255|254|252|248|240|224|192|128|0+)(\\\\.0+){3}))$"`
		// IP Address of the primary Domain Name System (DNS) server.
		PrimaryDns *string `json:"PrimaryDns,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// IP Address of the secondary Domain Name System (DNS) server.
		SecondaryDns *string `json:"SecondaryDns,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	}

	varIppoolIpV4ConfigWithoutEmbeddedStruct := IppoolIpV4ConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIppoolIpV4ConfigWithoutEmbeddedStruct)
	if err == nil {
		varIppoolIpV4Config := _IppoolIpV4Config{}
		varIppoolIpV4Config.ClassId = varIppoolIpV4ConfigWithoutEmbeddedStruct.ClassId
		varIppoolIpV4Config.ObjectType = varIppoolIpV4ConfigWithoutEmbeddedStruct.ObjectType
		varIppoolIpV4Config.Gateway = varIppoolIpV4ConfigWithoutEmbeddedStruct.Gateway
		varIppoolIpV4Config.Netmask = varIppoolIpV4ConfigWithoutEmbeddedStruct.Netmask
		varIppoolIpV4Config.PrimaryDns = varIppoolIpV4ConfigWithoutEmbeddedStruct.PrimaryDns
		varIppoolIpV4Config.SecondaryDns = varIppoolIpV4ConfigWithoutEmbeddedStruct.SecondaryDns
		*o = IppoolIpV4Config(varIppoolIpV4Config)
	} else {
		return err
	}

	varIppoolIpV4Config := _IppoolIpV4Config{}

	err = json.Unmarshal(data, &varIppoolIpV4Config)
	if err == nil {
		o.MoBaseComplexType = varIppoolIpV4Config.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "PrimaryDns")
		delete(additionalProperties, "SecondaryDns")

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

type NullableIppoolIpV4Config struct {
	value *IppoolIpV4Config
	isSet bool
}

func (v NullableIppoolIpV4Config) Get() *IppoolIpV4Config {
	return v.value
}

func (v *NullableIppoolIpV4Config) Set(val *IppoolIpV4Config) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpV4Config) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpV4Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpV4Config(val *IppoolIpV4Config) *NullableIppoolIpV4Config {
	return &NullableIppoolIpV4Config{value: val, isSet: true}
}

func (v NullableIppoolIpV4Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpV4Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
