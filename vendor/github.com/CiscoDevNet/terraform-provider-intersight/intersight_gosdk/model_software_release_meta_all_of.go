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
)

// SoftwareReleaseMetaAllOf Definition of the list of properties defined in 'software.ReleaseMeta', excluding properties defined in parent classes.
type SoftwareReleaseMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The subtype of the distributable image. For e.g. the firmware distributable is categorized according to the component it can upgrade - Standalone server, Intersight managed server or UCS Managed Fabric Interconnect.
	ImageType *string `json:"ImageType,omitempty"`
	// The name of the latest image file uploaded for this software type. It is populated as part of the image import operation.
	LatestFileName *string `json:"LatestFileName,omitempty"`
	// Latest version of the image avaiable for a specific software.
	LatestVersion *string `json:"LatestVersion,omitempty"`
	// The software type id of the image (For e.g. firmware.Distributable, software.ApplianceDistributable, software.HyperflexBundleDistributable, software.UcsdBundleDistributable).
	SoftwareTypeId       *string                                `json:"SoftwareTypeId,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	Image                *FirmwareBaseDistributableRelationship `json:"Image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareReleaseMetaAllOf SoftwareReleaseMetaAllOf

// NewSoftwareReleaseMetaAllOf instantiates a new SoftwareReleaseMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareReleaseMetaAllOf(classId string, objectType string) *SoftwareReleaseMetaAllOf {
	this := SoftwareReleaseMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareReleaseMetaAllOfWithDefaults instantiates a new SoftwareReleaseMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareReleaseMetaAllOfWithDefaults() *SoftwareReleaseMetaAllOf {
	this := SoftwareReleaseMetaAllOf{}
	var classId string = "software.ReleaseMeta"
	this.ClassId = classId
	var objectType string = "software.ReleaseMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareReleaseMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareReleaseMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareReleaseMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareReleaseMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *SoftwareReleaseMetaAllOf) GetImageType() string {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetImageTypeOk() (*string, bool) {
	if o == nil || o.ImageType == nil {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *SoftwareReleaseMetaAllOf) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *SoftwareReleaseMetaAllOf) SetImageType(v string) {
	o.ImageType = &v
}

// GetLatestFileName returns the LatestFileName field value if set, zero value otherwise.
func (o *SoftwareReleaseMetaAllOf) GetLatestFileName() string {
	if o == nil || o.LatestFileName == nil {
		var ret string
		return ret
	}
	return *o.LatestFileName
}

// GetLatestFileNameOk returns a tuple with the LatestFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetLatestFileNameOk() (*string, bool) {
	if o == nil || o.LatestFileName == nil {
		return nil, false
	}
	return o.LatestFileName, true
}

// HasLatestFileName returns a boolean if a field has been set.
func (o *SoftwareReleaseMetaAllOf) HasLatestFileName() bool {
	if o != nil && o.LatestFileName != nil {
		return true
	}

	return false
}

// SetLatestFileName gets a reference to the given string and assigns it to the LatestFileName field.
func (o *SoftwareReleaseMetaAllOf) SetLatestFileName(v string) {
	o.LatestFileName = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *SoftwareReleaseMetaAllOf) GetLatestVersion() string {
	if o == nil || o.LatestVersion == nil {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetLatestVersionOk() (*string, bool) {
	if o == nil || o.LatestVersion == nil {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *SoftwareReleaseMetaAllOf) HasLatestVersion() bool {
	if o != nil && o.LatestVersion != nil {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *SoftwareReleaseMetaAllOf) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetSoftwareTypeId returns the SoftwareTypeId field value if set, zero value otherwise.
func (o *SoftwareReleaseMetaAllOf) GetSoftwareTypeId() string {
	if o == nil || o.SoftwareTypeId == nil {
		var ret string
		return ret
	}
	return *o.SoftwareTypeId
}

// GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetSoftwareTypeIdOk() (*string, bool) {
	if o == nil || o.SoftwareTypeId == nil {
		return nil, false
	}
	return o.SoftwareTypeId, true
}

// HasSoftwareTypeId returns a boolean if a field has been set.
func (o *SoftwareReleaseMetaAllOf) HasSoftwareTypeId() bool {
	if o != nil && o.SoftwareTypeId != nil {
		return true
	}

	return false
}

// SetSoftwareTypeId gets a reference to the given string and assigns it to the SoftwareTypeId field.
func (o *SoftwareReleaseMetaAllOf) SetSoftwareTypeId(v string) {
	o.SoftwareTypeId = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareReleaseMetaAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareReleaseMetaAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareReleaseMetaAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SoftwareReleaseMetaAllOf) GetImage() FirmwareBaseDistributableRelationship {
	if o == nil || o.Image == nil {
		var ret FirmwareBaseDistributableRelationship
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareReleaseMetaAllOf) GetImageOk() (*FirmwareBaseDistributableRelationship, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SoftwareReleaseMetaAllOf) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given FirmwareBaseDistributableRelationship and assigns it to the Image field.
func (o *SoftwareReleaseMetaAllOf) SetImage(v FirmwareBaseDistributableRelationship) {
	o.Image = &v
}

func (o SoftwareReleaseMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ImageType != nil {
		toSerialize["ImageType"] = o.ImageType
	}
	if o.LatestFileName != nil {
		toSerialize["LatestFileName"] = o.LatestFileName
	}
	if o.LatestVersion != nil {
		toSerialize["LatestVersion"] = o.LatestVersion
	}
	if o.SoftwareTypeId != nil {
		toSerialize["SoftwareTypeId"] = o.SoftwareTypeId
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareReleaseMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareReleaseMetaAllOf := _SoftwareReleaseMetaAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwareReleaseMetaAllOf); err == nil {
		*o = SoftwareReleaseMetaAllOf(varSoftwareReleaseMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "LatestFileName")
		delete(additionalProperties, "LatestVersion")
		delete(additionalProperties, "SoftwareTypeId")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareReleaseMetaAllOf struct {
	value *SoftwareReleaseMetaAllOf
	isSet bool
}

func (v NullableSoftwareReleaseMetaAllOf) Get() *SoftwareReleaseMetaAllOf {
	return v.value
}

func (v *NullableSoftwareReleaseMetaAllOf) Set(val *SoftwareReleaseMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareReleaseMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareReleaseMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareReleaseMetaAllOf(val *SoftwareReleaseMetaAllOf) *NullableSoftwareReleaseMetaAllOf {
	return &NullableSoftwareReleaseMetaAllOf{value: val, isSet: true}
}

func (v NullableSoftwareReleaseMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareReleaseMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
