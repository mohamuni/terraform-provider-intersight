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
)

// FirmwareFirmwareSummaryAllOf Definition of the list of properties defined in 'firmware.FirmwareSummary', excluding properties defined in parent classes.
type FirmwareFirmwareSummaryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version details at the bundle level for the each of server.
	BundleVersion         *string                      `json:"BundleVersion,omitempty"`
	ComponentsFwInventory []FirmwareFirmwareInventory  `json:"ComponentsFwInventory,omitempty"`
	Server                *ComputePhysicalRelationship `json:"Server,omitempty"`
	TargetMo              *MoBaseMoRelationship        `json:"TargetMo,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _FirmwareFirmwareSummaryAllOf FirmwareFirmwareSummaryAllOf

// NewFirmwareFirmwareSummaryAllOf instantiates a new FirmwareFirmwareSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareFirmwareSummaryAllOf(classId string, objectType string) *FirmwareFirmwareSummaryAllOf {
	this := FirmwareFirmwareSummaryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareFirmwareSummaryAllOfWithDefaults instantiates a new FirmwareFirmwareSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareFirmwareSummaryAllOfWithDefaults() *FirmwareFirmwareSummaryAllOf {
	this := FirmwareFirmwareSummaryAllOf{}
	var classId string = "firmware.FirmwareSummary"
	this.ClassId = classId
	var objectType string = "firmware.FirmwareSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareFirmwareSummaryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummaryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareFirmwareSummaryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareFirmwareSummaryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummaryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareFirmwareSummaryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *FirmwareFirmwareSummaryAllOf) GetBundleVersion() string {
	if o == nil || o.BundleVersion == nil {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummaryAllOf) GetBundleVersionOk() (*string, bool) {
	if o == nil || o.BundleVersion == nil {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummaryAllOf) HasBundleVersion() bool {
	if o != nil && o.BundleVersion != nil {
		return true
	}

	return false
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *FirmwareFirmwareSummaryAllOf) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetComponentsFwInventory returns the ComponentsFwInventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareFirmwareSummaryAllOf) GetComponentsFwInventory() []FirmwareFirmwareInventory {
	if o == nil {
		var ret []FirmwareFirmwareInventory
		return ret
	}
	return o.ComponentsFwInventory
}

// GetComponentsFwInventoryOk returns a tuple with the ComponentsFwInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareFirmwareSummaryAllOf) GetComponentsFwInventoryOk() ([]FirmwareFirmwareInventory, bool) {
	if o == nil || o.ComponentsFwInventory == nil {
		return nil, false
	}
	return o.ComponentsFwInventory, true
}

// HasComponentsFwInventory returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummaryAllOf) HasComponentsFwInventory() bool {
	if o != nil && o.ComponentsFwInventory != nil {
		return true
	}

	return false
}

// SetComponentsFwInventory gets a reference to the given []FirmwareFirmwareInventory and assigns it to the ComponentsFwInventory field.
func (o *FirmwareFirmwareSummaryAllOf) SetComponentsFwInventory(v []FirmwareFirmwareInventory) {
	o.ComponentsFwInventory = v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *FirmwareFirmwareSummaryAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummaryAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummaryAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareFirmwareSummaryAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *FirmwareFirmwareSummaryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummaryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummaryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *FirmwareFirmwareSummaryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o FirmwareFirmwareSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BundleVersion != nil {
		toSerialize["BundleVersion"] = o.BundleVersion
	}
	if o.ComponentsFwInventory != nil {
		toSerialize["ComponentsFwInventory"] = o.ComponentsFwInventory
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareFirmwareSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareFirmwareSummaryAllOf := _FirmwareFirmwareSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareFirmwareSummaryAllOf); err == nil {
		*o = FirmwareFirmwareSummaryAllOf(varFirmwareFirmwareSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BundleVersion")
		delete(additionalProperties, "ComponentsFwInventory")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareFirmwareSummaryAllOf struct {
	value *FirmwareFirmwareSummaryAllOf
	isSet bool
}

func (v NullableFirmwareFirmwareSummaryAllOf) Get() *FirmwareFirmwareSummaryAllOf {
	return v.value
}

func (v *NullableFirmwareFirmwareSummaryAllOf) Set(val *FirmwareFirmwareSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareFirmwareSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareFirmwareSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareFirmwareSummaryAllOf(val *FirmwareFirmwareSummaryAllOf) *NullableFirmwareFirmwareSummaryAllOf {
	return &NullableFirmwareFirmwareSummaryAllOf{value: val, isSet: true}
}

func (v NullableFirmwareFirmwareSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareFirmwareSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
