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

// checks if the ConvergedinfraServerComplianceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvergedinfraServerComplianceDetails{}

// ConvergedinfraServerComplianceDetails The compliance details of a server in a converged infrastructure pod.
type ConvergedinfraServerComplianceDetails struct {
	ConvergedinfraBaseComplianceDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of ethernet NIC adapters in the server.
	AdapterCount *int64 `json:"AdapterCount,omitempty"`
	// The Cisco IMC firmware version of the server.
	Firmware *string `json:"Firmware,omitempty"`
	// The HCL compatibility status of the server. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
	HclStatus *string `json:"HclStatus,omitempty"`
	// The reason for server's HCL status. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information. * `Incompatible-Server` - The validation failed for this server because the server's model was not listed in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not listed for the given server model. * `Incompatible-Os-Info` - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination. * `Incompatible-Firmware` - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has returned a service error or unrecognized result. * `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted. * `Incompatible-Components` - The validation has failed for this server because one or more components have \"Not-Listed\" status. * `Compatible` - The validation has passed for this server's model, processor, OS vendor and version.
	HclStatusReason *string `json:"HclStatusReason,omitempty"`
	// The model information of the server.
	Model *string `json:"Model,omitempty"`
	// Details of name and version of the operating system running on the server.
	Os *string `json:"Os,omitempty"`
	// Details of platform of the server, examples are B-Series, C-Series, X-Series etc.
	Platform *string `json:"Platform,omitempty"`
	// The processor information of the server.
	Processor            *string                                             `json:"Processor,omitempty"`
	PodCompliance        NullableConvergedinfraPodComplianceInfoRelationship `json:"PodCompliance,omitempty"`
	Server               NullableComputePhysicalSummaryRelationship          `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraServerComplianceDetails ConvergedinfraServerComplianceDetails

// NewConvergedinfraServerComplianceDetails instantiates a new ConvergedinfraServerComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraServerComplianceDetails(classId string, objectType string) *ConvergedinfraServerComplianceDetails {
	this := ConvergedinfraServerComplianceDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraServerComplianceDetailsWithDefaults instantiates a new ConvergedinfraServerComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraServerComplianceDetailsWithDefaults() *ConvergedinfraServerComplianceDetails {
	this := ConvergedinfraServerComplianceDetails{}
	var classId string = "convergedinfra.ServerComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.ServerComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraServerComplianceDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraServerComplianceDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "convergedinfra.ServerComplianceDetails" of the ClassId field.
func (o *ConvergedinfraServerComplianceDetails) GetDefaultClassId() interface{} {
	return "convergedinfra.ServerComplianceDetails"
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraServerComplianceDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraServerComplianceDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "convergedinfra.ServerComplianceDetails" of the ObjectType field.
func (o *ConvergedinfraServerComplianceDetails) GetDefaultObjectType() interface{} {
	return "convergedinfra.ServerComplianceDetails"
}

// GetAdapterCount returns the AdapterCount field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetAdapterCount() int64 {
	if o == nil || IsNil(o.AdapterCount) {
		var ret int64
		return ret
	}
	return *o.AdapterCount
}

// GetAdapterCountOk returns a tuple with the AdapterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetAdapterCountOk() (*int64, bool) {
	if o == nil || IsNil(o.AdapterCount) {
		return nil, false
	}
	return o.AdapterCount, true
}

// HasAdapterCount returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasAdapterCount() bool {
	if o != nil && !IsNil(o.AdapterCount) {
		return true
	}

	return false
}

// SetAdapterCount gets a reference to the given int64 and assigns it to the AdapterCount field.
func (o *ConvergedinfraServerComplianceDetails) SetAdapterCount(v int64) {
	o.AdapterCount = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetFirmware() string {
	if o == nil || IsNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetFirmwareOk() (*string, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *ConvergedinfraServerComplianceDetails) SetFirmware(v string) {
	o.Firmware = &v
}

// GetHclStatus returns the HclStatus field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetHclStatus() string {
	if o == nil || IsNil(o.HclStatus) {
		var ret string
		return ret
	}
	return *o.HclStatus
}

// GetHclStatusOk returns a tuple with the HclStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetHclStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HclStatus) {
		return nil, false
	}
	return o.HclStatus, true
}

// HasHclStatus returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasHclStatus() bool {
	if o != nil && !IsNil(o.HclStatus) {
		return true
	}

	return false
}

// SetHclStatus gets a reference to the given string and assigns it to the HclStatus field.
func (o *ConvergedinfraServerComplianceDetails) SetHclStatus(v string) {
	o.HclStatus = &v
}

// GetHclStatusReason returns the HclStatusReason field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetHclStatusReason() string {
	if o == nil || IsNil(o.HclStatusReason) {
		var ret string
		return ret
	}
	return *o.HclStatusReason
}

// GetHclStatusReasonOk returns a tuple with the HclStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetHclStatusReasonOk() (*string, bool) {
	if o == nil || IsNil(o.HclStatusReason) {
		return nil, false
	}
	return o.HclStatusReason, true
}

// HasHclStatusReason returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasHclStatusReason() bool {
	if o != nil && !IsNil(o.HclStatusReason) {
		return true
	}

	return false
}

// SetHclStatusReason gets a reference to the given string and assigns it to the HclStatusReason field.
func (o *ConvergedinfraServerComplianceDetails) SetHclStatusReason(v string) {
	o.HclStatusReason = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConvergedinfraServerComplianceDetails) SetModel(v string) {
	o.Model = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ConvergedinfraServerComplianceDetails) SetOs(v string) {
	o.Os = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ConvergedinfraServerComplianceDetails) SetPlatform(v string) {
	o.Platform = &v
}

// GetProcessor returns the Processor field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetails) GetProcessor() string {
	if o == nil || IsNil(o.Processor) {
		var ret string
		return ret
	}
	return *o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetails) GetProcessorOk() (*string, bool) {
	if o == nil || IsNil(o.Processor) {
		return nil, false
	}
	return o.Processor, true
}

// HasProcessor returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasProcessor() bool {
	if o != nil && !IsNil(o.Processor) {
		return true
	}

	return false
}

// SetProcessor gets a reference to the given string and assigns it to the Processor field.
func (o *ConvergedinfraServerComplianceDetails) SetProcessor(v string) {
	o.Processor = &v
}

// GetPodCompliance returns the PodCompliance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraServerComplianceDetails) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship {
	if o == nil || IsNil(o.PodCompliance.Get()) {
		var ret ConvergedinfraPodComplianceInfoRelationship
		return ret
	}
	return *o.PodCompliance.Get()
}

// GetPodComplianceOk returns a tuple with the PodCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraServerComplianceDetails) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PodCompliance.Get(), o.PodCompliance.IsSet()
}

// HasPodCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasPodCompliance() bool {
	if o != nil && o.PodCompliance.IsSet() {
		return true
	}

	return false
}

// SetPodCompliance gets a reference to the given NullableConvergedinfraPodComplianceInfoRelationship and assigns it to the PodCompliance field.
func (o *ConvergedinfraServerComplianceDetails) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship) {
	o.PodCompliance.Set(&v)
}

// SetPodComplianceNil sets the value for PodCompliance to be an explicit nil
func (o *ConvergedinfraServerComplianceDetails) SetPodComplianceNil() {
	o.PodCompliance.Set(nil)
}

// UnsetPodCompliance ensures that no value is present for PodCompliance, not even an explicit nil
func (o *ConvergedinfraServerComplianceDetails) UnsetPodCompliance() {
	o.PodCompliance.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraServerComplianceDetails) GetServer() ComputePhysicalSummaryRelationship {
	if o == nil || IsNil(o.Server.Get()) {
		var ret ComputePhysicalSummaryRelationship
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraServerComplianceDetails) GetServerOk() (*ComputePhysicalSummaryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetails) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableComputePhysicalSummaryRelationship and assigns it to the Server field.
func (o *ConvergedinfraServerComplianceDetails) SetServer(v ComputePhysicalSummaryRelationship) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *ConvergedinfraServerComplianceDetails) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *ConvergedinfraServerComplianceDetails) UnsetServer() {
	o.Server.Unset()
}

func (o ConvergedinfraServerComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvergedinfraServerComplianceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConvergedinfraBaseComplianceDetails, errConvergedinfraBaseComplianceDetails := json.Marshal(o.ConvergedinfraBaseComplianceDetails)
	if errConvergedinfraBaseComplianceDetails != nil {
		return map[string]interface{}{}, errConvergedinfraBaseComplianceDetails
	}
	errConvergedinfraBaseComplianceDetails = json.Unmarshal([]byte(serializedConvergedinfraBaseComplianceDetails), &toSerialize)
	if errConvergedinfraBaseComplianceDetails != nil {
		return map[string]interface{}{}, errConvergedinfraBaseComplianceDetails
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdapterCount) {
		toSerialize["AdapterCount"] = o.AdapterCount
	}
	if !IsNil(o.Firmware) {
		toSerialize["Firmware"] = o.Firmware
	}
	if !IsNil(o.HclStatus) {
		toSerialize["HclStatus"] = o.HclStatus
	}
	if !IsNil(o.HclStatusReason) {
		toSerialize["HclStatusReason"] = o.HclStatusReason
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.Os) {
		toSerialize["Os"] = o.Os
	}
	if !IsNil(o.Platform) {
		toSerialize["Platform"] = o.Platform
	}
	if !IsNil(o.Processor) {
		toSerialize["Processor"] = o.Processor
	}
	if o.PodCompliance.IsSet() {
		toSerialize["PodCompliance"] = o.PodCompliance.Get()
	}
	if o.Server.IsSet() {
		toSerialize["Server"] = o.Server.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConvergedinfraServerComplianceDetails) UnmarshalJSON(data []byte) (err error) {
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
	type ConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of ethernet NIC adapters in the server.
		AdapterCount *int64 `json:"AdapterCount,omitempty"`
		// The Cisco IMC firmware version of the server.
		Firmware *string `json:"Firmware,omitempty"`
		// The HCL compatibility status of the server. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
		HclStatus *string `json:"HclStatus,omitempty"`
		// The reason for server's HCL status. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information. * `Incompatible-Server` - The validation failed for this server because the server's model was not listed in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not listed for the given server model. * `Incompatible-Os-Info` - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination. * `Incompatible-Firmware` - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has returned a service error or unrecognized result. * `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted. * `Incompatible-Components` - The validation has failed for this server because one or more components have \"Not-Listed\" status. * `Compatible` - The validation has passed for this server's model, processor, OS vendor and version.
		HclStatusReason *string `json:"HclStatusReason,omitempty"`
		// The model information of the server.
		Model *string `json:"Model,omitempty"`
		// Details of name and version of the operating system running on the server.
		Os *string `json:"Os,omitempty"`
		// Details of platform of the server, examples are B-Series, C-Series, X-Series etc.
		Platform *string `json:"Platform,omitempty"`
		// The processor information of the server.
		Processor     *string                                             `json:"Processor,omitempty"`
		PodCompliance NullableConvergedinfraPodComplianceInfoRelationship `json:"PodCompliance,omitempty"`
		Server        NullableComputePhysicalSummaryRelationship          `json:"Server,omitempty"`
	}

	varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct := ConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraServerComplianceDetails := _ConvergedinfraServerComplianceDetails{}
		varConvergedinfraServerComplianceDetails.ClassId = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.ClassId
		varConvergedinfraServerComplianceDetails.ObjectType = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.ObjectType
		varConvergedinfraServerComplianceDetails.AdapterCount = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.AdapterCount
		varConvergedinfraServerComplianceDetails.Firmware = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.Firmware
		varConvergedinfraServerComplianceDetails.HclStatus = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.HclStatus
		varConvergedinfraServerComplianceDetails.HclStatusReason = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.HclStatusReason
		varConvergedinfraServerComplianceDetails.Model = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.Model
		varConvergedinfraServerComplianceDetails.Os = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.Os
		varConvergedinfraServerComplianceDetails.Platform = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.Platform
		varConvergedinfraServerComplianceDetails.Processor = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.Processor
		varConvergedinfraServerComplianceDetails.PodCompliance = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.PodCompliance
		varConvergedinfraServerComplianceDetails.Server = varConvergedinfraServerComplianceDetailsWithoutEmbeddedStruct.Server
		*o = ConvergedinfraServerComplianceDetails(varConvergedinfraServerComplianceDetails)
	} else {
		return err
	}

	varConvergedinfraServerComplianceDetails := _ConvergedinfraServerComplianceDetails{}

	err = json.Unmarshal(data, &varConvergedinfraServerComplianceDetails)
	if err == nil {
		o.ConvergedinfraBaseComplianceDetails = varConvergedinfraServerComplianceDetails.ConvergedinfraBaseComplianceDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterCount")
		delete(additionalProperties, "Firmware")
		delete(additionalProperties, "HclStatus")
		delete(additionalProperties, "HclStatusReason")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Os")
		delete(additionalProperties, "Platform")
		delete(additionalProperties, "Processor")
		delete(additionalProperties, "PodCompliance")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectConvergedinfraBaseComplianceDetails := reflect.ValueOf(o.ConvergedinfraBaseComplianceDetails)
		for i := 0; i < reflectConvergedinfraBaseComplianceDetails.Type().NumField(); i++ {
			t := reflectConvergedinfraBaseComplianceDetails.Type().Field(i)

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

type NullableConvergedinfraServerComplianceDetails struct {
	value *ConvergedinfraServerComplianceDetails
	isSet bool
}

func (v NullableConvergedinfraServerComplianceDetails) Get() *ConvergedinfraServerComplianceDetails {
	return v.value
}

func (v *NullableConvergedinfraServerComplianceDetails) Set(val *ConvergedinfraServerComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraServerComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraServerComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraServerComplianceDetails(val *ConvergedinfraServerComplianceDetails) *NullableConvergedinfraServerComplianceDetails {
	return &NullableConvergedinfraServerComplianceDetails{value: val, isSet: true}
}

func (v NullableConvergedinfraServerComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraServerComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
