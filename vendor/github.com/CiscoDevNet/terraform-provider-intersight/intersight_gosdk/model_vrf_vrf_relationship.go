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
)

// VrfVrfRelationship - A relationship to the 'vrf.Vrf' resource, or the expanded 'vrf.Vrf' resource, or the 'null' value.
type VrfVrfRelationship struct {
	MoMoRef *MoMoRef
	VrfVrf  *VrfVrf
}

// MoMoRefAsVrfVrfRelationship is a convenience function that returns MoMoRef wrapped in VrfVrfRelationship
func MoMoRefAsVrfVrfRelationship(v *MoMoRef) VrfVrfRelationship {
	return VrfVrfRelationship{
		MoMoRef: v,
	}
}

// VrfVrfAsVrfVrfRelationship is a convenience function that returns VrfVrf wrapped in VrfVrfRelationship
func VrfVrfAsVrfVrfRelationship(v *VrfVrf) VrfVrfRelationship {
	return VrfVrfRelationship{
		VrfVrf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VrfVrfRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal VrfVrfRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vrf.Vrf'
	if jsonDict["ClassId"] == "vrf.Vrf" {
		// try to unmarshal JSON data into VrfVrf
		err = json.Unmarshal(data, &dst.VrfVrf)
		if err == nil {
			return nil // data stored in dst.VrfVrf, return on the first match
		} else {
			dst.VrfVrf = nil
			return fmt.Errorf("failed to unmarshal VrfVrfRelationship as VrfVrf: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VrfVrfRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VrfVrf != nil {
		return json.Marshal(&src.VrfVrf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VrfVrfRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VrfVrf != nil {
		return obj.VrfVrf
	}

	// all schemas are nil
	return nil
}

type NullableVrfVrfRelationship struct {
	value *VrfVrfRelationship
	isSet bool
}

func (v NullableVrfVrfRelationship) Get() *VrfVrfRelationship {
	return v.value
}

func (v *NullableVrfVrfRelationship) Set(val *VrfVrfRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVrfRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVrfRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVrfRelationship(val *VrfVrfRelationship) *NullableVrfVrfRelationship {
	return &NullableVrfVrfRelationship{value: val, isSet: true}
}

func (v NullableVrfVrfRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVrfRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
