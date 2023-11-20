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

// RecommendationPurchaseOrderEstimate Entity representing the estimate for the purchase order for user requested expansion.
type RecommendationPurchaseOrderEstimate struct {
	RecommendationBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be triggered for computing estimate. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
	Action *string `json:"Action,omitempty"`
	// The unique identification generated by the Cisco commerce APIs portal to identify the recommended bill of materials.
	EstimateId *string `json:"EstimateId,omitempty"`
	// The user visible message which indicates any errors encountered in making the external API call to get the estimate.
	Message *string `json:"Message,omitempty"`
	// The status of the external API call to get the estimate. * `None` - The Enum value None represents the default status value before any API call is made. * `Success` - The Enum value Success represents that the API call returned with success. * `Fail` - The Enum value Fail represents that the API call returned with a failure.
	Status *string `json:"Status,omitempty"`
	// The total cost of all the recommended hardware in the bill of materials for the corresponding estimate.
	// Deprecated
	TotalCost            *float32                                    `json:"TotalCost,omitempty"`
	ClusterExpansion     *RecommendationClusterExpansionRelationship `json:"ClusterExpansion,omitempty"`
	EstimateOwner        *IamUserRelationship                        `json:"EstimateOwner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationPurchaseOrderEstimate RecommendationPurchaseOrderEstimate

// NewRecommendationPurchaseOrderEstimate instantiates a new RecommendationPurchaseOrderEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationPurchaseOrderEstimate(classId string, objectType string) *RecommendationPurchaseOrderEstimate {
	this := RecommendationPurchaseOrderEstimate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewRecommendationPurchaseOrderEstimateWithDefaults instantiates a new RecommendationPurchaseOrderEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationPurchaseOrderEstimateWithDefaults() *RecommendationPurchaseOrderEstimate {
	this := RecommendationPurchaseOrderEstimate{}
	var classId string = "recommendation.PurchaseOrderEstimate"
	this.ClassId = classId
	var objectType string = "recommendation.PurchaseOrderEstimate"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationPurchaseOrderEstimate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationPurchaseOrderEstimate) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationPurchaseOrderEstimate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationPurchaseOrderEstimate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RecommendationPurchaseOrderEstimate) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RecommendationPurchaseOrderEstimate) SetAction(v string) {
	o.Action = &v
}

// GetEstimateId returns the EstimateId field value if set, zero value otherwise.
func (o *RecommendationPurchaseOrderEstimate) GetEstimateId() string {
	if o == nil || o.EstimateId == nil {
		var ret string
		return ret
	}
	return *o.EstimateId
}

// GetEstimateIdOk returns a tuple with the EstimateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetEstimateIdOk() (*string, bool) {
	if o == nil || o.EstimateId == nil {
		return nil, false
	}
	return o.EstimateId, true
}

// HasEstimateId returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasEstimateId() bool {
	if o != nil && o.EstimateId != nil {
		return true
	}

	return false
}

// SetEstimateId gets a reference to the given string and assigns it to the EstimateId field.
func (o *RecommendationPurchaseOrderEstimate) SetEstimateId(v string) {
	o.EstimateId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecommendationPurchaseOrderEstimate) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecommendationPurchaseOrderEstimate) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RecommendationPurchaseOrderEstimate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RecommendationPurchaseOrderEstimate) SetStatus(v string) {
	o.Status = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
// Deprecated
func (o *RecommendationPurchaseOrderEstimate) GetTotalCost() float32 {
	if o == nil || o.TotalCost == nil {
		var ret float32
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RecommendationPurchaseOrderEstimate) GetTotalCostOk() (*float32, bool) {
	if o == nil || o.TotalCost == nil {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasTotalCost() bool {
	if o != nil && o.TotalCost != nil {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given float32 and assigns it to the TotalCost field.
// Deprecated
func (o *RecommendationPurchaseOrderEstimate) SetTotalCost(v float32) {
	o.TotalCost = &v
}

// GetClusterExpansion returns the ClusterExpansion field value if set, zero value otherwise.
func (o *RecommendationPurchaseOrderEstimate) GetClusterExpansion() RecommendationClusterExpansionRelationship {
	if o == nil || o.ClusterExpansion == nil {
		var ret RecommendationClusterExpansionRelationship
		return ret
	}
	return *o.ClusterExpansion
}

// GetClusterExpansionOk returns a tuple with the ClusterExpansion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool) {
	if o == nil || o.ClusterExpansion == nil {
		return nil, false
	}
	return o.ClusterExpansion, true
}

// HasClusterExpansion returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasClusterExpansion() bool {
	if o != nil && o.ClusterExpansion != nil {
		return true
	}

	return false
}

// SetClusterExpansion gets a reference to the given RecommendationClusterExpansionRelationship and assigns it to the ClusterExpansion field.
func (o *RecommendationPurchaseOrderEstimate) SetClusterExpansion(v RecommendationClusterExpansionRelationship) {
	o.ClusterExpansion = &v
}

// GetEstimateOwner returns the EstimateOwner field value if set, zero value otherwise.
func (o *RecommendationPurchaseOrderEstimate) GetEstimateOwner() IamUserRelationship {
	if o == nil || o.EstimateOwner == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.EstimateOwner
}

// GetEstimateOwnerOk returns a tuple with the EstimateOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPurchaseOrderEstimate) GetEstimateOwnerOk() (*IamUserRelationship, bool) {
	if o == nil || o.EstimateOwner == nil {
		return nil, false
	}
	return o.EstimateOwner, true
}

// HasEstimateOwner returns a boolean if a field has been set.
func (o *RecommendationPurchaseOrderEstimate) HasEstimateOwner() bool {
	if o != nil && o.EstimateOwner != nil {
		return true
	}

	return false
}

// SetEstimateOwner gets a reference to the given IamUserRelationship and assigns it to the EstimateOwner field.
func (o *RecommendationPurchaseOrderEstimate) SetEstimateOwner(v IamUserRelationship) {
	o.EstimateOwner = &v
}

func (o RecommendationPurchaseOrderEstimate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecommendationBase, errRecommendationBase := json.Marshal(o.RecommendationBase)
	if errRecommendationBase != nil {
		return []byte{}, errRecommendationBase
	}
	errRecommendationBase = json.Unmarshal([]byte(serializedRecommendationBase), &toSerialize)
	if errRecommendationBase != nil {
		return []byte{}, errRecommendationBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.EstimateId != nil {
		toSerialize["EstimateId"] = o.EstimateId
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TotalCost != nil {
		toSerialize["TotalCost"] = o.TotalCost
	}
	if o.ClusterExpansion != nil {
		toSerialize["ClusterExpansion"] = o.ClusterExpansion
	}
	if o.EstimateOwner != nil {
		toSerialize["EstimateOwner"] = o.EstimateOwner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationPurchaseOrderEstimate) UnmarshalJSON(bytes []byte) (err error) {
	type RecommendationPurchaseOrderEstimateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action to be triggered for computing estimate. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
		Action *string `json:"Action,omitempty"`
		// The unique identification generated by the Cisco commerce APIs portal to identify the recommended bill of materials.
		EstimateId *string `json:"EstimateId,omitempty"`
		// The user visible message which indicates any errors encountered in making the external API call to get the estimate.
		Message *string `json:"Message,omitempty"`
		// The status of the external API call to get the estimate. * `None` - The Enum value None represents the default status value before any API call is made. * `Success` - The Enum value Success represents that the API call returned with success. * `Fail` - The Enum value Fail represents that the API call returned with a failure.
		Status *string `json:"Status,omitempty"`
		// The total cost of all the recommended hardware in the bill of materials for the corresponding estimate.
		// Deprecated
		TotalCost        *float32                                    `json:"TotalCost,omitempty"`
		ClusterExpansion *RecommendationClusterExpansionRelationship `json:"ClusterExpansion,omitempty"`
		EstimateOwner    *IamUserRelationship                        `json:"EstimateOwner,omitempty"`
	}

	varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct := RecommendationPurchaseOrderEstimateWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationPurchaseOrderEstimate := _RecommendationPurchaseOrderEstimate{}
		varRecommendationPurchaseOrderEstimate.ClassId = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.ClassId
		varRecommendationPurchaseOrderEstimate.ObjectType = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.ObjectType
		varRecommendationPurchaseOrderEstimate.Action = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.Action
		varRecommendationPurchaseOrderEstimate.EstimateId = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.EstimateId
		varRecommendationPurchaseOrderEstimate.Message = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.Message
		varRecommendationPurchaseOrderEstimate.Status = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.Status
		varRecommendationPurchaseOrderEstimate.TotalCost = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.TotalCost
		varRecommendationPurchaseOrderEstimate.ClusterExpansion = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.ClusterExpansion
		varRecommendationPurchaseOrderEstimate.EstimateOwner = varRecommendationPurchaseOrderEstimateWithoutEmbeddedStruct.EstimateOwner
		*o = RecommendationPurchaseOrderEstimate(varRecommendationPurchaseOrderEstimate)
	} else {
		return err
	}

	varRecommendationPurchaseOrderEstimate := _RecommendationPurchaseOrderEstimate{}

	err = json.Unmarshal(bytes, &varRecommendationPurchaseOrderEstimate)
	if err == nil {
		o.RecommendationBase = varRecommendationPurchaseOrderEstimate.RecommendationBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "EstimateId")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TotalCost")
		delete(additionalProperties, "ClusterExpansion")
		delete(additionalProperties, "EstimateOwner")

		// remove fields from embedded structs
		reflectRecommendationBase := reflect.ValueOf(o.RecommendationBase)
		for i := 0; i < reflectRecommendationBase.Type().NumField(); i++ {
			t := reflectRecommendationBase.Type().Field(i)

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

type NullableRecommendationPurchaseOrderEstimate struct {
	value *RecommendationPurchaseOrderEstimate
	isSet bool
}

func (v NullableRecommendationPurchaseOrderEstimate) Get() *RecommendationPurchaseOrderEstimate {
	return v.value
}

func (v *NullableRecommendationPurchaseOrderEstimate) Set(val *RecommendationPurchaseOrderEstimate) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationPurchaseOrderEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationPurchaseOrderEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationPurchaseOrderEstimate(val *RecommendationPurchaseOrderEstimate) *NullableRecommendationPurchaseOrderEstimate {
	return &NullableRecommendationPurchaseOrderEstimate{value: val, isSet: true}
}

func (v NullableRecommendationPurchaseOrderEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationPurchaseOrderEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
