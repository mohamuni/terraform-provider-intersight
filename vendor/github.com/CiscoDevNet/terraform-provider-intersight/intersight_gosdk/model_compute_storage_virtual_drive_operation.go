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
	"reflect"
	"strings"
)

// ComputeStorageVirtualDriveOperation The operation that can be performed on the Storage Virtual Drives on the servers.
type ComputeStorageVirtualDriveOperation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrative actions that can be performed on the Storage Virtual Drives. * `None` - No action on the selected Storage virtual Drives. * `Delete` - Delete action on the selected Storage Virtual Drives.
	AdminAction *string `json:"AdminAction,omitempty"`
	// Storage Controller Id of the storage Virtual Drives of the server.
	ControllerId         *string                      `json:"ControllerId,omitempty"`
	VirtualDrives        []ComputeStorageVirtualDrive `json:"VirtualDrives,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeStorageVirtualDriveOperation ComputeStorageVirtualDriveOperation

// NewComputeStorageVirtualDriveOperation instantiates a new ComputeStorageVirtualDriveOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeStorageVirtualDriveOperation(classId string, objectType string) *ComputeStorageVirtualDriveOperation {
	this := ComputeStorageVirtualDriveOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// NewComputeStorageVirtualDriveOperationWithDefaults instantiates a new ComputeStorageVirtualDriveOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeStorageVirtualDriveOperationWithDefaults() *ComputeStorageVirtualDriveOperation {
	this := ComputeStorageVirtualDriveOperation{}
	var classId string = "compute.StorageVirtualDriveOperation"
	this.ClassId = classId
	var objectType string = "compute.StorageVirtualDriveOperation"
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeStorageVirtualDriveOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDriveOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeStorageVirtualDriveOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeStorageVirtualDriveOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDriveOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeStorageVirtualDriveOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *ComputeStorageVirtualDriveOperation) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDriveOperation) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *ComputeStorageVirtualDriveOperation) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *ComputeStorageVirtualDriveOperation) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *ComputeStorageVirtualDriveOperation) GetControllerId() string {
	if o == nil || o.ControllerId == nil {
		var ret string
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDriveOperation) GetControllerIdOk() (*string, bool) {
	if o == nil || o.ControllerId == nil {
		return nil, false
	}
	return o.ControllerId, true
}

// HasControllerId returns a boolean if a field has been set.
func (o *ComputeStorageVirtualDriveOperation) HasControllerId() bool {
	if o != nil && o.ControllerId != nil {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given string and assigns it to the ControllerId field.
func (o *ComputeStorageVirtualDriveOperation) SetControllerId(v string) {
	o.ControllerId = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeStorageVirtualDriveOperation) GetVirtualDrives() []ComputeStorageVirtualDrive {
	if o == nil {
		var ret []ComputeStorageVirtualDrive
		return ret
	}
	return o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeStorageVirtualDriveOperation) GetVirtualDrivesOk() ([]ComputeStorageVirtualDrive, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *ComputeStorageVirtualDriveOperation) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []ComputeStorageVirtualDrive and assigns it to the VirtualDrives field.
func (o *ComputeStorageVirtualDriveOperation) SetVirtualDrives(v []ComputeStorageVirtualDrive) {
	o.VirtualDrives = v
}

func (o ComputeStorageVirtualDriveOperation) MarshalJSON() ([]byte, error) {
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
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.ControllerId != nil {
		toSerialize["ControllerId"] = o.ControllerId
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeStorageVirtualDriveOperation) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeStorageVirtualDriveOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Administrative actions that can be performed on the Storage Virtual Drives. * `None` - No action on the selected Storage virtual Drives. * `Delete` - Delete action on the selected Storage Virtual Drives.
		AdminAction *string `json:"AdminAction,omitempty"`
		// Storage Controller Id of the storage Virtual Drives of the server.
		ControllerId  *string                      `json:"ControllerId,omitempty"`
		VirtualDrives []ComputeStorageVirtualDrive `json:"VirtualDrives,omitempty"`
	}

	varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct := ComputeStorageVirtualDriveOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct)
	if err == nil {
		varComputeStorageVirtualDriveOperation := _ComputeStorageVirtualDriveOperation{}
		varComputeStorageVirtualDriveOperation.ClassId = varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct.ClassId
		varComputeStorageVirtualDriveOperation.ObjectType = varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct.ObjectType
		varComputeStorageVirtualDriveOperation.AdminAction = varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct.AdminAction
		varComputeStorageVirtualDriveOperation.ControllerId = varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct.ControllerId
		varComputeStorageVirtualDriveOperation.VirtualDrives = varComputeStorageVirtualDriveOperationWithoutEmbeddedStruct.VirtualDrives
		*o = ComputeStorageVirtualDriveOperation(varComputeStorageVirtualDriveOperation)
	} else {
		return err
	}

	varComputeStorageVirtualDriveOperation := _ComputeStorageVirtualDriveOperation{}

	err = json.Unmarshal(bytes, &varComputeStorageVirtualDriveOperation)
	if err == nil {
		o.MoBaseComplexType = varComputeStorageVirtualDriveOperation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminAction")
		delete(additionalProperties, "ControllerId")
		delete(additionalProperties, "VirtualDrives")

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

type NullableComputeStorageVirtualDriveOperation struct {
	value *ComputeStorageVirtualDriveOperation
	isSet bool
}

func (v NullableComputeStorageVirtualDriveOperation) Get() *ComputeStorageVirtualDriveOperation {
	return v.value
}

func (v *NullableComputeStorageVirtualDriveOperation) Set(val *ComputeStorageVirtualDriveOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeStorageVirtualDriveOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeStorageVirtualDriveOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeStorageVirtualDriveOperation(val *ComputeStorageVirtualDriveOperation) *NullableComputeStorageVirtualDriveOperation {
	return &NullableComputeStorageVirtualDriveOperation{value: val, isSet: true}
}

func (v NullableComputeStorageVirtualDriveOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeStorageVirtualDriveOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
