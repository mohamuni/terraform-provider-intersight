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

// checks if the OnpremResourceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnpremResourceInfo{}

// OnpremResourceInfo Common data structure used for capturing memory or storage usage information.
type OnpremResourceInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Available memory or storage in bytes.
	Avail *int64 `json:"Avail,omitempty"`
	// Name of the resource. In case of disk, it is the mount name of the disk.
	Name *string `json:"Name,omitempty"`
	// Total memory or storage in bytes.
	Total                *int64 `json:"Total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnpremResourceInfo OnpremResourceInfo

// NewOnpremResourceInfo instantiates a new OnpremResourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremResourceInfo(classId string, objectType string) *OnpremResourceInfo {
	this := OnpremResourceInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOnpremResourceInfoWithDefaults instantiates a new OnpremResourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremResourceInfoWithDefaults() *OnpremResourceInfo {
	this := OnpremResourceInfo{}
	var classId string = "onprem.ResourceInfo"
	this.ClassId = classId
	var objectType string = "onprem.ResourceInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OnpremResourceInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OnpremResourceInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OnpremResourceInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "onprem.ResourceInfo" of the ClassId field.
func (o *OnpremResourceInfo) GetDefaultClassId() interface{} {
	return "onprem.ResourceInfo"
}

// GetObjectType returns the ObjectType field value
func (o *OnpremResourceInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OnpremResourceInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OnpremResourceInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "onprem.ResourceInfo" of the ObjectType field.
func (o *OnpremResourceInfo) GetDefaultObjectType() interface{} {
	return "onprem.ResourceInfo"
}

// GetAvail returns the Avail field value if set, zero value otherwise.
func (o *OnpremResourceInfo) GetAvail() int64 {
	if o == nil || IsNil(o.Avail) {
		var ret int64
		return ret
	}
	return *o.Avail
}

// GetAvailOk returns a tuple with the Avail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremResourceInfo) GetAvailOk() (*int64, bool) {
	if o == nil || IsNil(o.Avail) {
		return nil, false
	}
	return o.Avail, true
}

// HasAvail returns a boolean if a field has been set.
func (o *OnpremResourceInfo) HasAvail() bool {
	if o != nil && !IsNil(o.Avail) {
		return true
	}

	return false
}

// SetAvail gets a reference to the given int64 and assigns it to the Avail field.
func (o *OnpremResourceInfo) SetAvail(v int64) {
	o.Avail = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OnpremResourceInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremResourceInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OnpremResourceInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OnpremResourceInfo) SetName(v string) {
	o.Name = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OnpremResourceInfo) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremResourceInfo) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OnpremResourceInfo) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *OnpremResourceInfo) SetTotal(v int64) {
	o.Total = &v
}

func (o OnpremResourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnpremResourceInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Avail) {
		toSerialize["Avail"] = o.Avail
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Total) {
		toSerialize["Total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnpremResourceInfo) UnmarshalJSON(data []byte) (err error) {
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
	type OnpremResourceInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Available memory or storage in bytes.
		Avail *int64 `json:"Avail,omitempty"`
		// Name of the resource. In case of disk, it is the mount name of the disk.
		Name *string `json:"Name,omitempty"`
		// Total memory or storage in bytes.
		Total *int64 `json:"Total,omitempty"`
	}

	varOnpremResourceInfoWithoutEmbeddedStruct := OnpremResourceInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOnpremResourceInfoWithoutEmbeddedStruct)
	if err == nil {
		varOnpremResourceInfo := _OnpremResourceInfo{}
		varOnpremResourceInfo.ClassId = varOnpremResourceInfoWithoutEmbeddedStruct.ClassId
		varOnpremResourceInfo.ObjectType = varOnpremResourceInfoWithoutEmbeddedStruct.ObjectType
		varOnpremResourceInfo.Avail = varOnpremResourceInfoWithoutEmbeddedStruct.Avail
		varOnpremResourceInfo.Name = varOnpremResourceInfoWithoutEmbeddedStruct.Name
		varOnpremResourceInfo.Total = varOnpremResourceInfoWithoutEmbeddedStruct.Total
		*o = OnpremResourceInfo(varOnpremResourceInfo)
	} else {
		return err
	}

	varOnpremResourceInfo := _OnpremResourceInfo{}

	err = json.Unmarshal(data, &varOnpremResourceInfo)
	if err == nil {
		o.MoBaseComplexType = varOnpremResourceInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Avail")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Total")

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

type NullableOnpremResourceInfo struct {
	value *OnpremResourceInfo
	isSet bool
}

func (v NullableOnpremResourceInfo) Get() *OnpremResourceInfo {
	return v.value
}

func (v *NullableOnpremResourceInfo) Set(val *OnpremResourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremResourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremResourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremResourceInfo(val *OnpremResourceInfo) *NullableOnpremResourceInfo {
	return &NullableOnpremResourceInfo{value: val, isSet: true}
}

func (v NullableOnpremResourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremResourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
