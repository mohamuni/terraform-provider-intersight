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
	"time"
)

// BulkResult The Result API is a status-monitor resource used to show the processing status of any bulk MO API when the request HTTP 'prefer' header is set to 'respond-async' value.
type BulkResult struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The timestamp in UTC when the request processing completed.
	CompletionTime *time.Time `json:"CompletionTime,omitempty"`
	// The number of subrequests received in this request.
	NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
	// The timestamp in UTC when the request was received.
	RequestReceivedTime *time.Time `json:"RequestReceivedTime,omitempty"`
	// The processing status of the request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
	Status *string `json:"Status,omitempty"`
	// The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// The URI on which this async operation is being performed.
	Uri          *string                               `json:"Uri,omitempty"`
	MoDeepCloner *BulkMoDeepClonerRelationship         `json:"MoDeepCloner,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	Results              []BulkSubRequestObjRelationship   `json:"Results,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkResult BulkResult

// NewBulkResult instantiates a new BulkResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkResult(classId string, objectType string) *BulkResult {
	this := BulkResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkResultWithDefaults instantiates a new BulkResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkResultWithDefaults() *BulkResult {
	this := BulkResult{}
	var classId string = "bulk.Result"
	this.ClassId = classId
	var objectType string = "bulk.Result"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *BulkResult) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *BulkResult) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *BulkResult) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetNumSubRequests returns the NumSubRequests field value if set, zero value otherwise.
func (o *BulkResult) GetNumSubRequests() int64 {
	if o == nil || o.NumSubRequests == nil {
		var ret int64
		return ret
	}
	return *o.NumSubRequests
}

// GetNumSubRequestsOk returns a tuple with the NumSubRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetNumSubRequestsOk() (*int64, bool) {
	if o == nil || o.NumSubRequests == nil {
		return nil, false
	}
	return o.NumSubRequests, true
}

// HasNumSubRequests returns a boolean if a field has been set.
func (o *BulkResult) HasNumSubRequests() bool {
	if o != nil && o.NumSubRequests != nil {
		return true
	}

	return false
}

// SetNumSubRequests gets a reference to the given int64 and assigns it to the NumSubRequests field.
func (o *BulkResult) SetNumSubRequests(v int64) {
	o.NumSubRequests = &v
}

// GetRequestReceivedTime returns the RequestReceivedTime field value if set, zero value otherwise.
func (o *BulkResult) GetRequestReceivedTime() time.Time {
	if o == nil || o.RequestReceivedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestReceivedTime
}

// GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetRequestReceivedTimeOk() (*time.Time, bool) {
	if o == nil || o.RequestReceivedTime == nil {
		return nil, false
	}
	return o.RequestReceivedTime, true
}

// HasRequestReceivedTime returns a boolean if a field has been set.
func (o *BulkResult) HasRequestReceivedTime() bool {
	if o != nil && o.RequestReceivedTime != nil {
		return true
	}

	return false
}

// SetRequestReceivedTime gets a reference to the given time.Time and assigns it to the RequestReceivedTime field.
func (o *BulkResult) SetRequestReceivedTime(v time.Time) {
	o.RequestReceivedTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkResult) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkResult) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkResult) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkResult) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkResult) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkResult) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkResult) SetUri(v string) {
	o.Uri = &v
}

// GetMoDeepCloner returns the MoDeepCloner field value if set, zero value otherwise.
func (o *BulkResult) GetMoDeepCloner() BulkMoDeepClonerRelationship {
	if o == nil || o.MoDeepCloner == nil {
		var ret BulkMoDeepClonerRelationship
		return ret
	}
	return *o.MoDeepCloner
}

// GetMoDeepClonerOk returns a tuple with the MoDeepCloner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetMoDeepClonerOk() (*BulkMoDeepClonerRelationship, bool) {
	if o == nil || o.MoDeepCloner == nil {
		return nil, false
	}
	return o.MoDeepCloner, true
}

// HasMoDeepCloner returns a boolean if a field has been set.
func (o *BulkResult) HasMoDeepCloner() bool {
	if o != nil && o.MoDeepCloner != nil {
		return true
	}

	return false
}

// SetMoDeepCloner gets a reference to the given BulkMoDeepClonerRelationship and assigns it to the MoDeepCloner field.
func (o *BulkResult) SetMoDeepCloner(v BulkMoDeepClonerRelationship) {
	o.MoDeepCloner = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkResult) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkResult) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkResult) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetResults() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetResultsOk() ([]BulkSubRequestObjRelationship, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkResult) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the Results field.
func (o *BulkResult) SetResults(v []BulkSubRequestObjRelationship) {
	o.Results = v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *BulkResult) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *BulkResult) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *BulkResult) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o BulkResult) MarshalJSON() ([]byte, error) {
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
	if o.CompletionTime != nil {
		toSerialize["CompletionTime"] = o.CompletionTime
	}
	if o.NumSubRequests != nil {
		toSerialize["NumSubRequests"] = o.NumSubRequests
	}
	if o.RequestReceivedTime != nil {
		toSerialize["RequestReceivedTime"] = o.RequestReceivedTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}
	if o.MoDeepCloner != nil {
		toSerialize["MoDeepCloner"] = o.MoDeepCloner
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkResult) UnmarshalJSON(bytes []byte) (err error) {
	type BulkResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The timestamp in UTC when the request processing completed.
		CompletionTime *time.Time `json:"CompletionTime,omitempty"`
		// The number of subrequests received in this request.
		NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
		// The timestamp in UTC when the request was received.
		RequestReceivedTime *time.Time `json:"RequestReceivedTime,omitempty"`
		// The processing status of the request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
		Status *string `json:"Status,omitempty"`
		// The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// The URI on which this async operation is being performed.
		Uri          *string                               `json:"Uri,omitempty"`
		MoDeepCloner *BulkMoDeepClonerRelationship         `json:"MoDeepCloner,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to bulkSubRequestObj resources.
		Results      []BulkSubRequestObjRelationship   `json:"Results,omitempty"`
		WorkflowInfo *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	}

	varBulkResultWithoutEmbeddedStruct := BulkResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkResultWithoutEmbeddedStruct)
	if err == nil {
		varBulkResult := _BulkResult{}
		varBulkResult.ClassId = varBulkResultWithoutEmbeddedStruct.ClassId
		varBulkResult.ObjectType = varBulkResultWithoutEmbeddedStruct.ObjectType
		varBulkResult.CompletionTime = varBulkResultWithoutEmbeddedStruct.CompletionTime
		varBulkResult.NumSubRequests = varBulkResultWithoutEmbeddedStruct.NumSubRequests
		varBulkResult.RequestReceivedTime = varBulkResultWithoutEmbeddedStruct.RequestReceivedTime
		varBulkResult.Status = varBulkResultWithoutEmbeddedStruct.Status
		varBulkResult.StatusMessage = varBulkResultWithoutEmbeddedStruct.StatusMessage
		varBulkResult.Uri = varBulkResultWithoutEmbeddedStruct.Uri
		varBulkResult.MoDeepCloner = varBulkResultWithoutEmbeddedStruct.MoDeepCloner
		varBulkResult.Organization = varBulkResultWithoutEmbeddedStruct.Organization
		varBulkResult.Results = varBulkResultWithoutEmbeddedStruct.Results
		varBulkResult.WorkflowInfo = varBulkResultWithoutEmbeddedStruct.WorkflowInfo
		*o = BulkResult(varBulkResult)
	} else {
		return err
	}

	varBulkResult := _BulkResult{}

	err = json.Unmarshal(bytes, &varBulkResult)
	if err == nil {
		o.MoBaseMo = varBulkResult.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompletionTime")
		delete(additionalProperties, "NumSubRequests")
		delete(additionalProperties, "RequestReceivedTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "MoDeepCloner")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Results")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableBulkResult struct {
	value *BulkResult
	isSet bool
}

func (v NullableBulkResult) Get() *BulkResult {
	return v.value
}

func (v *NullableBulkResult) Set(val *BulkResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkResult(val *BulkResult) *NullableBulkResult {
	return &NullableBulkResult{value: val, isSet: true}
}

func (v NullableBulkResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
