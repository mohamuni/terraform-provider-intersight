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

// checks if the ApplianceAppOpStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceAppOpStatus{}

// ApplianceAppOpStatus Status of an application running in Intersight Appliance.
type ApplianceAppOpStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string               `json:"ObjectType"`
	ApiStatuses []ApplianceApiStatus `json:"ApiStatuses,omitempty"`
	// Unique label to identify the application.
	AppLabel *string `json:"AppLabel,omitempty"`
	// Operational status of the application. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
	// Number of replicas ready.  The number of instances of the application currently ready to perform its intended functions.
	ReadyCount *int64 `json:"ReadyCount,omitempty"`
	// Number of replicas provisioned. The number of instances of the application provisioned to run on the Intersight appliance.
	ReplicaCount *int64 `json:"ReplicaCount,omitempty"`
	// Number of instance restarts in the last hour.
	RestartCount1Hour *int64 `json:"RestartCount1Hour,omitempty"`
	// Number of instance restarts in the last 24 hours.
	RestartCount24Hours *int64 `json:"RestartCount24Hours,omitempty"`
	// Number of instance restarts in the last 5 minutes.
	RestartCount5Mins *int64 `json:"RestartCount5Mins,omitempty"`
	// Total number of restarts since last deployment.
	RestartCountTotal *int64 `json:"RestartCountTotal,omitempty"`
	// Number of replicas running. The number of instances of the application currently running.
	RunningCount         *int64                                      `json:"RunningCount,omitempty"`
	GroupOpStatus        NullableApplianceGroupOpStatusRelationship  `json:"GroupOpStatus,omitempty"`
	SystemOpStatus       NullableApplianceSystemOpStatusRelationship `json:"SystemOpStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceAppOpStatus ApplianceAppOpStatus

// NewApplianceAppOpStatus instantiates a new ApplianceAppOpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAppOpStatus(classId string, objectType string) *ApplianceAppOpStatus {
	this := ApplianceAppOpStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceAppOpStatusWithDefaults instantiates a new ApplianceAppOpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAppOpStatusWithDefaults() *ApplianceAppOpStatus {
	this := ApplianceAppOpStatus{}
	var classId string = "appliance.AppOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.AppOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceAppOpStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceAppOpStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.AppOpStatus" of the ClassId field.
func (o *ApplianceAppOpStatus) GetDefaultClassId() interface{} {
	return "appliance.AppOpStatus"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceAppOpStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceAppOpStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.AppOpStatus" of the ObjectType field.
func (o *ApplianceAppOpStatus) GetDefaultObjectType() interface{} {
	return "appliance.AppOpStatus"
}

// GetApiStatuses returns the ApiStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceAppOpStatus) GetApiStatuses() []ApplianceApiStatus {
	if o == nil {
		var ret []ApplianceApiStatus
		return ret
	}
	return o.ApiStatuses
}

// GetApiStatusesOk returns a tuple with the ApiStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceAppOpStatus) GetApiStatusesOk() ([]ApplianceApiStatus, bool) {
	if o == nil || IsNil(o.ApiStatuses) {
		return nil, false
	}
	return o.ApiStatuses, true
}

// HasApiStatuses returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasApiStatuses() bool {
	if o != nil && !IsNil(o.ApiStatuses) {
		return true
	}

	return false
}

// SetApiStatuses gets a reference to the given []ApplianceApiStatus and assigns it to the ApiStatuses field.
func (o *ApplianceAppOpStatus) SetApiStatuses(v []ApplianceApiStatus) {
	o.ApiStatuses = v
}

// GetAppLabel returns the AppLabel field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetAppLabel() string {
	if o == nil || IsNil(o.AppLabel) {
		var ret string
		return ret
	}
	return *o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetAppLabelOk() (*string, bool) {
	if o == nil || IsNil(o.AppLabel) {
		return nil, false
	}
	return o.AppLabel, true
}

// HasAppLabel returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasAppLabel() bool {
	if o != nil && !IsNil(o.AppLabel) {
		return true
	}

	return false
}

// SetAppLabel gets a reference to the given string and assigns it to the AppLabel field.
func (o *ApplianceAppOpStatus) SetAppLabel(v string) {
	o.AppLabel = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetOperationalStatus() string {
	if o == nil || IsNil(o.OperationalStatus) {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetOperationalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OperationalStatus) {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasOperationalStatus() bool {
	if o != nil && !IsNil(o.OperationalStatus) {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceAppOpStatus) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetReadyCount returns the ReadyCount field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetReadyCount() int64 {
	if o == nil || IsNil(o.ReadyCount) {
		var ret int64
		return ret
	}
	return *o.ReadyCount
}

// GetReadyCountOk returns a tuple with the ReadyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetReadyCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadyCount) {
		return nil, false
	}
	return o.ReadyCount, true
}

// HasReadyCount returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasReadyCount() bool {
	if o != nil && !IsNil(o.ReadyCount) {
		return true
	}

	return false
}

// SetReadyCount gets a reference to the given int64 and assigns it to the ReadyCount field.
func (o *ApplianceAppOpStatus) SetReadyCount(v int64) {
	o.ReadyCount = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetReplicaCount() int64 {
	if o == nil || IsNil(o.ReplicaCount) {
		var ret int64
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetReplicaCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplicaCount) {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasReplicaCount() bool {
	if o != nil && !IsNil(o.ReplicaCount) {
		return true
	}

	return false
}

// SetReplicaCount gets a reference to the given int64 and assigns it to the ReplicaCount field.
func (o *ApplianceAppOpStatus) SetReplicaCount(v int64) {
	o.ReplicaCount = &v
}

// GetRestartCount1Hour returns the RestartCount1Hour field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetRestartCount1Hour() int64 {
	if o == nil || IsNil(o.RestartCount1Hour) {
		var ret int64
		return ret
	}
	return *o.RestartCount1Hour
}

// GetRestartCount1HourOk returns a tuple with the RestartCount1Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetRestartCount1HourOk() (*int64, bool) {
	if o == nil || IsNil(o.RestartCount1Hour) {
		return nil, false
	}
	return o.RestartCount1Hour, true
}

// HasRestartCount1Hour returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasRestartCount1Hour() bool {
	if o != nil && !IsNil(o.RestartCount1Hour) {
		return true
	}

	return false
}

// SetRestartCount1Hour gets a reference to the given int64 and assigns it to the RestartCount1Hour field.
func (o *ApplianceAppOpStatus) SetRestartCount1Hour(v int64) {
	o.RestartCount1Hour = &v
}

// GetRestartCount24Hours returns the RestartCount24Hours field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetRestartCount24Hours() int64 {
	if o == nil || IsNil(o.RestartCount24Hours) {
		var ret int64
		return ret
	}
	return *o.RestartCount24Hours
}

// GetRestartCount24HoursOk returns a tuple with the RestartCount24Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetRestartCount24HoursOk() (*int64, bool) {
	if o == nil || IsNil(o.RestartCount24Hours) {
		return nil, false
	}
	return o.RestartCount24Hours, true
}

// HasRestartCount24Hours returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasRestartCount24Hours() bool {
	if o != nil && !IsNil(o.RestartCount24Hours) {
		return true
	}

	return false
}

// SetRestartCount24Hours gets a reference to the given int64 and assigns it to the RestartCount24Hours field.
func (o *ApplianceAppOpStatus) SetRestartCount24Hours(v int64) {
	o.RestartCount24Hours = &v
}

// GetRestartCount5Mins returns the RestartCount5Mins field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetRestartCount5Mins() int64 {
	if o == nil || IsNil(o.RestartCount5Mins) {
		var ret int64
		return ret
	}
	return *o.RestartCount5Mins
}

// GetRestartCount5MinsOk returns a tuple with the RestartCount5Mins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetRestartCount5MinsOk() (*int64, bool) {
	if o == nil || IsNil(o.RestartCount5Mins) {
		return nil, false
	}
	return o.RestartCount5Mins, true
}

// HasRestartCount5Mins returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasRestartCount5Mins() bool {
	if o != nil && !IsNil(o.RestartCount5Mins) {
		return true
	}

	return false
}

// SetRestartCount5Mins gets a reference to the given int64 and assigns it to the RestartCount5Mins field.
func (o *ApplianceAppOpStatus) SetRestartCount5Mins(v int64) {
	o.RestartCount5Mins = &v
}

// GetRestartCountTotal returns the RestartCountTotal field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetRestartCountTotal() int64 {
	if o == nil || IsNil(o.RestartCountTotal) {
		var ret int64
		return ret
	}
	return *o.RestartCountTotal
}

// GetRestartCountTotalOk returns a tuple with the RestartCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetRestartCountTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.RestartCountTotal) {
		return nil, false
	}
	return o.RestartCountTotal, true
}

// HasRestartCountTotal returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasRestartCountTotal() bool {
	if o != nil && !IsNil(o.RestartCountTotal) {
		return true
	}

	return false
}

// SetRestartCountTotal gets a reference to the given int64 and assigns it to the RestartCountTotal field.
func (o *ApplianceAppOpStatus) SetRestartCountTotal(v int64) {
	o.RestartCountTotal = &v
}

// GetRunningCount returns the RunningCount field value if set, zero value otherwise.
func (o *ApplianceAppOpStatus) GetRunningCount() int64 {
	if o == nil || IsNil(o.RunningCount) {
		var ret int64
		return ret
	}
	return *o.RunningCount
}

// GetRunningCountOk returns a tuple with the RunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatus) GetRunningCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RunningCount) {
		return nil, false
	}
	return o.RunningCount, true
}

// HasRunningCount returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasRunningCount() bool {
	if o != nil && !IsNil(o.RunningCount) {
		return true
	}

	return false
}

// SetRunningCount gets a reference to the given int64 and assigns it to the RunningCount field.
func (o *ApplianceAppOpStatus) SetRunningCount(v int64) {
	o.RunningCount = &v
}

// GetGroupOpStatus returns the GroupOpStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceAppOpStatus) GetGroupOpStatus() ApplianceGroupOpStatusRelationship {
	if o == nil || IsNil(o.GroupOpStatus.Get()) {
		var ret ApplianceGroupOpStatusRelationship
		return ret
	}
	return *o.GroupOpStatus.Get()
}

// GetGroupOpStatusOk returns a tuple with the GroupOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceAppOpStatus) GetGroupOpStatusOk() (*ApplianceGroupOpStatusRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupOpStatus.Get(), o.GroupOpStatus.IsSet()
}

// HasGroupOpStatus returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasGroupOpStatus() bool {
	if o != nil && o.GroupOpStatus.IsSet() {
		return true
	}

	return false
}

// SetGroupOpStatus gets a reference to the given NullableApplianceGroupOpStatusRelationship and assigns it to the GroupOpStatus field.
func (o *ApplianceAppOpStatus) SetGroupOpStatus(v ApplianceGroupOpStatusRelationship) {
	o.GroupOpStatus.Set(&v)
}

// SetGroupOpStatusNil sets the value for GroupOpStatus to be an explicit nil
func (o *ApplianceAppOpStatus) SetGroupOpStatusNil() {
	o.GroupOpStatus.Set(nil)
}

// UnsetGroupOpStatus ensures that no value is present for GroupOpStatus, not even an explicit nil
func (o *ApplianceAppOpStatus) UnsetGroupOpStatus() {
	o.GroupOpStatus.Unset()
}

// GetSystemOpStatus returns the SystemOpStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceAppOpStatus) GetSystemOpStatus() ApplianceSystemOpStatusRelationship {
	if o == nil || IsNil(o.SystemOpStatus.Get()) {
		var ret ApplianceSystemOpStatusRelationship
		return ret
	}
	return *o.SystemOpStatus.Get()
}

// GetSystemOpStatusOk returns a tuple with the SystemOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceAppOpStatus) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemOpStatus.Get(), o.SystemOpStatus.IsSet()
}

// HasSystemOpStatus returns a boolean if a field has been set.
func (o *ApplianceAppOpStatus) HasSystemOpStatus() bool {
	if o != nil && o.SystemOpStatus.IsSet() {
		return true
	}

	return false
}

// SetSystemOpStatus gets a reference to the given NullableApplianceSystemOpStatusRelationship and assigns it to the SystemOpStatus field.
func (o *ApplianceAppOpStatus) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship) {
	o.SystemOpStatus.Set(&v)
}

// SetSystemOpStatusNil sets the value for SystemOpStatus to be an explicit nil
func (o *ApplianceAppOpStatus) SetSystemOpStatusNil() {
	o.SystemOpStatus.Set(nil)
}

// UnsetSystemOpStatus ensures that no value is present for SystemOpStatus, not even an explicit nil
func (o *ApplianceAppOpStatus) UnsetSystemOpStatus() {
	o.SystemOpStatus.Unset()
}

func (o ApplianceAppOpStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceAppOpStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ApiStatuses != nil {
		toSerialize["ApiStatuses"] = o.ApiStatuses
	}
	if !IsNil(o.AppLabel) {
		toSerialize["AppLabel"] = o.AppLabel
	}
	if !IsNil(o.OperationalStatus) {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if !IsNil(o.ReadyCount) {
		toSerialize["ReadyCount"] = o.ReadyCount
	}
	if !IsNil(o.ReplicaCount) {
		toSerialize["ReplicaCount"] = o.ReplicaCount
	}
	if !IsNil(o.RestartCount1Hour) {
		toSerialize["RestartCount1Hour"] = o.RestartCount1Hour
	}
	if !IsNil(o.RestartCount24Hours) {
		toSerialize["RestartCount24Hours"] = o.RestartCount24Hours
	}
	if !IsNil(o.RestartCount5Mins) {
		toSerialize["RestartCount5Mins"] = o.RestartCount5Mins
	}
	if !IsNil(o.RestartCountTotal) {
		toSerialize["RestartCountTotal"] = o.RestartCountTotal
	}
	if !IsNil(o.RunningCount) {
		toSerialize["RunningCount"] = o.RunningCount
	}
	if o.GroupOpStatus.IsSet() {
		toSerialize["GroupOpStatus"] = o.GroupOpStatus.Get()
	}
	if o.SystemOpStatus.IsSet() {
		toSerialize["SystemOpStatus"] = o.SystemOpStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceAppOpStatus) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceAppOpStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string               `json:"ObjectType"`
		ApiStatuses []ApplianceApiStatus `json:"ApiStatuses,omitempty"`
		// Unique label to identify the application.
		AppLabel *string `json:"AppLabel,omitempty"`
		// Operational status of the application. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
		OperationalStatus *string `json:"OperationalStatus,omitempty"`
		// Number of replicas ready.  The number of instances of the application currently ready to perform its intended functions.
		ReadyCount *int64 `json:"ReadyCount,omitempty"`
		// Number of replicas provisioned. The number of instances of the application provisioned to run on the Intersight appliance.
		ReplicaCount *int64 `json:"ReplicaCount,omitempty"`
		// Number of instance restarts in the last hour.
		RestartCount1Hour *int64 `json:"RestartCount1Hour,omitempty"`
		// Number of instance restarts in the last 24 hours.
		RestartCount24Hours *int64 `json:"RestartCount24Hours,omitempty"`
		// Number of instance restarts in the last 5 minutes.
		RestartCount5Mins *int64 `json:"RestartCount5Mins,omitempty"`
		// Total number of restarts since last deployment.
		RestartCountTotal *int64 `json:"RestartCountTotal,omitempty"`
		// Number of replicas running. The number of instances of the application currently running.
		RunningCount   *int64                                      `json:"RunningCount,omitempty"`
		GroupOpStatus  NullableApplianceGroupOpStatusRelationship  `json:"GroupOpStatus,omitempty"`
		SystemOpStatus NullableApplianceSystemOpStatusRelationship `json:"SystemOpStatus,omitempty"`
	}

	varApplianceAppOpStatusWithoutEmbeddedStruct := ApplianceAppOpStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceAppOpStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceAppOpStatus := _ApplianceAppOpStatus{}
		varApplianceAppOpStatus.ClassId = varApplianceAppOpStatusWithoutEmbeddedStruct.ClassId
		varApplianceAppOpStatus.ObjectType = varApplianceAppOpStatusWithoutEmbeddedStruct.ObjectType
		varApplianceAppOpStatus.ApiStatuses = varApplianceAppOpStatusWithoutEmbeddedStruct.ApiStatuses
		varApplianceAppOpStatus.AppLabel = varApplianceAppOpStatusWithoutEmbeddedStruct.AppLabel
		varApplianceAppOpStatus.OperationalStatus = varApplianceAppOpStatusWithoutEmbeddedStruct.OperationalStatus
		varApplianceAppOpStatus.ReadyCount = varApplianceAppOpStatusWithoutEmbeddedStruct.ReadyCount
		varApplianceAppOpStatus.ReplicaCount = varApplianceAppOpStatusWithoutEmbeddedStruct.ReplicaCount
		varApplianceAppOpStatus.RestartCount1Hour = varApplianceAppOpStatusWithoutEmbeddedStruct.RestartCount1Hour
		varApplianceAppOpStatus.RestartCount24Hours = varApplianceAppOpStatusWithoutEmbeddedStruct.RestartCount24Hours
		varApplianceAppOpStatus.RestartCount5Mins = varApplianceAppOpStatusWithoutEmbeddedStruct.RestartCount5Mins
		varApplianceAppOpStatus.RestartCountTotal = varApplianceAppOpStatusWithoutEmbeddedStruct.RestartCountTotal
		varApplianceAppOpStatus.RunningCount = varApplianceAppOpStatusWithoutEmbeddedStruct.RunningCount
		varApplianceAppOpStatus.GroupOpStatus = varApplianceAppOpStatusWithoutEmbeddedStruct.GroupOpStatus
		varApplianceAppOpStatus.SystemOpStatus = varApplianceAppOpStatusWithoutEmbeddedStruct.SystemOpStatus
		*o = ApplianceAppOpStatus(varApplianceAppOpStatus)
	} else {
		return err
	}

	varApplianceAppOpStatus := _ApplianceAppOpStatus{}

	err = json.Unmarshal(data, &varApplianceAppOpStatus)
	if err == nil {
		o.MoBaseMo = varApplianceAppOpStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiStatuses")
		delete(additionalProperties, "AppLabel")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "ReadyCount")
		delete(additionalProperties, "ReplicaCount")
		delete(additionalProperties, "RestartCount1Hour")
		delete(additionalProperties, "RestartCount24Hours")
		delete(additionalProperties, "RestartCount5Mins")
		delete(additionalProperties, "RestartCountTotal")
		delete(additionalProperties, "RunningCount")
		delete(additionalProperties, "GroupOpStatus")
		delete(additionalProperties, "SystemOpStatus")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableApplianceAppOpStatus struct {
	value *ApplianceAppOpStatus
	isSet bool
}

func (v NullableApplianceAppOpStatus) Get() *ApplianceAppOpStatus {
	return v.value
}

func (v *NullableApplianceAppOpStatus) Set(val *ApplianceAppOpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAppOpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAppOpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAppOpStatus(val *ApplianceAppOpStatus) *NullableApplianceAppOpStatus {
	return &NullableApplianceAppOpStatus{value: val, isSet: true}
}

func (v NullableApplianceAppOpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAppOpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
