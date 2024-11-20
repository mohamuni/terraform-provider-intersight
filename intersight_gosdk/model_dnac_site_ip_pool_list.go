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

// checks if the DnacSiteIpPoolList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacSiteIpPoolList{}

// DnacSiteIpPoolList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type DnacSiteIpPoolList struct {
	MoBaseResponse
	// The total number of 'dnac.SiteIpPool' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'dnac.SiteIpPool' resources matching the request.
	Results              []DnacSiteIpPool `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacSiteIpPoolList DnacSiteIpPoolList

// NewDnacSiteIpPoolList instantiates a new DnacSiteIpPoolList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacSiteIpPoolList(objectType string) *DnacSiteIpPoolList {
	this := DnacSiteIpPoolList{}
	this.ObjectType = objectType
	return &this
}

// NewDnacSiteIpPoolListWithDefaults instantiates a new DnacSiteIpPoolList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacSiteIpPoolListWithDefaults() *DnacSiteIpPoolList {
	this := DnacSiteIpPoolList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DnacSiteIpPoolList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacSiteIpPoolList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DnacSiteIpPoolList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DnacSiteIpPoolList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnacSiteIpPoolList) GetResults() []DnacSiteIpPool {
	if o == nil {
		var ret []DnacSiteIpPool
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnacSiteIpPoolList) GetResultsOk() ([]DnacSiteIpPool, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DnacSiteIpPoolList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DnacSiteIpPool and assigns it to the Results field.
func (o *DnacSiteIpPoolList) SetResults(v []DnacSiteIpPool) {
	o.Results = v
}

func (o DnacSiteIpPoolList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacSiteIpPoolList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return map[string]interface{}{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return map[string]interface{}{}, errMoBaseResponse
	}
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacSiteIpPoolList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type DnacSiteIpPoolListWithoutEmbeddedStruct struct {
		// The total number of 'dnac.SiteIpPool' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'dnac.SiteIpPool' resources matching the request.
		Results []DnacSiteIpPool `json:"Results,omitempty"`
	}

	varDnacSiteIpPoolListWithoutEmbeddedStruct := DnacSiteIpPoolListWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacSiteIpPoolListWithoutEmbeddedStruct)
	if err == nil {
		varDnacSiteIpPoolList := _DnacSiteIpPoolList{}
		varDnacSiteIpPoolList.Count = varDnacSiteIpPoolListWithoutEmbeddedStruct.Count
		varDnacSiteIpPoolList.Results = varDnacSiteIpPoolListWithoutEmbeddedStruct.Results
		*o = DnacSiteIpPoolList(varDnacSiteIpPoolList)
	} else {
		return err
	}

	varDnacSiteIpPoolList := _DnacSiteIpPoolList{}

	err = json.Unmarshal(data, &varDnacSiteIpPoolList)
	if err == nil {
		o.MoBaseResponse = varDnacSiteIpPoolList.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableDnacSiteIpPoolList struct {
	value *DnacSiteIpPoolList
	isSet bool
}

func (v NullableDnacSiteIpPoolList) Get() *DnacSiteIpPoolList {
	return v.value
}

func (v *NullableDnacSiteIpPoolList) Set(val *DnacSiteIpPoolList) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacSiteIpPoolList) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacSiteIpPoolList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacSiteIpPoolList(val *DnacSiteIpPoolList) *NullableDnacSiteIpPoolList {
	return &NullableDnacSiteIpPoolList{value: val, isSet: true}
}

func (v NullableDnacSiteIpPoolList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacSiteIpPoolList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
