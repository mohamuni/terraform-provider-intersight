/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ApplianceNodeIpInfo NodeInfo contains the hostname, ip address, and node number for a node.
type ApplianceNodeIpInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Gateway ip address of the cluster node.
	Gateway *string `json:"Gateway,omitempty"`
	// Hostname of the cluster node.
	Hostname *string `json:"Hostname,omitempty"`
	// Ip address of the cluster node.
	IpAddress *string `json:"IpAddress,omitempty"`
	// Netmask of the cluster node.
	Netmask *string `json:"Netmask,omitempty"`
	// Id number of the cluster node.
	NodeId *int64 `json:"NodeId,omitempty"`
	// Moid of the corresponding appliance.ClusterInfo or appliance.NodeInfo mo.
	NodeMoid *string `json:"NodeMoid,omitempty"`
	// Status of the cluster node. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNodeIpInfo ApplianceNodeIpInfo

// NewApplianceNodeIpInfo instantiates a new ApplianceNodeIpInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNodeIpInfo(classId string, objectType string) *ApplianceNodeIpInfo {
	this := ApplianceNodeIpInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNodeIpInfoWithDefaults instantiates a new ApplianceNodeIpInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNodeIpInfoWithDefaults() *ApplianceNodeIpInfo {
	this := ApplianceNodeIpInfo{}
	var classId string = "appliance.NodeIpInfo"
	this.ClassId = classId
	var objectType string = "appliance.NodeIpInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNodeIpInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNodeIpInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNodeIpInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNodeIpInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *ApplianceNodeIpInfo) SetGateway(v string) {
	o.Gateway = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceNodeIpInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ApplianceNodeIpInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *ApplianceNodeIpInfo) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceNodeIpInfo) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeMoid returns the NodeMoid field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetNodeMoid() string {
	if o == nil || o.NodeMoid == nil {
		var ret string
		return ret
	}
	return *o.NodeMoid
}

// GetNodeMoidOk returns a tuple with the NodeMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetNodeMoidOk() (*string, bool) {
	if o == nil || o.NodeMoid == nil {
		return nil, false
	}
	return o.NodeMoid, true
}

// HasNodeMoid returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasNodeMoid() bool {
	if o != nil && o.NodeMoid != nil {
		return true
	}

	return false
}

// SetNodeMoid gets a reference to the given string and assigns it to the NodeMoid field.
func (o *ApplianceNodeIpInfo) SetNodeMoid(v string) {
	o.NodeMoid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceNodeIpInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeIpInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceNodeIpInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceNodeIpInfo) SetStatus(v string) {
	o.Status = &v
}

func (o ApplianceNodeIpInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.NodeMoid != nil {
		toSerialize["NodeMoid"] = o.NodeMoid
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceNodeIpInfo) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceNodeIpInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Gateway ip address of the cluster node.
		Gateway *string `json:"Gateway,omitempty"`
		// Hostname of the cluster node.
		Hostname *string `json:"Hostname,omitempty"`
		// Ip address of the cluster node.
		IpAddress *string `json:"IpAddress,omitempty"`
		// Netmask of the cluster node.
		Netmask *string `json:"Netmask,omitempty"`
		// Id number of the cluster node.
		NodeId *int64 `json:"NodeId,omitempty"`
		// Moid of the corresponding appliance.ClusterInfo or appliance.NodeInfo mo.
		NodeMoid *string `json:"NodeMoid,omitempty"`
		// Status of the cluster node. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced.
		Status *string `json:"Status,omitempty"`
	}

	varApplianceNodeIpInfoWithoutEmbeddedStruct := ApplianceNodeIpInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceNodeIpInfoWithoutEmbeddedStruct)
	if err == nil {
		varApplianceNodeIpInfo := _ApplianceNodeIpInfo{}
		varApplianceNodeIpInfo.ClassId = varApplianceNodeIpInfoWithoutEmbeddedStruct.ClassId
		varApplianceNodeIpInfo.ObjectType = varApplianceNodeIpInfoWithoutEmbeddedStruct.ObjectType
		varApplianceNodeIpInfo.Gateway = varApplianceNodeIpInfoWithoutEmbeddedStruct.Gateway
		varApplianceNodeIpInfo.Hostname = varApplianceNodeIpInfoWithoutEmbeddedStruct.Hostname
		varApplianceNodeIpInfo.IpAddress = varApplianceNodeIpInfoWithoutEmbeddedStruct.IpAddress
		varApplianceNodeIpInfo.Netmask = varApplianceNodeIpInfoWithoutEmbeddedStruct.Netmask
		varApplianceNodeIpInfo.NodeId = varApplianceNodeIpInfoWithoutEmbeddedStruct.NodeId
		varApplianceNodeIpInfo.NodeMoid = varApplianceNodeIpInfoWithoutEmbeddedStruct.NodeMoid
		varApplianceNodeIpInfo.Status = varApplianceNodeIpInfoWithoutEmbeddedStruct.Status
		*o = ApplianceNodeIpInfo(varApplianceNodeIpInfo)
	} else {
		return err
	}

	varApplianceNodeIpInfo := _ApplianceNodeIpInfo{}

	err = json.Unmarshal(bytes, &varApplianceNodeIpInfo)
	if err == nil {
		o.MoBaseComplexType = varApplianceNodeIpInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "NodeMoid")
		delete(additionalProperties, "Status")

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

type NullableApplianceNodeIpInfo struct {
	value *ApplianceNodeIpInfo
	isSet bool
}

func (v NullableApplianceNodeIpInfo) Get() *ApplianceNodeIpInfo {
	return v.value
}

func (v *NullableApplianceNodeIpInfo) Set(val *ApplianceNodeIpInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNodeIpInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNodeIpInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNodeIpInfo(val *ApplianceNodeIpInfo) *NullableApplianceNodeIpInfo {
	return &NullableApplianceNodeIpInfo{value: val, isSet: true}
}

func (v NullableApplianceNodeIpInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNodeIpInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}