/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexClusterBackupPolicy Specifies cluster backup configuration for a HyperFlex Cluster.
type HyperflexClusterBackupPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up.
	BackupDataStoreName *string `json:"BackupDataStoreName,omitempty"`
	// Replication data store size in backupDataStoreSizeUnit.
	BackupDataStoreSize *int64 `json:"BackupDataStoreSize,omitempty"`
	// Replication data store size.
	BackupDataStoreSizeUnit *string `json:"BackupDataStoreSizeUnit,omitempty"`
	// Replication cluster pairing name prefix.
	ReplicationPairNamePrefix *string                              `json:"ReplicationPairNamePrefix,omitempty"`
	ReplicationSchedule       NullableHyperflexReplicationSchedule `json:"ReplicationSchedule,omitempty"`
	// Number of snapshots that will be retained as part of the Multi Point in Time support.
	SnapshotRetentionCount *int64                        `json:"SnapshotRetentionCount,omitempty"`
	BackupTarget           *HyperflexClusterRelationship `json:"BackupTarget,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterBackupPolicy HyperflexClusterBackupPolicy

// NewHyperflexClusterBackupPolicy instantiates a new HyperflexClusterBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterBackupPolicy(classId string, objectType string) *HyperflexClusterBackupPolicy {
	this := HyperflexClusterBackupPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var backupDataStoreName string = "backup-source-ds"
	this.BackupDataStoreName = &backupDataStoreName
	var backupDataStoreSize int64 = 2
	this.BackupDataStoreSize = &backupDataStoreSize
	var backupDataStoreSizeUnit string = "TB"
	this.BackupDataStoreSizeUnit = &backupDataStoreSizeUnit
	var replicationPairNamePrefix string = "backup"
	this.ReplicationPairNamePrefix = &replicationPairNamePrefix
	var snapshotRetentionCount int64 = 4
	this.SnapshotRetentionCount = &snapshotRetentionCount
	return &this
}

// NewHyperflexClusterBackupPolicyWithDefaults instantiates a new HyperflexClusterBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterBackupPolicyWithDefaults() *HyperflexClusterBackupPolicy {
	this := HyperflexClusterBackupPolicy{}
	var classId string = "hyperflex.ClusterBackupPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterBackupPolicy"
	this.ObjectType = objectType
	var backupDataStoreName string = "backup-source-ds"
	this.BackupDataStoreName = &backupDataStoreName
	var backupDataStoreSize int64 = 2
	this.BackupDataStoreSize = &backupDataStoreSize
	var backupDataStoreSizeUnit string = "TB"
	this.BackupDataStoreSizeUnit = &backupDataStoreSizeUnit
	var replicationPairNamePrefix string = "backup"
	this.ReplicationPairNamePrefix = &replicationPairNamePrefix
	var snapshotRetentionCount int64 = 4
	this.SnapshotRetentionCount = &snapshotRetentionCount
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterBackupPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterBackupPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterBackupPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterBackupPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupDataStoreName returns the BackupDataStoreName field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetBackupDataStoreName() string {
	if o == nil || o.BackupDataStoreName == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreName
}

// GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetBackupDataStoreNameOk() (*string, bool) {
	if o == nil || o.BackupDataStoreName == nil {
		return nil, false
	}
	return o.BackupDataStoreName, true
}

// HasBackupDataStoreName returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasBackupDataStoreName() bool {
	if o != nil && o.BackupDataStoreName != nil {
		return true
	}

	return false
}

// SetBackupDataStoreName gets a reference to the given string and assigns it to the BackupDataStoreName field.
func (o *HyperflexClusterBackupPolicy) SetBackupDataStoreName(v string) {
	o.BackupDataStoreName = &v
}

// GetBackupDataStoreSize returns the BackupDataStoreSize field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetBackupDataStoreSize() int64 {
	if o == nil || o.BackupDataStoreSize == nil {
		var ret int64
		return ret
	}
	return *o.BackupDataStoreSize
}

// GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetBackupDataStoreSizeOk() (*int64, bool) {
	if o == nil || o.BackupDataStoreSize == nil {
		return nil, false
	}
	return o.BackupDataStoreSize, true
}

// HasBackupDataStoreSize returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasBackupDataStoreSize() bool {
	if o != nil && o.BackupDataStoreSize != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSize gets a reference to the given int64 and assigns it to the BackupDataStoreSize field.
func (o *HyperflexClusterBackupPolicy) SetBackupDataStoreSize(v int64) {
	o.BackupDataStoreSize = &v
}

// GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetBackupDataStoreSizeUnit() string {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreSizeUnit
}

// GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetBackupDataStoreSizeUnitOk() (*string, bool) {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		return nil, false
	}
	return o.BackupDataStoreSizeUnit, true
}

// HasBackupDataStoreSizeUnit returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasBackupDataStoreSizeUnit() bool {
	if o != nil && o.BackupDataStoreSizeUnit != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSizeUnit gets a reference to the given string and assigns it to the BackupDataStoreSizeUnit field.
func (o *HyperflexClusterBackupPolicy) SetBackupDataStoreSizeUnit(v string) {
	o.BackupDataStoreSizeUnit = &v
}

// GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetReplicationPairNamePrefix() string {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.ReplicationPairNamePrefix
}

// GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetReplicationPairNamePrefixOk() (*string, bool) {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		return nil, false
	}
	return o.ReplicationPairNamePrefix, true
}

// HasReplicationPairNamePrefix returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasReplicationPairNamePrefix() bool {
	if o != nil && o.ReplicationPairNamePrefix != nil {
		return true
	}

	return false
}

// SetReplicationPairNamePrefix gets a reference to the given string and assigns it to the ReplicationPairNamePrefix field.
func (o *HyperflexClusterBackupPolicy) SetReplicationPairNamePrefix(v string) {
	o.ReplicationPairNamePrefix = &v
}

// GetReplicationSchedule returns the ReplicationSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterBackupPolicy) GetReplicationSchedule() HyperflexReplicationSchedule {
	if o == nil || o.ReplicationSchedule.Get() == nil {
		var ret HyperflexReplicationSchedule
		return ret
	}
	return *o.ReplicationSchedule.Get()
}

// GetReplicationScheduleOk returns a tuple with the ReplicationSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterBackupPolicy) GetReplicationScheduleOk() (*HyperflexReplicationSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationSchedule.Get(), o.ReplicationSchedule.IsSet()
}

// HasReplicationSchedule returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasReplicationSchedule() bool {
	if o != nil && o.ReplicationSchedule.IsSet() {
		return true
	}

	return false
}

// SetReplicationSchedule gets a reference to the given NullableHyperflexReplicationSchedule and assigns it to the ReplicationSchedule field.
func (o *HyperflexClusterBackupPolicy) SetReplicationSchedule(v HyperflexReplicationSchedule) {
	o.ReplicationSchedule.Set(&v)
}

// SetReplicationScheduleNil sets the value for ReplicationSchedule to be an explicit nil
func (o *HyperflexClusterBackupPolicy) SetReplicationScheduleNil() {
	o.ReplicationSchedule.Set(nil)
}

// UnsetReplicationSchedule ensures that no value is present for ReplicationSchedule, not even an explicit nil
func (o *HyperflexClusterBackupPolicy) UnsetReplicationSchedule() {
	o.ReplicationSchedule.Unset()
}

// GetSnapshotRetentionCount returns the SnapshotRetentionCount field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetSnapshotRetentionCount() int64 {
	if o == nil || o.SnapshotRetentionCount == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotRetentionCount
}

// GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetSnapshotRetentionCountOk() (*int64, bool) {
	if o == nil || o.SnapshotRetentionCount == nil {
		return nil, false
	}
	return o.SnapshotRetentionCount, true
}

// HasSnapshotRetentionCount returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasSnapshotRetentionCount() bool {
	if o != nil && o.SnapshotRetentionCount != nil {
		return true
	}

	return false
}

// SetSnapshotRetentionCount gets a reference to the given int64 and assigns it to the SnapshotRetentionCount field.
func (o *HyperflexClusterBackupPolicy) SetSnapshotRetentionCount(v int64) {
	o.SnapshotRetentionCount = &v
}

// GetBackupTarget returns the BackupTarget field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetBackupTarget() HyperflexClusterRelationship {
	if o == nil || o.BackupTarget == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.BackupTarget
}

// GetBackupTargetOk returns a tuple with the BackupTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetBackupTargetOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.BackupTarget == nil {
		return nil, false
	}
	return o.BackupTarget, true
}

// HasBackupTarget returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasBackupTarget() bool {
	if o != nil && o.BackupTarget != nil {
		return true
	}

	return false
}

// SetBackupTarget gets a reference to the given HyperflexClusterRelationship and assigns it to the BackupTarget field.
func (o *HyperflexClusterBackupPolicy) SetBackupTarget(v HyperflexClusterRelationship) {
	o.BackupTarget = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterBackupPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterBackupPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterBackupPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterBackupPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexClusterBackupPolicy) MarshalJSON() ([]byte, error) {
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
	if o.BackupDataStoreName != nil {
		toSerialize["BackupDataStoreName"] = o.BackupDataStoreName
	}
	if o.BackupDataStoreSize != nil {
		toSerialize["BackupDataStoreSize"] = o.BackupDataStoreSize
	}
	if o.BackupDataStoreSizeUnit != nil {
		toSerialize["BackupDataStoreSizeUnit"] = o.BackupDataStoreSizeUnit
	}
	if o.ReplicationPairNamePrefix != nil {
		toSerialize["ReplicationPairNamePrefix"] = o.ReplicationPairNamePrefix
	}
	if o.ReplicationSchedule.IsSet() {
		toSerialize["ReplicationSchedule"] = o.ReplicationSchedule.Get()
	}
	if o.SnapshotRetentionCount != nil {
		toSerialize["SnapshotRetentionCount"] = o.SnapshotRetentionCount
	}
	if o.BackupTarget != nil {
		toSerialize["BackupTarget"] = o.BackupTarget
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterBackupPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexClusterBackupPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up.
		BackupDataStoreName *string `json:"BackupDataStoreName,omitempty"`
		// Replication data store size in backupDataStoreSizeUnit.
		BackupDataStoreSize *int64 `json:"BackupDataStoreSize,omitempty"`
		// Replication data store size.
		BackupDataStoreSizeUnit *string `json:"BackupDataStoreSizeUnit,omitempty"`
		// Replication cluster pairing name prefix.
		ReplicationPairNamePrefix *string                              `json:"ReplicationPairNamePrefix,omitempty"`
		ReplicationSchedule       NullableHyperflexReplicationSchedule `json:"ReplicationSchedule,omitempty"`
		// Number of snapshots that will be retained as part of the Multi Point in Time support.
		SnapshotRetentionCount *int64                        `json:"SnapshotRetentionCount,omitempty"`
		BackupTarget           *HyperflexClusterRelationship `json:"BackupTarget,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexClusterBackupPolicyWithoutEmbeddedStruct := HyperflexClusterBackupPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexClusterBackupPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterBackupPolicy := _HyperflexClusterBackupPolicy{}
		varHyperflexClusterBackupPolicy.ClassId = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.ClassId
		varHyperflexClusterBackupPolicy.ObjectType = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterBackupPolicy.BackupDataStoreName = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.BackupDataStoreName
		varHyperflexClusterBackupPolicy.BackupDataStoreSize = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.BackupDataStoreSize
		varHyperflexClusterBackupPolicy.BackupDataStoreSizeUnit = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.BackupDataStoreSizeUnit
		varHyperflexClusterBackupPolicy.ReplicationPairNamePrefix = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.ReplicationPairNamePrefix
		varHyperflexClusterBackupPolicy.ReplicationSchedule = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.ReplicationSchedule
		varHyperflexClusterBackupPolicy.SnapshotRetentionCount = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.SnapshotRetentionCount
		varHyperflexClusterBackupPolicy.BackupTarget = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.BackupTarget
		varHyperflexClusterBackupPolicy.ClusterProfiles = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexClusterBackupPolicy.Organization = varHyperflexClusterBackupPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexClusterBackupPolicy(varHyperflexClusterBackupPolicy)
	} else {
		return err
	}

	varHyperflexClusterBackupPolicy := _HyperflexClusterBackupPolicy{}

	err = json.Unmarshal(bytes, &varHyperflexClusterBackupPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexClusterBackupPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupDataStoreName")
		delete(additionalProperties, "BackupDataStoreSize")
		delete(additionalProperties, "BackupDataStoreSizeUnit")
		delete(additionalProperties, "ReplicationPairNamePrefix")
		delete(additionalProperties, "ReplicationSchedule")
		delete(additionalProperties, "SnapshotRetentionCount")
		delete(additionalProperties, "BackupTarget")
		delete(additionalProperties, "ClusterProfiles")
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

type NullableHyperflexClusterBackupPolicy struct {
	value *HyperflexClusterBackupPolicy
	isSet bool
}

func (v NullableHyperflexClusterBackupPolicy) Get() *HyperflexClusterBackupPolicy {
	return v.value
}

func (v *NullableHyperflexClusterBackupPolicy) Set(val *HyperflexClusterBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterBackupPolicy(val *HyperflexClusterBackupPolicy) *NullableHyperflexClusterBackupPolicy {
	return &NullableHyperflexClusterBackupPolicy{value: val, isSet: true}
}

func (v NullableHyperflexClusterBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
