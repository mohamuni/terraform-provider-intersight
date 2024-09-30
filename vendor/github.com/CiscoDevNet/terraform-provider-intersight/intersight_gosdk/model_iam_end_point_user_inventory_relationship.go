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
)

// IamEndPointUserInventoryRelationship - A relationship to the 'iam.EndPointUserInventory' resource, or the expanded 'iam.EndPointUserInventory' resource, or the 'null' value.
type IamEndPointUserInventoryRelationship struct {
	IamEndPointUserInventory *IamEndPointUserInventory
	MoMoRef                  *MoMoRef
}

// IamEndPointUserInventoryAsIamEndPointUserInventoryRelationship is a convenience function that returns IamEndPointUserInventory wrapped in IamEndPointUserInventoryRelationship
func IamEndPointUserInventoryAsIamEndPointUserInventoryRelationship(v *IamEndPointUserInventory) IamEndPointUserInventoryRelationship {
	return IamEndPointUserInventoryRelationship{
		IamEndPointUserInventory: v,
	}
}

// MoMoRefAsIamEndPointUserInventoryRelationship is a convenience function that returns MoMoRef wrapped in IamEndPointUserInventoryRelationship
func MoMoRefAsIamEndPointUserInventoryRelationship(v *MoMoRef) IamEndPointUserInventoryRelationship {
	return IamEndPointUserInventoryRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamEndPointUserInventoryRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.EndPointUserInventory'
	if jsonDict["ClassId"] == "iam.EndPointUserInventory" {
		// try to unmarshal JSON data into IamEndPointUserInventory
		err = json.Unmarshal(data, &dst.IamEndPointUserInventory)
		if err == nil {
			return nil // data stored in dst.IamEndPointUserInventory, return on the first match
		} else {
			dst.IamEndPointUserInventory = nil
			return fmt.Errorf("failed to unmarshal IamEndPointUserInventoryRelationship as IamEndPointUserInventory: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamEndPointUserInventoryRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamEndPointUserInventoryRelationship) MarshalJSON() ([]byte, error) {
	if src.IamEndPointUserInventory != nil {
		return json.Marshal(&src.IamEndPointUserInventory)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamEndPointUserInventoryRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamEndPointUserInventory != nil {
		return obj.IamEndPointUserInventory
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamEndPointUserInventoryRelationship struct {
	value *IamEndPointUserInventoryRelationship
	isSet bool
}

func (v NullableIamEndPointUserInventoryRelationship) Get() *IamEndPointUserInventoryRelationship {
	return v.value
}

func (v *NullableIamEndPointUserInventoryRelationship) Set(val *IamEndPointUserInventoryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserInventoryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserInventoryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserInventoryRelationship(val *IamEndPointUserInventoryRelationship) *NullableIamEndPointUserInventoryRelationship {
	return &NullableIamEndPointUserInventoryRelationship{value: val, isSet: true}
}

func (v NullableIamEndPointUserInventoryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserInventoryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
