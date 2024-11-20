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

// checks if the KubernetesNodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodeInfo{}

// KubernetesNodeInfo Kubernetes Node Information.
type KubernetesNodeInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Node Operating System architecture (amd64, arm64).
	Architecture *string `json:"Architecture,omitempty"`
	// A Boot ID is an Identifier generated by the host everytimes it gets reboot.
	BootId *string `json:"BootId,omitempty"`
	// To run containers in Pods, Kubernetes uses a container runtime. This field describes Container Runtime Version.
	ContainerRuntimeVersion *string `json:"ContainerRuntimeVersion,omitempty"`
	// Node Operating System kernel version.
	KernelVersion *string `json:"KernelVersion,omitempty"`
	// The Kubernetes network proxy runs on each node. This reflects services as defined in the Kubernetes API on each node and can do simple TCP, UDP, and SCTP stream forwarding or round robin TCP, UDP, and SCTP forwarding across a set of backends. This field describes the kube-proxy version.
	KubeProxyVersion *string `json:"KubeProxyVersion,omitempty"`
	// The kubelet is the primary \"node agent\" that runs on each node. It can register the node with the apiserver using one of such as the hostname; a flag to override the hostname; or specific logic for a cloud provider. This field describes the kubelet version the node currently using.
	KubeletVersion *string `json:"KubeletVersion,omitempty"`
	// It is a node identifier in Kubernetes. When the node joins a kubernetes cluster, Kubernetes will assign a machine ID to that node. Learn more from man machine-id from http://man7.org/linux/man-pages/man5/machine-id.5.html.
	MachineId *string `json:"MachineId,omitempty"`
	// Operating System installed on this Node.
	OperatingSystem *string `json:"OperatingSystem,omitempty"`
	// Node current Operating System image.
	OsImage *string `json:"OsImage,omitempty"`
	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html.
	SystemUuid           *string `json:"SystemUuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeInfo KubernetesNodeInfo

// NewKubernetesNodeInfo instantiates a new KubernetesNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeInfo(classId string, objectType string) *KubernetesNodeInfo {
	this := KubernetesNodeInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeInfoWithDefaults instantiates a new KubernetesNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeInfoWithDefaults() *KubernetesNodeInfo {
	this := KubernetesNodeInfo{}
	var classId string = "kubernetes.NodeInfo"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.NodeInfo" of the ClassId field.
func (o *KubernetesNodeInfo) GetDefaultClassId() interface{} {
	return "kubernetes.NodeInfo"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.NodeInfo" of the ObjectType field.
func (o *KubernetesNodeInfo) GetDefaultObjectType() interface{} {
	return "kubernetes.NodeInfo"
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *KubernetesNodeInfo) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBootId returns the BootId field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetBootId() string {
	if o == nil || IsNil(o.BootId) {
		var ret string
		return ret
	}
	return *o.BootId
}

// GetBootIdOk returns a tuple with the BootId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetBootIdOk() (*string, bool) {
	if o == nil || IsNil(o.BootId) {
		return nil, false
	}
	return o.BootId, true
}

// HasBootId returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasBootId() bool {
	if o != nil && !IsNil(o.BootId) {
		return true
	}

	return false
}

// SetBootId gets a reference to the given string and assigns it to the BootId field.
func (o *KubernetesNodeInfo) SetBootId(v string) {
	o.BootId = &v
}

// GetContainerRuntimeVersion returns the ContainerRuntimeVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetContainerRuntimeVersion() string {
	if o == nil || IsNil(o.ContainerRuntimeVersion) {
		var ret string
		return ret
	}
	return *o.ContainerRuntimeVersion
}

// GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetContainerRuntimeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerRuntimeVersion) {
		return nil, false
	}
	return o.ContainerRuntimeVersion, true
}

// HasContainerRuntimeVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasContainerRuntimeVersion() bool {
	if o != nil && !IsNil(o.ContainerRuntimeVersion) {
		return true
	}

	return false
}

// SetContainerRuntimeVersion gets a reference to the given string and assigns it to the ContainerRuntimeVersion field.
func (o *KubernetesNodeInfo) SetContainerRuntimeVersion(v string) {
	o.ContainerRuntimeVersion = &v
}

// GetKernelVersion returns the KernelVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetKernelVersion() string {
	if o == nil || IsNil(o.KernelVersion) {
		var ret string
		return ret
	}
	return *o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetKernelVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KernelVersion) {
		return nil, false
	}
	return o.KernelVersion, true
}

// HasKernelVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasKernelVersion() bool {
	if o != nil && !IsNil(o.KernelVersion) {
		return true
	}

	return false
}

// SetKernelVersion gets a reference to the given string and assigns it to the KernelVersion field.
func (o *KubernetesNodeInfo) SetKernelVersion(v string) {
	o.KernelVersion = &v
}

// GetKubeProxyVersion returns the KubeProxyVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetKubeProxyVersion() string {
	if o == nil || IsNil(o.KubeProxyVersion) {
		var ret string
		return ret
	}
	return *o.KubeProxyVersion
}

// GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetKubeProxyVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubeProxyVersion) {
		return nil, false
	}
	return o.KubeProxyVersion, true
}

// HasKubeProxyVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasKubeProxyVersion() bool {
	if o != nil && !IsNil(o.KubeProxyVersion) {
		return true
	}

	return false
}

// SetKubeProxyVersion gets a reference to the given string and assigns it to the KubeProxyVersion field.
func (o *KubernetesNodeInfo) SetKubeProxyVersion(v string) {
	o.KubeProxyVersion = &v
}

// GetKubeletVersion returns the KubeletVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetKubeletVersion() string {
	if o == nil || IsNil(o.KubeletVersion) {
		var ret string
		return ret
	}
	return *o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetKubeletVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubeletVersion) {
		return nil, false
	}
	return o.KubeletVersion, true
}

// HasKubeletVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasKubeletVersion() bool {
	if o != nil && !IsNil(o.KubeletVersion) {
		return true
	}

	return false
}

// SetKubeletVersion gets a reference to the given string and assigns it to the KubeletVersion field.
func (o *KubernetesNodeInfo) SetKubeletVersion(v string) {
	o.KubeletVersion = &v
}

// GetMachineId returns the MachineId field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetMachineId() string {
	if o == nil || IsNil(o.MachineId) {
		var ret string
		return ret
	}
	return *o.MachineId
}

// GetMachineIdOk returns a tuple with the MachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetMachineIdOk() (*string, bool) {
	if o == nil || IsNil(o.MachineId) {
		return nil, false
	}
	return o.MachineId, true
}

// HasMachineId returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasMachineId() bool {
	if o != nil && !IsNil(o.MachineId) {
		return true
	}

	return false
}

// SetMachineId gets a reference to the given string and assigns it to the MachineId field.
func (o *KubernetesNodeInfo) SetMachineId(v string) {
	o.MachineId = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetOperatingSystem() string {
	if o == nil || IsNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *KubernetesNodeInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetOsImage returns the OsImage field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetOsImage() string {
	if o == nil || IsNil(o.OsImage) {
		var ret string
		return ret
	}
	return *o.OsImage
}

// GetOsImageOk returns a tuple with the OsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetOsImageOk() (*string, bool) {
	if o == nil || IsNil(o.OsImage) {
		return nil, false
	}
	return o.OsImage, true
}

// HasOsImage returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasOsImage() bool {
	if o != nil && !IsNil(o.OsImage) {
		return true
	}

	return false
}

// SetOsImage gets a reference to the given string and assigns it to the OsImage field.
func (o *KubernetesNodeInfo) SetOsImage(v string) {
	o.OsImage = &v
}

// GetSystemUuid returns the SystemUuid field value if set, zero value otherwise.
func (o *KubernetesNodeInfo) GetSystemUuid() string {
	if o == nil || IsNil(o.SystemUuid) {
		var ret string
		return ret
	}
	return *o.SystemUuid
}

// GetSystemUuidOk returns a tuple with the SystemUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfo) GetSystemUuidOk() (*string, bool) {
	if o == nil || IsNil(o.SystemUuid) {
		return nil, false
	}
	return o.SystemUuid, true
}

// HasSystemUuid returns a boolean if a field has been set.
func (o *KubernetesNodeInfo) HasSystemUuid() bool {
	if o != nil && !IsNil(o.SystemUuid) {
		return true
	}

	return false
}

// SetSystemUuid gets a reference to the given string and assigns it to the SystemUuid field.
func (o *KubernetesNodeInfo) SetSystemUuid(v string) {
	o.SystemUuid = &v
}

func (o KubernetesNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNodeInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Architecture) {
		toSerialize["Architecture"] = o.Architecture
	}
	if !IsNil(o.BootId) {
		toSerialize["BootId"] = o.BootId
	}
	if !IsNil(o.ContainerRuntimeVersion) {
		toSerialize["ContainerRuntimeVersion"] = o.ContainerRuntimeVersion
	}
	if !IsNil(o.KernelVersion) {
		toSerialize["KernelVersion"] = o.KernelVersion
	}
	if !IsNil(o.KubeProxyVersion) {
		toSerialize["KubeProxyVersion"] = o.KubeProxyVersion
	}
	if !IsNil(o.KubeletVersion) {
		toSerialize["KubeletVersion"] = o.KubeletVersion
	}
	if !IsNil(o.MachineId) {
		toSerialize["MachineId"] = o.MachineId
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["OperatingSystem"] = o.OperatingSystem
	}
	if !IsNil(o.OsImage) {
		toSerialize["OsImage"] = o.OsImage
	}
	if !IsNil(o.SystemUuid) {
		toSerialize["SystemUuid"] = o.SystemUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesNodeInfo) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesNodeInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Node Operating System architecture (amd64, arm64).
		Architecture *string `json:"Architecture,omitempty"`
		// A Boot ID is an Identifier generated by the host everytimes it gets reboot.
		BootId *string `json:"BootId,omitempty"`
		// To run containers in Pods, Kubernetes uses a container runtime. This field describes Container Runtime Version.
		ContainerRuntimeVersion *string `json:"ContainerRuntimeVersion,omitempty"`
		// Node Operating System kernel version.
		KernelVersion *string `json:"KernelVersion,omitempty"`
		// The Kubernetes network proxy runs on each node. This reflects services as defined in the Kubernetes API on each node and can do simple TCP, UDP, and SCTP stream forwarding or round robin TCP, UDP, and SCTP forwarding across a set of backends. This field describes the kube-proxy version.
		KubeProxyVersion *string `json:"KubeProxyVersion,omitempty"`
		// The kubelet is the primary \"node agent\" that runs on each node. It can register the node with the apiserver using one of such as the hostname; a flag to override the hostname; or specific logic for a cloud provider. This field describes the kubelet version the node currently using.
		KubeletVersion *string `json:"KubeletVersion,omitempty"`
		// It is a node identifier in Kubernetes. When the node joins a kubernetes cluster, Kubernetes will assign a machine ID to that node. Learn more from man machine-id from http://man7.org/linux/man-pages/man5/machine-id.5.html.
		MachineId *string `json:"MachineId,omitempty"`
		// Operating System installed on this Node.
		OperatingSystem *string `json:"OperatingSystem,omitempty"`
		// Node current Operating System image.
		OsImage *string `json:"OsImage,omitempty"`
		// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html.
		SystemUuid *string `json:"SystemUuid,omitempty"`
	}

	varKubernetesNodeInfoWithoutEmbeddedStruct := KubernetesNodeInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesNodeInfoWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNodeInfo := _KubernetesNodeInfo{}
		varKubernetesNodeInfo.ClassId = varKubernetesNodeInfoWithoutEmbeddedStruct.ClassId
		varKubernetesNodeInfo.ObjectType = varKubernetesNodeInfoWithoutEmbeddedStruct.ObjectType
		varKubernetesNodeInfo.Architecture = varKubernetesNodeInfoWithoutEmbeddedStruct.Architecture
		varKubernetesNodeInfo.BootId = varKubernetesNodeInfoWithoutEmbeddedStruct.BootId
		varKubernetesNodeInfo.ContainerRuntimeVersion = varKubernetesNodeInfoWithoutEmbeddedStruct.ContainerRuntimeVersion
		varKubernetesNodeInfo.KernelVersion = varKubernetesNodeInfoWithoutEmbeddedStruct.KernelVersion
		varKubernetesNodeInfo.KubeProxyVersion = varKubernetesNodeInfoWithoutEmbeddedStruct.KubeProxyVersion
		varKubernetesNodeInfo.KubeletVersion = varKubernetesNodeInfoWithoutEmbeddedStruct.KubeletVersion
		varKubernetesNodeInfo.MachineId = varKubernetesNodeInfoWithoutEmbeddedStruct.MachineId
		varKubernetesNodeInfo.OperatingSystem = varKubernetesNodeInfoWithoutEmbeddedStruct.OperatingSystem
		varKubernetesNodeInfo.OsImage = varKubernetesNodeInfoWithoutEmbeddedStruct.OsImage
		varKubernetesNodeInfo.SystemUuid = varKubernetesNodeInfoWithoutEmbeddedStruct.SystemUuid
		*o = KubernetesNodeInfo(varKubernetesNodeInfo)
	} else {
		return err
	}

	varKubernetesNodeInfo := _KubernetesNodeInfo{}

	err = json.Unmarshal(data, &varKubernetesNodeInfo)
	if err == nil {
		o.MoBaseComplexType = varKubernetesNodeInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Architecture")
		delete(additionalProperties, "BootId")
		delete(additionalProperties, "ContainerRuntimeVersion")
		delete(additionalProperties, "KernelVersion")
		delete(additionalProperties, "KubeProxyVersion")
		delete(additionalProperties, "KubeletVersion")
		delete(additionalProperties, "MachineId")
		delete(additionalProperties, "OperatingSystem")
		delete(additionalProperties, "OsImage")
		delete(additionalProperties, "SystemUuid")

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

type NullableKubernetesNodeInfo struct {
	value *KubernetesNodeInfo
	isSet bool
}

func (v NullableKubernetesNodeInfo) Get() *KubernetesNodeInfo {
	return v.value
}

func (v *NullableKubernetesNodeInfo) Set(val *KubernetesNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeInfo(val *KubernetesNodeInfo) *NullableKubernetesNodeInfo {
	return &NullableKubernetesNodeInfo{value: val, isSet: true}
}

func (v NullableKubernetesNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
