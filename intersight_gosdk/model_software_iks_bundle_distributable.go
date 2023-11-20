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

// SoftwareIksBundleDistributable An IKS image bundle distributed by Cisco for Private Appliance.
type SoftwareIksBundleDistributable struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                 `json:"ObjectType"`
	Catalog    *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	// An array of relationships to softwareSolutionDistributable resources.
	Images               []SoftwareSolutionDistributableRelationship `json:"Images,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareIksBundleDistributable SoftwareIksBundleDistributable

// NewSoftwareIksBundleDistributable instantiates a new SoftwareIksBundleDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareIksBundleDistributable(classId string, objectType string) *SoftwareIksBundleDistributable {
	this := SoftwareIksBundleDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var vendor string = "Cisco"
	this.Vendor = &vendor
	return &this
}

// NewSoftwareIksBundleDistributableWithDefaults instantiates a new SoftwareIksBundleDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareIksBundleDistributableWithDefaults() *SoftwareIksBundleDistributable {
	this := SoftwareIksBundleDistributable{}
	var classId string = "software.IksBundleDistributable"
	this.ClassId = classId
	var objectType string = "software.IksBundleDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareIksBundleDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareIksBundleDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareIksBundleDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareIksBundleDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareIksBundleDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareIksBundleDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareIksBundleDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareIksBundleDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareIksBundleDistributable) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareIksBundleDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareIksBundleDistributable) GetImages() []SoftwareSolutionDistributableRelationship {
	if o == nil {
		var ret []SoftwareSolutionDistributableRelationship
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareIksBundleDistributable) GetImagesOk() ([]SoftwareSolutionDistributableRelationship, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *SoftwareIksBundleDistributable) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []SoftwareSolutionDistributableRelationship and assigns it to the Images field.
func (o *SoftwareIksBundleDistributable) SetImages(v []SoftwareSolutionDistributableRelationship) {
	o.Images = v
}

func (o SoftwareIksBundleDistributable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Images != nil {
		toSerialize["Images"] = o.Images
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareIksBundleDistributable) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwareIksBundleDistributableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                 `json:"ObjectType"`
		Catalog    *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
		// An array of relationships to softwareSolutionDistributable resources.
		Images []SoftwareSolutionDistributableRelationship `json:"Images,omitempty"`
	}

	varSoftwareIksBundleDistributableWithoutEmbeddedStruct := SoftwareIksBundleDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwareIksBundleDistributableWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareIksBundleDistributable := _SoftwareIksBundleDistributable{}
		varSoftwareIksBundleDistributable.ClassId = varSoftwareIksBundleDistributableWithoutEmbeddedStruct.ClassId
		varSoftwareIksBundleDistributable.ObjectType = varSoftwareIksBundleDistributableWithoutEmbeddedStruct.ObjectType
		varSoftwareIksBundleDistributable.Catalog = varSoftwareIksBundleDistributableWithoutEmbeddedStruct.Catalog
		varSoftwareIksBundleDistributable.Images = varSoftwareIksBundleDistributableWithoutEmbeddedStruct.Images
		*o = SoftwareIksBundleDistributable(varSoftwareIksBundleDistributable)
	} else {
		return err
	}

	varSoftwareIksBundleDistributable := _SoftwareIksBundleDistributable{}

	err = json.Unmarshal(bytes, &varSoftwareIksBundleDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varSoftwareIksBundleDistributable.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Images")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableSoftwareIksBundleDistributable struct {
	value *SoftwareIksBundleDistributable
	isSet bool
}

func (v NullableSoftwareIksBundleDistributable) Get() *SoftwareIksBundleDistributable {
	return v.value
}

func (v *NullableSoftwareIksBundleDistributable) Set(val *SoftwareIksBundleDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareIksBundleDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareIksBundleDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareIksBundleDistributable(val *SoftwareIksBundleDistributable) *NullableSoftwareIksBundleDistributable {
	return &NullableSoftwareIksBundleDistributable{value: val, isSet: true}
}

func (v NullableSoftwareIksBundleDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareIksBundleDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
