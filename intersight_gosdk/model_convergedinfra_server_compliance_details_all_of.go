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

// ConvergedinfraServerComplianceDetailsAllOf Definition of the list of properties defined in 'convergedinfra.ServerComplianceDetails', excluding properties defined in parent classes.
type ConvergedinfraServerComplianceDetailsAllOf struct {
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
	Processor            *string                                      `json:"Processor,omitempty"`
	PodCompliance        *ConvergedinfraPodComplianceInfoRelationship `json:"PodCompliance,omitempty"`
	Server               *ComputePhysicalSummaryRelationship          `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraServerComplianceDetailsAllOf ConvergedinfraServerComplianceDetailsAllOf

// NewConvergedinfraServerComplianceDetailsAllOf instantiates a new ConvergedinfraServerComplianceDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraServerComplianceDetailsAllOf(classId string, objectType string) *ConvergedinfraServerComplianceDetailsAllOf {
	this := ConvergedinfraServerComplianceDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraServerComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraServerComplianceDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraServerComplianceDetailsAllOfWithDefaults() *ConvergedinfraServerComplianceDetailsAllOf {
	this := ConvergedinfraServerComplianceDetailsAllOf{}
	var classId string = "convergedinfra.ServerComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.ServerComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterCount returns the AdapterCount field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetAdapterCount() int64 {
	if o == nil || o.AdapterCount == nil {
		var ret int64
		return ret
	}
	return *o.AdapterCount
}

// GetAdapterCountOk returns a tuple with the AdapterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetAdapterCountOk() (*int64, bool) {
	if o == nil || o.AdapterCount == nil {
		return nil, false
	}
	return o.AdapterCount, true
}

// HasAdapterCount returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasAdapterCount() bool {
	if o != nil && o.AdapterCount != nil {
		return true
	}

	return false
}

// SetAdapterCount gets a reference to the given int64 and assigns it to the AdapterCount field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetAdapterCount(v int64) {
	o.AdapterCount = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetFirmwareOk() (*string, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetFirmware(v string) {
	o.Firmware = &v
}

// GetHclStatus returns the HclStatus field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatus() string {
	if o == nil || o.HclStatus == nil {
		var ret string
		return ret
	}
	return *o.HclStatus
}

// GetHclStatusOk returns a tuple with the HclStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatusOk() (*string, bool) {
	if o == nil || o.HclStatus == nil {
		return nil, false
	}
	return o.HclStatus, true
}

// HasHclStatus returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasHclStatus() bool {
	if o != nil && o.HclStatus != nil {
		return true
	}

	return false
}

// SetHclStatus gets a reference to the given string and assigns it to the HclStatus field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetHclStatus(v string) {
	o.HclStatus = &v
}

// GetHclStatusReason returns the HclStatusReason field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatusReason() string {
	if o == nil || o.HclStatusReason == nil {
		var ret string
		return ret
	}
	return *o.HclStatusReason
}

// GetHclStatusReasonOk returns a tuple with the HclStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatusReasonOk() (*string, bool) {
	if o == nil || o.HclStatusReason == nil {
		return nil, false
	}
	return o.HclStatusReason, true
}

// HasHclStatusReason returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasHclStatusReason() bool {
	if o != nil && o.HclStatusReason != nil {
		return true
	}

	return false
}

// SetHclStatusReason gets a reference to the given string and assigns it to the HclStatusReason field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetHclStatusReason(v string) {
	o.HclStatusReason = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetModel(v string) {
	o.Model = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetOs(v string) {
	o.Os = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetPlatform(v string) {
	o.Platform = &v
}

// GetProcessor returns the Processor field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetProcessor() string {
	if o == nil || o.Processor == nil {
		var ret string
		return ret
	}
	return *o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetProcessorOk() (*string, bool) {
	if o == nil || o.Processor == nil {
		return nil, false
	}
	return o.Processor, true
}

// HasProcessor returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasProcessor() bool {
	if o != nil && o.Processor != nil {
		return true
	}

	return false
}

// SetProcessor gets a reference to the given string and assigns it to the Processor field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetProcessor(v string) {
	o.Processor = &v
}

// GetPodCompliance returns the PodCompliance field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship {
	if o == nil || o.PodCompliance == nil {
		var ret ConvergedinfraPodComplianceInfoRelationship
		return ret
	}
	return *o.PodCompliance
}

// GetPodComplianceOk returns a tuple with the PodCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool) {
	if o == nil || o.PodCompliance == nil {
		return nil, false
	}
	return o.PodCompliance, true
}

// HasPodCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasPodCompliance() bool {
	if o != nil && o.PodCompliance != nil {
		return true
	}

	return false
}

// SetPodCompliance gets a reference to the given ConvergedinfraPodComplianceInfoRelationship and assigns it to the PodCompliance field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship) {
	o.PodCompliance = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetServer() ComputePhysicalSummaryRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalSummaryRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) GetServerOk() (*ComputePhysicalSummaryRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ConvergedinfraServerComplianceDetailsAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalSummaryRelationship and assigns it to the Server field.
func (o *ConvergedinfraServerComplianceDetailsAllOf) SetServer(v ComputePhysicalSummaryRelationship) {
	o.Server = &v
}

func (o ConvergedinfraServerComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterCount != nil {
		toSerialize["AdapterCount"] = o.AdapterCount
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
	if o.Os != nil {
		toSerialize["Os"] = o.Os
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}
	if o.Processor != nil {
		toSerialize["Processor"] = o.Processor
	}
	if o.PodCompliance != nil {
		toSerialize["PodCompliance"] = o.PodCompliance
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraServerComplianceDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraServerComplianceDetailsAllOf := _ConvergedinfraServerComplianceDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraServerComplianceDetailsAllOf); err == nil {
		*o = ConvergedinfraServerComplianceDetailsAllOf(varConvergedinfraServerComplianceDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraServerComplianceDetailsAllOf struct {
	value *ConvergedinfraServerComplianceDetailsAllOf
	isSet bool
}

func (v NullableConvergedinfraServerComplianceDetailsAllOf) Get() *ConvergedinfraServerComplianceDetailsAllOf {
	return v.value
}

func (v *NullableConvergedinfraServerComplianceDetailsAllOf) Set(val *ConvergedinfraServerComplianceDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraServerComplianceDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraServerComplianceDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraServerComplianceDetailsAllOf(val *ConvergedinfraServerComplianceDetailsAllOf) *NullableConvergedinfraServerComplianceDetailsAllOf {
	return &NullableConvergedinfraServerComplianceDetailsAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraServerComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraServerComplianceDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
