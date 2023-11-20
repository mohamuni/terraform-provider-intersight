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

// IwotenantMaintenanceNotificationAllOf Definition of the list of properties defined in 'iwotenant.MaintenanceNotification', excluding properties defined in parent classes.
type IwotenantMaintenanceNotificationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Any i18n (internationalization) key defined the message content. If the key already exists then the  message content will be picked based on the key, otherwise provided message value will be used.
	I18nKey *string `json:"I18nKey,omitempty"`
	// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
	IwoId *string `json:"IwoId,omitempty"`
	// The date/time from which the actual maintenance operations will be performed for a Customer's account.
	MaintenanceStartTime *time.Time `json:"MaintenanceStartTime,omitempty"`
	// The notification message content is to display in the UI banner after the Customer's login to inform about planned maintenance operations on IWO.
	NtfnMessage *string `json:"NtfnMessage,omitempty"`
	// The date/time from which the maintenance banner message will be shown to the Customer after login in to  Intersight UI.
	ShowFromTime *time.Time `json:"ShowFromTime,omitempty"`
	// The date/time until which the maintenance banner message will be shown to the Customer after login into  Intersight UI. This will also be the time actual maintenance operation is planned for the finish of a  Customer's account.
	ShowUntilTime        *time.Time `json:"ShowUntilTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IwotenantMaintenanceNotificationAllOf IwotenantMaintenanceNotificationAllOf

// NewIwotenantMaintenanceNotificationAllOf instantiates a new IwotenantMaintenanceNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwotenantMaintenanceNotificationAllOf(classId string, objectType string) *IwotenantMaintenanceNotificationAllOf {
	this := IwotenantMaintenanceNotificationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIwotenantMaintenanceNotificationAllOfWithDefaults instantiates a new IwotenantMaintenanceNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwotenantMaintenanceNotificationAllOfWithDefaults() *IwotenantMaintenanceNotificationAllOf {
	this := IwotenantMaintenanceNotificationAllOf{}
	var classId string = "iwotenant.MaintenanceNotification"
	this.ClassId = classId
	var objectType string = "iwotenant.MaintenanceNotification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IwotenantMaintenanceNotificationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IwotenantMaintenanceNotificationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IwotenantMaintenanceNotificationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IwotenantMaintenanceNotificationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetI18nKey returns the I18nKey field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotificationAllOf) GetI18nKey() string {
	if o == nil || o.I18nKey == nil {
		var ret string
		return ret
	}
	return *o.I18nKey
}

// GetI18nKeyOk returns a tuple with the I18nKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetI18nKeyOk() (*string, bool) {
	if o == nil || o.I18nKey == nil {
		return nil, false
	}
	return o.I18nKey, true
}

// HasI18nKey returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotificationAllOf) HasI18nKey() bool {
	if o != nil && o.I18nKey != nil {
		return true
	}

	return false
}

// SetI18nKey gets a reference to the given string and assigns it to the I18nKey field.
func (o *IwotenantMaintenanceNotificationAllOf) SetI18nKey(v string) {
	o.I18nKey = &v
}

// GetIwoId returns the IwoId field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotificationAllOf) GetIwoId() string {
	if o == nil || o.IwoId == nil {
		var ret string
		return ret
	}
	return *o.IwoId
}

// GetIwoIdOk returns a tuple with the IwoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetIwoIdOk() (*string, bool) {
	if o == nil || o.IwoId == nil {
		return nil, false
	}
	return o.IwoId, true
}

// HasIwoId returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotificationAllOf) HasIwoId() bool {
	if o != nil && o.IwoId != nil {
		return true
	}

	return false
}

// SetIwoId gets a reference to the given string and assigns it to the IwoId field.
func (o *IwotenantMaintenanceNotificationAllOf) SetIwoId(v string) {
	o.IwoId = &v
}

// GetMaintenanceStartTime returns the MaintenanceStartTime field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotificationAllOf) GetMaintenanceStartTime() time.Time {
	if o == nil || o.MaintenanceStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.MaintenanceStartTime
}

// GetMaintenanceStartTimeOk returns a tuple with the MaintenanceStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetMaintenanceStartTimeOk() (*time.Time, bool) {
	if o == nil || o.MaintenanceStartTime == nil {
		return nil, false
	}
	return o.MaintenanceStartTime, true
}

// HasMaintenanceStartTime returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotificationAllOf) HasMaintenanceStartTime() bool {
	if o != nil && o.MaintenanceStartTime != nil {
		return true
	}

	return false
}

// SetMaintenanceStartTime gets a reference to the given time.Time and assigns it to the MaintenanceStartTime field.
func (o *IwotenantMaintenanceNotificationAllOf) SetMaintenanceStartTime(v time.Time) {
	o.MaintenanceStartTime = &v
}

// GetNtfnMessage returns the NtfnMessage field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotificationAllOf) GetNtfnMessage() string {
	if o == nil || o.NtfnMessage == nil {
		var ret string
		return ret
	}
	return *o.NtfnMessage
}

// GetNtfnMessageOk returns a tuple with the NtfnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetNtfnMessageOk() (*string, bool) {
	if o == nil || o.NtfnMessage == nil {
		return nil, false
	}
	return o.NtfnMessage, true
}

// HasNtfnMessage returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotificationAllOf) HasNtfnMessage() bool {
	if o != nil && o.NtfnMessage != nil {
		return true
	}

	return false
}

// SetNtfnMessage gets a reference to the given string and assigns it to the NtfnMessage field.
func (o *IwotenantMaintenanceNotificationAllOf) SetNtfnMessage(v string) {
	o.NtfnMessage = &v
}

// GetShowFromTime returns the ShowFromTime field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotificationAllOf) GetShowFromTime() time.Time {
	if o == nil || o.ShowFromTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ShowFromTime
}

// GetShowFromTimeOk returns a tuple with the ShowFromTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetShowFromTimeOk() (*time.Time, bool) {
	if o == nil || o.ShowFromTime == nil {
		return nil, false
	}
	return o.ShowFromTime, true
}

// HasShowFromTime returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotificationAllOf) HasShowFromTime() bool {
	if o != nil && o.ShowFromTime != nil {
		return true
	}

	return false
}

// SetShowFromTime gets a reference to the given time.Time and assigns it to the ShowFromTime field.
func (o *IwotenantMaintenanceNotificationAllOf) SetShowFromTime(v time.Time) {
	o.ShowFromTime = &v
}

// GetShowUntilTime returns the ShowUntilTime field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotificationAllOf) GetShowUntilTime() time.Time {
	if o == nil || o.ShowUntilTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ShowUntilTime
}

// GetShowUntilTimeOk returns a tuple with the ShowUntilTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotificationAllOf) GetShowUntilTimeOk() (*time.Time, bool) {
	if o == nil || o.ShowUntilTime == nil {
		return nil, false
	}
	return o.ShowUntilTime, true
}

// HasShowUntilTime returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotificationAllOf) HasShowUntilTime() bool {
	if o != nil && o.ShowUntilTime != nil {
		return true
	}

	return false
}

// SetShowUntilTime gets a reference to the given time.Time and assigns it to the ShowUntilTime field.
func (o *IwotenantMaintenanceNotificationAllOf) SetShowUntilTime(v time.Time) {
	o.ShowUntilTime = &v
}

func (o IwotenantMaintenanceNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.I18nKey != nil {
		toSerialize["I18nKey"] = o.I18nKey
	}
	if o.IwoId != nil {
		toSerialize["IwoId"] = o.IwoId
	}
	if o.MaintenanceStartTime != nil {
		toSerialize["MaintenanceStartTime"] = o.MaintenanceStartTime
	}
	if o.NtfnMessage != nil {
		toSerialize["NtfnMessage"] = o.NtfnMessage
	}
	if o.ShowFromTime != nil {
		toSerialize["ShowFromTime"] = o.ShowFromTime
	}
	if o.ShowUntilTime != nil {
		toSerialize["ShowUntilTime"] = o.ShowUntilTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IwotenantMaintenanceNotificationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIwotenantMaintenanceNotificationAllOf := _IwotenantMaintenanceNotificationAllOf{}

	if err = json.Unmarshal(bytes, &varIwotenantMaintenanceNotificationAllOf); err == nil {
		*o = IwotenantMaintenanceNotificationAllOf(varIwotenantMaintenanceNotificationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "I18nKey")
		delete(additionalProperties, "IwoId")
		delete(additionalProperties, "MaintenanceStartTime")
		delete(additionalProperties, "NtfnMessage")
		delete(additionalProperties, "ShowFromTime")
		delete(additionalProperties, "ShowUntilTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIwotenantMaintenanceNotificationAllOf struct {
	value *IwotenantMaintenanceNotificationAllOf
	isSet bool
}

func (v NullableIwotenantMaintenanceNotificationAllOf) Get() *IwotenantMaintenanceNotificationAllOf {
	return v.value
}

func (v *NullableIwotenantMaintenanceNotificationAllOf) Set(val *IwotenantMaintenanceNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIwotenantMaintenanceNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIwotenantMaintenanceNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwotenantMaintenanceNotificationAllOf(val *IwotenantMaintenanceNotificationAllOf) *NullableIwotenantMaintenanceNotificationAllOf {
	return &NullableIwotenantMaintenanceNotificationAllOf{value: val, isSet: true}
}

func (v NullableIwotenantMaintenanceNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwotenantMaintenanceNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
