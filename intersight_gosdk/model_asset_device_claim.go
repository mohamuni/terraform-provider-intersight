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

// AssetDeviceClaim DeviceClaim captures the intent to claim a device to an Intersight account. A device can be unclaimed by performing a DELETE on a DeviceClaim instance. When performing a claim, a secret passphrase must be obtained from the device connector UI/API by a sufficiently privileged user. The passphrase is timebound and proves that the user currently has privileged administrative access to the device being claimed.
type AssetDeviceClaim struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Obtained from the device connector management UI or API (REST endpoint '/connector/SecurityTokens').
	SecurityToken *string `json:"SecurityToken,omitempty"`
	// Obtained from the device connector management UI or API (REST endpoint '/connector/DeviceIdentifiers').
	SerialNumber         *string                              `json:"SerialNumber,omitempty"`
	Account              *IamAccountRelationship              `json:"Account,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Reservation          *ResourceReservationRelationship     `json:"Reservation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceClaim AssetDeviceClaim

// NewAssetDeviceClaim instantiates a new AssetDeviceClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceClaim(classId string, objectType string) *AssetDeviceClaim {
	this := AssetDeviceClaim{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceClaimWithDefaults instantiates a new AssetDeviceClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceClaimWithDefaults() *AssetDeviceClaim {
	this := AssetDeviceClaim{}
	var classId string = "asset.DeviceClaim"
	this.ClassId = classId
	var objectType string = "asset.DeviceClaim"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceClaim) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceClaim) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceClaim) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceClaim) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetSecurityToken() string {
	if o == nil || o.SecurityToken == nil {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetSecurityTokenOk() (*string, bool) {
	if o == nil || o.SecurityToken == nil {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasSecurityToken() bool {
	if o != nil && o.SecurityToken != nil {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *AssetDeviceClaim) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *AssetDeviceClaim) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetDeviceClaim) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *AssetDeviceClaim) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetReservation() ResourceReservationRelationship {
	if o == nil || o.Reservation == nil {
		var ret ResourceReservationRelationship
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetReservationOk() (*ResourceReservationRelationship, bool) {
	if o == nil || o.Reservation == nil {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasReservation() bool {
	if o != nil && o.Reservation != nil {
		return true
	}

	return false
}

// SetReservation gets a reference to the given ResourceReservationRelationship and assigns it to the Reservation field.
func (o *AssetDeviceClaim) SetReservation(v ResourceReservationRelationship) {
	o.Reservation = &v
}

func (o AssetDeviceClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SecurityToken != nil {
		toSerialize["SecurityToken"] = o.SecurityToken
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Reservation != nil {
		toSerialize["Reservation"] = o.Reservation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceClaim) UnmarshalJSON(bytes []byte) (err error) {
	type AssetDeviceClaimWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Obtained from the device connector management UI or API (REST endpoint '/connector/SecurityTokens').
		SecurityToken *string `json:"SecurityToken,omitempty"`
		// Obtained from the device connector management UI or API (REST endpoint '/connector/DeviceIdentifiers').
		SerialNumber *string                              `json:"SerialNumber,omitempty"`
		Account      *IamAccountRelationship              `json:"Account,omitempty"`
		Device       *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Reservation  *ResourceReservationRelationship     `json:"Reservation,omitempty"`
	}

	varAssetDeviceClaimWithoutEmbeddedStruct := AssetDeviceClaimWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetDeviceClaimWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeviceClaim := _AssetDeviceClaim{}
		varAssetDeviceClaim.ClassId = varAssetDeviceClaimWithoutEmbeddedStruct.ClassId
		varAssetDeviceClaim.ObjectType = varAssetDeviceClaimWithoutEmbeddedStruct.ObjectType
		varAssetDeviceClaim.SecurityToken = varAssetDeviceClaimWithoutEmbeddedStruct.SecurityToken
		varAssetDeviceClaim.SerialNumber = varAssetDeviceClaimWithoutEmbeddedStruct.SerialNumber
		varAssetDeviceClaim.Account = varAssetDeviceClaimWithoutEmbeddedStruct.Account
		varAssetDeviceClaim.Device = varAssetDeviceClaimWithoutEmbeddedStruct.Device
		varAssetDeviceClaim.Reservation = varAssetDeviceClaimWithoutEmbeddedStruct.Reservation
		*o = AssetDeviceClaim(varAssetDeviceClaim)
	} else {
		return err
	}

	varAssetDeviceClaim := _AssetDeviceClaim{}

	err = json.Unmarshal(bytes, &varAssetDeviceClaim)
	if err == nil {
		o.MoBaseMo = varAssetDeviceClaim.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SecurityToken")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Reservation")

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

type NullableAssetDeviceClaim struct {
	value *AssetDeviceClaim
	isSet bool
}

func (v NullableAssetDeviceClaim) Get() *AssetDeviceClaim {
	return v.value
}

func (v *NullableAssetDeviceClaim) Set(val *AssetDeviceClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceClaim(val *AssetDeviceClaim) *NullableAssetDeviceClaim {
	return &NullableAssetDeviceClaim{value: val, isSet: true}
}

func (v NullableAssetDeviceClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
