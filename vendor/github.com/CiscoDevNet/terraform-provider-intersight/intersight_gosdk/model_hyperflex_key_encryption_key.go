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

// HyperflexKeyEncryptionKey Specifies a key encryption Key and parameters for the associated resource.
type HyperflexKeyEncryptionKey struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This defines whether we need to operate in an account recovery scenario or not. If yes, then most of the parameters will be populated from an internal MO. So, some of the input parameters MAY be ignored, if this value is set to true.
	IsAccountRecovery *bool `json:"IsAccountRecovery,omitempty"`
	// Indicates whether the value of the 'kek' property has been set.
	IsKekSet *bool `json:"IsKekSet,omitempty"`
	// Indicates whether the value of the 'passphrase' property has been set.
	IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
	// Number of iterations we want the hash to be run.
	Iteration *int64 `json:"Iteration,omitempty"`
	// Key encryption key used to encrypt the DEK's on the HyperFlex cluster.
	Kek *string `json:"Kek,omitempty"`
	// Resource id + time of creation used for retrieving the KEK.
	KeyId *string `json:"KeyId,omitempty"`
	// Last known Key encryption key state for this Key. * `NEW` - Key Encryption key is newly created. * `ACTIVE` - Key Encryption key is deployed on active resource. * `INACTIVE` - Key Encryption key is inactive and not used. * `INPROGRESS` - Key Encryption key is in a state where it was used on Intersight but did not receive confirmation from platform of success/failure.
	KeyState *string `json:"KeyState,omitempty"`
	// Initial passphrase for the encryption policy, password must contain a minimum of 12 characters, with at least 1 lowercase, 1 uppercase, 1 numeric.
	Passphrase *string `json:"Passphrase,omitempty"`
	// Resource type on which this key will be applied. * `CLUSTER` - Encryption is per HyperFlex cluster. * `DATASTORE` - Encryption is per dataStore on the HyperFlex cluster. * `DRIVE` - Encryption is per drive on the HyperFlex cluster.
	ResourceType *string `json:"ResourceType,omitempty"`
	// Copy of Key encryption key, which is used for sending the key over to the remote device endpoint. It is not persisited anywhere.
	TransitKek           *string                              `json:"TransitKek,omitempty"`
	ClusterProfile       *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
	ResourceMo           *MoBaseMoRelationship                `json:"ResourceMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexKeyEncryptionKey HyperflexKeyEncryptionKey

// NewHyperflexKeyEncryptionKey instantiates a new HyperflexKeyEncryptionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexKeyEncryptionKey(classId string, objectType string) *HyperflexKeyEncryptionKey {
	this := HyperflexKeyEncryptionKey{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isAccountRecovery bool = false
	this.IsAccountRecovery = &isAccountRecovery
	var keyState string = "NEW"
	this.KeyState = &keyState
	var resourceType string = "CLUSTER"
	this.ResourceType = &resourceType
	return &this
}

// NewHyperflexKeyEncryptionKeyWithDefaults instantiates a new HyperflexKeyEncryptionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexKeyEncryptionKeyWithDefaults() *HyperflexKeyEncryptionKey {
	this := HyperflexKeyEncryptionKey{}
	var classId string = "hyperflex.KeyEncryptionKey"
	this.ClassId = classId
	var objectType string = "hyperflex.KeyEncryptionKey"
	this.ObjectType = objectType
	var isAccountRecovery bool = false
	this.IsAccountRecovery = &isAccountRecovery
	var keyState string = "NEW"
	this.KeyState = &keyState
	var resourceType string = "CLUSTER"
	this.ResourceType = &resourceType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexKeyEncryptionKey) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexKeyEncryptionKey) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexKeyEncryptionKey) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexKeyEncryptionKey) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsAccountRecovery returns the IsAccountRecovery field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetIsAccountRecovery() bool {
	if o == nil || o.IsAccountRecovery == nil {
		var ret bool
		return ret
	}
	return *o.IsAccountRecovery
}

// GetIsAccountRecoveryOk returns a tuple with the IsAccountRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetIsAccountRecoveryOk() (*bool, bool) {
	if o == nil || o.IsAccountRecovery == nil {
		return nil, false
	}
	return o.IsAccountRecovery, true
}

// HasIsAccountRecovery returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasIsAccountRecovery() bool {
	if o != nil && o.IsAccountRecovery != nil {
		return true
	}

	return false
}

// SetIsAccountRecovery gets a reference to the given bool and assigns it to the IsAccountRecovery field.
func (o *HyperflexKeyEncryptionKey) SetIsAccountRecovery(v bool) {
	o.IsAccountRecovery = &v
}

// GetIsKekSet returns the IsKekSet field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetIsKekSet() bool {
	if o == nil || o.IsKekSet == nil {
		var ret bool
		return ret
	}
	return *o.IsKekSet
}

// GetIsKekSetOk returns a tuple with the IsKekSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetIsKekSetOk() (*bool, bool) {
	if o == nil || o.IsKekSet == nil {
		return nil, false
	}
	return o.IsKekSet, true
}

// HasIsKekSet returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasIsKekSet() bool {
	if o != nil && o.IsKekSet != nil {
		return true
	}

	return false
}

// SetIsKekSet gets a reference to the given bool and assigns it to the IsKekSet field.
func (o *HyperflexKeyEncryptionKey) SetIsKekSet(v bool) {
	o.IsKekSet = &v
}

// GetIsPassphraseSet returns the IsPassphraseSet field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetIsPassphraseSet() bool {
	if o == nil || o.IsPassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPassphraseSet
}

// GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetIsPassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsPassphraseSet == nil {
		return nil, false
	}
	return o.IsPassphraseSet, true
}

// HasIsPassphraseSet returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasIsPassphraseSet() bool {
	if o != nil && o.IsPassphraseSet != nil {
		return true
	}

	return false
}

// SetIsPassphraseSet gets a reference to the given bool and assigns it to the IsPassphraseSet field.
func (o *HyperflexKeyEncryptionKey) SetIsPassphraseSet(v bool) {
	o.IsPassphraseSet = &v
}

// GetIteration returns the Iteration field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetIteration() int64 {
	if o == nil || o.Iteration == nil {
		var ret int64
		return ret
	}
	return *o.Iteration
}

// GetIterationOk returns a tuple with the Iteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetIterationOk() (*int64, bool) {
	if o == nil || o.Iteration == nil {
		return nil, false
	}
	return o.Iteration, true
}

// HasIteration returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasIteration() bool {
	if o != nil && o.Iteration != nil {
		return true
	}

	return false
}

// SetIteration gets a reference to the given int64 and assigns it to the Iteration field.
func (o *HyperflexKeyEncryptionKey) SetIteration(v int64) {
	o.Iteration = &v
}

// GetKek returns the Kek field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetKek() string {
	if o == nil || o.Kek == nil {
		var ret string
		return ret
	}
	return *o.Kek
}

// GetKekOk returns a tuple with the Kek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetKekOk() (*string, bool) {
	if o == nil || o.Kek == nil {
		return nil, false
	}
	return o.Kek, true
}

// HasKek returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasKek() bool {
	if o != nil && o.Kek != nil {
		return true
	}

	return false
}

// SetKek gets a reference to the given string and assigns it to the Kek field.
func (o *HyperflexKeyEncryptionKey) SetKek(v string) {
	o.Kek = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *HyperflexKeyEncryptionKey) SetKeyId(v string) {
	o.KeyId = &v
}

// GetKeyState returns the KeyState field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetKeyState() string {
	if o == nil || o.KeyState == nil {
		var ret string
		return ret
	}
	return *o.KeyState
}

// GetKeyStateOk returns a tuple with the KeyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetKeyStateOk() (*string, bool) {
	if o == nil || o.KeyState == nil {
		return nil, false
	}
	return o.KeyState, true
}

// HasKeyState returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasKeyState() bool {
	if o != nil && o.KeyState != nil {
		return true
	}

	return false
}

// SetKeyState gets a reference to the given string and assigns it to the KeyState field.
func (o *HyperflexKeyEncryptionKey) SetKeyState(v string) {
	o.KeyState = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *HyperflexKeyEncryptionKey) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *HyperflexKeyEncryptionKey) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetTransitKek returns the TransitKek field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetTransitKek() string {
	if o == nil || o.TransitKek == nil {
		var ret string
		return ret
	}
	return *o.TransitKek
}

// GetTransitKekOk returns a tuple with the TransitKek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetTransitKekOk() (*string, bool) {
	if o == nil || o.TransitKek == nil {
		return nil, false
	}
	return o.TransitKek, true
}

// HasTransitKek returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasTransitKek() bool {
	if o != nil && o.TransitKek != nil {
		return true
	}

	return false
}

// SetTransitKek gets a reference to the given string and assigns it to the TransitKek field.
func (o *HyperflexKeyEncryptionKey) SetTransitKek(v string) {
	o.TransitKek = &v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetClusterProfile() HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfile == nil {
		var ret HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfile == nil {
		return nil, false
	}
	return o.ClusterProfile, true
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile != nil {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given HyperflexClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *HyperflexKeyEncryptionKey) SetClusterProfile(v HyperflexClusterProfileRelationship) {
	o.ClusterProfile = &v
}

// GetResourceMo returns the ResourceMo field value if set, zero value otherwise.
func (o *HyperflexKeyEncryptionKey) GetResourceMo() MoBaseMoRelationship {
	if o == nil || o.ResourceMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.ResourceMo
}

// GetResourceMoOk returns a tuple with the ResourceMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexKeyEncryptionKey) GetResourceMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.ResourceMo == nil {
		return nil, false
	}
	return o.ResourceMo, true
}

// HasResourceMo returns a boolean if a field has been set.
func (o *HyperflexKeyEncryptionKey) HasResourceMo() bool {
	if o != nil && o.ResourceMo != nil {
		return true
	}

	return false
}

// SetResourceMo gets a reference to the given MoBaseMoRelationship and assigns it to the ResourceMo field.
func (o *HyperflexKeyEncryptionKey) SetResourceMo(v MoBaseMoRelationship) {
	o.ResourceMo = &v
}

func (o HyperflexKeyEncryptionKey) MarshalJSON() ([]byte, error) {
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
	if o.IsAccountRecovery != nil {
		toSerialize["IsAccountRecovery"] = o.IsAccountRecovery
	}
	if o.IsKekSet != nil {
		toSerialize["IsKekSet"] = o.IsKekSet
	}
	if o.IsPassphraseSet != nil {
		toSerialize["IsPassphraseSet"] = o.IsPassphraseSet
	}
	if o.Iteration != nil {
		toSerialize["Iteration"] = o.Iteration
	}
	if o.Kek != nil {
		toSerialize["Kek"] = o.Kek
	}
	if o.KeyId != nil {
		toSerialize["KeyId"] = o.KeyId
	}
	if o.KeyState != nil {
		toSerialize["KeyState"] = o.KeyState
	}
	if o.Passphrase != nil {
		toSerialize["Passphrase"] = o.Passphrase
	}
	if o.ResourceType != nil {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if o.TransitKek != nil {
		toSerialize["TransitKek"] = o.TransitKek
	}
	if o.ClusterProfile != nil {
		toSerialize["ClusterProfile"] = o.ClusterProfile
	}
	if o.ResourceMo != nil {
		toSerialize["ResourceMo"] = o.ResourceMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexKeyEncryptionKey) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexKeyEncryptionKeyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This defines whether we need to operate in an account recovery scenario or not. If yes, then most of the parameters will be populated from an internal MO. So, some of the input parameters MAY be ignored, if this value is set to true.
		IsAccountRecovery *bool `json:"IsAccountRecovery,omitempty"`
		// Indicates whether the value of the 'kek' property has been set.
		IsKekSet *bool `json:"IsKekSet,omitempty"`
		// Indicates whether the value of the 'passphrase' property has been set.
		IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
		// Number of iterations we want the hash to be run.
		Iteration *int64 `json:"Iteration,omitempty"`
		// Key encryption key used to encrypt the DEK's on the HyperFlex cluster.
		Kek *string `json:"Kek,omitempty"`
		// Resource id + time of creation used for retrieving the KEK.
		KeyId *string `json:"KeyId,omitempty"`
		// Last known Key encryption key state for this Key. * `NEW` - Key Encryption key is newly created. * `ACTIVE` - Key Encryption key is deployed on active resource. * `INACTIVE` - Key Encryption key is inactive and not used. * `INPROGRESS` - Key Encryption key is in a state where it was used on Intersight but did not receive confirmation from platform of success/failure.
		KeyState *string `json:"KeyState,omitempty"`
		// Initial passphrase for the encryption policy, password must contain a minimum of 12 characters, with at least 1 lowercase, 1 uppercase, 1 numeric.
		Passphrase *string `json:"Passphrase,omitempty"`
		// Resource type on which this key will be applied. * `CLUSTER` - Encryption is per HyperFlex cluster. * `DATASTORE` - Encryption is per dataStore on the HyperFlex cluster. * `DRIVE` - Encryption is per drive on the HyperFlex cluster.
		ResourceType *string `json:"ResourceType,omitempty"`
		// Copy of Key encryption key, which is used for sending the key over to the remote device endpoint. It is not persisited anywhere.
		TransitKek     *string                              `json:"TransitKek,omitempty"`
		ClusterProfile *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
		ResourceMo     *MoBaseMoRelationship                `json:"ResourceMo,omitempty"`
	}

	varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct := HyperflexKeyEncryptionKeyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexKeyEncryptionKey := _HyperflexKeyEncryptionKey{}
		varHyperflexKeyEncryptionKey.ClassId = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.ClassId
		varHyperflexKeyEncryptionKey.ObjectType = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.ObjectType
		varHyperflexKeyEncryptionKey.IsAccountRecovery = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.IsAccountRecovery
		varHyperflexKeyEncryptionKey.IsKekSet = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.IsKekSet
		varHyperflexKeyEncryptionKey.IsPassphraseSet = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.IsPassphraseSet
		varHyperflexKeyEncryptionKey.Iteration = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.Iteration
		varHyperflexKeyEncryptionKey.Kek = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.Kek
		varHyperflexKeyEncryptionKey.KeyId = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.KeyId
		varHyperflexKeyEncryptionKey.KeyState = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.KeyState
		varHyperflexKeyEncryptionKey.Passphrase = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.Passphrase
		varHyperflexKeyEncryptionKey.ResourceType = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.ResourceType
		varHyperflexKeyEncryptionKey.TransitKek = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.TransitKek
		varHyperflexKeyEncryptionKey.ClusterProfile = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.ClusterProfile
		varHyperflexKeyEncryptionKey.ResourceMo = varHyperflexKeyEncryptionKeyWithoutEmbeddedStruct.ResourceMo
		*o = HyperflexKeyEncryptionKey(varHyperflexKeyEncryptionKey)
	} else {
		return err
	}

	varHyperflexKeyEncryptionKey := _HyperflexKeyEncryptionKey{}

	err = json.Unmarshal(bytes, &varHyperflexKeyEncryptionKey)
	if err == nil {
		o.MoBaseMo = varHyperflexKeyEncryptionKey.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsAccountRecovery")
		delete(additionalProperties, "IsKekSet")
		delete(additionalProperties, "IsPassphraseSet")
		delete(additionalProperties, "Iteration")
		delete(additionalProperties, "Kek")
		delete(additionalProperties, "KeyId")
		delete(additionalProperties, "KeyState")
		delete(additionalProperties, "Passphrase")
		delete(additionalProperties, "ResourceType")
		delete(additionalProperties, "TransitKek")
		delete(additionalProperties, "ClusterProfile")
		delete(additionalProperties, "ResourceMo")

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

type NullableHyperflexKeyEncryptionKey struct {
	value *HyperflexKeyEncryptionKey
	isSet bool
}

func (v NullableHyperflexKeyEncryptionKey) Get() *HyperflexKeyEncryptionKey {
	return v.value
}

func (v *NullableHyperflexKeyEncryptionKey) Set(val *HyperflexKeyEncryptionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexKeyEncryptionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexKeyEncryptionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexKeyEncryptionKey(val *HyperflexKeyEncryptionKey) *NullableHyperflexKeyEncryptionKey {
	return &NullableHyperflexKeyEncryptionKey{value: val, isSet: true}
}

func (v NullableHyperflexKeyEncryptionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexKeyEncryptionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
