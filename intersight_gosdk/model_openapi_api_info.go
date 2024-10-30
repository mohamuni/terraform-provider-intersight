/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the OpenapiApiInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenapiApiInfo{}

// OpenapiApiInfo Meta information about the API that will be generated as task.
type OpenapiApiInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType              string   `json:"ObjectType"`
	ApiPathObjectIdentifier *MoMoRef `json:"ApiPathObjectIdentifier,omitempty"`
	// A detailed description of the API.
	Description *string `json:"Description,omitempty"`
	// Display name of the selected API endpoint.
	DisplayLabel *string `json:"DisplayLabel,omitempty"`
	// Method Type of the selected API. * `GET` - Method type which indicates it is a GET API call. * `POST` - Method type which indicates it is a POST API call. * `PUT` - Method type which indicates it is a PUT API call. * `PATCH` - Method type which indicates it is a PATCH API call. * `DELETE` - Method type which indicates it is a DELETE API call.
	Method *string `json:"Method,omitempty"`
	// Name of the selected API endpoint.
	Name *string `json:"Name,omitempty"`
	// API Path of the selected API endpoint.
	Path *string `json:"Path,omitempty"`
	// Validation error messages will be captured by this property.
	ValidationError *string `json:"ValidationError,omitempty"`
	// Validation Status of selected API that indicates if the API validation passed or failed. * `none` - Indicates the default status. * `Success` - Denotes that the validation is successful. * `Failed` - Denotes that the validation failed. Validation could fail if information present in the OpenAPI specification is incomplete or missing.
	ValidationStatus     *string `json:"ValidationStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiApiInfo OpenapiApiInfo

// NewOpenapiApiInfo instantiates a new OpenapiApiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiApiInfo(classId string, objectType string) *OpenapiApiInfo {
	this := OpenapiApiInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var method string = "GET"
	this.Method = &method
	return &this
}

// NewOpenapiApiInfoWithDefaults instantiates a new OpenapiApiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiApiInfoWithDefaults() *OpenapiApiInfo {
	this := OpenapiApiInfo{}
	var classId string = "openapi.ApiInfo"
	this.ClassId = classId
	var objectType string = "openapi.ApiInfo"
	this.ObjectType = objectType
	var method string = "GET"
	this.Method = &method
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiApiInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiApiInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "openapi.ApiInfo" of the ClassId field.
func (o *OpenapiApiInfo) GetDefaultClassId() interface{} {
	return "openapi.ApiInfo"
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiApiInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiApiInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "openapi.ApiInfo" of the ObjectType field.
func (o *OpenapiApiInfo) GetDefaultObjectType() interface{} {
	return "openapi.ApiInfo"
}

// GetApiPathObjectIdentifier returns the ApiPathObjectIdentifier field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetApiPathObjectIdentifier() MoMoRef {
	if o == nil || IsNil(o.ApiPathObjectIdentifier) {
		var ret MoMoRef
		return ret
	}
	return *o.ApiPathObjectIdentifier
}

// GetApiPathObjectIdentifierOk returns a tuple with the ApiPathObjectIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetApiPathObjectIdentifierOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.ApiPathObjectIdentifier) {
		return nil, false
	}
	return o.ApiPathObjectIdentifier, true
}

// HasApiPathObjectIdentifier returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasApiPathObjectIdentifier() bool {
	if o != nil && !IsNil(o.ApiPathObjectIdentifier) {
		return true
	}

	return false
}

// SetApiPathObjectIdentifier gets a reference to the given MoMoRef and assigns it to the ApiPathObjectIdentifier field.
func (o *OpenapiApiInfo) SetApiPathObjectIdentifier(v MoMoRef) {
	o.ApiPathObjectIdentifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenapiApiInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetDisplayLabel() string {
	if o == nil || IsNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetDisplayLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasDisplayLabel() bool {
	if o != nil && !IsNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *OpenapiApiInfo) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *OpenapiApiInfo) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenapiApiInfo) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OpenapiApiInfo) SetPath(v string) {
	o.Path = &v
}

// GetValidationError returns the ValidationError field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetValidationError() string {
	if o == nil || IsNil(o.ValidationError) {
		var ret string
		return ret
	}
	return *o.ValidationError
}

// GetValidationErrorOk returns a tuple with the ValidationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetValidationErrorOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationError) {
		return nil, false
	}
	return o.ValidationError, true
}

// HasValidationError returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasValidationError() bool {
	if o != nil && !IsNil(o.ValidationError) {
		return true
	}

	return false
}

// SetValidationError gets a reference to the given string and assigns it to the ValidationError field.
func (o *OpenapiApiInfo) SetValidationError(v string) {
	o.ValidationError = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *OpenapiApiInfo) GetValidationStatus() string {
	if o == nil || IsNil(o.ValidationStatus) {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiInfo) GetValidationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationStatus) {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *OpenapiApiInfo) HasValidationStatus() bool {
	if o != nil && !IsNil(o.ValidationStatus) {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *OpenapiApiInfo) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

func (o OpenapiApiInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenapiApiInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ApiPathObjectIdentifier) {
		toSerialize["ApiPathObjectIdentifier"] = o.ApiPathObjectIdentifier
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DisplayLabel) {
		toSerialize["DisplayLabel"] = o.DisplayLabel
	}
	if !IsNil(o.Method) {
		toSerialize["Method"] = o.Method
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.ValidationError) {
		toSerialize["ValidationError"] = o.ValidationError
	}
	if !IsNil(o.ValidationStatus) {
		toSerialize["ValidationStatus"] = o.ValidationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenapiApiInfo) UnmarshalJSON(data []byte) (err error) {
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
	type OpenapiApiInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType              string   `json:"ObjectType"`
		ApiPathObjectIdentifier *MoMoRef `json:"ApiPathObjectIdentifier,omitempty"`
		// A detailed description of the API.
		Description *string `json:"Description,omitempty"`
		// Display name of the selected API endpoint.
		DisplayLabel *string `json:"DisplayLabel,omitempty"`
		// Method Type of the selected API. * `GET` - Method type which indicates it is a GET API call. * `POST` - Method type which indicates it is a POST API call. * `PUT` - Method type which indicates it is a PUT API call. * `PATCH` - Method type which indicates it is a PATCH API call. * `DELETE` - Method type which indicates it is a DELETE API call.
		Method *string `json:"Method,omitempty"`
		// Name of the selected API endpoint.
		Name *string `json:"Name,omitempty"`
		// API Path of the selected API endpoint.
		Path *string `json:"Path,omitempty"`
		// Validation error messages will be captured by this property.
		ValidationError *string `json:"ValidationError,omitempty"`
		// Validation Status of selected API that indicates if the API validation passed or failed. * `none` - Indicates the default status. * `Success` - Denotes that the validation is successful. * `Failed` - Denotes that the validation failed. Validation could fail if information present in the OpenAPI specification is incomplete or missing.
		ValidationStatus *string `json:"ValidationStatus,omitempty"`
	}

	varOpenapiApiInfoWithoutEmbeddedStruct := OpenapiApiInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOpenapiApiInfoWithoutEmbeddedStruct)
	if err == nil {
		varOpenapiApiInfo := _OpenapiApiInfo{}
		varOpenapiApiInfo.ClassId = varOpenapiApiInfoWithoutEmbeddedStruct.ClassId
		varOpenapiApiInfo.ObjectType = varOpenapiApiInfoWithoutEmbeddedStruct.ObjectType
		varOpenapiApiInfo.ApiPathObjectIdentifier = varOpenapiApiInfoWithoutEmbeddedStruct.ApiPathObjectIdentifier
		varOpenapiApiInfo.Description = varOpenapiApiInfoWithoutEmbeddedStruct.Description
		varOpenapiApiInfo.DisplayLabel = varOpenapiApiInfoWithoutEmbeddedStruct.DisplayLabel
		varOpenapiApiInfo.Method = varOpenapiApiInfoWithoutEmbeddedStruct.Method
		varOpenapiApiInfo.Name = varOpenapiApiInfoWithoutEmbeddedStruct.Name
		varOpenapiApiInfo.Path = varOpenapiApiInfoWithoutEmbeddedStruct.Path
		varOpenapiApiInfo.ValidationError = varOpenapiApiInfoWithoutEmbeddedStruct.ValidationError
		varOpenapiApiInfo.ValidationStatus = varOpenapiApiInfoWithoutEmbeddedStruct.ValidationStatus
		*o = OpenapiApiInfo(varOpenapiApiInfo)
	} else {
		return err
	}

	varOpenapiApiInfo := _OpenapiApiInfo{}

	err = json.Unmarshal(data, &varOpenapiApiInfo)
	if err == nil {
		o.MoBaseComplexType = varOpenapiApiInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiPathObjectIdentifier")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DisplayLabel")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "ValidationError")
		delete(additionalProperties, "ValidationStatus")

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

type NullableOpenapiApiInfo struct {
	value *OpenapiApiInfo
	isSet bool
}

func (v NullableOpenapiApiInfo) Get() *OpenapiApiInfo {
	return v.value
}

func (v *NullableOpenapiApiInfo) Set(val *OpenapiApiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiApiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiApiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiApiInfo(val *OpenapiApiInfo) *NullableOpenapiApiInfo {
	return &NullableOpenapiApiInfo{value: val, isSet: true}
}

func (v NullableOpenapiApiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiApiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
