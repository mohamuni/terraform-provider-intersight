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

// checks if the PartnerintegrationFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerintegrationFile{}

// PartnerintegrationFile A partner integration artifact which will be files containing development code uploaded by our partners to get the build for dc/microservice.
type PartnerintegrationFile struct {
	SoftwarerepositoryFile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Path of the file being uploaded.
	FilePath *string `json:"FilePath,omitempty"`
	// Type of the file being uploaded. * `None` - Invalid file type for partnerIntegration appliance. * `Model` - Model file of Generic Device. * `Etl` - ETL file of Generic Device. * `Ui` - UI file of Generic Device. * `DeviceConnector` - Generic Device Connector file.
	FileType *string `json:"FileType,omitempty"`
	// The partner integration workspace to use to upload the File.
	WorkspaceName        *string                                       `json:"WorkspaceName,omitempty"`
	Catalog              NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationFile PartnerintegrationFile

// NewPartnerintegrationFile instantiates a new PartnerintegrationFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationFile(classId string, objectType string) *PartnerintegrationFile {
	this := PartnerintegrationFile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var fileType string = "None"
	this.FileType = &fileType
	return &this
}

// NewPartnerintegrationFileWithDefaults instantiates a new PartnerintegrationFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationFileWithDefaults() *PartnerintegrationFile {
	this := PartnerintegrationFile{}
	var classId string = "partnerintegration.File"
	this.ClassId = classId
	var objectType string = "partnerintegration.File"
	this.ObjectType = objectType
	var fileType string = "None"
	this.FileType = &fileType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationFile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationFile) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "partnerintegration.File" of the ClassId field.
func (o *PartnerintegrationFile) GetDefaultClassId() interface{} {
	return "partnerintegration.File"
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationFile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationFile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "partnerintegration.File" of the ObjectType field.
func (o *PartnerintegrationFile) GetDefaultObjectType() interface{} {
	return "partnerintegration.File"
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *PartnerintegrationFile) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFile) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *PartnerintegrationFile) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *PartnerintegrationFile) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *PartnerintegrationFile) GetFileType() string {
	if o == nil || IsNil(o.FileType) {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFile) GetFileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *PartnerintegrationFile) HasFileType() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *PartnerintegrationFile) SetFileType(v string) {
	o.FileType = &v
}

// GetWorkspaceName returns the WorkspaceName field value if set, zero value otherwise.
func (o *PartnerintegrationFile) GetWorkspaceName() string {
	if o == nil || IsNil(o.WorkspaceName) {
		var ret string
		return ret
	}
	return *o.WorkspaceName
}

// GetWorkspaceNameOk returns a tuple with the WorkspaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationFile) GetWorkspaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.WorkspaceName) {
		return nil, false
	}
	return o.WorkspaceName, true
}

// HasWorkspaceName returns a boolean if a field has been set.
func (o *PartnerintegrationFile) HasWorkspaceName() bool {
	if o != nil && !IsNil(o.WorkspaceName) {
		return true
	}

	return false
}

// SetWorkspaceName gets a reference to the given string and assigns it to the WorkspaceName field.
func (o *PartnerintegrationFile) SetWorkspaceName(v string) {
	o.WorkspaceName = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationFile) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationFile) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *PartnerintegrationFile) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableSoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *PartnerintegrationFile) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *PartnerintegrationFile) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *PartnerintegrationFile) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o PartnerintegrationFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerintegrationFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFile, errSoftwarerepositoryFile := json.Marshal(o.SoftwarerepositoryFile)
	if errSoftwarerepositoryFile != nil {
		return map[string]interface{}{}, errSoftwarerepositoryFile
	}
	errSoftwarerepositoryFile = json.Unmarshal([]byte(serializedSoftwarerepositoryFile), &toSerialize)
	if errSoftwarerepositoryFile != nil {
		return map[string]interface{}{}, errSoftwarerepositoryFile
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FilePath) {
		toSerialize["FilePath"] = o.FilePath
	}
	if !IsNil(o.FileType) {
		toSerialize["FileType"] = o.FileType
	}
	if !IsNil(o.WorkspaceName) {
		toSerialize["WorkspaceName"] = o.WorkspaceName
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PartnerintegrationFile) UnmarshalJSON(data []byte) (err error) {
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
	type PartnerintegrationFileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Path of the file being uploaded.
		FilePath *string `json:"FilePath,omitempty"`
		// Type of the file being uploaded. * `None` - Invalid file type for partnerIntegration appliance. * `Model` - Model file of Generic Device. * `Etl` - ETL file of Generic Device. * `Ui` - UI file of Generic Device. * `DeviceConnector` - Generic Device Connector file.
		FileType *string `json:"FileType,omitempty"`
		// The partner integration workspace to use to upload the File.
		WorkspaceName *string                                       `json:"WorkspaceName,omitempty"`
		Catalog       NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varPartnerintegrationFileWithoutEmbeddedStruct := PartnerintegrationFileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPartnerintegrationFileWithoutEmbeddedStruct)
	if err == nil {
		varPartnerintegrationFile := _PartnerintegrationFile{}
		varPartnerintegrationFile.ClassId = varPartnerintegrationFileWithoutEmbeddedStruct.ClassId
		varPartnerintegrationFile.ObjectType = varPartnerintegrationFileWithoutEmbeddedStruct.ObjectType
		varPartnerintegrationFile.FilePath = varPartnerintegrationFileWithoutEmbeddedStruct.FilePath
		varPartnerintegrationFile.FileType = varPartnerintegrationFileWithoutEmbeddedStruct.FileType
		varPartnerintegrationFile.WorkspaceName = varPartnerintegrationFileWithoutEmbeddedStruct.WorkspaceName
		varPartnerintegrationFile.Catalog = varPartnerintegrationFileWithoutEmbeddedStruct.Catalog
		*o = PartnerintegrationFile(varPartnerintegrationFile)
	} else {
		return err
	}

	varPartnerintegrationFile := _PartnerintegrationFile{}

	err = json.Unmarshal(data, &varPartnerintegrationFile)
	if err == nil {
		o.SoftwarerepositoryFile = varPartnerintegrationFile.SoftwarerepositoryFile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "WorkspaceName")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectSoftwarerepositoryFile := reflect.ValueOf(o.SoftwarerepositoryFile)
		for i := 0; i < reflectSoftwarerepositoryFile.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFile.Type().Field(i)

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

type NullablePartnerintegrationFile struct {
	value *PartnerintegrationFile
	isSet bool
}

func (v NullablePartnerintegrationFile) Get() *PartnerintegrationFile {
	return v.value
}

func (v *NullablePartnerintegrationFile) Set(val *PartnerintegrationFile) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationFile) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationFile(val *PartnerintegrationFile) *NullablePartnerintegrationFile {
	return &NullablePartnerintegrationFile{value: val, isSet: true}
}

func (v NullablePartnerintegrationFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
