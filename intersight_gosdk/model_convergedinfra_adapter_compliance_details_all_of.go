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

// ConvergedinfraAdapterComplianceDetailsAllOf Definition of the list of properties defined in 'convergedinfra.AdapterComplianceDetails', excluding properties defined in parent classes.
type ConvergedinfraAdapterComplianceDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The driver name (e.g. nenic, nfnic) of the adapter.
	DriverName *string `json:"DriverName,omitempty"`
	// The driver version of the adapter.
	DriverVersion *string `json:"DriverVersion,omitempty"`
	// The firmware version of the adapter.
	Firmware *string `json:"Firmware,omitempty"`
	// The HCL compatibility status for the adapter. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
	HclStatus *string `json:"HclStatus,omitempty"`
	// The reason for the HCL status when it is not Approved. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID. * `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result. * `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated. * `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.
	HclStatusReason *string `json:"HclStatusReason,omitempty"`
	// The model information of the adapter.
	Model                   *string                                            `json:"Model,omitempty"`
	Adapter                 *AdapterUnitRelationship                           `json:"Adapter,omitempty"`
	ServerComplianceDetails *ConvergedinfraServerComplianceDetailsRelationship `json:"ServerComplianceDetails,omitempty"`
	// An array of relationships to convergedinfraStorageComplianceDetails resources.
	StorageCompliances   []ConvergedinfraStorageComplianceDetailsRelationship `json:"StorageCompliances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraAdapterComplianceDetailsAllOf ConvergedinfraAdapterComplianceDetailsAllOf

// NewConvergedinfraAdapterComplianceDetailsAllOf instantiates a new ConvergedinfraAdapterComplianceDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraAdapterComplianceDetailsAllOf(classId string, objectType string) *ConvergedinfraAdapterComplianceDetailsAllOf {
	this := ConvergedinfraAdapterComplianceDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraAdapterComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraAdapterComplianceDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraAdapterComplianceDetailsAllOfWithDefaults() *ConvergedinfraAdapterComplianceDetailsAllOf {
	this := ConvergedinfraAdapterComplianceDetailsAllOf{}
	var classId string = "convergedinfra.AdapterComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.AdapterComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriverName returns the DriverName field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetDriverName() string {
	if o == nil || o.DriverName == nil {
		var ret string
		return ret
	}
	return *o.DriverName
}

// GetDriverNameOk returns a tuple with the DriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetDriverNameOk() (*string, bool) {
	if o == nil || o.DriverName == nil {
		return nil, false
	}
	return o.DriverName, true
}

// HasDriverName returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasDriverName() bool {
	if o != nil && o.DriverName != nil {
		return true
	}

	return false
}

// SetDriverName gets a reference to the given string and assigns it to the DriverName field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetDriverName(v string) {
	o.DriverName = &v
}

// GetDriverVersion returns the DriverVersion field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetDriverVersion() string {
	if o == nil || o.DriverVersion == nil {
		var ret string
		return ret
	}
	return *o.DriverVersion
}

// GetDriverVersionOk returns a tuple with the DriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetDriverVersionOk() (*string, bool) {
	if o == nil || o.DriverVersion == nil {
		return nil, false
	}
	return o.DriverVersion, true
}

// HasDriverVersion returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasDriverVersion() bool {
	if o != nil && o.DriverVersion != nil {
		return true
	}

	return false
}

// SetDriverVersion gets a reference to the given string and assigns it to the DriverVersion field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetDriverVersion(v string) {
	o.DriverVersion = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetFirmwareOk() (*string, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetFirmware(v string) {
	o.Firmware = &v
}

// GetHclStatus returns the HclStatus field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetHclStatus() string {
	if o == nil || o.HclStatus == nil {
		var ret string
		return ret
	}
	return *o.HclStatus
}

// GetHclStatusOk returns a tuple with the HclStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetHclStatusOk() (*string, bool) {
	if o == nil || o.HclStatus == nil {
		return nil, false
	}
	return o.HclStatus, true
}

// HasHclStatus returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasHclStatus() bool {
	if o != nil && o.HclStatus != nil {
		return true
	}

	return false
}

// SetHclStatus gets a reference to the given string and assigns it to the HclStatus field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetHclStatus(v string) {
	o.HclStatus = &v
}

// GetHclStatusReason returns the HclStatusReason field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetHclStatusReason() string {
	if o == nil || o.HclStatusReason == nil {
		var ret string
		return ret
	}
	return *o.HclStatusReason
}

// GetHclStatusReasonOk returns a tuple with the HclStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetHclStatusReasonOk() (*string, bool) {
	if o == nil || o.HclStatusReason == nil {
		return nil, false
	}
	return o.HclStatusReason, true
}

// HasHclStatusReason returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasHclStatusReason() bool {
	if o != nil && o.HclStatusReason != nil {
		return true
	}

	return false
}

// SetHclStatusReason gets a reference to the given string and assigns it to the HclStatusReason field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetHclStatusReason(v string) {
	o.HclStatusReason = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetModel(v string) {
	o.Model = &v
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetAdapter() AdapterUnitRelationship {
	if o == nil || o.Adapter == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetAdapterOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.Adapter == nil {
		return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasAdapter() bool {
	if o != nil && o.Adapter != nil {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given AdapterUnitRelationship and assigns it to the Adapter field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetAdapter(v AdapterUnitRelationship) {
	o.Adapter = &v
}

// GetServerComplianceDetails returns the ServerComplianceDetails field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetServerComplianceDetails() ConvergedinfraServerComplianceDetailsRelationship {
	if o == nil || o.ServerComplianceDetails == nil {
		var ret ConvergedinfraServerComplianceDetailsRelationship
		return ret
	}
	return *o.ServerComplianceDetails
}

// GetServerComplianceDetailsOk returns a tuple with the ServerComplianceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetServerComplianceDetailsOk() (*ConvergedinfraServerComplianceDetailsRelationship, bool) {
	if o == nil || o.ServerComplianceDetails == nil {
		return nil, false
	}
	return o.ServerComplianceDetails, true
}

// HasServerComplianceDetails returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasServerComplianceDetails() bool {
	if o != nil && o.ServerComplianceDetails != nil {
		return true
	}

	return false
}

// SetServerComplianceDetails gets a reference to the given ConvergedinfraServerComplianceDetailsRelationship and assigns it to the ServerComplianceDetails field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetServerComplianceDetails(v ConvergedinfraServerComplianceDetailsRelationship) {
	o.ServerComplianceDetails = &v
}

// GetStorageCompliances returns the StorageCompliances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship {
	if o == nil {
		var ret []ConvergedinfraStorageComplianceDetailsRelationship
		return ret
	}
	return o.StorageCompliances
}

// GetStorageCompliancesOk returns a tuple with the StorageCompliances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) GetStorageCompliancesOk() ([]ConvergedinfraStorageComplianceDetailsRelationship, bool) {
	if o == nil || o.StorageCompliances == nil {
		return nil, false
	}
	return o.StorageCompliances, true
}

// HasStorageCompliances returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) HasStorageCompliances() bool {
	if o != nil && o.StorageCompliances != nil {
		return true
	}

	return false
}

// SetStorageCompliances gets a reference to the given []ConvergedinfraStorageComplianceDetailsRelationship and assigns it to the StorageCompliances field.
func (o *ConvergedinfraAdapterComplianceDetailsAllOf) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship) {
	o.StorageCompliances = v
}

func (o ConvergedinfraAdapterComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriverName != nil {
		toSerialize["DriverName"] = o.DriverName
	}
	if o.DriverVersion != nil {
		toSerialize["DriverVersion"] = o.DriverVersion
	}
	if o.Firmware != nil {
		toSerialize["Firmware"] = o.Firmware
	}
	if o.HclStatus != nil {
		toSerialize["HclStatus"] = o.HclStatus
	}
	if o.HclStatusReason != nil {
		toSerialize["HclStatusReason"] = o.HclStatusReason
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Adapter != nil {
		toSerialize["Adapter"] = o.Adapter
	}
	if o.ServerComplianceDetails != nil {
		toSerialize["ServerComplianceDetails"] = o.ServerComplianceDetails
	}
	if o.StorageCompliances != nil {
		toSerialize["StorageCompliances"] = o.StorageCompliances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraAdapterComplianceDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraAdapterComplianceDetailsAllOf := _ConvergedinfraAdapterComplianceDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraAdapterComplianceDetailsAllOf); err == nil {
		*o = ConvergedinfraAdapterComplianceDetailsAllOf(varConvergedinfraAdapterComplianceDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriverName")
		delete(additionalProperties, "DriverVersion")
		delete(additionalProperties, "Firmware")
		delete(additionalProperties, "HclStatus")
		delete(additionalProperties, "HclStatusReason")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Adapter")
		delete(additionalProperties, "ServerComplianceDetails")
		delete(additionalProperties, "StorageCompliances")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraAdapterComplianceDetailsAllOf struct {
	value *ConvergedinfraAdapterComplianceDetailsAllOf
	isSet bool
}

func (v NullableConvergedinfraAdapterComplianceDetailsAllOf) Get() *ConvergedinfraAdapterComplianceDetailsAllOf {
	return v.value
}

func (v *NullableConvergedinfraAdapterComplianceDetailsAllOf) Set(val *ConvergedinfraAdapterComplianceDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraAdapterComplianceDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraAdapterComplianceDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraAdapterComplianceDetailsAllOf(val *ConvergedinfraAdapterComplianceDetailsAllOf) *NullableConvergedinfraAdapterComplianceDetailsAllOf {
	return &NullableConvergedinfraAdapterComplianceDetailsAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraAdapterComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraAdapterComplianceDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
