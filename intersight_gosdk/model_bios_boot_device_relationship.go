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
)

// BiosBootDeviceRelationship - A relationship to the 'bios.BootDevice' resource, or the expanded 'bios.BootDevice' resource, or the 'null' value.
type BiosBootDeviceRelationship struct {
	BiosBootDevice *BiosBootDevice
	MoMoRef        *MoMoRef
}

// BiosBootDeviceAsBiosBootDeviceRelationship is a convenience function that returns BiosBootDevice wrapped in BiosBootDeviceRelationship
func BiosBootDeviceAsBiosBootDeviceRelationship(v *BiosBootDevice) BiosBootDeviceRelationship {
	return BiosBootDeviceRelationship{
		BiosBootDevice: v,
	}
}

// MoMoRefAsBiosBootDeviceRelationship is a convenience function that returns MoMoRef wrapped in BiosBootDeviceRelationship
func MoMoRefAsBiosBootDeviceRelationship(v *MoMoRef) BiosBootDeviceRelationship {
	return BiosBootDeviceRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BiosBootDeviceRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'bios.BootDevice'
	if jsonDict["ClassId"] == "bios.BootDevice" {
		// try to unmarshal JSON data into BiosBootDevice
		err = json.Unmarshal(data, &dst.BiosBootDevice)
		if err == nil {
			return nil // data stored in dst.BiosBootDevice, return on the first match
		} else {
			dst.BiosBootDevice = nil
			return fmt.Errorf("failed to unmarshal BiosBootDeviceRelationship as BiosBootDevice: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal BiosBootDeviceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BiosBootDeviceRelationship) MarshalJSON() ([]byte, error) {
	if src.BiosBootDevice != nil {
		return json.Marshal(&src.BiosBootDevice)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BiosBootDeviceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BiosBootDevice != nil {
		return obj.BiosBootDevice
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableBiosBootDeviceRelationship struct {
	value *BiosBootDeviceRelationship
	isSet bool
}

func (v NullableBiosBootDeviceRelationship) Get() *BiosBootDeviceRelationship {
	return v.value
}

func (v *NullableBiosBootDeviceRelationship) Set(val *BiosBootDeviceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosBootDeviceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosBootDeviceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosBootDeviceRelationship(val *BiosBootDeviceRelationship) *NullableBiosBootDeviceRelationship {
	return &NullableBiosBootDeviceRelationship{value: val, isSet: true}
}

func (v NullableBiosBootDeviceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosBootDeviceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
