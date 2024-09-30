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
)

// checks if the CatalystsdwanVedgeDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalystsdwanVedgeDevice{}

// CatalystsdwanVedgeDevice Details for the Catalyst SDWAN Vedge entities.
type CatalystsdwanVedgeDevice struct {
	EquipmentAbstractDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Catalyst SDWAN device config status message.
	ConfigStatusMessage *string `json:"ConfigStatusMessage,omitempty"`
	// The Catalyst SDWAN device id.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The Catalyst SDWAN device state.
	DeviceState *string `json:"DeviceState,omitempty"`
	// The Catalyst SDWAN device host name.
	HostName *string `json:"HostName,omitempty"`
	// The Catalyst SDWAN device platform family.
	PlatformFamily *string `json:"PlatformFamily,omitempty"`
	// The Catalyst SDWAN device reachability.
	Reachability *string `json:"Reachability,omitempty"`
	// The Catalyst SDWAN device site id.
	SiteId *string `json:"SiteId,omitempty"`
	// The Catalyst SDWAN device site name.
	SiteName *string `json:"SiteName,omitempty"`
	// The Catalyst SDWAN device sp organization name.
	SpOrganizationName *string `json:"SpOrganizationName,omitempty"`
	// The Catalyst SDWAN device system IP.
	SystemIp *string `json:"SystemIp,omitempty"`
	// The Catalyst SDWAN device template status.
	TemplateStatus *string `json:"TemplateStatus,omitempty"`
	// The Catalyst SDWAN device validity.
	Validity             *string                                     `json:"Validity,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalystsdwanVedgeDevice CatalystsdwanVedgeDevice

// NewCatalystsdwanVedgeDevice instantiates a new CatalystsdwanVedgeDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalystsdwanVedgeDevice(classId string, objectType string) *CatalystsdwanVedgeDevice {
	this := CatalystsdwanVedgeDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCatalystsdwanVedgeDeviceWithDefaults instantiates a new CatalystsdwanVedgeDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalystsdwanVedgeDeviceWithDefaults() *CatalystsdwanVedgeDevice {
	this := CatalystsdwanVedgeDevice{}
	var classId string = "catalystsdwan.VedgeDevice"
	this.ClassId = classId
	var objectType string = "catalystsdwan.VedgeDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CatalystsdwanVedgeDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CatalystsdwanVedgeDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "catalystsdwan.VedgeDevice" of the ClassId field.
func (o *CatalystsdwanVedgeDevice) GetDefaultClassId() interface{} {
	return "catalystsdwan.VedgeDevice"
}

// GetObjectType returns the ObjectType field value
func (o *CatalystsdwanVedgeDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CatalystsdwanVedgeDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "catalystsdwan.VedgeDevice" of the ObjectType field.
func (o *CatalystsdwanVedgeDevice) GetDefaultObjectType() interface{} {
	return "catalystsdwan.VedgeDevice"
}

// GetConfigStatusMessage returns the ConfigStatusMessage field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetConfigStatusMessage() string {
	if o == nil || IsNil(o.ConfigStatusMessage) {
		var ret string
		return ret
	}
	return *o.ConfigStatusMessage
}

// GetConfigStatusMessageOk returns a tuple with the ConfigStatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetConfigStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigStatusMessage) {
		return nil, false
	}
	return o.ConfigStatusMessage, true
}

// HasConfigStatusMessage returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasConfigStatusMessage() bool {
	if o != nil && !IsNil(o.ConfigStatusMessage) {
		return true
	}

	return false
}

// SetConfigStatusMessage gets a reference to the given string and assigns it to the ConfigStatusMessage field.
func (o *CatalystsdwanVedgeDevice) SetConfigStatusMessage(v string) {
	o.ConfigStatusMessage = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *CatalystsdwanVedgeDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceState returns the DeviceState field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetDeviceState() string {
	if o == nil || IsNil(o.DeviceState) {
		var ret string
		return ret
	}
	return *o.DeviceState
}

// GetDeviceStateOk returns a tuple with the DeviceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetDeviceStateOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceState) {
		return nil, false
	}
	return o.DeviceState, true
}

// HasDeviceState returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasDeviceState() bool {
	if o != nil && !IsNil(o.DeviceState) {
		return true
	}

	return false
}

// SetDeviceState gets a reference to the given string and assigns it to the DeviceState field.
func (o *CatalystsdwanVedgeDevice) SetDeviceState(v string) {
	o.DeviceState = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *CatalystsdwanVedgeDevice) SetHostName(v string) {
	o.HostName = &v
}

// GetPlatformFamily returns the PlatformFamily field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetPlatformFamily() string {
	if o == nil || IsNil(o.PlatformFamily) {
		var ret string
		return ret
	}
	return *o.PlatformFamily
}

// GetPlatformFamilyOk returns a tuple with the PlatformFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetPlatformFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformFamily) {
		return nil, false
	}
	return o.PlatformFamily, true
}

// HasPlatformFamily returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasPlatformFamily() bool {
	if o != nil && !IsNil(o.PlatformFamily) {
		return true
	}

	return false
}

// SetPlatformFamily gets a reference to the given string and assigns it to the PlatformFamily field.
func (o *CatalystsdwanVedgeDevice) SetPlatformFamily(v string) {
	o.PlatformFamily = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetReachability() string {
	if o == nil || IsNil(o.Reachability) {
		var ret string
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetReachabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Reachability) {
		return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasReachability() bool {
	if o != nil && !IsNil(o.Reachability) {
		return true
	}

	return false
}

// SetReachability gets a reference to the given string and assigns it to the Reachability field.
func (o *CatalystsdwanVedgeDevice) SetReachability(v string) {
	o.Reachability = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *CatalystsdwanVedgeDevice) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *CatalystsdwanVedgeDevice) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSpOrganizationName returns the SpOrganizationName field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetSpOrganizationName() string {
	if o == nil || IsNil(o.SpOrganizationName) {
		var ret string
		return ret
	}
	return *o.SpOrganizationName
}

// GetSpOrganizationNameOk returns a tuple with the SpOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetSpOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpOrganizationName) {
		return nil, false
	}
	return o.SpOrganizationName, true
}

// HasSpOrganizationName returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasSpOrganizationName() bool {
	if o != nil && !IsNil(o.SpOrganizationName) {
		return true
	}

	return false
}

// SetSpOrganizationName gets a reference to the given string and assigns it to the SpOrganizationName field.
func (o *CatalystsdwanVedgeDevice) SetSpOrganizationName(v string) {
	o.SpOrganizationName = &v
}

// GetSystemIp returns the SystemIp field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetSystemIp() string {
	if o == nil || IsNil(o.SystemIp) {
		var ret string
		return ret
	}
	return *o.SystemIp
}

// GetSystemIpOk returns a tuple with the SystemIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetSystemIpOk() (*string, bool) {
	if o == nil || IsNil(o.SystemIp) {
		return nil, false
	}
	return o.SystemIp, true
}

// HasSystemIp returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasSystemIp() bool {
	if o != nil && !IsNil(o.SystemIp) {
		return true
	}

	return false
}

// SetSystemIp gets a reference to the given string and assigns it to the SystemIp field.
func (o *CatalystsdwanVedgeDevice) SetSystemIp(v string) {
	o.SystemIp = &v
}

// GetTemplateStatus returns the TemplateStatus field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetTemplateStatus() string {
	if o == nil || IsNil(o.TemplateStatus) {
		var ret string
		return ret
	}
	return *o.TemplateStatus
}

// GetTemplateStatusOk returns a tuple with the TemplateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetTemplateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateStatus) {
		return nil, false
	}
	return o.TemplateStatus, true
}

// HasTemplateStatus returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasTemplateStatus() bool {
	if o != nil && !IsNil(o.TemplateStatus) {
		return true
	}

	return false
}

// SetTemplateStatus gets a reference to the given string and assigns it to the TemplateStatus field.
func (o *CatalystsdwanVedgeDevice) SetTemplateStatus(v string) {
	o.TemplateStatus = &v
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *CatalystsdwanVedgeDevice) GetValidity() string {
	if o == nil || IsNil(o.Validity) {
		var ret string
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanVedgeDevice) GetValidityOk() (*string, bool) {
	if o == nil || IsNil(o.Validity) {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasValidity() bool {
	if o != nil && !IsNil(o.Validity) {
		return true
	}

	return false
}

// SetValidity gets a reference to the given string and assigns it to the Validity field.
func (o *CatalystsdwanVedgeDevice) SetValidity(v string) {
	o.Validity = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalystsdwanVedgeDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalystsdwanVedgeDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CatalystsdwanVedgeDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CatalystsdwanVedgeDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *CatalystsdwanVedgeDevice) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *CatalystsdwanVedgeDevice) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o CatalystsdwanVedgeDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalystsdwanVedgeDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentAbstractDevice, errEquipmentAbstractDevice := json.Marshal(o.EquipmentAbstractDevice)
	if errEquipmentAbstractDevice != nil {
		return map[string]interface{}{}, errEquipmentAbstractDevice
	}
	errEquipmentAbstractDevice = json.Unmarshal([]byte(serializedEquipmentAbstractDevice), &toSerialize)
	if errEquipmentAbstractDevice != nil {
		return map[string]interface{}{}, errEquipmentAbstractDevice
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConfigStatusMessage) {
		toSerialize["ConfigStatusMessage"] = o.ConfigStatusMessage
	}
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.DeviceState) {
		toSerialize["DeviceState"] = o.DeviceState
	}
	if !IsNil(o.HostName) {
		toSerialize["HostName"] = o.HostName
	}
	if !IsNil(o.PlatformFamily) {
		toSerialize["PlatformFamily"] = o.PlatformFamily
	}
	if !IsNil(o.Reachability) {
		toSerialize["Reachability"] = o.Reachability
	}
	if !IsNil(o.SiteId) {
		toSerialize["SiteId"] = o.SiteId
	}
	if !IsNil(o.SiteName) {
		toSerialize["SiteName"] = o.SiteName
	}
	if !IsNil(o.SpOrganizationName) {
		toSerialize["SpOrganizationName"] = o.SpOrganizationName
	}
	if !IsNil(o.SystemIp) {
		toSerialize["SystemIp"] = o.SystemIp
	}
	if !IsNil(o.TemplateStatus) {
		toSerialize["TemplateStatus"] = o.TemplateStatus
	}
	if !IsNil(o.Validity) {
		toSerialize["Validity"] = o.Validity
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalystsdwanVedgeDevice) UnmarshalJSON(data []byte) (err error) {
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
	type CatalystsdwanVedgeDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Catalyst SDWAN device config status message.
		ConfigStatusMessage *string `json:"ConfigStatusMessage,omitempty"`
		// The Catalyst SDWAN device id.
		DeviceId *string `json:"DeviceId,omitempty"`
		// The Catalyst SDWAN device state.
		DeviceState *string `json:"DeviceState,omitempty"`
		// The Catalyst SDWAN device host name.
		HostName *string `json:"HostName,omitempty"`
		// The Catalyst SDWAN device platform family.
		PlatformFamily *string `json:"PlatformFamily,omitempty"`
		// The Catalyst SDWAN device reachability.
		Reachability *string `json:"Reachability,omitempty"`
		// The Catalyst SDWAN device site id.
		SiteId *string `json:"SiteId,omitempty"`
		// The Catalyst SDWAN device site name.
		SiteName *string `json:"SiteName,omitempty"`
		// The Catalyst SDWAN device sp organization name.
		SpOrganizationName *string `json:"SpOrganizationName,omitempty"`
		// The Catalyst SDWAN device system IP.
		SystemIp *string `json:"SystemIp,omitempty"`
		// The Catalyst SDWAN device template status.
		TemplateStatus *string `json:"TemplateStatus,omitempty"`
		// The Catalyst SDWAN device validity.
		Validity         *string                                     `json:"Validity,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct := CatalystsdwanVedgeDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct)
	if err == nil {
		varCatalystsdwanVedgeDevice := _CatalystsdwanVedgeDevice{}
		varCatalystsdwanVedgeDevice.ClassId = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.ClassId
		varCatalystsdwanVedgeDevice.ObjectType = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.ObjectType
		varCatalystsdwanVedgeDevice.ConfigStatusMessage = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.ConfigStatusMessage
		varCatalystsdwanVedgeDevice.DeviceId = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.DeviceId
		varCatalystsdwanVedgeDevice.DeviceState = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.DeviceState
		varCatalystsdwanVedgeDevice.HostName = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.HostName
		varCatalystsdwanVedgeDevice.PlatformFamily = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.PlatformFamily
		varCatalystsdwanVedgeDevice.Reachability = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.Reachability
		varCatalystsdwanVedgeDevice.SiteId = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.SiteId
		varCatalystsdwanVedgeDevice.SiteName = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.SiteName
		varCatalystsdwanVedgeDevice.SpOrganizationName = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.SpOrganizationName
		varCatalystsdwanVedgeDevice.SystemIp = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.SystemIp
		varCatalystsdwanVedgeDevice.TemplateStatus = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.TemplateStatus
		varCatalystsdwanVedgeDevice.Validity = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.Validity
		varCatalystsdwanVedgeDevice.RegisteredDevice = varCatalystsdwanVedgeDeviceWithoutEmbeddedStruct.RegisteredDevice
		*o = CatalystsdwanVedgeDevice(varCatalystsdwanVedgeDevice)
	} else {
		return err
	}

	varCatalystsdwanVedgeDevice := _CatalystsdwanVedgeDevice{}

	err = json.Unmarshal(data, &varCatalystsdwanVedgeDevice)
	if err == nil {
		o.EquipmentAbstractDevice = varCatalystsdwanVedgeDevice.EquipmentAbstractDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigStatusMessage")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "DeviceState")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "PlatformFamily")
		delete(additionalProperties, "Reachability")
		delete(additionalProperties, "SiteId")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SpOrganizationName")
		delete(additionalProperties, "SystemIp")
		delete(additionalProperties, "TemplateStatus")
		delete(additionalProperties, "Validity")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentAbstractDevice := reflect.ValueOf(o.EquipmentAbstractDevice)
		for i := 0; i < reflectEquipmentAbstractDevice.Type().NumField(); i++ {
			t := reflectEquipmentAbstractDevice.Type().Field(i)

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

type NullableCatalystsdwanVedgeDevice struct {
	value *CatalystsdwanVedgeDevice
	isSet bool
}

func (v NullableCatalystsdwanVedgeDevice) Get() *CatalystsdwanVedgeDevice {
	return v.value
}

func (v *NullableCatalystsdwanVedgeDevice) Set(val *CatalystsdwanVedgeDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalystsdwanVedgeDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalystsdwanVedgeDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalystsdwanVedgeDevice(val *CatalystsdwanVedgeDevice) *NullableCatalystsdwanVedgeDevice {
	return &NullableCatalystsdwanVedgeDevice{value: val, isSet: true}
}

func (v NullableCatalystsdwanVedgeDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalystsdwanVedgeDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
