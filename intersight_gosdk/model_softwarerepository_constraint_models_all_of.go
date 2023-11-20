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

// SoftwarerepositoryConstraintModelsAllOf Definition of the list of properties defined in 'softwarerepository.ConstraintModels', excluding properties defined in parent classes.
type SoftwarerepositoryConstraintModelsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Minimum version of the image.
	MinVersion *string `json:"MinVersion,omitempty"`
	// Name of the contraint, used to identify constriant type.
	Name *string `json:"Name,omitempty"`
	// Regular expression of the image name.
	PlatformRegex        *string  `json:"PlatformRegex,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryConstraintModelsAllOf SoftwarerepositoryConstraintModelsAllOf

// NewSoftwarerepositoryConstraintModelsAllOf instantiates a new SoftwarerepositoryConstraintModelsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryConstraintModelsAllOf(classId string, objectType string) *SoftwarerepositoryConstraintModelsAllOf {
	this := SoftwarerepositoryConstraintModelsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryConstraintModelsAllOfWithDefaults instantiates a new SoftwarerepositoryConstraintModelsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryConstraintModelsAllOfWithDefaults() *SoftwarerepositoryConstraintModelsAllOf {
	this := SoftwarerepositoryConstraintModelsAllOf{}
	var classId string = "softwarerepository.ConstraintModels"
	this.ClassId = classId
	var objectType string = "softwarerepository.ConstraintModels"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryConstraintModelsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryConstraintModelsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryConstraintModelsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryConstraintModelsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMinVersion returns the MinVersion field value if set, zero value otherwise.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetMinVersion() string {
	if o == nil || o.MinVersion == nil {
		var ret string
		return ret
	}
	return *o.MinVersion
}

// GetMinVersionOk returns a tuple with the MinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetMinVersionOk() (*string, bool) {
	if o == nil || o.MinVersion == nil {
		return nil, false
	}
	return o.MinVersion, true
}

// HasMinVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) HasMinVersion() bool {
	if o != nil && o.MinVersion != nil {
		return true
	}

	return false
}

// SetMinVersion gets a reference to the given string and assigns it to the MinVersion field.
func (o *SoftwarerepositoryConstraintModelsAllOf) SetMinVersion(v string) {
	o.MinVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwarerepositoryConstraintModelsAllOf) SetName(v string) {
	o.Name = &v
}

// GetPlatformRegex returns the PlatformRegex field value if set, zero value otherwise.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetPlatformRegex() string {
	if o == nil || o.PlatformRegex == nil {
		var ret string
		return ret
	}
	return *o.PlatformRegex
}

// GetPlatformRegexOk returns a tuple with the PlatformRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) GetPlatformRegexOk() (*string, bool) {
	if o == nil || o.PlatformRegex == nil {
		return nil, false
	}
	return o.PlatformRegex, true
}

// HasPlatformRegex returns a boolean if a field has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) HasPlatformRegex() bool {
	if o != nil && o.PlatformRegex != nil {
		return true
	}

	return false
}

// SetPlatformRegex gets a reference to the given string and assigns it to the PlatformRegex field.
func (o *SoftwarerepositoryConstraintModelsAllOf) SetPlatformRegex(v string) {
	o.PlatformRegex = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryConstraintModelsAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryConstraintModelsAllOf) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryConstraintModelsAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryConstraintModelsAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o SoftwarerepositoryConstraintModelsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MinVersion != nil {
		toSerialize["MinVersion"] = o.MinVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PlatformRegex != nil {
		toSerialize["PlatformRegex"] = o.PlatformRegex
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryConstraintModelsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryConstraintModelsAllOf := _SoftwarerepositoryConstraintModelsAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryConstraintModelsAllOf); err == nil {
		*o = SoftwarerepositoryConstraintModelsAllOf(varSoftwarerepositoryConstraintModelsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MinVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PlatformRegex")
		delete(additionalProperties, "SupportedModels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryConstraintModelsAllOf struct {
	value *SoftwarerepositoryConstraintModelsAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryConstraintModelsAllOf) Get() *SoftwarerepositoryConstraintModelsAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryConstraintModelsAllOf) Set(val *SoftwarerepositoryConstraintModelsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryConstraintModelsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryConstraintModelsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryConstraintModelsAllOf(val *SoftwarerepositoryConstraintModelsAllOf) *NullableSoftwarerepositoryConstraintModelsAllOf {
	return &NullableSoftwarerepositoryConstraintModelsAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryConstraintModelsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryConstraintModelsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
