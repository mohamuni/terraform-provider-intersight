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

// NotificationTriggerWebhookAllOf Definition of the list of properties defined in 'notification.TriggerWebhook', excluding properties defined in parent classes.
type NotificationTriggerWebhookAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Holds the timestamp of the first failed request. The first time the notification is not delivered to the webhook server, the user will have the Warning alarm in the system. Next 48 hours the system still will try to notify the webhook server. If after 48 hours the server is not recovered, the system will mark this webhook as Inactive, and the user will have a critical alarm in the system.
	FirstFailedRequest *time.Time `json:"FirstFailedRequest,omitempty"`
	// Indicates whether the value of the 'secret' property has been set.
	IsSecretSet *bool `json:"IsSecretSet,omitempty"`
	// Holds the error message for the user of the last response.
	LastNetworkError *string `json:"LastNetworkError,omitempty"`
	// Holds the code of the last response, which helps to debug the issue in case if webhook server is not reachable.
	LastResponseCode *int64 `json:"LastResponseCode,omitempty"`
	// The secret is used to build the Authorization header, which will be attached to each webhook notification. By this header developers of the webhooks servers can make sure that events are received from the trusted source - Intersight.
	Secret *string `json:"Secret,omitempty"`
	// State of the action shows whether this action passes the verification or not. If this property holds 'Inactive' value, this action will not be executed. To verify action again, use the Verify property from the MO. * `Inactive` - Inactive state means action didn't pass the verification and it won't be executed. * `Active` - Active state means that action successfully passed the verification and it is ready to be performed.
	State *string `json:"State,omitempty"`
	// Payload URL of the recipient app, which is intended to serve the events that happens in Intersight.
	Url                  *string `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationTriggerWebhookAllOf NotificationTriggerWebhookAllOf

// NewNotificationTriggerWebhookAllOf instantiates a new NotificationTriggerWebhookAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTriggerWebhookAllOf(classId string, objectType string) *NotificationTriggerWebhookAllOf {
	this := NotificationTriggerWebhookAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationTriggerWebhookAllOfWithDefaults instantiates a new NotificationTriggerWebhookAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTriggerWebhookAllOfWithDefaults() *NotificationTriggerWebhookAllOf {
	this := NotificationTriggerWebhookAllOf{}
	var classId string = "notification.TriggerWebhook"
	this.ClassId = classId
	var objectType string = "notification.TriggerWebhook"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationTriggerWebhookAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationTriggerWebhookAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationTriggerWebhookAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationTriggerWebhookAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFirstFailedRequest returns the FirstFailedRequest field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetFirstFailedRequest() time.Time {
	if o == nil || o.FirstFailedRequest == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstFailedRequest
}

// GetFirstFailedRequestOk returns a tuple with the FirstFailedRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetFirstFailedRequestOk() (*time.Time, bool) {
	if o == nil || o.FirstFailedRequest == nil {
		return nil, false
	}
	return o.FirstFailedRequest, true
}

// HasFirstFailedRequest returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasFirstFailedRequest() bool {
	if o != nil && o.FirstFailedRequest != nil {
		return true
	}

	return false
}

// SetFirstFailedRequest gets a reference to the given time.Time and assigns it to the FirstFailedRequest field.
func (o *NotificationTriggerWebhookAllOf) SetFirstFailedRequest(v time.Time) {
	o.FirstFailedRequest = &v
}

// GetIsSecretSet returns the IsSecretSet field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetIsSecretSet() bool {
	if o == nil || o.IsSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.IsSecretSet
}

// GetIsSecretSetOk returns a tuple with the IsSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetIsSecretSetOk() (*bool, bool) {
	if o == nil || o.IsSecretSet == nil {
		return nil, false
	}
	return o.IsSecretSet, true
}

// HasIsSecretSet returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasIsSecretSet() bool {
	if o != nil && o.IsSecretSet != nil {
		return true
	}

	return false
}

// SetIsSecretSet gets a reference to the given bool and assigns it to the IsSecretSet field.
func (o *NotificationTriggerWebhookAllOf) SetIsSecretSet(v bool) {
	o.IsSecretSet = &v
}

// GetLastNetworkError returns the LastNetworkError field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetLastNetworkError() string {
	if o == nil || o.LastNetworkError == nil {
		var ret string
		return ret
	}
	return *o.LastNetworkError
}

// GetLastNetworkErrorOk returns a tuple with the LastNetworkError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetLastNetworkErrorOk() (*string, bool) {
	if o == nil || o.LastNetworkError == nil {
		return nil, false
	}
	return o.LastNetworkError, true
}

// HasLastNetworkError returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasLastNetworkError() bool {
	if o != nil && o.LastNetworkError != nil {
		return true
	}

	return false
}

// SetLastNetworkError gets a reference to the given string and assigns it to the LastNetworkError field.
func (o *NotificationTriggerWebhookAllOf) SetLastNetworkError(v string) {
	o.LastNetworkError = &v
}

// GetLastResponseCode returns the LastResponseCode field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetLastResponseCode() int64 {
	if o == nil || o.LastResponseCode == nil {
		var ret int64
		return ret
	}
	return *o.LastResponseCode
}

// GetLastResponseCodeOk returns a tuple with the LastResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetLastResponseCodeOk() (*int64, bool) {
	if o == nil || o.LastResponseCode == nil {
		return nil, false
	}
	return o.LastResponseCode, true
}

// HasLastResponseCode returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasLastResponseCode() bool {
	if o != nil && o.LastResponseCode != nil {
		return true
	}

	return false
}

// SetLastResponseCode gets a reference to the given int64 and assigns it to the LastResponseCode field.
func (o *NotificationTriggerWebhookAllOf) SetLastResponseCode(v int64) {
	o.LastResponseCode = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NotificationTriggerWebhookAllOf) SetSecret(v string) {
	o.Secret = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NotificationTriggerWebhookAllOf) SetState(v string) {
	o.State = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NotificationTriggerWebhookAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWebhookAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationTriggerWebhookAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NotificationTriggerWebhookAllOf) SetUrl(v string) {
	o.Url = &v
}

func (o NotificationTriggerWebhookAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FirstFailedRequest != nil {
		toSerialize["FirstFailedRequest"] = o.FirstFailedRequest
	}
	if o.IsSecretSet != nil {
		toSerialize["IsSecretSet"] = o.IsSecretSet
	}
	if o.LastNetworkError != nil {
		toSerialize["LastNetworkError"] = o.LastNetworkError
	}
	if o.LastResponseCode != nil {
		toSerialize["LastResponseCode"] = o.LastResponseCode
	}
	if o.Secret != nil {
		toSerialize["Secret"] = o.Secret
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationTriggerWebhookAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationTriggerWebhookAllOf := _NotificationTriggerWebhookAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationTriggerWebhookAllOf); err == nil {
		*o = NotificationTriggerWebhookAllOf(varNotificationTriggerWebhookAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FirstFailedRequest")
		delete(additionalProperties, "IsSecretSet")
		delete(additionalProperties, "LastNetworkError")
		delete(additionalProperties, "LastResponseCode")
		delete(additionalProperties, "Secret")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationTriggerWebhookAllOf struct {
	value *NotificationTriggerWebhookAllOf
	isSet bool
}

func (v NullableNotificationTriggerWebhookAllOf) Get() *NotificationTriggerWebhookAllOf {
	return v.value
}

func (v *NullableNotificationTriggerWebhookAllOf) Set(val *NotificationTriggerWebhookAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTriggerWebhookAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTriggerWebhookAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTriggerWebhookAllOf(val *NotificationTriggerWebhookAllOf) *NullableNotificationTriggerWebhookAllOf {
	return &NullableNotificationTriggerWebhookAllOf{value: val, isSet: true}
}

func (v NullableNotificationTriggerWebhookAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTriggerWebhookAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
