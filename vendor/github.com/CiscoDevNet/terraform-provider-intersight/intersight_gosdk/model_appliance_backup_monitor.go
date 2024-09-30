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

// checks if the ApplianceBackupMonitor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceBackupMonitor{}

// ApplianceBackupMonitor BackupMonitor keeps track of the Appliance's backup history and sets the status to BackupFound, BackupFailed, or  BackupNotFound based on when the last backup was scheduled.
type ApplianceBackupMonitor struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Filename of the backup for the backup monitor.
	Filename *string `json:"Filename,omitempty"`
	// Status of the oldest Intersight Appliance backup cleanup. * `BackupFound` - Backup is successful and complete. * `BackupFailed` - The current Backup failed. * `BackupOutdated` - Backup is old and outdated. * `BackupCleanupFailed` - Cleanup of the old backup has failed.
	LastBackupRotationStatus *string `json:"LastBackupRotationStatus,omitempty"`
	// Status of the most recent Intersight Appliance backup. * `BackupFound` - Backup is successful and complete. * `BackupFailed` - The current Backup failed. * `BackupOutdated` - Backup is old and outdated. * `BackupCleanupFailed` - Cleanup of the old backup has failed.
	LastBackupStatus     *string                        `json:"LastBackupStatus,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceBackupMonitor ApplianceBackupMonitor

// NewApplianceBackupMonitor instantiates a new ApplianceBackupMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceBackupMonitor(classId string, objectType string) *ApplianceBackupMonitor {
	this := ApplianceBackupMonitor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceBackupMonitorWithDefaults instantiates a new ApplianceBackupMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceBackupMonitorWithDefaults() *ApplianceBackupMonitor {
	this := ApplianceBackupMonitor{}
	var classId string = "appliance.BackupMonitor"
	this.ClassId = classId
	var objectType string = "appliance.BackupMonitor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceBackupMonitor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceBackupMonitor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.BackupMonitor" of the ClassId field.
func (o *ApplianceBackupMonitor) GetDefaultClassId() interface{} {
	return "appliance.BackupMonitor"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceBackupMonitor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceBackupMonitor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.BackupMonitor" of the ObjectType field.
func (o *ApplianceBackupMonitor) GetDefaultObjectType() interface{} {
	return "appliance.BackupMonitor"
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ApplianceBackupMonitor) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitor) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ApplianceBackupMonitor) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ApplianceBackupMonitor) SetFilename(v string) {
	o.Filename = &v
}

// GetLastBackupRotationStatus returns the LastBackupRotationStatus field value if set, zero value otherwise.
func (o *ApplianceBackupMonitor) GetLastBackupRotationStatus() string {
	if o == nil || IsNil(o.LastBackupRotationStatus) {
		var ret string
		return ret
	}
	return *o.LastBackupRotationStatus
}

// GetLastBackupRotationStatusOk returns a tuple with the LastBackupRotationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitor) GetLastBackupRotationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastBackupRotationStatus) {
		return nil, false
	}
	return o.LastBackupRotationStatus, true
}

// HasLastBackupRotationStatus returns a boolean if a field has been set.
func (o *ApplianceBackupMonitor) HasLastBackupRotationStatus() bool {
	if o != nil && !IsNil(o.LastBackupRotationStatus) {
		return true
	}

	return false
}

// SetLastBackupRotationStatus gets a reference to the given string and assigns it to the LastBackupRotationStatus field.
func (o *ApplianceBackupMonitor) SetLastBackupRotationStatus(v string) {
	o.LastBackupRotationStatus = &v
}

// GetLastBackupStatus returns the LastBackupStatus field value if set, zero value otherwise.
func (o *ApplianceBackupMonitor) GetLastBackupStatus() string {
	if o == nil || IsNil(o.LastBackupStatus) {
		var ret string
		return ret
	}
	return *o.LastBackupStatus
}

// GetLastBackupStatusOk returns a tuple with the LastBackupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitor) GetLastBackupStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastBackupStatus) {
		return nil, false
	}
	return o.LastBackupStatus, true
}

// HasLastBackupStatus returns a boolean if a field has been set.
func (o *ApplianceBackupMonitor) HasLastBackupStatus() bool {
	if o != nil && !IsNil(o.LastBackupStatus) {
		return true
	}

	return false
}

// SetLastBackupStatus gets a reference to the given string and assigns it to the LastBackupStatus field.
func (o *ApplianceBackupMonitor) SetLastBackupStatus(v string) {
	o.LastBackupStatus = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceBackupMonitor) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceBackupMonitor) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceBackupMonitor) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *ApplianceBackupMonitor) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *ApplianceBackupMonitor) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ApplianceBackupMonitor) UnsetAccount() {
	o.Account.Unset()
}

func (o ApplianceBackupMonitor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceBackupMonitor) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Filename) {
		toSerialize["Filename"] = o.Filename
	}
	if !IsNil(o.LastBackupRotationStatus) {
		toSerialize["LastBackupRotationStatus"] = o.LastBackupRotationStatus
	}
	if !IsNil(o.LastBackupStatus) {
		toSerialize["LastBackupStatus"] = o.LastBackupStatus
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceBackupMonitor) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceBackupMonitorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Filename of the backup for the backup monitor.
		Filename *string `json:"Filename,omitempty"`
		// Status of the oldest Intersight Appliance backup cleanup. * `BackupFound` - Backup is successful and complete. * `BackupFailed` - The current Backup failed. * `BackupOutdated` - Backup is old and outdated. * `BackupCleanupFailed` - Cleanup of the old backup has failed.
		LastBackupRotationStatus *string `json:"LastBackupRotationStatus,omitempty"`
		// Status of the most recent Intersight Appliance backup. * `BackupFound` - Backup is successful and complete. * `BackupFailed` - The current Backup failed. * `BackupOutdated` - Backup is old and outdated. * `BackupCleanupFailed` - Cleanup of the old backup has failed.
		LastBackupStatus *string                        `json:"LastBackupStatus,omitempty"`
		Account          NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceBackupMonitorWithoutEmbeddedStruct := ApplianceBackupMonitorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceBackupMonitorWithoutEmbeddedStruct)
	if err == nil {
		varApplianceBackupMonitor := _ApplianceBackupMonitor{}
		varApplianceBackupMonitor.ClassId = varApplianceBackupMonitorWithoutEmbeddedStruct.ClassId
		varApplianceBackupMonitor.ObjectType = varApplianceBackupMonitorWithoutEmbeddedStruct.ObjectType
		varApplianceBackupMonitor.Filename = varApplianceBackupMonitorWithoutEmbeddedStruct.Filename
		varApplianceBackupMonitor.LastBackupRotationStatus = varApplianceBackupMonitorWithoutEmbeddedStruct.LastBackupRotationStatus
		varApplianceBackupMonitor.LastBackupStatus = varApplianceBackupMonitorWithoutEmbeddedStruct.LastBackupStatus
		varApplianceBackupMonitor.Account = varApplianceBackupMonitorWithoutEmbeddedStruct.Account
		*o = ApplianceBackupMonitor(varApplianceBackupMonitor)
	} else {
		return err
	}

	varApplianceBackupMonitor := _ApplianceBackupMonitor{}

	err = json.Unmarshal(data, &varApplianceBackupMonitor)
	if err == nil {
		o.MoBaseMo = varApplianceBackupMonitor.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "LastBackupRotationStatus")
		delete(additionalProperties, "LastBackupStatus")
		delete(additionalProperties, "Account")

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

type NullableApplianceBackupMonitor struct {
	value *ApplianceBackupMonitor
	isSet bool
}

func (v NullableApplianceBackupMonitor) Get() *ApplianceBackupMonitor {
	return v.value
}

func (v *NullableApplianceBackupMonitor) Set(val *ApplianceBackupMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceBackupMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceBackupMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceBackupMonitor(val *ApplianceBackupMonitor) *NullableApplianceBackupMonitor {
	return &NullableApplianceBackupMonitor{value: val, isSet: true}
}

func (v NullableApplianceBackupMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceBackupMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
