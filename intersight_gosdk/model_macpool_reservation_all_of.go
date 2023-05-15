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
)

// MacpoolReservationAllOf Definition of the list of properties defined in 'macpool.Reservation', excluding properties defined in parent classes.
type MacpoolReservationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MAC identity to be reserved.
	Identity             *string                               `json:"Identity,omitempty"`
	MemberOf             []MacpoolMemberOf                     `json:"MemberOf,omitempty"`
	BlockHead            *MacpoolIdBlockRelationship           `json:"BlockHead,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Pool                 *MacpoolPoolRelationship              `json:"Pool,omitempty"`
	PoolMember           *MacpoolPoolMemberRelationship        `json:"PoolMember,omitempty"`
	Universe             *MacpoolUniverseRelationship          `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolReservationAllOf MacpoolReservationAllOf

// NewMacpoolReservationAllOf instantiates a new MacpoolReservationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolReservationAllOf(classId string, objectType string) *MacpoolReservationAllOf {
	this := MacpoolReservationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMacpoolReservationAllOfWithDefaults instantiates a new MacpoolReservationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolReservationAllOfWithDefaults() *MacpoolReservationAllOf {
	this := MacpoolReservationAllOf{}
	var classId string = "macpool.Reservation"
	this.ClassId = classId
	var objectType string = "macpool.Reservation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolReservationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolReservationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolReservationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolReservationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MacpoolReservationAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *MacpoolReservationAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetMemberOf returns the MemberOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolReservationAllOf) GetMemberOf() []MacpoolMemberOf {
	if o == nil {
		var ret []MacpoolMemberOf
		return ret
	}
	return o.MemberOf
}

// GetMemberOfOk returns a tuple with the MemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolReservationAllOf) GetMemberOfOk() ([]MacpoolMemberOf, bool) {
	if o == nil || o.MemberOf == nil {
		return nil, false
	}
	return o.MemberOf, true
}

// HasMemberOf returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasMemberOf() bool {
	if o != nil && o.MemberOf != nil {
		return true
	}

	return false
}

// SetMemberOf gets a reference to the given []MacpoolMemberOf and assigns it to the MemberOf field.
func (o *MacpoolReservationAllOf) SetMemberOf(v []MacpoolMemberOf) {
	o.MemberOf = v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *MacpoolReservationAllOf) GetBlockHead() MacpoolIdBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret MacpoolIdBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetBlockHeadOk() (*MacpoolIdBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given MacpoolIdBlockRelationship and assigns it to the BlockHead field.
func (o *MacpoolReservationAllOf) SetBlockHead(v MacpoolIdBlockRelationship) {
	o.BlockHead = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *MacpoolReservationAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *MacpoolReservationAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *MacpoolReservationAllOf) GetPool() MacpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given MacpoolPoolRelationship and assigns it to the Pool field.
func (o *MacpoolReservationAllOf) SetPool(v MacpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *MacpoolReservationAllOf) GetPoolMember() MacpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret MacpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetPoolMemberOk() (*MacpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given MacpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *MacpoolReservationAllOf) SetPoolMember(v MacpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *MacpoolReservationAllOf) GetUniverse() MacpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret MacpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationAllOf) GetUniverseOk() (*MacpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *MacpoolReservationAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given MacpoolUniverseRelationship and assigns it to the Universe field.
func (o *MacpoolReservationAllOf) SetUniverse(v MacpoolUniverseRelationship) {
	o.Universe = &v
}

func (o MacpoolReservationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.MemberOf != nil {
		toSerialize["MemberOf"] = o.MemberOf
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MacpoolReservationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMacpoolReservationAllOf := _MacpoolReservationAllOf{}

	if err = json.Unmarshal(bytes, &varMacpoolReservationAllOf); err == nil {
		*o = MacpoolReservationAllOf(varMacpoolReservationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "MemberOf")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMacpoolReservationAllOf struct {
	value *MacpoolReservationAllOf
	isSet bool
}

func (v NullableMacpoolReservationAllOf) Get() *MacpoolReservationAllOf {
	return v.value
}

func (v *NullableMacpoolReservationAllOf) Set(val *MacpoolReservationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolReservationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolReservationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolReservationAllOf(val *MacpoolReservationAllOf) *NullableMacpoolReservationAllOf {
	return &NullableMacpoolReservationAllOf{value: val, isSet: true}
}

func (v NullableMacpoolReservationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolReservationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}