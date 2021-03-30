package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricSwitchProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricSwitchProfileRead,
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Defines the type of the profile. Accepted value is instance.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricSwitchProfile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricSwitchProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricSwitchProfile{}
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricSwitchProfile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricSwitchProfileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricSwitchProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricSwitchProfileList.GetCount()
	var i int32
	var fabricSwitchProfileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricSwitchProfileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricSwitchProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricSwitchProfileList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricSwitchProfile data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["action"] = (s.GetAction())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["assigned_switch"] = flattenMapNetworkElementRelationship(s.GetAssignedSwitch(), d)

				temp["associated_switch"] = flattenMapNetworkElementRelationship(s.GetAssociatedSwitch(), d)
				temp["class_id"] = (s.GetClassId())

				temp["config_change_details"] = flattenListFabricConfigChangeDetailRelationship(s.GetConfigChangeDetails(), d)

				temp["config_changes"] = flattenMapPolicyConfigChange(s.GetConfigChanges(), d)

				temp["config_context"] = flattenMapPolicyConfigContext(s.GetConfigContext(), d)

				temp["config_result"] = flattenMapFabricConfigResultRelationship(s.GetConfigResult(), d)
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["policy_bucket"] = flattenListPolicyAbstractPolicyRelationship(s.GetPolicyBucket(), d)

				temp["running_workflows"] = flattenListWorkflowWorkflowInfoRelationship(s.GetRunningWorkflows(), d)

				temp["src_template"] = flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)

				temp["switch_cluster_profile"] = flattenMapFabricSwitchClusterProfileRelationship(s.GetSwitchClusterProfile(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				fabricSwitchProfileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricSwitchProfileResults))
	if err := d.Set("results", fabricSwitchProfileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricSwitchProfileResults[0]["moid"].(string))
	return de
}
