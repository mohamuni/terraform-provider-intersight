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

// MarketplaceUseCaseVersionResourcesAllOf Definition of the list of properties defined in 'marketplace.UseCaseVersionResources', excluding properties defined in parent classes.
type MarketplaceUseCaseVersionResourcesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A string ID for each use case
	ResourceId *string `json:"ResourceId,omitempty"`
	// A string resource type for each use case
	ResourceType         *string `json:"ResourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseVersionResourcesAllOf MarketplaceUseCaseVersionResourcesAllOf

// NewMarketplaceUseCaseVersionResourcesAllOf instantiates a new MarketplaceUseCaseVersionResourcesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseVersionResourcesAllOf(classId string, objectType string) *MarketplaceUseCaseVersionResourcesAllOf {
	this := MarketplaceUseCaseVersionResourcesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseVersionResourcesAllOfWithDefaults instantiates a new MarketplaceUseCaseVersionResourcesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseVersionResourcesAllOfWithDefaults() *MarketplaceUseCaseVersionResourcesAllOf {
	this := MarketplaceUseCaseVersionResourcesAllOf{}
	var classId string = "marketplace.UseCaseVersionResources"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseVersionResources"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseVersionResourcesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseVersionResourcesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *MarketplaceUseCaseVersionResourcesAllOf) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *MarketplaceUseCaseVersionResourcesAllOf) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResourcesAllOf) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *MarketplaceUseCaseVersionResourcesAllOf) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *MarketplaceUseCaseVersionResourcesAllOf) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o MarketplaceUseCaseVersionResourcesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ResourceId != nil {
		toSerialize["ResourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["ResourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarketplaceUseCaseVersionResourcesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMarketplaceUseCaseVersionResourcesAllOf := _MarketplaceUseCaseVersionResourcesAllOf{}

	if err = json.Unmarshal(bytes, &varMarketplaceUseCaseVersionResourcesAllOf); err == nil {
		*o = MarketplaceUseCaseVersionResourcesAllOf(varMarketplaceUseCaseVersionResourcesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ResourceId")
		delete(additionalProperties, "ResourceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketplaceUseCaseVersionResourcesAllOf struct {
	value *MarketplaceUseCaseVersionResourcesAllOf
	isSet bool
}

func (v NullableMarketplaceUseCaseVersionResourcesAllOf) Get() *MarketplaceUseCaseVersionResourcesAllOf {
	return v.value
}

func (v *NullableMarketplaceUseCaseVersionResourcesAllOf) Set(val *MarketplaceUseCaseVersionResourcesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseVersionResourcesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseVersionResourcesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseVersionResourcesAllOf(val *MarketplaceUseCaseVersionResourcesAllOf) *NullableMarketplaceUseCaseVersionResourcesAllOf {
	return &NullableMarketplaceUseCaseVersionResourcesAllOf{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseVersionResourcesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseVersionResourcesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}