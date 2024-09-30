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
	"reflect"
	"strings"
)

// checks if the SdaaciConnectionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdaaciConnectionDetail{}

// SdaaciConnectionDetail Peer connection details for each connection.
type SdaaciConnectionDetail struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of this connection between two peers.
	Description *string `json:"Description,omitempty"`
	// Id of the ip pool configured for this connection.
	IpPool *string `json:"IpPool,omitempty"`
	// Id of layer 3 handoff configured between a border node and a border leaf.
	Layer3HandoffId *string `json:"Layer3HandoffId,omitempty"`
	// Interface id configured on Peer A.
	PeerAinterface *string `json:"PeerAinterface,omitempty"`
	// The IP Address of the device used as the local peer.
	PeerAipAddress *string `json:"PeerAipAddress,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2][0-9]|3[0-2])$"`
	// Type of device used as Peer A for this peer connection.
	PeerAtype *string `json:"PeerAtype,omitempty"`
	// Interface id configured on Peer B.
	PeerBinterface *string `json:"PeerBinterface,omitempty"`
	// The IP Address of the device used as the remote peer.
	PeerBipAddress *string `json:"PeerBipAddress,omitempty"`
	// Type of device used as Peer B for this peer connection.
	PeerBtype *string `json:"PeerBtype,omitempty"`
	// First peer of the connection.
	Peera *string `json:"Peera,omitempty"`
	// Second Peer of the connection.
	Peerb *string `json:"Peerb,omitempty"`
	// Router id defined for this peer connection.
	RouterId *string `json:"RouterId,omitempty"`
	// Connection status between the peers. * `NotConnected` - Connection Status NotConnected. * `Connected` - Connection Status Connected.
	Status               *string                              `json:"Status,omitempty"`
	Connection           NullableSdaaciConnectionRelationship `json:"Connection,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdaaciConnectionDetail SdaaciConnectionDetail

// NewSdaaciConnectionDetail instantiates a new SdaaciConnectionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdaaciConnectionDetail(classId string, objectType string) *SdaaciConnectionDetail {
	this := SdaaciConnectionDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "NotConnected"
	this.Status = &status
	return &this
}

// NewSdaaciConnectionDetailWithDefaults instantiates a new SdaaciConnectionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdaaciConnectionDetailWithDefaults() *SdaaciConnectionDetail {
	this := SdaaciConnectionDetail{}
	var classId string = "sdaaci.ConnectionDetail"
	this.ClassId = classId
	var objectType string = "sdaaci.ConnectionDetail"
	this.ObjectType = objectType
	var status string = "NotConnected"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdaaciConnectionDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdaaciConnectionDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "sdaaci.ConnectionDetail" of the ClassId field.
func (o *SdaaciConnectionDetail) GetDefaultClassId() interface{} {
	return "sdaaci.ConnectionDetail"
}

// GetObjectType returns the ObjectType field value
func (o *SdaaciConnectionDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdaaciConnectionDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "sdaaci.ConnectionDetail" of the ObjectType field.
func (o *SdaaciConnectionDetail) GetDefaultObjectType() interface{} {
	return "sdaaci.ConnectionDetail"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SdaaciConnectionDetail) SetDescription(v string) {
	o.Description = &v
}

// GetIpPool returns the IpPool field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetIpPool() string {
	if o == nil || IsNil(o.IpPool) {
		var ret string
		return ret
	}
	return *o.IpPool
}

// GetIpPoolOk returns a tuple with the IpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetIpPoolOk() (*string, bool) {
	if o == nil || IsNil(o.IpPool) {
		return nil, false
	}
	return o.IpPool, true
}

// HasIpPool returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasIpPool() bool {
	if o != nil && !IsNil(o.IpPool) {
		return true
	}

	return false
}

// SetIpPool gets a reference to the given string and assigns it to the IpPool field.
func (o *SdaaciConnectionDetail) SetIpPool(v string) {
	o.IpPool = &v
}

// GetLayer3HandoffId returns the Layer3HandoffId field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetLayer3HandoffId() string {
	if o == nil || IsNil(o.Layer3HandoffId) {
		var ret string
		return ret
	}
	return *o.Layer3HandoffId
}

// GetLayer3HandoffIdOk returns a tuple with the Layer3HandoffId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetLayer3HandoffIdOk() (*string, bool) {
	if o == nil || IsNil(o.Layer3HandoffId) {
		return nil, false
	}
	return o.Layer3HandoffId, true
}

// HasLayer3HandoffId returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasLayer3HandoffId() bool {
	if o != nil && !IsNil(o.Layer3HandoffId) {
		return true
	}

	return false
}

// SetLayer3HandoffId gets a reference to the given string and assigns it to the Layer3HandoffId field.
func (o *SdaaciConnectionDetail) SetLayer3HandoffId(v string) {
	o.Layer3HandoffId = &v
}

// GetPeerAinterface returns the PeerAinterface field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerAinterface() string {
	if o == nil || IsNil(o.PeerAinterface) {
		var ret string
		return ret
	}
	return *o.PeerAinterface
}

// GetPeerAinterfaceOk returns a tuple with the PeerAinterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerAinterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.PeerAinterface) {
		return nil, false
	}
	return o.PeerAinterface, true
}

// HasPeerAinterface returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerAinterface() bool {
	if o != nil && !IsNil(o.PeerAinterface) {
		return true
	}

	return false
}

// SetPeerAinterface gets a reference to the given string and assigns it to the PeerAinterface field.
func (o *SdaaciConnectionDetail) SetPeerAinterface(v string) {
	o.PeerAinterface = &v
}

// GetPeerAipAddress returns the PeerAipAddress field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerAipAddress() string {
	if o == nil || IsNil(o.PeerAipAddress) {
		var ret string
		return ret
	}
	return *o.PeerAipAddress
}

// GetPeerAipAddressOk returns a tuple with the PeerAipAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerAipAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PeerAipAddress) {
		return nil, false
	}
	return o.PeerAipAddress, true
}

// HasPeerAipAddress returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerAipAddress() bool {
	if o != nil && !IsNil(o.PeerAipAddress) {
		return true
	}

	return false
}

// SetPeerAipAddress gets a reference to the given string and assigns it to the PeerAipAddress field.
func (o *SdaaciConnectionDetail) SetPeerAipAddress(v string) {
	o.PeerAipAddress = &v
}

// GetPeerAtype returns the PeerAtype field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerAtype() string {
	if o == nil || IsNil(o.PeerAtype) {
		var ret string
		return ret
	}
	return *o.PeerAtype
}

// GetPeerAtypeOk returns a tuple with the PeerAtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerAtypeOk() (*string, bool) {
	if o == nil || IsNil(o.PeerAtype) {
		return nil, false
	}
	return o.PeerAtype, true
}

// HasPeerAtype returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerAtype() bool {
	if o != nil && !IsNil(o.PeerAtype) {
		return true
	}

	return false
}

// SetPeerAtype gets a reference to the given string and assigns it to the PeerAtype field.
func (o *SdaaciConnectionDetail) SetPeerAtype(v string) {
	o.PeerAtype = &v
}

// GetPeerBinterface returns the PeerBinterface field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerBinterface() string {
	if o == nil || IsNil(o.PeerBinterface) {
		var ret string
		return ret
	}
	return *o.PeerBinterface
}

// GetPeerBinterfaceOk returns a tuple with the PeerBinterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerBinterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.PeerBinterface) {
		return nil, false
	}
	return o.PeerBinterface, true
}

// HasPeerBinterface returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerBinterface() bool {
	if o != nil && !IsNil(o.PeerBinterface) {
		return true
	}

	return false
}

// SetPeerBinterface gets a reference to the given string and assigns it to the PeerBinterface field.
func (o *SdaaciConnectionDetail) SetPeerBinterface(v string) {
	o.PeerBinterface = &v
}

// GetPeerBipAddress returns the PeerBipAddress field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerBipAddress() string {
	if o == nil || IsNil(o.PeerBipAddress) {
		var ret string
		return ret
	}
	return *o.PeerBipAddress
}

// GetPeerBipAddressOk returns a tuple with the PeerBipAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerBipAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PeerBipAddress) {
		return nil, false
	}
	return o.PeerBipAddress, true
}

// HasPeerBipAddress returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerBipAddress() bool {
	if o != nil && !IsNil(o.PeerBipAddress) {
		return true
	}

	return false
}

// SetPeerBipAddress gets a reference to the given string and assigns it to the PeerBipAddress field.
func (o *SdaaciConnectionDetail) SetPeerBipAddress(v string) {
	o.PeerBipAddress = &v
}

// GetPeerBtype returns the PeerBtype field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerBtype() string {
	if o == nil || IsNil(o.PeerBtype) {
		var ret string
		return ret
	}
	return *o.PeerBtype
}

// GetPeerBtypeOk returns a tuple with the PeerBtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerBtypeOk() (*string, bool) {
	if o == nil || IsNil(o.PeerBtype) {
		return nil, false
	}
	return o.PeerBtype, true
}

// HasPeerBtype returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerBtype() bool {
	if o != nil && !IsNil(o.PeerBtype) {
		return true
	}

	return false
}

// SetPeerBtype gets a reference to the given string and assigns it to the PeerBtype field.
func (o *SdaaciConnectionDetail) SetPeerBtype(v string) {
	o.PeerBtype = &v
}

// GetPeera returns the Peera field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeera() string {
	if o == nil || IsNil(o.Peera) {
		var ret string
		return ret
	}
	return *o.Peera
}

// GetPeeraOk returns a tuple with the Peera field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeeraOk() (*string, bool) {
	if o == nil || IsNil(o.Peera) {
		return nil, false
	}
	return o.Peera, true
}

// HasPeera returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeera() bool {
	if o != nil && !IsNil(o.Peera) {
		return true
	}

	return false
}

// SetPeera gets a reference to the given string and assigns it to the Peera field.
func (o *SdaaciConnectionDetail) SetPeera(v string) {
	o.Peera = &v
}

// GetPeerb returns the Peerb field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetPeerb() string {
	if o == nil || IsNil(o.Peerb) {
		var ret string
		return ret
	}
	return *o.Peerb
}

// GetPeerbOk returns a tuple with the Peerb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetPeerbOk() (*string, bool) {
	if o == nil || IsNil(o.Peerb) {
		return nil, false
	}
	return o.Peerb, true
}

// HasPeerb returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasPeerb() bool {
	if o != nil && !IsNil(o.Peerb) {
		return true
	}

	return false
}

// SetPeerb gets a reference to the given string and assigns it to the Peerb field.
func (o *SdaaciConnectionDetail) SetPeerb(v string) {
	o.Peerb = &v
}

// GetRouterId returns the RouterId field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetRouterId() string {
	if o == nil || IsNil(o.RouterId) {
		var ret string
		return ret
	}
	return *o.RouterId
}

// GetRouterIdOk returns a tuple with the RouterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetRouterIdOk() (*string, bool) {
	if o == nil || IsNil(o.RouterId) {
		return nil, false
	}
	return o.RouterId, true
}

// HasRouterId returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasRouterId() bool {
	if o != nil && !IsNil(o.RouterId) {
		return true
	}

	return false
}

// SetRouterId gets a reference to the given string and assigns it to the RouterId field.
func (o *SdaaciConnectionDetail) SetRouterId(v string) {
	o.RouterId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SdaaciConnectionDetail) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdaaciConnectionDetail) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SdaaciConnectionDetail) SetStatus(v string) {
	o.Status = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdaaciConnectionDetail) GetConnection() SdaaciConnectionRelationship {
	if o == nil || IsNil(o.Connection.Get()) {
		var ret SdaaciConnectionRelationship
		return ret
	}
	return *o.Connection.Get()
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdaaciConnectionDetail) GetConnectionOk() (*SdaaciConnectionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connection.Get(), o.Connection.IsSet()
}

// HasConnection returns a boolean if a field has been set.
func (o *SdaaciConnectionDetail) HasConnection() bool {
	if o != nil && o.Connection.IsSet() {
		return true
	}

	return false
}

// SetConnection gets a reference to the given NullableSdaaciConnectionRelationship and assigns it to the Connection field.
func (o *SdaaciConnectionDetail) SetConnection(v SdaaciConnectionRelationship) {
	o.Connection.Set(&v)
}

// SetConnectionNil sets the value for Connection to be an explicit nil
func (o *SdaaciConnectionDetail) SetConnectionNil() {
	o.Connection.Set(nil)
}

// UnsetConnection ensures that no value is present for Connection, not even an explicit nil
func (o *SdaaciConnectionDetail) UnsetConnection() {
	o.Connection.Unset()
}

func (o SdaaciConnectionDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdaaciConnectionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.IpPool) {
		toSerialize["IpPool"] = o.IpPool
	}
	if !IsNil(o.Layer3HandoffId) {
		toSerialize["Layer3HandoffId"] = o.Layer3HandoffId
	}
	if !IsNil(o.PeerAinterface) {
		toSerialize["PeerAinterface"] = o.PeerAinterface
	}
	if !IsNil(o.PeerAipAddress) {
		toSerialize["PeerAipAddress"] = o.PeerAipAddress
	}
	if !IsNil(o.PeerAtype) {
		toSerialize["PeerAtype"] = o.PeerAtype
	}
	if !IsNil(o.PeerBinterface) {
		toSerialize["PeerBinterface"] = o.PeerBinterface
	}
	if !IsNil(o.PeerBipAddress) {
		toSerialize["PeerBipAddress"] = o.PeerBipAddress
	}
	if !IsNil(o.PeerBtype) {
		toSerialize["PeerBtype"] = o.PeerBtype
	}
	if !IsNil(o.Peera) {
		toSerialize["Peera"] = o.Peera
	}
	if !IsNil(o.Peerb) {
		toSerialize["Peerb"] = o.Peerb
	}
	if !IsNil(o.RouterId) {
		toSerialize["RouterId"] = o.RouterId
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.Connection.IsSet() {
		toSerialize["Connection"] = o.Connection.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SdaaciConnectionDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
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
	type SdaaciConnectionDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of this connection between two peers.
		Description *string `json:"Description,omitempty"`
		// Id of the ip pool configured for this connection.
		IpPool *string `json:"IpPool,omitempty"`
		// Id of layer 3 handoff configured between a border node and a border leaf.
		Layer3HandoffId *string `json:"Layer3HandoffId,omitempty"`
		// Interface id configured on Peer A.
		PeerAinterface *string `json:"PeerAinterface,omitempty"`
		// The IP Address of the device used as the local peer.
		PeerAipAddress *string `json:"PeerAipAddress,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2][0-9]|3[0-2])$"`
		// Type of device used as Peer A for this peer connection.
		PeerAtype *string `json:"PeerAtype,omitempty"`
		// Interface id configured on Peer B.
		PeerBinterface *string `json:"PeerBinterface,omitempty"`
		// The IP Address of the device used as the remote peer.
		PeerBipAddress *string `json:"PeerBipAddress,omitempty"`
		// Type of device used as Peer B for this peer connection.
		PeerBtype *string `json:"PeerBtype,omitempty"`
		// First peer of the connection.
		Peera *string `json:"Peera,omitempty"`
		// Second Peer of the connection.
		Peerb *string `json:"Peerb,omitempty"`
		// Router id defined for this peer connection.
		RouterId *string `json:"RouterId,omitempty"`
		// Connection status between the peers. * `NotConnected` - Connection Status NotConnected. * `Connected` - Connection Status Connected.
		Status     *string                              `json:"Status,omitempty"`
		Connection NullableSdaaciConnectionRelationship `json:"Connection,omitempty"`
	}

	varSdaaciConnectionDetailWithoutEmbeddedStruct := SdaaciConnectionDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSdaaciConnectionDetailWithoutEmbeddedStruct)
	if err == nil {
		varSdaaciConnectionDetail := _SdaaciConnectionDetail{}
		varSdaaciConnectionDetail.ClassId = varSdaaciConnectionDetailWithoutEmbeddedStruct.ClassId
		varSdaaciConnectionDetail.ObjectType = varSdaaciConnectionDetailWithoutEmbeddedStruct.ObjectType
		varSdaaciConnectionDetail.Description = varSdaaciConnectionDetailWithoutEmbeddedStruct.Description
		varSdaaciConnectionDetail.IpPool = varSdaaciConnectionDetailWithoutEmbeddedStruct.IpPool
		varSdaaciConnectionDetail.Layer3HandoffId = varSdaaciConnectionDetailWithoutEmbeddedStruct.Layer3HandoffId
		varSdaaciConnectionDetail.PeerAinterface = varSdaaciConnectionDetailWithoutEmbeddedStruct.PeerAinterface
		varSdaaciConnectionDetail.PeerAipAddress = varSdaaciConnectionDetailWithoutEmbeddedStruct.PeerAipAddress
		varSdaaciConnectionDetail.PeerAtype = varSdaaciConnectionDetailWithoutEmbeddedStruct.PeerAtype
		varSdaaciConnectionDetail.PeerBinterface = varSdaaciConnectionDetailWithoutEmbeddedStruct.PeerBinterface
		varSdaaciConnectionDetail.PeerBipAddress = varSdaaciConnectionDetailWithoutEmbeddedStruct.PeerBipAddress
		varSdaaciConnectionDetail.PeerBtype = varSdaaciConnectionDetailWithoutEmbeddedStruct.PeerBtype
		varSdaaciConnectionDetail.Peera = varSdaaciConnectionDetailWithoutEmbeddedStruct.Peera
		varSdaaciConnectionDetail.Peerb = varSdaaciConnectionDetailWithoutEmbeddedStruct.Peerb
		varSdaaciConnectionDetail.RouterId = varSdaaciConnectionDetailWithoutEmbeddedStruct.RouterId
		varSdaaciConnectionDetail.Status = varSdaaciConnectionDetailWithoutEmbeddedStruct.Status
		varSdaaciConnectionDetail.Connection = varSdaaciConnectionDetailWithoutEmbeddedStruct.Connection
		*o = SdaaciConnectionDetail(varSdaaciConnectionDetail)
	} else {
		return err
	}

	varSdaaciConnectionDetail := _SdaaciConnectionDetail{}

	err = json.Unmarshal(data, &varSdaaciConnectionDetail)
	if err == nil {
		o.MoBaseMo = varSdaaciConnectionDetail.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "IpPool")
		delete(additionalProperties, "Layer3HandoffId")
		delete(additionalProperties, "PeerAinterface")
		delete(additionalProperties, "PeerAipAddress")
		delete(additionalProperties, "PeerAtype")
		delete(additionalProperties, "PeerBinterface")
		delete(additionalProperties, "PeerBipAddress")
		delete(additionalProperties, "PeerBtype")
		delete(additionalProperties, "Peera")
		delete(additionalProperties, "Peerb")
		delete(additionalProperties, "RouterId")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Connection")

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

type NullableSdaaciConnectionDetail struct {
	value *SdaaciConnectionDetail
	isSet bool
}

func (v NullableSdaaciConnectionDetail) Get() *SdaaciConnectionDetail {
	return v.value
}

func (v *NullableSdaaciConnectionDetail) Set(val *SdaaciConnectionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSdaaciConnectionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSdaaciConnectionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdaaciConnectionDetail(val *SdaaciConnectionDetail) *NullableSdaaciConnectionDetail {
	return &NullableSdaaciConnectionDetail{value: val, isSet: true}
}

func (v NullableSdaaciConnectionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdaaciConnectionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
