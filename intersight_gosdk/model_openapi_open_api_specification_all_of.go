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

// OpenapiOpenApiSpecificationAllOf Definition of the list of properties defined in 'openapi.OpenApiSpecification', excluding properties defined in parent classes.
type OpenapiOpenApiSpecificationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The path of the file in S3/minio bucket.
	FilePath *string `json:"FilePath,omitempty"`
	// A unique tracking ID for the file being uploaded.
	FileUploadId         *string                               `json:"FileUploadId,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiOpenApiSpecificationAllOf OpenapiOpenApiSpecificationAllOf

// NewOpenapiOpenApiSpecificationAllOf instantiates a new OpenapiOpenApiSpecificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiOpenApiSpecificationAllOf(classId string, objectType string) *OpenapiOpenApiSpecificationAllOf {
	this := OpenapiOpenApiSpecificationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOpenapiOpenApiSpecificationAllOfWithDefaults instantiates a new OpenapiOpenApiSpecificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiOpenApiSpecificationAllOfWithDefaults() *OpenapiOpenApiSpecificationAllOf {
	this := OpenapiOpenApiSpecificationAllOf{}
	var classId string = "openapi.OpenApiSpecification"
	this.ClassId = classId
	var objectType string = "openapi.OpenApiSpecification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiOpenApiSpecificationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecificationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiOpenApiSpecificationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiOpenApiSpecificationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecificationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiOpenApiSpecificationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *OpenapiOpenApiSpecificationAllOf) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecificationAllOf) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *OpenapiOpenApiSpecificationAllOf) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *OpenapiOpenApiSpecificationAllOf) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileUploadId returns the FileUploadId field value if set, zero value otherwise.
func (o *OpenapiOpenApiSpecificationAllOf) GetFileUploadId() string {
	if o == nil || o.FileUploadId == nil {
		var ret string
		return ret
	}
	return *o.FileUploadId
}

// GetFileUploadIdOk returns a tuple with the FileUploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecificationAllOf) GetFileUploadIdOk() (*string, bool) {
	if o == nil || o.FileUploadId == nil {
		return nil, false
	}
	return o.FileUploadId, true
}

// HasFileUploadId returns a boolean if a field has been set.
func (o *OpenapiOpenApiSpecificationAllOf) HasFileUploadId() bool {
	if o != nil && o.FileUploadId != nil {
		return true
	}

	return false
}

// SetFileUploadId gets a reference to the given string and assigns it to the FileUploadId field.
func (o *OpenapiOpenApiSpecificationAllOf) SetFileUploadId(v string) {
	o.FileUploadId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OpenapiOpenApiSpecificationAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecificationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OpenapiOpenApiSpecificationAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OpenapiOpenApiSpecificationAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o OpenapiOpenApiSpecificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FileUploadId != nil {
		toSerialize["FileUploadId"] = o.FileUploadId
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenapiOpenApiSpecificationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOpenapiOpenApiSpecificationAllOf := _OpenapiOpenApiSpecificationAllOf{}

	if err = json.Unmarshal(bytes, &varOpenapiOpenApiSpecificationAllOf); err == nil {
		*o = OpenapiOpenApiSpecificationAllOf(varOpenapiOpenApiSpecificationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FileUploadId")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenapiOpenApiSpecificationAllOf struct {
	value *OpenapiOpenApiSpecificationAllOf
	isSet bool
}

func (v NullableOpenapiOpenApiSpecificationAllOf) Get() *OpenapiOpenApiSpecificationAllOf {
	return v.value
}

func (v *NullableOpenapiOpenApiSpecificationAllOf) Set(val *OpenapiOpenApiSpecificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiOpenApiSpecificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiOpenApiSpecificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiOpenApiSpecificationAllOf(val *OpenapiOpenApiSpecificationAllOf) *NullableOpenapiOpenApiSpecificationAllOf {
	return &NullableOpenapiOpenApiSpecificationAllOf{value: val, isSet: true}
}

func (v NullableOpenapiOpenApiSpecificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiOpenApiSpecificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
