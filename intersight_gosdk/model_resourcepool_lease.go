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

// checks if the ResourcepoolLease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcepoolLease{}

// ResourcepoolLease Lease API invoked by passing resource pool, lease API will reserve or un-reserve the resource from the pool.
type ResourcepoolLease struct {
	PoolAbstractLease
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string             `json:"ObjectType"`
	Condition  []ResourceSelector `json:"Condition,omitempty"`
	// Lease opertion applied for the feature.
	Feature *string `json:"Feature,omitempty"`
	// Indicates whether a lease allocation is exclusive based on the Assigned Entity, if the AssignedEntity holds any lease then not allowed to create new lease later.
	IsExclusiveAtAssignedEntity *bool                               `json:"IsExclusiveAtAssignedEntity,omitempty"`
	LeaseParameters             NullableResourcepoolLeaseParameters `json:"LeaseParameters,omitempty"`
	Resource                    *MoBaseMo                           `json:"Resource,omitempty"`
	// The type of resource present in the pool, such as 'server' can be a RackUnit or Blade. * `None` - The resource cannot consider for Resource Pool. * `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.
	ResourceType         *string                                       `json:"ResourceType,omitempty"`
	AssignedToEntity     NullableMoBaseMoRelationship                  `json:"AssignedToEntity,omitempty"`
	LeasedResource       NullableResourcepoolLeaseResourceRelationship `json:"LeasedResource,omitempty"`
	Pool                 NullableResourcepoolPoolRelationship          `json:"Pool,omitempty"`
	PoolMember           NullableResourcepoolPoolMemberRelationship    `json:"PoolMember,omitempty"`
	Universe             NullableResourcepoolUniverseRelationship      `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolLease ResourcepoolLease

// NewResourcepoolLease instantiates a new ResourcepoolLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolLease(classId string, objectType string) *ResourcepoolLease {
	this := ResourcepoolLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	var hasDuplicate bool = false
	this.HasDuplicate = &hasDuplicate
	var isExclusiveAtAssignedEntity bool = false
	this.IsExclusiveAtAssignedEntity = &isExclusiveAtAssignedEntity
	var resourceType string = "None"
	this.ResourceType = &resourceType
	return &this
}

// NewResourcepoolLeaseWithDefaults instantiates a new ResourcepoolLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolLeaseWithDefaults() *ResourcepoolLease {
	this := ResourcepoolLease{}
	var classId string = "resourcepool.Lease"
	this.ClassId = classId
	var objectType string = "resourcepool.Lease"
	this.ObjectType = objectType
	var isExclusiveAtAssignedEntity bool = false
	this.IsExclusiveAtAssignedEntity = &isExclusiveAtAssignedEntity
	var resourceType string = "None"
	this.ResourceType = &resourceType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolLease) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resourcepool.Lease" of the ClassId field.
func (o *ResourcepoolLease) GetDefaultClassId() interface{} {
	return "resourcepool.Lease"
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resourcepool.Lease" of the ObjectType field.
func (o *ResourcepoolLease) GetDefaultObjectType() interface{} {
	return "resourcepool.Lease"
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetCondition() []ResourceSelector {
	if o == nil {
		var ret []ResourceSelector
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetConditionOk() ([]ResourceSelector, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given []ResourceSelector and assigns it to the Condition field.
func (o *ResourcepoolLease) SetCondition(v []ResourceSelector) {
	o.Condition = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *ResourcepoolLease) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLease) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *ResourcepoolLease) SetFeature(v string) {
	o.Feature = &v
}

// GetIsExclusiveAtAssignedEntity returns the IsExclusiveAtAssignedEntity field value if set, zero value otherwise.
func (o *ResourcepoolLease) GetIsExclusiveAtAssignedEntity() bool {
	if o == nil || IsNil(o.IsExclusiveAtAssignedEntity) {
		var ret bool
		return ret
	}
	return *o.IsExclusiveAtAssignedEntity
}

// GetIsExclusiveAtAssignedEntityOk returns a tuple with the IsExclusiveAtAssignedEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLease) GetIsExclusiveAtAssignedEntityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExclusiveAtAssignedEntity) {
		return nil, false
	}
	return o.IsExclusiveAtAssignedEntity, true
}

// HasIsExclusiveAtAssignedEntity returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasIsExclusiveAtAssignedEntity() bool {
	if o != nil && !IsNil(o.IsExclusiveAtAssignedEntity) {
		return true
	}

	return false
}

// SetIsExclusiveAtAssignedEntity gets a reference to the given bool and assigns it to the IsExclusiveAtAssignedEntity field.
func (o *ResourcepoolLease) SetIsExclusiveAtAssignedEntity(v bool) {
	o.IsExclusiveAtAssignedEntity = &v
}

// GetLeaseParameters returns the LeaseParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetLeaseParameters() ResourcepoolLeaseParameters {
	if o == nil || IsNil(o.LeaseParameters.Get()) {
		var ret ResourcepoolLeaseParameters
		return ret
	}
	return *o.LeaseParameters.Get()
}

// GetLeaseParametersOk returns a tuple with the LeaseParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetLeaseParametersOk() (*ResourcepoolLeaseParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeaseParameters.Get(), o.LeaseParameters.IsSet()
}

// HasLeaseParameters returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasLeaseParameters() bool {
	if o != nil && o.LeaseParameters.IsSet() {
		return true
	}

	return false
}

// SetLeaseParameters gets a reference to the given NullableResourcepoolLeaseParameters and assigns it to the LeaseParameters field.
func (o *ResourcepoolLease) SetLeaseParameters(v ResourcepoolLeaseParameters) {
	o.LeaseParameters.Set(&v)
}

// SetLeaseParametersNil sets the value for LeaseParameters to be an explicit nil
func (o *ResourcepoolLease) SetLeaseParametersNil() {
	o.LeaseParameters.Set(nil)
}

// UnsetLeaseParameters ensures that no value is present for LeaseParameters, not even an explicit nil
func (o *ResourcepoolLease) UnsetLeaseParameters() {
	o.LeaseParameters.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourcepoolLease) GetResource() MoBaseMo {
	if o == nil || IsNil(o.Resource) {
		var ret MoBaseMo
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLease) GetResourceOk() (*MoBaseMo, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMo and assigns it to the Resource field.
func (o *ResourcepoolLease) SetResource(v MoBaseMo) {
	o.Resource = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourcepoolLease) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLease) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourcepoolLease) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *ResourcepoolLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *ResourcepoolLease) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *ResourcepoolLease) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetLeasedResource returns the LeasedResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetLeasedResource() ResourcepoolLeaseResourceRelationship {
	if o == nil || IsNil(o.LeasedResource.Get()) {
		var ret ResourcepoolLeaseResourceRelationship
		return ret
	}
	return *o.LeasedResource.Get()
}

// GetLeasedResourceOk returns a tuple with the LeasedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetLeasedResourceOk() (*ResourcepoolLeaseResourceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeasedResource.Get(), o.LeasedResource.IsSet()
}

// HasLeasedResource returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasLeasedResource() bool {
	if o != nil && o.LeasedResource.IsSet() {
		return true
	}

	return false
}

// SetLeasedResource gets a reference to the given NullableResourcepoolLeaseResourceRelationship and assigns it to the LeasedResource field.
func (o *ResourcepoolLease) SetLeasedResource(v ResourcepoolLeaseResourceRelationship) {
	o.LeasedResource.Set(&v)
}

// SetLeasedResourceNil sets the value for LeasedResource to be an explicit nil
func (o *ResourcepoolLease) SetLeasedResourceNil() {
	o.LeasedResource.Set(nil)
}

// UnsetLeasedResource ensures that no value is present for LeasedResource, not even an explicit nil
func (o *ResourcepoolLease) UnsetLeasedResource() {
	o.LeasedResource.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetPool() ResourcepoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret ResourcepoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetPoolOk() (*ResourcepoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableResourcepoolPoolRelationship and assigns it to the Pool field.
func (o *ResourcepoolLease) SetPool(v ResourcepoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *ResourcepoolLease) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *ResourcepoolLease) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetPoolMember() ResourcepoolPoolMemberRelationship {
	if o == nil || IsNil(o.PoolMember.Get()) {
		var ret ResourcepoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember.Get()
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetPoolMemberOk() (*ResourcepoolPoolMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolMember.Get(), o.PoolMember.IsSet()
}

// HasPoolMember returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasPoolMember() bool {
	if o != nil && o.PoolMember.IsSet() {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given NullableResourcepoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *ResourcepoolLease) SetPoolMember(v ResourcepoolPoolMemberRelationship) {
	o.PoolMember.Set(&v)
}

// SetPoolMemberNil sets the value for PoolMember to be an explicit nil
func (o *ResourcepoolLease) SetPoolMemberNil() {
	o.PoolMember.Set(nil)
}

// UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
func (o *ResourcepoolLease) UnsetPoolMember() {
	o.PoolMember.Unset()
}

// GetUniverse returns the Universe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLease) GetUniverse() ResourcepoolUniverseRelationship {
	if o == nil || IsNil(o.Universe.Get()) {
		var ret ResourcepoolUniverseRelationship
		return ret
	}
	return *o.Universe.Get()
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLease) GetUniverseOk() (*ResourcepoolUniverseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Universe.Get(), o.Universe.IsSet()
}

// HasUniverse returns a boolean if a field has been set.
func (o *ResourcepoolLease) HasUniverse() bool {
	if o != nil && o.Universe.IsSet() {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given NullableResourcepoolUniverseRelationship and assigns it to the Universe field.
func (o *ResourcepoolLease) SetUniverse(v ResourcepoolUniverseRelationship) {
	o.Universe.Set(&v)
}

// SetUniverseNil sets the value for Universe to be an explicit nil
func (o *ResourcepoolLease) SetUniverseNil() {
	o.Universe.Set(nil)
}

// UnsetUniverse ensures that no value is present for Universe, not even an explicit nil
func (o *ResourcepoolLease) UnsetUniverse() {
	o.Universe.Unset()
}

func (o ResourcepoolLease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcepoolLease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractLease, errPoolAbstractLease := json.Marshal(o.PoolAbstractLease)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	errPoolAbstractLease = json.Unmarshal([]byte(serializedPoolAbstractLease), &toSerialize)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}
	if !IsNil(o.Feature) {
		toSerialize["Feature"] = o.Feature
	}
	if !IsNil(o.IsExclusiveAtAssignedEntity) {
		toSerialize["IsExclusiveAtAssignedEntity"] = o.IsExclusiveAtAssignedEntity
	}
	if o.LeaseParameters.IsSet() {
		toSerialize["LeaseParameters"] = o.LeaseParameters.Get()
	}
	if !IsNil(o.Resource) {
		toSerialize["Resource"] = o.Resource
	}
	if !IsNil(o.ResourceType) {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.LeasedResource.IsSet() {
		toSerialize["LeasedResource"] = o.LeasedResource.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.PoolMember.IsSet() {
		toSerialize["PoolMember"] = o.PoolMember.Get()
	}
	if o.Universe.IsSet() {
		toSerialize["Universe"] = o.Universe.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourcepoolLease) UnmarshalJSON(data []byte) (err error) {
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
	type ResourcepoolLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string             `json:"ObjectType"`
		Condition  []ResourceSelector `json:"Condition,omitempty"`
		// Lease opertion applied for the feature.
		Feature *string `json:"Feature,omitempty"`
		// Indicates whether a lease allocation is exclusive based on the Assigned Entity, if the AssignedEntity holds any lease then not allowed to create new lease later.
		IsExclusiveAtAssignedEntity *bool                               `json:"IsExclusiveAtAssignedEntity,omitempty"`
		LeaseParameters             NullableResourcepoolLeaseParameters `json:"LeaseParameters,omitempty"`
		Resource                    *MoBaseMo                           `json:"Resource,omitempty"`
		// The type of resource present in the pool, such as 'server' can be a RackUnit or Blade. * `None` - The resource cannot consider for Resource Pool. * `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.
		ResourceType     *string                                       `json:"ResourceType,omitempty"`
		AssignedToEntity NullableMoBaseMoRelationship                  `json:"AssignedToEntity,omitempty"`
		LeasedResource   NullableResourcepoolLeaseResourceRelationship `json:"LeasedResource,omitempty"`
		Pool             NullableResourcepoolPoolRelationship          `json:"Pool,omitempty"`
		PoolMember       NullableResourcepoolPoolMemberRelationship    `json:"PoolMember,omitempty"`
		Universe         NullableResourcepoolUniverseRelationship      `json:"Universe,omitempty"`
	}

	varResourcepoolLeaseWithoutEmbeddedStruct := ResourcepoolLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourcepoolLeaseWithoutEmbeddedStruct)
	if err == nil {
		varResourcepoolLease := _ResourcepoolLease{}
		varResourcepoolLease.ClassId = varResourcepoolLeaseWithoutEmbeddedStruct.ClassId
		varResourcepoolLease.ObjectType = varResourcepoolLeaseWithoutEmbeddedStruct.ObjectType
		varResourcepoolLease.Condition = varResourcepoolLeaseWithoutEmbeddedStruct.Condition
		varResourcepoolLease.Feature = varResourcepoolLeaseWithoutEmbeddedStruct.Feature
		varResourcepoolLease.IsExclusiveAtAssignedEntity = varResourcepoolLeaseWithoutEmbeddedStruct.IsExclusiveAtAssignedEntity
		varResourcepoolLease.LeaseParameters = varResourcepoolLeaseWithoutEmbeddedStruct.LeaseParameters
		varResourcepoolLease.Resource = varResourcepoolLeaseWithoutEmbeddedStruct.Resource
		varResourcepoolLease.ResourceType = varResourcepoolLeaseWithoutEmbeddedStruct.ResourceType
		varResourcepoolLease.AssignedToEntity = varResourcepoolLeaseWithoutEmbeddedStruct.AssignedToEntity
		varResourcepoolLease.LeasedResource = varResourcepoolLeaseWithoutEmbeddedStruct.LeasedResource
		varResourcepoolLease.Pool = varResourcepoolLeaseWithoutEmbeddedStruct.Pool
		varResourcepoolLease.PoolMember = varResourcepoolLeaseWithoutEmbeddedStruct.PoolMember
		varResourcepoolLease.Universe = varResourcepoolLeaseWithoutEmbeddedStruct.Universe
		*o = ResourcepoolLease(varResourcepoolLease)
	} else {
		return err
	}

	varResourcepoolLease := _ResourcepoolLease{}

	err = json.Unmarshal(data, &varResourcepoolLease)
	if err == nil {
		o.PoolAbstractLease = varResourcepoolLease.PoolAbstractLease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		delete(additionalProperties, "Feature")
		delete(additionalProperties, "IsExclusiveAtAssignedEntity")
		delete(additionalProperties, "LeaseParameters")
		delete(additionalProperties, "Resource")
		delete(additionalProperties, "ResourceType")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "LeasedResource")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolAbstractLease := reflect.ValueOf(o.PoolAbstractLease)
		for i := 0; i < reflectPoolAbstractLease.Type().NumField(); i++ {
			t := reflectPoolAbstractLease.Type().Field(i)

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

type NullableResourcepoolLease struct {
	value *ResourcepoolLease
	isSet bool
}

func (v NullableResourcepoolLease) Get() *ResourcepoolLease {
	return v.value
}

func (v *NullableResourcepoolLease) Set(val *ResourcepoolLease) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolLease) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolLease(val *ResourcepoolLease) *NullableResourcepoolLease {
	return &NullableResourcepoolLease{value: val, isSet: true}
}

func (v NullableResourcepoolLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
