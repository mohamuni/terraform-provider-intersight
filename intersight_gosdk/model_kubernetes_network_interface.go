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

// KubernetesNetworkInterface Abstract Network Interface Configuration.
type KubernetesNetworkInterface struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string   `json:"ObjectType"`
	Addresses  []string `json:"Addresses,omitempty"`
	// Deprecated. This will add a default route as long as the first default route in Routes is not different. If is different, Gateway will be replaced with that default route. If there is no default Route and this is set, then Routes will be updated with the first entry as a default with this default gateway. If there is only one default Route and this gateway becomes empty, then the default routes will all be removed. Do not set if using Ip Pools, as the gateway is configured in the pool. This will be removed in the future.
	// Deprecated
	Gateway     *string                `json:"Gateway,omitempty"`
	IpV4Configs []KubernetesIpV4Config `json:"IpV4Configs,omitempty"`
	// The MTU to assign to this Network Interface.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Name for this network interface.
	Name                 *string           `json:"Name,omitempty"`
	Routes               []KubernetesRoute `json:"Routes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNetworkInterface KubernetesNetworkInterface

// NewKubernetesNetworkInterface instantiates a new KubernetesNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNetworkInterface(classId string, objectType string) *KubernetesNetworkInterface {
	this := KubernetesNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNetworkInterfaceWithDefaults instantiates a new KubernetesNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNetworkInterfaceWithDefaults() *KubernetesNetworkInterface {
	this := KubernetesNetworkInterface{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNetworkInterface) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNetworkInterface) GetAddressesOk() ([]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *KubernetesNetworkInterface) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *KubernetesNetworkInterface) SetAddresses(v []string) {
	o.Addresses = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
// Deprecated
func (o *KubernetesNetworkInterface) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *KubernetesNetworkInterface) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *KubernetesNetworkInterface) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
// Deprecated
func (o *KubernetesNetworkInterface) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpV4Configs returns the IpV4Configs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNetworkInterface) GetIpV4Configs() []KubernetesIpV4Config {
	if o == nil {
		var ret []KubernetesIpV4Config
		return ret
	}
	return o.IpV4Configs
}

// GetIpV4ConfigsOk returns a tuple with the IpV4Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNetworkInterface) GetIpV4ConfigsOk() ([]KubernetesIpV4Config, bool) {
	if o == nil || o.IpV4Configs == nil {
		return nil, false
	}
	return o.IpV4Configs, true
}

// HasIpV4Configs returns a boolean if a field has been set.
func (o *KubernetesNetworkInterface) HasIpV4Configs() bool {
	if o != nil && o.IpV4Configs != nil {
		return true
	}

	return false
}

// SetIpV4Configs gets a reference to the given []KubernetesIpV4Config and assigns it to the IpV4Configs field.
func (o *KubernetesNetworkInterface) SetIpV4Configs(v []KubernetesIpV4Config) {
	o.IpV4Configs = v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *KubernetesNetworkInterface) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterface) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *KubernetesNetworkInterface) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *KubernetesNetworkInterface) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesNetworkInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNetworkInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesNetworkInterface) SetName(v string) {
	o.Name = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNetworkInterface) GetRoutes() []KubernetesRoute {
	if o == nil {
		var ret []KubernetesRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNetworkInterface) GetRoutesOk() ([]KubernetesRoute, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *KubernetesNetworkInterface) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []KubernetesRoute and assigns it to the Routes field.
func (o *KubernetesNetworkInterface) SetRoutes(v []KubernetesRoute) {
	o.Routes = v
}

func (o KubernetesNetworkInterface) MarshalJSON() ([]byte, error) {
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
	if o.Addresses != nil {
		toSerialize["Addresses"] = o.Addresses
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.IpV4Configs != nil {
		toSerialize["IpV4Configs"] = o.IpV4Configs
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Routes != nil {
		toSerialize["Routes"] = o.Routes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNetworkInterface) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string   `json:"ObjectType"`
		Addresses  []string `json:"Addresses,omitempty"`
		// Deprecated. This will add a default route as long as the first default route in Routes is not different. If is different, Gateway will be replaced with that default route. If there is no default Route and this is set, then Routes will be updated with the first entry as a default with this default gateway. If there is only one default Route and this gateway becomes empty, then the default routes will all be removed. Do not set if using Ip Pools, as the gateway is configured in the pool. This will be removed in the future.
		// Deprecated
		Gateway     *string                `json:"Gateway,omitempty"`
		IpV4Configs []KubernetesIpV4Config `json:"IpV4Configs,omitempty"`
		// The MTU to assign to this Network Interface.
		Mtu *int64 `json:"Mtu,omitempty"`
		// Name for this network interface.
		Name   *string           `json:"Name,omitempty"`
		Routes []KubernetesRoute `json:"Routes,omitempty"`
	}

	varKubernetesNetworkInterfaceWithoutEmbeddedStruct := KubernetesNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNetworkInterface := _KubernetesNetworkInterface{}
		varKubernetesNetworkInterface.ClassId = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varKubernetesNetworkInterface.ObjectType = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varKubernetesNetworkInterface.Addresses = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.Addresses
		varKubernetesNetworkInterface.Gateway = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.Gateway
		varKubernetesNetworkInterface.IpV4Configs = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.IpV4Configs
		varKubernetesNetworkInterface.Mtu = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.Mtu
		varKubernetesNetworkInterface.Name = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.Name
		varKubernetesNetworkInterface.Routes = varKubernetesNetworkInterfaceWithoutEmbeddedStruct.Routes
		*o = KubernetesNetworkInterface(varKubernetesNetworkInterface)
	} else {
		return err
	}

	varKubernetesNetworkInterface := _KubernetesNetworkInterface{}

	err = json.Unmarshal(bytes, &varKubernetesNetworkInterface)
	if err == nil {
		o.MoBaseComplexType = varKubernetesNetworkInterface.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Addresses")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "IpV4Configs")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Routes")

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

type NullableKubernetesNetworkInterface struct {
	value *KubernetesNetworkInterface
	isSet bool
}

func (v NullableKubernetesNetworkInterface) Get() *KubernetesNetworkInterface {
	return v.value
}

func (v *NullableKubernetesNetworkInterface) Set(val *KubernetesNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNetworkInterface(val *KubernetesNetworkInterface) *NullableKubernetesNetworkInterface {
	return &NullableKubernetesNetworkInterface{value: val, isSet: true}
}

func (v NullableKubernetesNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
