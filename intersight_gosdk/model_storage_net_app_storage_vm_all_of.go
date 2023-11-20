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

// StorageNetAppStorageVmAllOf Definition of the list of properties defined in 'storage.NetAppStorageVm', excluding properties defined in parent classes.
type StorageNetAppStorageVmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                  `json:"ObjectType"`
	Aggregates            []string                                `json:"Aggregates,omitempty"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	// Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers.
	CifsEnabled *bool    `json:"CifsEnabled,omitempty"`
	DnsDomains  []string `json:"DnsDomains,omitempty"`
	// Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers.
	FcpEnabled *bool `json:"FcpEnabled,omitempty"`
	// IPspace name. IPspaces are distinct IP address spaces in which storage virtual machines (SVMs) reside.
	Ipspace *string `json:"Ipspace,omitempty"`
	// Specifies whether the Storage VM is a SnapMirror source Storage VM, using SnapMirror to protect its data.
	IsProtected *string `json:"IsProtected,omitempty"`
	// Status for iSCSI protocol allowed to run on Vservers.
	IscsiEnabled *bool `json:"IscsiEnabled,omitempty"`
	// Unique identifier of VServer across data center.
	Key         *string  `json:"Key,omitempty"`
	NameServers []string `json:"NameServers,omitempty"`
	// The number of native FPolicy engines enabled on this SVM.
	NativeFpolicyCount *int64 `json:"NativeFpolicyCount,omitempty"`
	// Status for Network File System Protocol ( NFS ) allowed to run on  Vservers.
	NfsEnabled *bool `json:"NfsEnabled,omitempty"`
	// Status for NVME protocol allowed to run on Vservers.
	NvmeEnabled *bool `json:"NvmeEnabled,omitempty"`
	// SVM subtype (default, dp_destination, sync_source, or sync_destination). The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster.
	Subtype *string                           `json:"Subtype,omitempty"`
	Array   *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	// An array of relationships to storageNetAppSvmEvent resources.
	Events               []StorageNetAppSvmEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppStorageVmAllOf StorageNetAppStorageVmAllOf

// NewStorageNetAppStorageVmAllOf instantiates a new StorageNetAppStorageVmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppStorageVmAllOf(classId string, objectType string) *StorageNetAppStorageVmAllOf {
	this := StorageNetAppStorageVmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppStorageVmAllOfWithDefaults instantiates a new StorageNetAppStorageVmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppStorageVmAllOfWithDefaults() *StorageNetAppStorageVmAllOf {
	this := StorageNetAppStorageVmAllOf{}
	var classId string = "storage.NetAppStorageVm"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageVm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppStorageVmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppStorageVmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppStorageVmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppStorageVmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregates returns the Aggregates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVmAllOf) GetAggregates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVmAllOf) GetAggregatesOk() ([]string, bool) {
	if o == nil || o.Aggregates == nil {
		return nil, false
	}
	return o.Aggregates, true
}

// HasAggregates returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasAggregates() bool {
	if o != nil && o.Aggregates != nil {
		return true
	}

	return false
}

// SetAggregates gets a reference to the given []string and assigns it to the Aggregates field.
func (o *StorageNetAppStorageVmAllOf) SetAggregates(v []string) {
	o.Aggregates = v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppStorageVmAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetCifsEnabled returns the CifsEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetCifsEnabled() bool {
	if o == nil || o.CifsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CifsEnabled
}

// GetCifsEnabledOk returns a tuple with the CifsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetCifsEnabledOk() (*bool, bool) {
	if o == nil || o.CifsEnabled == nil {
		return nil, false
	}
	return o.CifsEnabled, true
}

// HasCifsEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasCifsEnabled() bool {
	if o != nil && o.CifsEnabled != nil {
		return true
	}

	return false
}

// SetCifsEnabled gets a reference to the given bool and assigns it to the CifsEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetCifsEnabled(v bool) {
	o.CifsEnabled = &v
}

// GetDnsDomains returns the DnsDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVmAllOf) GetDnsDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsDomains
}

// GetDnsDomainsOk returns a tuple with the DnsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVmAllOf) GetDnsDomainsOk() ([]string, bool) {
	if o == nil || o.DnsDomains == nil {
		return nil, false
	}
	return o.DnsDomains, true
}

// HasDnsDomains returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasDnsDomains() bool {
	if o != nil && o.DnsDomains != nil {
		return true
	}

	return false
}

// SetDnsDomains gets a reference to the given []string and assigns it to the DnsDomains field.
func (o *StorageNetAppStorageVmAllOf) SetDnsDomains(v []string) {
	o.DnsDomains = v
}

// GetFcpEnabled returns the FcpEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetFcpEnabled() bool {
	if o == nil || o.FcpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FcpEnabled
}

// GetFcpEnabledOk returns a tuple with the FcpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetFcpEnabledOk() (*bool, bool) {
	if o == nil || o.FcpEnabled == nil {
		return nil, false
	}
	return o.FcpEnabled, true
}

// HasFcpEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasFcpEnabled() bool {
	if o != nil && o.FcpEnabled != nil {
		return true
	}

	return false
}

// SetFcpEnabled gets a reference to the given bool and assigns it to the FcpEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetFcpEnabled(v bool) {
	o.FcpEnabled = &v
}

// GetIpspace returns the Ipspace field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetIpspace() string {
	if o == nil || o.Ipspace == nil {
		var ret string
		return ret
	}
	return *o.Ipspace
}

// GetIpspaceOk returns a tuple with the Ipspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetIpspaceOk() (*string, bool) {
	if o == nil || o.Ipspace == nil {
		return nil, false
	}
	return o.Ipspace, true
}

// HasIpspace returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasIpspace() bool {
	if o != nil && o.Ipspace != nil {
		return true
	}

	return false
}

// SetIpspace gets a reference to the given string and assigns it to the Ipspace field.
func (o *StorageNetAppStorageVmAllOf) SetIpspace(v string) {
	o.Ipspace = &v
}

// GetIsProtected returns the IsProtected field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetIsProtected() string {
	if o == nil || o.IsProtected == nil {
		var ret string
		return ret
	}
	return *o.IsProtected
}

// GetIsProtectedOk returns a tuple with the IsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetIsProtectedOk() (*string, bool) {
	if o == nil || o.IsProtected == nil {
		return nil, false
	}
	return o.IsProtected, true
}

// HasIsProtected returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasIsProtected() bool {
	if o != nil && o.IsProtected != nil {
		return true
	}

	return false
}

// SetIsProtected gets a reference to the given string and assigns it to the IsProtected field.
func (o *StorageNetAppStorageVmAllOf) SetIsProtected(v string) {
	o.IsProtected = &v
}

// GetIscsiEnabled returns the IscsiEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabled() bool {
	if o == nil || o.IscsiEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IscsiEnabled
}

// GetIscsiEnabledOk returns a tuple with the IscsiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabledOk() (*bool, bool) {
	if o == nil || o.IscsiEnabled == nil {
		return nil, false
	}
	return o.IscsiEnabled, true
}

// HasIscsiEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasIscsiEnabled() bool {
	if o != nil && o.IscsiEnabled != nil {
		return true
	}

	return false
}

// SetIscsiEnabled gets a reference to the given bool and assigns it to the IscsiEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetIscsiEnabled(v bool) {
	o.IscsiEnabled = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppStorageVmAllOf) SetKey(v string) {
	o.Key = &v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVmAllOf) GetNameServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVmAllOf) GetNameServersOk() ([]string, bool) {
	if o == nil || o.NameServers == nil {
		return nil, false
	}
	return o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasNameServers() bool {
	if o != nil && o.NameServers != nil {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *StorageNetAppStorageVmAllOf) SetNameServers(v []string) {
	o.NameServers = v
}

// GetNativeFpolicyCount returns the NativeFpolicyCount field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetNativeFpolicyCount() int64 {
	if o == nil || o.NativeFpolicyCount == nil {
		var ret int64
		return ret
	}
	return *o.NativeFpolicyCount
}

// GetNativeFpolicyCountOk returns a tuple with the NativeFpolicyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetNativeFpolicyCountOk() (*int64, bool) {
	if o == nil || o.NativeFpolicyCount == nil {
		return nil, false
	}
	return o.NativeFpolicyCount, true
}

// HasNativeFpolicyCount returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasNativeFpolicyCount() bool {
	if o != nil && o.NativeFpolicyCount != nil {
		return true
	}

	return false
}

// SetNativeFpolicyCount gets a reference to the given int64 and assigns it to the NativeFpolicyCount field.
func (o *StorageNetAppStorageVmAllOf) SetNativeFpolicyCount(v int64) {
	o.NativeFpolicyCount = &v
}

// GetNfsEnabled returns the NfsEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetNfsEnabled() bool {
	if o == nil || o.NfsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsEnabled
}

// GetNfsEnabledOk returns a tuple with the NfsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetNfsEnabledOk() (*bool, bool) {
	if o == nil || o.NfsEnabled == nil {
		return nil, false
	}
	return o.NfsEnabled, true
}

// HasNfsEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasNfsEnabled() bool {
	if o != nil && o.NfsEnabled != nil {
		return true
	}

	return false
}

// SetNfsEnabled gets a reference to the given bool and assigns it to the NfsEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetNfsEnabled(v bool) {
	o.NfsEnabled = &v
}

// GetNvmeEnabled returns the NvmeEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabled() bool {
	if o == nil || o.NvmeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NvmeEnabled
}

// GetNvmeEnabledOk returns a tuple with the NvmeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabledOk() (*bool, bool) {
	if o == nil || o.NvmeEnabled == nil {
		return nil, false
	}
	return o.NvmeEnabled, true
}

// HasNvmeEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasNvmeEnabled() bool {
	if o != nil && o.NvmeEnabled != nil {
		return true
	}

	return false
}

// SetNvmeEnabled gets a reference to the given bool and assigns it to the NvmeEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetNvmeEnabled(v bool) {
	o.NvmeEnabled = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *StorageNetAppStorageVmAllOf) SetSubtype(v string) {
	o.Subtype = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppStorageVmAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVmAllOf) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVmAllOf) GetDiskPoolOk() ([]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppStorageVmAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVmAllOf) GetEvents() []StorageNetAppSvmEventRelationship {
	if o == nil {
		var ret []StorageNetAppSvmEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVmAllOf) GetEventsOk() ([]StorageNetAppSvmEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppSvmEventRelationship and assigns it to the Events field.
func (o *StorageNetAppStorageVmAllOf) SetEvents(v []StorageNetAppSvmEventRelationship) {
	o.Events = v
}

func (o StorageNetAppStorageVmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Aggregates != nil {
		toSerialize["Aggregates"] = o.Aggregates
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.CifsEnabled != nil {
		toSerialize["CifsEnabled"] = o.CifsEnabled
	}
	if o.DnsDomains != nil {
		toSerialize["DnsDomains"] = o.DnsDomains
	}
	if o.FcpEnabled != nil {
		toSerialize["FcpEnabled"] = o.FcpEnabled
	}
	if o.Ipspace != nil {
		toSerialize["Ipspace"] = o.Ipspace
	}
	if o.IsProtected != nil {
		toSerialize["IsProtected"] = o.IsProtected
	}
	if o.IscsiEnabled != nil {
		toSerialize["IscsiEnabled"] = o.IscsiEnabled
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.NameServers != nil {
		toSerialize["NameServers"] = o.NameServers
	}
	if o.NativeFpolicyCount != nil {
		toSerialize["NativeFpolicyCount"] = o.NativeFpolicyCount
	}
	if o.NfsEnabled != nil {
		toSerialize["NfsEnabled"] = o.NfsEnabled
	}
	if o.NvmeEnabled != nil {
		toSerialize["NvmeEnabled"] = o.NvmeEnabled
	}
	if o.Subtype != nil {
		toSerialize["Subtype"] = o.Subtype
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppStorageVmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppStorageVmAllOf := _StorageNetAppStorageVmAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppStorageVmAllOf); err == nil {
		*o = StorageNetAppStorageVmAllOf(varStorageNetAppStorageVmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Aggregates")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "CifsEnabled")
		delete(additionalProperties, "DnsDomains")
		delete(additionalProperties, "FcpEnabled")
		delete(additionalProperties, "Ipspace")
		delete(additionalProperties, "IsProtected")
		delete(additionalProperties, "IscsiEnabled")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "NameServers")
		delete(additionalProperties, "NativeFpolicyCount")
		delete(additionalProperties, "NfsEnabled")
		delete(additionalProperties, "NvmeEnabled")
		delete(additionalProperties, "Subtype")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "DiskPool")
		delete(additionalProperties, "Events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppStorageVmAllOf struct {
	value *StorageNetAppStorageVmAllOf
	isSet bool
}

func (v NullableStorageNetAppStorageVmAllOf) Get() *StorageNetAppStorageVmAllOf {
	return v.value
}

func (v *NullableStorageNetAppStorageVmAllOf) Set(val *StorageNetAppStorageVmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppStorageVmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppStorageVmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppStorageVmAllOf(val *StorageNetAppStorageVmAllOf) *NullableStorageNetAppStorageVmAllOf {
	return &NullableStorageNetAppStorageVmAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppStorageVmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppStorageVmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
