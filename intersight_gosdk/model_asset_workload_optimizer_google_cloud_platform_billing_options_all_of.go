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

// AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Google Cloud Platform (GCP) Billing Account ID.
	BillingAccountId *string `json:"BillingAccountId,omitempty"`
	// Name of Google BigQuery Cost Export Data Set which is the dataset for the billed cost export.
	CostExportDataSetName *string `json:"CostExportDataSetName,omitempty"`
	// Google BigQuery Cost Export Table Name. This table will store the exported data from the Cost Export.
	CostExportTableName *string `json:"CostExportTableName,omitempty"`
	// Name of the BigQuery Pricing Export Data Set which is the dataset for negotiated pricing.
	PricingExportDataSetName *string `json:"PricingExportDataSetName,omitempty"`
	// The Google BigQuery Pricing Export Table Name field is auto-populated with the table used in BigQuery, cloud_pricing_export. There is no need to modify this name, unless you use a different table for negotiated pricing. The Default name is \"cloud_pricing_export\".
	PricingExportTableName *string `json:"PricingExportTableName,omitempty"`
	// The unique ID assigned to the project containing the cost and pricing exports. If the exports are in separate projects, multiple billing targets will be necessary.
	ProjectId *string `json:"ProjectId,omitempty"`
	// This flag will enable querying of detailed usage cost with resource level information included.  If not enabled, cost export data will be queried, if dataset and table name are provided.
	ResourceLevelCostEnabled *bool `json:"ResourceLevelCostEnabled,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf

// NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf instantiates a new AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf {
	this := AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOfWithDefaults() *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf {
	this := AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetBillingAccountId() string {
	if o == nil || o.BillingAccountId == nil {
		var ret string
		return ret
	}
	return *o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetBillingAccountIdOk() (*string, bool) {
	if o == nil || o.BillingAccountId == nil {
		return nil, false
	}
	return o.BillingAccountId, true
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasBillingAccountId() bool {
	if o != nil && o.BillingAccountId != nil {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given string and assigns it to the BillingAccountId field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetBillingAccountId(v string) {
	o.BillingAccountId = &v
}

// GetCostExportDataSetName returns the CostExportDataSetName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetCostExportDataSetName() string {
	if o == nil || o.CostExportDataSetName == nil {
		var ret string
		return ret
	}
	return *o.CostExportDataSetName
}

// GetCostExportDataSetNameOk returns a tuple with the CostExportDataSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetCostExportDataSetNameOk() (*string, bool) {
	if o == nil || o.CostExportDataSetName == nil {
		return nil, false
	}
	return o.CostExportDataSetName, true
}

// HasCostExportDataSetName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasCostExportDataSetName() bool {
	if o != nil && o.CostExportDataSetName != nil {
		return true
	}

	return false
}

// SetCostExportDataSetName gets a reference to the given string and assigns it to the CostExportDataSetName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetCostExportDataSetName(v string) {
	o.CostExportDataSetName = &v
}

// GetCostExportTableName returns the CostExportTableName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetCostExportTableName() string {
	if o == nil || o.CostExportTableName == nil {
		var ret string
		return ret
	}
	return *o.CostExportTableName
}

// GetCostExportTableNameOk returns a tuple with the CostExportTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetCostExportTableNameOk() (*string, bool) {
	if o == nil || o.CostExportTableName == nil {
		return nil, false
	}
	return o.CostExportTableName, true
}

// HasCostExportTableName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasCostExportTableName() bool {
	if o != nil && o.CostExportTableName != nil {
		return true
	}

	return false
}

// SetCostExportTableName gets a reference to the given string and assigns it to the CostExportTableName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetCostExportTableName(v string) {
	o.CostExportTableName = &v
}

// GetPricingExportDataSetName returns the PricingExportDataSetName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetPricingExportDataSetName() string {
	if o == nil || o.PricingExportDataSetName == nil {
		var ret string
		return ret
	}
	return *o.PricingExportDataSetName
}

// GetPricingExportDataSetNameOk returns a tuple with the PricingExportDataSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetPricingExportDataSetNameOk() (*string, bool) {
	if o == nil || o.PricingExportDataSetName == nil {
		return nil, false
	}
	return o.PricingExportDataSetName, true
}

// HasPricingExportDataSetName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasPricingExportDataSetName() bool {
	if o != nil && o.PricingExportDataSetName != nil {
		return true
	}

	return false
}

// SetPricingExportDataSetName gets a reference to the given string and assigns it to the PricingExportDataSetName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetPricingExportDataSetName(v string) {
	o.PricingExportDataSetName = &v
}

// GetPricingExportTableName returns the PricingExportTableName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetPricingExportTableName() string {
	if o == nil || o.PricingExportTableName == nil {
		var ret string
		return ret
	}
	return *o.PricingExportTableName
}

// GetPricingExportTableNameOk returns a tuple with the PricingExportTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetPricingExportTableNameOk() (*string, bool) {
	if o == nil || o.PricingExportTableName == nil {
		return nil, false
	}
	return o.PricingExportTableName, true
}

// HasPricingExportTableName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasPricingExportTableName() bool {
	if o != nil && o.PricingExportTableName != nil {
		return true
	}

	return false
}

// SetPricingExportTableName gets a reference to the given string and assigns it to the PricingExportTableName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetPricingExportTableName(v string) {
	o.PricingExportTableName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceLevelCostEnabled returns the ResourceLevelCostEnabled field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetResourceLevelCostEnabled() bool {
	if o == nil || o.ResourceLevelCostEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceLevelCostEnabled
}

// GetResourceLevelCostEnabledOk returns a tuple with the ResourceLevelCostEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) GetResourceLevelCostEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceLevelCostEnabled == nil {
		return nil, false
	}
	return o.ResourceLevelCostEnabled, true
}

// HasResourceLevelCostEnabled returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) HasResourceLevelCostEnabled() bool {
	if o != nil && o.ResourceLevelCostEnabled != nil {
		return true
	}

	return false
}

// SetResourceLevelCostEnabled gets a reference to the given bool and assigns it to the ResourceLevelCostEnabled field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) SetResourceLevelCostEnabled(v bool) {
	o.ResourceLevelCostEnabled = &v
}

func (o AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillingAccountId != nil {
		toSerialize["BillingAccountId"] = o.BillingAccountId
	}
	if o.CostExportDataSetName != nil {
		toSerialize["CostExportDataSetName"] = o.CostExportDataSetName
	}
	if o.CostExportTableName != nil {
		toSerialize["CostExportTableName"] = o.CostExportTableName
	}
	if o.PricingExportDataSetName != nil {
		toSerialize["PricingExportDataSetName"] = o.PricingExportDataSetName
	}
	if o.PricingExportTableName != nil {
		toSerialize["PricingExportTableName"] = o.PricingExportTableName
	}
	if o.ProjectId != nil {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if o.ResourceLevelCostEnabled != nil {
		toSerialize["ResourceLevelCostEnabled"] = o.ResourceLevelCostEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf := _AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf(varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingAccountId")
		delete(additionalProperties, "CostExportDataSetName")
		delete(additionalProperties, "CostExportTableName")
		delete(additionalProperties, "PricingExportDataSetName")
		delete(additionalProperties, "PricingExportTableName")
		delete(additionalProperties, "ProjectId")
		delete(additionalProperties, "ResourceLevelCostEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf struct {
	value *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) Get() *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) Set(val *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf(val *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf {
	return &NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
