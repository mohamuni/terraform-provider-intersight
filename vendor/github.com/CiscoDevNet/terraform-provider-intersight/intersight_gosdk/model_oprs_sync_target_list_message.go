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
	"time"
)

// checks if the OprsSyncTargetListMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OprsSyncTargetListMessage{}

// OprsSyncTargetListMessage The targets sync messages are sent to assist and back to euclid for reconciliation.
type OprsSyncTargetListMessage struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The unique id of the assist which is associated with this message.
	AssistId *string `json:"AssistId,omitempty"`
	// The target type for which this sync message is meant for. * `` - An unrecognized platform type. * `APIC` - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * `CAPIC` - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * `DCNM` - A Cisco Data Center Network Manager (DCNM) instance. * `UCSFI` - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * `UCSFIISM` - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * `IMC` - A standalone Cisco UCS rack server (Deprecated). * `IMCM4` - A standalone Cisco UCS C-Series or S-Series M4 server. * `IMCM5` - A standalone Cisco UCS C-Series or S-Series M5 server. * `IMCRack` - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * `UCSIOM` - A Cisco UCS Blade Chassis I/O Module (IOM). * `HX` - A Cisco HyperFlex (HX) cluster. * `UCSD` - A Cisco UCS Director (UCSD) instance. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance instance. * `IntersightAssist` - A Cisco Intersight Assist instance. * `PureStorageFlashArray` - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * `NexusDevice` - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * `ACISwitch` - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * `NexusSwitch` - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * `MDSSwitch` - A Cisco MDS Switch that is managed using the embedded Device Connector. * `MDSDevice` - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * `UCSC890` - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * `RedfishServer` - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * `NetAppOntap` - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * `EmcScaleIo` - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVmax` - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVplex` - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcXtremIo` - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `VmwareVcenter` - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * `MicrosoftHyperV` - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * `AppDynamics` - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * `Dynatrace` - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `NewRelic` - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `Umbrella` - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `MySqlServer` - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `OracleDatabaseServer` - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `IBMWebSphereApplicationServer` - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * `ApacheTomcatServer` - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `JavaVirtualMachine` - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * `Kubernetes` - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * `AmazonWebService` - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * `AmazonWebServiceBilling` - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatform` - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatformBilling` - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureBilling` - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `DellCompellent` - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `HPE3Par` - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * `NutanixPrismCentral` - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * `HPEOneView` - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * `GenericTarget` - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * `IMCBlade` - A Cisco UCS blade server managed by Cisco Intersight. * `TerraformCloud` - A Terraform Cloud Business Tier account. * `TerraformAgent` - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * `CustomTarget` - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * `AnsibleEndpoint` - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * `HTTPEndpoint` - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * `SSHEndpoint` - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * `CiscoDNAC` - A Cisco Digital Network Architecture (DNA) Center appliance. * `CiscoFMC` - A Cisco Secure Firewall Management Center. * `ViptelaCloud` - A Cisco Viptela SD-WAN Cloud. * `MerakiCloud` - A Cisco Meraki Organization. * `CiscoISE` - A Cisco Identity Services Engine (ISE) target.
	TargetType *string  `json:"TargetType,omitempty"`
	Targets    []string `json:"Targets,omitempty"`
	// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
	TimeStamp            *time.Time                                  `json:"TimeStamp,omitempty"`
	Assist               NullableAssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OprsSyncTargetListMessage OprsSyncTargetListMessage

// NewOprsSyncTargetListMessage instantiates a new OprsSyncTargetListMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOprsSyncTargetListMessage(classId string, objectType string) *OprsSyncTargetListMessage {
	this := OprsSyncTargetListMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewOprsSyncTargetListMessageWithDefaults instantiates a new OprsSyncTargetListMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOprsSyncTargetListMessageWithDefaults() *OprsSyncTargetListMessage {
	this := OprsSyncTargetListMessage{}
	var classId string = "oprs.SyncTargetListMessage"
	this.ClassId = classId
	var objectType string = "oprs.SyncTargetListMessage"
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OprsSyncTargetListMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OprsSyncTargetListMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "oprs.SyncTargetListMessage" of the ClassId field.
func (o *OprsSyncTargetListMessage) GetDefaultClassId() interface{} {
	return "oprs.SyncTargetListMessage"
}

// GetObjectType returns the ObjectType field value
func (o *OprsSyncTargetListMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OprsSyncTargetListMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "oprs.SyncTargetListMessage" of the ObjectType field.
func (o *OprsSyncTargetListMessage) GetDefaultObjectType() interface{} {
	return "oprs.SyncTargetListMessage"
}

// GetAssistId returns the AssistId field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetAssistId() string {
	if o == nil || IsNil(o.AssistId) {
		var ret string
		return ret
	}
	return *o.AssistId
}

// GetAssistIdOk returns a tuple with the AssistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetAssistIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssistId) {
		return nil, false
	}
	return o.AssistId, true
}

// HasAssistId returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasAssistId() bool {
	if o != nil && !IsNil(o.AssistId) {
		return true
	}

	return false
}

// SetAssistId gets a reference to the given string and assigns it to the AssistId field.
func (o *OprsSyncTargetListMessage) SetAssistId(v string) {
	o.AssistId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *OprsSyncTargetListMessage) SetTargetType(v string) {
	o.TargetType = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OprsSyncTargetListMessage) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OprsSyncTargetListMessage) GetTargetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *OprsSyncTargetListMessage) SetTargets(v []string) {
	o.Targets = v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetTimeStamp() time.Time {
	if o == nil || IsNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *OprsSyncTargetListMessage) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OprsSyncTargetListMessage) GetAssist() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Assist.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Assist.Get()
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OprsSyncTargetListMessage) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assist.Get(), o.Assist.IsSet()
}

// HasAssist returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasAssist() bool {
	if o != nil && o.Assist.IsSet() {
		return true
	}

	return false
}

// SetAssist gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Assist field.
func (o *OprsSyncTargetListMessage) SetAssist(v AssetDeviceRegistrationRelationship) {
	o.Assist.Set(&v)
}

// SetAssistNil sets the value for Assist to be an explicit nil
func (o *OprsSyncTargetListMessage) SetAssistNil() {
	o.Assist.Set(nil)
}

// UnsetAssist ensures that no value is present for Assist, not even an explicit nil
func (o *OprsSyncTargetListMessage) UnsetAssist() {
	o.Assist.Unset()
}

func (o OprsSyncTargetListMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OprsSyncTargetListMessage) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AssistId) {
		toSerialize["AssistId"] = o.AssistId
	}
	if !IsNil(o.TargetType) {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["TimeStamp"] = o.TimeStamp
	}
	if o.Assist.IsSet() {
		toSerialize["Assist"] = o.Assist.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OprsSyncTargetListMessage) UnmarshalJSON(data []byte) (err error) {
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
	type OprsSyncTargetListMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The unique id of the assist which is associated with this message.
		AssistId *string `json:"AssistId,omitempty"`
		// The target type for which this sync message is meant for. * `` - An unrecognized platform type. * `APIC` - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * `CAPIC` - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * `DCNM` - A Cisco Data Center Network Manager (DCNM) instance. * `UCSFI` - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * `UCSFIISM` - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * `IMC` - A standalone Cisco UCS rack server (Deprecated). * `IMCM4` - A standalone Cisco UCS C-Series or S-Series M4 server. * `IMCM5` - A standalone Cisco UCS C-Series or S-Series M5 server. * `IMCRack` - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * `UCSIOM` - A Cisco UCS Blade Chassis I/O Module (IOM). * `HX` - A Cisco HyperFlex (HX) cluster. * `UCSD` - A Cisco UCS Director (UCSD) instance. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance instance. * `IntersightAssist` - A Cisco Intersight Assist instance. * `PureStorageFlashArray` - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * `NexusDevice` - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * `ACISwitch` - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * `NexusSwitch` - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * `MDSSwitch` - A Cisco MDS Switch that is managed using the embedded Device Connector. * `MDSDevice` - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * `UCSC890` - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * `RedfishServer` - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * `NetAppOntap` - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * `EmcScaleIo` - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVmax` - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVplex` - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcXtremIo` - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `VmwareVcenter` - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * `MicrosoftHyperV` - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * `AppDynamics` - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * `Dynatrace` - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `NewRelic` - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `Umbrella` - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `MySqlServer` - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `OracleDatabaseServer` - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `IBMWebSphereApplicationServer` - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * `ApacheTomcatServer` - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `JavaVirtualMachine` - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * `Kubernetes` - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * `AmazonWebService` - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * `AmazonWebServiceBilling` - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatform` - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatformBilling` - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureBilling` - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `DellCompellent` - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `HPE3Par` - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * `NutanixPrismCentral` - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * `HPEOneView` - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * `GenericTarget` - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * `IMCBlade` - A Cisco UCS blade server managed by Cisco Intersight. * `TerraformCloud` - A Terraform Cloud Business Tier account. * `TerraformAgent` - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * `CustomTarget` - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * `AnsibleEndpoint` - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * `HTTPEndpoint` - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * `SSHEndpoint` - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * `CiscoDNAC` - A Cisco Digital Network Architecture (DNA) Center appliance. * `CiscoFMC` - A Cisco Secure Firewall Management Center. * `ViptelaCloud` - A Cisco Viptela SD-WAN Cloud. * `MerakiCloud` - A Cisco Meraki Organization. * `CiscoISE` - A Cisco Identity Services Engine (ISE) target.
		TargetType *string  `json:"TargetType,omitempty"`
		Targets    []string `json:"Targets,omitempty"`
		// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
		TimeStamp *time.Time                                  `json:"TimeStamp,omitempty"`
		Assist    NullableAssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	}

	varOprsSyncTargetListMessageWithoutEmbeddedStruct := OprsSyncTargetListMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOprsSyncTargetListMessageWithoutEmbeddedStruct)
	if err == nil {
		varOprsSyncTargetListMessage := _OprsSyncTargetListMessage{}
		varOprsSyncTargetListMessage.ClassId = varOprsSyncTargetListMessageWithoutEmbeddedStruct.ClassId
		varOprsSyncTargetListMessage.ObjectType = varOprsSyncTargetListMessageWithoutEmbeddedStruct.ObjectType
		varOprsSyncTargetListMessage.AssistId = varOprsSyncTargetListMessageWithoutEmbeddedStruct.AssistId
		varOprsSyncTargetListMessage.TargetType = varOprsSyncTargetListMessageWithoutEmbeddedStruct.TargetType
		varOprsSyncTargetListMessage.Targets = varOprsSyncTargetListMessageWithoutEmbeddedStruct.Targets
		varOprsSyncTargetListMessage.TimeStamp = varOprsSyncTargetListMessageWithoutEmbeddedStruct.TimeStamp
		varOprsSyncTargetListMessage.Assist = varOprsSyncTargetListMessageWithoutEmbeddedStruct.Assist
		*o = OprsSyncTargetListMessage(varOprsSyncTargetListMessage)
	} else {
		return err
	}

	varOprsSyncTargetListMessage := _OprsSyncTargetListMessage{}

	err = json.Unmarshal(data, &varOprsSyncTargetListMessage)
	if err == nil {
		o.MoBaseMo = varOprsSyncTargetListMessage.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssistId")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "TimeStamp")
		delete(additionalProperties, "Assist")

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

type NullableOprsSyncTargetListMessage struct {
	value *OprsSyncTargetListMessage
	isSet bool
}

func (v NullableOprsSyncTargetListMessage) Get() *OprsSyncTargetListMessage {
	return v.value
}

func (v *NullableOprsSyncTargetListMessage) Set(val *OprsSyncTargetListMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOprsSyncTargetListMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOprsSyncTargetListMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOprsSyncTargetListMessage(val *OprsSyncTargetListMessage) *NullableOprsSyncTargetListMessage {
	return &NullableOprsSyncTargetListMessage{value: val, isSet: true}
}

func (v NullableOprsSyncTargetListMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOprsSyncTargetListMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
