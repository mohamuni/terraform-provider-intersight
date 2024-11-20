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

// checks if the PoolAbstractPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolAbstractPool{}

// PoolAbstractPool Pool represents a  collection of identifiers that can be allocated to server profiles.
type PoolAbstractPool struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Number of IDs that are currently assigned (in use).
	Assigned *int64 `json:"Assigned,omitempty"`
	// Assignment order decides the order in which the next identifier is allocated. * `sequential` - Identifiers are assigned in a sequential order. * `default` - Assignment order is decided by the system.
	AssignmentOrder *string `json:"AssignmentOrder,omitempty"`
	// Number of IDs that are currently reserved (and not in use).
	Reserved *int64 `json:"Reserved,omitempty"`
	// Total number of identifiers in this pool.
	Size                 *int64 `json:"Size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractPool PoolAbstractPool

// NewPoolAbstractPool instantiates a new PoolAbstractPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractPool(classId string, objectType string) *PoolAbstractPool {
	this := PoolAbstractPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewPoolAbstractPoolWithDefaults instantiates a new PoolAbstractPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractPoolWithDefaults() *PoolAbstractPool {
	this := PoolAbstractPool{}
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetAssigned() int64 {
	if o == nil || IsNil(o.Assigned) {
		var ret int64
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetAssignedOk() (*int64, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int64 and assigns it to the Assigned field.
func (o *PoolAbstractPool) SetAssigned(v int64) {
	o.Assigned = &v
}

// GetAssignmentOrder returns the AssignmentOrder field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetAssignmentOrder() string {
	if o == nil || IsNil(o.AssignmentOrder) {
		var ret string
		return ret
	}
	return *o.AssignmentOrder
}

// GetAssignmentOrderOk returns a tuple with the AssignmentOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetAssignmentOrderOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentOrder) {
		return nil, false
	}
	return o.AssignmentOrder, true
}

// HasAssignmentOrder returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasAssignmentOrder() bool {
	if o != nil && !IsNil(o.AssignmentOrder) {
		return true
	}

	return false
}

// SetAssignmentOrder gets a reference to the given string and assigns it to the AssignmentOrder field.
func (o *PoolAbstractPool) SetAssignmentOrder(v string) {
	o.AssignmentOrder = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetReserved() int64 {
	if o == nil || IsNil(o.Reserved) {
		var ret int64
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetReservedOk() (*int64, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int64 and assigns it to the Reserved field.
func (o *PoolAbstractPool) SetReserved(v int64) {
	o.Reserved = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *PoolAbstractPool) SetSize(v int64) {
	o.Size = &v
}

func (o PoolAbstractPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolAbstractPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Assigned) {
		toSerialize["Assigned"] = o.Assigned
	}
	if !IsNil(o.AssignmentOrder) {
		toSerialize["AssignmentOrder"] = o.AssignmentOrder
	}
	if !IsNil(o.Reserved) {
		toSerialize["Reserved"] = o.Reserved
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PoolAbstractPool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
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
	type PoolAbstractPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Number of IDs that are currently assigned (in use).
		Assigned *int64 `json:"Assigned,omitempty"`
		// Assignment order decides the order in which the next identifier is allocated. * `sequential` - Identifiers are assigned in a sequential order. * `default` - Assignment order is decided by the system.
		AssignmentOrder *string `json:"AssignmentOrder,omitempty"`
		// Number of IDs that are currently reserved (and not in use).
		Reserved *int64 `json:"Reserved,omitempty"`
		// Total number of identifiers in this pool.
		Size *int64 `json:"Size,omitempty"`
	}

	varPoolAbstractPoolWithoutEmbeddedStruct := PoolAbstractPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPoolAbstractPoolWithoutEmbeddedStruct)
	if err == nil {
		varPoolAbstractPool := _PoolAbstractPool{}
		varPoolAbstractPool.ClassId = varPoolAbstractPoolWithoutEmbeddedStruct.ClassId
		varPoolAbstractPool.ObjectType = varPoolAbstractPoolWithoutEmbeddedStruct.ObjectType
		varPoolAbstractPool.Assigned = varPoolAbstractPoolWithoutEmbeddedStruct.Assigned
		varPoolAbstractPool.AssignmentOrder = varPoolAbstractPoolWithoutEmbeddedStruct.AssignmentOrder
		varPoolAbstractPool.Reserved = varPoolAbstractPoolWithoutEmbeddedStruct.Reserved
		varPoolAbstractPool.Size = varPoolAbstractPoolWithoutEmbeddedStruct.Size
		*o = PoolAbstractPool(varPoolAbstractPool)
	} else {
		return err
	}

	varPoolAbstractPool := _PoolAbstractPool{}

	err = json.Unmarshal(data, &varPoolAbstractPool)
	if err == nil {
		o.PolicyAbstractPolicy = varPoolAbstractPool.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Assigned")
		delete(additionalProperties, "AssignmentOrder")
		delete(additionalProperties, "Reserved")
		delete(additionalProperties, "Size")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullablePoolAbstractPool struct {
	value *PoolAbstractPool
	isSet bool
}

func (v NullablePoolAbstractPool) Get() *PoolAbstractPool {
	return v.value
}

func (v *NullablePoolAbstractPool) Set(val *PoolAbstractPool) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractPool) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractPool(val *PoolAbstractPool) *NullablePoolAbstractPool {
	return &NullablePoolAbstractPool{value: val, isSet: true}
}

func (v NullablePoolAbstractPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
