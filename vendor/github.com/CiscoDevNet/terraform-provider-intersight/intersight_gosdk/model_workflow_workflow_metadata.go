/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowWorkflowMetadata Workflow metadata is a collection of properties that are common across all the versions of a workflow.
type WorkflowWorkflowMetadata struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description for this workflow.
	Description *string `json:"Description,omitempty"`
	// A user friendly short name to identify the workflow metadata. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// The name for this workflow metadata. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
	Name *string `json:"Name,omitempty"`
	// An array of relationships to iamRole resources.
	AssociatedRoles      []IamRoleRelationship        `json:"AssociatedRoles,omitempty"`
	Catalog              *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowMetadata WorkflowWorkflowMetadata

// NewWorkflowWorkflowMetadata instantiates a new WorkflowWorkflowMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowMetadata(classId string, objectType string) *WorkflowWorkflowMetadata {
	this := WorkflowWorkflowMetadata{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowWorkflowMetadataWithDefaults instantiates a new WorkflowWorkflowMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowMetadataWithDefaults() *WorkflowWorkflowMetadata {
	this := WorkflowWorkflowMetadata{}
	var classId string = "workflow.WorkflowMetadata"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowMetadata"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowMetadata) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMetadata) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowMetadata) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowMetadata) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowMetadata) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWorkflowMetadata) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWorkflowMetadata) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWorkflowMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowWorkflowMetadata) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMetadata) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowWorkflowMetadata) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowWorkflowMetadata) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWorkflowMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWorkflowMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWorkflowMetadata) SetName(v string) {
	o.Name = &v
}

// GetAssociatedRoles returns the AssociatedRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowMetadata) GetAssociatedRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.AssociatedRoles
}

// GetAssociatedRolesOk returns a tuple with the AssociatedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowMetadata) GetAssociatedRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.AssociatedRoles == nil {
		return nil, false
	}
	return &o.AssociatedRoles, true
}

// HasAssociatedRoles returns a boolean if a field has been set.
func (o *WorkflowWorkflowMetadata) HasAssociatedRoles() bool {
	if o != nil && o.AssociatedRoles != nil {
		return true
	}

	return false
}

// SetAssociatedRoles gets a reference to the given []IamRoleRelationship and assigns it to the AssociatedRoles field.
func (o *WorkflowWorkflowMetadata) SetAssociatedRoles(v []IamRoleRelationship) {
	o.AssociatedRoles = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowWorkflowMetadata) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMetadata) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowWorkflowMetadata) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowWorkflowMetadata) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

func (o WorkflowWorkflowMetadata) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.AssociatedRoles != nil {
		toSerialize["AssociatedRoles"] = o.AssociatedRoles
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowMetadata) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWorkflowMetadataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description for this workflow.
		Description *string `json:"Description,omitempty"`
		// A user friendly short name to identify the workflow metadata. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
		Label *string `json:"Label,omitempty"`
		// The name for this workflow metadata. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
		Name *string `json:"Name,omitempty"`
		// An array of relationships to iamRole resources.
		AssociatedRoles []IamRoleRelationship        `json:"AssociatedRoles,omitempty"`
		Catalog         *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	}

	varWorkflowWorkflowMetadataWithoutEmbeddedStruct := WorkflowWorkflowMetadataWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowMetadataWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWorkflowMetadata := _WorkflowWorkflowMetadata{}
		varWorkflowWorkflowMetadata.ClassId = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.ClassId
		varWorkflowWorkflowMetadata.ObjectType = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.ObjectType
		varWorkflowWorkflowMetadata.Description = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.Description
		varWorkflowWorkflowMetadata.Label = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.Label
		varWorkflowWorkflowMetadata.Name = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.Name
		varWorkflowWorkflowMetadata.AssociatedRoles = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.AssociatedRoles
		varWorkflowWorkflowMetadata.Catalog = varWorkflowWorkflowMetadataWithoutEmbeddedStruct.Catalog
		*o = WorkflowWorkflowMetadata(varWorkflowWorkflowMetadata)
	} else {
		return err
	}

	varWorkflowWorkflowMetadata := _WorkflowWorkflowMetadata{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowMetadata)
	if err == nil {
		o.MoBaseMo = varWorkflowWorkflowMetadata.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "AssociatedRoles")
		delete(additionalProperties, "Catalog")

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

type NullableWorkflowWorkflowMetadata struct {
	value *WorkflowWorkflowMetadata
	isSet bool
}

func (v NullableWorkflowWorkflowMetadata) Get() *WorkflowWorkflowMetadata {
	return v.value
}

func (v *NullableWorkflowWorkflowMetadata) Set(val *WorkflowWorkflowMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowMetadata(val *WorkflowWorkflowMetadata) *NullableWorkflowWorkflowMetadata {
	return &NullableWorkflowWorkflowMetadata{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
