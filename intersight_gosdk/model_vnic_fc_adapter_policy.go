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

// checks if the VnicFcAdapterPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicFcAdapterPolicy{}

// VnicFcAdapterPolicy A Fibre Channel Adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. You can enable FCP Error Recovery, change the default settings of Queues and Interrupt handling for performance enhancement.
type VnicFcAdapterPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred.
	ErrorDetectionTimeout *int64                              `json:"ErrorDetectionTimeout,omitempty"`
	ErrorRecoverySettings NullableVnicFcErrorRecoverySettings `json:"ErrorRecoverySettings,omitempty"`
	FlogiSettings         NullableVnicFlogiSettings           `json:"FlogiSettings,omitempty"`
	InterruptSettings     NullableVnicFcInterruptSettings     `json:"InterruptSettings,omitempty"`
	// The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed.
	IoThrottleCount *int64 `json:"IoThrottleCount,omitempty"`
	// The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. Lun Count value can exceed 1024 only for vHBA of type 'FC Initiator' and on servers having supported firmware version.
	LunCount *int64 `json:"LunCount,omitempty"`
	// The number of commands that the HBA can send and receive in a single transmission per LUN.
	LunQueueDepth *int64                    `json:"LunQueueDepth,omitempty"`
	PlogiSettings NullableVnicPlogiSettings `json:"PlogiSettings,omitempty"`
	// Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated.
	ResourceAllocationTimeout *int64                                       `json:"ResourceAllocationTimeout,omitempty"`
	RxQueueSettings           NullableVnicFcQueueSettings                  `json:"RxQueueSettings,omitempty"`
	ScsiQueueSettings         NullableVnicScsiQueueSettings                `json:"ScsiQueueSettings,omitempty"`
	TxQueueSettings           NullableVnicFcQueueSettings                  `json:"TxQueueSettings,omitempty"`
	Organization              NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _VnicFcAdapterPolicy VnicFcAdapterPolicy

// NewVnicFcAdapterPolicy instantiates a new VnicFcAdapterPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcAdapterPolicy(classId string, objectType string) *VnicFcAdapterPolicy {
	this := VnicFcAdapterPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var errorDetectionTimeout int64 = 2000
	this.ErrorDetectionTimeout = &errorDetectionTimeout
	var ioThrottleCount int64 = 512
	this.IoThrottleCount = &ioThrottleCount
	var lunCount int64 = 1024
	this.LunCount = &lunCount
	var lunQueueDepth int64 = 20
	this.LunQueueDepth = &lunQueueDepth
	var resourceAllocationTimeout int64 = 10000
	this.ResourceAllocationTimeout = &resourceAllocationTimeout
	return &this
}

// NewVnicFcAdapterPolicyWithDefaults instantiates a new VnicFcAdapterPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcAdapterPolicyWithDefaults() *VnicFcAdapterPolicy {
	this := VnicFcAdapterPolicy{}
	var classId string = "vnic.FcAdapterPolicy"
	this.ClassId = classId
	var objectType string = "vnic.FcAdapterPolicy"
	this.ObjectType = objectType
	var errorDetectionTimeout int64 = 2000
	this.ErrorDetectionTimeout = &errorDetectionTimeout
	var ioThrottleCount int64 = 512
	this.IoThrottleCount = &ioThrottleCount
	var lunCount int64 = 1024
	this.LunCount = &lunCount
	var lunQueueDepth int64 = 20
	this.LunQueueDepth = &lunQueueDepth
	var resourceAllocationTimeout int64 = 10000
	this.ResourceAllocationTimeout = &resourceAllocationTimeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcAdapterPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcAdapterPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.FcAdapterPolicy" of the ClassId field.
func (o *VnicFcAdapterPolicy) GetDefaultClassId() interface{} {
	return "vnic.FcAdapterPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcAdapterPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcAdapterPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.FcAdapterPolicy" of the ObjectType field.
func (o *VnicFcAdapterPolicy) GetDefaultObjectType() interface{} {
	return "vnic.FcAdapterPolicy"
}

// GetErrorDetectionTimeout returns the ErrorDetectionTimeout field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicy) GetErrorDetectionTimeout() int64 {
	if o == nil || IsNil(o.ErrorDetectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ErrorDetectionTimeout
}

// GetErrorDetectionTimeoutOk returns a tuple with the ErrorDetectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetErrorDetectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ErrorDetectionTimeout) {
		return nil, false
	}
	return o.ErrorDetectionTimeout, true
}

// HasErrorDetectionTimeout returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasErrorDetectionTimeout() bool {
	if o != nil && !IsNil(o.ErrorDetectionTimeout) {
		return true
	}

	return false
}

// SetErrorDetectionTimeout gets a reference to the given int64 and assigns it to the ErrorDetectionTimeout field.
func (o *VnicFcAdapterPolicy) SetErrorDetectionTimeout(v int64) {
	o.ErrorDetectionTimeout = &v
}

// GetErrorRecoverySettings returns the ErrorRecoverySettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetErrorRecoverySettings() VnicFcErrorRecoverySettings {
	if o == nil || IsNil(o.ErrorRecoverySettings.Get()) {
		var ret VnicFcErrorRecoverySettings
		return ret
	}
	return *o.ErrorRecoverySettings.Get()
}

// GetErrorRecoverySettingsOk returns a tuple with the ErrorRecoverySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetErrorRecoverySettingsOk() (*VnicFcErrorRecoverySettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorRecoverySettings.Get(), o.ErrorRecoverySettings.IsSet()
}

// HasErrorRecoverySettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasErrorRecoverySettings() bool {
	if o != nil && o.ErrorRecoverySettings.IsSet() {
		return true
	}

	return false
}

// SetErrorRecoverySettings gets a reference to the given NullableVnicFcErrorRecoverySettings and assigns it to the ErrorRecoverySettings field.
func (o *VnicFcAdapterPolicy) SetErrorRecoverySettings(v VnicFcErrorRecoverySettings) {
	o.ErrorRecoverySettings.Set(&v)
}

// SetErrorRecoverySettingsNil sets the value for ErrorRecoverySettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetErrorRecoverySettingsNil() {
	o.ErrorRecoverySettings.Set(nil)
}

// UnsetErrorRecoverySettings ensures that no value is present for ErrorRecoverySettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetErrorRecoverySettings() {
	o.ErrorRecoverySettings.Unset()
}

// GetFlogiSettings returns the FlogiSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetFlogiSettings() VnicFlogiSettings {
	if o == nil || IsNil(o.FlogiSettings.Get()) {
		var ret VnicFlogiSettings
		return ret
	}
	return *o.FlogiSettings.Get()
}

// GetFlogiSettingsOk returns a tuple with the FlogiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetFlogiSettingsOk() (*VnicFlogiSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlogiSettings.Get(), o.FlogiSettings.IsSet()
}

// HasFlogiSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasFlogiSettings() bool {
	if o != nil && o.FlogiSettings.IsSet() {
		return true
	}

	return false
}

// SetFlogiSettings gets a reference to the given NullableVnicFlogiSettings and assigns it to the FlogiSettings field.
func (o *VnicFcAdapterPolicy) SetFlogiSettings(v VnicFlogiSettings) {
	o.FlogiSettings.Set(&v)
}

// SetFlogiSettingsNil sets the value for FlogiSettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetFlogiSettingsNil() {
	o.FlogiSettings.Set(nil)
}

// UnsetFlogiSettings ensures that no value is present for FlogiSettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetFlogiSettings() {
	o.FlogiSettings.Unset()
}

// GetInterruptSettings returns the InterruptSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetInterruptSettings() VnicFcInterruptSettings {
	if o == nil || IsNil(o.InterruptSettings.Get()) {
		var ret VnicFcInterruptSettings
		return ret
	}
	return *o.InterruptSettings.Get()
}

// GetInterruptSettingsOk returns a tuple with the InterruptSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetInterruptSettingsOk() (*VnicFcInterruptSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterruptSettings.Get(), o.InterruptSettings.IsSet()
}

// HasInterruptSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasInterruptSettings() bool {
	if o != nil && o.InterruptSettings.IsSet() {
		return true
	}

	return false
}

// SetInterruptSettings gets a reference to the given NullableVnicFcInterruptSettings and assigns it to the InterruptSettings field.
func (o *VnicFcAdapterPolicy) SetInterruptSettings(v VnicFcInterruptSettings) {
	o.InterruptSettings.Set(&v)
}

// SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetInterruptSettingsNil() {
	o.InterruptSettings.Set(nil)
}

// UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetInterruptSettings() {
	o.InterruptSettings.Unset()
}

// GetIoThrottleCount returns the IoThrottleCount field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicy) GetIoThrottleCount() int64 {
	if o == nil || IsNil(o.IoThrottleCount) {
		var ret int64
		return ret
	}
	return *o.IoThrottleCount
}

// GetIoThrottleCountOk returns a tuple with the IoThrottleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetIoThrottleCountOk() (*int64, bool) {
	if o == nil || IsNil(o.IoThrottleCount) {
		return nil, false
	}
	return o.IoThrottleCount, true
}

// HasIoThrottleCount returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasIoThrottleCount() bool {
	if o != nil && !IsNil(o.IoThrottleCount) {
		return true
	}

	return false
}

// SetIoThrottleCount gets a reference to the given int64 and assigns it to the IoThrottleCount field.
func (o *VnicFcAdapterPolicy) SetIoThrottleCount(v int64) {
	o.IoThrottleCount = &v
}

// GetLunCount returns the LunCount field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicy) GetLunCount() int64 {
	if o == nil || IsNil(o.LunCount) {
		var ret int64
		return ret
	}
	return *o.LunCount
}

// GetLunCountOk returns a tuple with the LunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetLunCountOk() (*int64, bool) {
	if o == nil || IsNil(o.LunCount) {
		return nil, false
	}
	return o.LunCount, true
}

// HasLunCount returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasLunCount() bool {
	if o != nil && !IsNil(o.LunCount) {
		return true
	}

	return false
}

// SetLunCount gets a reference to the given int64 and assigns it to the LunCount field.
func (o *VnicFcAdapterPolicy) SetLunCount(v int64) {
	o.LunCount = &v
}

// GetLunQueueDepth returns the LunQueueDepth field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicy) GetLunQueueDepth() int64 {
	if o == nil || IsNil(o.LunQueueDepth) {
		var ret int64
		return ret
	}
	return *o.LunQueueDepth
}

// GetLunQueueDepthOk returns a tuple with the LunQueueDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetLunQueueDepthOk() (*int64, bool) {
	if o == nil || IsNil(o.LunQueueDepth) {
		return nil, false
	}
	return o.LunQueueDepth, true
}

// HasLunQueueDepth returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasLunQueueDepth() bool {
	if o != nil && !IsNil(o.LunQueueDepth) {
		return true
	}

	return false
}

// SetLunQueueDepth gets a reference to the given int64 and assigns it to the LunQueueDepth field.
func (o *VnicFcAdapterPolicy) SetLunQueueDepth(v int64) {
	o.LunQueueDepth = &v
}

// GetPlogiSettings returns the PlogiSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetPlogiSettings() VnicPlogiSettings {
	if o == nil || IsNil(o.PlogiSettings.Get()) {
		var ret VnicPlogiSettings
		return ret
	}
	return *o.PlogiSettings.Get()
}

// GetPlogiSettingsOk returns a tuple with the PlogiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetPlogiSettingsOk() (*VnicPlogiSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlogiSettings.Get(), o.PlogiSettings.IsSet()
}

// HasPlogiSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasPlogiSettings() bool {
	if o != nil && o.PlogiSettings.IsSet() {
		return true
	}

	return false
}

// SetPlogiSettings gets a reference to the given NullableVnicPlogiSettings and assigns it to the PlogiSettings field.
func (o *VnicFcAdapterPolicy) SetPlogiSettings(v VnicPlogiSettings) {
	o.PlogiSettings.Set(&v)
}

// SetPlogiSettingsNil sets the value for PlogiSettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetPlogiSettingsNil() {
	o.PlogiSettings.Set(nil)
}

// UnsetPlogiSettings ensures that no value is present for PlogiSettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetPlogiSettings() {
	o.PlogiSettings.Unset()
}

// GetResourceAllocationTimeout returns the ResourceAllocationTimeout field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicy) GetResourceAllocationTimeout() int64 {
	if o == nil || IsNil(o.ResourceAllocationTimeout) {
		var ret int64
		return ret
	}
	return *o.ResourceAllocationTimeout
}

// GetResourceAllocationTimeoutOk returns a tuple with the ResourceAllocationTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicy) GetResourceAllocationTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceAllocationTimeout) {
		return nil, false
	}
	return o.ResourceAllocationTimeout, true
}

// HasResourceAllocationTimeout returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasResourceAllocationTimeout() bool {
	if o != nil && !IsNil(o.ResourceAllocationTimeout) {
		return true
	}

	return false
}

// SetResourceAllocationTimeout gets a reference to the given int64 and assigns it to the ResourceAllocationTimeout field.
func (o *VnicFcAdapterPolicy) SetResourceAllocationTimeout(v int64) {
	o.ResourceAllocationTimeout = &v
}

// GetRxQueueSettings returns the RxQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetRxQueueSettings() VnicFcQueueSettings {
	if o == nil || IsNil(o.RxQueueSettings.Get()) {
		var ret VnicFcQueueSettings
		return ret
	}
	return *o.RxQueueSettings.Get()
}

// GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetRxQueueSettingsOk() (*VnicFcQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.RxQueueSettings.Get(), o.RxQueueSettings.IsSet()
}

// HasRxQueueSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasRxQueueSettings() bool {
	if o != nil && o.RxQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetRxQueueSettings gets a reference to the given NullableVnicFcQueueSettings and assigns it to the RxQueueSettings field.
func (o *VnicFcAdapterPolicy) SetRxQueueSettings(v VnicFcQueueSettings) {
	o.RxQueueSettings.Set(&v)
}

// SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetRxQueueSettingsNil() {
	o.RxQueueSettings.Set(nil)
}

// UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetRxQueueSettings() {
	o.RxQueueSettings.Unset()
}

// GetScsiQueueSettings returns the ScsiQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetScsiQueueSettings() VnicScsiQueueSettings {
	if o == nil || IsNil(o.ScsiQueueSettings.Get()) {
		var ret VnicScsiQueueSettings
		return ret
	}
	return *o.ScsiQueueSettings.Get()
}

// GetScsiQueueSettingsOk returns a tuple with the ScsiQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetScsiQueueSettingsOk() (*VnicScsiQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScsiQueueSettings.Get(), o.ScsiQueueSettings.IsSet()
}

// HasScsiQueueSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasScsiQueueSettings() bool {
	if o != nil && o.ScsiQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetScsiQueueSettings gets a reference to the given NullableVnicScsiQueueSettings and assigns it to the ScsiQueueSettings field.
func (o *VnicFcAdapterPolicy) SetScsiQueueSettings(v VnicScsiQueueSettings) {
	o.ScsiQueueSettings.Set(&v)
}

// SetScsiQueueSettingsNil sets the value for ScsiQueueSettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetScsiQueueSettingsNil() {
	o.ScsiQueueSettings.Set(nil)
}

// UnsetScsiQueueSettings ensures that no value is present for ScsiQueueSettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetScsiQueueSettings() {
	o.ScsiQueueSettings.Unset()
}

// GetTxQueueSettings returns the TxQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetTxQueueSettings() VnicFcQueueSettings {
	if o == nil || IsNil(o.TxQueueSettings.Get()) {
		var ret VnicFcQueueSettings
		return ret
	}
	return *o.TxQueueSettings.Get()
}

// GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetTxQueueSettingsOk() (*VnicFcQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxQueueSettings.Get(), o.TxQueueSettings.IsSet()
}

// HasTxQueueSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasTxQueueSettings() bool {
	if o != nil && o.TxQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetTxQueueSettings gets a reference to the given NullableVnicFcQueueSettings and assigns it to the TxQueueSettings field.
func (o *VnicFcAdapterPolicy) SetTxQueueSettings(v VnicFcQueueSettings) {
	o.TxQueueSettings.Set(&v)
}

// SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil
func (o *VnicFcAdapterPolicy) SetTxQueueSettingsNil() {
	o.TxQueueSettings.Set(nil)
}

// UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetTxQueueSettings() {
	o.TxQueueSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicFcAdapterPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *VnicFcAdapterPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *VnicFcAdapterPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o VnicFcAdapterPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicFcAdapterPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ErrorDetectionTimeout) {
		toSerialize["ErrorDetectionTimeout"] = o.ErrorDetectionTimeout
	}
	if o.ErrorRecoverySettings.IsSet() {
		toSerialize["ErrorRecoverySettings"] = o.ErrorRecoverySettings.Get()
	}
	if o.FlogiSettings.IsSet() {
		toSerialize["FlogiSettings"] = o.FlogiSettings.Get()
	}
	if o.InterruptSettings.IsSet() {
		toSerialize["InterruptSettings"] = o.InterruptSettings.Get()
	}
	if !IsNil(o.IoThrottleCount) {
		toSerialize["IoThrottleCount"] = o.IoThrottleCount
	}
	if !IsNil(o.LunCount) {
		toSerialize["LunCount"] = o.LunCount
	}
	if !IsNil(o.LunQueueDepth) {
		toSerialize["LunQueueDepth"] = o.LunQueueDepth
	}
	if o.PlogiSettings.IsSet() {
		toSerialize["PlogiSettings"] = o.PlogiSettings.Get()
	}
	if !IsNil(o.ResourceAllocationTimeout) {
		toSerialize["ResourceAllocationTimeout"] = o.ResourceAllocationTimeout
	}
	if o.RxQueueSettings.IsSet() {
		toSerialize["RxQueueSettings"] = o.RxQueueSettings.Get()
	}
	if o.ScsiQueueSettings.IsSet() {
		toSerialize["ScsiQueueSettings"] = o.ScsiQueueSettings.Get()
	}
	if o.TxQueueSettings.IsSet() {
		toSerialize["TxQueueSettings"] = o.TxQueueSettings.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicFcAdapterPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type VnicFcAdapterPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred.
		ErrorDetectionTimeout *int64                              `json:"ErrorDetectionTimeout,omitempty"`
		ErrorRecoverySettings NullableVnicFcErrorRecoverySettings `json:"ErrorRecoverySettings,omitempty"`
		FlogiSettings         NullableVnicFlogiSettings           `json:"FlogiSettings,omitempty"`
		InterruptSettings     NullableVnicFcInterruptSettings     `json:"InterruptSettings,omitempty"`
		// The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed.
		IoThrottleCount *int64 `json:"IoThrottleCount,omitempty"`
		// The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. Lun Count value can exceed 1024 only for vHBA of type 'FC Initiator' and on servers having supported firmware version.
		LunCount *int64 `json:"LunCount,omitempty"`
		// The number of commands that the HBA can send and receive in a single transmission per LUN.
		LunQueueDepth *int64                    `json:"LunQueueDepth,omitempty"`
		PlogiSettings NullableVnicPlogiSettings `json:"PlogiSettings,omitempty"`
		// Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated.
		ResourceAllocationTimeout *int64                                       `json:"ResourceAllocationTimeout,omitempty"`
		RxQueueSettings           NullableVnicFcQueueSettings                  `json:"RxQueueSettings,omitempty"`
		ScsiQueueSettings         NullableVnicScsiQueueSettings                `json:"ScsiQueueSettings,omitempty"`
		TxQueueSettings           NullableVnicFcQueueSettings                  `json:"TxQueueSettings,omitempty"`
		Organization              NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicFcAdapterPolicyWithoutEmbeddedStruct := VnicFcAdapterPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicFcAdapterPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcAdapterPolicy := _VnicFcAdapterPolicy{}
		varVnicFcAdapterPolicy.ClassId = varVnicFcAdapterPolicyWithoutEmbeddedStruct.ClassId
		varVnicFcAdapterPolicy.ObjectType = varVnicFcAdapterPolicyWithoutEmbeddedStruct.ObjectType
		varVnicFcAdapterPolicy.ErrorDetectionTimeout = varVnicFcAdapterPolicyWithoutEmbeddedStruct.ErrorDetectionTimeout
		varVnicFcAdapterPolicy.ErrorRecoverySettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.ErrorRecoverySettings
		varVnicFcAdapterPolicy.FlogiSettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.FlogiSettings
		varVnicFcAdapterPolicy.InterruptSettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.InterruptSettings
		varVnicFcAdapterPolicy.IoThrottleCount = varVnicFcAdapterPolicyWithoutEmbeddedStruct.IoThrottleCount
		varVnicFcAdapterPolicy.LunCount = varVnicFcAdapterPolicyWithoutEmbeddedStruct.LunCount
		varVnicFcAdapterPolicy.LunQueueDepth = varVnicFcAdapterPolicyWithoutEmbeddedStruct.LunQueueDepth
		varVnicFcAdapterPolicy.PlogiSettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.PlogiSettings
		varVnicFcAdapterPolicy.ResourceAllocationTimeout = varVnicFcAdapterPolicyWithoutEmbeddedStruct.ResourceAllocationTimeout
		varVnicFcAdapterPolicy.RxQueueSettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.RxQueueSettings
		varVnicFcAdapterPolicy.ScsiQueueSettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.ScsiQueueSettings
		varVnicFcAdapterPolicy.TxQueueSettings = varVnicFcAdapterPolicyWithoutEmbeddedStruct.TxQueueSettings
		varVnicFcAdapterPolicy.Organization = varVnicFcAdapterPolicyWithoutEmbeddedStruct.Organization
		*o = VnicFcAdapterPolicy(varVnicFcAdapterPolicy)
	} else {
		return err
	}

	varVnicFcAdapterPolicy := _VnicFcAdapterPolicy{}

	err = json.Unmarshal(data, &varVnicFcAdapterPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicFcAdapterPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorDetectionTimeout")
		delete(additionalProperties, "ErrorRecoverySettings")
		delete(additionalProperties, "FlogiSettings")
		delete(additionalProperties, "InterruptSettings")
		delete(additionalProperties, "IoThrottleCount")
		delete(additionalProperties, "LunCount")
		delete(additionalProperties, "LunQueueDepth")
		delete(additionalProperties, "PlogiSettings")
		delete(additionalProperties, "ResourceAllocationTimeout")
		delete(additionalProperties, "RxQueueSettings")
		delete(additionalProperties, "ScsiQueueSettings")
		delete(additionalProperties, "TxQueueSettings")
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

type NullableVnicFcAdapterPolicy struct {
	value *VnicFcAdapterPolicy
	isSet bool
}

func (v NullableVnicFcAdapterPolicy) Get() *VnicFcAdapterPolicy {
	return v.value
}

func (v *NullableVnicFcAdapterPolicy) Set(val *VnicFcAdapterPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcAdapterPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcAdapterPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcAdapterPolicy(val *VnicFcAdapterPolicy) *NullableVnicFcAdapterPolicy {
	return &NullableVnicFcAdapterPolicy{value: val, isSet: true}
}

func (v NullableVnicFcAdapterPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcAdapterPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
