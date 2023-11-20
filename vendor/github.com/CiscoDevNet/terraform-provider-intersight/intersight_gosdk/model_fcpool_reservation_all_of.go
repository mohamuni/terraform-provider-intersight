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

// FcpoolReservationAllOf Definition of the list of properties defined in 'fcpool.Reservation', excluding properties defined in parent classes.
type FcpoolReservationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Purpose of this WWN ID. Purpose can be WWPN or WWNN.
	IdPurpose *string `json:"IdPurpose,omitempty"`
	// WWN ID that needs to be reserved.
	Identity             *string                               `json:"Identity,omitempty"`
	Block                *FcpoolFcBlockRelationship            `json:"Block,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Pool                 *FcpoolPoolRelationship               `json:"Pool,omitempty"`
	PoolMember           *FcpoolPoolMemberRelationship         `json:"PoolMember,omitempty"`
	Universe             *FcpoolUniverseRelationship           `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolReservationAllOf FcpoolReservationAllOf

// NewFcpoolReservationAllOf instantiates a new FcpoolReservationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolReservationAllOf(classId string, objectType string) *FcpoolReservationAllOf {
	this := FcpoolReservationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcpoolReservationAllOfWithDefaults instantiates a new FcpoolReservationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolReservationAllOfWithDefaults() *FcpoolReservationAllOf {
	this := FcpoolReservationAllOf{}
	var classId string = "fcpool.Reservation"
	this.ClassId = classId
	var objectType string = "fcpool.Reservation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolReservationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolReservationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolReservationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolReservationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdPurpose returns the IdPurpose field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetIdPurpose() string {
	if o == nil || o.IdPurpose == nil {
		var ret string
		return ret
	}
	return *o.IdPurpose
}

// GetIdPurposeOk returns a tuple with the IdPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetIdPurposeOk() (*string, bool) {
	if o == nil || o.IdPurpose == nil {
		return nil, false
	}
	return o.IdPurpose, true
}

// HasIdPurpose returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasIdPurpose() bool {
	if o != nil && o.IdPurpose != nil {
		return true
	}

	return false
}

// SetIdPurpose gets a reference to the given string and assigns it to the IdPurpose field.
func (o *FcpoolReservationAllOf) SetIdPurpose(v string) {
	o.IdPurpose = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *FcpoolReservationAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetBlock() FcpoolFcBlockRelationship {
	if o == nil || o.Block == nil {
		var ret FcpoolFcBlockRelationship
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetBlockOk() (*FcpoolFcBlockRelationship, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasBlock() bool {
	if o != nil && o.Block != nil {
		return true
	}

	return false
}

// SetBlock gets a reference to the given FcpoolFcBlockRelationship and assigns it to the Block field.
func (o *FcpoolReservationAllOf) SetBlock(v FcpoolFcBlockRelationship) {
	o.Block = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FcpoolReservationAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolReservationAllOf) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetPoolMember() FcpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret FcpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given FcpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *FcpoolReservationAllOf) SetPoolMember(v FcpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *FcpoolReservationAllOf) GetUniverse() FcpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret FcpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolReservationAllOf) GetUniverseOk() (*FcpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *FcpoolReservationAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given FcpoolUniverseRelationship and assigns it to the Universe field.
func (o *FcpoolReservationAllOf) SetUniverse(v FcpoolUniverseRelationship) {
	o.Universe = &v
}

func (o FcpoolReservationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IdPurpose != nil {
		toSerialize["IdPurpose"] = o.IdPurpose
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Block != nil {
		toSerialize["Block"] = o.Block
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

func (o *FcpoolReservationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcpoolReservationAllOf := _FcpoolReservationAllOf{}

	if err = json.Unmarshal(bytes, &varFcpoolReservationAllOf); err == nil {
		*o = FcpoolReservationAllOf(varFcpoolReservationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdPurpose")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Block")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcpoolReservationAllOf struct {
	value *FcpoolReservationAllOf
	isSet bool
}

func (v NullableFcpoolReservationAllOf) Get() *FcpoolReservationAllOf {
	return v.value
}

func (v *NullableFcpoolReservationAllOf) Set(val *FcpoolReservationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolReservationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolReservationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolReservationAllOf(val *FcpoolReservationAllOf) *NullableFcpoolReservationAllOf {
	return &NullableFcpoolReservationAllOf{value: val, isSet: true}
}

func (v NullableFcpoolReservationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolReservationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
