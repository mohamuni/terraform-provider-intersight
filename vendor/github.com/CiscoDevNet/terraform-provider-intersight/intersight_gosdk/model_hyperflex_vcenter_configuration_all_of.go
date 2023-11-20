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

// HyperflexVcenterConfigurationAllOf Definition of the list of properties defined in 'hyperflex.VcenterConfiguration', excluding properties defined in parent classes.
type HyperflexVcenterConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The vCenter compute cluster identifier for the HyperFlex Cluster.
	ClusterId *string `json:"ClusterId,omitempty"`
	// The vCenter compute cluster name for the HyperFlex cluster.
	ClusterName *string `json:"ClusterName,omitempty"`
	// The identifier of the datacenter in vCenter that the HyperFlex cluster belongs to.
	DatacenterId *string `json:"DatacenterId,omitempty"`
	// The name of the datacenter in vCenter that the HyperFlex cluster belongs to.
	DatacenterName *string `json:"DatacenterName,omitempty"`
	// The URL of the vCenter for this HyperFlex Cluster.
	Url                  *string `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVcenterConfigurationAllOf HyperflexVcenterConfigurationAllOf

// NewHyperflexVcenterConfigurationAllOf instantiates a new HyperflexVcenterConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVcenterConfigurationAllOf(classId string, objectType string) *HyperflexVcenterConfigurationAllOf {
	this := HyperflexVcenterConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVcenterConfigurationAllOfWithDefaults instantiates a new HyperflexVcenterConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVcenterConfigurationAllOfWithDefaults() *HyperflexVcenterConfigurationAllOf {
	this := HyperflexVcenterConfigurationAllOf{}
	var classId string = "hyperflex.VcenterConfiguration"
	this.ClassId = classId
	var objectType string = "hyperflex.VcenterConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVcenterConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVcenterConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVcenterConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVcenterConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigurationAllOf) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigurationAllOf) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *HyperflexVcenterConfigurationAllOf) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigurationAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigurationAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HyperflexVcenterConfigurationAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigurationAllOf) GetDatacenterId() string {
	if o == nil || o.DatacenterId == nil {
		var ret string
		return ret
	}
	return *o.DatacenterId
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetDatacenterIdOk() (*string, bool) {
	if o == nil || o.DatacenterId == nil {
		return nil, false
	}
	return o.DatacenterId, true
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigurationAllOf) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given string and assigns it to the DatacenterId field.
func (o *HyperflexVcenterConfigurationAllOf) SetDatacenterId(v string) {
	o.DatacenterId = &v
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigurationAllOf) GetDatacenterName() string {
	if o == nil || o.DatacenterName == nil {
		var ret string
		return ret
	}
	return *o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetDatacenterNameOk() (*string, bool) {
	if o == nil || o.DatacenterName == nil {
		return nil, false
	}
	return o.DatacenterName, true
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigurationAllOf) HasDatacenterName() bool {
	if o != nil && o.DatacenterName != nil {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given string and assigns it to the DatacenterName field.
func (o *HyperflexVcenterConfigurationAllOf) SetDatacenterName(v string) {
	o.DatacenterName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigurationAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigurationAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigurationAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HyperflexVcenterConfigurationAllOf) SetUrl(v string) {
	o.Url = &v
}

func (o HyperflexVcenterConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterId != nil {
		toSerialize["ClusterId"] = o.ClusterId
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.DatacenterId != nil {
		toSerialize["DatacenterId"] = o.DatacenterId
	}
	if o.DatacenterName != nil {
		toSerialize["DatacenterName"] = o.DatacenterName
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVcenterConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVcenterConfigurationAllOf := _HyperflexVcenterConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVcenterConfigurationAllOf); err == nil {
		*o = HyperflexVcenterConfigurationAllOf(varHyperflexVcenterConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterId")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "DatacenterId")
		delete(additionalProperties, "DatacenterName")
		delete(additionalProperties, "Url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVcenterConfigurationAllOf struct {
	value *HyperflexVcenterConfigurationAllOf
	isSet bool
}

func (v NullableHyperflexVcenterConfigurationAllOf) Get() *HyperflexVcenterConfigurationAllOf {
	return v.value
}

func (v *NullableHyperflexVcenterConfigurationAllOf) Set(val *HyperflexVcenterConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVcenterConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVcenterConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVcenterConfigurationAllOf(val *HyperflexVcenterConfigurationAllOf) *NullableHyperflexVcenterConfigurationAllOf {
	return &NullableHyperflexVcenterConfigurationAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVcenterConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVcenterConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
