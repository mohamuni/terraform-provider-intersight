/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// MarketplaceUseCaseDependencyAllOf Definition of the list of properties defined in 'marketplace.UseCaseDependency', excluding properties defined in parent classes.
type MarketplaceUseCaseDependencyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The string field to hold the key name
	Name *string `json:"Name,omitempty"`
	// The string field to hold the value
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseDependencyAllOf MarketplaceUseCaseDependencyAllOf

// NewMarketplaceUseCaseDependencyAllOf instantiates a new MarketplaceUseCaseDependencyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseDependencyAllOf(classId string, objectType string) *MarketplaceUseCaseDependencyAllOf {
	this := MarketplaceUseCaseDependencyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseDependencyAllOfWithDefaults instantiates a new MarketplaceUseCaseDependencyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseDependencyAllOfWithDefaults() *MarketplaceUseCaseDependencyAllOf {
	this := MarketplaceUseCaseDependencyAllOf{}
	var classId string = "marketplace.UseCaseDependency"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseDependency"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseDependencyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseDependencyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseDependencyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseDependencyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseDependencyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseDependencyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketplaceUseCaseDependencyAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseDependencyAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketplaceUseCaseDependencyAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketplaceUseCaseDependencyAllOf) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MarketplaceUseCaseDependencyAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseDependencyAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MarketplaceUseCaseDependencyAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MarketplaceUseCaseDependencyAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o MarketplaceUseCaseDependencyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarketplaceUseCaseDependencyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMarketplaceUseCaseDependencyAllOf := _MarketplaceUseCaseDependencyAllOf{}

	if err = json.Unmarshal(bytes, &varMarketplaceUseCaseDependencyAllOf); err == nil {
		*o = MarketplaceUseCaseDependencyAllOf(varMarketplaceUseCaseDependencyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketplaceUseCaseDependencyAllOf struct {
	value *MarketplaceUseCaseDependencyAllOf
	isSet bool
}

func (v NullableMarketplaceUseCaseDependencyAllOf) Get() *MarketplaceUseCaseDependencyAllOf {
	return v.value
}

func (v *NullableMarketplaceUseCaseDependencyAllOf) Set(val *MarketplaceUseCaseDependencyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseDependencyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseDependencyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseDependencyAllOf(val *MarketplaceUseCaseDependencyAllOf) *NullableMarketplaceUseCaseDependencyAllOf {
	return &NullableMarketplaceUseCaseDependencyAllOf{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseDependencyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseDependencyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}