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
	"time"
)

// AssetDeploymentAllOf Definition of the list of properties defined in 'asset.Deployment', excluding properties defined in parent classes.
type AssetDeploymentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                           `json:"ObjectType"`
	AlarmInfo  NullableAssetDeploymentAlarmInfo `json:"AlarmInfo,omitempty"`
	// Identifies the consumption-based subscription's deployment.
	DeploymentRefId *string                          `json:"DeploymentRefId,omitempty"`
	EndCustomer     NullableAssetCustomerInformation `json:"EndCustomer,omitempty"`
	// End Date for the consumption-based subscription's deployment.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// Active license tier for the purchased Cisco device's deployment. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type. * `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud. * `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud. * `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud. * `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud. * `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud. * `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud. * `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers.
	LicenseType *string `json:"LicenseType,omitempty"`
	// Identifier for consumption based pricing. MLB refers to MultiLine Bundle.
	MlbOfferType *string `json:"MlbOfferType,omitempty"`
	// Start Date for the consumption-based subscription's deployment.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Identifies the consumption-based subscription.
	SubscriptionRefId *string             `json:"SubscriptionRefId,omitempty"`
	UnitOfMeasure     []AssetMeteringType `json:"UnitOfMeasure,omitempty"`
	Workloads         []string            `json:"Workloads,omitempty"`
	// An array of relationships to assetDeploymentDevice resources.
	Devices              []AssetDeploymentDeviceRelationship `json:"Devices,omitempty"`
	Subscription         *AssetSubscriptionRelationship      `json:"Subscription,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeploymentAllOf AssetDeploymentAllOf

// NewAssetDeploymentAllOf instantiates a new AssetDeploymentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeploymentAllOf(classId string, objectType string) *AssetDeploymentAllOf {
	this := AssetDeploymentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeploymentAllOfWithDefaults instantiates a new AssetDeploymentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeploymentAllOfWithDefaults() *AssetDeploymentAllOf {
	this := AssetDeploymentAllOf{}
	var classId string = "asset.Deployment"
	this.ClassId = classId
	var objectType string = "asset.Deployment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeploymentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeploymentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeploymentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeploymentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmInfo returns the AlarmInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentAllOf) GetAlarmInfo() AssetDeploymentAlarmInfo {
	if o == nil || o.AlarmInfo.Get() == nil {
		var ret AssetDeploymentAlarmInfo
		return ret
	}
	return *o.AlarmInfo.Get()
}

// GetAlarmInfoOk returns a tuple with the AlarmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentAllOf) GetAlarmInfoOk() (*AssetDeploymentAlarmInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmInfo.Get(), o.AlarmInfo.IsSet()
}

// HasAlarmInfo returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasAlarmInfo() bool {
	if o != nil && o.AlarmInfo.IsSet() {
		return true
	}

	return false
}

// SetAlarmInfo gets a reference to the given NullableAssetDeploymentAlarmInfo and assigns it to the AlarmInfo field.
func (o *AssetDeploymentAllOf) SetAlarmInfo(v AssetDeploymentAlarmInfo) {
	o.AlarmInfo.Set(&v)
}

// SetAlarmInfoNil sets the value for AlarmInfo to be an explicit nil
func (o *AssetDeploymentAllOf) SetAlarmInfoNil() {
	o.AlarmInfo.Set(nil)
}

// UnsetAlarmInfo ensures that no value is present for AlarmInfo, not even an explicit nil
func (o *AssetDeploymentAllOf) UnsetAlarmInfo() {
	o.AlarmInfo.Unset()
}

// GetDeploymentRefId returns the DeploymentRefId field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetDeploymentRefId() string {
	if o == nil || o.DeploymentRefId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentRefId
}

// GetDeploymentRefIdOk returns a tuple with the DeploymentRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetDeploymentRefIdOk() (*string, bool) {
	if o == nil || o.DeploymentRefId == nil {
		return nil, false
	}
	return o.DeploymentRefId, true
}

// HasDeploymentRefId returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasDeploymentRefId() bool {
	if o != nil && o.DeploymentRefId != nil {
		return true
	}

	return false
}

// SetDeploymentRefId gets a reference to the given string and assigns it to the DeploymentRefId field.
func (o *AssetDeploymentAllOf) SetDeploymentRefId(v string) {
	o.DeploymentRefId = &v
}

// GetEndCustomer returns the EndCustomer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentAllOf) GetEndCustomer() AssetCustomerInformation {
	if o == nil || o.EndCustomer.Get() == nil {
		var ret AssetCustomerInformation
		return ret
	}
	return *o.EndCustomer.Get()
}

// GetEndCustomerOk returns a tuple with the EndCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentAllOf) GetEndCustomerOk() (*AssetCustomerInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndCustomer.Get(), o.EndCustomer.IsSet()
}

// HasEndCustomer returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasEndCustomer() bool {
	if o != nil && o.EndCustomer.IsSet() {
		return true
	}

	return false
}

// SetEndCustomer gets a reference to the given NullableAssetCustomerInformation and assigns it to the EndCustomer field.
func (o *AssetDeploymentAllOf) SetEndCustomer(v AssetCustomerInformation) {
	o.EndCustomer.Set(&v)
}

// SetEndCustomerNil sets the value for EndCustomer to be an explicit nil
func (o *AssetDeploymentAllOf) SetEndCustomerNil() {
	o.EndCustomer.Set(nil)
}

// UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
func (o *AssetDeploymentAllOf) UnsetEndCustomer() {
	o.EndCustomer.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *AssetDeploymentAllOf) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *AssetDeploymentAllOf) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetMlbOfferType returns the MlbOfferType field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetMlbOfferType() string {
	if o == nil || o.MlbOfferType == nil {
		var ret string
		return ret
	}
	return *o.MlbOfferType
}

// GetMlbOfferTypeOk returns a tuple with the MlbOfferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetMlbOfferTypeOk() (*string, bool) {
	if o == nil || o.MlbOfferType == nil {
		return nil, false
	}
	return o.MlbOfferType, true
}

// HasMlbOfferType returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasMlbOfferType() bool {
	if o != nil && o.MlbOfferType != nil {
		return true
	}

	return false
}

// SetMlbOfferType gets a reference to the given string and assigns it to the MlbOfferType field.
func (o *AssetDeploymentAllOf) SetMlbOfferType(v string) {
	o.MlbOfferType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *AssetDeploymentAllOf) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetSubscriptionRefId returns the SubscriptionRefId field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetSubscriptionRefId() string {
	if o == nil || o.SubscriptionRefId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionRefId
}

// GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetSubscriptionRefIdOk() (*string, bool) {
	if o == nil || o.SubscriptionRefId == nil {
		return nil, false
	}
	return o.SubscriptionRefId, true
}

// HasSubscriptionRefId returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasSubscriptionRefId() bool {
	if o != nil && o.SubscriptionRefId != nil {
		return true
	}

	return false
}

// SetSubscriptionRefId gets a reference to the given string and assigns it to the SubscriptionRefId field.
func (o *AssetDeploymentAllOf) SetSubscriptionRefId(v string) {
	o.SubscriptionRefId = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentAllOf) GetUnitOfMeasure() []AssetMeteringType {
	if o == nil {
		var ret []AssetMeteringType
		return ret
	}
	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentAllOf) GetUnitOfMeasureOk() ([]AssetMeteringType, bool) {
	if o == nil || o.UnitOfMeasure == nil {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasUnitOfMeasure() bool {
	if o != nil && o.UnitOfMeasure != nil {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given []AssetMeteringType and assigns it to the UnitOfMeasure field.
func (o *AssetDeploymentAllOf) SetUnitOfMeasure(v []AssetMeteringType) {
	o.UnitOfMeasure = v
}

// GetWorkloads returns the Workloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentAllOf) GetWorkloads() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentAllOf) GetWorkloadsOk() ([]string, bool) {
	if o == nil || o.Workloads == nil {
		return nil, false
	}
	return o.Workloads, true
}

// HasWorkloads returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasWorkloads() bool {
	if o != nil && o.Workloads != nil {
		return true
	}

	return false
}

// SetWorkloads gets a reference to the given []string and assigns it to the Workloads field.
func (o *AssetDeploymentAllOf) SetWorkloads(v []string) {
	o.Workloads = v
}

// GetDevices returns the Devices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentAllOf) GetDevices() []AssetDeploymentDeviceRelationship {
	if o == nil {
		var ret []AssetDeploymentDeviceRelationship
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentAllOf) GetDevicesOk() ([]AssetDeploymentDeviceRelationship, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []AssetDeploymentDeviceRelationship and assigns it to the Devices field.
func (o *AssetDeploymentAllOf) SetDevices(v []AssetDeploymentDeviceRelationship) {
	o.Devices = v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *AssetDeploymentAllOf) GetSubscription() AssetSubscriptionRelationship {
	if o == nil || o.Subscription == nil {
		var ret AssetSubscriptionRelationship
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentAllOf) GetSubscriptionOk() (*AssetSubscriptionRelationship, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *AssetDeploymentAllOf) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given AssetSubscriptionRelationship and assigns it to the Subscription field.
func (o *AssetDeploymentAllOf) SetSubscription(v AssetSubscriptionRelationship) {
	o.Subscription = &v
}

func (o AssetDeploymentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlarmInfo.IsSet() {
		toSerialize["AlarmInfo"] = o.AlarmInfo.Get()
	}
	if o.DeploymentRefId != nil {
		toSerialize["DeploymentRefId"] = o.DeploymentRefId
	}
	if o.EndCustomer.IsSet() {
		toSerialize["EndCustomer"] = o.EndCustomer.Get()
	}
	if o.EndDate != nil {
		toSerialize["EndDate"] = o.EndDate
	}
	if o.LicenseType != nil {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.MlbOfferType != nil {
		toSerialize["MlbOfferType"] = o.MlbOfferType
	}
	if o.StartDate != nil {
		toSerialize["StartDate"] = o.StartDate
	}
	if o.SubscriptionRefId != nil {
		toSerialize["SubscriptionRefId"] = o.SubscriptionRefId
	}
	if o.UnitOfMeasure != nil {
		toSerialize["UnitOfMeasure"] = o.UnitOfMeasure
	}
	if o.Workloads != nil {
		toSerialize["Workloads"] = o.Workloads
	}
	if o.Devices != nil {
		toSerialize["Devices"] = o.Devices
	}
	if o.Subscription != nil {
		toSerialize["Subscription"] = o.Subscription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeploymentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetDeploymentAllOf := _AssetDeploymentAllOf{}

	if err = json.Unmarshal(bytes, &varAssetDeploymentAllOf); err == nil {
		*o = AssetDeploymentAllOf(varAssetDeploymentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmInfo")
		delete(additionalProperties, "DeploymentRefId")
		delete(additionalProperties, "EndCustomer")
		delete(additionalProperties, "EndDate")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "MlbOfferType")
		delete(additionalProperties, "StartDate")
		delete(additionalProperties, "SubscriptionRefId")
		delete(additionalProperties, "UnitOfMeasure")
		delete(additionalProperties, "Workloads")
		delete(additionalProperties, "Devices")
		delete(additionalProperties, "Subscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDeploymentAllOf struct {
	value *AssetDeploymentAllOf
	isSet bool
}

func (v NullableAssetDeploymentAllOf) Get() *AssetDeploymentAllOf {
	return v.value
}

func (v *NullableAssetDeploymentAllOf) Set(val *AssetDeploymentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeploymentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeploymentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeploymentAllOf(val *AssetDeploymentAllOf) *NullableAssetDeploymentAllOf {
	return &NullableAssetDeploymentAllOf{value: val, isSet: true}
}

func (v NullableAssetDeploymentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeploymentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
