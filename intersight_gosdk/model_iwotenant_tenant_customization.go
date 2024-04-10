/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IwotenantTenantCustomization Tenant customization that end user can do such as enabling data extractor.
type IwotenantTenantCustomization struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables IWO tenant data to be available for reporting.  This will start 'extractor' pod.
	EnableDataExtractor *bool `json:"EnableDataExtractor,omitempty"`
	// Indicates whether the value of the 'writeUserAccessKeyId' property has been set.
	IsWriteUserAccessKeyIdSet *bool `json:"IsWriteUserAccessKeyIdSet,omitempty"`
	// Indicates whether the value of the 'writeUserSecretAccessKey' property has been set.
	IsWriteUserSecretAccessKeySet *bool `json:"IsWriteUserSecretAccessKeySet,omitempty"`
	// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
	IwoId *string `json:"IwoId,omitempty"`
	// MSK cluster endpoint that data extractor can send reporting data to. This  MS cluster in turn populates data into tables in Redshift cluster.
	MskServerForDataExtractor *string `json:"MskServerForDataExtractor,omitempty"`
	// AWS access key Id to write data to redshift.  Refer to AWS cloud formation stack 'Output' of the tenant.
	WriteUserAccessKeyId *string `json:"WriteUserAccessKeyId,omitempty"`
	// AWS secret access key to write data to redshift.  Refer to AWS cloud formation stack 'Output' of the tenant.
	WriteUserSecretAccessKey *string                 `json:"WriteUserSecretAccessKey,omitempty"`
	Account                  *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _IwotenantTenantCustomization IwotenantTenantCustomization

// NewIwotenantTenantCustomization instantiates a new IwotenantTenantCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwotenantTenantCustomization(classId string, objectType string) *IwotenantTenantCustomization {
	this := IwotenantTenantCustomization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIwotenantTenantCustomizationWithDefaults instantiates a new IwotenantTenantCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwotenantTenantCustomizationWithDefaults() *IwotenantTenantCustomization {
	this := IwotenantTenantCustomization{}
	var classId string = "iwotenant.TenantCustomization"
	this.ClassId = classId
	var objectType string = "iwotenant.TenantCustomization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IwotenantTenantCustomization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IwotenantTenantCustomization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IwotenantTenantCustomization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IwotenantTenantCustomization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableDataExtractor returns the EnableDataExtractor field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetEnableDataExtractor() bool {
	if o == nil || o.EnableDataExtractor == nil {
		var ret bool
		return ret
	}
	return *o.EnableDataExtractor
}

// GetEnableDataExtractorOk returns a tuple with the EnableDataExtractor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetEnableDataExtractorOk() (*bool, bool) {
	if o == nil || o.EnableDataExtractor == nil {
		return nil, false
	}
	return o.EnableDataExtractor, true
}

// HasEnableDataExtractor returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasEnableDataExtractor() bool {
	if o != nil && o.EnableDataExtractor != nil {
		return true
	}

	return false
}

// SetEnableDataExtractor gets a reference to the given bool and assigns it to the EnableDataExtractor field.
func (o *IwotenantTenantCustomization) SetEnableDataExtractor(v bool) {
	o.EnableDataExtractor = &v
}

// GetIsWriteUserAccessKeyIdSet returns the IsWriteUserAccessKeyIdSet field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetIsWriteUserAccessKeyIdSet() bool {
	if o == nil || o.IsWriteUserAccessKeyIdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsWriteUserAccessKeyIdSet
}

// GetIsWriteUserAccessKeyIdSetOk returns a tuple with the IsWriteUserAccessKeyIdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetIsWriteUserAccessKeyIdSetOk() (*bool, bool) {
	if o == nil || o.IsWriteUserAccessKeyIdSet == nil {
		return nil, false
	}
	return o.IsWriteUserAccessKeyIdSet, true
}

// HasIsWriteUserAccessKeyIdSet returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasIsWriteUserAccessKeyIdSet() bool {
	if o != nil && o.IsWriteUserAccessKeyIdSet != nil {
		return true
	}

	return false
}

// SetIsWriteUserAccessKeyIdSet gets a reference to the given bool and assigns it to the IsWriteUserAccessKeyIdSet field.
func (o *IwotenantTenantCustomization) SetIsWriteUserAccessKeyIdSet(v bool) {
	o.IsWriteUserAccessKeyIdSet = &v
}

// GetIsWriteUserSecretAccessKeySet returns the IsWriteUserSecretAccessKeySet field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetIsWriteUserSecretAccessKeySet() bool {
	if o == nil || o.IsWriteUserSecretAccessKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsWriteUserSecretAccessKeySet
}

// GetIsWriteUserSecretAccessKeySetOk returns a tuple with the IsWriteUserSecretAccessKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetIsWriteUserSecretAccessKeySetOk() (*bool, bool) {
	if o == nil || o.IsWriteUserSecretAccessKeySet == nil {
		return nil, false
	}
	return o.IsWriteUserSecretAccessKeySet, true
}

// HasIsWriteUserSecretAccessKeySet returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasIsWriteUserSecretAccessKeySet() bool {
	if o != nil && o.IsWriteUserSecretAccessKeySet != nil {
		return true
	}

	return false
}

// SetIsWriteUserSecretAccessKeySet gets a reference to the given bool and assigns it to the IsWriteUserSecretAccessKeySet field.
func (o *IwotenantTenantCustomization) SetIsWriteUserSecretAccessKeySet(v bool) {
	o.IsWriteUserSecretAccessKeySet = &v
}

// GetIwoId returns the IwoId field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetIwoId() string {
	if o == nil || o.IwoId == nil {
		var ret string
		return ret
	}
	return *o.IwoId
}

// GetIwoIdOk returns a tuple with the IwoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetIwoIdOk() (*string, bool) {
	if o == nil || o.IwoId == nil {
		return nil, false
	}
	return o.IwoId, true
}

// HasIwoId returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasIwoId() bool {
	if o != nil && o.IwoId != nil {
		return true
	}

	return false
}

// SetIwoId gets a reference to the given string and assigns it to the IwoId field.
func (o *IwotenantTenantCustomization) SetIwoId(v string) {
	o.IwoId = &v
}

// GetMskServerForDataExtractor returns the MskServerForDataExtractor field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetMskServerForDataExtractor() string {
	if o == nil || o.MskServerForDataExtractor == nil {
		var ret string
		return ret
	}
	return *o.MskServerForDataExtractor
}

// GetMskServerForDataExtractorOk returns a tuple with the MskServerForDataExtractor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetMskServerForDataExtractorOk() (*string, bool) {
	if o == nil || o.MskServerForDataExtractor == nil {
		return nil, false
	}
	return o.MskServerForDataExtractor, true
}

// HasMskServerForDataExtractor returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasMskServerForDataExtractor() bool {
	if o != nil && o.MskServerForDataExtractor != nil {
		return true
	}

	return false
}

// SetMskServerForDataExtractor gets a reference to the given string and assigns it to the MskServerForDataExtractor field.
func (o *IwotenantTenantCustomization) SetMskServerForDataExtractor(v string) {
	o.MskServerForDataExtractor = &v
}

// GetWriteUserAccessKeyId returns the WriteUserAccessKeyId field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetWriteUserAccessKeyId() string {
	if o == nil || o.WriteUserAccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.WriteUserAccessKeyId
}

// GetWriteUserAccessKeyIdOk returns a tuple with the WriteUserAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetWriteUserAccessKeyIdOk() (*string, bool) {
	if o == nil || o.WriteUserAccessKeyId == nil {
		return nil, false
	}
	return o.WriteUserAccessKeyId, true
}

// HasWriteUserAccessKeyId returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasWriteUserAccessKeyId() bool {
	if o != nil && o.WriteUserAccessKeyId != nil {
		return true
	}

	return false
}

// SetWriteUserAccessKeyId gets a reference to the given string and assigns it to the WriteUserAccessKeyId field.
func (o *IwotenantTenantCustomization) SetWriteUserAccessKeyId(v string) {
	o.WriteUserAccessKeyId = &v
}

// GetWriteUserSecretAccessKey returns the WriteUserSecretAccessKey field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetWriteUserSecretAccessKey() string {
	if o == nil || o.WriteUserSecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.WriteUserSecretAccessKey
}

// GetWriteUserSecretAccessKeyOk returns a tuple with the WriteUserSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetWriteUserSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.WriteUserSecretAccessKey == nil {
		return nil, false
	}
	return o.WriteUserSecretAccessKey, true
}

// HasWriteUserSecretAccessKey returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasWriteUserSecretAccessKey() bool {
	if o != nil && o.WriteUserSecretAccessKey != nil {
		return true
	}

	return false
}

// SetWriteUserSecretAccessKey gets a reference to the given string and assigns it to the WriteUserSecretAccessKey field.
func (o *IwotenantTenantCustomization) SetWriteUserSecretAccessKey(v string) {
	o.WriteUserSecretAccessKey = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IwotenantTenantCustomization) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantCustomization) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IwotenantTenantCustomization) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IwotenantTenantCustomization) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IwotenantTenantCustomization) MarshalJSON() ([]byte, error) {
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
	if o.EnableDataExtractor != nil {
		toSerialize["EnableDataExtractor"] = o.EnableDataExtractor
	}
	if o.IsWriteUserAccessKeyIdSet != nil {
		toSerialize["IsWriteUserAccessKeyIdSet"] = o.IsWriteUserAccessKeyIdSet
	}
	if o.IsWriteUserSecretAccessKeySet != nil {
		toSerialize["IsWriteUserSecretAccessKeySet"] = o.IsWriteUserSecretAccessKeySet
	}
	if o.IwoId != nil {
		toSerialize["IwoId"] = o.IwoId
	}
	if o.MskServerForDataExtractor != nil {
		toSerialize["MskServerForDataExtractor"] = o.MskServerForDataExtractor
	}
	if o.WriteUserAccessKeyId != nil {
		toSerialize["WriteUserAccessKeyId"] = o.WriteUserAccessKeyId
	}
	if o.WriteUserSecretAccessKey != nil {
		toSerialize["WriteUserSecretAccessKey"] = o.WriteUserSecretAccessKey
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IwotenantTenantCustomization) UnmarshalJSON(bytes []byte) (err error) {
	type IwotenantTenantCustomizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables IWO tenant data to be available for reporting.  This will start 'extractor' pod.
		EnableDataExtractor *bool `json:"EnableDataExtractor,omitempty"`
		// Indicates whether the value of the 'writeUserAccessKeyId' property has been set.
		IsWriteUserAccessKeyIdSet *bool `json:"IsWriteUserAccessKeyIdSet,omitempty"`
		// Indicates whether the value of the 'writeUserSecretAccessKey' property has been set.
		IsWriteUserSecretAccessKeySet *bool `json:"IsWriteUserSecretAccessKeySet,omitempty"`
		// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
		IwoId *string `json:"IwoId,omitempty"`
		// MSK cluster endpoint that data extractor can send reporting data to. This  MS cluster in turn populates data into tables in Redshift cluster.
		MskServerForDataExtractor *string `json:"MskServerForDataExtractor,omitempty"`
		// AWS access key Id to write data to redshift.  Refer to AWS cloud formation stack 'Output' of the tenant.
		WriteUserAccessKeyId *string `json:"WriteUserAccessKeyId,omitempty"`
		// AWS secret access key to write data to redshift.  Refer to AWS cloud formation stack 'Output' of the tenant.
		WriteUserSecretAccessKey *string                 `json:"WriteUserSecretAccessKey,omitempty"`
		Account                  *IamAccountRelationship `json:"Account,omitempty"`
	}

	varIwotenantTenantCustomizationWithoutEmbeddedStruct := IwotenantTenantCustomizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIwotenantTenantCustomizationWithoutEmbeddedStruct)
	if err == nil {
		varIwotenantTenantCustomization := _IwotenantTenantCustomization{}
		varIwotenantTenantCustomization.ClassId = varIwotenantTenantCustomizationWithoutEmbeddedStruct.ClassId
		varIwotenantTenantCustomization.ObjectType = varIwotenantTenantCustomizationWithoutEmbeddedStruct.ObjectType
		varIwotenantTenantCustomization.EnableDataExtractor = varIwotenantTenantCustomizationWithoutEmbeddedStruct.EnableDataExtractor
		varIwotenantTenantCustomization.IsWriteUserAccessKeyIdSet = varIwotenantTenantCustomizationWithoutEmbeddedStruct.IsWriteUserAccessKeyIdSet
		varIwotenantTenantCustomization.IsWriteUserSecretAccessKeySet = varIwotenantTenantCustomizationWithoutEmbeddedStruct.IsWriteUserSecretAccessKeySet
		varIwotenantTenantCustomization.IwoId = varIwotenantTenantCustomizationWithoutEmbeddedStruct.IwoId
		varIwotenantTenantCustomization.MskServerForDataExtractor = varIwotenantTenantCustomizationWithoutEmbeddedStruct.MskServerForDataExtractor
		varIwotenantTenantCustomization.WriteUserAccessKeyId = varIwotenantTenantCustomizationWithoutEmbeddedStruct.WriteUserAccessKeyId
		varIwotenantTenantCustomization.WriteUserSecretAccessKey = varIwotenantTenantCustomizationWithoutEmbeddedStruct.WriteUserSecretAccessKey
		varIwotenantTenantCustomization.Account = varIwotenantTenantCustomizationWithoutEmbeddedStruct.Account
		*o = IwotenantTenantCustomization(varIwotenantTenantCustomization)
	} else {
		return err
	}

	varIwotenantTenantCustomization := _IwotenantTenantCustomization{}

	err = json.Unmarshal(bytes, &varIwotenantTenantCustomization)
	if err == nil {
		o.MoBaseMo = varIwotenantTenantCustomization.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableDataExtractor")
		delete(additionalProperties, "IsWriteUserAccessKeyIdSet")
		delete(additionalProperties, "IsWriteUserSecretAccessKeySet")
		delete(additionalProperties, "IwoId")
		delete(additionalProperties, "MskServerForDataExtractor")
		delete(additionalProperties, "WriteUserAccessKeyId")
		delete(additionalProperties, "WriteUserSecretAccessKey")
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

type NullableIwotenantTenantCustomization struct {
	value *IwotenantTenantCustomization
	isSet bool
}

func (v NullableIwotenantTenantCustomization) Get() *IwotenantTenantCustomization {
	return v.value
}

func (v *NullableIwotenantTenantCustomization) Set(val *IwotenantTenantCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableIwotenantTenantCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableIwotenantTenantCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwotenantTenantCustomization(val *IwotenantTenantCustomization) *NullableIwotenantTenantCustomization {
	return &NullableIwotenantTenantCustomization{value: val, isSet: true}
}

func (v NullableIwotenantTenantCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwotenantTenantCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}