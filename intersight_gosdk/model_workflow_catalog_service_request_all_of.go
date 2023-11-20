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

// WorkflowCatalogServiceRequestAllOf Definition of the list of properties defined in 'workflow.CatalogServiceRequest', excluding properties defined in parent classes.
type WorkflowCatalogServiceRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description for this catalog service request which provides information on request from the user.
	Description *string `json:"Description,omitempty"`
	// The time at which the service request completed or stopped.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Inputs for a catalog service request instance creation, format is specified by input definition of the catalog item definition.
	Input interface{} `json:"Input,omitempty"`
	// A user friendly short name to identify the catalog service request instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label    *string                 `json:"Label,omitempty"`
	Messages []ServicerequestMessage `json:"Messages,omitempty"`
	// A name of the catalog service request instance.
	Name                    *string                             `json:"Name,omitempty"`
	Operation               NullableWorkflowBaseOperation       `json:"Operation,omitempty"`
	SelectionCriteriaInputs []ServiceitemSelectionCriteriaInput `json:"SelectionCriteriaInputs,omitempty"`
	// Status of the catalog service request instance which determines the actions that are allowed on this instance. * `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * `Okay` - The last action on the service item instance completed and the service item instance is in Okay state. * `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned.
	Status *string `json:"Status,omitempty"`
	// The user identifier who invoked the request to create the catalog service request.
	UserIdOrEmail         *string                                    `json:"UserIdOrEmail,omitempty"`
	CatalogItemDefinition *WorkflowCatalogItemDefinitionRelationship `json:"CatalogItemDefinition,omitempty"`
	Idp                   *IamIdpRelationship                        `json:"Idp,omitempty"`
	IdpReference          *IamIdpReferenceRelationship               `json:"IdpReference,omitempty"`
	Organization          *OrganizationOrganizationRelationship      `json:"Organization,omitempty"`
	// An array of relationships to workflowServiceItemActionInstance resources.
	ServiceItemActionInstances []WorkflowServiceItemActionInstanceRelationship `json:"ServiceItemActionInstances,omitempty"`
	// An array of relationships to workflowServiceItemInstance resources.
	ServiceItemInstances []WorkflowServiceItemInstanceRelationship `json:"ServiceItemInstances,omitempty"`
	User                 *IamUserRelationship                      `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCatalogServiceRequestAllOf WorkflowCatalogServiceRequestAllOf

// NewWorkflowCatalogServiceRequestAllOf instantiates a new WorkflowCatalogServiceRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCatalogServiceRequestAllOf(classId string, objectType string) *WorkflowCatalogServiceRequestAllOf {
	this := WorkflowCatalogServiceRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCatalogServiceRequestAllOfWithDefaults instantiates a new WorkflowCatalogServiceRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCatalogServiceRequestAllOfWithDefaults() *WorkflowCatalogServiceRequestAllOf {
	this := WorkflowCatalogServiceRequestAllOf{}
	var classId string = "workflow.CatalogServiceRequest"
	this.ClassId = classId
	var objectType string = "workflow.CatalogServiceRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCatalogServiceRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCatalogServiceRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCatalogServiceRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCatalogServiceRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowCatalogServiceRequestAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *WorkflowCatalogServiceRequestAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogServiceRequestAllOf) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogServiceRequestAllOf) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowCatalogServiceRequestAllOf) SetInput(v interface{}) {
	o.Input = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowCatalogServiceRequestAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogServiceRequestAllOf) GetMessages() []ServicerequestMessage {
	if o == nil {
		var ret []ServicerequestMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogServiceRequestAllOf) GetMessagesOk() ([]ServicerequestMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ServicerequestMessage and assigns it to the Messages field.
func (o *WorkflowCatalogServiceRequestAllOf) SetMessages(v []ServicerequestMessage) {
	o.Messages = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowCatalogServiceRequestAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogServiceRequestAllOf) GetOperation() WorkflowBaseOperation {
	if o == nil || o.Operation.Get() == nil {
		var ret WorkflowBaseOperation
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogServiceRequestAllOf) GetOperationOk() (*WorkflowBaseOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableWorkflowBaseOperation and assigns it to the Operation field.
func (o *WorkflowCatalogServiceRequestAllOf) SetOperation(v WorkflowBaseOperation) {
	o.Operation.Set(&v)
}

// SetOperationNil sets the value for Operation to be an explicit nil
func (o *WorkflowCatalogServiceRequestAllOf) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *WorkflowCatalogServiceRequestAllOf) UnsetOperation() {
	o.Operation.Unset()
}

// GetSelectionCriteriaInputs returns the SelectionCriteriaInputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogServiceRequestAllOf) GetSelectionCriteriaInputs() []ServiceitemSelectionCriteriaInput {
	if o == nil {
		var ret []ServiceitemSelectionCriteriaInput
		return ret
	}
	return o.SelectionCriteriaInputs
}

// GetSelectionCriteriaInputsOk returns a tuple with the SelectionCriteriaInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogServiceRequestAllOf) GetSelectionCriteriaInputsOk() ([]ServiceitemSelectionCriteriaInput, bool) {
	if o == nil || o.SelectionCriteriaInputs == nil {
		return nil, false
	}
	return o.SelectionCriteriaInputs, true
}

// HasSelectionCriteriaInputs returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasSelectionCriteriaInputs() bool {
	if o != nil && o.SelectionCriteriaInputs != nil {
		return true
	}

	return false
}

// SetSelectionCriteriaInputs gets a reference to the given []ServiceitemSelectionCriteriaInput and assigns it to the SelectionCriteriaInputs field.
func (o *WorkflowCatalogServiceRequestAllOf) SetSelectionCriteriaInputs(v []ServiceitemSelectionCriteriaInput) {
	o.SelectionCriteriaInputs = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowCatalogServiceRequestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetUserIdOrEmail() string {
	if o == nil || o.UserIdOrEmail == nil {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || o.UserIdOrEmail == nil {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasUserIdOrEmail() bool {
	if o != nil && o.UserIdOrEmail != nil {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *WorkflowCatalogServiceRequestAllOf) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetCatalogItemDefinition returns the CatalogItemDefinition field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetCatalogItemDefinition() WorkflowCatalogItemDefinitionRelationship {
	if o == nil || o.CatalogItemDefinition == nil {
		var ret WorkflowCatalogItemDefinitionRelationship
		return ret
	}
	return *o.CatalogItemDefinition
}

// GetCatalogItemDefinitionOk returns a tuple with the CatalogItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetCatalogItemDefinitionOk() (*WorkflowCatalogItemDefinitionRelationship, bool) {
	if o == nil || o.CatalogItemDefinition == nil {
		return nil, false
	}
	return o.CatalogItemDefinition, true
}

// HasCatalogItemDefinition returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasCatalogItemDefinition() bool {
	if o != nil && o.CatalogItemDefinition != nil {
		return true
	}

	return false
}

// SetCatalogItemDefinition gets a reference to the given WorkflowCatalogItemDefinitionRelationship and assigns it to the CatalogItemDefinition field.
func (o *WorkflowCatalogServiceRequestAllOf) SetCatalogItemDefinition(v WorkflowCatalogItemDefinitionRelationship) {
	o.CatalogItemDefinition = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *WorkflowCatalogServiceRequestAllOf) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpReference returns the IdpReference field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetIdpReference() IamIdpReferenceRelationship {
	if o == nil || o.IdpReference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.IdpReference
}

// GetIdpReferenceOk returns a tuple with the IdpReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.IdpReference == nil {
		return nil, false
	}
	return o.IdpReference, true
}

// HasIdpReference returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasIdpReference() bool {
	if o != nil && o.IdpReference != nil {
		return true
	}

	return false
}

// SetIdpReference gets a reference to the given IamIdpReferenceRelationship and assigns it to the IdpReference field.
func (o *WorkflowCatalogServiceRequestAllOf) SetIdpReference(v IamIdpReferenceRelationship) {
	o.IdpReference = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *WorkflowCatalogServiceRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetServiceItemActionInstances returns the ServiceItemActionInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemActionInstances() []WorkflowServiceItemActionInstanceRelationship {
	if o == nil {
		var ret []WorkflowServiceItemActionInstanceRelationship
		return ret
	}
	return o.ServiceItemActionInstances
}

// GetServiceItemActionInstancesOk returns a tuple with the ServiceItemActionInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemActionInstancesOk() ([]WorkflowServiceItemActionInstanceRelationship, bool) {
	if o == nil || o.ServiceItemActionInstances == nil {
		return nil, false
	}
	return o.ServiceItemActionInstances, true
}

// HasServiceItemActionInstances returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasServiceItemActionInstances() bool {
	if o != nil && o.ServiceItemActionInstances != nil {
		return true
	}

	return false
}

// SetServiceItemActionInstances gets a reference to the given []WorkflowServiceItemActionInstanceRelationship and assigns it to the ServiceItemActionInstances field.
func (o *WorkflowCatalogServiceRequestAllOf) SetServiceItemActionInstances(v []WorkflowServiceItemActionInstanceRelationship) {
	o.ServiceItemActionInstances = v
}

// GetServiceItemInstances returns the ServiceItemInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemInstances() []WorkflowServiceItemInstanceRelationship {
	if o == nil {
		var ret []WorkflowServiceItemInstanceRelationship
		return ret
	}
	return o.ServiceItemInstances
}

// GetServiceItemInstancesOk returns a tuple with the ServiceItemInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemInstancesOk() ([]WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil || o.ServiceItemInstances == nil {
		return nil, false
	}
	return o.ServiceItemInstances, true
}

// HasServiceItemInstances returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasServiceItemInstances() bool {
	if o != nil && o.ServiceItemInstances != nil {
		return true
	}

	return false
}

// SetServiceItemInstances gets a reference to the given []WorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstances field.
func (o *WorkflowCatalogServiceRequestAllOf) SetServiceItemInstances(v []WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstances = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *WorkflowCatalogServiceRequestAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogServiceRequestAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *WorkflowCatalogServiceRequestAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *WorkflowCatalogServiceRequestAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o WorkflowCatalogServiceRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Operation.IsSet() {
		toSerialize["Operation"] = o.Operation.Get()
	}
	if o.SelectionCriteriaInputs != nil {
		toSerialize["SelectionCriteriaInputs"] = o.SelectionCriteriaInputs
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UserIdOrEmail != nil {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if o.CatalogItemDefinition != nil {
		toSerialize["CatalogItemDefinition"] = o.CatalogItemDefinition
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.IdpReference != nil {
		toSerialize["IdpReference"] = o.IdpReference
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ServiceItemActionInstances != nil {
		toSerialize["ServiceItemActionInstances"] = o.ServiceItemActionInstances
	}
	if o.ServiceItemInstances != nil {
		toSerialize["ServiceItemInstances"] = o.ServiceItemInstances
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCatalogServiceRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowCatalogServiceRequestAllOf := _WorkflowCatalogServiceRequestAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowCatalogServiceRequestAllOf); err == nil {
		*o = WorkflowCatalogServiceRequestAllOf(varWorkflowCatalogServiceRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Operation")
		delete(additionalProperties, "SelectionCriteriaInputs")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "CatalogItemDefinition")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "IdpReference")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "ServiceItemActionInstances")
		delete(additionalProperties, "ServiceItemInstances")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowCatalogServiceRequestAllOf struct {
	value *WorkflowCatalogServiceRequestAllOf
	isSet bool
}

func (v NullableWorkflowCatalogServiceRequestAllOf) Get() *WorkflowCatalogServiceRequestAllOf {
	return v.value
}

func (v *NullableWorkflowCatalogServiceRequestAllOf) Set(val *WorkflowCatalogServiceRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCatalogServiceRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCatalogServiceRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCatalogServiceRequestAllOf(val *WorkflowCatalogServiceRequestAllOf) *NullableWorkflowCatalogServiceRequestAllOf {
	return &NullableWorkflowCatalogServiceRequestAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCatalogServiceRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCatalogServiceRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
