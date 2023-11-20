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
	"reflect"
	"strings"
)

// MonitoringCategoryStatus Status of an Intersight category such as 'Licensing' with additional details. Several categories and their respective status is reported as part of the health status API.
type MonitoringCategoryStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the category for which status is being reported (for e.g. 'Licensing', 'Advisories').
	CategoryLabel *string `json:"CategoryLabel,omitempty"`
	// Link to the corresponding category specific page in Intersight to get additional information and troubleshoot. for e.g. 'Alarms' category would have the deeplink as 'https://intersight.com/an/cond/alarms/active'.
	Deeplink *string `json:"Deeplink,omitempty"`
	// Additional information regarding category status that may be displayed in UI related to the category status. Optional and currently unused.
	Details *string `json:"Details,omitempty"`
	// Additional parameter to be used for traceability and troubleshooting, currently unused.
	SourceId *string `json:"SourceId,omitempty"`
	// Aggregated status of the category being reported. For e.g., a given Intersight account might have a combination of high and low severity Advisories applicable to managed devices. However, even if one of the devices is impacted by a high severity Advisories, the category status is reported as 'critical'. * `Unknown` - The current status for the respective category cannot be determined. This may be due to some intermittent issue or due to the fact that the user making the request does not have appropriate privileges/entitlements for the concerned category. * `Critical` - The Intersight account is impacted by a high severity issue  for the applicable category that will need to be addressed immediately. for e.g. critical status for 'Advisories' category would mean that the Intersight account is impacted by an urgent field notice or a vulnerability with high severity level. * `Warning` - The Intersight account is impacted by a high severity issue  for the applicable category that will need to be addressed soon. for e.g. warning status for 'Advisories' category might mean that the Intersight account is impacted by an a vulnerability with moderate severity level. * `Healthy` - The Intersight account is not impacted by any high severity issue for the applicable category.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MonitoringCategoryStatus MonitoringCategoryStatus

// NewMonitoringCategoryStatus instantiates a new MonitoringCategoryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringCategoryStatus(classId string, objectType string) *MonitoringCategoryStatus {
	this := MonitoringCategoryStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMonitoringCategoryStatusWithDefaults instantiates a new MonitoringCategoryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringCategoryStatusWithDefaults() *MonitoringCategoryStatus {
	this := MonitoringCategoryStatus{}
	var classId string = "monitoring.CategoryStatus"
	this.ClassId = classId
	var objectType string = "monitoring.CategoryStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MonitoringCategoryStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MonitoringCategoryStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MonitoringCategoryStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MonitoringCategoryStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategoryLabel returns the CategoryLabel field value if set, zero value otherwise.
func (o *MonitoringCategoryStatus) GetCategoryLabel() string {
	if o == nil || o.CategoryLabel == nil {
		var ret string
		return ret
	}
	return *o.CategoryLabel
}

// GetCategoryLabelOk returns a tuple with the CategoryLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetCategoryLabelOk() (*string, bool) {
	if o == nil || o.CategoryLabel == nil {
		return nil, false
	}
	return o.CategoryLabel, true
}

// HasCategoryLabel returns a boolean if a field has been set.
func (o *MonitoringCategoryStatus) HasCategoryLabel() bool {
	if o != nil && o.CategoryLabel != nil {
		return true
	}

	return false
}

// SetCategoryLabel gets a reference to the given string and assigns it to the CategoryLabel field.
func (o *MonitoringCategoryStatus) SetCategoryLabel(v string) {
	o.CategoryLabel = &v
}

// GetDeeplink returns the Deeplink field value if set, zero value otherwise.
func (o *MonitoringCategoryStatus) GetDeeplink() string {
	if o == nil || o.Deeplink == nil {
		var ret string
		return ret
	}
	return *o.Deeplink
}

// GetDeeplinkOk returns a tuple with the Deeplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetDeeplinkOk() (*string, bool) {
	if o == nil || o.Deeplink == nil {
		return nil, false
	}
	return o.Deeplink, true
}

// HasDeeplink returns a boolean if a field has been set.
func (o *MonitoringCategoryStatus) HasDeeplink() bool {
	if o != nil && o.Deeplink != nil {
		return true
	}

	return false
}

// SetDeeplink gets a reference to the given string and assigns it to the Deeplink field.
func (o *MonitoringCategoryStatus) SetDeeplink(v string) {
	o.Deeplink = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MonitoringCategoryStatus) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MonitoringCategoryStatus) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *MonitoringCategoryStatus) SetDetails(v string) {
	o.Details = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *MonitoringCategoryStatus) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *MonitoringCategoryStatus) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *MonitoringCategoryStatus) SetSourceId(v string) {
	o.SourceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitoringCategoryStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCategoryStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitoringCategoryStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MonitoringCategoryStatus) SetStatus(v string) {
	o.Status = &v
}

func (o MonitoringCategoryStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CategoryLabel != nil {
		toSerialize["CategoryLabel"] = o.CategoryLabel
	}
	if o.Deeplink != nil {
		toSerialize["Deeplink"] = o.Deeplink
	}
	if o.Details != nil {
		toSerialize["Details"] = o.Details
	}
	if o.SourceId != nil {
		toSerialize["SourceId"] = o.SourceId
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MonitoringCategoryStatus) UnmarshalJSON(bytes []byte) (err error) {
	type MonitoringCategoryStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the category for which status is being reported (for e.g. 'Licensing', 'Advisories').
		CategoryLabel *string `json:"CategoryLabel,omitempty"`
		// Link to the corresponding category specific page in Intersight to get additional information and troubleshoot. for e.g. 'Alarms' category would have the deeplink as 'https://intersight.com/an/cond/alarms/active'.
		Deeplink *string `json:"Deeplink,omitempty"`
		// Additional information regarding category status that may be displayed in UI related to the category status. Optional and currently unused.
		Details *string `json:"Details,omitempty"`
		// Additional parameter to be used for traceability and troubleshooting, currently unused.
		SourceId *string `json:"SourceId,omitempty"`
		// Aggregated status of the category being reported. For e.g., a given Intersight account might have a combination of high and low severity Advisories applicable to managed devices. However, even if one of the devices is impacted by a high severity Advisories, the category status is reported as 'critical'. * `Unknown` - The current status for the respective category cannot be determined. This may be due to some intermittent issue or due to the fact that the user making the request does not have appropriate privileges/entitlements for the concerned category. * `Critical` - The Intersight account is impacted by a high severity issue  for the applicable category that will need to be addressed immediately. for e.g. critical status for 'Advisories' category would mean that the Intersight account is impacted by an urgent field notice or a vulnerability with high severity level. * `Warning` - The Intersight account is impacted by a high severity issue  for the applicable category that will need to be addressed soon. for e.g. warning status for 'Advisories' category might mean that the Intersight account is impacted by an a vulnerability with moderate severity level. * `Healthy` - The Intersight account is not impacted by any high severity issue for the applicable category.
		Status *string `json:"Status,omitempty"`
	}

	varMonitoringCategoryStatusWithoutEmbeddedStruct := MonitoringCategoryStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMonitoringCategoryStatusWithoutEmbeddedStruct)
	if err == nil {
		varMonitoringCategoryStatus := _MonitoringCategoryStatus{}
		varMonitoringCategoryStatus.ClassId = varMonitoringCategoryStatusWithoutEmbeddedStruct.ClassId
		varMonitoringCategoryStatus.ObjectType = varMonitoringCategoryStatusWithoutEmbeddedStruct.ObjectType
		varMonitoringCategoryStatus.CategoryLabel = varMonitoringCategoryStatusWithoutEmbeddedStruct.CategoryLabel
		varMonitoringCategoryStatus.Deeplink = varMonitoringCategoryStatusWithoutEmbeddedStruct.Deeplink
		varMonitoringCategoryStatus.Details = varMonitoringCategoryStatusWithoutEmbeddedStruct.Details
		varMonitoringCategoryStatus.SourceId = varMonitoringCategoryStatusWithoutEmbeddedStruct.SourceId
		varMonitoringCategoryStatus.Status = varMonitoringCategoryStatusWithoutEmbeddedStruct.Status
		*o = MonitoringCategoryStatus(varMonitoringCategoryStatus)
	} else {
		return err
	}

	varMonitoringCategoryStatus := _MonitoringCategoryStatus{}

	err = json.Unmarshal(bytes, &varMonitoringCategoryStatus)
	if err == nil {
		o.MoBaseComplexType = varMonitoringCategoryStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CategoryLabel")
		delete(additionalProperties, "Deeplink")
		delete(additionalProperties, "Details")
		delete(additionalProperties, "SourceId")
		delete(additionalProperties, "Status")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableMonitoringCategoryStatus struct {
	value *MonitoringCategoryStatus
	isSet bool
}

func (v NullableMonitoringCategoryStatus) Get() *MonitoringCategoryStatus {
	return v.value
}

func (v *NullableMonitoringCategoryStatus) Set(val *MonitoringCategoryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringCategoryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringCategoryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringCategoryStatus(val *MonitoringCategoryStatus) *NullableMonitoringCategoryStatus {
	return &NullableMonitoringCategoryStatus{value: val, isSet: true}
}

func (v NullableMonitoringCategoryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringCategoryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
