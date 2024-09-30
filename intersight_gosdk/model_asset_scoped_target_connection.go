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

// checks if the AssetScopedTargetConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetScopedTargetConnection{}

// AssetScopedTargetConnection ScopedTargetConnection provides the necessary details for Intersight to connect to and authenticate with a managed scoped target such as msSQL, mySQL and Oracle Database etc.
type AssetScopedTargetConnection struct {
	AssetConnection
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When this flag is set to true, every IWO entity in the scope targets will be checked and discovery of the scope target will be regarded as a failure when anyone of these entities cannot be connected and validated.
	FullValidation *bool `json:"FullValidation,omitempty"`
	// Indicates whether a connection to the target should be established using TLS.
	IsSecure *bool `json:"IsSecure,omitempty"`
	// The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
	Port *int64 `json:"Port,omitempty"`
	// The group id of IWO entities upon which the discover of a scoped target is performed. For example, a database target may be scoped to the group of virtual machines upon which the database application is running. Scope value is group id created for all those virtual machines in this scope.
	Scope                *string `json:"Scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetScopedTargetConnection AssetScopedTargetConnection

// NewAssetScopedTargetConnection instantiates a new AssetScopedTargetConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetScopedTargetConnection(classId string, objectType string) *AssetScopedTargetConnection {
	this := AssetScopedTargetConnection{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isSecure bool = true
	this.IsSecure = &isSecure
	return &this
}

// NewAssetScopedTargetConnectionWithDefaults instantiates a new AssetScopedTargetConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetScopedTargetConnectionWithDefaults() *AssetScopedTargetConnection {
	this := AssetScopedTargetConnection{}
	var classId string = "asset.ScopedTargetConnection"
	this.ClassId = classId
	var objectType string = "asset.ScopedTargetConnection"
	this.ObjectType = objectType
	var isSecure bool = true
	this.IsSecure = &isSecure
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetScopedTargetConnection) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnection) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetScopedTargetConnection) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.ScopedTargetConnection" of the ClassId field.
func (o *AssetScopedTargetConnection) GetDefaultClassId() interface{} {
	return "asset.ScopedTargetConnection"
}

// GetObjectType returns the ObjectType field value
func (o *AssetScopedTargetConnection) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnection) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetScopedTargetConnection) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.ScopedTargetConnection" of the ObjectType field.
func (o *AssetScopedTargetConnection) GetDefaultObjectType() interface{} {
	return "asset.ScopedTargetConnection"
}

// GetFullValidation returns the FullValidation field value if set, zero value otherwise.
func (o *AssetScopedTargetConnection) GetFullValidation() bool {
	if o == nil || IsNil(o.FullValidation) {
		var ret bool
		return ret
	}
	return *o.FullValidation
}

// GetFullValidationOk returns a tuple with the FullValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnection) GetFullValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.FullValidation) {
		return nil, false
	}
	return o.FullValidation, true
}

// HasFullValidation returns a boolean if a field has been set.
func (o *AssetScopedTargetConnection) HasFullValidation() bool {
	if o != nil && !IsNil(o.FullValidation) {
		return true
	}

	return false
}

// SetFullValidation gets a reference to the given bool and assigns it to the FullValidation field.
func (o *AssetScopedTargetConnection) SetFullValidation(v bool) {
	o.FullValidation = &v
}

// GetIsSecure returns the IsSecure field value if set, zero value otherwise.
func (o *AssetScopedTargetConnection) GetIsSecure() bool {
	if o == nil || IsNil(o.IsSecure) {
		var ret bool
		return ret
	}
	return *o.IsSecure
}

// GetIsSecureOk returns a tuple with the IsSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnection) GetIsSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecure) {
		return nil, false
	}
	return o.IsSecure, true
}

// HasIsSecure returns a boolean if a field has been set.
func (o *AssetScopedTargetConnection) HasIsSecure() bool {
	if o != nil && !IsNil(o.IsSecure) {
		return true
	}

	return false
}

// SetIsSecure gets a reference to the given bool and assigns it to the IsSecure field.
func (o *AssetScopedTargetConnection) SetIsSecure(v bool) {
	o.IsSecure = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AssetScopedTargetConnection) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnection) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AssetScopedTargetConnection) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *AssetScopedTargetConnection) SetPort(v int64) {
	o.Port = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AssetScopedTargetConnection) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnection) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AssetScopedTargetConnection) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AssetScopedTargetConnection) SetScope(v string) {
	o.Scope = &v
}

func (o AssetScopedTargetConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetScopedTargetConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetConnection, errAssetConnection := json.Marshal(o.AssetConnection)
	if errAssetConnection != nil {
		return map[string]interface{}{}, errAssetConnection
	}
	errAssetConnection = json.Unmarshal([]byte(serializedAssetConnection), &toSerialize)
	if errAssetConnection != nil {
		return map[string]interface{}{}, errAssetConnection
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FullValidation) {
		toSerialize["FullValidation"] = o.FullValidation
	}
	if !IsNil(o.IsSecure) {
		toSerialize["IsSecure"] = o.IsSecure
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	if !IsNil(o.Scope) {
		toSerialize["Scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetScopedTargetConnection) UnmarshalJSON(data []byte) (err error) {
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
	type AssetScopedTargetConnectionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When this flag is set to true, every IWO entity in the scope targets will be checked and discovery of the scope target will be regarded as a failure when anyone of these entities cannot be connected and validated.
		FullValidation *bool `json:"FullValidation,omitempty"`
		// Indicates whether a connection to the target should be established using TLS.
		IsSecure *bool `json:"IsSecure,omitempty"`
		// The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
		Port *int64 `json:"Port,omitempty"`
		// The group id of IWO entities upon which the discover of a scoped target is performed. For example, a database target may be scoped to the group of virtual machines upon which the database application is running. Scope value is group id created for all those virtual machines in this scope.
		Scope *string `json:"Scope,omitempty"`
	}

	varAssetScopedTargetConnectionWithoutEmbeddedStruct := AssetScopedTargetConnectionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetScopedTargetConnectionWithoutEmbeddedStruct)
	if err == nil {
		varAssetScopedTargetConnection := _AssetScopedTargetConnection{}
		varAssetScopedTargetConnection.ClassId = varAssetScopedTargetConnectionWithoutEmbeddedStruct.ClassId
		varAssetScopedTargetConnection.ObjectType = varAssetScopedTargetConnectionWithoutEmbeddedStruct.ObjectType
		varAssetScopedTargetConnection.FullValidation = varAssetScopedTargetConnectionWithoutEmbeddedStruct.FullValidation
		varAssetScopedTargetConnection.IsSecure = varAssetScopedTargetConnectionWithoutEmbeddedStruct.IsSecure
		varAssetScopedTargetConnection.Port = varAssetScopedTargetConnectionWithoutEmbeddedStruct.Port
		varAssetScopedTargetConnection.Scope = varAssetScopedTargetConnectionWithoutEmbeddedStruct.Scope
		*o = AssetScopedTargetConnection(varAssetScopedTargetConnection)
	} else {
		return err
	}

	varAssetScopedTargetConnection := _AssetScopedTargetConnection{}

	err = json.Unmarshal(data, &varAssetScopedTargetConnection)
	if err == nil {
		o.AssetConnection = varAssetScopedTargetConnection.AssetConnection
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FullValidation")
		delete(additionalProperties, "IsSecure")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Scope")

		// remove fields from embedded structs
		reflectAssetConnection := reflect.ValueOf(o.AssetConnection)
		for i := 0; i < reflectAssetConnection.Type().NumField(); i++ {
			t := reflectAssetConnection.Type().Field(i)

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

type NullableAssetScopedTargetConnection struct {
	value *AssetScopedTargetConnection
	isSet bool
}

func (v NullableAssetScopedTargetConnection) Get() *AssetScopedTargetConnection {
	return v.value
}

func (v *NullableAssetScopedTargetConnection) Set(val *AssetScopedTargetConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetScopedTargetConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetScopedTargetConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetScopedTargetConnection(val *AssetScopedTargetConnection) *NullableAssetScopedTargetConnection {
	return &NullableAssetScopedTargetConnection{value: val, isSet: true}
}

func (v NullableAssetScopedTargetConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetScopedTargetConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
