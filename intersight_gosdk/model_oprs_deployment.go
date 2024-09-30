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

// checks if the OprsDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OprsDeployment{}

// OprsDeployment Monitors the status of operator deployed in the assist.
type OprsDeployment struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Available number of replicas.
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
	// The expected number of replicas.
	DesiredReplicas *int64 `json:"DesiredReplicas,omitempty"`
	// The type of event which was triggered.
	Event  *string      `json:"Event,omitempty"`
	Labels []OprsKvpair `json:"Labels,omitempty"`
	// Agent name for which the event is triggered.
	Name *string `json:"Name,omitempty"`
	// Name space in which the agents are running.
	Namespace *string `json:"Namespace,omitempty"`
	// Status which shows if the resource is healthy or not. * `` - An Unknown status indicates that the resource status is not known. * `Healthy` - A healthy status indicates that the resource is healthy and running as per spec. * `Unhealthy` - An unhealthy status indicates that the resource is down.
	Status *string `json:"Status,omitempty"`
	// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
	TimeStamp *time.Time `json:"TimeStamp,omitempty"`
	// Number of replicas Unavailable.
	UnavailableReplicas  *int64                                      `json:"UnavailableReplicas,omitempty"`
	Assist               NullableAssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OprsDeployment OprsDeployment

// NewOprsDeployment instantiates a new OprsDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOprsDeployment(classId string, objectType string) *OprsDeployment {
	this := OprsDeployment{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// NewOprsDeploymentWithDefaults instantiates a new OprsDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOprsDeploymentWithDefaults() *OprsDeployment {
	this := OprsDeployment{}
	var classId string = "oprs.Deployment"
	this.ClassId = classId
	var objectType string = "oprs.Deployment"
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *OprsDeployment) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OprsDeployment) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "oprs.Deployment" of the ClassId field.
func (o *OprsDeployment) GetDefaultClassId() interface{} {
	return "oprs.Deployment"
}

// GetObjectType returns the ObjectType field value
func (o *OprsDeployment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OprsDeployment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "oprs.Deployment" of the ObjectType field.
func (o *OprsDeployment) GetDefaultObjectType() interface{} {
	return "oprs.Deployment"
}

// GetAvailableReplicas returns the AvailableReplicas field value if set, zero value otherwise.
func (o *OprsDeployment) GetAvailableReplicas() int64 {
	if o == nil || IsNil(o.AvailableReplicas) {
		var ret int64
		return ret
	}
	return *o.AvailableReplicas
}

// GetAvailableReplicasOk returns a tuple with the AvailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetAvailableReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.AvailableReplicas) {
		return nil, false
	}
	return o.AvailableReplicas, true
}

// HasAvailableReplicas returns a boolean if a field has been set.
func (o *OprsDeployment) HasAvailableReplicas() bool {
	if o != nil && !IsNil(o.AvailableReplicas) {
		return true
	}

	return false
}

// SetAvailableReplicas gets a reference to the given int64 and assigns it to the AvailableReplicas field.
func (o *OprsDeployment) SetAvailableReplicas(v int64) {
	o.AvailableReplicas = &v
}

// GetDesiredReplicas returns the DesiredReplicas field value if set, zero value otherwise.
func (o *OprsDeployment) GetDesiredReplicas() int64 {
	if o == nil || IsNil(o.DesiredReplicas) {
		var ret int64
		return ret
	}
	return *o.DesiredReplicas
}

// GetDesiredReplicasOk returns a tuple with the DesiredReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetDesiredReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.DesiredReplicas) {
		return nil, false
	}
	return o.DesiredReplicas, true
}

// HasDesiredReplicas returns a boolean if a field has been set.
func (o *OprsDeployment) HasDesiredReplicas() bool {
	if o != nil && !IsNil(o.DesiredReplicas) {
		return true
	}

	return false
}

// SetDesiredReplicas gets a reference to the given int64 and assigns it to the DesiredReplicas field.
func (o *OprsDeployment) SetDesiredReplicas(v int64) {
	o.DesiredReplicas = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *OprsDeployment) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *OprsDeployment) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *OprsDeployment) SetEvent(v string) {
	o.Event = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OprsDeployment) GetLabels() []OprsKvpair {
	if o == nil {
		var ret []OprsKvpair
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OprsDeployment) GetLabelsOk() ([]OprsKvpair, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *OprsDeployment) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []OprsKvpair and assigns it to the Labels field.
func (o *OprsDeployment) SetLabels(v []OprsKvpair) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OprsDeployment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OprsDeployment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OprsDeployment) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *OprsDeployment) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *OprsDeployment) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *OprsDeployment) SetNamespace(v string) {
	o.Namespace = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OprsDeployment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OprsDeployment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OprsDeployment) SetStatus(v string) {
	o.Status = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *OprsDeployment) GetTimeStamp() time.Time {
	if o == nil || IsNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *OprsDeployment) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *OprsDeployment) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetUnavailableReplicas returns the UnavailableReplicas field value if set, zero value otherwise.
func (o *OprsDeployment) GetUnavailableReplicas() int64 {
	if o == nil || IsNil(o.UnavailableReplicas) {
		var ret int64
		return ret
	}
	return *o.UnavailableReplicas
}

// GetUnavailableReplicasOk returns a tuple with the UnavailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeployment) GetUnavailableReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.UnavailableReplicas) {
		return nil, false
	}
	return o.UnavailableReplicas, true
}

// HasUnavailableReplicas returns a boolean if a field has been set.
func (o *OprsDeployment) HasUnavailableReplicas() bool {
	if o != nil && !IsNil(o.UnavailableReplicas) {
		return true
	}

	return false
}

// SetUnavailableReplicas gets a reference to the given int64 and assigns it to the UnavailableReplicas field.
func (o *OprsDeployment) SetUnavailableReplicas(v int64) {
	o.UnavailableReplicas = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OprsDeployment) GetAssist() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Assist.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Assist.Get()
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OprsDeployment) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assist.Get(), o.Assist.IsSet()
}

// HasAssist returns a boolean if a field has been set.
func (o *OprsDeployment) HasAssist() bool {
	if o != nil && o.Assist.IsSet() {
		return true
	}

	return false
}

// SetAssist gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Assist field.
func (o *OprsDeployment) SetAssist(v AssetDeviceRegistrationRelationship) {
	o.Assist.Set(&v)
}

// SetAssistNil sets the value for Assist to be an explicit nil
func (o *OprsDeployment) SetAssistNil() {
	o.Assist.Set(nil)
}

// UnsetAssist ensures that no value is present for Assist, not even an explicit nil
func (o *OprsDeployment) UnsetAssist() {
	o.Assist.Unset()
}

func (o OprsDeployment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OprsDeployment) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AvailableReplicas) {
		toSerialize["AvailableReplicas"] = o.AvailableReplicas
	}
	if !IsNil(o.DesiredReplicas) {
		toSerialize["DesiredReplicas"] = o.DesiredReplicas
	}
	if !IsNil(o.Event) {
		toSerialize["Event"] = o.Event
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["Namespace"] = o.Namespace
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["TimeStamp"] = o.TimeStamp
	}
	if !IsNil(o.UnavailableReplicas) {
		toSerialize["UnavailableReplicas"] = o.UnavailableReplicas
	}
	if o.Assist.IsSet() {
		toSerialize["Assist"] = o.Assist.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OprsDeployment) UnmarshalJSON(data []byte) (err error) {
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
	type OprsDeploymentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Available number of replicas.
		AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
		// The expected number of replicas.
		DesiredReplicas *int64 `json:"DesiredReplicas,omitempty"`
		// The type of event which was triggered.
		Event  *string      `json:"Event,omitempty"`
		Labels []OprsKvpair `json:"Labels,omitempty"`
		// Agent name for which the event is triggered.
		Name *string `json:"Name,omitempty"`
		// Name space in which the agents are running.
		Namespace *string `json:"Namespace,omitempty"`
		// Status which shows if the resource is healthy or not. * `` - An Unknown status indicates that the resource status is not known. * `Healthy` - A healthy status indicates that the resource is healthy and running as per spec. * `Unhealthy` - An unhealthy status indicates that the resource is down.
		Status *string `json:"Status,omitempty"`
		// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
		TimeStamp *time.Time `json:"TimeStamp,omitempty"`
		// Number of replicas Unavailable.
		UnavailableReplicas *int64                                      `json:"UnavailableReplicas,omitempty"`
		Assist              NullableAssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	}

	varOprsDeploymentWithoutEmbeddedStruct := OprsDeploymentWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOprsDeploymentWithoutEmbeddedStruct)
	if err == nil {
		varOprsDeployment := _OprsDeployment{}
		varOprsDeployment.ClassId = varOprsDeploymentWithoutEmbeddedStruct.ClassId
		varOprsDeployment.ObjectType = varOprsDeploymentWithoutEmbeddedStruct.ObjectType
		varOprsDeployment.AvailableReplicas = varOprsDeploymentWithoutEmbeddedStruct.AvailableReplicas
		varOprsDeployment.DesiredReplicas = varOprsDeploymentWithoutEmbeddedStruct.DesiredReplicas
		varOprsDeployment.Event = varOprsDeploymentWithoutEmbeddedStruct.Event
		varOprsDeployment.Labels = varOprsDeploymentWithoutEmbeddedStruct.Labels
		varOprsDeployment.Name = varOprsDeploymentWithoutEmbeddedStruct.Name
		varOprsDeployment.Namespace = varOprsDeploymentWithoutEmbeddedStruct.Namespace
		varOprsDeployment.Status = varOprsDeploymentWithoutEmbeddedStruct.Status
		varOprsDeployment.TimeStamp = varOprsDeploymentWithoutEmbeddedStruct.TimeStamp
		varOprsDeployment.UnavailableReplicas = varOprsDeploymentWithoutEmbeddedStruct.UnavailableReplicas
		varOprsDeployment.Assist = varOprsDeploymentWithoutEmbeddedStruct.Assist
		*o = OprsDeployment(varOprsDeployment)
	} else {
		return err
	}

	varOprsDeployment := _OprsDeployment{}

	err = json.Unmarshal(data, &varOprsDeployment)
	if err == nil {
		o.MoBaseMo = varOprsDeployment.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvailableReplicas")
		delete(additionalProperties, "DesiredReplicas")
		delete(additionalProperties, "Event")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TimeStamp")
		delete(additionalProperties, "UnavailableReplicas")
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

type NullableOprsDeployment struct {
	value *OprsDeployment
	isSet bool
}

func (v NullableOprsDeployment) Get() *OprsDeployment {
	return v.value
}

func (v *NullableOprsDeployment) Set(val *OprsDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableOprsDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableOprsDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOprsDeployment(val *OprsDeployment) *NullableOprsDeployment {
	return &NullableOprsDeployment{value: val, isSet: true}
}

func (v NullableOprsDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOprsDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
