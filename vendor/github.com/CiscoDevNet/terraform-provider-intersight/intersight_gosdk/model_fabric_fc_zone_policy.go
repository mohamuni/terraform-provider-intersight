/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricFcZonePolicy List of target path entries of storage arrays that are used to configure zones on the switch.
type FabricFcZonePolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string               `json:"ObjectType"`
	FcTargetMembers []FabricFcZoneMember `json:"FcTargetMembers,omitempty"`
	// Type of FC zoning. Allowed values are SIST, SIMT and None. * `SIST` - The system automatically creates one zone for each vHBA and storage port pair. Each zone has two members. * `SIMT` - The system automatically creates one zone for each vHBA. Configure this type of zoning if the number of zones created is likely to exceed the maximum supported number of zones. * `None` - FC zoning is not configured.
	FcTargetZoningType   *string                               `json:"FcTargetZoningType,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcZonePolicy FabricFcZonePolicy

// NewFabricFcZonePolicy instantiates a new FabricFcZonePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcZonePolicy(classId string, objectType string) *FabricFcZonePolicy {
	this := FabricFcZonePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fcTargetZoningType string = "SIST"
	this.FcTargetZoningType = &fcTargetZoningType
	return &this
}

// NewFabricFcZonePolicyWithDefaults instantiates a new FabricFcZonePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcZonePolicyWithDefaults() *FabricFcZonePolicy {
	this := FabricFcZonePolicy{}
	var classId string = "fabric.FcZonePolicy"
	this.ClassId = classId
	var objectType string = "fabric.FcZonePolicy"
	this.ObjectType = objectType
	var fcTargetZoningType string = "SIST"
	this.FcTargetZoningType = &fcTargetZoningType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcZonePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcZonePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcZonePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcZonePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcZonePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcZonePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFcTargetMembers returns the FcTargetMembers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricFcZonePolicy) GetFcTargetMembers() []FabricFcZoneMember {
	if o == nil {
		var ret []FabricFcZoneMember
		return ret
	}
	return o.FcTargetMembers
}

// GetFcTargetMembersOk returns a tuple with the FcTargetMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricFcZonePolicy) GetFcTargetMembersOk() ([]FabricFcZoneMember, bool) {
	if o == nil || o.FcTargetMembers == nil {
		return nil, false
	}
	return o.FcTargetMembers, true
}

// HasFcTargetMembers returns a boolean if a field has been set.
func (o *FabricFcZonePolicy) HasFcTargetMembers() bool {
	if o != nil && o.FcTargetMembers != nil {
		return true
	}

	return false
}

// SetFcTargetMembers gets a reference to the given []FabricFcZoneMember and assigns it to the FcTargetMembers field.
func (o *FabricFcZonePolicy) SetFcTargetMembers(v []FabricFcZoneMember) {
	o.FcTargetMembers = v
}

// GetFcTargetZoningType returns the FcTargetZoningType field value if set, zero value otherwise.
func (o *FabricFcZonePolicy) GetFcTargetZoningType() string {
	if o == nil || o.FcTargetZoningType == nil {
		var ret string
		return ret
	}
	return *o.FcTargetZoningType
}

// GetFcTargetZoningTypeOk returns a tuple with the FcTargetZoningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcZonePolicy) GetFcTargetZoningTypeOk() (*string, bool) {
	if o == nil || o.FcTargetZoningType == nil {
		return nil, false
	}
	return o.FcTargetZoningType, true
}

// HasFcTargetZoningType returns a boolean if a field has been set.
func (o *FabricFcZonePolicy) HasFcTargetZoningType() bool {
	if o != nil && o.FcTargetZoningType != nil {
		return true
	}

	return false
}

// SetFcTargetZoningType gets a reference to the given string and assigns it to the FcTargetZoningType field.
func (o *FabricFcZonePolicy) SetFcTargetZoningType(v string) {
	o.FcTargetZoningType = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricFcZonePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcZonePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricFcZonePolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricFcZonePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricFcZonePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FcTargetMembers != nil {
		toSerialize["FcTargetMembers"] = o.FcTargetMembers
	}
	if o.FcTargetZoningType != nil {
		toSerialize["FcTargetZoningType"] = o.FcTargetZoningType
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcZonePolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcZonePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string               `json:"ObjectType"`
		FcTargetMembers []FabricFcZoneMember `json:"FcTargetMembers,omitempty"`
		// Type of FC zoning. Allowed values are SIST, SIMT and None. * `SIST` - The system automatically creates one zone for each vHBA and storage port pair. Each zone has two members. * `SIMT` - The system automatically creates one zone for each vHBA. Configure this type of zoning if the number of zones created is likely to exceed the maximum supported number of zones. * `None` - FC zoning is not configured.
		FcTargetZoningType *string                               `json:"FcTargetZoningType,omitempty"`
		Organization       *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varFabricFcZonePolicyWithoutEmbeddedStruct := FabricFcZonePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcZonePolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcZonePolicy := _FabricFcZonePolicy{}
		varFabricFcZonePolicy.ClassId = varFabricFcZonePolicyWithoutEmbeddedStruct.ClassId
		varFabricFcZonePolicy.ObjectType = varFabricFcZonePolicyWithoutEmbeddedStruct.ObjectType
		varFabricFcZonePolicy.FcTargetMembers = varFabricFcZonePolicyWithoutEmbeddedStruct.FcTargetMembers
		varFabricFcZonePolicy.FcTargetZoningType = varFabricFcZonePolicyWithoutEmbeddedStruct.FcTargetZoningType
		varFabricFcZonePolicy.Organization = varFabricFcZonePolicyWithoutEmbeddedStruct.Organization
		*o = FabricFcZonePolicy(varFabricFcZonePolicy)
	} else {
		return err
	}

	varFabricFcZonePolicy := _FabricFcZonePolicy{}

	err = json.Unmarshal(bytes, &varFabricFcZonePolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricFcZonePolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FcTargetMembers")
		delete(additionalProperties, "FcTargetZoningType")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableFabricFcZonePolicy struct {
	value *FabricFcZonePolicy
	isSet bool
}

func (v NullableFabricFcZonePolicy) Get() *FabricFcZonePolicy {
	return v.value
}

func (v *NullableFabricFcZonePolicy) Set(val *FabricFcZonePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcZonePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcZonePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcZonePolicy(val *FabricFcZonePolicy) *NullableFabricFcZonePolicy {
	return &NullableFabricFcZonePolicy{value: val, isSet: true}
}

func (v NullableFabricFcZonePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcZonePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}