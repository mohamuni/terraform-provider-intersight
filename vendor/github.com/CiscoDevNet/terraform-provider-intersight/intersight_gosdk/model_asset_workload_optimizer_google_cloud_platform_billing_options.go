/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions{}

// AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions Captures configuration specific to the Google Cloud Platform (GCP) Billing target for Workload Optimizer service.
type AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions struct {
	AssetServiceOptions
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
	// This flag will enable querying of detailed usage cost with resource level information included. If not enabled, cost export data will be queried, if dataset and table name are provided.
	ResourceLevelCostEnabled *bool `json:"ResourceLevelCostEnabled,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions

// NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions instantiates a new AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions(classId string, objectType string) *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions {
	this := AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithDefaults instantiates a new AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithDefaults() *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions {
	this := AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions{}
	var classId string = "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions" of the ClassId field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetDefaultClassId() interface{} {
	return "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions" of the ObjectType field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetDefaultObjectType() interface{} {
	return "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetBillingAccountId() string {
	if o == nil || IsNil(o.BillingAccountId) {
		var ret string
		return ret
	}
	return *o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetBillingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAccountId) {
		return nil, false
	}
	return o.BillingAccountId, true
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasBillingAccountId() bool {
	if o != nil && !IsNil(o.BillingAccountId) {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given string and assigns it to the BillingAccountId field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetBillingAccountId(v string) {
	o.BillingAccountId = &v
}

// GetCostExportDataSetName returns the CostExportDataSetName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetCostExportDataSetName() string {
	if o == nil || IsNil(o.CostExportDataSetName) {
		var ret string
		return ret
	}
	return *o.CostExportDataSetName
}

// GetCostExportDataSetNameOk returns a tuple with the CostExportDataSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetCostExportDataSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.CostExportDataSetName) {
		return nil, false
	}
	return o.CostExportDataSetName, true
}

// HasCostExportDataSetName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasCostExportDataSetName() bool {
	if o != nil && !IsNil(o.CostExportDataSetName) {
		return true
	}

	return false
}

// SetCostExportDataSetName gets a reference to the given string and assigns it to the CostExportDataSetName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetCostExportDataSetName(v string) {
	o.CostExportDataSetName = &v
}

// GetCostExportTableName returns the CostExportTableName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetCostExportTableName() string {
	if o == nil || IsNil(o.CostExportTableName) {
		var ret string
		return ret
	}
	return *o.CostExportTableName
}

// GetCostExportTableNameOk returns a tuple with the CostExportTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetCostExportTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.CostExportTableName) {
		return nil, false
	}
	return o.CostExportTableName, true
}

// HasCostExportTableName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasCostExportTableName() bool {
	if o != nil && !IsNil(o.CostExportTableName) {
		return true
	}

	return false
}

// SetCostExportTableName gets a reference to the given string and assigns it to the CostExportTableName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetCostExportTableName(v string) {
	o.CostExportTableName = &v
}

// GetPricingExportDataSetName returns the PricingExportDataSetName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportDataSetName() string {
	if o == nil || IsNil(o.PricingExportDataSetName) {
		var ret string
		return ret
	}
	return *o.PricingExportDataSetName
}

// GetPricingExportDataSetNameOk returns a tuple with the PricingExportDataSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportDataSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.PricingExportDataSetName) {
		return nil, false
	}
	return o.PricingExportDataSetName, true
}

// HasPricingExportDataSetName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasPricingExportDataSetName() bool {
	if o != nil && !IsNil(o.PricingExportDataSetName) {
		return true
	}

	return false
}

// SetPricingExportDataSetName gets a reference to the given string and assigns it to the PricingExportDataSetName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetPricingExportDataSetName(v string) {
	o.PricingExportDataSetName = &v
}

// GetPricingExportTableName returns the PricingExportTableName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportTableName() string {
	if o == nil || IsNil(o.PricingExportTableName) {
		var ret string
		return ret
	}
	return *o.PricingExportTableName
}

// GetPricingExportTableNameOk returns a tuple with the PricingExportTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.PricingExportTableName) {
		return nil, false
	}
	return o.PricingExportTableName, true
}

// HasPricingExportTableName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasPricingExportTableName() bool {
	if o != nil && !IsNil(o.PricingExportTableName) {
		return true
	}

	return false
}

// SetPricingExportTableName gets a reference to the given string and assigns it to the PricingExportTableName field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetPricingExportTableName(v string) {
	o.PricingExportTableName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceLevelCostEnabled returns the ResourceLevelCostEnabled field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetResourceLevelCostEnabled() bool {
	if o == nil || IsNil(o.ResourceLevelCostEnabled) {
		var ret bool
		return ret
	}
	return *o.ResourceLevelCostEnabled
}

// GetResourceLevelCostEnabledOk returns a tuple with the ResourceLevelCostEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetResourceLevelCostEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ResourceLevelCostEnabled) {
		return nil, false
	}
	return o.ResourceLevelCostEnabled, true
}

// HasResourceLevelCostEnabled returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasResourceLevelCostEnabled() bool {
	if o != nil && !IsNil(o.ResourceLevelCostEnabled) {
		return true
	}

	return false
}

// SetResourceLevelCostEnabled gets a reference to the given bool and assigns it to the ResourceLevelCostEnabled field.
func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetResourceLevelCostEnabled(v bool) {
	o.ResourceLevelCostEnabled = &v
}

func (o AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BillingAccountId) {
		toSerialize["BillingAccountId"] = o.BillingAccountId
	}
	if !IsNil(o.CostExportDataSetName) {
		toSerialize["CostExportDataSetName"] = o.CostExportDataSetName
	}
	if !IsNil(o.CostExportTableName) {
		toSerialize["CostExportTableName"] = o.CostExportTableName
	}
	if !IsNil(o.PricingExportDataSetName) {
		toSerialize["PricingExportDataSetName"] = o.PricingExportDataSetName
	}
	if !IsNil(o.PricingExportTableName) {
		toSerialize["PricingExportTableName"] = o.PricingExportTableName
	}
	if !IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if !IsNil(o.ResourceLevelCostEnabled) {
		toSerialize["ResourceLevelCostEnabled"] = o.ResourceLevelCostEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) UnmarshalJSON(data []byte) (err error) {
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
	type AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct struct {
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
		// This flag will enable querying of detailed usage cost with resource level information included. If not enabled, cost export data will be queried, if dataset and table name are provided.
		ResourceLevelCostEnabled *bool `json:"ResourceLevelCostEnabled,omitempty"`
	}

	varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions := _AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions{}
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.ClassId = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.ObjectType = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.BillingAccountId = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.BillingAccountId
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.CostExportDataSetName = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.CostExportDataSetName
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.CostExportTableName = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.CostExportTableName
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.PricingExportDataSetName = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.PricingExportDataSetName
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.PricingExportTableName = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.PricingExportTableName
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.ProjectId = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.ProjectId
		varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.ResourceLevelCostEnabled = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithoutEmbeddedStruct.ResourceLevelCostEnabled
		*o = AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions(varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions := _AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions{}

	err = json.Unmarshal(data, &varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingAccountId")
		delete(additionalProperties, "CostExportDataSetName")
		delete(additionalProperties, "CostExportTableName")
		delete(additionalProperties, "PricingExportDataSetName")
		delete(additionalProperties, "PricingExportTableName")
		delete(additionalProperties, "ProjectId")
		delete(additionalProperties, "ResourceLevelCostEnabled")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions struct {
	value *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) Get() *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) Set(val *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions(val *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions {
	return &NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
