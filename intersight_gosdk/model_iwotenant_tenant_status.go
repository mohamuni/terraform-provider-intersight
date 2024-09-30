/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// checks if the IwotenantTenantStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IwotenantTenantStatus{}

// IwotenantTenantStatus When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.
type IwotenantTenantStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The deployStatus provides the current status of this deployment. * `NotStarted` - The workflow to deploy the tenant cluster has not yet started. * `InProgress` - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running. * `Completed` - The workflow to deploy the tenant cluster has completed and health checks have passed. * `Failed` - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus.
	DeployStatus *string `json:"DeployStatus,omitempty"`
	// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
	IwoId *string `json:"IwoId,omitempty"`
	// During IWO tenant upgrade (or reconfiguration), deployStatus is set to InProgress and referenceTime set to current time. When tenant upgrade (or reconfiguration) does not complete within a pre-defined time using this as reference, deployStatus is set as Failed.
	ReferenceTime        *time.Time                     `json:"ReferenceTime,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IwotenantTenantStatus IwotenantTenantStatus

// NewIwotenantTenantStatus instantiates a new IwotenantTenantStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwotenantTenantStatus(classId string, objectType string) *IwotenantTenantStatus {
	this := IwotenantTenantStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIwotenantTenantStatusWithDefaults instantiates a new IwotenantTenantStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwotenantTenantStatusWithDefaults() *IwotenantTenantStatus {
	this := IwotenantTenantStatus{}
	var classId string = "iwotenant.TenantStatus"
	this.ClassId = classId
	var objectType string = "iwotenant.TenantStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IwotenantTenantStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IwotenantTenantStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iwotenant.TenantStatus" of the ClassId field.
func (o *IwotenantTenantStatus) GetDefaultClassId() interface{} {
	return "iwotenant.TenantStatus"
}

// GetObjectType returns the ObjectType field value
func (o *IwotenantTenantStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IwotenantTenantStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iwotenant.TenantStatus" of the ObjectType field.
func (o *IwotenantTenantStatus) GetDefaultObjectType() interface{} {
	return "iwotenant.TenantStatus"
}

// GetDeployStatus returns the DeployStatus field value if set, zero value otherwise.
func (o *IwotenantTenantStatus) GetDeployStatus() string {
	if o == nil || IsNil(o.DeployStatus) {
		var ret string
		return ret
	}
	return *o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatus) GetDeployStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DeployStatus) {
		return nil, false
	}
	return o.DeployStatus, true
}

// HasDeployStatus returns a boolean if a field has been set.
func (o *IwotenantTenantStatus) HasDeployStatus() bool {
	if o != nil && !IsNil(o.DeployStatus) {
		return true
	}

	return false
}

// SetDeployStatus gets a reference to the given string and assigns it to the DeployStatus field.
func (o *IwotenantTenantStatus) SetDeployStatus(v string) {
	o.DeployStatus = &v
}

// GetIwoId returns the IwoId field value if set, zero value otherwise.
func (o *IwotenantTenantStatus) GetIwoId() string {
	if o == nil || IsNil(o.IwoId) {
		var ret string
		return ret
	}
	return *o.IwoId
}

// GetIwoIdOk returns a tuple with the IwoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatus) GetIwoIdOk() (*string, bool) {
	if o == nil || IsNil(o.IwoId) {
		return nil, false
	}
	return o.IwoId, true
}

// HasIwoId returns a boolean if a field has been set.
func (o *IwotenantTenantStatus) HasIwoId() bool {
	if o != nil && !IsNil(o.IwoId) {
		return true
	}

	return false
}

// SetIwoId gets a reference to the given string and assigns it to the IwoId field.
func (o *IwotenantTenantStatus) SetIwoId(v string) {
	o.IwoId = &v
}

// GetReferenceTime returns the ReferenceTime field value if set, zero value otherwise.
func (o *IwotenantTenantStatus) GetReferenceTime() time.Time {
	if o == nil || IsNil(o.ReferenceTime) {
		var ret time.Time
		return ret
	}
	return *o.ReferenceTime
}

// GetReferenceTimeOk returns a tuple with the ReferenceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatus) GetReferenceTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReferenceTime) {
		return nil, false
	}
	return o.ReferenceTime, true
}

// HasReferenceTime returns a boolean if a field has been set.
func (o *IwotenantTenantStatus) HasReferenceTime() bool {
	if o != nil && !IsNil(o.ReferenceTime) {
		return true
	}

	return false
}

// SetReferenceTime gets a reference to the given time.Time and assigns it to the ReferenceTime field.
func (o *IwotenantTenantStatus) SetReferenceTime(v time.Time) {
	o.ReferenceTime = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IwotenantTenantStatus) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IwotenantTenantStatus) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IwotenantTenantStatus) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IwotenantTenantStatus) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IwotenantTenantStatus) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IwotenantTenantStatus) UnsetAccount() {
	o.Account.Unset()
}

func (o IwotenantTenantStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IwotenantTenantStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DeployStatus) {
		toSerialize["DeployStatus"] = o.DeployStatus
	}
	if !IsNil(o.IwoId) {
		toSerialize["IwoId"] = o.IwoId
	}
	if !IsNil(o.ReferenceTime) {
		toSerialize["ReferenceTime"] = o.ReferenceTime
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IwotenantTenantStatus) UnmarshalJSON(data []byte) (err error) {
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
	type IwotenantTenantStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The deployStatus provides the current status of this deployment. * `NotStarted` - The workflow to deploy the tenant cluster has not yet started. * `InProgress` - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running. * `Completed` - The workflow to deploy the tenant cluster has completed and health checks have passed. * `Failed` - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus.
		DeployStatus *string `json:"DeployStatus,omitempty"`
		// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
		IwoId *string `json:"IwoId,omitempty"`
		// During IWO tenant upgrade (or reconfiguration), deployStatus is set to InProgress and referenceTime set to current time. When tenant upgrade (or reconfiguration) does not complete within a pre-defined time using this as reference, deployStatus is set as Failed.
		ReferenceTime *time.Time                     `json:"ReferenceTime,omitempty"`
		Account       NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varIwotenantTenantStatusWithoutEmbeddedStruct := IwotenantTenantStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIwotenantTenantStatusWithoutEmbeddedStruct)
	if err == nil {
		varIwotenantTenantStatus := _IwotenantTenantStatus{}
		varIwotenantTenantStatus.ClassId = varIwotenantTenantStatusWithoutEmbeddedStruct.ClassId
		varIwotenantTenantStatus.ObjectType = varIwotenantTenantStatusWithoutEmbeddedStruct.ObjectType
		varIwotenantTenantStatus.DeployStatus = varIwotenantTenantStatusWithoutEmbeddedStruct.DeployStatus
		varIwotenantTenantStatus.IwoId = varIwotenantTenantStatusWithoutEmbeddedStruct.IwoId
		varIwotenantTenantStatus.ReferenceTime = varIwotenantTenantStatusWithoutEmbeddedStruct.ReferenceTime
		varIwotenantTenantStatus.Account = varIwotenantTenantStatusWithoutEmbeddedStruct.Account
		*o = IwotenantTenantStatus(varIwotenantTenantStatus)
	} else {
		return err
	}

	varIwotenantTenantStatus := _IwotenantTenantStatus{}

	err = json.Unmarshal(data, &varIwotenantTenantStatus)
	if err == nil {
		o.MoBaseMo = varIwotenantTenantStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeployStatus")
		delete(additionalProperties, "IwoId")
		delete(additionalProperties, "ReferenceTime")
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

type NullableIwotenantTenantStatus struct {
	value *IwotenantTenantStatus
	isSet bool
}

func (v NullableIwotenantTenantStatus) Get() *IwotenantTenantStatus {
	return v.value
}

func (v *NullableIwotenantTenantStatus) Set(val *IwotenantTenantStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIwotenantTenantStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIwotenantTenantStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwotenantTenantStatus(val *IwotenantTenantStatus) *NullableIwotenantTenantStatus {
	return &NullableIwotenantTenantStatus{value: val, isSet: true}
}

func (v NullableIwotenantTenantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwotenantTenantStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
