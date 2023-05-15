/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// MonitoringHealthStatus High level, aggregated status of Intersight components for a given Intersight account user. Meant to inform the user if there's an issue with Intersight components that needs her attention. At this point, Aggregated status is reported for 'Licensing', 'Advisories' and 'Alarms' components. Specifically designed to be easily consumed by external dashboards to display an at-a-glance status of Intersight components. This conforms to the health data API schema published as part of Cisco PlatformSuite.
type MonitoringHealthStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string                     `json:"ObjectType"`
	CategoryStatus []MonitoringCategoryStatus `json:"CategoryStatus,omitempty"`
	// Version of compliant health data API schema.
	HealthDataSchemaVersion *string `json:"HealthDataSchemaVersion,omitempty"`
	// Set as 'Intersight'. Especially useful in cases such as when this API is consumed by an external dashboard. This field allows such dashboards to aggregate health status across multiple  sources (Intersight, Meraki etc.).
	Source *string `json:"Source,omitempty"`
	// Time stamp when the status was generated. The status reported by this API may lag the real time status by up to 5 minutes.
	StatusTimeStamp      *time.Time                            `json:"StatusTimeStamp,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MonitoringHealthStatus MonitoringHealthStatus

// NewMonitoringHealthStatus instantiates a new MonitoringHealthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringHealthStatus(classId string, objectType string) *MonitoringHealthStatus {
	this := MonitoringHealthStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMonitoringHealthStatusWithDefaults instantiates a new MonitoringHealthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringHealthStatusWithDefaults() *MonitoringHealthStatus {
	this := MonitoringHealthStatus{}
	var classId string = "monitoring.HealthStatus"
	this.ClassId = classId
	var objectType string = "monitoring.HealthStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MonitoringHealthStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MonitoringHealthStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MonitoringHealthStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MonitoringHealthStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MonitoringHealthStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MonitoringHealthStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategoryStatus returns the CategoryStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringHealthStatus) GetCategoryStatus() []MonitoringCategoryStatus {
	if o == nil {
		var ret []MonitoringCategoryStatus
		return ret
	}
	return o.CategoryStatus
}

// GetCategoryStatusOk returns a tuple with the CategoryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringHealthStatus) GetCategoryStatusOk() ([]MonitoringCategoryStatus, bool) {
	if o == nil || o.CategoryStatus == nil {
		return nil, false
	}
	return o.CategoryStatus, true
}

// HasCategoryStatus returns a boolean if a field has been set.
func (o *MonitoringHealthStatus) HasCategoryStatus() bool {
	if o != nil && o.CategoryStatus != nil {
		return true
	}

	return false
}

// SetCategoryStatus gets a reference to the given []MonitoringCategoryStatus and assigns it to the CategoryStatus field.
func (o *MonitoringHealthStatus) SetCategoryStatus(v []MonitoringCategoryStatus) {
	o.CategoryStatus = v
}

// GetHealthDataSchemaVersion returns the HealthDataSchemaVersion field value if set, zero value otherwise.
func (o *MonitoringHealthStatus) GetHealthDataSchemaVersion() string {
	if o == nil || o.HealthDataSchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.HealthDataSchemaVersion
}

// GetHealthDataSchemaVersionOk returns a tuple with the HealthDataSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringHealthStatus) GetHealthDataSchemaVersionOk() (*string, bool) {
	if o == nil || o.HealthDataSchemaVersion == nil {
		return nil, false
	}
	return o.HealthDataSchemaVersion, true
}

// HasHealthDataSchemaVersion returns a boolean if a field has been set.
func (o *MonitoringHealthStatus) HasHealthDataSchemaVersion() bool {
	if o != nil && o.HealthDataSchemaVersion != nil {
		return true
	}

	return false
}

// SetHealthDataSchemaVersion gets a reference to the given string and assigns it to the HealthDataSchemaVersion field.
func (o *MonitoringHealthStatus) SetHealthDataSchemaVersion(v string) {
	o.HealthDataSchemaVersion = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MonitoringHealthStatus) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringHealthStatus) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MonitoringHealthStatus) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *MonitoringHealthStatus) SetSource(v string) {
	o.Source = &v
}

// GetStatusTimeStamp returns the StatusTimeStamp field value if set, zero value otherwise.
func (o *MonitoringHealthStatus) GetStatusTimeStamp() time.Time {
	if o == nil || o.StatusTimeStamp == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusTimeStamp
}

// GetStatusTimeStampOk returns a tuple with the StatusTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringHealthStatus) GetStatusTimeStampOk() (*time.Time, bool) {
	if o == nil || o.StatusTimeStamp == nil {
		return nil, false
	}
	return o.StatusTimeStamp, true
}

// HasStatusTimeStamp returns a boolean if a field has been set.
func (o *MonitoringHealthStatus) HasStatusTimeStamp() bool {
	if o != nil && o.StatusTimeStamp != nil {
		return true
	}

	return false
}

// SetStatusTimeStamp gets a reference to the given time.Time and assigns it to the StatusTimeStamp field.
func (o *MonitoringHealthStatus) SetStatusTimeStamp(v time.Time) {
	o.StatusTimeStamp = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *MonitoringHealthStatus) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringHealthStatus) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *MonitoringHealthStatus) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *MonitoringHealthStatus) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o MonitoringHealthStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CategoryStatus != nil {
		toSerialize["CategoryStatus"] = o.CategoryStatus
	}
	if o.HealthDataSchemaVersion != nil {
		toSerialize["HealthDataSchemaVersion"] = o.HealthDataSchemaVersion
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.StatusTimeStamp != nil {
		toSerialize["StatusTimeStamp"] = o.StatusTimeStamp
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MonitoringHealthStatus) UnmarshalJSON(bytes []byte) (err error) {
	type MonitoringHealthStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string                     `json:"ObjectType"`
		CategoryStatus []MonitoringCategoryStatus `json:"CategoryStatus,omitempty"`
		// Version of compliant health data API schema.
		HealthDataSchemaVersion *string `json:"HealthDataSchemaVersion,omitempty"`
		// Set as 'Intersight'. Especially useful in cases such as when this API is consumed by an external dashboard. This field allows such dashboards to aggregate health status across multiple  sources (Intersight, Meraki etc.).
		Source *string `json:"Source,omitempty"`
		// Time stamp when the status was generated. The status reported by this API may lag the real time status by up to 5 minutes.
		StatusTimeStamp *time.Time                            `json:"StatusTimeStamp,omitempty"`
		Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varMonitoringHealthStatusWithoutEmbeddedStruct := MonitoringHealthStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMonitoringHealthStatusWithoutEmbeddedStruct)
	if err == nil {
		varMonitoringHealthStatus := _MonitoringHealthStatus{}
		varMonitoringHealthStatus.ClassId = varMonitoringHealthStatusWithoutEmbeddedStruct.ClassId
		varMonitoringHealthStatus.ObjectType = varMonitoringHealthStatusWithoutEmbeddedStruct.ObjectType
		varMonitoringHealthStatus.CategoryStatus = varMonitoringHealthStatusWithoutEmbeddedStruct.CategoryStatus
		varMonitoringHealthStatus.HealthDataSchemaVersion = varMonitoringHealthStatusWithoutEmbeddedStruct.HealthDataSchemaVersion
		varMonitoringHealthStatus.Source = varMonitoringHealthStatusWithoutEmbeddedStruct.Source
		varMonitoringHealthStatus.StatusTimeStamp = varMonitoringHealthStatusWithoutEmbeddedStruct.StatusTimeStamp
		varMonitoringHealthStatus.Organization = varMonitoringHealthStatusWithoutEmbeddedStruct.Organization
		*o = MonitoringHealthStatus(varMonitoringHealthStatus)
	} else {
		return err
	}

	varMonitoringHealthStatus := _MonitoringHealthStatus{}

	err = json.Unmarshal(bytes, &varMonitoringHealthStatus)
	if err == nil {
		o.MoBaseMo = varMonitoringHealthStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CategoryStatus")
		delete(additionalProperties, "HealthDataSchemaVersion")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "StatusTimeStamp")
		delete(additionalProperties, "Organization")

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

type NullableMonitoringHealthStatus struct {
	value *MonitoringHealthStatus
	isSet bool
}

func (v NullableMonitoringHealthStatus) Get() *MonitoringHealthStatus {
	return v.value
}

func (v *NullableMonitoringHealthStatus) Set(val *MonitoringHealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringHealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringHealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringHealthStatus(val *MonitoringHealthStatus) *NullableMonitoringHealthStatus {
	return &NullableMonitoringHealthStatus{value: val, isSet: true}
}

func (v NullableMonitoringHealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringHealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}