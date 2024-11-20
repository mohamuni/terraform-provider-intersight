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

// checks if the HyperflexHxNetworkAddressDt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexHxNetworkAddressDt{}

// HyperflexHxNetworkAddressDt A network address. If both IP and FQDN are specified, IP is used. If only IP is specified, IP is used. If only FQDN is specified, FQDN is used. If both are not specified, it is an error.
type HyperflexHxNetworkAddressDt struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The network address as an FQDN or IPv4 address.
	Address *string `json:"Address,omitempty"`
	// The fully qualified domain name for the network address.
	Fqdn *string `json:"Fqdn,omitempty"`
	// The network address as an IPv4 address.
	Ip                   *string `json:"Ip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxNetworkAddressDt HyperflexHxNetworkAddressDt

// NewHyperflexHxNetworkAddressDt instantiates a new HyperflexHxNetworkAddressDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxNetworkAddressDt(classId string, objectType string) *HyperflexHxNetworkAddressDt {
	this := HyperflexHxNetworkAddressDt{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxNetworkAddressDtWithDefaults instantiates a new HyperflexHxNetworkAddressDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxNetworkAddressDtWithDefaults() *HyperflexHxNetworkAddressDt {
	this := HyperflexHxNetworkAddressDt{}
	var classId string = "hyperflex.HxNetworkAddressDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxNetworkAddressDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxNetworkAddressDt) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxNetworkAddressDt) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.HxNetworkAddressDt" of the ClassId field.
func (o *HyperflexHxNetworkAddressDt) GetDefaultClassId() interface{} {
	return "hyperflex.HxNetworkAddressDt"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxNetworkAddressDt) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxNetworkAddressDt) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.HxNetworkAddressDt" of the ObjectType field.
func (o *HyperflexHxNetworkAddressDt) GetDefaultObjectType() interface{} {
	return "hyperflex.HxNetworkAddressDt"
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDt) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDt) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HyperflexHxNetworkAddressDt) SetAddress(v string) {
	o.Address = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDt) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDt) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *HyperflexHxNetworkAddressDt) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDt) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDt) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *HyperflexHxNetworkAddressDt) SetIp(v string) {
	o.Ip = &v
}

func (o HyperflexHxNetworkAddressDt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexHxNetworkAddressDt) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Address) {
		toSerialize["Address"] = o.Address
	}
	if !IsNil(o.Fqdn) {
		toSerialize["Fqdn"] = o.Fqdn
	}
	if !IsNil(o.Ip) {
		toSerialize["Ip"] = o.Ip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexHxNetworkAddressDt) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexHxNetworkAddressDtWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The network address as an FQDN or IPv4 address.
		Address *string `json:"Address,omitempty"`
		// The fully qualified domain name for the network address.
		Fqdn *string `json:"Fqdn,omitempty"`
		// The network address as an IPv4 address.
		Ip *string `json:"Ip,omitempty"`
	}

	varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct := HyperflexHxNetworkAddressDtWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxNetworkAddressDt := _HyperflexHxNetworkAddressDt{}
		varHyperflexHxNetworkAddressDt.ClassId = varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct.ClassId
		varHyperflexHxNetworkAddressDt.ObjectType = varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct.ObjectType
		varHyperflexHxNetworkAddressDt.Address = varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct.Address
		varHyperflexHxNetworkAddressDt.Fqdn = varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct.Fqdn
		varHyperflexHxNetworkAddressDt.Ip = varHyperflexHxNetworkAddressDtWithoutEmbeddedStruct.Ip
		*o = HyperflexHxNetworkAddressDt(varHyperflexHxNetworkAddressDt)
	} else {
		return err
	}

	varHyperflexHxNetworkAddressDt := _HyperflexHxNetworkAddressDt{}

	err = json.Unmarshal(data, &varHyperflexHxNetworkAddressDt)
	if err == nil {
		o.MoBaseComplexType = varHyperflexHxNetworkAddressDt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Fqdn")
		delete(additionalProperties, "Ip")

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

type NullableHyperflexHxNetworkAddressDt struct {
	value *HyperflexHxNetworkAddressDt
	isSet bool
}

func (v NullableHyperflexHxNetworkAddressDt) Get() *HyperflexHxNetworkAddressDt {
	return v.value
}

func (v *NullableHyperflexHxNetworkAddressDt) Set(val *HyperflexHxNetworkAddressDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxNetworkAddressDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxNetworkAddressDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxNetworkAddressDt(val *HyperflexHxNetworkAddressDt) *NullableHyperflexHxNetworkAddressDt {
	return &NullableHyperflexHxNetworkAddressDt{value: val, isSet: true}
}

func (v NullableHyperflexHxNetworkAddressDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxNetworkAddressDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
