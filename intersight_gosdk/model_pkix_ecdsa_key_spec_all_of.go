/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// PkixEcdsaKeySpecAllOf Definition of the list of properties defined in 'pkix.EcdsaKeySpec', excluding properties defined in parent classes.
type PkixEcdsaKeySpecAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A specific set of Elliptic Curve parameters, as recommended by NIST in FIPS 186-4. * `P256` - P256 returns a Curve which implements P-256, as defined in FIPS 186-4, section D.2.3. * `P224` - P224 returns a Curve which implements P-224, as defined in FIPS 186-4, section D.2.2. * `P384` - P384 returns a Curve which implements P-384, as defined in FIPS 186-4, section D.2.4. * `P521` - P521 returns a Curve which implements P-521, as defined in FIPS 186-4, section D.2.5.
	Curve                *string `json:"Curve,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixEcdsaKeySpecAllOf PkixEcdsaKeySpecAllOf

// NewPkixEcdsaKeySpecAllOf instantiates a new PkixEcdsaKeySpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixEcdsaKeySpecAllOf(classId string, objectType string) *PkixEcdsaKeySpecAllOf {
	this := PkixEcdsaKeySpecAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var curve string = "P256"
	this.Curve = &curve
	return &this
}

// NewPkixEcdsaKeySpecAllOfWithDefaults instantiates a new PkixEcdsaKeySpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixEcdsaKeySpecAllOfWithDefaults() *PkixEcdsaKeySpecAllOf {
	this := PkixEcdsaKeySpecAllOf{}
	var classId string = "pkix.EcdsaKeySpec"
	this.ClassId = classId
	var objectType string = "pkix.EcdsaKeySpec"
	this.ObjectType = objectType
	var curve string = "P256"
	this.Curve = &curve
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixEcdsaKeySpecAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixEcdsaKeySpecAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixEcdsaKeySpecAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixEcdsaKeySpecAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixEcdsaKeySpecAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixEcdsaKeySpecAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurve returns the Curve field value if set, zero value otherwise.
func (o *PkixEcdsaKeySpecAllOf) GetCurve() string {
	if o == nil || o.Curve == nil {
		var ret string
		return ret
	}
	return *o.Curve
}

// GetCurveOk returns a tuple with the Curve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixEcdsaKeySpecAllOf) GetCurveOk() (*string, bool) {
	if o == nil || o.Curve == nil {
		return nil, false
	}
	return o.Curve, true
}

// HasCurve returns a boolean if a field has been set.
func (o *PkixEcdsaKeySpecAllOf) HasCurve() bool {
	if o != nil && o.Curve != nil {
		return true
	}

	return false
}

// SetCurve gets a reference to the given string and assigns it to the Curve field.
func (o *PkixEcdsaKeySpecAllOf) SetCurve(v string) {
	o.Curve = &v
}

func (o PkixEcdsaKeySpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Curve != nil {
		toSerialize["Curve"] = o.Curve
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixEcdsaKeySpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPkixEcdsaKeySpecAllOf := _PkixEcdsaKeySpecAllOf{}

	if err = json.Unmarshal(bytes, &varPkixEcdsaKeySpecAllOf); err == nil {
		*o = PkixEcdsaKeySpecAllOf(varPkixEcdsaKeySpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Curve")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePkixEcdsaKeySpecAllOf struct {
	value *PkixEcdsaKeySpecAllOf
	isSet bool
}

func (v NullablePkixEcdsaKeySpecAllOf) Get() *PkixEcdsaKeySpecAllOf {
	return v.value
}

func (v *NullablePkixEcdsaKeySpecAllOf) Set(val *PkixEcdsaKeySpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixEcdsaKeySpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixEcdsaKeySpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixEcdsaKeySpecAllOf(val *PkixEcdsaKeySpecAllOf) *NullablePkixEcdsaKeySpecAllOf {
	return &NullablePkixEcdsaKeySpecAllOf{value: val, isSet: true}
}

func (v NullablePkixEcdsaKeySpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixEcdsaKeySpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
