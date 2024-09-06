/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18369
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

// checks if the NotificationDistributableMoCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationDistributableMoCondition{}

// NotificationDistributableMoCondition This condition type gives the ability to subscribe to the software.HyperflexBundleDistributable or firmware.Distributable or appliance.ApplianceDistributable or software.HciBundleDistributable ObjectType.
type NotificationDistributableMoCondition struct {
	NotificationAbstractMoCondition
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Image type for which user wants to create notification. The list of valid values for Image type of notification.DistributableMoCondition is software.HyperflexBundleDistributable, software.ApplianceDistributableIntelligence, firmware.Distributable, software.HciBundleDistributable. * `None` - Default value for DistributableObjectType enum. * `software.ApplianceDistributableIntelligence` - Object type for Appliance Intelligence updates related distributable. * `firmware.Distributable` - Object type for firmware related distributable. * `software.HyperflexBundleDistributable` - Object type for HyperFlex related distributable. * `software.ApplianceDistributableConnected` - Object type for Appliance related distributable. By Choosing this, user shows intent to use the downloaded image for Connected appliance.  * `software.ApplianceDistributablePrivate` - Object type for Appliance related distributable. By Choosing this, user shows intent to use the downloaded image for Private appliance.  * `software.UcsdBundleDistributable` - Object type for UCSD related distributable. * `software.HciBundleDistributable` - Object type for HCI Bundle related distributable.
	ImageType *string `json:"ImageType,omitempty"`
	// Odata filter string managed internally. It is built with specific ObjectType properties.
	OdataFilter          *string `json:"OdataFilter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationDistributableMoCondition NotificationDistributableMoCondition

// NewNotificationDistributableMoCondition instantiates a new NotificationDistributableMoCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationDistributableMoCondition(classId string, objectType string) *NotificationDistributableMoCondition {
	this := NotificationDistributableMoCondition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var imageType string = "None"
	this.ImageType = &imageType
	return &this
}

// NewNotificationDistributableMoConditionWithDefaults instantiates a new NotificationDistributableMoCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDistributableMoConditionWithDefaults() *NotificationDistributableMoCondition {
	this := NotificationDistributableMoCondition{}
	var classId string = "notification.DistributableMoCondition"
	this.ClassId = classId
	var objectType string = "notification.DistributableMoCondition"
	this.ObjectType = objectType
	var imageType string = "None"
	this.ImageType = &imageType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationDistributableMoCondition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationDistributableMoCondition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationDistributableMoCondition) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "notification.DistributableMoCondition" of the ClassId field.
func (o *NotificationDistributableMoCondition) GetDefaultClassId() interface{} {
	return "notification.DistributableMoCondition"
}

// GetObjectType returns the ObjectType field value
func (o *NotificationDistributableMoCondition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationDistributableMoCondition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationDistributableMoCondition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "notification.DistributableMoCondition" of the ObjectType field.
func (o *NotificationDistributableMoCondition) GetDefaultObjectType() interface{} {
	return "notification.DistributableMoCondition"
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *NotificationDistributableMoCondition) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDistributableMoCondition) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *NotificationDistributableMoCondition) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *NotificationDistributableMoCondition) SetImageType(v string) {
	o.ImageType = &v
}

// GetOdataFilter returns the OdataFilter field value if set, zero value otherwise.
func (o *NotificationDistributableMoCondition) GetOdataFilter() string {
	if o == nil || IsNil(o.OdataFilter) {
		var ret string
		return ret
	}
	return *o.OdataFilter
}

// GetOdataFilterOk returns a tuple with the OdataFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDistributableMoCondition) GetOdataFilterOk() (*string, bool) {
	if o == nil || IsNil(o.OdataFilter) {
		return nil, false
	}
	return o.OdataFilter, true
}

// HasOdataFilter returns a boolean if a field has been set.
func (o *NotificationDistributableMoCondition) HasOdataFilter() bool {
	if o != nil && !IsNil(o.OdataFilter) {
		return true
	}

	return false
}

// SetOdataFilter gets a reference to the given string and assigns it to the OdataFilter field.
func (o *NotificationDistributableMoCondition) SetOdataFilter(v string) {
	o.OdataFilter = &v
}

func (o NotificationDistributableMoCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationDistributableMoCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationAbstractMoCondition, errNotificationAbstractMoCondition := json.Marshal(o.NotificationAbstractMoCondition)
	if errNotificationAbstractMoCondition != nil {
		return map[string]interface{}{}, errNotificationAbstractMoCondition
	}
	errNotificationAbstractMoCondition = json.Unmarshal([]byte(serializedNotificationAbstractMoCondition), &toSerialize)
	if errNotificationAbstractMoCondition != nil {
		return map[string]interface{}{}, errNotificationAbstractMoCondition
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ImageType) {
		toSerialize["ImageType"] = o.ImageType
	}
	if !IsNil(o.OdataFilter) {
		toSerialize["OdataFilter"] = o.OdataFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationDistributableMoCondition) UnmarshalJSON(data []byte) (err error) {
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
	type NotificationDistributableMoConditionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Image type for which user wants to create notification. The list of valid values for Image type of notification.DistributableMoCondition is software.HyperflexBundleDistributable, software.ApplianceDistributableIntelligence, firmware.Distributable, software.HciBundleDistributable. * `None` - Default value for DistributableObjectType enum. * `software.ApplianceDistributableIntelligence` - Object type for Appliance Intelligence updates related distributable. * `firmware.Distributable` - Object type for firmware related distributable. * `software.HyperflexBundleDistributable` - Object type for HyperFlex related distributable. * `software.ApplianceDistributableConnected` - Object type for Appliance related distributable. By Choosing this, user shows intent to use the downloaded image for Connected appliance.  * `software.ApplianceDistributablePrivate` - Object type for Appliance related distributable. By Choosing this, user shows intent to use the downloaded image for Private appliance.  * `software.UcsdBundleDistributable` - Object type for UCSD related distributable. * `software.HciBundleDistributable` - Object type for HCI Bundle related distributable.
		ImageType *string `json:"ImageType,omitempty"`
		// Odata filter string managed internally. It is built with specific ObjectType properties.
		OdataFilter *string `json:"OdataFilter,omitempty"`
	}

	varNotificationDistributableMoConditionWithoutEmbeddedStruct := NotificationDistributableMoConditionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNotificationDistributableMoConditionWithoutEmbeddedStruct)
	if err == nil {
		varNotificationDistributableMoCondition := _NotificationDistributableMoCondition{}
		varNotificationDistributableMoCondition.ClassId = varNotificationDistributableMoConditionWithoutEmbeddedStruct.ClassId
		varNotificationDistributableMoCondition.ObjectType = varNotificationDistributableMoConditionWithoutEmbeddedStruct.ObjectType
		varNotificationDistributableMoCondition.ImageType = varNotificationDistributableMoConditionWithoutEmbeddedStruct.ImageType
		varNotificationDistributableMoCondition.OdataFilter = varNotificationDistributableMoConditionWithoutEmbeddedStruct.OdataFilter
		*o = NotificationDistributableMoCondition(varNotificationDistributableMoCondition)
	} else {
		return err
	}

	varNotificationDistributableMoCondition := _NotificationDistributableMoCondition{}

	err = json.Unmarshal(data, &varNotificationDistributableMoCondition)
	if err == nil {
		o.NotificationAbstractMoCondition = varNotificationDistributableMoCondition.NotificationAbstractMoCondition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "OdataFilter")

		// remove fields from embedded structs
		reflectNotificationAbstractMoCondition := reflect.ValueOf(o.NotificationAbstractMoCondition)
		for i := 0; i < reflectNotificationAbstractMoCondition.Type().NumField(); i++ {
			t := reflectNotificationAbstractMoCondition.Type().Field(i)

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

type NullableNotificationDistributableMoCondition struct {
	value *NotificationDistributableMoCondition
	isSet bool
}

func (v NullableNotificationDistributableMoCondition) Get() *NotificationDistributableMoCondition {
	return v.value
}

func (v *NullableNotificationDistributableMoCondition) Set(val *NotificationDistributableMoCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDistributableMoCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDistributableMoCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDistributableMoCondition(val *NotificationDistributableMoCondition) *NullableNotificationDistributableMoCondition {
	return &NullableNotificationDistributableMoCondition{value: val, isSet: true}
}

func (v NullableNotificationDistributableMoCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDistributableMoCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}